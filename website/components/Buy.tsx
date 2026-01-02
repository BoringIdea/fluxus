'use client'
import React, { useEffect, useState, useRef, useCallback } from "react";
import {
  useReadContract,
  useWriteContract,
  useWaitForTransactionReceipt,
  useWatchContractEvent,
  useWatchBlockNumber,
  useChainId
} from "wagmi";
import { Button } from "./ui/button";
import { NFTItemCardBuy } from "@/components/Collection/NFTItemCardBuy";
import TransactionDialog from "@/components/shared/TransactionDialog";
import { getTradeContractAddress } from "@/src/contract";
import { Fluxus_EVM_CROSS_CHAIN_ABI, PRICE_ABI, TRADE_ABI } from "@/src/contract";
import { Collection } from "@/src/api/types";
import { SupportedChainId } from "@/src/alchemy";
import { BatchBuyPanel } from "@/components/Collection/BatchBuyPanel"
import { getBatchBuyPrice } from "@/lib/price";
import { getChainSymbol, formatNumberWithMaxDecimalsAndRounding } from "@/src/utils";
import Loading from './ui/Loading';
import { useTransactionDialog } from "@/src/hooks/useTransactionDialog";
import { buildBuyTx, buildBulkBuyTx, buildBulkQuickBuyTx, buildQuickBuyTx } from "@/src/onchain/tradeTx";
import { batchFetchTokenImageUrlsByUri } from "@/src/media/tokenImage";

interface BuyProps {
  contractAddress: string;
  collection?: Collection;
}

