// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {ISertnAggregator} from "../interfaces/ISertnAggregator.sol";
import {ISertnTaskManager} from "../interfaces/ISertnTaskManager.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

/**
 * @title Sertn Aggregator
 * @author Inference Labs, Inc.
 * @notice Aggregator for Sertn tasks.
 */
contract SertnAggregator is OwnableUpgradeable, ISertnAggregator {
    address public aggregatorEOA;
    ISertnTaskManager public sertnTaskManager;

    function initialize(
        address _aggregatorEOA,
        address _sertnTaskManager
    ) public initializer {
        __Ownable_init();
        aggregatorEOA = _aggregatorEOA;
        sertnTaskManager = ISertnTaskManager(_sertnTaskManager);
    }

    function updateAggregatorEOA(address _aggregatorEOA) external onlyOwner {
        if (_aggregatorEOA == address(0)) {
            revert ZeroAddress();
        }
        aggregatorEOA = _aggregatorEOA;
    }

    function updateSertnTaskManager(
        address _sertnTaskManager
    ) external onlyOwner {
        if (_sertnTaskManager == address(0)) {
            revert ZeroAddress();
        }
        sertnTaskManager = ISertnTaskManager(_sertnTaskManager);
    }

    function submitTask(
        ISertnTaskManager.Task memory _task,
        bytes memory _proof
    ) external {
        bytes32 messageHash = keccak256(abi.encode(_task));
        address signer = ECDSA.recover(
            MessageHashUtils.toEthSignedMessageHash(messageHash),
            _proof
        );

        if (signer != aggregatorEOA) {
            revert InvalidEOASignature();
        }

        sertnTaskManager.sendTask(_task);
        emit TaskSubmitted(messageHash, msg.sender);
    }
}
