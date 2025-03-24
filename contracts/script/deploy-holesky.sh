#!/usr/bin/env bash

set -e

RPC_URL=https://1rpc.io/holesky
CHAIN_ID=17000

if [ -z "$HOLESKY_DEPLOYER_KEY" ]; then
    echo "Error: HOLESKY_DEPLOYER_KEY environment variable is not set"
    exit 1
fi

ADDRESS=$(cast wallet address $HOLESKY_DEPLOYER_KEY)
echo "Deploying from address: $ADDRESS"

BALANCE=$(cast balance $ADDRESS --rpc-url $RPC_URL)
BALANCE_ETH=$(cast --from-wei $BALANCE)
echo "Account balance: $BALANCE_ETH ETH"
echo "Account balance in wei: $BALANCE"

if [ "$BALANCE" -eq "0" ]; then
    echo "Error: Account has no ETH"
    exit 1
fi

BASE_FEE=$(cast gas-price --rpc-url $RPC_URL)
GAS_PRICE=$((BASE_FEE + 100000))
echo "Using gas price: $(cast --to-unit $GAS_PRICE gwei) gwei"

NONCE=$(cast nonce $ADDRESS --rpc-url $RPC_URL)
echo "Current nonce: $NONCE"

parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)
cd "$parent_path"

cd ../

echo "Deploying Service Manager..."
FOUNDRY_PROFILE=debug forge script script/SertnDeployer.s.sol:SertnDeployer \
    --sig "deployServiceManager()" \
    --rpc-url $RPC_URL \
    --broadcast \
    --skip-simulation \
    --gas-limit 15000000 \
    --gas-price $GAS_PRICE \
    --private-key $HOLESKY_DEPLOYER_KEY \
    -vv

echo "Deploying Task Manager..."
forge script script/SertnDeployer.s.sol:SertnDeployer \
    --sig "deployTaskManager()" \
    --rpc-url $RPC_URL \
    --broadcast \
    --verify \
    --skip-simulation \
    --gas-limit 15000000 \
    --gas-price $GAS_PRICE \
    --private-key $HOLESKY_DEPLOYER_KEY \
    -vvvv
