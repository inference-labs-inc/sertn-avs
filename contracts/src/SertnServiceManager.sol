// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

// Sertn
import {ISertnServiceManager} from "../interfaces/ISertnServiceManager.sol";
import {ISertnTaskManager} from "../interfaces/ISertnTaskManager.sol";
import {ISertnRegistrar} from "../interfaces/ISertnRegistrar.sol";
import {IModelRegistry} from "../interfaces/IModelRegistry.sol";
import {IVerifier} from "../interfaces/IVerifier.sol";
// EigenLayer
import {IRewardsCoordinator} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {IAllocationManager, IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";

import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
// OpenZeppelin
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgradeable/contracts/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin-upgradeable/contracts/utils/ReentrancyGuardUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title Primary entrypoint for procuring services from Sertn.
 * @author Inference Labs, Inc.
 */

contract SertnServiceManager is
    ISertnServiceManager,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    IAllocationManager public allocationManager;
    IDelegationManager public delegationManager;
    IRewardsCoordinator public rewardsCoordinator;
    ISertnTaskManager public sertnTaskManager;
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

    // List of aggregators
    address[] public aggregators;

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

    /// @inheritdoc ISertnServiceManager
    function initialize(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        address _sertnRegistrar,
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) public initializer {
        __Ownable_init(msg.sender);
        __ReentrancyGuard_init();
        // Set the deployer as an aggregator
        isAggregator[msg.sender] = true;

        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);
        sertnRegistrar = ISertnRegistrar(_sertnRegistrar);
        _registerToEigen(_strategies, _avsMetadata);
    }

    /// @inheritdoc ISertnServiceManager
    function updateTaskManager(address _sertnTaskManager) external onlyOwner {
        if (_sertnTaskManager == address(0)) {
            revert InvalidTaskManager();
        }
        sertnTaskManager = ISertnTaskManager(_sertnTaskManager);
    }

    /// @inheritdoc ISertnServiceManager
    function updateModelRegistry(address _modelRegistry) external onlyOwner {
        if (_modelRegistry == address(0)) {
            revert ZeroAddress();
        }
        modelRegistry = IModelRegistry(_modelRegistry);
    }

    /// @inheritdoc ISertnServiceManager
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
        allocationManager.setAVSRegistrar(address(this), sertnRegistrar);
    }

    /// @inheritdoc ISertnServiceManager
    function addStrategies(
        IStrategy[] memory _strategies,
        uint32 operatorSetId
    ) external onlyOwner {
        allocationManager.addStrategiesToOperatorSet(
            address(this),
            operatorSetId,
            _strategies
        );
    }

    /// @inheritdoc ISertnServiceManager
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

    /// @inheritdoc ISertnServiceManager
    function removeAggregator(address _aggregator) external onlyOwner {
        if (!isAggregator[_aggregator]) {
            revert NotAggregator();
        }
        for (uint256 i = 0; i < aggregators.length; i++) {
            if (aggregators[i] == _aggregator) {
                aggregators[i] = aggregators[aggregators.length - 1];
                aggregators.pop();
            }
        }
    }
}
