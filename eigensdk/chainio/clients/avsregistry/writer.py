import logging
from typing import List

from eth_account import Account
from eth_account.signers.local import LocalAccount
from eth_typing import Address
from web3 import Web3
from web3.contract.contract import Contract
from web3.types import TxReceipt

from eigensdk.chainio import utils
from eigensdk.crypto.bls.attestation import G1Point, KeyPair

from ...utils import send_transaction
from ..elcontracts.reader import ELReader


class AvsRegistryWriter:
    def __init__(
        self,
        service_manager_addr: Address,
        registry_coordinator: Contract,
        operator_state_retriever: Contract,
        stake_registry: Contract,
        bls_apk_registry: Contract,
        el_reader: ELReader,
        logger: logging.Logger,
        eth_http_client: Web3,
        pk_wallet: LocalAccount,
    ):
        self.service_manager_addr: Address = service_manager_addr
        self.registry_coordinator: Contract = registry_coordinator
        self.operator_state_retriever: Contract = operator_state_retriever
        self.stake_registry: Contract = stake_registry
        self.bls_apk_registry: Contract = bls_apk_registry
        self.el_reader: ELReader = el_reader
        self.logger: logging.Logger = logger
        self.eth_http_client: Web3 = eth_http_client
        self.pk_wallet: LocalAccount = pk_wallet

    def register_operator_in_quorum_with_avs_registry_coordinator(
        self,
        operator_ecdsa_private_key: str,
        operator_to_avs_registration_sig_salt: bytes,
        operator_to_avs_registration_sig_expiry: int,
        bls_key_pair: KeyPair,
        quorum_numbers: List[int],
        socket: str,
    ) -> TxReceipt:
        account = Account.from_key(operator_ecdsa_private_key)
        operator_addr = account.address
        self.logger.info(
            "Registering operator with the AVS's registry coordinator",
            extra={
                "avs-service-manager": self.service_manager_addr,
                "operator": operator_addr,
                "quorumNumbers": quorum_numbers,
                "socket": socket,
            },
        )
        g1_hashed_msg_to_sign = (
            self.registry_coordinator.functions.pubkeyRegistrationMessageHash(
                operator_addr
            ).call()
        )
        signed_msg = bls_key_pair.sign_hashed_to_curve_message(
            G1Point(*g1_hashed_msg_to_sign)
        )

        pubkey_reg_params = (
            (
                int(signed_msg.getX().getStr()),
                int(signed_msg.getY().getStr()),
            ),
            (
                int(bls_key_pair.pub_g1.getX().getStr()),
                int(bls_key_pair.pub_g1.getY().getStr()),
            ),
            (
                (
                    int(bls_key_pair.pub_g2.getX().get_a().getStr()),
                    int(bls_key_pair.pub_g2.getX().get_b().getStr()),
                ),
                (
                    int(bls_key_pair.pub_g2.getY().get_a().getStr()),
                    int(bls_key_pair.pub_g2.getY().get_b().getStr()),
                ),
            ),
        )

        msg_to_sign = self.el_reader.calculate_operator_avs_registration_digest_hash(
            operator_addr,
            self.service_manager_addr,
            operator_to_avs_registration_sig_salt,
            operator_to_avs_registration_sig_expiry,
        )
        operator_signature = account.unsafe_sign_hash(msg_to_sign)["signature"]
        operator_signature_with_salt_and_expiry = (
            operator_signature,
            operator_to_avs_registration_sig_salt,
            operator_to_avs_registration_sig_expiry,
        )

        func = self.registry_coordinator.functions.registerOperator(
            utils.nums_to_bytes(quorum_numbers),
            socket,
            pubkey_reg_params,
            operator_signature_with_salt_and_expiry,
        )
        receipt = send_transaction(func, self.pk_wallet, self.eth_http_client)

        self.logger.info(
            "Successfully registered operator with AVS registry coordinator",
            extra={
                "txHash": receipt.transactionHash.hex(),
                "avs-service-manager": self.service_manager_addr,
                "operator": operator_addr,
                "quorumNumbers": quorum_numbers,
            },
        )
        return receipt

    def update_stakes_of_entire_operator_set_for_quorums(
        self,
        operators_per_quorum: List[List[Address]],
        quorum_numbers: List[int],
    ) -> TxReceipt:
        self.logger.info(
            "Updating stakes for entire operator set",
            extra={"quorumNumbers": quorum_numbers},
        )

        func = self.registry_coordinator.functions.updateOperatorsForQuorum(
            operators_per_quorum, utils.nums_to_bytes(quorum_numbers)
        )
        receipt = send_transaction(func, self.pk_wallet, self.eth_http_client)

        self.logger.info(
            "Successfully updated stakes for entire operator set",
            extra={
                "txHash": receipt.transactionHash.hex(),
                "quorumNumbers": quorum_numbers,
            },
        )
        return receipt

    def update_stakes_of_operator_subset_for_all_quorums(
        self, operators: List[Address]
    ) -> TxReceipt:
        self.logger.info(
            "Updating stakes of operator subset for all quorums",
            extra={"operators": operators},
        )

        func = self.registry_coordinator.functions.updateOperators(operators)
        receipt = send_transaction(func, self.pk_wallet, self.eth_http_client)

        self.logger.info(
            "Successfully updated stakes of operator subset for all quorums",
            extra={
                "txHash": receipt.transactionHash.hex(),
                "operators": operators,
            },
        )
        return receipt

    def deregister_operator(self, quorum_numbers: List[int]) -> TxReceipt:
        self.logger.info("Deregistering operator with the AVS's registry coordinator")

        func = self.registry_coordinator.functions.deregisterOperator(
            utils.nums_to_bytes(quorum_numbers)
        )
        receipt = send_transaction(func, self.pk_wallet, self.eth_http_client)

        self.logger.info(
            "Successfully deregistered operator with the AVS's registry coordinator",
            extra={"txHash": receipt.transactionHash.hex()},
        )
        return receipt

    def update_socket(self, socket: str) -> TxReceipt:
        self.logger.info(
            "Updating socket",
            extra={"socket": socket},
        )
        func = self.registry_coordinator.functions.updateSocket(socket)
        receipt = send_transaction(func, self.pk_wallet, self.eth_http_client)

        self.logger.info(
            "Successfully updated socket",
            extra={"txHash": receipt.transactionHash.hex()},
        )
        return receipt
