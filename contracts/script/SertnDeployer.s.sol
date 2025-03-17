// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/Test.sol";
import {SertnDeploymentLib} from "./utils/SertnDeploymentLib.sol";
import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";
import {UpgradeableProxyLib} from "./utils/UpgradeableProxyLib.sol";
import {StrategyBase} from "@eigenlayer/contracts/strategies/StrategyBase.sol";
import {ERC20Mock} from "../test/ERC20Mock.sol";
import {TransparentUpgradeableProxy} from
    "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";
import {StrategyManager} from "@eigenlayer/contracts/core/StrategyManager.sol";
import {IRewardsCoordinator} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";

import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {IECDSAStakeRegistryTypes} from "@eigenlayer-middleware/src/interfaces/IECDSAStakeRegistry.sol";

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
    SertnDeploymentLib.DeploymentData sertnDeployment;
    SertnDeploymentLib.DeploymentConfigData sertnConfig;
    IECDSAStakeRegistryTypes.Quorum internal quorum;
    ERC20Mock token;

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");

        sertnConfig =
            SertnDeploymentLib.readDeploymentConfigValues("config/sertn/", block.chainid);

        coreDeployment = CoreDeploymentLib.readDeploymentJson("deployments/core/", block.chainid);
    }

    function run() external {
        vm.startBroadcast(deployer);
        rewardsOwner = sertnConfig.rewardsOwner;
        rewardsInitiator = sertnConfig.rewardsInitiator;

        token = new ERC20Mock();
        sertnStrategy =
            IStrategy(StrategyFactory(coreDeployment.strategyFactory).deployNewStrategy(token));

        quorum.strategies.push(IECDSAStakeRegistryTypes.StrategyParams({strategy: sertnStrategy, multiplier: 10_000}));

        proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();

        sertnDeployment = SertnDeploymentLib.deployContracts(
            proxyAdmin, coreDeployment, quorum, rewardsInitiator, rewardsOwner
        );

        sertnDeployment.strategy = address(sertnStrategy);
        sertnDeployment.token = address(token);

        vm.stopBroadcast();
        verifyDeployment();
        SertnDeploymentLib.writeDeploymentJson(sertnDeployment);
    }

    function verifyDeployment() internal view {
        require(
            sertnDeployment.stakeRegistry != address(0), "StakeRegistry address cannot be zero"
        );
        require(
            sertnDeployment.sertnServiceManager != address(0),
            "SertnServiceManager address cannot be zero"
        );
        require(sertnDeployment.strategy != address(0), "Strategy address cannot be zero");
        require(proxyAdmin != address(0), "ProxyAdmin address cannot be zero");
        require(
            coreDeployment.delegationManager != address(0),
            "DelegationManager address cannot be zero"
        );
    }
}
