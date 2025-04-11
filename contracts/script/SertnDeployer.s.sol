// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/Test.sol";
// import {SertnDeploymentLib} from "./utils/SertnDeploymentLib.sol";
import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";
import {UpgradeableProxyLib} from "./utils/UpgradeableProxyLib.sol";
import {StrategyBase} from "@eigenlayer/contracts/strategies/StrategyBase.sol";
import {ERC20Mock} from "../test/ERC20Mock.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";
import {StrategyManager} from "@eigenlayer/contracts/core/StrategyManager.sol";
import {IRewardsCoordinator} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";

import {SertnServiceManager} from "../src/SertnServiceManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";

import {IERC20, StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";
import {SertnTaskManager} from "../src/SertnTaskManager.sol";
import {ModelRegistry} from "../src/ModelRegistry.sol";
import "forge-std/Test.sol";

contract SertnDeployer is Script, Test {
    using CoreDeploymentLib for *;
    using UpgradeableProxyLib for address;

    address private deployer;
    address proxyAdmin;
    address rewardsOwner;
    address rewardsInitiator;
    IStrategy sertnStrategy;
    CoreDeploymentLib.DeploymentData coreDeployment;
    // SertnDeploymentLib.DeploymentData sertnDeployment;
    // SertnDeploymentLib.DeploymentConfigData sertnConfig;
    // IECDSAStakeRegistryTypes.Quorum internal quorum;
    ERC20Mock token;

    IStrategy[] _tokenToStrategy;
    IStrategy[] _ethStrategies;
    IStrategy _serStrategy;
    IStrategy[] strategies;

    uint32[] opSetIds;

    ERC20Mock public ethToken1;
    ERC20Mock public ethToken2;
    ERC20Mock public serToken;

    SertnServiceManager sertnServiceManager;
    SertnTaskManager sertnTaskManager;
    ModelRegistry modelRegistry;

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");

        coreDeployment = CoreDeploymentLib.readDeploymentJson(
            "deployments/core/",
            block.chainid
        );
    }

    function run() external {
        vm.startBroadcast(deployer);
        ethToken1 = new ERC20Mock();
        ethToken2 = new ERC20Mock();
        serToken = new ERC20Mock();

        IStrategy strategy1 = addStrategy(address(ethToken1));
        IStrategy strategy2 = addStrategy(address(ethToken2));
        _serStrategy = addStrategy(address(serToken));
        strategies.push(_serStrategy);
        strategies.push(strategy1);
        strategies.push(strategy2);

        _ethStrategies.push(strategy1);
        _ethStrategies.push(strategy2);

        // quorum.strategies.push(IECDSAStakeRegistryTypes.StrategyParams({strategy: strategy, multiplier: 10_000}));

        // sertnDeployment = SertnDeploymentLib.deployContracts(
        //     proxyAdmin, coreDeployment, quorum, owner.key.addr, owner.key.addr
        // );
        sertnServiceManager = new SertnServiceManager();
        modelRegistry = new ModelRegistry();
        sertnTaskManager = new SertnTaskManager();

        // Initialize contracts
        sertnServiceManager.initialize(
            coreDeployment.rewardsCoordinator,
            coreDeployment.delegationManager,
            coreDeployment.allocationManager,
            address(0),
            strategies,
            ""
        );

        modelRegistry.initialize();

        sertnTaskManager.initialize(
            coreDeployment.rewardsCoordinator,
            coreDeployment.delegationManager,
            coreDeployment.allocationManager,
            address(sertnServiceManager),
            address(modelRegistry)
        );

        sertnServiceManager.updateTaskManager(address(sertnTaskManager));
        sertnServiceManager.updateModelRegistry(address(modelRegistry));

        // sertnServiceManager.setTaskManager(address(sertnTaskManager));

        // Save addresses of the contracts to a JSON file:
        string memory json = vm.serializeAddress(
            "SertnDeployment",
            "sertnServiceManager",
            address(sertnServiceManager)
        );
        json = vm.serializeAddress(
            "SertnDeployment",
            "sertnTaskManager",
            address(sertnTaskManager)
        );
        for (uint256 i = 0; i < strategies.length; i++) {
            json = vm.serializeAddress(
                "SertnDeployment",
                string.concat("strategy_", Strings.toString(i)),
                address(strategies[i])
            );
        }
        for (uint256 i = 0; i < _ethStrategies.length; i++) {
            json = vm.serializeAddress(
                "SertnDeployment",
                string.concat("eth_strategy_", Strings.toString(i)),
                address(_ethStrategies[i])
            );
        }
        vm.writeFile("deployments/sertnDeployment.json", json);

        vm.stopBroadcast();
        // verifyDeployment();
        // SertnDeploymentLib.writeDeploymentJson(sertnDeployment);
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

    //     function verifyDeployment() internal view {y

    //         require(
    //             sertnDeployment.stakeRegistry != address(0), "StakeRegistry address cannot be zero"
    //         );
    //         require(
    //             sertnDeployment.sertnServiceManager != address(0),
    //             "SertnServiceManager address cannot be zero"
    //         );
    //         require(sertnDeployment.strategy != address(0), "Strategy address cannot be zero");
    //         require(proxyAdmin != address(0), "ProxyAdmin address cannot be zero");
    //         require(
    //             coreDeployment.delegationManager != address(0),
    //             "DelegationManager address cannot be zero"
    //         );
    //     }
}
