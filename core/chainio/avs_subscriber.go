package chainio

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	cstaskmanager "github.com/inference-labs-inc/omron-avs/contracts/bindings/OmronTaskManager"
	"github.com/inference-labs-inc/omron-avs/core/config"
)

type AvsSubscriberer interface {
	SubscribeToNewTasks(newTaskCreatedChan chan *cstaskmanager.ContractOmronTaskManagerNewTaskCreated) event.Subscription
	SubscribeToTaskResponses(taskResponseLogs chan *cstaskmanager.ContractOmronTaskManagerTaskResponded) event.Subscription
	SubscribeToTaskChallenge(taskChallengeLogs chan *cstaskmanager.ContractOmronTaskManagerTaskChallenged) event.Subscription
	ParseTaskResponded(rawLog types.Log) (*cstaskmanager.ContractOmronTaskManagerTaskResponded, error)
}

// Subscribers use a ws connection instead of http connection like Readers
// kind of stupid that the geth client doesn't have a unified interface for both...
// it takes a single url, so the bindings, even though they have watcher functions, those can't be used
// with the http connection... seems very very stupid. Am I missing something?
type AvsSubscriber struct {
	AvsContractBindings *AvsManagersBindings
	logger              sdklogging.Logger
}

func BuildAvsSubscriberFromConfig(config *config.Config) (*AvsSubscriber, error) {
	return BuildAvsSubscriber(
		config.OmronRegistryCoordinatorAddr,
		config.OperatorStateRetrieverAddr,
		config.EthWsClient,
		config.Logger,
	)
}

func BuildAvsSubscriber(registryCoordinatorAddr, blsOperatorStateRetrieverAddr gethcommon.Address, ethclient eth.Client, logger sdklogging.Logger) (*AvsSubscriber, error) {
	avsContractBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, blsOperatorStateRetrieverAddr, ethclient, logger)
	if err != nil {
		logger.Errorf("Failed to create contract bindings", "err", err)
		return nil, err
	}
	return NewAvsSubscriber(avsContractBindings, logger), nil
}

func NewAvsSubscriber(avsContractBindings *AvsManagersBindings, logger sdklogging.Logger) *AvsSubscriber {
	return &AvsSubscriber{
		AvsContractBindings: avsContractBindings,
		logger:              logger,
	}
}

func (s *AvsSubscriber) SubscribeToNewTasks(newTaskCreatedChan chan *cstaskmanager.ContractOmronTaskManagerNewTaskCreated) event.Subscription {
	sub, err := s.AvsContractBindings.TaskManager.WatchNewTaskCreated(
		&bind.WatchOpts{}, newTaskCreatedChan, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to new TaskManager tasks", "err", err)
	}
	s.logger.Infof("Subscribed to new TaskManager tasks")
	return sub
}

func (s *AvsSubscriber) SubscribeToTaskResponses(taskResponseChan chan *cstaskmanager.ContractOmronTaskManagerTaskResponded) event.Subscription {
	sub, err := s.AvsContractBindings.TaskManager.WatchTaskResponded(
		&bind.WatchOpts{}, taskResponseChan,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to TaskResponded events", "err", err)
	}
	s.logger.Infof("Subscribed to TaskResponded events")
	return sub
}

func (s *AvsSubscriber) SubscribeToTaskChallenge(taskChallengedChan chan *cstaskmanager.ContractOmronTaskManagerTaskChallenged) event.Subscription {
	var taskIndexes []uint32
	sub, err := s.AvsContractBindings.TaskManager.WatchTaskChallenged(
		&bind.WatchOpts{}, taskChallengedChan, taskIndexes,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to TaskChallenged events", "err", err)
	}
	s.logger.Infof("Subscribed to TaskChallenged events")
	return sub
}

func (s *AvsSubscriber) ParseTaskResponded(rawLog types.Log) (*cstaskmanager.ContractOmronTaskManagerTaskResponded, error) {
	return s.AvsContractBindings.TaskManager.ContractOmronTaskManagerFilterer.ParseTaskResponded(rawLog)
}
