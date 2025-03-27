from logging import Logger
from prometheus_client.core import GaugeMetricFamily, CounterMetricFamily
from prometheus_client import registry
from eigensdk.chainio.clients.avsregistry.reader import AvsRegistryReader
from eigensdk.chainio.clients.elcontracts.reader import ELReader


class Collector(registry.Collector):
    def __init__(
        self,
        el_reader: ELReader,
        avs_registry_reader: AvsRegistryReader,
        avs_name: str,
        logger: Logger,
        operator_addr: str,
        quorum_names: dict,
    ):
        self.el_reader = el_reader
        self.avs_registry_reader = avs_registry_reader
        self.logger = logger
        self.operator_addr = operator_addr
        self.operator_id = None
        self.quorum_names = quorum_names

        self.slashing_status = CounterMetricFamily(
            "eigen_slashing_status",
            "Whether the operator has been slashed",
        )
        self.registered_stake = GaugeMetricFamily(
            "eigen_registered_stakes",
            "Operator stake in <quorum> of <avs_name>'s StakeRegistry contract",
            labels=["quorum_number", "quorum_name"],
        )

    def describe(self):
        yield self.slashing_status
        yield self.registered_stake

    def init_operator_id(self):
        if self.operator_id is None:
            self.operator_id = self.avs_registry_reader.get_operator_id(
                self.operator_addr
            )

        return self.operator_id is not None  # true means success

    def collect(self):
        # Collect slashingStatus metric
        operator_is_frozen = self.el_reader.operator_is_frozen(self.operator_addr)
        if operator_is_frozen is None:
            self.logger.error("Failed to get slashing incurred")
        else:
            operator_is_frozen_value = 1.0 if operator_is_frozen else 0.0
            self.slashing_status.set(operator_is_frozen_value)
            yield self.slashing_status

        # Collect registeredStake metric
        if not self.init_operator_id():
            self.logger.warn(
                "Failed to fetch and cache operator id. Skipping collection of registeredStake metric."
            )
        else:
            quorum_stake_map = self.avs_registry_reader.get_operator_stake_in_quorums(
                self.operator_id
            )
            for quorum_num, stake in quorum_stake_map.items():
                stake_value = float(stake)
                self.registered_stake.labels(
                    quorum_num=str(quorum_num),
                    quorum_name=self.quorum_names[quorum_num],
                ).set(stake_value)
                yield self.registered_stake
