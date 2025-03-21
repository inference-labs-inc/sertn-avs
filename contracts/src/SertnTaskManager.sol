// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./ISertnServiceManager.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {SertnServiceManagerStorage} from "./SertnServiceManagerStorage.sol";

import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";

import {IAllocationManager} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IERC20} from "@openzeppelin/contracts/interfaces/IERC20.sol";

import {IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IVerifier} from "./IVerifier.sol";
import "@eigenlayer/contracts/libraries/OperatorSetLib.sol";

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";

import {Test, console2 as console} from "forge-std/Test.sol";

import {SertnServiceManager} from "./SertnServiceManager.sol";

contract SertnTaskManager is
    SertnServiceManagerStorage,
    ISertnServiceManagerErrors,
    ISertnServiceManagerEvents,
    OwnableUpgradeable
{
    address[] aggregators;
    address[] operators;
    bytes[] slashingQueue;

    IERC20 ser;

    IAllocationManager allocationManager;
    IDelegationManager delegationManager;
    IRewardsCoordinator rewardsCoordinator;
    OperatorSet opSet;

    SertnServiceManager sertnServiceManager;

    modifier onlyAggregators() {
        if (!sertnServiceManager.isAggregator(msg.sender)) {
            revert NotAggregator();
        }
        _;
    }

    constructor(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        address _sertnServiceManager,
        address _ser
    ) OwnableUpgradeable() {
        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);
        sertnServiceManager = SertnServiceManager(_sertnServiceManager);

        opSet = OperatorSet({avs: address(sertnServiceManager), id: 0});
        ser = IERC20(_ser);
    }

    function sendTask(Task memory _task) external {
        uint8 _modelId = _task.modelId_;
        _task.startTime_ = block.timestamp;
        _task.startingBlock_ = block.number;

        if (0 > _modelId || sertnServiceManager.getNumModels() < _modelId) {
            revert NotModelId("Not Model Id");
        }
        Model memory _model = sertnServiceManager.getModelInfo(_modelId);

        if (_task.proveOnResponse_ && !_model.proveOnResponse_) {
            revert NoProofOnResponse("Prove On Response Not Available");
        }

        if (_task.proveOnResponse_) {
            _checkFinancialSecurity(_task.poc_, _model, _model.maxBlocks_);
        } else {
            _checkFinancialSecurity(
                _task.poc_ * 10,
                _model,
                uint32(TASK_EXPIRY_BLOCKS)
            );
        }

        Operator memory _operator = sertnServiceManager.getOperatorInfo(_model.operator_);

        if (
            _model.available_ &&
            sertnServiceManager.getComputeUnits(_model.operator_,_model.computeType_) > 0 &&
            _operator.pausedBlock_ == 0
        ) {
            uint256 transferAmount = (1.5e3 * (_model.baseFee_ + _task.poc_)) / 1e3;
            if (!ser.transferFrom(msg.sender, address(sertnServiceManager), transferAmount)) {
                revert TransferFailed();
            }

            bytes memory _taskId = abi.encode(_task);
            _operator.openTasks_.push(_taskId);
            if (_task.proveOnResponse_) {
                _operator.proofRequests_.push(_taskId);
            }
            sertnServiceManager.setComputeUnits(_model.operator_,_model.computeType_, false);

            for (uint8 i = 0; i < _operator.allocatedEth_.length; i++) {
                _operator.allocatedEth_[i] += _model.ethShares_[i];
            }
            _operator.allocatedSer_ += 10 * _task.poc_;
            sertnServiceManager.updateOperator(_model.operator_, _operator);
            emit newTask(_model.operator_, _taskId);
        } else {
            revert TaskCouldNotBeSent("Unknown error sending task");
        }
    }

    function _checkFinancialSecurity(
        uint256 _poc,
        Model memory _model,
        uint32 _securityDuration
    ) internal view {
        address[] memory _operators = new address[](1);
        _operators[0] = _model.operator_;
        //check if secure using max blocks as model param
        uint256[] memory _ethSecurity = allocationManager
            .getMinimumSlashableStake(
                opSet,
                _operators,
                _model.ethStrategies_,
                uint32(block.number) + _securityDuration
            )[0];

        Operator memory _operator = sertnServiceManager.getOperatorInfo(_model.operator_);

        for (uint8 i = 0; i < _ethSecurity.length; i++) {
            if (
                _ethSecurity[i] <
                _model.ethShares_[i] + _operator.allocatedEth_[i]
            ) {
                revert TaskCouldNotBeSent("Not enough eth backed security");
            }
        }
        IStrategy[] memory _pocStrategy = new IStrategy[](1);
        _pocStrategy[0] = tokenToStrategy[address(ser)];
        uint256 _serSecurity = allocationManager.getMinimumSlashableStake(
            opSet,
            _operators,
            _pocStrategy,
            uint32(block.number) + _securityDuration
        )[0][0];

        if (
            _model.maxSer_ < _poc ||
            _serSecurity < _poc + _operator.allocatedSer_
        ) {
            revert TaskCouldNotBeSent("Insufficient economic security backing");
        }
    }

    function submitTask(TaskResponse memory _taskResponse, bool _verification, bytes memory _proof) external {

        Task memory _task = abi.decode(_taskResponse.taskId_, (Task));

        uint256 _modelId = _task.modelId_;

        if (0 > _modelId || sertnServiceManager.getNumModels() < _modelId) {
            revert NotModelId("Invalid Model ID");
        }

        Model memory _model = modelInfo[_modelId];

        if (msg.sender != _model.operator_) {
            revert IncorrectOperator();
        }

        Operator memory _operator = opInfo[_model.operator_];

        if (_model.maxBlocks_ > uint32(block.number - _task.startingBlock_)) {
            revert TaskExpired();
        }

        if (_verification) {
            _checkFinancialSecurity(_task.poc_, _model, 0);
        } else {
            _checkFinancialSecurity(10*_task.poc_, _model, uint32(TASK_EXPIRY_BLOCKS - (block.number - _task.startingBlock_)));
        }

        if (_taskResponse.proven_) {
            if (sertnServiceManager.isTaskVerified(_taskResponse.taskId_)) {
                _taskResponse.proven_ = true;
                sertnServiceManager.clearTask(_taskResponse.taskId_);
            }

            else {
                sertnServiceManager.pushToOperatorSlashingQueue(msg.sender, _taskResponse.taskId_);
                sertnServiceManager.pushToSlashingQueue(_taskResponse.taskId_);
                emit upForSlashing(_model.operator_, _taskResponse.taskId_);
                return;
            }
        } else {
            if (_verification || _task.proveOnResponse_) {
            sertnServiceManager._verifyTask(_taskResponse.taskId_, _proof);

            if (sertnServiceManager.isTaskVerified(_taskResponse.taskId_)) {
                _taskResponse.proven_ = true;
                sertnServiceManager.clearTask(_taskResponse.taskId_);
            }

            else {
                sertnServiceManager.pushToOperatorSlashingQueue(msg.sender, _taskResponse.taskId_);
                sertnServiceManager.pushToSlashingQueue(_taskResponse.taskId_);
                emit upForSlashing(_model.operator_, _taskResponse.taskId_);
                return;
            }
        }
        }

        _operator.submittedTasks_.push(_taskResponse.taskId_);

        sertnServiceManager.updateOperator(_model.operator_, _operator);
        sertnServiceManager.setTaskResponse(_taskResponse);
        sertnServiceManager.setComputeUnits(_model.operator_, _model.computeType_, true);

        emit taskResponded(_modelId, _taskResponse.taskId_, _taskResponse);
    }

    function requestProof(bytes memory _taskId) external {
        //add probabilistic case
        if (sertnServiceManager.getBountyHunter(_taskId) != address(0)) {
            revert BountyAlreadySet();
        }
        Task memory _task = abi.decode(_taskId, (Task));
        Model memory _model = sertnServiceManager.getModelInfo(_task.modelId_);
        Operator memory _operator = sertnServiceManager.getOperatorInfo(_model.operator_);

        uint256 _amount = PROOF_REQUEST_COST*(_operator.proofRequestExponents_[0]/_operator.proofRequestExponents_[1]);
        ser.transferFrom(msg.sender, address(sertnServiceManager), _amount);

        _operator.proofRequests_.push(_taskId);

        _operator.proofRequestExponents_[0] += 500;

        sertnServiceManager.updateOperator(_model.operator_, _operator);
        sertnServiceManager.setBountyHunter(_taskId, msg.sender);

        emit proofRequested(_model.operator_, _taskId);

    }

    function sendRewards(IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] calldata operatorDirectedRewardsSubmissions, bytes[][] memory _rewardedTasks) external onlyAggregators() {
        uint256 _approvalAmount = 0;
        for (uint8 i = 0; i < operatorDirectedRewardsSubmissions[0].operatorRewards.length; i ++) {
            Operator memory _operator = sertnServiceManager.getOperatorInfo(operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator);
            for (uint8 j = 0; j < _rewardedTasks[i].length; j ++) {
                Task memory _task = abi.decode(_rewardedTasks[i][j], (Task));
                require(_task.startingBlock_ + TASK_EXPIRY_BLOCKS < block.number || sertnServiceManager.isTaskVerified(_rewardedTasks[i][j]), "Task has not expired");
                require(_inBytesArray(_operator.submittedTasks_, _rewardedTasks[i][j]), "Not in submitted Tasks");
                _operator.submittedTasks_ = _removeBytesElement(_operator.submittedTasks_, _rewardedTasks[i][j]);
                sertnServiceManager.updateOperator(operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator, _operator);
                if (_inBytesArray(_operator.openTasks_, _rewardedTasks[i][j])) {
                    sertnServiceManager.clearTask(_rewardedTasks[i][j]);
                }
            }

            _approvalAmount += operatorDirectedRewardsSubmissions[0].operatorRewards[i].amount;
        }

        ser.approve(address(rewardsCoordinator), _approvalAmount);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(address(sertnServiceManager), operatorDirectedRewardsSubmissions);
    }

    function slashOperator(bytes memory _taskId, string memory _whySlashed) external onlyAggregators() {

        Task memory _task = abi.decode(_taskId, (Task));

        uint8 _modelId = _task.modelId_;

        if (0 > _modelId || sertnServiceManager.getNumModels() < _modelId) {
            revert NotModelId("Not Model Id");
        }
        Model memory _model = sertnServiceManager.getModelInfo(_modelId);

        address[] memory _operator = new address[](1);
        _operator[0] = _model.operator_;

        IStrategy[] memory _strategies = new IStrategy[](_model.ethStrategies_.length + 1);
        for (uint8 i = 0; i < _model.ethStrategies_.length; i++) {
            _strategies[i] = _model.ethStrategies_[i];
        }

        _strategies[_model.ethStrategies_.length] = sertnServiceManager.getTokenToStrategy(address(ser));

        uint256[] memory _wadsToSlash = new uint256[](_strategies.length);

        uint256[] memory _operatorShares = delegationManager.getOperatorShares(_model.operator_, _strategies);

        for (uint8 i = 0; i < _model.ethStrategies_.length; i++) {
            _wadsToSlash[i] = 1 ether * 1 ether * (_model.ethShares_[i]) / (allocationManager.getAllocation(_model.operator_, opSet, _strategies[i]).currentMagnitude * _operatorShares[i]);
        }

        _wadsToSlash[_model.ethStrategies_.length] = (1 ether * 1 ether * 10 * _task.poc_)/(allocationManager.getAllocation(_model.operator_, opSet, _strategies[_model.ethStrategies_.length]).currentMagnitude * _operatorShares[_model.ethStrategies_.length]);

        IAllocationManagerTypes.SlashingParams memory _slashParams = IAllocationManagerTypes.SlashingParams({operator: _model.operator_, operatorSetId: 0, strategies: _strategies, wadsToSlash: _wadsToSlash, description: _whySlashed});

        allocationManager.slashOperator(address(sertnServiceManager), _slashParams);

        emit operatorSlashed(_model.operator_, _taskId);
        // TODO: Custom logic to change bounty amount
        if (sertnServiceManager.getBountyHunter(_taskId) != address(0)) {
            ser.transferFrom(address(this), sertnServiceManager.getBountyHunter(_taskId), BOUNTY);
            Operator memory _operator = sertnServiceManager.getOperatorInfo(_model.operator_);
            _operator.proofRequestExponents_[0] -= 500;
            sertnServiceManager.updateOperator(_model.operator_, _operator);
        }

        sertnServiceManager.clearTask(_taskId);
    }

    function _inBytesArray(bytes[] memory _byteArray, bytes memory _byteElement) internal pure returns(bool) {
        bytes32 elementHash = keccak256(_byteElement);
        for (uint256 i = 0; i < _byteArray.length; i++) {
            if (_byteArray[i].length > 0 && elementHash == keccak256(_byteArray[i])) {
                return true;
            }
        }
        return false;
    }

    function _removeBytesElement(bytes[] memory _byteArray, bytes memory _byteElement) internal pure returns(bytes[] memory) {

        bytes[] memory newArray = new bytes[](_byteArray.length - 1);
        bytes32 elementHash = keccak256(_byteElement);
        bool found = false;
        uint256 writeIndex = 0;

        for (uint256 readIndex = 0; readIndex < _byteArray.length; readIndex++) {
            if (!found && elementHash == keccak256(_byteArray[readIndex])) {
                found = true;
                continue;
            }

            if (writeIndex >= newArray.length) {
                revert("Array length mismatch");
            }

            newArray[writeIndex] = _byteArray[readIndex];
            writeIndex++;
        }


        return newArray;
    }
}
