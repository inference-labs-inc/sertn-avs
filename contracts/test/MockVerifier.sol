// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IVerifier} from "../src/IVerifier.sol";

contract MockVerifier is IVerifier {
    function verifyProof(bytes memory _proof, bytes32 _modelKey) public view returns (bool) {
        if (keccak256(bytes("1")) == keccak256(_proof)) {
            return true;
        }
        return false;
    }
}