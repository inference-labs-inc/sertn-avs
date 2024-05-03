package challenger

import (
	"context"
	"math/big"
	"testing"

	mockethclient "github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	aggtypes "github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	"github.com/Layr-Labs/incredible-squaring-avs/challenger/mocks"
	chtypes "github.com/Layr-Labs/incredible-squaring-avs/challenger/types"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	chainiomocks "github.com/Layr-Labs/incredible-squaring-avs/core/chainio/mocks"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var MOCK_OPERATOR_ID = [32]byte{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}
var MOCK_OPERATOR_STAKE = big.NewInt(100)
var MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING = "50"

// @samlaf I tried pulling the MockTask struct froma ggregator_test but getting error: "undefined: aggregator.MockTask"
type MockTask struct {
	TaskNum        uint32
	BlockNumber    uint32
	NumberToSquare uint32
}

func TestCallChallengeModule(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	challenger, mockAvsWriterer, _, _, _, err := createMockChallenger(mockCtrl)
	assert.Nil(t, err)

	const TASK_INDEX = 1
	const BLOCK_NUMBER = uint32(100)

	challenger.tasks[TASK_INDEX] = cstaskmanager.IIncredibleSquaringTaskManagerTask{
		NumberToBeSquared:         big.NewInt(3),
		TaskCreatedBlock:          1000,
		QuorumNumbers:             aggtypes.QUORUM_NUMBERS.UnderlyingType(),
		QuorumThresholdPercentage: uint32(aggtypes.QUORUM_THRESHOLD_NUMERATOR),
	}

	challenger.taskResponses[TASK_INDEX] = chtypes.TaskResponseData{
		TaskResponse: cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
			ReferenceTaskIndex: TASK_INDEX,
			NumberSquared:      big.NewInt(2),
		},
		TaskResponseMetadata: cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata{
			TaskResponsedBlock: 1001,
			HashOfNonSigners:   [32]byte{},
		},
		NonSigningOperatorPubKeys: []cstaskmanager.BN254G1Point{},
	}

	mockAvsWriterer.EXPECT().RaiseChallenge(
		context.Background(),
		challenger.tasks[TASK_INDEX],
		challenger.taskResponses[TASK_INDEX].TaskResponse,
		challenger.taskResponses[TASK_INDEX].TaskResponseMetadata,
		challenger.taskResponses[TASK_INDEX].NonSigningOperatorPubKeys,
	).Return(mocks.MockRaiseAndResolveChallengeCall(BLOCK_NUMBER, TASK_INDEX), nil)

	msg := challenger.callChallengeModule(TASK_INDEX)
	assert.Equal(t, msg, nil)

}

func TestRaiseChallenge(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	challenger, mockAvsWriterer, _, _, _, err := createMockChallenger(mockCtrl)
	assert.Nil(t, err)

	const TASK_INDEX = 1
	const BLOCK_NUMBER = uint32(100)

	challenger.tasks[TASK_INDEX] = cstaskmanager.IIncredibleSquaringTaskManagerTask{
		NumberToBeSquared:         big.NewInt(3),
		TaskCreatedBlock:          1000,
		QuorumNumbers:             aggtypes.QUORUM_NUMBERS.UnderlyingType(),
		QuorumThresholdPercentage: uint32(aggtypes.QUORUM_THRESHOLD_NUMERATOR),
	}

	challenger.taskResponses[TASK_INDEX] = chtypes.TaskResponseData{
		TaskResponse: cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
			ReferenceTaskIndex: TASK_INDEX,
			NumberSquared:      big.NewInt(9),
		},
		TaskResponseMetadata: cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata{
			TaskResponsedBlock: 1001,
			HashOfNonSigners:   [32]byte{},
		},
		NonSigningOperatorPubKeys: []cstaskmanager.BN254G1Point{},
	}

	mockAvsWriterer.EXPECT().RaiseChallenge(
		context.Background(),
		challenger.tasks[TASK_INDEX],
		challenger.taskResponses[TASK_INDEX].TaskResponse,
		challenger.taskResponses[TASK_INDEX].TaskResponseMetadata,
		challenger.taskResponses[TASK_INDEX].NonSigningOperatorPubKeys,
	).Return(mocks.MockRaiseAndResolveChallengeCall(BLOCK_NUMBER, TASK_INDEX), nil)

	err = challenger.raiseChallenge(TASK_INDEX)
	assert.Nil(t, err)
}

