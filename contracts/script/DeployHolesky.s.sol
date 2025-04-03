
pragma solidity ^0.8.13;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {SertnServiceManager} from "../src/SertnServiceManager.sol";
import {SertnTaskManager} from "../src/SertnTaskManager.sol";
import {ModelRegistry} from "../src/ModelRegistry.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IAllocationManager} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {ISignatureUtils} from "@eigenlayer/contracts/interfaces/ISignatureUtils.sol";
import {IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";

contract IAVSRegistrar {
    function registerOperator(address operator, address avs, uint32[] calldata operatorSetIds, bytes calldata data) external {}
    function deregisterOperator(address operator, address avs, uint32[] calldata operatorSetIds) external {}
    function supportsAVS(address avs) external view returns (bool) {}
    fallback () external {}
}

contract DeployHolesky is Script {
    function run() external {
        uint256 deployerKey = vm.envUint("HOLESKY_DEPLOYER_KEY");
        address avsDirectory = vm.envAddress("AVS_DIRECTORY");
        address allocationManager = vm.envAddress("ALLOCATION_MANAGER");
        address strategy = vm.envAddress("STRATEGY_ADDRESS");

        vm.startBroadcast(deployerKey);


        ProxyAdmin proxyAdmin = new ProxyAdmin();
        console2.log("Deployed ProxyAdmin at:", address(proxyAdmin));


        ModelRegistry modelRegistryImpl = new ModelRegistry();
        SertnTaskManager taskManagerImpl = new SertnTaskManager();
        SertnServiceManager serviceManagerImpl = new SertnServiceManager();

        console2.log("Deployed implementations:");
        console2.log("- ModelRegistry:", address(modelRegistryImpl));
        console2.log("- TaskManager:", address(taskManagerImpl));
        console2.log("- ServiceManager:", address(serviceManagerImpl));


        bytes memory empty = "";
        TransparentUpgradeableProxy modelRegistryProxy = new TransparentUpgradeableProxy(
            address(modelRegistryImpl),
            address(proxyAdmin),
            empty
        );

        TransparentUpgradeableProxy taskManagerProxy = new TransparentUpgradeableProxy(
            address(taskManagerImpl),
            address(proxyAdmin),
            empty
        );

        TransparentUpgradeableProxy serviceManagerProxy = new TransparentUpgradeableProxy(
            address(serviceManagerImpl),
            address(proxyAdmin),
            empty
        );

        console2.log("Deployed proxies:");
        console2.log("- ModelRegistry:", address(modelRegistryProxy));
        console2.log("- TaskManager:", address(taskManagerProxy));
        console2.log("- ServiceManager:", address(serviceManagerProxy));


        ModelRegistry(address(modelRegistryProxy)).initialize();
        SertnTaskManager(address(taskManagerProxy)).initialize();
        SertnServiceManager(address(serviceManagerProxy)).initialize(
            address(taskManagerProxy),
            address(modelRegistryProxy)
        );


        IAVSDirectory avsDir = IAVSDirectory(avsDirectory);
        IAllocationManager allocMgr = IAllocationManager(allocationManager);


        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = IStrategy(strategy);

        IAllocationManagerTypes.CreateSetParams[] memory sets = new IAllocationManagerTypes.CreateSetParams[](1);
        sets[0] = IAllocationManagerTypes.CreateSetParams({
            operatorSetId: 1,
            strategies: strategies
        });
        allocMgr.createOperatorSets(address(serviceManagerProxy), sets);


        uint256 expiry = type(uint256).max;
        bytes32 salt = bytes32(uint256(0) + 1);
        bytes32 digest = avsDir.calculateOperatorAVSRegistrationDigestHash(
            address(serviceManagerProxy),
            address(serviceManagerProxy),
            salt,
            expiry
        );
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(deployerKey, digest);

        avsDir.registerOperatorToAVS(
            address(serviceManagerProxy),
            ISignatureUtils.SignatureWithSaltAndExpiry(
                abi.encodePacked(r, s, v),
                salt,
                expiry
            )
        );


        allocMgr.setAVSRegistrar(address(serviceManagerProxy), new IAVSRegistrar());


        uint32[] memory operatorSetIds = new uint32[](1);
        operatorSetIds[0] = 1;

        IAllocationManagerTypes.RegisterParams memory register = IAllocationManagerTypes.RegisterParams({
            avs: address(serviceManagerProxy),
            operatorSetIds: operatorSetIds,
            data: ""
        });
        allocMgr.registerForOperatorSets(address(serviceManagerProxy), register);


        string memory deploymentInfo = string(abi.encodePacked(
            "{\n",
            '    "proxyAdmin": "', vm.toString(address(proxyAdmin)), '",\n',
            '    "modelRegistry": {\n',
            '        "implementation": "', vm.toString(address(modelRegistryImpl)), '",\n',
            '        "proxy": "', vm.toString(address(modelRegistryProxy)), '"\n',
            "    },\n",
            '    "taskManager": {\n',
            '        "implementation": "', vm.toString(address(taskManagerImpl)), '",\n',
            '        "proxy": "', vm.toString(address(taskManagerProxy)), '"\n',
            "    },\n",
            '    "serviceManager": {\n',
            '        "implementation": "', vm.toString(address(serviceManagerImpl)), '",\n',
            '        "proxy": "', vm.toString(address(serviceManagerProxy)), '"\n',
            "    }\n",
            "}"
        ));

        vm.writeFile("deployments/holesky.json", deploymentInfo);

        vm.stopBroadcast();
    }
}
