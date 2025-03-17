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
import {IVerifier} from "./IVerifier.sol";
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
            modelsByName[_models[i].modelName_] = _pushToUint8Array(modelNum, modelsByName[_models[i].modelName_]);
            
        }

        for (uint8 i = 0; i < _computeUnitNames.length; i++) {
            computeUnits[operator][_computeUnitNames[i]] = _computeUnits[i];
        }

        opInfo[operator] = Operator({
            models_: _modelIds,
            computeUnits_: _computeUnitNames
        });

        allocatedEth[operator] = new uint256[](_models[0].ethShares_.length);
        for (uint8 i = 0; i < allocatedEth[operator].length; i++) {
            allocatedEth[operator][i] = 0;
        }
        
        allocatedSer[operator] = 0;

        operators.push(operator);
        isOperator[operator] = true;
        uint32[2] memory tempProofRequestExponents = [uint32(1e3),uint32(1e3)];
        proofRequestExponents[operator] = tempProofRequestExponents;

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

        _checkFinancialSecurity(_task.poc_*10, _model, _model.maxBlocks_);

        if (_task.proveOnResponse_ && !_model.proveOnResponse_) {
            revert NoProofOnResponse("Prove On Response Not Available");
        }

        if (
            _model.available_ &&
            (computeUnits[_model.operator_][_model.computeType_] > 0 && 
            //payment, should probably implement rounding
            ser.transferFrom(msg.sender, address(this), 1.5e3*(_model.baseFee_ + _task.poc_)/1e3))
        ) {
            bytes memory _taskId = abi.encode(_task);
            // Note: open tasks also denotes tasks which have been completed but are still slashable
            openTasks[_model.operator_] = _pushToByteArray(_taskId, openTasks[_model.operator_]);
            if (_task.proveOnResponse_) {
                proofRequests[_model.operator_] = _pushToByteArray(_taskId, proofRequests[_model.operator_]);
            }

            computeUnits[_model.operator_][_model.computeType_] -= 1;
            
            for (uint8 i = 0; i < allocatedEth[_model.operator_].length; i++) {
                allocatedEth[_model.operator_][i] += _model.ethShares_[i];
            }

            allocatedSer[_model.operator_] += 10*_task.poc_;

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
        Task memory _task = abi.decode(_taskId, (Task));

        uint8 _modelId = _task.modelId_;

        if (0 > _modelId || numModels < _modelId) {
            revert NotModelId("Not Model Id");
        }

        Model memory _model = modelInfo[_modelId];

        taskVerified[_taskId] = IVerifier(_model.verifier_).verifyProof(_proof);
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
                _clearTask(_taskResponse.taskId_, false);
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

        require(bountyHunter[_taskId] == address(0), "bounty already set");
        Task memory _task = abi.decode(_taskId, (Task));
        Model memory _model = modelInfo[_task.modelId_];


        uint256 _amount = PROOF_REQUEST_COST*(proofRequestExponents[_model.operator_][0]/proofRequestExponents[_model.operator_][1]);
        ser.transferFrom(msg.sender, address(this), _amount);

        proofRequests[_model.operator_] = _pushToByteArray(_taskId, proofRequests[_model.operator_]);

        proofRequestExponents[_model.operator_][0] += 500;

        bountyHunter[_taskId] = msg.sender;

        emit proofRequested(_model.operator_, _taskId);

    }

    function clearTask(bytes memory _taskId) external onlyAggregators() {
        _clearTask(_taskId, false);
    }

    function _clearTask(bytes memory _taskId, bool _slashed) internal {
        Task memory _task = abi.decode(_taskId, (Task));
        require(_task.startingBlock_ + TASK_EXPIRY_BLOCKS < block.number || taskVerified[_taskId] || _slashed, "Task has not expired");

        uint8 _modelId = _task.modelId_;

        if (0 > _modelId || numModels < _modelId) {
            revert NotModelId("Not Model Id");
        }

        Model memory _model = modelInfo[_modelId];

        bytes[] memory _taskArr = openTasks[_model.operator_];
        for (uint8 i = 0; i < _taskArr.length; i++) {
            if (_taskId.length != _taskArr[i].length) {
                continue;
            }
            for (uint8 j = 0; j < _taskId.length; j++) {
                if (_taskId[j] != _taskArr[i][j]) {
                    break;
                }
            }
            delete openTasks[_model.operator_][i];
            break;
            }
        
        _taskArr = proofRequests[_model.operator_];
        for (uint8 i = 0; i < _taskArr.length; i++) {
            if (_taskId.length != _taskArr[i].length) {
                continue;
            }
            for (uint8 j = 0; j < _taskId.length; j++) {
                if (_taskId[j] != _taskArr[i][j]) {
                    break;
                }
            }
            delete proofRequests[_model.operator_][i];
            break;
            }
        _taskArr = submittedTasks[_model.operator_];
        for (uint8 i = 0; i < _taskArr.length; i++) {
            if (_taskId.length != _taskArr[i].length) {
                continue;
            }
            for (uint8 j = 0; j < _taskId.length; j++) {
                if (_taskId[j] != _taskArr[i][j]) {
                    break;
                }
            }
            delete submittedTasks[_model.operator_][i];
            break;
            }
        
        _taskArr = slashingQueue[_model.operator_];
        for (uint8 i = 0; i < _taskArr.length; i++) {
            if (_taskId.length != _taskArr[i].length) {
                continue;
            }
            for (uint8 j = 0; j < _taskId.length; j++) {
                if (_taskId[j] != _taskArr[i][j]) {
                    break;
                }
            }
            delete slashingQueue[_model.operator_][i];
            break;
            }

        for (uint8 i = 0; i < allocatedEth[_model.operator_].length; i++) {
                allocatedEth[_model.operator_][i] -= _model.ethShares_[i];
            }
        
        allocatedSer[_model.operator_] -= 10*_task.poc_;

        

    }

    function _slashOperator(bytes memory _taskId, string memory _whySlashed) external onlyAggregators() {

        Task memory _task = abi.decode(_taskId, (Task));

        uint8 _modelId = _task.modelId_;

        if (0 > _modelId || numModels < _modelId) {
            revert NotModelId("Not Model Id");
        }

        Model memory _model = modelInfo[_modelId];

        address[] memory _operator = new address[](1);
        _operator[0] = _model.operator_;

        IStrategy[] memory _strategies = new IStrategy[](_model.ethStrategies_.length + 1);
        for (uint8 i = 0; i < _model.ethStrategies_.length; i++) {
            _strategies[i] = _model.ethStrategies_[i];
        }

        _strategies[_model.ethStrategies_.length] = tokenToStrategy[address(ser)];

        uint256[] memory _wadsToSlash = new uint256[](_strategies.length);

        uint256[] memory _operatorShares = delegationManager.getOperatorShares(_model.operator_, _strategies);
        
        for (uint8 i = 0; i < _model.ethStrategies_.length; i++) {
            _wadsToSlash[i] = 1 ether * (_model.ethShares_[i]) / (allocationManager.getAllocation(_model.operator_, opSet, _strategies[i]).currentMagnitude * _operatorShares[i]);
        }

        _wadsToSlash[_model.ethStrategies_.length] = (1 ether * 10 * _task.poc_)/(allocationManager.getAllocation(_model.operator_, opSet, _strategies[_model.ethStrategies_.length]).currentMagnitude * _operatorShares[_model.ethStrategies_.length]);

        IAllocationManagerTypes.SlashingParams memory _slashParams = IAllocationManagerTypes.SlashingParams({operator: _model.operator_, operatorSetId: 0, strategies: _strategies, wadsToSlash: _wadsToSlash, description: _whySlashed});

        allocationManager.slashOperator(address(this), _slashParams);

        emit operatorSlashed(_model.operator_, _taskId);

        _clearTask(_taskId, true);
    }

    function _pushToByteArray(bytes memory _element, bytes[] memory _arr) internal returns (bytes[] memory){
        bytes[] memory _tempBytesArr = new bytes[](_arr.length + 1);
        for (uint8 i = 0; i < _arr.length; i++) {
            _tempBytesArr[i] = _arr[i];
        }
        _tempBytesArr[_arr.length] = _element;
        return _tempBytesArr;
    }

    function _pushToUint8Array(uint8 _element, uint8[] memory _arr) internal returns (uint8[] memory){
        uint8[] memory _tempArr = new uint8[](_arr.length + 1);
        for (uint8 i = 0; i < _arr.length; i++) {
            _tempArr[i] = _arr[i];
        }
        _tempArr[_arr.length] = _element;
        return _tempArr;
    }

    function deregisterOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds
    ) external override {}
}
