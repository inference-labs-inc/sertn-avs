import os
import time
import json
import logging
from random import randbytes
import requests
import yaml
from web3 import Web3
import eth_abi
from eth_account import Account
from eigensdk.chainio.clients.builder import BuildAllConfig, build_all
from eigensdk.crypto.bls.attestation import KeyPair
from eigensdk._types import Operator

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


class SquaringOperator:
    def __init__(self, config):
        self.config = config
        self.__load_bls_key()
        self.__load_ecdsa_key()
        self.__load_clients()
        self.__load_task_manager()
        if config["register_operator_on_startup"] == "true":
            self.register()
        # operator id can only be loaded after registration
        self.__load_operator_id()

    def register(self):
        operator = Operator(
            address=self.config["operator_address"],
            earnings_receiver_address=self.config["operator_address"],
            delegation_approver_address="0x0000000000000000000000000000000000000000",
            staker_opt_out_window_blocks=0,
            metadata_url="",
        )
        self.clients.el_writer.register_as_operator(operator)
        self.clients.avs_registry_writer.register_operator_in_quorum_with_avs_registry_coordinator(
            operator_ecdsa_private_key=self.operator_ecdsa_private_key,
            operator_to_avs_registration_sig_salt=randbytes(32),
            operator_to_avs_registration_sig_expiry=int(time.time()) + 3600,
            bls_key_pair=self.bls_key_pair,
            quorum_numbers=[0],
            socket="Not Needed",
        )

    def start(self):
        logger.info("Starting Operator...")
        event_filter = self.task_manager.events.NewTaskCreated.create_filter(
            from_block="latest"
        )
        while True:
            for event in event_filter.get_new_entries():
                logger.info(f"New task created: {event}")
                self.process_task_event(event)
            time.sleep(3)

    def process_task_event(self, event):
        task_id = event["args"]["taskIndex"]
        number_to_be_squared = event["args"]["task"]["numberToBeSquared"]
        number_squared = number_to_be_squared**2
        encoded = eth_abi.encode(["uint32", "uint256"], [task_id, number_squared])
        hash_bytes = Web3.keccak(encoded)
        signature = self.bls_key_pair.sign_message(msg_bytes=hash_bytes).to_json()
        logger.info(
            f"Signature generated, task id: {task_id}, number squared: {number_squared}, signature: {signature}"
        )
        print("operator data id", self.operator_id.hex())
        data = {
            "task_id": task_id,
            "number_to_be_squared": number_to_be_squared,
            "number_squared": number_squared,
            "signature": signature,
            "block_number": event["blockNumber"],
            "operator_id": "0x" + self.operator_id.hex(),
        }
        logger.info(f"Submitting result for task to aggregator {data}")
        # prevent submitting task before initialize_new_task gets completed on aggregator
        time.sleep(3)
        url = f'http://{self.config["aggregator_server_ip_port_address"]}/signature'
        requests.post(url, json=data)

    def __load_bls_key(self):
        bls_key_password = os.environ.get("OPERATOR_BLS_KEY_PASSWORD", "")
        if not bls_key_password:
            logger.warning("OPERATOR_BLS_KEY_PASSWORD not set. using empty string.")

        self.bls_key_pair = KeyPair.read_from_file(
            self.config["bls_private_key_store_path"], bls_key_password
        )

    def __load_ecdsa_key(self):
        ecdsa_key_password = os.environ.get("OPERATOR_ECDSA_KEY_PASSWORD", "")
        if not ecdsa_key_password:
            logger.warning("OPERATOR_ECDSA_KEY_PASSWORD not set. using empty string.")

        with open(self.config["ecdsa_private_key_store_path"], "r") as f:
            keystore = json.load(f)
        self.operator_ecdsa_private_key = Account.decrypt(
            keystore, ecdsa_key_password
        ).hex()

    def __load_clients(self):
        cfg = BuildAllConfig(
            eth_http_url=self.config["eth_rpc_url"],
            eth_ws_url=self.config["eth_ws_url"],
            avs_name="incredible-squaring",
            registry_coordinator_addr=self.config["avs_registry_coordinator_address"],
            operator_state_retriever_addr=self.config[
                "operator_state_retriever_address"
            ],
            prom_metrics_ip_port_address=self.config["eigen_metrics_ip_port_address"],
        )
        self.clients = build_all(cfg, self.operator_ecdsa_private_key, logger)

    def __load_task_manager(self):
        web3 = Web3(Web3.HTTPProvider(self.config["eth_rpc_url"]))

        service_manager_address = self.clients.avs_registry_writer.service_manager_addr
        with open("abis/IncredibleSquaringServiceManager.json") as f:
            service_manager_abi = f.read()
        service_manager = web3.eth.contract(
            address=service_manager_address, abi=service_manager_abi
        )

        task_manager_address = (
            service_manager.functions.incredibleSquaringTaskManager().call()
        )
        with open("abis/IncredibleSquaringTaskManager.json") as f:
            task_manager_abi = f.read()
        self.task_manager = web3.eth.contract(
            address=task_manager_address, abi=task_manager_abi
        )

    def __load_operator_id(self):
        self.operator_id = self.clients.avs_registry_reader.get_operator_id(
            self.config["operator_address"]
        )


if __name__ == "__main__":
    with open("config-files/operator.anvil.yaml", "r") as f:
        config = yaml.load(f, Loader=yaml.BaseLoader)

    SquaringOperator(config=config).start()
