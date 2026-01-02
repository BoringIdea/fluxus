// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract UniversalNFTEvents {
    event SetUniversal(address indexed universalAddress, address indexed contractAddress);
    event SetConnected(address indexed zrc20, address indexed contractAddress);
    event TokenTransfer(
        address contractAddress,
        address indexed destination,
        address indexed receiver,
        uint256 indexed tokenId,
        string uri
    );
    event TokenTransferReceived(
        address indexed contractAddress,
        address indexed receiver,
        uint256 indexed tokenId,
        string uri
    );
    event TokenTransferReverted(
        address indexed contractAddress,
        address indexed sender,
        uint256 indexed tokenId,
        string uri
    );
    event TokenTransferToDestination(
        address contractAddress,
        address indexed destination,
        address indexed sender,
        uint256 indexed tokenId,
        string uri
    );
}
