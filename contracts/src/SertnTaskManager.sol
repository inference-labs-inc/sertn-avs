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
import "@openzeppelin-upgrades/contracts/utils/math/MathUpgradeable.sol";

import {Test, console2 as console} from "forge-std/Test.sol";

import {SertnServiceManager} from "./SertnServiceManager.sol";
import {ModelStorage} from "./ModelStorage.sol";



contract SertnTaskManager is
    SertnServiceManagerStorage,
    ISertnServiceManagerErrors,
    ISertnServiceManagerEvents,
    OwnableUpgradeable
{
    address[] public aggregators;
    address[] public operators;
    bytes[] public slashingQueue;

    IERC20 public ser;

    IAllocationManager public allocationManager;
    IDelegationManager public delegationManager;
    IRewardsCoordinator public rewardsCoordinator;
    OperatorSet public opSet;

    SertnServiceManager public sertnServiceManager;
    ModelStorage public modelStorage;


    modifier onlyAggregators() {
        if (!sertnServiceManager.isAggregator(msg.sender) && msg.sender != address(sertnServiceManager)) {
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

    constructor(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        address _sertnServiceManager,
        address _modelStorage,
        address _ser
    ) OwnableUpgradeable() {
        _transferOwnership(msg.sender);
        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);
        sertnServiceManager = SertnServiceManager(_sertnServiceManager);
        modelStorage = ModelStorage(_modelStorage);
        



        opSet = OperatorSet({avs: address(sertnServiceManager), id: 0});
        ser = IERC20(_ser);
    }

    function sendTask(Task memory _task) external {

        uint256 _operatorModelId = _task.operatorModelId_;
        _task.startTime_ = block.timestamp;
        _task.startingBlock_ = uint32(block.number);

        if (sertnServiceManager.numOperatorModels() < _operatorModelId) {
            revert NotModelId();
        }
        OperatorModel memory _operatorModel = abi.decode(sertnServiceManager.operatorModelInfo(_operatorModelId), (OperatorModel));
        Operator memory _operator = abi.decode(sertnServiceManager.opInfo(_operatorModel.operator_), (Operator));
        // OperatorModel memory _operatorModel = sertnServiceManager.getModelInfo(_operatorModelId);
        if (_task.proveOnResponse_ && !_operatorModel.proveOnResponse_) {
            revert NoProofOnResponse();
        }

        if (_task.proveOnResponse_) {
            _checkFinancialSecurity(_task.poc_, _operatorModel, _operator, _operatorModel.maxBlocks_);

        } else {
            _checkFinancialSecurity(
                _task.poc_ * 10,
                _operatorModel,
                _operator,
                uint32(TASK_EXPIRY_BLOCKS)
            );
        }
        
        if (
            _operatorModel.available_ &&
            sertnServiceManager.computeUnits(_operatorModel.operator_,_operatorModel.computeType_) > 0 &&
            _operator.pausedBlock_ == 0
        ) {
            // uint256 transferAmount = (1.5e3 * (_operatorModel.baseFee_ + _task.poc_)) / 1e3;
            uint256 transferAmount = MathUpgradeable.mulDiv(1.5e9, _operatorModel.baseFee_ + _task.poc_, 1e9);
            if (!ser.transferFrom(msg.sender, address(sertnServiceManager), transferAmount)) {
                revert TransferFailed();
            }

            bytes memory _taskId = abi.encode(_task);
            _operator.openTasks_ = _pushToBytesArray(_taskId, _operator.openTasks_);
            if (_task.proveOnResponse_) {
                _operator.proofRequests_ = _pushToBytesArray(_taskId, _operator.proofRequests_);
                // _operator.proofRequests_.push(_taskId);
            }

            sertnServiceManager.setComputeUnits(_operatorModel.operator_,_operatorModel.computeType_, false);

            for (uint256 i; i < _operator.allocatedEth_.length;) {
                _operator.allocatedEth_[i] += _operatorModel.ethShares_[i];
                unchecked { ++i; }
            }
            _operator.allocatedSer_ += _task.poc_;
            if (!_task.proveOnResponse_) {
                _operator.allocatedSer_ += 9 * _task.poc_;
            }
            sertnServiceManager.updateOperator(_operatorModel.operator_, _operator);
            emit NewTask(_operatorModel.operator_, _taskId);
        } else {
            revert TaskCouldNotBeSent();
        }
    }

    function _checkFinancialSecurity(
        uint256 _poc,
        OperatorModel memory _operatorModel,
        Operator memory _operator,
        uint32 _securityDuration
    ) internal view {
        address[] memory _operators = new address[](1);
        _operators[0] = _operatorModel.operator_;
        //check if secure using max blocks as model param
        uint256[] memory _ethSecurity = allocationManager
            .getMinimumSlashableStake(
                opSet,
                _operators,
                _operatorModel.ethStrategies_,
                uint32(block.number) + _securityDuration
            )[0];

        for (uint256 i; i < _ethSecurity.length;) {
            if (
                _ethSecurity[i] <
                _operatorModel.ethShares_[i] + _operator.allocatedEth_[i]
            ) {
                revert TaskCouldNotBeSent();
            }
            unchecked { ++i; }
        }
        IStrategy[] memory _pocStrategy = new IStrategy[](1);
        _pocStrategy[0] = sertnServiceManager.tokenToStrategy(address(ser));
        uint256 _serSecurity = allocationManager.getMinimumSlashableStake(
            opSet,
            _operators,
            _pocStrategy,
            uint32(block.number) + _securityDuration
        )[0][0];

        if (
            _operatorModel.maxSer_ < _poc ||
            _serSecurity < _poc + _operator.allocatedSer_
        ) {
            revert TaskCouldNotBeSent();
        }
    }

    function submitTask(TaskResponse memory _taskResponse, bool _verification, bytes memory _proof) external {

        Task memory _task = abi.decode(_taskResponse.taskId_, (Task));

        uint256 _operatorModelId = _task.operatorModelId_;

        if (sertnServiceManager.numOperatorModels() < _operatorModelId) {
            revert NotModelId();
        }
        OperatorModel memory _operatorModel = abi.decode(sertnServiceManager.operatorModelInfo(_operatorModelId), (OperatorModel));
        // OperatorModel memory _operatorModel = operatorModelInfo[_operatorModelId];

        if (msg.sender != _operatorModel.operator_) {
            revert IncorrectOperator();
        }

        Operator memory _operator = abi.decode(sertnServiceManager.opInfo(_operatorModel.operator_),(Operator));

        if (_operatorModel.maxBlocks_ < uint32(block.number) - _task.startingBlock_) {
            revert TaskExpired();
        }
        if (_verification) {
            _checkFinancialSecurity(_task.poc_, _operatorModel, _operator, 0);
        } else {
            _checkFinancialSecurity(10*_task.poc_, _operatorModel, _operator, TASK_EXPIRY_BLOCKS + _task.startingBlock_ - uint32(block.number));
        }

        if (_taskResponse.proven_) {
            if (taskVerified[_taskResponse.taskId_]) {
                _taskResponse.proven_ = true;
                _clearTask(_taskResponse.taskId_, false);
            }

            else {
                sertnServiceManager.pushToOperatorSlashingQueue(msg.sender, _taskResponse.taskId_);
                sertnServiceManager.pushToSlashingQueue(_taskResponse.taskId_);
                emit UpForSlashing(_operatorModel.operator_, _taskResponse.taskId_);
                return;
            }
        } else {
            if (_verification || _task.proveOnResponse_) {
            _verifyTask(_taskResponse.taskId_, _proof);

            if (taskVerified[_taskResponse.taskId_]) {
                _taskResponse.proven_ = true;
                _clearTask(_taskResponse.taskId_, false);
            }

            else {
                sertnServiceManager.pushToOperatorSlashingQueue(msg.sender, _taskResponse.taskId_);
                sertnServiceManager.pushToSlashingQueue(_taskResponse.taskId_);
                emit UpForSlashing(_operatorModel.operator_, _taskResponse.taskId_);
                return;
            }
        }
        }
        _operator.submittedTasks_ = _pushToBytesArray(_taskResponse.taskId_, _operator.submittedTasks_);
        // _operator.submittedTasks_.push(_taskResponse.taskId_);

        sertnServiceManager.updateOperator(_operatorModel.operator_, _operator);
        sertnServiceManager.setTaskResponse(_taskResponse);
        sertnServiceManager.setComputeUnits(_operatorModel.operator_, _operatorModel.computeType_, true);

        emit TaskResponded(_operatorModelId, _taskResponse.taskId_, _taskResponse);
    }

    function requestProof(bytes memory _taskId) external {
        //add probabilistic case
        if (sertnServiceManager.bountyHunter(_taskId) != address(0)) {
            revert BountyAlreadySet();
        }
        Task memory _task = abi.decode(_taskId, (Task));
        OperatorModel memory _operatorModel = abi.decode(sertnServiceManager.operatorModelInfo(_task.operatorModelId_), (OperatorModel));
        // OperatorModel memory _operatorModel = sertnServiceManager.getModelInfo(_task.operatorModelId_);
        Operator memory _operator = abi.decode(sertnServiceManager.opInfo(_operatorModel.operator_),(Operator));

        uint256 _amount = PROOF_REQUEST_COST*(_operator.proofRequestExponents_[0]/_operator.proofRequestExponents_[1]);
        if (!ser.transferFrom(msg.sender, address(sertnServiceManager), _amount)) {
            revert TransferFailed();
        }

        _operator.proofRequests_ = _pushToBytesArray(_taskId, _operator.proofRequests_);

        _operator.proofRequestExponents_[0] += 500;

        sertnServiceManager.updateOperator(_operatorModel.operator_, _operator);
        sertnServiceManager.setBountyHunter(_taskId, msg.sender);

        emit ProofRequested(_operatorModel.operator_, _taskId);

    }

    function _verifyTask(bytes memory _taskId, bytes memory _proof) internal {
        //logic to verify task
        Task memory _task = abi.decode(_taskId, (Task));

        uint256 _operatorModelId = _task.operatorModelId_;

        if (sertnServiceManager.numOperatorModels() < _operatorModelId) {
            revert NotModelId();
        }

        OperatorModel memory _operatorModel = abi.decode(sertnServiceManager.operatorModelInfo(_operatorModelId), (OperatorModel));
        taskVerified[_taskId] = IVerifier(modelStorage.modelAddresses(_operatorModel.modelId_)).verifyProof(_proof);
    }

    function verifyTask(bytes memory _taskId, bytes memory _proof) external onlyOperators() {
        //logic to verify task
        Task memory _task = abi.decode(_taskId, (Task));

        uint256 _operatorModelId = _task.operatorModelId_;

        if (sertnServiceManager.numOperatorModels() < _operatorModelId) {
            revert NotModelId();
        }
        OperatorModel memory _operatorModel = abi.decode(sertnServiceManager.operatorModelInfo(_operatorModelId), (OperatorModel));
        // OperatorModel memory _operatorModel = operatorModelInfo[_operatorModelId];

        if (_operatorModel.operator_ != msg.sender) {
            revert IncorrectOperator();
        }

        taskVerified[_taskId] = IVerifier(modelStorage.modelAddresses(_operatorModel.modelId_)).verifyProof(_proof);
    }

    function verifiedOffChain(bytes memory _taskId, bool _verified) external onlyAggregators() {
        taskVerified[_taskId] = _verified;
    }

    function _clearTask(bytes memory _taskId, bool _slashed) internal {
        Task memory _task = abi.decode(_taskId, (Task));
        if (_task.startingBlock_ + TASK_EXPIRY_BLOCKS > uint32(block.number) && !taskVerified[_taskId] && !_slashed) {
            revert TaskNotExpired();
        }

        uint256 _operatorModelId = _task.operatorModelId_;

        if (sertnServiceManager.numOperatorModels() < _operatorModelId) {
            revert NotModelId();
        }
        OperatorModel memory _operatorModel = abi.decode(sertnServiceManager.operatorModelInfo(_operatorModelId), (OperatorModel));
        Operator memory _operator = abi.decode(sertnServiceManager.opInfo(_operatorModel.operator_), (Operator));
        _operator.openTasks_ = _removeBytesElement(_operator.openTasks_, _taskId);
        _operator.proofRequests_ = _removeBytesElement(_operator.proofRequests_, _taskId);
        if(_slashed) {
            _operator.submittedTasks_ = _removeBytesElement(_operator.submittedTasks_, _taskId);
        }
        operatorSlashingQueue[_operatorModel.operator_] = _removeBytesElement(operatorSlashingQueue[_operatorModel.operator_], _taskId);
        slashingQueue = _removeBytesElement(slashingQueue, _taskId);
        for (uint256 i; i < _operator.allocatedEth_.length;) {
            _operator.allocatedEth_[i] -= _operatorModel.ethShares_[i];
            unchecked { ++i; }
        }

        _operator.allocatedSer_ -= _task.poc_;
        if (!_task.proveOnResponse_) {
            _operator.allocatedSer_ -= 9*_task.poc_;
        }
        

        sertnServiceManager.updateOperator(_operatorModel.operator_, _operator);

    }

    function clearTask(bytes memory _taskId, bool _slashed) external onlyAggregators() {
        _clearTask(_taskId, _slashed);
    }

    function _inBytesArray(bytes[] memory _byteArray, bytes memory _byteElement) internal pure returns(bool) {
        bytes32 elementHash = keccak256(_byteElement);
        for (uint256 i; i < _byteArray.length;) {
            if (_byteArray[i].length > 0 && elementHash == keccak256(_byteArray[i])) {
                return true;
            }
            unchecked { ++i; }
        }
        return false;
    }

    function _removeBytesElement(bytes[] memory _byteArray, bytes memory _byteElement) internal pure returns(bytes[] memory) {
        if (_inBytesArray(_byteArray, _byteElement)) {
            bytes[] memory newArray;
            if (_byteArray.length ==1) {
                return newArray;
            }
            newArray = new bytes[](_byteArray.length - 1);
            bytes32 elementHash = keccak256(_byteElement);
            bool found = false;
            uint256 writeIndex = 0;

            for (uint256 readIndex; readIndex < _byteArray.length;) {
                if (!found && elementHash == keccak256(_byteArray[readIndex])) {
                    found = true;
                    continue;
                }
                if (writeIndex >= newArray.length) {
                    revert("Array length mismatch");
                }

                newArray[writeIndex] = _byteArray[readIndex];
                writeIndex++;
                unchecked { ++readIndex; }
            }


            return newArray;
        } else {
            return _byteArray;
        }
        
    }
    function _pushToBytesArray(bytes memory _element, bytes[] memory _array) internal returns (bytes[] memory) {
        bytes[] memory _newArray = new bytes[](_array.length + 1);
        for (uint256 i; i < _array.length;) {
            _newArray[i] = _array[i];
            unchecked { ++i; }
        }
        _newArray[_array.length] = _element;
        return _newArray;
    }

    
}
