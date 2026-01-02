import {
  assert,
  describe,
  test,
  clearStore,
  beforeAll,
  afterAll
} from "matchstick-as/assembly/index"
import { Address, BigInt } from "@graphprotocol/graph-ts"
import { FluxusCreated } from "../generated/schema"
import { FluxusCreated as FluxusCreatedEvent } from "../generated/Factory/Factory"
import { handleFluxusCreated } from "../src/factory"
import { createFluxusCreatedEvent } from "./factory-utils"

// Tests structure (matchstick-as >=0.5.0)
// https://thegraph.com/docs/en/subgraphs/developing/creating/unit-testing-framework/#tests-structure

describe("Describe entity assertions", () => {
  beforeAll(() => {
    let creator = Address.fromString(
      "0x0000000000000000000000000000000000000001"
    )
    let fluxusAddress = Address.fromString(
      "0x0000000000000000000000000000000000000001"
    )
    let priceAddress = Address.fromString(
      "0x0000000000000000000000000000000000000001"
    )
    let name = "Example string value"
    let symbol = "Example string value"
    let initialPrice = BigInt.fromI32(234)
    let maxSupply = BigInt.fromI32(234)
    let maxPrice = BigInt.fromI32(234)
    let creatorFeePercent = BigInt.fromI32(234)
    let baseUri = "Example string value"
    let newFluxusCreatedEvent = createFluxusCreatedEvent(
      creator,
      fluxusAddress,
      priceAddress,
      name,
      symbol,
      initialPrice,
      maxSupply,
      maxPrice,
      creatorFeePercent,
      baseUri
    )
    handleFluxusCreated(newFluxusCreatedEvent)
  })

  afterAll(() => {
    clearStore()
  })

  // For more test scenarios, see:
  // https://thegraph.com/docs/en/subgraphs/developing/creating/unit-testing-framework/#write-a-unit-test

  test("FluxusCreated created and stored", () => {
    assert.entityCount("FluxusCreated", 1)

    // 0xa16081f360e3847006db660bae1c6d1b2e17ec2a is the default address used in newMockEvent() function
    assert.fieldEquals(
      "FluxusCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "creator",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "FluxusCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "fluxusAddress",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "FluxusCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "priceAddress",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "FluxusCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "name",
      "Example string value"
    )
    assert.fieldEquals(
      "FluxusCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "symbol",
      "Example string value"
    )
    assert.fieldEquals(
      "FluxusCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "initialPrice",
      "234"
    )
    assert.fieldEquals(
      "FluxusCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "maxSupply",
      "234"
    )
    assert.fieldEquals(
      "FluxusCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "maxPrice",
      "234"
    )
    assert.fieldEquals(
      "FluxusCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "creatorFeePercent",
      "234"
    )
    assert.fieldEquals(
      "FluxusCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "baseUri",
      "Example string value"
    )

    // More assert options:
    // https://thegraph.com/docs/en/subgraphs/developing/creating/unit-testing-framework/#asserts
  })
})
