# 💽 Core

This document introduces the core concepts and mechanisms of the Fluxus platform.

## Bonding Curve Algorithm

The Bonding Curve is the core pricing mechanism of the Fluxus platform, which associates NFT supply with price through mathematical functions.

### How It Works

1. **Initial Price**: A baseline price is set when the NFT project is created
2. **Price Growth**: As the number of minted NFTs increases, the price gradually rises according to the curve function
3. **Price Cap**: After reaching the preset maximum price, the price stops rising
4. **Instant Pricing**: Each time a buy or sell occurs, the price is calculated in real-time by the algorithm

### Advantages

- **Automatic Pricing**: No need for manual listing, prices are automatically determined
- **Liquidity Guarantee**: Always has a tradable price
- **Price Discovery**: Market supply and demand automatically reflected in prices

## Trading Pool Mechanism

### How It Works

- All sold NFTs automatically enter the trading pool
- NFTs in the trading pool can be directly purchased by other users
- No need to wait for order matching, enabling instant trading

### Advantages

- **Instant Liquidity**: Can buy or sell at any time
- **Simplified Process**: No complex orderbook system required
- **Reduced Friction**: Minimizes transaction waiting time

## Royalty Mechanism

### Design Principles

- **Continuous Incentives**: Creators receive royalties from each transaction
- **Fair Distribution**: Royalty percentage is set when the project is created
- **Automatic Execution**: Automatically distributed through smart contracts

### Royalty Distribution

- Royalties are automatically calculated for each NFT transaction
- Royalties are directly distributed to project creators
- Supports multiple recipient addresses

## No Reserved Rights Mechanism

### Core Principles

- **Equal Participation**: Project creators and regular users participate in purchases on equal terms
- **Fair Issuance**: No reserved or priority purchase rights
- **Market-driven**: Prices are completely determined by the market

### Advantages

- **Enhanced Trust**: Eliminates the possibility of price manipulation by project creators
- **Fair Competition**: Equal opportunities for all participants
- **Market Transparency**: Completely transparent price formation process

## Cross-chain Mechanism

### ZetaChain Integration

- Supports NFT transfers between different chains
- Unified cross-chain interface
- Secure message passing mechanism

### Superchain-ERC721 Standard

- Unified NFT standard
- Cross-chain compatibility
- Future SuperChain ecosystem support

## Smart Contract Architecture

### Modular Design

- **Factory**: Project creation and management
- **Fluxus**: Core trading logic
- **Registry**: Project registration and queries
- **Trade**: Transaction execution
- **FeeVault**: Fee management

### Upgradeability

- Uses proxy pattern to support contract upgrades
- Maintains state unchanged
- Backward compatible

## Gas Optimization

### Optimization Strategies

- Batch operation support
- State variable optimization
- Event optimization
- Storage optimization

### Cost Control

- Minimize on-chain operations
- Efficient algorithm implementation
- Reasonable storage layout

## Security

### Security Measures

- Code audits
- Use of verified libraries
- Access control mechanisms
- Timelock protection

### Risk Management

- Price limiting mechanisms
- Trading volume limits
- Emergency pause functionality
