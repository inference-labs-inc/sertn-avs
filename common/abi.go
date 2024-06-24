package common

import (
	_ "embed"
)

//go:embed abis/TaskManager.abi.json
var TaskManagerAbi []byte
