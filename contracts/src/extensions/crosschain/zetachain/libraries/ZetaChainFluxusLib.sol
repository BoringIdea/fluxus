// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import { IPrice } from "../../../../interfaces/core/IPrice.sol";
import { IFeeVault } from "../../../../interfaces/core/IFeeVault.sol";

library ZetaChainFluxusLib {
    error TransferFailed();
    error RefundFailed();
    error InsufficientPayment();
    error NotSupportMint();
    error OverMaxSupply();
    error NotOwner();
    error NotOnSale();

    function distributeFee(address feeVault, address creator, uint256 fee) internal {
        uint256 protocolFee = (fee * IFeeVault(feeVault).protocolFeePercent()) / 1 ether;
        uint256 creatorFee = fee - protocolFee;

        (bool success1,) = creator.call{value: creatorFee}("");
        (bool success2,) = feeVault.call{value: protocolFee}("");
        if (!success1 || !success2) revert TransferFailed();
    }

    function refundExcess(address sender, uint256 price) internal {
        uint256 refundAmount = msg.value - price;
        if (refundAmount > 0) {
            (bool success,) = sender.call{value: refundAmount}("");
            if (!success) revert RefundFailed();
        }
    }

    function getPriceAndFee(address priceContract, address nftContract) internal view returns (uint256 price, uint256 creatorFee) {
        return IPrice(priceContract).getBuyPriceAndFee(nftContract);
    }

    function getSellPriceAndFee(address priceContract, address nftContract) internal view returns (uint256 price, uint256 creatorFee) {
        return IPrice(priceContract).getSellPriceAndFee(nftContract);
    }

    function validateMint(bool supportMint, uint256 totalSupply, uint256 maxSupply) internal pure {
        if (!supportMint) revert NotSupportMint();
        if (totalSupply >= maxSupply) revert OverMaxSupply();
    }

    function validateBuy(address owner, uint256 price, uint256 creatorFee, uint256 msgValue) internal pure {
        if (owner != address(0)) revert NotOnSale();
        if (msgValue < price + creatorFee) revert InsufficientPayment();
    }

    function validateSell(address owner, address sender) internal pure {
        if (owner != sender) revert NotOwner();
    }

    function validatePayment(uint256 price, uint256 creatorFee, uint256 msgValue) internal pure {
        if (msgValue < price + creatorFee) revert InsufficientPayment();
    }
}
