from typing import Optional
from tqdm import tqdm

from console import console, styles


def run_aggregator(config: dict) -> None:
    console.print("Starting Sertn Aggregator...", style=styles.agg_info)

    with tqdm(total=100, desc="Initializing aggregator") as pbar:
        pbar.update(50)
        pbar.update(50)
