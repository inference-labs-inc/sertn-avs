import os
import json
from dotenv import load_dotenv

ROOT_DIR = os.path.dirname(os.path.abspath(__file__))
load_dotenv(dotenv_path=os.path.join(ROOT_DIR, "contracts", ".env"))

RPC_SERVER_URL = f"{os.getenv("RPC_HOST")}:{os.getenv("RPC_PORT")}"

OPERATOR_KEY = os.environ.get("OPERATOR_ECDSA_PRIV_KEY")
AGGREGATOR_KEY = os.environ.get("AGGREGATOR_ECDSA_PRIV_KEY")
BLS_PRIVATE_KEY_STORE_PATH = os.environ.get("BLS_PRIVATE_KEY_STORE_PATH")
OPERATOR_BLS_KEY_PASSWORD = os.environ.get("OPERATOR_BLS_KEY_PASSWORD")

AGGREGATOR_SERVER_HOST = os.environ.get("AGGREGATOR_SERVER_HOST")

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

# contracts ABIs:
with open(os.path.join(ROOT_DIR, "abis/SertnServiceManager.abi.json")) as f:
    SERVICE_MANAGER_ABI = f.read()

with open("abis/SertnTaskManager.abi.json") as f:
    TASK_MANAGER_ABI = json.loads(f.read())
