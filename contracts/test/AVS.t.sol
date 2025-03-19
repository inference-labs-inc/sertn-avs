// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {SertnServiceManager} from "../src/SertnServiceManager.sol";
import "../src/ISertnServiceManager.sol";
import {MockAVSDeployer} from "@eigenlayer-middleware/test/utils/MockAVSDeployer.sol";
import {ECDSAStakeRegistry} from "@eigenlayer-middleware/src/unaudited/ECDSAStakeRegistry.sol";
import {Vm} from "forge-std/Vm.sol";
// import {SertnDeploymentLib} from "../script/utils/SertnDeploymentLib.sol";
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
import {AllocationManager} from "@eigenlayer/contracts/core/AllocationManager.sol";
import {IAllocationManager} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";
import {OperatorSet} from "@eigenlayer/contracts/libraries/OperatorSetLib.sol";

import {Test, console2 as console} from "forge-std/Test.sol";
import {ISertnServiceManager} from "../src/ISertnServiceManager.sol";
import {ECDSAUpgradeable} from "@openzeppelin-upgrades/contracts/utils/cryptography/ECDSAUpgradeable.sol";

contract AVSSetup is Test {
    IECDSAStakeRegistryTypes.Quorum internal quorum;

    struct Operator {
        Vm.Wallet key;
        Vm.Wallet signingKey;
    }

    struct AVSOwner {
        Vm.Wallet key;
    }

    struct User {
        Vm.Wallet key;
    }

    Operator[] internal operators;
    AVSOwner internal owner;
    User internal user;

    OperatorSet internal opSet;

    // SertnDeploymentLib.DeploymentData internal sertnDeployment;
    CoreDeploymentLib.DeploymentData internal coreDeployment;
    CoreDeploymentLib.DeploymentConfigData coreConfigData;

    IStrategy[] _tokenToStrategy;
    IStrategy[] _ethStrategies;
    IStrategy _serStrategy;
    IStrategy[] strategies;

    uint32[] opSetIds;

    ERC20Mock public ethToken1;
    ERC20Mock public ethToken2;
    ERC20Mock public serToken;

    SertnServiceManager sertnServiceManager;

    mapping(address => IStrategy) public tokenToStrategy;

    function setUp() public virtual {
        owner = AVSOwner({key: vm.createWallet("owner_wallet")});

        address proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();

        coreConfigData = CoreDeploymentLib.readDeploymentConfigValues(
            "test/mockData/config/core/",
            1337
        ); // TODO: Fix this to correct path
        coreDeployment = CoreDeploymentLib.deployContracts(
            proxyAdmin,
            coreConfigData
        );

        ethToken1 = new ERC20Mock();
        ethToken2 = new ERC20Mock();
        serToken = new ERC20Mock();

        IStrategy strategy1 = addStrategy(address(ethToken1));
        IStrategy strategy2 = addStrategy(address(ethToken2));
        _serStrategy = addStrategy(address(serToken));
        strategies.push(_serStrategy);
        strategies.push(strategy1);
        strategies.push(strategy2);

        _ethStrategies.push(strategy1);
        _ethStrategies.push(strategy2);

        // quorum.strategies.push(IECDSAStakeRegistryTypes.StrategyParams({strategy: strategy, multiplier: 10_000}));

        // sertnDeployment = SertnDeploymentLib.deployContracts(
        //     proxyAdmin, coreDeployment, quorum, owner.key.addr, owner.key.addr
        // );
        vm.startPrank(owner.key.addr);
        sertnServiceManager = new SertnServiceManager(
            coreDeployment.rewardsCoordinator,
            coreDeployment.delegationManager,
            coreDeployment.allocationManager,
            strategies,
            ""
        );
        vm.stopPrank();
        labelContracts(coreDeployment);
        opSetIds.push(0);
    }

    function addStrategy(address token) public returns (IStrategy) {
        if (tokenToStrategy[token] != IStrategy(address(0))) {
            return tokenToStrategy[token];
        }

        StrategyFactory strategyFactory = StrategyFactory(
            coreDeployment.strategyFactory
        );
        IStrategy newStrategy = strategyFactory.deployNewStrategy(
            IERC20(token)
        );
        tokenToStrategy[token] = newStrategy;
        return newStrategy;
    }

    function labelContracts(
        CoreDeploymentLib.DeploymentData memory coreDeployment
    ) internal {
        vm.label(coreDeployment.delegationManager, "DelegationManager");
        vm.label(coreDeployment.allocationManager, "AllocationManager");
        vm.label(coreDeployment.strategyManager, "StrategyManager");
        vm.label(coreDeployment.eigenPodManager, "EigenPodManager");
        vm.label(coreDeployment.rewardsCoordinator, "RewardsCoordinator");
        vm.label(coreDeployment.eigenPodBeacon, "EigenPodBeacon");
        vm.label(coreDeployment.pauserRegistry, "PauserRegistry");
        vm.label(coreDeployment.strategyFactory, "StrategyFactory");
        vm.label(coreDeployment.strategyBeacon, "StrategyBeacon");
        vm.label(address(sertnServiceManager), "SertnServiceManager");
        // vm.label(sertnDeployment.stakeRegistry, "StakeRegistry");
    }

    function signWithOperatorKey(
        Operator memory operator,
        bytes32 digest
    ) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(
            operator.key.privateKey,
            digest
        );
        return abi.encodePacked(r, s, v);
    }

    function signWithSigningKey(
        Operator memory operator,
        bytes32 digest
    ) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(
            operator.signingKey.privateKey,
            digest
        );
        return abi.encodePacked(r, s, v);
    }

    function mintMockTokens(Operator memory operator, uint256 amount) internal {
        ethToken1.mint(operator.key.addr, amount);
        ethToken2.mint(operator.key.addr, amount);
        serToken.mint(operator.key.addr, amount);
    }

    function depositTokenIntoStrategy(
        Operator memory operator,
        address token,
        uint256 amount
    ) internal returns (uint256) {
        IStrategy strategy = IStrategy(tokenToStrategy[token]);
        require(address(strategy) != address(0), "Strategy was not found");
        IStrategyManager strategyManager = IStrategyManager(
            coreDeployment.strategyManager
        );

        vm.startPrank(operator.key.addr);
        ethToken1.approve(address(strategyManager), amount);
        ethToken2.approve(address(strategyManager), amount);
        serToken.approve(address(strategyManager), amount);

        uint256 shares = strategyManager.depositIntoStrategy(
            strategy,
            IERC20(token),
            amount
        );
        vm.stopPrank();

        return shares;
    }

    function registerAsOperator(Operator memory operator) internal {
        IDelegationManager delegationManager = IDelegationManager(
            coreDeployment.delegationManager
        );

        IDelegationManagerTypes.OperatorDetails
            memory operatorDetails = IDelegationManagerTypes.OperatorDetails({
                __deprecated_earningsReceiver: operator.key.addr,
                delegationApprover: address(0),
                __deprecated_stakerOptOutWindowBlocks: 0
            });

        vm.prank(operator.key.addr);
        delegationManager.registerAsOperator(address(operator.key.addr), 0, "");
    }

    function registerAsOperatorToAVS(Operator memory operator) internal {
        IAllocationManager allocationManager = IAllocationManager(
            coreDeployment.allocationManager
        );
        ISertnServiceManager sm = ISertnServiceManager(
            address(sertnServiceManager)
        );
        ISertnServiceManagerTypes.Model[]
            memory _model = new ISertnServiceManagerTypes.Model[](1);
        uint256[] memory _ethShares = new uint256[](2);
        _ethShares[0] = 10;
        _ethShares[1] = 10;
        _model[0] = ISertnServiceManagerTypes.Model({
            modelName_: bytes32("1"),
            operator_: operator.key.addr,
            verifier_: address(0),
            benchData_: "",
            inputs_: bytes(""),
            output_: bytes(""),
            maxBlocks_: 1000,
            ethStrategies_: _ethStrategies,
            ethShares_: _ethShares,
            baseFee_: 1e2,
            maxSer_: 1e4,
            computeType_: bytes32("model1"),
            proveOnResponse_: false,
            available_: true
        });
        bytes32[] memory _computeUnitNames = new bytes32[](1);
        _computeUnitNames[0] = bytes32("model1");
        uint8[] memory _computeUnits = new uint8[](1);
        _computeUnits[0] = 10;
        bytes memory _data = abi.encode(
            _model,
            _computeUnitNames,
            _computeUnits
        );
        IAllocationManagerTypes.RegisterParams
            memory registerParams = IAllocationManagerTypes.RegisterParams({
                avs: address(sm),
                operatorSetIds: opSetIds,
                data: _data
            });
        vm.startPrank(operator.key.addr);
        allocationManager.registerForOperatorSets(
            operator.key.addr,
            registerParams
        );
        vm.stopPrank();
    }

    function createAndAddOperator() internal returns (Operator memory) {
        Vm.Wallet memory operatorKey = vm.createWallet(
            string.concat("operator", vm.toString(operators.length))
        );
        Vm.Wallet memory signingKey = vm.createWallet(
            string.concat("signing", vm.toString(operators.length))
        );

        Operator memory newOperator = Operator({
            key: operatorKey,
            signingKey: signingKey
        });

        operators.push(newOperator);
        return newOperator;
    }
}

