pragma solidity ^0.8.0;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/Test.sol";

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
import {SertnRegistrar} from "../src/SertnRegistrar.sol";
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
    SertnRegistrar sertnRegistrar;

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

        sertnServiceManager = new SertnServiceManager();
        modelRegistry = new ModelRegistry();
        sertnRegistrar = new SertnRegistrar();
        sertnTaskManager = new SertnTaskManager();

        sertnRegistrar.initialize(address(sertnServiceManager));

        sertnServiceManager.initialize(
            coreDeployment.rewardsCoordinator,
            coreDeployment.delegationManager,
            coreDeployment.allocationManager,
            address(sertnRegistrar),
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
        json = vm.serializeAddress(
            "SertnDeployment",
            "sertnRegistrar",
            address(sertnRegistrar)
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
