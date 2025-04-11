// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";
import {IAllocationManager} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IRewardsCoordinator} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {ISertnServiceManager} from "../interfaces/ISertnServiceManager.sol";
import {ISertnTaskManager} from "../interfaces/ISertnTaskManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {IVerifier} from "../interfaces/IVerifier.sol";
import {ModelRegistry} from "./ModelRegistry.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract SertnTaskManager is OwnableUpgradeable, ISertnTaskManager {
    // queue of tasks that are waiting to be assigned to an operator
    address[] public operators;
    // queue of tasks that are waiting to be challenged
    bytes[] public slashingQueue;
    // nonce for tasks to ensure uniqueness
    uint256 taskNonce;

    IERC20 public ser;

    IAllocationManager public allocationManager;
    IDelegationManager public delegationManager;
    IRewardsCoordinator public rewardsCoordinator;

    ISertnServiceManager public sertnServiceManager;
    ModelRegistry public modelRegistry;

    modifier onlyAggregators() {
        if (!sertnServiceManager.isAggregator(msg.sender)) {
            revert NotAggregator();
        }
        _;
    }

    modifier onlyOperators() {
        if (!sertnServiceManager.isOperator(msg.sender)) {
            revert NotOperator();
        }
        _;
    }

    function initialize(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        address _sertnServiceManager,
        address _modelRegistry
    ) public initializer {
        __Ownable_init();
        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);
        sertnServiceManager = ISertnServiceManager(_sertnServiceManager);
        modelRegistry = ModelRegistry(_modelRegistry);
    }

    function sendTask(Task memory _task) external onlyAggregators {
        // TODO: Implement
    }

    function submitTaskOutput(
        bytes32 taskId,
        bytes calldata output
    ) external onlyOperators {
        // TODO: Implement
    }

    function challengeTask(bytes32 taskId) external onlyAggregators {
        // TODO: Implement
    }

    function submitProofForTask(
        bytes32 taskId,
        bytes calldata proof
    ) external onlyOperators {
        // TODO: Implement
    }
}
