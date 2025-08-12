// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";
import {IAllocationManager, IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IRewardsCoordinator} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {OperatorSet} from "@eigenlayer/contracts/libraries/OperatorSetLib.sol";
import {ISertnServiceManager} from "../interfaces/ISertnServiceManager.sol";
import {ISertnTaskManager} from "../interfaces/ISertnTaskManager.sol";
import {ISertnNodesManager} from "../interfaces/ISertnNodesManager.sol";
import {IVerifier} from "../interfaces/IVerifier.sol";
import {IModelRegistry} from "../interfaces/IModelRegistry.sol";
import {ModelRegistry} from "./ModelRegistry.sol";
import {SertnNodesManager} from "./SertnNodesManager.sol";


contract SertnTaskManager is OwnableUpgradeable, ISertnTaskManager {
    using EnumerableSet for EnumerableSet.UintSet;
    // queue of tasks that are waiting to be assigned to an operator
    address[] public operators;
    // queue of tasks that are waiting to be challenged
    bytes[] public slashingQueue;
    // nonce for tasks to ensure uniqueness
    uint256 public taskNonce = 1;
    // Mapping from taskId to Task struct
    mapping(uint256 => Task) public tasks;
    // all assigned tasks IDs, which are not resolved and not rejected
    EnumerableSet.UintSet private pendingTasks;

    IERC20 public ser;

    IAllocationManager public allocationManager;
    IDelegationManager public delegationManager;
    IRewardsCoordinator public rewardsCoordinator;
    ISertnNodesManager public sertnNodesManager;
    ISertnServiceManager public sertnServiceManager;
    ModelRegistry public modelRegistry;

    modifier onlyAggregators() {
        if (!sertnServiceManager.isAggregator(msg.sender)) {
            revert NotAggregator();
        }
        _;
    }

    function initialize(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        address _sertnServiceManager,
        address _modelRegistry,
        address _sertnNodesManager
    ) public initializer {
        __Ownable_init();
        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);
        sertnServiceManager = ISertnServiceManager(_sertnServiceManager);
        modelRegistry = ModelRegistry(_modelRegistry);
        sertnNodesManager = ISertnNodesManager(_sertnNodesManager);
        taskNonce = 1; // Start task nonce at 1 to avoid zero-indexing issues
    }

    function sendTask(Task memory task) external onlyAggregators {
        // check if the task is valid
        if (modelRegistry.modelVerifier(task.modelId) == address(0)) revert InvalidModelId();
        task.startTimestamp = uint32(block.timestamp);
        task.nonce = taskNonce;
        tasks[taskNonce] = task;
        taskNonce++;
        IStrategy strategy = allocationManager.getAllocatedStrategies(
            task.operator,
            allocationManager.getAllocatedSets(task.operator)[0]
        )[0];
        IERC20 token = strategy.underlyingToken();
        sertnServiceManager.pullFeeFromUser(task.user, token, task.fee);

        emit TaskCreated(task.nonce, task.user);
        tasks[task.nonce].state = TaskState.ASSIGNED;
        pendingTasks.add(task.nonce);

        // Allocate FUCUs for this task
        _allocateFucusForTask(task.nonce);

        emit TaskAssigned(task.nonce, task.operator);
    }

    function getTask(uint256 taskId) external view returns (Task memory) {
        if (taskId > taskNonce) {
            revert TaskDoesNotExist();
        }
        return tasks[taskId];
    }

    function submitTaskOutput(uint256 taskId, bytes calldata output) external {
        if (taskId > taskNonce) {
            revert TaskDoesNotExist();
        }
        Task memory task = tasks[taskId];
        if (task.operator != msg.sender) {
            revert NotAssignedToTask();
        }
        if (task.state != TaskState.ASSIGNED) {
            revert TaskStateIncorrect(task.state);
        }

        tasks[taskId].output = output;
        tasks[taskId].state = TaskState.COMPLETED;
        emit TaskCompleted(taskId, task.operator);

        OperatorSet memory operatorSet = allocationManager.getAllocatedSets(task.operator)[0];
        IStrategy strategy = allocationManager.getAllocatedStrategies(task.operator, operatorSet)[
            0
        ];
    }

    function challengeTask(uint256 taskId) external onlyAggregators {
        Task memory task = tasks[taskId];
        if (taskId > taskNonce) {
            revert TaskDoesNotExist();
        }
        // TODO: configurable timeout
        if (
            (task.state == TaskState.ASSIGNED && task.startBlock + 300 > block.number) ||
            (task.state == TaskState.CHALLENGED && task.startBlock + 600 > block.number)
        ) {
            OperatorSet memory operatorSet = allocationManager.getAllocatedSets(task.operator)[0];
            IStrategy strategy = allocationManager.getAllocatedStrategies(
                task.operator,
                operatorSet
            )[0];

            sertnServiceManager.slashOperator(task.operator, task.fee, operatorSet.id, strategy);
        }
        if (task.state != TaskState.COMPLETED) {
            revert TaskStateIncorrect(TaskState.COMPLETED);
        }
        tasks[taskId].state = TaskState.CHALLENGED;
        // TODO: maybe operator address instead of msg.sender?
        emit TaskChallenged(taskId, msg.sender);
    }

    function _resolveTask(uint256 taskId, bool success) internal {
        if (taskId > taskNonce) {
            revert TaskDoesNotExist();
        }
        Task memory task = tasks[taskId];
        if (task.state != TaskState.CHALLENGED && task.state != TaskState.COMPLETED) {
            revert TaskStateIncorrect(TaskState.CHALLENGED);
        }
        OperatorSet memory operatorSet = allocationManager.getAllocatedSets(task.operator)[0];
        IStrategy strategy = allocationManager.getAllocatedStrategies(task.operator, operatorSet)[
            0
        ];

        if (success) {
            sertnServiceManager.taskCompleted(
                task.operator,
                task.fee,
                strategy,
                task.startTimestamp
            );
            tasks[taskId].state = TaskState.RESOLVED;
            emit TaskResolved(taskId, task.operator);
        } else {
            sertnServiceManager.slashOperator(task.operator, task.fee, operatorSet.id, strategy);
            tasks[taskId].state = TaskState.REJECTED;
            emit TaskRejected(taskId, task.operator);
        }

        // Release FUCUs regardless of success or failure
        _releaseFucusForTask(taskId);

        pendingTasks.remove(taskId);
    }

    function resolveTask(uint256 taskId, bool success) public onlyAggregators {
        if (taskId > taskNonce) {
            revert TaskDoesNotExist();
        }
        Task memory task = tasks[taskId];
        if (task.state != TaskState.CHALLENGED && task.state != TaskState.COMPLETED) {
            revert TaskStateIncorrect(task.state);
        }
        _resolveTask(taskId, success);
    }

    function submitProofForTask(uint256 taskId, bytes calldata proof) external {
        if (taskId > taskNonce) {
            revert TaskDoesNotExist();
        }
        Task memory task = tasks[taskId];
        if (task.operator != msg.sender) {
            revert NotAssignedToTask();
        }
        if (task.state != TaskState.CHALLENGED) {
            revert TaskStateIncorrect(TaskState.CHALLENGED);
        }

        // Verify the proof using the model verifier
        IVerifier verifier = IVerifier(modelRegistry.modelVerifier(task.modelId));

        bytes32 proofHash = keccak256(proof);
        tasks[taskId].proofHash = proofHash;
        if (
            modelRegistry.verificationStrategy(task.modelId) ==
            IModelRegistry.VerificationStrategy.Onchain
        ) {
            // On-chain verification
            if (verifier.verifyProof(proof)) {} else {
                revert InvalidProof(taskId, proofHash);
            }
            emit ProofSubmitted(taskId, proofHash);
            _resolveTask(taskId, true);
        } else if (
            modelRegistry.verificationStrategy(task.modelId) ==
            IModelRegistry.VerificationStrategy.Offchain
        ) {
            // Off-chain verification
            // The proof is expected to be verified off-chain by the aggregator
            // Operator must submit the proof to the aggregator
            // Here we just emit an event with the proof hash
            // TODO: check isn't it a second time submitted proof?
            emit ProofSubmitted(taskId, proofHash);
        } else {
            revert InvalidVerificationStrategy(taskId);
        }
    }

    function getPendingTasksIds() public view returns (uint256[] memory) {
        uint256[] memory result = new uint256[](pendingTasks.length());
        for (uint i = 0; i < pendingTasks.length(); i++) {
            result[i] = pendingTasks.at(i);
        }
        return result;
    }

    /**
     * @notice Select an operator for a model based on available FUCUs
     * @param modelId The ID of the model to run
     * @param requiredFucus The required compute units (if 0, uses model's compute cost)
     * @return selectedOperator The address of the selected operator (address(0) if none available)
     */
    function selectOperatorForModel(
        uint256 modelId,
        uint256 requiredFucus
    ) external view returns (address selectedOperator) {
        // If no specific FUCU requirement, use model's compute cost
        if (requiredFucus == 0) {
            requiredFucus = modelRegistry.computeCost(modelId);
        }

        // Get available operators for this model
        (address[] memory availableOperators, ) = sertnNodesManager.getAvailableOperatorsForModel(
            modelId,
            requiredFucus
        );

        if (availableOperators.length == 0) {
            return address(0); // No operators available
        }

        // Simple selection: return the first operator with enough capacity
        // TODO: Implement more sophisticated selection logic (round-robin, load balancing, etc.)
        return availableOperators[0];
    }

    /**
     * @notice Get available operators for a model with their capacity
     * @param modelId The ID of the model
     * @param requiredFucus The required compute units (if 0, uses model's compute cost)
     * @return availableOperators Array of available operators
     * @return availableFucus Array of available FUCUs for each operator
     */
    function getAvailableOperators(
        uint256 modelId,
        uint256 requiredFucus
    ) external view returns (address[] memory availableOperators, uint256[] memory availableFucus) {
        // If no specific FUCU requirement, use model's compute cost
        if (requiredFucus == 0) {
            requiredFucus = modelRegistry.computeCost(modelId);
        }

        return sertnNodesManager.getAvailableOperatorsForModel(modelId, requiredFucus);
    }

    /**
     * @notice Allocate FUCUs for a task when it's assigned to an operator
     * @param taskId The ID of the task
     */
    function _allocateFucusForTask(uint256 taskId) internal {
        Task memory task = tasks[taskId];
        uint256 requiredFucus = modelRegistry.requiredFUCUs(task.modelId);

        // Try to allocate FUCUs for this task
        bool success = sertnNodesManager.allocateFucusForTask(
            task.operator,
            task.modelId,
            requiredFucus
        );

        if (!success) {
            revert ISertnTaskManager.InsufficientFucusCapacity(
                task.operator,
                task.modelId,
                requiredFucus
            );
        }
    }

    /**
     * @notice Release FUCUs when a task is completed or rejected
     * @param taskId The ID of the task
     */
    function _releaseFucusForTask(uint256 taskId) internal {
        Task memory task = tasks[taskId];
        uint256 requiredFucus = modelRegistry.requiredFUCUs(task.modelId);

        // Release the allocated FUCUs
        sertnNodesManager.releaseFucusForTask(task.operator, task.modelId, requiredFucus);
    }
}
