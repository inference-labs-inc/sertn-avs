from prometheus_client import Histogram, Counter, REGISTRY


class Collector:
    def __init__(self, avs_name, registry=None):
        if registry is None:
            registry = REGISTRY

        self.rpc_request_duration_seconds = Histogram(
            "rpc_request_duration_seconds",
            "Duration of json-rpc <method> in seconds",
            ["method", "client_version"],
            namespace="eigen_prom_namespace",
            labelnames={"avs_name": avs_name},
            registry=registry,
        )

        self.rpc_request_total = Counter(
            "rpc_request_total",
            "Total number of json-rpc <method> requests",
            ["method", "client_version"],
            namespace="eigen_prom_namespace",
            labelnames={"avs_name": avs_name},
            registry=registry,
        )

    def observe_rpc_request_duration_seconds(self, duration, method, client_version):
        self.rpc_request_duration_seconds.labels(
            method=method, client_version=client_version
        ).observe(duration)

    def add_rpc_request_total(self, method, client_version):
        self.rpc_request_total.labels(
            method=method, client_version=client_version
        ).inc()
