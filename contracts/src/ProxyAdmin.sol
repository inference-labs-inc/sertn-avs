// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgradeable/contracts/access/OwnableUpgradeable.sol";

contract ProxyAdmin is OwnableUpgradeable {
    function upgrade(
        ITransparentUpgradeableProxy proxy,
        address implementation
    ) external onlyOwner {
        proxy.upgradeToAndCall(implementation, "");
    }
}
