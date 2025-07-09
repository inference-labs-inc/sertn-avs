import asyncio
import os
import threading
import time

import pytest
import uvicorn
from dotenv import load_dotenv
from eth_account import Account
from eth_account.datastructures import SignedMessage, SignedTransaction

from aggregator.main import Aggregator
from avs_operator.main import TaskOperator
from common.constants import CONTRACTS_DIR, STRATEGIES_ADDRESSES
from common.contract_constants import TaskStateMap, TaskStructMap

load_dotenv(os.path.join(CONTRACTS_DIR, ".env"))  # Load environment variables


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

    def test_process_task(self, aggregator: Aggregator, operator: TaskOperator):
        """Just a smoke test to ensure send_new_task runs without errors"""

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
        init_rewards: int = (
            aggregator.eth_client.service_manager.functions.getIntervalRewards(
                currentInterval,
                operator.operator_address,
                STRATEGIES_ADDRESSES[0],
            ).call()
        )

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
                currentInterval
            ).call()
        )
        strategies_in_interval: list = (
            aggregator.eth_client.service_manager.functions.getStrategiesInInterval(
                currentInterval
            ).call()
        )
        # Get the rewards for the operator after processing the task
        rewards = aggregator.eth_client.service_manager.functions.getIntervalRewards(
            currentInterval,
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
            rewards - init_rewards == model_cost
        ), "Operator should receive rewards equal to the model cost"

        # Fast forward blockchain time by one interval
        aggregator.eth_client.w3.provider.make_request(
            "evm_increaseTime", [interval_seconds]
        )
        aggregator.eth_client.w3.provider.make_request("evm_mine", [])

        # Submit rewards for the interval
        tx = aggregator.eth_client.service_manager.functions.submitRewardsForInterval(
            currentInterval
        ).build_transaction(
            {
                "from": owner_address,
                "gas": 2000000,
                "gasPrice": aggregator.eth_client.w3.to_wei("20", "gwei"),
                "nonce": aggregator.eth_client.w3.eth.get_transaction_count(
                    owner_address
                ),
                "chainId": aggregator.eth_client.w3.eth.chain_id,
            }
        )
        signed_tx: SignedTransaction = (
            aggregator.eth_client.w3.eth.account.sign_transaction(
                tx, private_key=os.getenv("PRIVATE_KEY")
            )
        )
        tx_hash = aggregator.eth_client.w3.eth.send_raw_transaction(
            signed_tx.raw_transaction
        )
        receipt = aggregator.eth_client.w3.eth.wait_for_transaction_receipt(tx_hash)
        if receipt.status != 1:
            raise Exception("Failed to submit rewards for the interval")

        self.stop_aggregator_server()
        if aggregator_server:
            aggregator_server.join(timeout=5)  # Wait up to 5 seconds
