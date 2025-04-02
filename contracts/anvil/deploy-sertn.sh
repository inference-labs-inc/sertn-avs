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

forge script script/SertnDeployer.s.sol --rpc-url $RPC_HOST:$RPC_PORT --broadcast

# Format the JSON file using Python
if [ -f deployments/sertnDeployment.json ]; then
    python3 -c "
import json
with open('deployments/sertnDeployment.json', 'r') as f:
    data = json.load(f)
with open('deployments/sertnDeployment.json', 'w') as f:
    json.dump(data, f, indent=4)
"
    echo "Formatted contracts/deployments/sertnDeployment.json"
else
    echo "contracts/deployments/sertnDeployment.json not found!"
fi
