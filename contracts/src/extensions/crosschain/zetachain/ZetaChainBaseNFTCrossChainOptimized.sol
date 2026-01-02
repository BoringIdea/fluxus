// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseNFT} from "../../../core/BaseNFT.sol";
import {UniversalNFTEvents} from "@zetachain/standard-contracts/contracts/nft/contracts/shared/UniversalNFTEvents.sol";
import {UniversalContract} from "@zetachain/protocol-contracts/contracts/zevm/interfaces/UniversalContract.sol";
import {IGatewayZEVM, MessageContext, RevertContext} from "@zetachain/protocol-contracts/contracts/zevm/interfaces/IGatewayZEVM.sol";
import {GatewayZEVM} from "@zetachain/protocol-contracts/contracts/zevm/GatewayZEVM.sol";
import { ZetaChainCrossChainLib } from "./libraries/ZetaChainCrossChainLib.sol";

contract ZetaChainBaseNFTCrossChainOptimized is UniversalContract, BaseNFT, UniversalNFTEvents {
    
    using ZetaChainCrossChainLib for *;
    
    bool public constant IS_UNIVERSAL = true;

    struct CrossChainStorage {
        GatewayZEVM gateway;
        address uniswapRouter;
        uint256 gasLimitAmount;
        bool supportMint;
        mapping(address => address) connected;
        mapping(uint256 tokenId => bool) isCrossChainTransfer;
    }

    error InvalidAddress();
    error InvalidGasLimit();

    bytes32 private constant CROSS_CHAIN_STORAGE_LOCATION = 0xa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebf;

    function _getCrossChainStorageData() internal pure returns (CrossChainStorage storage $) {
        assembly {
            $.slot := CROSS_CHAIN_STORAGE_LOCATION
        }
    }

    modifier onlyCrossChainGateway() {
        _onlyCrossChainGateway();
        _;
    }

    function _onlyCrossChainGateway() internal view {
        if (msg.sender != address(_getCrossChainStorageData().gateway)) revert Unauthorized();
    }

    function setGateway(address gatewayAddress) external onlyCreator {
        if (gatewayAddress == address(0)) revert InvalidAddress();
        _getCrossChainStorageData().gateway = GatewayZEVM(payable(gatewayAddress));
    }

    function setGasLimit(uint256 gas) external onlyCreator {
        if (gas == 0) revert InvalidGasLimit();
        _getCrossChainStorageData().gasLimitAmount = gas;
    }

    function setConnected(address zrc20, address contractAddress) external onlyOwner {
        if (zrc20 == address(0) || contractAddress == address(0)) revert InvalidAddress();
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
        return _getCrossChainStorageData().supportMint;
    }

    function isCrossChainTransfer(uint256 tokenId) public view returns (bool) {
        return _getCrossChainStorageData().isCrossChainTransfer[tokenId];
    }

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

    function transferCrossChain(
        uint256 tokenId,
        address receiver,
        address destination
    ) public payable {
        string memory uri = tokenURI(tokenId);
        CrossChainStorage storage $ = _getCrossChainStorageData();
        
        _transfer(msg.sender, address(this), tokenId);
        $.isCrossChainTransfer[tokenId] = true;

        ZetaChainCrossChainLib.CrossChainParams memory params = ZetaChainCrossChainLib.CrossChainParams({
            gateway: $.gateway,
            uniswapRouter: $.uniswapRouter,
            gasLimitAmount: $.gasLimitAmount,
            connectedContract: $.connected[destination],
            isCrossChainTransfer: $.isCrossChainTransfer[tokenId]
        });

        ZetaChainCrossChainLib.transferCrossChain(
            params,
            tokenId,
            receiver,
            destination,
            uri,
            msg.sender
        );
    }

    function onCall(MessageContext calldata context, bytes calldata message) external payable virtual override onlyCrossChainGateway {
        // Handle calls with native tokens if needed, otherwise revert
        revert Unauthorized();
    }

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
        sender; // Silence unused variable warning

        ZetaChainCrossChainLib.CrossChainParams memory params = ZetaChainCrossChainLib.CrossChainParams({
            gateway: $.gateway,
            uniswapRouter: $.uniswapRouter,
            gasLimitAmount: $.gasLimitAmount,
            connectedContract: $.connected[destination],
            isCrossChainTransfer: $.isCrossChainTransfer[tokenId]
        });
        
        bool isTransferBack = ZetaChainCrossChainLib.handleCrossChainCall(
            params,
            zrc20,
            amount,
            message
        );

        if (destination == address(0)) {
            if (isTransferBack) {
                _transfer(address(this), receiver, tokenId);
                $.isCrossChainTransfer[tokenId] = false;
            } else {
                _safeMint(receiver, tokenId);
            }
            emit TokenTransferReceived(receiver, tokenId, uri);
        }
        emit TokenTransferToDestination(receiver, destination, tokenId, uri);
    }

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

    function supportsInterface(bytes4 interfaceId) public view virtual override(BaseNFT) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
