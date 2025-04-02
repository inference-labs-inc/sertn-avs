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
        bytes32 _nodeOperatorId = keccak256(abi.encode(_task.operator_,_task.node_));
        bytes32 _nodeModelId = keccak256(abi.encode(_task.operator_,_task.node_,_task.modelId_));
        if (maxBlocks[_nodeModelId] == 0) {
            revert NotNodeModel();
        }
        // NodeModel memory _nodeModel = sertnServiceManager.getModelInfo(_operatorModelId);
        if (_task.proveOnResponse_ && !proveOnResponse[_nodeModelId]) {
            revert NoProofOnResponse();
        }

        if (_task.proveOnResponse_) {
            _checkFinancialSecurity(_task.poc_, _task.operator_, _nodeModelId, maxBlocks[_nodeModelId]);

        } else {
            _checkFinancialSecurity(
                _task.poc_ * 10,
                _task.operator_,
                _nodeModelId,
                uint32(taskExpiryBlocks)
            );
        }

        if (
            nodeModelAvailable[_nodeModelId] &&
            nodeComputeUnits[_nodeOperatorId] > compute[_nodeModelId] &&
            operatorPausedBlock[_task.operator_] == 0
        ) {
            // uint256 transferAmount = (1.5e3 * (sertnServiceManager.baseFee_ + _task.poc_)) / 1e3;
            uint256 transferAmount = MathUpgradeable.mulDiv(1.5e9, baseFee[_nodeModelId] + _task.poc_, 1e9);
            if (!ser.transferFrom(msg.sender, address(sertnServiceManager), transferAmount)) {
                revert TransferFailed();
            }


            
            bytes32 _taskId = keccak256(abi.encode(_task));
            openTasks[_nodeOperatorId] = _pushToBytes32Array(_taskId, openTasks[_nodeOperatorId]);
            isOpenTask[_nodeOperatorId] = true;

            if (_task.proveOnResponse_) {
                proofRequests[_nodeOperatorId] = _pushToBytes32Array(_taskId, proofRequests[_nodeOperatorId]);
                proofRequested[_nodeOperatorId] = true;
            }

            nodeComputeUnits[_nodeOperatorId] -= compute[_nodeModelId];

            for (uint256 i = 1; i < allocatedStake[_task.operator_].length;) {
                allocatedStake[_task.operator_][i] += ethShares[_nodeModelId][i];
                slashableAmount[_taskId][i] += ethShares[_nodeModelId][i];
                unchecked { ++i; }
            }
            allocatedStake[_task.operator_][0] += _task.poc_;
            slashableAmount[_taskId] += _task.poc_;
            taskPoc[_taskId] = _task.poc_;
            if (!_task.proveOnResponse_) {
                allocatedStake[_task.operator_][0] += 9*_task.poc_;
                slashableAmount[_taskId] += 9*_task.poc_;
                taskPoc[_taskId] = 10*_task.poc_;
            }
            taskOperator[_taskId] = _task.operator_;
            taskNodeOperatorId[_taskId] = _nodeOperatorId;
            taskNodeModelId[_taskId] = _nodeModelId;
            startingBlock[_taskId] = _task.startingBlock_;
            

            emit NewTask(_nodeModelId, _task);
        } else {
            revert TaskCouldNotBeSent();
        }
    }

    function _checkFinancialSecurity(
        uint256 _poc,
        address _operator,
        bytes32 _nodeModelId,
        uint32 _securityDuration
    ) internal view {
        address[] memory _operators = new address[](1);
        _operators[0] = _operator;
        //check if secure using max blocks as model param

        IStrategy[] memory _pocStrategy = new IStrategy[](1);
        _pocStrategy[0] = sertnServiceManager.tokenToStrategy(address(ser));
        uint256 _serSecurity = allocationManager.getMinimumSlashableStake(
            opSet,
            _operators,
            _pocStrategy,
            uint32(block.number) + _securityDuration
        )[0][0];

        if (
            maxSer[_nodeModelId] < _poc ||
            _serSecurity < _poc + allocatedStake[_operator][0]
        ) {
            revert TaskCouldNotBeSent();
        }

        uint256[] memory _ethSecurity = allocationManager
            .getMinimumSlashableStake(
                opSet,
                _operators,
                ethStrategies,
                uint32(block.number) + _securityDuration
            )[0];

        for (uint256 i = 1; i < _ethSecurity.length + 1;) {
            if (
                _ethSecurity[i] <
                ethShares[_nodeModelId][i] + allocatedStake[_operator][i]
            ) {
                revert TaskCouldNotBeSent();
            }
            unchecked { ++i; }
        }
    }

    function submitTask(TaskResponse memory _taskResponse, bool _verification, bytes memory _proof) external {

        bytes32 _taskId = keccak256(abi.encode(_taskResponse.task_));

        if (!isOpenTask[_taskId]) {
            revert NotOpenTask();
        }

        bytes32 _nodeModelId = keccak256(abi.encodePacked(_taskResponse.task_.operator_,_taskResponse.task_.node_,_taskResponse.task_.modelId_));

        if (msg.sender != _taskResponse.task_.operator_) {
            revert IncorrectOperator();
        }

        if (maxBlocks[_nodeModelId] < uint32(block.number) - _taskResponse.task_.startingBlock_) {
            revert TaskExpired();
        }
        if (_verification) {
            _checkFinancialSecurity(_taskResponse.task_.poc_, _taskResponse.task_.operator_, _nodeModelId, 0);
        } else {
            _checkFinancialSecurity(10*_taskResponse.task_.poc_, _taskResponse.task_.operator_, _nodeModelId, taskExpiryBlocks + _taskResponse.task_.startingBlock_ - uint32(block.number));
        }

        if (_taskResponse.proven_) {
            if (taskVerified[_taskId]) {
                _taskResponse.proven_ = true;
                _clearTask(_taskResponse.taskId_, false);
            }

            else {
                sertnServiceManager.pushToOperatorSlashingQueue(msg.sender, _taskId);
                sertnServiceManager.pushToSlashingQueue(_taskId);
                emit UpForSlashing(_taskResponse.task_.operator_, _taskId);
                return;
            }
        } else {
            if (_verification || _taskResponse.task_.proveOnResponse_) {
            _verifyTask(_taskResponse.taskId_, _taskResponse.task_.modelId_, _proof);

            if (taskVerified[_taskId]) {
                _taskResponse.proven_ = true;
                _clearTask(_taskResponse.taskId_, false);
            }

            else {
                sertnServiceManager.pushToOperatorSlashingQueue(msg.sender, _taskId);
                sertnServiceManager.pushToSlashingQueue(_taskId);
                emit UpForSlashing(_taskResponse.task_.operator_, _taskId);
                return;
            }
        }
        }
        taskSubmitted[_taskId] = true;

        nodeComputeUnits[_nodeModelId] += compute[_nodeModelId];
        rewardsAmount[_taskResponse.task_.operator_] += MathUpgradeable.mulDiv(_taskResponse.task_.poc_ + baseFee[_nodeModelId], 7e9, 1e10);
        slashableAmount[_taskId][0] += MathUpgradeable.mulDiv(_taskResponse.task_.poc_ + baseFee[_nodeModelId], 7e9, 1e10);

        emit TaskResponded(_taskId, _taskResponse);
    }

    function requestProof(bytes32 memory _taskId) external {
        //add probabilistic case
        if (bountyHunter[_taskId] != address(0)) {
            revert BountyAlreadySet();
        }

        uint256 _amount = MathUpgradeable.mulDiv(PROOF_REQUEST_COST, proofRequestCoefficients[taskOperator[_taskId]][0], proofRequestCoefficients[taskOperator[_taskId]][1]);
        if (!ser.transferFrom(msg.sender, address(sertnServiceManager), _amount)) {
            revert TransferFailed();
        }

        proofRequests[taskNodeOperatorId[_taskId]] = _pushToBytes32Array(_taskId, proofRequests[taskNodeOperatorId[_taskId]]);
        proofRequested[taskNodeOperatorId[_taskId]] = true;

        proofRequestCoefficients[taskOperator[_taskId]][0] += proofRequestIncrement;

        bountyHunter[_taskId] = msg.sender;

        emit ProofRequested(_taskId);

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

    function _clearTask(bytes32 memory _taskId, bool _slashed) internal {

        if (startingBlock[_taskId] + taskExpiryBlocks > uint32(block.number) && !taskVerified[_taskId] && !_slashed) {
            revert TaskNotExpired();
        }

        openTasks[taskNodeOperatorId[_taskId]] = _removeBytes32Element(openTasks[taskNodeOperatorId[_taskId]], _taskId);
        delete isOpenTask[_taskId];
        proofRequests[taskNodeOperatorId[_taskId]] = _removeBytes32Element(proofRequests[taskNodeOperatorId[_taskId]], _taskId);
        delete proofRequested[_taskId];
        if(_slashed) {
            delete taskSubmitted[_taskId];
        }
        
        sertnServiceManager.removeFromOperatorSlashingQueue(taskOperator[_taskId], _taskId);
        sertnServiceManager.removeFromSlashingQueue(_taskId);
        for (uint256 i = 1; i < allocatedStake[taskOperator[_taskId]].length+1;) {
            allocatedStake[taskOperator[_taskId]][i] -= slashableAmount[_taskId][i];
            unchecked { ++i; }
        }

        allocatedStake[taskOperator[_taskId]][0] -= taskPoc[_taskId];

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
