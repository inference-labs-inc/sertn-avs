// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";
import {IAllocationManager, IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
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
import {OperatorSet} from "@eigenlayer/contracts/libraries/OperatorSetLib.sol";

contract SertnTaskManager is OwnableUpgradeable, ISertnTaskManager {
    // queue of tasks that are waiting to be assigned to an operator
    address[] public operators;
    // queue of tasks that are waiting to be challenged
    bytes[] public slashingQueue;
    // nonce for tasks to ensure uniqueness
    uint256 public taskNonce;
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
        IStrategy strategy = allocationManager.getAllocatedStrategies(
            task.operator,
            allocationManager.getAllocatedSets(task.operator)[0]
        )[0];
        IERC20 token = strategy.underlyingToken();
        sertnServiceManager.pullFeeFromUser(task.user, token, task.fee);
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
        IStrategy strategy = allocationManager.getAllocatedStrategies(
            task.operator,
            operatorSet
        )[0];
        IERC20 token = strategy.underlyingToken();

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
        // TODO: configurable timeout
        if (
            task.state != TaskState.COMPLETED &&
            task.startBlock + 300 > block.number
        ) {
            OperatorSet memory operatorSet = allocationManager.getAllocatedSets(
                task.operator
            )[0];
            IStrategy strategy = allocationManager.getAllocatedStrategies(
                task.operator,
                operatorSet
            )[0];

            sertnServiceManager.slashOperator(
                task.operator,
                task.fee,
                operatorSet.id,
                strategy
            );
        }
        if (task.state != TaskState.COMPLETED) {
            revert TaskStateIncorrect(TaskState.COMPLETED);
        }
        task.state = TaskState.CHALLENGED;
        emit TaskChallenged(taskId, task.user);
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
        // TODO: Solidity hashes bytes fields longer than 32 bytes in events
        //       so if the proof is longer than 32 bytes, nobody will be able to
        //       verify it.
        emit ProofSubmitted(taskId, proof);
    }

    function resolveTask(
        uint256 taskId,
        bool success
    ) external onlyAggregators {
        if (taskId > taskNonce) {
            revert TaskDoesNotExist();
        }
        Task memory task = tasks[taskId];
        if (task.state != TaskState.CHALLENGED) {
            revert TaskStateIncorrect(TaskState.CHALLENGED);
        }
        OperatorSet memory operatorSet = allocationManager.getAllocatedSets(
            task.operator
        )[0];
        IStrategy strategy = allocationManager.getAllocatedStrategies(
            task.operator,
            operatorSet
        )[0];
        IERC20 token = strategy.underlyingToken();

        if (success) {
            sertnServiceManager.taskCompleted(
                task.operator,
                task.fee,
                strategy,
                token
            );
            task.state = TaskState.RESOLVED;
            emit TaskResolved(taskId, task.operator);
        } else {
            sertnServiceManager.slashOperator(
                task.operator,
                task.fee,
                operatorSet.id,
                strategy
            );
            task.state = TaskState.REJECTED;
            emit TaskRejected(taskId, task.operator);
        }
    }
}
