from web3 import Web3
import json
from pathlib import Path

PROJECT_ROOT = Path(__file__).parent.parent

contract_path = (
    PROJECT_ROOT / "contracts/out/SertnServiceManager.sol/SertnServiceManager.json"
)
with open(contract_path) as f:
    contract_data = json.load(f)
abi = contract_data["abi"]

w3 = Web3(Web3.HTTPProvider("http://localhost:8545"))
contract = w3.eth.contract(
    address="0x51fb710F469415FB6A539af8ab6757Ac6Aff0292", abi=abi
)

operators = contract.functions.getOperators().call()
aggregators = contract.functions.getAggregators().call()

print("Operators:", operators)
print("Aggregators:", aggregators)
