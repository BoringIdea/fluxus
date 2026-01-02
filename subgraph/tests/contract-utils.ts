import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt } from "@graphprotocol/graph-ts"
import {
  Approval,
  ApprovalForAll,
  BatchMetadataUpdate,
  Initialized,
  MetadataUpdate,
  OwnershipTransferred,
  SetConnected,
  SetUniversal,
  TokenMinted,
  TokenTransfer,
  TokenTransferReceived,
  TokenTransferReverted,
  TokenTransferToDestination,
  Transfer,
  Upgraded
} from "../generated/Contract/Contract"

export function createApprovalEvent(
  owner: Address,
  approved: Address,
  tokenId: BigInt
): Approval {
  let approvalEvent = changetype<Approval>(newMockEvent())

  approvalEvent.parameters = new Array()

  approvalEvent.parameters.push(
    new ethereum.EventParam("owner", ethereum.Value.fromAddress(owner))
  )
  approvalEvent.parameters.push(
    new ethereum.EventParam("approved", ethereum.Value.fromAddress(approved))
  )
  approvalEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )

  return approvalEvent
}

export function createApprovalForAllEvent(
  owner: Address,
  operator: Address,
  approved: boolean
): ApprovalForAll {
  let approvalForAllEvent = changetype<ApprovalForAll>(newMockEvent())

  approvalForAllEvent.parameters = new Array()

  approvalForAllEvent.parameters.push(
    new ethereum.EventParam("owner", ethereum.Value.fromAddress(owner))
  )
  approvalForAllEvent.parameters.push(
    new ethereum.EventParam("operator", ethereum.Value.fromAddress(operator))
  )
  approvalForAllEvent.parameters.push(
    new ethereum.EventParam("approved", ethereum.Value.fromBoolean(approved))
  )

  return approvalForAllEvent
}

export function createBatchMetadataUpdateEvent(
  _fromTokenId: BigInt,
  _toTokenId: BigInt
): BatchMetadataUpdate {
  let batchMetadataUpdateEvent = changetype<BatchMetadataUpdate>(newMockEvent())

  batchMetadataUpdateEvent.parameters = new Array()

  batchMetadataUpdateEvent.parameters.push(
    new ethereum.EventParam(
      "_fromTokenId",
      ethereum.Value.fromUnsignedBigInt(_fromTokenId)
    )
  )
  batchMetadataUpdateEvent.parameters.push(
    new ethereum.EventParam(
      "_toTokenId",
      ethereum.Value.fromUnsignedBigInt(_toTokenId)
    )
  )

  return batchMetadataUpdateEvent
}

export function createInitializedEvent(version: BigInt): Initialized {
  let initializedEvent = changetype<Initialized>(newMockEvent())

  initializedEvent.parameters = new Array()

  initializedEvent.parameters.push(
    new ethereum.EventParam(
      "version",
      ethereum.Value.fromUnsignedBigInt(version)
    )
  )

  return initializedEvent
}

export function createMetadataUpdateEvent(_tokenId: BigInt): MetadataUpdate {
  let metadataUpdateEvent = changetype<MetadataUpdate>(newMockEvent())

  metadataUpdateEvent.parameters = new Array()

  metadataUpdateEvent.parameters.push(
    new ethereum.EventParam(
      "_tokenId",
      ethereum.Value.fromUnsignedBigInt(_tokenId)
    )
  )

  return metadataUpdateEvent
}

export function createOwnershipTransferredEvent(
  previousOwner: Address,
  newOwner: Address
): OwnershipTransferred {
  let ownershipTransferredEvent =
    changetype<OwnershipTransferred>(newMockEvent())

  ownershipTransferredEvent.parameters = new Array()

  ownershipTransferredEvent.parameters.push(
    new ethereum.EventParam(
      "previousOwner",
      ethereum.Value.fromAddress(previousOwner)
    )
  )
  ownershipTransferredEvent.parameters.push(
    new ethereum.EventParam("newOwner", ethereum.Value.fromAddress(newOwner))
  )

  return ownershipTransferredEvent
}

export function createSetConnectedEvent(
  zrc20: Address,
  contractAddress: Address
): SetConnected {
  let setConnectedEvent = changetype<SetConnected>(newMockEvent())

  setConnectedEvent.parameters = new Array()

  setConnectedEvent.parameters.push(
    new ethereum.EventParam("zrc20", ethereum.Value.fromAddress(zrc20))
  )
  setConnectedEvent.parameters.push(
    new ethereum.EventParam(
      "contractAddress",
      ethereum.Value.fromAddress(contractAddress)
    )
  )

  return setConnectedEvent
}

