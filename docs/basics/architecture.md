# 🏴‍☠️ Architecture

This document introduces the overall architecture design of the Fluxus platform.

## System Architecture Overview

Fluxus adopts a decentralized architecture, implementing NFT issuance, trading, and liquidity management based on smart contracts.

## Core Components

### 1. Smart Contract Layer

#### Factory Contract
- Responsible for creating and managing NFT projects
- Handles project initialization parameters
- Manages project registry

#### Fluxus Contract
- Implements the Bonding Curve algorithm
- Handles NFT minting and trading
- Manages trading pool and liquidity

#### Registry Contract
- Maintains project registration information
- Provides project query interfaces
- Manages project metadata

#### Trade Contract
- Handles NFT buy and sell transactions
- Calculates dynamic prices
- Manages trading pool

#### FeeVault Contract
- Manages platform fees
- Handles royalty distribution
- Fee withdrawal and distribution

### 2. Cross-chain Components

#### ZetaChain Integration
- Implements cross-chain NFT transfers
- Supports multi-chain interactions
- Cross-chain message passing

#### Superchain-ERC721 Standard
- Unified cross-chain NFT standard
- Enables seamless experience within the SuperChain ecosystem

### 3. Frontend Application Layer

#### Web Interface
- User interaction interface
- Wallet connection
- Trading operation interface

#### API Services
- Backend API services
- WebSocket gateway
- Real-time data push

### 4. Data Layer

#### The Graph Subgraph
- Indexes on-chain data
- Provides GraphQL query interfaces
- Real-time data synchronization

## Technology Stack

### Smart Contracts
- **Solidity**: Primary programming language
- **Hardhat/Foundry**: Development framework
- **OpenZeppelin**: Secure contract library

### Frontend
- **Next.js**: Web application framework
- **React**: UI library
- **Wagmi**: Ethereum React Hooks

### Backend
- **NestJS**: Node.js framework
- **WebSocket**: Real-time communication

### Data Indexing
- **The Graph**: Decentralized indexing protocol

## Data Flow

### NFT Minting Flow

1. User initiates minting request through frontend
2. Frontend calls smart contract
3. Bonding Curve algorithm calculates price
4. User confirms and pays
5. NFT minting completed
6. Events triggered, Subgraph indexes data
7. Frontend updates display

### NFT Trading Flow

1. User selects NFT to sell
2. Smart contract calculates selling price
3. NFT enters trading pool
4. Other users can purchase directly
5. Transaction completed, royalties automatically distributed
6. Data synchronized and updated

## Security Mechanisms

### Smart Contract Security
- Code audits
- Use of verified libraries
- Upgrade mechanisms

### Access Control
- Role-based permission management
- Multi-signature wallet support
- Timelock mechanisms

## Scalability

### Horizontal Scaling
- Support for multi-chain deployment
- Cross-chain interoperability
- Modular design

### Performance Optimization
- Gas optimization
- Batch operation support
- Caching mechanisms

## Future Plans

- Support for more EVM-compatible chains
- Enhanced cross-chain functionality
- Improved user experience
- New feature additions
