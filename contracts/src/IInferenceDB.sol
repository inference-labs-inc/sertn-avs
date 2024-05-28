// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;
import "@openzeppelin/contracts/access/Ownable.sol";

import {IZKVerifier} from "./IZKVerifier.sol";

abstract contract IInferenceDB is Ownable {
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
        uint32 referenceTaskIndex,
        uint256[5] calldata inputs,
        uint256 output
    ) public virtual;

    function proveResultAccurate(
        uint32 taskId,
        uint256[] calldata instances,
        bytes calldata proof
    ) public virtual;

    function confirmChallenge(uint32 referenceTaskIndex) public virtual;

    function updateTaskManager(address newTaskManagerAddress) external virtual;

    modifier onlyTaskManager() {
        require(
            msg.sender == taskManagerAddress,
            "Function restricted to task manager"
        );
        _;
    }
}
