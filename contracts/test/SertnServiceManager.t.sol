// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {SertnServiceManager} from "../src/SertnServiceManager.sol";
import {MockAVSDeployer} from "@eigenlayer-middleware/test/utils/MockAVSDeployer.sol";
import {ECDSAStakeRegistry} from "@eigenlayer-middleware/src/unaudited/ECDSAStakeRegistry.sol";
import {Vm} from "forge-std/Vm.sol";
import {console2} from "forge-std/Test.sol";
import {SertnDeploymentLib} from "../script/utils/SertnDeploymentLib.sol";
import {CoreDeploymentLib} from "../script/utils/CoreDeploymentLib.sol";
import {UpgradeableProxyLib} from "../script/utils/UpgradeableProxyLib.sol";
import {ERC20Mock} from "./ERC20Mock.sol";
import {IERC20, StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";
import {IPermissionController} from "@eigenlayer/contracts/interfaces/IPermissionController.sol";
import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {IECDSAStakeRegistryTypes} from "@eigenlayer-middleware/src/interfaces/IECDSAStakeRegistry.sol";

import {IStrategyManager} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {IDelegationManagerTypes} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {ISignatureUtils} from "@eigenlayer/contracts/interfaces/ISignatureUtils.sol";
import {AllocationManager} from "@eigenlayer/contracts/core/AllocationManager.sol";
import {IAllocationManager} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {Test, console2 as console} from "forge-std/Test.sol";
import {ISertnServiceManager} from "../src/ISertnServiceManager.sol";
import {ECDSAUpgradeable} from
    "@openzeppelin-upgrades/contracts/utils/cryptography/ECDSAUpgradeable.sol";

contract SertnTaskManagerSetup is Test {
    IECDSAStakeRegistryTypes.Quorum internal quorum;

    struct Operator {
        Vm.Wallet key;
        Vm.Wallet signingKey;
    }

    struct TrafficGenerator {
        Vm.Wallet key;
    }

    struct AVSOwner {
        Vm.Wallet key;
    }

    Operator[] internal operators;
    TrafficGenerator internal generator;
    AVSOwner internal owner;

    SertnDeploymentLib.DeploymentData internal sertnDeployment;
    CoreDeploymentLib.DeploymentData internal coreDeployment;
    CoreDeploymentLib.DeploymentConfigData coreConfigData;

    IStrategy[] _tokenToStrategy;

    ERC20Mock public mockToken;

    mapping(address => IStrategy) public tokenToStrategy;

    function setUp() public virtual {
        generator = TrafficGenerator({key: vm.createWallet("generator_wallet")});
        owner = AVSOwner({key: vm.createWallet("owner_wallet")});

        address proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();

        coreConfigData =
            CoreDeploymentLib.readDeploymentConfigValues("test/mockData/config/core/", 1337); // TODO: Fix this to correct path
        coreDeployment = CoreDeploymentLib.deployContracts(proxyAdmin, coreConfigData);

        mockToken = new ERC20Mock();

        IStrategy strategy = addStrategy(address(mockToken));
        quorum.strategies.push(IECDSAStakeRegistryTypes.StrategyParams({strategy: strategy, multiplier: 10_000}));

        sertnDeployment = SertnDeploymentLib.deployContracts(
            proxyAdmin, coreDeployment, quorum, owner.key.addr, owner.key.addr
        );
        labelContracts(coreDeployment, sertnDeployment);
    }

    function addStrategy(
        address token
    ) public returns (IStrategy) {
        if (tokenToStrategy[token] != IStrategy(address(0))) {
            return tokenToStrategy[token];
        }

        StrategyFactory strategyFactory = StrategyFactory(coreDeployment.strategyFactory);
        IStrategy newStrategy = strategyFactory.deployNewStrategy(IERC20(token));
        tokenToStrategy[token] = newStrategy;
        return newStrategy;
    }

    function labelContracts(
        CoreDeploymentLib.DeploymentData memory coreDeployment,
        SertnDeploymentLib.DeploymentData memory sertnDeployment
    ) internal {
        vm.label(coreDeployment.delegationManager, "DelegationManager");
        // vm.label(coreDeployment.avsDirectory, "AVSDirectory");
        vm.label(coreDeployment.allocationManager, "AllocationManager");
        vm.label(coreDeployment.strategyManager, "StrategyManager");
        vm.label(coreDeployment.eigenPodManager, "EigenPodManager");
        vm.label(coreDeployment.rewardsCoordinator, "RewardsCoordinator");
        vm.label(coreDeployment.eigenPodBeacon, "EigenPodBeacon");
        vm.label(coreDeployment.pauserRegistry, "PauserRegistry");
        vm.label(coreDeployment.strategyFactory, "StrategyFactory");
        vm.label(coreDeployment.strategyBeacon, "StrategyBeacon");
        vm.label(sertnDeployment.sertnServiceManager, "SertnServiceManager");
        vm.label(sertnDeployment.stakeRegistry, "StakeRegistry");
    }

    function signWithOperatorKey(
        Operator memory operator,
        bytes32 digest
    ) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(operator.key.privateKey, digest);
        return abi.encodePacked(r, s, v);
    }

    function signWithSigningKey(
        Operator memory operator,
        bytes32 digest
    ) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(operator.signingKey.privateKey, digest);
        return abi.encodePacked(r, s, v);
    }

    function mintMockTokens(Operator memory operator, uint256 amount) internal {
        mockToken.mint(operator.key.addr, amount);
    }

    function depositTokenIntoStrategy(
        Operator memory operator,
        address token,
        uint256 amount
    ) internal returns (uint256) {
        IStrategy strategy = IStrategy(tokenToStrategy[token]);
        require(address(strategy) != address(0), "Strategy was not found");
        IStrategyManager strategyManager = IStrategyManager(coreDeployment.strategyManager);

        vm.startPrank(operator.key.addr);
        mockToken.approve(address(strategyManager), amount);
        uint256 shares = strategyManager.depositIntoStrategy(strategy, IERC20(token), amount);
        vm.stopPrank();

        return shares;
    }

    function registerAsOperator(
        Operator memory operator
    ) internal {
        IDelegationManager delegationManager = IDelegationManager(coreDeployment.delegationManager);

        IDelegationManagerTypes.OperatorDetails memory operatorDetails = IDelegationManagerTypes.OperatorDetails({
            __deprecated_earningsReceiver: operator.key.addr,
            delegationApprover: address(0),
            __deprecated_stakerOptOutWindowBlocks: 0
        });

        vm.prank(operator.key.addr);
        delegationManager.registerAsOperator(address(operator.key.addr), 0, "");
        // operatorDetails, "");
    }
