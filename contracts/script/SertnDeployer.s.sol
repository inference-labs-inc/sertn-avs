pragma solidity ^0.8.0;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/Test.sol";

import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";
import {UpgradeableProxyLib} from "./utils/UpgradeableProxyLib.sol";
import {StrategyBase} from "@eigenlayer/contracts/strategies/StrategyBase.sol";
import {ERC20Mock} from "../test/ERC20Mock.sol";
import {TransparentUpgradeableProxy} from
    "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";
import {StrategyManager} from "@eigenlayer/contracts/core/StrategyManager.sol";
import {IRewardsCoordinator} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";

import {SertnServiceManager} from "../src/SertnServiceManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {IECDSAStakeRegistryTypes} from "@eigenlayer-middleware/src/interfaces/IECDSAStakeRegistry.sol";

import {IERC20} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";
import {SertnTaskManager} from "../src/SertnTaskManager.sol";

import "forge-std/Test.sol";

contract SertnDeployer is Script, Test {
    using CoreDeploymentLib for *;
    using UpgradeableProxyLib for address;

    address private deployer;
    address proxyAdmin;
    address rewardsOwner;
    address rewardsInitiator;
    CoreDeploymentLib.DeploymentData coreDeployment;

    IStrategy[] strategies;
    IStrategy[] _ethStrategies;

    address constant STETH_STRATEGY = 0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3;
    address constant RETH_STRATEGY = 0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0;

    SertnServiceManager sertnServiceManager;
    SertnTaskManager sertnTaskManager;
    address public serviceManagerProxyAddr;

    function setUp() public virtual {
        uint256 deployerKey = vm.envUint("HOLESKY_DEPLOYER_KEY");
        console.log("Got deployer key:", deployerKey != 0);
        deployer = vm.rememberKey(deployerKey);
        console.log("Actual deployer address:", deployer);

        proxyAdmin = 0x33C1aE65ee7801270a823DF4F7FeE017cDE5A5A0;
        vm.label(proxyAdmin, "ProxyAdmin");

        if (block.chainid == 17000) {
            console.log("Setting up for Holesky deployment");
            coreDeployment.delegationManager = 0xA44151489861Fe9e3055d95adC98FbD462B948e7;
            coreDeployment.strategyManager = 0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6;
            coreDeployment.eigenPodManager = 0x30770d7E3e71112d7A6b7259542D1f680a70e315;
            coreDeployment.rewardsCoordinator = 0xAcc1fb458a1317E886dB376Fc8141540537E68fE;
            coreDeployment.allocationManager = 0x78469728304326CBc65f8f95FA756B0B73164462;
            coreDeployment.permissionController = 0x598cb226B591155F767dA17AfE7A2241a68C5C10;

            delete strategies;
            strategies.push(IStrategy(STETH_STRATEGY));
            strategies.push(IStrategy(RETH_STRATEGY));

            delete _ethStrategies;
            _ethStrategies.push(IStrategy(STETH_STRATEGY));
            _ethStrategies.push(IStrategy(RETH_STRATEGY));
        } else {
            revert("Only Holesky deployment supported");
        }

        require(coreDeployment.delegationManager != address(0), "DelegationManager not set");
        require(coreDeployment.rewardsCoordinator != address(0), "RewardsCoordinator not set");
        require(coreDeployment.allocationManager != address(0), "AllocationManager not set");
    }

    function run() external {
        if (block.chainid == 17000) {
            deployServiceManager();
            deployTaskManager();
        } else {
            revert("Only Holesky deployment supported");
        }
    }

    function deployServiceManager() public {
        vm.startBroadcast(deployer);
        console.log("Starting ServiceManager deployment...");
        console.log("Deployer address:", deployer);

        console.log("Creating implementation...");
        SertnServiceManager serviceManagerImpl = new SertnServiceManager();
        console.log("Implementation deployed at:", address(serviceManagerImpl));

        console.log("Preparing init data...");
        console.log("Strategies length:", strategies.length);
        for(uint i = 0; i < strategies.length; i++) {
            console.log("Strategy", i, ":", address(strategies[i]));
            try IStrategy(strategies[i]).underlyingToken() returns (IERC20 token) {
                console.log("Underlying token:", address(token));
            } catch {
                console.log("Failed to get underlying token for strategy", i);
            }
        }

        bytes memory serviceManagerInitData = abi.encodeWithSelector(
            SertnServiceManager.initialize.selector,
            coreDeployment.rewardsCoordinator,
            coreDeployment.delegationManager,
            coreDeployment.allocationManager,
            strategies,
            "https://raw.githubusercontent.com/inference-labs-inc/sertn-avs/v1/el-info.json"
        );
        console.log("Init data prepared");

        console.log("Creating proxy...");
        console.log("Proxy admin:", proxyAdmin);
        TransparentUpgradeableProxy serviceManagerProxy = new TransparentUpgradeableProxy(
            address(serviceManagerImpl),
            proxyAdmin,
            serviceManagerInitData
        );
        console.log("Proxy deployed at:", address(serviceManagerProxy));

        sertnServiceManager = SertnServiceManager(address(serviceManagerProxy));
        serviceManagerProxyAddr = address(serviceManagerProxy);

        console.log("SertnServiceManager implementation deployed at:", address(serviceManagerImpl));
        console.log("SertnServiceManager proxy deployed at:", address(serviceManagerProxy));

        vm.stopBroadcast();
    }

    function deployTaskManager() public {
        require(serviceManagerProxyAddr != address(0), "ServiceManager must be deployed first");

        vm.startBroadcast(deployer);

        SertnTaskManager taskManagerImpl = new SertnTaskManager();

        bytes memory taskManagerInitData = abi.encodeWithSelector(
            SertnTaskManager.initialize.selector,
            coreDeployment.rewardsCoordinator,
            coreDeployment.delegationManager,
            coreDeployment.allocationManager,
            serviceManagerProxyAddr,
            address(0)
        );

        TransparentUpgradeableProxy taskManagerProxy = new TransparentUpgradeableProxy(
            address(taskManagerImpl),
            proxyAdmin,
            taskManagerInitData
        );

        sertnTaskManager = SertnTaskManager(address(taskManagerProxy));

        sertnServiceManager = SertnServiceManager(serviceManagerProxyAddr);
        sertnServiceManager.setTaskManager(address(sertnTaskManager));

        console.log("SertnTaskManager implementation deployed at:", address(taskManagerImpl));
        console.log("SertnTaskManager proxy deployed at:", address(taskManagerProxy));
        console.log("Proxy admin address:", proxyAdmin);

        vm.stopBroadcast();
    }

    function addStrategy(address token) public returns (IStrategy) {
        StrategyFactory strategyFactory = StrategyFactory(
            coreDeployment.strategyFactory
        );
        IStrategy newStrategy = strategyFactory.deployNewStrategy(
            IERC20(token)
        );
        return newStrategy;
    }
}
