from typing import Optional

import typer

from avs_operator import run_operator
from aggregator import run_aggregator
from common.config import load_config
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

    if not config:
        console.print(
            "Config file path is required. Use --config to specify the path.",
            style=styles.error,
        )
        raise typer.Exit(1)

    try:
        if mode == "operator":
            config_obj = load_config(config, "operator")
            run_operator(config_obj)
        elif mode == "aggregator":
            config_obj = load_config(config, "aggregator")
            run_aggregator(config_obj)
        else:
            console.print(
                f"Invalid mode: {mode}. Use 'operator' or 'aggregator'",
                style=styles.error,
            )
            raise typer.Exit(1)
    except Exception as e:
        console.print(
            f"Error loading or validating config: {e}",
            style=styles.error,
        )
        raise typer.Exit(1)


if __name__ == "__main__":
    app()
