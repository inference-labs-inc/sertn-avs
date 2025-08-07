// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;

import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {ISignatureUtilsMixin} from "@eigenlayer/contracts/interfaces/ISignatureUtilsMixin.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title DelegationManagerMock
 * @notice Minimal mock implementation of IDelegationManager for testing
 */
contract DelegationManagerMock is IDelegationManager {
    // Minimal state to track basic delegation relationships
    mapping(address => address) public override delegatedTo;
    mapping(address => bool) public registeredOperators;

    // ============ TEST STUFF ============
    bool public isAllUsersAreOperators = true;

    function setAllUsersAreOperators(bool _isAllUsersAreOperators) external {
        isAllUsersAreOperators = _isAllUsersAreOperators;
    }

    // ============ INITIALIZATION ============

    function initialize(address, uint256) external pure override {
        // Mock implementation - do nothing
    }

    // ============ OPERATOR REGISTRATION ============

    function registerAsOperator(address operator, uint32, string calldata) external override {
        registeredOperators[operator] = true;
        delegatedTo[operator] = operator; // Operators delegate to themselves
        emit OperatorRegistered(operator, address(0));
    }

    function modifyOperatorDetails(address, address) external pure override {
        // Mock implementation - do nothing
    }

    function updateOperatorMetadataURI(address, string calldata) external pure override {
        // Mock implementation - do nothing
    }

    // ============ DELEGATION ============

    function delegateTo(address operator, SignatureWithExpiry memory, bytes32) external override {
        require(registeredOperators[operator], "Operator not registered");
        delegatedTo[msg.sender] = operator;
        emit StakerDelegated(msg.sender, operator);
    }

    function undelegate(address) external pure override returns (bytes32[] memory) {
        // Mock implementation - return empty array
        return new bytes32[](0);
    }

    function redelegate(
        address,
        SignatureWithExpiry memory,
        bytes32
    ) external pure override returns (bytes32[] memory) {
        // Mock implementation - return empty array
        return new bytes32[](0);
    }

    // ============ WITHDRAWALS ============

    function queueWithdrawals(
        QueuedWithdrawalParams[] calldata
    ) external pure override returns (bytes32[] memory) {
        // Mock implementation - return empty array
        return new bytes32[](0);
    }

    function completeQueuedWithdrawal(
        Withdrawal calldata,
        IERC20[] calldata,
        bool
    ) external pure override {
        // Mock implementation - do nothing
    }

    function completeQueuedWithdrawals(
        Withdrawal[] calldata,
        IERC20[][] calldata,
        bool[] calldata
    ) external pure override {
        // Mock implementation - do nothing
    }

    function cumulativeWithdrawalsQueued(address staker) external view returns (uint256) {
        // Mock implementation - return 0
        return 0;
    }

    function delegationApproverSaltIsSpent(
        address _delegationApprover,
        bytes32 salt
    ) external view returns (bool) {
        // Mock implementation - return false
        return false;
    }

    // ============ STRATEGY MANAGER INTERACTIONS ============

    function increaseDelegatedShares(address, IStrategy, uint256, uint256) external pure override {
        // Mock implementation - do nothing
    }

    function decreaseDelegatedShares(address, uint256, uint64) external pure override {
        // Mock implementation - do nothing
    }

    function slashOperatorShares(address, IStrategy, uint64, uint64) external pure override {
        // Mock implementation - do nothing
    }

    function beaconChainETHStrategy() external view returns (IStrategy) {
        return IStrategy(address(0));
    }

    // ============ VIEW FUNCTIONS ============

    function isDelegated(address staker) external view override returns (bool) {
        return delegatedTo[staker] != address(0);
    }

    function isOperator(address operator) external view override returns (bool) {
        return isAllUsersAreOperators || registeredOperators[operator];
    }

    function delegationApprover(address) external pure override returns (address) {
        return address(0);
    }

    function depositScalingFactor(address, IStrategy) external pure override returns (uint256) {
        return 1e18; // Return 1 as default scaling factor
    }

    function getOperatorShares(
        address,
        IStrategy[] memory strategies
    ) external pure override returns (uint256[] memory) {
        return new uint256[](strategies.length);
    }

    function getOperatorsShares(
        address[] memory operators,
        IStrategy[] memory strategies
    ) external pure override returns (uint256[][] memory) {
        uint256[][] memory shares = new uint256[][](operators.length);
        for (uint256 i = 0; i < operators.length; i++) {
            shares[i] = new uint256[](strategies.length);
        }
        return shares;
    }

    function getSlashableSharesInQueue(
        address,
        IStrategy
    ) external pure override returns (uint256) {
        return 0;
    }

    function getWithdrawableShares(
        address,
        IStrategy[] memory strategies
    ) external pure override returns (uint256[] memory, uint256[] memory) {
        uint256[] memory withdrawableShares = new uint256[](strategies.length);
        uint256[] memory depositShares = new uint256[](strategies.length);
        return (withdrawableShares, depositShares);
    }

    function getDepositedShares(
        address
    ) external pure override returns (IStrategy[] memory, uint256[] memory) {
        return (new IStrategy[](0), new uint256[](0));
    }

    function queuedWithdrawals(bytes32) external pure override returns (Withdrawal memory) {
        return
            Withdrawal({
                staker: address(0),
                delegatedTo: address(0),
                withdrawer: address(0),
                nonce: 0,
                startBlock: 0,
                strategies: new IStrategy[](0),
                scaledShares: new uint256[](0)
            });
    }

    function getQueuedWithdrawal(
        bytes32
    ) external pure override returns (Withdrawal memory, uint256[] memory) {
        Withdrawal memory withdrawal = Withdrawal({
            staker: address(0),
            delegatedTo: address(0),
            withdrawer: address(0),
            nonce: 0,
            startBlock: 0,
            strategies: new IStrategy[](0),
            scaledShares: new uint256[](0)
        });
        return (withdrawal, new uint256[](0));
    }

    function getQueuedWithdrawals(
        address
    ) external pure override returns (Withdrawal[] memory, uint256[][] memory) {
        return (new Withdrawal[](0), new uint256[][](0));
    }

    function getQueuedWithdrawalRoots(address) external pure override returns (bytes32[] memory) {
        return new bytes32[](0);
    }

    function convertToDepositShares(
        address,
        IStrategy[] memory strategies,
        uint256[] memory
    ) external pure override returns (uint256[] memory) {
        return new uint256[](strategies.length);
    }

    function calculateWithdrawalRoot(Withdrawal memory) external pure override returns (bytes32) {
        return bytes32(0);
    }

    function minWithdrawalDelayBlocks() external pure override returns (uint32) {
        return 0;
    }

    function calculateDelegationApprovalDigestHash(
        address,
        address,
        address,
        bytes32,
        uint256
    ) external pure override returns (bytes32) {
        return bytes32(0);
    }

    // ============ SIGNATURE UTILS MIXIN ============

    function domainSeparator() external pure override returns (bytes32) {
        return bytes32(0);
    }

    function version() external pure override returns (string memory) {
        return "v1.0.0-mock";
    }

    function DELEGATION_APPROVAL_TYPEHASH() external view returns (bytes32) {
        return
            keccak256(
                "DelegationApproval(address staker,address operator,address delegationApprover,bytes32 salt,uint256 expiry)"
            );
    }
}
