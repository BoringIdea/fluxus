// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test, console} from "forge-std/Test.sol";
import {Factory} from "../src/Factory.sol";
import {Registry} from "../src/Registry.sol";
import {Fluxus} from "../src/Fluxus.sol";
import {EVMFluxusCrossChain} from "../src/extensions/crosschain/evm/EVMFluxusCrossChain.sol";
import {FeeVault} from "../src/core/FeeVault.sol";
import {Price} from "../src/core/Price.sol";
import {IPrice} from "../src/interfaces/core/IPrice.sol";
import {IFactory} from "../src/interfaces/IFactory.sol";
import {Validator} from "../src/libs/Validator.sol";

contract FactoryTest is Test {
    Factory public factory;
    Registry public registry;
    FeeVault public feeVault;
    Price public priceContract;
    Fluxus public fluxusImplementation;
    EVMFluxusCrossChain public fluxusCrossChainImplementation;
    
    address public alice = address(0x1);
    address public bob = address(0x2);
    address public mockGateway = address(0x3);
    address public mockUniversal = address(0x4);
    uint256 public mockGasLimit = 500000;
    
    receive() external payable {}
        
    function setUp() public {
        registry = new Registry();
        feeVault = new FeeVault();
        priceContract = new Price();
        fluxusImplementation = new Fluxus();
        fluxusCrossChainImplementation = new EVMFluxusCrossChain();
        factory = new Factory(
            address(registry), 
            address(feeVault), 
            address(priceContract), 
            address(fluxusImplementation),
            address(fluxusCrossChainImplementation)
        );
        
        // Give test accounts some ETH
        vm.deal(alice, 10 ether);
        vm.deal(bob, 10 ether);
    }

    function test_createFluxus() public {
        address contractAddress = factory.createFluxus(
            "Fluxus", 
            "Fluxus", 
            0.001 ether, 
            10000, 
            0,
            0.05 ether, 
            "https://fluxus.com/image.png"
            );
        Fluxus fluxus = Fluxus(contractAddress);
        
        address expectedAddress = factory.calculateFluxusAddress(
            "Fluxus", 
            "Fluxus", 
            0.001 ether, 
            10000, 
            0,
            0.05 ether,
            "https://fluxus.com/image.png"
            );
        console.log("Deployed address:", address(fluxus));
        console.log("Expected address:", expectedAddress);
        assertEq(address(fluxus), expectedAddress);

        address[] memory creatorContracts = registry.getCreatorContracts(address(this));
        assertEq(creatorContracts.length, 1);
        assertEq(creatorContracts[0], expectedAddress);

        address creator = registry.getContractCreator(expectedAddress);
        assertEq(creator, address(this));

        uint256 buyPrice = IPrice(fluxus.priceContract()).getBuyPriceAfterFee(address(fluxus));
        vm.prank(alice);
        fluxus.mint{value: buyPrice}();
        assertEq(fluxus.balanceOf(alice), 1);
    }

    function test_createFluxusCrossChain() public {
        address payable contractAddress = payable(factory.createFluxusCrossChain(
            "FluxusTest1", 
            "FluxusTEST1", 
            1000000000000000, 
            10000, 
            1000000000000000000,
            50000000000000000, 
            "https://fluxuscrosschain.com/image.png",
            address(0x0c487a766110c85d301D96E33579C5B317Fa4995),
            12000000,
            true
        ));
        EVMFluxusCrossChain fluxusCrossChain = EVMFluxusCrossChain(contractAddress);
        
        address expectedAddress = factory.calculateFluxusCrossChainAddress(
            "FluxusTest1", 
            "FluxusTEST1", 
            1000000000000000, 
            10000, 
            1000000000000000000,
            50000000000000000, 
            "https://fluxuscrosschain.com/image.png",
            address(0x0c487a766110c85d301D96E33579C5B317Fa4995),
            12000000,
            true
        );
        console.log("CrossChain Deployed address:", address(fluxusCrossChain));
        console.log("CrossChain Expected address:", expectedAddress);
        assertEq(address(fluxusCrossChain), expectedAddress);

        address[] memory creatorContracts = registry.getCreatorContracts(address(this));
        assertEq(creatorContracts.length, 1);
        assertEq(creatorContracts[0], expectedAddress);

        address creator = registry.getContractCreator(expectedAddress);
        assertEq(creator, address(this));

        // Test basic properties
        // assertEq(fluxusCrossChain.name(), "FluxusCrossChain");
        // assertEq(fluxusCrossChain.symbol(), "FLICC");
        // assertEq(fluxusCrossChain.initialPrice(), 0.001 ether);
        // assertEq(fluxusCrossChain.maxSupply(), 10000);
        // assertEq(fluxusCrossChain.creatorFeePercent(), 0.05 ether);
        // assertEq(fluxusCrossChain.creator(), address(this));
        // assertEq(fluxusCrossChain.feeVault(), address(feeVault));
        // assertEq(fluxusCrossChain.priceContract(), address(priceContract));
        // assertEq(fluxusCrossChain.isSupportMint(), true);
        
        // Test cross-chain specific properties
        // assertEq(fluxusCrossChain.gateway(), mockGateway);
        // assertEq(fluxusCrossChain.gasLimit(), mockGasLimit);
    }

    function test_factoryState() public {
        assertEq(factory.registry(), address(registry));
        assertEq(factory.feeVault(), address(feeVault));
        assertEq(factory.priceContract(), address(priceContract));
        assertEq(factory.Fluxus_IMPLEMENTATION(), address(fluxusImplementation));
        assertEq(factory.Fluxus_CROSS_CHAIN_IMPLEMENTATION(), address(fluxusCrossChainImplementation));
    }

    function test_calculateFluxusAddress() public {
        address calculatedAddress = factory.calculateFluxusAddress(
            "Test", 
            "TEST", 
            0.001 ether, 
            1000, 
            0,
            0.1 ether,
            "https://test.com"
        );
        
        address actualAddress = factory.createFluxus(
            "Test", 
            "TEST", 
            0.001 ether, 
            1000, 
            0,
            0.1 ether, 
            "https://test.com"
        );
        
        assertEq(calculatedAddress, actualAddress);
    }

    function test_calculateFluxusCrossChainAddress() public {
        address calculatedAddress = factory.calculateFluxusCrossChainAddress(
            "TestCrossChain", 
            "TESTCC", 
            0.001 ether, 
            1000, 
            0,
            0.1 ether,
            "https://testcc.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        address actualAddress = factory.createFluxusCrossChain(
            "TestCrossChain", 
            "TESTCC", 
            0.001 ether, 
            1000, 
            0,
            0.1 ether, 
            "https://testcc.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        assertEq(calculatedAddress, actualAddress);
    }

    function test_createMultipleFluxuss() public {
        // Create first Fluxus
        address fluxus1 = factory.createFluxus(
            "Fluxus1", 
            "Fluxus1", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxus1.com"
        );
        
        // Create second Fluxus
        address fluxus2 = factory.createFluxus(
            "Fluxus2", 
            "Fluxus2", 
            0.002 ether, 
            2000, 
            0,
            0.1 ether, 
            "https://fluxus2.com"
        );
        
        // Verify both are registered
        address[] memory creatorContracts = registry.getCreatorContracts(address(this));
        assertEq(creatorContracts.length, 2);
        assertTrue(creatorContracts[0] == fluxus1 || creatorContracts[1] == fluxus1);
        assertTrue(creatorContracts[0] == fluxus2 || creatorContracts[1] == fluxus2);
        
        // Verify different addresses
        assertFalse(fluxus1 == fluxus2);
    }

    function test_createMultipleFluxusCrossChains() public {
        // Create first Fluxus CrossChain
        address fluxusCc1 = factory.createFluxusCrossChain(
            "FluxusCC1", 
            "FluxusCC1", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxuscc1.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        // Create second Fluxus CrossChain with different parameters
        address fluxusCc2 = factory.createFluxusCrossChain(
            "FluxusCC2", 
            "FluxusCC2", 
            0.002 ether, 
            2000, 
            0,
            0.1 ether, 
            "https://fluxuscc2.com",
            mockGateway,
            mockGasLimit + 100000, // Different gas limit
            true
        );
        
        // Verify both are registered
        address[] memory creatorContracts = registry.getCreatorContracts(address(this));
        assertEq(creatorContracts.length, 2);
        assertTrue(creatorContracts[0] == fluxusCc1 || creatorContracts[1] == fluxusCc1);
        assertTrue(creatorContracts[0] == fluxusCc2 || creatorContracts[1] == fluxusCc2);
        
        // Verify different addresses
        assertFalse(fluxusCc1 == fluxusCc2);
    }

    function test_createMixedFluxuss() public {
        // Create regular Fluxus
        address fluxus = factory.createFluxus(
            "RegularFluxus", 
            "RFluxus", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://regular.com"
        );
        
        // Create CrossChain Fluxus
        address fluxusCc = factory.createFluxusCrossChain(
            "CrossChainFluxus", 
            "CCFluxus", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://crosschain.com",
            mockGateway,    
            mockGasLimit,
            true
        );
        
        // Verify both are registered
        address[] memory creatorContracts = registry.getCreatorContracts(address(this));
        assertEq(creatorContracts.length, 2);
        assertTrue(creatorContracts[0] == fluxus || creatorContracts[1] == fluxus);
        assertTrue(creatorContracts[0] == fluxusCc || creatorContracts[1] == fluxusCc);
        
        // Verify different addresses
        assertFalse(fluxus == fluxusCc);
    }

    function test_createFluxusByDifferentCreators() public {
        // Create Fluxus by this contract
        address fluxus1 = factory.createFluxus(
            "Fluxus1", 
            "Fluxus1", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxus1.com"
        );
        
        // Create Fluxus by alice
        vm.prank(alice);
        address fluxus2 = factory.createFluxus(
            "Fluxus2", 
            "Fluxus2", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxus2.com"
        );
        
        // Verify registry entries
        address[] memory thisContracts = registry.getCreatorContracts(address(this));
        address[] memory aliceContracts = registry.getCreatorContracts(alice);
        
        assertEq(thisContracts.length, 1);
        assertEq(aliceContracts.length, 1);
        assertEq(thisContracts[0], fluxus1);
        assertEq(aliceContracts[0], fluxus2);
        
        assertEq(registry.getContractCreator(fluxus1), address(this));
        assertEq(registry.getContractCreator(fluxus2), alice);
    }

    function test_createFluxusCrossChainByDifferentCreators() public {
        // Create Fluxus CrossChain by this contract
        address fluxusCc1 = factory.createFluxusCrossChain(
            "FluxusCC1", 
            "FluxusCC1", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxuscc1.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        // Create Fluxus CrossChain by alice
        vm.prank(alice);
        address fluxusCc2 = factory.createFluxusCrossChain(
            "FluxusCC2", 
            "FluxusCC2", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxuscc2.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        // Verify registry entries
        address[] memory thisContracts = registry.getCreatorContracts(address(this));
        address[] memory aliceContracts = registry.getCreatorContracts(alice);
        
        assertEq(thisContracts.length, 1);
        assertEq(aliceContracts.length, 1);
        assertEq(thisContracts[0], fluxusCc1);
        assertEq(aliceContracts[0], fluxusCc2);
        
        assertEq(registry.getContractCreator(fluxusCc1), address(this));
        assertEq(registry.getContractCreator(fluxusCc2), alice);
    }

    function test_fluxusContractProperties() public {
        address fluxusAddress = factory.createFluxus(
            "TestFluxus", 
            "TFluxus", 
            0.001 ether, 
            5000, 
            0,
            0.03 ether, 
            "https://testfluxus.com"
        );
        
        Fluxus fluxus = Fluxus(fluxusAddress);
        
        // Test basic properties
        assertEq(fluxus.name(), "TestFluxus");
        assertEq(fluxus.symbol(), "TFluxus");
        assertEq(fluxus.initialPrice(), 0.001 ether);
        assertEq(fluxus.maxSupply(), 5000);
        assertEq(fluxus.creatorFeePercent(), 0.03 ether);
        assertEq(fluxus.creator(), address(this));
        assertEq(fluxus.feeVault(), address(feeVault));
        assertEq(fluxus.priceContract(), address(priceContract));
        assertEq(fluxus.totalSupply(), 0);
        assertEq(fluxus.currentSupply(), 0);
    }

    function test_fluxusCrossChainContractProperties() public {
        address payable fluxusCcAddress = payable(factory.createFluxusCrossChain(
            "TestFluxusCC", 
            "TFLICC", 
            0.001 ether, 
            5000, 
            0,
            0.03 ether, 
            "https://testfluxuscc.com",
            mockGateway,
            mockGasLimit,
            true
        ));
        
        EVMFluxusCrossChain fluxusCc = EVMFluxusCrossChain(fluxusCcAddress);
        
        // Test basic properties
        assertEq(fluxusCc.name(), "TestFluxusCC");
        assertEq(fluxusCc.symbol(), "TFLICC");
        assertEq(fluxusCc.initialPrice(), 0.001 ether);
        assertEq(fluxusCc.maxSupply(), 5000);
        assertEq(fluxusCc.creatorFeePercent(), 0.03 ether);
        assertEq(fluxusCc.creator(), address(this));
        assertEq(fluxusCc.feeVault(), address(feeVault));
        assertEq(fluxusCc.priceContract(), address(priceContract));
        assertEq(fluxusCc.totalSupply(), 0);
        assertEq(fluxusCc.currentSupply(), 0);
        
        // Test cross-chain specific properties
        assertEq(fluxusCc.gateway(), mockGateway);
        assertEq(fluxusCc.gasLimit(), mockGasLimit);
    }

    function test_priceCalculation() public {
        address fluxusAddress = factory.createFluxus(
            "PriceTest", 
            "PRICE", 
            0.001 ether, 
            1000, 
            0,
            0.1 ether, 
            "https://price.com"
        );
        
        Fluxus fluxus = Fluxus(fluxusAddress);
        
        // Test initial price
        uint256 initialBuyPrice = IPrice(fluxus.priceContract()).getBuyPriceAfterFee(address(fluxus));
        assertGt(initialBuyPrice, 0.001 ether); // Should be initial price + fee
        
        // Mint one NFT and check price increase
        vm.prank(alice);
        fluxus.mint{value: initialBuyPrice}();
        
        uint256 secondBuyPrice = IPrice(fluxus.priceContract()).getBuyPriceAfterFee(address(fluxus));
        assertGt(secondBuyPrice, initialBuyPrice); // Price should increase
        
        assertEq(fluxus.totalSupply(), 1);
        assertEq(fluxus.currentSupply(), 1);
    }

    function test_invalidParametersShouldRevert() public {
        // Test empty name
        vm.expectRevert(abi.encodeWithSelector(Validator.EmptyString.selector, "name"));
        factory.createFluxus(
            "", 
            "Fluxus", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxus.com"
        );
        
        // Test empty symbol
        vm.expectRevert(abi.encodeWithSelector(Validator.EmptyString.selector, "symbol"));
        factory.createFluxus(
            "Fluxus", 
            "", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxus.com"
        );
        
        // Test invalid initial price (too low)
        vm.expectRevert(abi.encodeWithSelector(Validator.InvalidPrice.selector, 0.00001 ether));
        factory.createFluxus(
            "Fluxus", 
            "Fluxus", 
            0.00001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxus.com"
        );
        
        // Test invalid initial price (too high)
        vm.expectRevert(abi.encodeWithSelector(Validator.InvalidPrice.selector, 2 ether));
        factory.createFluxus(
            "Fluxus", 
            "Fluxus", 
            2 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxus.com"
        );
        
        // Test invalid supply (too low)
        vm.expectRevert(abi.encodeWithSelector(Validator.InvalidSupply.selector, 50));
        factory.createFluxus(
            "Fluxus", 
            "Fluxus", 
            0.001 ether, 
            50, 
            0,
            0.05 ether, 
            "https://fluxus.com"
        );
        
        // Test invalid supply (too high)
        vm.expectRevert(abi.encodeWithSelector(Validator.InvalidSupply.selector, 2000000));
        factory.createFluxus(
            "Fluxus", 
            "Fluxus", 
            0.001 ether, 
            2000000, 
            0,
            0.05 ether, 
            "https://fluxus.com"
        );
        
        // Test invalid creator fee percent (too low)
        vm.expectRevert(abi.encodeWithSelector(Validator.InvalidFeePercent.selector, 0.005 ether));
        factory.createFluxus(
            "Fluxus", 
            "Fluxus", 
            0.001 ether, 
            1000, 
            0,
            0.005 ether, 
            "https://fluxus.com"
        );
        
        // Test invalid creator fee percent (too high)
        vm.expectRevert(abi.encodeWithSelector(Validator.InvalidFeePercent.selector, 1.1 ether));
        factory.createFluxus(
            "Fluxus", 
            "Fluxus", 
            0.001 ether, 
            1000, 
            0,
            1.1 ether, 
            "https://fluxus.com"
        );
    }

    function test_invalidCrossChainParametersShouldRevert() public {
        // Test same invalid parameters as regular Fluxus - should still apply to CrossChain
        vm.expectRevert(abi.encodeWithSelector(Validator.EmptyString.selector, "name"));
        factory.createFluxusCrossChain(
            "", 
            "FluxusCC", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxuscc.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        // Test invalid gateway address (zero address)
        vm.expectRevert(abi.encodeWithSelector(IFactory.InvalidGatewayAddress.selector));
        factory.createFluxusCrossChain(
            "FluxusCC", 
            "FluxusCC", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxuscc.com",
            address(0), 
            mockGasLimit,
            true
        );
        
        // Test invalid gas limit (zero)
        vm.expectRevert(abi.encodeWithSelector(IFactory.InvalidGasLimit.selector));
        factory.createFluxusCrossChain(
            "FluxusCC", 
            "FluxusCC", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://fluxuscc.com",
            mockGateway,
            0,
            true
        );
    }

    function test_eventEmission() public {
        // Calculate the expected address first
        address expectedAddress = factory.calculateFluxusAddress(
            "EventTest", 
            "EVENT", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://event.com"
        );
        
        // Test FluxusCreated event
        vm.expectEmit(true, true, false, false);
        emit IFactory.FluxusCreated(
            address(this),
            expectedAddress,
            address(priceContract),
            "EventTest",
            "EVENT",
            0.001 ether,
            1000,
            0,
            0.05 ether,
            "https://event.com"
        );
        
        factory.createFluxus(
            "EventTest", 
            "EVENT", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://event.com"
        );
    }

    function test_crossChainEventEmission() public {
        // Calculate the expected address first
        address expectedAddress = factory.calculateFluxusCrossChainAddress(
            "EventTestCC", 
            "EVENTCC", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://eventcc.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        // Test FluxusCrossChainCreated event
        vm.expectEmit(true, true, false, false);
        emit IFactory.FluxusCrossChainCreated(
            address(this),
            expectedAddress,
            address(priceContract),
            "EventTestCC",
            "EVENTCC",
            0.001 ether,
            1000,
            0,
            0.05 ether,
            "https://eventcc.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        factory.createFluxusCrossChain(
            "EventTestCC", 
            "EVENTCC", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether, 
            "https://eventcc.com",
            mockGateway,
            mockGasLimit,
            true
        );
    }

    function test_deterministic_addresses() public {
        // Same parameters should produce same calculated address
        address calc1 = factory.calculateFluxusAddress(
            "Same", 
            "SAME", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether,
            "https://same.com"
        );
        
        address calc2 = factory.calculateFluxusAddress(
            "Same", 
            "SAME", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether,
            "https://same.com"
        );
        
        assertEq(calc1, calc2);
        
        // Different parameters should produce different addresses
        address calc3 = factory.calculateFluxusAddress(
            "Different", 
            "DIFF", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether,
            "https://different.com"
        );
        
        assertFalse(calc1 == calc3);
    }

    function test_deterministic_crosschain_addresses() public {
        // Same parameters should produce same calculated address
        address calc1 = factory.calculateFluxusCrossChainAddress(
            "SameCC", 
            "SAMECC", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether,
            "https://samecc.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        address calc2 = factory.calculateFluxusCrossChainAddress(
            "SameCC", 
            "SAMECC", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether,
            "https://samecc.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        assertEq(calc1, calc2);
        
        // Different CrossChain parameters should produce different addresses
        address calc3 = factory.calculateFluxusCrossChainAddress(
            "DifferentCC", 
            "DIFFCC", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether,
            "https://differentcc.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        assertFalse(calc1 == calc3);
        
        // Same basic params but different crosschain params should produce different addresses
        address calc4 = factory.calculateFluxusCrossChainAddress(
            "SameCC", 
            "SAMECC", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether,
            "https://samecc.com",
            address(0x5), // Different gateway
            mockGasLimit,
            true
        );
        
        assertFalse(calc1 == calc4);
    }

    function test_crosschain_vs_regular_addresses_different() public {
        // Regular Fluxus and CrossChain Fluxus with same basic params should have different addresses
        address regularAddress = factory.calculateFluxusAddress(
            "Test", 
            "TEST", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether,
            "https://test.com"
        );
        
        address crossChainAddress = factory.calculateFluxusCrossChainAddress(
            "Test", 
            "TEST", 
            0.001 ether, 
            1000, 
            0,
            0.05 ether,
            "https://test.com",
            mockGateway,
            mockGasLimit,
            true
        );
        
        assertFalse(regularAddress == crossChainAddress);
    }
}
