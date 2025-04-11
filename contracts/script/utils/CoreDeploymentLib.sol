// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import {console2} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {DelegationManager} from "@eigenlayer/contracts/core/DelegationManager.sol";
import {StrategyManager} from "@eigenlayer/contracts/core/StrategyManager.sol";
//Adding allocation manager
import {AllocationManager} from "@eigenlayer/contracts/core/AllocationManager.sol";
//Adding permission controller
import {PermissionController} from "@eigenlayer/contracts/permissions/PermissionController.sol";
import {EigenPodManager} from "@eigenlayer/contracts/pods/EigenPodManager.sol";
import {RewardsCoordinator} from "@eigenlayer/contracts/core/RewardsCoordinator.sol";
import {StrategyBase} from "@eigenlayer/contracts/strategies/StrategyBase.sol";
import {EigenPod} from "@eigenlayer/contracts/pods/EigenPod.sol";
import {IETHPOSDeposit} from "@eigenlayer/contracts/interfaces/IETHPOSDeposit.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import {PauserRegistry} from "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IBeacon} from "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import {IStrategyManager} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {IEigenPodManager} from "@eigenlayer/contracts/interfaces/IEigenPodManager.sol";
//Adding Allocation Manager interface
import {IAllocationManager} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
//Adding Permission controller interface
import {IPermissionController} from "@eigenlayer/contracts/interfaces/IPermissionController.sol";
import {IPauserRegistry} from "@eigenlayer/contracts/interfaces/IPauserRegistry.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";
import {IRewardsCoordinatorTypes} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";

import {UpgradeableProxyLib} from "./UpgradeableProxyLib.sol";

