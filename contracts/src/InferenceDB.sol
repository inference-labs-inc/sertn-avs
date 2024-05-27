// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {IZKVerifier} from "./IZKVerifier.sol";
import {IInferenceDB} from "./IInferenceDB.sol";

contract InferenceDB is OwnableUpgradeable, IInferenceDB {
    uint32 TASK_RESPONSE_WINDOW_BLOCK = 0;
    IZKVerifier zkVerifier;

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
    mapping(uint32 => TaskChallengeMetadata) challengeData;
    mapping(uint32 => mapping(uint256 => uint256)) challengeInstances;

    constructor(uint32 _taskResponseWindowBlock, address _taskManagerAddress) {
        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
        taskManagerAddress = _taskManagerAddress;
    }

    function initialize(address _zkVerifier) public override initializer {
        zkVerifier = IZKVerifier(_zkVerifier);
    }

    function challenge(
        uint32 referenceTaskIndex,
        uint256[5] calldata inputs,
        uint256 output
    ) public override onlyTaskManager {
        require(
            challengeData[referenceTaskIndex].taskProven ==
                ChallengeStatus.NotChallenged
        );

        for (uint i = 0; i < inputs.length; i++) {
            challengeInstances[referenceTaskIndex][i] = inputs[i];
        }

        challengeInstances[referenceTaskIndex][5] = output;

        challengeData[referenceTaskIndex].challenger = msg.sender;
        challengeData[referenceTaskIndex].timeChallenged = block.timestamp;
        challengeData[referenceTaskIndex].taskProven = ChallengeStatus
            .ChallengedAndPendingConfirmation;
    }

    function proveResultAccurate(
        uint32 taskId,
        uint256[] calldata instances,
        bytes calldata proof
    ) public override onlyTaskManager {
        TaskChallengeMetadata storage challengeMetadata = challengeData[taskId];

        for (uint256 i = 0; i < 6; i++) {
            require(challengeInstances[taskId][i] == instances[i]);
        }

        require(zkVerifier.verifyProof(proof, instances));

        challengeMetadata.taskProven = ChallengeStatus.ProofConfirmed;
    }

    function confirmChallenge(
        uint32 referenceTaskIndex
    ) public override onlyTaskManager {
        require(
            challengeData[referenceTaskIndex].timeChallenged < block.timestamp
        );

        require(
            challengeData[referenceTaskIndex].taskProven ==
                ChallengeStatus.ChallengedAndPendingConfirmation
        );

        challengeData[referenceTaskIndex].taskProven = ChallengeStatus
            .ProofRejected;
    }

    function updateTaskManager(
        address newTaskManagerAddress
    ) external override onlyOwner {
        taskManagerAddress = newTaskManagerAddress;
    }
}
