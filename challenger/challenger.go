package challenger

import (
	"bytes"
	"context"
	"encoding/hex"
	"math/big"
	"os/exec"
	"strconv"
	"strings"

	_ "embed"

	ethclient "github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/inference-labs-inc/omron-avs/common"
	"github.com/inference-labs-inc/omron-avs/core"
	"github.com/inference-labs-inc/omron-avs/core/config"

	"github.com/inference-labs-inc/omron-avs/challenger/types"
	cstaskmanager "github.com/inference-labs-inc/omron-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/inference-labs-inc/omron-avs/core/chainio"
)

type Challenger struct {
	logger             logging.Logger
	ethClient          ethclient.Client
	avsReader          chainio.AvsReaderer
	avsWriter          chainio.AvsWriterer
	avsSubscriber      chainio.AvsSubscriberer
	tasks              map[uint32]cstaskmanager.IIncredibleSquaringTaskManagerTask
	taskResponses      map[uint32]types.TaskResponseData
	taskResponseChan   chan *cstaskmanager.ContractIncredibleSquaringTaskManagerTaskResponded
	newTaskCreatedChan chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated
}

func NewChallenger(c *config.Config) (*Challenger, error) {

	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create EthWriter", "err", err)
		return nil, err
	}
	avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create EthReader", "err", err)
		return nil, err
	}
	avsSubscriber, err := chainio.BuildAvsSubscriberFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create EthSubscriber", "err", err)
		return nil, err
	}

	challenger := &Challenger{
		ethClient:          c.EthHttpClient,
		logger:             c.Logger,
		avsWriter:          avsWriter,
		avsReader:          avsReader,
		avsSubscriber:      avsSubscriber,
		tasks:              make(map[uint32]cstaskmanager.IIncredibleSquaringTaskManagerTask),
		taskResponses:      make(map[uint32]types.TaskResponseData),
		taskResponseChan:   make(chan *cstaskmanager.ContractIncredibleSquaringTaskManagerTaskResponded),
		newTaskCreatedChan: make(chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated),
	}

	return challenger, nil
}

func (c *Challenger) Start(ctx context.Context) error {
	c.logger.Infof("Starting Challenger.")

	newTaskSub := c.avsSubscriber.SubscribeToNewTasks(c.newTaskCreatedChan)
	c.logger.Infof("Subscribed to new tasks")

	taskResponseSub := c.avsSubscriber.SubscribeToTaskResponses(c.taskResponseChan)
	c.logger.Infof("Subscribed to task responses")

	for {
		select {
		case err := <-newTaskSub.Err():
			// TODO(samlaf): Copied from operator. There was a comment about this on when should exactly do these errors occur? do we need to restart the websocket
			c.logger.Error("Error in websocket subscription for new Task", "err", err)
			newTaskSub.Unsubscribe()
			newTaskSub = c.avsSubscriber.SubscribeToNewTasks(c.newTaskCreatedChan)

		case err := <-taskResponseSub.Err():
			// TODO(samlaf): Copied from operator. There was a comment about this on when should exactly do these errors occur? do we need to restart the websocket
			c.logger.Error("Error in websocket subscription for task response", "err", err)
			taskResponseSub.Unsubscribe()
			taskResponseSub = c.avsSubscriber.SubscribeToTaskResponses(c.taskResponseChan)

		case newTaskCreatedLog := <-c.newTaskCreatedChan:
			c.logger.Info("New task created log received", "newTaskCreatedLog", newTaskCreatedLog)
			taskIndex := c.processNewTaskCreatedLog(newTaskCreatedLog)

			if _, found := c.taskResponses[taskIndex]; found {
				err := c.callChallengeModule(taskIndex)
				if err != nil {
					c.logger.Error("Error calling challenge module", "err", err)
				}
				continue
			}

		case taskResponseLog := <-c.taskResponseChan:
			c.logger.Info("Task response log received", "taskResponseLog", taskResponseLog)
			taskIndex := c.processTaskResponseLog(taskResponseLog)

			if _, found := c.tasks[taskIndex]; found {
				err := c.callChallengeModule(taskIndex)
				if err != nil {
					c.logger.Info("Info:", "err", err)
				}
				continue
			}
		}
	}

}

func (c *Challenger) processNewTaskCreatedLog(newTaskCreatedLog *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated) uint32 {
	c.tasks[newTaskCreatedLog.TaskIndex] = newTaskCreatedLog.Task
	return newTaskCreatedLog.TaskIndex
}

func (c *Challenger) processTaskResponseLog(taskResponseLog *cstaskmanager.ContractIncredibleSquaringTaskManagerTaskResponded) uint32 {
	taskResponseRawLog, err := c.avsSubscriber.ParseTaskResponded(taskResponseLog.Raw)
	if err != nil {
		c.logger.Error("Error parsing task response. skipping task (this is probably bad and should be investigated)", "err", err)
	}

	// get the inputs necessary for raising a challenge
	nonSigningOperatorPubKeys := c.getNonSigningOperatorPubKeys(taskResponseLog)

	taskResponseData := types.TaskResponseData{
		TaskResponse:              taskResponseLog.TaskResponse,
		TaskResponseMetadata:      taskResponseLog.TaskResponseMetadata,
		NonSigningOperatorPubKeys: nonSigningOperatorPubKeys,
	}

	c.taskResponses[taskResponseRawLog.TaskResponse.ReferenceTaskIndex] = taskResponseData
	return taskResponseRawLog.TaskResponse.ReferenceTaskIndex
}

func (c *Challenger) FormatInstancesForSolidity(inputs [5]*big.Int, output *big.Int) [6]*big.Int {
	var instances [6]*big.Int
	for i := 0; i < 5; i++ {
		instances[i] = big.NewInt(0).Mul(inputs[i], big.NewInt(4))
	}
	instances[5] = output
	return instances
}

