'use client'
import { useEffect, useState, useCallback, useMemo } from "react";
import useSWR from 'swr';
import {
  useReadContract,
  useAccount,
  useChainId,
  useWriteContract,
  useWaitForTransactionReceipt,
} from "wagmi";
import TransactionDialog from "@/components/shared/TransactionDialog";
import { useTransactionDialog } from "@/src/hooks/useTransactionDialog";
import { buildBulkSellTx, buildSellTx, buildSetApprovalForAllTx } from "@/src/onchain/tradeTx";
import { NFTItemCardSell } from "@/components/Collection/NFTItemCardSell";
import { Fluxus_EVM_CROSS_CHAIN_ABI, TRADE_ABI, getTradeContractAddress } from "@/src/contract";
import { SupportedChainId } from "@/src/alchemy";
import { Collection } from "@/src/api/types";
import { getNFTsByOwnerForCollection } from "@/src/api";
import { BatchSellPanel } from "@/components/Collection/BatchSellPanel"
import { getBatchSellPrice } from "@/lib/price";
import { getChainSymbol } from "@/src/utils";
import Loading from './ui/Loading';
import { batchFetchTokenImageUrlsByUri } from "@/src/media/tokenImage";

export default function Sell({ contractAddress, collection }: { contractAddress: string, collection?: Collection }) {
  const chainId = useChainId();
  const { address } = useAccount();
  
  // Contract interaction hooks
  const { data: hash, writeContract, isError: isWriteContractError, error: writeContractError, isPending } = useWriteContract();
  const { data: alchemyUserNFTs, error, isLoading, mutate: refetchUserTokens } = useSWR(
    address && collection?.address ? ["alchemy-nfts", chainId, address, collection.address] : null,
    () => getNFTsByOwnerForCollection(chainId as SupportedChainId, address as string, collection!.address, true, 100),
    { revalidateOnFocus: false }
  );
  if (error) {
    console.error('Error fetching alchemy user NFTs:', error);
  }
  
  // Component state
  const [soldTokens, setSoldTokens] = useState<number[]>([]);
  const [selectedTokens, setSelectedTokens] = useState<number[]>([]);
  const [batchCount, setBatchCount] = useState<number>(0);
  const [pendingApprovalTokenId, setPendingApprovalTokenId] = useState<number | null>(null);
  const [pendingBatchSell, setPendingBatchSell] = useState<boolean>(false);
  const [isBatchOperation, setIsBatchOperation] = useState<boolean>(false);
  const [hasInitialLoad, setHasInitialLoad] = useState(false);
  // Process user NFTs
  const userNFTs = useMemo(() => {
    if (!alchemyUserNFTs?.ownedNfts) return [];
    return alchemyUserNFTs.ownedNfts.map((nft: any) => {
      const tokenIdNum = Number(nft.tokenId);
      return {
        tokenId: tokenIdNum,
        name: collection?.symbol,
        tokenUri: `${collection?.base_uri}/${tokenIdNum}.json`,
      };
    }).filter((token: any) => !soldTokens.includes(token.tokenId));
  }, [alchemyUserNFTs, soldTokens, collection?.symbol, collection?.base_uri]);

  const maxSweep = userNFTs?.length || 0;
  
  // Transaction status
  const { isLoading: isConfirming, isSuccess: isConfirmed, isError, error: contractError } = useWaitForTransactionReceipt({ hash });
  const { dialogState, onOpenChange, setDialogState } = useTransactionDialog({
    hash,
    isPending,
    isConfirming,
    isConfirmed,
    isWriteContractError,
    writeContractError,
    onConfirmed: () => {
      if (pendingApprovalTokenId !== null) {
        refetchIsApproved().then(() => {
          executeSell(pendingApprovalTokenId);
          setPendingApprovalTokenId(null);
        });
      } else if (pendingBatchSell) {
        refetchIsApproved().then(() => {
          executeBatchSell(selectedTokens);
          setPendingBatchSell(false);
        });
      } else {
        setSoldTokens(prev => [...prev, ...selectedTokens]);
        setSelectedTokens([]);
        setIsBatchOperation(false);
      }
    }
  });
  
  // Check if user has approved the contract to spend NFTs
  const { data: isApproved, refetch: refetchIsApproved } = useReadContract({
    address: contractAddress as `0x${string}`,
    abi: Fluxus_EVM_CROSS_CHAIN_ABI,
    functionName: 'isApprovedForAll',
    args: [address, getTradeContractAddress(chainId as SupportedChainId)],
    query: { enabled: !!address },
  });

  // Read user's NFT balance
  const { data: balance, refetch: refetchBalance } = useReadContract({
    address: contractAddress as `0x${string}`,
    abi: Fluxus_EVM_CROSS_CHAIN_ABI,
    functionName: 'balanceOf',
    args: address ? [address] : undefined,
    query: { enabled: !!address },
  });

  // Refresh balance when contract or address changes
  useEffect(() => {
    refetchBalance();
  }, [contractAddress, address, refetchBalance]);

  // Refetch user tokens when balance changes
  useEffect(() => {
    if (balance !== undefined) {
      refetchUserTokens();
    }
  }, [balance, refetchUserTokens]);

  // Reset loading state when address or collection changes
  useEffect(() => {
    setHasInitialLoad(false);
    // Also manually trigger refetch when address changes
    if (address) {
      refetchBalance();
      refetchUserTokens();
    }
  }, [address, collection?.address, refetchBalance, refetchUserTokens]);

  // Handle initial data loading
  useEffect(() => {
    if (!address) return;
    if (!isLoading) {
      setHasInitialLoad(true);
    }
  }, [address, isLoading]);

  // Handle transaction errors surfaced by receipt hook
  useEffect(() => {
    if (isError) {
      handleTransactionError(contractError?.toString() || 'Unknown error');
    }
  }, [isError, contractError]);

  // Helper function to handle transaction errors
  const handleTransactionError = (errorMessage: string) => {
    setDialogState({ isOpen: true, status: 'error', error: errorMessage });
    setPendingApprovalTokenId(null);
    setPendingBatchSell(false);
    setIsBatchOperation(false);
  };

  // Execute sell operation
  const executeSell = (tokenId: number) => {
    setSelectedTokens([tokenId]);
    setIsBatchOperation(false);
    writeContract(
      buildSellTx({
        tradeContractAddress: getTradeContractAddress(chainId as SupportedChainId),
        collectionAddress: contractAddress,
        tokenId,
      }) as any
    );
  };

  // Execute approval operation
  const executeApproval = (tokenId: number) => {
    setPendingApprovalTokenId(tokenId);
    setIsBatchOperation(false);
    writeContract(
      buildSetApprovalForAllTx({
        nftContractAddress: contractAddress,
        operatorAddress: getTradeContractAddress(chainId as SupportedChainId),
        approved: true,
      }) as any
    );
  };

  // Main sell function
  const sellNFT = (tokenId: number) => {
    if (!isApproved) {
      executeApproval(tokenId);
    } else {
      executeSell(tokenId);
    }
  };



  // Handle manual token selection
  const handleSelect = (tokenId: number, selected: boolean) => {
    const newSelected = selected
      ? [...selectedTokens, tokenId]
      : selectedTokens.filter(id => id !== tokenId);
    setSelectedTokens(newSelected);
    setBatchCount(newSelected.length);
  };

  // Handle batch count change
  const handleCountChange = (count: number) => {
    const validCount = Math.max(0, Math.min(maxSweep, count));
    setBatchCount(validCount);
    const newSelected = userNFTs.slice(0, validCount).map(t => t.tokenId);
    setSelectedTokens(newSelected);
  };

  // Execute batch sell operation
  const executeBatchSell = (tokenIds: number[]) => {
    setIsBatchOperation(true);
    writeContract(
      buildBulkSellTx({
        tradeContractAddress: getTradeContractAddress(chainId as SupportedChainId),
        collectionAddress: contractAddress,
        tokenIds,
      }) as any
    );
  };

  // Execute batch approval operation
  const executeBatchApproval = () => {
    setPendingBatchSell(true);
    writeContract(
      buildSetApprovalForAllTx({
        nftContractAddress: contractAddress,
        operatorAddress: getTradeContractAddress(chainId as SupportedChainId),
        approved: true,
      }) as any
    );
  };

  // Main batch sell function
  const batchSellTokens = (tokenIds: number[]) => {
    if (!tokenIds.length) return;

    if (!isApproved) {
      executeBatchApproval();
    } else {
      executeBatchSell(tokenIds);
    }
  };

  // Calculate total profit from selected tokens
  const calculateTotalSellProfit = () => {
    if (!selectedTokens.length) return BigInt(0);
    return getBatchSellPrice(
      collection?.max_supply.toString() || '0',
      collection?.current_supply.toString() || '0', 
      collection?.initial_price.toString() || '0',
      selectedTokens.length,
      collection?.creator_fee || '0'
    );
  };

  // Image loading state for NFT previews
  const [imageLoadStatus, setImageLoadStatus] = useState<Record<string, boolean>>({});

  // Preload NFT images
  const preloadImages = useCallback((tokens: any[]) => {
    const uris = tokens.map((t: any) => t.tokenUri);
    batchFetchTokenImageUrlsByUri(uris, 8).then((images) => {
      const newLoadStatus: Record<string, boolean> = {};
      const tokenIdByUri: Record<string, number> = {};
      tokens.forEach((t: any) => { tokenIdByUri[t.tokenUri] = t.tokenId; });
      setImageLoadStatus(prev => {
        const merged = { ...prev };
        Object.entries(images).forEach(([uri, url]) => {
          const tokenId = tokenIdByUri[uri];
          newLoadStatus[tokenId] = !!url;
          merged[tokenId] = !!url;
        });
        return merged;
      });
      // Update token objects' image for cards
      tokens.forEach((t: any) => {
        const url = images[t.tokenUri];
        if (url) t.image = { originalUrl: url };
      });
    });
  }, []);

  // Preload images when NFTs change
  useEffect(() => {
    if (userNFTs?.length > 0) {
      preloadImages(userNFTs);
    }
  }, [userNFTs, preloadImages]);

  return (
    <div className="flex flex-col pb-32">
      {/* Header */}
      <div className="flex flex-col justify-between gap-2 border-b border-black/10 bg-[color:var(--bg-muted)] px-3 py-4 sm:flex-row sm:items-center sm:px-4">
        <div>
          <p className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">Holding NFTs</p>
          <p className="font-heading text-[26px] leading-none text-[color:var(--text-primary)]">{balance?.toString() || '0'}</p>
        </div>
        <p className="max-w-md text-sm text-[color:var(--text-secondary)]">
          Select items to sell individually or sweep multiple at once.
        </p>
      </div>
      
      {/* NFT Grid */}
      <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-4 gap-3 sm:gap-4 p-3 sm:p-4 min-h-[200px]">
        {!address ? (
          <div className="col-span-full flex flex-col justify-center items-center min-h-[200px] text-secondary">
            <div className="text-6xl mb-4">🔌</div>
            <div className="text-xl font-bold text-primary mb-2">Connect Wallet</div>
            <div className="text-sm text-secondary">Please connect your wallet to view your NFTs.</div>
          </div>
        ) : !hasInitialLoad ? (
          <div className="col-span-full flex justify-center items-center min-h-[200px]">
            <Loading />
          </div>
        ) : userNFTs.length > 0 ? (
          userNFTs.map((token: any) => (
            <NFTItemCardSell
              key={token.tokenId}
              token={token}
              collectionSymbol={collection?.symbol}
              contractAddress={contractAddress}
              isPending={isPending}
              onSell={sellNFT}
              onSelect={handleSelect}
              selected={selectedTokens.includes(token.tokenId)}
              isApproved={isApproved as boolean}
              pendingApprovalTokenId={pendingApprovalTokenId}
              isImageLoaded={imageLoadStatus[token.tokenId]}
            />
          ))
        ) : isLoading ? (
          <div className="col-span-full flex justify-center items-center min-h-[200px]">
            <Loading />
          </div>
        ) : (
          <div className="col-span-full flex flex-col justify-center items-center min-h-[200px] text-secondary">
            <div className="text-6xl mb-4">💼</div>
            <div className="text-xl font-bold text-primary mb-2">No NFTs to Sell</div>
            <div className="text-sm text-secondary">You don't own any NFTs in this collection.</div>
          </div>
        )}
      </div>
      
      {/* Transaction status dialog */}
      <TransactionDialog
        isOpen={dialogState.isOpen}
        onOpenChange={onOpenChange}
        status={dialogState.status}
        hash={dialogState.hash}
        error={dialogState.error}
        title={
          pendingApprovalTokenId !== null 
            ? "Approve NFT success" 
            : pendingBatchSell 
              ? "Approve NFTs success"
              : isBatchOperation
                ? "Sell NFTs success"
                : "Sell NFT success"
        }
        chainId={chainId}
      />
      {/* Batch sell panel */}
      <BatchSellPanel
        pricerSymbol={getChainSymbol(chainId)}
        maxSupply={collection?.max_supply.toString() || '0'}
        currentSupply={collection?.current_supply.toString() || '0'}
        initialPrice={collection?.initial_price.toString() || '0'}
        creatorFee={collection?.creator_fee || '0'}
        maxSweep={maxSweep}
        batchCount={batchCount}
        selectedTokens={selectedTokens}
        onCountChange={handleCountChange}
        onBatchSell={() => batchSellTokens(selectedTokens)}
        isApproved={isApproved as boolean}
        isPending={isPending}
        pendingBatchSell={pendingBatchSell}
      />
    </div>
  );
}
