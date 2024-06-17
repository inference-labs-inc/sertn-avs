// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;
import {SertnServiceManager, IServiceManager} from "../src/SertnServiceManager.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/UpdateMetadata.s.sol:SertnDeployer --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract SertnDeployer is Script, Utils {
    SertnServiceManager serviceManager =
        SertnServiceManager(0x91eBe0B98aA1801868AcE6a66fA8D7cEE00d8eA3);

    function run() external {
        vm.startBroadcast();
        console.log(serviceManager.owner());
        console.log(msg.sender);
        serviceManager.updateAVSMetadataURI(
            "https://raw.githubusercontent.com/inference-labs-inc/sertn-avs/main/el-info.json"
        );
        vm.stopBroadcast();
    }
}
