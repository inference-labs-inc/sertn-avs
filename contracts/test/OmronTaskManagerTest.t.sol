// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "../src/OmronServiceManager.sol" as omrm;
import {InferenceDB} from "../src/InferenceDb.sol";
import {OmronTaskManager} from "../src/OmronTaskManager.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract OmronTaskManagerTest is Test, BLSMockAVSDeployer {
    omrm.OmronServiceManager sm;
    omrm.OmronServiceManager smImplementation;
    OmronTaskManager tm;
    OmronTaskManager tmImplementation;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator =
        address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator =
        address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();
        InferenceDB inferenceDB = new InferenceDB(0, address(0));

        tmImplementation = new OmronTaskManager(
            omrm.IRegistryCoordinator(address(registryCoordinator)),
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
                        generator,
                        address(inferenceDB)
                    )
                )
            )
        );
    }

    function testCreateNewTask() public {
        bytes memory quorumNumbers = new bytes(0);
        cheats.prank(generator);

        uint256[5] memory input = [0, 0, 0, 0, uint256(2)];
        tm.createNewTask(input, 100, quorumNumbers);
        assertEq(tm.latestTaskNum(), 1);
    }
}
