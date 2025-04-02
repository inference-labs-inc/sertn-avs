import json
import logging
import os
import time
from random import randbytes
from typing import Optional

import eth_abi
import requests
import yaml
from eth_abi import encode
from eth_account import Account
from web3 import Web3

from eigensdk._types import Operator
from eigensdk.chainio.clients.builder import BuildAllConfig, build_all
from eigensdk.crypto.bls.attestation import KeyPair

from constants import (
    ROOT_DIR,
    OPERATOR_KEY,
    AGGREGATOR_SERVER_HOST,
    BLS_PRIVATE_KEY_STORE_PATH,
    OPERATOR_BLS_KEY_PASSWORD,
    RPC_SERVER_URL,
    TASK_MANAGER_ABI,
    TASK_MANAGER_ADDRESS,
    SERVICE_MANAGER_ABI,
    SERVICE_MANAGER_ADDRESS,
    ETH_STRATEGY_ADDRESSES,
)

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

if not OPERATOR_BLS_KEY_PASSWORD:
    logger.warning("OPERATOR_BLS_KEY_PASSWORD not set. using empty string.")


class TaskOperator:
    def __init__(self):
        self.bls_key_pair = KeyPair.read_from_file(
            BLS_PRIVATE_KEY_STORE_PATH, OPERATOR_BLS_KEY_PASSWORD
        )
        # self.__load_clients()
        self.web3 = Web3(Web3.HTTPProvider(RPC_SERVER_URL))

        self.service_manager = self.web3.eth.contract(
            address=SERVICE_MANAGER_ADDRESS, abi=SERVICE_MANAGER_ABI
        )
        self.task_manager = self.web3.eth.contract(
            address=TASK_MANAGER_ADDRESS, abi=TASK_MANAGER_ABI
        )

    def register(self):
        """
        Register operator with the AVS registry
        TODO: verify isn't the operator already registered
        """
        operator_address = Account.from_key(OPERATOR_KEY).address

        # `data` argument for `registerOperator` function is just `bytes` type
        # and is encoded as a tuple of the following types:
        data_schema = [
            # Model[] memory _models
            "tuple(string title_, string description_, address modelVerifier_, address[] operators_)[]",
            # OperatorModel[] memory _operatorModels
            (
                "tuple(address operator_, uint256 modelId_, uint256 maxBlocks_, "
                "address[] ethStrategies_, uint256[] ethShares_, uint256 baseFee_, uint256 maxSer_, "
                "bytes32 computeType_, bool proveOnResponse_, bool available_)[]"
            ),
            # bytes32[] memory _computeUnitNames
            "bytes32[]",
            # uint256[] memory _computeUnits
            "uint256[]",
        ]

        # Define the structs
        # `Model[] memory _models`:
        models = [
            {
                "title_": "Example Model Title",
                "description_": "Example Model Description",
                "modelVerifier_": "0x0000000000000000000000000000000000000000",  # Example address
                "operators_": [
                    "0x0000000000000000000000000000000000000000"
                ],  # Example addresses
            }
        ]

        # OperatorModel[] memory _operatorModels:
        operator_models = [
            {
                "operator_": "0xabcdef1234567890abcdef1234567890abcdef12",  # Example operator address
                "modelId_": 1,
                "maxBlocks_": 1000,
                "ethShares_": [100, 200],
                "ethStrategies_": [
                    ETH_STRATEGY_ADDRESSES[0],
                ],  # IStrategy[]
                "baseFee_": 0,
                "maxSer_": 1000,
                "computeType_": "0xabcdefabcdefabcdefabcdefabcdefabcdef",  # dummy bytes32 value
                "proveOnResponse_": True,
                "available_": True,
            }
        ]

        # bytes32[] memory _computeUnitNames
        compute_unit_names = [
            Web3.toBytes(hexstr="0x636f6d707574655f756e6974")  # Example bytes32 value
        ]

        # uint256[] memory _computeUnits
        compute_units = [1000]  # Example compute unit values

        # so we need to encode the data as a tuple of the above types:
        data = encode(
            data_schema,
            [models, operator_models, compute_unit_names, compute_units],
        )

        # Define other the function arguments:
        # required, but not used by the contract
        avs = "0x000000000"
        # Example operator set IDs, according to the contract, must be `[0]`
        operator_set_ids = [0]

        self.service_manager.functions.registerOperator(
            operator_address, avs, operator_set_ids, data
        ).build_transaction(
            {
                "from": operator_address,
                "gas": 3000000,
                "gasPrice": self.web3.toWei("20", "gwei"),
                "nonce": self.web3.eth.get_transaction_count("0xYourAccountAddress"),
            }
        )

    def start(self):
        logger.info("Checking Operator registration...")
        self.register()
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

        model_id: int = event["args"]["task"]["modelId_"]
        inputs: bytes = event["args"]["task"]["inputs_"]
        # XXX: WTF is poc? Proof of computation value? What to do with it?
        poc: int = event["args"]["task"]["poc_"]
        proveOnResponse: bool = event["args"]["task"]["proveOnResponse_"]

        ##################################
        # TODO: Actual work here........
        output: bytes = "".encode()
        prove: Optional[bytes] = None
        if proveOnResponse:
            prove = "".encode()
        ##################################

        encoded = eth_abi.encode(
            ["uint32", "bytes", "uint256"], [task_id, output, model_id]
        )
        hash_bytes = Web3.keccak(encoded)
        signature = self.bls_key_pair.sign_message(msg_bytes=hash_bytes).to_json()
        logger.info(
            f"Signature generated, task id: {task_id}, model: {model_id}, "
            f"output length: {len(output)} signature: {signature}"
        )
        # logger.info("operator data id", self.operator_id.hex())
        data = {
            "task_id": task_id,
            "inputs": inputs,
            "output": output,
            "prove": prove,
            "model_id": model_id,
            "signature": signature,
            "poc": poc,
            "proveOnResponse": proveOnResponse,
            "block_number": event["blockNumber"],
            # "operator_id": "0x" + self.operator_id.hex(),
        }
        logger.info(f"Submitting result for task to aggregator {data}")
        url = f"http://{AGGREGATOR_SERVER_HOST}/signature"
        requests.post(url, json=data)

    # def __load_clients(self):
    #     cfg = BuildAllConfig(
    #         eth_http_url=RPC_SERVER_URL,
    #         registry_coordinator_addr=self.config["avs_registry_coordinator_address"],
    #         operator_state_retriever_addr=self.config[
    #             "operator_state_retriever_address"
    #         ],
    #         avs_name="incredible-squaring",
    #         prom_metrics_ip_port_address=self.config["eigen_metrics_ip_port_address"],
    #         # eth_ws_url=self.config["eth_ws_url"],
    #     )
    #     self.clients = build_all(cfg, OPERATOR_KEY, logger)


if __name__ == "__main__":
    # with open("configs/operator.anvil.yaml", "r") as f:
    #     config = yaml.load(f, Loader=yaml.BaseLoader)

    TaskOperator().start()
