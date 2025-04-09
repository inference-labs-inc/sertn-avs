import os
import time
from typing import Optional

import requests
from web3 import Web3
from tqdm import tqdm
from eth_account import Account
from eth_account.datastructures import SignedTransaction
from eth_abi import encode

from common.console import console, styles
from common.eth import EthereumClient, load_ecdsa_private_key
from common.constants import (
    ETH_STRATEGY_ADDRESSES,
)


def run_operator(config: dict) -> None:
    console.print("Starting Sertn Operator...", style=styles.op_info)

    with tqdm(total=100, desc="Initializing operator") as pbar:
        task_operator = TaskOperator(config)
        pbar.update(25)
        task_operator.register()
        pbar.update(100)
    task_operator.start()


class TaskOperator:
    def __init__(self, config: dict):
        self.config = config
        self.eth_client = EthereumClient(rpc_url=self.config["eth_rpc_url"])
        self.private_key = load_ecdsa_private_key(
            keystore_path=self.config["ecdsa_private_key_store_path"],
            password=os.environ.get("OPERATOR_ECDSA_KEY_PASSWORD", ""),
        )
        self.operator_address = Account.from_key(self.private_key).address

    def register(self):
        """
        Register operator with the AVS registry
        TODO: verify isn't the operator already registered
        """
        operator_address = Account.from_key(self.private_key).address

        # `data` argument for `registerOperator` function is just `bytes` type
        # and is encoded as a tuple of the following types:
        data_schema = [
            # Model[] memory _models
            "(string,string,address,address[])[]",
            # OperatorModel[] memory _operatorModels
            "(address,uint256,uint256,address[],uint256[],uint256,uint256,bytes32,bool,bool)[]",
            # bytes32[] memory _computeUnitNames
            "bytes32[]",
            # uint256[] memory _computeUnits
            "uint256[]",
        ]

        # Define the structs
        # `Model[] memory _models`:
        models = [
            (
                "Example Model Title",  # title_
                "Example Model Description",  # description_
                "0x0000000000000000000000000000000000000000",  # modelVerifier_ - Example address
                [
                    "0x0000000000000000000000000000000000000000"
                ],  # operators_ - Example addresses
            )
        ]

        # OperatorModel[] memory _operatorModels:
        operator_models = [
            (
                operator_address,  # operator_ - value will be set in the contract
                1,  # modelId_
                100,  # maxBlocks_
                [ETH_STRATEGY_ADDRESSES[0]],  # ethStrategies_ - IStrategy[]
                [100, 200],  # ethShares_
                0,  # baseFee_
                1000,  # maxSer_
                "0".encode(),  # computeType_ - dummy bytes32 value
                True,  # proveOnResponse_
                True,  # available_
            )
        ]

        # bytes32[] memory _computeUnitNames
        compute_unit_names = [
            Web3.to_bytes(hexstr="0x636f6d707574655f756e6974")  # Example bytes32 value
        ]

        # uint256[] memory _computeUnits
        compute_units = [1000]  # Example compute unit values

        # so we need to encode the data as a tuple of the above types:
        data = encode(
            data_schema,
            [
                models,
                operator_models,
                compute_unit_names,
                compute_units,
            ],
        )

        # Define other the function arguments:
        # required, but not used by the contract
        avs = "0x0000000000000000000000000000000000000000"  # Example AVS address
        # Example operator set IDs, according to the contract, must be `[0]`
        operator_set_ids = [0]

        tx = self.eth_client.service_manager.functions.registerOperator(
            operator_address, avs, operator_set_ids, data
        ).build_transaction(
            {
                "from": operator_address,
                "gas": 2000000,
                "gasPrice": self.eth_client.w3.to_wei("20", "gwei"),
                "nonce": self.eth_client.w3.eth.get_transaction_count(operator_address),
                "chainId": self.eth_client.w3.eth.chain_id,
            }
        )

        try:
            # Simulate the transaction
            self.eth_client.w3.eth.call(tx)
        except Exception as e:
            print(f"Transaction simulation failed: {e}")

        signed_tx: SignedTransaction = self.eth_client.w3.eth.account.sign_transaction(
            tx, private_key=self.private_key
        )
        tx_hash = self.eth_client.w3.eth.send_raw_transaction(signed_tx.raw_transaction)
        receipt = self.eth_client.w3.eth.wait_for_transaction_receipt(tx_hash)

        if not receipt["logs"]:
            console.print(
                "Failed to register operator, no logs found in transaction receipt.",
                style=styles.error,
            )
            exit(1)

        console.print(
            f"Operator registered successfully, tx hash: {tx_hash.hex()}",
            style=styles.op_info,
        )

    def start(self):
        console.print("Checking Operator registration...", style=styles.op_info)
        self.register()
        console.print("Starting Operator...", style=styles.op_info)
        event_filter = self.eth_client.task_manager.events.NewTask.create_filter(
            from_block="latest"
        )
        while True:
            for event in event_filter.get_new_entries():
                console.print(f"New task created: {event}", style=styles.op_info)
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

        encoded = encode(["uint32", "bytes", "uint256"], [task_id, output, model_id])
        hash_bytes = Web3.keccak(encoded)
        signature = self.bls_key_pair.sign_message(msg_bytes=hash_bytes).to_json()
        console.print(
            f"Signature generated, task id: {task_id}, model: {model_id}, "
            f"output length: {len(output)} signature: {signature}",
            style=styles.debug,
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
        console.print(
            f"Submitting result for task to aggregator {data}", style=styles.debug
        )
        url = f"http://{self.config['aggregator_server_ip_port_address']}/signature"
        requests.post(url, json=data)
