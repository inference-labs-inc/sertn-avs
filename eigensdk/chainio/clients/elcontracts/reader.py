import logging
from typing import Tuple

from eth_typing import Address
from web3 import Web3
from web3.contract.contract import Contract

from eigensdk._types import Operator
from eigensdk.contracts import ABIs


class ELReader:
    def __init__(
        self,
        slasher: Contract,
        delegation_manager: Contract,
        strategy_manager: Contract,
        avs_directory: Contract,
        logger: logging.Logger,
        eth_http_client: Web3,
    ):
        self.slasher: Contract = slasher
        self.delegation_manager: Contract = delegation_manager
        self.strategy_manager: Contract = strategy_manager
        self.avs_directory: Contract = avs_directory
        self.eth_http_client: Web3 = eth_http_client
        self.logger: logging.Logger = logger

    def is_operator_registered(self, operator_addr: Address) -> bool:
        return self.delegation_manager.functions.isOperator(operator_addr).call()

    def get_operator_details(self, operator_addr: Address) -> Operator:
        operator_details = self.delegation_manager.functions.operatorDetails(
            operator_addr
        ).call()

        return Operator(
            address=operator_addr,
            earnings_receiver_address=self.eth_http_client.to_checksum_address(
                operator_details[0]
            ),
            staker_opt_out_window_blocks=operator_details[2],
            delegation_approver_address=self.eth_http_client.to_checksum_address(
                operator_details[1]
            ),
        )

    def get_strategy_and_underlying_token(
        self, strategy_addr: Address
    ) -> Tuple[Contract, str]:
        strategy: Contract = self.eth_http_client.eth.contract(
            address=strategy_addr, abi=ABIs.STRATEGY
        )
        underlying_token_addr = strategy.functions.underlyingToken().call()
        return strategy, underlying_token_addr

    def get_strategy_and_underlying_erc20_token(
        self, strategy_addr: Address
    ) -> Tuple[Contract, Contract, Address]:
        strategy: Contract = self.eth_http_client.eth.contract(
            address=strategy_addr, abi=ABIs.STRATEGY
        )
        underlying_token_addr = strategy.functions.underlyingToken().call()
        underlying_token: Contract = self.eth_http_client.eth.contract(
            address=underlying_token_addr, abi=ABIs.ERC20
        )
        return strategy, underlying_token, underlying_token_addr

    def service_manager_can_slash_operator_until_block(
        self, operator_addr: Address, service_manager_addr: Address
    ) -> int:
        return self.slasher.functions.contractCanSlashOperatorUntilBlock(
            operator_addr, service_manager_addr
        ).call()

    def operator_is_frozen(self, operator_addr: Address) -> bool:
        return self.slasher.functions.isFrozen(operator_addr).call()

    def get_operator_shares_in_strategy(
        self, operator_addr: Address, strategy_addr: Address
    ) -> int:
        return self.delegation_manager.functions.operatorShares(
            operator_addr, strategy_addr
        ).call()

    def calculate_delegation_approval_digest_hash(
        self,
        staker: Address,
        operator_addr: Address,
        delegation_approver: Address,
        approver_salt: bytes,
        expiry: int,
    ) -> bytes:
        return self.delegation_manager.functions.calculateDelegationApprovalDigestHash(
            staker, operator_addr, delegation_approver, approver_salt, expiry
        ).call()

    def calculate_operator_avs_registration_digest_hash(
        self, operator_addr: Address, avs: Address, salt: bytes, expiry: int
    ) -> bytes:
        return self.avs_directory.functions.calculateOperatorAVSRegistrationDigestHash(
            operator_addr, avs, salt, expiry
        ).call()
