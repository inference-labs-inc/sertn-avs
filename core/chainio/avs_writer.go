package chainio

import (
	"context"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	cstaskmanager "github.com/inference-labs-inc/zklayer-avs/contracts/bindings/ZklayerTaskManager"
	"github.com/inference-labs-inc/zklayer-avs/core/config"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	SendNewTaskInput(
		ctx context.Context,
		inputs [5]*big.Int,
		quorumThresholdPercentage sdktypes.QuorumThresholdPercentage,
		quorumNumbers sdktypes.QuorumNums,
	) (cstaskmanager.IZklayerTaskManagerTask, uint32, error)
	RaiseChallenge(
		ctx context.Context,
		task cstaskmanager.IZklayerTaskManagerTask,
		taskResponse cstaskmanager.IZklayerTaskManagerTaskResponse,
		taskResponseMetadata cstaskmanager.IZklayerTaskManagerTaskResponseMetadata,
	) (*types.Receipt, error)
	SendAggregatedResponse(ctx context.Context,
		task cstaskmanager.IZklayerTaskManagerTask,
		taskResponse cstaskmanager.IZklayerTaskManagerTaskResponse,
		nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (*types.Receipt, error)
	ConfirmChallenge(
		ctx context.Context,
		task cstaskmanager.IZklayerTaskManagerTask,
		taskResponse cstaskmanager.IZklayerTaskManagerTaskResponse,
		taskResponseMetadata cstaskmanager.IZklayerTaskManagerTaskResponseMetadata,
		pubkeysOfNonSigningOperators []cstaskmanager.BN254G1Point,
	) (*types.Receipt, error)
	ProveResponseAccurate(
		ctx context.Context,
		taskId uint32,
		instances []*big.Int,
		proof []byte,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsregistry.AvsRegistryWriter
	AvsContractBindings *AvsManagersBindings
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	client              eth.Client
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return BuildAvsWriter(c.TxMgr, c.ZklayerRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.Client, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}
	return NewAvsWriter(avsRegistryWriter, avsServiceBindings, logger, txMgr), nil
}
func NewAvsWriter(avsRegistryWriter avsregistry.AvsRegistryWriter, avsServiceBindings *AvsManagersBindings, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		logger:              logger,
		TxMgr:               txMgr,
	}
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewTaskInput(ctx context.Context, inputs [5]*big.Int, quorumThresholdPercentage sdktypes.QuorumThresholdPercentage, quorumNumbers sdktypes.QuorumNums) (cstaskmanager.IZklayerTaskManagerTask, uint32, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return cstaskmanager.IZklayerTaskManagerTask{}, 0, err
	}
	tx, err := w.AvsContractBindings.TaskManager.CreateNewTask(txOpts, inputs, uint32(quorumThresholdPercentage), quorumNumbers.UnderlyingType())
	if err != nil {
		w.logger.Errorf("Error assembling CreateNewTask tx")
		return cstaskmanager.IZklayerTaskManagerTask{}, 0, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateNewTask tx")
		return cstaskmanager.IZklayerTaskManagerTask{}, 0, err
	}
	newTaskCreatedEvent, err := w.AvsContractBindings.TaskManager.ContractZklayerTaskManagerFilterer.ParseNewTaskCreated(*receipt.Logs[0])
	if err != nil {
		w.logger.Error("Aggregator failed to parse new task created event", "err", err)
		return cstaskmanager.IZklayerTaskManagerTask{}, 0, err
	}
	return newTaskCreatedEvent.Task, newTaskCreatedEvent.TaskIndex, nil
}

func (w *AvsWriter) SendAggregatedResponse(
	ctx context.Context, task cstaskmanager.IZklayerTaskManagerTask,
	taskResponse cstaskmanager.IZklayerTaskManagerTaskResponse,
	nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RespondToTask(txOpts, task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting respondToTask tx")
		return nil, err
	}
	return receipt, nil
}

func (w *AvsWriter) RaiseChallenge(
	ctx context.Context,
	task cstaskmanager.IZklayerTaskManagerTask,
	taskResponse cstaskmanager.IZklayerTaskManagerTaskResponse,
	taskResponseMetadata cstaskmanager.IZklayerTaskManagerTaskResponseMetadata,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RaiseChallenger(txOpts, task, taskResponse, taskResponseMetadata)
	if err != nil {
		w.logger.Errorf("Error assembling RaiseChallenge tx")
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting RaiseChallenge tx")
		return nil, err
	}
	return receipt, nil
}

func (w *AvsWriter) ProveResponseAccurate(
	ctx context.Context,
	taskId uint32,
	instances []*big.Int,
	proof []byte,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.ProveResultAccurate(txOpts, taskId, instances[:], proof)
	if err != nil {
		w.logger.Errorf("Error assembling ProveResultAccurate tx")
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting ProveResultAccurate tx")
		return nil, err
	}
	return receipt, nil
}

func (w *AvsWriter) ConfirmChallenge(
	ctx context.Context,
	task cstaskmanager.IZklayerTaskManagerTask,
	taskResponse cstaskmanager.IZklayerTaskManagerTaskResponse,
	taskResponseMetadata cstaskmanager.IZklayerTaskManagerTaskResponseMetadata,
	pubkeysOfNonSigningOperators []cstaskmanager.BN254G1Point,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.ConfirmChallenge(txOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
	if err != nil {
		w.logger.Errorf("Error assembling ConfirmChallenge tx")
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting ConfirmChallenge tx")
		return nil, err
	}
	return receipt, nil
}
