// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title Factory Interface
/// @notice This interface is used to create Fluxus contracts
interface IFactory {
    /// @notice Error thrown when an invalid gateway address is provided
    error InvalidGatewayAddress();

    /// @notice Error thrown when an invalid gas limit is provided
    error InvalidGasLimit();

    /// @notice Error thrown when a zero address is provided
    error ZeroAddress();

    /// @notice Error thrown when an unauthorized action is attempted
    error Unauthorized();
    
    /// @notice The address of the registry contract
    function registry() external view returns (address);

    /// @notice The address of the fee vault contract
    function feeVault() external view returns (address);

    /// @notice The address of the price contract
    function priceContract() external view returns (address);

    /// @notice Event emitted when a Fluxus contract is created
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

    /// @notice Event emitted when the universal contract is set for a Fluxus
    /// @param fluxus The address of the Fluxus contract
    /// @param universal The address of the universal contract
    event SetUniversal(address indexed fluxus, address universal);

    /// @notice Event emitted when the gateway is set for a Fluxus
    /// @param fluxus The address of the Fluxus contract
    /// @param gateway The address of the gateway
    event SetGateway(address indexed fluxus, address gateway);

    /// @notice Event emitted when the gas limit is set for a Fluxus
    /// @param fluxus The address of the Fluxus contract
    /// @param gasLimit The gas limit value
    event SetGasLimit(address indexed fluxus, uint256 gasLimit);

    /// @notice Event emitted when a Fluxus CrossChain contract is created
    /// @param creator The creator of the Fluxus CrossChain
    /// @param fluxusAddress The address of the Fluxus CrossChain
    /// @param priceAddress The address of the price contract
    /// @param name The name of the Fluxus CrossChain
    /// @param symbol The symbol of the Fluxus CrossChain
    /// @param initialPrice The initial price of the Fluxus CrossChain
    /// @param maxSupply The maximum supply of the Fluxus CrossChain
    /// @param maxPrice The maximum price of the Fluxus CrossChain
    /// @param creatorFeePercent The creator fee percent of the Fluxus CrossChain
    /// @param baseUri The base URI of the Fluxus CrossChain
    /// @param gatewayAddress The address of the gateway contract
    /// @param gasLimit The gas limit for cross-chain operations
    /// @param supportMint Whether the contract supports minting
    event FluxusCrossChainCreated(
        address indexed creator,
        address indexed fluxusAddress,
        address priceAddress,
        string name,
        string symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent,
        string baseUri,
        address gatewayAddress,
        uint256 gasLimit,
        bool supportMint
    );

    /// @notice Create a Fluxus contract
    /// @param name The name of the Fluxus
    /// @param symbol The symbol of the Fluxus
    /// @param initialPrice The initial price of the Fluxus
    /// @param maxSupply The maximum supply of the Fluxus
    /// @param maxPrice The maximum price of the Fluxus
    /// @param creatorFeePercent The creator fee percent of the Fluxus
    /// @param baseUri The base URI of the Fluxus
    function createFluxus(
        string memory name,
        string memory symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent,
        string memory baseUri
    ) external returns (address);

    /// @notice Create a Fluxus CrossChain contract
    /// @param name The name of the Fluxus CrossChain
    /// @param symbol The symbol of the Fluxus CrossChain
    /// @param initialPrice The initial price of the Fluxus CrossChain
    /// @param maxSupply The maximum supply of the Fluxus CrossChain
    /// @param maxPrice The maximum price of the Fluxus CrossChain
    /// @param creatorFeePercent The creator fee percent of the Fluxus CrossChain
    /// @param baseUri The base URI of the Fluxus CrossChain
    /// @param gatewayAddress The address of the gateway contract
    /// @param gasLimit The gas limit for cross-chain operations
    /// @param supportMint Whether the contract supports minting
    function createFluxusCrossChain(
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

    /// @notice Calculate the address of a Fluxus contract
    /// @param name The name of the Fluxus
    /// @param symbol The symbol of the Fluxus
    /// @param initialPrice The initial price of the Fluxus
    /// @param maxSupply The maximum supply of the Fluxus
    /// @param maxPrice The maximum price of the Fluxus
    /// @param creatorFeePercent The creator fee percent of the Fluxus
    /// @param baseUri The base URI of the Fluxus
    /// @return The address of the Fluxus contract
    function calculateFluxusAddress(
        string memory name,
        string memory symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent,
        string memory baseUri
    ) external view returns (address);

    /// @notice Calculate the address of a Fluxus CrossChain contract
    /// @param name The name of the Fluxus CrossChain
    /// @param symbol The symbol of the Fluxus CrossChain
    /// @param initialPrice The initial price of the Fluxus CrossChain
    /// @param maxSupply The maximum supply of the Fluxus CrossChain
    /// @param maxPrice The maximum price of the Fluxus CrossChain
    /// @param creatorFeePercent The creator fee percent of the Fluxus CrossChain
    /// @param baseUri The base URI of the Fluxus CrossChain
    /// @param gatewayAddress The address of the gateway contract
    /// @param gasLimit The gas limit for cross-chain operations
    /// @return The address of the Fluxus CrossChain contract
    /// @param supportMint Whether the contract supports minting
    function calculateFluxusCrossChainAddress(
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
    ) external view returns (address);

    /// @notice Set the universal contract address for a Fluxus CrossChain contract
    /// @param fluxus The address of the Fluxus CrossChain contract
    /// @param universal The address of the universal contract
    function setUniversal(address payable fluxus, address universal) external;

    /// @notice Set the gateway contract address for a Fluxus CrossChain contract
    /// @param fluxus The address of the Fluxus CrossChain contract
    /// @param gateway The address of the gateway contract
    function setGateway(address payable fluxus, address gateway) external;

    /// @notice Set the gas limit for a Fluxus CrossChain contract
    /// @param fluxus The address of the Fluxus CrossChain contract
    /// @param gasLimit The gas limit for the contract
    function setGasLimit(address payable fluxus, uint256 gasLimit) external;
}