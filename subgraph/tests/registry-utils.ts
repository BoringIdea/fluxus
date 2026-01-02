import { newMockEvent } from "matchstick-as"
import { ethereum, Address } from "@graphprotocol/graph-ts"
import { ContractRegistered } from "../generated/Registry/Registry"

export function createContractRegisteredEvent(
  creator: Address,
  contractAddress: Address
): ContractRegistered {
  let contractRegisteredEvent = changetype<ContractRegistered>(newMockEvent())

  contractRegisteredEvent.parameters = new Array()

  contractRegisteredEvent.parameters.push(
    new ethereum.EventParam("creator", ethereum.Value.fromAddress(creator))
  )
  contractRegisteredEvent.parameters.push(
    new ethereum.EventParam(
      "contractAddress",
      ethereum.Value.fromAddress(contractAddress)
    )
  )

  return contractRegisteredEvent
}
