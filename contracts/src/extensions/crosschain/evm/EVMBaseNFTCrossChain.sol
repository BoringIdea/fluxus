// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseNFT} from "../../../core/BaseNFT.sol";
import {GatewayEVM, RevertOptions, MessageContext, RevertContext} from "@zetachain/protocol-contracts/contracts/evm/GatewayEVM.sol";
import {UniversalNFTEvents} from "@zetachain/standard-contracts/contracts/nft/contracts/shared/UniversalNFTEvents.sol";
import {ERC721URIStorageUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";

contract EVMBaseNFTCrossChain is BaseNFT, ERC721URIStorageUpgradeable, UniversalNFTEvents {

    struct CrossChainStorage {
      // Address of the EVM gateway contract
      GatewayEVM gateway;

      // The address of the Universal NFT contract on ZetaChain. This contract serves
      // as a key component for handling all cross-chain transfers while also functioning
      // as an ERC-721 Universal NFT.
      address universal;

      // The amount of gas used when making cross-chain transfers
      uint256 gasLimitAmount;

      // Whether the contract supports minting.
      bool supportMint;

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
        uint256 gasLimit;
        bool supportMint;
    }

    error InvalidAddress();
    error Unauthorized();
    error InvalidGasLimit();
    error GasTokenTransferFailed();
    error NotSupportMint();
    error CallFailed();

    modifier onlyGateway() {
        _onlyGateway();
        _;
    }

    function _onlyGateway() internal view {
        CrossChainStorage storage $ = _getCrossChainStorageData();
        if (msg.sender != address($.gateway)) {
            revert Unauthorized();
        }
    }

    // keccak256(abi.encode(uint256(keccak256("fluxus.extensions.BaseNFTCrossChain")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CROSS_CHAIN_STORAGE_LOCATION = 0x9f9d9c9b9a99989796959493929190908f8e8d8c8b8a89888786858483828180;

    function _getCrossChainStorageData() private pure returns (CrossChainStorage storage $) {
        assembly {
            $.slot := CROSS_CHAIN_STORAGE_LOCATION
        }
    }

    /**
     * @notice Gets the gateway address.
     * @return The address of the gateway contract.
     */
    function gateway() public view returns (address) {
        CrossChainStorage storage $ = _getCrossChainStorageData();
        return address($.gateway);
    }

    /**
     * @notice Gets the universal contract address.
     * @return The address of the universal contract.
     */
    function universal() public view returns (address) {
        CrossChainStorage storage $ = _getCrossChainStorageData();
        return $.universal;
    }

    /**
     * @notice Gets the gas limit amount for cross-chain operations.
     * @return The gas limit amount.
     */
    function gasLimit() public view returns (uint256) {
        CrossChainStorage storage $ = _getCrossChainStorageData();
        return $.gasLimitAmount;
    }

    /**
     * @notice Sets the universal contract address.
     * @dev Can only be called by the contract owner.
     * @param contractAddress The address of the universal contract.
     */
    function setUniversal(address contractAddress) external onlyOwner {
        if (contractAddress == address(0)) revert InvalidAddress();
        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.universal = contractAddress;
        emit SetUniversal(contractAddress);
    }

    /**
     * @notice Sets the EVM gateway contract address.
     * @dev Can only be called by the contract owner.
     * @param gatewayAddress The address of the gateway contract.
     */
    function setGateway(address gatewayAddress) external onlyOwner {
        if (gatewayAddress == address(0)) revert InvalidAddress();
        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.gateway = GatewayEVM(gatewayAddress);
    }

    /**
     * @notice Sets the gas limit for cross-chain transfers.
     * @dev Can only be called by the contract owner.
     * @param _gasLimit New gas limit value.
     */
    function setGasLimit(uint256 _gasLimit) external onlyOwner {
        if (_gasLimit == 0) revert InvalidGasLimit();
        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.gasLimitAmount = _gasLimit;
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
        _initCrossChainParams(params.gatewayAddress, params.gasLimit, params.supportMint);
    }

    /**
     * @notice Initialize cross-chain specific parameters.
     * @dev Internal function to avoid stack too deep error.
     */
    function _initCrossChainParams(
        address _gatewayAddress,
        uint256 _gasLimit,
        bool _supportMint
    ) internal {
        if (_gatewayAddress == address(0)) revert InvalidAddress();
        if (_gasLimit == 0) revert InvalidGasLimit();

        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.gateway = GatewayEVM(_gatewayAddress);
        $.gasLimitAmount = _gasLimit;
        $.supportMint = _supportMint;
    }

    /**
     * @notice Transfers an NFT to another chain.
     * @dev Burns the NFT locally, then uses the Gateway to send a message to
     *      mint the same NFT on the destination chain. If the destination is the zero
     *      address, transfers the NFT to ZetaChain.
     * @param tokenId The ID of the NFT to transfer.
     * @param receiver The address on the destination chain that will receive the NFT.
     * @param destination The ZRC-20 address of the gas token of the destination chain.
     */
    function transferCrossChain(
        uint256 tokenId,
        address receiver,
        address destination
    ) external payable virtual {
        if (receiver == address(0)) revert InvalidAddress();

        string memory uri = tokenURI(tokenId);

        CrossChainStorage storage $ = _getCrossChainStorageData();
        $.isCrossChainTransfer[tokenId] = true;

        _transfer(_msgSender(), address(this), tokenId);

        bytes memory message = abi.encode(
            destination,
            receiver,
            tokenId,
            uri,
            msg.sender
        );

        emit TokenTransfer(destination, receiver, tokenId, uri);

        if (destination == address(0)) {
            try $.gateway.call(
                $.universal,
                message,
                RevertOptions(address(this), false, address(0), message, 0)
            ) {
                // Call succeeded
            } catch {
                revert CallFailed();
            }
        } else {
            $.gateway.depositAndCall{value: msg.value}(
                $.universal,
                message,
                RevertOptions(
                    address(this),
                    true,
                    address(0),
                    abi.encode(receiver, tokenId, uri, msg.sender),
                    $.gasLimitAmount
                )
            );
        }
    }

    /**
     * @notice Mint an NFT in response to an incoming cross-chain transfer.
     * @dev Called by the Gateway upon receiving a message.
     * @param context The message context.
     * @param message The encoded message containing information about the NFT.
     * @return A constant indicating the function was successfully handled.
     */
    function onCall(
        MessageContext calldata context,
        bytes calldata message
    ) external payable onlyGateway returns (bytes4) {
        CrossChainStorage storage $ = _getCrossChainStorageData();
        if (context.sender != $.universal) revert Unauthorized();

        (
            address receiver,
            uint256 tokenId,
            string memory uri,
            uint256 gasAmount,
            address sender
        ) = abi.decode(message, (address, uint256, string, uint256, address));

        // If the token is a cross-chain transfer, transfer it to the receiver
        // Otherwise, mint the token to the receiver
        if ($.isCrossChainTransfer[tokenId]) {
            _transfer(address(this), receiver, tokenId);
            $.isCrossChainTransfer[tokenId] = false;
        } else {
            _safeMint(receiver, tokenId);
            _setTokenURI(tokenId, uri);
        }
        if (gasAmount > 0) {
            if (sender == address(0)) revert InvalidAddress();
            (bool success, ) = payable(sender).call{value: gasAmount}("");
            if (!success) revert GasTokenTransferFailed();
        }
        emit TokenTransferReceived(receiver, tokenId, uri);
        return "";
    }

    /**
     * @notice Mint an NFT and send it back to the sender if a cross-chain transfer fails.
     * @dev Called by the Gateway if a call fails.
     * @param context The revert context containing metadata and revert message.
     */
    function onRevert(RevertContext calldata context) external onlyGateway {
        (, uint256 tokenId, string memory uri, address sender) = abi.decode(
            context.revertMessage,
            (address, uint256, string, address)
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

    receive() external payable {}
}