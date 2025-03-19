// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Vm} from "forge-std/Vm.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {EmptyContract} from "@eigenlayer/test/mocks/EmptyContract.sol";

library UpgradeableProxyLib {
    bytes32 internal constant IMPLEMENTATION_SLOT =
        0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
    bytes32 internal constant ADMIN_SLOT =
        0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;

    Vm internal constant vm =
        Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    function deployProxyAdmin() internal returns (address) {
        return address(new ProxyAdmin());
    }

    function setUpEmptyProxy(address admin) internal returns (address) {
        address emptyContract = address(new EmptyContract());
        return
            address(new TransparentUpgradeableProxy(emptyContract, admin, ""));
    }

    function upgrade(address proxy, address impl) internal {
        ProxyAdmin admin = getProxyAdmin(proxy);
        admin.upgrade(TransparentUpgradeableProxy(payable(proxy)), impl);
    }

    function upgradeAndCall(
        address proxy,
        address impl,
        bytes memory initData
    ) internal {
        ProxyAdmin admin = getProxyAdmin(proxy);
        admin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(proxy)),
            impl,
            initData
        );
    }

    function getImplementation(address proxy) internal view returns (address) {
        bytes32 value = vm.load(proxy, IMPLEMENTATION_SLOT);
        return address(uint160(uint256(value)));
    }

    function getProxyAdmin(address proxy) internal view returns (ProxyAdmin) {
        bytes32 value = vm.load(proxy, ADMIN_SLOT);
        return ProxyAdmin(address(uint160(uint256(value))));
    }
}
