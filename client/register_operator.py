from web3 import AsyncWeb3
from eth_account import Account
from web3.utils.abi import get_abi_element_info
import asyncio
import json
from pathlib import Path


PROJECT_ROOT = Path(__file__).parent.parent


OPERATOR_PRIV_KEY = "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
RPC_URL = "http://localhost:8545"
SERVICE_MANAGER_ADDRESS = "0x7a2088a1bFc9d81c55368AE168C2C02570cB814F"


async def main():
    w3 = AsyncWeb3(AsyncWeb3.AsyncHTTPProvider(RPC_URL))

    contract_path = (
        PROJECT_ROOT / "contracts/out/SertnServiceManager.sol/SertnServiceManager.json"
    )
    with open(contract_path) as f:
        contract_data = json.load(f)
    abi = contract_data["abi"]

    contract = w3.eth.contract(address=SERVICE_MANAGER_ADDRESS, abi=abi)

    account = Account.from_key(OPERATOR_PRIV_KEY)
    w3.eth.default_account = account.address

    with open(PROJECT_ROOT / "contracts/deployments/core/31337.json") as f:
        core_deployment = json.load(f)

    model = {
        "modelName_": w3.to_bytes(text="gpt-4").ljust(32, b"\0"),
        "operator_": account.address,
        "verifier_": "0x0000000000000000000000000000000000000000",
        "benchData_": "",
        "inputs_": b"",
        "output_": b"",
        "maxBlocks_": 1000,
        "ethStrategies_": [core_deployment["addresses"]["strategyManager"]],
        "ethShares_": [10, 10],
        "baseFee_": 100,
        "maxSer_": 10000,
        "computeType_": w3.to_bytes(text="gpu").ljust(32, b"\0"),
        "proveOnResponse_": False,
        "available_": True,
    }

    compute_unit_names = [w3.to_bytes(text="gpu").ljust(32, b"\0")]
    compute_units = [10]

    model_tuple = (
        model["modelName_"],
        model["operator_"],
        model["verifier_"],
        model["benchData_"],
        model["inputs_"],
        model["output_"],
        model["maxBlocks_"],
        model["ethStrategies_"],
        model["ethShares_"],
        model["baseFee_"],
        model["maxSer_"],
        model["computeType_"],
        model["proveOnResponse_"],
        model["available_"],
    )

    model_data = w3.codec.encode(
        [
            "(bytes32,address,address,string,bytes,bytes,uint32,address[],uint256[],uint256,uint256,bytes32,bool,bool)[]",
            "bytes32[]",
            "uint8[]",
        ],
        [[model_tuple], compute_unit_names, compute_units],
    )

    tx = await contract.functions.registerOperator(
        account.address, SERVICE_MANAGER_ADDRESS, [0], model_data
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
    print(f"Operator registered in transaction: {receipt['transactionHash'].hex()}")


if __name__ == "__main__":
    asyncio.run(main())
