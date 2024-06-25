// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@eigenlayer-middleware/src/libraries/BN254.sol";
import "./ISertnTaskManager.sol";
import {IInferenceDB} from "./IInferenceDB.sol";

contract SertnTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    ITaskStruct
{
    using BN254 for BN254.G1Point;

    /* CONSTANT */
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    /* STORAGE */

    address public aggregator;
    address public generator;
    IInferenceDB inferenceDB;
    /* MODIFIERS */
    modifier onlyAggregator() {
        require(msg.sender == aggregator);
        _;
    }

    // onlyTaskGenerator is used to restrict createNewTask from only being called by a permissioned entity
    // in a real world scenario, this would be removed by instead making createNewTask a payable function
    modifier onlyTaskGenerator() {
        require(msg.sender == generator);
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator
    ) BLSSignatureChecker(_registryCoordinator) {}

    function challengeInstances(
        uint32 id,
        uint256 index
    ) public view returns (uint256) {
        return inferenceDB.challengeInstances(id, index);
    }

    function setNewInferenceDB(address _inferenceDB) public onlyTaskGenerator {
        inferenceDB = IInferenceDB(_inferenceDB);
    }

    function setNewAggregator(address _aggregator) public onlyAggregator {
        aggregator = _aggregator;
        generator = _aggregator;
    }

    /* FUNCTIONS */
    // NOTE: this function creates new task, assigns it a taskId
    function createNewTask(
        uint256[5] calldata inputs,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers,
        bool provenOnResponce,
        address modelVerifier
    ) external onlyTaskGenerator {
        // create a new task struct
        Task memory newTask;
        newTask.inputs = inputs;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumThresholdPercentage = quorumThresholdPercentage;
        newTask.quorumNumbers = quorumNumbers;
        newTask.provenOnResponse = provenOnResponce;
        newTask.modelVerifier = modelVerifier;

        uint32 taskNum = inferenceDB.createNewTask(
            keccak256(abi.encode(newTask))
        );

        emit NewTaskCreated(taskNum, newTask);
    }

    // NOTE: this function responds to existing tasks.
    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;
        bytes calldata quorumNumbers = task.quorumNumbers;
        uint32 quorumThresholdPercentage = task.quorumThresholdPercentage;

        /* CHECKING SIGNATURES & WHETHER THRESHOLD IS MET OR NOT */
        // calculate message which operators signed
        bytes32 message = keccak256(abi.encode(taskResponse));

        // check the BLS signature
        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
                message,
                quorumNumbers,
                taskCreatedBlock,
                nonSignerStakesAndSignature
            );

        // check that signatories own at least a threshold percentage of each quourm
        for (uint i = 0; i < quorumNumbers.length; i++) {
            // we don't check that the quorumThresholdPercentages are not >100 because a greater value would trivially fail the check, implying
            // signed stake > total stake
            require(
                quorumStakeTotals.signedStakeForQuorum[i] *
                    _THRESHOLD_DENOMINATOR >=
                    quorumStakeTotals.totalStakeForQuorum[i] *
                        uint8(quorumThresholdPercentage)
            );
        }

        TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(
            uint32(block.number),
            hashOfNonSigners
        );
        inferenceDB.respondToTask(task, taskResponse, taskResponseMetadata);
        // updating the storage with task responsea
        // emitting event
        emit TaskResponded(taskResponse, taskResponseMetadata);
    }

    function respondToTaskWithProof(
        Task calldata task,
        uint32 taskIndex,
        uint256[] calldata instances,
        bytes calldata proof
    ) public {
        inferenceDB.respondToTaskWithProof(task, taskIndex, instances, proof);
        emit TaskRespondedWithProof(
            taskIndex,
            instances[instances.length - 1],
            msg.sender
        );
    }

    function raiseChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata
    ) external {
        inferenceDB.challenge(task, taskResponse, taskResponseMetadata);
        emit TaskChallenged(taskResponse.referenceTaskIndex, task);
    }

    function proveResultAccurate(
        uint32 taskId,
        uint256[] calldata instances,
        bytes calldata proof
    ) external {
        inferenceDB.proveResultAccurate(taskId, instances, proof);
        emit TaskChallengedUnsuccessfully(taskId, msg.sender);
    }

    function confirmChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external {
        uint32 referenceTaskIndex = taskResponse.referenceTaskIndex;
        // get the list of hash of pubkeys of operators who weren't part of the task response submitted by the aggregator
        bytes32[] memory hashesOfPubkeysOfNonSigningOperators = new bytes32[](
            pubkeysOfNonSigningOperators.length
        );

        for (uint i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            hashesOfPubkeysOfNonSigningOperators[
                i
            ] = pubkeysOfNonSigningOperators[i].hashG1Point();
        }

        // verify whether the pubkeys of "claimed" non-signers supplied by challenger are actually non-signers as recorded before
        // when the aggregator responded to the task
        // currently inlined, as the MiddlewareUtils.computeSignatoryRecordHash function was removed from BLSSignatureChecker
        // in this PR: https://github.com/Layr-Labs/eigenlayer-contracts/commit/c836178bf57adaedff37262dff1def18310f3dce#diff-8ab29af002b60fc80e3d6564e37419017c804ae4e788f4c5ff468ce2249b4386L155-L158
        // TODO(samlaf): contracts team will add this function back in the BLSSignatureChecker, which we should use to prevent potential bugs from code duplication
        bytes32 signatoryRecordHash = keccak256(
            abi.encodePacked(
                task.taskCreatedBlock,
                hashesOfPubkeysOfNonSigningOperators
            )
        );
        require(signatoryRecordHash == taskResponseMetadata.hashOfNonSigners);

        // get the address of operators who didn't sign
        address[] memory addresssOfNonSigningOperators = new address[](
            pubkeysOfNonSigningOperators.length
        );
        for (uint i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            addresssOfNonSigningOperators[i] = BLSApkRegistry(
                address(blsApkRegistry)
            ).pubkeyHashToOperator(hashesOfPubkeysOfNonSigningOperators[i]);
        }

        inferenceDB.confirmChallenge(referenceTaskIndex);

        // @dev the below code is commented out for the upcoming M2 release
        //      in which there will be no slashing. The slasher is also being redesigned
        //      so its interface may very well change.
        // ==========================================
        // // get the list of all operators who were active when the task was initialized
        // Operator[][] memory allOperatorInfo = getOperatorState(
        //     IRegistryCoordinator(address(registryCoordinator)),
        //     task.quorumNumbers,
        //     task.taskCreatedBlock
        // );
        // // freeze the operators who signed adversarially
        // for (uint i = 0; i < allOperatorInfo.length; i++) {
        //     // first for loop iterate over quorums

        //     for (uint j = 0; j < allOperatorInfo[i].length; j++) {
        //         // second for loop iterate over operators active in the quorum when the task was initialized

        //         // get the operator address
        //         bytes32 operatorID = allOperatorInfo[i][j].operatorId;
        //         address operatorAddress = BLSPubkeyRegistry(
        //             address(blsPubkeyRegistry)
        //         ).pubkeyCompendium().pubkeyHashToOperator(operatorID);

        //         // check if the operator has already NOT been frozen
        //         if (
        //             IServiceManager(
        //                 address(
        //                     BLSRegistryCoordinatorWithIndices(
        //                         address(registryCoordinator)
        //                     ).serviceManager()
        //                 )
        //             ).slasher().isFrozen(operatorAddress) == false
        //         ) {
        //             // check whether the operator was a signer for the task
        //             bool wasSigningOperator = true;
        //             for (
        //                 uint k = 0;
        //                 k < addresssOfNonSigningOperators.length;
        //                 k++
        //             ) {
        //                 if (
        //                     operatorAddress == addresssOfNonSigningOperators[k]
        //                 ) {
        //                     // if the operator was a non-signer, then we set the flag to false
        //                     wasSigningOperator == false;
        //                     break;
        //                 }
        //             }

        //             if (wasSigningOperator == true) {
        //                 BLSRegistryCoordinatorWithIndices(
        //                     address(registryCoordinator)
        //                 ).serviceManager().freezeOperator(operatorAddress);
        //             }
        //         }
        //     }
        // }

        emit TaskChallengedSuccessfully(referenceTaskIndex, msg.sender);
    }

    event NewTaskCreated(uint32 indexed taskIndex, Task task);

    event TaskResponded(
        TaskResponse taskResponse,
        TaskResponseMetadata taskResponseMetadata
    );

    event TaskCompleted(uint32 indexed taskIndex);

    event TaskChallengedSuccessfully(
        uint32 indexed taskIndex,
        address indexed challenger
    );

    event TaskChallengedUnsuccessfully(
        uint32 indexed taskIndex,
        address indexed prover
    );

    event TaskChallenged(uint32 indexed taskIndex, Task task);

    event TaskRespondedWithProof(
        uint32 indexed taskIndex,
        uint256 output,
        address indexed prover
    );
}
