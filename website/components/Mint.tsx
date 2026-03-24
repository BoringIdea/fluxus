"use client";
import { useEffect, useMemo, useState } from "react";
import {
  useReadContract,
  useWriteContract,
  useWaitForTransactionReceipt,
  useWatchBlockNumber,
  useChainId
} from "wagmi";
import { ethers } from "ethers";
import { Button } from "./ui/button";
import { Progress } from "@/components/ui/progress"
import { PRICE_ABI, TRADE_ABI, getTradeContractAddress } from "@/src/contract";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
} from 'chart.js';
import { SupportedChainId } from "@/src/alchemy";
import { Collection } from "@/src/api/types";
import TransactionDialog from "@/components/shared/TransactionDialog"
import PriceChart from "@/components/Collection/PriceChart";
import { useTransactionDialog } from "@/src/hooks/useTransactionDialog";
import { getBatchBuyPrice } from "@/lib/price";
import { getChainSymbol, formatNumberWithMaxDecimalsAndRounding } from "@/src/utils";
import { fetchCollectionImage } from "@/lib/utils";

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

interface MintProps {
  contractAddress: string;
  collection: Collection | undefined;
}

export default function Mint({ contractAddress, collection }: MintProps) {
  const chainId = useChainId();

  let alchemy: any;
  // if (chainId !== monadTestnet.id) {
  //   alchemy = getAlchemy(chainId as SupportedChainId);
  // }

  const [mintAmount, setMintAmount] = useState<number>(1);
  const { data: hash, writeContract, isError: isWriteContractError, error: writeContractError, isPending } = useWriteContract();

  const [collectionImage, setCollectionImage] = useState<string | null>(null);
  useEffect(() => {
    if (collection) {
      fetchCollectionImage(collection).then((image) => {
        setCollectionImage(image);
      }).catch((error) => {
        console.error('Error fetching collection image:', error);
        setCollectionImage(null);
      });
    }
  }, [collection]);

  // Read contract mintPrice
  const { data: mintPrice, refetch: refetchMintPrice } = useReadContract({
    // @ts-ignore
    address: collection?.price_contract,
    abi: PRICE_ABI,
    functionName: 'getBuyPriceAfterFee',
    args: [
      collection?.address,
    ],
    scopeKey: `mintPrice:${collection?.address}`,
    query: {
      enabled: !!collection?.price_contract && !!collection?.address,
      staleTime: 30000,
      gcTime: 300000,
      refetchOnWindowFocus: false,
      retry: 1,
      refetchOnMount: 'always',
    }
  });

  const fallbackMintPrice = useMemo(() => {
    if (!collection?.max_supply || !collection?.current_supply || !collection?.initial_price) return null;
    return getBatchBuyPrice(
      collection?.max_supply?.toString() || '0',
      collection?.current_supply?.toString() || '0',
      collection?.initial_price?.toString() || '0',
      1,
      collection?.creator_fee?.toString() || '0'
    );
  }, [collection?.max_supply, collection?.current_supply, collection?.initial_price, collection?.creator_fee]);

  const { isLoading: isConfirming, isSuccess: isConfirmed } = useWaitForTransactionReceipt({
    hash,
  });

  const [progress, setProgress] = useState(0);

  // Calculate mint progress
  useEffect(() => {
    if (collection?.current_supply !== undefined) {
      const supply = Number(collection?.current_supply);
      if (!isNaN(supply)) {
        const newProgress = (supply / 10000) * 100;
        setProgress(newProgress);
      }
    }
  }, [collection]);

  // Mint NFT function
  async function mint() {
    console.log('mintPrice', mintPrice);
    if (mintPrice && mintAmount == 1) {
      writeContract({
        // @ts-ignore
        address: getTradeContractAddress(chainId as SupportedChainId),
        abi: TRADE_ABI,
        functionName: 'mint',
        args: [
          contractAddress,
        ],
        value: BigInt(mintPrice.toString()),
      });
    } else if (mintAmount > 1) {
      const mintCost = getBatchBuyPrice(
        collection?.max_supply?.toString() || '0',
        collection?.current_supply?.toString() || '0',
        collection?.initial_price?.toString() || '0',
        mintAmount,
        collection?.creator_fee?.toString() || '0'
      )
      console.log('bulkMint cost', mintCost);
      writeContract({
        // @ts-ignore
        address: getTradeContractAddress(chainId as SupportedChainId),
        abi: TRADE_ABI,
        functionName: 'bulkMint',
        args: [
          contractAddress,
          mintAmount,
        ],
        value: mintCost,
      })
    }
  }

  const { dialogState, onOpenChange } = useTransactionDialog({
    hash,
    isPending,
    isConfirming,
    isConfirmed,
    isWriteContractError,
    writeContractError,
  });


  // Add a state to handle hydration
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  const formattedMintPrice = formatNumberWithMaxDecimalsAndRounding(
    ethers.formatEther((mintPrice ?? fallbackMintPrice ?? 0).toString())
  );

  return (
    <div className="flex w-full flex-col gap-6 border border-black/10 bg-[color:var(--bg-surface)] p-4 sm:flex-row sm:p-6">
      {/* Left Section */}
      <div className="flex flex-col gap-6 pb-12 lg:w-1/2">
        <div className="flex items-center justify-center border border-black/10 bg-[color:var(--bg-muted)] p-4">
          {collectionImage ? (
            <div className="w-[240px] h-[240px] sm:w-[300px] sm:h-[300px] overflow-hidden">
              <img src={collectionImage} alt={collection?.name} className="w-full h-full object-cover" />
            </div>
          ) : (
            <div className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">No preview</div>
          )}
        </div>

        <div className="flex flex-col gap-4 border border-black/10 bg-[color:var(--bg-muted)] p-5">
          <div>
            <p className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">Mint Price</p>
            <div className="flex flex-wrap items-baseline gap-2">
              <span className="font-heading text-[34px] leading-none text-[color:var(--text-primary)]">{formattedMintPrice}</span>
              <span className="text-sm text-[color:var(--text-secondary)]">{getChainSymbol(chainId)}</span>
            </div>
            <p className="mt-2 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
              Includes Fee:{" "}
              {collection?.creator_fee
                ? (Number(ethers.formatEther(collection?.creator_fee.toString())) * 100).toFixed(2)
                : "0"}
              %
            </p>
          </div>

          <Button
            onClick={() => mint()}
            disabled={isPending}
            className="h-12 w-full"
          >
            {isPending ? "Processing..." : "Mint"}
          </Button>

          <div className="flex flex-col gap-3">
            <div className="flex flex-col items-center justify-between gap-4 sm:flex-row">
              <div className="flex items-center border border-black/10 bg-[color:var(--bg-surface)]">
                <button
                  className="px-4 py-2 text-xl text-[color:var(--text-muted)] transition-colors hover:text-[color:var(--text-primary)]"
                  onClick={() => setMintAmount((prev) => Math.max(1, prev - 1))}
                >
                  -
                </button>
                <span className="border-x border-black/10 px-6 py-2 font-heading text-[28px] leading-none text-[color:var(--text-primary)]">
                  {mintAmount}
                </span>
                <button
                  className="px-4 py-2 text-xl text-[color:var(--text-muted)] transition-colors hover:text-[color:var(--text-primary)]"
                  onClick={() => setMintAmount((prev) => Math.min(20, prev + 1))}
                >
                  +
                </button>
              </div>
              <div className="w-full border border-black/10 bg-[color:var(--bg-surface)] px-4 py-3 sm:w-auto">
                <p className="mb-1 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Expected Cost</p>
                <p className="font-heading text-[24px] leading-none text-[color:var(--text-primary)]">
                  {mintAmount > 1
                    ? formatNumberWithMaxDecimalsAndRounding(
                        Number(
                          getBatchBuyPrice(
                            collection?.max_supply?.toString() || "0",
                            collection?.current_supply?.toString() || "0",
                            collection?.initial_price?.toString() || "0",
                            mintAmount,
                            collection?.creator_fee?.toString() || "0"
                          )
                        ) / 1e18,
                        2
                      )
                    : formatNumberWithMaxDecimalsAndRounding(
                        Number((mintPrice ?? fallbackMintPrice ?? 0).toString()) / 1e18,
                        2
                      )}{" "}
                  <span className="text-sm text-[color:var(--text-secondary)]">{getChainSymbol(chainId)}</span>
                </p>
              </div>
            </div>

            <input
              type="range"
              min="1"
              max="20"
              value={mintAmount}
              onChange={(e) => setMintAmount(Number(e.target.value))}
              className="h-2 w-full cursor-pointer appearance-none border border-black/10 bg-[color:var(--bg-surface)] [&::-moz-range-thumb]:h-4 [&::-moz-range-thumb]:w-4 [&::-moz-range-thumb]:border-0 [&::-moz-range-thumb]:bg-[color:var(--color-primary)] [&::-webkit-slider-thumb]:h-4 [&::-webkit-slider-thumb]:w-4 [&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:bg-[color:var(--color-primary)]"
            />
          </div>
        </div>

        <div className="border border-black/10 bg-[color:var(--bg-muted)] p-4">
          <Progress 
            value={progress} 
            className="h-3 w-full bg-[color:var(--bg-surface)]" 
            indicatorClassName="bg-[color:var(--color-primary)]"
          />
          <p className="mt-2 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)] sm:text-[11px]">
            MINTED: {collection?.total_supply?.toString()} / {collection?.max_supply?.toString()}
            <span className="ml-2 text-[color:var(--text-muted)]">
              (
              {collection?.max_supply && Number(collection?.max_supply) > 0
                ? (
                    (Number(collection?.total_supply || 0) / Number(collection?.max_supply)) * 100
                  ).toFixed(2).replace(/\.?0+$/, "")
                : "0"
              }
              %
              )
            </span>
          </p>
        </div>
      </div>

      <div className="flex w-full flex-col gap-4 lg:w-1/2">
        <div className="border-b border-black/10 pb-4">
          <p className="flux-kicker mb-2">Analytics</p>
          <h2 className="flux-h2">Mint Price Curve</h2>
          <p className="mt-2 text-sm text-[color:var(--text-secondary)]">
            Based on the set price, supply, and current mint position.
          </p>
        </div>
        <div className="flex h-[620px] flex-col overflow-hidden border border-black/10 bg-[color:var(--bg-muted)] p-4 sm:mb-12">
          <div className="relative w-full flex-1">
            <div className="absolute inset-0 overflow-hidden pr-2 pb-6">
              <PriceChart
                collection={{
                  max_supply: collection?.max_supply,
                  initial_price: collection?.initial_price,
                  current_supply: collection?.current_supply,
                  floor_price: collection?.floor_price,
                }}
                chainId={chainId}
              />
            </div>
          </div>
        </div>
      </div>

      <TransactionDialog
        isOpen={dialogState.isOpen}
        onOpenChange={onOpenChange}
        status={dialogState.status}
        hash={dialogState.hash}
        error={dialogState.error}
        title="Mint NFT success"
        chainId={chainId}
      />
    </div>
  );
}
