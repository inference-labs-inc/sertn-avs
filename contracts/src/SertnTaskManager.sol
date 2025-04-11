// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";
import {IAllocationManager} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IRewardsCoordinator} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {ISertnServiceManager} from "../interfaces/ISertnServiceManager.sol";
import {ISertnTaskManager} from "../interfaces/ISertnTaskManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {IVerifier} from "../interfaces/IVerifier.sol";
import {ModelRegistry} from "./ModelRegistry.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract SertnTaskManager is OwnableUpgradeable, ISertnTaskManager {
    // queue of tasks that are waiting to be assigned to an operator
    address[] public operators;
    // queue of tasks that are waiting to be challenged
    bytes[] public slashingQueue;
    // nonce for tasks to ensure uniqueness
    uint256 taskNonce;
    // Mapping from taskId to Task struct
    mapping(uint256 => Task) public tasks;

    IERC20 public ser;

    IAllocationManager public allocationManager;
    IDelegationManager public delegationManager;
    IRewardsCoordinator public rewardsCoordinator;

    ISertnServiceManager public sertnServiceManager;
    ModelRegistry public modelRegistry;

    modifier onlyAggregators() {
        if (!sertnServiceManager.isAggregator(msg.sender)) {
            revert NotAggregator();
        }
        _;
    }

    modifier onlyOperators() {
        // if (!sertnServiceManager.isOperator(msg.sender)) {
        //     revert NotOperator();
        // }
        _;
    }

    function initialize(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        address _sertnServiceManager,
        address _modelRegistry
    ) public initializer {
        __Ownable_init();
        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);
        sertnServiceManager = ISertnServiceManager(_sertnServiceManager);
        modelRegistry = ModelRegistry(_modelRegistry);
    }

    function sendTask(Task memory task) external onlyAggregators {
        tasks[taskNonce] = task;
        taskNonce++;
        emit TaskCreated(taskNonce, task.user);
        emit TaskAssigned(taskNonce, task.operator);
    }

    function submitTaskOutput(uint256 taskId, bytes calldata output) external {
        Task memory task = tasks[taskId];
        if (taskId > taskNonce) {
            revert TaskDoesNotExist();
        }
        if (task.operator != msg.sender) {
            revert NotAssignedToTask();
        }
        if (task.state != TaskState.ASSIGNED) {
            revert TaskStateIncorrect(TaskState.ASSIGNED);
        }

        task.output = output;
        task.state = TaskState.COMPLETED;
        emit TaskCompleted(taskId, task.operator);

        OperatorSet memory operatorSet = allocationManager.getAllocatedSets(
            task.operator
        )[0];
        IStrategy memory strategy = allocationManager.getAllocatedStrategies(
            task.operator,
            operatorSet
        )[0];
        IERC20 token = allocationManager.getAllocatedTokens(
            task.operator,
            operatorSet
        )[0];

        sertnServiceManager.taskCompleted(
            task.operator,
            task.fee,
            strategy,
            token
        );
    }

    function challengeTask(uint256 taskId) external onlyAggregators {
        Task memory task = tasks[taskId];
        if (taskId > taskNonce) {
            revert TaskDoesNotExist();
        }
        if (
            task.state != TaskState.COMPLETED &&
            task.blockNumber + TIMEOUT > block.number
        ) {
            sertnServiceManager.slashOperator(task.operator, task.fee);
        }
        if (task.state != TaskState.COMPLETED) {
            revert TaskStateIncorrect(TaskState.COMPLETED);
        }
        task.state = TaskState.CHALLENGED;
        emit TaskChallenged(taskId, task.user);
    }

    function submitProofForTask(
        uint256 taskId,
        bytes calldata proof
    ) external onlyOperators {
        if (taskId > taskNonce) {
            revert TaskDoesNotExist();
        }
        if (tasks[taskId].operator != msg.sender) {
            revert NotAssignedToTask();
        }
        if (tasks[taskId].state != TaskState.CHALLENGED) {
            revert TaskStateIncorrect(TaskState.CHALLENGED);
        }

        tasks[taskId].state = TaskState.RESOLVED;
        emit TaskResolved(taskId, tasks[taskId].operator);
    }
}
