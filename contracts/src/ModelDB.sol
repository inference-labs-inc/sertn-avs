// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;
import "@openzeppelin/contracts/access/Ownable.sol";

contract ModelDB {
    struct Model {
        string title;
        string description;
        address modelVerifier;
    }

    uint256 public currentModelId = 0;

    mapping(uint256 => address) public modelAddresses;
    mapping(address => Model) public modelVerifiers;

    function createNewModel(
        address modelVerifierAddress,
        string calldata title,
        string calldata description
    ) public {
        Model memory newModel;
        newModel.description = description;
        newModel.title = title;
        newModel.modelVerifier = modelVerifierAddress;

        modelVerifiers[modelVerifierAddress] = newModel;
        modelAddresses[currentModelId] = modelVerifierAddress;

        currentModelId++;
    }
}
