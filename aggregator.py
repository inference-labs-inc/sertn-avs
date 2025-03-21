import os
import logging
import json
import time
import threading
import asyncio
from dataclasses import dataclass
from typing import List
import yaml
from web3 import Web3
import eth_abi
from eth_account import Account
from flask import Flask, request
from eigensdk.chainio.clients.builder import BuildAllConfig, build_all
from eigensdk.services.avsregistry import AvsRegistryService
from eigensdk.services.operatorsinfo.operatorsinfo_inmemory import (
    OperatorsInfoServiceInMemory,
)
from eigensdk.services.bls_aggregation.blsagg import (
    BlsAggregationService,
    BlsAggregationServiceResponse,
)
from eigensdk.chainio.utils import nums_to_bytes
from eigensdk.crypto.bls.attestation import (
    Signature,
    G1Point,
    G2Point,
    g1_to_tupple,
    g2_to_tupple,
)

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


class Aggregator:
    def __init__(self, config):
        self.config = config
        self.web3 = Web3(Web3.HTTPProvider(self.config["eth_rpc_url"]))
        self.__load_ecdsa_key()
        self.__load_clients()
        self.__load_task_manager()
        self.__load_bls_aggregation_service()
        self.tasks = {}
        self.taskResponses = {}
        self.app = Flask(__name__)
        self.app.add_url_rule(
            "/signature", "signature", self.submit_signature, methods=["POST"]
        )

    def submit_signature(self):
        data = request.get_json()
        signature = Signature(data["signature"]["X"], data["signature"]["Y"])
        task_index = data["task_id"]
        task_response = {
            "task_index": task_index,
            "number_squared": data["number_squared"],
            "number_to_be_squared": data["number_to_be_squared"],
            "block_number": data["block_number"],
        }
        print("data", data["operator_id"])
        try:
            self.bls_aggregation_service.process_new_signature(
                task_index, task_response, signature, data["operator_id"]
            )
            return "true", 200
        except Exception as e:
            logger.error(f"Submitting signature failed: {e}")
            return "false", 500

    def start_server(self):
        host, port = self.config["aggregator_server_ip_port_address"].split(":")
        self.app.run(host=host, port=port)

    def send_new_task(self, i):
        tx = self.task_manager.functions.createNewTask(
            i, 100, nums_to_bytes([0])
        ).build_transaction(
            {
                "from": self.aggregator_address,
                "gas": 2000000,
                "gasPrice": self.web3.to_wei("20", "gwei"),
                "nonce": self.web3.eth.get_transaction_count(self.aggregator_address),
                "chainId": self.web3.eth.chain_id,
            }
        )
        signed_tx = self.web3.eth.account.sign_transaction(
            tx, private_key=self.aggregator_ecdsa_private_key
        )
        tx_hash = self.web3.eth.send_raw_transaction(signed_tx.rawTransaction)
        receipt = self.web3.eth.wait_for_transaction_receipt(tx_hash)
        event = self.task_manager.events.NewTaskCreated().process_log(
            receipt["logs"][0]
        )

        task_index = event["args"]["taskIndex"]
        logger.info(f"Successfully sent the new task {task_index}")
        self.bls_aggregation_service.initialize_new_task(
            task_index=task_index,
            task_created_block=receipt["blockNumber"],
            quorum_numbers=nums_to_bytes([0]),
            quorum_threshold_percentages=[100],
            time_to_expiry=60000,
        )
        return event["args"]["taskIndex"]

    def start_sending_new_tasks(self):
        i = 0
        while True:
            logger.info("Sending new task")
            task_index = self.send_new_task(i)
            time.sleep(10)
            i += 1

    def start_submitting_signatures(self):
        while True:
            logger.info("Waiting for response")
            aggregated_response = next(
                self.bls_aggregation_service.get_aggregated_responses()
            )
            logger.info(f"Aggregated response {aggregated_response}")
            response = aggregated_response.task_response
            task = [
                response["number_to_be_squared"],
                response["block_number"],
                nums_to_bytes([0]),
                100,
            ]
            task_response = [response["task_index"], response["number_squared"]]
            non_signers_stakes_and_signature = [
                aggregated_response.non_signer_quorum_bitmap_indices,
                [g1_to_tupple(g1) for g1 in aggregated_response.non_signers_pubkeys_g1],
                [g1_to_tupple(g1) for g1 in aggregated_response.quorum_apks_g1],
                g2_to_tupple(aggregated_response.signers_apk_g2),
                g1_to_tupple(aggregated_response.signers_agg_sig_g1),
                aggregated_response.quorum_apk_indices,
                aggregated_response.total_stake_indices,
                aggregated_response.non_signer_stake_indices,
            ]

            tx = self.task_manager.functions.respondToTask(
                task, task_response, non_signers_stakes_and_signature
            ).build_transaction(
                {
                    "from": self.aggregator_address,
                    "gas": 2000000,
                    "gasPrice": self.web3.to_wei("20", "gwei"),
                    "nonce": self.web3.eth.get_transaction_count(
                        self.aggregator_address
                    ),
                    "chainId": self.web3.eth.chain_id,
                }
            )
            signed_tx = self.web3.eth.account.sign_transaction(
                tx, private_key=self.aggregator_ecdsa_private_key
            )
            tx_hash = self.web3.eth.send_raw_transaction(signed_tx.rawTransaction)
            receipt = self.web3.eth.wait_for_transaction_receipt(tx_hash)

    def __load_ecdsa_key(self):
        ecdsa_key_password = os.environ.get("AGGREGATOR_ECDSA_KEY_PASSWORD", "")
        if not ecdsa_key_password:
            logger.warning("AGGREGATOR_ECDSA_KEY_PASSWORD not set. using empty string.")

        with open(self.config["ecdsa_private_key_store_path"], "r") as f:
            keystore = json.load(f)
        self.aggregator_ecdsa_private_key = Account.decrypt(
            keystore, ecdsa_key_password
        ).hex()
        self.aggregator_address = Account.from_key(
            self.aggregator_ecdsa_private_key
        ).address

    def __load_clients(self):
        cfg = BuildAllConfig(
            eth_http_url=self.config["eth_rpc_url"],
            avs_name="incredible-squaring",
            registry_coordinator_addr=self.config["avs_registry_coordinator_address"],
            operator_state_retriever_addr=self.config[
                "operator_state_retriever_address"
            ],
            prom_metrics_ip_port_address="",
        )
        self.clients = build_all(cfg, self.aggregator_ecdsa_private_key, logger)

    def __load_task_manager(self):
        service_manager_address = self.clients.avs_registry_writer.service_manager_addr
        with open("abis/IncredibleSquaringServiceManager.json") as f:
            service_manager_abi = f.read()
        service_manager = self.web3.eth.contract(
            address=service_manager_address, abi=service_manager_abi
        )

        task_manager_address = (
            service_manager.functions.incredibleSquaringTaskManager().call()
        )
        with open("abis/IncredibleSquaringTaskManager.json") as f:
            task_manager_abi = f.read()
        self.task_manager = self.web3.eth.contract(
            address=task_manager_address, abi=task_manager_abi
        )

    def __load_bls_aggregation_service(self):
        operator_info_service = OperatorsInfoServiceInMemory(
            avs_registry_reader=self.clients.avs_registry_reader,
            start_block_pub=0,
            start_block_socket=0,
            logger=logger,
        )

        avs_registry_service = AvsRegistryService(
            self.clients.avs_registry_reader, operator_info_service, logger
        )

        def hasher(task):
            encoded = eth_abi.encode(
                ["uint32", "uint256"], [task["task_index"], task["number_squared"]]
            )
            return Web3.keccak(encoded)

        self.bls_aggregation_service = BlsAggregationService(
            avs_registry_service, hasher
        )


if __name__ == "__main__":
    with open("config-files/aggregator.yaml", "r") as f:
        config = yaml.load(f, Loader=yaml.BaseLoader)
    aggregator = Aggregator(config)
    threading.Thread(target=aggregator.start_submitting_signatures, args=[]).start()
    threading.Thread(target=aggregator.start_sending_new_tasks, args=[]).start()
    aggregator.start_server()
