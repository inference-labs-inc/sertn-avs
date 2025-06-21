import asyncio
import threading

import pytest
import uvicorn

from aggregator.main import Aggregator
from avs_operator.main import TaskOperator
from common.contract_constants import TaskStateMap, TaskStructMap


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

        aggregator_server = threading.Thread(
            target=self.start_aggregator_server, args=[aggregator]
        )
        aggregator_server.start()

        # check events by the operator, it should process the challenge and generate proof
        processed_count = operator.listen_for_events(loop_running=False)
        assert processed_count == 1, "Operator should process one challenge"
        # the proof should be sent to the aggregator

        # here the task should be resolved by the aggregator
        task = aggregator.eth_client.task_manager.functions.getTask(task_id).call()
        assert task[TaskStructMap.STATE] == TaskStateMap.RESOLVED

        self.stop_aggregator_server()
        if aggregator_server:
            aggregator_server.join(timeout=5)  # Wait up to 5 seconds
