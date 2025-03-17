// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

interface ISertnServiceManager {
    // event NewTaskCreated(uint32 indexed taskIndex, Task task);

    // event TaskResponded(uint32 indexed taskIndex, Task task, address operator);

    // struct Task {
    //     string name;
    //     uint32 taskCreatedBlock;
    // }

    // function latestTaskNum() external view returns (uint32);

    // function allTaskHashes(
    //     uint32 taskIndex
    // ) external view returns (bytes32);

    // function allTaskResponses(
    //     address operator,
    //     uint32 taskIndex
    // ) external view returns (bytes memory);

    // function createNewTask(
    //     string memory name
    // ) external returns (Task memory);

    // function respondToTask(
    //     Task calldata task,
    //     uint32 referenceTaskIndex,
    //     bytes calldata signature
    // ) external;
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
}

interface ISertnServiceManagerTypes {

    struct Operator {
        uint8[] models_;
        bytes32[] computeUnits_;
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



