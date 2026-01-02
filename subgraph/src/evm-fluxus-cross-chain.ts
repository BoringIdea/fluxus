import {
  Approval as ApprovalEvent,
  ApprovalForAll as ApprovalForAllEvent,
  BatchMetadataUpdate as BatchMetadataUpdateEvent,
  Initialized as InitializedEvent,
  MetadataUpdate as MetadataUpdateEvent,
  OwnershipTransferred as OwnershipTransferredEvent,
  SetConnected as SetConnectedEvent,
  SetUniversal as SetUniversalEvent,
  TokenMinted as TokenMintedEvent,
  TokenTransfer as TokenTransferEvent,
  TokenTransferReceived as TokenTransferReceivedEvent,
  TokenTransferReverted as TokenTransferRevertedEvent,
  TokenTransferToDestination as TokenTransferToDestinationEvent,
  Transfer as TransferEvent,
  Upgraded as UpgradedEvent,
} from "../generated/EVMFluxusCrossChain/EVMFluxusCrossChain"
import {
  EVMFluxusCrossChainApproval as Approval,
  EVMFluxusCrossChainApprovalForAll as ApprovalForAll,
  BatchMetadataUpdate,
  EVMFluxusCrossChainInitialized as Initialized,
  MetadataUpdate,
  EVMFluxusCrossChainOwnershipTransferred as OwnershipTransferred,
  SetConnected,
  SetUniversal,
  EVMFluxusCrossChainTokenMinted as TokenMinted,
  TokenTransfer,
  TokenTransferReceived,
  TokenTransferReverted,
  TokenTransferToDestination,
  EVMFluxusCrossChainTransfer as Transfer,
  EVMFluxusCrossChainUpgraded as Upgraded,
} from "../generated/schema"

export function handleApproval(event: ApprovalEvent): void {
  let entity = new Approval(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.owner = event.params.owner
  entity.approved = event.params.approved
  entity.tokenId = event.params.tokenId

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleApprovalForAll(event: ApprovalForAllEvent): void {
  let entity = new ApprovalForAll(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.owner = event.params.owner
  entity.operator = event.params.operator
  entity.approved = event.params.approved

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleBatchMetadataUpdate(
  event: BatchMetadataUpdateEvent,
): void {
  let entity = new BatchMetadataUpdate(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity._fromTokenId = event.params._fromTokenId
  entity._toTokenId = event.params._toTokenId

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleInitialized(event: InitializedEvent): void {
  let entity = new Initialized(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.version = event.params.version

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleMetadataUpdate(event: MetadataUpdateEvent): void {
  let entity = new MetadataUpdate(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity._tokenId = event.params._tokenId

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleOwnershipTransferred(
  event: OwnershipTransferredEvent,
): void {
  let entity = new OwnershipTransferred(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.previousOwner = event.params.previousOwner
  entity.newOwner = event.params.newOwner

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleSetConnected(event: SetConnectedEvent): void {
  let entity = new SetConnected(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.zrc20 = event.params.zrc20
  entity.contractAddress = event.params.contractAddress

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleSetUniversal(event: SetUniversalEvent): void {
  let entity = new SetUniversal(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.universal = event.params.universalAddress

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleTokenMinted(event: TokenMintedEvent): void {
  let entity = new TokenMinted(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.to = event.params.to
  entity.tokenId = event.params.tokenId
  entity.uri = event.params.uri

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleTokenTransfer(event: TokenTransferEvent): void {
  let entity = new TokenTransfer(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.destination = event.params.destination
  entity.receiver = event.params.receiver
  entity.tokenId = event.params.tokenId
  entity.uri = event.params.uri

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleTokenTransferReceived(
  event: TokenTransferReceivedEvent,
): void {
  let entity = new TokenTransferReceived(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.receiver = event.params.receiver
  entity.tokenId = event.params.tokenId
  entity.uri = event.params.uri

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleTokenTransferReverted(
  event: TokenTransferRevertedEvent,
): void {
  let entity = new TokenTransferReverted(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.sender = event.params.sender
  entity.tokenId = event.params.tokenId
  entity.uri = event.params.uri

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleTokenTransferToDestination(
  event: TokenTransferToDestinationEvent,
): void {
  let entity = new TokenTransferToDestination(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.destination = event.params.destination
  entity.sender = event.params.sender
  entity.tokenId = event.params.tokenId
  entity.uri = event.params.uri

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleTransfer(event: TransferEvent): void {
  let entity = new Transfer(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.from = event.params.from
  entity.to = event.params.to
  entity.tokenId = event.params.tokenId

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleUpgraded(event: UpgradedEvent): void {
  let entity = new Upgraded(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.implementation = event.params.implementation

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
