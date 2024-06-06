package types

import (
	"errors"

	cstaskmanager "github.com/inference-labs-inc/zklayer-avs/contracts/bindings/ZklayerTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.ITaskStructTaskResponse
	TaskResponseMetadata      cstaskmanager.ITaskStructTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
