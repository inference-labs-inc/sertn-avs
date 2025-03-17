// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {ISertnServiceManager} from "./ISertnServiceManager.sol";
import {ISertnServiceManagerTypes} from "./ISertnServiceManager.sol";
import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

abstract contract SertnServiceManagerStorage is ISertnServiceManager, ISertnServiceManagerTypes {

    uint256 PROOF_REQUEST_COST = 100;
    mapping(address => IStrategy) public tokenToStrategy;
    mapping(address => Operator) internal opInfo;
    mapping(address => bool) internal isAggregator;
    mapping(address => bool) internal isOperator;
    mapping(uint8 => Model) internal modelInfo;
    mapping(address => mapping(bytes32 => uint8)) internal computeUnits;
    mapping(bytes32 => uint8[]) internal modelsByName;
    mapping(address => bytes[]) internal openTasks;
    mapping(address => bytes[]) internal submittedTasks;
    mapping(bytes => bool) internal taskVerified;
    mapping(address => mapping(bytes => Task)) internal taskInfo;
    mapping(address => mapping(bytes => TaskResponse)) internal taskResponse;
    mapping(address => bytes[]) internal proofRequests;
    mapping(address => bytes[]) internal slashingQueue;
    mapping(address => uint256[]) internal allocatedEth;
    mapping(address => uint256) internal allocatedSer;
    mapping(address => uint32[2]) internal proofRequestExponents;
}