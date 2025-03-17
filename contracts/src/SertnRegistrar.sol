// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";

contract SertnRegistrar is IAVSRegistrar {
    error WrongAVS();

    address sertnAVS;

    constructor() {
        sertnAVS = msg.sender;
    }

    function registerOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds,
        bytes calldata data
    ) external {
        
    }

    function deregisterOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds
    ) external override {}

    function supportsAVS(address avs) external view override returns (bool) {
        require(avs == sertnAVS, WrongAVS());
        return true;
    }
}