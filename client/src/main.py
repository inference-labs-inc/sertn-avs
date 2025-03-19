import typer
from rich.console import Console
from typing import Optional

from avs_operator import run_operator
from aggregator import run_aggregator

app = typer.Typer(
    name="sertn",
    help="Sertn AVS Client",
    add_completion=False,
)
console = Console()


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
    try:
        console.print(f"Starting Sertn in {mode} mode...")

        if mode == "operator":
            run_operator(config)
        elif mode == "aggregator":
            run_aggregator(config)
        else:
            console.print(
                f"[red]Invalid mode: {mode}. Use 'operator' or 'aggregator'[/red]"
            )
            raise typer.Exit(1)
    except Exception as e:
        console.print(f"[red]Error: {str(e)}[/red]")
        raise typer.Exit(1)


if __name__ == "__main__":
    app()