// replace with allocation manager implementation
    // function registerOperatorToAVS(
    //     Operator memory operator
    // ) internal {
    //     ECDSAStakeRegistry stakeRegistry = ECDSAStakeRegistry(sertnDeployment.stakeRegistry);
    //     AllocationManager allocationManager = AllocationManager(coreDeployment.allocationManager);

    //     bytes32 salt = keccak256(abi.encodePacked(block.timestamp, operator.key.addr));
    //     uint256 expiry = block.timestamp + 1 hours;
    //     allocationManager.calculateOperatorAVSRegistrationDigestHash();
    //     bytes32 operatorRegistrationDigestHash = allocationManager
    //         .calculateOperatorAVSRegistrationDigestHash(
    //         operator.key.addr, address(sertnDeployment.sertnServiceManager), salt, expiry
    //     );

    //     bytes memory signature = signWithOperatorKey(operator, operatorRegistrationDigestHash);

    //     ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = ISignatureUtils
    //         .SignatureWithSaltAndExpiry({signature: signature, salt: salt, expiry: expiry});

    //     vm.prank(address(operator.key.addr));
    //     stakeRegistry.registerOperatorWithSignature(operatorSignature, operator.signingKey.addr);
    // }

    // function deregisterOperatorFromAVS(
    //     Operator memory operator
    // ) internal {
    //     ECDSAStakeRegistry stakeRegistry = ECDSAStakeRegistry(sertnDeployment.stakeRegistry);

    //     vm.prank(operator.key.addr);
    //     stakeRegistry.deregisterOperator();
    // }

    function createAndAddOperator() internal returns (Operator memory) {
        Vm.Wallet memory operatorKey =
            vm.createWallet(string.concat("operator", vm.toString(operators.length)));
        Vm.Wallet memory signingKey =
            vm.createWallet(string.concat("signing", vm.toString(operators.length)));

        Operator memory newOperator = Operator({key: operatorKey, signingKey: signingKey});

        operators.push(newOperator);
        return newOperator;
    }

    function updateOperatorWeights(
        Operator[] memory _operators
    ) internal {
        ECDSAStakeRegistry stakeRegistry = ECDSAStakeRegistry(sertnDeployment.stakeRegistry);

        address[] memory operatorAddresses = new address[](_operators.length);
        for (uint256 i = 0; i < _operators.length; i++) {
            operatorAddresses[i] = _operators[i].key.addr;
        }

        stakeRegistry.updateOperators(operatorAddresses);
    }

    function getSortedOperatorSignatures(
        Operator[] memory _operators,
        bytes32 digest
    ) internal pure returns (bytes[] memory) {
        uint256 length = _operators.length;
        bytes[] memory signatures = new bytes[](length);
        address[] memory addresses = new address[](length);

        for (uint256 i = 0; i < length; i++) {
            addresses[i] = _operators[i].key.addr;
            signatures[i] = signWithOperatorKey(_operators[i], digest);
        }

        for (uint256 i = 0; i < length - 1; i++) {
            for (uint256 j = 0; j < length - i - 1; j++) {
                if (addresses[j] > addresses[j + 1]) {
                    // Swap addresses
                    address tempAddr = addresses[j];
                    addresses[j] = addresses[j + 1];
                    addresses[j + 1] = tempAddr;

                    // Swap signatures
                    bytes memory tempSig = signatures[j];
                    signatures[j] = signatures[j + 1];
                    signatures[j + 1] = tempSig;
                }
            }
        }

        return signatures;
    }

    function createTask(TrafficGenerator memory generator, string memory taskName) internal {
        ISertnServiceManager sertnServiceManager =
            ISertnServiceManager(sertnDeployment.sertnServiceManager);

        vm.prank(generator.key.addr);
        sertnServiceManager.createNewTask(taskName);
    }

    function respondToTask(
        Operator memory operator,
        ISertnServiceManager.Task memory task,
        uint32 referenceTaskIndex
    ) internal {
        // Create the message hash
        bytes32 messageHash = keccak256(abi.encodePacked("Hello, ", task.name));

        bytes memory signature = signWithSigningKey(operator, messageHash);

        address[] memory operators = new address[](1);
        operators[0] = operator.key.addr;
        bytes[] memory signatures = new bytes[](1);
        signatures[0] = signature;

        bytes memory signedTask = abi.encode(operators, signatures, uint32(block.number));

        ISertnServiceManager(sertnDeployment.sertnServiceManager).respondToTask(
            task, referenceTaskIndex, signedTask
        );
    }
}

