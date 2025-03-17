// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Test} from "forge-std/Test.sol";
import {console2} from "forge-std/console2.sol";
import {CoreDeploymentLib} from "../script/utils/CoreDeploymentLib.sol";
import {UpgradeableProxyLib} from "../script/utils/UpgradeableProxyLib.sol";

contract CoreDeploymentLibTest is Test {
    using CoreDeploymentLib for *;
    using UpgradeableProxyLib for address;

    address proxyAdmin;
    CoreDeploymentLib.DeploymentData deploymentData;
    CoreDeploymentLib.DeploymentConfigData configData;

    function setUp() public {
        proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();
    }

    /// won't test specific functionality/values. Testing behavior of the library
    function test_ReadConfig() public view {
        CoreDeploymentLib.readDeploymentConfigValues("test/mockData/config/core/", 1337);
    }

    /// forge-config: default.allow_internal_expect_revert = true
    function test_ReadConfig_Reverts() public {
        vm.expectRevert();
        /// Incorrect path
        CoreDeploymentLib.readDeploymentConfigValues("test/mockData/deployments/core/", 1337);
    }

    function test_ReadDeployment() public view {
        CoreDeploymentLib.readDeploymentJson("test/mockData/deployments/core/", 1337);
    }

    /// forge-config: default.allow_internal_expect_revert = true
    function test_ReadDeployment_Reverts() public {
        vm.expectRevert();
        /// Incorrect path
        CoreDeploymentLib.readDeploymentJson("test/mockData/config/core/", 1337);
    }

    function test_DeployContracts() public {
        configData =
            CoreDeploymentLib.readDeploymentConfigValues("test/mockData/config/core/", 1337);
        CoreDeploymentLib.DeploymentData memory data =
            CoreDeploymentLib.deployContracts(proxyAdmin, configData);

        assertTrue(data.delegationManager != address(0), "DelegationManager not deployed");
        assertTrue(data.strategyManager != address(0), "StrategyManager not deployed");
        assertTrue(data.allocationManager != address(0), "AllocationyManager not deployed");
    }

    function test_WriteDeploymentJson() public {
        configData =
            CoreDeploymentLib.readDeploymentConfigValues("test/mockData/config/core/", 1337);
        CoreDeploymentLib.DeploymentData memory data =
            CoreDeploymentLib.deployContracts(proxyAdmin, configData);

        string memory scratchPath = "test/mockData/scratch/test_WriteDeploymentJson/";
        CoreDeploymentLib.writeDeploymentJson(scratchPath, block.chainid, data);

        string memory fileName = string.concat(scratchPath, vm.toString(block.chainid), ".json");
        assertTrue(vm.exists(fileName), "Deployment file not created");

        vm.removeFile(fileName);
    }

    function test_WriteAndReadDeploymentJson() public {
        configData =
            CoreDeploymentLib.readDeploymentConfigValues("test/mockData/config/core/", 1337);
        CoreDeploymentLib.DeploymentData memory initialData =
            CoreDeploymentLib.deployContracts(proxyAdmin, configData);

        string memory scratchPath = "test/mockData/scratch/test_WriteAndReadDeploymentJson/";

        CoreDeploymentLib.writeDeploymentJson(scratchPath, block.chainid, initialData);

        string memory fileName = string.concat(vm.toString(block.chainid), ".json");

        CoreDeploymentLib.readDeploymentJson(scratchPath, fileName);

        vm.removeFile(string.concat(scratchPath, fileName));
    }

    function test_ReadConfigFromM2DeploymentData() public {
        /// TODO: Deployment json is missing the strategy factory
        vm.skip(true);
        // Path to the M2 deployment data JSON file
        string memory m2DeploymentDataPath =
            "lib/eigenlayer-middleware/lib/eigenlayer-contracts/script/output/devnet/";
        string memory m2DeploymentFilename = "M2_from_scratch_deployment_data.json";

        CoreDeploymentLib.readDeploymentJson(m2DeploymentDataPath, m2DeploymentFilename);
    }
}
