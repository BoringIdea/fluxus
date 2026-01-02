import {
  Bought as BoughtEvent,
  BulkBuyExecuted as BulkBuyExecutedEvent,
  BulkMintExecuted as BulkMintExecutedEvent,
  BulkQuickBuyExecuted as BulkQuickBuyExecutedEvent,
  BulkSellExecuted as BulkSellExecutedEvent,
  Minted as MintedEvent,
  QuickBuyExecuted as QuickBuyExecutedEvent,
  Sold as SoldEvent,
  TransferCrossChain as TransferCrossChainEvent,
} from "../generated/Trade/Trade"
import {
  Bought,
  BulkBuyExecuted,
  BulkMintExecuted,
  BulkQuickBuyExecuted,
  BulkSellExecuted,
  CollectionInfo,
  CollectionStats,
  Minted,
  NFTOwnership,
  OwnershipSummary,
  QuickBuyExecuted,
  Sold,
  TransferCrossChain,
  Txs,
  CrossChainStatus,
} from "../generated/schema"

import { BigInt, Bytes, log } from "@graphprotocol/graph-ts"

// Transaction type constants
const TX_TYPE_MINT = 1
const TX_TYPE_BUY = 2  
const TX_TYPE_SELL = 3
const TX_TYPE_BULK_BUY = 4
const TX_TYPE_BULK_SELL = 5
const TX_TYPE_BULK_MINT = 6
const TX_TYPE_QUICK_BUY = 7

// Helper function to create Txs record
function createTxsRecord(
  txHash: Bytes,
  collectionAddress: Bytes,
  txType: i32,
  sender: Bytes,
  price: BigInt,
  tokenIds: BigInt[],
  blockNumber: BigInt,
  blockTimestamp: BigInt
): void {
  let txsId = txHash
  let txs = new Txs(txsId)
  
  txs.collection_info = collectionAddress // reference to CollectionInfo
  txs.tx_type = txType
  txs.sender = sender
  txs.price = price
  txs.token_ids = tokenIds
  txs.blockNumber = blockNumber
  txs.blockTimestamp = blockTimestamp
  
  txs.save()
}

// Helper function to generate consistent CrossChainStatus ID
function generateCrossChainStatusId(fluxusContract: Bytes, tokenId: BigInt): Bytes {
  let idString = fluxusContract.toHexString().concat("-").concat(tokenId.toString())
  return Bytes.fromUTF8(idString)
}

// Helper function to generate consistent NFT ownership ID
function generateNFTOwnershipId(contractAddress: Bytes, tokenId: BigInt): Bytes {
  let idString = contractAddress.toHexString().concat("-").concat(tokenId.toString())
  return Bytes.fromUTF8(idString)
}

// Helper function to generate consistent ownership summary ID
function generateOwnershipSummaryId(contractAddress: Bytes, owner: string): Bytes {
  let idString = contractAddress.toHexString().concat("-").concat(owner)
  return Bytes.fromUTF8(idString)
}

// Helper function to get or create CollectionStats
function getOrCreateCollectionStats(contractAddress: Bytes, blockTimestamp: BigInt): CollectionStats {
  let stats = CollectionStats.load(contractAddress)
  if (!stats) {
    stats = new CollectionStats(contractAddress)
    stats.collection_info = contractAddress // reference to CollectionInfo
    stats.current_supply = 0
    stats.total_supply = 0
    stats.owners = 0
    stats.total_volume = BigInt.fromI32(0)
    stats.floor_price = BigInt.fromI32(0)
    stats.total_transactions = 0
    stats.last_updated = blockTimestamp
  }
  return stats
}

