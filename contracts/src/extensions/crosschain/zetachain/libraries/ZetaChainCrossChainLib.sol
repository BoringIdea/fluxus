// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IZRC20} from "@zetachain/protocol-contracts/contracts/zevm/interfaces/IZRC20.sol";
import {IWETH9} from "@zetachain/protocol-contracts/contracts/zevm/interfaces/IWZETA.sol";
import {GatewayZEVM, CallOptions, RevertOptions} from "@zetachain/protocol-contracts/contracts/zevm/GatewayZEVM.sol";
import {SwapHelperLib} from "@zetachain/toolkit/contracts/SwapHelperLib.sol";

library ZetaChainCrossChainLib {
    error InvalidAddress();
    error ApproveFailed();
    error ZeroMsgValue();
    error CallFailed();

    struct CrossChainParams {
        GatewayZEVM gateway;
        address uniswapRouter;
        uint256 gasLimitAmount;
        address connectedContract;
        bool isCrossChainTransfer;
    }

    function transferCrossChain(
        CrossChainParams memory params,
        uint256 tokenId,
        address receiver,
        address destination,
        string memory uri,
        address sender
    ) external {
        if (msg.value == 0) revert ZeroMsgValue();
        if (receiver == address(0)) revert InvalidAddress();

        emit TokenTransfer(receiver, destination, tokenId, uri);

        (address gasZrc20, uint256 gasFee) = IZRC20(destination)
            .withdrawGasFeeWithGasLimit(params.gasLimitAmount);
        if (destination != gasZrc20) revert InvalidAddress();

        address wzeta = params.gateway.zetaToken();
        IWETH9(wzeta).deposit{value: msg.value}();
        if (!IWETH9(wzeta).approve(params.uniswapRouter, msg.value)) {
            revert ApproveFailed();
        }
        
        SwapHelperLib.swapTokensForExactTokens(
            params.uniswapRouter,
            wzeta,
            gasFee,
            gasZrc20,
            msg.value
        );

        bytes memory message = abi.encode(
            receiver,
            tokenId,
            uri,
            0,
            sender
        );
        CallOptions memory callOptions = CallOptions(params.gasLimitAmount, false);

        RevertOptions memory revertOptions = RevertOptions(
            address(this),
            true,
            address(0),
            abi.encode(tokenId, uri, sender),
            params.gasLimitAmount
        );

        if (!IZRC20(gasZrc20).approve(address(params.gateway), gasFee)) {
            revert ApproveFailed();
        }

        try params.gateway.call(
            abi.encodePacked(params.connectedContract),
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

    function handleCrossChainCall(
        CrossChainParams memory params,
        address zrc20,
        uint256 amount,
        bytes calldata message
    ) external returns (bool) {
        (
            address destination,
            address receiver,
            uint256 tokenId,
            string memory uri,
            address sender
        ) = abi.decode(message, (address, address, uint256, string, address));

        if (destination == address(0)) {
            if (params.isCrossChainTransfer) {
                // Transfer back to receiver
                return true; // Indicates transfer back
            } else {
                // Mint new token
                return false; // Indicates mint
            }
        } else {
            (address gasZrc20, uint256 gasFee) = IZRC20(destination)
                .withdrawGasFeeWithGasLimit(params.gasLimitAmount);
            if (destination != gasZrc20) revert InvalidAddress();

            uint256 out = SwapHelperLib.swapExactTokensForTokens(
                params.uniswapRouter,
                zrc20,
                amount,
                destination,
                0
            );

            if (!IZRC20(destination).approve(address(params.gateway), out)) {
                revert ApproveFailed();
            }
            params.gateway.withdrawAndCall(
                abi.encodePacked(params.connectedContract),
                out - gasFee,
                destination,
                abi.encode(receiver, tokenId, uri, out - gasFee, sender),
                CallOptions(params.gasLimitAmount, false),
                RevertOptions(
                    address(this),
                    true,
                    address(0),
                    abi.encode(tokenId, uri, sender),
                    0
                )
            );
        }
        return false;
    }

    event TokenTransfer(address receiver, address destination, uint256 tokenId, string uri);
}
