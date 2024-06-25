import ezkl
import sys
import os

sys.path.append(os.path.join(os.path.dirname(__file__), "../"))

model_path = "models/model_1"

from utils import relative_file_path
vk_path = relative_file_path(model_path + "/test.vk")
settings_path = relative_file_path(model_path + "/settings.json")

sol_code_path = relative_file_path('../contracts/src/'+model_path+'/ZKVerifier.sol')
abi_path = relative_file_path('../contracts/abi/ZKVerifier.abi.json')

print ("INFO: Generating solidity contract")

res = ezkl.create_evm_verifier(
        vk_path,
        settings_path,
        sol_code_path,
        abi_path
    )

print("SUCCESS: Solidity contract generated at", sol_code_path)
print("SUCCESS: Solidity contract abi now at", abi_path)
