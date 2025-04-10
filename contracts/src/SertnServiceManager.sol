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

    // Operator info
    mapping(address => bytes) public opInfo;
    // Mapping of aggregators
    mapping(address => bool) public isAggregator;
    // The number of nodes that a given operator has
    mapping(address => uint256) public operatorNodeCount;
    // Compute units for a given operator-node
    mapping(address => mapping(uint256 => uint256))
        public operatorNodeComputeUnits;
    // Which models a given operator node supports
    mapping(address => mapping(uint256 => mapping(uint256 => bool)))
        public operatorNodeModelIds;

    /**
     * @notice Modifier to ensure the caller is an aggregator
     */
    modifier onlyAggregators() {
        if (!isAggregator[msg.sender]) {
            revert NotAggregator();
        }
        _;
    }

    /**
     * @notice Modifier to ensure the caller is the task manager
     */
    modifier onlyTaskManager() {
        if (msg.sender != address(sertnTaskManager)) {
            revert NotTaskManager();
        }
        _;
    }

    /**
     * @notice Modifier to ensure the caller is the sertn registrar
     */
    modifier onlySertnRegistrar() {
        if (msg.sender != address(sertnRegistrar)) {
            revert NotSertnRegistrar();
        }
        _;
    }

    /**
     * @notice Initializes the SertnServiceManager
     * @param _rewardsCoordinator The address of the rewards coordinator
     * @param _delegationManager The address of the delegation manager
     * @param _allocationManager The address of the allocation manager
     * @param _sertnRegistrar The address of the sertn registrar
     * @param _strategies The strategies to add to the operator set
     * @param _avsMetadata The metadata for the AVS
     */
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

    /**
     * @notice Updates the task manager
     * @param _sertnTaskManager The address of the task manager
     */
    function updateTaskManager(address _sertnTaskManager) external onlyOwner {
        if (_sertnTaskManager == address(0)) {
            revert InvalidTaskManager();
        }
        sertnTaskManager = SertnTaskManager(_sertnTaskManager);
    }

    /**
     * @notice Updates the model registry
     * @param _modelRegistry The address of the model registry
     */
    function updateModelRegistry(address _modelRegistry) external onlyOwner {
        if (_modelRegistry == address(0)) {
            revert InvalidModelRegistry();
        }
        modelRegistry = IModelRegistry(_modelRegistry);
    }

    /**
     * @notice Registers the service manager to EigenLayer
     * @param _strategies The strategies to add to the operator set
     * @param _avsMetadata The metadata for the AVS
     */
    function _registerToEigen(
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) internal {
        allocationManager.updateAVSMetadataURI(address(this), _avsMetadata);
        allocationManager.createOperatorSets(
            address(this),
            abi.decode(
                abi.encode(
                    IAllocationManagerTypes.CreateSetParams({
                        operatorSetId: 0,
                        strategies: _strategies
                    })
                ),
                (IAllocationManagerTypes.CreateSetParams[])
            )
        );
        allocationManager.setAVSRegistrar(address(sertnRegistrar));
    }

    function addStrategies(
        IStrategy[] memory _strategies,
        uint256 operatorSetId
    ) external onlyOwner {
        allocationManager.addStrategiesToOperatorSet(
            address(this),
            operatorSetId,
            _strategies
        );
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
}
