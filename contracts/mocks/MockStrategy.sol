// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@eigenlayer/contracts/interfaces/IStrategy.sol";

contract MockStrategy is IStrategy {
    IERC20 public override underlyingToken;
    uint256 public override totalShares;
    mapping(address => uint256) public userShares;
    uint256 public constant EXCHANGE_RATE = 1e18; // 1:1 exchange rate for simplicity

    constructor(IERC20 _underlyingToken) {
        underlyingToken = _underlyingToken;
        emit StrategyTokenSet(_underlyingToken, 18); // Assuming 18 decimals for simplicity
    }

    function deposit(IERC20 token, uint256 amount) external override returns (uint256) {
        require(token == underlyingToken, "Invalid token");
        uint256 newShares = amount;
        totalShares += newShares;
        userShares[msg.sender] += newShares;
        emit ExchangeRateEmitted(EXCHANGE_RATE);
        return newShares;
    }

    function withdraw(address recipient, IERC20 token, uint256 amountShares) external override {
        require(token == underlyingToken, "Invalid token");
        require(userShares[msg.sender] >= amountShares, "Insufficient shares");
        userShares[msg.sender] -= amountShares;
        totalShares -= amountShares;
        underlyingToken.transfer(recipient, amountShares);
    }

    function sharesToUnderlying(uint256 amountShares) external pure override returns (uint256) {
        return amountShares;
    }

    function underlyingToShares(uint256 amountUnderlying) external pure override returns (uint256) {
        return amountUnderlying;
    }

    function userUnderlying(address user) external view override returns (uint256) {
        return userShares[user];
    }

    function shares(address user) external view override returns (uint256) {
        return userShares[user];
    }

    function sharesToUnderlyingView(uint256 amountShares) external pure override returns (uint256) {
        return amountShares;
    }

    function underlyingToSharesView(uint256 amountUnderlying) external pure override returns (uint256) {
        return amountUnderlying;
    }

    function userUnderlyingView(address user) external view override returns (uint256) {
        return userShares[user];
    }

    function explanation() external pure override returns (string memory) {
        return "Mock Strategy for testing purposes";
    }
}