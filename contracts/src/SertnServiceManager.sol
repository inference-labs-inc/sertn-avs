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
    OwnableUpgradeable
{
    uint256 public numOperatorModels;
    address[] public aggregators;
    address[] public operators;
    bytes[] public slashingQueue;

    uint256[] _tempOpModels;
    bytes32[] _tempComputeUnitNames;

    IERC20 public ser;

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
        ser.approve(address(sertnTaskManager), 100 ether);

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
        (   Model[] memory _models,
            OperatorModel[] memory _operatorModels,
            bytes32[] memory _computeUnitNames,
            uint256[] memory _computeUnits
        ) = abi.decode(data, (Model[], OperatorModel[], bytes32[], uint256[]));
        uint256[] memory _operatorModelIds = new uint256[](_operatorModels.length);
        for (uint256 i = 0; i < _operatorModels.length; i++) {
            uint256 operatorModelNum = numOperatorModels;
            numOperatorModels++;
            _operatorModelIds[i] = operatorModelNum;
            _operatorModels[i].operator_ = operator;
            if (!(_operatorModels[i].maxBlocks_ < TASK_EXPIRY_BLOCKS - 1e2)) {
                revert MaxBlocksTooLong();
            }

            if (_operatorModels[i].modelId_ > modelStorage.numModels()) {
                _operatorModels[i].modelId_ = modelStorage.createNewModel(_models[i]);

            } else {
                if (modelStorage.modelAddresses(_operatorModels[i].modelId_) != _models[i].modelVerifier_) {
                    revert NotModelId();
                }   
                modelStorage.JoinOperatorList(_operatorModels[i].modelId_, msg.sender);
            }

            // require(_operatorModels[i].maxBlocks_ < TASK_EXPIRY_BLOCKS - 1e2, "max blocks too long");
            
            operatorModelInfo[operatorModelNum] = abi.encode(_operatorModels[i]);
            // operatorModelInfo[modelNum] = _operatorModels[i];

        }

        for (uint256 i = 0; i < _computeUnitNames.length; i++) {
            computeUnits[operator][_computeUnitNames[i]] = _computeUnits[i];
        }

        Operator memory _operator = Operator({
            models_: _operatorModelIds,
            computeUnits_: _computeUnitNames,
            openTasks_: new bytes[](0),
            submittedTasks_: new bytes[](0),
            proofRequests_: new bytes[](0),
            allocatedEth_: new uint256[](_operatorModels[0].ethShares_.length),
            allocatedSer_: 0,
            proofRequestExponents_: [uint32(1e3),uint32(1e3)],
            pausedBlock_: 0
        });

        for (uint8 i = 0; i < _operatorModels[0].ethShares_.length; i++) {
            _operator.allocatedEth_[i] = 0;
        }

        opInfo[operator] = abi.encode(_operator);
        //was this for lower gas costs?
        // opInfo[operator] = Operator({
        //     models_: _operatorModelIds,
        //     computeUnits_: _computeUnitNames,
        //     openTasks_: new bytes[](0),
        //     submittedTasks_: new bytes[](0),
        //     proofRequests_: new bytes[](0),
        //     allocatedEth_: new uint256[](_operatorModels[0].ethShares_.length),
        //     allocatedSer_: 0,
        //     proofRequestExponents_: [uint32(1e3),uint32(1e3)],
        //     pausedBlock_: 0
        // });

        // for (uint8 i = 0; i < _operatorModels[0].ethShares_.length; i++) {
        //     opInfo[operator].allocatedEth_[i] = 0;
        // }

        operators.push(operator);
        isOperator[operator] = true;

        emit NewOperator(operator);
        emit NewModels(_operatorModelIds);
    }

    function sendRewards(IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] calldata operatorDirectedRewardsSubmissions, bytes[][] memory _rewardedTasks) external onlyAggregators() {
        uint256 _approvalAmount = 0;
        for (uint8 i = 0; i < operatorDirectedRewardsSubmissions[0].operatorRewards.length; i ++) {
            Operator memory _operator = abi.decode(opInfo[operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator],(Operator));
            for (uint8 j = 0; j < _rewardedTasks[i].length; j ++) {
                Task memory _task = abi.decode(_rewardedTasks[i][j], (Task));
                if (!(_task.startingBlock_ + TASK_EXPIRY_BLOCKS < block.number) || sertnTaskManager.taskVerified(_rewardedTasks[i][j])) {
                    revert TaskNotExpired();
                }
                // require(_task.startingBlock_ + TASK_EXPIRY_BLOCKS < block.number || sertnTaskManager.taskVerified(_rewardedTasks[i][j]), "Task has not expired");
                if (!_inBytesArray(_operator.submittedTasks_, _rewardedTasks[i][j])) {
                    revert NotInSubmittedTasks();
                }
                // require(_inBytesArray(_operator.submittedTasks_, _rewardedTasks[i][j]), "Not in submitted Tasks");
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

    function modifyModelParameters(OperatorModel memory _operatorModel, uint256 _operatorModelId) external onlyOperators() {
        OperatorModel memory _oldOperatorModel = abi.decode(operatorModelInfo[_operatorModelId], (OperatorModel));
        // OperatorModel storage _oldModel = operatorModelInfo[_operatorModelId];
        if (_oldOperatorModel.operator_ != msg.sender) {
            revert IncorrectOperator();
        }
        Model memory _model = abi.decode(modelStorage.modelVerifiers(modelStorage.modelAddresses(_operatorModel.modelId_)), (Model));
        bool _registered;
        for (uint256 i = 0; i < _model.operators_.length; i ++) {
            if (_model.operators_[i] == msg.sender) {
                _registered = true;
            }
        }

        if (!_registered) {
            modelStorage.JoinOperatorList(_operatorModel.modelId_, msg.sender);
        }
        operatorModelInfo[_operatorModelId] = abi.encode(_operatorModel);
        //Assuming this saves on gas fees?
        
        emit ModelUpdated(_operatorModelId, _operatorModel);
    }

    function slashOperator(bytes memory _taskId, string memory _whySlashed) external onlyAggregators() {

        Task memory _task = abi.decode(_taskId, (Task));
        uint256 _operatorModelId = _task.operatorModelId_;
        if (0 > _operatorModelId || numOperatorModels < _operatorModelId) {
            revert NotModelId();
        }

        OperatorModel memory _operatorModel = abi.decode(operatorModelInfo[_operatorModelId], (OperatorModel));
        // OperatorModel memory _operatorModel = sertnServiceManager.getModelInfo(_operatorModelId);
        address[] memory _operator = new address[](1);
        _operator[0] = _operatorModel.operator_;
        //Should this be included?
        if (!(block.number - _task.startingBlock_ < TASK_EXPIRY_BLOCKS)) {
            revert TaskExpired();
        }
        // require(block.number - _task.startingBlock_ < TASK_EXPIRY_BLOCKS, "Task already expired");
        IStrategy[] memory _strategies = new IStrategy[](_operatorModel.ethStrategies_.length + 1);
        for (uint8 i = 0; i < _operatorModel.ethStrategies_.length; i++) {
            _strategies[i] = _operatorModel.ethStrategies_[i];
        }
        _strategies[_operatorModel.ethStrategies_.length] = tokenToStrategy[address(ser)];
        uint256[] memory _wadsToSlash = new uint256[](_strategies.length);
        uint256[] memory _operatorShares = delegationManager.getOperatorShares(_operatorModel.operator_, _strategies);
        for (uint8 i = 0; i < _operatorModel.ethStrategies_.length; i++) {
            _wadsToSlash[i] = 1 ether * 1 ether * (_operatorModel.ethShares_[i]) / (allocationManager.getAllocation(_operatorModel.operator_, opSet, _strategies[i]).currentMagnitude * _operatorShares[i]);
        }
        if (_task.proveOnResponse_) {
            _wadsToSlash[_operatorModel.ethStrategies_.length] = (1 ether * 1 ether * _task.poc_)/(allocationManager.getAllocation(_operatorModel.operator_, opSet, _strategies[_operatorModel.ethStrategies_.length]).currentMagnitude * _operatorShares[_operatorModel.ethStrategies_.length]);
        } else {
            _wadsToSlash[_operatorModel.ethStrategies_.length] = (1 ether * 1 ether * 10 * _task.poc_)/(allocationManager.getAllocation(_operatorModel.operator_, opSet, _strategies[_operatorModel.ethStrategies_.length]).currentMagnitude * _operatorShares[_operatorModel.ethStrategies_.length]);
        }
        IAllocationManagerTypes.SlashingParams memory _slashParams = IAllocationManagerTypes.SlashingParams({operator: _operatorModel.operator_, operatorSetId: 0, strategies: _strategies, wadsToSlash: _wadsToSlash, description: _whySlashed});
        allocationManager.slashOperator(address(this), _slashParams);

        emit OperatorSlashed(_operatorModel.operator_, _taskId);
        // TODO: Custom logic to change bounty amount
        if (bountyHunter[_taskId] != address(0)) {
            if (!ser.transferFrom(address(this), bountyHunter[_taskId], BOUNTY)) {
                revert TransferFailed();
            }
            // require(ser.transferFrom(address(this), bountyHunter[_taskId], BOUNTY), "Couldn't pay bounty hunter");
            Operator memory _operator = abi.decode(opInfo[_operatorModel.operator_], (Operator));
            _operator.proofRequestExponents_[0] = 1e3;
            opInfo[_operatorModel.operator_] = abi.encode(_operator);
        }
        sertnTaskManager.clearTask(_taskId, true);
    }

    function addModels(OperatorModel[] memory _operatorModels) external onlyOperators() {
        uint256[] memory _operatorModelIds = new uint256[](_operatorModels.length);
        for (uint256 i = 0; i < _operatorModels.length; i++) {
            uint256 operatorModelNum = numOperatorModels;
            numOperatorModels++;
            _operatorModelIds[i] = operatorModelNum;
            _operatorModels[i].operator_ = msg.sender;
            // require(_operatorModels[i].maxBlocks_ < TASK_EXPIRY_BLOCKS - 1e2, "max blocks too long");
            if (!(_operatorModels[i].maxBlocks_ < TASK_EXPIRY_BLOCKS - 1e2)) {
                revert MaxBlocksTooLong();
            }

            Model memory _model = abi.decode(modelStorage.modelVerifiers(modelStorage.modelAddresses(_operatorModels[i].modelId_)), (Model));
            bool _registered;
            for (uint256 j = 0; j < _model.operators_.length; j ++) {
                if (_model.operators_[j] == msg.sender) {
                    _registered = true;
                }
            }
            if (!_registered) {
                revert NotRegisteredToModel();
            }
            operatorModelInfo[operatorModelNum] = abi.encode(_operatorModels[i]);
            // operatorModelInfo[modelNum] = _operatorModels[i];
            Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
            _tempOpModels = _operator.models_;
            _tempOpModels.push(operatorModelNum);
            _operator.models_ = _tempOpModels;
            opInfo[msg.sender] = abi.encode(_operator);
            // opInfo[msg.sender].models_.push(uint8(modelNum));
        }
        emit NewModels(_operatorModelIds);
    }

    function modifyCompute(bytes32[] memory _computeUnitNames, uint8[] memory _computeUnits) external onlyOperators() {
        for (uint8 i = 0; i < _computeUnitNames.length; i++) {
            computeUnits[msg.sender][_computeUnitNames[i]] = _computeUnits[i];
        }

        Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));
        delete _operator.computeUnits_;
        // delete opInfo[msg.sender].computeUnits_;
        for (uint8 i = 0; i < _computeUnitNames.length; i++) {
            _tempComputeUnitNames = _operator.computeUnits_;
            _tempComputeUnitNames.push(_computeUnitNames[i]);
            _operator.computeUnits_ = _tempComputeUnitNames;
        }
        opInfo[msg.sender] = abi.encode(_operator);
        emit OpInfoChanged(msg.sender, opInfo[msg.sender]);
    }

    function updateOperator(address _operator, Operator memory _opInfo) external onlyAggregators() {
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

        // require(isAggregator[msg.sender] || msg.sender == operator, "No permission to call");
        if (!(isAggregator[msg.sender] || msg.sender == operator)) {
            revert NoPermission();
        }
        if (address(this) != avs) {
            revert IncorrectAVS();
        }
        Operator memory _operator = abi.decode(opInfo[operator], (Operator));
        if (isAggregator[msg.sender] || uint256(_operator.pausedBlock_) + 2*TASK_EXPIRY_BLOCKS < block.number || _operator.pausedBlock_ > 0) {
            revert NotPausedLongEnough();
        }
        for (uint8 i = 0; i < _operator.models_.length; i ++) {
            modelStorage.RemoveFromOperatorList(abi.decode(operatorModelInfo[_operator.models_[i]],(OperatorModel)).modelId_, operator);
            delete operatorModelInfo[_operator.models_[i]];

        }
        delete opInfo[operator];
        for (uint8 i = 0; i < operators.length; i ++) {
            if (operators[i] == operator) {
                delete operators[i];
            }
        }
        emit OperatorDeleted(operator, operatorSetIds);
    }
}
