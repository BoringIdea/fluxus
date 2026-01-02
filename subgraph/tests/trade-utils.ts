import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt } from "@graphprotocol/graph-ts"
import {
  Bought,
  BulkBuyExecuted,
  BulkMintExecuted,
  BulkQuickBuyExecuted,
  BulkSellExecuted,
  Minted,
  QuickBuyExecuted,
  Sold,
  TransferCrossChain
} from "../generated/Trade/Trade"

export function createBoughtEvent(
  fluxusContract: Address,
  buyer: Address,
  tokenId: BigInt,
  price: BigInt
): Bought {
  let boughtEvent = changetype<Bought>(newMockEvent())

  boughtEvent.parameters = new Array()

  boughtEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusContract",
      ethereum.Value.fromAddress(fluxusContract)
    )
  )
  boughtEvent.parameters.push(
    new ethereum.EventParam("buyer", ethereum.Value.fromAddress(buyer))
  )
  boughtEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  boughtEvent.parameters.push(
    new ethereum.EventParam("price", ethereum.Value.fromUnsignedBigInt(price))
  )

  return boughtEvent
}

export function createBulkBuyExecutedEvent(
  fluxusContract: Address,
  buyer: Address,
  tokenIds: Array<BigInt>,
  totalPrice: BigInt
): BulkBuyExecuted {
  let bulkBuyExecutedEvent = changetype<BulkBuyExecuted>(newMockEvent())

  bulkBuyExecutedEvent.parameters = new Array()

  bulkBuyExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusContract",
      ethereum.Value.fromAddress(fluxusContract)
    )
  )
  bulkBuyExecutedEvent.parameters.push(
    new ethereum.EventParam("buyer", ethereum.Value.fromAddress(buyer))
  )
  bulkBuyExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "tokenIds",
      ethereum.Value.fromUnsignedBigIntArray(tokenIds)
    )
  )
  bulkBuyExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "totalPrice",
      ethereum.Value.fromUnsignedBigInt(totalPrice)
    )
  )

  return bulkBuyExecutedEvent
}

export function createBulkMintExecutedEvent(
  fluxusContract: Address,
  buyer: Address,
  tokenIds: Array<BigInt>,
  totalPrice: BigInt
): BulkMintExecuted {
  let bulkMintExecutedEvent = changetype<BulkMintExecuted>(newMockEvent())

  bulkMintExecutedEvent.parameters = new Array()

  bulkMintExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusContract",
      ethereum.Value.fromAddress(fluxusContract)
    )
  )
  bulkMintExecutedEvent.parameters.push(
    new ethereum.EventParam("buyer", ethereum.Value.fromAddress(buyer))
  )
  bulkMintExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "tokenIds",
      ethereum.Value.fromUnsignedBigIntArray(tokenIds)
    )
  )
  bulkMintExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "totalPrice",
      ethereum.Value.fromUnsignedBigInt(totalPrice)
    )
  )

  return bulkMintExecutedEvent
}

export function createBulkQuickBuyExecutedEvent(
  fluxusContract: Address,
  buyer: Address,
  tokenIds: Array<BigInt>,
  totalPrice: BigInt
): BulkQuickBuyExecuted {
  let bulkQuickBuyExecutedEvent =
    changetype<BulkQuickBuyExecuted>(newMockEvent())

  bulkQuickBuyExecutedEvent.parameters = new Array()

  bulkQuickBuyExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusContract",
      ethereum.Value.fromAddress(fluxusContract)
    )
  )
  bulkQuickBuyExecutedEvent.parameters.push(
    new ethereum.EventParam("buyer", ethereum.Value.fromAddress(buyer))
  )
  bulkQuickBuyExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "tokenIds",
      ethereum.Value.fromUnsignedBigIntArray(tokenIds)
    )
  )
  bulkQuickBuyExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "totalPrice",
      ethereum.Value.fromUnsignedBigInt(totalPrice)
    )
  )

  return bulkQuickBuyExecutedEvent
}

