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
    bytes32[] public slashingQueue;

    IStrategy[] public strategies;

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

        strategies = _strategies;

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
            revert IncorrectAVS();
        }
        if (operatorSetIds.length != 1 || operatorSetIds[0] != 0) {
            revert IncorrectOperatorSetIds();
        }
        
        
        Operator memory _operator = Operator({
            nodesAndModels_: new uint256[][](0),
            nodeComputeUnits_: new uint256[](0), 
            submittedTasks_: new bytes32[](0),
            allocatedEth_: new uint256[](strategies.length - 1),
            allocatedSer_: 0,
            proofRequestCoefficients_: new uint256[](0),
            pausedBlock_: 0
        });

        opInfo[operator] = abi.encode(_operator);

        operators.push(operator);
        isOperator[operator] = true;

        emit NewOperator(operator);
    }

    function sendRewards(IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] calldata operatorDirectedRewardsSubmissions, bytes[][] calldata _rewardedTasks) external onlyAggregators() {
        uint256 _approvalAmount;
        for (uint256 i; i < operatorDirectedRewardsSubmissions[0].operatorRewards.length;) {
            Operator memory _operator = abi.decode(opInfo[operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator],(Operator));
            for (uint256 j; j < _rewardedTasks[i].length;) {
                Task memory _task = abi.decode(_rewardedTasks[i][j], (Task));
                if (!(_task.startingBlock_ + taskExpiryBlocks < uint32(block.number)) || sertnTaskManager.taskVerified(keccak256(_rewardedTasks[i][j]))) {
                    revert TaskNotExpired();
                }
                if (!_inBytes32Array(_operator.submittedTasks_, keccak256(_rewardedTasks[i][j]))) {
                    revert NotInSubmittedTasks();
                }
                _operator.submittedTasks_ = _removeBytes32Element(_operator.submittedTasks_, keccak256(_rewardedTasks[i][j]));
                opInfo[operatorDirectedRewardsSubmissions[0].operatorRewards[i].operator] = abi.encode(_operator);
                sertnTaskManager.clearTask(_rewardedTasks[i][j], false);
                unchecked { ++j; }
            }

            _approvalAmount += operatorDirectedRewardsSubmissions[0].operatorRewards[i].amount;
            unchecked { ++i; }
        }

        ser.approve(address(rewardsCoordinator), _approvalAmount);
        rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(address(this), operatorDirectedRewardsSubmissions);
    }

    function _removeBytesElement(bytes[] memory _byteArray, bytes calldata _byteElement) internal pure returns(bytes[] memory) {

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
    function _inBytesArray(bytes[] memory _byteArray, bytes calldata _byteElement) internal pure returns(bool) {
        bytes32 elementHash = keccak256(_byteElement);
        for (uint256 i; i < _byteArray.length;) {
            if (elementHash == keccak256(_byteArray[i])) {
                return true;
            }
            unchecked { ++i; }
        }
        return false;
    }

    //Want to have the option to deregister a model rather than just pause?
    function updateNodeModels(NodeModel calldata _nodeModel, uint256 _nodeId, uint256 _modelId) external onlyOperators() {
        bytes32 _nodeModelId = keccak256(abi.encodePacked(msg.sender,_nodeId,_modelId));
        if (nodeModelInfo[_nodeModelId].length==0) {
            revert NotRegisteredToModel();
        }

        nodeModelInfo[_nodeModelId] = abi.encode(_nodeModel);
        emit ModelUpdated(_nodeModelId, _nodeModel);
    }

    function slashOperator(Task calldata _task, string calldata _whySlashed, OperatorSet memory _opSet) external onlyAggregators() nonReentrant {

        NodeModel memory _nodeModel = abi.decode(nodeModelInfo[keccak256(abi.encodePacked(_task.operator_,_task.node_,_task.modelId_))], (NodeModel));
        //Should this be included?
        if (!(uint32(block.number) - _task.startingBlock_ < taskExpiryBlocks)) {
            revert TaskExpired();
        }
       
        uint256[] memory _wadsToSlash = new uint256[](strategies.length);
        uint256[] memory _operatorShares = delegationManager.getOperatorShares(_task.operator_, strategies);
        uint256 normalizingConstant = 1e36;
        uint256 amountToSlash;
        uint256 slashableShares;

        // Wads to slash = normalizingConstant*amountToSlash/slashableShares
        if (_task.proveOnResponse_) {
            // _wadsToSlash[0] = MathUpgradeable.mulDiv(1 ether * 1 ether, _task.poc_, allocationManager.getAllocation(_task.operator_, _opSet, strategies[0]).currentMagnitude * _operatorShares[0]);
            amountToSlash = _task.poc_;
            slashableShares = allocationManager.getAllocation(_task.operator_, _opSet, strategies[0]).currentMagnitude * _operatorShares[0];
            _wadsToSlash[0] = MathUpgradeable.mulDiv(normalizingConstant, amountToSlash, slashableShares);
        } else {
            // _wadsToSlash[0] = MathUpgradeable.mulDiv(1 ether * 1 ether, 10*_task.poc_, allocationManager.getAllocation(_task.operator_, _opSet, strategies[0]).currentMagnitude * _operatorShares[0]);
            amountToSlash = 10*_task.poc_;
            slashableShares = allocationManager.getAllocation(_task.operator_, _opSet, strategies[0]).currentMagnitude * _operatorShares[0];
            _wadsToSlash[0] = MathUpgradeable.mulDiv(normalizingConstant, amountToSlash, slashableShares);
        }
        for (uint256 i = 1; i < strategies.length;) {
            // _wadsToSlash[i] = MathUpgradeable.mulDiv(1 ether * 1 ether , _nodeModel.ethShares_[i],(allocationManager.getAllocation(_task.operator_, _opSet, strategies[i]).currentMagnitude * _operatorShares[i]));
            amountToSlash = _nodeModel.ethShares_[i];
            slashableShares = allocationManager.getAllocation(_task.operator_, _opSet, strategies[i]).currentMagnitude * _operatorShares[i];
            _wadsToSlash[i] = MathUpgradeable.mulDiv(normalizingConstant, amountToSlash, slashableShares);
            unchecked { ++i; }
        }
        
        IAllocationManagerTypes.SlashingParams memory _slashParams = IAllocationManagerTypes.SlashingParams({operator: _task.operator_, operatorSetId: _opSet.id, strategies: strategies, wadsToSlash: _wadsToSlash, description: _whySlashed});
        allocationManager.slashOperator(address(this), _slashParams);

        bytes32 _taskId = keccak256(abi.encode(_task));
        emit OperatorSlashed(_taskId);
        // TODO: Custom logic to change bounty amount
        if (bountyHunter[_taskId] != address(0)) {
            if (!ser.transferFrom(address(this), bountyHunter[_taskId], BOUNTY)) {
                revert TransferFailed();
            }
            Operator memory _operator = abi.decode(opInfo[_task.operator_], (Operator));
            _operator.proofRequestCoefficients_[_task.node_] = startingCoefficients;
            opInfo[_task.operator_] = abi.encode(_operator);
        }
        sertnTaskManager.clearTask(abi.encode(_task), true);
    }

    function addNodeModels(NodeModel[][] memory _nodeModels, uint256[] memory _nodeIds, uint256[][] memory _nodeAndModelIds) external onlyOperators() {
        Operator memory _operator = abi.decode(opInfo[msg.sender], (Operator));

        uint256[][] memory _tempNodesAndModels;
        uint256[] memory _tempProofRequestCoefficients;
        uint256 _numNewNodes;
        for (uint256 i; i < _nodeIds.length; ) {
            if (_nodeIds[i] + 1 > _operator.nodesAndModels_.length) {
                _nodeIds[i] = _operator.nodesAndModels_.length + _numNewNodes;
                _numNewNodes ++;
            }
            unchecked {++i;}
        }

        _tempProofRequestCoefficients = new uint256[](_operator.proofRequestCoefficients_.length + 2*_numNewNodes);
        _tempNodesAndModels = new uint256[][]( _operator.nodesAndModels_.length + _numNewNodes);

        for (uint256 i; i < _operator.proofRequestCoefficients_.length; ) {
            _tempProofRequestCoefficients[i] = _operator.proofRequestCoefficients_[i];
            unchecked {++i;}
        }
        for (uint256 i; i < 2*_numNewNodes; ) {
            _tempProofRequestCoefficients[i] = startingCoefficients;
            unchecked {++i;}
        }
            
        for (uint256 i; i < _nodeIds.length; ) {

            _tempNodesAndModels[_nodeIds[i]]  = new uint256[](_operator.nodesAndModels_[_nodeIds[i]].length + _nodeAndModelIds[i].length);
            for (uint256 j; j < _operator.nodesAndModels_[_nodeIds[i]].length; ) {
                    _tempNodesAndModels[_nodeIds[i]][j] = _operator.nodesAndModels_[_nodeIds[i]][j];
                    unchecked {++j;}
                }
            bytes32 _nodeModelId;
            for (uint256 j; j < _nodeAndModelIds[i].length; ) {
                _nodeModelId = keccak256(abi.encodePacked(msg.sender,_nodeIds[i],_nodeAndModelIds[i][j]));  
                if (nodeModelInfo[_nodeModelId].length!=0) {
                    revert NodeModelExists();
                }
                if (!modelStorage.operatorRegistered(keccak256(abi.encodePacked(msg.sender, _nodeAndModelIds[i][j])))) {
                    revert NotRegisteredToModel();
                }
                if (_nodeModels[i][j].maxBlocks_ + taskExpiryBuffer > taskExpiryBlocks) {
                    revert MaxBlocksTooLong();
                }
                _tempNodesAndModels[_nodeIds[i]][j + _operator.nodesAndModels_[_nodeIds[i]].length] = _nodeAndModelIds[i][j];
                nodeModelInfo[_nodeModelId] = abi.encode(_nodeModels[i][j]);
                unchecked {++j;} 
            }

            unchecked { ++i; }
        }
        _operator.nodesAndModels_ = _tempNodesAndModels;
        
        opInfo[msg.sender] = abi.encode(_operator);
        
        emit NewModels(_nodeAndModelIds);
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

    function pushToSlashingQueue(bytes32 _taskId) external onlyTaskManager() {
        slashingQueue.push(_taskId);
    }

    function pushToOperatorSlashingQueue(address _operator, bytes32 _taskId) external onlyTaskManager() {
        operatorSlashingQueue[_operator].push(_taskId);
    }

    function removeFromOperatorSlashingQueue(address _operator, bytes32 _taskId) external onlyTaskManager() {
        operatorSlashingQueue[_operator] = _removeBytes32Element(operatorSlashingQueue[_operator], _taskId);
    }

    function removeFromSlashingQueue(bytes32 _taskId) external onlyTaskManager() {
        slashingQueue = _removeBytes32Element(slashingQueue, _taskId);
    }

    function setTaskResponse(TaskResponse calldata _taskResponse) external onlyTaskManager() {
        taskResponse[keccak256(_taskResponse.taskId_)] = abi.encode(_taskResponse);
    }

    function setBountyHunter(bytes32 _taskId, address _bountyHunter) external onlyTaskManager() {
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
        if (!(isAggregator[msg.sender] || msg.sender == operator) || !isOperator[msg.sender]) {
            revert NoPermission();
        }
        Operator memory _operator = abi.decode(opInfo[operator], (Operator));
        if (!isAggregator[msg.sender] || _operator.pausedBlock_ + 2*taskExpiryBlocks > uint32(block.number)) {
            revert NotPausedLongEnough();
        }
        for (uint256 i; i < _operator.nodesAndModels_.length;) {
            for (uint256 j; j < _operator.nodesAndModels_[i].length; ) {
                modelStorage.RemoveFromOperatorList(_operator.nodesAndModels_[i][j], operator);
                unchecked { ++j; }
                delete nodeModelInfo[keccak256(abi.encodePacked(msg.sender,i,_operator.nodesAndModels_[i][j]))];
            }
            
            unchecked { ++i; }
        }

        delete opInfo[operator];
        bool pastOpIndex;
        for (uint256 i; i < operators.length -1;) {
            if (operators[i] == operator) {
                pastOpIndex = true;
            }
            if (pastOpIndex) {
                operators[i] = operators[i+1];
            }
            unchecked { ++i; }
        }
        if (pastOpIndex) {
            operators.pop();
        }
        isOperator[operator] = false;
        emit OperatorDeleted(operator, operatorSetIds);
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

}
