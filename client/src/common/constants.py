import os
import json

CLIENT_PATH = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", ".."))
ROOT_DIR = os.path.abspath(os.path.join(CLIENT_PATH, ".."))

# contracts addresses:
with open(
    os.path.join(ROOT_DIR, "contracts", "deployments", "sertnDeployment.json")
) as f:
    deployment_info = json.load(f)
    TASK_MANAGER_ADDRESS = deployment_info["sertnTaskManager"]
    SERVICE_MANAGER_ADDRESS = deployment_info["sertnServiceManager"]
    STRATEGIES_ADDRESSES = [
        deployment_info["strategy_0"],
        deployment_info["strategy_1"],
        deployment_info["strategy_2"],
    ]
    ETH_STRATEGY_ADDRESSES = [
        deployment_info["eth_strategy_0"],
        deployment_info["eth_strategy_1"],
    ]
