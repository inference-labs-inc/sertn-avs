// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {ITransparentUpgradeableProxy} from "@openzeppelin/lib/openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts/access/OwnableUpgradeable.sol";

contract ProxyAdmin is OwnableUpgradeable {
    function upgrade(
        ITransparentUpgradeableProxy proxy,
        address implementation
    ) external onlyOwner {
        proxy.upgradeToAndCall(implementation, "");
    }
}
