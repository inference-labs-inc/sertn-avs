// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;
import "@openzeppelin/contracts/access/Ownable.sol";
import "./ITaskStruct.sol";
import {IZKVerifier} from "./IZKVerifier.sol";

abstract contract IInferenceDB is Ownable, ITaskStruct {
    address taskManagerAddress;
    mapping(uint32 => TaskChallengeMetadata) public challengeData;
    mapping(uint32 => mapping(uint256 => uint256)) public challengeInstances;

    enum ChallengeStatus {
        NotChallenged,
        ChallengedAndPendingConfirmation,
        ProofConfirmed,
        ProofRejected
    }

    struct TaskChallengeMetadata {
        address challenger;
        ChallengeStatus taskProven;
        uint256 timeChallenged;
    }

    function challenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata
    ) public virtual;

    function proveResultAccurate(
        uint32 taskId,
        uint256[] calldata instances,
        bytes calldata proof
    ) public virtual;

    function confirmChallenge(uint32 referenceTaskIndex) public virtual;

    function createNewTask(bytes32 taskHash) public virtual returns (uint32);

    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata
    ) public virtual;

    function updateTaskManager(address newTaskManagerAddress) public virtual;

    function respondToTaskWithProof(
        Task calldata task,
        uint32 taskIndex,
        uint256[] calldata instances,
        bytes calldata proof
    ) public virtual;
}
