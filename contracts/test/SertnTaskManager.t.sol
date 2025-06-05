// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

import {Test, console2 as console} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import {StrategyBase} from "@eigenlayer/contracts/strategies/StrategyBase.sol";
import {IStrategyManager} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import "@eigenlayer/contracts/libraries/OperatorSetLib.sol";
import {SertnTaskManager} from "../src/SertnTaskManager.sol";
import {ISertnTaskManager} from "../interfaces/ISertnTaskManager.sol";
import {ISertnServiceManager} from "../interfaces/ISertnServiceManager.sol";
import {ModelRegistry} from "../src/ModelRegistry.sol";
import {IModelRegistry} from "../interfaces/IModelRegistry.sol";
import {MockVerifier} from "./mockContracts/VerifierMock.sol";
import {MockAllocationManager} from "./mockContracts/AllocationManagerMock.sol";
import {MockPauserRegistry} from "./mockContracts/PauserRegistryMock.sol";
import {MockSertnServiceManager} from "./mockContracts/SertnServiceManagerMock.sol";
import {ERC20Mock} from "./mockContracts/ERC20Mock.sol";

import {Test, console2 as console} from "forge-std/Test.sol";

contract MockDelegationManager {
    // Empty implementation
}

contract MockRewardsCoordinator {
    // Empty implementation
}

