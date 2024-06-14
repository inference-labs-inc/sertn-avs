package common

import (
	_ "embed"
)

//go:embed abis/SertnTaskManager.json
var TaskManagerAbi []byte
