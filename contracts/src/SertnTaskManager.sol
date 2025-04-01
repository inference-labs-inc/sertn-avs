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
    IStrategy[] public ethStrategies;

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
        address _ser,
        IStrategy[] memory _ethStrategies
    ) OwnableUpgradeable() {
        _transferOwnership(msg.sender);
        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);
        sertnServiceManager = SertnServiceManager(_sertnServiceManager);
        modelStorage = ModelStorage(_modelStorage);




        opSet = OperatorSet({avs: address(sertnServiceManager), id: 0});
        ser = IERC20(_ser);
        ethStrategies = _ethStrategies;

    }

    function sendTask(Task memory _task) external {
        _task.startTime_ = block.timestamp;
        _task.startingBlock_ = uint32(block.number);

        bytes32 _nodeModelId = keccak256(abi.encodePacked(_task.operator_,_task.node_,_task.modelId_));
        bytes memory _emptyBytes;
        if (keccak256(sertnServiceManager.nodeModelInfo(_nodeModelId)) == keccak256(_emptyBytes)) {
            revert NotNodeModel();
        }
        NodeModel memory _nodeModel = abi.decode(sertnServiceManager.nodeModelInfo(_nodeModelId), (NodeModel));
        Operator memory _operator = abi.decode(sertnServiceManager.opInfo(_task.operator_), (Operator));
        // NodeModel memory _nodeModel = sertnServiceManager.getModelInfo(_operatorModelId);
        if (_task.proveOnResponse_ && !_nodeModel.proveOnResponse_) {
            revert NoProofOnResponse();
        }

        if (_task.proveOnResponse_) {
            _checkFinancialSecurity(_task.poc_, _task.operator_, _nodeModel, _operator, _nodeModel.maxBlocks_);

        } else {
            _checkFinancialSecurity(
                _task.poc_ * 10,
                _task.operator_,
                _nodeModel,
                _operator,
                uint32(taskExpiryBlocks)
            );
        }

        if (
            _nodeModel.available_ &&
            _operator.nodeComputeUnits_[_task.node_] > _nodeModel.compute_ &&
            _operator.pausedBlock_ == 0
        ) {
            // uint256 transferAmount = (1.5e3 * (_nodeModel.baseFee_ + _task.poc_)) / 1e3;
            uint256 transferAmount = MathUpgradeable.mulDiv(1.5e9, _nodeModel.baseFee_ + _task.poc_, 1e9);
            if (!ser.transferFrom(msg.sender, address(sertnServiceManager), transferAmount)) {
                revert TransferFailed();
            }
            
            bytes memory _taskId = abi.encode(_task);

            NodeTasks memory _nodeTasks;
            if (keccak256(_emptyBytes) == keccak256(nodeTasks[_nodeModelId])) {
                _nodeTasks.openTasks_ = _pushToBytesArray(_taskId, _nodeTasks.openTasks_);
                _nodeTasks.proofRequests_ = _pushToBytes32Array(keccak256(_taskId), _nodeTasks.proofRequests_);
            } else {
                _nodeTasks = abi.decode(nodeTasks[_nodeModelId],(NodeTasks));
                _nodeTasks.openTasks_ = _pushToBytesArray(_taskId, _nodeTasks.openTasks_);
                _nodeTasks.proofRequests_ = _pushToBytes32Array(keccak256(_taskId), _nodeTasks.proofRequests_);

            }

            _operator.nodeComputeUnits_[_task.node_] -= _nodeModel.compute_;

            for (uint256 i; i < _operator.allocatedEth_.length;) {
                _operator.allocatedEth_[i] += _nodeModel.ethShares_[i];
                unchecked { ++i; }
            }
            _operator.allocatedSer_ += _task.poc_;
            if (!_task.proveOnResponse_) {
                _operator.allocatedSer_ += 9 * _task.poc_;
            }
            sertnServiceManager.updateOperator(_task.operator_, _operator);
            emit NewTask(_task.operator_, _taskId);
        } else {
            revert TaskCouldNotBeSent();
        }
    }

    function _checkFinancialSecurity(
        uint256 _poc,
        address _operatorAddr,
        NodeModel memory _nodeModel,
        Operator memory _operator,
        uint32 _securityDuration
    ) internal view {
        address[] memory _operators = new address[](1);
        _operators[0] = _operatorAddr;
        //check if secure using max blocks as model param
        uint256[] memory _ethSecurity = allocationManager
            .getMinimumSlashableStake(
                opSet,
                _operators,
                ethStrategies,
                uint32(block.number) + _securityDuration
            )[0];

        for (uint256 i; i < _ethSecurity.length;) {
            if (
                _ethSecurity[i] <
                _nodeModel.ethShares_[i] + _operator.allocatedEth_[i]
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
            _nodeModel.maxSer_ < _poc ||
            _serSecurity < _poc + _operator.allocatedSer_
        ) {
            revert TaskCouldNotBeSent();
        }
    }

    function submitTask(TaskResponse memory _taskResponse, bool _verification, bytes memory _proof) external {

        Task memory _task = abi.decode(_taskResponse.taskId_, (Task));
        bytes32 _taskHash = keccak256(_taskResponse.taskId_);

        bytes32 _nodeModelId = keccak256(abi.encodePacked(_task.operator_,_task.node_,_task.modelId_));

        NodeModel memory _nodeModel = abi.decode(sertnServiceManager.nodeModelInfo(_nodeModelId), (NodeModel));

        if (msg.sender != _task.operator_) {
            revert IncorrectOperator();
        }

        Operator memory _operator = abi.decode(sertnServiceManager.opInfo(_task.operator_),(Operator));

        if (_nodeModel.maxBlocks_ < uint32(block.number) - _task.startingBlock_) {
            revert TaskExpired();
        }
        if (_verification) {
            _checkFinancialSecurity(_task.poc_, _task.operator_, _nodeModel, _operator, 0);
        } else {
            _checkFinancialSecurity(10*_task.poc_, _task.operator_, _nodeModel, _operator, taskExpiryBlocks + _task.startingBlock_ - uint32(block.number));
        }

        if (_taskResponse.proven_) {
            if (taskVerified[_taskHash]) {
                _taskResponse.proven_ = true;
                _clearTask(_taskResponse.taskId_, false);
            }

            else {
                sertnServiceManager.pushToOperatorSlashingQueue(msg.sender, _taskHash);
                sertnServiceManager.pushToSlashingQueue(_taskHash);
                emit UpForSlashing(_task.operator_, _taskHash);
                return;
            }
        } else {
            if (_verification || _task.proveOnResponse_) {
            _verifyTask(_taskResponse.taskId_, _task.modelId_, _proof);

            if (taskVerified[_taskHash]) {
                _taskResponse.proven_ = true;
                _clearTask(_taskResponse.taskId_, false);
            }

            else {
                sertnServiceManager.pushToOperatorSlashingQueue(msg.sender, _taskHash);
                sertnServiceManager.pushToSlashingQueue(_taskHash);
                emit UpForSlashing(_task.operator_, _taskHash);
                return;
            }
        }
        }
        _operator.submittedTasks_ = _pushToBytes32Array(_taskHash, _operator.submittedTasks_);

        _operator.nodeComputeUnits_[_task.node_] += _nodeModel.compute_;
        sertnServiceManager.updateOperator(_task.operator_, _operator);
        sertnServiceManager.setTaskResponse(_taskResponse);

        emit TaskResponded(_task.modelId_, _taskHash, _taskResponse);
    }

    function requestProof(bytes memory _taskId) external {
        //add probabilistic case
        if (sertnServiceManager.bountyHunter(keccak256(_taskId)) != address(0)) {
            revert BountyAlreadySet();
        }
        Task memory _task = abi.decode(_taskId, (Task));
        bytes32 _nodeModelId = keccak256(abi.encodePacked(_task.operator_,_task.node_,_task.modelId_));
        Operator memory _operator = abi.decode(sertnServiceManager.opInfo(_task.operator_),(Operator));

        uint256 _amount = PROOF_REQUEST_COST*(_operator.proofRequestCoefficients_[_task.node_]/_operator.proofRequestCoefficients_[_task.node_ + 1]);
        if (!ser.transferFrom(msg.sender, address(sertnServiceManager), _amount)) {
            revert TransferFailed();
        }
        NodeTasks memory _nodeTasks = abi.decode(nodeTasks[_nodeModelId],(NodeTasks));
        _nodeTasks.proofRequests_ = _pushToBytes32Array(keccak256(_taskId), _nodeTasks.proofRequests_);
        nodeTasks[_nodeModelId] = abi.encode(_nodeTasks);

        _operator.proofRequestCoefficients_[_task.node_] += 500;

        sertnServiceManager.updateOperator(_task.operator_, _operator);
        sertnServiceManager.setBountyHunter(keccak256(_taskId), msg.sender);

        emit ProofRequested(_task.operator_, _taskId);

    }

    function _verifyTask(bytes memory _taskId, uint256 _modelId, bytes memory _proof) internal {
        //logic to verify task
        Model memory _model = abi.decode(modelStorage.modelInfo(_modelId), (Model));
        taskVerified[keccak256(_taskId)] = IVerifier(_model.modelVerifier_).verifyProof(_proof, _model.modelKey_);
    }

    function verifyTask(bytes memory _taskId, bytes memory _proof) external onlyOperators() {
        //logic to verify task
        Task memory _task = abi.decode(_taskId, (Task));

        if (_task.operator_ != msg.sender) {
            revert IncorrectOperator();
        }

        _verifyTask(_taskId, _task.modelId_, _proof);
    }

    function verifiedOffChain(bytes memory _taskId, bool _verified) external onlyAggregators() {
        taskVerified[keccak256(_taskId)] = _verified;
    }

    function _clearTask(bytes memory _taskId, bool _slashed) internal {
        Task memory _task = abi.decode(_taskId, (Task));
        bytes32 _taskHash = keccak256(_taskId);

        if (_task.startingBlock_ + taskExpiryBlocks > uint32(block.number) && !taskVerified[_taskHash] && !_slashed) {
            revert TaskNotExpired();
        }

        bytes32 _nodeModelId = keccak256(abi.encodePacked(_task.operator_,_task.node_,_task.modelId_));
        NodeModel memory _nodeModel = abi.decode(sertnServiceManager.nodeModelInfo(_nodeModelId), (NodeModel));
        Operator memory _operator = abi.decode(sertnServiceManager.opInfo(_task.operator_), (Operator));
        NodeTasks memory _nodeTasks = abi.decode(nodeTasks[_nodeModelId],(NodeTasks));

        _nodeTasks.openTasks_ = _removeBytesElement(_nodeTasks.openTasks_, _taskId);
        _nodeTasks.proofRequests_ = _removeBytes32Element( _nodeTasks.proofRequests_, _taskHash);
        if(_slashed) {
            _operator.submittedTasks_ = _removeBytes32Element(_operator.submittedTasks_, _taskHash);
        }
        
        sertnServiceManager.removeFromOperatorSlashingQueue(_task.operator_, _taskHash);
        sertnServiceManager.removeFromSlashingQueue(_taskHash);
        for (uint256 i; i < _operator.allocatedEth_.length;) {
            _operator.allocatedEth_[i] -= _nodeModel.ethShares_[i];
            unchecked { ++i; }
        }

        _operator.allocatedSer_ -= _task.poc_;
        if (!_task.proveOnResponse_) {
            _operator.allocatedSer_ -= 9*_task.poc_;
        }


        sertnServiceManager.updateOperator(_task.operator_, _operator);

    }

    function clearTask(bytes memory _taskId, bool _slashed) external onlyAggregators() {
        _clearTask(_taskId, _slashed);
    }

    function _inBytesArray(bytes[] memory _byteArray, bytes memory _byteElement) internal pure returns(bool) {
        bytes32 elementHash = keccak256(_byteElement);
        for (uint256 i; i < _byteArray.length;) {
            if (elementHash == keccak256(_byteArray[i])) {
                return true;
            }
            unchecked { ++i; }
        }
        return false;
    }

    function _inBytes32Array(bytes32[] memory _byteArray, bytes32 _byteElement) internal pure returns(bool) {
        for (uint256 i; i < _byteArray.length;) {
            if (_byteElement == _byteArray[i]) {
                return true;
            }
            unchecked { ++i; }
        }
        return false;
    }

    function _removeBytesElement(bytes[] memory _byteArray, bytes memory _byteElement) internal pure returns(bytes[] memory) {

        if (_inBytesArray(_byteArray, _byteElement)) {
            bool pastElementIndex;
            bytes[] memory _newArray = new bytes[](_byteArray.length - 1);

            for (uint256 i; i < _byteArray.length -1;) {
                if (keccak256(_byteArray[i]) == keccak256(_byteElement)) {
                    pastElementIndex = true;
                }
                if (pastElementIndex) {
                    _newArray[i] = _byteArray[i+1];
                } else {
                    _newArray[i] = _byteArray[i];
                }
                unchecked { ++i; }
            }
            return _newArray;
        } else {
            return _byteArray;
        }

    }

    function _removeBytes32Element(bytes32[] memory _byteArray, bytes32 _byteElement) internal pure returns(bytes32[] memory) {

        if (_inBytes32Array(_byteArray, _byteElement)) {
            bool pastElementIndex;
            bytes32[] memory _newArray = new bytes32[](_byteArray.length - 1);

            for (uint256 i; i < _byteArray.length -1;) {
                if (_byteArray[i] == _byteElement) {
                    pastElementIndex = true;
                }
                if (pastElementIndex) {
                    _newArray[i] = _byteArray[i+1];
                } else {
                    _newArray[i] = _byteArray[i];
                }
                unchecked { ++i; }
            }
            return _newArray;
        } else {
            return _byteArray;
        }

    }

    function _pushToBytesArray(bytes memory _element, bytes[] memory _array) internal pure returns (bytes[] memory) {
        bytes[] memory _newArray = new bytes[](_array.length + 1);
        for (uint256 i; i < _array.length;) {
            _newArray[i] = _array[i];
            unchecked { ++i; }
        }
        _newArray[_array.length] = _element;
        return _newArray;
    }

    function _pushToBytes32Array(bytes32 _element, bytes32[] memory _array) internal pure returns (bytes32[] memory) {
        bytes32[] memory _newArray = new bytes32[](_array.length + 1);
        for (uint256 i; i < _array.length;) {
            _newArray[i] = _array[i];
            unchecked { ++i; }
        }
        _newArray[_array.length] = _element;
        return _newArray;
    }


}
