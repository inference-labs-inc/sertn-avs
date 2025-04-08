// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

/**
 * @title ISertnAggregator
 * @author Inference Labs, Inc.
 * @notice Interface for the Aggregators within Sertn
 */
interface ISertnAggregator {
    /**
     * @notice Thrown when the taskId is invalid
     */
    error InvalidTaskId();

    /**
     * @notice Thrown when the EOA signature is invalid
     */
    error InvalidEOASignature();

    /**
     * @notice Thrown when the address is zero
     */
    error ZeroAddress();

    /**
     * @notice Emitted when a task is submitted
     */
    event TaskSubmitted(bytes32 indexed taskId, address indexed submitter);

    /**
     * @notice Submits a task to the Aggregator
     * @param _taskId The taskId of the task to submit
     * @param _proof The proof of the task
     */
    function submitTask(bytes32 _taskId, bytes memory _proof) external;
}
