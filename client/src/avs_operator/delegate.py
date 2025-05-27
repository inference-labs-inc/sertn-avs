import os
import secrets
import time

import yaml
from eth_account import Account
from eth_account.messages import encode_typed_data
from eth_account.datastructures import SignedMessage, SignedTransaction
from hexbytes import HexBytes
from web3.types import TxReceipt

from common.abis import STRATEGY_ABI, ERC20_MOCK_ABI
from common.constants import ETH_STRATEGY_ADDRESSES, CLIENT_PATH
from common.console import console, styles
from common.eth import EthereumClient, load_ecdsa_private_key


def delegate():
    """
    Delegate tokens to the delegation manager
    Only for testing purposes. This function should not be used in production.
    This function mints 1000 AVS tokens for the operator and delegates them to the delegation manager.
    """
    amount = 1 * 10**9  # 1 AVS token
    console.print(
        f"Init...",
        style=styles.op_info,
    )

    # Load config
    config_file = os.path.join(CLIENT_PATH, "configs", "operator.yaml")
    with open(config_file, "r") as f:
        config = yaml.load(f, Loader=yaml.BaseLoader)

    # Load ECDSA private key
    eth_client = EthereumClient(rpc_url=config["eth_rpc_url"])
    private_key: str = load_ecdsa_private_key(
        keystore_path=os.path.join(CLIENT_PATH, config["ecdsa_private_key_store_path"]),
        password=os.environ.get("OPERATOR_ECDSA_KEY_PASSWORD", ""),
    )
    operator_address: str = Account.from_key(private_key).address
    console.print(f"Operator address: {operator_address}", style=styles.op_info)

    # Check ETH balance (just for being sure the operator is in the network)
    eth_balance: int = eth_client.w3.eth.get_balance(operator_address)
    eth_balance_in_ether = eth_client.w3.from_wei(eth_balance, "ether")
    console.print(f"ETH Balance: {eth_balance_in_ether} ETH", style=styles.op_info)

    # Approve tokens
    # Get the strategy contract and underlying token address
    strategy_contract = eth_client.w3.eth.contract(
        address=ETH_STRATEGY_ADDRESSES[0],
        abi=STRATEGY_ABI,
    )
    token_address = strategy_contract.functions.underlyingToken().call()
    staking_token = eth_client.w3.eth.contract(
        address=token_address,
        abi=ERC20_MOCK_ABI,
    )

    # Mint tokens for the operator.
    console.print("Minting tokens for the operator...", style=styles.op_info)
    tx = staking_token.functions.mint(operator_address, amount).build_transaction(
        {
            "from": operator_address,
            "gas": 200000,
            "gasPrice": eth_client.w3.to_wei("20", "gwei"),
            "nonce": eth_client.w3.eth.get_transaction_count(operator_address),
        }
    )
    signed_tx: SignedTransaction = eth_client.w3.eth.account.sign_transaction(
        tx, private_key
    )
    eth_client.w3.eth.send_raw_transaction(signed_tx.raw_transaction)
    eth_client.check_transaction_success(signed_tx)
    console.print(
        f"Minted {amount / 10**18} AVS tokens for the operator", style=styles.op_info
    )

    # Approve the delegation manager to spend the tokens
    console.print("Approving tokens for staking...", style=styles.op_info)
    tx = staking_token.functions.approve(
        eth_client.strategy_manager.address, amount
    ).build_transaction(
        {
            "from": operator_address,
            "gas": 200000,
            "gasPrice": eth_client.w3.to_wei("20", "gwei"),
            "nonce": eth_client.w3.eth.get_transaction_count(operator_address),
        }
    )
    signed_tx: SignedTransaction = eth_client.w3.eth.account.sign_transaction(
        tx, private_key
    )
    eth_client.w3.eth.send_raw_transaction(signed_tx.raw_transaction)
    eth_client.check_transaction_success(signed_tx)

    # Deposit tokens into the strategy
    console.print("Depositing tokens into the strategy...", style=styles.op_info)
    tx = eth_client.strategy_manager.functions.depositIntoStrategy(
        strategy_contract.address, token_address, amount
    ).build_transaction(
        {
            "from": operator_address,
            "gas": 200000,
            "gasPrice": eth_client.w3.to_wei("20", "gwei"),
            "nonce": eth_client.w3.eth.get_transaction_count(operator_address),
        }
    )
    signed_tx: SignedTransaction = eth_client.w3.eth.account.sign_transaction(
        tx, private_key
    )
    eth_client.w3.eth.send_raw_transaction(signed_tx.raw_transaction)
    eth_client.check_transaction_success(signed_tx)

    # create a "Delegate approval digest" for the delegation manager
    console.print("Creating a Delegate approval digest...", style=styles.op_info)
    signature, expiry, approver_salt = create_eip712_signature(
        operator_address=operator_address,
        private_key=private_key,
        eth_client=eth_client,
    )

    # Delegate tokens to the delegation manager
    console.print(
        "Delegating tokens to the delegation manager...", style=styles.op_info
    )

    tx = eth_client.delegation_manager.functions.delegateTo(
        operator_address,
        (signature, expiry),
        approver_salt,
    ).build_transaction(
        {
            "from": operator_address,
            "gas": 200000,
            "gasPrice": eth_client.w3.to_wei("20", "gwei"),
            "nonce": eth_client.w3.eth.get_transaction_count(operator_address),
        }
    )
    signed_tx: SignedTransaction = eth_client.w3.eth.account.sign_transaction(
        tx, private_key
    )
    eth_client.w3.eth.send_raw_transaction(signed_tx.raw_transaction)
    eth_client.check_transaction_success(signed_tx)
    console.print(
        f"Successfully Delegated {amount / 10**18} AVS tokens to the delegation manager",
        style=styles.op_info,
    )

    # Wait for the transaction to be mined
    receipt: TxReceipt = eth_client.w3.eth.wait_for_transaction_receipt(signed_tx.hash)

    if receipt.status == 1:
        # console.print("Transaction mined successfully.", style=styles.op_info)
        pass
    else:
        console.print(f"Transaction failed. {receipt}", style=styles.error)
        try:
            eth_client.w3.eth.call(signed_tx)
        except Exception as e:
            console.print(f"Transaction call failed: {e}", style=styles.error)