// Helper function to update ownership count and collection owners
function updateOwnershipCount(
  contractAddress: Bytes, 
  owner: string, 
  delta: i32,
  stats: CollectionStats,
  blockTimestamp: BigInt
): void {
  let summaryId = generateOwnershipSummaryId(contractAddress, owner)
  let summary = OwnershipSummary.load(summaryId)
  
  if (!summary) {
    summary = new OwnershipSummary(summaryId)
    summary.collection_info = contractAddress // reference to CollectionInfo
    summary.owner = owner
    summary.nft_count = 0
    summary.first_owned_at = blockTimestamp
    summary.last_updated = blockTimestamp
  }
  
  let oldCount = summary.nft_count
  let newCount = oldCount + delta
  
  // Ensure count doesn't go below 0
  if (newCount < 0) {
    newCount = 0
  }
  
  summary.nft_count = newCount
  summary.last_updated = blockTimestamp
  summary.save()
  
  // Update collection owners count based on ownership changes
  if (oldCount == 0 && newCount > 0) {
    // New owner
    stats.owners = stats.owners + 1
  } else if (oldCount > 0 && newCount == 0) {
    // Owner no longer has any NFTs
    stats.owners = stats.owners - 1
  }
  // If oldCount > 0 && newCount > 0, no change to owners count
}

