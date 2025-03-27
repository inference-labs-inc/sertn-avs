from fastapi import APIRouter, FastAPI, HTTPException, Response
from pydantic import BaseModel
from enum import Enum
import uvicorn
import logging

# Constants
BASE_URL = "/eigen"
SPEC_SEM_VER = "v0.0.1"


# Enums and Models
class NodeHealth(Enum):
    Healthy = 0
    PartiallyHealthy = 1
    Unhealthy = 2


class ServiceStatus(Enum):
    Up = "Up"
    Down = "Down"
    Initializing = "Initializing"


class NodeService(BaseModel):
    id: str
    name: str
    description: str
    status: ServiceStatus


class NodeAPI:
    def __init__(
        self,
        avs_node_name: str,
        avs_node_sem_ver: str,
        ip_port_addr: str,
        logger: logging.Logger,
    ):
        self.avs_node_name = avs_node_name
        self.avs_node_sem_ver = avs_node_sem_ver
        self.health = NodeHealth.Healthy
        self.node_services: list[NodeService] = []
        self.ip_port_addr = ip_port_addr
        self.logger = logger

        # self.logger = logging.getLogger("NodeAPI")
        # self.logger.setLevel(logging.INFO)
        # handler = logging.StreamHandler()
        # handler.setFormatter(
        #     logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s")
        # )
        # self.logger.addHandler(handler)
        self.router = APIRouter(prefix=BASE_URL)

    def update_health(self, health: NodeHealth):
        self.health = health

    def register_new_service(
        self,
        service_id: str,
        service_name: str,
        service_description: str,
        service_status: ServiceStatus,
    ):
        new_service = NodeService(
            id=service_id,
            name=service_name,
            description=service_description,
            status=service_status,
        )
        self.node_services.append(new_service)

    def update_service_status(self, service_id: str, new_status: ServiceStatus):
        for service in self.node_services:
            if service.id == service_id:
                service.status = new_status
                return
        raise ValueError(f"Service with serviceId {service_id} not found")

    def deregister_service(self, service_id: str):
        self.node_services = [
            service for service in self.node_services if service.id != service_id
        ]


def node_handler(node_api: NodeAPI):
    async def f():
        return {
            "node_name": node_api.avs_node_name,
            "spec_version": SPEC_SEM_VER,
            "node_version": node_api.avs_node_sem_ver,
        }

    return f


def health_handler(node_api: NodeAPI):
    async def f():
        if node_api.health == NodeHealth.Healthy:
            return Response(status_code=200)
        elif node_api.health == NodeHealth.PartiallyHealthy:
            return Response(status_code=206)
        elif node_api.health == NodeHealth.Unhealthy:
            return Response(status_code=503)
        else:
            node_api.logger.error(
                "Unknown health status", extra={"health": node_api.health}
            )
            return Response(status_code=503)

    return f


def services_handler(node_api: NodeAPI):
    async def f():
        return {"services": node_api.node_services}

    return f


def service_health_handler(node_api: NodeAPI):
    async def f(service_id: str):

        for service in node_api.node_services:
            if service.id == service_id:
                return {"service_id": service_id, "status": service.status.value}
        raise HTTPException(status_code=404, detail="Service not found")

    return f


def run(node_api: NodeAPI):
    app = FastAPI()
    router = APIRouter(prefix=BASE_URL)

    router.add_api_route("/node", endpoint=node_handler(node_api))
    router.add_api_route("/node/health", endpoint=health_handler(node_api))
    router.add_api_route("/node/services", endpoint=services_handler(node_api))
    router.add_api_route(
        "/node/services/{service_id}/health", endpoint=service_health_handler(node_api)
    )

    host, port = node_api.ip_port_addr.split(':')
    uvicorn.run(app, host=host, port=int(port))
