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
}

interface ISertnServiceManagerEvents {
    event NewOperator(address opAddr_);
    event NewModels(uint256[] modelId_);
    event NewStrategies(address[] newSupportedTokens_);
    event NewTask(address indexed opAddr_, bytes indexed taskId_);
    event TaskResponded(uint256 indexed model, bytes indexed taskId, ISertnServiceManagerTypes.TaskResponse task);
    event UpForSlashing(address indexed operator, bytes indexed taskId);
    event ProofRequested(address indexed operator, bytes indexed taskId);
    event OperatorSlashed(address indexed operator, bytes indexed taskId);
    event ModelUpdated(uint256 indexed modelId, ISertnServiceManagerTypes.OperatorModel operatorModel);
    event OpInfoChanged(address indexed _operator, bytes _opInfo);
    event OperatorDeleted(address indexed _operator, uint32[] opSetIds);
}

interface ISertnServiceManagerTypes {

    struct Model {
        string title_;
        string description_;
        address modelVerifier_;
        address[] operators_;
    }

    struct Operator {
        uint256[] models_;
        bytes32[] computeUnits_;
        bytes[] openTasks_;
        bytes[] submittedTasks_;
        bytes[] proofRequests_;
        uint256[] allocatedEth_;
        uint256 allocatedSer_;
        uint256[2] proofRequestExponents_;
        uint32 pausedBlock_;
    }

    struct OperatorModel {
        address operator_;
        uint256 modelId_;
        uint32 maxBlocks_;
        IStrategy[] ethStrategies_;
        uint256[] ethShares_;
        uint256 baseFee_;
        //May want to swap out in future for an operator's unique token.  For now will just use SER
        uint256 maxSer_;
        bytes32 computeType_;
        bool proveOnResponse_;
        bool available_;
    }

    struct Task {
        uint256 operatorModelId_;
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
