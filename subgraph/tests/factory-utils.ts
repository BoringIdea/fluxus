import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt } from "@graphprotocol/graph-ts"
import {
  FluxusCreated,
  FluxusCrossChainCreated,
  SetGasLimit,
  SetGateway,
  SetUniversal
} from "../generated/Factory/Factory"

export function createFluxusCreatedEvent(
  creator: Address,
  fluxusAddress: Address,
  priceAddress: Address,
  name: string,
  symbol: string,
  initialPrice: BigInt,
  maxSupply: BigInt,
  maxPrice: BigInt,
  creatorFeePercent: BigInt,
  baseUri: string
): FluxusCreated {
  let fluxusCreatedEvent = changetype<FluxusCreated>(newMockEvent())

  fluxusCreatedEvent.parameters = new Array()

  fluxusCreatedEvent.parameters.push(
    new ethereum.EventParam("creator", ethereum.Value.fromAddress(creator))
  )
  fluxusCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusAddress",
      ethereum.Value.fromAddress(fluxusAddress)
    )
  )
  fluxusCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "priceAddress",
      ethereum.Value.fromAddress(priceAddress)
    )
  )
  fluxusCreatedEvent.parameters.push(
    new ethereum.EventParam("name", ethereum.Value.fromString(name))
  )
  fluxusCreatedEvent.parameters.push(
    new ethereum.EventParam("symbol", ethereum.Value.fromString(symbol))
  )
  fluxusCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "initialPrice",
      ethereum.Value.fromUnsignedBigInt(initialPrice)
    )
  )
  fluxusCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "maxSupply",
      ethereum.Value.fromUnsignedBigInt(maxSupply)
    )
  )
  fluxusCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "maxPrice",
      ethereum.Value.fromUnsignedBigInt(maxPrice)
    )
  )
  fluxusCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "creatorFeePercent",
      ethereum.Value.fromUnsignedBigInt(creatorFeePercent)
    )
  )
  fluxusCreatedEvent.parameters.push(
    new ethereum.EventParam("baseUri", ethereum.Value.fromString(baseUri))
  )

  return fluxusCreatedEvent
}

export function createFluxusCrossChainCreatedEvent(
  creator: Address,
  fluxusAddress: Address,
  priceAddress: Address,
  name: string,
  symbol: string,
  initialPrice: BigInt,
  maxSupply: BigInt,
  maxPrice: BigInt,
  creatorFeePercent: BigInt,
  baseUri: string,
  gatewayAddress: Address,
  gasLimit: BigInt,
  supportMint: boolean
): FluxusCrossChainCreated {
  let fluxusCrossChainCreatedEvent =
    changetype<FluxusCrossChainCreated>(newMockEvent())

  fluxusCrossChainCreatedEvent.parameters = new Array()

  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam("creator", ethereum.Value.fromAddress(creator))
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "fluxusAddress",
      ethereum.Value.fromAddress(fluxusAddress)
    )
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "priceAddress",
      ethereum.Value.fromAddress(priceAddress)
    )
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam("name", ethereum.Value.fromString(name))
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam("symbol", ethereum.Value.fromString(symbol))
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "initialPrice",
      ethereum.Value.fromUnsignedBigInt(initialPrice)
    )
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "maxSupply",
      ethereum.Value.fromUnsignedBigInt(maxSupply)
    )
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "maxPrice",
      ethereum.Value.fromUnsignedBigInt(maxPrice)
    )
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "creatorFeePercent",
      ethereum.Value.fromUnsignedBigInt(creatorFeePercent)
    )
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam("baseUri", ethereum.Value.fromString(baseUri))
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "gatewayAddress",
      ethereum.Value.fromAddress(gatewayAddress)
    )
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "gasLimit",
      ethereum.Value.fromUnsignedBigInt(gasLimit)
    )
  )
  fluxusCrossChainCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "supportMint",
      ethereum.Value.fromBoolean(supportMint)
    )
  )

  return fluxusCrossChainCreatedEvent
}

export function createSetGasLimitEvent(
  fluxus: Address,
  gasLimit: BigInt
): SetGasLimit {
  let setGasLimitEvent = changetype<SetGasLimit>(newMockEvent())

  setGasLimitEvent.parameters = new Array()

  setGasLimitEvent.parameters.push(
    new ethereum.EventParam("fluxus", ethereum.Value.fromAddress(fluxus))
  )
  setGasLimitEvent.parameters.push(
    new ethereum.EventParam(
      "gasLimit",
      ethereum.Value.fromUnsignedBigInt(gasLimit)
    )
  )

  return setGasLimitEvent
}

export function createSetGatewayEvent(
  fluxus: Address,
  gateway: Address
): SetGateway {
  let setGatewayEvent = changetype<SetGateway>(newMockEvent())

  setGatewayEvent.parameters = new Array()

  setGatewayEvent.parameters.push(
    new ethereum.EventParam("fluxus", ethereum.Value.fromAddress(fluxus))
  )
  setGatewayEvent.parameters.push(
    new ethereum.EventParam("gateway", ethereum.Value.fromAddress(gateway))
  )

  return setGatewayEvent
}

export function createSetUniversalEvent(
  fluxus: Address,
  universal: Address
): SetUniversal {
  let setUniversalEvent = changetype<SetUniversal>(newMockEvent())

  setUniversalEvent.parameters = new Array()

  setUniversalEvent.parameters.push(
    new ethereum.EventParam("fluxus", ethereum.Value.fromAddress(fluxus))
  )
  setUniversalEvent.parameters.push(
    new ethereum.EventParam("universal", ethereum.Value.fromAddress(universal))
  )

  return setUniversalEvent
}
