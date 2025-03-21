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

    uint256[] tempOpModels;
    bytes32[] tempComputeUnitNames;

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

    // constructor(
    //     address _rewardsCoordinator,
    //     address _delegationManager,
    //     address _allocationManager,
    //     IStrategy[] memory _strategies,
    //     string memory _avsMetadata
    // ) OwnableUpgradeable() {
    //     // Set the deployer as an aggregator
    //     isAggregator[msg.sender] = true;
    //     console.log(owner(), msg.sender);

    //     allocationManager = IAllocationManager(_allocationManager);
    //     delegationManager = IDelegationManager(_delegationManager);
    //     rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);

    //     _registerToEigen(_strategies, _avsMetadata);

    //     opSet = OperatorSet({avs: address(this), id: 0});
    //     ser = _strategies[0].underlyingToken();
    // }

    constructor(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) {
        // Set the deployer as an aggregator
        isAggregator[msg.sender] = true;

        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);

        _registerToEigen(_strategies, _avsMetadata);

        opSet = OperatorSet({avs: address(this), id: 0});
        ser = _strategies[0].underlyingToken();
    }

    // function setTaskManager(address _sertnTaskManager) external onlyOwner {
    //     if (_sertnTaskManager == address(0)) {
    //         revert InvalidTaskManager();
    //     }
    //     // Check if the task manager is already set, if so revoke approval first
    //     if (address(sertnTaskManager) != address(0)) {
    //         ser.approve(address(sertnTaskManager), 0);
    //     }
    //     sertnTaskManager = SertnTaskManager(_sertnTaskManager);
    //     isAggregator[_sertnTaskManager] = true;
    //     ser.approve(address(sertnTaskManager), 100 ether);
    // }

    function setTaskManager(address _sertnTaskManager) external onlyAggregators() {
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
        // TODO: Is this needed:?
        // Yes, part of format for AVS Registrar
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

            modelInfo[modelNum] = abi.encode(_models[i]);
            // modelInfo[modelNum] = _models[i];

        }

        for (uint8 i = 0; i < _computeUnitNames.length; i++) {
            computeUnits[operator][_computeUnitNames[i]] = _computeUnits[i];
        }

        Operator memory _operator = Operator({
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
            _operator.allocatedEth_[i] = 0;
        }

        opInfo[operator] = abi.encode(_operator);

        // opInfo[operator] = Operator({
        //     models_: _modelIds,
        //     computeUnits_: _computeUnitNames,
        //     openTasks_: new bytes[](0),
        //     submittedTasks_: new bytes[](0),
        //     proofRequests_: new bytes[](0),
        //     allocatedEth_: new uint256[](_models[0].ethShares_.length),
        //     allocatedSer_: 0,
        //     proofRequestExponents_: [uint32(1e3),uint32(1e3)],
        //     pausedBlock_: 0
        // });

        // for (uint8 i = 0; i < _models[0].ethShares_.length; i++) {
        //     opInfo[operator].allocatedEth_[i] = 0;
        // }

        operators.push(operator);
        isOperator[operator] = true;

        emit newOperator(operator);
        emit newModels(_modelIds);
    }

    function sendRewards(IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] calldata operatorDirectedRewardsSubmissions, bytes[][] memory _rewardedTasks) external onlyAggregators() {
        uint256 _approvalAmount = 0;
        for (uint8 i = 0; i < operatorDirectedRewardsSubmissions[0].operatorRewards.length; i ++) {
            Operator memory _operator = abi.decode(opInfo[operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator],(Operator));
            for (uint8 j = 0; j < _rewardedTasks[i].length; j ++) {
                Task memory _task = abi.decode(_rewardedTasks[i][j], (Task));
                require(_task.startingBlock_ + TASK_EXPIRY_BLOCKS < block.number || sertnTaskManager.taskVerified(_rewardedTasks[i][j]), "Task has not expired");
                require(_inBytesArray(_operator.submittedTasks_, _rewardedTasks[i][j]), "Not in submitted Tasks");
                _operator.submittedTasks_ = _removeBytesElement(_operator.submittedTasks_, _rewardedTasks[i][j]);
                opInfo[operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator] = abi.encode(_operator);
                if (_inBytesArray(_operator.openTasks_, _rewardedTasks[i][j])) {
                    sertnTaskManager.clearTask(_rewardedTasks[i][j], false);
                }
            }

            _approvalAmount += operatorDirectedRewardsSubmissions[0].operatorRewards[i].amount;
        }

        ser.approve(address(rewardsCoordinator), _approvalAmount);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(address(this), operatorDirectedRewardsSubmissions);
    }

    // function _verifyTask(bytes memory _taskId, bytes memory _proof) external onlyTaskManager() {
    //     //logic to verify task
    //     Task memory _task = abi.decode(_taskId, (Task));

    //     uint256 _modelId = _task.modelId_;

    //     if (0 > _modelId || numModels < _modelId) {
    //         revert NotModelId("Not Model Id");
    //     }

    //     Model memory _model = abi.decode(modelInfo[_modelId], (Model));

    //     taskVerified[_taskId] = IVerifier(_model.verifier_).verifyProof(_proof);
    // }

    // function verifyTask(bytes memory _taskId, bytes memory _proof) external onlyOperators() {
    //     //logic to verify task
    //     Task memory _task = abi.decode(_taskId, (Task));

    //     uint256 _modelId = _task.modelId_;

    //     if (0 > _modelId || numModels < _modelId) {
    //         revert NotModelId("Not Model Id");
    //     }
    //     Model memory _model = abi.decode(modelInfo[_modelId], (Model));
    //     // Model memory _model = modelInfo[_modelId];

    //     if (_model.operator_ != msg.sender) {
    //         revert IncorrectOperator();
    //     }

    //     taskVerified[_taskId] = IVerifier(_model.verifier_).verifyProof(_proof);
    // }

    // function verifiedOffChain(bytes memory _taskId, bool _verified) external onlyAggregators() {
    //     taskVerified[_taskId] = _verified;
    // }
    // function clearTask(bytes memory _taskId) external onlyAggregators() {
    //     _clearTask(_taskId, false);
    // }

    // function _clearTask(bytes memory _taskId, bool _slashed) internal {
    //     Task memory _task = abi.decode(_taskId, (Task));

    //     if (_task.startingBlock_ + TASK_EXPIRY_BLOCKS < block.number || taskVerified[_taskId] || _slashed) {
    //         revert TaskNotExpired();
    //     }

    //     uint256 _modelId = _task.modelId_;

    //     if (0 > _modelId || numModels < _modelId) {
    //         revert NotModelId("Not Model Id");
    //     }
    //     Model memory _model = abi.decode(modelInfo[_modelId], (Model));
    //     // Model memory _model = modelInfo[_modelId];
    //     Operator memory _operator = abi.decode(opInfo[_model.operator_], (Operator));
    //     _operator.openTasks_ = _removeBytesElement(_operator.openTasks_, _taskId);
    //     _operator.proofRequests_ = _removeBytesElement(_operator.proofRequests_, _taskId);
    //     if(_slashed) {
    //         _operator.submittedTasks_ = _removeBytesElement(_operator.submittedTasks_, _taskId);
    //     }

    //     operatorSlashingQueue[_model.operator_] = _removeBytesElement(operatorSlashingQueue[_model.operator_], _taskId);
    //     slashingQueue = _removeBytesElement(slashingQueue, _taskId);

    //     for (uint8 i = 0; i < _operator.allocatedEth_.length; i++) {
    //         _operator.allocatedEth_[i] -= _model.ethShares_[i];
    //     }

    //     _operator.allocatedSer_ -= 10*_task.poc_;

    //     opInfo[_model.operator_] = abi.encode(_operator);
    // }

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
    function modifyModelParameters(Model memory _model, uint256 _modelId) external onlyOperators() {
        Model memory _oldModel = abi.decode(modelInfo[_modelId], (Model));
        // Model storage _oldModel = modelInfo[_modelId];
        if (_oldModel.operator_ != msg.sender) {
            revert IncorrectOperator();
        }
        modelInfo[_modelId] = abi.encode(_model);
        //Assuming this saves on gas fees?
        // _oldModel.modelName_ = _model.modelName_;
        // _oldModel.verifier_ = _model.verifier_;
        // _oldModel.benchData_ = _model.benchData_;
        // _oldModel.inputs_ = _model.inputs_;
        // _oldModel.output_ = _model.output_;
        // _oldModel.maxBlocks_ = _model.maxBlocks_;
        // _oldModel.ethStrategies_ = _model.ethStrategies_;
        // _oldModel.ethShares_ = _model.ethShares_;
        // _oldModel.baseFee_ = _model.baseFee_;
        // _oldModel.maxSer_ = _model.maxSer_;
        // _oldModel.computeType_ = _model.computeType_;
        // _oldModel.proveOnResponse_ = _model.proveOnResponse_;
        // _oldModel.available_ = _model.available_;

        emit modelUpdated(_modelId, _model);
    }

    function slashOperator(bytes memory _taskId, string memory _whySlashed) external onlyAggregators() {

        Task memory _task = abi.decode(_taskId, (Task));
        uint256 _modelId = _task.modelId_;
        if (0 > _modelId || numModels < _modelId) {
            revert NotModelId("Not Model Id");
        }
        Model memory _model = abi.decode(modelInfo[_modelId], (Model));
        // Model memory _model = sertnServiceManager.getModelInfo(_modelId);
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
            _wadsToSlash[i] = 1 ether * 1 ether * (_model.ethShares_[i]) / (allocationManager.getAllocation(_model.operator_, opSet, _strategies[i]).currentMagnitude * _operatorShares[i]);
        }
        _wadsToSlash[_model.ethStrategies_.length] = (1 ether * 1 ether * 10 * _task.poc_)/(allocationManager.getAllocation(_model.operator_, opSet, _strategies[_model.ethStrategies_.length]).currentMagnitude * _operatorShares[_model.ethStrategies_.length]);
        IAllocationManagerTypes.SlashingParams memory _slashParams = IAllocationManagerTypes.SlashingParams({operator: _model.operator_, operatorSetId: 0, strategies: _strategies, wadsToSlash: _wadsToSlash, description: _whySlashed});
        allocationManager.slashOperator(address(this), _slashParams);

        emit operatorSlashed(_model.operator_, _taskId);
        // TODO: Custom logic to change bounty amount
        if (bountyHunter[_taskId] != address(0)) {
            ser.transferFrom(address(this), bountyHunter[_taskId], BOUNTY);
            Operator memory _operator = abi.decode(opInfo[_model.operator_], (Operator));
            _operator.proofRequestExponents_[0] -= 500;
            opInfo[_model.operator_] = abi.encode(_operator);
        }
        sertnTaskManager.clearTask(_taskId, true);
    }

    function addModels(Model[] memory _models) external onlyOperators() {
        uint256[] memory _modelIds = new uint256[](_models.length);
        for (uint256 i = 0; i < _models.length; i++) {
            numModels++;
            uint256 modelNum = numModels;
            _modelIds[i] = modelNum;
            _models[i].operator_ = msg.sender;
            modelInfo[modelNum] = abi.encode(_models[i]);
            // modelInfo[modelNum] = _models[i];
            Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
            tempOpModels = _operator.models_;
            tempOpModels.push(modelNum);
            _operator.models_ = tempOpModels;
            opInfo[msg.sender] = abi.encode(_operator);
            // opInfo[msg.sender].models_.push(uint8(modelNum));
        }
        emit newModels(_modelIds);
    }

    function modifyCompute(bytes32[] memory _computeUnitNames, uint8[] memory _computeUnits) external onlyOperators() {
        for (uint8 i = 0; i < _computeUnitNames.length; i++) {
            computeUnits[msg.sender][_computeUnitNames[i]] = _computeUnits[i];
        }

        Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
        delete _operator.computeUnits_;
        // delete opInfo[msg.sender].computeUnits_;
        for (uint8 i = 0; i < _computeUnitNames.length; i++) {
            tempComputeUnitNames = _operator.computeUnits_;
            tempComputeUnitNames.push(_computeUnitNames[i]);
            _operator.computeUnits_ = tempComputeUnitNames;
        }
        opInfo[msg.sender] = abi.encode(_operator);
        emit opInfoChanged(msg.sender, opInfo[msg.sender]);
    }

    function updateOperator(address _operator, Operator memory _opInfo) external onlyAggregators() {
        opInfo[_operator] = abi.encode(_opInfo);
        emit opInfoChanged(_operator, abi.encode(_opInfo));
    }

    function pauseOperator() external onlyOperators() {
        Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
        _operator.pausedBlock_ = uint32(block.number);
        opInfo[msg.sender] = abi.encode(_operator);

        // opInfo[msg.sender].pausedBlock_ = uint32(block.number);
        emit opInfoChanged(msg.sender, opInfo[msg.sender]);
    }

    function _unpauseOperator() external onlyOperators() {
         Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
        _operator.pausedBlock_ = 0;
        opInfo[msg.sender] = abi.encode(_operator);
        emit opInfoChanged(msg.sender, opInfo[msg.sender]);
    }

    // function getComputeUnits(address _operator, bytes32 _computeType) external view returns (uint8) {
    //     return computeUnits[_operator][_computeType];
    // }

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
        operatorSlashingQueue[_operator].push(_taskId);
    }

    function setTaskResponse(TaskResponse memory _taskResponse) external onlyTaskManager() {
        taskResponse[_taskResponse.taskId_] = abi.encode(_taskResponse);
    }

    function setBountyHunter(bytes memory _taskId, address _bountyHunter) external onlyTaskManager() {
        bountyHunter[_taskId] = _bountyHunter;
    }


    function deregisterOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds
    ) external {

        require(isAggregator[msg.sender] || msg.sender == operator, "No permission to call");
        if (address(this) != avs) {
            revert IncorrectAVS();
        }
        Operator memory _operator = abi.decode(opInfo[operator], (Operator));
        if (isAggregator[msg.sender] || uint256(_operator.pausedBlock_) + 2*TASK_EXPIRY_BLOCKS < block.number || _operator.pausedBlock_ > 0) {
            revert NotPausedLongEnough();
        }
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
