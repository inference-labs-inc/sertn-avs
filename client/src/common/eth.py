import json
from typing import Optional

from eth_account import Account
from eth_account.datastructures import SignedTransaction
from web3 import Web3
from web3.types import TxReceipt

from common.abis import (
    SERVICE_MANAGER_ABI,
    TASK_MANAGER_ABI,
    DELEGATION_MANAGER_ABI,
    STRATEGY_MANAGER_ABI,
    ALLOCATION_MANAGER_ABI,
    MODEL_REGISTRY_ABI,
    REWARDS_COORDINATOR_ABI,
)
from common.constants import (
    SERVICE_MANAGER_ADDRESS,
    TASK_MANAGER_ADDRESS,
    ALLOCATION_MANAGER_ADDRESS,
)
from common.console import console, styles


def load_ecdsa_private_key(keystore_path: str, password: str) -> str:
    """
    Load ECDSA private key from keystore file
    """
    if not password:
        console.print(
            "ECDSA key password is not set. using empty string.",
            style=styles.debug,
        )

    with open(keystore_path, "r") as f:
        keystore = json.load(f)
    return Account.decrypt(keystore, password).hex()


class EthereumClient:
    """
    Ethereum client wrapper
    XXX: Actually a classic piece of code for being a singleton.
    Maybe a good idea to refactor it later.
    """

    def __init__(self, rpc_url: Optional[str] = None):
        self.w3 = Web3(Web3.HTTPProvider(rpc_url or "http://localhost:8545"))
        if not self.is_connected():
            raise ConnectionError("Failed to connect to Ethereum node")
        self.service_manager = None
        self.task_manager = None
        self.delegation_manager = None
        self.strategy_manager = None
        self.init_contracts()

    def init_contracts(self) -> None:
        """
        Initialize contracts
        """
        # Service manager
        self.check_contract_deployed(SERVICE_MANAGER_ADDRESS)
        self.service_manager = self.w3.eth.contract(
            address=SERVICE_MANAGER_ADDRESS, abi=SERVICE_MANAGER_ABI
        )
        # Task manager
        self.check_contract_deployed(TASK_MANAGER_ADDRESS)
        self.task_manager = self.w3.eth.contract(
            address=TASK_MANAGER_ADDRESS, abi=TASK_MANAGER_ABI
        )
        # Delegation manager
        delegation_manager_address = (
            self.service_manager.functions.delegationManager().call()
        )
        self.check_contract_deployed(delegation_manager_address)
        self.delegation_manager = self.w3.eth.contract(
            address=delegation_manager_address,
            abi=DELEGATION_MANAGER_ABI,
        )
        # Strategy manager
        strategy_manager_address = (
            self.delegation_manager.functions.strategyManager().call()
        )
        self.check_contract_deployed(strategy_manager_address)
        self.strategy_manager = self.w3.eth.contract(
            address=strategy_manager_address,
            abi=STRATEGY_MANAGER_ABI,
        )
        # allocation manager
        self.check_contract_deployed(ALLOCATION_MANAGER_ADDRESS)
        self.allocation_manager = self.w3.eth.contract(
            address=ALLOCATION_MANAGER_ADDRESS,
            abi=ALLOCATION_MANAGER_ABI,
        )
        # Model registry
        self.model_registry = self.w3.eth.contract(
            address=self.task_manager.functions.modelRegistry().call(),
            abi=MODEL_REGISTRY_ABI,
        )
        self.check_contract_deployed(self.model_registry.address)
        # Rewards coordinator
        self.rewards_coordinator = self.w3.eth.contract(
            address=self.service_manager.functions.rewardsCoordinator().call(),
            abi=REWARDS_COORDINATOR_ABI,
        )
        self.check_contract_deployed(self.rewards_coordinator.address)

    def is_connected(self) -> bool:
        return self.w3.is_connected()

    def get_chain_id(self) -> int:
        return self.w3.eth.chain_id

    def check_contract_deployed(self, address: str) -> None:
        """
        Check if a contract is deployed at the given address.
        This is kinda optional, but useful for debugging.
        I'm tired by errors like "execution reverted: revert: Address: call to non-contract"

        :param address: Address to check
        :return: None
        """
        code = self.w3.eth.get_code(address)
        if code == b"0x":
            raise ValueError(f"Address {address} is not a contract. Code: {code}")

    def check_transaction_success(self, tx: SignedTransaction) -> None:
        """
        Check if a transaction was successful
        :param tx: Transaction to check
        :type tx: SignedTransaction
        :raises Exception: If the transaction call failed
        :return: None
        """
        receipt: TxReceipt = self.w3.eth.wait_for_transaction_receipt(
            transaction_hash=tx.hash,
            timeout=120,
            poll_latency=0.5,
        )
        if receipt.status != 1:
            console.print(
                f"Transaction {tx.hash.hex()} failed. Receipt: {receipt}",
                style=styles.error,
            )
            try:
                self.w3.eth.call(tx)
            except ValueError as e:
                console.print(f"Transaction call failed: {e}", style=styles.error)
                raise e
