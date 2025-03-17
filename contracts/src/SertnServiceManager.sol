// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {ECDSAServiceManagerBase} from
    "@eigenlayer-middleware/src/unaudited/ECDSAServiceManagerBase.sol";
import {ECDSAStakeRegistry} from "@eigenlayer-middleware/src/unaudited/ECDSAStakeRegistry.sol";
import {IServiceManager} from "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {ECDSAUpgradeable} from
    "@openzeppelin-upgrades/contracts/utils/cryptography/ECDSAUpgradeable.sol";
import {IERC1271Upgradeable} from
    "@openzeppelin-upgrades/contracts/interfaces/IERC1271Upgradeable.sol";
import "./ISertnServiceManager.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {SertnServiceManagerStorage} from "./SertnServiceManagerStorage.sol";

import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";

import {IAllocationManager} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IERC20} from "@openzeppelin/contracts/interfaces/IERC20.sol";

import {IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import "@eigenlayer/contracts/libraries/OperatorSetLib.sol";

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

/**
 * @title Primary entrypoint for procuring services from Sertn.
 * @author Inference Labs, Inc.
 */
contract SertnServiceManager is
    ISertnServiceManager,
    SertnServiceManagerStorage,
    IAVSRegistrar,
    ISertnServiceManagerErrors,
    ISertnServiceManagerEvents,
    OwnableUpgradeable
{
    uint8 numModels = 0;
    address[] aggregators;
    address[] operators;

    IERC20 ser;

    IAllocationManager allocationManager;
    IDelegationManager delegationManager;
    OperatorSet opSet;

    modifier onlyAggregators() {
        require(isAggregator[msg.sender]);
        _;
    }

    modifier onlyOperators() {
        require(isOperator[msg.sender]);
        _;
    }

    constructor(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) OwnableUpgradeable() {
        isAggregator[msg.sender] = true;

        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);

        _registerToEigen(_strategies, _avsMetadata);

        opSet = OperatorSet({avs: address(this), id: 0});
    }

    function _registerToEigen(
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) internal {
        allocationManager.updateAVSMetadataURI(address(this), _avsMetadata);
        // allocationManager.setAVSRegistrar(address(this), address(this));
        IAllocationManagerTypes.CreateSetParams[]
            memory setParams = new IAllocationManagerTypes.CreateSetParams[](1);
        setParams[0] = IAllocationManagerTypes.CreateSetParams({
            operatorSetId: 0,
            strategies: _strategies
        });
        allocationManager.createOperatorSets(address(this), setParams);
        _addStrategies(_strategies);
    }

    function _addStrategies(
        IStrategy[] memory _strategies
    ) internal {
        for (uint8 i = 0; i < _strategies.length; i++) {
            tokenToStrategy[
                address(_strategies[i].underlyingToken())
            ] = _strategies[i];
        }
        allocationManager.addStrategiesToOperatorSet(
            address(this),
            0,
            _strategies
        );
    }

    function addAggregator(address _aggregator) external onlyAggregators {
        aggregators.push(_aggregator);
        isAggregator[_aggregator] = true;
    }

    function supportsAVS(address avs) external view override returns (bool) {
        require(avs == address(this), WrongAVS());
        return true;
    }

    function registerOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds,
        bytes calldata data
    ) external {
        (
            Model[] memory _models,
            bytes32[] memory _computeUnitNames,
            uint8[] memory _computeUnits
        ) = abi.decode(data, (Model[], bytes32[], uint8[]));
        uint8[] memory _modelIds = new uint8[](_models.length);
        for (uint8 i = 0; i < _models.length; i++) {
            numModels++;
            uint8 modelNum = numModels;
            _modelIds[i] = modelNum;
            _models[i].operator_ = operator;
            modelInfo[modelNum] = _models[i];
            //Allows users to access similar model names, for instance "llama" could correspond to model id's 1, 3, and 7
            modelsByName[_models[i].modelName_].push(modelNum);
            
        }

        for (uint8 i = 0; i < _computeUnitNames.length; i++) {
            computeUnits[operator][_computeUnitNames[i]] = _computeUnits[i];
        }

        opInfo[operator] = Operator({
            models_: _modelIds,
            computeUnits_: _computeUnitNames
        });

        bytes[] memory tempBytesArr;
        // Note: open tasks also denotes tasks which have been completed but are still slashable

        openTasks[operator] = tempBytesArr;
        submittedTasks[operator] = tempBytesArr;
        slashingQueue[operator] = tempBytesArr;
        proofRequests[operator] = tempBytesArr;

        allocatedEth[operator] = new uint256[](_models[0].ethShares_.length);
        for (uint8 i = 0; i < allocatedEth[operator].length; i++) {
            allocatedEth[operator][i] = 0;
        }
        
        allocatedSer[operator] = 0;

        operators.push(operator);
        isOperator[operator] = true;

        // allocationManager.getAllocatedStrategies(operator, )

        emit newOperator(operator);
        emit newModels(_modelIds);
    }

    function sendTask(Task memory _task) external {
        uint8 _modelId = _task.modelId_;
        _task.startTime_ = block.timestamp;
        _task.startingBlock_ = block.number;
        

        if (0 > _modelId || numModels < _modelId) {
            revert NotModelId("Not Model Id");
        }

        Model memory _model = modelInfo[_modelId];

        _checkFinancialSecurity(_task.poc_, _model, _model.maxBlocks_);

        if (_task.proveOnResponse_ && !_model.proveOnResponse_) {
            revert NoProofOnResponse("Prove On Response Not Available");
        }

        if (
            _model.available_ &&
            (computeUnits[_model.operator_][_model.computeType_] > 0)
        ) {
            bytes memory _taskId = abi.encode(_task);
            openTasks[_model.operator_] = _pushToByteArray(_taskId, openTasks[_model.operator_]);
            if (_task.proveOnResponse_) {
                proofRequests[_model.operator_] = _pushToByteArray(_taskId, proofRequests[_model.operator_]);
            }

            computeUnits[_model.operator_][_model.computeType_] -= 1;
            
            for (uint8 i = 0; i < allocatedEth[_model.operator_].length; i++) {
                allocatedEth[_model.operator_][i] += _model.ethShares_[i];
            }

            allocatedSer[_model.operator_] += _task.poc_;

            emit newTask(_model.operator_, _taskId);
        } else {
            revert TaskCouldNotBeSent("Task Could Not Be Sent");
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
        for (uint8 i = 0; i < _ethSecurity.length; i++) {
            if (_ethSecurity[i] > _model.ethShares_[i] + allocatedEth[_model.operator_][i]) {
                revert TaskCouldNotBeSent("Not enough eth backed security");
            }
        }
        IStrategy[] memory _pocStrategy = new IStrategy[](1);
        _pocStrategy[0] = tokenToStrategy[address(ser)];
        if (
            _model.maxSer_ > _poc &&
            allocationManager.getMinimumSlashableStake(
                opSet,
                _operators,
                _pocStrategy,
                uint32(block.number) + _securityDuration
            )[0][0] >
            _poc + allocatedSer[_model.operator_]
        ) {
            revert TaskCouldNotBeSent("Not enough ser backed security");
        }
    }

    function _verifyTask(bytes memory _taskId, bytes memory _proof) internal {
        //logic to verify task
        taskVerified[_taskId] = true;
    }

    function submitTask(TaskResponse memory _taskResponse, bool _verification, bytes memory _proof) external {
        
        Task memory _task = abi.decode(_taskResponse.taskId_, (Task));

        uint8 _modelId = _task.modelId_;

        if (0 > _modelId || numModels < _modelId) {
            revert NotModelId("Not Model Id");
        }

        Model memory _model = modelInfo[_modelId];

        require(msg.sender == _model.operator_, "Not operator assigned to task");

        _checkFinancialSecurity(_task.poc_, _model, 0);

        if (_verification || _task.proveOnResponse_) {
            _verifyTask(_taskResponse.taskId_, _proof);

            if (taskVerified[_taskResponse.taskId_]) {
                bytes[] memory _taskId = openTasks[msg.sender];
                for (uint8 i = 0; i < openTasks[msg.sender].length; i++) {
                    if (_taskResponse.taskId_.length != _taskId[i].length) {
                        continue;
                    }
                    for (uint8 j = 0; j < openTasks[msg.sender].length; j++) {
                        if (_taskResponse.taskId_[j] != _taskId[i][j]) {
                            break;
                        }
                    }
                    delete openTasks[msg.sender][i];
                    break;
                    }
                
                _taskId = proofRequests[msg.sender];
                for (uint8 i = 0; i < proofRequests[msg.sender].length; i++) {
                    if (_taskResponse.taskId_.length != _taskId[i].length) {
                        continue;
                    }
                    for (uint8 j = 0; j < proofRequests[msg.sender].length; j++) {
                        if (_taskResponse.taskId_[j] != _taskId[i][j]) {
                            break;
                        }
                    }
                    delete proofRequests[msg.sender][i];
                    break;
                    }
            }

            else {
                slashingQueue[msg.sender] = _pushToByteArray(_taskResponse.taskId_, slashingQueue[msg.sender]);
                emit upForSlashing(_model.operator_, _taskResponse.taskId_);
            }

        }
        
        
        submittedTasks[msg.sender] = _pushToByteArray(_taskResponse.taskId_, submittedTasks[msg.sender]);

        taskResponse[msg.sender][_taskResponse.taskId_] = _taskResponse;

        emit taskResponded(_modelId, _taskResponse.taskId_, _taskResponse);
    }

    function requestProof(bytes memory _taskId) external {
        
        Task memory _task = abi.decode(_taskId, (Task));
        Model memory _model = modelInfo[_task.modelId_];


        uint256 _amount = 
        ser.transferFrom(msg.sender, address(this), _amount);
    }

    function _pushToByteArray(bytes memory _element, bytes[] memory _arr) internal returns (bytes[] memory){
        bytes[] memory _tempBytesArr = new bytes[](_arr.length + 1);
        for (uint8 i = 0; i < _arr.length; i++) {
            _tempBytesArr[i] = _arr[i];
        }
        _tempBytesArr[_arr.length] = _element;
        return _tempBytesArr;
    }

    function deregisterOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds
    ) external override {}
}
