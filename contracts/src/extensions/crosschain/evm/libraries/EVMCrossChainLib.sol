// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {GatewayEVM, RevertOptions} from "@zetachain/protocol-contracts/contracts/evm/GatewayEVM.sol";

library EVMCrossChainLib {
    error InvalidAddress();
    error Unauthorized();
    error GasTokenTransferFailed();
    error CallFailed();

    struct CrossChainParams {
        GatewayEVM gateway;
        address universal;
        uint256 gasLimitAmount;
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
        if (receiver == address(0)) revert InvalidAddress();

        bytes memory message = abi.encode(
            destination,
            receiver,
            tokenId,
            uri,
            sender
        );

        emit TokenTransfer(destination, receiver, tokenId, uri);

        if (destination == address(0)) {
            try params.gateway.call(
                params.universal,
                message,
                RevertOptions(address(this), false, address(0), message, 0)
            ) {
                // Call succeeded
            } catch {
                revert CallFailed();
            }
        } else {
            params.gateway.depositAndCall{value: msg.value}(
                params.universal,
                message,
                RevertOptions(
                    address(this),
                    true,
                    address(0),
                    abi.encode(receiver, tokenId, uri, sender),
                    params.gasLimitAmount
                )
            );
        }
    }

    function handleCrossChainCall(
        CrossChainParams memory params,
        address contextSender,
        bytes calldata message
    ) external returns (bool, address, uint256, string memory, uint256, address) {
        if (contextSender != params.universal) revert Unauthorized();

        (
            address receiver,
            uint256 tokenId,
            string memory uri,
            uint256 gasAmount,
            address sender
        ) = abi.decode(message, (address, uint256, string, uint256, address));

        if (params.isCrossChainTransfer) {
            // Transfer back to receiver
            return (true, receiver, tokenId, uri, gasAmount, sender);
        } else {
            // Mint new token
            return (false, receiver, tokenId, uri, gasAmount, sender);
        }
    }

    function handleGasTransfer(uint256 gasAmount, address sender) external {
        if (gasAmount > 0) {
            if (sender == address(0)) revert InvalidAddress();
            (bool success, ) = payable(sender).call{value: gasAmount}("");
            if (!success) revert GasTokenTransferFailed();
        }
    }

    event TokenTransfer(address destination, address receiver, uint256 tokenId, string uri);
}
