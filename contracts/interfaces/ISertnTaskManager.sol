// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

interface ISertnTaskManager {
    /// @notice Thrown when the model id is invalid
    error InvalidModelId();

    /// @notice Emitted when a task is created
    event TaskCreated(bytes32 indexed taskId, address indexed user);

    /// @notice Emitted when a task is assigned to an operator
    event TaskAssigned(bytes32 indexed taskId, address indexed operator);

    /// @notice Emitted when a task is completed by an operator
    event TaskCompleted(bytes32 indexed taskId, address indexed operator);

    /// @notice Emitted when a task is challenged by a user
    event TaskChallenged(bytes32 indexed taskId, address indexed user);

    /// @notice Emitted when a task is rejected and the operator is slashed
    event TaskRejected(bytes32 indexed taskId, address indexed operator);

    /// @notice Emitted when a task is resolved and the operator is rewarded
    event TaskResolved(bytes32 indexed taskId, address indexed operator);

    /**
     * @notice The task struct
     * @param startBlock The block number when the task was created
     * @param modelId The model id
     * @param inputs The inputs to the model
     * @param proof The proof of completion
     * @param proveOnResponse Whether to prove on response
     * @param user The user who created the task
     * @param nonce The nonce for the task
     * @param operator The operator who is assigned to the task
     */
    struct Task {
        uint32 startBlock;
        uint256 modelId;
        bytes inputs;
        uint256 proof;
        address user;
        uint256 nonce;
        address operator;
    }

    /**
     * @notice The task state enum
     * @param CREATED Task has been created and is waiting to be assigned to an operator
     * @param ASSIGNED Task has been assigned to an operator
     * @param COMPLETED Task has been completed by an operator
     * @param CHALLENGED Task has been challenged by a user
     * @param REJECTED Task has been rejected, and the operator has been slashed
     * @param RESOLVED Task has been resolved, and the operator has been rewarded
     */
    enum TaskState {
        CREATED,
        ASSIGNED,
        COMPLETED,
        CHALLENGED,
        REJECTED,
        RESOLVED
    }

    /**
     * @notice Send a task to the task manager
     * @param _task The task to send
     */
    function sendTask(Task memory _task) external;

    /**
     * @notice Submit a task to the task manager
     * @param _taskId The task id
     * @param _proof The proof of completion
     */
    function submitTask(bytes32 _taskId, bytes memory _proof) external;

    /**
     * @notice Challenge a task
     * @param _taskId The task id
     */
    function challengeTask(bytes32 _taskId) external;

    /**
     * @notice Submit a task response to the task manager
     * @param _taskId The task id
     * @param _proof The proof of completion
     */
    function submitTaskResponse(bytes32 _taskId, bytes memory _proof) external;
}
