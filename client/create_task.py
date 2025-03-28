from web3 import Web3
from eth_account import Account
from pathlib import Path
import json

RPC_URL = "http://localhost:8545"
PRIVATE_KEY = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
SERVICE_MANAGER_ADDRESS = "0x5FbDB2315678afecb367f032d93F642f64180aa3"


def create_task():
    w3 = Web3(Web3.HTTPProvider(RPC_URL))
    account = Account.from_key(PRIVATE_KEY)

    ser_token_path = Path("contracts/out/ERC20Mock.sol/ERC20Mock.json")
    with open(ser_token_path) as f:
        ser_token_json = json.load(f)
        ser_token = w3.eth.contract(
            address="0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512",
            abi=ser_token_json["abi"],
        )

    contract_path = Path(
        "contracts/out/SertnServiceManager.sol/SertnServiceManager.json"
    )
    with open(contract_path) as f:
        contract_json = json.load(f)
        contract = w3.eth.contract(
            address=SERVICE_MANAGER_ADDRESS, abi=contract_json["abi"]
        )

    gas_price = w3.eth.gas_price
    nonce = w3.eth.get_transaction_count(account.address)

    mint_tx = ser_token.functions.mint(account.address, int(1e6)).build_transaction(
        {"from": account.address, "nonce": nonce, "gas": 500000, "gasPrice": gas_price}
    )

    signed = w3.eth.account.sign_transaction(mint_tx, PRIVATE_KEY)
    tx_hash = w3.eth.send_raw_transaction(signed.raw_transaction)
    w3.eth.wait_for_transaction_receipt(tx_hash)
    print(f"Minted SER tokens with transaction hash: {tx_hash.hex()}")

    amount = int(1.5e3 * (1 + 100) / 1e3)
    nonce = w3.eth.get_transaction_count(account.address)

    approve_tx = ser_token.functions.approve(
        SERVICE_MANAGER_ADDRESS, amount
    ).build_transaction(
        {"from": account.address, "nonce": nonce, "gas": 500000, "gasPrice": gas_price}
    )

    signed = w3.eth.account.sign_transaction(approve_tx, PRIVATE_KEY)
    tx_hash = w3.eth.send_raw_transaction(signed.raw_transaction)
    w3.eth.wait_for_transaction_receipt(tx_hash)
    print(f"Approved SER tokens with transaction hash: {tx_hash.hex()}")

    task = {
        "modelId_": 1,
        "inputs_": b"",
        "poc_": int(1e2),
        "startTime_": 0,
        "startingBlock_": 0,
        "proveOnResponse_": False,
        "user_": account.address,
    }

    nonce = w3.eth.get_transaction_count(account.address)
    tx = contract.functions.sendTask(task).build_transaction(
        {"from": account.address, "nonce": nonce, "gas": 500000, "gasPrice": gas_price}
    )

    signed = w3.eth.account.sign_transaction(tx, PRIVATE_KEY)
    tx_hash = w3.eth.send_raw_transaction(signed.raw_transaction)
    receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
    print(f"Created task with transaction hash: {tx_hash.hex()}")

    task_events = contract.events.newTask().process_receipt(receipt)
    if task_events:
        task_id = task_events[0].args.taskId_
        print(f"Task ID: {task_id.hex()}")
    else:
        print("No task ID found in receipt")


if __name__ == "__main__":
    create_task()
