package common

import (
	_ "embed"
)

//go:embed abis/ZklayerTaskManager.json
var TaskManagerAbi []byte
