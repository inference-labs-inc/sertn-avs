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
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
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
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableSet for EnumerableSet.AddressSet;

    uint32 deployTimestamp;

    IAllocationManager public allocationManager;
    IDelegationManager public delegationManager;
    IRewardsCoordinator public rewardsCoordinator;
    ISertnTaskManager public sertnTaskManager;
    IModelRegistry public modelRegistry;
    ISertnRegistrar public sertnRegistrar;

    // Operator info
    // mapping(address => bytes) public opInfo;
    // The number of nodes that a given operator has
    // mapping(address => uint256) public operatorNodeCount;
    // Compute units for a given operator-node
    // mapping(address => mapping(uint256 => uint256)) public operatorNodeComputeUnits;
    // Which models a given operator node supports
    // mapping(address => mapping(uint256 => mapping(uint256 => bool))) public operatorNodeModelIds;

    // Set of aggregators
    EnumerableSet.AddressSet internal aggregators;

    // Reward tracking
    // Here we track the rewards for each operator in each interval (interval is a time period, kinda one week)
    // We pay rewards to operators for each interval, so we need to track them separately
    // We might have multiple strategies, each one with its own reward token
    mapping(uint32 => EnumerableSet.AddressSet) internal operatorsInInterval; // all operators participated in the interval
    mapping(uint32 => EnumerableSet.AddressSet) internal strategiesInInterval; // strategies used in the interval
    // here we track the rewards in the interval for each operator, by strategy
    // schema: mapping(interval_id => mapping(operator => mapping(strategy => reward_count)))
    mapping(uint32 => mapping(address => mapping(address => uint256))) public intervalRewards;

    mapping(uint32 => bool) public intervalSubmitted;

    /**
     * @notice Modifier to ensure the caller is an aggregator
     */
    modifier onlyAggregators() {
        if (aggregators.contains(msg.sender)) {
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
        aggregators.add(msg.sender);

        deployTimestamp = uint32(block.timestamp);

        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);
        sertnRegistrar = ISertnRegistrar(_sertnRegistrar);
        _registerToEigen(_strategies, _avsMetadata);
    }

    /// @inheritdoc ISertnServiceManager
    function canSubmitRewardsForInterval(uint32 interval) external view returns (bool) {
        return
            interval < this.getCurrentInterval() &&
            !intervalSubmitted[interval] &&
            operatorsInInterval[interval].length() > 0;
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
    function _registerToEigen(IStrategy[] memory _strategies, string memory _avsMetadata) internal {
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
        allocationManager.addStrategiesToOperatorSet(address(this), operatorSetId, _strategies);
    }

    /// @inheritdoc ISertnServiceManager
    function addAggregator(address _aggregator) external onlyOwner {
        if (_aggregator == address(0)) {
            revert ZeroAddress();
        }
        if (aggregators.contains(_aggregator)) {
            revert AggregatorAlreadyExists();
        }
        aggregators.add(_aggregator);
    }

    /// @inheritdoc ISertnServiceManager
    function removeAggregator(address _aggregator) external onlyOwner {
        if (!aggregators.contains(_aggregator)) {
            revert NotAggregator();
        }
        aggregators.remove(_aggregator);
    }

    /// @inheritdoc ISertnServiceManager
    function isAggregator(address _aggregator) external view returns (bool) {
        return aggregators.contains(_aggregator);
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

    /// @inheritdoc ISertnServiceManager
    function taskCompleted(
        address _operator,
        uint256 _fee,
        IStrategy _strategy,
        uint32 _startTimestamp
    ) external onlyTaskManager nonReentrant {
        uint32 currentInterval = this.getCurrentInterval();

        if (!operatorsInInterval[currentInterval].contains(_operator)) {
            // If this operator is not in the current interval, we need to add it
            operatorsInInterval[currentInterval].add(_operator);
            // intervalRewards[currentInterval][_operator] = new mapping(address => uint256)();
        }
        if (!strategiesInInterval[currentInterval].contains(address(_strategy))) {
            // If this strategy is not in the current interval, we need to add it
            strategiesInInterval[currentInterval].add(address(_strategy));
        }

        // Accumulate rewards for this operator in this interval
        if (intervalRewards[currentInterval][_operator][address(_strategy)] == 0) {
            // First reward for this operator in this interval
            intervalRewards[currentInterval][_operator][address(_strategy)] = 0;
        }

        intervalRewards[currentInterval][_operator][address(_strategy)] += _fee;

        emit TaskRewardAccumulated(_operator, _fee, currentInterval);
    }

    /// @inheritdoc ISertnServiceManager
    function submitRewardsForInterval(uint32 interval) external onlyOwner {
        // require(!intervalSubmitted[interval], "Interval already submitted");
        // require(interval < rewardsCoordinator.getCurrentInterval(), "Interval not finished");
        // address[] memory operators = operatorsInInterval[interval];
        // require(operators.length > 0, "No operators to reward");
        // // Calculate total rewards for this interval
        // uint256 totalRewards = 0;
        // for (uint256 i = 0; i < operators.length; i++) {
        //     totalRewards += intervalRewards[interval][operators[i]];
        // }
        // // Create operator rewards array
        // IRewardsCoordinatorTypes.OperatorReward[]
        //     memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](
        //         operators.length
        //     );
        // for (uint256 i = 0; i < operators.length; i++) {
        //     operatorRewards[i] = IRewardsCoordinatorTypes.OperatorReward({
        //         operator: operators[i],
        //         amount: intervalRewards[interval][operators[i]]
        //     });
        //     // Deduct from pending rewards
        //     pendingRewards[operators[i]] -= intervalRewards[interval][operators[i]];
        // }
        // // Create strategies array
        // IRewardsCoordinatorTypes.StrategyAndMultiplier[]
        //     memory strategiesAndMultipliers = new IRewardsCoordinatorTypes.StrategyAndMultiplier[](
        //         1
        //     );
        // strategiesAndMultipliers[0] = IRewardsCoordinatorTypes.StrategyAndMultiplier({
        //     strategy: defaultStrategy,
        //     multiplier: 1e18 // 1.0 in WAD format. Maybe later we'll need different multipliers for different strategies
        // });
        // // Create submission
        // IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[]
        //     memory submissions = new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](
        //         1
        //     );
        // submissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
        //     strategiesAndMultipliers: strategiesAndMultipliers,
        //     token: rewardsToken,
        //     operatorRewards: operatorRewards,
        //     startTimestamp: interval * rewardsCoordinator.CALCULATION_INTERVAL_SECONDS(),
        //     duration: rewardsCoordinator.CALCULATION_INTERVAL_SECONDS(),
        //     description: string(abi.encodePacked("Sertn AVS rewards for interval ", interval))
        // });
        // // Approve tokens
        // rewardsToken.approve(address(rewardsCoordinator), totalRewards);
        // // Submit to EigenLayer
        // rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(address(this), submissions);
        // intervalSubmitted[interval] = true;
        // emit RewardsSubmittedForInterval(interval, totalRewards, operators.length);
    }

    // function autoSubmitPreviousIntervalRewards() external {
    //     uint32 currentInterval = rewardsCoordinator.getCurrentInterval();
    //     if (currentInterval > 0) {
    //         uint32 previousInterval = currentInterval - 1;
    //         if (canSubmitRewardsForInterval(previousInterval)) {
    //             submitRewardsForInterval(previousInterval);
    //         }
    //     }
    // }

    /// @inheritdoc ISertnServiceManager
    function getCurrentInterval() external view returns (uint32) {
        return
            uint32(
                (uint32(block.timestamp) - deployTimestamp) /
                    rewardsCoordinator.CALCULATION_INTERVAL_SECONDS()
            );
    }

    /// @inheritdoc ISertnServiceManager
    function getIntervalRewards(
        uint32 interval,
        address operator,
        address strategy
    ) external view returns (uint256) {
        return intervalRewards[interval][operator][strategy];
    }

    /// @inheritdoc ISertnServiceManager
    function getOperatorsInInterval(uint32 interval) external view returns (address[] memory) {
        return operatorsInInterval[interval].values();
    }

    /// @inheritdoc ISertnServiceManager
    function getStrategiesInInterval(uint32 interval) external view returns (address[] memory) {
        return strategiesInInterval[interval].values();
    }
}