contract SertnServiceManagerInitialization is SertnTaskManagerSetup {
    function testInitialization() public view {
        ECDSAStakeRegistry stakeRegistry = ECDSAStakeRegistry(sertnDeployment.stakeRegistry);

        IECDSAStakeRegistryTypes.Quorum memory quorum = stakeRegistry.quorum();

        assertGt(quorum.strategies.length, 0, "No strategies in quorum");
        assertEq(
            address(quorum.strategies[0].strategy),
            address(tokenToStrategy[address(mockToken)]),
            "First strategy doesn't match mock token strategy"
        );

        assertTrue(sertnDeployment.stakeRegistry != address(0), "StakeRegistry not deployed");
        assertTrue(
            sertnDeployment.sertnServiceManager != address(0),
            "SertnServiceManager not deployed"
        );
        assertTrue(coreDeployment.delegationManager != address(0), "DelegationManager not deployed");
        // assertTrue(coreDeployment.avsDirectory != address(0), "AVSDirectory not deployed");
        assertTrue(coreDeployment.strategyManager != address(0), "StrategyManager not deployed");
        assertTrue(coreDeployment.eigenPodManager != address(0), "EigenPodManager not deployed");
        assertTrue(coreDeployment.strategyFactory != address(0), "StrategyFactory not deployed");
        assertTrue(coreDeployment.strategyBeacon != address(0), "StrategyBeacon not deployed");
    }
}

