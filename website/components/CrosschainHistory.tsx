'use client'
import React, { useState, useMemo, useEffect } from "react";
import { useAccount, useChainId } from "wagmi";
import Loading from './ui/Loading';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { sliceAddress } from "@/src/utils";
import Pagination from '@/components/ui/pagination';
import { useCrossChainStatus } from "@/src/api/hooks";
import { defaultCCTXTracker, CCTXRecord } from '@/lib/cctxTracker';
import { getChainName, mapCCTXStatus } from '@/lib/chains';

// Simple module-level cache to persist processed history between navigations
type HistoryCacheEntry = { items: CrosschainHistoryItem[]; hashes: Set<string>; timestamp: number };
const historyCache = new Map<string, HistoryCacheEntry>();

interface CrosschainHistoryItem {
  id: string;
  tokenId: string;
  sourceChain: string;
  targetChain: string;
  status: 'pending' | 'completed' | 'failed';
  transactionHash: string;
  blockTimestamp: number;
  userAddress: string;
  sender?: string;
  receiver?: string;
  destination?: string;
}

// Mock data for demonstration
const mockHistoryData: CrosschainHistoryItem[] = [
  {
    id: '1',
    tokenId: '123',
    sourceChain: 'Base Sepolia',
    targetChain: 'BSC Testnet',
    status: 'completed',
    transactionHash: '0x1234567890abcdef1234567890abcdef12345678',
    blockTimestamp: Date.now() - 3600000, // 1 hour ago
    userAddress: '0xabcdef1234567890abcdef1234567890abcdef12',
    sender: '0xabcdef1234567890abcdef1234567890abcdef12',
    receiver: '0x9876543210fedcba9876543210fedcba98765432'
  },
  {
    id: '2',
    tokenId: '456',
    sourceChain: 'BSC Testnet',
    targetChain: 'Base Sepolia',
    status: 'pending',
    transactionHash: '0x9876543210fedcba9876543210fedcba98765432',
    blockTimestamp: Date.now() - 1800000, // 30 minutes ago
    userAddress: '0xabcdef1234567890abcdef1234567890abcdef12',
    sender: '0x9876543210fedcba9876543210fedcba98765432',
    receiver: '0xabcdef1234567890abcdef1234567890abcdef12'
  },
  {
    id: '3',
    tokenId: '789',
    sourceChain: 'Base Sepolia',
    targetChain: 'BSC Testnet',
    status: 'failed',
    transactionHash: '0x5555555555555555555555555555555555555555',
    blockTimestamp: Date.now() - 7200000, // 2 hours ago
    userAddress: '0xabcdef1234567890abcdef1234567890abcdef12',
    sender: '0xabcdef1234567890abcdef1234567890abcdef12',
    receiver: '0x1111222233334444555566667777888899990000'
  },
  {
    id: '4',
    tokenId: '234',
    sourceChain: 'BSC Testnet',
    targetChain: 'Base Sepolia',
    status: 'completed',
    transactionHash: '0xa1b2c3d4e5f6789012345678901234567890abcd',
    blockTimestamp: Date.now() - 10800000, // 3 hours ago
    userAddress: '0xabcdef1234567890abcdef1234567890abcdef12'
  },
  {
    id: '5',
    tokenId: '567',
    sourceChain: 'Base Sepolia',
    targetChain: 'BSC Testnet',
    status: 'completed',
    transactionHash: '0xfedcba0987654321fedcba0987654321fedcba09',
    blockTimestamp: Date.now() - 14400000, // 4 hours ago
    userAddress: '0xabcdef1234567890abcdef1234567890abcdef12'
  },
  {
    id: '6',
    tokenId: '890',
    sourceChain: 'BSC Testnet',
    targetChain: 'Base Sepolia',
    status: 'pending',
    transactionHash: '0x1111222233334444555566667777888899990000',
    blockTimestamp: Date.now() - 900000, // 15 minutes ago
    userAddress: '0xabcdef1234567890abcdef1234567890abcdef12'
  },
  {
    id: '7',
    tokenId: '345',
    sourceChain: 'Base Sepolia',
    targetChain: 'BSC Testnet',
    status: 'completed',
    transactionHash: '0xabcdef1234567890abcdef1234567890abcdef12',
    blockTimestamp: Date.now() - 18000000, // 5 hours ago
    userAddress: '0xabcdef1234567890abcdef1234567890abcdef12'
  },
  {
    id: '8',
    tokenId: '678',
    sourceChain: 'BSC Testnet',
    targetChain: 'Base Sepolia',
    status: 'failed',
    transactionHash: '0x9999888877776666555544443333222211110000',
    blockTimestamp: Date.now() - 21600000, // 6 hours ago
    userAddress: '0xabcdef1234567890abcdef1234567890abcdef12'
  },
  {
    id: '9',
    tokenId: '901',
    sourceChain: 'Base Sepolia',
    targetChain: 'BSC Testnet',
    status: 'completed',
    transactionHash: '0x77776666555544443333222211110000abcdef12',
    blockTimestamp: Date.now() - 25200000, // 7 hours ago
    userAddress: '0xabcdef1234567890abcdef1234567890abcdef12'
  },
  {
    id: '10',
    tokenId: '112',
    sourceChain: 'BSC Testnet',
    targetChain: 'Base Sepolia',
    status: 'completed',
    transactionHash: '0x33334444555566667777888899990000abcdef12',
    blockTimestamp: Date.now() - 28800000, // 8 hours ago
    userAddress: '0xabcdef1234567890abcdef1234567890abcdef12'
  }
];

