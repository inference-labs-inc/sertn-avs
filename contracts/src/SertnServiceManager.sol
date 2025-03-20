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
    bytes[] slashingQueue;

    IERC20 ser;

    IAllocationManager allocationManager;
    IDelegationManager delegationManager;
    IRewardsCoordinator rewardsCoordinator;
    OperatorSet opSet;

    modifier onlyAggregators() {
        require(isAggregator[msg.sender], "Not registered aggregator");
        _;
    }

    modifier onlyAggregatorsOrOperator(address _operator) {
        require(isAggregator[msg.sender] || _operator == msg.sender, "Not registered aggregator or concerned operator");
        _;
    }

    modifier onlyOperators() {
        require(isOperator[msg.sender], "Not registered operator");
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
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);

        _registerToEigen(_strategies, _avsMetadata);

        opSet = OperatorSet({avs: address(this), id: 0});
        ser = _strategies[0].underlyingToken();
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
        _addStrategies(_strategies, false);
    }

    function _addStrategies(
        IStrategy[] memory _strategies, bool _newStrategies
    ) internal {
        for (uint8 i = 0; i < _strategies.length; i++) {
            tokenToStrategy[
                address(_strategies[i].underlyingToken())
            ] = _strategies[i];
        }
        if (_newStrategies) {
            allocationManager.addStrategiesToOperatorSet(
            address(this),
            0,
            _strategies
        );
        }
    }
    function addStrategies(
        IStrategy[] memory _strategies, bool _newStrategies
    ) external onlyAggregators() {
        _addStrategies(_strategies, _newStrategies);
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
        require(avs==address(this), "wrong AVS");
        require(operatorSetIds.length == 1 && operatorSetIds[0] == 0, "wrong operatorSetIds");
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
            computeUnits_: _computeUnitNames,
            openTasks_: new bytes[](0),
            submittedTasks_: new bytes[](0),
            proofRequests_: new bytes[](0),
            allocatedEth_: new uint256[](_models[0].ethShares_.length),
            allocatedSer_: 0,
            proofRequestExponents_: [uint32(1e3),uint32(1e3)],
            pausedBlock_: 0
        });

        for (uint8 i = 0; i < _models[0].ethShares_.length; i++) {
            opInfo[operator].allocatedEth_[i] = 0;
        }

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

        if (_task.proveOnResponse_ && !_model.proveOnResponse_) {
            revert NoProofOnResponse("Prove On Response Not Available");
        }

        if (_task.proveOnResponse_) { 

            _checkFinancialSecurity(_task.poc_, _model, _model.maxBlocks_);

        } else {

             _checkFinancialSecurity(_task.poc_*10, _model, _model.maxBlocks_);

        }

        Operator memory _operator = opInfo[_model.operator_];
       
        if (
            _model.available_ &&
            computeUnits[_model.operator_][_model.computeType_] > 0 && 
            //payment, should probably implement rounding
            ser.transferFrom(msg.sender, address(this), 1.5e3*(_model.baseFee_ + _task.poc_)/1e3)
            && _operator.pausedBlock_ == 0
        ) {
            bytes memory _taskId = abi.encode(_task);
            // Note: open tasks also denotes tasks which have been completed but are still slashable
            _operator.openTasks_ = _pushToByteArray(_taskId, _operator.openTasks_);
            // openTasks[_model.operator_] = _pushToByteArray(_taskId, openTasks[_model.operator_]);
            if (_task.proveOnResponse_) {
                _operator.proofRequests_ = _pushToByteArray(_taskId, _operator.proofRequests_);
                // proofRequests[_model.operator_] = _pushToByteArray(_taskId, proofRequests[_model.operator_]);
            }


            computeUnits[_model.operator_][_model.computeType_] -= 1;
            
            for (uint8 i = 0; i < _operator.allocatedEth_.length; i++) {
                _operator.allocatedEth_[i] += _model.ethShares_[i];
            }
            // for (uint8 i = 0; i < allocatedEth[_model.operator_].length; i++) {
            //     allocatedEth[_model.operator_][i] += _model.ethShares_[i];
            // }
            _operator.allocatedSer_ += 10*_task.poc_;
            // allocatedSer[_model.operator_] += 10*_task.poc_;
            opInfo[_model.operator_] = _operator; 
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
        Operator memory _operator = opInfo[_model.operator_];
        for (uint8 i = 0; i < _ethSecurity.length; i++) {

            if (_ethSecurity[i] < _model.ethShares_[i] + _operator.allocatedEth_[i]) {
            // if (_ethSecurity[i] > _model.ethShares_[i] + allocatedEth[_model.operator_][i]) {
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
            _serSecurity <
            _poc + _operator.allocatedSer_
            // _poc + allocatedSer[_model.operator_]
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

    function verifyTask(bytes memory _taskId, bytes memory _proof) external onlyOperators() {
        //logic to verify task
        Task memory _task = abi.decode(_taskId, (Task));

        uint8 _modelId = _task.modelId_;

        if (0 > _modelId || numModels < _modelId) {
            revert NotModelId("Not Model Id");
        }

        Model memory _model = modelInfo[_modelId];

        require(_model.operator_ == msg.sender, "Wrong operator for task");

        taskVerified[_taskId] = IVerifier(_model.verifier_).verifyProof(_proof);
    }

    function verifiedOffChain(bytes memory _taskId, bool _verified) external onlyAggregators() {
        taskVerified[_taskId] = _verified;
    }

    function submitTask(TaskResponse memory _taskResponse, bool _verification, bytes memory _proof) external {
        
        Task memory _task = abi.decode(_taskResponse.taskId_, (Task));

        uint8 _modelId = _task.modelId_;

        if (0 > _modelId || numModels < _modelId) {
            revert NotModelId("Not Model Id");
        }

        Model memory _model = modelInfo[_modelId];

        Operator memory _operator = opInfo[_model.operator_];

        require(msg.sender == _model.operator_, "Not operator assigned to task");

        if (_verification) {
            _checkFinancialSecurity(_task.poc_, _model, 0);
        } else {
            _checkFinancialSecurity(10*_task.poc_, _model, 0);
        }
        
        if (_taskResponse.proven_) {
            if (taskVerified[_taskResponse.taskId_]) {
                _taskResponse.proven_ = true;
                _clearTask(_taskResponse.taskId_, false);
            }

            else {
                operatorSlashingQueue[msg.sender] = _pushToByteArray(_taskResponse.taskId_, operatorSlashingQueue[msg.sender]);
                slashingQueue.push(_taskResponse.taskId_);
                emit upForSlashing(_model.operator_, _taskResponse.taskId_);
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
                operatorSlashingQueue[msg.sender] = _pushToByteArray(_taskResponse.taskId_, operatorSlashingQueue[msg.sender]);
                slashingQueue.push(_taskResponse.taskId_);
                emit upForSlashing(_model.operator_, _taskResponse.taskId_);
                return;
            }
        }
        }
        
        _operator.submittedTasks_ = _pushToByteArray(_taskResponse.taskId_, _operator.submittedTasks_);
        // submittedTasks[msg.sender] = _pushToByteArray(_taskResponse.taskId_, submittedTasks[msg.sender]);

        opInfo[_model.operator_] = _operator;

        taskResponse[_taskResponse.taskId_] = _taskResponse;

        computeUnits[_model.operator_][_model.computeType_] += 1;

        emit taskResponded(_modelId, _taskResponse.taskId_, _taskResponse);
    }

    function requestProof(bytes memory _taskId) external {

        require(bountyHunter[_taskId] == address(0), "bounty already set");
        Task memory _task = abi.decode(_taskId, (Task));
        Model memory _model = modelInfo[_task.modelId_];
        Operator memory _operator = opInfo[_model.operator_];

        uint256 _amount = PROOF_REQUEST_COST*(_operator.proofRequestExponents_[0]/_operator.proofRequestExponents_[1]);
        // uint256 _amount = PROOF_REQUEST_COST*(proofRequestExponents[_model.operator_][0]/proofRequestExponents[_model.operator_][1]);
        ser.transferFrom(msg.sender, address(this), _amount);

        _operator.proofRequests_ = _pushToByteArray(_taskId, _operator.proofRequests_);
        // proofRequests[_model.operator_] = _pushToByteArray(_taskId, proofRequests[_model.operator_]);

        _operator.proofRequestExponents_[0] += 500;
        // proofRequestExponents[_model.operator_][0] += 500;

        opInfo[_model.operator_] = _operator;
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
        Operator memory _operator = opInfo[_model.operator_];
        _operator.openTasks_ = _removeBytesElement(_operator.openTasks_, _taskId);
        // openTasks[_model.operator_] = _removeBytesElement(openTasks[_model.operator_], _taskId);
        _operator.proofRequests_ = _removeBytesElement(_operator.proofRequests_, _taskId);
        // proofRequests[_model.operator_] = _removeBytesElement(proofRequests[_model.operator_], _taskId);
        if(_slashed) {
            _operator.submittedTasks_ = _removeBytesElement(_operator.submittedTasks_, _taskId);
            // submittedTasks[_model.operator_] = _removeBytesElement(submittedTasks[_model.operator_], _taskId);
        }
        
        operatorSlashingQueue[_model.operator_] = _removeBytesElement(operatorSlashingQueue[_model.operator_], _taskId);
        slashingQueue = _removeBytesElement(slashingQueue, _taskId);
        
        // for (uint8 i = 0; i < allocatedEth[_model.operator_].length; i++) {
        for (uint8 i = 0; i < _operator.allocatedEth_.length; i++) {
            _operator.allocatedEth_[i] -= _model.ethShares_[i];
            // allocatedEth[_model.operator_][i] -= _model.ethShares_[i];
        }
        
        _operator.allocatedSer_ -= 10*_task.poc_;
        // allocatedSer[_model.operator_] -= 10*_task.poc_;

        opInfo[_model.operator_] = _operator;
    }

    function _removeBytesElement(bytes[] memory _byteArray, bytes memory _byteElement) internal pure returns(bytes[] memory) {
        for (uint8 i = 0; i < _byteArray.length; i++) {
            if (keccak256(_byteElement) == keccak256(_byteArray[i])) {
                delete _byteArray[i];
                return _byteArray;
            }
        }
    }

    function _inBytesArray(bytes[] memory _byteArray, bytes memory _byteElement) internal pure returns(bool) {
        for (uint8 i = 0; i < _byteArray.length; i++) {
            if (keccak256(_byteElement) == keccak256(_byteArray[i])) {
                return true;
            }
        }
        return false;
    }

    function slashOperator(bytes memory _taskId, string memory _whySlashed) external onlyAggregators() {

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
            // _wadsToSlash[i] = 1 ether * (_model.ethShares_[i]) / (allocationManager.getAllocation(_model.operator_, opSet, _strategies[i]).currentMagnitude * _operatorShares[i]);
            _wadsToSlash[i] = 1 ether * 1 ether * (_model.ethShares_[i]) / (allocationManager.getAllocation(_model.operator_, opSet, _strategies[i]).currentMagnitude * _operatorShares[i]);
        }

        _wadsToSlash[_model.ethStrategies_.length] = (1 ether * 1 ether * 10 * _task.poc_)/(allocationManager.getAllocation(_model.operator_, opSet, _strategies[_model.ethStrategies_.length]).currentMagnitude * _operatorShares[_model.ethStrategies_.length]);

        IAllocationManagerTypes.SlashingParams memory _slashParams = IAllocationManagerTypes.SlashingParams({operator: _model.operator_, operatorSetId: 0, strategies: _strategies, wadsToSlash: _wadsToSlash, description: _whySlashed});
        
        allocationManager.slashOperator(address(this), _slashParams);

        emit operatorSlashed(_model.operator_, _taskId);
        //May want to implements some custom logic to change bounty amount
        if (bountyHunter[_taskId] != address(0)) {
            ser.transferFrom(address(this), bountyHunter[_taskId], BOUNTY);
            opInfo[_model.operator_].proofRequestExponents_[0] -= 500;
            // proofRequestExponents[_model.operator_][0] -= 500;
        }

        _clearTask(_taskId, true);
    }

    function sendRewards(IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] calldata operatorDirectedRewardsSubmissions, bytes[][] memory _rewardedTasks) external onlyAggregators() {
        uint256 _approvalAmount = 0;
        for (uint8 i = 0; i < operatorDirectedRewardsSubmissions[0].operatorRewards.length; i ++) {
            Operator memory _operator = opInfo[operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator];
            for (uint8 j = 0; j < _rewardedTasks[i].length; j ++) {
                Task memory _task = abi.decode(_rewardedTasks[i][j], (Task));
                require(_task.startingBlock_ + TASK_EXPIRY_BLOCKS < block.number || taskVerified[_rewardedTasks[i][j]], "Task has not expired");
                // require(_removeBytesElement(_operator.submittedTasks_, _rewardedTasks[i][j]).length == (_operator.submittedTasks_.length - 1), "Not in submitted Tasks");
                require(_inBytesArray(_operator.submittedTasks_, _rewardedTasks[i][j]), "Not in submitted Tasks");
                _operator.submittedTasks_ = _removeBytesElement(_operator.submittedTasks_, _rewardedTasks[i][j]);
                opInfo[operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator] = _operator;
                if (_inBytesArray(_operator.openTasks_, _rewardedTasks[i][j])) {
                    _clearTask(_rewardedTasks[i][j], false);
                }
            }
            
            _approvalAmount += operatorDirectedRewardsSubmissions[0].operatorRewards[i].amount;
        }
        
        ser.approve(address(rewardsCoordinator), _approvalAmount);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(address(this), operatorDirectedRewardsSubmissions);
    }

    function modifyModelParameters(Model memory _model, uint8 _modelId) external onlyOperators() {
        Model memory _oldModel = modelInfo[_modelId];
        require(_oldModel.operator_ == msg.sender, "Sender is not model owner");
        if (_oldModel.modelName_ != _model.modelName_) {
            for (uint8 i = 0; i < modelsByName[_oldModel.modelName_].length; i ++) {
                if (modelsByName[_oldModel.modelName_][i] == _modelId) {
                    delete modelsByName[_oldModel.modelName_][i];
                }
            modelsByName[_model.modelName_] = _pushToUint8Array(_modelId, modelsByName[_model.modelName_]);
            }
        }
        modelInfo[_modelId] = _model;
        emit modelUpdated(_modelId, _model);
    }

    function addModels(Model[] memory _models) external onlyOperators() {
        uint8[] memory _modelIds = new uint8[](_models.length);
        for (uint8 i = 0; i < _models.length; i++) {
            numModels++;
            uint8 modelNum = numModels;
            _modelIds[i] = modelNum;
            _models[i].operator_ = msg.sender;
            modelInfo[modelNum] = _models[i];
            modelsByName[_models[i].modelName_] = _pushToUint8Array(modelNum, modelsByName[_models[i].modelName_]);
            opInfo[msg.sender].models_ = _pushToUint8Array(modelNum, opInfo[msg.sender].models_);
        }
        emit newModels(_modelIds);
    }

    function modifyCompute(bytes32[] memory _computeUnitNames, uint8[] memory _computeUnits) external onlyOperators() {
        for (uint8 i = 0; i < _computeUnitNames.length; i++) {
            computeUnits[msg.sender][_computeUnitNames[i]] = _computeUnits[i];
        }
        opInfo[msg.sender].computeUnits_ = _computeUnitNames;
        emit opInfoChanged(msg.sender, opInfo[msg.sender]);
    }

    function updateOperator(address _operator, Operator memory _opInfo) external onlyAggregators() {
        opInfo[_operator] = _opInfo;
        emit opInfoChanged(_operator, _opInfo);
    }

    function pauseOperator() external onlyOperators() {
        opInfo[msg.sender].pausedBlock_ = uint32(block.number);
        emit opInfoChanged(msg.sender, opInfo[msg.sender]);
    }

    function _unpauseOperator() external onlyOperators() {
        opInfo[msg.sender].pausedBlock_ = 0;
        emit opInfoChanged(msg.sender, opInfo[msg.sender]);
    }

    function _pushToByteArray(bytes memory _element, bytes[] memory _arr) internal pure returns (bytes[] memory){
        bytes[] memory _tempBytesArr = new bytes[](_arr.length + 1);
        for (uint8 i = 0; i < _arr.length; i++) {
            _tempBytesArr[i] = _arr[i];
        }
        _tempBytesArr[_arr.length] = _element;
        return _tempBytesArr;
    }

    function _pushToByte32Array(bytes32 _element, bytes32[] memory _arr) internal pure returns (bytes32[] memory){
        bytes32[] memory _tempBytesArr = new bytes32[](_arr.length + 1);
        for (uint8 i = 0; i < _arr.length; i++) {
            _tempBytesArr[i] = _arr[i];
        }
        _tempBytesArr[_arr.length] = _element;
        return _tempBytesArr;
    }

    function _pushToUint8Array(uint8 _element, uint8[] memory _arr) internal pure returns (uint8[] memory){
        uint8[] memory _tempArr = new uint8[](_arr.length + 1);
        for (uint8 i = 0; i < _arr.length; i++) {
            _tempArr[i] = _arr[i];
        }
        _tempArr[_arr.length] = _element;
        return _tempArr;
    }

    function getModelsByName(bytes32 _name) external view returns (uint8[] memory) {
        return modelsByName[_name];
    }

    function getModelInfo(uint8 _modelId) external view returns (Model memory) {
        return modelInfo[_modelId];
    }

    function getOperators() external view returns (address[] memory) {
        return operators;
    }

    function getOperatorInfo(address _operator) external view returns (Operator memory) {
        return opInfo[_operator];
    }
    
    function getAggregators() external view returns (address[] memory) {
        return aggregators;
    }

    // function getOpenTasks(address _operator) external view returns (bytes[] memory) {
    //     return openTasks[_operator];
    // }

    // function getSubmittedTasks(address _operator) external view returns (bytes[] memory) {
    //     return submittedTasks[_operator];
    // }

    function getSlashingQueue() external view returns (bytes[] memory) {
        return slashingQueue;
    }

    function getOperatorSlashingQueue(address _operator) external view returns (bytes[] memory) {
        return operatorSlashingQueue[_operator];
    }

    function getTaskResponse(bytes memory _taskId) external view returns (TaskResponse memory) {
        return taskResponse[_taskId];
    }



    function deregisterOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds
    ) external onlyAggregatorsOrOperator(operator) {
        require(address(this) == avs, "Wrong AVS");
        Operator memory _operator = opInfo[operator];
        require(isAggregator[msg.sender] || uint256(_operator.pausedBlock_) + 2*TASK_EXPIRY_BLOCKS < block.number || _operator.pausedBlock_ > 0, "Not paused long enough");
        for (uint8 i = 0; i < _operator.models_.length; i ++) {
            delete modelInfo[_operator.models_[i]];
        }
        delete opInfo[operator];
        for (uint8 i = 0; i < operators.length; i ++) {
            if (operators[i] == operator) {
                delete operators[i];
            }
        }
        emit operatorDeleted(operator, operatorSetIds);
    }
}
