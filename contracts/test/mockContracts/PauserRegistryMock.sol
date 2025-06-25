// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

import "@eigenlayer/contracts/permissions/Pausable.sol";

contract MockPauserRegistry {
    function isPauser(address) external pure returns (bool) {
        return false;
    }

    function unregisterAsOperator() external pure {
        // Empty implementation
    }
}
