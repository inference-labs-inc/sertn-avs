package challenger

import (
	"bytes"
	"context"
	"math/big"
	"os/exec"
	"strconv"
	"strings"
	"time"

	_ "embed"

	ethclient "github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	sertnCommon "github.com/inference-labs-inc/sertn-avs/common"
	"github.com/inference-labs-inc/sertn-avs/core"
	"github.com/inference-labs-inc/sertn-avs/core/config"

	"github.com/inference-labs-inc/sertn-avs/challenger/types"
	cstaskmanager "github.com/inference-labs-inc/sertn-avs/contracts/bindings/SertnTaskManager"
	"github.com/inference-labs-inc/sertn-avs/core/chainio"
)

type Challenger struct {
	logger             logging.Logger
	ethClient          ethclient.Client
	avsReader          chainio.AvsReaderer
	avsWriter          chainio.AvsWriterer
	avsSubscriber      chainio.AvsSubscriberer
	tasks              map[uint32]cstaskmanager.ITaskStructTask
	taskResponses      map[uint32]types.TaskResponseData
	taskResponseChan   chan *cstaskmanager.ContractSertnTaskManagerTaskResponded
	newTaskCreatedChan chan *cstaskmanager.ContractSertnTaskManagerNewTaskCreated
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
		tasks:              make(map[uint32]cstaskmanager.ITaskStructTask),
		taskResponses:      make(map[uint32]types.TaskResponseData),
		taskResponseChan:   make(chan *cstaskmanager.ContractSertnTaskManagerTaskResponded),
		newTaskCreatedChan: make(chan *cstaskmanager.ContractSertnTaskManagerNewTaskCreated),
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

func (c *Challenger) processNewTaskCreatedLog(newTaskCreatedLog *cstaskmanager.ContractSertnTaskManagerNewTaskCreated) uint32 {
	c.tasks[newTaskCreatedLog.TaskIndex] = newTaskCreatedLog.Task
	return newTaskCreatedLog.TaskIndex
}

func (c *Challenger) processTaskResponseLog(taskResponseLog *cstaskmanager.ContractSertnTaskManagerTaskResponded) uint32 {
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
		instances[i] = inputs[i]
	}
	instances[5] = output
	return instances
}

func (c *Challenger) RunModelFromBigIntInputs(rawInputs [5]*big.Int, modelVerifier common.Address) *big.Int {
	inputs := core.FormatBigIntInputsToString(rawInputs)
	cmd := exec.Command("python", core.RelativeUrl("python/run.py"), "-i", inputs, "-m", modelVerifier.Hex())

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		c.logger.Error(err.Error())
	}
	// convert output string into int64
	output, err := strconv.ParseInt(strings.Split(string(stdout), "\n")[0], 10, 64)
	if err != nil {
		c.logger.Error("Error parsing integer from python response", "error", err.Error(), "python output", string(stdout))
	}

	return big.NewInt(output)
}

func (c *Challenger) callChallengeModule(taskIndex uint32) error {
	response := c.taskResponses[taskIndex].TaskResponse.Output
	output := c.RunModelFromBigIntInputs(c.tasks[taskIndex].Inputs, c.tasks[taskIndex].ModelVerifier)
	c.logger.Info("CHALLENGER - OPERATOR-OUTPUT", "response", response)
	c.logger.Info("CHALLENGER - REAL-OUTPUT", "output", output)
	// checking if the answer in the response submitted by aggregator is correct
	if output.Cmp(response) != 0 || true {
		c.logger.Infof("OUTPUT FROM OPERATOR INCORRECT")
		// raise challenge
		c.raiseChallenge(taskIndex, output)
		return nil
	}
	c.logger.Infof("OUTPUT FROM OPERATOR CORRECT")
	return nil
	// return types.NoErrorInTaskResponse
}

func (c *Challenger) getNonSigningOperatorPubKeys(vLog *cstaskmanager.ContractSertnTaskManagerTaskResponded) []cstaskmanager.BN254G1Point {
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

	cstmAbi, err := abi.JSON(bytes.NewReader(sertnCommon.TaskManagerAbi))

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

func (c *Challenger) raiseChallenge(taskIndex uint32, output *big.Int) error {
	c.logger.Info("Challenger raising challenge.", "taskIndex", taskIndex)
	c.logger.Info("Task", "Task", c.tasks[taskIndex])
	c.logger.Info("TaskResponse", "TaskResponse", c.taskResponses[taskIndex].TaskResponse)
	c.logger.Info("TaskResponseMetadata", "TaskResponseMetadata", c.taskResponses[taskIndex].TaskResponseMetadata)
	c.logger.Info("NonSigningOperatorPubKeys", "NonSigningOperatorPubKeys", c.taskResponses[taskIndex].NonSigningOperatorPubKeys)

	inputs := core.FormatBigIntInputsToString(c.tasks[taskIndex].Inputs)
	c.logger.Info("Task Inputs", "Task inputs to prove", inputs)

	instances := c.FormatInstancesForSolidity(c.tasks[taskIndex].Inputs, output)

	c.logger.Info("CHALLENGER - INSTANCES", "instances", instances)

	receipt, err := c.avsWriter.RaiseChallenge(
		context.Background(),
		c.tasks[taskIndex],
		c.taskResponses[taskIndex].TaskResponse,
		c.taskResponses[taskIndex].TaskResponseMetadata,
	)

	if err != nil {
		c.logger.Error("Challenger failed to raise challenge:", "err", err)
		return err
	}
	c.logger.Infof("Tx hash of the challenge tx: %v", receipt.TxHash.Hex())
	go c.confirmChallenge(taskIndex)
	return nil
}

func (c *Challenger) confirmChallenge(taskIndex uint32) error {
	time.Sleep(time.Second * 10)

	receipt, err := c.avsWriter.ConfirmChallenge(
		context.Background(),
		c.tasks[taskIndex],
		c.taskResponses[taskIndex].TaskResponse,
		c.taskResponses[taskIndex].TaskResponseMetadata,
		c.taskResponses[taskIndex].NonSigningOperatorPubKeys,
	)

	if err != nil {
		c.logger.Error("Challenger failed to confirm challenge:", "err", err)
		return err
	}
	c.logger.Infof("Tx hash of the confirmed challenge tx: %v", receipt.TxHash.Hex())

	return nil
}
