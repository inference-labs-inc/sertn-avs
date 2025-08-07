// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

contract MockSertnTaskManager {
    address public sertnServiceManager;

    function setServiceManager(address _serviceManager) external {
        sertnServiceManager = _serviceManager;
    }
}
