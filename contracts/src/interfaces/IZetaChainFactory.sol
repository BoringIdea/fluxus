// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title ZetaChain Factory Interface
/// @notice This interface is used to create Fluxus contracts on ZetaChain
interface IZetaChainFactory {
    /// @notice The address of the registry contract
    function registry() external view returns (address);

    /// @notice The address of the fee vault contract
    function feeVault() external view returns (address);

    /// @notice The address of the price contract
    function priceContract() external view returns (address);

    /// @notice The address of the uniswap router contract
    function uniswapRouter() external view returns (address);

    /// @notice The address of the universal contract
    function universalContract() external view returns (address);

    /// @notice Error thrown when an invalid gateway address is provided
    error InvalidGatewayAddress();

    /// @notice Error thrown when an invalid gas limit is provided
    error InvalidGasLimit();

    /// @notice Error thrown when a zero address is provided
    error ZeroAddress();

    /// @notice Error thrown when an unauthorized action is attempted
    error Unauthorized();

    /// @notice Event emitted when a Fluxus contract is created on ZetaChain
    /// @param creator The creator of the Fluxus
    /// @param fluxusAddress The address of the Fluxus
    /// @param priceAddress The address of the price contract
    /// @param name The name of the Fluxus
    /// @param symbol The symbol of the Fluxus
    /// @param initialPrice The initial price of the Fluxus
    /// @param maxSupply The maximum supply of the Fluxus
    /// @param maxPrice The maximum price of the Fluxus
    /// @param creatorFeePercent The creator fee percent of the Fluxus
    /// @param baseUri The base URI of the Fluxus
    event FluxusCreated(
        address indexed creator,
        address indexed fluxusAddress,
        address priceAddress,
        string name,
        string symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent,
        string baseUri
    );

    /// @notice Event emitted when the connected contract is set for a Fluxus
    /// @param fluxus The address of the Fluxus contract
    /// @param zrc20 The address of the ZRC20 contract
    /// @param connected The address of the connected contract
    event SetConnected(address indexed fluxus, address zrc20, address connected);

    /// @notice Event emitted when the gateway is set for a Fluxus
    /// @param fluxus The address of the Fluxus contract
    /// @param gateway The address of the gateway
    event SetGateway(address indexed fluxus, address gateway);

    /// @notice Event emitted when the gas limit is set for a Fluxus
    /// @param fluxus The address of the Fluxus contract
    /// @param gasLimit The gas limit value
    event SetGasLimit(address indexed fluxus, uint256 gasLimit);

    /// @notice Create a Fluxus contract on ZetaChain with cross-chain support
    /// @param name The name of the Fluxus
    /// @param symbol The symbol of the Fluxus
    /// @param initialPrice The initial price of the Fluxus
    /// @param maxSupply The maximum supply of the Fluxus
    /// @param maxPrice The maximum price of the Fluxus
    /// @param creatorFeePercent The creator fee percent of the Fluxus
    /// @param baseUri The base URI of the Fluxus
    /// @param gatewayAddress The address of the gateway contract
    /// @param gasLimit The gas limit for cross-chain operations
    /// @param supportMint Whether the contract supports minting
    /// @return The address of the created Fluxus contract
    function createFluxusZetaChainCrossChain(
        string memory name,
        string memory symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent,
        string memory baseUri,
        address gatewayAddress,
        uint256 gasLimit,
        bool supportMint
    ) external returns (address);

    /// @notice Set the connected contract for a Fluxus
    /// @param fluxus The address of the Fluxus contract
    /// @param zrc20 The address of the ZRC20 contract
    /// @param connected The address of the connected contract
    function setConnected(address payable fluxus, address zrc20, address connected) external;

    /// @notice Set the gateway address for a Fluxus
    /// @param fluxus The address of the Fluxus contract
    /// @param gateway The address of the gateway contract
    function setGateway(address payable fluxus, address gateway) external;

    /// @notice Set the gas limit for a Fluxus
    /// @param fluxus The address of the Fluxus contract
    /// @param gasLimit The gas limit for the contract
    function setGasLimit(address payable fluxus, uint256 gasLimit) external;
}
