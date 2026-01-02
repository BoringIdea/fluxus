'use client'
import React from "react"
import { Button } from "@/components/ui/button"
import { getBatchBuyPrice } from "@/lib/price"
import { formatNumberWithMaxDecimalsAndRounding } from "@/src/utils"

interface BatchBuyPanelProps {
  priceSymbol: string
  maxSupply: string
  currentSupply: string
  initialPrice: string
  creatorFee: string
  maxSweep: number
  batchCount: number
  selectedTokens: number[]
  onCountChange: (count: number) => void
  onBatchBuy: () => void
}

export function BatchBuyPanel({
  priceSymbol,
  maxSupply,
  currentSupply,
  initialPrice,
  creatorFee,
  maxSweep,
  batchCount,
  selectedTokens,
  onCountChange,
  onBatchBuy
}: BatchBuyPanelProps) {
  return (
    <div className="fixed bottom-[53px] left-0 right-0 bg-background/95 backdrop-blur md:p-3 p-2 border-t border-border flex flex-col sm:flex-row items-stretch sm:items-center justify-center gap-3 sm:gap-6 z-50 text-xs">
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
      {/* Batch buy button */}
      <Button
        onClick={onBatchBuy}
        disabled={selectedTokens.length === 0}
        className="w-full sm:w-auto mt-1 sm:mt-0 text-xs tracking-[0.2em]"
      >
        Buy {batchCount} Items ({formatNumberWithMaxDecimalsAndRounding(Number(getBatchBuyPrice(maxSupply, currentSupply, initialPrice, batchCount, creatorFee)) / 1e18, 2)} {priceSymbol})
      </Button>
    </div>
  )
}
