// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {SertnServiceManager} from "./SertnServiceManager.sol";
import "./ISertnServiceManager.sol";
import "./SertnServiceManagerStorage.sol";

contract ModelStorage is ISertnServiceManagerTypes, ISertnServiceManagerErrors, OwnableUpgradeable, SertnServiceManagerStorage {

    modifier onlyOperators() {
        if (!(sertnServiceManager.isOperator(msg.sender)|| msg.sender == address(sertnServiceManager))) {
            revert NotOperator();
        }
        _;
    }

    uint256 public numModels;

    SertnServiceManager public sertnServiceManager;

    constructor(
        address _sertnServiceManager
    ) OwnableUpgradeable() {
        sertnServiceManager = SertnServiceManager(_sertnServiceManager);
    }

    function createNewModel(
        Model calldata _newModel
    ) external onlyOperators() returns(uint256) {
        // Model memory newModel = Model({description_: description, title_: title, modelVerifier_: modelVerifierAddress, operators_: operators});
        // newModel.description = description;
        // newModel.title = title;
        // newModel.modelVerifier = modelVerifierAddress;
        

        modelVerifiers[_newModel.modelVerifier_] = abi.encode(_newModel);
        modelAddresses[numModels] = _newModel.modelVerifier_;


        numModels++;
        return numModels - 1;
    }

    function JoinOperatorList(uint256 _modelId, address _operator) external onlyOperators() {
        Model memory model = abi.decode(modelVerifiers[modelAddresses[_modelId]],(Model));

        address[] memory newOperators = new address[](model.operators_.length + 1);

        for (uint256 i; i < model.operators_.length;) {
            newOperators[i] = model.operators_[i];
            unchecked { ++i; }
        }
        if (msg.sender != _operator || msg.sender != address(sertnServiceManager)) {
            revert NoPermission();
        }
        newOperators[model.operators_.length] = _operator;
        model.operators_ = newOperators;
        modelVerifiers[modelAddresses[_modelId]] = abi.encode(model);

    }

    function RemoveFromOperatorList(uint256 _modelId, address _operator) external {
        if (msg.sender != address(sertnServiceManager)) {
            revert NoPermission();
        }

        Model memory model = abi.decode(modelVerifiers[modelAddresses[_modelId]],(Model));

        for (uint256 i = 0; i < model.operators_.length;) {
            if (model.operators_[i] == _operator) {
                delete model.operators_[i];
            }
            unchecked { ++i; }
        }
        
        modelVerifiers[modelAddresses[_modelId]] = abi.encode(model);


    }
}