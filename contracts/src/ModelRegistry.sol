// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {IModelRegistry} from "../interfaces/IModelRegistry.sol";

/**
 * @title ModelRegistry
 * @author Inference Labs, Inc.
 * @notice ModelRegistry is a contract that stores models for the Sertn network.
 */
contract ModelRegistry is OwnableUpgradeable, IModelRegistry {
    // Current model index
    uint256 public modelIndex;

    // modelVerifier => modelId
    mapping(address => uint256) public modelVerifiers;

    // modelId => modelVerifier
    mapping(uint256 => address) public modelVerifier;
    // modelId => modelURI
    mapping(uint256 => string) public modelURI;
    // modelId => verificationStrategy
    mapping(uint256 => VerificationStrategy) public verificationStrategy;
    // modelId => computeCost
    mapping(uint256 => uint256) public computeCost;

    function initialize() public initializer {
        __Ownable_init();
    }

    /// @inheritdoc IModelRegistry
    function createNewModel(
        address _modelVerifier,
        VerificationStrategy _verificationStrategy,
        string memory _modelURI,
        uint256 _computeCost
    ) external onlyOwner {
        // Validation checks for model verifier
        if (_modelVerifier == address(0)) revert InvalidModelVerifier();
        if (modelVerifiers[_modelVerifier] != 0)
            revert ModelAlreadyExists(modelVerifiers[_modelVerifier]);
        // Assign verifier and URI
        modelVerifier[modelIndex] = _modelVerifier;
        modelURI[modelIndex] = _modelURI;
        verificationStrategy[modelIndex] = _verificationStrategy;
        computeCost[modelIndex] = _computeCost;
        // Assign reverse mapping values
        modelVerifiers[_modelVerifier] = modelIndex;

        emit ModelCreated(
            modelIndex,
            _modelVerifier,
            _verificationStrategy,
            _modelURI,
            _computeCost
        );
        // Increment model index
        modelIndex++;
    }

    /// @inheritdoc IModelRegistry
    function updateModelURI(
        uint256 modelId,
        string memory _modelURI
    ) external onlyOwner {
        if (modelId >= modelIndex) revert ModelDoesNotExist();
        modelURI[modelId] = _modelURI;
        emit ModelURIUpdated(modelId, _modelURI);
    }

    /// @inheritdoc IModelRegistry
    function updateComputeCost(
        uint256 modelId,
        uint256 _computeCost
    ) external onlyOwner {
        if (modelId >= modelIndex) revert ModelDoesNotExist();
        computeCost[modelId] = _computeCost;
        emit ComputeCostUpdated(modelId, _computeCost);
    }
}
