#!/usr/bin/env bash

# Load environment variables
source contracts/anvil/load-env.sh

# cd to the directory of this script so that this can be run from anywhere
parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)
cd "$parent_path"

cd ../

forge script script/DeployEigenLayerCore.s.sol --rpc-url $RPC_HOST:$RPC_PORT --broadcast