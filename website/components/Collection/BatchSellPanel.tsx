'use client'
import React from "react"
import { Button } from "@/components/ui/button"
import { getBatchSellPrice } from "@/lib/price"
import { formatNumberWithMaxDecimalsAndRounding } from "@/src/utils"

interface BatchSellPanelProps {
  pricerSymbol: string | 'ETH'
  maxSupply: string
  currentSupply: string
  initialPrice: string
  creatorFee: string
  maxSweep: number
  batchCount: number
  selectedTokens: number[]
  onCountChange: (count: number) => void
  onBatchSell: () => void
  isApproved?: boolean
  isPending?: boolean
  pendingBatchSell?: boolean
}

export function BatchSellPanel({
  pricerSymbol,
  maxSupply,
  currentSupply,
  initialPrice,
  creatorFee,
  maxSweep,
  batchCount,
  selectedTokens,
  onCountChange,
  onBatchSell,
  isApproved,
  isPending,
  pendingBatchSell
}: BatchSellPanelProps) {
  return (
    <div className="fixed bottom-[53px] left-0 right-0 bg-background/95 backdrop-blur md:p-3 p-2 border-t border-border flex flex-col sm:flex-row items-stretch sm:items-center justify-center gap-3 sm:gap-6 z-50 text-xs">
      {/* Slider with count display */}
      <div className="flex flex-col sm:flex-row items-center border border-border px-2 py-2 rounded-none w-full sm:w-auto gap-2 bg-bg-tertiary">
        <div className="text-[10px] uppercase tracking-[0.3em] text-secondary">Sweep</div>
        <div className="flex-1 w-full sm:min-w-[320px]">
          <input
            type="range"
            min={0}
            max={maxSweep}
            step={1}
            value={batchCount}
            onChange={(event) => onCountChange(Number(event.target.value))}
            className="w-full accent-fluxus-primary"
          />
        </div>
        <span className="text-primary font-bold w-10 text-center">{batchCount}</span>
      </div>

      {/* Batch sell button */}
      <Button
        onClick={onBatchSell}
        disabled={selectedTokens.length === 0 || isPending}
        className="w-full sm:w-auto mt-1 sm:mt-0 text-xs tracking-[0.2em]"
      >
        {isPending
          ? (pendingBatchSell ? 'Approving...' : 'Selling...')
          : (isApproved 
              ? `Sell ${batchCount} Items (${formatNumberWithMaxDecimalsAndRounding(Number(getBatchSellPrice(maxSupply, currentSupply, initialPrice, batchCount, creatorFee)) / 1e18, 2)} ${pricerSymbol})`
              : `Approve and Sell ${batchCount} Items`
            )
        }
      </Button>
    </div>
  )
}
