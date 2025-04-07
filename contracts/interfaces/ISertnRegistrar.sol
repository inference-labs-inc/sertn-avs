// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";

interface ISertnRegistrar is IAVSRegistrar {
    address public avsAddress;
}
