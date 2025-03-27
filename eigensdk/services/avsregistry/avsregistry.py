import logging
from typing import Dict, List, Union

from eigensdk._types import OperatorAvsState, OperatorInfo
from eigensdk.chainio.clients.avsregistry.reader import AvsRegistryReader
from eigensdk.crypto.bls.attestation import G1Point, new_zero_g1_point
from eigensdk.services.operatorsinfo.operatorsinfo_inmemory import (
    OperatorsInfoServiceInMemory,
)


class AvsRegistryService:
    def __init__(
        self,
        avs_registry_reader: AvsRegistryReader,
        operator_info_service: OperatorsInfoServiceInMemory,
        logger: logging.Logger,
    ):
        self.avs_registry_reader: AvsRegistryReader = avs_registry_reader
        self.operator_info_service: OperatorsInfoServiceInMemory = operator_info_service
        self.logger: logging.Logger = logger

    def get_operators_avs_state_at_block(
        self, quorum_numbers: List[int], block_number: int
    ) -> Dict[bytes, OperatorAvsState]:
        operators_avs_state: Dict[bytes, OperatorAvsState] = {}

        operators_stakes_in_quorums = (
            self.avs_registry_reader.get_operators_stake_in_quorums_at_block(
                quorum_numbers, block_number
            )
        )
        num_quorums = len(quorum_numbers)
        if len(operators_stakes_in_quorums) != num_quorums:
            self.logger.error(
                "Number of quorums returned from get_operators_stake_in_quorums_at_block does not match number of quorums requested. Probably pointing to old contract or wrong implementation.",
                extra={"service": "AvsRegistryServiceChainCaller"},
            )

        for quorum_idx, quorum_num in enumerate(quorum_numbers):
            for operator in operators_stakes_in_quorums[quorum_idx]:
                try:
                    info = self.get_operator_info(operator.operator_id)
                except:
                    self.logger.error(f"Operator {operator.operator_id} info not found. The operator is skipped.")
                    continue

                if operator.operator_id not in operators_avs_state:
                    operators_avs_state[operator.operator_id] = OperatorAvsState(
                        operator_id=operator.operator_id,
                        operator_info=info,
                        stake_per_quorum={},
                        block_number=block_number,
                    )

                operators_avs_state[operator.operator_id].stake_per_quorum[
                    quorum_num
                ] = operator.stake
        return operators_avs_state

    def get_quorums_avs_state_at_block(
        self, quorum_numbers: List[int], block_number: int
    ) -> Dict[int, Dict[str, Union[int, G1Point]]]:
        operators_avs_state = self.get_operators_avs_state_at_block(
            quorum_numbers, block_number
        )

        quorums_avs_state: Dict[int, Dict[str, Union[int, G1Point]]] = {}
        for quorum_num in quorum_numbers:
            agg_pubkey_g1 = new_zero_g1_point()
            total_stake = 0

            for operator_state in operators_avs_state.values():
                if quorum_num in operator_state.stake_per_quorum:
                    agg_pubkey_g1 = agg_pubkey_g1.add(
                        operator_state.operator_info.pub_keys.g1_pub_key
                    )
                    stake = operator_state.stake_per_quorum[quorum_num]
                    total_stake += stake

            quorums_avs_state[quorum_num] = {
                "quorum_number": quorum_num,
                "agg_pubkey_g1": agg_pubkey_g1,
                "total_stake": total_stake,
                "block_number": block_number,
            }

        return quorums_avs_state

    def get_operator_info(self, operator_id: bytes) -> OperatorInfo:
        operator_addr = self.avs_registry_reader.get_operator_from_id(operator_id)
        return self.operator_info_service.get_operator_info(operator_addr)
