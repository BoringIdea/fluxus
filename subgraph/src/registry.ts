import { ContractRegistered as ContractRegisteredEvent } from "../generated/Registry/Registry"
import { ContractRegistered } from "../generated/schema"

export function handleContractRegistered(event: ContractRegisteredEvent): void {
  let entity = new ContractRegistered(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.creator = event.params.creator
  entity.contractAddress = event.params.contractAddress

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
