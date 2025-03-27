from dataclasses import dataclass, field
from typing import Any, Dict, List, Optional

from eth_typing import Address

from eigensdk.crypto.bls.attestation import G1Point, G2Point, Signature


@dataclass
class Operator:
    address: Address
    earnings_receiver_address: Address = field(default="")
    delegation_approver_address: Address = field(default="")
    staker_opt_out_window_blocks: Optional[int] = field(default=None)
    metadata_url: str = field(default="")


@dataclass
class OperatorPubkeys:
    # G1 signatures are used to verify signatures onchain (since G1 is cheaper to verify onchain via precompiles)
    g1_pub_key: G1Point
    # G2 is used to verify signatures offchain (signatures are on G1)
    g2_pub_key: G2Point


@dataclass
class OperatorInfo:
    socket: str
    pub_keys: OperatorPubkeys


@dataclass
class OperatorAvsState:
    operator_id: bytes
    operator_info: OperatorInfo
    stake_per_quorum: Dict[int, int]
    block_number: int


@dataclass
class QuorumAvsState:
    quorum_number: int
    total_stake: int
    agg_pub_key_g1: G1Point
    block_number: int


@dataclass
class OperatorStateRetrieverCheckSignaturesIndices:
    non_signer_quorum_bitmap_indices: List[int]
    quorum_apk_indices: List[int]
    total_stake_indices: List[int]
    non_signer_stake_indices: List[List[int]]


@dataclass
class SignedTaskResponseDigest:
    task_response: Any
    bls_signature: Signature
    operator_id: int


@dataclass
class OperatorStateRetrieverOperator:
    operator: Address
    operator_id: bytes
    stake: int
