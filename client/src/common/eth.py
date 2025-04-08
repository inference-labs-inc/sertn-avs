import json
from typing import Optional

from eth_account import Account
from web3 import Web3

from common.abis import SERVICE_MANAGER_ABI, TASK_MANAGER_ABI
from common.constants import (
    SERVICE_MANAGER_ADDRESS,
    TASK_MANAGER_ADDRESS,
)
from console import console, styles


def load_ecdsa_private_key(keystore_path: str, password: str) -> str:
    """
    Load ECDSA private key from keystore file
    """
    if not password:
        console.print(
            "ECDSA key not set. using empty string.",
            style=styles.debug,
        )

    with open(keystore_path, "r") as f:
        keystore = json.load(f)
    return Account.decrypt(keystore, password).hex()


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
