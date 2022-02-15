// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

interface IERC20 {
    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address to, uint256 amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) external returns (bool);

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract BalanceGetter {
    constructor() {}

    function readBalances(address token, address[] calldata addresses) public virtual returns (uint[] memory) {
        IERC20 erc20 = IERC20(token);
        uint[] memory balances = new uint[](addresses.length);

        for (uint i = 0; i < addresses.length; i++) {
            balances[i] = erc20.balanceOf(addresses[i]);
        }

        // return 16;
        return balances;
    }
}