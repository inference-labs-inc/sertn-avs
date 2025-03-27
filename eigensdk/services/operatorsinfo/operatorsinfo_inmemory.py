import logging
import time
from threading import Thread
from typing import Any, Dict, Optional

from eth_typing import Address
from eth_utils.encoding import int_to_big_endian
from web3 import Web3

from eigensdk._types import OperatorInfo, OperatorPubkeys
from eigensdk.chainio.clients.avsregistry.reader import AvsRegistryReader
from eigensdk.crypto.bls.attestation import G1Point


class OperatorsInfoServiceInMemory:
    def __init__(
        self,
        avs_registry_reader: AvsRegistryReader,
        start_block_pub: int = 0,
        start_block_socket: int = 0,
        check_interval: int = 10,
        log_filter_query_block_range: int = 10000,
        logger: Optional[logging.Logger] = None,
    ):
        self.avs_registry_reader: AvsRegistryReader = avs_registry_reader
        self.start_block_pub: int = start_block_pub
        self.start_block_socket: int = start_block_socket
        self.check_interval: int = check_interval
        self.log_filter_query_block_range: int = log_filter_query_block_range
        self.logger: Optional[logging.Logger] = logger or logging.getLogger(__name__)
        self.eth_http_client: Any = self.avs_registry_reader.eth_http_client

        self.pubkey_dict: Dict[bytes, OperatorPubkeys] = {}
        self.operator_addr_to_id: Dict[Address, bytes] = {}
        self.socket_dict: Dict[bytes, str] = {}

        # Start the service in a separate thread
        self.get_events()
        self.thread = Thread(target=self._service_thread)
        self.thread.start()

    @staticmethod
    def operator_id_from_g1_pubkey(g1: G1Point) -> bytes:
        x_bytes = int_to_big_endian(int(g1.x.getStr()))
        y_bytes = int_to_big_endian(int(g1.y.getStr()))
        concatenated = x_bytes + y_bytes
        return Web3.keccak(concatenated)

    def _service_thread(self) -> None:
        while True:
            try:
                self.get_events()
            except Exception as e:
                self.logger.error(f"Get event Error: {e}")
                pass
            time.sleep(self.check_interval)

    def get_events(self) -> None:
        operator_addresses, operator_pubkeys, to_block_pub = (
            self.avs_registry_reader.query_existing_registered_operator_pubkeys(
                start_block=self.start_block_pub
            )
        )
        operator_sockets, to_block_socket = (
            self.avs_registry_reader.query_existing_registered_operator_sockets(
                start_block=self.start_block_socket
            )
        )

        for i, operator_addr in enumerate(operator_addresses):
            operator_pubkeys = operator_pubkeys[i]
            operator_id = self.operator_id_from_g1_pubkey(operator_pubkeys.g1_pub_key)
            self.pubkey_dict[operator_id] = operator_pubkeys
            self.operator_addr_to_id[operator_addr] = operator_id

        self.socket_dict.update(operator_sockets)
        self.logger.debug(f"Queried operator registration events: {operator_pubkeys}")

        self.start_block_pub = to_block_pub
        self.start_block_socket = to_block_socket

    def get_operator_info(self, operator_addr: Address) -> OperatorInfo:
        operator_id = self.operator_addr_to_id.get(operator_addr)
        if not operator_id:
            raise Exception("Not found")
        return OperatorInfo(
            socket=self.socket_dict[operator_id],
            pub_keys=self.pubkey_dict[operator_id],
        )
