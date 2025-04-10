// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {ISertnTaskManager} from "./ISertnTaskManager.sol";

/**
 * @title ISertnAggregator
 * @author Inference Labs, Inc.
 * @notice Interface for the Aggregators within Sertn
 */
interface ISertnAggregator {
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
     * @param taskId The hash of the task
     * @param submitter The address that submitted the task
     */
    event TaskSubmitted(bytes32 indexed taskId, address indexed submitter);

    /**
     * @notice Submits a task to the Aggregator
     * @param _task The task to submit
     * @param _proof The proof of the task signed by aggregatorEOA
     */
    function submitTask(
        ISertnTaskManager.Task memory _task,
        bytes memory _proof
    ) external;

    /**
     * @notice Updates the Aggregator EOA
     * @param _aggregatorEOA The new Aggregator EOA
     */
    function updateAggregatorEOA(address _aggregatorEOA) external;

    /**
     * @notice Updates the SertnTaskManager contract address
     * @param _sertnTaskManager The new SertnTaskManager contract address
     */
    function updateSertnTaskManager(address _sertnTaskManager) external;
}
