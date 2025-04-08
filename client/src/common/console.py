from dataclasses import dataclass

from rich.console import Console
from rich.style import Style


console = Console()


@dataclass
class ApplicationStyles:
    error = Style(
        color="red",
        bold=True,
    )
    agg_info = Style(
        color="green",
        bold=True,
    )
    op_info = Style(
        color="blue",
        bold=True,
    )
    warning = Style(
        color="yellow",
        bold=True,
    )
    debug = Style(
        color="grey0",
        bold=True,
    )


styles = ApplicationStyles()
