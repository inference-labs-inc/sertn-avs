#!/bin/bash

# Load environment variables
source contracts/anvil/load-env.sh
RUN_ANVIL_IN_BACKGROUND=${1:-false}

if [ $RUN_ANVIL_IN_BACKGROUND = true ]; then
  echo "Starting Anvil with state dump in background"
  anvil --dump-state "$STATE_FILE" --host 0.0.0.0 --port $RPC_PORT --base-fee 0 --gas-price 0 &
  ANVIL_PID=$!

  # Save the PID to a file
  echo $ANVIL_PID > $PID_FILE
  echo "Anvil started with PID $ANVIL_PID"
else
  # Start Anvil in foreground
  anvil --dump-state "$STATE_FILE" --host 0.0.0.0 --port $RPC_PORT --base-fee 0 --gas-price 0 -vvv
fi
