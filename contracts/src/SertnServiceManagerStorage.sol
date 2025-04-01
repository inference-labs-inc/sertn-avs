// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {ISertnServiceManager} from "./ISertnServiceManager.sol";
import {ISertnServiceManagerTypes} from "./ISertnServiceManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";

abstract contract SertnServiceManagerStorage is ISertnServiceManager, ISertnServiceManagerTypes {

    uint256 public PROOF_REQUEST_COST = 100;
    uint32 public taskExpiryBlocks = 1e3;
    uint256 public BOUNTY = 500;

    mapping(address => IStrategy) public tokenToStrategy;
    mapping(address => bytes) public opInfo;
    mapping(address => bool) public isAggregator;
    mapping(address => bool) public isOperator;
    mapping(bytes32 => bytes) public nodeModelInfo;
    mapping(bytes32 => bytes) public nodeTasks;
    mapping(bytes32 => bool) public taskVerified;
    mapping(bytes32  => bytes) public taskResponse;
    mapping(address => bytes32[]) public operatorSlashingQueue;
    mapping(bytes32 => address) public bountyHunter;
    mapping(uint256 => bytes) public modelInfo;
}
