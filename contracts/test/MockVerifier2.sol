// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IVerifier} from "../interfaces/IVerifier.sol";

contract MockVerifier2 is IVerifier {
    function verifyProof(bytes memory _proof) public view returns (bool) {
        if (keccak256(bytes("2")) == keccak256(_proof)) {
            return true;
        }
        return false;
    }
}
