// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IZKVerifier {
    function verifyProof(
        bytes calldata proof,
        uint256[] calldata instances
    ) external returns (bool);
}