export function handleBought(event: BoughtEvent): void {
  let entity = new Bought(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxusContract = event.params.fluxusContract
  entity.buyer = event.params.buyer
  entity.tokenId = event.params.tokenId
  entity.price = event.params.price

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create Txs record
  createTxsRecord(
    event.transaction.hash,
    event.params.fluxusContract,
    TX_TYPE_BUY,
    event.params.buyer,
    event.params.price,
    [event.params.tokenId],
    event.block.number,
    event.block.timestamp
  )

  // Check if CollectionInfo exists
  let collectionInfo = CollectionInfo.load(event.params.fluxusContract)
  if (!collectionInfo) {
    log.error("CollectionInfo not found: {}", [event.params.fluxusContract.toHex()])
    return
  }

  // Update collection stats
  let stats = getOrCreateCollectionStats(event.params.fluxusContract, event.block.timestamp)
  stats.total_volume = stats.total_volume.plus(event.params.price)
  stats.current_supply = stats.current_supply + 1
  stats.total_transactions = stats.total_transactions + 1
  stats.floor_price = event.params.price // Update floor price with latest transaction price
  stats.last_updated = event.block.timestamp

  let nftOwnership = NFTOwnership.load(generateNFTOwnershipId(event.params.fluxusContract, event.params.tokenId))
  if (!nftOwnership) {
    log.error("NFT ownership not found: {}-{}", [event.params.fluxusContract.toHexString(), event.params.tokenId.toString()])
    return
  }
  
  let previousOwner = nftOwnership.owner
  let newOwner = event.params.buyer.toHex()
  
  // Update ownership summary for previous owner (decrease count)
  updateOwnershipCount(event.params.fluxusContract, previousOwner, -1, stats, event.block.timestamp)
  
  // Update ownership summary for new owner (increase count)
  updateOwnershipCount(event.params.fluxusContract, newOwner, 1, stats, event.block.timestamp)
  
  nftOwnership.owner = newOwner
  nftOwnership.save()
  stats.save()
}

export function handleBulkBuyExecuted(event: BulkBuyExecutedEvent): void {
  let entity = new BulkBuyExecuted(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxusContract = event.params.fluxusContract
  entity.buyer = event.params.buyer
  entity.tokenIds = event.params.tokenIds
  entity.totalPrice = event.params.totalPrice

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create Txs record
  createTxsRecord(
    event.transaction.hash,
    event.params.fluxusContract,
    TX_TYPE_BULK_BUY,
    event.params.buyer,
    event.params.totalPrice,
    event.params.tokenIds,
    event.block.number,
    event.block.timestamp
  )

  // Check if CollectionInfo exists
  let collectionInfo = CollectionInfo.load(event.params.fluxusContract)
  if (!collectionInfo) {
    log.error("CollectionInfo not found: {}", [event.params.fluxusContract.toHex()])
    return
  }

  // Update collection stats
  let stats = getOrCreateCollectionStats(event.params.fluxusContract, event.block.timestamp)
  stats.total_volume = stats.total_volume.plus(event.params.totalPrice)
  stats.current_supply = stats.current_supply + event.params.tokenIds.length
  stats.total_transactions = stats.total_transactions + 1
  stats.floor_price = event.params.totalPrice.div(BigInt.fromI32(event.params.tokenIds.length)) // Average price per token
  stats.last_updated = event.block.timestamp

  let newOwner = event.params.buyer.toHex()
  let ownershipChanges = new Map<string, i32>()

  for (let i = 0; i < event.params.tokenIds.length; i++) {
    let nftOwnership = NFTOwnership.load(generateNFTOwnershipId(event.params.fluxusContract, event.params.tokenIds[i]))
    if (!nftOwnership) {
      log.error("NFT ownership not found: {}-{}", [event.params.fluxusContract.toHexString(), event.params.tokenIds[i].toString()])
      return
    }
    
    let previousOwner = nftOwnership.owner
    
    // Track ownership changes
    if (!ownershipChanges.has(previousOwner)) {
      ownershipChanges.set(previousOwner, 0)
    }
    ownershipChanges.set(previousOwner, ownershipChanges.get(previousOwner) - 1)
    
    if (!ownershipChanges.has(newOwner)) {
      ownershipChanges.set(newOwner, 0)
    }
    ownershipChanges.set(newOwner, ownershipChanges.get(newOwner) + 1)
    
    nftOwnership.owner = newOwner
    nftOwnership.save()
  }

  // Apply all ownership changes
  let owners = ownershipChanges.keys()
  for (let i = 0; i < owners.length; i++) {
    let owner = owners[i]
    let delta = ownershipChanges.get(owner)
    updateOwnershipCount(event.params.fluxusContract, owner, delta, stats, event.block.timestamp)
  }

  stats.save()
}

export function handleBulkMintExecuted(event: BulkMintExecutedEvent): void {
  let entity = new BulkMintExecuted(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxusContract = event.params.fluxusContract
  entity.buyer = event.params.buyer
  entity.tokenIds = event.params.tokenIds
  entity.totalPrice = event.params.totalPrice

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create Txs record
  createTxsRecord(
    event.transaction.hash,
    event.params.fluxusContract,
    TX_TYPE_BULK_MINT,
    event.params.buyer,
    event.params.totalPrice,
    event.params.tokenIds,
    event.block.number,
    event.block.timestamp
  )

  // Check if CollectionInfo exists
  let collectionInfo = CollectionInfo.load(event.params.fluxusContract)
  if (!collectionInfo) {
    log.error("CollectionInfo not found: {}", [event.params.fluxusContract.toHex()])
    return
  }

  // Update collection stats
  let stats = getOrCreateCollectionStats(event.params.fluxusContract, event.block.timestamp)
  stats.total_volume = stats.total_volume.plus(event.params.totalPrice)
  stats.current_supply = stats.current_supply + event.params.tokenIds.length
  stats.total_supply = stats.total_supply + event.params.tokenIds.length
  stats.total_transactions = stats.total_transactions + 1
  stats.last_updated = event.block.timestamp

  let newOwner = event.params.buyer.toHex()

  for (let i = 0; i < event.params.tokenIds.length; i++) {
    let ownershipId = generateNFTOwnershipId(event.params.fluxusContract, event.params.tokenIds[i])
    let nftOwnership = NFTOwnership.load(ownershipId)
    if (nftOwnership) {
      log.error("NFT ownership already exists: {}-{}", [event.params.fluxusContract.toHexString(), event.params.tokenIds[i].toString()])
      return
    }
    nftOwnership = new NFTOwnership(ownershipId)
    nftOwnership.collection_info = event.params.fluxusContract
    nftOwnership.token_id = event.params.tokenIds[i].toI32()
    nftOwnership.owner = newOwner
    nftOwnership.save()
  }

  // Update ownership count for the new owner (add all minted tokens)
  updateOwnershipCount(event.params.fluxusContract, newOwner, event.params.tokenIds.length, stats, event.block.timestamp)
  
  stats.save()
}

export function handleBulkQuickBuyExecuted(
  event: BulkQuickBuyExecutedEvent,
): void {
  let entity = new BulkQuickBuyExecuted(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxusContract = event.params.fluxusContract
  entity.buyer = event.params.buyer
  entity.tokenIds = event.params.tokenIds
  entity.totalPrice = event.params.totalPrice

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create Txs record (using QUICK_BUY type since it's a variation of bulk buy)
  createTxsRecord(
    event.transaction.hash,
    event.params.fluxusContract,
    TX_TYPE_QUICK_BUY,
    event.params.buyer,
    event.params.totalPrice,
    event.params.tokenIds,
    event.block.number,
    event.block.timestamp
  )

  // Check if CollectionInfo exists
  let collectionInfo = CollectionInfo.load(event.params.fluxusContract)
  if (!collectionInfo) {
    log.error("CollectionInfo not found: {}", [event.params.fluxusContract.toHex()])
    return
  }

  let stats = getOrCreateCollectionStats(event.params.fluxusContract, event.block.timestamp)
  stats.total_volume = stats.total_volume.plus(event.params.totalPrice)
  stats.current_supply = stats.current_supply + event.params.tokenIds.length
  stats.total_transactions = stats.total_transactions + 1
  stats.floor_price = event.params.totalPrice.div(BigInt.fromI32(event.params.tokenIds.length)) // Average price per token
  stats.last_updated = event.block.timestamp

  let newOwner = event.params.buyer.toHex()
  let ownershipChanges = new Map<string, i32>()

  for (let i = 0; i < event.params.tokenIds.length; i++) {
    let nftOwnership = NFTOwnership.load(generateNFTOwnershipId(event.params.fluxusContract, event.params.tokenIds[i]))
    if (!nftOwnership) {
      log.error("NFT ownership not found: {}-{}", [event.params.fluxusContract.toHexString(), event.params.tokenIds[i].toString()])
      return
    }
    
    let previousOwner = nftOwnership.owner
    
    // Track ownership changes
    if (!ownershipChanges.has(previousOwner)) {
      ownershipChanges.set(previousOwner, 0)
    }
    ownershipChanges.set(previousOwner, ownershipChanges.get(previousOwner) - 1)
    
    if (!ownershipChanges.has(newOwner)) {
      ownershipChanges.set(newOwner, 0)
    }
    ownershipChanges.set(newOwner, ownershipChanges.get(newOwner) + 1)
    
    nftOwnership.owner = newOwner
    nftOwnership.save()
  }

  // Apply all ownership changes
  let owners = ownershipChanges.keys()
  for (let i = 0; i < owners.length; i++) {
    let owner = owners[i]
    let delta = ownershipChanges.get(owner)
    updateOwnershipCount(event.params.fluxusContract, owner, delta, stats, event.block.timestamp)
  }

  stats.save()
}

export function handleBulkSellExecuted(event: BulkSellExecutedEvent): void {
  let entity = new BulkSellExecuted(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxusContract = event.params.fluxusContract
  entity.seller = event.params.seller
  entity.tokenIds = event.params.tokenIds
  entity.totalPrice = event.params.totalPrice

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create Txs record
  createTxsRecord(
    event.transaction.hash,
    event.params.fluxusContract,
    TX_TYPE_BULK_SELL,
    event.params.seller,
    event.params.totalPrice,
    event.params.tokenIds,
    event.block.number,
    event.block.timestamp
  )

  // Check if CollectionInfo exists
  let collectionInfo = CollectionInfo.load(event.params.fluxusContract)
  if (!collectionInfo) {
    log.error("CollectionInfo not found: {}", [event.params.fluxusContract.toHex()])
    return
  }

  // Update collection stats
  let stats = getOrCreateCollectionStats(event.params.fluxusContract, event.block.timestamp)
  stats.total_volume = stats.total_volume.plus(event.params.totalPrice)
  stats.current_supply = stats.current_supply - event.params.tokenIds.length
  stats.total_transactions = stats.total_transactions + 1
  stats.last_updated = event.block.timestamp

  let seller = event.params.seller.toHex()
  let contractAddress = event.params.fluxusContract.toHex()

  for (let i = 0; i < event.params.tokenIds.length; i++) {
    let nftOwnership = NFTOwnership.load(generateNFTOwnershipId(event.params.fluxusContract, event.params.tokenIds[i]))
    if (!nftOwnership) {
      log.error("NFT ownership not found: {}-{}", [event.params.fluxusContract.toHexString(), event.params.tokenIds[i].toString()])
      return
    }
    nftOwnership.owner = contractAddress
    nftOwnership.save()
  }

  // Update ownership count for seller (decrease by number of tokens sold)
  updateOwnershipCount(event.params.fluxusContract, seller, -event.params.tokenIds.length, stats, event.block.timestamp)
  
  // Update ownership count for contract (increase by number of tokens received)
  updateOwnershipCount(event.params.fluxusContract, contractAddress, event.params.tokenIds.length, stats, event.block.timestamp)

  stats.save()
}

export function handleMinted(event: MintedEvent): void {
  let entity = new Minted(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxusContract = event.params.fluxusContract
  entity.to = event.params.to
  entity.tokenId = event.params.tokenId
  entity.price = event.params.price

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create Txs record
  createTxsRecord(
    event.transaction.hash,
    event.params.fluxusContract,
    TX_TYPE_MINT,
    event.params.to,
    event.params.price,
    [event.params.tokenId],
    event.block.number,
    event.block.timestamp
  )

  // Check if CollectionInfo exists
  let collectionInfo = CollectionInfo.load(event.params.fluxusContract)
  if (!collectionInfo) {
    log.error("CollectionInfo not found: {}", [event.params.fluxusContract.toHex()])
    return
  }

  let stats = getOrCreateCollectionStats(event.params.fluxusContract, event.block.timestamp)
  stats.current_supply = stats.current_supply + 1
  stats.total_supply = stats.total_supply + 1
  stats.total_volume = stats.total_volume.plus(event.params.price)
  stats.total_transactions = stats.total_transactions + 1
  stats.floor_price = event.params.price // Update floor price with latest mint price
  stats.last_updated = event.block.timestamp

  let ownershipId = generateNFTOwnershipId(event.params.fluxusContract, event.params.tokenId)
  let nftOwnership = NFTOwnership.load(ownershipId)
  let newOwner = event.params.to.toHex()
  
  if (!nftOwnership) {
    // New NFT being minted
    nftOwnership = new NFTOwnership(ownershipId)
    nftOwnership.collection_info = event.params.fluxusContract
    nftOwnership.token_id = event.params.tokenId.toI32()
    nftOwnership.owner = newOwner
    nftOwnership.save()
    
    // Update ownership count for new owner (add 1)
    updateOwnershipCount(event.params.fluxusContract, newOwner, 1, stats, event.block.timestamp)
  } else {
    // Existing NFT ownership being updated (re-mint scenario)
    let previousOwner = nftOwnership.owner
    
    // Update ownership summary for previous owner (decrease count)
    updateOwnershipCount(event.params.fluxusContract, previousOwner, -1, stats, event.block.timestamp)
    
    // Update ownership summary for new owner (increase count)  
    updateOwnershipCount(event.params.fluxusContract, newOwner, 1, stats, event.block.timestamp)
    
    nftOwnership.owner = newOwner
    nftOwnership.save()
  }
  
  stats.save()
}

export function handleQuickBuyExecuted(event: QuickBuyExecutedEvent): void {
  let entity = new QuickBuyExecuted(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxusContract = event.params.fluxusContract
  entity.buyer = event.params.buyer
  entity.tokenId = event.params.tokenId
  entity.price = event.params.price

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create Txs record
  createTxsRecord(
    event.transaction.hash,
    event.params.fluxusContract,
    TX_TYPE_QUICK_BUY,
    event.params.buyer,
    event.params.price,
    [event.params.tokenId],
    event.block.number,
    event.block.timestamp
  )

  // Check if CollectionInfo exists
  let collectionInfo = CollectionInfo.load(event.params.fluxusContract)
  if (!collectionInfo) {
    log.error("CollectionInfo not found: {}", [event.params.fluxusContract.toHex()])
    return
  }

  let stats = getOrCreateCollectionStats(event.params.fluxusContract, event.block.timestamp)
  stats.total_volume = stats.total_volume.plus(event.params.price)
  stats.current_supply = stats.current_supply + 1
  stats.total_transactions = stats.total_transactions + 1
  stats.floor_price = event.params.price // Update floor price with latest transaction price
  stats.last_updated = event.block.timestamp

  let nftOwnership = NFTOwnership.load(generateNFTOwnershipId(event.params.fluxusContract, event.params.tokenId))
  if (!nftOwnership) {
    log.error("NFT ownership not found: {}-{}", [event.params.fluxusContract.toHexString(), event.params.tokenId.toString()])
    return
  }
  
  let previousOwner = nftOwnership.owner
  let newOwner = event.params.buyer.toHex()
  
  // Update ownership summary for previous owner (decrease count)
  updateOwnershipCount(event.params.fluxusContract, previousOwner, -1, stats, event.block.timestamp)
  
  // Update ownership summary for new owner (increase count)
  updateOwnershipCount(event.params.fluxusContract, newOwner, 1, stats, event.block.timestamp)
  
  nftOwnership.owner = newOwner
  nftOwnership.save()
  stats.save()
}

export function handleSold(event: SoldEvent): void {
  let entity = new Sold(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxusContract = event.params.fluxusContract
  entity.seller = event.params.seller
  entity.tokenId = event.params.tokenId
  entity.price = event.params.price

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create Txs record
  createTxsRecord(
    event.transaction.hash,
    event.params.fluxusContract,
    TX_TYPE_SELL,
    event.params.seller,
    event.params.price,
    [event.params.tokenId],
    event.block.number,
    event.block.timestamp
  )

  // Check if CollectionInfo exists
  let collectionInfo = CollectionInfo.load(event.params.fluxusContract)
  if (!collectionInfo) {
    log.error("CollectionInfo not found: {}", [event.params.fluxusContract.toHex()])
    return
  }

  let stats = getOrCreateCollectionStats(event.params.fluxusContract, event.block.timestamp)
  stats.total_volume = stats.total_volume.plus(event.params.price)
  stats.current_supply = stats.current_supply - 1
  stats.total_transactions = stats.total_transactions + 1
  stats.floor_price = event.params.price // Update floor price with latest sell price
  stats.last_updated = event.block.timestamp

  let nftOwnership = NFTOwnership.load(generateNFTOwnershipId(event.params.fluxusContract, event.params.tokenId))
  if (!nftOwnership) {
    log.error("NFT ownership not found: {}-{}", [event.params.fluxusContract.toHexString(), event.params.tokenId.toString()])
    return
  }
  
  let seller = event.params.seller.toHex()
  let contractAddress = event.params.fluxusContract.toHex()
  
  // Update ownership summary for seller (decrease count)
  updateOwnershipCount(event.params.fluxusContract, seller, -1, stats, event.block.timestamp)
  
  // Update ownership summary for contract (increase count)
  updateOwnershipCount(event.params.fluxusContract, contractAddress, 1, stats, event.block.timestamp)
  
  // When NFT is sold back to contract, set owner to contract address
  nftOwnership.owner = contractAddress
  nftOwnership.save()
  stats.save()
}

export function handleTransferCrossChain(event: TransferCrossChainEvent): void {
  let entity = new TransferCrossChain(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxusContract = event.params.fluxusContract
  entity.sender = event.params.sender
  entity.tokenId = event.params.tokenId
  entity.receiver = event.params.receiver
  entity.destination = event.params.destination

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create or update CrossChainStatus record
  let crossChainStatusId = generateCrossChainStatusId(event.params.fluxusContract, event.params.tokenId)
  let crossChainStatus = CrossChainStatus.load(crossChainStatusId)
  if (!crossChainStatus) {
    crossChainStatus = new CrossChainStatus(crossChainStatusId)
  }
  crossChainStatus.fluxusContract = event.params.fluxusContract
  crossChainStatus.sender = event.params.sender
  crossChainStatus.tokenId = event.params.tokenId
  crossChainStatus.receiver = event.params.receiver
  crossChainStatus.destination = event.params.destination
  crossChainStatus.isTransfered = true

  crossChainStatus.blockNumber = event.block.number
  crossChainStatus.blockTimestamp = event.block.timestamp
  crossChainStatus.transactionHash = event.transaction.hash

  crossChainStatus.save()

  // Update NFT ownership
  let nftOwnership = NFTOwnership.load(generateNFTOwnershipId(event.params.fluxusContract, event.params.tokenId))
  if (!nftOwnership) {
    log.error("NFT ownership not found: {}-{}", [event.params.fluxusContract.toHexString(), event.params.tokenId.toString()])
    return
  }

  nftOwnership.owner = event.params.fluxusContract.toHex()
  nftOwnership.save()

}
