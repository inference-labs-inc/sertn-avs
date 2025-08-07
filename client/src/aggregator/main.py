import os
import random
import threading
import time

import eth_abi
import uvicorn
from eth_account import Account
from eth_account.datastructures import SignedTransaction
from eth_account.messages import encode_defunct
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from web3 import Web3

from common.abis import ERC20_ABI, STRATEGY_ABI
from common.config import AggregatorConfig
from common.console import console, styles
from common.constants import (
    DEFAULT_PROOF_REQUEST_PROBABILITY,
    RESOLVE_BLOCKS_DELAY,
    OPERATOR_SET_ID,
)
from common.contract_constants import (
    TaskStateMap,
    TaskStructMap,
    VerificationStrategiesMap,
)
from common.eth import EthereumClient, load_ecdsa_private_key
from models.proof.ezkl_handler import EZKLHandler


def run_aggregator(config: AggregatorConfig) -> None:
    console.print("Starting Sertn Aggregator...", style=styles.agg_info)
    aggregator = Aggregator(config=config)
    threading.Thread(target=aggregator.start_sending_new_tasks, args=[]).start()
    threading.Thread(target=aggregator.listen_for_events, args=[]).start()
    threading.Thread(target=aggregator.check_pending_tasks, args=[]).start()
    aggregator.start_server()

    # with tqdm(total=100, desc="Initializing aggregator") as pbar:
    #     aggregator = Aggregator(config=config)
    #     pbar.update(25)
    #     # threading.Thread(target=aggregator.start_submitting_signatures, args=[]).start()
    #     threading.Thread(target=aggregator.start_sending_new_tasks, args=[]).start()
    #     aggregator.start_server()
    #     pbar.update(75)
    #     while True:
    #         time.sleep(10)


class InvalidProofError(ValueError):
    pass


class ProofRequest(BaseModel):
    """
    Pydantic model for operator submits proof to the aggregator
    """

    task_id: int
    proof: str
    signature: str


