from web3 import Web3
from web3.eth import Eth

from common.config import GasStrategy
from common.console import console, styles


def get_gas_config(
    w3_eth: Eth, strategy: GasStrategy = GasStrategy.STANDARD
) -> tuple[int | None, int]:
    """
    Get recommended gas configuration based on strategy

    :param strategy: Gas strategy - "fast", "standard", "slow", or "priority"
    :return: Tuple with (max_priority_fee_per_gas, max_fee_per_gas)
    """
    try:
        latest_block = w3_eth.get_block("latest")
        if (
            hasattr(latest_block, "baseFeePerGas")
            and latest_block.baseFeePerGas is not None
        ):
            # EIP-1559 network
            base_fee = latest_block.baseFeePerGas

            strategies = {
                GasStrategy.SLOW: (
                    Web3.to_wei("1", "gwei"),  # max_priority_fee_per_gas
                    int(base_fee * 1.1) + Web3.to_wei("1", "gwei"),  # max_fee_per_gas
                ),
                GasStrategy.STANDARD: (
                    Web3.to_wei("2", "gwei"),  # max_priority_fee_per_gas
                    int(base_fee * 1.5) + Web3.to_wei("2", "gwei"),  # max_fee_per_gas
                ),
                GasStrategy.FAST: (
                    Web3.to_wei("3", "gwei"),  # max_priority_fee_per_gas
                    int(base_fee * 2) + Web3.to_wei("3", "gwei"),  # max_fee_per_gas
                ),
                GasStrategy.PRIORITY: (
                    Web3.to_wei("5", "gwei"),  # max_priority_fee_per_gas
                    int(base_fee * 3) + Web3.to_wei("5", "gwei"),  # max_fee_per_gas
                ),
            }
            return strategies.get(strategy, strategies["standard"])
        else:
            # Legacy network
            gas_price = w3_eth.gas_price
            multipliers = {
                GasStrategy.SLOW: 1.0,
                GasStrategy.STANDARD: 1.1,
                GasStrategy.FAST: 1.25,
                GasStrategy.PRIORITY: 1.5,
            }
            return (
                None,  # max_priority_fee_per_gas
                int(gas_price * multipliers.get(strategy, 1.1)),  # max_fee_per_gas
            )

    except Exception as e:
        console.print(f"Failed to get gas config: {e}", style=styles.error)
        return (None, None)  # Default to None if error occurs
