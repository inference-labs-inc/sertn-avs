package integration_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/inference-labs-inc/zklayer-avs/aggregator"
	"github.com/inference-labs-inc/zklayer-avs/core/chainio"
	"github.com/inference-labs-inc/zklayer-avs/core/config"
	"github.com/inference-labs-inc/zklayer-avs/operator"
	"github.com/inference-labs-inc/zklayer-avs/types"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type IntegrationClients struct {
	Sdkclients clients.Clients
}

func TestIntegration(t *testing.T) {
	log.Println("This test takes ~50 seconds to run...")

	/* Start the anvil chain */
	anvilC := startAnvilTestContainer()
	// Not sure why but deferring anvilC.Terminate() causes a panic when the test finishes...
	// so letting it terminate silently for now
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "")
	if err != nil {
		t.Error(err)
	}

	/* Prepare the config file for aggregator */
	var aggConfigRaw config.ConfigRaw
	aggConfigFilePath := "../../config-files/aggregator.yaml"
	sdkutils.ReadYamlConfig(aggConfigFilePath, &aggConfigRaw)
	aggConfigRaw.EthRpcUrl = "http://" + anvilEndpoint
	aggConfigRaw.EthWsUrl = "ws://" + anvilEndpoint

	var zklayerDeploymentRaw config.ZklayerDeploymentRaw
	zklayerDeploymentFilePath := "../../contracts/script/output/31337/zklayer_avs_deployment_output.json"
	sdkutils.ReadJsonConfig(zklayerDeploymentFilePath, &zklayerDeploymentRaw)

	logger, err := sdklogging.NewZapLogger(aggConfigRaw.Environment)
	if err != nil {
		t.Fatalf("Failed to create logger: %s", err.Error())
	}
	ethRpcClient, err := eth.NewClient(aggConfigRaw.EthRpcUrl)
	if err != nil {
		t.Fatalf("Failed to create eth client: %s", err.Error())
	}
	ethWsClient, err := eth.NewClient(aggConfigRaw.EthWsUrl)
	if err != nil {
		t.Fatalf("Failed to create eth client: %s", err.Error())
	}

	aggregatorEcdsaPrivateKeyString := "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	if aggregatorEcdsaPrivateKeyString[:2] == "0x" {
		aggregatorEcdsaPrivateKeyString = aggregatorEcdsaPrivateKeyString[2:]
	}
	aggregatorEcdsaPrivateKey, err := crypto.HexToECDSA(aggregatorEcdsaPrivateKeyString)
	if err != nil {
		t.Fatalf("Cannot parse ecdsa private key: %s", err.Error())
	}
	aggregatorAddr, err := sdkutils.EcdsaPrivateKeyToAddress(aggregatorEcdsaPrivateKey)
	if err != nil {
		t.Fatalf("Cannot get operator address: %s", err.Error())
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		t.Fatalf("Cannot get chainId: %s", err.Error())
	}

	privateKeySigner, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: aggregatorEcdsaPrivateKey}, chainId)
	if err != nil {
		t.Fatalf("Cannot create signer: %s", err.Error())
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, privateKeySigner, aggregatorAddr, logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, aggregatorAddr)

	config := &config.Config{
		EcdsaPrivateKey:                aggregatorEcdsaPrivateKey,
		Logger:                         logger,
		EthHttpRpcUrl:                  aggConfigRaw.EthRpcUrl,
		EthHttpClient:                  ethRpcClient,
		EthWsRpcUrl:                    aggConfigRaw.EthWsUrl,
		EthWsClient:                    ethWsClient,
		OperatorStateRetrieverAddr:     common.HexToAddress(zklayerDeploymentRaw.Addresses.OperatorStateRetrieverAddr),
		ZklayerRegistryCoordinatorAddr: common.HexToAddress(zklayerDeploymentRaw.Addresses.RegistryCoordinatorAddr),
		AggregatorServerIpPortAddr:     aggConfigRaw.AggregatorServerIpPortAddr,
		RegisterOperatorOnStartup:      aggConfigRaw.RegisterOperatorOnStartup,
		TxMgr:                          txMgr,
		AggregatorAddress:              aggregatorAddr,
	}

	/* Prepare the config file for operator */
	nodeConfig := types.NodeConfig{}
	nodeConfigFilePath := "../../config-files/operator.anvil.yaml"
	err = sdkutils.ReadYamlConfig(nodeConfigFilePath, &nodeConfig)
	if err != nil {
		t.Fatalf("Failed to read yaml config: %s", err.Error())
	}
	/* Register operator*/
	// log.Println("registering operator for integration tests")
	// we need to do this dynamically and can't just hardcode a registered operator into the anvil
	// state because the anvil state dump doesn't also dump the receipts tree so we lose events,
	// and the aggregator thus can't get the operator's pubkey
	// operatorRegistrationCmd := exec.Command("bash", "./operator-registration.sh")
	// err = operatorRegistrationCmd.Run()
	// if err != nil {
	// 	t.Fatalf("Failed to register operator: %s", err.Error())
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 65*time.Second)
	defer cancel()
	/* start operator */
	// the passwords are set to empty strings
	log.Println("starting operator for integration tests")
	os.Setenv("OPERATOR_BLS_KEY_PASSWORD", "")
	os.Setenv("OPERATOR_ECDSA_KEY_PASSWORD", "")
	nodeConfig.BlsPrivateKeyStorePath = "../keys/test.bls.key.json"
	nodeConfig.EcdsaPrivateKeyStorePath = "../keys/test.ecdsa.key.json"
	nodeConfig.RegisterOperatorOnStartup = true
	nodeConfig.EthRpcUrl = "http://" + anvilEndpoint
	nodeConfig.EthWsUrl = "ws://" + anvilEndpoint
	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		t.Fatalf("Failed to create operator: %s", err.Error())
	}
	go operator.Start(ctx)
	log.Println("Started operator. Sleeping 15 seconds to give it time to register...")
	time.Sleep(15 * time.Second)

	/* start aggregator */
	log.Println("starting aggregator for integration tests")
	agg, err := aggregator.NewAggregator(config)
	if err != nil {
		t.Fatalf("Failed to create aggregator: %s", err.Error())
	}
	go agg.Start(ctx)
	log.Println("Started aggregator. Sleeping 20 seconds to give operator time to answer task 1...")
	time.Sleep(20 * time.Second)

	// get avsRegistry client to interact with the chain
	avsReader, err := chainio.BuildAvsReaderFromConfig(config)
	if err != nil {
		t.Fatalf("Cannot create AVS Reader: %s", err.Error())
	}

	// check if the task is recorded in the contract for task index 1
	taskHash, err := avsReader.AvsServiceBindings.TaskManager.AllTaskHashes(&bind.CallOpts{}, 1)
	if err != nil {
		t.Fatalf("Cannot get task hash: %s", err.Error())
	}
	if taskHash == [32]byte{} {
		t.Fatalf("Task hash is empty")
	}

	// check if the task response is recorded in the contract for task index 1
	taskResponseHash, err := avsReader.AvsServiceBindings.TaskManager.AllTaskResponses(&bind.CallOpts{}, 1)
	log.Printf("taskResponseHash: %v", taskResponseHash)
	if err != nil {
		t.Fatalf("Cannot get task response hash: %s", err.Error())
	}
	if taskResponseHash == [32]byte{} {
		t.Fatalf("Task response hash is empty")
	}

}

