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

    address public avsAddress;
  function initialize(address _avsAddress) public initializer {
    __Ownable_init();
    avsAddress = _avsAddress;
  }


    function registerOperator(address operator, address avs, uint32[] calldata operatorSetIds, bytes calldata data) external {

    }

    function deregisterOperator(address operator, address avs, uint32[] calldata operatorSetIds) external {

    }

    function supportsAVS(address avs) external view returns (bool) {
        return avs == avsAddress;
    }


}
