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
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";

import {IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IVerifier} from "./IVerifier.sol";
import "@eigenlayer/contracts/libraries/OperatorSetLib.sol";

import "@openzeppelin-upgrades/contracts/utils/math/MathUpgradeable.sol";

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {SertnTaskManager} from "./SertnTaskManager.sol";
import {ModelStorage} from "./ModelStorage.sol";

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
    OwnableUpgradeable,
    ReentrancyGuard
{

    address[] public aggregators;
    address[] public operators;
    bytes[] public slashingQueue;

    IERC20 public ser;
    uint256 public BASE_APPROVAL;

    IAllocationManager public allocationManager;
    IDelegationManager public delegationManager;
    IRewardsCoordinator public rewardsCoordinator;
    OperatorSet public opSet;
    SertnTaskManager public sertnTaskManager;
    ModelStorage public modelStorage;

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

    constructor(
        address _rewardsCoordinator,
        address _delegationManager,
        address _allocationManager,
        IStrategy[] memory _strategies,
        string memory _avsMetadata
    ) OwnableUpgradeable() {
        // Set the deployer as an aggregator
        isAggregator[msg.sender] = true;
        BASE_APPROVAL = 100 ether;
        _transferOwnership(msg.sender);

        allocationManager = IAllocationManager(_allocationManager);
        delegationManager = IDelegationManager(_delegationManager);
        rewardsCoordinator = IRewardsCoordinator(_rewardsCoordinator);

        _registerToEigen(_strategies, _avsMetadata);

        opSet = OperatorSet({avs: address(this), id: 0});
        ser = _strategies[0].underlyingToken();
    }

    function setTaskManagerandModelStorage(address _sertnTaskManager, address _modelStorage) external onlyOwner {
        if (_sertnTaskManager == address(0)) {
            revert InvalidTaskManager();
        }
        // Check if the task manager is already set, if so revoke approval first
        if (address(sertnTaskManager) != address(0)) {
            ser.approve(address(sertnTaskManager), 0);
        }
        sertnTaskManager = SertnTaskManager(_sertnTaskManager);
        isAggregator[_sertnTaskManager] = true;
        ser.approve(address(sertnTaskManager), BASE_APPROVAL);

        modelStorage = ModelStorage(_modelStorage);
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
        for (uint256 i; i < _strategies.length;) {
            tokenToStrategy[
                address(_strategies[i].underlyingToken())
            ] = _strategies[i];
            unchecked { ++i; }
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
        if (_aggregator == address(0)) {
            revert InvalidAggregator();
        }
        if (isAggregator[_aggregator]) {
            revert AggregatorAlreadyExists();
        }
        aggregators.push(_aggregator);
        isAggregator[_aggregator] = true;
    }

    function supportsAVS(address avs) external view returns (bool) {
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
    ) external override {
        if (avs != address(this)) {
            revert NoPermission();
        }
        if (operatorSetIds.length != 1 || operatorSetIds[0] != 0) {
            revert IncorrectOperatorSetIds();
        }
        (   Model[] memory _models,
            OperatorModel[] memory _operatorModels,
            uint256[] memory _modelIds,
            bytes32[] memory _computeUnitNames,
            uint256[] memory _computeUnits
        ) = abi.decode(data, (Model[], OperatorModel[], uint256[], bytes32[], uint256[]));
        for (uint256 i; i < _operatorModels.length;) {

            if (!(_operatorModels[i].maxBlocks_ < taskExpiryBlocks - 1e2)) {
                revert MaxBlocksTooLong();
            }

            if (_modelIds[i] > modelStorage.numModels()) {
                _modelIds[i] = modelStorage.createNewModel(_models[i]);

            } else {
                Model memory _model = abi.decode(modelStorage.modelInfo(_modelIds[i]),(Model));
                if (_model.modelVerifier_ != _models[i].modelVerifier_ || _model.modelKey_ != _models[i].modelKey_) {
                    revert NotModelId();
                }
                modelStorage.JoinOperatorList(_modelIds[i], operator);
            }

            operatorModelInfo[_modelIds[i]][operator] = abi.encode(_operatorModels[i]);
            unchecked { ++i; }
        }

        for (uint256 i; i < _computeUnitNames.length;) {
            computeUnits[operator][_computeUnitNames[i]] = _computeUnits[i];
            unchecked { ++i; }
        }

        Operator memory _operator = Operator({
            models_: _modelIds,
            computeUnits_: _computeUnitNames,
            openTasks_: new bytes[](0),
            submittedTasks_: new bytes[](0),
            proofRequests_: new bytes[](0),
            allocatedEth_: new uint256[](_operatorModels[0].ethShares_.length),
            allocatedSer_: 0,
            proofRequestExponents_: [uint256(1e3),uint256(1e3)],
            pausedBlock_: 0
        });

        for (uint256 i; i < _operatorModels[0].ethShares_.length;) {
            _operator.allocatedEth_[i] = 0;
            unchecked { ++i; }
        }

        opInfo[operator] = abi.encode(_operator);

        operators.push(operator);
        isOperator[operator] = true;

        emit NewOperator(operator);
        emit NewModels(_modelIds);
    }

    function sendRewards(IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] calldata operatorDirectedRewardsSubmissions, bytes[][] calldata _rewardedTasks) external onlyAggregators() {
        uint256 _approvalAmount;
        for (uint256 i; i < operatorDirectedRewardsSubmissions[0].operatorRewards.length;) {
            Operator memory _operator = abi.decode(opInfo[operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator],(Operator));
            for (uint256 j; j < _rewardedTasks[i].length;) {
                Task memory _task = abi.decode(_rewardedTasks[i][j], (Task));
                if (!(_task.startingBlock_ + taskExpiryBlocks < uint32(block.number)) || sertnTaskManager.taskVerified(_rewardedTasks[i][j])) {
                    revert TaskNotExpired();
                }
                if (!_inBytesArray(_operator.submittedTasks_, _rewardedTasks[i][j])) {
                    revert NotInSubmittedTasks();
                }
                _operator.submittedTasks_ = _removeBytesElement(_operator.submittedTasks_, _rewardedTasks[i][j]);
                opInfo[operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator] = abi.encode(_operator);
                if (_inBytesArray(_operator.openTasks_, _rewardedTasks[i][j])) {
                    sertnTaskManager.clearTask(_rewardedTasks[i][j], false);
                }
                unchecked { ++j; }
            }

            _approvalAmount += operatorDirectedRewardsSubmissions[0].operatorRewards[i].amount;
            unchecked { ++i; }
        }

        ser.approve(address(rewardsCoordinator), _approvalAmount);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(address(this), operatorDirectedRewardsSubmissions);
    }

    function _removeBytesElement(bytes[] memory _byteArray, bytes calldata _byteElement) internal pure returns(bytes[] memory) {
        for (uint256 i; i < _byteArray.length;) {
            if (keccak256(_byteElement) == keccak256(_byteArray[i])) {
                delete _byteArray[i];
                return _byteArray;
            }
            unchecked { ++i; }
        }
    }

    function _inBytesArray(bytes[] memory _byteArray, bytes calldata _byteElement) internal pure returns(bool) {
        for (uint256 i; i < _byteArray.length; ) {
            if (keccak256(_byteElement) == keccak256(_byteArray[i])) {
                return true;
            }
            unchecked { ++i; }
        }
        return false;
    }

    //Want to have the option to deregister a model rather than just pause?
    function modifyOperatorModelParameters(OperatorModel calldata _operatorModel, uint256 _modelId) external onlyOperators() {
        Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
        bool _registered;
        for (uint256 i; i < _operator.models_.length;) {
            if (_operator.models_[i] == _modelId) {
                _registered = true;
            }
            unchecked { ++i; }
        }

        if (!_registered) {
            revert NotRegisteredToModel();
        }

        operatorModelInfo[_modelId][msg.sender] = abi.encode(_operatorModel);

        emit ModelUpdated(_modelId, _operatorModel);
    }

    function slashOperator(bytes calldata _taskId, string calldata _whySlashed) external onlyAggregators() nonReentrant {

        Task memory _task = abi.decode(_taskId, (Task));

        OperatorModel memory _operatorModel = abi.decode(operatorModelInfo[_task.modelId_][_task.operator_], (OperatorModel));
        //Should this be included?
        if (!(uint32(block.number) - _task.startingBlock_ < taskExpiryBlocks)) {
            revert TaskExpired();
        }
        IStrategy[] memory _strategies = new IStrategy[](_operatorModel.ethStrategies_.length + 1);
        for (uint256 i; i < _operatorModel.ethStrategies_.length;) {
            _strategies[i] = _operatorModel.ethStrategies_[i];
            unchecked { ++i; }
        }
        _strategies[_operatorModel.ethStrategies_.length] = tokenToStrategy[address(ser)];
        uint256[] memory _wadsToSlash = new uint256[](_strategies.length);
        uint256[] memory _operatorShares = delegationManager.getOperatorShares(_task.operator_, _strategies);
        for (uint256 i = 0; i < _operatorModel.ethStrategies_.length;) {
        
            _wadsToSlash[i] = MathUpgradeable.mulDiv(1 ether * 1 ether , _operatorModel.ethShares_[i],(allocationManager.getAllocation(_task.operator_, opSet, _strategies[i]).currentMagnitude * _operatorShares[i]));
            unchecked { ++i; }
        }
        if (_task.proveOnResponse_) {
            
            _wadsToSlash[_operatorModel.ethStrategies_.length] = MathUpgradeable.mulDiv(1 ether * 1 ether, _task.poc_, allocationManager.getAllocation(_task.operator_, opSet, _strategies[_operatorModel.ethStrategies_.length]).currentMagnitude * _operatorShares[_operatorModel.ethStrategies_.length]);
        } else {
            
            _wadsToSlash[_operatorModel.ethStrategies_.length] = MathUpgradeable.mulDiv(1 ether * 1 ether, 10*_task.poc_, allocationManager.getAllocation(_task.operator_, opSet, _strategies[_operatorModel.ethStrategies_.length]).currentMagnitude * _operatorShares[_operatorModel.ethStrategies_.length]);
        }
        IAllocationManagerTypes.SlashingParams memory _slashParams = IAllocationManagerTypes.SlashingParams({operator: _task.operator_, operatorSetId: 0, strategies: _strategies, wadsToSlash: _wadsToSlash, description: _whySlashed});
        allocationManager.slashOperator(address(this), _slashParams);

        emit OperatorSlashed(_taskId);
        // TODO: Custom logic to change bounty amount
        if (bountyHunter[_taskId] != address(0)) {
            if (!ser.transferFrom(address(this), bountyHunter[_taskId], BOUNTY)) {
                revert TransferFailed();
            }
            Operator memory _operator = abi.decode(opInfo[_task.operator_], (Operator));
            _operator.proofRequestExponents_[0] = 1e3;
            opInfo[_task.operator_] = abi.encode(_operator);
        }
        sertnTaskManager.clearTask(_taskId, true);
    }

    function addModels(OperatorModel[] memory _operatorModels, uint256[] memory _modelIds) external onlyOperators() {
        Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
        uint256[] memory _tempOpModels = new uint256[](_operator.models_.length + _modelIds.length);

        for (uint256 i; i < _operator.models_.length;) {
                _tempOpModels[i] = _operator.models_[i];
                unchecked { ++i; }
            }

        for (uint256 i; i < _operatorModels.length;) {

            if (_operatorModels[i].maxBlocks_ + 1e2 > taskExpiryBlocks) {
                revert MaxBlocksTooLong();
            }

            Model memory _model = abi.decode(modelStorage.modelInfo(_modelIds[i]), (Model));
            bool _registered;
            for (uint256 j; j < _model.operators_.length;) {
                if (_model.operators_[j] == msg.sender) {
                    _registered = true;
                }
                unchecked { ++j; }
            }
            if (!_registered) {
                revert NotRegisteredToModel();
            }
            operatorModelInfo[_modelIds[i]][msg.sender] = abi.encode(_operatorModels[i]);
            
            
            _tempOpModels[_operator.models_.length + i] = _modelIds[i];
            unchecked { ++i; }
        }
        _operator.models_ = _tempOpModels;
        opInfo[msg.sender] = abi.encode(_operator);
            
        
        emit NewModels(_modelIds);
    }

    function modifyCompute(bytes32[] calldata _computeUnitNames, uint256[] calldata _computeUnits) external onlyOperators() {
        for (uint256 i; i < _computeUnitNames.length;) {
            computeUnits[msg.sender][_computeUnitNames[i]] = _computeUnits[i];
            unchecked { ++i; }
        }
        Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
        _operator.computeUnits_ = _computeUnitNames;
        opInfo[msg.sender] = abi.encode(_operator);
        emit OpInfoChanged(msg.sender, opInfo[msg.sender]);
    }

    function updateOperator(address _operator, Operator calldata _opInfo) external onlyAggregators() {
        opInfo[_operator] = abi.encode(_opInfo);
        emit OpInfoChanged(_operator, abi.encode(_opInfo));
    }

    function pauseOperator() external onlyOperators() {
        Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
        _operator.pausedBlock_ = uint32(block.number);
        opInfo[msg.sender] = abi.encode(_operator);

        // opInfo[msg.sender].pausedBlock_ = uint32(block.number);
        emit OpInfoChanged(msg.sender, opInfo[msg.sender]);
    }

    function unpauseOperator() external onlyOperators() {
         Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
        _operator.pausedBlock_ = 0;
        opInfo[msg.sender] = abi.encode(_operator);
        emit OpInfoChanged(msg.sender, opInfo[msg.sender]);
    }

    function setComputeUnits(address _operator, bytes32 _computeType, bool increment) external onlyTaskManager() {
        if (increment) {
            computeUnits[_operator][_computeType] += 1;
        } else {
            computeUnits[_operator][_computeType] -= 1;
        }
    }

    function pushToSlashingQueue(bytes calldata _taskId) external onlyTaskManager() {
        slashingQueue.push(_taskId);
    }

    function pushToOperatorSlashingQueue(address _operator, bytes calldata _taskId) external onlyTaskManager() {
        operatorSlashingQueue[_operator].push(_taskId);
    }

    function setTaskResponse(TaskResponse calldata _taskResponse) external onlyTaskManager() {
        taskResponse[_taskResponse.taskId_] = abi.encode(_taskResponse);
    }

    function setBountyHunter(bytes calldata _taskId, address _bountyHunter) external onlyTaskManager() {
        bountyHunter[_taskId] = _bountyHunter;
    }

    function setTaskExpiryBlocks(uint32 _taskExpiryBlocks) external onlyOwner() {
        taskExpiryBlocks = _taskExpiryBlocks;
        emit TaskExpiryChanged(taskExpiryBlocks);
    }


    function deregisterOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds
    ) external override {
        if (!(isAggregator[msg.sender] || msg.sender == operator)) {
            revert NoPermission();
        }
        Operator memory _operator = abi.decode(opInfo[operator], (Operator));
        if (!isAggregator[msg.sender] || _operator.pausedBlock_ + 2*taskExpiryBlocks > uint32(block.number)) {
            revert NotPausedLongEnough();
        }
        for (uint256 i; i < _operator.models_.length;) {
            modelStorage.RemoveFromOperatorList(_operator.models_[i], operator);
            delete operatorModelInfo[_operator.models_[i]][operator];
            unchecked { ++i; }
        }
        delete opInfo[operator];
        for (uint256 i; i < operators.length;) {
            if (operators[i] == operator) {
                delete operators[i];
            }
            unchecked { ++i; }
        }
        isOperator[operator] = false;
        emit OperatorDeleted(operator, operatorSetIds);
    }
}
