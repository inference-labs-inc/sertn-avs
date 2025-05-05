import os
import json

from common.constants import ROOT_DIR


def load_abi(file_name):
    with open(os.path.join(ROOT_DIR, "abis", file_name)) as f:
        return json.loads(f.read())


SERVICE_MANAGER_ABI = load_abi("SertnServiceManager.abi.json")

TASK_MANAGER_ABI = load_abi("SertnTaskManager.abi.json")

STRATEGY_ABI = load_abi("StrategyBase.abi.json")

ERC20_MOCK_ABI = load_abi("ERC20Mock.abi.json")

DELEGATION_MANAGER_ABI = load_abi("DelegationManager.abi.json")

STRATEGY_MANAGER_ABI = load_abi("StrategyManager.abi.json")
