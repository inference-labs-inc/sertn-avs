package aggregator

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
)

func TestProcessSignedTaskResponse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)
	var NUMBER_TO_SQUARE = uint32(3)

	MOCK_OPERATOR_BLS_PRIVATE_KEY, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING)
	assert.Nil(t, err)
	MOCK_OPERATOR_KEYPAIR := bls.NewKeyPair(MOCK_OPERATOR_BLS_PRIVATE_KEY)
	MOCK_OPERATOR_G1PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG1()
	MOCK_OPERATOR_G2PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG2()

	operatorPubkeyDict := map[sdktypes.OperatorId]types.OperatorInfo{
		MOCK_OPERATOR_ID: {
			OperatorPubkeys: sdktypes.OperatorPubkeys{
				G1Pubkey: MOCK_OPERATOR_G1PUBKEY,
				G2Pubkey: MOCK_OPERATOR_G2PUBKEY,
			},
			OperatorAddr: common.Address{},
		},
	}
	aggregator, _, mockBlsAggServ, err := createMockAggregator(mockCtrl, operatorPubkeyDict)
	assert.Nil(t, err)

	signedTaskResponse, err := createMockSignedTaskResponse(MockTask{
		TaskNum:        TASK_INDEX,
		BlockNumber:    BLOCK_NUMBER,
		NumberToSquare: NUMBER_TO_SQUARE,
	}, *MOCK_OPERATOR_KEYPAIR)
	assert.Nil(t, err)
	signedTaskResponseDigest, err := core.GetTaskResponseDigest(&signedTaskResponse.TaskResponse)
	assert.Nil(t, err)

	// TODO(samlaf): is this the right way to test writing to external service?
	// or is there some wisdom to "don't mock 3rd party code"?
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockBlsAggServ.EXPECT().ProcessNewSignature(context.Background(), TASK_INDEX, signedTaskResponseDigest,
		&signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId)
	err = aggregator.ProcessSignedTaskResponse(signedTaskResponse, nil)
	assert.Nil(t, err)
}

// mocks an operator signing on a task response
func createMockSignedTaskResponse(mockTask MockTask, keypair bls.KeyPair) (*SignedTaskResponse, error) {
	numberToSquareBigInt := big.NewInt(int64(mockTask.NumberToSquare))
	taskResponse := &cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
		ReferenceTaskIndex: mockTask.TaskNum,
		NumberSquared:      numberToSquareBigInt.Mul(numberToSquareBigInt, numberToSquareBigInt),
	}
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		return nil, err
	}
	blsSignature := keypair.SignMessage(taskResponseHash)
	signedTaskResponse := &SignedTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   MOCK_OPERATOR_ID,
	}
	return signedTaskResponse, nil
}
