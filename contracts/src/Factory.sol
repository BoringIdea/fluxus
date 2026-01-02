// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**********************************************************************
 *       ███████╗██╗     ██╗   ██╗██╗  ██╗██╗   ██╗ ██████╗            *
 *       ██╔════╝██║     ██║   ██║╚██╗██╔╝██║   ██║██╔════╝            *
 *       █████╗  ██║     ██║   ██║ ╚███╔╝ ██║   ██║╚█████╗             *
 *       ██╔══╝  ██║     ██║   ██║ ██╔██╗ ██║   ██║ ╚═══██╗            *
 *       ██║     ███████╗╚██████╔╝██╔╝ ██╗╚██████╔╝██████╔╝            *
 *       ╚═╝     ╚══════╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═════╝             *
 *                        ⭐️ FLUXUS.ART ⭐️                              *
 **********************************************************************/

import {Fluxus} from "./Fluxus.sol";
import {EVMFluxusCrossChain} from "./extensions/crosschain/evm/EVMFluxusCrossChain.sol";
import {IRegistry} from "./interfaces/IRegistry.sol";
import {IFactory} from "./interfaces/IFactory.sol";
import {Validator} from "./libs/Validator.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/**
 * @title Factory Contract
 * @author @lukema95
 * @notice Factory contract for creating Fluxus contracts
 */
