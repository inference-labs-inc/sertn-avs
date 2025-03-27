from typing import List

from eth_account.signers.local import LocalAccount
from web3 import Web3
from web3.contract.contract import ContractFunction
# from web3.middleware.geth_poa import geth_poa_middleware
from web3.types import TxReceipt


def nums_to_bytes(nums: List[int]) -> bytes:
    return "".join(map(chr, nums)).encode()


def bitmap_to_quorum_ids(bitmap: int) -> List[int]:
    quorum_ids = []
    for i in range(256):
        if bitmap & (1 << i):
            quorum_ids.append(int(i))
    return quorum_ids


def send_transaction(
    func: ContractFunction, pk_wallet: LocalAccount, eth_http_client: Web3
) -> TxReceipt:
    # try:
    #     eth_http_client.middleware_onion.inject(geth_poa_middleware, layer=0)
    # except Exception:
    #     pass

    try:
        gas_estimate = func.estimate_gas({"from": pk_wallet.address})
    except Exception as e:
        raise Exception(f"Gas estimation failed: {e}")

    current_gas_price = eth_http_client.eth.gas_price

    tx = func.build_transaction(
        {
            "from": pk_wallet.address,
            "gas": gas_estimate,
            "gasPrice": current_gas_price,
            "nonce": eth_http_client.eth.get_transaction_count(pk_wallet.address),
            "chainId": eth_http_client.eth.chain_id,
        }
    )
    signed_tx = eth_http_client.eth.account.sign_transaction(
        tx, private_key=pk_wallet.key
    )
    tx_hash = eth_http_client.eth.send_raw_transaction(signed_tx.raw_transaction)
    receipt = eth_http_client.eth.wait_for_transaction_receipt(tx_hash)
    return receipt
