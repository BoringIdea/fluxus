import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt } from "@graphprotocol/graph-ts"
import {
  Approval,
  ApprovalForAll,
  Initialized,
  OwnershipTransferred,
  SetConnected,
  SetUniversal,
  TokenBought,
  TokenMinted,
  TokenMinted1,
  TokenSold,
  TokenTransfer,
  TokenTransferReceived,
  TokenTransferReverted,
  TokenTransferToDestination,
  Transfer,
  Upgraded
} from "../generated/ZetaChainFluxusCrossChainOptimized/ZetaChainFluxusCrossChainOptimized"

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

export function createTokenBoughtEvent(
  buyer: Address,
  tokenId: BigInt,
  price: BigInt,
  creatorFee: BigInt
): TokenBought {
  let tokenBoughtEvent = changetype<TokenBought>(newMockEvent())

  tokenBoughtEvent.parameters = new Array()

  tokenBoughtEvent.parameters.push(
    new ethereum.EventParam("buyer", ethereum.Value.fromAddress(buyer))
  )
  tokenBoughtEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  tokenBoughtEvent.parameters.push(
    new ethereum.EventParam("price", ethereum.Value.fromUnsignedBigInt(price))
  )
  tokenBoughtEvent.parameters.push(
    new ethereum.EventParam(
      "creatorFee",
      ethereum.Value.fromUnsignedBigInt(creatorFee)
    )
  )

  return tokenBoughtEvent
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

export function createTokenMinted1Event(
  to: Address,
  tokenId: BigInt,
  price: BigInt,
  creatorFee: BigInt
): TokenMinted1 {
  let tokenMinted1Event = changetype<TokenMinted1>(newMockEvent())

  tokenMinted1Event.parameters = new Array()

  tokenMinted1Event.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  tokenMinted1Event.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  tokenMinted1Event.parameters.push(
    new ethereum.EventParam("price", ethereum.Value.fromUnsignedBigInt(price))
  )
  tokenMinted1Event.parameters.push(
    new ethereum.EventParam(
      "creatorFee",
      ethereum.Value.fromUnsignedBigInt(creatorFee)
    )
  )

  return tokenMinted1Event
}

export function createTokenSoldEvent(
  seller: Address,
  tokenId: BigInt,
  price: BigInt,
  creatorFee: BigInt
): TokenSold {
  let tokenSoldEvent = changetype<TokenSold>(newMockEvent())

  tokenSoldEvent.parameters = new Array()

  tokenSoldEvent.parameters.push(
    new ethereum.EventParam("seller", ethereum.Value.fromAddress(seller))
  )
  tokenSoldEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  tokenSoldEvent.parameters.push(
    new ethereum.EventParam("price", ethereum.Value.fromUnsignedBigInt(price))
  )
  tokenSoldEvent.parameters.push(
    new ethereum.EventParam(
      "creatorFee",
      ethereum.Value.fromUnsignedBigInt(creatorFee)
    )
  )

  return tokenSoldEvent
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
