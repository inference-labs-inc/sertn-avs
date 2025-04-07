// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../interfaces/ISertnServiceManager.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {ISertnRegistrar} from "../interfaces/ISertnRegistrar.sol";

import {IAllocationManager} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IERC20} from "@openzeppelin/contracts/interfaces/IERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";

import {IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IVerifier} from "../interfaces/IVerifier.sol";
import "@eigenlayer/contracts/libraries/OperatorSetLib.sol";

import "@openzeppelin-upgrades/contracts/utils/math/MathUpgradeable.sol";

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {SertnTaskManager} from "./SertnTaskManager.sol";
import {ModelRegistry} from "./ModelRegistry.sol";

import {Test, console2 as console} from "forge-std/Test.sol";

/**
 * @title Primary entrypoint for procuring services from Sertn.
 * @author Inference Labs, Inc.
 */

contract SertnServiceManager is
    ISertnServiceManager,
    SertnServiceManagerStorage,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    IAllocationManager public allocationManager;
    IDelegationManager public delegationManager;
    IRewardsCoordinator public rewardsCoordinator;
    SertnTaskManager public sertnTaskManager;
    IModelRegistry public modelRegistry;
    ISertnRegistrar public sertnRegistrar;

    uint256 public PROOF_REQUEST_COST = 100;
    uint32 public TASK_EXPIRY_BLOCKS = 1e3;
    uint256 public BOUNTY = 500;

    mapping(address => IStrategy) public tokenToStrategy;
    mapping(address => bytes) public opInfo;
    mapping(address => bool) public isAggregator;
    mapping(address => bool) public isOperator;
    mapping(uint256 => bytes) public operatorModelInfo;
    mapping(address => mapping(bytes32 => uint256)) public computeUnits;
    mapping(uint256 => uint256[]) public modelsByName;
    mapping(bytes => bool) public taskVerified;
    mapping(bytes  => bytes) public taskResponse;
    mapping(address => bytes[]) public operatorSlashingQueue;
    mapping(bytes => address) public bountyHunter;
    mapping(uint256 => address) public modelAddresses;
    mapping(address => bytes) public modelVerifiers;

    // The number of nodes that a given operator has
    mapping(address => uint256) public operatorNodeCount;
    // Compute units for a given operator-node
    mapping(address => mapping(uint256 => uint256)) public operatorNodeComputeUnits;
    // Which models a given operator node supports
    mapping(address => mapping(uint256 => mapping(uint256 => bool))) public operatorNodeModelIds;

    modifier onlyAggregators() {
        if (!isAggregator[msg.sender]) {
            revert NotAggregator();
        }
        _;
    }

    modifier onlyTaskManager() {
        if (msg.sender != address(sertnTaskManager)) {
            revert NotTaskManager();
        }
        _;
    }

    modifier onlySertnRegistrar() {
        if (msg.sender != address(sertnRegistrar)) {
            revert NotSertnRegistrar();
        }
        _;
    }

    function initialize(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        address _sertnRegistrar,
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) public initializer {
        __Ownable_init();
        __ReentrancyGuard_init();
        // Set the deployer as an aggregator
        isAggregator[msg.sender] = true;

        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);
        sertnRegistrar = ISertnRegistrar(_sertnRegistrar);
        _registerToEigen(_strategies, _avsMetadata);
    }

    function updateTaskManager(address _sertnTaskManager) external onlyOwner {
        if (_sertnTaskManager == address(0)) {
            revert InvalidTaskManager();
        }
        sertnTaskManager = SertnTaskManager(_sertnTaskManager);
    }

    function updateModelRegistry(address _modelRegistry) external onlyOwner {
        if (_modelRegistry == address(0)) {
            revert InvalidModelRegistry();
        }
        modelRegistry = IModelRegistry(_modelRegistry);
    }

    function _registerToEigen(
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) internal {
        allocationManager.updateAVSMetadataURI(address(this), _avsMetadata);
        IAllocationManagerTypes.CreateSetParams[]
            memory setParams = new IAllocationManagerTypes.CreateSetParams[](1);
        setParams[0] = IAllocationManagerTypes.CreateSetParams({
            operatorSetId: 0,
            strategies: _strategies
        });
        allocationManager.createOperatorSets(address(this), setParams);
        allocationManager.setAVSRegistrar(address(sertnRegistrar));
        _addStrategies(_strategies, false);
    }

    function _addStrategies(
        IStrategy[] memory _strategies, bool _newStrategies
    ) internal {
        for (uint256 i; i < _strategies.length;) {
            tokenToStrategy[
                address(_strategies[i].underlyingToken())
            ] = _strategies[i];
            unchecked { ++i; }
        }
        if (_newStrategies) {
            allocationManager.addStrategiesToOperatorSet(
            address(this),
            0,
            _strategies
        );
        }
    }
    function addStrategies(
        IStrategy[] memory _strategies, bool _newStrategies
    ) external onlyAggregators() {
        _addStrategies(_strategies, _newStrategies);
    }

    function addAggregator(address _aggregator) external onlyOwner {
        if (_aggregator == address(0)) {
            revert InvalidAggregator();
        }
        if (isAggregator[_aggregator]) {
            revert AggregatorAlreadyExists();
        }
        aggregators.push(_aggregator);
        isAggregator[_aggregator] = true;
    }


    function pushToSlashingQueue(bytes calldata _taskId) external onlyTaskManager() {
        slashingQueue.push(_taskId);
    }

    function pushToOperatorSlashingQueue(address _operator, bytes calldata _taskId) external onlyTaskManager() {
        operatorSlashingQueue[_operator].push(_taskId);
    }

    function setTaskResponse(TaskResponse calldata _taskResponse) external onlyTaskManager() {
        taskResponse[_taskResponse.taskId_] = abi.encode(_taskResponse);
    }

    function setBountyHunter(bytes calldata _taskId, address _bountyHunter) external onlyTaskManager() {
        bountyHunter[_taskId] = _bountyHunter;
    }

    function getOperatorInfo(address _operator) public view returns (Operator memory) {
        return abi.decode(opInfo[_operator], (Operator));
    }
}