export function createSetUniversalEvent(
  universalAddress: Address
): SetUniversal {
  let setUniversalEvent = changetype<SetUniversal>(newMockEvent())

  setUniversalEvent.parameters = new Array()

  setUniversalEvent.parameters.push(
    new ethereum.EventParam(
      "universalAddress",
      ethereum.Value.fromAddress(universalAddress)
    )
  )

  return setUniversalEvent
}

export function createTokenMintedEvent(
  to: Address,
  tokenId: BigInt,
  uri: string
): TokenMinted {
  let tokenMintedEvent = changetype<TokenMinted>(newMockEvent())

  tokenMintedEvent.parameters = new Array()

  tokenMintedEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  tokenMintedEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  tokenMintedEvent.parameters.push(
    new ethereum.EventParam("uri", ethereum.Value.fromString(uri))
  )

  return tokenMintedEvent
}

export function createTokenTransferEvent(
  destination: Address,
  receiver: Address,
  tokenId: BigInt,
  uri: string
): TokenTransfer {
  let tokenTransferEvent = changetype<TokenTransfer>(newMockEvent())

  tokenTransferEvent.parameters = new Array()

  tokenTransferEvent.parameters.push(
    new ethereum.EventParam(
      "destination",
      ethereum.Value.fromAddress(destination)
    )
  )
  tokenTransferEvent.parameters.push(
    new ethereum.EventParam("receiver", ethereum.Value.fromAddress(receiver))
  )
  tokenTransferEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  tokenTransferEvent.parameters.push(
    new ethereum.EventParam("uri", ethereum.Value.fromString(uri))
  )

  return tokenTransferEvent
}

export function createTokenTransferReceivedEvent(
  receiver: Address,
  tokenId: BigInt,
  uri: string
): TokenTransferReceived {
  let tokenTransferReceivedEvent =
    changetype<TokenTransferReceived>(newMockEvent())

  tokenTransferReceivedEvent.parameters = new Array()

  tokenTransferReceivedEvent.parameters.push(
    new ethereum.EventParam("receiver", ethereum.Value.fromAddress(receiver))
  )
  tokenTransferReceivedEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  tokenTransferReceivedEvent.parameters.push(
    new ethereum.EventParam("uri", ethereum.Value.fromString(uri))
  )

  return tokenTransferReceivedEvent
}

export function createTokenTransferRevertedEvent(
  sender: Address,
  tokenId: BigInt,
  uri: string
): TokenTransferReverted {
  let tokenTransferRevertedEvent =
    changetype<TokenTransferReverted>(newMockEvent())

  tokenTransferRevertedEvent.parameters = new Array()

  tokenTransferRevertedEvent.parameters.push(
    new ethereum.EventParam("sender", ethereum.Value.fromAddress(sender))
  )
  tokenTransferRevertedEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  tokenTransferRevertedEvent.parameters.push(
    new ethereum.EventParam("uri", ethereum.Value.fromString(uri))
  )

  return tokenTransferRevertedEvent
}

export function createTokenTransferToDestinationEvent(
  destination: Address,
  sender: Address,
  tokenId: BigInt,
  uri: string
): TokenTransferToDestination {
  let tokenTransferToDestinationEvent =
    changetype<TokenTransferToDestination>(newMockEvent())

  tokenTransferToDestinationEvent.parameters = new Array()

  tokenTransferToDestinationEvent.parameters.push(
    new ethereum.EventParam(
      "destination",
      ethereum.Value.fromAddress(destination)
    )
  )
  tokenTransferToDestinationEvent.parameters.push(
    new ethereum.EventParam("sender", ethereum.Value.fromAddress(sender))
  )
  tokenTransferToDestinationEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  tokenTransferToDestinationEvent.parameters.push(
    new ethereum.EventParam("uri", ethereum.Value.fromString(uri))
  )

  return tokenTransferToDestinationEvent
}

export function createTransferEvent(
  from: Address,
  to: Address,
  tokenId: BigInt
): Transfer {
  let transferEvent = changetype<Transfer>(newMockEvent())

  transferEvent.parameters = new Array()

  transferEvent.parameters.push(
    new ethereum.EventParam("from", ethereum.Value.fromAddress(from))
  )
  transferEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  transferEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )

  return transferEvent
}

export function createUpgradedEvent(implementation: Address): Upgraded {
  let upgradedEvent = changetype<Upgraded>(newMockEvent())

  upgradedEvent.parameters = new Array()

  upgradedEvent.parameters.push(
    new ethereum.EventParam(
      "implementation",
      ethereum.Value.fromAddress(implementation)
    )
  )

  return upgradedEvent
}
