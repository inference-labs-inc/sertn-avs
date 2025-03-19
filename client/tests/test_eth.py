import pytest
from unittest.mock import Mock, patch
from common.eth import EthereumClient


@patch("common.eth.Web3")
def test_ethereum_client_init(mock_web3):
    mock_web3.return_value.eth = Mock()
    client = EthereumClient()
    assert client.w3 is not None


def test_ethereum_client_custom_rpc():
    custom_rpc = "http://localhost:8546"
    client = EthereumClient(rpc_url=custom_rpc)
    assert client.w3.provider.endpoint_uri == custom_rpc


@patch("common.eth.Web3")
def test_ethereum_client_chain_id(mock_web3):
    mock_web3.return_value.eth = Mock()
    mock_web3.return_value.eth.chain_id = 1
    client = EthereumClient()
    chain_id = client.get_chain_id()
    assert isinstance(chain_id, int)
    assert chain_id == 1
