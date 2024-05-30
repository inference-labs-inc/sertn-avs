// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "../src/ZklayerServiceManager.sol" as zklm;
import {InferenceDB} from "../src/InferenceDb.sol";
import {ZklayerTaskManager} from "../src/ZklayerTaskManager.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract ZklayerTaskManagerTest is Test, BLSMockAVSDeployer {
    zklm.ZklayerServiceManager sm;
    zklm.ZklayerServiceManager smImplementation;
    ZklayerTaskManager tm;
    ZklayerTaskManager tmImplementation;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator =
        address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator =
        address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();
        InferenceDB inferenceDB = new InferenceDB(0, address(0));

        tmImplementation = new ZklayerTaskManager(
            zklm.IRegistryCoordinator(address(registryCoordinator)),
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = ZklayerTaskManager(
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
