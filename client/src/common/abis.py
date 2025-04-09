import os
import json

from common.constants import ROOT_DIR

with open(os.path.join(ROOT_DIR, "abis", "SertnServiceManager.abi.json")) as f:
    SERVICE_MANAGER_ABI = json.loads(f.read())

with open(os.path.join(ROOT_DIR, "abis", "SertnTaskManager.abi.json")) as f:
    TASK_MANAGER_ABI = json.loads(f.read())
