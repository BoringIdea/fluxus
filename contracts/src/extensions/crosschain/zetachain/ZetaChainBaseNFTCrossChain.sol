// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseNFT} from "../../../core/BaseNFT.sol";
import {UniversalNFTEvents} from "@zetachain/standard-contracts/contracts/nft/contracts/shared/UniversalNFTEvents.sol";
import {UniversalContract, MessageContext} from "@zetachain/protocol-contracts/contracts/zevm/interfaces/UniversalContract.sol";
import {IGatewayZEVM, RevertContext} from "@zetachain/protocol-contracts/contracts/zevm/interfaces/IGatewayZEVM.sol";
import {IWETH9} from "@zetachain/protocol-contracts/contracts/zevm/interfaces/IWZETA.sol";
import {GatewayZEVM, CallOptions, RevertOptions} from "@zetachain/protocol-contracts/contracts/zevm/GatewayZEVM.sol";
import {ERC721URIStorageUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import {SwapHelperLib} from "@zetachain/toolkit/contracts/SwapHelperLib.sol";
import {IZRC20} from "@zetachain/protocol-contracts/contracts/zevm/interfaces/IZRC20.sol";

contract ZetaChainBaseNFTCrossChain is UniversalContract, BaseNFT, ERC721URIStorageUpgradeable, UniversalNFTEvents {
    
    // Indicates this contract implements a Universal Contract
    bool public constant IS_UNIVERSAL = true;

    struct CrossChainStorage {
      // Address of the ZetaChain gateway contract
      GatewayZEVM gateway;

      // Address of the Uniswap Router for token swaps
      address uniswapRouter;

      // The amount of gas used when making cross-chain transfers
      uint256 gasLimitAmount;

      // Whether the contract supports minting.
      bool supportMint;

      // Mapping of connected ZRC-20 tokens to their respective contracts
      mapping(address => address) connected;

      // Whether the token is a cross-chain transfer
      mapping(uint256 tokenId => bool) isCrossChainTransfer;
    }

    struct InitParams {
        string name;
        string symbol;
        uint256 initialPrice;
        uint256 maxSupply;
        uint256 maxPrice;
        uint256 creatorFeePercent;
        address feeVault;
        address priceContract;
        address creator;
        string uri;
        address gatewayAddress;
        address uniswapRouterAddress;
        uint256 gasLimit;
        bool supportMint;
    }

    error TransferTokenFailed();
    error InvalidAddress();
    error InvalidGasLimit();
    error ApproveFailed();
    error ZeroMsgValue();
    error NotSupportMint();
    error CallFailed();

    modifier onlyCrossChainGateway() {
        _onlyCrossChainGateway();
        _;
    }

    function _onlyCrossChainGateway() internal view {
        CrossChainStorage storage $ = _getCrossChainStorageData();
        if (msg.sender != address($.gateway)) {
            revert Unauthorized();
        }
    }

    // keccak256(abi.encode(uint256(keccak256("fluxus.extensions.BaseNFTCrossChain.zetachain")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CROSS_CHAIN_STORAGE_LOCATION = 0xa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebf;

    function _getCrossChainStorageData() private pure returns (CrossChainStorage storage $) {
        assembly {
            $.slot := CROSS_CHAIN_STORAGE_LOCATION
        }
    }

    /**
     * @notice Sets the EVM gateway contract address.
     * @dev Can only be called by the contract owner.
     * @param gatewayAddress The address of the gateway contract.
     */
    function setGateway(address gatewayAddress) external onlyCreator {
        if (gatewayAddress == address(0)) revert InvalidAddress();
        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.gateway = GatewayZEVM(payable(gatewayAddress));
    }

    /**
     * @notice Sets the gas limit for cross-chain transfers.
     * @dev Can only be called by the contract owner.
     * @param gas New gas limit value.
     */
    function setGasLimit(uint256 gas) external onlyCreator {
        if (gas == 0) revert InvalidGasLimit();
        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.gasLimitAmount = gas;
    }

    /**
     * @notice Links a ZRC-20 gas token address to an NFT contract
     *         on the corresponding chain.
     * @dev Can only be called by the contract owner.
     * @param zrc20 Address of the ZRC-20 token.
     * @param contractAddress Address of the corresponding contract.
     */
    function setConnected(
        address zrc20,
        address contractAddress
    ) external onlyCreator {
        if (zrc20 == address(0)) revert InvalidAddress();
        if (contractAddress == address(0)) revert InvalidAddress();
        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.connected[zrc20] = contractAddress;
        emit SetConnected(zrc20, contractAddress);
    }

    function getZetaGateway() public view returns (address) {
        return address(_getCrossChainStorageData().gateway);
    }

    function uniswapRouter() public view returns (address) {
        return _getCrossChainStorageData().uniswapRouter;
    }

    function gasLimitAmount() public view returns (uint256) {
        return _getCrossChainStorageData().gasLimitAmount;
    }

    function connected(address zrc20) public view returns (address) {
        return _getCrossChainStorageData().connected[zrc20];
    }

    function isSupportMint() public view returns (bool) {
        CrossChainStorage storage $ = _getCrossChainStorageData();
        return $.supportMint;
    }

    /**
     * @notice Initializes the contract.
     * @dev Should be called during contract deployment.
     * @param params Struct containing all initialization parameters.
     */
    function __BaseNFTCrossChain_init(InitParams memory params) internal {
        // Initialize BaseNFT first
        __BaseNFT_init(
            params.name, 
            params.symbol, 
            params.initialPrice, 
            params.maxSupply, 
            params.maxPrice, 
            params.creatorFeePercent, 
            params.feeVault, 
            params.priceContract, 
            params.creator, 
            params.uri
        );

        // Initialize cross-chain specific parameters
        _initCrossChainParams(params.gatewayAddress, params.uniswapRouterAddress, params.gasLimit, params.supportMint);
    }

    /**
     * @notice Initialize cross-chain specific parameters.
     * @dev Internal function to avoid stack too deep error.
     */
    function _initCrossChainParams(
        address _gatewayAddress,
        address _uniswapRouterAddress,
        uint256 _gasLimit,
        bool _supportMint
    ) internal {
        if (_gatewayAddress == address(0) || _uniswapRouterAddress == address(0)) revert InvalidAddress();
        if (_gasLimit == 0) revert InvalidGasLimit();

        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.gateway = GatewayZEVM(payable(_gatewayAddress));
        $.uniswapRouter = _uniswapRouterAddress;
        $.gasLimitAmount = _gasLimit;
        $.supportMint = _supportMint;
    }


    /**
     * @notice Transfers an NFT to a connected chain.
     * @dev This function accepts native ZETA tokens as gas fees, which are swapped
     *      for the corresponding ZRC-20 gas token of the destination chain. The NFT is then
     *      transferred to the destination chain using the ZetaChain Gateway.
     * @param tokenId The ID of the NFT to transfer.
     * @param receiver Address of the recipient on the destination chain.
     * @param destination Address of the ZRC-20 gas token for the destination chain.
     */
    function transferCrossChain(
        uint256 tokenId,
        address receiver,
        address destination
    ) public payable {

        if (msg.value == 0) revert ZeroMsgValue();
        if (receiver == address(0)) revert InvalidAddress();

        string memory uri = tokenURI(tokenId);

        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.isCrossChainTransfer[tokenId] = true;

        _transfer(msg.sender, address(this), tokenId);

        emit TokenTransfer(receiver, destination, tokenId, uri);

        (address gasZrc20, uint256 gasFee) = IZRC20(destination)
            .withdrawGasFeeWithGasLimit($.gasLimitAmount);
        if (destination != gasZrc20) revert InvalidAddress();

        address wzeta = $.gateway.zetaToken();
        IWETH9(wzeta).deposit{value: msg.value}();
        if (!IWETH9(wzeta).approve($.uniswapRouter, msg.value)) {
            revert ApproveFailed();
        }
        
        SwapHelperLib.swapTokensForExactTokens(
            $.uniswapRouter,
            wzeta,
            gasFee,
            gasZrc20,
            msg.value
        );

        // uint256 remaining = msg.value > out ? msg.value - out : 0;
        // if (remaining > 0) {                 
        //     IWETH9(wzeta).withdraw(remaining);
        //     (bool success, ) = msg.sender.call{value: remaining}("");
        //     if (!success) revert TransferTokenFailed();
        // }

        bytes memory message = abi.encode(
            receiver,
            tokenId,
            uri,
            0,
            msg.sender
        );
        CallOptions memory callOptions = CallOptions($.gasLimitAmount, false);

        RevertOptions memory revertOptions = RevertOptions(
            address(this),
            true,
            address(0),
            abi.encode(tokenId, uri, msg.sender),
            $.gasLimitAmount
        );

        if (!IZRC20(gasZrc20).approve(address($.gateway), gasFee)) {
            revert ApproveFailed();
        }

        try $.gateway.call(
            abi.encodePacked($.connected[destination]),
            destination,
            message,
            callOptions,
            revertOptions
        ) {
            // Call succeeded
        } catch {
            revert CallFailed();
        }
    }

    function onCall(MessageContext calldata context, bytes calldata message) external payable virtual override onlyCrossChainGateway {
        // Handle calls with native tokens if needed, otherwise revert
        revert Unauthorized();
    }

    /**
     * @notice Handles cross-chain NFT transfers with ZRC20 tokens.
     * @dev This function is called by the Gateway contract upon receiving a message.
     *      If the destination is ZetaChain, mint an NFT and set its URI.
     *      If the destination is another chain, swap the gas token for the corresponding
     *      ZRC-20 token and use the Gateway to send a message to mint an NFT on the
     *      destination chain.
     * @param context Message context metadata.
     * @param zrc20 ZRC-20 token address.
     * @param amount Amount of token provided.
     * @param message Encoded payload containing NFT metadata.
     */
    function onCall(
        MessageContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    ) external override onlyCrossChainGateway {
        CrossChainStorage storage $ = _getCrossChainStorageData();
        if (context.senderEVM != $.connected[zrc20]) revert Unauthorized();

        (
            address destination,
            address receiver,
            uint256 tokenId,
            string memory uri,
            address sender
        ) = abi.decode(message, (address, address, uint256, string, address));

        if (destination == address(0)) {
            if ($.isCrossChainTransfer[tokenId]) {
                _transfer(address(this), receiver, tokenId);
                $.isCrossChainTransfer[tokenId] = false;
            } else {
                _safeMint(receiver, tokenId);
                _setTokenURI(tokenId, uri);
            }
            emit TokenTransferReceived(receiver, tokenId, uri);
        } else {
            (address gasZrc20, uint256 gasFee) = IZRC20(destination)
                .withdrawGasFeeWithGasLimit($.gasLimitAmount);
            if (destination != gasZrc20) revert InvalidAddress();

            uint256 out = SwapHelperLib.swapExactTokensForTokens(
                $.uniswapRouter,
                zrc20,
                amount,
                destination,
                0
            );

            if (!IZRC20(destination).approve(address($.gateway), out)) {
                revert ApproveFailed();
            }
            $.gateway.withdrawAndCall(
                abi.encodePacked($.connected[destination]),
                out - gasFee,
                destination,
                abi.encode(receiver, tokenId, uri, out - gasFee, sender),
                CallOptions($.gasLimitAmount, false),
                RevertOptions(
                    address(this),
                    true,
                    address(0),
                    abi.encode(tokenId, uri, sender),
                    0
                )
            );
        }
        emit TokenTransferToDestination(receiver, destination, tokenId, uri);
    }

    /**
     * @notice Handles a cross-chain call failure and reverts the NFT transfer.
     * @param context Metadata about the failed call.
     */
    function onRevert(RevertContext calldata context) external onlyCrossChainGateway {
        (uint256 tokenId, string memory uri, address sender) = abi.decode(
            context.revertMessage,
            (uint256, string, address)
        );

        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.isCrossChainTransfer[tokenId] = false;
        _transfer(address(this), sender, tokenId);
        emit TokenTransferReverted(sender, tokenId, uri);
    }

    /**
     * @notice Gets the token URI for a given token ID.
     * @param tokenId The ID of the token.
     * @return The token URI as a string.
     */
    function tokenURI(
        uint256 tokenId
    )
        public
        view
        virtual
        override(BaseNFT, ERC721URIStorageUpgradeable)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function isCrossChainTransfer(uint256 tokenId) public view returns (bool) {
        CrossChainStorage storage $ = _getCrossChainStorageData();
        return $.isCrossChainTransfer[tokenId];
    }

    function _baseURI() internal view virtual override(BaseNFT, ERC721Upgradeable) returns (string memory) {
        return super._baseURI();
    }

    function _increaseBalance(address account, uint128 amount) internal override(BaseNFT, ERC721Upgradeable) {
        super._increaseBalance(account, amount);
    }

    function _update(address to, uint256 tokenId, address auth) internal override(BaseNFT, ERC721Upgradeable) returns (address) {
        return super._update(to, tokenId, auth);
    }

    /**
     * @notice Checks if the contract supports a specific interface.
     * @param interfaceId The interface identifier to check.
     * @return True if the interface is supported, false otherwise.
     */
    function supportsInterface(
        bytes4 interfaceId
    )
        public
        view
        virtual
        override(BaseNFT, ERC721URIStorageUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}