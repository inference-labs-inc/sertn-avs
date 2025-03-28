from web3 import Web3
from eth_account import Account
from pathlib import Path
import json


PROJECT_ROOT = Path(__file__).parent.parent


RPC_URL = "http://localhost:8545"
SERVICE_MANAGER_ADDRESS = Web3.to_checksum_address(
    "0x3c1cb427d20f15563ada8c249e71db76d7183b6c"
)
OPERATOR_PRIV_KEY = "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"


def main():

    w3 = Web3(Web3.HTTPProvider(RPC_URL))
    account = Account.from_key(OPERATOR_PRIV_KEY)

    contract_path = (
        PROJECT_ROOT
        / "contracts"
        / "out"
        / "SertnServiceManager.sol"
        / "SertnServiceManager.json"
    )
    with open(contract_path) as f:
        contract_json = json.load(f)
        contract = w3.eth.contract(
            address=SERVICE_MANAGER_ADDRESS, abi=contract_json["abi"]
        )

    operator_manager = contract.functions.operatorManager().call()
    model_manager = contract.functions.modelManager().call()

    operator_manager_path = (
        PROJECT_ROOT
        / "contracts"
        / "out"
        / "SertnOperatorManager.sol"
        / "SertnOperatorManager.json"
    )
    with open(operator_manager_path) as f:
        operator_manager_json = json.load(f)
        operator_manager_contract = w3.eth.contract(
            address=operator_manager, abi=operator_manager_json["abi"]
        )

    model_manager_path = (
        PROJECT_ROOT
        / "contracts"
        / "out"
        / "SertnModelManager.sol"
        / "SertnModelManager.json"
    )
    with open(model_manager_path) as f:
        model_manager_json = json.load(f)
        model_manager_contract = w3.eth.contract(
            address=model_manager, abi=model_manager_json["abi"]
        )

    models = [1]
    compute_types = [w3.to_bytes(text="gpu").ljust(32, b"\0")]
    proof_request_exponents = [1, 1]

    tx = operator_manager_contract.functions.registerOperator(
        account.address, models, compute_types, proof_request_exponents
    ).build_transaction(
        {
            "from": account.address,
            "nonce": w3.eth.get_transaction_count(account.address),
            "gas": 500000,
            "gasPrice": w3.eth.gas_price,
        }
    )

    signed = w3.eth.account.sign_transaction(tx, OPERATOR_PRIV_KEY)
    tx_hash = w3.eth.send_raw_transaction(signed.raw_transaction)
    print(f"Registering operator... tx hash: {tx_hash.hex()}")
    w3.eth.wait_for_transaction_receipt(tx_hash)

    model = (
        w3.to_bytes(text="test-model").ljust(32, b"\0"),
        account.address,
        account.address,
        "test-bench-data",
        b"",
        b"",
        1000,
        [],
        [],
        100,
        1000,
        w3.to_bytes(text="gpu").ljust(32, b"\0"),
        False,
        True,
    )

    tx = model_manager_contract.functions.registerModel(model).build_transaction(
        {
            "from": account.address,
            "nonce": w3.eth.get_transaction_count(account.address),
            "gas": 500000,
            "gasPrice": w3.eth.gas_price,
        }
    )

    signed = w3.eth.account.sign_transaction(tx, OPERATOR_PRIV_KEY)
    tx_hash = w3.eth.send_raw_transaction(signed.raw_transaction)
    print(f"Registering model... tx hash: {tx_hash.hex()}")
    w3.eth.wait_for_transaction_receipt(tx_hash)


if __name__ == "__main__":
    main()
