import {
  AdminChanged as AdminChangedEvent,
  FeeWithdrawn as FeeWithdrawnEvent,
} from "../generated/FeeVault/FeeVault"
import { AdminChanged, FeeWithdrawn } from "../generated/schema"

export function handleAdminChanged(event: AdminChangedEvent): void {
  let entity = new AdminChanged(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.previousAdmin = event.params.previousAdmin
  entity.newAdmin = event.params.newAdmin

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleFeeWithdrawn(event: FeeWithdrawnEvent): void {
  let entity = new FeeWithdrawn(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.recipient = event.params.recipient
  entity.amount = event.params.amount

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
