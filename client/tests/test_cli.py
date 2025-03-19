from typer.testing import CliRunner
from main import app

runner = CliRunner()


def test_cli_operator_mode():
    result = runner.invoke(app, ["--mode", "operator"])
    print(f"stdout: {result.stdout}")
    assert result.exit_code == 0
    assert "Starting Sertn in operator mode" in result.stdout


def test_cli_aggregator_mode():
    result = runner.invoke(app, ["--mode", "aggregator"])
    print(f"stdout: {result.stdout}")
    assert result.exit_code == 0
    assert "Starting Sertn in aggregator mode" in result.stdout


def test_cli_invalid_mode():
    result = runner.invoke(app, ["--mode", "invalid"])
    print(f"stdout: {result.stdout}")
    assert result.exit_code == 1
    assert "Invalid mode: invalid" in result.stdout
