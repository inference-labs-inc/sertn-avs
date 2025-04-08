// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ISertnTaskManager {
    error InvalidModelId();

    struct Task {
        uint256 modelId;
        bytes inputs;
        uint256 poc;
        uint256 startTime;
        uint32 startingBlock;
        bool proveOnResponse;
        address user;
        uint256 nonce;
    }

    function sendTask(Task memory _task) external;

    function submitTask(bytes32 _taskId, bytes memory _proof) external;

    function submitTaskResponse(bytes32 _taskId, bytes memory _proof) external;
}
