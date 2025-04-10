// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {TransparentUpgradeableProxy} from "@openzeppelin/lib/openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract ModelRegistryProxy is TransparentUpgradeableProxy {
    constructor(
        address _logic,
        address _admin,
        bytes memory _data
    ) TransparentUpgradeableProxy(_logic, _admin, _data) {}
}
