// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IZKVerifier} from "./IZKVerifier.sol";
import {IInferenceDB} from "./IInferenceDB.sol";

contract InferenceDB is IInferenceDB {
    /* STORAGE */
    // The latest task index
    uint32 public latestTaskNum;

    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK = 30;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 30;

    // mapping of task indices to all tasks hashes
    // when a task is created, task hash is stored here,
    // and responses need to pass the actual task,
    // which is hashed onchain and checked against this mapping
    mapping(uint32 => bytes32) public allTaskHashes;

    // mapping of task indices to hash of abi.encode(taskResponse, taskResponseMetadata)
    mapping(uint32 => bytes32) public allTaskResponses;

    mapping(uint32 => address) public modelVerifiersForTasks;

    uint32 TASK_CHALLENGE_RESPONSE_WINDOW_BLOCK = 0;
    IZKVerifier zkVerifier;

    constructor(uint32 _taskChallengeResponseWindowBlock, address _zkVerifier) {
        TASK_CHALLENGE_RESPONSE_WINDOW_BLOCK = _taskChallengeResponseWindowBlock;
        zkVerifier = IZKVerifier(_zkVerifier);
    }

    function challenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata
    ) public override onlyTaskManager {
        uint32 referenceTaskIndex = taskResponse.referenceTaskIndex;
        // some logical checks
        require(allTaskResponses[referenceTaskIndex] != bytes32(0));
        require(
            allTaskResponses[referenceTaskIndex] ==
                keccak256(abi.encode(taskResponse, taskResponseMetadata))
        );

        require(
            uint32(block.number) <=
                taskResponseMetadata.taskResponsedBlock +
                    TASK_CHALLENGE_WINDOW_BLOCK
        );

        require(
            challengeData[referenceTaskIndex].taskProven ==
                ChallengeStatus.NotChallenged
        );

        for (uint i = 0; i < task.inputs.length; i++) {
            challengeInstances[referenceTaskIndex][i] = task.inputs[i];
        }

        challengeInstances[referenceTaskIndex][5] = taskResponse.output;
        if (task.modelVerifier == address(0)) {
            modelVerifiersForTasks[referenceTaskIndex] = address(zkVerifier);
        } else {
            modelVerifiersForTasks[referenceTaskIndex] = task.modelVerifier;
        }

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

        require(
            IZKVerifier(modelVerifiersForTasks[taskId]).verifyProof(
                proof,
                instances
            )
        );

        challengeMetadata.taskProven = ChallengeStatus.ProofConfirmed;
    }

    function confirmChallenge(
        uint32 referenceTaskIndex
    ) public override onlyTaskManager {
        require(
            challengeData[referenceTaskIndex].timeChallenged +
                TASK_CHALLENGE_RESPONSE_WINDOW_BLOCK *
                12 seconds <
                block.timestamp
        );

        require(
            challengeData[referenceTaskIndex].taskProven ==
                ChallengeStatus.ChallengedAndPendingConfirmation
        );

        challengeData[referenceTaskIndex].taskProven = ChallengeStatus
            .ProofRejected;
    }

    function createNewTask(
        bytes32 taskHash
    ) public override onlyTaskManager returns (uint32) {
        // store hash of task onchain, emit event, and increase taskNum
        allTaskHashes[latestTaskNum] = taskHash;
        latestTaskNum = latestTaskNum + 1;
        return (latestTaskNum - 1);
    }

    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata
    ) public override onlyTaskManager {
        uint32 taskCreatedBlock = task.taskCreatedBlock;
        require(task.provenOnResponse == false);
        // check that the task is valid, hasn't been responsed yet, and is being responsed in time
        require(
            keccak256(abi.encode(task)) ==
                allTaskHashes[taskResponse.referenceTaskIndex]
        );
        // some logical checks
        require(
            allTaskResponses[taskResponse.referenceTaskIndex] == bytes32(0)
        );

        require(
            uint32(block.number) <=
                taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK
        );

        allTaskResponses[taskResponse.referenceTaskIndex] = keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        );
    }

    function updateTaskManager(
        address newTaskManagerAddress
    ) public override onlyOwner {
        taskManagerAddress = newTaskManagerAddress;
    }

    function respondToTaskWithProof(
        Task calldata task,
        uint32 taskIndex,
        uint256[] calldata instances,
        bytes calldata proof
    ) public override {
        uint32 taskCreatedBlock = task.taskCreatedBlock;

        require(task.provenOnResponse == true);
        // check that the task is valid, hasn't been responsed yet, and is being responsed in time
        require(keccak256(abi.encode(task)) == allTaskHashes[taskIndex]);
        // some logical checks
        require(allTaskResponses[taskIndex] == bytes32(0));
        require(
            uint32(block.number) <=
                taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK
        );
        for (uint256 i = 0; i < task.inputs.length; i++) {
            require(instances[i] == task.inputs[i]);
        }
        address modelVerifier = task.modelVerifier;
        if (modelVerifier == address(0)) modelVerifier = address(zkVerifier);
        require(IZKVerifier(modelVerifier).verifyProof(proof, instances));

        challengeData[taskIndex].taskProven = ChallengeStatus.ProofRejected;
    }

    modifier onlyTaskManager() {
        require(
            msg.sender == taskManagerAddress,
            "Function restricted to task manager"
        );
        _;
    }
}
