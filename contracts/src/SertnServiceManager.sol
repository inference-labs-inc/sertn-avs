// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

// Sertn
import {ISertnServiceManager} from "../interfaces/ISertnServiceManager.sol";
import {ISertnTaskManager} from "../interfaces/ISertnTaskManager.sol";
import {ISertnRegistrar} from "../interfaces/ISertnRegistrar.sol";
import {IModelRegistry} from "../interfaces/IModelRegistry.sol";
import {IVerifier} from "../interfaces/IVerifier.sol";
// EigenLayer
import {IRewardsCoordinator, IRewardsCoordinatorTypes} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {IAllocationManager, IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
// OpenZeppelin
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgradeable/contracts/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin-upgradeable/contracts/security/ReentrancyGuardUpgradeable.sol";
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

    /// @inheritdoc ISertnServiceManager
    function updateTaskManager(address _sertnTaskManager) external onlyOwner {
        if (_sertnTaskManager == address(0)) {
            revert ZeroAddress();
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

    /**
     * @notice Register the AVS to EigenLayer
     * @param _strategies The strategies to register
     * @param _avsMetadata The AVS metadata to register
     */
    function _registerToEigen(
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) internal {
        allocationManager.updateAVSMetadataURI(address(this), _avsMetadata);

        // prepare the params for creating operator sets:
        IAllocationManagerTypes.CreateSetParams[]
            memory params = new IAllocationManagerTypes.CreateSetParams[](1);
        params[0] = IAllocationManagerTypes.CreateSetParams({
            operatorSetId: 0,
            strategies: _strategies
        });
        allocationManager.createOperatorSets(address(this), params);
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
            revert ZeroAddress();
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

    /// @inheritdoc ISertnServiceManager
    function pullFeeFromUser(
        address _user,
        IERC20 _token,
        uint256 _fee
    ) external onlyTaskManager nonReentrant {
        _token.transferFrom(_user, address(this), _fee);
    }

    /// @inheritdoc ISertnServiceManager
    function slashOperator(
        address _operator,
        uint256 _fee,
        uint32 _operatorSetId,
        IStrategy _strategy
    ) external onlyTaskManager nonReentrant {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = _strategy;
        uint256[] memory wadsToSlash = new uint256[](1);
        wadsToSlash[0] = _fee;
        allocationManager.slashOperator(
            address(this),
            IAllocationManagerTypes.SlashingParams({
                operator: _operator,
                operatorSetId: _operatorSetId,
                strategies: strategies,
                wadsToSlash: wadsToSlash,
                description: "Failure to provide a proof or output for a task"
            })
        );
    }

    function taskCompleted(
        address _operator,
        uint256 _fee,
        IStrategy _strategy,
        IERC20 _token
    ) external onlyTaskManager nonReentrant {
        IRewardsCoordinatorTypes.StrategyAndMultiplier[]
            memory strategiesAndMultipliers = new IRewardsCoordinatorTypes.StrategyAndMultiplier[](
                1
            );
        strategiesAndMultipliers[0] = IRewardsCoordinatorTypes
            .StrategyAndMultiplier({strategy: _strategy, multiplier: 1});

        IRewardsCoordinatorTypes.OperatorReward[]
            memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](
                1
            );
        operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
            operator: _operator,
            amount: _fee
        });

        IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[]
            memory submissions = new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](
                1
            );
        submissions[0] = IRewardsCoordinatorTypes
            .OperatorDirectedRewardsSubmission({
                strategiesAndMultipliers: strategiesAndMultipliers,
                token: _token,
                operatorRewards: operatorRewards,
                // TODO: I'm not sure how to figure out params below
                //       so they just picked up to have tests running :-P
                //       values here must be compatible with the constraints
                //       of the rewards coordinator
                // @see "@eigenlayer/contracts/core/RewardsCoordinator.sol" - _validateCommonRewardsSubmission method
                startTimestamp: rewardsCoordinator
                    .CALCULATION_INTERVAL_SECONDS(), // uint32(block.timestamp),
                duration: rewardsCoordinator.CALCULATION_INTERVAL_SECONDS(), // 12
                description: "Compensation for task completed"
            });

        // Approve the rewards coordinator to spend the fee
        for (uint256 i = 0; i < submissions.length; i++) {
            uint256 _rewards = 0;
            for (
                uint256 j = 0;
                j < submissions[i].operatorRewards.length;
                j++
            ) {
                _rewards += submissions[i].operatorRewards[j].amount;
            }
            submissions[i].token.approve(address(rewardsCoordinator), _rewards);
        }

        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(
            address(this),
            submissions
        );
    }
}