contract SertnTaskManagerTest is Test {
    SertnTaskManager taskManager;
    MockAllocationManager mockAllocationManager;
    MockDelegationManager mockDelegationManager;
    MockRewardsCoordinator mockRewardsCoordinator;
    MockSertnServiceManager mockServiceManager;
    ModelRegistry modelRegistry;
    MockVerifier mockVerifier;
    ERC20Mock mockToken;
    StrategyBase mockStrategy;

    address owner = vm.addr(1);
    address aggregator = vm.addr(2);
    address operator = vm.addr(3);
    address user = vm.addr(4);
    address nonAggregator = vm.addr(5);

    uint256 modelId;

    function setUp() public {
        vm.startPrank(owner);

        // Deploy mock contracts
        mockAllocationManager = new MockAllocationManager();
        mockDelegationManager = new MockDelegationManager();
        mockRewardsCoordinator = new MockRewardsCoordinator();
        mockServiceManager = new MockSertnServiceManager();
        mockToken = new ERC20Mock();
        mockStrategy = new StrategyBase(
            IStrategyManager(address(0)),
            IPauserRegistry(address(new MockPauserRegistry())),
            "0"
        );
        // mockStrategy.initialize(ERC20(address(mockToken)));

        // Deploy and initialize ModelRegistry
        modelRegistry = new ModelRegistry();
        modelRegistry.initialize();

        mockVerifier = new MockVerifier();

        // Create a model for testing
        modelId = modelRegistry.createNewModel(
            address(mockVerifier),
            IModelRegistry.VerificationStrategy.Onchain,
            "test_model",
            100
        );

        // Deploy and initialize TaskManager
        taskManager = new SertnTaskManager();
        taskManager.initialize(
            address(mockRewardsCoordinator),
            address(mockDelegationManager),
            address(mockAllocationManager),
            address(mockServiceManager),
            address(modelRegistry)
        );

        // Setup mock data
        mockServiceManager.addAggregator(aggregator);

        // Setup allocation data for operator
        OperatorSet[] memory sets = new OperatorSet[](1);
        sets[0] = OperatorSet({id: 1, avs: address(0)});
        mockAllocationManager.setAllocatedSets(operator, sets);

        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = mockStrategy;
        mockAllocationManager.setAllocatedStrategies(
            operator,
            sets[0],
            strategies
        );

        vm.stopPrank();
    }

    function test_initialize() public view {
        assertEq(
            address(taskManager.allocationManager()),
            address(mockAllocationManager)
        );
        assertEq(
            address(taskManager.delegationManager()),
            address(mockDelegationManager)
        );
        assertEq(
            address(taskManager.rewardsCoordinator()),
            address(mockRewardsCoordinator)
        );
        assertEq(
            address(taskManager.sertnServiceManager()),
            address(mockServiceManager)
        );
        assertEq(address(taskManager.modelRegistry()), address(modelRegistry));
        assertEq(taskManager.taskNonce(), 0);
    }

    function test_sendTask_success() public {
        ISertnTaskManager.Task memory task = _createValidTask();

        vm.expectEmit(true, true, false, true);
        emit ISertnTaskManager.TaskCreated(1, user);

        vm.expectEmit(true, true, false, true);
        emit ISertnTaskManager.TaskAssigned(1, operator);

        vm.prank(aggregator);
        taskManager.sendTask(task);

        assertEq(taskManager.taskNonce(), 1);

        ISertnTaskManager.Task memory storedTask = taskManager.getTask(0);
        assertEq(
            uint8(storedTask.state),
            uint8(ISertnTaskManager.TaskState.ASSIGNED)
        );
        assertEq(storedTask.operator, operator);
        assertEq(storedTask.user, user);
        assertEq(storedTask.modelId, modelId);
    }

    function test_sendTask_revertNotAggregator() public {
        ISertnTaskManager.Task memory task = _createValidTask();

        vm.expectRevert(ISertnTaskManager.NotAggregator.selector);
        vm.prank(nonAggregator);
        taskManager.sendTask(task);
    }

    function test_sendTask_revertInvalidModelId() public {
        ISertnTaskManager.Task memory task = _createValidTask();
        task.modelId = 999; // Invalid model ID

        vm.expectRevert(ISertnTaskManager.InvalidModelId.selector);
        vm.prank(aggregator);
        taskManager.sendTask(task);
    }

    function test_getTask_revertTaskDoesNotExist() public {
        vm.expectRevert(ISertnTaskManager.TaskDoesNotExist.selector);
        taskManager.getTask(999);
    }

    function test_submitTaskOutput_success() public {
        // First send a task
        ISertnTaskManager.Task memory task = _createValidTask();
        vm.prank(aggregator);
        taskManager.sendTask(task);

        bytes memory output = "test output";

        vm.expectEmit(true, true, false, true);
        emit ISertnTaskManager.TaskCompleted(0, operator);

        vm.prank(operator);
        taskManager.submitTaskOutput(0, output);

        assertEq(
            uint8(taskManager.getTask(0).state),
            uint8(ISertnTaskManager.TaskState.COMPLETED)
        );
        assertEq(taskManager.getTask(0).output, output);
    }

    function test_submitTaskOutput_revertTaskDoesNotExist() public {
        vm.expectRevert(ISertnTaskManager.TaskDoesNotExist.selector);
        vm.prank(operator);
        taskManager.submitTaskOutput(999, "output");
    }

    function test_submitTaskOutput_revertNotAssignedToTask() public {
        // Send a task assigned to operator
        ISertnTaskManager.Task memory task = _createValidTask();
        vm.prank(aggregator);
        taskManager.sendTask(task);

        // Try to submit output from different address
        vm.expectRevert(ISertnTaskManager.NotAssignedToTask.selector);
        vm.prank(user);
        taskManager.submitTaskOutput(0, "output");
    }

    function test_submitTaskOutput_revertTaskStateIncorrect() public {
        // Send and complete a task
        ISertnTaskManager.Task memory task = _createValidTask();
        vm.prank(aggregator);
        taskManager.sendTask(task);

        vm.prank(operator);
        taskManager.submitTaskOutput(0, "output");

        // Try to submit output again
        vm.expectRevert(
            abi.encodeWithSelector(
                ISertnTaskManager.TaskStateIncorrect.selector,
                ISertnTaskManager.TaskState.COMPLETED
            )
        );
        vm.prank(operator);
        taskManager.submitTaskOutput(0, "output2");
    }

    function test_challengeTask_success() public {
        // Send and complete a task
        ISertnTaskManager.Task memory task = _createValidTask();
        vm.prank(aggregator);
        taskManager.sendTask(task);

        vm.prank(operator);
        taskManager.submitTaskOutput(0, "output");

        vm.expectEmit(true, true, false, true);
        emit ISertnTaskManager.TaskChallenged(0, user);

        vm.prank(aggregator);
        taskManager.challengeTask(0);
    }

    function test_challengeTask_revertNotAggregator() public {
        vm.expectRevert(ISertnTaskManager.NotAggregator.selector);
        vm.prank(nonAggregator);
        taskManager.challengeTask(0);
    }

    function test_challengeTask_revertTaskDoesNotExist() public {
        vm.expectRevert(ISertnTaskManager.TaskDoesNotExist.selector);
        vm.prank(aggregator);
        taskManager.challengeTask(999);
    }

    function test_submitProofForTask_success() public {
        // Send, complete, and challenge a task
        ISertnTaskManager.Task memory task = _createValidTask();
        vm.prank(aggregator);
        taskManager.sendTask(task);

        vm.prank(operator);
        taskManager.submitTaskOutput(0, "output");

        vm.prank(aggregator);
        taskManager.challengeTask(0);

        bytes memory proof = bytes("1");

        vm.expectEmit(true, false, false, true);
        emit ISertnTaskManager.ProofSubmitted(0, keccak256(proof));

        vm.prank(operator);
        taskManager.submitProofForTask(0, proof);

        assertEq(
            uint8(taskManager.getTask(0).state),
            uint8(ISertnTaskManager.TaskState.RESOLVED)
        );
    }

    function test_submitProofForTask_revertTaskDoesNotExist() public {
        vm.expectRevert(ISertnTaskManager.TaskDoesNotExist.selector);
        vm.prank(operator);
        taskManager.submitProofForTask(999, "proof");
    }

    function test_submitProofForTask_revertNotAssignedToTask() public {
        // Send, complete, and challenge a task
        ISertnTaskManager.Task memory task = _createValidTask();
        vm.prank(aggregator);
        taskManager.sendTask(task);

        vm.prank(operator);
        taskManager.submitTaskOutput(0, "output");

        vm.prank(aggregator);
        taskManager.challengeTask(0);

        vm.expectRevert(ISertnTaskManager.NotAssignedToTask.selector);
        vm.prank(user);
        taskManager.submitProofForTask(0, "proof");
    }

    function test_submitProofForTask_revertTaskStateIncorrect() public {
        // Send a task but don't challenge it
        ISertnTaskManager.Task memory task = _createValidTask();
        vm.prank(aggregator);
        taskManager.sendTask(task);

        vm.prank(operator);
        taskManager.submitTaskOutput(0, "output");

        vm.expectRevert(
            abi.encodeWithSelector(
                ISertnTaskManager.TaskStateIncorrect.selector,
                ISertnTaskManager.TaskState.CHALLENGED
            )
        );
        vm.prank(operator);
        taskManager.submitProofForTask(0, "proof");
    }

    function test_resolveTask_success() public {
        // Send, complete, and challenge a task
        ISertnTaskManager.Task memory task = _createValidTask();
        vm.prank(aggregator);
        taskManager.sendTask(task);

        vm.prank(operator);
        taskManager.submitTaskOutput(0, "output");

        vm.prank(aggregator);
        taskManager.challengeTask(0);

        vm.expectEmit(true, true, false, true);
        emit ISertnTaskManager.TaskResolved(0, operator);

        vm.prank(aggregator);
        taskManager.resolveTask(0, true);

        assertEq(
            uint8(taskManager.getTask(0).state),
            uint8(ISertnTaskManager.TaskState.RESOLVED)
        );
    }

    function test_resolveTask_reject() public {
        // Send, complete, and challenge a task
        ISertnTaskManager.Task memory task = _createValidTask();
        vm.prank(aggregator);
        taskManager.sendTask(task);

        vm.prank(operator);
        taskManager.submitTaskOutput(0, "output");

        vm.prank(aggregator);
        taskManager.challengeTask(0);

        vm.expectEmit(true, true, false, true);
        emit ISertnTaskManager.TaskRejected(0, operator);

        vm.prank(aggregator);
        taskManager.resolveTask(0, false);

        assertEq(
            uint8(taskManager.getTask(0).state),
            uint8(ISertnTaskManager.TaskState.REJECTED)
        );
    }

    function test_resolveTask_revertNotAggregator() public {
        vm.expectRevert(ISertnTaskManager.NotAggregator.selector);
        vm.prank(nonAggregator);
        taskManager.resolveTask(0, true);
    }

    function test_resolveTask_revertTaskDoesNotExist() public {
        vm.expectRevert(ISertnTaskManager.TaskDoesNotExist.selector);
        vm.prank(aggregator);
        taskManager.resolveTask(999, true);
    }

    function test_resolveTask_revertTaskStateIncorrect() public {
        // Send a task but don't challenge it
        ISertnTaskManager.Task memory task = _createValidTask();
        vm.prank(aggregator);
        taskManager.sendTask(task);

        vm.expectRevert(
            abi.encodeWithSelector(
                ISertnTaskManager.TaskStateIncorrect.selector,
                ISertnTaskManager.TaskState.CHALLENGED
            )
        );
        vm.prank(aggregator);
        taskManager.resolveTask(0, true);
    }

    function test_taskNonceIncrementsCorrectly() public {
        ISertnTaskManager.Task memory task = _createValidTask();

        assertEq(taskManager.taskNonce(), 0);

        vm.prank(aggregator);
        taskManager.sendTask(task);
        assertEq(taskManager.taskNonce(), 1);

        vm.prank(aggregator);
        taskManager.sendTask(task);
        assertEq(taskManager.taskNonce(), 2);
    }

    // Helper function to create a valid task
    function _createValidTask()
        internal
        view
        returns (ISertnTaskManager.Task memory)
    {
        return
            ISertnTaskManager.Task({
                startBlock: block.number,
                modelId: modelId,
                inputs: "test inputs",
                proof: 100,
                user: user,
                nonce: 0,
                operator: operator,
                state: ISertnTaskManager.TaskState.CREATED,
                output: "",
                fee: 1000
            });
    }
}
