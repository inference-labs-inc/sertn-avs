// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {IVerifier} from "../../interfaces/IVerifier.sol";

contract MockVerifier is IVerifier {
    function verifyProof(bytes memory _proof) public pure returns (bool) {
        if (keccak256(bytes("1")) == keccak256(_proof)) {
            return true;
        }
        return false;
    }
}
