#!/bin/bash

# Load environment variables
source contracts/anvil/load-env.sh
RUN_ANVIL_IN_BACKGROUND=${1:-false}

if [ $RUN_ANVIL_IN_BACKGROUND = true ]; then
  echo "Starting Anvil with state dump in background"
  nohup anvil \
      --host 0.0.0.0 \
      --port $RPC_PORT \
      --base-fee 0 \
      --gas-price 0 \
      --timestamp 1672531200 > anvil.log 2>&1 &
  ANVIL_PID=$!

  # Save the PID to a file
  echo $ANVIL_PID > $ANVIL_PID_FILE
  echo "Anvil started with PID '$ANVIL_PID' pid file is '$ANVIL_PID_FILE'"
else
  # Start Anvil in foreground
  echo "Starting Anvil with state dump in foreground"
  anvil \
      --host 0.0.0.0 \
      --port $RPC_PORT \
      --base-fee 0 \
      --gas-price 0 \
      --timestamp 1672531200 \
      -vvv
fi
