
pragma solidity ^0.8.0;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/Test.sol";
import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

import {SertnServiceManager} from "../src/SertnServiceManager.sol";
import {SertnTaskManager} from "../src/SertnTaskManager.sol";
import {ModelStorage} from "../src/ModelStorage.sol";

contract UpgradeHolesky is Script {
    using stdJson for *;
    using Strings for uint256;

    struct DeploymentData {
        address serviceManager;
        address taskManager;
        address modelStorage;
    }

    address private deployer;
    address proxyAdmin;

    function setUp() public virtual {
        console2.log("Starting setup...");
        string memory rpcUrl = vm.envString("RPC_URL");
        console2.log("Using RPC URL:", rpcUrl);

        uint256 forkId = vm.createFork(rpcUrl);
        vm.selectFork(forkId);
        console2.log("Fork created and selected");

        uint256 deployerKey = vm.envUint("HOLESKY_DEPLOYER_KEY");
        console2.log("Got deployer key:", deployerKey != 0);
        deployer = vm.rememberKey(deployerKey);
        console2.log("Actual deployer address:", deployer);

        proxyAdmin = vm.envAddress("PROXY_ADMIN");
        console2.log("Proxy admin from env:", proxyAdmin);
    }

    function readDeploymentJson() internal view returns (DeploymentData memory) {
        string memory filePath = string.concat("deployments/sertn/", vm.toString(block.chainid), ".json");
        console2.log("Reading Sertn deployment from:", filePath);

        require(vm.exists(filePath), "Deployment file does not exist");
        console2.log("File exists, reading contents...");

        string memory json = vm.readFile(filePath);
        console2.log("File contents read successfully");

        DeploymentData memory data;
        data.serviceManager = json.readAddress(".addresses.serviceManager");
        data.taskManager = json.readAddress(".addresses.taskManager");
        data.modelStorage = json.readAddress(".addresses.modelStorage");

        console2.log("Deployment data decoded successfully");
        console2.log("Service manager:", data.serviceManager);
        console2.log("Task manager:", data.taskManager);
        console2.log("Model storage:", data.modelStorage);

        return data;
    }

    function run() external {
        console2.log("Starting upgrade script...");
        vm.startBroadcast(deployer);

        DeploymentData memory deployment = readDeploymentJson();


        console2.log("Deploying new implementations...");
        SertnServiceManager newServiceManagerImpl = new SertnServiceManager();
        console2.log("New service manager implementation:", address(newServiceManagerImpl));

        SertnTaskManager newTaskManagerImpl = new SertnTaskManager();
        console2.log("New task manager implementation:", address(newTaskManagerImpl));

        ModelStorage newModelStorageImpl = new ModelStorage();
        console2.log("New model storage implementation:", address(newModelStorageImpl));


        console2.log("Upgrading service manager...");
        TransparentUpgradeableProxy(payable(deployment.serviceManager)).upgradeTo(address(newServiceManagerImpl));
        console2.log("Service manager upgraded successfully");


        console2.log("Upgrading task manager...");
        TransparentUpgradeableProxy(payable(deployment.taskManager)).upgradeTo(address(newTaskManagerImpl));
        console2.log("Task manager upgraded successfully");


        if (deployment.modelStorage != address(0)) {
            console2.log("Upgrading model storage...");
            TransparentUpgradeableProxy(payable(deployment.modelStorage)).upgradeTo(address(newModelStorageImpl));
            console2.log("Model storage upgraded successfully");
        }

        vm.stopBroadcast();


        string memory deploymentPath = string.concat("deployments/sertn/", vm.toString(block.chainid), ".json");
        string memory json = vm.readFile(deploymentPath);


        string memory addresses = string.concat(
            '"addresses":{',
            '"serviceManager":"', vm.toString(deployment.serviceManager), '",',
            '"taskManager":"', vm.toString(deployment.taskManager), '",',
            '"modelStorage":"', vm.toString(deployment.modelStorage), '"}'
        );


        string memory newJson = string.concat(
            '{',
            addresses, ',',
            '"implementations":{',
            '"serviceManager":"', vm.toString(address(newServiceManagerImpl)), '",',
            '"taskManager":"', vm.toString(address(newTaskManagerImpl)), '"'
        );

        if (deployment.modelStorage != address(0)) {
            newJson = string.concat(
                newJson,
                ',"modelStorage":"', vm.toString(address(newModelStorageImpl)), '"'
            );
        }
        newJson = string.concat(newJson, '}}');

        vm.writeFile(deploymentPath, newJson);
        console2.log("Updated implementation addresses in:", deploymentPath);

        console2.log("\nUpgrade complete!");
    }
}
