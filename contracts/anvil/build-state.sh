#!/bin/bash

STATE_FILE="contracts/anvil/state.json"

mkdir -p "$(dirname "$STATE_FILE")"

echo "Starting Anvil with state dump in background"
anvil --dump-state "$STATE_FILE" --host 0.0.0.0 --port 8545 --base-fee 0 --gas-price 0 > /dev/null 2>&1 &
ANVIL_PID=$!

sleep 3

cp .env.example .env
cp contracts/.env.example contracts/.env

echo "Building contracts"
make build-contracts > /dev/null 2>&1

echo "Deploying EigenLayer contracts."
make deploy-eigenlayer-contracts > /dev/null 2>&1

echo "Deploying Sertn contracts."
make deploy-sertn-contracts > /dev/null 2>&1

echo "Killed Anvil"
kill $ANVIL_PID
