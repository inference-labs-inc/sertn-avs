package types

import (
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

// all operators in quorum0 must sign the task response in order for it to be accepted
// TODO: our contracts require uint8 but right now sdktypes.QuorumThresholdPercentage is uint8
//       prob need to update our inc-sq contracts to use uint8 as well?
const QUORUM_THRESHOLD_NUMERATOR = sdktypes.QuorumThresholdPercentage(100)
const QUORUM_THRESHOLD_DENOMINATOR = sdktypes.QuorumThresholdPercentage(100)

const QUERY_FILTER_FROM_BLOCK = uint64(1)

// we only use a single quorum (quorum 0) for incredible squaring
var QUORUM_NUMBERS = sdktypes.QuorumNums{0}

type BlockNumber = uint32
type TaskIndex = uint32

type OperatorInfo struct {
	OperatorPubkeys sdktypes.OperatorPubkeys
	OperatorAddr    common.Address
}
