from web3 import Web3
from typing import Optional


class EthereumClient:
    def __init__(self, rpc_url: Optional[str] = None):
        self.w3 = Web3(Web3.HTTPProvider(rpc_url or "http://localhost:8545"))

    def is_connected(self) -> bool:
        return self.w3.is_connected()

    def get_chain_id(self) -> int:
        return self.w3.eth.chain_id
