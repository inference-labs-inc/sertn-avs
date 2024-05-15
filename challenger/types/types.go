package types

import (
	"errors"

	cstaskmanager "github.com/inference-labs-inc/omron-avs/contracts/bindings/OmronTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IOmronTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IOmronTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