contract RegisterOperatorToAVS is AVSSetup {
    uint256 internal constant INITIAL_BALANCE = 100 ether;
    uint256 internal constant DEPOSIT_AMOUNT = 1 ether;
    uint256 internal constant OPERATOR_COUNT = 4;
    IDelegationManager internal delegationManager;
    AllocationManager internal allocationManager;
    IPermissionController internal permissionController;

    IStrategy[] strat;
    IAllocationManagerTypes.CreateSetParams[] setParams;

    function setUp() public virtual override {
        super.setUp();
        /// Setting to internal state for convenience
        delegationManager = IDelegationManager(
            coreDeployment.delegationManager
        );
        allocationManager = AllocationManager(coreDeployment.allocationManager);
        permissionController = IPermissionController(
            coreDeployment.permissionController
        );

        // vm.startPrank(address(sm));

        // allocationManager.updateAVSMetadataURI(address(sm), "");
        // allocationManager.setAVSRegistrar(address(sm), sertnRegistrar);
        // delete strat;
        // strat.push(tokenToStrategy[address(mockToken)]);
        // delete setParams;
        // setParams.push(IAllocationManagerTypes.CreateSetParams({operatorSetId:0, strategies: strat}));
        // allocationManager.createOperatorSets(address(sm), setParams);

        // vm.stopPrank();

        // addStrategy(address(mockToken));
        opSet = OperatorSet({avs: address(sertnServiceManager), id: 0});

        while (operators.length < OPERATOR_COUNT) {
            createAndAddOperator();
        }

        uint64[] memory _newMags = new uint64[](3);
        _newMags[0] = 1 ether;
        _newMags[1] = 1 ether;
        _newMags[2] = 1 ether;

        IAllocationManagerTypes.AllocateParams[]
            memory allocateParams = new IAllocationManagerTypes.AllocateParams[](
                1
            );
        allocateParams[0] = IAllocationManagerTypes.AllocateParams({
            operatorSet: opSet,
            strategies: strategies,
            newMagnitudes: _newMags
        });

        address[] memory _operatorKeys = new address[](OPERATOR_COUNT);
        for (uint8 i = 0; i < OPERATOR_COUNT; i++) {
            _operatorKeys[i] = operators[i].key.addr;
        }

        for (uint256 i = 0; i < OPERATOR_COUNT; i++) {
            mintMockTokens(operators[i], INITIAL_BALANCE);
            depositTokenIntoStrategy(
                operators[i],
                address(ethToken1),
                DEPOSIT_AMOUNT
            );
            depositTokenIntoStrategy(
                operators[i],
                address(ethToken2),
                DEPOSIT_AMOUNT
            );
            depositTokenIntoStrategy(
                operators[i],
                address(serToken),
                DEPOSIT_AMOUNT
            );
            registerAsOperator(operators[i]);
            registerAsOperatorToAVS(operators[i]);
            vm.startPrank(_operatorKeys[i]);
            allocationManager.modifyAllocations(
                _operatorKeys[i],
                allocateParams
            );
            vm.stopPrank();
        }
    }

    function test_sendTask() public {
        vm.roll(1e10);
        user = User({key: vm.createWallet("user_wallet")});
        vm.startPrank(user.key.addr);

        ethToken1.mint(user.key.addr, 1 ether);
        ethToken2.mint(user.key.addr, 1 ether);
        serToken.mint(user.key.addr, 1 ether);

        ethToken1.approve(address(sertnServiceManager), 1e4);
        ethToken2.approve(address(sertnServiceManager), 1e4);
        serToken.approve(address(sertnServiceManager), 1e4);

        ISertnServiceManagerTypes.Task memory task = ISertnServiceManagerTypes
            .Task({
                modelId_: 1,
                inputs_: bytes(""),
                poc_: 1e2,
                startTime_: 0,
                startingBlock_: 0,
                proveOnResponse_: false,
                user_: user.key.addr
            });
        sertnServiceManager.sendTask(task);
        vm.stopPrank();

        vm.startPrank(operators[0].key.addr);
        ISertnServiceManagerTypes.Operator
            memory _operator = sertnServiceManager.getOperatorInfo(
                operators[0].key.addr
            );
        ISertnServiceManagerTypes.Task memory _task = abi.decode(
            _operator.openTasks_[0],
            (ISertnServiceManagerTypes.Task)
        );
        require(_task.user_ == user.key.addr);
        ISertnServiceManagerTypes.TaskResponse
            memory _taskResponse = ISertnServiceManagerTypes.TaskResponse({
                taskId_: _operator.openTasks_[0],
                output_: bytes("hello world"),
                proven_: false
            });
        sertnServiceManager.submitTask(_taskResponse, false, bytes(""));
        vm.stopPrank();

        vm.startPrank(user.key.addr);
        _taskResponse = sertnServiceManager.getTaskResponse(
            _operator.openTasks_[0]
        );
        string memory _outputData = string(_taskResponse.output_);
        console.log(_outputData);
        vm.startPrank(operators[0].key.addr);
        address[] memory _operatorKeys = new address[](OPERATOR_COUNT);
        for (uint8 i = 0; i < OPERATOR_COUNT; i++) {
            _operatorKeys[i] = operators[i].key.addr;
        }
        vm.stopPrank();
        vm.startPrank(owner.key.addr);
        console.log(
            allocationManager.getMinimumSlashableStake(
                opSet,
                _operatorKeys,
                strategies,
                uint32(vm.getBlockNumber())
            )[0][1],
            "min slashable"
        );
        sertnServiceManager.slashOperator(_operator.openTasks_[0], "test");
        console.log(
            allocationManager.getMinimumSlashableStake(
                opSet,
                _operatorKeys,
                strategies,
                uint32(vm.getBlockNumber())
            )[0][1],
            "min slashable"
        );
        vm.stopPrank();
    }

    function test_sendTaskb() public {
        vm.roll(1e10);
        user = User({key: vm.createWallet("user_wallet")});
        vm.startPrank(user.key.addr);

        ethToken1.mint(user.key.addr, 1 ether);
        ethToken2.mint(user.key.addr, 1 ether);
        serToken.mint(user.key.addr, 1 ether);

        ethToken1.approve(address(sertnServiceManager), 1e4);
        ethToken2.approve(address(sertnServiceManager), 1e4);
        serToken.approve(address(sertnServiceManager), 1e4);

        ISertnServiceManagerTypes.Task memory task = ISertnServiceManagerTypes
            .Task({
                modelId_: 1,
                inputs_: bytes(""),
                poc_: 1e2,
                startTime_: 0,
                startingBlock_: 0,
                proveOnResponse_: false,
                user_: user.key.addr
            });
        sertnServiceManager.sendTask(task);
        vm.stopPrank();

        vm.startPrank(operators[0].key.addr);
        ISertnServiceManagerTypes.Operator
            memory _operator = sertnServiceManager.getOperatorInfo(
                operators[0].key.addr
            );
        ISertnServiceManagerTypes.Task memory _task = abi.decode(
            _operator.openTasks_[0],
            (ISertnServiceManagerTypes.Task)
        );
        require(_task.user_ == user.key.addr);
        ISertnServiceManagerTypes.TaskResponse
            memory _taskResponse = ISertnServiceManagerTypes.TaskResponse({
                taskId_: _operator.openTasks_[0],
                output_: bytes("hello world"),
                proven_: false
            });
        sertnServiceManager.submitTask(_taskResponse, false, bytes(""));
        vm.stopPrank();

        vm.startPrank(user.key.addr);
        _taskResponse = sertnServiceManager.getTaskResponse(
            _operator.openTasks_[0]
        );
        string memory _outputData = string(_taskResponse.output_);
        console.log(_outputData);
        vm.startPrank(operators[0].key.addr);
        address[] memory _operatorKeys = new address[](OPERATOR_COUNT);
        for (uint8 i = 0; i < OPERATOR_COUNT; i++) {
            _operatorKeys[i] = operators[i].key.addr;
        }
        vm.stopPrank();
        vm.startPrank(owner.key.addr);
        vm.roll(1e10 + 1e4);
        _operator = sertnServiceManager.getOperatorInfo(operators[0].key.addr);
        console.log(_operator.allocatedSer_);
        sertnServiceManager.clearTask(_operator.openTasks_[0]);
        _operator = sertnServiceManager.getOperatorInfo(operators[0].key.addr);
        console.log(_operator.allocatedSer_);
        vm.stopPrank();
    }
}