contract RegisterOperator is SertnTaskManagerSetup {
    uint256 internal constant INITIAL_BALANCE = 100 ether;
    uint256 internal constant DEPOSIT_AMOUNT = 1 ether;
    uint256 internal constant OPERATOR_COUNT = 4;
    IDelegationManager internal delegationManager;
    AllocationManager internal allocationManager;
    ISertnServiceManager internal sm;
    ECDSAStakeRegistry internal stakeRegistry;
    IPermissionController internal permissionController;

    function setUp() public virtual override {
        super.setUp();
        /// Setting to internal state for convenience
        delegationManager = IDelegationManager(coreDeployment.delegationManager);
        allocationManager = AllocationManager(coreDeployment.allocationManager);
        permissionController = IPermissionController(coreDeployment.permissionController);
        
       
        sm = ISertnServiceManager(sertnDeployment.sertnServiceManager);
        
        vm.startPrank(address(sm));
        
        allocationManager.updateAVSMetadataURI(address(sm), "");
        vm.stopPrank();
        stakeRegistry = ECDSAStakeRegistry(sertnDeployment.stakeRegistry);

        addStrategy(address(mockToken));

        while (operators.length < OPERATOR_COUNT) {
            createAndAddOperator();
        }

        for (uint256 i = 0; i < OPERATOR_COUNT; i++) {
            mintMockTokens(operators[i], INITIAL_BALANCE);

            depositTokenIntoStrategy(operators[i], address(mockToken), DEPOSIT_AMOUNT);

            registerAsOperator(operators[i]);
        }
    }

    function testVerifyOperatorStates() public {
        for (uint256 i = 0; i < OPERATOR_COUNT; i++) {
            address operatorAddr = operators[i].key.addr;
            _tokenToStrategy.push(tokenToStrategy[address(mockToken)]);
            uint256[] memory operatorShares = 
                delegationManager.getOperatorShares(operatorAddr, _tokenToStrategy);
            delete _tokenToStrategy;
            assertEq(operatorShares[0], DEPOSIT_AMOUNT, "Operator shares in DelegationManager incorrect");
        }
    }

    // function test_RegisterOperatorToAVS() public {
    //     address operatorAddr = operators[0].key.addr;
    //     registerOperatorToAVS(operators[0]);
    //     assertTrue(
    //         allocationManager.avsOperatorStatus(address(sm), operatorAddr)
    //             == IAllocationManager.OperatorAVSRegistrationStatus.REGISTERED,
    //         "Operator not registered in AVSDirectory"
    //     );

    //     address signingKey = stakeRegistry.getLastestOperatorSigningKey(operatorAddr);
    //     assertTrue(signingKey != address(0), "Operator signing key not set in ECDSAStakeRegistry");

    //     uint256 operatorWeight = stakeRegistry.getLastCheckpointOperatorWeight(operatorAddr);
    //     assertTrue(operatorWeight > 0, "Operator weight not set in ECDSAStakeRegistry");
    // }
}

contract CreateTask is SertnTaskManagerSetup {
    ISertnServiceManager internal sm;

    function setUp() public override {
        super.setUp();
        sm = ISertnServiceManager(sertnDeployment.sertnServiceManager);
    }

    function testCreateTask() public {
        string memory taskName = "Test Task";

        vm.prank(generator.key.addr);
        ISertnServiceManager.Task memory newTask = sm.createNewTask(taskName);

        require(
            sha256(abi.encodePacked(newTask.name)) == sha256(abi.encodePacked(taskName)),
            "Task name not set correctly"
        );
        require(
            newTask.taskCreatedBlock == uint32(block.number), "Task created block not set correctly"
        );
    }
}

contract RespondToTask is SertnTaskManagerSetup {
    using ECDSAUpgradeable for bytes32;

    uint256 internal constant INITIAL_BALANCE = 100 ether;
    uint256 internal constant DEPOSIT_AMOUNT = 1 ether;
    uint256 internal constant OPERATOR_COUNT = 4;

    IDelegationManager internal delegationManager;
    AllocationManager internal allocationManager;
    ISertnServiceManager internal sm;
    ECDSAStakeRegistry internal stakeRegistry;

    function setUp() public override {
        super.setUp();

        delegationManager = IDelegationManager(coreDeployment.delegationManager);
        allocationManager = AllocationManager(coreDeployment.allocationManager);
        sm = ISertnServiceManager(sertnDeployment.sertnServiceManager);
        stakeRegistry = ECDSAStakeRegistry(sertnDeployment.stakeRegistry);

        addStrategy(address(mockToken));

        while (operators.length < OPERATOR_COUNT) {
            createAndAddOperator();
        }

        for (uint256 i = 0; i < OPERATOR_COUNT; i++) {
            mintMockTokens(operators[i], INITIAL_BALANCE);

            depositTokenIntoStrategy(operators[i], address(mockToken), DEPOSIT_AMOUNT);

            registerAsOperator(operators[i]);

            // registerOperatorToAVS(operators[i]);
        }
    }

    function testRespondToTask() public {
        string memory taskName = "TestTask";
        ISertnServiceManager.Task memory newTask = sm.createNewTask(taskName);
        uint32 taskIndex = sm.latestTaskNum() - 1;

        bytes32 messageHash = keccak256(abi.encodePacked("Hello, ", taskName));
        bytes32 ethSignedMessageHash = messageHash.toEthSignedMessageHash();
        bytes memory signature = signWithSigningKey(operators[0], ethSignedMessageHash); // TODO: Use signing key after changes to service manager

        address[] memory operatorsMem = new address[](1);
        operatorsMem[0] = operators[0].key.addr;
        bytes[] memory signatures = new bytes[](1);
        signatures[0] = signature;

        bytes memory signedTask = abi.encode(operatorsMem, signatures, uint32(block.number));

        vm.roll(block.number + 1);
        sm.respondToTask(newTask, taskIndex, signedTask);
    }
}
