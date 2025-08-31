// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;
import {OwnableUpgradeable} from "@openzeppelin-upgradeable/contracts/access/OwnableUpgradeable.sol";
import {IModelRegistry} from "../interfaces/IModelRegistry.sol";

/**
 * @title ModelRegistry
 * @author Inference Labs, Inc.
 * @notice ModelRegistry is a contract that stores models for the Sertn network.
 */
contract ModelRegistry is OwnableUpgradeable, IModelRegistry {
    // Current model index (starts at 1 to avoid zero-indexing issues)
    uint256 public modelIndex = 1;

    // modelVerifier => modelId
    mapping(address => uint256) public modelVerifiers;

    // modelId => modelVerifier
    mapping(uint256 => address) public modelVerifier;
    // modelId => modelName
    mapping(uint256 => string) public modelName;
    // modelName => modelId (enforce uniqueness of modelName values)
    mapping(string => uint256) public modelNameToId;
    // modelId => verificationStrategy
    mapping(uint256 => VerificationStrategy) public verificationStrategy;
    // modelId => computeCost
    mapping(uint256 => uint256) public computeCost;
    // modelId => requiredFUCUs
    mapping(uint256 => uint256) public requiredFUCUs;

    function initialize() public initializer {
        __Ownable_init();
        modelIndex = 1; // Ensure we start at 1
    }

    /// @inheritdoc IModelRegistry
    function createNewModel(
        address _modelVerifier,
        VerificationStrategy _verificationStrategy,
        string memory _modelName,
        uint256 _computeCost,
        uint256 _requiredFUCUs
    ) external onlyOwner returns (uint256 modelId) {
        // Validation checks for model verifier
        if (_modelVerifier == address(0)) revert InvalidModelVerifier();
        if (modelVerifiers[_modelVerifier] != 0)
            revert ModelAlreadyExists(modelVerifiers[_modelVerifier]);
        // Ensure modelName is unique
        if (modelNameToId[_modelName] != 0) revert ModelAlreadyExists(modelNameToId[_modelName]);
        // Assign verifier and Name
        modelVerifier[modelIndex] = _modelVerifier;
        modelName[modelIndex] = _modelName;
        modelNameToId[_modelName] = modelIndex;
        verificationStrategy[modelIndex] = _verificationStrategy;
        computeCost[modelIndex] = _computeCost;
        requiredFUCUs[modelIndex] = _requiredFUCUs;
        // Assign reverse mapping values
        modelVerifiers[_modelVerifier] = modelIndex;

        emit ModelCreated(
            modelIndex,
            _modelVerifier,
            _verificationStrategy,
            _modelName,
            _computeCost,
            _requiredFUCUs
        );
        modelId = modelIndex;
        // Increment model index
        modelIndex++;
    }

    /// @inheritdoc IModelRegistry
    function updateModelName(uint256 modelId, string memory _modelName) external onlyOwner {
        if (modelId >= modelIndex) revert ModelDoesNotExist();
        // Enforce uniqueness: allow same modelId to keep its own name
        uint256 existing = modelNameToId[_modelName];
        if (existing != 0 && existing != modelId) revert ModelAlreadyExists(existing);
        // Update reverse mapping: remove old name mapping
        string memory oldName = modelName[modelId];
        if (bytes(oldName).length != 0) {
            // Clear only if previously set
            delete modelNameToId[oldName];
        }
        modelName[modelId] = _modelName;
        modelNameToId[_modelName] = modelId;
        emit ModelNameUpdated(modelId, _modelName);
    }

    /// @inheritdoc IModelRegistry
    function updateComputeCost(uint256 modelId, uint256 _computeCost) external onlyOwner {
        if (modelId >= modelIndex) revert ModelDoesNotExist();
        computeCost[modelId] = _computeCost;
        emit ComputeCostUpdated(modelId, _computeCost);
    }

    /// @inheritdoc IModelRegistry
    function updateRequiredFUCUs(uint256 modelId, uint256 fucus) external onlyOwner {
        if (modelId >= modelIndex) revert ModelDoesNotExist();
        requiredFUCUs[modelId] = fucus;
        emit RequiredFUCUsUpdated(modelId, fucus);
    }

    /// @inheritdoc IModelRegistry
    function updateModelVerifier(uint256 modelId, address _modelVerifier) external onlyOwner {
        if (modelId >= modelIndex) revert ModelDoesNotExist();
        if (_modelVerifier == address(0)) revert InvalidModelVerifier();
        if (modelVerifiers[_modelVerifier] != 0)
            revert ModelAlreadyExists(modelVerifiers[_modelVerifier]);
        // Update the verifier
        address oldVerifier = modelVerifier[modelId];
        if (oldVerifier != address(0)) {
            delete modelVerifiers[oldVerifier];
        }
        modelVerifiers[_modelVerifier] = modelId;
        modelVerifier[modelId] = _modelVerifier;
        emit ModelVerifierUpdated(modelId, _modelVerifier);
    }

    /// @inheritdoc IModelRegistry
    function updateVerificationStrategy(
        uint256 modelId,
        VerificationStrategy _verificationStrategy
    ) external onlyOwner {
        if (modelId >= modelIndex) revert ModelDoesNotExist();
        verificationStrategy[modelId] = _verificationStrategy;
        emit VerificationStrategyUpdated(modelId, _verificationStrategy);
    }
}
