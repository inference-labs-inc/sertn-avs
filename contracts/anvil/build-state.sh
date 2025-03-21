#!/bin/bash

STATE_FILE="contracts/anvil/state.json"

mkdir -p "$(dirname "$STATE_FILE")"

echo "Starting Anvil with state dump in background"
anvil --dump-state "$STATE_FILE" --host 0.0.0.0 --port 8545 --base-fee 0 --gas-price 0 &
ANVIL_PID=$!

sleep 3

cp .env.example .env
cp contracts/.env.example contracts/.env

echo "Building contracts"
make build-contracts

echo "Deploying EigenLayer contracts."
make deploy-eigenlayer-contracts

echo "Deploying Sertn contracts."
make deploy-sertn-contracts

echo "Killed Anvil"
kill $ANVIL_PID