export default function Buy({ contractAddress, collection }: BuyProps) {
  const chainId = useChainId();
  const tradeContractAddress = getTradeContractAddress(chainId as SupportedChainId);

  // State for NFTs
  const [listedTokens, setListedTokens] = useState<any[]>([]);
  const [displayedTokens, setDisplayedTokens] = useState<any[]>([]);
  const [displayCount, setDisplayCount] = useState(20);
  const [boughtTokens, setBoughtTokens] = useState<any[]>([]);
  const [hasInitialLoad, setHasInitialLoad] = useState(false);
  const [loadingTimeout, setLoadingTimeout] = useState(false);
  const [isProcessingMetadata, setIsProcessingMetadata] = useState(false);
  // Infinite scroll ref
  const observer = useRef<IntersectionObserver | null>(null);
  const lastTokenRef = useCallback((node: HTMLDivElement) => {
    if (observer.current) observer.current.disconnect();
    observer.current = new IntersectionObserver(entries => {
      if (entries[0].isIntersecting && displayedTokens.length < listedTokens.length) {
        setDisplayCount(c => c + 20);
      }
    });
    if (node) observer.current.observe(node);
  }, [displayedTokens.length, listedTokens.length]);

  // Contract interaction hooks
  const { data: hash, writeContract, isError: isWriteContractError, error: writeContractError, isPending } = useWriteContract();

  const { data: availableTokens, refetch: refetchAvailableTokens, isLoading: isLoadingAvailableTokens } = useReadContract({
    address: contractAddress as any,
    abi: Fluxus_EVM_CROSS_CHAIN_ABI,
    functionName: 'getAvailableTokensPaginated',
    args: [0, 100],
    scopeKey: `available:${contractAddress}`,
    query: {
      staleTime: 30000,
      gcTime: 300000,
      refetchOnWindowFocus: false,
      retry: 1,
    }
  });

  const { data: buyPrice, refetch: refetchBuyPrice } = useReadContract({
    address: collection?.price_contract as any,
    abi: PRICE_ABI,
    functionName: 'getBuyPriceAfterFee',
    args: [collection?.address],
    scopeKey: `buyPrice:${collection?.address}`,
    query: {
      enabled: !!collection?.price_contract && !!collection?.address,
      staleTime: 30000,
      gcTime: 300000,
      refetchOnWindowFocus: false,
      retry: 1,
    }
  });

  // Refetch on new blocks (throttled) for better timeliness than fixed interval
  const lastRefetchTsRef = useRef(0);
  useWatchBlockNumber({
    onBlockNumber: () => {
      const now = Date.now();
      if (now - lastRefetchTsRef.current < 1000) return;
      refetchAvailableTokens();
      refetchBuyPrice();
      lastRefetchTsRef.current = now;
    }
  });
  // Watch for TokenBought events to refresh data
  useWatchContractEvent({
    address: contractAddress as any,
    abi: Fluxus_EVM_CROSS_CHAIN_ABI,
    eventName: 'TokenBought',
    onLogs: () => {
      refetchBuyPrice();
      refetchAvailableTokens();
    },
  });

  // Refresh data when contract or collection changes
  useEffect(() => {
    refetchBuyPrice();
    refetchAvailableTokens();
    setHasInitialLoad(false);
    setLoadingTimeout(false);
    setIsProcessingMetadata(false);
  }, [contractAddress, collection?.address, collection?.price_contract]);

  // Keep a gentle polling (2s) alongside event refresh to improve timeliness
  useEffect(() => {
    if (!contractAddress) return;
    const intervalId = setInterval(() => {
      if (typeof document !== 'undefined' && document.hidden) return;
      refetchAvailableTokens();
    }, 2000);
    return () => clearInterval(intervalId);
  }, [contractAddress, refetchAvailableTokens]);

  // Soften loading timeout and avoid switching to timeout too early
  useEffect(() => {
    if (isLoadingAvailableTokens && !hasInitialLoad) {
      const timeoutId = setTimeout(() => {
        setLoadingTimeout(true);
        setHasInitialLoad(true);
      }, 8000);
      return () => clearTimeout(timeoutId);
    }
  }, [isLoadingAvailableTokens, hasInitialLoad]);

  const { isLoading: isConfirming, isSuccess: isConfirmed } = useWaitForTransactionReceipt({ hash });

  // Metadata cache for NFT data - use ref to avoid re-renders
  const metadataCacheRef = useRef<Record<string, any>>({});
  const [metadataCache, setMetadataCache] = useState<Record<string, any>>({});

  // Clean cache when collection changes
  useEffect(() => {
    if (collection?.base_uri) {
      const newCache: Record<string, any> = {};
      // Keep only current collection cache
      Object.entries(metadataCacheRef.current).forEach(([key, value]) => {
        if (key.startsWith(collection.base_uri)) {
          newCache[key] = value;
        }
      });
      metadataCacheRef.current = newCache;
      setMetadataCache(newCache);
    }
  }, [collection?.base_uri]);

  // Cache size limit and update function
  const MAX_CACHE_SIZE = 1000;
  const updateCache = useCallback((key: string, value: any) => {
    metadataCacheRef.current = {
      ...metadataCacheRef.current,
      [key]: value
    };
    
    // Remove oldest entry if cache exceeds size limit
    if (Object.keys(metadataCacheRef.current).length > MAX_CACHE_SIZE) {
      const oldestKey = Object.keys(metadataCacheRef.current)[0];
      delete metadataCacheRef.current[oldestKey];
    }
    
    setMetadataCache(metadataCacheRef.current);
  }, []);

  // Build initial token list fast and lazy-load metadata for visible range
  useEffect(() => {
    if (!availableTokens || !Array.isArray(availableTokens)) {
      setListedTokens([]);
      setHasInitialLoad(true);
      setIsProcessingMetadata(false);
      return;
    }

    // Initialize with placeholders for fast first paint
    const initial = availableTokens.map((tokenId: any) => ({
      tokenId: Number(tokenId),
      image: { originalUrl: '/fluxus-logo.png' },
    }));
    setListedTokens(initial);
    setDisplayedTokens(initial.slice(0, displayCount));
    setHasInitialLoad(true);

    // Prefetch metadata for the first visible batch
    const limit = displayCount;
    const visibleIds = initial.slice(0, limit).map(t => t.tokenId);
    prefetchMetadata(visibleIds);
  }, [availableTokens]);

  // Prefetch metadata for newly visible tokens when displayCount increases
  useEffect(() => {
    if (!listedTokens.length) return;
    const needIds = displayedTokens
      .filter(t => t && t.image && t.image.originalUrl === '/fluxus-logo.png')
      .map(t => t.tokenId);
    if (needIds.length) prefetchMetadata(needIds);
  }, [displayedTokens, listedTokens.length]);

  // Concurrency-limited metadata prefetch (extracted util)
  const prefetchMetadata = async (tokenIds: number[], concurrency = 8) => {
    if (!collection?.base_uri) return;
    setIsProcessingMetadata(true);
    const uris = tokenIds.map(id => `${collection.base_uri}/${id}.json`);
    const images = await batchFetchTokenImageUrlsByUri(uris, concurrency);
    // Update both listedTokens and displayedTokens for quicker visual update
    setListedTokens(prev => prev.map(t => {
      const uri = `${collection.base_uri}/${t.tokenId}.json`;
      const url = images[uri];
      if (url) return { ...t, image: { originalUrl: url }, needsRetry: false };
      return t;
    }));
    setDisplayedTokens(prev => prev.map(t => {
      const uri = `${collection.base_uri}/${t.tokenId}.json`;
      const url = images[uri];
      if (url) return { ...t, image: { originalUrl: url }, needsRetry: false };
      return t;
    }));
    setIsProcessingMetadata(false);
  };

  // Retry mechanism for failed metadata fetches
  useEffect(() => {
    const retryFailedTokens = async () => {
      const failedTokens = listedTokens.filter(token => token.needsRetry);
      if (failedTokens.length === 0) return;

      const retryResults = await Promise.all(
        failedTokens.map(async (token) => {
          const cacheKey = `${collection?.base_uri}/${token.tokenId}.json`;
          try {
            const res = await fetch(cacheKey);
            if (!res.ok) {
              throw new Error(`Failed to fetch metadata: ${res.status}`);
            }
            
            const md = await res.json();
            
            if (!md.image) {
              throw new Error('No image URL in metadata');
            }

            updateCache(cacheKey, md);
            
            return {
              tokenId: token.tokenId,
              image: { originalUrl: md.image }
            };
          } catch (error) {
            console.warn(`Retry failed for token ${token.tokenId}:`, error);
            return {
              ...token,
              retryCount: (token.retryCount || 0) + 1,
              lastError: error instanceof Error ? error.message : 'Unknown error'
            };
          }
        })
      );

      setListedTokens(prev => 
        prev.map(token => {
          const retryResult = retryResults.find(r => r.tokenId === token.tokenId);
          return retryResult || token;
        })
      );
    };

    // Only retry if we have listedTokens and they contain failed tokens
    if (listedTokens.length > 0 && listedTokens.some(token => token.needsRetry)) {
      // Retry after 500ms delay
      const timeoutId = setTimeout(() => {
        retryFailedTokens();
      }, 500);

      return () => clearTimeout(timeoutId);
    }
  }, [collection?.base_uri]);

  useEffect(() => {
    setDisplayedTokens(listedTokens.slice(0, displayCount));
  }, [listedTokens, displayCount]);

  // Selection and batch state
  const [selectedTokens, setSelectedTokens] = useState<number[]>([]);
  const [batchCount, setBatchCount] = useState<number>(0);
  const [quickBuyCount, setQuickBuyCount] = useState<number | ''>(1);

  const maxSweep = displayedTokens.length;

  const { dialogState, onOpenChange } = useTransactionDialog({
    hash,
    isPending,
    isConfirming,
    isConfirmed,
    isWriteContractError,
    writeContractError,
    onConfirmed: () => {
      setDisplayedTokens(prev => prev.filter(token => !boughtTokens.includes(token.tokenId)));
      setListedTokens(prev => prev.filter(token => !boughtTokens.includes(token.tokenId)));
    }
  });

  // Handle manual token selection
  const handleSelect = (tokenId: number, selected: boolean) => {
    const newSelected = selected
      ? [...selectedTokens, tokenId]
      : selectedTokens.filter(id => id !== tokenId);
    setSelectedTokens(newSelected);
    setBatchCount(newSelected.length);
  };

  // Handle batch count change from input/slider
  const handleCountChange = (count: number) => {
    const validCount = Math.max(0, Math.min(maxSweep, count));
    setBatchCount(validCount);
    const newSelected = displayedTokens.slice(0, validCount).map(t => t.tokenId);
    setSelectedTokens(newSelected);
  };

  // Buy single token
  async function buyToken(tokenId: number) {
    if (!buyPrice) return;
    await writeContract(
      buildBuyTx({
        tradeContractAddress,
        collectionAddress: contractAddress,
        tokenId,
        buyPriceWei: BigInt(buyPrice.toString()),
      }) as any
    );
    setBoughtTokens([tokenId]);
  }
  // Quick buy function
  async function quickBuyToken() {
    if (!buyPrice) return;
    
    if (quickBuyCount === 1) {
      await writeContract(
        buildQuickBuyTx({
          tradeContractAddress,
          collectionAddress: contractAddress,
          buyPriceWei: BigInt(buyPrice.toString()),
          // include tokenId to satisfy structural typing in case of widened param inference
          tokenId: 0,
        } as any) as any
      );
    } else if (quickBuyCount && quickBuyCount > 1) {
      const totalCost = getBatchBuyPrice(
        collection?.max_supply.toString() || '0',
        collection?.current_supply.toString() || '0',
        collection?.initial_price.toString() || '0',
        quickBuyCount,
        collection?.creator_fee || '0'
      );
      await writeContract(
        buildBulkQuickBuyTx({
          tradeContractAddress,
          collectionAddress: contractAddress,
          amount: quickBuyCount,
          totalCostWei: totalCost,
        }) as any
      );
    }
  }
  // Calculate total cost for batch purchase
  const calculateTotalCost = () => {
    if (!buyPrice || batchCount === 0) return BigInt(0);
    return getBatchBuyPrice(
      collection?.max_supply.toString() || '0',
      collection?.current_supply.toString() || '0',
      collection?.initial_price.toString() || '0',
      batchCount,
      collection?.creator_fee || '0'
    );
  };

  // Batch buy selected tokens
  async function batchBuyTokens(tokenIds: number[]) {
    if (!tokenIds.length || !buyPrice) return;
    
    await writeContract(
      buildBulkBuyTx({
        tradeContractAddress,
        collectionAddress: contractAddress,
        tokenIds,
        totalCostWei: calculateTotalCost(),
      }) as any
    );
    setBoughtTokens(tokenIds);
  }

  return (
    <div className="flex flex-col pb-32">
      {/* Header with quick buy controls */}
      <div className="flex flex-col sm:flex-row justify-between items-stretch sm:items-center gap-3 border-b border-border px-3 sm:px-4 py-4 bg-bg-card">
        <div>
          <p className="text-[11px] font-bold uppercase tracking-[0.3em] text-secondary">Listed NFTs</p>
          <p className="text-lg font-black text-primary">{listedTokens.length}</p>
        </div>
        <div className="flex flex-col sm:flex-row gap-2 w-full sm:w-auto sm:ml-auto text-xs">
          <div className="flex items-center border border-border bg-bg-tertiary text-primary w-full sm:w-auto">
            <button
              type="button"
              className="h-8 w-8 border-r border-border font-black hover:bg-bg-card-hover transition-colors"
              onClick={() => setQuickBuyCount(prev => {
                if (typeof prev === 'number') return Math.max(1, prev - 1);
                return 1;
              })}
            >
              -
            </button>
            <input
              type="number"
              min={1}
              max={20}
              value={quickBuyCount}
              onChange={e => { 
                const value = e.target.value ? Number(e.target.value) : '';
                if (typeof value === 'number' && value > 20) {
                  setQuickBuyCount(20);
                } else {
                  setQuickBuyCount(value);
                }
              }}
              className="w-full sm:w-14 bg-transparent text-primary text-center font-semibold focus:outline-none focus:ring-0 text-sm border-0 [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
            />
            <button
              type="button"
              className="h-8 w-8 border-l border-border font-black hover:bg-bg-card-hover transition-colors"
              onClick={() => setQuickBuyCount(prev => {
                if (typeof prev === 'number') return Math.min(Math.min(20, maxSweep), prev + 1);
                return 1;
              })}
            >
              +
            </button>
          </div>
          <Button
            className="w-full sm:w-auto text-xs tracking-[0.2em] px-3 py-2"
            onClick={quickBuyToken}
          >
            Quick buy ({formatNumberWithMaxDecimalsAndRounding(Number(getBatchBuyPrice(
              collection?.max_supply.toString() || '0',
              collection?.current_supply.toString() || '0',
              collection?.initial_price.toString() || '0',
              Number(quickBuyCount),
              collection?.creator_fee || '0'
            )) / 1e18, 2)} {getChainSymbol(chainId)})
          </Button>
        </div>
      </div>
      {/* NFT Grid */}
      <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-4 gap-3 sm:gap-4 p-3 sm:p-4 min-h-[200px]">
        {(isLoadingAvailableTokens || !hasInitialLoad) && !loadingTimeout ? (
          <div className="col-span-full flex justify-center items-center min-h-[200px]">
            <Loading />
          </div>
        ) : loadingTimeout ? (
          <div className="col-span-full flex flex-col justify-center items-center min-h-[200px] text-secondary">
            <div className="text-6xl mb-4">⏰</div>
            <div className="text-xl font-bold mb-2 text-primary">Loading Timeout</div>
            <div className="text-sm text-secondary text-center mb-4">
              Unable to load NFTs. Please check your connection and try again.
            </div>
            <Button
              onClick={() => {
                setLoadingTimeout(false);
                setHasInitialLoad(false);
                setIsProcessingMetadata(false);
                refetchAvailableTokens();
              }}
            >
              Retry
            </Button>
          </div>
        ) : displayedTokens.length === 0 ? (
          <div className="col-span-full flex flex-col justify-center items-center min-h-[200px] text-secondary">
            <div className="text-6xl mb-4">🛒</div>
            <div className="text-xl font-bold mb-2 text-primary">No NFTs Available</div>
            <div className="text-sm text-secondary">All NFTs have been sold or none are listed yet.</div>
          </div>
        ) : (
          displayedTokens.map((token, idx) => (
            <NFTItemCardBuy
              key={token.tokenId}
              token={token}
              isPending={isPending}
              collectionSymbol={collection?.symbol}
              contractAddress={contractAddress}
              onBuy={buyToken}
              onSelect={handleSelect}
              selected={selectedTokens.includes(token.tokenId)}
              ref={idx === displayedTokens.length - 1 ? lastTokenRef : undefined}
            />
          ))
        )}
      </div>
      {/* Transaction status dialog */}
      <TransactionDialog
        isOpen={dialogState.isOpen}
        onOpenChange={onOpenChange}
        status={dialogState.status}
        hash={dialogState.hash}
        error={dialogState.error}
        title="Buy NFT success"
        chainId={chainId}
      />
      {/* Batch buy panel */}
      <BatchBuyPanel
        priceSymbol={getChainSymbol(chainId)}
        maxSupply={collection?.max_supply.toString() || '0'}
        currentSupply={collection?.current_supply.toString() || '0'}
        initialPrice={collection?.initial_price.toString() || '0'}
        creatorFee={collection?.creator_fee || '0'}
        maxSweep={maxSweep}
        batchCount={batchCount}
        selectedTokens={selectedTokens}
        onCountChange={handleCountChange}
        onBatchBuy={() => batchBuyTokens(selectedTokens)}
      />
    </div>
  )
}
