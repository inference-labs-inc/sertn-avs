import os
import json

CLIENT_PATH = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", ".."))
ROOT_DIR = os.path.abspath(os.path.join(CLIENT_PATH, ".."))
MODELS_DATA_DIR = os.path.join(CLIENT_PATH, "src", "models", "models_data")
TEMP_FOLDER = os.path.join(CLIENT_PATH, "src", "models", "temp")
PROOFS_FOLDER = os.path.join(CLIENT_PATH, "src", "models", "generated_proofs")
MODELS_FOLDER = os.path.join(CLIENT_PATH, "src", "models", "models")

os.makedirs(MODELS_DATA_DIR, exist_ok=True)
os.makedirs(TEMP_FOLDER, exist_ok=True)
os.makedirs(PROOFS_FOLDER, exist_ok=True)

# contracts addresses:
with open(
    os.path.join(ROOT_DIR, "contracts", "deployments", "sertnDeployment.json")
) as f:
    deployment_info = json.load(f)
    TASK_MANAGER_ADDRESS = deployment_info["sertnTaskManager"]
    SERVICE_MANAGER_ADDRESS = deployment_info["sertnServiceManager"]
    ALLOCATION_MANAGER_ADDRESS = deployment_info["allocationManager"]
    STRATEGIES_ADDRESSES = [
        deployment_info["strategy_0"],
        deployment_info["strategy_1"],
        deployment_info["strategy_2"],
    ]
    ETH_STRATEGY_ADDRESSES = [
        deployment_info["eth_strategy_0"],
        deployment_info["eth_strategy_1"],
    ]

IGNORED_MODEL_HASHES = []

# Queue size limits
MAX_EVALUATION_ITEMS = 1024

# the aggregator requests a proof for some responses randomly,
# the probability of that is configurable, but here we have a default value
DEFAULT_PROOF_REQUEST_PROBABILITY = 0.4

# how many blocks are we going to wait for an operator response before rejecting the task?
# TODO: share this constant with contracts (we need that there too)
RESOLVE_BLOCKS_DELAY = 300

# Default proof size when we're unable to determine the actual size
DEFAULT_PROOF_SIZE = 5000
# Size in percent of the sample to be used for the maximum score median
MAXIMUM_SCORE_MEDIAN_SAMPLE = 0.05
# The maximum timespan allowed for miners to process through a circuit
CIRCUIT_TIMEOUT_SECONDS = 60

# Operator set ID
OPERATOR_SET_ID = 0