class Aggregator:
    def __init__(self, config: AggregatorConfig = None):
        self.config = config
        self.eth_client = EthereumClient(rpc_url=self.config.eth_rpc_url)
        self.private_key = load_ecdsa_private_key(
            keystore_path=str(self.config.ecdsa_private_key_store_path),
            password=os.environ.get("AGGREGATOR_ECDSA_KEY_PASSWORD", ""),
        )
        # the aggregator gonna request proofs for some responses randomly, so probability:
        self.proof_req_probability = self.config.proof_request_probability
        self.aggregator_address = Account.from_key(self.private_key).address

        self.tasks = {}
        self.taskResponses = {}

        # Initialize FastAPI app
        self.server_thread: threading.Thread | None = None
        self.app = FastAPI()

        # Add FastAPI route
        @self.app.post("/proof")
        async def _(data: ProofRequest):
            try:
                self.process_submitted_proof(data.task_id, data.proof, data.signature)
            except InvalidProofError as exc:
                raise HTTPException(status_code=400, detail=str(exc))

    def start_server(self):
        host, port = self.config.aggregator_server_ip_port_address.split(":")
        uvicorn.run(
            self.app,
            host=host,
            port=int(port),
            # ssl_keyfile=self.config["ssl_keyfile"],
            # ssl_certfile=self.config["ssl_certfile"],
        )
        # no need to run in a separate thread for now, maybe add some CLI arg to start it like a daemon
        # self.server_thread = threading.Thread(
        #     target=uvicorn.run,
        #     args=(self.app,),
        #     kwargs={
        #         "host": host,
        #         "port": int(port),
        #         # "ssl_keyfile": ...,
        #         # "ssl_certfile": ...,
        #     },
        #     daemon=True,
        # )
        # self.server_thread.start()

    def start_sending_new_tasks(self, loop_running: bool = True):
        """
        Runs in a separate thread and sends new tasks to the task manager.
        TODO: make sleep time configurable
        """
        i = 0
        while True:
            console.print("Sending new task", style=styles.debug)
            self.send_new_task(i)
            if not loop_running:
                console.print("Stopping sending new tasks", style=styles.debug)
                break
            time.sleep(60)
            i += 1

    def send_new_task(self, i) -> int | None:
        # create some generic input data
        inputs = " ".join(str(random.uniform(0.0, 0.85)) for _ in range(5))

        model_id = self.get_model_id()
        if model_id is None:
            console.print(
                "No models found in the model registry, cannot send a new task",
                style=styles.error,
            )
            return None

        operator_address = self.get_random_operator()
        if operator_address is None:
            console.print(
                "No operators found, cannot send a new task", style=styles.error
            )
            return None

        # Define the _task struct, the dict should correspond to the struct in the contract
        task = {
            "startBlock": self.eth_client.w3.eth.block_number,  # uint256 - Current block number
            "startTimestamp": self.eth_client.w3.eth.get_block("latest")[
                "timestamp"
            ],  # uint256 - Current timestamp
            "modelId": model_id,  # uint256
            "inputs": inputs.encode(),  # bytes - actual input data
            "proofHash": b"\x00" * 32,  # bytes32 - Hash of the proof, empty for now
            "user": self.aggregator_address,  # at the moment the aggregator is the end user
            "nonce": i,  # uint256 - Nonce for the task, can be any number
            "operator": operator_address,  # address - Operator address, can be any address
            "state": TaskStateMap.CREATED.value,  # uint8 - gonna be changed by the contract anyway
            "output": b"",  # bytes - Output of the task, empty for now
            "fee": self.get_model_cost(
                model_id
            ),  # uint256 - Fee for the task, can be any number
        }

        token = self.get_token(task)

        print(f"Approving funds for the task - {task['fee']} tokens")
        tx = token.functions.approve(
            self.eth_client.service_manager.address,  # Approve the task manager to spend tokens
            # TODO: check is the fee correct after converting to wei
            Web3.to_wei(task["fee"], "ether"),  # Approve the fee amount
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
            console.print("Failed to approve funds for the task", style=styles.error)
            return
        console.print(
            f"Successfully approved {task['fee']} tokens for the task",
            style=styles.agg_info,
        )

        # XXX: Approve funds for the task?
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

        console.print(f"New task sent!", style=styles.agg_info)

        # also `TaskCreated` event is emitted, do we need it?
        event = self.eth_client.task_manager.events.TaskAssigned().process_log(
            receipt["logs"][-1]
        )

        task_index = event["args"]["taskId"]
        console.print(
            f"Successfully sent the new task {task_index}", style=styles.agg_info
        )

        return task_index

    def get_random_operator(self) -> str | None:
        """
        Get a random operator address from the allocation manager.
        """
        operators = self.eth_client.allocation_manager.functions.getMembers(
            {
                "avs": self.eth_client.service_manager.address,
                "id": OPERATOR_SET_ID,
            }
        ).call()
        if operators:
            return random.choice(operators)
        else:
            console.print(
                "No operators found in the allocation manager", style=styles.error
            )
            return None

    def get_model_id(self) -> int:
        """
        Get a model ID from the model registry.
        """
        models_count = self.eth_client.model_registry.functions.modelIndex().call() - 1
        if models_count <= 0:
            console.print("No models found in the model registry", style=styles.error)
            return None
        # return a random model ID from 1 to models_count
        return random.randint(1, models_count)

    def get_model_cost(self, model_id: int) -> int:
        """
        Get the compute cost for a model by its ID.
        This is just a placeholder function, you might want to implement a more sophisticated logic.
        """
        cost: int = self.eth_client.model_registry.functions.computeCost(
            model_id
        ).call()
        if not cost:
            console.print(f"Model {model_id} has zero cost!!!", style=styles.error)
            cost = 0
        return cost

    def get_token(self, task: dict):
        """
        Get the token address used for the task fees.
        """
        allocated_sets = self.eth_client.allocation_manager.functions.getAllocatedSets(
            task["operator"]
        ).call()
        print(
            f"Got {len(allocated_sets)} allocated sets for operator {task['operator']}"
        )
        strategy_addresses = (
            self.eth_client.allocation_manager.functions.getAllocatedStrategies(
                task["operator"], allocated_sets[0]
            ).call()
        )
        print(
            f"Got {len(strategy_addresses)} allocated strategies for operator {task['operator']}"
        )

        strategy = self.eth_client.w3.eth.contract(
            address=strategy_addresses[0],
            abi=STRATEGY_ABI,
        )
        print(f"Got strategy {strategy.address} for operator {task['operator']}")

        token_address = strategy.functions.underlyingToken().call()
        print(f"Got token {token_address} for strategy {strategy.address}")

        return self.eth_client.w3.eth.contract(
            address=token_address,
            abi=ERC20_ABI,
        )

    def listen_for_events(self, loop_running: bool = True) -> int | None:
        """
        Listen for task manager events in a separate thread.
        """
        console.print("Listening for task manager events...", style=styles.debug)
        task_completed_events = (
            self.eth_client.task_manager.events.TaskCompleted.create_filter(
                from_block="latest"
            )
        )
        processed_count: int = 0
        while True:
            for event in task_completed_events.get_new_entries():
                console.print(
                    f"Some task has been completed. Processing...", style=styles.debug
                )
                self.process_completed_task(event)
                processed_count += 1

            if not loop_running:
                console.print("Stopping event listener...", style=styles.debug)
                return processed_count

            time.sleep(3)

    def process_completed_task(self, event):
        """
        Process a completed task event.
        This function is called when a task is completed and the `TaskCompleted` event is emitted.
        It checks the task state and either resolves the task or challenges it for a proof.
        """
        task_id = event["args"]["taskId"]
        console.print(f"Processing completed task #{task_id}", style=styles.agg_info)
        # get task details from the contract
        task = self.eth_client.task_manager.functions.getTask(task_id).call()
        task_state = task[TaskStructMap.STATE]  # TaskStateMap
        if task_state != TaskStateMap.COMPLETED.value:
            console.print(
                f"Task {task_id} has invalid state, current state: {task_state}",
                style=styles.error,
            )
            return
        # no we are going to mark the task resolved or challenge it, requesting a proof
        # that's gonna be random. With configurable probability of proof requests
        if random.random() < self.proof_req_probability:
            # gonna request a proof
            self.challenge_task(task_id)
        else:
            # here we have a winner, gonna just resolve the task
            self.resolve_task(task_id=task_id, resolved=True)

    def challenge_task(self, task_id: int):
        """
        Challenge a task by its ID.
        Two cases of using this function:
        1. The task is not completed for a long time (300 blocks or something). So we are slashing the operator
        2. The task is completed, but we want to get a proof from the operator.
        """
        tx = self.eth_client.task_manager.functions.challengeTask(
            task_id
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
            console.print("Failed to challenge the task", style=styles.error)
            return
        console.print(
            f"Successfully challenged the task #{task_id}",
            style=styles.agg_info,
        )

    def process_submitted_proof(self, task_id: int, proof: str, signature: str) -> bool:
        console.print(
            f"Processing proof submitted for the task #{task_id}...",
            style=styles.agg_info,
        )
        task: tuple = self.eth_client.task_manager.functions.getTask(task_id).call()
        operator_address: str = task[TaskStructMap.OPERATOR]

        # check the signature (the origin operator submitted us the data)
        encoded = eth_abi.encode(
            ["uint32", "bytes", "address"],
            [task_id, proof.encode(), operator_address],
        )
        message = encode_defunct(primitive=encoded)
        recovered_address = Account.recover_message(
            message, signature=bytes.fromhex(signature)
        )
        if recovered_address != operator_address:
            console.print(
                f"Signature verification failed: {recovered_address} != {operator_address}",
                style=styles.error,
            )
            raise InvalidProofError("Signature verification failed")

        # check the task state
        state: int = task[TaskStructMap.STATE]
        if state in (TaskStateMap.RESOLVED.value, TaskStateMap.REJECTED.value):
            console.print(
                f"The task #{task_id} is already {TaskStateMap.from_int(state)}, skipping...",
                style=styles.error,
            )
            return InvalidProofError("The task is already resolved")
        if state != TaskStateMap.CHALLENGED.value:
            console.print(
                f"The task #{task_id} isn't challenged. State is {TaskStateMap.from_int(state)}, skipping...",
                style=styles.error,
            )
            return InvalidProofError("The task isn't challenged")

        # check the task verification strategy
        model_id = task[TaskStructMap.MODEL_ID]
        verification_strategy: str = (
            self.eth_client.model_registry.functions.verificationStrategy(
                model_id
            ).call()
        )
        if verification_strategy != VerificationStrategiesMap.OFFCHAIN:
            console.print(
                f"The task #{task_id} is not offchain verifiable...",
                style=styles.error,
            )
            return InvalidProofError("Not verifiable")
        model_uri = self.eth_client.model_registry.functions.modelURI(model_id).call()
        if not model_uri:
            console.print(
                f"The model {model_id} has no URI, cannot verify the proof...",
                style=styles.error,
            )
            return InvalidProofError("Model URI is empty")

        inputs: bytes = task[TaskStructMap.INPUTS]
        output: bytes = task[TaskStructMap.OUTPUT]
        proof_hash: bytes = task[TaskStructMap.PROOF_HASH]

        # check the task hash (is the same one submitted to the chain?)
        if proof_hash != Web3.keccak(proof.encode()):
            print(
                "Proof hashes mismatch! Proof, submitted to the aggregator differs from"
                f" one submitted to the chain. Task #{task_id} Rejecting...",
                style=styles.error,
            )
            self.resolve_task(task_id=task_id, resolved=False)  # slash the operator
            raise InvalidProofError("Invalid proof hash")

        # OK, we are good to verify the task
        validator_input = [float(i) for i in inputs.decode().split(" ")]
        proof_generator = EZKLHandler(
            model_id=model_uri,
            task_id=str(task_id),
            inputs=validator_input,
        )
        verified: bool = proof_generator.verify_proof(
            validator_inputs=[validator_input], proof=proof
        )
        if verified:
            # cool, the operator might receive the reward
            console.print(f"The task #{task_id} is verified!", style=styles.agg_info)
            self.resolve_task(task_id=task_id, resolved=True)
            return True
        else:
            # not good... the proof is not verified
            console.print(
                f"The proof for task #{task_id} is invalid!", style=styles.error
            )
            self.resolve_task(task_id=task_id, resolved=False)
            raise InvalidProofError("Proof is not verified")

    def resolve_task(self, task_id: int, resolved: bool):
        """
        Resolve a task by its ID.
        If `resolved` is True, the task is resolved as completed.
        If `resolved` is False, the task is resolved as rejected and the operator is slashed.
        """
        tx = self.eth_client.task_manager.functions.resolveTask(
            task_id, resolved
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
            console.print("Failed to resolve the task", style=styles.error)
            return
        console.print(
            f"Successfully resolved the task #{task_id}",
            style=styles.agg_info,
        )

    def check_pending_tasks(self):
        """
        Check for pending tasks that are not completed for a long time.
        Intended to run in a separate thread.
        """
        while True:
            pending_ids: list[int] = (
                self.eth_client.task_manager.functions.getPendingTasksIds().call()
            )
            current_block: int = self.eth_client.w3.eth.block_number
            for task_id in pending_ids:
                console.print(
                    f"Checking pending task #{task_id}...", style=styles.debug
                )
                task: tuple = self.eth_client.task_manager.functions.getTask(
                    task_id
                ).call()
                task_state: int = task[TaskStructMap.STATE]
                start_block: int = task[TaskStructMap.START_BLOCK]
                if (
                    task_state == TaskStateMap.ASSIGNED.value
                    and start_block + RESOLVE_BLOCKS_DELAY < current_block
                ):
                    # the task is assigned, but not completed for a long time
                    # we are going to challenge it and slash the operator
                    console.print(
                        f"Task #{task_id} is assigned, but not completed for a long time. Challenging...",
                        style=styles.agg_info,
                    )
                    self.challenge_task(task_id)
                elif (
                    task_state == TaskStateMap.CHALLENGED.value
                    and start_block + RESOLVE_BLOCKS_DELAY * 2 < current_block
                ):
                    # the task is challenged, but not completed for a long time
                    # we are going to resolve it as rejected and slash the operator
                    console.print(
                        f"Task #{task_id} is challenged, but a proof is not completed for a long time. "
                        "Challenging again... Contract will slash the operator.",
                        style=styles.agg_info,
                    )
                    self.challenge_task(task_id=task_id)
            time.sleep(60)