contract Factory is IFactory {
    address public override registry;
    address public override feeVault;
    address public override priceContract;
    address public immutable Fluxus_IMPLEMENTATION;
    address public immutable Fluxus_CROSS_CHAIN_IMPLEMENTATION;


    constructor(
        address _registry, 
        address _feeVault, 
        address _priceContract, 
        address _fluxusImplementation,
        address _fluxusCrossChainImplementation
    ) {
        if (_registry == address(0) || 
            _feeVault == address(0) || 
            _priceContract == address(0) || 
            _fluxusImplementation == address(0) || 
            _fluxusCrossChainImplementation == address(0)) 
        {
            revert ZeroAddress();
        }

        registry = _registry;
        feeVault = _feeVault;
        priceContract = _priceContract;
        Fluxus_IMPLEMENTATION = _fluxusImplementation;
        Fluxus_CROSS_CHAIN_IMPLEMENTATION = _fluxusCrossChainImplementation;
    }

    /// @dev Modifier to validate the parameters for factory to create a Fluxus contract
    modifier validateParams(
        string memory name,
        string memory symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent
    ) {
        Validator.validateFluxusParams(
            name,
            symbol,
            initialPrice,
            maxSupply,
            maxPrice,
            creatorFeePercent
        );
        _;
    }

    /// @dev Modifier to validate cross-chain specific parameters
    modifier validateCrossChainParams(address gatewayAddress, uint256 gasLimit) {
        _validateCrossChainParams(gatewayAddress, gasLimit);
        _;
    }

    function _validateCrossChainParams(address gatewayAddress, uint256 gasLimit) internal pure {
        if (gatewayAddress == address(0)) {
            revert InvalidGatewayAddress();
        }
        if (gasLimit == 0) {
            revert InvalidGasLimit();
        }
    }

    /// @dev Modifier to validate the creator of the Fluxus contract
    modifier onlyCreator(address fluxus) {
        _onlyCreator(fluxus);
        _;
    }

    function _onlyCreator(address fluxus) internal view {
        if (msg.sender != Fluxus(fluxus).creator()) {
            revert Unauthorized();
        }
    }

    /// @inheritdoc IFactory
    function createFluxus(
        string memory name,
        string memory symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent,
        string memory baseUri
    )
        external
        validateParams(name, symbol, initialPrice, maxSupply, maxPrice, creatorFeePercent)
        returns (address)
    {
        address creator = msg.sender;
        bytes32 salt = getSalt(name, symbol, initialPrice, maxSupply, maxPrice);
        bytes memory initData = abi.encodeWithSelector(
            Fluxus.initialize.selector,
            name,
            symbol,
            initialPrice,
            maxSupply,
            maxPrice,
            creatorFeePercent,
            feeVault,
            priceContract,
            creator,
            baseUri
        );

        ERC1967Proxy proxy = new ERC1967Proxy{
            salt: salt
        }(Fluxus_IMPLEMENTATION, initData);
        
        address fluxusAddress = address(proxy);
        
        emit FluxusCreated(
            creator,
            fluxusAddress,
            priceContract,
            name,
            symbol,
            initialPrice,
            maxSupply,
            maxPrice,
            creatorFeePercent,
            baseUri
        );

        IRegistry(registry).register(creator, fluxusAddress);
        return fluxusAddress;
    }

    /// @inheritdoc IFactory
    function createFluxusCrossChain(
        string memory name,
        string memory symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent,
        string memory baseUri,
        address gatewayAddress,
        uint256 gasLimit,
        bool supportMint
    )
        external
        validateParams(name, symbol, initialPrice, maxSupply, maxPrice, creatorFeePercent)
        validateCrossChainParams(gatewayAddress, gasLimit)
        returns (address)
    {
        address creator = msg.sender;
        bytes32 salt = getCrossChainSalt(
            name, 
            symbol, 
            initialPrice, 
            maxSupply, 
            maxPrice, 
            creatorFeePercent,
            baseUri,
            gatewayAddress,
            gasLimit,
            supportMint
            );

        bytes memory initData = abi.encodeWithSelector(
            EVMFluxusCrossChain.initialize.selector,
            name,
            symbol,
            initialPrice,
            maxSupply,
            maxPrice,
            creatorFeePercent,
            feeVault,
            priceContract,
            creator,
            baseUri,
            gatewayAddress,
            gasLimit,
            supportMint
        );

        ERC1967Proxy proxy = new ERC1967Proxy{
            salt: salt
        }(Fluxus_CROSS_CHAIN_IMPLEMENTATION, initData);
        
        address fluxusAddress = address(proxy);
        
        emit FluxusCrossChainCreated(
            creator,
            fluxusAddress,
            priceContract,
            name,
            symbol,
            initialPrice,
            maxSupply,
            maxPrice,
            creatorFeePercent,
            baseUri,
            gatewayAddress,
            gasLimit,
            supportMint
        );

        IRegistry(registry).register(creator, fluxusAddress);
        return fluxusAddress;
    }

    /// @inheritdoc IFactory
    function calculateFluxusAddress(
        string memory name,
        string memory symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent,
        string memory baseUri
    ) public view returns (address fluxusAddress) {
        address creator = msg.sender;
        
        bytes memory initData = abi.encodeWithSelector(
            Fluxus.initialize.selector,
            name,
            symbol,
            initialPrice,
            maxSupply,
            maxPrice,
            creatorFeePercent,
            feeVault,
            priceContract,
            creator,
            baseUri
        );

        bytes memory proxyInitParams = abi.encode(Fluxus_IMPLEMENTATION, initData);
        bytes32 salt = getSalt(name, symbol, initialPrice, maxSupply, maxPrice);
        fluxusAddress = address(
            uint160(
                uint256(
                    keccak256(
                        abi.encodePacked(
                            bytes1(0xff),
                            address(this),
                            salt,
                            keccak256(
                                abi.encodePacked(
                                    type(ERC1967Proxy).creationCode,
                                    proxyInitParams
                                )
                            )
                        )
                    )
                )
            )
        );
    }

    /// @inheritdoc IFactory
    function setUniversal(address payable fluxus, address universal) external onlyCreator(fluxus) {
        EVMFluxusCrossChain(fluxus).setUniversal(universal);
        emit SetUniversal(fluxus, universal);
    }

    /// @inheritdoc IFactory
    function setGateway(address payable fluxus, address gateway) external onlyCreator(fluxus) {
        EVMFluxusCrossChain(fluxus).setGateway(gateway);
        emit SetGateway(fluxus, gateway);
    }

    /// @inheritdoc IFactory
    function setGasLimit(address payable fluxus, uint256 gasLimit) external onlyCreator(fluxus) {
        EVMFluxusCrossChain(fluxus).setGasLimit(gasLimit);
        emit SetGasLimit(fluxus, gasLimit);
    }

    /// @inheritdoc IFactory
    function calculateFluxusCrossChainAddress(
        string memory name,
        string memory symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent,
        string memory baseUri,
        address gatewayAddress,
        uint256 gasLimit,
        bool supportMint
    ) public view returns (address fluxusAddress) {
        address creator = msg.sender;
        
        bytes memory initData = abi.encodeWithSelector(
            EVMFluxusCrossChain.initialize.selector,
            name,
            symbol,
            initialPrice,
            maxSupply,
            maxPrice,
            creatorFeePercent,
            feeVault,
            priceContract,
            creator,
            baseUri,
            gatewayAddress,
            gasLimit,
            supportMint
        );

        bytes memory proxyInitParams = abi.encode(Fluxus_CROSS_CHAIN_IMPLEMENTATION, initData);
        bytes32 salt = getCrossChainSalt(
            name, 
            symbol, 
            initialPrice, 
            maxSupply, 
            maxPrice, 
            creatorFeePercent,
            baseUri,
            gatewayAddress,
            gasLimit,
            supportMint);

        fluxusAddress = address(
            uint160(
                uint256(
                    keccak256(
                        abi.encodePacked(
                            bytes1(0xff),
                            address(this),
                            salt,
                            keccak256(
                                abi.encodePacked(
                                    type(ERC1967Proxy).creationCode,
                                    proxyInitParams
                                )
                            )
                        )
                    )
                )
            )
        );
    }

    function getSalt(
        string memory name,
        string memory symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice
    ) internal pure returns (bytes32) {
        return keccak256(abi.encode(name, symbol, initialPrice, maxSupply, maxPrice));
    }

    function getCrossChainSalt(
        string memory name,
        string memory symbol,
        uint256 initialPrice,
        uint256 maxSupply,
        uint256 maxPrice,
        uint256 creatorFeePercent,
        string memory baseUri,
        address gatewayAddress,
        uint256 gasLimit,
        bool supportMint
    ) internal pure returns (bytes32) {
        return keccak256(
            abi.encode(
                name, 
                symbol, 
                initialPrice, 
                maxSupply, 
                maxPrice, 
                creatorFeePercent,
                baseUri,
                gatewayAddress,
                gasLimit,
                supportMint
            )
        );
    }
}

