from web3 import AsyncWeb3
from eth_account import Account
import asyncio
import json
from pathlib import Path


PROJECT_ROOT = Path(__file__).parent.parent


OPERATOR_PRIV_KEY = "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
RPC_URL = "http://localhost:8545"


async def main():
    w3 = AsyncWeb3(AsyncWeb3.AsyncHTTPProvider(RPC_URL))

    with open(PROJECT_ROOT / "contracts/deployments/core/31337.json") as f:
        core_deployment = json.load(f)

    with open(
        PROJECT_ROOT / "contracts/out/DelegationManager.sol/DelegationManager.json"
    ) as f:
        delegation_manager_abi = json.load(f)["abi"]

    delegation_address = w3.to_checksum_address(
        core_deployment["addresses"]["delegation"]
    )
    delegation_manager = w3.eth.contract(
        address=delegation_address, abi=delegation_manager_abi
    )

    account = Account.from_key(OPERATOR_PRIV_KEY)
    w3.eth.default_account = account.address

    tx = await delegation_manager.functions.registerAsOperator(
        "0x0000000000000000000000000000000000000000",
        0,
        "",
    ).build_transaction(
        {
            "from": account.address,
            "gas": 2000000,
            "gasPrice": await w3.eth.gas_price,
            "nonce": await w3.eth.get_transaction_count(account.address),
            "chainId": await w3.eth.chain_id,
        }
    )

    signed_tx = w3.eth.account.sign_transaction(tx, OPERATOR_PRIV_KEY)
    tx_hash = await w3.eth.send_raw_transaction(signed_tx.raw_transaction)
    receipt = await w3.eth.wait_for_transaction_receipt(tx_hash)
    print(
        f"Registered as operator in EigenLayer in transaction: {receipt['transactionHash'].hex()}"
    )


if __name__ == "__main__":
    asyncio.run(main())