export function createBulkSellExecutedEvent(
  fluxusContract: Address,
  seller: Address,
  tokenIds: Array<BigInt>,
  totalPrice: BigInt
): BulkSellExecuted {
  let bulkSellExecutedEvent = changetype<BulkSellExecuted>(newMockEvent())

  bulkSellExecutedEvent.parameters = new Array()

  bulkSellExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusContract",
      ethereum.Value.fromAddress(fluxusContract)
    )
  )
  bulkSellExecutedEvent.parameters.push(
    new ethereum.EventParam("seller", ethereum.Value.fromAddress(seller))
  )
  bulkSellExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "tokenIds",
      ethereum.Value.fromUnsignedBigIntArray(tokenIds)
    )
  )
  bulkSellExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "totalPrice",
      ethereum.Value.fromUnsignedBigInt(totalPrice)
    )
  )

  return bulkSellExecutedEvent
}

export function createMintedEvent(
  fluxusContract: Address,
  to: Address,
  tokenId: BigInt,
  price: BigInt
): Minted {
  let mintedEvent = changetype<Minted>(newMockEvent())

  mintedEvent.parameters = new Array()

  mintedEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusContract",
      ethereum.Value.fromAddress(fluxusContract)
    )
  )
  mintedEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  mintedEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  mintedEvent.parameters.push(
    new ethereum.EventParam("price", ethereum.Value.fromUnsignedBigInt(price))
  )

  return mintedEvent
}

export function createQuickBuyExecutedEvent(
  fluxusContract: Address,
  buyer: Address,
  tokenId: BigInt,
  price: BigInt
): QuickBuyExecuted {
  let quickBuyExecutedEvent = changetype<QuickBuyExecuted>(newMockEvent())

  quickBuyExecutedEvent.parameters = new Array()

  quickBuyExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusContract",
      ethereum.Value.fromAddress(fluxusContract)
    )
  )
  quickBuyExecutedEvent.parameters.push(
    new ethereum.EventParam("buyer", ethereum.Value.fromAddress(buyer))
  )
  quickBuyExecutedEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  quickBuyExecutedEvent.parameters.push(
    new ethereum.EventParam("price", ethereum.Value.fromUnsignedBigInt(price))
  )

  return quickBuyExecutedEvent
}

export function createSoldEvent(
  fluxusContract: Address,
  seller: Address,
  tokenId: BigInt,
  price: BigInt
): Sold {
  let soldEvent = changetype<Sold>(newMockEvent())

  soldEvent.parameters = new Array()

  soldEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusContract",
      ethereum.Value.fromAddress(fluxusContract)
    )
  )
  soldEvent.parameters.push(
    new ethereum.EventParam("seller", ethereum.Value.fromAddress(seller))
  )
  soldEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  soldEvent.parameters.push(
    new ethereum.EventParam("price", ethereum.Value.fromUnsignedBigInt(price))
  )

  return soldEvent
}

export function createTransferCrossChainEvent(
  fluxusContract: Address,
  sender: Address,
  tokenId: BigInt,
  receiver: Address,
  destination: Address
): TransferCrossChain {
  let transferCrossChainEvent = changetype<TransferCrossChain>(newMockEvent())

  transferCrossChainEvent.parameters = new Array()

  transferCrossChainEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusContract",
      ethereum.Value.fromAddress(fluxusContract)
    )
  )
  transferCrossChainEvent.parameters.push(
    new ethereum.EventParam("sender", ethereum.Value.fromAddress(sender))
  )
  transferCrossChainEvent.parameters.push(
    new ethereum.EventParam(
      "tokenId",
      ethereum.Value.fromUnsignedBigInt(tokenId)
    )
  )
  transferCrossChainEvent.parameters.push(
    new ethereum.EventParam("receiver", ethereum.Value.fromAddress(receiver))
  )
  transferCrossChainEvent.parameters.push(
    new ethereum.EventParam(
      "destination",
      ethereum.Value.fromAddress(destination)
    )
  )

  return transferCrossChainEvent
}
