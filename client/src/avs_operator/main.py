import json
import os
import time
from typing import Optional

import eth_abi
import requests
from eth_account import Account
from eth_account.datastructures import SignedMessage, SignedTransaction
from eth_account.messages import encode_defunct
from tqdm import tqdm
from web3 import Web3

from avs_operator.nodes import OperatorNodesManager
from common.console import console, styles
from common.contract_constants import TaskStructMap
from common.eth import EthereumClient, load_ecdsa_private_key
from models.onnx_run import run_onnx
from models.proof.ezkl_handler import EZKLHandler


def run_operator(config: dict) -> None:
    console.print("Starting Sertn Operator...", style=styles.op_info)

    with tqdm(total=100, desc="Initializing operator") as pbar:
        task_operator = TaskOperator(config)
        pbar.update(25)
        task_operator.nodes_manager.sync_nodes()
        task_operator.nodes_manager.print_nodes()
        # task_operator.register()
        pbar.update(100)
    task_operator.listen_for_events()


class TaskOperator:
    def __init__(self, config: dict):
        self.config = config
        self.eth_client = EthereumClient(rpc_url=self.config["eth_rpc_url"])
        self.private_key = load_ecdsa_private_key(
            keystore_path=self.config["ecdsa_private_key_store_path"],
            password=os.environ.get("OPERATOR_ECDSA_KEY_PASSWORD", ""),
        )
        self.operator_address = Account.from_key(self.private_key).address

        self.nodes_manager = OperatorNodesManager(
            private_key=self.private_key,
            operator_address=self.operator_address,
            eth_client=self.eth_client,
            nodes_config=self.config["nodes"],
        )

    def listen_for_events(self, loop_running: bool = True) -> int | None:
        console.print("Starting Operator...", style=styles.op_info)
        task_assigned_events = (
            self.eth_client.task_manager.events.TaskAssigned.create_filter(
                from_block="latest"
            )
        )
        task_challenged_events = (
            self.eth_client.task_manager.events.TaskChallenged.create_filter(
                from_block="latest"
            )
        )
        processed_count: int = 0
        while True:
            for event in task_assigned_events.get_new_entries():
                console.print(
                    f"New assignedTask received. Processing...", style=styles.op_info
                )
                self.process_assigned_task(
                    event["args"]["taskId"], event["args"]["operator"]
                )
                processed_count += 1
            for event in task_challenged_events.get_new_entries():
                console.print(
                    f"Challenged received. Processing...", style=styles.op_info
                )
                self.process_challenged_task(event["args"]["taskId"])
                processed_count += 1
            if not loop_running:
                console.print("Stopping Operator...", style=styles.op_info)
                return processed_count
            time.sleep(3)

    def process_assigned_task(self, task_id: int, operator_address: bytes):
        """Process assigned task by task ID and operator address."""

        if operator_address != self.operator_address:
            console.print(
                f"Operator address {operator_address} does not match "
                f"the operator address {self.operator_address}, skipping...",
                style=styles.debug,
            )
            return

        # get task details, `struct ISertnTaskManager.Task`
        task = self.eth_client.task_manager.functions.tasks(task_id).call()
        model_id: int = task[TaskStructMap.MODEL_ID]  # uint256 modelId
        inputs: bytes = task[TaskStructMap.INPUTS]  # bytes inputs
        operator_address: str = task[TaskStructMap.OPERATOR]  # address operator

        model_uri: str = self.eth_client.model_registry.functions.modelURI(
            model_id
        ).call()
        input_data = [float(i) for i in inputs.decode().split(" ") if i]
        output = run_onnx(
            model_id=model_uri,
            input_data=input_data,
        )
        output_bytes: bytes = json.dumps(output).encode()
        # post task output to the task manager contract
        tx = self.eth_client.task_manager.functions.submitTaskOutput(
            task_id, output_bytes
        ).build_transaction(
            {
                "from": self.operator_address,
                "gas": 2000000,
                "gasPrice": Web3.to_wei("20", "gwei"),
                "nonce": self.eth_client.w3.eth.get_transaction_count(
                    self.operator_address
                ),
                "chainId": self.eth_client.w3.eth.chain_id,
            }
        )
        signed_tx: SignedTransaction = self.eth_client.w3.eth.account.sign_transaction(
            tx, private_key=self.private_key
        )
        tx_hash = self.eth_client.w3.eth.send_raw_transaction(signed_tx.raw_transaction)
        receipt = self.eth_client.w3.eth.wait_for_transaction_receipt(tx_hash)
        if receipt.status != 1:
            console.print(
                "Failed to post task output to the contract", style=styles.error
            )
            raise Exception("Failed to post task output to the contract")

        console.print(
            f"Successfully posted an output for the {task_id} task",
            style=styles.agg_info,
        )

    def process_challenged_task(self, task_id: int):
        # get task details, `struct ISertnTaskManager.Task`
        task = self.eth_client.task_manager.functions.tasks(task_id).call()

        starting_block: int = task[TaskStructMap.START_BLOCK]  # uint256 startBlock
        model_id: int = task[TaskStructMap.MODEL_ID]  # uint256 modelId
        inputs: bytes = task[TaskStructMap.INPUTS]  # bytes inputs
        operator_address: str = task[TaskStructMap.OPERATOR]  # address operator

        if operator_address != self.operator_address:
            console.print(
                f"Operator address {operator_address} does not match "
                f"the operator address {self.operator_address}, skipping...",
                style=styles.debug,
            )
            return

        # generate proof for the task
        proof = self.generate_proof_for_task(inputs, task_id, model_id)
        proof_encoded = proof.encode()

        # submit proof to the task manager contract
        tx = self.eth_client.task_manager.functions.submitProofForTask(
            task_id, proof_encoded
        ).build_transaction(
            {
                "from": self.operator_address,
                "gas": 2000000,
                "gasPrice": Web3.to_wei("20", "gwei"),
                "nonce": self.eth_client.w3.eth.get_transaction_count(
                    self.operator_address
                ),
                "chainId": self.eth_client.w3.eth.chain_id,
            }
        )
        signed_tx: SignedTransaction = self.eth_client.w3.eth.account.sign_transaction(
            tx, private_key=self.private_key
        )
        tx_hash = self.eth_client.w3.eth.send_raw_transaction(signed_tx.raw_transaction)
        receipt = self.eth_client.w3.eth.wait_for_transaction_receipt(tx_hash)
        if receipt.status != 1:
            console.print("Failed to post proof to the contract", style=styles.error)
            return
        console.print(
            f"Successfully posted a proof for the {task_id} task",
            style=styles.agg_info,
        )

        # sign the proof
        encoded = eth_abi.encode(
            ["uint32", "bytes", "address"],
            [task_id, proof_encoded, operator_address],
        )
        message = encode_defunct(primitive=encoded)
        signed_message: SignedMessage = Account.sign_message(
            message, private_key=self.private_key
        )
        signature = signed_message.signature.hex()
        console.print(
            f"Signature generated, task_id: {starting_block}, model: {model_id}, "
            f"signature: {signature}",
            style=styles.debug,
        )

        # send the proof and signature to the aggregator server
        resp = requests.post(
            f"http://{self.config['aggregator_server_ip_port_address']}/proof",
            json={
                "task_id": task_id,
                "proof": proof,
                "signature": signature,
            },
        )
        try:
            resp.raise_for_status()
            console.print(
                f"Successfully posted a proof for the {task_id} task",
                style=styles.agg_info,
            )
        except requests.HTTPError:
            console.print(
                f"Failed to post a proof for the {task_id} task: {resp.text}",
                style=styles.error,
            )

    def generate_proof_for_task(
        self, inputs: bytes, task_id: int, model_id: int
    ) -> Optional[str]:
        """
        Generate proof for the task with the given task_id.
        """
        # generate proof for the task
        model_uri: str = self.eth_client.model_registry.functions.modelURI(
            model_id
        ).call()
        proof_generator = EZKLHandler(
            model_id=model_uri,
            task_id=str(task_id),
            inputs=[float(i) for i in inputs.decode().split(" ")],
        )
        proof_generator.gen_input_file()
        proof_result = proof_generator.gen_proof()
        if not proof_result or len(proof_result) == 0:
            return None
        proof: str = proof_result[0]
        return proof

    def post_task_output(
        self,
        task_id: int,
        output: bytes,
    ):
        """
        Post task output to the task manager contract.
        """
        tx = self.eth_client.task_manager.functions.submitTaskOutput(
            task_id, output
        ).build_transaction(
            {
                "from": self.aggregator_address,
                "gas": 2000000,
                "gasPrice": Web3.to_wei("20", "gwei"),
                "nonce": self.eth_client.w3.eth.get_transaction_count(
                    self.aggregator_address
                ),
                "chainId": self.eth_client.w3.eth.chain_id,
            }
        )
        signed_tx: SignedTransaction = self.eth_client.w3.eth.account.sign_transaction(
            tx, private_key=self.private_key
        )
        tx_hash = self.eth_client.w3.eth.send_raw_transaction(signed_tx.raw_transaction)
        receipt = self.eth_client.w3.eth.wait_for_transaction_receipt(tx_hash)
        if receipt.status != 1:
            console.print(
                "Failed to post task output to the contract", style=styles.error
            )
            return
        console.print(
            f"Successfully posted an output to the {task_id} task",
            style=styles.agg_info,
        )
