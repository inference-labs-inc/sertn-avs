// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract InferenceDB is OwnableUpgradeable {
    uint32 TASK_RESPONSE_WINDOW_BLOCK = 0;
    address taskManagerAddress;

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

    modifier onlyTaskManager() {
        require(
            msg.sender == taskManagerAddress,
            "Function restricted to task manager"
        );
        _;
    }

    function initialize() public initializer {}

    function challenge(
        uint32 referenceTaskIndex,
        uint256[5] calldata inputs,
        uint256 output
    ) public onlyTaskManager {
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
    ) public onlyTaskManager {
        TaskChallengeMetadata storage challengeMetadata = challengeData[taskId];

        for (uint256 i = 0; i < 6; i++) {
            require(challengeInstances[taskId][i] == instances[i]);
        }

        require(zkVerifier.verifyProof(proof, instances));

        challengeMetadata.taskProven = ChallengeStatus.ProofConfirmed;
    }

    function proveChallenge() public onlyTaskManager {}

    function updateTaskManager(
        address newTaskManagerAddress
    ) external onlyOwner {
        taskManagerAddress = newTaskManagerAddress;
    }
}
