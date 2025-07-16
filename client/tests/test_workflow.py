import asyncio
import os
import threading
import time

import pytest
import requests
import uvicorn
from dotenv import load_dotenv
from eth_account import Account
from eth_account.datastructures import SignedMessage, SignedTransaction

from aggregator.main import Aggregator
from avs_operator.main import TaskOperator
from common.constants import ROOT_DIR, STRATEGIES_ADDRESSES
from common.contract_constants import TaskStateMap, TaskStructMap
from owner.main import AvsOwner

load_dotenv(os.path.join(ROOT_DIR, ".env"))  # Load environment variables


class TestWorkflow:
    @pytest.fixture(scope="session")
    def aggregator(self):
        return Aggregator(
            {
                "eth_rpc_url": "http://localhost:8545",
                "aggregator_server_ip_port_address": "localhost:8090",
                "ecdsa_private_key_store_path": "tests/keys/aggregator.ecdsa.key.json",
                "proof_request_probability": 1.0,  # challenge every task
            }
        )

    @pytest.fixture(scope="session")
    def operator(self):
        return TaskOperator(
            {
                "eth_rpc_url": "http://localhost:8545",
                "aggregator_server_ip_port_address": "localhost:8090",
                "ecdsa_private_key_store_path": "tests/keys/operator.ecdsa.key.json",
            }
        )

    @pytest.fixture(scope="session")
    def owner(self):
        return AvsOwner(
            private_key=os.getenv("PRIVATE_KEY"), eth_rpc_url="http://localhost:8545"
        )

    @pytest.fixture(scope="function")
    def init_environment(self, aggregator: Aggregator):
        """Ensure the Anvil node is running and has the correct state"""
        # Fast-forward the blockchain to 400 blocks,
        # because current block number must be greater than MIN_WITHDRAWAL_DELAY_BLOCKS
        try:
            requests.post(
                "http://localhost:8545",
                json={
                    "jsonrpc": "2.0",
                    "method": "anvil_mine",
                    "params": ["0x190"],
                    "id": 1,
                },
                headers={"Content-Type": "application/json"},
                timeout=60,  # Set a timeout for the request
            )
        except requests.exceptions.ReadTimeout:
            pass
        # Get the current blockchain timestamp
        current_block = aggregator.eth_client.w3.eth.get_block("latest")
        current_timestamp = current_block["timestamp"]

        # Get interval duration once
        interval_seconds = (
            aggregator.eth_client.rewards_coordinator.functions.CALCULATION_INTERVAL_SECONDS().call()
        )
        init_interval_datetime = (
            aggregator.eth_client.service_manager.functions.firstIntervalTimestamp().call()
        )

        # Fast forward time by the interval duration
        aggregator.eth_client.w3.provider.make_request(
            "evm_increaseTime",
            [init_interval_datetime + interval_seconds - current_timestamp + 1],
        )
        aggregator.eth_client.w3.provider.make_request("evm_mine", [])

        # Get the owner address from the private key
        owner_address = Account.from_key(os.getenv("PRIVATE_KEY")).address
        # Get the current interval and initial operator rewards
        currentInterval: int = (
            aggregator.eth_client.service_manager.functions.getCurrentInterval().call()
        )
        return {
            "current_interval": currentInterval,
            "owner_address": owner_address,
            "interval_seconds": interval_seconds,
        }

    def start_aggregator_server(self, aggregator: Aggregator):
        self._stop_event = threading.Event()
        config = uvicorn.Config(
            app=aggregator.app, host="0.0.0.0", port=8090, log_level="info"
        )
        self.server = uvicorn.Server(config)

        # Run server until stop event is set
        loop = asyncio.new_event_loop()
        asyncio.set_event_loop(loop)

        try:
            loop.run_until_complete(self.server.serve())
        except asyncio.CancelledError:
            pass
        finally:
            loop.close()

    def stop_aggregator_server(self):
        """Stop server thread"""
        if self.server:
            self.server.should_exit = True
        self._stop_event.set()

    def test_process_task(
        self,
        aggregator: Aggregator,
        operator: TaskOperator,
        owner: AvsOwner,
        init_environment: dict,
    ):
        """Just a smoke test to ensure send_new_task runs without errors"""

        # create a new task
        task_id = aggregator.send_new_task(1)
        assert task_id is not None, "Task ID should not be None"
        # here task  should be assigned to the operator

        # process the task by the operator
        processed_count = operator.listen_for_events(loop_running=False)
        assert processed_count == 1, "Operator should process one task"
        # the task should be marked as completed

        # checkout the completed task
        processed_count = aggregator.listen_for_events(loop_running=False)
        assert processed_count == 1, "Aggregator should process one task"
        # At this point the task should be challenged

        # start aggregator server in a separate thread
        # This is necessary to allow the aggregator to listen for events and process the challenge
        aggregator_server = threading.Thread(
            target=self.start_aggregator_server, args=[aggregator]
        )
        aggregator_server.start()
        time.sleep(5)

        # check events by the operator, it should process the challenge and generate proof
        processed_count = operator.listen_for_events(loop_running=False)
        assert processed_count == 1, "Operator should process one challenge"
        # the proof should be sent to the aggregator

        # here the task should be resolved by the aggregator
        task = aggregator.eth_client.task_manager.functions.getTask(task_id).call()
        assert task[TaskStructMap.STATE] == TaskStateMap.RESOLVED
        model_id = task[TaskStructMap.MODEL_ID]

        # check rewards collected for the operator
        operators_in_interval: list = (
            aggregator.eth_client.service_manager.functions.getOperatorsInInterval(
                init_environment["current_interval"]
            ).call()
        )
        strategies_in_interval: list = (
            aggregator.eth_client.service_manager.functions.getStrategiesInInterval(
                init_environment["current_interval"]
            ).call()
        )
        # Get the rewards for the operator after processing the task
        rewards = aggregator.eth_client.service_manager.functions.getIntervalRewards(
            init_environment["current_interval"],
            operator.operator_address,
            STRATEGIES_ADDRESSES[0],
        ).call()
        model_cost: int = aggregator.eth_client.model_registry.functions.computeCost(
            model_id
        ).call()
        assert operators_in_interval == [
            operator.operator_address,
        ], "Operator should be in the current interval"
        assert strategies_in_interval == [
            STRATEGIES_ADDRESSES[0],
        ], "Aggregator should be in the current interval"
        assert (
            rewards == model_cost
        ), "Operator should receive rewards equal to the model cost"

        # Fast forward blockchain time by one interval
        aggregator.eth_client.w3.provider.make_request(
            "evm_increaseTime", [init_environment["interval_seconds"]]
        )
        aggregator.eth_client.w3.provider.make_request("evm_mine", [])

        # Submit rewards for the interval
        owner.submit_rewards_for_interval(init_environment["current_interval"])

        self.stop_aggregator_server()
        if aggregator_server:
            aggregator_server.join(timeout=5)  # Wait up to 5 seconds

    def test_task_incorrect_proof(
        self,
        aggregator: Aggregator,
        operator: TaskOperator,
        owner: AvsOwner,
        init_environment: dict,
    ):
        """Just a smoke test to ensure send_new_task runs without errors"""

        # Mock the generate_proof_for_task method to return incorrect proof
        def mock_generate_proof_for_task(task_id):
            return "incorrect_proof"

        operator.generate_proof_for_task = mock_generate_proof_for_task
        initial_shares = (
            aggregator.eth_client.delegation_manager.functions.operatorShares(
                operator.operator_address, STRATEGIES_ADDRESSES[0]
            ).call()
        )
        initial_rewards = (
            aggregator.eth_client.service_manager.functions.getIntervalRewards(
                init_environment["current_interval"],
                operator.operator_address,
                STRATEGIES_ADDRESSES[0],
            ).call()
        )

        # create a new task
        task_id = aggregator.send_new_task(2)
        assert task_id is not None, "Task ID should not be None"
        # here task  should be assigned to the operator

        # process the task by the operator
        processed_count = operator.listen_for_events(loop_running=False)
        assert processed_count == 1, "Operator should process one task"
        # the task should be marked as completed

        # checkout the completed task
        processed_count = aggregator.listen_for_events(loop_running=False)
        assert processed_count == 1, "Aggregator should process one task"
        # At this point the task should be challenged

        # start aggregator server in a separate thread
        # This is necessary to allow the aggregator to listen for events and process the challenge
        aggregator_server = threading.Thread(
            target=self.start_aggregator_server, args=[aggregator]
        )
        aggregator_server.start()
        time.sleep(5)

        # check events by the operator, it should process the challenge and generate proof
        processed_count = operator.listen_for_events(loop_running=False)
        assert processed_count == 1, "Operator should process one challenge"
        # the proof should be sent to the aggregator

        # here the task should be resolved by the aggregator
        task = aggregator.eth_client.task_manager.functions.getTask(task_id).call()
        assert task[TaskStructMap.STATE] == TaskStateMap.REJECTED.value
        # model_id = task[TaskStructMap.MODEL_ID]

        # check rewards collected for the operator
        operators_in_interval: list = (
            aggregator.eth_client.service_manager.functions.getOperatorsInInterval(
                init_environment["current_interval"]
            ).call()
        )
        assert (
            operator.operator_address not in operators_in_interval
        ), "Operator should NOT be in the current interval"

        final_shares = (
            aggregator.eth_client.delegation_manager.functions.operatorShares(
                operator.operator_address, STRATEGIES_ADDRESSES[0]
            ).call()
        )
        assert (
            initial_shares - final_shares
        ) / initial_shares == 0.1, "Operator's shares are reduced by 10%"

        final_rewards = (
            aggregator.eth_client.service_manager.functions.getIntervalRewards(
                init_environment["current_interval"],
                operator.operator_address,
                STRATEGIES_ADDRESSES[0],
            ).call()
        )
        assert final_rewards == initial_rewards, "No new rewards for the operator"

        self.stop_aggregator_server()
        if aggregator_server:
            aggregator_server.join(timeout=5)  # Wait up to 5 seconds
