// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import {IBLSApkRegistry, IIndexRegistry, IStakeRegistry} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {SertnServiceManager, IServiceManager} from "../src/SertnServiceManager.sol";
import {SertnTaskManager} from "../src/SertnTaskManager.sol";
import {ISertnTaskManager} from "../src/ISertnTaskManager.sol";
import "../src/ERC20Mock.sol";

import {Halo2Verifier} from "../src/models/model_0/ZKVerifier.sol";
import {Halo2Verifier as Halo2Verifier_2} from "../src/models/model_1/ZKVerifier.sol";
import {ModelDB} from "../src/ModelDB.sol";

import {InferenceDB} from "../src/InferenceDB.sol";
import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/SertnDeployer.s.sol:SertnDeployer --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract UpdateContracts is Script, Utils {
    // DEPLOYMENT CONSTANTS
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 100;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 30;
    // TODO: right now hardcoding these (this address is anvil's default address 9)
    address public constant AGGREGATOR_ADDR =
        0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;
    address public constant TASK_GENERATOR_ADDR =
        0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;

    ERC20Mock public erc20Mock;
    StrategyBaseTVLLimits public erc20MockStrategy;

    // Sertn contracts
    ProxyAdmin public sertnProxyAdmin;
    PauserRegistry public sertnPauserReg;

    regcoord.RegistryCoordinator public registryCoordinator;
    regcoord.IRegistryCoordinator public registryCoordinatorImplementation;

    IBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;

    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;

    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;

    OperatorStateRetriever public operatorStateRetriever;

    SertnServiceManager public sertnServiceManager;
    IServiceManager public sertnServiceManagerImplementation;

    SertnTaskManager public sertnTaskManager;
    SertnTaskManager public sertnTaskManagerImplementation;

    function run() external {
        vm.startBroadcast(0x32530349A795378D7e2276670D9D991281F5BfCA);
        string memory avsDeployedContracts = readOutput(
            "sertn_avs_deployment_output"
        );

        // deploy proxy admin for ability to upgrade proxy contracts
        sertnProxyAdmin = ProxyAdmin(
            stdJson.readAddress(avsDeployedContracts, ".addresses.proxyAdmin")
        );

        registryCoordinator = regcoord.RegistryCoordinator(
            stdJson.readAddress(
                avsDeployedContracts,
                ".addresses.registryCoordinator"
            )
        );

        sertnTaskManager = SertnTaskManager(
            stdJson.readAddress(
                avsDeployedContracts,
                ".addresses.sertnTaskManager"
            )
        );

        operatorStateRetriever = OperatorStateRetriever(
            stdJson.readAddress(
                avsDeployedContracts,
                ".addresses.operatorStateRetriever"
            )
        );

        sertnTaskManagerImplementation = new SertnTaskManager(
            registryCoordinator
        );

        Halo2Verifier minModel = new Halo2Verifier();
        Halo2Verifier_2 maxModel = new Halo2Verifier_2();
        ModelDB modelDB = new ModelDB();

        modelDB.createNewModel(
            address(minModel),
            "Minimum",
            "This model returns the minimum value of the 5 inputs"
        );

        modelDB.createNewModel(
            address(maxModel),
            "Maximum",
            "This model returns the maximum value of the inputs"
        );

        InferenceDB inferenceDB = new InferenceDB(
            TASK_CHALLENGE_WINDOW_BLOCK,
            address(minModel)
        );
        inferenceDB.updateTaskManager(address(sertnTaskManager));

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        sertnProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(sertnTaskManager))),
            address(sertnTaskManagerImplementation)
        );
        vm.stopBroadcast();

        vm.startBroadcast(sertnTaskManager.generator());

        sertnTaskManager.setNewInferenceDB(address(inferenceDB));
        sertnTaskManager.setNewAggregator(AGGREGATOR_ADDR);

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";

        vm.serializeAddress(deployed_addresses, "model_0", address(minModel));

        vm.serializeAddress(deployed_addresses, "model_1", address(maxModel));

        vm.serializeAddress(
            deployed_addresses,
            "inferenceDB",
            address(inferenceDB)
        );

        vm.serializeAddress(deployed_addresses, "modelDB", address(modelDB));

        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "operatorStateRetriever",
            address(operatorStateRetriever)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "updated_contracts");
        vm.stopBroadcast();
    }
}
