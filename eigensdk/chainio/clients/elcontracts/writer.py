import logging
from typing import Any

from eth_account.signers.local import LocalAccount
from eth_typing import Address
from web3 import Web3
from web3.contract.contract import Contract
from web3.types import TxReceipt

from eigensdk._types import Operator

from ...utils import send_transaction
from .reader import ELReader


class ELWriter:
    def __init__(
        self,
        slasher: Contract,
        delegation_manager: Contract,
        strategy_manager: Contract,
        strategy_manager_addr: Address,
        avs_directory: Contract,
        el_reader: ELReader,
        logger: logging.Logger,
        eth_http_client: Web3,
        pk_wallet: LocalAccount,
    ):
        self.slasher: Contract = slasher
        self.delegation_manager: Contract = delegation_manager
        self.strategy_manager: Contract = strategy_manager
        self.strategy_manager_addr: Address = strategy_manager_addr
        self.avs_directory: Contract = avs_directory
        self.el_reader: ELReader = el_reader
        self.logger: logging.Logger = logger
        self.eth_http_client: Web3 = eth_http_client
        self.pk_wallet: Any = pk_wallet

    def register_as_operator(self, operator: Operator) -> TxReceipt:
        self.logger.info(f"Registering operator {operator.address} to EigenLayer")
        op_details = {
            "earningsReceiver": Web3.to_checksum_address(
                operator.earnings_receiver_address
            ),
            "stakerOptOutWindowBlocks": operator.staker_opt_out_window_blocks,
            "delegationApprover": Web3.to_checksum_address(
                operator.delegation_approver_address
            ),
        }
        func = self.delegation_manager.functions.registerAsOperator(
            op_details, operator.metadata_url
        )
        receipt = send_transaction(func, self.pk_wallet, self.eth_http_client)

        self.logger.info(
            "Transaction successfully included",
            extra={"txHash": receipt["transactionHash"].hex()},
        )
        return receipt

    def update_operator_details(self, operator: Operator) -> TxReceipt:
        self.logger.info(
            f"Updating operator details of operator {operator.address} to EigenLayer"
        )
        op_details = {
            "earningsReceiver": Web3.to_checksum_address(
                operator.earnings_receiver_address
            ),
            "delegationApprover": Web3.to_checksum_address(
                operator.delegation_approver_address
            ),
            "stakerOptOutWindowBlocks": operator.staker_opt_out_window_blocks,
        }
        func = self.delegation_manager.functions.modifyOperatorDetails(op_details)
        try:
            receipt = send_transaction(func, self.pk_wallet, self.eth_http_client)
        except Exception as e:
            self.logger.error(e)
            return None

        self.logger.info(
            "Successfully updated operator details",
            extra={
                "txHash": receipt["transactionHash"].hex(),
                "operator": operator.address,
            },
        )

        func = self.delegation_manager.functions.updateOperatorMetadataURI(
            operator.metadata_url
        )
        receipt = send_transaction(func, self.pk_wallet, self.eth_http_client)

        self.logger.info(
            "Successfully updated operator metadata URI",
            extra={
                "txHash": receipt["transactionHash"].hex(),
                "operator": operator.address,
            },
        )
        return receipt

    def deposit_erc20_into_strategy(
        self, strategy_addr: Address, amount: int
    ) -> TxReceipt:
        self.logger.info(f"Depositing {amount} tokens into strategy {strategy_addr}")

        _, underlying_token_contract, underlying_token_addr = (
            self.el_reader.get_strategy_and_underlying_erc20_token(strategy_addr)
        )

        approve_func = underlying_token_contract.functions.approve(
            self.strategy_manager_addr, amount
        )
        try:
            send_transaction(approve_func, self.pk_wallet, self.eth_http_client)
        except Exception as e:
            self.logger.error(e)
            return None

        deposit_func = self.strategy_manager.functions.depositIntoStrategy(
            strategy_addr, underlying_token_addr, amount
        )
        receipt = send_transaction(deposit_func, self.pk_wallet, self.eth_http_client)

        self.logger.info(
            "Successfully deposited the token into the strategy",
            extra={
                "txHash": receipt["transactionHash"].hex(),
                "strategy": strategy_addr,
                "token": underlying_token_addr,
                "amount": amount,
            },
        )
        return receipt