// TODO(samlaf): have to advance chain to a block where the task is answered
func startAnvilTestContainer() testcontainers.Container {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "ghcr.io/foundry-rs/foundry:latest",
		Mounts: testcontainers.ContainerMounts{
			testcontainers.ContainerMount{
				Source: testcontainers.GenericBindMountSource{
					HostPath: filepath.Join(integrationDir, "../anvil/avs-and-eigenlayer-deployed-anvil-state.json"),
				},
				Target: "/root/.anvil/state.json",
			},
		},
		Entrypoint:   []string{"anvil"},
		Cmd:          []string{"--host", "0.0.0.0", "--load-state", "/root/.anvil/state.json"},
		ExposedPorts: []string{"8545/tcp"},
		WaitingFor:   wait.ForLog("Listening on"),
	}
	anvilC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	// this is needed temporarily because anvil restarts at 0 block when we load a state...
	// see comment in start-anvil-chain-with-el-and-avs-deployed.sh
	advanceChain(anvilC)
	return anvilC
}

func advanceChain(anvilC testcontainers.Container) {
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "")
	if err != nil {
		panic(err)
	}
	rpcUrl := "http://" + anvilEndpoint
	privateKey := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	cmd := exec.Command("bash", "-c",
		fmt.Sprintf(
			`forge script script/utils/Utils.sol --sig "advanceChainByNBlocks(uint256)" 100 --rpc-url %s --private-key %s --broadcast`,
			rpcUrl, privateKey),
	)
	cmd.Dir = "../../contracts"
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
