from typing import Optional

from tqdm import tqdm
import yaml

from console import console, styles


def run_operator(config: dict) -> None:
    console.print("Starting Sertn Operator...", style=styles.op_info)

    with tqdm(total=100, desc="Initializing operator") as pbar:

        pbar.update(50)
        pbar.update(50)
