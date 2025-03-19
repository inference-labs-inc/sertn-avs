from typing import Optional
from tqdm import tqdm
from rich.console import Console

console = Console()


def run_aggregator(config: Optional[str] = None) -> None:
    console.print("[bold blue]Starting Sertn Aggregator...[/bold blue]")

    with tqdm(total=100, desc="Initializing aggregator") as pbar:
        pbar.update(50)
        pbar.update(50)