function getStatusDisplay(status: string) {
  switch (status) {
    case 'completed':
      return { text: "Completed", color: "text-green-400", bgColor: "bg-green-400/10", icon: "✅" };
    case 'pending':
      return { text: "Pending", color: "text-yellow-400", bgColor: "bg-yellow-400/10", icon: "⏳" };
    case 'failed':
      return { text: "Failed", color: "text-red-400", bgColor: "bg-red-400/10", icon: "❌" };
    default:
      return { text: "Unknown", color: "text-gray-400", bgColor: "bg-gray-400/10", icon: "❓" };
  }
}

function formatTime(timestamp: number) {
  const date = new Date(timestamp);
  const now = new Date();
  const diff = (now.getTime() - date.getTime()) / 1000;
  
  if (diff < 60) return `${Math.floor(diff)}s ago`;
  if (diff < 3600) return `${Math.floor(diff / 60)}m ago`;
  if (diff < 86400) return `${Math.floor(diff / 3600)}h ago`;
  return date.toLocaleDateString();
}

const PAGE_SIZE = 5;

const mockTxHash = '0xf11ffae73c222fd259bf5cf8518eb4a14a827bb898b48ae7046d5b67a679cdb5';

export default function CrosschainHistory({ 
  baseContractAddress, 
  bscContractAddress 
}: { 
  baseContractAddress: string;
  bscContractAddress: string;
}) {
  const { address } = useAccount();
  const [isLoading, setIsLoading] = useState(false);
  const [page, setPage] = useState(1);
  const cacheKey = `${baseContractAddress}|${bscContractAddress}`;
  const cached = historyCache.get(cacheKey);
  const [processedTransactions, setProcessedTransactions] = useState<CrosschainHistoryItem[]>(cached?.items || []);
  const [processedTxHashes, setProcessedTxHashes] = useState<Set<string>>(cached?.hashes ? new Set(Array.from(cached.hashes)) : new Set());
  const [isProcessing, setIsProcessing] = useState(false);
  const [hasInitialLoaded, setHasInitialLoaded] = useState<boolean>(false);
  const chainId = useChainId();
  const getCachedHashesArray = () => Array.from(historyCache.get(cacheKey)?.hashes ?? []);
  
  // Get cross-chain status from Base Sepolia (84532)
  const { data: baseCrossChainStatus, isLoading: isLoadingBaseCrossChainStatus, isValidating: isValidatingBaseCrossChainStatus, mutate: mutateBaseCrossChainStatus } = useCrossChainStatus(84532, baseContractAddress);
  
  // Get cross-chain status from BSC Testnet (97)
  const { data: bscCrossChainStatus, isLoading: isLoadingBscCrossChainStatus, isValidating: isValidatingBscCrossChainStatus, mutate: mutateBscCrossChainStatus } = useCrossChainStatus(97, bscContractAddress);
  const [isRefreshing, setIsRefreshing] = useState(false);

  // Poll every 5s to fetch latest data (explicit mutate to override hook's default 15s)
  useEffect(() => {
    if (!baseContractAddress && !bscContractAddress) return;
    let isMounted = true;
    const interval = setInterval(async () => {
      try {
        if (!isMounted) return;
        setIsRefreshing(true);
        await Promise.all([
          mutateBaseCrossChainStatus?.(),
          mutateBscCrossChainStatus?.()
        ]);
      } finally {
        if (isMounted) setIsRefreshing(false);
      }
    }, 5000);
    return () => {
      isMounted = false;
      clearInterval(interval);
    };
  }, [baseContractAddress, bscContractAddress, mutateBaseCrossChainStatus, mutateBscCrossChainStatus]);

  // Immediately fetch newest data on mount and when contract addresses change
  useEffect(() => {
    let isMounted = true;
    const run = async () => {
      if (!baseContractAddress && !bscContractAddress) return;
      try {
        setIsRefreshing(true);
        await Promise.all([
          mutateBaseCrossChainStatus?.(),
          mutateBscCrossChainStatus?.()
        ]);
      } finally {
        if (isMounted) setIsRefreshing(false);
      }
    };
    run();
    return () => {
      isMounted = false;
    };
  }, [baseContractAddress, bscContractAddress, mutateBaseCrossChainStatus, mutateBscCrossChainStatus]);
  
  // Copy to clipboard helper with feedback
  const copyToClipboard = async (text: string, label: string) => {
    try {
      await navigator.clipboard.writeText(text);
      console.log(`📋 Copied ${label}: ${text}`);
      // You could add toast notification here if desired
    } catch (error) {
      console.error(`❌ Failed to copy ${label}:`, error);
    }
  };

  // Process cross-chain transactions with CCTX tracking
  const processCrossChainData = async (transactions: any[]) => {
    if (!transactions || transactions.length === 0) return;

    setIsProcessing(true);
    console.log(`🔄 Processing ${transactions.length} transactions for CCTX data`);

    try {
      for (const tx of transactions) {
        const txHash = tx.transactionHash;
        if (!txHash || processedTxHashes.has(txHash)) {
          continue; // Skip already processed transactions
        }

        console.log(`📡 Processing transaction: ${txHash}`);

        try {
          const cctxRecords = await defaultCCTXTracker.trackCCTX(txHash, {
            onCCTXUpdate: (cctxRecord: CCTXRecord) => {
              console.log(`📋 CCTX Update received for ${txHash}:`, cctxRecord);
            },
            onTrackingStart: () => {
              console.log(`📡 Started tracking: ${txHash}`);
            },
            onTrackingComplete: () => {
              console.log(`✅ Completed tracking: ${txHash}`);
            },
            onError: (error) => {
              console.error(`❌ CCTX tracking failed for ${txHash}:`, error);
            }
          });

          console.log(`🔍 CCTX Records received for ${txHash}:`, cctxRecords);

          // Process the returned CCTX records
          if (cctxRecords && cctxRecords.length > 0) {
            console.log(`🔗 Processing ${cctxRecords.length} CCTX records for ${txHash}`);
            
            // Analyze all CCTX records to understand the full cross-chain flow
            const inboundRecord = cctxRecords.find(record => record.sender_chain_id !== '7001'); // Not from ZetaChain
            const outboundRecord = cctxRecords.find(record => record.sender_chain_id === '7001'); // From ZetaChain
            
            console.log(`📥 Inbound record (to ZetaChain):`, inboundRecord);
            console.log(`📤 Outbound record (from ZetaChain):`, outboundRecord);
            
            // Determine the overall transaction status
            let overallStatus: 'pending' | 'completed' | 'failed' = 'pending';
            let sourceChain = 'Unknown Chain';
            let targetChain = 'Unknown Chain';
            
            if (inboundRecord && outboundRecord) {
              // Both legs exist - check final status based on outbound transaction
              sourceChain = getChainName(inboundRecord.sender_chain_id) || 'Unknown Chain';
              targetChain = getChainName(outboundRecord.receiver_chainId) || 'Unknown Chain';
              
              // The overall status is determined by the outbound transaction status
              overallStatus = mapCCTXStatus(outboundRecord.status);
              
              console.log(`🔗 Full cross-chain flow: ${sourceChain} → ZetaChain → ${targetChain}, Status: ${overallStatus}`);
              console.log(`📊 Inbound status: ${inboundRecord.status}, Outbound status: ${outboundRecord.status}`);
            } else if (inboundRecord) {
              // Only inbound leg - use inbound status to determine overall status
              sourceChain = getChainName(inboundRecord.sender_chain_id) || 'Unknown Chain';
              targetChain = getChainName(inboundRecord.receiver_chainId) || 'ZetaChain';
              overallStatus = mapCCTXStatus(inboundRecord.status);
              
              console.log(`📥 Inbound only: ${sourceChain} → ${targetChain}, Status: ${overallStatus}`);
              console.log(`📊 Inbound status: ${inboundRecord.status}`);
            } else {
              // Use the first available record
              const firstRecord = cctxRecords[0];
              sourceChain = getChainName(firstRecord.sender_chain_id) || 'Unknown Chain';
              targetChain = getChainName(firstRecord.receiver_chainId) || 'Unknown Chain';
              overallStatus = mapCCTXStatus(firstRecord.status);
              
              console.log(`📄 Single record: ${sourceChain} → ${targetChain}, Status: ${overallStatus}`);
            }
            
            const historyItem: CrosschainHistoryItem = {
              id: tx.id || txHash,
              tokenId: tx.tokenId?.toString() || 'Unknown',
              sourceChain,
              targetChain,
              status: overallStatus,
              transactionHash: txHash,
              blockTimestamp: tx.blockTimestamp * 1000,
              userAddress: tx.sender || address || '',
              sender: tx.sender,
              receiver: tx.receiver,
              destination: tx.destination
            };

            console.log(`✨ Created comprehensive history item for ${txHash}:`, historyItem);

            setProcessedTransactions(prev => {
              const exists = prev.find(item => item.id === historyItem.id);
              if (!exists) {
                const newList = [historyItem, ...prev].sort((a, b) => b.blockTimestamp - a.blockTimestamp);
                console.log(`📝 Added new item to processed transactions. Total: ${newList.length}`);
                // update cache
                historyCache.set(cacheKey, { items: newList, hashes: new Set<string>([...getCachedHashesArray(), txHash]), timestamp: Date.now() });
                return newList;
              }
              const updated = prev.map(item => item.id === historyItem.id ? historyItem : item);
              historyCache.set(cacheKey, { items: updated, hashes: new Set<string>([...getCachedHashesArray(), txHash]), timestamp: Date.now() });
              return updated;
            });
          } else {
            console.log(`⚠️ No CCTX records found for ${txHash}, using basic data`);
            // Use basic transaction data if no CCTX data is available
            const basicItem: CrosschainHistoryItem = {
              id: tx.id || txHash,
              tokenId: tx.tokenId?.toString() || 'Unknown',
              sourceChain: 'Unknown Chain',
              targetChain: tx.destination || 'Unknown Chain',
              status: tx.isTransfered ? 'completed' : 'pending',
              transactionHash: txHash,
              blockTimestamp: tx.blockTimestamp * 1000,
              userAddress: tx.sender || address || '',
              sender: tx.sender,
              receiver: tx.receiver,
              destination: tx.destination
            };
            
            setProcessedTransactions(prev => {
              const exists = prev.find(item => item.id === basicItem.id);
              if (!exists) {
                const newList = [basicItem, ...prev].sort((a, b) => b.blockTimestamp - a.blockTimestamp);
                console.log(`📝 Added basic item to processed transactions. Total: ${newList.length}`);
                historyCache.set(cacheKey, { items: newList, hashes: new Set<string>([...getCachedHashesArray(), txHash]), timestamp: Date.now() });
                return newList;
              }
              return prev;
            });
          }
          
          setProcessedTxHashes(prev => {
            const next = new Set<string>([...Array.from(prev), txHash]);
            historyCache.set(cacheKey, { items: historyCache.get(cacheKey)?.items || processedTransactions, hashes: next, timestamp: Date.now() });
            return next;
          });
        } catch (error) {
          console.error(`❌ Error processing ${txHash}:`, error);
          setProcessedTxHashes(prev => {
            const next = new Set<string>([...Array.from(prev), txHash]);
            historyCache.set(cacheKey, { items: historyCache.get(cacheKey)?.items || processedTransactions, hashes: next, timestamp: Date.now() });
            return next;
          });
        }
      }
    } finally {
      setIsProcessing(false);
    }
  };



  // Process cross-chain data when it's available from either chain
  useEffect(() => {
    const allTransactions = [
      ...(baseCrossChainStatus?.data || []),
      ...(bscCrossChainStatus?.data || [])
    ];
    
    if (allTransactions.length > 0) {
      const unprocessedTransactions = allTransactions.filter(
        tx => (tx?.transactionHash) && !processedTxHashes.has(tx.transactionHash)
      );
      
      if (unprocessedTransactions.length > 0 && !isProcessing) {
        console.log(`🔄 Processing ${unprocessedTransactions.length} transactions from both chains:`, {
          baseChainTxs: baseCrossChainStatus?.data?.length || 0,
          bscChainTxs: bscCrossChainStatus?.data?.length || 0,
          totalUnprocessed: unprocessedTransactions.length,
          baseContractAddress,
          bscContractAddress,
          contractAddressesValid: {
            base: !!baseContractAddress,
            bsc: !!bscContractAddress
          }
        });
        processCrossChainData(unprocessedTransactions);
      }
    }
  }, [baseCrossChainStatus, bscCrossChainStatus, processedTxHashes.size, isProcessing, baseContractAddress, bscContractAddress]);

  // On contract address change, try hydrate from cache
  useEffect(() => {
    const entry = historyCache.get(cacheKey);
    if (entry) {
      setProcessedTransactions(entry.items);
      setProcessedTxHashes(new Set(Array.from(entry.hashes)));
      setHasInitialLoaded(true);
    }
  }, [cacheKey]);

  // Mark initial load done only when:
  // 1) we already have processed transactions, or
  // 2) both chains have responded AND there is definitively no data
  useEffect(() => {
    if (hasInitialLoaded) return;
    const baseLen = baseCrossChainStatus?.data?.length || 0;
    const bscLen = bscCrossChainStatus?.data?.length || 0;
    const hasProcessed = processedTransactions.length > 0;
    const bothResolved = (baseCrossChainStatus !== undefined && bscCrossChainStatus !== undefined);
    const bothResolvedAndEmpty = bothResolved && (baseLen + bscLen === 0);

    if (hasProcessed || bothResolvedAndEmpty) setHasInitialLoaded(true);
  }, [hasInitialLoaded, baseCrossChainStatus, bscCrossChainStatus, processedTransactions.length, isProcessing]);

  // Use processed transactions or fall back to mock data
  const userHistory = useMemo(() => {
    const totalRawTransactions = (baseCrossChainStatus?.data?.length || 0) + (bscCrossChainStatus?.data?.length || 0);
    
    console.log(`🔄 Computing userHistory:`, {
      processedTransactionsLength: processedTransactions.length,
      baseChainTransactions: baseCrossChainStatus?.data?.length || 0,
      bscChainTransactions: bscCrossChainStatus?.data?.length || 0,
      totalRawTransactions,
      processedTxHashesSize: processedTxHashes.size,
      contractAddresses: {
        base: baseContractAddress,
        bsc: bscContractAddress
      }
    });
    
    if (processedTransactions.length > 0) {
      console.log(`✅ Using processed transactions:`, processedTransactions);
      return processedTransactions;
    }
    
    // Fall back to mock data for demonstration when no real data is processed yet
    if (totalRawTransactions === 0) {
      console.log(`📝 Using mock data (no cross-chain data from either chain)`);
      // return mockHistoryData;
      return [];
    }
    
    console.log(`⏳ No processed transactions yet, but we have ${totalRawTransactions} raw transactions. Waiting for processing...`);
    return [];
  }, [processedTransactions, baseCrossChainStatus, bscCrossChainStatus, processedTxHashes.size]);
  
  // Determine initial fetching state to avoid empty flicker
  const baseLen = baseCrossChainStatus?.data?.length || 0;
  const bscLen = bscCrossChainStatus?.data?.length || 0;
  const bothFetchedOnce = (baseCrossChainStatus !== undefined && bscCrossChainStatus !== undefined);

  // Build the set of raw tx hashes we expect to process for this snapshot
  const expectedRawHashes = useMemo(() => {
    const hashes = new Set<string>();
    if (baseCrossChainStatus?.data) {
      for (const tx of baseCrossChainStatus.data) {
        if (tx?.transactionHash) hashes.add(tx.transactionHash);
      }
    }
    if (bscCrossChainStatus?.data) {
      for (const tx of bscCrossChainStatus.data) {
        if (tx?.transactionHash) hashes.add(tx.transactionHash);
      }
    }
    return hashes;
  }, [baseCrossChainStatus, bscCrossChainStatus]);

  // Processing is complete when all expected hashes have been processed
  const hasProcessedAllExpected = useMemo(() => {
    if (expectedRawHashes.size === 0) return true; // nothing to process
    const hashes = Array.from(expectedRawHashes);
    for (let i = 0; i < hashes.length; i++) {
      if (!processedTxHashes.has(hashes[i])) return false;
    }
    return true;
  }, [expectedRawHashes, processedTxHashes]);

  // First-load pending if: not yet marked loaded AND (not both fetched OR networks still loading OR processing not complete)
  const networksLoading = isValidatingBaseCrossChainStatus || isValidatingBscCrossChainStatus || isLoadingBaseCrossChainStatus || isLoadingBscCrossChainStatus;
  const isFirstLoadPending = !hasInitialLoaded && (!bothFetchedOnce || networksLoading || isProcessing || !hasProcessedAllExpected);

  // Keep table loading while header shows loading and list is empty
  const isHeaderLoading = isValidatingBaseCrossChainStatus || isValidatingBscCrossChainStatus;

  // Pagination
  const totalPages = Math.ceil(userHistory.length / PAGE_SIZE);
  const paginatedHistory = useMemo(() => {
    const startIndex = (page - 1) * PAGE_SIZE;
    return userHistory.slice(startIndex, startIndex + PAGE_SIZE);
  }, [userHistory, page]);

  return (
    <div className="space-y-4">
      <div className="border border-border bg-bg-card px-4 py-4 flex flex-col gap-1 text-xs">
        <div className="flex items-center justify-between">
          <div>
            <p className="text-[11px] font-bold uppercase tracking-[0.3em] text-secondary">History</p>
          </div>
          <div className="flex flex-col items-end gap-1 text-secondary">
            {(isValidatingBaseCrossChainStatus || isValidatingBscCrossChainStatus) && (
              <div className="flex items-center gap-2 text-[10px] uppercase tracking-[0.2em]">
                <div className="w-3 h-3 border-2 border-secondary border-t-transparent rounded-full animate-spin" />
                Syncing
              </div>
            )}
            {isProcessing && (
              <div className="flex items-center gap-2 text-[10px] uppercase tracking-[0.2em] text-blue-400">
                <div className="w-3 h-3 border-2 border-blue-400 border-t-transparent rounded-full animate-spin" />
                Processing
              </div>
            )}
            <span className="text-[11px] tracking-[0.25em]">{userHistory.length} transfers</span>
          </div>
        </div>
        <div className="text-[11px] tracking-[0.2em] text-secondary">
          Base: {baseCrossChainStatus?.data?.length || 0} • BSC: {bscCrossChainStatus?.data?.length || 0} • Processed: {processedTxHashes.size}
        </div>
      </div>

      {/* Desktop table */}
      <div className="hidden sm:block border border-border bg-bg-card">
        <div className="overflow-x-auto">
          <Table className="w-full min-w-[560px] text-[13px]">
            <TableHeader>
              <TableRow className="border-b border-border/60">
                <TableHead className="px-4 py-3 text-left text-secondary uppercase tracking-[0.2em] text-[11px]">Status</TableHead>
                <TableHead className="px-4 py-3 text-center text-secondary uppercase tracking-[0.2em] text-[11px]">Token</TableHead>
                <TableHead className="px-3 py-3 text-center text-secondary uppercase tracking-[0.2em] text-[11px]">Route</TableHead>
                <TableHead className="px-4 py-3 text-center text-secondary uppercase tracking-[0.2em] text-[11px]">Sender</TableHead>
                <TableHead className="px-4 py-3 text-center text-secondary uppercase tracking-[0.2em] text-[11px]">Receiver</TableHead>
                <TableHead className="px-4 py-3 text-center text-secondary uppercase tracking-[0.2em] text-[11px]">Transaction</TableHead>
                <TableHead className="px-4 py-3 text-right text-secondary uppercase tracking-[0.2em] text-[11px]">Time</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {isFirstLoadPending || (isHeaderLoading && userHistory.length === 0) ? (
                <TableRow>
                  <TableCell colSpan={7} className="py-12 text-center text-secondary">
                    <div className="flex flex-col items-center gap-3">
                      <div className="w-8 h-8 border border-fluxus-primary border-t-transparent rounded-full animate-spin" />
                      <span className="tracking-[0.2em] uppercase text-[11px]">Loading history</span>
                    </div>
                  </TableCell>
                </TableRow>
              ) : userHistory.length === 0 ? (
                <TableRow>
                  <TableCell colSpan={7} className="py-12 text-center text-secondary">
                    <div className="flex flex-col items-center gap-2">
                      <div className="text-4xl">🌉</div>
                      <div className="flex items-center gap-2 text-secondary">
                        <div className="w-4 h-4 border border-secondary border-t-fluxus-primary rounded-full animate-spin" />
                        <span className="text-sm">No history yet — syncing latest data...</span>
                      </div>
                      <span className="text-xs text-secondary/70">Recent transfers may take up to a minute to appear</span>
                    </div>
                  </TableCell>
                </TableRow>
              ) : (
                paginatedHistory.map((item) => {
                  const statusDisplay = getStatusDisplay(item.status);
                  return (
                    <TableRow key={item.id} className="border-border/60 hover:bg-bg-elevated transition-colors">
                      <TableCell className="px-4 py-3">
                        <div className={`inline-flex items-center gap-2 px-3 py-1 text-xs font-semibold rounded-full ${statusDisplay.bgColor} ${statusDisplay.color}`}>
                          <span>{statusDisplay.icon}</span>
                          {statusDisplay.text}
                        </div>
                      </TableCell>
                      <TableCell className="px-4 py-3 text-center">
                        <span className="inline-flex font-mono text-xs bg-bg-elevated px-2 py-1 border border-border/70">
                          #{item.tokenId}
                        </span>
                      </TableCell>
                      <TableCell className="px-3 py-3 text-center text-secondary">
                        <div className="flex flex-col items-center gap-1 text-[11px] uppercase tracking-[0.18em] leading-tight">
                          <span>{item.sourceChain}</span>
                          <span className="text-border text-[10px] tracking-[0.3em]">↓</span>
                          <span>{item.targetChain}</span>
                        </div>
                      </TableCell>
                      <TableCell className="px-4 py-3 text-center">
                        {item.sender ? (
                          <button
                            className="text-primary font-mono text-[11px] border border-border/70 px-2 py-1 hover:border-fluxus-primary hover:text-fluxus-primary transition-colors"
                            onClick={() => copyToClipboard(item.sender!, 'sender address')}
                            title="Copy sender address"
                          >
                            {sliceAddress(item.sender)}
                          </button>
                        ) : (
                          <span className="text-secondary text-xs">-</span>
                        )}
                      </TableCell>
                      <TableCell className="px-4 py-3 text-center">
                        {item.receiver ? (
                          <button
                            className="text-primary font-mono text-[11px] border border-border/70 px-2 py-1 hover:border-fluxus-primary hover:text-fluxus-primary transition-colors"
                            onClick={() => copyToClipboard(item.receiver!, 'receiver address')}
                            title="Copy receiver address"
                          >
                            {sliceAddress(item.receiver)}
                          </button>
                        ) : (
                          <span className="text-secondary text-xs">-</span>
                        )}
                      </TableCell>
                      <TableCell className="px-4 py-3">
                        <div className="flex items-center justify-center gap-2">
                          <span className="text-primary font-mono text-[11px]">{sliceAddress(item.transactionHash)}</span>
                          <button
                            className="p-1 border border-transparent hover:border-border/70 hover:bg-bg-elevated transition-colors"
                            onClick={() => copyToClipboard(item.transactionHash, 'transaction hash')}
                            title="Copy transaction hash"
                          >
                            ⧉
                          </button>
                        </div>
                      </TableCell>
                      <TableCell className="px-4 py-3 text-right text-secondary font-mono text-[11px] whitespace-nowrap">
                        {formatTime(item.blockTimestamp)}
                      </TableCell>
                    </TableRow>
                  );
                })
              )}
            </TableBody>
          </Table>
        </div>
      </div>
      {/* Mobile cards */}
      <div className="sm:hidden space-y-3">
        {isFirstLoadPending || (isHeaderLoading && userHistory.length === 0) ? (
          <div className="border border-border bg-bg-card py-10 text-center">
            <div className="flex flex-col items-center gap-3">
              <div className="w-8 h-8 border-2 border-fluxus-primary border-t-transparent rounded-full animate-spin"></div>
              <span className="text-secondary">Loading history...</span>
            </div>
          </div>
        ) : userHistory.length === 0 ? (
          <div className="border border-border bg-bg-card py-10 text-center">
            <div className="flex flex-col items-center gap-2">
              <div className="text-4xl">🌉</div>
              <div className="flex items-center gap-2 text-secondary">
                <div className="w-4 h-4 border-2 border-secondary border-t-fluxus-primary rounded-full animate-spin"></div>
                <span className="text-sm">No history yet — syncing latest data...</span>
              </div>
              <span className="text-xs text-secondary">Recent transfers may take up to a minute to appear</span>
            </div>
          </div>
        ) : (
          paginatedHistory.map((item) => {
            const statusDisplay = getStatusDisplay(item.status);
            return (
              <div key={item.id} className="border border-border bg-bg-card p-4">
                <div className="flex items-start justify-between mb-2">
                  <div className={`inline-flex items-center gap-2 px-3 py-1.5 text-xs font-bold uppercase tracking-[0.2em] ${statusDisplay.bgColor} ${statusDisplay.color}`}>
                    <span className="text-xs">{statusDisplay.icon}</span>
                    {statusDisplay.text}
                  </div>
                  <div className="text-secondary text-xs">{formatTime(item.blockTimestamp)}</div>
                </div>
                <div className="grid grid-cols-2 gap-3 text-sm">
                  <div className="border border-border px-3 py-2">
                    <div className="text-xs text-secondary mb-1 uppercase tracking-[0.2em]">Token</div>
                    <div className="text-primary font-mono">#{item.tokenId}</div>
                  </div>
                  <div className="border border-border px-3 py-2">
                    <div className="text-xs text-secondary mb-1 uppercase tracking-[0.2em]">Route</div>
                    <div className="text-primary">
                      {item.sourceChain} → {item.targetChain}
                    </div>
                  </div>
                </div>
                
                {/* Sender and Receiver for mobile */}
                {(item.sender || item.receiver) && (
                  <div className="grid grid-cols-1 gap-3 mt-3">
                    {item.sender && (
                      <button
                        className="border border-border px-3 py-2 text-left hover:text-primary"
                        onClick={() => copyToClipboard(item.sender!, 'sender address')}
                        title="Copy sender address"
                      >
                        <div className="text-xs text-secondary mb-1 uppercase tracking-[0.2em]">Sender</div>
                        <div className="text-sm text-primary font-mono">
                          {sliceAddress(item.sender)}
                        </div>
                      </button>
                    )}
                    {item.receiver && (
                      <button
                        className="border border-border px-3 py-2 text-left hover:text-primary"
                        onClick={() => copyToClipboard(item.receiver!, 'receiver address')}
                        title="Copy receiver address"
                      >
                        <div className="text-xs text-secondary mb-1 uppercase tracking-[0.2em]">Receiver</div>
                        <div className="text-sm text-primary font-mono">
                          {sliceAddress(item.receiver)}
                        </div>
                      </button>
                    )}
                  </div>
                )}
                <div className="flex items-center justify-between mt-3">
                  <div className="text-secondary font-mono text-sm">{sliceAddress(item.transactionHash)}</div>
                  <button
                    className="p-1 border border-border text-secondary hover:text-primary"
                    onClick={() => copyToClipboard(item.transactionHash, 'transaction hash')}
                    title="Copy transaction hash"
                  >
                    ⧉
                  </button>
                </div>
              </div>
            );
          })
        )}
      </div>

      {/* Pagination */}
      {userHistory.length > PAGE_SIZE && (
        <div className="flex justify-center mt-6">
          <Pagination
            totalPages={totalPages}
            offset={page - 1}
            onPageChange={setPage}
          />
        </div>
      )}
    </div>
  );
}
