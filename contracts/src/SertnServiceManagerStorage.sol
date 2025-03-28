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
    mapping(uint256 => mapping(address => bytes)) public operatorModelInfo;
    // mapping(uint256 => Model) public modelInfo;
    mapping(address => mapping(bytes32 => uint256)) public computeUnits;
    mapping(bytes => bool) public taskVerified;
    mapping(bytes  => bytes) public taskResponse;
    mapping(address => bytes[]) public operatorSlashingQueue;
    mapping(bytes => address) public bountyHunter;
    mapping(uint256 => bytes) public modelInfo;
}
