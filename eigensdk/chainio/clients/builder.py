import logging
from typing import Any, Tuple, Optional

from eth_account import Account
from eth_account.signers.local import LocalAccount
from eth_typing import Address
from web3 import Web3

from eigensdk.contracts import ABIs

from .avsregistry import reader as avs_reader
from .avsregistry import writer as avs_writer
from .elcontracts import reader as el_reader
from .elcontracts import writer as el_writer


class BuildAllConfig:
    def __init__(
        self,
        eth_http_url: str,
        registry_coordinator_addr: Address,
        operator_state_retriever_addr: Address,
        avs_name: str = '',
        prom_metrics_ip_port_address: str = '',
    ):
        self.eth_http_url: str = eth_http_url
        self.registry_coordinator_addr: Address = registry_coordinator_addr
        self.operator_state_retriever_addr: Address = operator_state_retriever_addr
        self.avs_name: str = avs_name
        self.prom_metrics_ip_port_address: str = prom_metrics_ip_port_address

    def build_el_clients(
        self, pk_wallet: LocalAccount, logger: logging.Logger
    ) -> Tuple[el_reader.ELReader, el_writer.ELWriter]:
        eth_http_client = Web3(Web3.HTTPProvider(self.eth_http_url))
        registry_coordinator = eth_http_client.eth.contract(
            address=self.registry_coordinator_addr,
            abi=ABIs.REGISTRY_COORDINATOR,
        )

        stake_registry_addr = registry_coordinator.functions.stakeRegistry().call()
        stake_registry = eth_http_client.eth.contract(
            address=stake_registry_addr,
            abi=ABIs.STAKE_REGISTRY,
        )

        delegation_manager_addr = stake_registry.functions.delegation().call()
        delegation_manager = eth_http_client.eth.contract(
            address=delegation_manager_addr,
            abi=ABIs.DELEGATION_MANAGER,
        )

        # fixme: slasher seems to be removed in testnet newer version
        # these codes should be updated based on latest changes to eigensdk-go

        # slasher_addr = delegation_manager.functions.slasher().call()
        # slasher = eth_http_client.eth.contract(
        #     address=slasher_addr,
        #     abi=ABIs.SLASHER,
        # )

        strategy_manager_addr = delegation_manager.functions.strategyManager().call()
        strategy_manager = eth_http_client.eth.contract(
            address=strategy_manager_addr,
            abi=ABIs.STRATEGY_MANAGER,
        )

        service_manager_addr = registry_coordinator.functions.serviceManager().call()
        service_manager = eth_http_client.eth.contract(
            address=service_manager_addr,
            abi=ABIs.SERVICE_MANAGER,
        )

        avs_directory_addr = service_manager.functions.avsDirectory().call()
        avs_directory = eth_http_client.eth.contract(
            address=avs_directory_addr,
            abi=ABIs.AVS_DIRECTORY,
        )

        el_reader_instance = el_reader.ELReader(
            # slasher,
            "0x0000000000000000000000000000000000000000",
            delegation_manager,
            strategy_manager,
            avs_directory,
            logger,
            eth_http_client,
        )

        el_writer_instance = el_writer.ELWriter(
            # slasher,
            "0x0000000000000000000000000000000000000000",
            delegation_manager,
            strategy_manager,
            strategy_manager_addr,
            avs_directory,
            el_reader_instance,
            logger,
            eth_http_client,
            pk_wallet,
        )

        return el_reader_instance, el_writer_instance

    def build_avs_registry_clients(
        self,
        el_reader: el_reader.ELReader,
        logger: logging.Logger,
        pk_wallet: LocalAccount,
    ) -> Tuple[avs_reader.AvsRegistryReader, avs_writer.AvsRegistryWriter]:
        eth_http_client = Web3(Web3.HTTPProvider(self.eth_http_url))
        registry_coordinator = eth_http_client.eth.contract(
            address=self.registry_coordinator_addr,
            abi=ABIs.REGISTRY_COORDINATOR,
        )
        service_manager_addr = registry_coordinator.functions.serviceManager().call()

        bls_apk_registry_addr = registry_coordinator.functions.blsApkRegistry().call()
        bls_apk_registry = eth_http_client.eth.contract(
            address=bls_apk_registry_addr,
            abi=ABIs.BLS_APK_REGISTRY,
        )

        operator_state_retriever = eth_http_client.eth.contract(
            address=self.operator_state_retriever_addr,
            abi=ABIs.OPERATOR_STATE_RETRIEVER,
        )

        stake_registry_addr = registry_coordinator.functions.stakeRegistry().call()
        stake_registry = eth_http_client.eth.contract(
            address=stake_registry_addr,
            abi=ABIs.STAKE_REGISTRY,
        )

        avs_registry_reader = avs_reader.AvsRegistryReader(
            self.registry_coordinator_addr,
            registry_coordinator,
            bls_apk_registry_addr,
            bls_apk_registry,
            operator_state_retriever,
            stake_registry,
            logger,
            eth_http_client,
        )

        avs_registry_writer = avs_writer.AvsRegistryWriter(
            service_manager_addr,
            registry_coordinator,
            operator_state_retriever,
            stake_registry,
            bls_apk_registry,
            el_reader,
            logger,
            eth_http_client,
            pk_wallet,
        )

        return avs_registry_reader, avs_registry_writer


class Clients:
    def __init__(
        self,
        avs_registry_reader: avs_reader.AvsRegistryReader,
        avs_registry_writer: avs_writer.AvsRegistryWriter,
        el_reader: el_reader.ELReader,
        el_writer: el_writer.ELWriter,
        eth_http_client: Web3,
        wallet: LocalAccount,
        metrics: Optional[Any],
    ):
        self.avs_registry_reader = avs_registry_reader
        self.avs_registry_writer = avs_registry_writer
        self.el_reader = el_reader
        self.el_writer = el_writer
        self.eth_http_client = eth_http_client
        self.wallet = wallet
        self.metrics = metrics


def build_all(
    config: BuildAllConfig, ecdsa_private_key: str = '', logger: logging.Logger = logging.getLogger(__name__)
) -> Clients:
    eth_http_client = Web3(Web3.HTTPProvider(config.eth_http_url))

    pk_wallet: LocalAccount = Account.from_key(ecdsa_private_key) if ecdsa_private_key else None

    el_reader, el_writer = config.build_el_clients(pk_wallet, logger)

    (
        avs_registry_reader,
        avs_registry_writer,
    ) = config.build_avs_registry_clients(el_reader, logger, pk_wallet)

    return Clients(
        avs_registry_reader=avs_registry_reader,
        avs_registry_writer=avs_registry_writer,
        el_reader=el_reader,
        el_writer=el_writer,
        eth_http_client=eth_http_client,
        wallet=pk_wallet,
        metrics=None,
    )