def create_eip712_signature(
    # chain_id: int,
    # delegator_address: str,
    operator_address: str,
    # delegation_manager_address: str,
    private_key: str,
    eth_client: EthereumClient,
) -> tuple[str, int, str]:
    """
    Create a EIP712 signature for the delegation manager
    """
    approver_salt = HexBytes(secrets.token_bytes(32))
    expiry = int(time.time()) + 3600  # valid for 1 hour

    domain = {
        "name": "DelegationManager",
        "version": "1",
        "chainId": eth_client.get_chain_id(),  # or whatever your local chain ID is
        "verifyingContract": eth_client.delegation_manager.address,
    }

    typed_message = {
        "types": {
            "EIP712Domain": [
                {"name": "name", "type": "string"},
                {"name": "version", "type": "string"},
                {"name": "chainId", "type": "uint256"},
                {"name": "verifyingContract", "type": "address"},
            ],
            "DelegationApproval": [
                {"name": "operator", "type": "address"},
                {"name": "approver", "type": "address"},
                {"name": "approverSalt", "type": "bytes32"},
                {"name": "expiry", "type": "uint256"},
            ],
        },
        "primaryType": "DelegationApproval",
        "domain": domain,
        "message": {
            "operator": operator_address,
            "approver": operator_address,
            "approverSalt": approver_salt.to_0x_hex(),
            "expiry": expiry,
        },
    }

    # Sign the message
    message = encode_typed_data(full_message=typed_message)
    signature: SignedMessage = Account.sign_message(
        message,
        private_key=private_key,
    )
    return signature.signature.to_0x_hex(), expiry, approver_salt.to_0x_hex()


if __name__ == "__main__":
    delegate()
