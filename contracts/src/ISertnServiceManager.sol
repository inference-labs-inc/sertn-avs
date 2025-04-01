// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";

interface ISertnServiceManager {
}

interface ISertnServiceManagerErrors {
    error IncorrectAVS();
    error IncorrectOperatorSetIds();
    error NotModelId();
    error NotAggregator();
    error NotOperator();
    error NotTaskManager();
    error InvalidTaskManager();
    error AggregatorAlreadyExists();
    error NoProofOnResponse();
    error TaskCouldNotBeSent();
    error IncorrectOperator();
    error TaskNotExpired();
    error TaskExpired();
    error NotPausedLongEnough();
    error NotRegisteredAggregatorOrOperator();
    error TransferFailed();
    error BountyAlreadySet();
    error NotInSubmittedTasks();
    error MaxBlocksTooLong();
    error NoPermission();
    error NotRegisteredToModel();
    error InvalidAggregator();
    error NodeModelExists();
    error NotNodeModel();
}

interface ISertnServiceManagerEvents {
    event NewOperator(address opAddr_);
    event NewModels(uint256[][] nodeAndModelIds_);
    event NewStrategies(address[] newSupportedTokens_);
    event NewTask(address indexed opAddr_, bytes indexed taskId_);
    event TaskResponded(uint256 indexed model, bytes32 indexed taskId, ISertnServiceManagerTypes.TaskResponse taskResponse);
    event UpForSlashing(address indexed operator, bytes32 indexed taskId);
    event ProofRequested(address indexed operator, bytes indexed taskId);
    event OperatorSlashed(bytes indexed taskId);
    event ModelUpdated(bytes32 indexed nodeModelId, ISertnServiceManagerTypes.NodeModel nodeModel);
    event OpInfoChanged(address indexed _operator, bytes _opInfo);
    event OperatorDeleted(address indexed _operator, uint32[] opSetIds);
    event TaskExpiryChanged(uint256 _taskExpiryBlocks);
}

interface ISertnServiceManagerTypes {

    struct Model {
        string title_;
        string description_;
        address modelVerifier_;
        bytes32 modelKey_;
        address[] operators_;
    }

    struct Operator {
        uint256[][] nodesAndModels_;
        uint256[] nodeComputeUnits_;
        bytes32[] submittedTasks_;
        uint256[] allocatedEth_;
        uint256 allocatedSer_;
        //will get deeper into proof request function in future, right now for node 0 cost of proof request = pRC[0]/pRC[1] * PROOF_REQUEST_COST
        uint256[] proofRequestCoefficients_;
        uint32 pausedBlock_;
    }

    struct NodeTasks {
        bytes[] openTasks_;
        bytes32[] proofRequests_;
    }

    struct NodeModel {
        uint32 maxBlocks_;
        uint256[] ethShares_;
        uint256 baseFee_;
        //May want to swap out in future for an operator's unique token.  For now will just use SER
        uint256 maxSer_;
        uint256 compute_;
        bool proveOnResponse_;
        bool available_;
    }

    struct Task {
        address operator_;
        uint256 node_;
        uint256 modelId_;
        bytes inputs_;
        uint256 poc_;
        uint256 startTime_;
        uint32 startingBlock_;
        bool proveOnResponse_;
        address user_;
    }

    struct TaskResponse {
        bytes taskId_;
        bytes output_;
        bool proven_;
    }

}
