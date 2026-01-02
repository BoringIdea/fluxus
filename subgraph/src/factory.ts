import {
  FluxusCreated as FluxusCreatedEvent,
  FluxusCrossChainCreated as FluxusCrossChainCreatedEvent,
  SetGasLimit as SetGasLimitEvent,
  SetGateway as SetGatewayEvent,
  SetUniversal as SetUniversalEvent,
} from "../generated/Factory/Factory"
import {
  FluxusCreated,
  FluxusCrossChainCreated,
  SetGasLimit,
  SetGateway,
  SetUniversal,
  CollectionStats,
  CollectionInfo,
} from "../generated/schema"
import { BigInt, Bytes } from "@graphprotocol/graph-ts"

function initializeCollectionStats(
  collectionAddress: Bytes,
  initialPrice: BigInt,
  blockTimestamp: BigInt,
): void {
  let stats = CollectionStats.load(collectionAddress)
  if (!stats) {
    stats = new CollectionStats(collectionAddress)
    stats.collection_info = collectionAddress
    stats.current_supply = 0
    stats.total_supply = 0
    stats.owners = 0
    stats.total_volume = BigInt.fromI32(0)
    stats.floor_price = initialPrice
    stats.total_transactions = 0
    stats.last_updated = blockTimestamp
    stats.save()
  }
}

export function handleFluxusCreated(event: FluxusCreatedEvent): void {
  let entity = new FluxusCreated(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.creator = event.params.creator
  entity.fluxusAddress = event.params.fluxusAddress
  entity.priceAddress = event.params.priceAddress
  entity.name = event.params.name
  entity.symbol = event.params.symbol
  entity.initialPrice = event.params.initialPrice
  entity.maxSupply = event.params.maxSupply
  entity.maxPrice = event.params.maxPrice
  entity.creatorFeePercent = event.params.creatorFeePercent
  entity.baseURI = event.params.baseUri

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create or update CollectionInfo
  let collectionId = event.params.fluxusAddress
  let collection = CollectionInfo.load(collectionId)
  if (!collection) {
    collection = new CollectionInfo(collectionId)
  }
  collection.address = event.params.fluxusAddress.toHex()
  collection.name = event.params.name
  collection.symbol = event.params.symbol
  collection.creator = event.params.creator.toHex()
  collection.creator_fee = event.params.creatorFeePercent
  collection.base_uri = event.params.baseUri

  collection.initial_price = event.params.initialPrice
  collection.max_supply = event.params.maxSupply.toI32()
  collection.max_price = event.params.maxPrice
  collection.support_crosschain = false
  collection.support_mint = true
  collection.gas_limit = BigInt.fromI32(0)
  collection.price_contract = event.params.priceAddress.toHex()
  collection.is_registered = true
  collection.created_at = event.block.timestamp
  collection.block_number = event.block.number
  collection.save()

  // Initialize stats for this collection (idempotent)
  initializeCollectionStats(
    collectionId,
    event.params.initialPrice,
    event.block.timestamp,
  )
}

export function handleFluxusCrossChainCreated(
  event: FluxusCrossChainCreatedEvent,
): void {
  let entity = new FluxusCrossChainCreated(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.creator = event.params.creator
  entity.fluxusAddress = event.params.fluxusAddress
  entity.priceAddress = event.params.priceAddress
  entity.name = event.params.name
  entity.symbol = event.params.symbol
  entity.initialPrice = event.params.initialPrice
  entity.maxSupply = event.params.maxSupply
  entity.maxPrice = event.params.maxPrice
  entity.creatorFeePercent = event.params.creatorFeePercent
  entity.baseURI = event.params.baseUri
  entity.gatewayAddress = event.params.gatewayAddress
  entity.gasLimit = event.params.gasLimit
  entity.supportMint = event.params.supportMint

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Create or update CollectionInfo
  let collectionId = event.params.fluxusAddress
  let collection = CollectionInfo.load(collectionId)
  if (!collection) {
    collection = new CollectionInfo(collectionId)
  }
  collection.address = event.params.fluxusAddress.toHex()
  collection.name = event.params.name
  collection.symbol = event.params.symbol
  collection.creator = event.params.creator.toHex()
  collection.creator_fee = event.params.creatorFeePercent
  collection.base_uri = event.params.baseUri
  // TODO: add gateway address
  collection.initial_price = event.params.initialPrice
  collection.max_supply = event.params.maxSupply.toI32()
  collection.max_price = event.params.maxPrice
  collection.support_crosschain = true
  collection.support_mint = event.params.supportMint
  collection.gas_limit = event.params.gasLimit
  collection.price_contract = event.params.priceAddress.toHex()
  collection.is_registered = true
  collection.created_at = event.block.timestamp
  collection.block_number = event.block.number
  collection.save()

  // Initialize stats for this collection (idempotent)
  initializeCollectionStats(
    collectionId,
    event.params.initialPrice,
    event.block.timestamp,
  )
}

export function handleSetGasLimit(event: SetGasLimitEvent): void {
  let entity = new SetGasLimit(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxus = event.params.fluxus
  entity.gasLimit = event.params.gasLimit

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Update CollectionInfo
  let collectionId = event.params.fluxus
  let collection = CollectionInfo.load(collectionId)
  if (collection) {
    collection.gas_limit = event.params.gasLimit
    collection.save()
  }
}

export function handleSetGateway(event: SetGatewayEvent): void {
  let entity = new SetGateway(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxus = event.params.fluxus
  entity.gateway = event.params.gateway

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Update CollectionInfo
  let collectionId = event.params.fluxus
  let collection = CollectionInfo.load(collectionId)
  if (collection) {
    collection.gateway = event.params.gateway.toHex()
    collection.save()
  }
}

export function handleSetUniversal(event: SetUniversalEvent): void {
  let entity = new SetUniversal(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.fluxus = event.params.fluxus
  entity.universal = event.params.universal

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()

  // Update CollectionInfo
  let collectionId = event.params.fluxus
  let collection = CollectionInfo.load(collectionId)
  if (collection) {
    collection.universal = event.params.universal.toHex()
    collection.save()
  }
}