func (c *Challenger) OutputAndProofFromInputs(inputs string) (*big.Int, []byte) {
	cmd := exec.Command("python", "python/prove.py", "--input", inputs)
	stdout, err := cmd.Output()
	if err != nil {
		c.logger.Error("Challenger failed to prove computation:", "err", err)
	}

	result := string(stdout)
	instancesAndProof := strings.Split(result, "\n")

	proof, _ := hex.DecodeString(instancesAndProof[1])
	output, _ := strconv.ParseInt(instancesAndProof[0], 16, 64)

	return big.NewInt(output), proof
}

func (c *Challenger) callChallengeModule(taskIndex uint32) error {
	inputs := core.FormatBigIntInputsToString(c.tasks[taskIndex].Inputs)
	responce := big.NewInt(0).Mul(c.taskResponses[taskIndex].TaskResponse.Output, big.NewInt(4))
	output, proof := c.OutputAndProofFromInputs(inputs)

	//checking if the answer in the response submitted by aggregator is correct
	if output.Cmp(responce) != 0 {
		c.logger.Infof("The output from operator was not correct")
		// raise challenge
		c.raiseChallenge(taskIndex, output, proof)
		return nil
	}
	c.logger.Infof("The ouput from operator was correct")
	return nil
	//return types.NoErrorInTaskResponse
}

func (c *Challenger) getNonSigningOperatorPubKeys(vLog *cstaskmanager.ContractIncredibleSquaringTaskManagerTaskResponded) []cstaskmanager.BN254G1Point {
	c.logger.Info("vLog.Raw is", "vLog.Raw", vLog.Raw)

	// get the nonSignerStakesAndSignature
	txHash := vLog.Raw.TxHash
	c.logger.Info("txHash", "txHash", txHash)
	tx, _, err := c.ethClient.TransactionByHash(context.Background(), txHash)
	_ = tx
	if err != nil {
		c.logger.Error("Error getting transaction by hash",
			"txHash", txHash,
			"err", err,
		)
	}
	calldata := tx.Data()
	c.logger.Info("calldata", "calldata", calldata)

	cstmAbi, err := abi.JSON(bytes.NewReader(common.TaskManagerAbi))

	if err != nil {
		c.logger.Error("Error getting Abi", "err", err)
	}
	methodSig := calldata[:4]
	c.logger.Info("methodSig", "methodSig", methodSig)
	method, err := cstmAbi.MethodById(methodSig)
	if err != nil {
		c.logger.Error("Error getting method", "err", err)
	}

	inputs, err := method.Inputs.Unpack(calldata[4:])
	if err != nil {
		c.logger.Error("Error unpacking calldata", "err", err)
	}

	nonSignerStakesAndSignatureInput := inputs[2].(struct {
		NonSignerQuorumBitmapIndices []uint32 "json:\"nonSignerQuorumBitmapIndices\""
		NonSignerPubkeys             []struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"nonSignerPubkeys\""
		QuorumApks []struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"quorumApks\""
		ApkG2 struct {
			X [2]*big.Int "json:\"X\""
			Y [2]*big.Int "json:\"Y\""
		} "json:\"apkG2\""
		Sigma struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"sigma\""
		QuorumApkIndices      []uint32   "json:\"quorumApkIndices\""
		TotalStakeIndices     []uint32   "json:\"totalStakeIndices\""
		NonSignerStakeIndices [][]uint32 "json:\"nonSignerStakeIndices\""
	})

	// get pubkeys of non-signing operators and submit them to the contract
	nonSigningOperatorPubKeys := make([]cstaskmanager.BN254G1Point, len(nonSignerStakesAndSignatureInput.NonSignerPubkeys))
	for i, pubkey := range nonSignerStakesAndSignatureInput.NonSignerPubkeys {
		nonSigningOperatorPubKeys[i] = cstaskmanager.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}
	}

	return nonSigningOperatorPubKeys
}

func (c *Challenger) raiseChallenge(taskIndex uint32, output *big.Int, proof []byte) error {
	c.logger.Info("Challenger raising challenge.", "taskIndex", taskIndex)
	c.logger.Info("Task", "Task", c.tasks[taskIndex])
	c.logger.Info("TaskResponse", "TaskResponse", c.taskResponses[taskIndex].TaskResponse)
	c.logger.Info("TaskResponseMetadata", "TaskResponseMetadata", c.taskResponses[taskIndex].TaskResponseMetadata)
	c.logger.Info("NonSigningOperatorPubKeys", "NonSigningOperatorPubKeys", c.taskResponses[taskIndex].NonSigningOperatorPubKeys)

	inputs := core.FormatBigIntInputsToString(c.tasks[taskIndex].Inputs)
	c.logger.Info("Task Inputs", "Task inputs to prove", inputs)

	instances := c.FormatInstancesForSolidity(c.tasks[taskIndex].Inputs, output)

	c.logger.Info("Instances", "instances", instances)
	c.logger.Info("Proof", "proof", instances)

	receipt, err := c.avsWriter.RaiseChallenge(
		context.Background(),
		c.tasks[taskIndex],
		c.taskResponses[taskIndex].TaskResponse,
		c.taskResponses[taskIndex].TaskResponseMetadata,
		c.taskResponses[taskIndex].NonSigningOperatorPubKeys,
		instances[:], // need to convert from array to slice
		proof,
	)
	if err != nil {
		c.logger.Error("Challenger failed to raise challenge:", "err", err)
		return err
	}
	c.logger.Infof("Tx hash of the challenge tx: %v", receipt.TxHash.Hex())
	return nil
}
