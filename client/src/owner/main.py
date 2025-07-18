from eth_account import Account
from eth_account.datastructures import SignedTransaction

from common.eth import EthereumClient


class AvsOwner:
    def __init__(self, private_key: str, eth_rpc_url: str):
        self.eth_client = EthereumClient(rpc_url=eth_rpc_url)
        self.private_key = private_key
        self.owner_address = Account.from_key(self.private_key).address

    def submit_rewards_for_interval(
        self, current_interval: int, gas_limit: int = 2000000, gas_price_gwei: int = 20
    ):
        """
        Submit rewards for the current interval.
        This function is called by the owner of the aggregator.
        """
        tx = self.eth_client.service_manager.functions.submitRewardsForInterval(
            current_interval
        ).build_transaction(
            {
                "from": self.owner_address,
                "gas": gas_limit,
                "gasPrice": self.eth_client.w3.to_wei(str(gas_price_gwei), "gwei"),
                "nonce": self.eth_client.w3.eth.get_transaction_count(
                    self.owner_address
                ),
                "chainId": self.eth_client.w3.eth.chain_id,
            }
        )
        signed_tx: SignedTransaction = self.eth_client.w3.eth.account.sign_transaction(
            tx, private_key=self.private_key
        )
        tx_hash = self.eth_client.w3.eth.send_raw_transaction(signed_tx.raw_transaction)
        receipt = self.eth_client.w3.eth.wait_for_transaction_receipt(tx_hash)
        if receipt.status != 1:
            raise Exception(
                f"Failed to submit rewards for interval {current_interval}. Transaction hash: {tx_hash.hex()}"
            )
