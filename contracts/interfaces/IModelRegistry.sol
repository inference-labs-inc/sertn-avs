// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

interface IModelRegistry {

  event ModelCreated(uint256 indexed modelId, address indexed modelVerifier, string indexed modelURI);
  event ModelUpdated(uint256 indexed modelId, string indexed modelURI);

  error ModelAlreadyExists(uint256 modelId);
  error InvalidModelVerifier();
  error ModelDoesNotExist();

  function createNewModel(address modelVerifier, string memory modelURI) external;
  function updateModelURI(uint256 modelId, string memory modelURI) external;
}
