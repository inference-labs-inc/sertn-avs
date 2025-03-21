// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";

interface ISertnServiceManager {
}

interface ISertnServiceManagerErrors {
    error WrongAVS();
    error NotModelId(string);
    error NoProofOnResponse(string);
    error TaskCouldNotBeSent(string);
}

interface ISertnServiceManagerEvents {
    event newOperator(address opAddr_);
    event newModels(uint8[] modelId_);
    event newStrategies(address[] newSupportedTokens_);
    event newTask(address indexed opAddr_, bytes indexed taskId_);
    event taskResponded(uint8 indexed model, bytes indexed taskId, ISertnServiceManagerTypes.TaskResponse task);
    event upForSlashing(address indexed operator, bytes indexed taskId);
    event proofRequested(address indexed operator, bytes indexed taskId);
    event operatorSlashed(address indexed operator, bytes indexed taskId);
    event modelUpdated(uint8 indexed modelId, ISertnServiceManagerTypes.Model model);
    event opInfoChanged(address indexed _operator, ISertnServiceManagerTypes.Operator _opInfo);
    event operatorDeleted(address indexed _operator, uint32[] opSetIds);
}

interface ISertnServiceManagerTypes {

    struct Operator {
        uint8[] models_;
        bytes32[] computeUnits_;
        bytes[] openTasks_;
        bytes[] submittedTasks_;
        bytes[] proofRequests_;
        uint256[] allocatedEth_;
        uint256 allocatedSer_;
        uint32[2] proofRequestExponents_;
        uint32 pausedBlock_;
    }

    struct Model {
        bytes32 modelName_;
        address operator_;
        address verifier_;
        string benchData_;
        bytes inputs_;
        bytes output_;
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
        uint8 modelId_;
        bytes inputs_;
        uint256 poc_;
        uint256 startTime_;
        uint256 startingBlock_;
        bool proveOnResponse_; 
        address user_;
    }

    struct TaskResponse {
        bytes taskId_;
        bytes output_;
        bool proven_;
    }

}



