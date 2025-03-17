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
import {IAllocationManagerTypes} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";
import {SertnRegistrar} from "../src/SertnRegistrar.sol";
import {OperatorSet} from "@eigenlayer/contracts/libraries/OperatorSetLib.sol";

import {Test, console2 as console} from "forge-std/Test.sol";
import {ISertnServiceManager} from "../src/ISertnServiceManager.sol";
import {ECDSAUpgradeable} from
    "@openzeppelin-upgrades/contracts/utils/cryptography/ECDSAUpgradeable.sol";

contract AVSSetup is Test {
    IECDSAStakeRegistryTypes.Quorum internal quorum;

    struct Operator {
        Vm.Wallet key;
        Vm.Wallet signingKey;
    }

    struct AVSOwner {
        Vm.Wallet key;
    }

    Operator[] internal operators;
    AVSOwner internal owner;

    SertnDeploymentLib.DeploymentData internal sertnDeployment;
    CoreDeploymentLib.DeploymentData internal coreDeployment;
    CoreDeploymentLib.DeploymentConfigData coreConfigData;

    IStrategy[] _tokenToStrategy;

    uint32[] opSetIds;

    ERC20Mock public mockToken;

    mapping(address => IStrategy) public tokenToStrategy;

    function setUp() public virtual {
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
    }

    function registerAsOperatorToAVS(Operator memory operator) internal {
        IAllocationManager allocationManager = IAllocationManager(coreDeployment.allocationManager);
        ISertnServiceManager sm = ISertnServiceManager(sertnDeployment.sertnServiceManager);
        delete opSetIds;
        opSetIds.push(0);
        IAllocationManagerTypes.RegisterParams memory registerParams = IAllocationManagerTypes.RegisterParams({avs:address(sm), operatorSetIds: opSetIds, data: bytes("")});
        vm.startPrank(operator.key.addr);
        allocationManager.registerForOperatorSets(operator.key.addr, registerParams);
        vm.stopPrank();
    }


    function createAndAddOperator() internal returns (Operator memory) {
        Vm.Wallet memory operatorKey =
            vm.createWallet(string.concat("operator", vm.toString(operators.length)));
        Vm.Wallet memory signingKey =
            vm.createWallet(string.concat("signing", vm.toString(operators.length)));

        Operator memory newOperator = Operator({key: operatorKey, signingKey: signingKey});

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
    ISertnServiceManager internal sm;
    ECDSAStakeRegistry internal stakeRegistry;
    IPermissionController internal permissionController;
    SertnRegistrar internal sertnRegistrar;

    IStrategy[] strat;
    IAllocationManagerTypes.CreateSetParams[] setParams;

    function setUp() public virtual override {
        super.setUp();
        /// Setting to internal state for convenience
        delegationManager = IDelegationManager(coreDeployment.delegationManager);
        allocationManager = AllocationManager(coreDeployment.allocationManager);
        permissionController = IPermissionController(coreDeployment.permissionController);
        
       
        sm = ISertnServiceManager(sertnDeployment.sertnServiceManager);
        
        vm.startPrank(address(sm));

        sertnRegistrar = new SertnRegistrar();
        allocationManager.updateAVSMetadataURI(address(sm), "");
        // allocationManager.setAVSRegistrar(address(sm), sertnRegistrar);
        delete strat;
        strat.push(tokenToStrategy[address(mockToken)]);
        delete setParams;
        setParams.push(IAllocationManagerTypes.CreateSetParams({operatorSetId:0, strategies: strat}));
        allocationManager.createOperatorSets(address(sm), setParams);
        
        vm.stopPrank();

        addStrategy(address(mockToken));

        while (operators.length < OPERATOR_COUNT) {
            createAndAddOperator();
        }

        for (uint256 i = 0; i < OPERATOR_COUNT; i++) {
            mintMockTokens(operators[i], INITIAL_BALANCE);

            depositTokenIntoStrategy(operators[i], address(mockToken), DEPOSIT_AMOUNT);

            registerAsOperator(operators[i]);

            registerAsOperatorToAVS(operators[i]);
        }
    }
    function test_run() public {
        console2.log(allocationManager.getMemberCount(opSet[0]));
    }

}
