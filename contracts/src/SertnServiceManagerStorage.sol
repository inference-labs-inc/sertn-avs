// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {ISertnServiceManager} from "./ISertnServiceManager.sol";
import {ISertnServiceManagerTypes} from "./ISertnServiceManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";

abstract contract SertnServiceManagerStorage is ISertnServiceManager, ISertnServiceManagerTypes {

    uint256 PROOF_REQUEST_COST = 100;
    uint256 TASK_EXPIRY_BLOCKS = 1e3;
    uint256 BOUNTY = 500;

    mapping(address => IStrategy) public tokenToStrategy;
    mapping(address => Operator) internal opInfo;
    mapping(address => bool) internal isAggregator;
    mapping(address => bool) internal isOperator;
    mapping(uint8 => Model) internal modelInfo;
    mapping(address => mapping(bytes32 => uint8)) internal computeUnits;
    mapping(bytes32 => uint8[]) internal modelsByName;
    mapping(bytes => bool) internal taskVerified;
    mapping(bytes  => TaskResponse) internal taskResponse;
    mapping(address => bytes[]) internal operatorSlashingQueue;
    mapping(bytes => address) internal bountyHunter;
}