func TestProcessTaskResponseLog(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	challenger, _, _, mockAvsSubscriber, mockEthClient, err := createMockChallenger(mockCtrl)
	assert.Nil(t, err)

	const TASK_INDEX = 1

	challenger.tasks[TASK_INDEX] = cstaskmanager.IIncredibleSquaringTaskManagerTask{
		NumberToBeSquared:         big.NewInt(3),
		TaskCreatedBlock:          1000,
		QuorumNumbers:             aggtypes.QUORUM_NUMBERS.UnderlyingType(),
		QuorumThresholdPercentage: uint32(aggtypes.QUORUM_THRESHOLD_NUMERATOR),
	}

	challenger.taskResponses[TASK_INDEX] = chtypes.TaskResponseData{
		TaskResponse: cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
			ReferenceTaskIndex: TASK_INDEX,
			NumberSquared:      big.NewInt(9),
		},
		TaskResponseMetadata: cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata{
			TaskResponsedBlock: 1001,
			HashOfNonSigners:   [32]byte{},
		},
		NonSigningOperatorPubKeys: []cstaskmanager.BN254G1Point{},
	}

	taskResponseLog := cstaskmanager.ContractIncredibleSquaringTaskManagerTaskResponded{
		TaskResponse:         challenger.taskResponses[TASK_INDEX].TaskResponse,
		TaskResponseMetadata: challenger.taskResponses[TASK_INDEX].TaskResponseMetadata,
		Raw: gethtypes.Log{
			Address: common.HexToAddress("0x9e545e3c0baab3e08cdfd552c960a1050f373042"),
			Topics: []common.Hash{
				// this is the actual TaskChallengedSuccessfully event from calling the func "raiseAndResolveChallenge"
				common.HexToHash("0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a"),
			},
			Data:        common.Hex2Bytes("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a50fe07922f57ae3b4553201bfd7c11aca85e1541f91db8e62dca9c418dc5feae"),
			BlockNumber: uint64(100),
			TxHash:      common.HexToHash("0x6d3b7741fef7cfb22d943cb1c3d221b71253acdf7e56925969d54a18a8566480"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x0"),
			Index:       1,
			Removed:     false,
		},
	}

	mockAvsSubscriber.EXPECT().ParseTaskResponded(taskResponseLog.Raw).Return(&taskResponseLog, nil)

	mockEthClient.EXPECT().TransactionByHash(
		context.Background(), taskResponseLog.Raw.TxHash,
	).Return(mocks.MockTransactionByHash(), true, nil)

	taskIndex := challenger.processTaskResponseLog(&taskResponseLog)
	assert.Equal(t, taskIndex, uint32(TASK_INDEX))

}

func createMockChallenger(mockCtrl *gomock.Controller) (*Challenger, *chainiomocks.MockAvsWriterer, *chainiomocks.MockAvsReaderer, *chainiomocks.MockAvsSubscriberer, *mockethclient.MockEthClient, error) {
	logger := sdklogging.NewNoopLogger()
	mockAvsWriter := chainiomocks.NewMockAvsWriterer(mockCtrl)
	mockAvsReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
	mockAvsSubscriber := chainiomocks.NewMockAvsSubscriberer(mockCtrl)
	mockEthClient := mockethclient.NewMockEthClient(mockCtrl)
	challenger := &Challenger{
		logger:             logger,
		avsWriter:          mockAvsWriter,
		avsReader:          mockAvsReader,
		ethClient:          mockEthClient,
		avsSubscriber:      mockAvsSubscriber,
		tasks:              make(map[uint32]cstaskmanager.IIncredibleSquaringTaskManagerTask),
		taskResponses:      make(map[uint32]chtypes.TaskResponseData),
		taskResponseChan:   make(chan *cstaskmanager.ContractIncredibleSquaringTaskManagerTaskResponded),
		newTaskCreatedChan: make(chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated),
	}
	return challenger, mockAvsWriter, mockAvsReader, mockAvsSubscriber, mockEthClient, nil
}
