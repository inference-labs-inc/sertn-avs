// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/OmronServiceManager.sol" as incsqsm;
import {OmronTaskManager} from "../src/OmronTaskManager.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract OmronTaskManagerTest is BLSMockAVSDeployer {
    incsqsm.OmronServiceManager sm;
    incsqsm.OmronServiceManager smImplementation;
    OmronTaskManager tm;
    OmronTaskManager tmImplementation;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator =
        address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator =
        address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        tmImplementation = new OmronTaskManager(
            incsqsm.IRegistryCoordinator(address(registryCoordinator)),
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = OmronTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(tmImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        tm.initialize.selector,
                        pauserRegistry,
                        registryCoordinatorOwner,
                        aggregator,
                        generator
                    )
                )
            )
        );
    }

    function testCreateNewTask() public {
        bytes memory quorumNumbers = new bytes(0);
        cheats.prank(generator, generator);
        uint256[5] memory input = [0, 0, 0, 0, uint256(2)];
        tm.createNewTask(input, 100, quorumNumbers);
        assertEq(tm.latestTaskNum(), 1);
    }
}
