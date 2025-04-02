// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {ISertnServiceManager} from "./ISertnServiceManager.sol";
import {ISertnServiceManagerTypes} from "./ISertnServiceManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";

abstract contract SertnServiceManagerStorage is ISertnServiceManager, ISertnServiceManagerTypes {

    uint256 public PROOF_REQUEST_COST = 100;
    uint32 public taskExpiryBlocks = 1e3;
    uint256 public BOUNTY = 500;
    uint32 public taskExpiryBuffer = 1e2;
    uint256 public proofRequestIncrement = 500;
    uint256 public startingCoefficients = 1e3;

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
    mapping(bytes32 => bool) public operatorRegistered;
    mapping(address => uint256) public rewardsAmount;
    mapping(bytes32 => uint256[]) public slashableAmount;

    //Operator mappings
    mapping(bytes32 => uint256[]) nodesAndModels;
    mapping(bytes32 => uint256) public nodeComputeUnits;
    mapping(bytes32 => bool) public taskSubmitted;
    mapping(address => uint256[]) public allocatedStake;
    mapping(address => uint256[2]) public proofRequestCoefficients;
    mapping(address => uint32) public operatorPausedBlock; 

    //NodeTask mappings
    mapping(bytes32 => bytes32[]) public openTasks;
    mapping(bytes32 => bool) public isOpenTask;
    mapping(bytes32 => bytes32[]) public proofRequests;
    mapping(bytes32 => bool) public proofRequested;

    //NodeModel mappings
    mapping(bytes32 => uint32) public maxBlocks;
    mapping(bytes32 => uint256) public ethShares;
    mapping(bytes32 => uint256) public baseFee;
    mapping(bytes32 => uint256) public maxSer;
    mapping(bytes32 => uint256) public compute;
    mapping(bytes32 => bool) public proveOnResponse;
    mapping(bytes32 => bool) public nodeModelAvailable;

    //Task mapping
    mapping(bytes32 => address) public taskOperator;
    mapping(bytes32 => bytes32) public taskNodeOperatorId;
    mapping(bytes32 => bytes32) public taskNodeModelId;
    mapping(bytes32 => uint32) public startingBlock;
    mapping(bytes32 => uint256) public taskPoc;
}