library CoreDeploymentLib {
    using stdJson for *;
    using Strings for *;
    using UpgradeableProxyLib for address;

    Vm internal constant vm =
        Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    struct StrategyManagerConfig {
        uint256 initPausedStatus;
    }

    struct DelegationManagerConfig {
        uint256 initPausedStatus;
        uint32 withdrawalDelayBlocks;
    }

    struct AllocationManagerConfig {
        uint256 initPausedStatus;
        uint32 deallocationDelay;
        uint32 allocationConfigurationDelay;
    }

    struct EigenPodManagerConfig {
        uint256 initPausedStatus;
    }

    struct RewardsCoordinatorConfig {
        uint256 initPausedStatus;
        uint32 maxRewardsDuration;
        uint32 maxRetroactiveLength;
        uint32 maxFutureLength;
        uint32 genesisRewardsTimestamp;
        address updater;
        uint256 updaterKey;
        uint256 activationDelay;
        uint32 calculationIntervalSeconds;
        uint256 globalOperatorCommissionBips;
    }

    struct StrategyFactoryConfig {
        uint256 initPausedStatus;
    }

    struct DeploymentData {
        address delegationManager;
        address allocationManager;
        address strategyManager;
        address eigenPodManager;
        address rewardsCoordinator;
        address eigenPodBeacon;
        address pauserRegistry;
        address permissionController;
        address strategyFactory;
        address strategyBeacon;
        address initialOwner;
    }

    function deployContracts(
        address proxyAdmin,
        DeploymentConfigData memory configData
    ) internal returns (DeploymentData memory) {
        DeploymentData memory result;

        result.delegationManager = UpgradeableProxyLib.setUpEmptyProxy(
            proxyAdmin
        );
        result.allocationManager = UpgradeableProxyLib.setUpEmptyProxy(
            proxyAdmin
        );
        result.permissionController = UpgradeableProxyLib.setUpEmptyProxy(
            proxyAdmin
        );
        result.strategyManager = UpgradeableProxyLib.setUpEmptyProxy(
            proxyAdmin
        );
        result.eigenPodManager = UpgradeableProxyLib.setUpEmptyProxy(
            proxyAdmin
        );
        result.rewardsCoordinator = UpgradeableProxyLib.setUpEmptyProxy(
            proxyAdmin
        );
        result.eigenPodBeacon = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.pauserRegistry = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.strategyFactory = UpgradeableProxyLib.setUpEmptyProxy(
            proxyAdmin
        );

        // Deploy the implementation contracts, using the proxy contracts as inputs
        address delegationManagerImpl = address(
            new DelegationManager(
                IStrategyManager(result.strategyManager),
                IEigenPodManager(result.eigenPodManager),
                IAllocationManager(result.allocationManager),
                IPauserRegistry(result.pauserRegistry),
                IPermissionController(result.permissionController),
                configData.delegationManager.withdrawalDelayBlocks,
                //Middleware will likely update to need version in constructor
                configData.version
            )
        );

        //Allocation Manager
        address allocationManagerImpl = address(
            new AllocationManager(
                IDelegationManager(result.delegationManager),
                IPauserRegistry(result.pauserRegistry),
                IPermissionController(result.permissionController),
                configData.allocationManager.deallocationDelay,
                configData.allocationManager.allocationConfigurationDelay,
                configData.version
            )
        );

        string memory _strategyVersion = "";

        address strategyManagerImpl = address(
            new StrategyManager(
                IDelegationManager(result.delegationManager),
                IPauserRegistry(result.pauserRegistry),
                configData.version
            )
        );

        address strategyFactoryImpl = address(
            new StrategyFactory(
                IStrategyManager(result.strategyManager),
                IPauserRegistry(result.pauserRegistry),
                configData.version
            )
        );

        address ethPOSDeposit;
        if (block.chainid == 1) {
            ethPOSDeposit = 0x00000000219ab540356cBB839Cbe05303d7705Fa;
        } else {
            // For non-mainnet chains, you might want to deploy a mock or read from a config
            // This assumes you have a similar config setup as in M2_Deploy_From_Scratch.s.sol
            /// TODO: Handle Eth pos
        }

        address eigenPodManagerImpl = address(
            new EigenPodManager(
                IETHPOSDeposit(ethPOSDeposit),
                IBeacon(result.eigenPodBeacon),
                IDelegationManager(result.delegationManager),
                IPauserRegistry(result.pauserRegistry),
                configData.version
            )
        );

        address rewardsCoordinatorImpl = address(
            new RewardsCoordinator(
                IRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams(
                    IDelegationManager(result.delegationManager),
                    IStrategyManager(result.strategyManager),
                    IAllocationManager(result.allocationManager),
                    IPauserRegistry(result.pauserRegistry),
                    IPermissionController(result.permissionController),
                    configData.rewardsCoordinator.calculationIntervalSeconds,
                    configData.rewardsCoordinator.maxRewardsDuration,
                    configData.rewardsCoordinator.maxRetroactiveLength,
                    configData.rewardsCoordinator.maxFutureLength,
                    configData.rewardsCoordinator.genesisRewardsTimestamp,
                    configData.version
                )
            )
        );

        /// TODO: Get actual genesis time
        uint64 GENESIS_TIME = 1_564_000;

        address eigenPodImpl = address(
            new EigenPod(
                IETHPOSDeposit(ethPOSDeposit),
                IEigenPodManager(result.eigenPodManager),
                GENESIS_TIME,
                configData.version
            )
        );

        address eigenPodBeaconImpl = address(
            new UpgradeableBeacon(eigenPodImpl)
        );

        address baseStrategyImpl = address(
            new StrategyBase(
                IStrategyManager(result.strategyManager),
                IPauserRegistry(result.pauserRegistry),
                configData.version
            )
        );
        /// TODO: PauserRegistry isn't upgradeable
        address pauserRegistryImpl = address(
            new PauserRegistry(
                new address[](0), // Empty array for pausers
                proxyAdmin // ProxyAdmin as the unpauser
            )
        );

        address permissionControllerImpl = address(
            new PermissionController(configData.version)
        );

        // Deploy and configure the strategy beacon
        result.strategyBeacon = address(
            new UpgradeableBeacon(baseStrategyImpl)
        );

        // Upgrade contracts
        /// TODO: Get from config
        bytes memory upgradeCall = abi.encodeCall(
            DelegationManager.initialize,
            (proxyAdmin, configData.delegationManager.initPausedStatus)
        );
        UpgradeableProxyLib.upgradeAndCall(
            result.delegationManager,
            delegationManagerImpl,
            upgradeCall
        );

        upgradeCall = abi.encodeCall(
            AllocationManager.initialize,
            (proxyAdmin, configData.allocationManager.initPausedStatus)
        );
        UpgradeableProxyLib.upgradeAndCall(
            result.allocationManager,
            allocationManagerImpl,
            upgradeCall
        );

        // Upgrade StrategyManager contract
        upgradeCall = abi.encodeCall(
            StrategyManager.initialize,
            (
                proxyAdmin, // initialOwner
                result.strategyFactory, // initialStrategyWhitelister
                configData.strategyManager.initPausedStatus // initialPausedStatus
            )
        );
        UpgradeableProxyLib.upgradeAndCall(
            result.strategyManager,
            strategyManagerImpl,
            upgradeCall
        );

        // Upgrade StrategyFactory contract
        upgradeCall = abi.encodeCall(
            StrategyFactory.initialize,
            (
                proxyAdmin, // initialOwner
                configData.strategyFactory.initPausedStatus, // initialPausedStatus
                IBeacon(result.strategyBeacon)
            )
        );
        UpgradeableProxyLib.upgradeAndCall(
            result.strategyFactory,
            strategyFactoryImpl,
            upgradeCall
        );

        // Upgrade EigenPodManager contract
        upgradeCall = abi.encodeCall(
            EigenPodManager.initialize,
            (
                proxyAdmin, // initialOwner
                configData.eigenPodManager.initPausedStatus // initialPausedStatus
            )
        );
        UpgradeableProxyLib.upgradeAndCall(
            result.eigenPodManager,
            eigenPodManagerImpl,
            upgradeCall
        );

        // Upgrade RewardsCoordinator contract
        upgradeCall = abi.encodeCall(
            RewardsCoordinator.initialize,
            (
                proxyAdmin, // initialOwner
                configData.rewardsCoordinator.initPausedStatus, // initialPausedStatus
                configData.rewardsCoordinator.updater,
                uint32(configData.rewardsCoordinator.activationDelay), // _activationDelay
                uint16(
                    configData.rewardsCoordinator.globalOperatorCommissionBips
                ) // _globalCommissionBips
            )
        );
        UpgradeableProxyLib.upgradeAndCall(
            result.rewardsCoordinator,
            rewardsCoordinatorImpl,
            upgradeCall
        );

        UpgradeableProxyLib.upgrade(
            result.permissionController,
            permissionControllerImpl
        );

        // Upgrade EigenPod contract
        upgradeCall = abi.encodeCall(
            EigenPod.initialize,
            // TODO: Double check this
            (address(result.eigenPodManager)) // _podOwner
        );
        UpgradeableProxyLib.upgradeAndCall(
            result.eigenPodBeacon,
            eigenPodImpl,
            upgradeCall
        );

        return result;
    }

    struct DeploymentConfigData {
        string version;
        StrategyManagerConfig strategyManager;
        DelegationManagerConfig delegationManager;
        AllocationManagerConfig allocationManager;
        EigenPodManagerConfig eigenPodManager;
        RewardsCoordinatorConfig rewardsCoordinator;
        StrategyFactoryConfig strategyFactory;
    }

    // StrategyConfig[] strategies;

    function readDeploymentConfigValues(
        string memory directoryPath,
        string memory fileName
    ) internal view returns (DeploymentConfigData memory) {
        string memory pathToFile = string.concat(directoryPath, fileName);

        require(
            vm.exists(pathToFile),
            "CoreDeployment: Deployment config file does not exist"
        );

        string memory json = vm.readFile(pathToFile);

        DeploymentConfigData memory data;

        data.version = json.readString("._version");

        // StrategyManager config start
        data.strategyManager.initPausedStatus = json.readUint(
            ".strategyManager.init_paused_status"
        );
        // StrategyManager config end

        // AllocationManager config start
        data.allocationManager.initPausedStatus = json.readUint(
            ".allocation.init_paused_status"
        );
        data.allocationManager.allocationConfigurationDelay = uint32(
            json.readUint(".allocation.allocation_config_delay")
        );
        data.allocationManager.deallocationDelay = uint32(
            json.readUint(".allocation.deallocation_delay")
        );
        // AllocationManager config end

        // DelegationManager config start
        data.delegationManager.initPausedStatus = json.readUint(
            ".delegation.init_paused_status"
        );
        data.delegationManager.withdrawalDelayBlocks = uint32(
            json.readUint(".delegation.init_withdrawal_delay_blocks")
        );
        // DelegationManager config end

        // EigenPodManager config start
        data.eigenPodManager.initPausedStatus = json.readUint(
            ".eigenPodManager.init_paused_status"
        );
        // EigenPodManager config end

        // RewardsCoordinator config start
        data.rewardsCoordinator.initPausedStatus = json.readUint(
            ".rewardsCoordinator.init_paused_status"
        );
        data.rewardsCoordinator.maxRewardsDuration = uint32(
            json.readUint(".rewardsCoordinator.MAX_REWARDS_DURATION")
        );
        data.rewardsCoordinator.maxRetroactiveLength = uint32(
            json.readUint(".rewardsCoordinator.MAX_RETROACTIVE_LENGTH")
        );
        data.rewardsCoordinator.maxFutureLength = uint32(
            json.readUint(".rewardsCoordinator.MAX_FUTURE_LENGTH")
        );
        data.rewardsCoordinator.genesisRewardsTimestamp = uint32(
            json.readUint(".rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP")
        );
        data.rewardsCoordinator.updater = json.readAddress(
            ".rewardsCoordinator.rewards_updater_address"
        );
        data.rewardsCoordinator.updaterKey = json.readUint(
            ".rewardsCoordinator.rewards_updater_key"
        );
        data.rewardsCoordinator.activationDelay = json.readUint(
            ".rewardsCoordinator.activation_delay"
        );
        data.rewardsCoordinator.calculationIntervalSeconds = uint32(
            json.readUint(".rewardsCoordinator.calculation_interval_seconds")
        );
        data.rewardsCoordinator.globalOperatorCommissionBips = json.readUint(
            ".rewardsCoordinator.global_operator_commission_bips"
        );
        // RewardsCoordinator config end

        // StrategyFactory config start
        data.strategyFactory.initPausedStatus = json.readUint(
            ".strategyFactory.init_paused_status"
        );
        // StrategyFactory config end

        return data;
    }

    function readDeploymentConfigValues(
        string memory directoryPath,
        uint256 chainId
    ) internal view returns (DeploymentConfigData memory) {
        return
            readDeploymentConfigValues(
                directoryPath,
                string.concat(vm.toString(chainId), ".json")
            );
    }

    function readDeploymentJson(
        string memory directoryPath,
        uint256 chainId
    ) internal view returns (DeploymentData memory) {
        return
            readDeploymentJson(
                directoryPath,
                string.concat(vm.toString(chainId), ".json")
            );
    }

    function readDeploymentJson(
        string memory path,
        string memory fileName
    ) internal view returns (DeploymentData memory) {
        string memory pathToFile = string.concat(path, fileName);

        require(
            vm.exists(pathToFile),
            "CoreDeployment: Deployment file does not exist"
        );

        string memory json = vm.readFile(pathToFile);

        DeploymentData memory data;
        data.strategyFactory = json.readAddress(".addresses.strategyFactory");
        data.strategyManager = json.readAddress(".addresses.strategyManager");
        data.eigenPodManager = json.readAddress(".addresses.eigenPodManager");
        data.delegationManager = json.readAddress(".addresses.delegation");
        data.allocationManager = json.readAddress(".addresses.allocation");
        data.rewardsCoordinator = json.readAddress(
            ".addresses.rewardsCoordinator"
        );
        data.initialOwner = json.readAddress(".addresses.proxyAdmin");

        return data;
    }

    /// TODO: Need to be able to read json from eigenlayer-contracts repo for holesky/mainnet and output the json here
    function writeDeploymentJson(DeploymentData memory data) internal {
        writeDeploymentJson("deployments/core/", block.chainid, data);
    }

    function writeDeploymentJson(
        string memory path,
        uint256 chainId,
        DeploymentData memory data
    ) internal {
        address proxyAdmin = address(
            UpgradeableProxyLib.getProxyAdmin(data.strategyManager)
        );

        string memory deploymentData = _generateDeploymentJson(
            data,
            proxyAdmin
        );

        string memory fileName = string.concat(
            path,
            vm.toString(chainId),
            ".json"
        );
        if (!vm.exists(path)) {
            vm.createDir(path, true);
        }

        vm.writeFile(fileName, deploymentData);
        console2.log("Deployment artifacts written to:", fileName);
    }

    function _generateDeploymentJson(
        DeploymentData memory data,
        address proxyAdmin
    ) private view returns (string memory) {
        return
            string.concat(
                '{"lastUpdate":{"timestamp":"',
                vm.toString(block.timestamp),
                '","block_number":"',
                vm.toString(block.number),
                '"},"addresses":',
                _generateContractsJson(data, proxyAdmin),
                "}"
            );
    }

    function _generateContractsJson(
        DeploymentData memory data,
        address proxyAdmin
    ) private view returns (string memory) {
        /// TODO: namespace contracts -> {avs, core}
        return
            string.concat(
                '{"proxyAdmin":"',
                proxyAdmin.toHexString(),
                '","delegation":"',
                data.delegationManager.toHexString(),
                '","delegationManagerImpl":"',
                data.delegationManager.getImplementation().toHexString(),
                '","allocation":"',
                data.allocationManager.toHexString(),
                '","allocationManagerImpl":"',
                data.allocationManager.getImplementation().toHexString(),
                '","strategyManager":"',
                data.strategyManager.toHexString(),
                '","strategyManagerImpl":"',
                data.strategyManager.getImplementation().toHexString(),
                '","eigenPodManager":"',
                data.eigenPodManager.toHexString(),
                '","eigenPodManagerImpl":"',
                data.eigenPodManager.getImplementation().toHexString(),
                '","strategyFactory":"',
                data.strategyFactory.toHexString(),
                '","strategyFactoryImpl":"',
                data.strategyFactory.getImplementation().toHexString(),
                '","strategyBeacon":"',
                data.strategyBeacon.toHexString(),
                '","rewardsCoordinator":"',
                data.rewardsCoordinator.toHexString(),
                '","permissionController":"',
                data.permissionController.toHexString(),
                '"}'
            );
    }
}
