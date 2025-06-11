import json
import os
import time
from importlib import import_module
from typing import Optional

import requests
from eth_abi import decode, encode
from eth_account import Account
from eth_account.datastructures import SignedMessage, SignedTransaction
from eth_account.messages import encode_defunct
from tqdm import tqdm
from web3 import Web3

from avs_operator.utils import parse_input
from common.console import console, styles
from common.constants import ETH_STRATEGY_ADDRESSES, PROOFS_FOLDER
from common.eth import EthereumClient, load_ecdsa_private_key
from models.proof.ezkl_handler import EZKLHandler
from models.onnx_run import run_onnx

task_schema = [
    # ISertnServiceManagerTypes.Task structure
    "uint256",  # operatorModelId_
    "bytes",  # inputs_
    "uint256",  # poc_
    "uint256",  # startTime_
    "uint32",  # startingBlock_
    "bool",  # proveOnResponse_
    "address",  # operator_
]


def run_operator(config: dict) -> None:
    console.print("Starting Sertn Operator...", style=styles.op_info)

    with tqdm(total=100, desc="Initializing operator") as pbar:
        task_operator = TaskOperator(config)
        pbar.update(25)
        # task_operator.register()
        pbar.update(100)
    task_operator.start()


def relative_file_path(file_path: str):
    return os.path.join(os.path.dirname(__file__), file_path)


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

        # try:
        #     # Simulate the transaction
        #     self.eth_client.w3.eth.call(tx)
        # except Exception as e:
        #     print(f"Transaction simulation failed: {e}")

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
        # console.print("Checking Operator registration...", style=styles.op_info)
        # self.register()
        console.print("Starting Operator...", style=styles.op_info)
        event_filter = self.eth_client.task_manager.events.TaskCreated.create_filter(
            from_block="latest"
        )
        while True:
            for event in event_filter.get_new_entries():
                console.print(
                    f"New task event received. Processing...", style=styles.op_info
                )
                self.process_task_event(event)
            time.sleep(3)

    def process_task_event(self, event):
        task_id_bytes: bytes = event["args"]["taskId_"]
        task_id_hex: str = task_id_bytes.hex()
        operator_address: bytes = event["args"]["opAddr_"]

        task = self.eth_client.task_manager.functions.tasks(task_id_bytes).call()

        operator_model_id: str = task[0]
        inputs: bytes = task[1]
        # XXX: WTF is poc? Proof of computation value? What to do with it?
        poc: int = task[2]
        start_time: int = task[3]
        starting_block: int = task[4]
        prove_on_response: bool = task[5]

        if operator_address != self.operator_address:
            console.print(
                f"Operator address {operator_address} does not match "
                f"the operator address {self.operator_address}, skipping...",
                style=styles.debug,
            )
            return

        input_data = [float(i) for i in inputs.decode().split(" ")]
        output = run_onnx(
            model_id=operator_model_id,
            input_data=input_data,
        )

        proof_generator = EZKLHandler(
            model_id=operator_model_id,
            task_id=task_id_hex,
            inputs=input_data,
        )
        # just save input data to the file
        proof_generator.gen_input_file()

        prove: Optional[bytes]
        if prove_on_response:
            # generate proof
            prove = json.dumps(proof_generator.gen_proof()).encode()

        output_bytes: bytes = json.dumps(output).encode()

        encoded = encode(
            ["uint32", "bytes", "uint256", "address"],
            [starting_block, output_bytes, operator_model_id, self.operator_address],
        )
        message = encode_defunct(primitive=encoded)
        signed_message: SignedMessage = Account.sign_message(
            message, private_key=self.private_key
        )
        signature = signed_message.signature.hex()

        console.print(
            f"Signature generated, block: {starting_block}, model: {operator_model_id}, "
            f"Signature generated, block: {starting_block}, model: {operator_model_id}, "
            f"output length: {len(output)} signature: {signature}",
            f"output length: {len(output)} signature: {signature}",
            style=styles.debug,
        )
        # logger.info("operator data id", self.operator_id.hex())
        data = {
            "task_id": task_id_bytes.hex(),
            "operator_model_id": operator_model_id,
            "inputs": inputs.hex(),
            "output": output_bytes.hex(),
            "prove": prove.hex(),
            "start_time": start_time,
            "starting_block": starting_block,
            "signature": signature,
            "poc": poc,
            "prove_on_response": prove_on_response,
            "operator_address": self.operator_address,
        }
        console.print(
            f"Submitting result for task to aggregator task ID - {task_id_hex}",
            style=styles.debug,
        )
        url = f"http://{self.config['aggregator_server_ip_port_address']}/signature"
        requests.post(url, json=data)
