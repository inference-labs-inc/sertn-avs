from .avsregistry import AvsRegistryService as ARSInterface
from .avsregistry import (
    OperatorAvsState,
    OperatorInfo,
    OperatorPubkeys,
    QuorumAvsState,
    CallOpts,
    OperatorStateRetrieverCheckSignaturesIndices,
)

from dataclasses import dataclass
from eigensdk.crypto.bls.attestation import KeyPair, G1Point, G2Point
from eigensdk.crypto import bls


@dataclass
class TestOperator:
    operator_id: int
    stake_per_quorum: dict[int, int]
    bls_key_pair: KeyPair


class FakeAvsRegistryService(ARSInterface):
    operators: dict[int, dict[int, OperatorAvsState]]

    def __init__(self, block_number: int, operators: list[TestOperator]) -> None:
        self.operators = {block_number: {}}
        for operator in operators:
            self.operators[block_number][operator.operator_id] = OperatorAvsState(
                operator_id=operator.operator_id,
                operator_info=OperatorInfo(
                    socket="localhost:9090",
                    pub_keys=OperatorPubkeys(
                        g1_pub_key=operator.bls_key_pair.pub_g1,
                        g2_pub_key=operator.bls_key_pair.pub_g2,
                    ),
                ),
                stake_per_quorum=operator.stake_per_quorum,
                block_number=block_number,
            )

    async def get_operators_avs_state_at_block(
        self, quorum_numbers: list[int], block_number: int
    ) -> dict[int, OperatorAvsState]:
        if block_number not in self.operators:
            raise ValueError(f"No data for the block {block_number}")
        return self.operators[block_number]

    async def get_quorums_avs_state_at_block(
        self, quorum_numbers: list[int], block_number: int
    ) -> dict[int, QuorumAvsState]:
        if block_number not in self.operators:
            raise ValueError(f"No data for the block {block_number}")
        state = {}
        for qn in quorum_numbers:
            agg_pub_key_g1 = G1Point(0, 0)
            total_stake = 0
            for operator_id, operator_avs_state in self.operators[block_number].items():
                agg_pub_key_g1 = (
                    agg_pub_key_g1
                    + operator_avs_state.operator_info.pub_keys.g1_pub_key
                )
                total_stake += operator_avs_state.stake_per_quorum[qn]
            state[qn] = QuorumAvsState(
                quorum_number=qn,
                total_stake=total_stake,
                agg_pub_key_g1=agg_pub_key_g1,
                block_number=block_number,
            )
        return state

    async def get_check_signatures_indices(
        self,
        opts: CallOpts,
        reference_block_number: int,
        quorum_numbers: list[int],
        non_nigner_operator_ids: list[int],
    ) -> OperatorStateRetrieverCheckSignaturesIndices:
        result = OperatorStateRetrieverCheckSignaturesIndices(
            non_signer_quorum_bitmap_indices=[],
            quorum_apk_indices=[],
            total_stake_indices=[],
            non_signer_stake_indices=[],
        )
        return result, None
