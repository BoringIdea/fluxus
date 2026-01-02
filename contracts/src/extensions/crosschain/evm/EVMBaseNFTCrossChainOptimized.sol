// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseNFT} from "../../../core/BaseNFT.sol";
import {GatewayEVM, MessageContext, RevertContext} from "@zetachain/protocol-contracts/contracts/evm/GatewayEVM.sol";
import {UniversalNFTEvents} from "@zetachain/standard-contracts/contracts/nft/contracts/shared/UniversalNFTEvents.sol";
import {ERC721URIStorageUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import { EVMCrossChainLib } from "./libraries/EVMCrossChainLib.sol";

contract EVMBaseNFTCrossChainOptimized is BaseNFT, ERC721URIStorageUpgradeable, UniversalNFTEvents {
    
    using EVMCrossChainLib for *;

    struct CrossChainStorage {
        GatewayEVM gateway;
        address universal;
        uint256 gasLimitAmount;
        bool supportMint;
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

    bytes32 private constant CROSS_CHAIN_STORAGE_LOCATION = 0x9f9d9c9b9a99989796959493929190908f8e8d8c8b8a89888786858483828180;

    function _getCrossChainStorageData() internal pure returns (CrossChainStorage storage $) {
        assembly {
            $.slot := CROSS_CHAIN_STORAGE_LOCATION
        }
    }

    modifier onlyGateway() {
        _onlyGateway();
        _;
    }

    function _onlyGateway() internal view {
        if (msg.sender != address(_getCrossChainStorageData().gateway)) revert Unauthorized();
    }

    function gateway() public view returns (address) {
        return address(_getCrossChainStorageData().gateway);
    }

    function universal() public view returns (address) {
        return _getCrossChainStorageData().universal;
    }

    function gasLimit() public view returns (uint256) {
        return _getCrossChainStorageData().gasLimitAmount;
    }

    function setUniversal(address contractAddress) external onlyOwner {
        if (contractAddress == address(0)) revert InvalidAddress();
        _getCrossChainStorageData().universal = contractAddress;
        emit SetUniversal(contractAddress);
    }

    function setGateway(address gatewayAddress) external onlyOwner {
        if (gatewayAddress == address(0)) revert InvalidAddress();
        _getCrossChainStorageData().gateway = GatewayEVM(gatewayAddress);
    }

    function setGasLimit(uint256 _gasLimit) external onlyOwner {
        if (_gasLimit == 0) revert InvalidGasLimit();
        _getCrossChainStorageData().gasLimitAmount = _gasLimit;
    }

    function isSupportMint() public view returns (bool) {
        return _getCrossChainStorageData().supportMint;
    }

    function isCrossChainTransfer(uint256 tokenId) public view returns (bool) {
        return _getCrossChainStorageData().isCrossChainTransfer[tokenId];
    }

    function __BaseNFTCrossChain_init(InitParams memory params) internal {
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
        _initCrossChainParams(params.gatewayAddress, params.gasLimit, params.supportMint);
    }

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

    function transferCrossChain(
        uint256 tokenId,
        address receiver,
        address destination
    ) external payable virtual {
        string memory uri = tokenURI(tokenId);
        CrossChainStorage storage $ = _getCrossChainStorageData();
        
        _transfer(_msgSender(), address(this), tokenId);
        $.isCrossChainTransfer[tokenId] = true;

        EVMCrossChainLib.CrossChainParams memory params = EVMCrossChainLib.CrossChainParams({
            gateway: $.gateway,
            universal: $.universal,
            gasLimitAmount: $.gasLimitAmount,
            isCrossChainTransfer: $.isCrossChainTransfer[tokenId]
        });

        EVMCrossChainLib.transferCrossChain(
            params,
            tokenId,
            receiver,
            destination,
            uri,
            msg.sender
        );
    }

    function onCall(
        MessageContext calldata context,
        bytes calldata message
    ) external payable onlyGateway returns (bytes4) {
        CrossChainStorage storage $ = _getCrossChainStorageData();
        
        EVMCrossChainLib.CrossChainParams memory params = EVMCrossChainLib.CrossChainParams({
            gateway: $.gateway,
            universal: $.universal,
            gasLimitAmount: $.gasLimitAmount,
            isCrossChainTransfer: false // Will be updated based on tokenId
        });

        (
            address receiver,
            uint256 tokenId,
            string memory uri,
            uint256 gasAmount,
            address sender
        ) = abi.decode(message, (address, uint256, string, uint256, address));

        // Update params with correct token-specific data
        params.isCrossChainTransfer = $.isCrossChainTransfer[tokenId];
        
        (bool isTransferBack, , , , , ) = EVMCrossChainLib.handleCrossChainCall(
            params,
            context.sender,
            message
        );

        if (isTransferBack) {
            _transfer(address(this), receiver, tokenId);
            $.isCrossChainTransfer[tokenId] = false;
        } else {
            _safeMint(receiver, tokenId);
            _setTokenURI(tokenId, uri);
        }

        EVMCrossChainLib.handleGasTransfer(gasAmount, sender);
        emit TokenTransferReceived(receiver, tokenId, uri);
        return "";
    }

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

    function _baseURI() internal view virtual override(BaseNFT, ERC721Upgradeable) returns (string memory) {
        return super._baseURI();
    }

    function _increaseBalance(address account, uint128 amount) internal virtual override(BaseNFT, ERC721Upgradeable) {
        super._increaseBalance(account, amount);
    }

    function _update(address to, uint256 tokenId, address auth) internal virtual override(BaseNFT, ERC721Upgradeable) returns (address) {
        return super._update(to, tokenId, auth);
    }

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
