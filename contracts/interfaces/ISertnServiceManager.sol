// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;
import {IStrategy} from "../lib/eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Initializable} from "@openzeppelin-upgradeable/contracts/proxy/utils/Initializable.sol";

interface ISertnServiceManager {
    /// @notice Thrown when the caller is not an aggregator
    error NotAggregator();

    /// @notice Thrown when the caller is not the task manager
    error NotTaskManager();

    /// @notice Thrown when attempting to set a zero address
    error ZeroAddress();

    /// @notice Thrown when attempting to add an aggregator that already exists
    error AggregatorAlreadyExists();

    /**
     * @notice Update the task manager
     * @param _sertnTaskManager The address of the task manager to update to
     */
    function updateTaskManager(address _sertnTaskManager) external;

    /**
     * @notice Pull fee from a user
     * @param _user The address of the user to pull fee from
     * @param _fee The amount of fee to pull
     */
    function pullFeeFromUser(address _user, IERC20 _token, uint256 _fee) external;

    /**
     * @notice Update the model registry
     * @param _modelRegistry The address of the model registry to update to
     */
    function updateModelRegistry(address _modelRegistry) external;

    /**
     * @notice Add strategies to an operator set
     * @param _strategies The strategies to add
     * @param _operatorSetId The operator set id to add the strategies to
     */
    function addStrategies(IStrategy[] memory _strategies, uint32 _operatorSetId) external;

    /**
     * @notice Add an aggregator to the service manager
     * @param _aggregator The address of the aggregator to add
     */
    function addAggregator(address _aggregator) external;

    /**
     * @notice Remove an aggregator from the service manager
     * @param _aggregator The address of the aggregator to remove
     */
    function removeAggregator(address _aggregator) external;

    /**
     * @notice Task completed
     */
    function taskCompleted(
        address _operator,
        uint256 _fee,
        IStrategy _strategy,
        IERC20 _token,
        uint32 _startTimestamp
    ) external;

    /**
     * @notice Slash an operator for a task
     * @param _operator The address of the operator to slash
     * @param _fee The amount of tokens to slash
     * @param _operatorSetId The operator set id to slash
     * @param _strategy The strategy to slash
     */
    function slashOperator(
        address _operator,
        uint256 _fee,
        uint32 _operatorSetId,
        IStrategy _strategy
    ) external;

    /**
     * @notice Check if an address is registered as an aggregator
     * @param _aggregator The address to check
     * @return True if the address is an aggregator, false otherwise
     */
    function isAggregator(address _aggregator) external view returns (bool);
}
