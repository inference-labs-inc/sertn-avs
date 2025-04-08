from web3 import Web3
from typing import Optional

from common.abis import TASK_MANAGER_ABI, SERVICE_MANAGER_ABI
from common.constants import (
    TASK_MANAGER_ADDRESS,
    SERVICE_MANAGER_ADDRESS,
)


class EthereumClient:
    def __init__(self, rpc_url: Optional[str] = None):
        self.w3 = Web3(Web3.HTTPProvider(rpc_url or "http://localhost:8545"))
        self.service_manager = self.w3.eth.contract(
            address=SERVICE_MANAGER_ADDRESS, abi=SERVICE_MANAGER_ABI
        )
        self.task_manager = self.w3.eth.contract(
            address=TASK_MANAGER_ADDRESS, abi=TASK_MANAGER_ABI
        )

    def is_connected(self) -> bool:
        return self.w3.is_connected()

    def get_chain_id(self) -> int:
        return self.w3.eth.chain_id
