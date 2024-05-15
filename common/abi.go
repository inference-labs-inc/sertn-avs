package common

import (
	_ "embed"
)

//go:embed abis/OmronTaskManager.json
var TaskManagerAbi []byte
