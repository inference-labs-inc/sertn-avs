// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

/**
 * @title SertnRegistrar
 * @author Inference Labs, Inc.
 * @notice SertnRegistrar is a contract that allows operators to register to operate on the Sertn network.
 */
contract SertnRegistrar is IAVSRegistrar, OwnableUpgradeable {
    // The address of the AVS contract (SertnServiceManager)
    address public avsAddress;

    /**
     * @notice Initialize the contract
     * @param _avsAddress The address of the AVS contract
     */
    function initialize(address _avsAddress) public initializer {
        __Ownable_init();
        avsAddress = _avsAddress;
    }

    /**
     * @inheritdoc IAVSRegistrar
     */
    function registerOperator(
        address operator,
        uint32[] calldata operatorSetIds,
        bytes calldata data
    ) external {}

    /**
     * @inheritdoc IAVSRegistrar
     */
    function deregisterOperator(
        address operator,
        uint32[] calldata operatorSetIds
    ) external {}

    /**
     * @inheritdoc IAVSRegistrar
     */
    function supportsAVS(address avs) external view returns (bool) {
        return avs == avsAddress;
    }
}
