from typing import Optional
from tqdm import tqdm
from rich.console import Console

console = Console()


def run_operator(config: Optional[str] = None) -> None:
    console.print("[bold green]Starting Sertn Operator...[/bold green]")

    with tqdm(total=100, desc="Initializing operator") as pbar:
        pbar.update(50)
        pbar.update(50)
