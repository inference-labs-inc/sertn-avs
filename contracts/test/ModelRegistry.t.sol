// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Test, console2 as console} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
import {CoreDeploymentLib} from "../script/utils/CoreDeploymentLib.sol";
import {UpgradeableProxyLib} from "../script/utils/UpgradeableProxyLib.sol";
import {ModelRegistry} from "../src/ModelRegistry.sol";
import "../interfaces/IModelRegistry.sol";
import {MockVerifier} from "./mockContracts/VerifierMock.sol";
import {MockVerifier2} from "./mockContracts/VerifierMock2.sol";

contract ModelRegistryTest is Test {
    struct AVSOwner {
        Vm.Wallet key;
    }

    AVSOwner internal owner;

    CoreDeploymentLib.DeploymentData internal coreDeployment;
    CoreDeploymentLib.DeploymentConfigData coreConfigData;

    ModelRegistry modelRegistry;
    MockVerifier mockVerifier;
    MockVerifier2 mockVerifier2;
    MockVerifier2 mockVerifier3;

    uint256 modelId;

    function setUp() public {
        console.log("ModelRegistryTest setUp");

        // Create owner wallet
        owner = AVSOwner({key: vm.createWallet("owner_wallet")});

        // Deploy proxy admin
        address proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();
        console.log("ProxyAdmin deployed");

        // Read config and deploy core contracts
        coreConfigData = CoreDeploymentLib.readDeploymentConfigValues(
            "test/mockData/config/core/",
            1337
        );
        console.log("Config read");

        coreDeployment = CoreDeploymentLib.deployContracts(proxyAdmin, coreConfigData);
        console.log("CoreDeploymentLib deployed");

        // Deploy contracts as owner
        vm.startPrank(owner.key.addr);

        modelRegistry = new ModelRegistry();
        modelRegistry.initialize();
        mockVerifier = new MockVerifier();
        mockVerifier2 = new MockVerifier2();
        mockVerifier3 = new MockVerifier2();

        modelId = modelRegistry.createNewModel(
            address(mockVerifier),
            IModelRegistry.VerificationStrategy.Onchain,
            "model1",
            100,
            10
        );

        vm.stopPrank();

        console.log("ModelRegistry and verifiers deployed");
    }

    function test_createSecondModel() public {
        vm.startPrank(owner.key.addr);

        // Fail to create second model with same verifier
        vm.expectRevert(
            abi.encodeWithSelector(IModelRegistry.ModelAlreadyExists.selector, modelId)
        );
        modelRegistry.createNewModel(
            address(mockVerifier),
            IModelRegistry.VerificationStrategy.Offchain,
            "model2",
            200,
            15
        );

        // Create second model
        vm.expectEmit(true, true, true, true, address(modelRegistry)); // Expect ModelCreated event
        emit IModelRegistry.ModelCreated(
            2, // modelIndex
            address(mockVerifier2),
            IModelRegistry.VerificationStrategy.Offchain,
            "model2",
            200,
            15
        );
        uint256 modelId2 = modelRegistry.createNewModel(
            address(mockVerifier2),
            IModelRegistry.VerificationStrategy.Offchain,
            "model2",
            200,
            15
        );

        require(modelId == 1, "First model ID should be 1");
        require(modelId2 == 2, "Second model ID should be 2");

        // Verify model names
        require(
            keccak256(abi.encodePacked(modelRegistry.modelName(modelId))) ==
                keccak256(abi.encodePacked("model1")),
            "First model name should be model1"
        );

        require(
            keccak256(abi.encodePacked(modelRegistry.modelName(modelId2))) ==
                keccak256(abi.encodePacked("model2")),
            "Second model name should be model2"
        );

        vm.stopPrank();
    }

    function test_updateModelName() public {
        vm.startPrank(owner.key.addr);

        // Update the model name
        vm.expectEmit(true, true, false, false, address(modelRegistry));
        emit IModelRegistry.ModelNameUpdated(modelId, "updatedModel1");
        modelRegistry.updateModelName(modelId, "updatedModel1");

        // Verify the update
        require(
            keccak256(abi.encodePacked(modelRegistry.modelName(modelId))) ==
                keccak256(abi.encodePacked("updatedModel1")),
            "Model name should be updated to updatedModel1"
        );

        vm.stopPrank();
    }

    function test_updateComputeCost() public {
        vm.startPrank(owner.key.addr);

        // Update the compute cost
        vm.expectEmit(true, true, false, false, address(modelRegistry));
        emit IModelRegistry.ComputeCostUpdated(modelId, 200);
        modelRegistry.updateComputeCost(modelId, 200);

        // Verify the update
        require(modelRegistry.computeCost(modelId) == 200, "Compute cost should be updated to 200");

        vm.stopPrank();
    }

    function test_updateModelVerifier() public {
        vm.startPrank(owner.key.addr);

        // Fail to update model verifier to the same address
        vm.expectRevert(
            abi.encodeWithSelector(IModelRegistry.ModelAlreadyExists.selector, modelId)
        );
        modelRegistry.updateModelVerifier(modelId, address(mockVerifier));
        // Fail to update model verifier to an invalid address
        vm.expectRevert(abi.encodeWithSelector(IModelRegistry.InvalidModelVerifier.selector));
        modelRegistry.updateModelVerifier(modelId, address(0));

        // Update the model verifier
        vm.expectEmit(true, true, false, false, address(modelRegistry));
        emit IModelRegistry.ModelVerifierUpdated(modelId, address(mockVerifier2));
        modelRegistry.updateModelVerifier(modelId, address(mockVerifier2));

        // Verify the update
        require(
            modelRegistry.modelVerifier(modelId) == address(mockVerifier2),
            "Model verifier should be updated to mockVerifier2"
        );

        vm.stopPrank();
    }

    function test_updateVerificationStrategy() public {
        vm.startPrank(owner.key.addr);

        // Update the verification strategy
        vm.expectEmit(true, true, false, false, address(modelRegistry));
        emit IModelRegistry.VerificationStrategyUpdated(
            modelId,
            IModelRegistry.VerificationStrategy.Offchain
        );
        modelRegistry.updateVerificationStrategy(
            modelId,
            IModelRegistry.VerificationStrategy.Offchain
        );

        // Verify the update
        require(
            modelRegistry.verificationStrategy(modelId) ==
                IModelRegistry.VerificationStrategy.Offchain,
            "Verification strategy should be updated to Offchain"
        );

        vm.stopPrank();
    }

    function test_Permissions() public {
    // Check that only the owner can update model name
        vm.startPrank(address(0x123));
        vm.expectRevert("Ownable: caller is not the owner");
    modelRegistry.updateModelName(modelId, "unauthorizedUpdate");
        vm.stopPrank();

        // Check that only the owner can update compute cost
        vm.startPrank(address(0x123));
        vm.expectRevert("Ownable: caller is not the owner");
        modelRegistry.updateComputeCost(modelId, 300);
        vm.stopPrank();

        // Check that only the owner can update model verifier
        vm.startPrank(address(0x123));
        vm.expectRevert("Ownable: caller is not the owner");
        modelRegistry.updateModelVerifier(modelId, address(mockVerifier3));
        vm.stopPrank();

        // Check that only the owner can update verification strategy
        vm.startPrank(address(0x123));
        vm.expectRevert("Ownable: caller is not the owner");
        modelRegistry.updateVerificationStrategy(
            modelId,
            IModelRegistry.VerificationStrategy.Onchain
        );
        vm.stopPrank();
    }
}
