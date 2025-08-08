from eth_account import Account
from eth_account.datastructures import SignedTransaction

from common.eth import EthereumClient
from common.config import GasStrategy


class AvsOwner:
    def __init__(self, private_key: str, eth_rpc_url: str):
        self.eth_client = EthereumClient(
            eth_rpc_url=eth_rpc_url, gas_strategy=GasStrategy.STANDARD.value
        )
        self.private_key = private_key
        self.owner_address = Account.from_key(self.private_key).address

    def submit_rewards_for_interval(self, current_interval: int, **gas_kwargs):
        """
        Submit rewards for the current interval.
        This function is called by the owner of the aggregator.

        :param current_interval: The interval for which to submit rewards
        :param gas_kwargs: Optional gas parameters (gas_limit, max_fee_per_gas, etc.)
        """
        self.eth_client.execute_transaction(
            self.eth_client.service_manager,
            "submitRewardsForInterval",
            self.private_key,
            [current_interval],
            **gas_kwargs
        )
