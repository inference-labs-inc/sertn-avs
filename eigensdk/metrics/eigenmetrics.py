from prometheus_client import Counter, Gauge, make_asgi_app, start_http_server, REGISTRY
from fastapi import FastAPI
import logging
import threading
import time

import uvicorn

EIGEN_PROM_NAMESPACE = "eigen"


class EigenMetrics:
    def __init__(self, avs_name, ip_port_address, logger: logging.Logger, reg=REGISTRY):
        self.ip_port_address = ip_port_address
        self.logger = logger

        # Metrics
        self.fee_earned_total = Counter(
            "fees_earned_total",
            "The amount of fees earned in <token>",
            ["token"],
            namespace=EIGEN_PROM_NAMESPACE,
            registry=reg,
        )
        self.performance_score = Gauge(
            "performance_score",
            "The performance metric is a score between 0 and 100 and each developer can define their own way of calculating the score. The score is calculated based on the performance of the Node and the performance of the backing services.",
            namespace=EIGEN_PROM_NAMESPACE,
            registry=reg,
        )

        self.init_metrics()

    def init_metrics(self):
        # Performance score starts as 100, and goes down if node doesn't perform well
        self.performance_score.set(100)
        # TODO: Initialize fee_earned_total if needed

    def add_fee_earned_total(self, amount, token):
        self.fee_earned_total.labels(token=token).inc(amount)

    def set_performance_score(self, score):
        self.performance_score.set(score)

    def start(self):
        self.logger.info(f"Starting metrics server at port {self.ip_port_address}")

        try:
            app = FastAPI()
            metrics_app = make_asgi_app()
            app.mount("/metrics", metrics_app)

            uvicorn.run(
                app,
                host=self.ip_port_address.split(":")[0],
                port=int(self.ip_port_address.split(":")[1]),
            )
        except Exception as e:
            self.logger.error(f"Prometheus server failed: {e}")
            return


# Usage example:
if __name__ == "__main__":
    import queue
    import signal

    logging.basicConfig(level=logging.INFO)

    logger = logging.getLogger("NodeAPI")
    logger.setLevel(logging.INFO)
    handler = logging.StreamHandler()
    handler.setFormatter(
        logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s")
    )
    logger.addHandler(handler)

    metrics = EigenMetrics("example_avs", "0.0.0.0:8000", logger)
    metrics.start()
