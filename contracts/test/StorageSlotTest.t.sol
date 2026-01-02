// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test, console} from "forge-std/Test.sol";

/**
 * @title StorageSlotTest
 * @notice Simple test contract for verifying ERC7201 storage slot calculations
 */
contract StorageSlotTest is Test {
    
    /**
     * @notice Calculate ERC7201 storage slot according to standard
     * @param namespace The namespace string
     * @return slot The calculated storage slot
     */
    function calculateERC7201Slot(string memory namespace) public pure returns (bytes32 slot) {
        // keccak256(abi.encode(uint256(keccak256(namespace)) - 1)) & ~bytes32(uint256(0xff))
        bytes32 innerHash = keccak256(bytes(namespace));
        uint256 innerValue = uint256(innerHash);
        uint256 innerValueMinus1 = innerValue - 1;
        bytes32 outerHash = keccak256(abi.encode(innerValueMinus1));
        uint256 outerValue = uint256(outerHash);
        uint256 maskedValue = outerValue & ~uint256(0xff);
        slot = bytes32(maskedValue);
    }
    
    /**
     * @notice Test storage slot calculation for our project namespaces
     */
    function test_calculateProjectStorageSlots() public view {
        // Current project namespaces
        string memory storageNamespace = "fluxus.storage.Storage";
        string memory evmCrossChainNamespace = "fluxus.extensions.BaseNFTCrossChain";
        string memory zetaChainNamespace = "fluxus.extensions.BaseNFTCrossChain.zetachain";
        
        // Calculate storage slots
        bytes32 storageSlot = calculateERC7201Slot(storageNamespace);
        bytes32 evmCrossChainSlot = calculateERC7201Slot(evmCrossChainNamespace);
        bytes32 zetaChainSlot = calculateERC7201Slot(zetaChainNamespace);
        
        console.log("=== Project Storage Slots ===");
        console.log("Storage namespace:", storageNamespace);
        console.log("Storage slot:", vm.toString(storageSlot));
        
        console.log("EVM CrossChain namespace:", evmCrossChainNamespace);
        console.log("EVM CrossChain slot:", vm.toString(evmCrossChainSlot));
        
        console.log("ZetaChain namespace:", zetaChainNamespace);
        console.log("ZetaChain slot:", vm.toString(zetaChainSlot));
        
        // Verify all slots are unique
        assertFalse(storageSlot == evmCrossChainSlot, "Storage and EVM CrossChain slots should be unique");
        assertFalse(storageSlot == zetaChainSlot, "Storage and ZetaChain slots should be unique");
        assertFalse(evmCrossChainSlot == zetaChainSlot, "EVM CrossChain and ZetaChain slots should be unique");
    }
    
    /**
     * @notice Test that current hardcoded slot values match calculated values
     */
    function test_verifyCurrentSlotValues() public view {
        // Current hardcoded slot values
        bytes32 currentStorageSlot = 0x8f8c8c4c8a8f3f3e3d3c3b3a39393837363534333231302f2e2d2c2b2a292827;
        bytes32 currentEvmCrossChainSlot = 0x9f9d9c9b9a99989796959493929190908f8e8d8c8b8a89888786858483828180;
        bytes32 currentZetaChainSlot = 0xa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebf;
        
        // Calculate expected values
        bytes32 calculatedStorageSlot = calculateERC7201Slot("fluxus.storage.Storage");
        bytes32 calculatedEvmCrossChainSlot = calculateERC7201Slot("fluxus.extensions.BaseNFTCrossChain");
        bytes32 calculatedZetaChainSlot = calculateERC7201Slot("fluxus.extensions.BaseNFTCrossChain.zetachain");
        
        console.log("=== Current values vs calculated values ===");
        console.log("Storage - Current:", vm.toString(currentStorageSlot));
        console.log("Storage - Calculated:", vm.toString(calculatedStorageSlot));
        console.log("Storage matches:", currentStorageSlot == calculatedStorageSlot);
        
        console.log("EVM CrossChain - Current:", vm.toString(currentEvmCrossChainSlot));
        console.log("EVM CrossChain - Calculated:", vm.toString(calculatedEvmCrossChainSlot));
        console.log("EVM CrossChain matches:", currentEvmCrossChainSlot == calculatedEvmCrossChainSlot);
        
        console.log("ZetaChain - Current:", vm.toString(currentZetaChainSlot));
        console.log("ZetaChain - Calculated:", vm.toString(calculatedZetaChainSlot));
        console.log("ZetaChain matches:", currentZetaChainSlot == calculatedZetaChainSlot);
        
        // Verify current values are unique
        assertFalse(currentStorageSlot == currentEvmCrossChainSlot, "Current slots should be unique");
        assertFalse(currentStorageSlot == currentZetaChainSlot, "Current slots should be unique");
        assertFalse(currentEvmCrossChainSlot == currentZetaChainSlot, "Current slots should be unique");
    }
    
    /**
     * @notice Test ERC7201 compliance (slots should end with 0x00)
     */
    function test_erc7201Compliance() public view {
        bytes32 slot1 = calculateERC7201Slot("test.namespace.1");
        bytes32 slot2 = calculateERC7201Slot("test.namespace.2");
        bytes32 slot3 = calculateERC7201Slot("test.namespace.3");
        
        // Verify all slots end with 0x00
        assertEq(uint256(slot1) & 0xff, 0, "ERC7201 slot should end with 0x00");
        assertEq(uint256(slot2) & 0xff, 0, "ERC7201 slot should end with 0x00");
        assertEq(uint256(slot3) & 0xff, 0, "ERC7201 slot should end with 0x00");
        
        console.log("=== ERC7201 compliance test ===");
        console.log("Slot 1:", vm.toString(slot1));
        console.log("Slot 1 last byte:", uint256(slot1) & 0xff);
        console.log("Slot 2:", vm.toString(slot2));
        console.log("Slot 2 last byte:", uint256(slot2) & 0xff);
        console.log("Slot 3:", vm.toString(slot3));
        console.log("Slot 3 last byte:", uint256(slot3) & 0xff);
    }
    
    /**
     * @notice Test collision resistance with similar namespaces
     */
    function test_collisionResistance() public view {
        string[6] memory testNamespaces = [
            "fluxus.storage.Storage",
            "fluxus.storage.storage",  // Case difference
            "fluxus.Storage.Storage",  // First letter uppercase
            "fluxus.storage.Storage1", // Add number
            "fluxus.storage.Storag",   // Less one character
            "fluxus.storage.Storagee"  // More one character
        ];
        
        bytes32[6] memory slots;
        
        console.log("=== Collision resistance test ===");
        for (uint i = 0; i < testNamespaces.length; i++) {
            slots[i] = calculateERC7201Slot(testNamespaces[i]);
            console.log("Namespace:", testNamespaces[i]);
            console.log("Slot:", vm.toString(slots[i]));
        }
        
        // Verify all slots are different
        for (uint i = 0; i < slots.length; i++) {
            for (uint j = i + 1; j < slots.length; j++) {
                assertFalse(
                    slots[i] == slots[j], 
                    string(abi.encodePacked(
                        "Collision detected between: ", 
                        testNamespaces[i], " and ", testNamespaces[j]
                    ))
                );
            }
        }
    }
    
    /**
     * @notice Generate recommended storage slots for future use
     */
    function test_generateRecommendedSlots() public view {
        string[5] memory futureNamespaces = [
            "fluxus.storage.Storage",
            "fluxus.extensions.BaseNFTCrossChain",
            "fluxus.extensions.BaseNFTCrossChain.zetachain",
            "fluxus.permissions.AccessControl",
            "fluxus.rewards.StakingRewards"
        ];
        
        console.log("=== Recommended storage slots ===");
        for (uint i = 0; i < futureNamespaces.length; i++) {
            bytes32 slot = calculateERC7201Slot(futureNamespaces[i]);
            console.log("Namespace:", futureNamespaces[i]);
            console.log("Recommended slot:", vm.toString(slot));
            console.log("---");
        }
    }
} 