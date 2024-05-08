import ezkl
import os

vk_path = os.path.join("./model_data/test.vk")
settings_path = os.path.join("./model_data/settings.json")

sol_code_path = os.path.join('../contracts/src/ZKVerifier.sol')
abi_path = os.path.join('../contracts/abi/ZKVerifier.abi.json')

print ("INFO: Generating solidity contract")

res = ezkl.create_evm_verifier(
        vk_path,
        settings_path,
        sol_code_path,
        abi_path
    )

print("SUCCESS: Solidity contract generated at", sol_code_path)
print("SUCCESS: Solidity contract abi now at", abi_path)
