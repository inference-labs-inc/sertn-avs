#!/bin/bash

RPC_URL=http://localhost:8545
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# cd to the directory of this script so that this can be run from anywhere
parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)
cd $parent_path/..

CHAINID=$(cast chain-id)
DEPLOYMENT_OUTPUT_FILE=./contracts/script/output/${CHAINID}/credible_squaring_avs_deployment_output.json
STRATEGY_ADDRESS=$(jq -r '.addresses.erc20MockStrategy' $DEPLOYMENT_OUTPUT_FILE)

go run cli/main.go --config config-files/operator.anvil.yaml deposit-into-strategy --strategy-addr ${STRATEGY_ADDRESS} --amount 100
