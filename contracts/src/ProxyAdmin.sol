// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
contract ProxyAdmin is OwnableUpgradeable {

    function upgrade(TransparentUpgradeableProxy proxy, address implementation) external onlyOwner {
        proxy.upgradeTo(implementation);
    }
}
