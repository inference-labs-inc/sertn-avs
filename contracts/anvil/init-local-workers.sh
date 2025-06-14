#!/bin/bash

# Load environment variables
source contracts/anvil/load-env.sh
cd contracts

# Run InitLocalWorkers script
# It inits operator and aggregator for local testing
echo "Running first stage to initialize local workers..."
STAGE=1 forge script script/InitLocalWorkers.t.sol --rpc-url $RPC_HOST:$RPC_PORT --broadcast

# Wait for a few blocks...
start_block=$(cast block-number --rpc-url $RPC_HOST:$RPC_PORT)
echo "Current block: $start_block, waiting for a few blocks to prevent AllocationDelay error..."

target_blocks=5
for i in $(seq 1 $target_blocks); do
  printf "\rWaiting... %2d blocks left" $(($target_blocks - i + 1))
  cast rpc --rpc-url $RPC_HOST:$RPC_PORT evm_mine > /dev/null 2>&1
  sleep 1
done
echo ""

# Run the second stage to complete the initialization
echo "Running the second stage to complete initialization..."

STAGE=2 forge script script/InitLocalWorkers.t.sol --rpc-url $RPC_HOST:$RPC_PORT --broadcast
