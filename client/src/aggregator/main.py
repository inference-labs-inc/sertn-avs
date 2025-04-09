import os
import threading
import time

import eth_abi
import uvicorn
from eth_account import Account
from eth_account.datastructures import SignedTransaction
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from tqdm import tqdm
from web3 import Web3

from common.console import console, styles
from common.eth import EthereumClient, load_ecdsa_private_key


def run_aggregator(config: dict) -> None:
    console.print("Starting Sertn Aggregator...", style=styles.agg_info)

    with tqdm(total=100, desc="Initializing aggregator") as pbar:
        aggregator = Aggregator(config=config)
        pbar.update(25)
        # threading.Thread(target=aggregator.start_submitting_signatures, args=[]).start()
        threading.Thread(target=aggregator.start_sending_new_tasks, args=[]).start()
        aggregator.start_server()
        pbar.update(75)
        while True:
            time.sleep(10)


class SignatureRequest(BaseModel):
    """
    Pydantic model for operator request validation
    """

    task_id: int
    block_number: int
    # operator_id: str
    signature: dict
    inputs: bytes
    output: bytes
    prove: bytes
    model_id: int
    poc: int
    proveOnResponse: bool


def hasher(task: SignatureRequest) -> bytes:
    """
    Encode the data returned by an operator into a hash for signature verification
    """
    encoded = eth_abi.encode(
        ["uint32", "bytes", "uint256"],
        [task.task_id, task.output, task.model_id],
    )
    return Web3.keccak(encoded)


class Aggregator:
    def __init__(self, config: dict = None):
        self.config = config
        self.eth_client = EthereumClient(rpc_url=self.config["eth_rpc_url"])
        self.private_key = load_ecdsa_private_key(
            keystore_path=self.config["ecdsa_private_key_store_path"],
            password=os.environ.get("AGGREGATOR_ECDSA_KEY_PASSWORD", ""),
        )
        self.aggregator_address = Account.from_key(self.private_key).address

        self.tasks = {}
        self.taskResponses = {}

        # Initialize FastAPI app
        self.server_thread: threading.Thread | None = None
        self.app = FastAPI()

        # Add FastAPI route
        @self.app.post("/signature")
        async def submit_signature(data: SignatureRequest):
            try:
                # Signature class from eigensdk...
                # signature = Signature(data.signature["X"], data.signature["Y"])
                task_index = data.task_id
                console.print(
                    f"Received signature for task {task_index}", style=styles.agg_info
                )
                # TODO: check if the signature is valid somehow
                # self.bls_aggregation_service.process_new_signature(
                #     task_index, data, signature, data.operator_id
                # )
                return "true"
            except Exception as e:
                console.print(f"Submitting signature failed: {e}", style=styles.error)
                raise HTTPException(
                    status_code=500, detail="Failed to process signature"
                )

    def start_server(self):
        host, port = self.config["aggregator_server_ip_port_address"].split(":")
        self.server_thread = threading.Thread(
            target=uvicorn.run,
            args=(self.app,),
            kwargs={
                "host": host,
                "port": int(port),
                # "ssl_keyfile": ...,
                # "ssl_certfile": ...,
            },
            daemon=True,
        )
        self.server_thread.start()

    def send_new_task(self, i):
        # TODO: here we send input data to the task
        # Define the _task struct, the dict should correspond to the struct in the contract
        # here: `contracts/src/ISertnServiceManager.sol` -> `ISertnServiceManagerTypes.Task`
        task = {
            "operatorModelId_": 0,  # uint256 - Replace with the actual model ID
            # TODO: INPUT DATA here:
            "inputs_": "".encode(),  # bytes - actual input data
            "poc_": 100,  # uint256 - Proof of computation value (WTF is this?)
            "startTime_": 0,  # uint256 - Will be set in the contract
            "startingBlock_": 0,  # uint256 - Will be set in the contract
            "proveOnResponse_": True,  # bool - Whether the task is a proof of stake task
            "user_": self.aggregator_address,  # address TODO: what address should be here?
        }

        tx = self.eth_client.task_manager.functions.sendTask(task).build_transaction(
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
            console.print("Failed to send new task", style=styles.error)
            return

        event = self.task_manager.events.NewTaskCreated().process_log(
            receipt["logs"][0]
        )

        task_index = event["args"]["taskIndex"]
        console.print(
            f"Successfully sent the new task {task_index}", style=styles.agg_info
        )
        # TODO: ???
        # self.bls_aggregation_service.initialize_new_task(
        #     task_index=task_index,
        #     task_created_block=receipt["blockNumber"],
        #     quorum_numbers=nums_to_bytes([0]),
        #     quorum_threshold_percentages=[100],
        #     time_to_expiry=60000,
        # )
        return event["args"]["taskIndex"]

    def start_sending_new_tasks(self):
        i = 0
        while True:
            console.print("Sending new task", style=styles.debug)
            self.send_new_task(i)
            time.sleep(10)
            i += 1

    # NOTE: `RespondToTask` function is not implemented in the contract
    # def start_submitting_signatures(self):
    #     while True:
    #         logger.info("Waiting for response")
    #         aggregated_response = next(
    #             self.bls_aggregation_service.get_aggregated_responses()
    #         )
    #         logger.info(f"Aggregated response {aggregated_response}")
    #         response = aggregated_response.task_response
    #         # TODO: WTF is this?
    #         task = [
    #             response["number_to_be_squared"],
    #             response["block_number"],
    #             nums_to_bytes([0]),
    #             100,
    #         ]
    #         task_response = [response["task_index"], response["number_squared"]]
    #         non_signers_stakes_and_signature = [
    #             aggregated_response.non_signer_quorum_bitmap_indices,
    #             [g1_to_tupple(g1) for g1 in aggregated_response.non_signers_pubkeys_g1],
    #             [g1_to_tupple(g1) for g1 in aggregated_response.quorum_apks_g1],
    #             g2_to_tupple(aggregated_response.signers_apk_g2),
    #             g1_to_tupple(aggregated_response.signers_agg_sig_g1),
    #             aggregated_response.quorum_apk_indices,
    #             aggregated_response.total_stake_indices,
    #             aggregated_response.non_signer_stake_indices,
    #         ]

    #         tx = self.task_manager.functions.respondToTask(
    #             task, task_response, non_signers_stakes_and_signature
    #         ).build_transaction(
    #             {
    #                 "from": self.aggregator_address,
    #                 "gas": 2000000,
    #                 "gasPrice": self.eth_client.w3.to_wei("20", "gwei"),
    #                 "nonce": self.eth_client.w3.eth.get_transaction_count(
    #                     self.aggregator_address
    #                 ),
    #                 "chainId": self.eth_client.w3.eth.chain_id,
    #             }
    #         )
    #         signed_tx = self.eth_client.w3.eth.account.sign_transaction(
    #             tx, private_key=AGGREGATOR_KEY
    #         )
    #         tx_hash = self.eth_client.w3.eth.send_raw_transaction(signed_tx.raw_transaction)
    #         receipt = self.eth_client.w3.eth.wait_for_transaction_receipt(tx_hash)
