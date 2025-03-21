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
import {SertnTaskManager} from "./SertnTaskManager.sol";

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
    uint256 public numModels = 0;
    address[] aggregators;
    address[] operators;
    bytes[] slashingQueue;

    IERC20 ser;

    IAllocationManager allocationManager;
    IDelegationManager delegationManager;
    IRewardsCoordinator rewardsCoordinator;
    OperatorSet opSet;
    SertnTaskManager sertnTaskManager;

    modifier onlyAggregators() {
        if (!isAggregator[msg.sender]) {
            revert NotAggregator();
        }
        _;
    }

    modifier onlyAggregatorsOrOperator(address _operator) {
        require(isAggregator[msg.sender] || _operator == msg.sender, "Not registered aggregator or concerned operator");
        _;
    }

    modifier onlyOperators() {
        if (!isOperator[msg.sender]) {
            revert NotOperator();
        }
        _;
    }

    modifier onlyTaskManager() {
        if (msg.sender != address(sertnTaskManager)) {
            revert NotTaskManager();
        }
        _;
    }

    constructor(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) OwnableUpgradeable() {
        // Set the deployer as an aggregator
        isAggregator[msg.sender] = true;

        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);

        _registerToEigen(_strategies, _avsMetadata);

        opSet = OperatorSet({avs: address(this), id: 0});
        ser = _strategies[0].underlyingToken();
    }

    function setTaskManager(address _sertnTaskManager) external onlyOwner {
        if (_sertnTaskManager == address(0)) {
            revert InvalidTaskManager();
        }
        // Check if the task manager is already set, if so revoke approval first
        if (address(sertnTaskManager) != address(0)) {
            ser.approve(address(sertnTaskManager), 0);
        }
        sertnTaskManager = SertnTaskManager(_sertnTaskManager);
        isAggregator[_sertnTaskManager] = true;
        ser.approve(address(sertnTaskManager), 100 ether);
    }

    function _registerToEigen(
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) internal {
        allocationManager.updateAVSMetadataURI(address(this), _avsMetadata);
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

    function addAggregator(address _aggregator) external onlyOwner {
        if (isAggregator[_aggregator]) {
            revert AggregatorAlreadyExists();
        }
        aggregators.push(_aggregator);
        isAggregator[_aggregator] = true;
    }

    function supportsAVS(address avs) external view override returns (bool) {
        // TODO: Is this needed?
        if (avs != address(this)) {
            revert IncorrectAVS();
        }
        return true;
    }

    function registerOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds,
        bytes calldata data
    ) external {
        if (avs != address(this)) {
            revert IncorrectAVS();
        }
        if (operatorSetIds.length != 1 || operatorSetIds[0] != 0) {
            revert IncorrectOperatorSetIds();
        }
        (
            Model[] memory _models,
            bytes32[] memory _computeUnitNames,
            uint8[] memory _computeUnits
        ) = abi.decode(data, (Model[], bytes32[], uint8[]));
        uint256[] memory _modelIds = new uint256[](_models.length);
        for (uint256 i = 0; i < _models.length; i++) {
            numModels++;
            uint256 modelNum = numModels;
            _modelIds[i] = modelNum;
            _models[i].operator_ = operator;
            modelInfo[modelNum] = _models[i];

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

        emit newOperator(operator);
        emit newModels(_modelIds);
    }

    function _verifyTask(bytes memory _taskId, bytes memory _proof) external onlyTaskManager() {
        //logic to verify task
        Task memory _task = abi.decode(_taskId, (Task));

        uint256 _modelId = _task.modelId_;

        if (0 > _modelId || numModels < _modelId) {
            revert NotModelId("Not Model Id");
        }

        Model memory _model = modelInfo[_modelId];

        taskVerified[_taskId] = IVerifier(_model.verifier_).verifyProof(_proof);
    }

    function verifyTask(bytes memory _taskId, bytes memory _proof) external onlyOperators() {
        //logic to verify task
        Task memory _task = abi.decode(_taskId, (Task));

        uint256 _modelId = _task.modelId_;

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
    function clearTask(bytes memory _taskId) external onlyAggregators() {
        _clearTask(_taskId, false);
    }

    function _clearTask(bytes memory _taskId, bool _slashed) internal {
        Task memory _task = abi.decode(_taskId, (Task));
        require(_task.startingBlock_ + TASK_EXPIRY_BLOCKS < block.number || taskVerified[_taskId] || _slashed, "Task has not expired");

        uint256 _modelId = _task.modelId_;

        if (0 > _modelId || numModels < _modelId) {
            revert NotModelId("Not Model Id");
        }

        Model memory _model = modelInfo[_modelId];
        Operator memory _operator = opInfo[_model.operator_];
        _operator.openTasks_ = _removeBytesElement(_operator.openTasks_, _taskId);
        _operator.proofRequests_ = _removeBytesElement(_operator.proofRequests_, _taskId);
        if(_slashed) {
            _operator.submittedTasks_ = _removeBytesElement(_operator.submittedTasks_, _taskId);
        }

        operatorSlashingQueue[_model.operator_] = _removeBytesElement(operatorSlashingQueue[_model.operator_], _taskId);
        slashingQueue = _removeBytesElement(slashingQueue, _taskId);

        for (uint8 i = 0; i < _operator.allocatedEth_.length; i++) {
            _operator.allocatedEth_[i] -= _model.ethShares_[i];
        }

        _operator.allocatedSer_ -= 10*_task.poc_;

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
    function modifyModelParameters(Model memory _model, uint8 _modelId) external onlyOperators() {
        Model memory _oldModel = modelInfo[_modelId];
        require(_oldModel.operator_ == msg.sender, "Sender is not model owner");
        modelInfo[_modelId] = _model;
        emit modelUpdated(_modelId, _model);
    }

    function addModels(Model[] memory _models) external onlyOperators() {
        uint256[] memory _modelIds = new uint256[](_models.length);
        for (uint256 i = 0; i < _models.length; i++) {
            numModels++;
            uint256 modelNum = numModels;
            _modelIds[i] = modelNum;
            _models[i].operator_ = msg.sender;
            modelInfo[modelNum] = _models[i];
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

    function getComputeUnits(address _operator, bytes32 _computeType) external view returns (uint8) {
        return computeUnits[_operator][_computeType];
    }

    function setComputeUnits(address _operator, bytes32 _computeType, bool increment) external onlyTaskManager() {
        if (increment) {
            computeUnits[_operator][_computeType] += 1;
        } else {
            computeUnits[_operator][_computeType] -= 1;
        }
    }

    function pushToSlashingQueue(bytes memory _taskId) external onlyTaskManager() {
        slashingQueue.push(_taskId);
    }

    function pushToOperatorSlashingQueue(address _operator, bytes memory _taskId) external onlyTaskManager() {
        operatorSlashingQueue[_operator] = _pushToByteArray(_taskId, operatorSlashingQueue[_operator]);
    }

    function setTaskResponse(TaskResponse memory _taskResponse) external onlyTaskManager() {
        taskResponse[_taskResponse.taskId_] = _taskResponse;
    }

    function setBountyHunter(bytes memory _taskId, address _bountyHunter) external onlyTaskManager() {
        bountyHunter[_taskId] = _bountyHunter;
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
