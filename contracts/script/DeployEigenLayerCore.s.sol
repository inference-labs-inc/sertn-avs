// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Script} from "forge-std/Script.sol";

import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";
import {UpgradeableProxyLib} from "./utils/UpgradeableProxyLib.sol";

import {IRewardsCoordinator} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";

import "forge-std/Test.sol";

contract DeployEigenLayerCore is Script, Test {
    using CoreDeploymentLib for *;
    using UpgradeableProxyLib for address;

    address internal deployer;
    address internal proxyAdmin;
    CoreDeploymentLib.DeploymentData internal deploymentData;
    CoreDeploymentLib.DeploymentConfigData internal configData;

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");
    }

    function run() external {
        vm.startBroadcast(deployer);
        //set the rewards updater to the deployer address for payment flow
        configData = CoreDeploymentLib.readDeploymentConfigValues("config/core/", block.chainid);
        configData.rewardsCoordinator.updater = deployer;
        proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();
        deploymentData = CoreDeploymentLib.deployContracts(proxyAdmin, configData);
        vm.stopBroadcast();
        string memory deploymentPath = "deployments/core/";
        CoreDeploymentLib.writeDeploymentJson(deploymentPath, block.chainid, deploymentData);
    }
}
