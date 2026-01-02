// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {ERC721Holder} from "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import { IPrice } from "../interfaces/core/IPrice.sol";
import { ITrade } from "../interfaces/periphery/ITrade.sol";
import { Fluxus } from "../Fluxus.sol";
import { EVMFluxusCrossChain } from "../extensions/crosschain/evm/EVMFluxusCrossChain.sol";

/**
 * @title Trade Contract
 * @author @lukema95
 * @notice Trade contract for Fluxus NFTs
 */
contract Trade is ITrade, ERC721Holder {

    /// @inheritdoc ITrade
    function mint(address _fluxusContractAddress) external payable {
        Fluxus fluxusContract = Fluxus(_fluxusContractAddress);
        IPrice priceContract = IPrice(fluxusContract.priceContract());
        uint256 price = priceContract.getBuyPriceAfterFee(_fluxusContractAddress);
        uint256 tokenId = fluxusContract.mint{value: msg.value}();
        fluxusContract.safeTransferFrom(address(this), msg.sender, tokenId);

        emit Minted(_fluxusContractAddress, msg.sender, tokenId, price);
    }

    /// @inheritdoc ITrade
    function buy(address _fluxusContractAddress, uint256 tokenId) external payable {
        Fluxus fluxusContract = Fluxus(_fluxusContractAddress);
        IPrice priceContract = IPrice(fluxusContract.priceContract());
        uint256 price = priceContract.getBuyPriceAfterFee(_fluxusContractAddress);
        fluxusContract.buy{value: msg.value}(tokenId);
        fluxusContract.safeTransferFrom(address(this), msg.sender, tokenId);

        emit Bought(_fluxusContractAddress, msg.sender, tokenId, price);
    }

    /// @inheritdoc ITrade
    function sell(address _fluxusContractAddress, uint256 tokenId) external {
        Fluxus fluxusContract = Fluxus(_fluxusContractAddress);
        IPrice priceContract = IPrice(fluxusContract.priceContract());
        uint256 price = priceContract.getSellPriceAfterFee(_fluxusContractAddress);
        fluxusContract.safeTransferFrom(msg.sender, address(this), tokenId);
        fluxusContract.sell(tokenId);

        (bool success, ) = msg.sender.call{value: price}("");
        if (!success) revert TransferFailed();

        emit Sold(_fluxusContractAddress, msg.sender, tokenId, price);
    }

    /// @inheritdoc ITrade
    function quickBuy(address _fluxusContractAddress) public payable {
        Fluxus fluxusContract = Fluxus(_fluxusContractAddress);
        IPrice priceContract = IPrice(fluxusContract.priceContract());
        if (fluxusContract.getAvailableTokensCount() == 0) revert NoTokensAvailable();
        uint256 tokenId = fluxusContract.getAvailableTokenByIndex(fluxusContract.getAvailableTokensCount() - 1);
        uint256 price = priceContract.getBuyPriceAfterFee(_fluxusContractAddress);
        fluxusContract.buy{value: price}(tokenId);
        fluxusContract.safeTransferFrom(address(this), msg.sender, tokenId);
        
        emit QuickBuyExecuted(_fluxusContractAddress, msg.sender, tokenId, price);
    }

    /// @inheritdoc ITrade
    function bulkBuy(address _fluxusContractAddress, uint256[] calldata tokenIds) external payable {
        Fluxus fluxusContract = Fluxus(_fluxusContractAddress);
        IPrice priceContract = IPrice(fluxusContract.priceContract());
        uint256 totalPrice = 0;
        for (uint256 i = 0; i < tokenIds.length; i++) {
            uint256 price = priceContract.getBuyPriceAfterFee(_fluxusContractAddress);
            totalPrice += price;
            if (msg.value < totalPrice) revert InsufficientPayment();
            fluxusContract.buy{value: price}(tokenIds[i]);
            fluxusContract.safeTransferFrom(address(this), msg.sender, tokenIds[i]);
        }
        
        emit BulkBuyExecuted(_fluxusContractAddress, msg.sender, tokenIds, totalPrice);

        // Refund excess ETH
        uint256 excess = msg.value - totalPrice;
        if (excess > 0) {
            (bool success, ) = msg.sender.call{value: excess}("");
            if (!success) revert RefundFailed();
        }
    }

    /// @inheritdoc ITrade
    function bulkQuickBuy(address _fluxusContractAddress, uint256 quantity) external payable {
        Fluxus fluxusContract = Fluxus(_fluxusContractAddress);
        IPrice priceContract = IPrice(fluxusContract.priceContract());
        uint256 totalPrice = 0;
        uint256 availableTokensCount = fluxusContract.getAvailableTokensCount();
        uint256[] memory tokenIds = new uint256[](quantity);
        if (availableTokensCount < quantity) revert NoTokensAvailable();
        for (uint256 i = 0; i < quantity; i++) {
            uint256 tokenId = fluxusContract.getAvailableTokenByIndex(availableTokensCount - 1);
            uint256 price = priceContract.getBuyPriceAfterFee(_fluxusContractAddress);
            totalPrice += price;
            if (msg.value < totalPrice) revert InsufficientPayment();
            fluxusContract.buy{value: price}(tokenId);
            fluxusContract.safeTransferFrom(address(this), msg.sender, tokenId);
            availableTokensCount--;
            tokenIds[i] = tokenId;
        }
        
        emit BulkQuickBuyExecuted(_fluxusContractAddress, msg.sender, tokenIds, totalPrice);

        // Refund excess ETH
        uint256 excess = msg.value - totalPrice;
        if (excess > 0) {
            (bool success, ) = msg.sender.call{value: excess}("");
            if (!success) revert RefundFailed();
        }
    }

    /// @inheritdoc ITrade
    function bulkSell(address _fluxusContractAddress, uint256[] calldata tokenIds) external {
        Fluxus fluxusContract = Fluxus(_fluxusContractAddress);
        IPrice priceContract = IPrice(fluxusContract.priceContract());
        uint256 totalPrice = 0;
        for (uint256 i = 0; i < tokenIds.length; i++) {
            uint256 tokenId = tokenIds[i];
            if (fluxusContract.ownerOf(tokenId) != msg.sender) revert NotTokenOwner();
            uint256 price = priceContract.getSellPriceAfterFee(_fluxusContractAddress);
            totalPrice += price;

            fluxusContract.safeTransferFrom(msg.sender, address(this), tokenId);
            fluxusContract.sell(tokenId);
        }
        
        emit BulkSellExecuted(_fluxusContractAddress, msg.sender, tokenIds, totalPrice);

        (bool success, ) = msg.sender.call{value: totalPrice}("");
        if (!success) revert TransferFailed();
    }

    /// @inheritdoc ITrade
    function bulkMint(address _fluxusContractAddress, uint256 quantity) external payable {
        Fluxus fluxusContract = Fluxus(_fluxusContractAddress);
        IPrice priceContract = IPrice(fluxusContract.priceContract());
        uint256 totalPrice = 0;
        uint256[] memory tokenIds = new uint256[](quantity);
        for (uint256 i = 0; i < quantity; i++) {
            uint256 price = priceContract.getBuyPriceAfterFee(_fluxusContractAddress);
            totalPrice += price;
            if (msg.value < totalPrice) revert InsufficientPayment();
            uint256 tokenId = fluxusContract.mint{value: price}();
            fluxusContract.safeTransferFrom(address(this), msg.sender, tokenId);
            tokenIds[i] = tokenId;
        }
        
        emit BulkMintExecuted(_fluxusContractAddress, msg.sender, tokenIds, totalPrice);

        // Refund excess ETH
        uint256 excess = msg.value - totalPrice;
        if (excess > 0) {
            (bool success, ) = msg.sender.call{value: excess}("");
            if (!success) revert RefundFailed();
        }
    }

    /// @inheritdoc ITrade
    function transferCrossChain(address payable _fluxusContractAddress, uint256 tokenId, address receiver, address destination) external payable {
        EVMFluxusCrossChain fluxusContract = EVMFluxusCrossChain(_fluxusContractAddress);
        
        // First transfer the NFT from the user to this contract (using the approval)
        fluxusContract.safeTransferFrom(msg.sender, address(this), tokenId);
        
        // Then call transferCrossChain from this contract (which now owns the NFT)
        fluxusContract.transferCrossChain{value: msg.value}(tokenId, receiver, destination);

        emit TransferCrossChain(_fluxusContractAddress, msg.sender, tokenId, receiver, destination);
    }

    receive() external payable {}
}