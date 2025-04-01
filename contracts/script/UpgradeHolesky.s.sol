// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/Test.sol";
import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {SertnServiceManager} from "../src/SertnServiceManager.sol";
import {SertnTaskManager} from "../src/SertnTaskManager.sol";
import {ModelStorage} from "../src/ModelStorage.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";

contract UpgradeHolesky is Script {
    using CoreDeploymentLib for *;

    address private deployer;
    address proxyAdmin;
    CoreDeploymentLib.DeploymentData coreDeployment;
    IStrategy[] strategies;

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");

        coreDeployment = CoreDeploymentLib.readDeploymentJson("deployments/core/", block.chainid);
    }

    function run() external {
        vm.startBroadcast(deployer);

        SertnServiceManager newServiceManager = new SertnServiceManager(
            coreDeployment.rewardsCoordinator,
            coreDeployment.delegationManager,
            coreDeployment.allocationManager,
            strategies,
            ""
        );

        ModelStorage newModelStorage = new ModelStorage(address(newServiceManager));

        SertnTaskManager newTaskManager = new SertnTaskManager(
            coreDeployment.rewardsCoordinator,
            coreDeployment.delegationManager,
            coreDeployment.allocationManager,
            address(newServiceManager),
            address(newModelStorage),
            address(0) // serToken address will be set after deployment
        );

        ProxyAdmin proxyAdminContract = ProxyAdmin(proxyAdmin);

        proxyAdminContract.upgrade(
            TransparentUpgradeableProxy(payable(address(newServiceManager))),
            address(newServiceManager)
        );

        proxyAdminContract.upgrade(
            TransparentUpgradeableProxy(payable(address(newTaskManager))),
            address(newTaskManager)
        );

        proxyAdminContract.upgrade(
            TransparentUpgradeableProxy(payable(address(newModelStorage))),
            address(newModelStorage)
        );

        vm.stopBroadcast();
    }
}
