from dataclasses import dataclass
from typing import Optional

import typer
import yaml

from avs_operator import run_operator
from aggregator import run_aggregator
from common.console import console, styles

app = typer.Typer(
    name="sertn",
    help="Sertn AVS Client",
    add_completion=False,
)


@app.command()
def start(
    mode: str = typer.Option(
        ...,
        "--mode",
        "-m",
        help="Operation mode: 'operator' or 'aggregator'",
    ),
    config: Optional[str] = typer.Option(
        None,
        "--config",
        "-c",
        help="Path to config file",
    ),
) -> None:
    console.print(f"Starting Sertn in {mode} mode...", style=styles.debug)

    try:
        with open(config, "r") as f:
            config_dict = config = yaml.load(f, Loader=yaml.BaseLoader)
    except Exception as e:
        console.print(
            f"Error loading config file. Please check the path and format.",
            style=styles.error,
        )
        raise typer.Exit(1)

    if mode == "operator":
        run_operator(config_dict)
    elif mode == "aggregator":
        run_aggregator(config_dict)
    else:
        console.print(
            f"Invalid mode: {mode}. Use 'operator' or 'aggregator'",
            style=styles.error,
        )
        raise typer.Exit(1)


if __name__ == "__main__":
    app()
