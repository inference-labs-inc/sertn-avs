// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ISertnAggregator} from "../interfaces/ISertnAggregator.sol";
import {ISertnTaskManager} from "../interfaces/ISertnTaskManager.sol";

contract SertnAggregator is OwnableUpgradeable, ISertnAggregator {
    address public aggregatorEOA;
    ISertnTaskManager public sertnTaskManager;

    function initialize(
        address _aggregatorEOA,
        address _sertnTaskManager
    ) public initializer {
        __Ownable_init(msg.sender);
        aggregatorEOA = _aggregatorEOA;
        sertnTaskManager = ISertnTaskManager(_sertnTaskManager);
    }

    function updateAggregatorEOA(address _aggregatorEOA) external onlyOwner {
        if (_aggregatorEOA == address(0)) {
            revert InvalidEOA();
        }
        aggregatorEOA = _aggregatorEOA;
    }

    function updateSertnTaskManager(
        address _sertnTaskManager
    ) external onlyOwner {
        if (_sertnTaskManager == address(0)) {
            revert InvalidTaskManager();
        }
        sertnTaskManager = ISertnTaskManager(_sertnTaskManager);
    }

    function submitTask(bytes32 _taskId, bytes memory _proof) external {
        bytes32 messageHash = keccak256(abi.encodePacked(_taskId));
        address signer = recoverSigner(messageHash, _proof);

        if (signer != aggregatorEOA) {
            revert InvalidEOASignature();
        }

        emit TaskSubmitted(_taskId, msg.sender);
    }

    function recoverSigner(
        bytes32 _messageHash,
        bytes memory _signature
    ) internal pure returns (address) {
        bytes32 ethSignedMessageHash = keccak256(
            abi.encodePacked("\x19Ethereum Signed Message:\n32", _messageHash)
        );
        (bytes32 r, bytes32 s, uint8 v) = splitSignature(_signature);
        return ecrecover(ethSignedMessageHash, v, r, s);
    }

    function splitSignature(
        bytes memory _signature
    ) internal pure returns (bytes32 r, bytes32 s, uint8 v) {
        require(_signature.length == 65, "Invalid signature length");
        assembly {
            r := mload(add(_signature, 32))
            s := mload(add(_signature, 64))
            v := byte(0, mload(add(_signature, 96)))
        }
    }
}
