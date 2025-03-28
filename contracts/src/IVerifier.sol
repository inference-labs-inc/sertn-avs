// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IVerifier {

    function verifyProof(bytes memory _proof, bytes32 _modelKey) external view returns (bool);

}