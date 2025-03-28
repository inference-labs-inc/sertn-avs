from web3 import AsyncWeb3
from eth_account import Account
from web3.utils.abi import get_abi_element_info
import asyncio
import json
from pathlib import Path


PROJECT_ROOT = Path(__file__).parent.parent


AGGREGATOR_PRIV_KEY = (
    "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
)
RPC_URL = "http://localhost:8545"
SERVICE_MANAGER_ADDRESS = AsyncWeb3.to_checksum_address(
    "0x51fb710F469415FB6A539af8ab6757Ac6Aff0292"
)


async def main():
    w3 = AsyncWeb3(AsyncWeb3.AsyncHTTPProvider(RPC_URL))

    contract_path = (
        PROJECT_ROOT / "contracts/out/SertnServiceManager.sol/SertnServiceManager.json"
    )
    with open(contract_path) as f:
        contract_data = json.load(f)
    abi = contract_data["abi"]

    contract = w3.eth.contract(address=SERVICE_MANAGER_ADDRESS, abi=abi)

    account = Account.from_key(AGGREGATOR_PRIV_KEY)
    w3.eth.default_account = account.address

    with open(PROJECT_ROOT / "contracts/deployments/core/31337.json") as f:
        core_deployment = json.load(f)

    tx = await contract.functions.addAggregator(account.address).build_transaction(
        {
            "from": account.address,
            "gas": 2000000,
            "gasPrice": await w3.eth.gas_price,
            "nonce": await w3.eth.get_transaction_count(account.address),
            "chainId": await w3.eth.chain_id,
        }
    )

    signed_tx = w3.eth.account.sign_transaction(tx, AGGREGATOR_PRIV_KEY)
    tx_hash = await w3.eth.send_raw_transaction(signed_tx.raw_transaction)
    receipt = await w3.eth.wait_for_transaction_receipt(tx_hash)
    print(f"Aggregator registered in transaction: {receipt['transactionHash'].hex()}")


if __name__ == "__main__":
    asyncio.run(main())
