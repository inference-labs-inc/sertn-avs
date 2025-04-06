import asyncio
import json
import logging
import os
import threading
import time
from dataclasses import dataclass
from typing import List

import eth_abi
import uvicorn
import yaml
from eth_account import Account
from eth_account.datastructures import SignedTransaction
from eth_account.signers.local import LocalAccount
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from web3 import Web3

from eigensdk.chainio.clients.builder import BuildAllConfig, build_all
from eigensdk.chainio.utils import nums_to_bytes
from eigensdk.crypto.bls.attestation import (
    Signature,
    g1_to_tupple,
    g2_to_tupple,
)
from eigensdk.services.avsregistry import AvsRegistryService
from eigensdk.services.bls_aggregation.blsagg import (
    BlsAggregationService,
)
from eigensdk.services.operatorsinfo.operatorsinfo_inmemory import (
    OperatorsInfoServiceInMemory,
)
from constants import (
    TASK_MANAGER_ADDRESS,
    TASK_MANAGER_ABI,
    AGGREGATOR_KEY,
    AGGREGATOR_SERVER_HOST,
    RPC_SERVER_URL,
)

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


# Pydantic model for request validation
class SignatureRequest(BaseModel):
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
    XXX: maybe makes sense to move this to some `utils` module, so operator would be able to use it too
    """
    encoded = eth_abi.encode(
        ["uint32", "bytes", "uint256"],
        [task.task_id, task.output, task.model_id],
    )
    return Web3.keccak(encoded)


class Aggregator:
    def __init__(self):
        # self.config = config
        self.web3 = Web3(Web3.HTTPProvider(RPC_SERVER_URL))
        self.aggregator_address = Account.from_key(AGGREGATOR_KEY).address
        # self.__load_clients()
        self.pk_wallet: LocalAccount = (
            Account.from_key(AGGREGATOR_KEY) if AGGREGATOR_KEY else None
        )
        self.task_manager = self.web3.eth.contract(
            address=TASK_MANAGER_ADDRESS, abi=TASK_MANAGER_ABI
        )
        # self.__load_bls_aggregation_service()
        self.tasks = {}
        self.taskResponses = {}

        # Initialize FastAPI app
        self.server_thread: threading.Thread | None = None
        self.app = FastAPI()

        # Add FastAPI route
        @self.app.post("/signature")
        async def submit_signature(data: SignatureRequest):
            try:
                signature = Signature(data.signature["X"], data.signature["Y"])
                task_index = data.task_id
                logger.info(f"Received signature for task {task_index}")
                # TODO: check if the signature is valid somehow
                # self.bls_aggregation_service.process_new_signature(
                #     task_index, data, signature, data.operator_id
                # )
                return "true"
            except Exception as e:
                logger.error(f"Submitting signature failed: {e}")
                raise HTTPException(
                    status_code=500, detail="Failed to process signature"
                )

    def start_server(self):
        host, port = AGGREGATOR_SERVER_HOST.split(":")
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
            "user_": self.pk_wallet.address,  # address TODO: what address should be here?
        }

        tx = self.task_manager.functions.sendTask(task).build_transaction(
            {
                "from": self.aggregator_address,
                "gas": 2000000,
                "gasPrice": self.web3.to_wei("20", "gwei"),
                "nonce": self.web3.eth.get_transaction_count(self.aggregator_address),
                "chainId": self.web3.eth.chain_id,
            }
        )

        # try:
        #     # Simulate the transaction
        #     self.web3.eth.call(tx)
        # except Exception as e:
        #     print(f"Transaction simulation failed: {e}")

        signed_tx: SignedTransaction = self.web3.eth.account.sign_transaction(
            tx, private_key=AGGREGATOR_KEY
        )
        tx_hash = self.web3.eth.send_raw_transaction(signed_tx.raw_transaction)
        receipt = self.web3.eth.wait_for_transaction_receipt(tx_hash)

        if not receipt["logs"]:
            logger.error("Failed to send new task")
            return

        event = self.task_manager.events.NewTaskCreated().process_log(
            receipt["logs"][0]
        )

        task_index = event["args"]["taskIndex"]
        logger.info(f"Successfully sent the new task {task_index}")
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
            logger.info("Sending new task")
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
    #                 "gasPrice": self.web3.to_wei("20", "gwei"),
    #                 "nonce": self.web3.eth.get_transaction_count(
    #                     self.aggregator_address
    #                 ),
    #                 "chainId": self.web3.eth.chain_id,
    #             }
    #         )
    #         signed_tx = self.web3.eth.account.sign_transaction(
    #             tx, private_key=AGGREGATOR_KEY
    #         )
    #         tx_hash = self.web3.eth.send_raw_transaction(signed_tx.raw_transaction)
    #         receipt = self.web3.eth.wait_for_transaction_receipt(tx_hash)

    def __load_clients(self):
        pass
        # cfg = BuildAllConfig(
        #     eth_http_url=self.config["eth_rpc_url"],
        #     avs_name="incredible-squaring",
        #     registry_coordinator_addr=self.config["avs_registry_coordinator_address"],
        #     operator_state_retriever_addr=self.config[
        #         "operator_state_retriever_address"
        #     ],
        #     prom_metrics_ip_port_address="",
        # )
        # self.clients = build_all(cfg, AGGREGATOR_KEY, logger)

        # self.eth_http_client = Web3(Web3.HTTPProvider(self.config["eth_rpc_url"]))

    # def __load_bls_aggregation_service(self):
    #     operator_info_service = OperatorsInfoServiceInMemory(
    #         avs_registry_reader=self.clients.avs_registry_reader,
    #         start_block_pub=0,
    #         start_block_socket=0,
    #         logger=logger,
    #     )

    #     avs_registry_service = AvsRegistryService(
    #         self.clients.avs_registry_reader, operator_info_service, logger
    #     )

    #     self.bls_aggregation_service = BlsAggregationService(
    #         avs_registry_service, hasher
    #     )


if __name__ == "__main__":
    # sertnServiceManager address - 0x9d4454B023096f34B160D6B654540c56A1F81688
    aggregator = Aggregator()
    # threading.Thread(target=aggregator.start_submitting_signatures, args=[]).start()
    threading.Thread(target=aggregator.start_sending_new_tasks, args=[]).start()
    aggregator.start_server()
    while True:
        time.sleep(10)
