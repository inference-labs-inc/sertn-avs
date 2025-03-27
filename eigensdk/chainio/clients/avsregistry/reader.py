import logging
import math
from typing import Dict, List, Optional, Tuple

from eth_typing import Address
from web3 import Web3
from web3.contract.contract import Contract

from eigensdk._types import (
    OperatorPubkeys,
    OperatorStateRetrieverCheckSignaturesIndices,
    OperatorStateRetrieverOperator,
)
from eigensdk.chainio import utils
from eigensdk.crypto.bls.attestation import G1Point, G2Point

DEFAULT_QUERY_BLOCK_RANGE = 10_000


class AvsRegistryReader:
    def __init__(
        self,
        registry_coordinator_addr: Address,
        registry_coordinator: Contract,
        bls_apk_registry_addr: Address,
        bls_apk_registry: Contract,
        operator_state_retriever: Contract,
        stake_registry: Contract,
        logger: logging.Logger,
        eth_http_client: Web3,
    ):
        self.logger: logging.Logger = logger
        self.bls_apk_registry_addr: Address = bls_apk_registry_addr
        self.bls_apk_registry: Contract = bls_apk_registry
        self.registry_coordinator_addr: Address = registry_coordinator_addr
        self.registry_coordinator: Contract = registry_coordinator
        self.operator_state_retriever: Contract = operator_state_retriever
        self.stake_registry: Contract = stake_registry
        self.eth_http_client: Web3 = eth_http_client

    def get_quorum_count(self) -> int:
        return self.registry_coordinator.functions.quorumCount().call()

    def get_operators_stake_in_quorums_at_current_block(
        self, quorum_numbers: List[int]
    ) -> List[List[OperatorStateRetrieverOperator]]:
        cur_block = self.eth_http_client.eth.block_number
        if cur_block > math.pow(2, 32) - 1:
            raise ValueError(
                "Current block number is too large to be converted to uint32"
            )
        return self.get_operators_stake_in_quorums_at_block(quorum_numbers, cur_block)

    def get_operators_stake_in_quorums_at_block(
        self, quorum_numbers: List[int], block_number: int
    ) -> List[List[OperatorStateRetrieverOperator]]:
        operator_stakes = self.operator_state_retriever.functions.getOperatorState(
            registryCoordinator=self.registry_coordinator_addr,
            quorumNumbers=utils.nums_to_bytes(quorum_numbers),
            blockNumber=block_number,
        ).call()
        return [
            [
                OperatorStateRetrieverOperator(
                    operator=operator[0],
                    operator_id="0x" + operator[1].hex(),
                    stake=operator[2],
                )
                for operator in quorum
            ]
            for quorum in operator_stakes
        ]

    def get_operator_addrs_in_quorums_at_current_block(
        self, quorum_numbers: List[int]
    ) -> List[List[Address]]:
        cur_block = self.eth_http_client.eth.block_number
        if cur_block > math.pow(2, 32) - 1:
            raise ValueError(
                "Current block number is too large to be converted to uint32"
            )

        operator_stakes = self.operator_state_retriever.functions.getOperatorState(
            registryCoordinator=self.registry_coordinator_addr,
            quorumNumbers=utils.nums_to_bytes(quorum_numbers),
            blockNumber=cur_block,
        ).call()
        return [[operator[0] for operator in quorum] for quorum in operator_stakes]

    def get_operators_stake_in_quorums_of_operator_at_block(
        self, operator_id: bytes, block_number: int
    ) -> Tuple[List[int], List[List[OperatorStateRetrieverOperator]]]:
        quorum_bitmap, operator_stakes = (
            self.operator_state_retriever.functions.getOperatorState(
                registryCoordinator=self.registry_coordinator_addr,
                operatorId=operator_id,
                blockNumber=block_number,
            ).call()
        )
        quorums = utils.bitmap_to_quorum_ids(quorum_bitmap)
        operator_stakes = [
            [
                OperatorStateRetrieverOperator(
                    operator=operator[0],
                    operator_id="0x" + operator[1].hex(),
                    stake=operator[2],
                )
                for operator in quorum
            ]
            for quorum in operator_stakes
        ]
        return quorums, operator_stakes

    def get_operators_stake_in_quorums_of_operator_at_current_block(
        self, operator_id: bytes
    ) -> Tuple[List[int], List[List[OperatorStateRetrieverOperator]]]:
        cur_block = self.eth_http_client.eth.block_number
        if cur_block > math.pow(2, 32) - 1:
            raise ValueError(
                "Current block number is too large to be converted to uint32"
            )
        return self.get_operators_stake_in_quorums_of_operator_at_block(
            operator_id, cur_block
        )

    def get_operator_stake_in_quorums_of_operator_at_current_block(
        self, operator_id: bytes
    ) -> Dict[int, int]:
        quorum_bitmap = self.registry_coordinator.functions.getCurrentQuorumBitmap(
            operator_id
        ).call()
        quorums = utils.bitmap_to_quorum_ids(quorum_bitmap)
        quorum_stakes: Dict[int, int] = {}
        for quorum in quorums:
            stake = self.stake_registry.functions.getCurrentStake(
                operator_id, quorum
            ).call()
            quorum_stakes[quorum] = stake
        return quorum_stakes

    def get_check_signatures_indices(
        self,
        reference_block_number: int,
        quorum_numbers: List[int],
        non_signer_operator_ids: List[int],
    ) -> OperatorStateRetrieverCheckSignaturesIndices:
        non_signer_operator_ids_bytes = [
            operator_id.to_bytes(32, "big") for operator_id in non_signer_operator_ids
        ]
        check_signature_indices = (
            self.operator_state_retriever.functions.getCheckSignaturesIndices(
                self.registry_coordinator_addr,
                reference_block_number,
                utils.nums_to_bytes(quorum_numbers),
                non_signer_operator_ids_bytes,
            ).call()
        )
        return OperatorStateRetrieverCheckSignaturesIndices(
            non_signer_quorum_bitmap_indices=check_signature_indices[0],
            quorum_apk_indices=check_signature_indices[1],
            total_stake_indices=check_signature_indices[2],
            non_signer_stake_indices=check_signature_indices[3],
        )

    def get_operator_id(self, operator_address: Address) -> bytes:
        operator_id = self.registry_coordinator.functions.getOperatorId(
            operator_address
        ).call()
        return operator_id

    def get_operator_from_id(self, operator_id: bytes) -> Address:
        operator_address = self.registry_coordinator.functions.getOperatorFromId(
            operator_id
        ).call()
        return operator_address

    def is_operator_registered(self, operator_address: Address) -> bool:
        operator_status = self.registry_coordinator.functions.getOperatorStatus(
            operator_address
        ).call()
        registered_with_avs = operator_status == 1
        return registered_with_avs

    def query_existing_registered_operator_pubkeys(
        self,
        start_block: int = 0,
        stop_block: Optional[int] = None,
        block_range: int = DEFAULT_QUERY_BLOCK_RANGE,
    ) -> Tuple[List[Address], List[OperatorPubkeys], int]:
        if stop_block is None:
            stop_block = self.eth_http_client.eth.block_number

        operator_pubkeys: List[OperatorPubkeys] = []
        operator_addresses: List[Address] = []
        for i in range(start_block, stop_block + 1, block_range):
            to_block: int = min(i + block_range - 1, stop_block)
            pubkey_updates = (
                self.bls_apk_registry.events.NewPubkeyRegistration.create_filter(
                    from_block=i, to_block=to_block, argument_filters={}, topics=[]
                ).get_all_entries()
            )
            self.logger.debug(
                "avsRegistryChainReader.query_existing_registered_operator_pubkeys",
                extra={
                    "numTransactionLogs": len(pubkey_updates),
                    "fromBlock": i,
                    "toBlock": to_block,
                },
            )
            for update in pubkey_updates:
                operator_addr = update["args"]["operator"]
                pubkey_g1 = update["args"]["pubkeyG1"]
                pubkey_g2 = update["args"]["pubkeyG2"]
                operator_pubkeys.append(
                    OperatorPubkeys(
                        g1_pub_key=G1Point(pubkey_g1["X"], pubkey_g1["Y"]),
                        g2_pub_key=G2Point(*pubkey_g2["X"], *pubkey_g2["Y"]),
                    )
                )
                operator_addresses.append(operator_addr)
        return operator_addresses, operator_pubkeys, to_block

    def query_existing_registered_operator_sockets(
        self,
        start_block: int = 0,
        stop_block: Optional[int] = None,
        block_range: int = DEFAULT_QUERY_BLOCK_RANGE,
    ) -> Tuple[Dict[bytes, str], int]:
        if stop_block is None:
            stop_block = self.eth_http_client.eth.block_number

        operator_id_to_socket_map: Dict[bytes, str] = {}
        for i in range(start_block, stop_block + 1, block_range):
            to_block = min(i + block_range - 1, stop_block)
            socket_updates = (
                self.registry_coordinator.events.OperatorSocketUpdate.create_filter(
                    from_block=i, to_block=to_block, argument_filters={}, topics=[]
                ).get_all_entries()
            )
            num_socket_updates = 0
            for update in socket_updates:
                operator_id_to_socket_map[update["args"]["operatorId"]] = update[
                    "args"
                ]["socket"]
                num_socket_updates += 1
            self.logger.debug(
                "avsRegistryChainReader.query_existing_registered_operator_sockets",
                extra={
                    "numTransactionLogs": num_socket_updates,
                    "fromBlock": i,
                    "toBlock": to_block,
                },
            )
        return operator_id_to_socket_map, to_block
