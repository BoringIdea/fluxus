'use client'
import React, { useState, useEffect, useMemo, useRef } from "react";
import useSWR from 'swr';
import { 
  useAccount, 
  useChainId, 
  useReadContract,
  useWriteContract,
  useWaitForTransactionReceipt,
  useSwitchChain 
} from "wagmi";
import { parseEther } from "viem";
import { Button } from "./ui/button";
import { cn } from "@/lib/utils";
import TransactionDialog from "@/components/shared/TransactionDialog";
import { useTransactionDialog } from "@/src/hooks/useTransactionDialog";
import { Collection } from "@/src/api/types";
import { useCrossChainStatus } from "@/src/api";
import { getNFTsByOwnerForCollection } from "@/src/api";
import Loading from './ui/Loading';
import { ChevronDownIcon } from '@radix-ui/react-icons';
import CrosschainHistory from './CrosschainHistory';
import { defaultCCTXTracker, CCTXRecord } from '@/lib/cctxTracker';
import { getChainName } from '@/lib/chains';
import { Fluxus_EVM_CROSS_CHAIN_ABI, TRADE_ABI, getTradeContractAddress, FACTORY_CONTRACT_ABI, getFactoryContractAddress } from '@/src/contract';
import { SupportedChainId } from '@/src/alchemy';
import { buildSetApprovalForAllTx, buildTransferCrossChainTx } from '@/src/onchain/tradeTx';
import { useContractAddress } from '@/hooks/useContractAddress';

interface CrossChainProps {
  contractAddress: string;
  collection?: Collection;
}

interface ChainInfo {
  id: number;
  name: string;
  symbol: string;
  icon: string;
}

const CHAINS: Record<string, ChainInfo> = {
  'base-sepolia': {
    id: 84532,
    name: 'Base Sepolia',
    symbol: 'ETH',
    icon: '🔵'
  },
  'bsc-testnet': {
    id: 97,
    name: 'BSC Testnet',
    symbol: 'BNB',
    icon: '🟡'
  }
};

const mockTxHash = '0xf11ffae73c222fd259bf5cf8518eb4a14a827bb898b48ae7046d5b67a679cdb5';

// ZRC20 destination address
const DESTINATION_ADDRESS_BSC = '0xd97B1de3619ed2c6BEb3860147E30cA8A7dC9891';
const DESTINATION_ADDRESS_BASE = '0x236b0DE675cC8F46AE186897fCCeFe3370C9eDeD';

const DEFAULT_CROSS_CHAIN_GAS_FEE = parseEther("0.001"); // 0.001 ETH
const DEFAULT_GATEWAY_ADDRESS = "0x0c487a766110c85d301d96e33579c5b317fa4995";
const DEFAULT_GAS_LIMIT = BigInt(12000000);

const tabs = [
  { key: 'transfer', label: 'Transfer' },
  { key: 'history', label: 'History' },
] as const;

const TabsContext = React.createContext<{ value: string; onValueChange: (value: string) => void } | null>(null);

function Tabs({
  value,
  onValueChange,
  children,
}: {
  value: string;
  onValueChange: (value: string) => void;
  children: React.ReactNode;
}) {
  return (
    <TabsContext.Provider value={{ value, onValueChange }}>
      {children}
    </TabsContext.Provider>
  );
}

function TabsList({ children, className }: { children: React.ReactNode; className?: string }) {
  return <div className={className}>{children}</div>;
}

function TabsTrigger({
  value,
  children,
  className,
}: {
  value: string;
  children: React.ReactNode;
  className?: string;
}) {
  const context = React.useContext(TabsContext);
  if (!context) return null;
  const isActive = context.value === value;
  const activeClasses = isActive ? "text-fluxus-primary bg-bg-tertiary border-l-2 border-fluxus-primary" : "text-secondary border-l-2 border-transparent";
  return (
    <button
      type="button"
      className={cn("w-full text-left uppercase tracking-[0.3em] text-[11px] px-4 py-4", activeClasses, className)}
      onClick={() => context.onValueChange(value)}
    >
      {children}
    </button>
  );
}

function TabsContent({
  value,
  children,
  className,
}: {
  value: string;
  children: React.ReactNode;
  className?: string;
}) {
  const context = React.useContext(TabsContext);
  if (!context || context.value !== value) return null;
  return <div className={className}>{children}</div>;
}

const getDestinationAddress = (chainId: number) => {
  if (chainId === 97) {
    // cross to base
    return DESTINATION_ADDRESS_BASE;
  } else if (chainId === 84532) {
    // cross to bsc
    return DESTINATION_ADDRESS_BSC;
  }
  return DESTINATION_ADDRESS_BSC;
};

export default function CrossChain({ contractAddress, collection }: CrossChainProps) {
  const { address } = useAccount();
  const chainId = useChainId();
  
  // State management
  const [activeTab, setActiveTab] = useState<'transfer' | 'history'>('transfer');
  const [sourceChain, setSourceChain] = useState<ChainInfo>(CHAINS['base-sepolia']);
  const [targetChain, setTargetChain] = useState<ChainInfo>(CHAINS['bsc-testnet']);
  const [selectedTokenId, setSelectedTokenId] = useState<string>('');
  const [receiverAddress, setReceiverAddress] = useState<string>('');
  const [isTransferring, setIsTransferring] = useState(false);
  const [tokenDropdownOpen, setTokenDropdownOpen] = useState(false);
  const [cctxData, setCctxData] = useState<CCTXRecord | null>(null);
  const [isTrackingCctx, setIsTrackingCctx] = useState(false);
  const [cctxHistory, setCctxHistory] = useState<CCTXRecord[]>([]);
  const [pendingApprovalTokenId, setPendingApprovalTokenId] = useState<string | null>(null);
  const [isApproving, setIsApproving] = useState(false);
  const tokenDropdownRef = useRef<HTMLDivElement>(null);
  
  // Contract interaction hooks
  const { data: hash, writeContract, isError: isWriteContractError, error: writeContractError, isPending } = useWriteContract();
  const { isLoading: isConfirming, isSuccess: isConfirmed, isError, error: contractError } = useWaitForTransactionReceipt({ hash });
  const { switchChain, isPending: isSwitchingChain } = useSwitchChain();
  
  // Cross-chain status and user tokens
  const { data: crossChainStatus, isLoading: isLoadingCrossChainStatus } = useCrossChainStatus(chainId, collection?.address || '', address as string);
  
  // Check if user has approved the contract to spend NFTs
  const { data: isApproved, refetch: refetchIsApproved } = useReadContract({
    address: contractAddress as `0x${string}`,
    abi: Fluxus_EVM_CROSS_CHAIN_ABI,
    functionName: 'isApprovedForAll',
    args: [address, getTradeContractAddress(chainId as SupportedChainId)],
    query: { 
      enabled: !!address && !!contractAddress,
      retry: false 
    },
  });

  // Memoize contract params to prevent unnecessary re-renders
  const contractParams = useMemo(() => {
    if (!collection) return undefined;
    
    return {
      name: collection.name,
      symbol: collection.symbol,
      initialPrice: BigInt(collection.initial_price),
      maxSupply: BigInt(collection.max_supply),
      maxPrice: BigInt(collection.max_price),
      creatorFeePercent: BigInt(collection.creator_fee),
      imageUrl: collection.base_uri,
      gatewayAddress: DEFAULT_GATEWAY_ADDRESS,
      gasLimit: BigInt(DEFAULT_GAS_LIMIT),
      supportMint: collection.support_mint,
      creator: collection.creator, // 🎯 Add creator address from collection
    };
  }, [collection]);

  // Use custom hook to get contract addresses with ethers.js
  const { 
    data: sourceChainContractAddress, 
    error: sourceChainError, 
    isLoading: sourceChainLoading 
  } = useContractAddress(sourceChain.id, contractParams);

  const { 
    data: targetChainContractAddress, 
    error: targetChainError, 
    isLoading: targetChainLoading 
  } = useContractAddress(targetChain.id, contractParams);

  useEffect(() => {
    const hashToTrack = crossChainStatus?.data[0]?.transactionHash;
    
    if (hashToTrack && !isTrackingCctx) {
      const updateCCTXRecord = (record: CCTXRecord) => {
        // Always use blockTimestamp from crossChainStatus API
        const originalData = crossChainStatus?.data[0];
        if (originalData && originalData.blockTimestamp) {
          record = {
            ...record,
            blockTimestamp: new Date(originalData.blockTimestamp * 1000)
          };
        } else {
          console.warn('⚠️ No blockTimestamp available from crossChainStatus API');
        }
        
        setCctxData(record);
        setCctxHistory(prev => {
          const exists = prev.find(item => item.transactionHash === record.transactionHash);
          if (!exists) {
            return [record, ...prev];
          }
          return prev.map(item => item.transactionHash === record.transactionHash ? record : item);
        });
      };

      defaultCCTXTracker.trackCCTX(hashToTrack, {
        onCCTXUpdate: updateCCTXRecord,
        onTrackingStart: () => setIsTrackingCctx(true),
        onTrackingComplete: () => setIsTrackingCctx(false),
        onError: (error) => {
          console.error('CCTX tracking failed:', error);
        }
      }).catch(error => {
        console.error('CCTX tracking promise rejected:', error);
      });
    }
  }, [crossChainStatus, isTrackingCctx]);

  useEffect(() => {
    if (!tokenDropdownOpen) return;
    const handleClick = (event: MouseEvent) => {
      if (tokenDropdownRef.current && !tokenDropdownRef.current.contains(event.target as Node)) {
        setTokenDropdownOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClick);
    return () => document.removeEventListener("mousedown", handleClick);
  }, [tokenDropdownOpen]);

  // Handle transaction errors
  useEffect(() => {
    if (isError) {
      handleTransactionError(contractError?.toString() || 'Unknown error');
    }
  }, [isError, contractError]);
  
  // Get user's NFTs from the selected source chain via Alchemy
  const { data: alchemyUserNFTs, isLoading: isLoadingTokens } = useSWR(
    address && sourceChainContractAddress ? ["alchemy-nfts", sourceChain.id, address, sourceChainContractAddress] : null,
    () => getNFTsByOwnerForCollection(sourceChain.id as SupportedChainId, address as string, sourceChainContractAddress as string, true, 100),
    { revalidateOnFocus: false, revalidateIfStale: true, keepPreviousData: false }
  );
  
  // Process user NFTs for dropdown
  const userNFTs = useMemo(() => {
    if (!alchemyUserNFTs?.ownedNfts) {
      return [];
    }
    return alchemyUserNFTs.ownedNfts.map((nft: any) => ({
      tokenId: nft.tokenId.toString(),
      name: `${collection?.symbol || 'NFT'} #${nft.tokenId}`,
      image: nft.image?.cachedUrl || 'https://red-naval-coyote-720.mypinata.cloud/ipfs/bafybeibczf5smgrsb5dje6z2na5xvdb5brkl6apbkoy3myjukvfusi5hmm/Drug-nft.png',
    }));
  }, [alchemyUserNFTs, collection?.symbol, chainId]);

  // When source chain contract address changes, reset selected token and close dropdown
  useEffect(() => {
    setSelectedTokenId('');
    setTokenDropdownOpen(false);
  }, [sourceChainContractAddress]);
  
  // Transaction dialog
  const { dialogState, onOpenChange, setDialogState } = useTransactionDialog({
    hash,
    isPending,
    isConfirming,
    isConfirmed,
    isWriteContractError,
    writeContractError,
    onConfirmed: () => {
      if (pendingApprovalTokenId !== null) {
        // Approval completed, now execute cross-chain transfer
        refetchIsApproved().then(() => {
          executeCrossChainTransfer(pendingApprovalTokenId);
          setPendingApprovalTokenId(null);
        });
      } else {
        // Cross-chain transfer completed
        setIsTransferring(false);
        setSelectedTokenId('');
      }
    }
  });
  
  // Handle chain selection
  const handleSourceChainSelect = (chainKey: string) => {
    const selectedChain = CHAINS[chainKey];
    setSourceChain(selectedChain);
    // Reset selected token when source chain changes
    setSelectedTokenId('');
    setTokenDropdownOpen(false);
  };
  
  const handleTargetChainSelect = (chainKey: string) => {
    setTargetChain(CHAINS[chainKey]);
  };
  
  // Handle token selection
  const handleTokenSelect = (tokenId: string) => {
    setSelectedTokenId(tokenId);
    setTokenDropdownOpen(false);
  };
  
  // Handle receiver address input
  const handleReceiverChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setReceiverAddress(event.target.value);
  };
  
  // Fill receiver with user's own address
  const fillSelfAddress = () => {
    if (address) {
      setReceiverAddress(address);
    }
  };
  
  // Handle chain swap
  const handleChainSwap = () => {
    setSourceChain(targetChain);
    setTargetChain(sourceChain);
    // Reset selected token when chains are swapped
    setSelectedTokenId('');
    setTokenDropdownOpen(false);
  };
  
  // Helper function to handle transaction errors
  const handleTransactionError = (errorMessage: string) => {
    setDialogState({ isOpen: true, status: 'error', error: errorMessage });
    setPendingApprovalTokenId(null);
    setIsTransferring(false);
    setIsApproving(false);
  };

  // Execute approval operation
  const executeApproval = (tokenId: string) => {
    setPendingApprovalTokenId(tokenId);
    setIsApproving(true);
    writeContract(
      buildSetApprovalForAllTx({
        nftContractAddress: contractAddress,
        operatorAddress: getTradeContractAddress(chainId as SupportedChainId),
        approved: true,
      }) as any
    );
  };

  // Execute cross-chain transfer
  const executeCrossChainTransfer = (tokenId: string) => {
    if (!address) return;
    
    setIsTransferring(true);
    
    const finalReceiver = receiverAddress || address; // Use custom receiver or fallback to user's address
    
    writeContract(
      buildTransferCrossChainTx({
        tradeContractAddress: getTradeContractAddress(chainId as SupportedChainId),
        fluxusContractAddress: contractAddress,
        tokenId: Number(tokenId),
        receiver: finalReceiver,
        destination: getDestinationAddress(chainId as SupportedChainId),
        gasFee: DEFAULT_CROSS_CHAIN_GAS_FEE,
      }) as any
    );
  };

  // Main cross-chain transfer function
  const handleCrossChainTransfer = async () => {
    try {
      if (!selectedTokenId || !address) return;
      
      // Check if user is on the correct network
      if (chainId !== sourceChain?.id) {
        try {
          // Show wallet network switch dialog
          if (sourceChain?.id && switchChain) {
            await switchChain({ chainId: sourceChain.id });
          }
          // After successful switch, the component will re-render with new chainId
          return;
        } catch (switchError) {
          console.error('User rejected network switch or switch failed:', switchError);
          return;
        }
      }
      
      const finalReceiver = receiverAddress || address;
      
      console.log('Cross-chain transfer:', {
        tokenId: selectedTokenId,
        sourceChain: sourceChain.name,
        targetChain: targetChain.name,
        userAddress: address,
        receiver: finalReceiver,
        destination: getDestinationAddress(chainId as SupportedChainId),
        isApproved: isApproved
      });

      if (!isApproved) {
        executeApproval(selectedTokenId);
      } else {
        executeCrossChainTransfer(selectedTokenId);
      }
    } catch (error) {
      console.error('Error in handleCrossChainTransfer:', error);
      // Reset any loading states
      setIsTransferring(false);
      setIsApproving(false);
    }
  };
  
  return (
    <div className="px-3 sm:px-6 pb-24 mt-4">
      <Tabs value={activeTab} onValueChange={(value) => setActiveTab(value as 'transfer' | 'history')}>
        <div className="border border-border bg-bg-card flex flex-col lg:flex-row">
          <TabsList className="lg:w-40 border-b lg:border-b-0 lg:border-r border-border flex lg:flex-col">
            {tabs.map((tab) => (
              <TabsTrigger key={tab.key} value={tab.key}>
                {tab.label}
              </TabsTrigger>
            ))}
          </TabsList>
          <div className="flex-1">
          <TabsContent value="transfer" className="p-4 lg:p-6">
            {/* Main Content */}
            <div className="flex flex-col lg:flex-row gap-6 mt-3">
              {/* Left Section: Operation Controls */}
              <div className="lg:w-1/2 w-full flex justify-center lg:justify-start">
                <div className="w-full max-w-md">
                  <div className="p-5 space-y-6 border border-border/80 bg-gradient-to-b from-black/80 via-bg-card to-bg-card/90 overflow-hidden">
                    {/* Source Chain Selection */}
                    <div className="space-y-3">
                      <div className="flex items-center justify-between text-[11px] font-semibold uppercase tracking-[0.35em] text-secondary">
                        <span>Source Chain</span>
                        {address && sourceChain?.id && chainId !== sourceChain.id && (
                          <span className="text-[10px] text-yellow-300 bg-yellow-500/10 px-3 py-1 tracking-[0.25em]">
                            Switch required
                          </span>
                        )}
                      </div>
                      <div className="space-y-3 rounded border border-fluxus-primary/40 bg-black/30 p-4">
                        <button
                          className="w-full flex items-center justify-between px-4 py-3 border border-fluxus-primary/50 bg-black/40 text-left uppercase tracking-[0.2em] text-primary/80 cursor-not-allowed"
                          disabled
                          aria-disabled="true"
                          title="Source chain selection is temporarily disabled"
                        >
                          <div className="flex items-center gap-3">
                            <span className="text-xl">{sourceChain.icon}</span>
                            <span className="text-sm text-primary tracking-[0.2em]">{sourceChain.name}</span>
                            {isSwitchingChain && (
                              <div className="w-4 h-4 border border-fluxus-primary border-t-transparent rounded-full animate-spin" />
                            )}
                          </div>
                          <ChevronDownIcon className="w-4 h-4 text-secondary" />
                        </button>

                        {/* Source Chain Contract Address */}
                        <div className="text-[11px] uppercase tracking-[0.3em] text-secondary">Contract Address</div>
                        {sourceChainLoading ? (
                          <div className="px-3 py-2 border border-border/70 bg-black/40 text-[11px] text-yellow-400 tracking-[0.2em]">
                            Loading...
                          </div>
                        ) : sourceChainError ? (
                          <div className="px-3 py-2 border border-red-500/30 bg-red-500/5 text-[11px] text-red-400 tracking-[0.2em]">
                            Error: Failed to calculate address
                          </div>
                        ) : sourceChainContractAddress ? (
                          <div className="flex items-center gap-2 px-3 py-2 border border-fluxus-primary/40 bg-black/40">
                            <span className="text-xs text-white font-mono break-all flex-1">
                              {sourceChainContractAddress as string}
                            </span>
                            <button
                              onClick={() => navigator.clipboard.writeText(sourceChainContractAddress as string)}
                              className="text-[10px] uppercase tracking-[0.3em] text-fluxus-primary px-3 py-1 border border-fluxus-primary/50 bg-fluxus-primary/10 hover:bg-fluxus-primary/20 transition-colors"
                            >
                              Copy
                            </button>
                          </div>
                        ) : (
                          <div className="px-3 py-2 border border-border/70 bg-black/40 text-[11px] tracking-[0.2em] text-secondary">
                            Not available
                          </div>
                        )}
                      </div>
                    </div>

                    {/* Transfer Direction Arrow */}
                    <div className="flex justify-center">
                      <button
                        onClick={(e) => e.preventDefault()}
                        className="w-12 h-12 border border-fluxus-primary/30 bg-black/40 text-fluxus-primary text-xl flex items-center justify-center tracking-[0.3em] cursor-not-allowed"
                        disabled
                        aria-disabled="true"
                        title="Chain swap is temporarily disabled"
                      >
                        ↓
                      </button>
                    </div>

                    {/* Target Chain Selection */}
                    <div className="space-y-3">
                      <div className="text-[11px] font-semibold uppercase tracking-[0.35em] text-secondary">Target Chain</div>
                      <div className="space-y-3 rounded border border-fluxus-primary/40 bg-black/30 p-4">
                        <button
                          className="w-full flex items-center justify-between px-4 py-3 border border-fluxus-primary/50 bg-black/40 text-left uppercase tracking-[0.2em] text-primary/80 cursor-not-allowed"
                          disabled
                          aria-disabled="true"
                          title="Target chain selection is temporarily disabled"
                        >
                          <div className="flex items-center gap-3">
                            <span className="text-xl">{targetChain.icon}</span>
                            <span className="text-sm text-primary tracking-[0.2em]">{targetChain.name}</span>
                          </div>
                          <ChevronDownIcon className="w-4 h-4 text-secondary" />
                        </button>

                        <div className="text-[11px] uppercase tracking-[0.3em] text-secondary">Contract Address</div>
                        {targetChainLoading ? (
                          <div className="px-3 py-2 border border-border/70 bg-black/40 text-[11px] text-yellow-400 tracking-[0.2em]">
                            Loading...
                          </div>
                        ) : targetChainContractAddress ? (
                          <div className="flex items-center gap-2 px-3 py-2 border border-fluxus-primary/40 bg-black/40">
                            <span className="text-xs text-white font-mono break-all flex-1">
                              {targetChainContractAddress as string}
                            </span>
                            <button
                              onClick={() => navigator.clipboard.writeText(targetChainContractAddress as string)}
                              className="text-[10px] uppercase tracking-[0.3em] text-fluxus-primary px-3 py-1 border border-fluxus-primary/50 bg-fluxus-primary/10 hover:bg-fluxus-primary/20 transition-colors"
                            >
                              Copy
                            </button>
                          </div>
                        ) : (
                          <div className="px-3 py-2 border border-border/70 bg-black/40 text-[11px] tracking-[0.2em] text-secondary">
                            Not available
                          </div>
                        )}
                      </div>
                    </div>

                    {/* Token ID Selection */}
                    <div className="space-y-3">
                      <div className="text-[11px] font-semibold uppercase tracking-[0.35em] text-secondary">Select NFT</div>
                      <div className="relative" ref={tokenDropdownRef}>
                        <button
                          className="w-full flex items-center justify-between px-4 py-3 border border-border/70 bg-black/40 hover:border-fluxus-primary/50 transition-colors"
                          disabled={sourceChainLoading || isLoadingTokens}
                          onClick={() => !sourceChainLoading && !isLoadingTokens && setTokenDropdownOpen(prev => !prev)}
                          type="button"
                        >
                          <div className="flex items-center gap-3">
                            {(() => {
                              const selected = userNFTs.find((n: any) => n.tokenId === selectedTokenId);
                              return selected?.image ? (
                                <div className="w-7 h-7 bg-bg-card overflow-hidden flex-shrink-0 border border-border/70">
                                  <img src={selected.image} alt={`NFT #${selectedTokenId}`} className="w-full h-full object-cover" />
                                </div>
                              ) : (
                                <span className="text-lg">🎨</span>
                              );
                            })()}
                            <span className="text-primary font-medium text-sm tracking-[0.1em]">
                              {selectedTokenId ? `NFT #${selectedTokenId}` : 'Select NFT to transfer'}
                            </span>
                          </div>
                          <ChevronDownIcon className="w-4 h-4 text-secondary" />
                        </button>
                        {tokenDropdownOpen && (
                          <div className="absolute z-50 mt-2 w-full max-h-60 overflow-y-auto border border-border bg-black/95 p-2 space-y-2">
                            {sourceChainLoading || isLoadingTokens ? (
                              <div className="p-4 text-center text-secondary">
                                <Loading />
                              </div>
                            ) : userNFTs.length === 0 ? (
                              <div className="p-4 text-center text-secondary text-sm tracking-[0.2em] uppercase">
                                No NFTs found
                              </div>
                            ) : (
                              userNFTs.map((nft: any) => (
                                <button
                                  key={nft.tokenId}
                                  onClick={() => {
                                    handleTokenSelect(nft.tokenId);
                                    setTokenDropdownOpen(false);
                                  }}
                                  className="w-full flex items-center gap-3 p-3 border border-border/60 bg-black/40 text-left hover:border-fluxus-primary/40 hover:bg-black/60 transition-colors"
                                  type="button"
                                >
                                  {nft.image ? (
                                    <div className="w-9 h-9 border border-border/70 bg-bg-card overflow-hidden flex-shrink-0">
                                      <img src={nft.image} alt={nft.name} className="w-full h-full object-cover" />
                                    </div>
                                  ) : (
                                    <span className="text-lg">🎨</span>
                                  )}
                                  <div>
                                    <div className="text-sm text-primary font-medium">{nft.name || `NFT #${nft.tokenId}`}</div>
                                    <div className="text-[11px] text-secondary tracking-[0.2em]">#{nft.tokenId}</div>
                                  </div>
                                </button>
                              ))
                            )}
                          </div>
                        )}
                      </div>
                    </div>

                    {/* Receiver Address Input */}
                    <div className="space-y-3">
                      <div className="text-[11px] font-semibold uppercase tracking-[0.35em] text-secondary">Receiver Address</div>
                      <div className="rounded border border-border/70 bg-black/30 p-3 space-y-3 overflow-hidden">
                        <div className="flex flex-col sm:flex-row gap-2 min-w-0">
                          <input
                            type="text"
                            value={receiverAddress}
                            onChange={handleReceiverChange}
                            placeholder="Enter receiver address (0x...)"
                            className="flex-1 px-3 py-2 border border-border/60 bg-black/40 text-primary placeholder-secondary focus:border-fluxus-primary focus:outline-none min-w-0"
                          />
                          <Button
                            variant="outline"
                            onClick={fillSelfAddress}
                            disabled={!address}
                            className="px-4 py-2 text-[11px] tracking-[0.25em] uppercase bg-black/40 border-fluxus-primary/40 text-primary hover:bg-black/60 w-full sm:w-auto"
                          >
                            Use My Address
                          </Button>
                        </div>
                        <div className="text-[11px] tracking-[0.2em] text-secondary break-words leading-relaxed">
                          {receiverAddress ? (
                            <>NFT will be sent to: {receiverAddress}</>
                          ) : (
                            <>If empty, NFT will be sent to {address ? `${address.slice(0, 6)}...${address.slice(-4)}` : 'your address'}</>
                          )}
                        </div>
                      </div>
                    </div>

                    {/* Transfer Summary */}
                    {selectedTokenId && (
                      <div className="p-4 border border-fluxus-primary/30 bg-gradient-to-b from-black/50 to-bg-card/30 space-y-2 text-sm tracking-[0.08em] uppercase text-secondary">
                        <div className="flex items-center justify-between text-[11px] tracking-[0.3em]">
                          <span>Transfer Summary</span>
                          <span className="text-fluxus-primary">Ready</span>
                        </div>
                        <div className="grid grid-cols-2 gap-3 text-xs text-primary">
                          <div>
                            <p className="text-secondary/70 text-[10px] tracking-[0.2em] mb-1">From</p>
                            <p>{sourceChain.name}</p>
                          </div>
                          <div>
                            <p className="text-secondary/70 text-[10px] tracking-[0.2em] mb-1">To</p>
                            <p>{targetChain.name}</p>
                          </div>
                          <div>
                            <p className="text-secondary/70 text-[10px] tracking-[0.2em] mb-1">Token</p>
                            <p>#{selectedTokenId}</p>
                          </div>
                          <div>
                            <p className="text-secondary/70 text-[10px] tracking-[0.2em] mb-1">Receiver</p>
                            <p>
                              {(() => {
                                const finalReceiver = receiverAddress || address || '';
                                return finalReceiver ? `${finalReceiver.slice(0, 6)}...${finalReceiver.slice(-4)}` : 'Not set';
                              })()}
                            </p>
                          </div>
                        </div>
                        <div className="text-[10px] tracking-[0.3em] text-secondary/80">
                          Estimated time: <span className="text-primary ml-2">2-5 minutes</span>
                        </div>
                      </div>
                    )}

                    {/* Transfer Button */}
                    <Button
                      onClick={handleCrossChainTransfer}
                      disabled={!selectedTokenId || isTransferring || isApproving || isPending || !address || isSwitchingChain}
                      className="w-full border border-fluxus-primary/60 bg-fluxus-primary text-black hover:bg-[#1FB455]"
                    >
                      {isSwitchingChain ? (
                        <div className="flex items-center gap-2">
                          <div className="w-4 h-4 border-2 border-black border-t-transparent rounded-full animate-spin"></div>
                          Switching Network...
                        </div>
                      ) : isApproving || (isPending && pendingApprovalTokenId !== null) ? (
                        <div className="flex items-center gap-2">
                          <div className="w-4 h-4 border-2 border-black border-t-transparent rounded-full animate-spin"></div>
                          Approving...
                        </div>
                      ) : isTransferring || (isPending && pendingApprovalTokenId === null) ? (
                        <div className="flex items-center gap-2">
                          <div className="w-4 h-4 border-2 border-black border-t-transparent rounded-full animate-spin"></div>
                          Transferring...
                        </div>
                      ) : address && sourceChain?.id && chainId !== sourceChain.id ? (
                        `Switch to ${sourceChain.name} & Transfer`
                      ) : (
                        (!isApproved ? 'Approve & Transfer' : 'Cross-Chain Transfer')
                      )}
                    </Button>

                    {/* Connect Wallet Notice */}
                    {!address && (
                      <div className="text-center p-3 border border-yellow-500/40 bg-yellow-500/5 text-[11px] uppercase tracking-[0.25em] text-yellow-300">
                        Connect wallet to enable transfers
                      </div>
                    )}
                  </div>
                </div>
              </div>
              
              {/* Right Section: Latest Transaction History (Desktop only) */}
              <div className="hidden lg:block lg:w-1/2">
                <div className="w-full">
                  <div className="p-4 space-y-3 border border-border bg-bg-card min-h-[600px]">
                    <div className="flex items-center justify-between">
                      <h3 className="text-sm font-medium text-gray-300">
                        YOUR CROSS-CHAIN TRANSACTION HISTORY
                      </h3>
                      {isTrackingCctx && (
                        <div className="flex items-center gap-2 text-xs text-blue-400">
                          <div className="w-3 h-3 border-2 border-blue-400 border-t-transparent rounded-full animate-spin"></div>
                          Tracking...
                        </div>
                      )}
                    </div>
                    
                    {/* Transaction History List */}
                    {(cctxData || cctxHistory.length > 0) ? (
                      <div className="space-y-3 max-h-[500px] overflow-y-auto">
                        {(cctxHistory.length > 0 ? cctxHistory : (cctxData ? [cctxData] : [])).map((transaction, index) => (
                          <div key={transaction.transactionHash || index} className={`p-3 rounded-lg border ${
                            index === 0 ? 'border-blue-500/50 bg-blue-500/5' : 'border-[#555] bg-[#1a1b1e]'
                          }`}>
                            {/* Transaction Header */}
                            <div className="flex items-center justify-between mb-2">
                              <div className="text-xs text-gray-400">
                                {index === 0 && 'Latest Transaction'}
                                {index > 0 && `Transaction #${index + 1}`}
                              </div>
                              <div className={`px-2 py-1 rounded text-xs font-medium ${
                                transaction.status === 'Aborted' 
                                  ? 'bg-red-500/20 text-red-400 border border-red-500/30' 
                                  : transaction.status === 'OutboundMined'
                                  ? 'bg-green-500/20 text-green-400 border border-green-500/30'
                                  : 'bg-yellow-500/20 text-yellow-400 border border-yellow-500/30'
                              }`}>
                                {transaction.status === 'Aborted' && '❌ '}
                                {transaction.status === 'OutboundMined' && '✅ '}
                                {transaction.status !== 'Aborted' && transaction.status !== 'OutboundMined' && '⏳ '}
                                {transaction.status}
                              </div>
                            </div>
                            
                            {/* Transaction Hash */}
                            <div className="space-y-1 mb-2">
                              <div className="text-xs text-gray-400">Transaction Hash</div>
                              <div className="font-mono text-xs text-primary bg-bg-tertiary p-2 border border-border break-all overflow-hidden">
                                {transaction.transactionHash}
                              </div>
                            </div>
                            
                            {/* Status Message */}
                            {transaction.status_message && (
                              <div className="space-y-1 mb-2">
                                <div className="text-xs text-gray-400">Message</div>
                                <div className="text-xs text-orange-400 bg-orange-500/10 p-2 rounded border border-orange-500/30">
                                  {transaction.status_message}
                                </div>
                              </div>
                            )}
                            
                            {/* Chain Information */}
                            <div className="grid grid-cols-2 gap-3 mb-2">
                              <div className="space-y-1">
                                <div className="text-xs text-gray-400">From Chain</div>
                                <div className="text-xs text-white font-medium">
                                  {getChainName(transaction.sender_chain_id)}
                                </div>
                              </div>
                              <div className="space-y-1">
                                <div className="text-xs text-gray-400">To Chain</div>
                                <div className="text-xs text-white font-medium">
                                  {getChainName(transaction.receiver_chainId)}
                                </div>
                              </div>
                            </div>
                            
                            {/* Outbound Transaction */}
                            {transaction.outbound_tx_hash && (
                              <div className="space-y-1 mb-2">
                                <div className="text-xs text-gray-400">Outbound Transaction</div>
                                <div className="font-mono text-xs text-purple-400 bg-purple-500/10 p-2 rounded border border-purple-500/30 break-all overflow-hidden">
                                  {transaction.outbound_tx_hash}
                                </div>
                              </div>
                            )}
                            
                            {/* Confirmation Status */}
                            {/* <div className="flex items-center justify-between mb-2">
                              <div className="text-xs text-gray-400">Confirmed on Destination</div>
                              <div className={`text-xs font-medium ${
                                transaction.confirmed_on_destination ? 'text-green-400' : 'text-red-400'
                              }`}>
                                {transaction.confirmed_on_destination ? '✅ Yes' : '❌ No'}
                              </div>
                            </div> */}
                            
                            {/* Timestamp */}
                            {transaction.blockTimestamp && (
                              <div className="pt-2 border-t border-border">
                                <div className="text-xs text-gray-500">
                                  Last updated: {(() => {
                                    const date = new Date(transaction.blockTimestamp);
                                    return isNaN(date.getTime()) ? 'Unknown' : date.toLocaleTimeString();
                                  })()}
                                </div>
                              </div>
                            )}
                          </div>
                        ))}
                      </div>
                    ) : (
                      <div className="flex flex-col items-center justify-center py-12 text-center">
                        <div className="text-4xl mb-2">🌉</div>
                        <div className="text-gray-400 text-sm">No cross-chain transactions yet</div>
                        <div className="text-gray-500 text-xs mt-2">Start a transfer to see it here</div>
                      </div>
                    )}
                  </div>
                </div>
              </div>
              
              {/* Mobile Transaction History */}
              <div className="lg:hidden w-full flex justify-center">
                {(cctxData || cctxHistory.length > 0) && (
                  <div className="w-full max-w-md mt-4">
                    <div className="p-4 space-y-3 border border-border bg-bg-card">
                      <div className="flex items-center justify-between">
                        <h3 className="text-sm font-medium text-gray-300">
                          {cctxHistory.length > 1 ? 'Your Cross-Chain Transaction History' : 'Your Latest Cross-Chain Transaction'}
                        </h3>
                        {isTrackingCctx && (
                          <div className="flex items-center gap-2 text-xs text-blue-400">
                            <div className="w-3 h-3 border-2 border-blue-400 border-t-transparent rounded-full animate-spin"></div>
                            Tracking...
                          </div>
                        )}
                      </div>
                      
                      {/* Show only latest transaction on mobile */}
                      <div className="space-y-3">
                        {(() => {
                          const latestTransaction = cctxHistory.length > 0 ? cctxHistory[0] : cctxData;
                          if (!latestTransaction) return null;
                          
                          return (
                            <div className="p-3 rounded-lg border border-blue-500/50 bg-blue-500/5">
                              {/* Transaction Header */}
                              <div className="flex items-center justify-between mb-2">
                                <div className="text-xs text-gray-400">Latest Transaction</div>
                                <div className={`px-2 py-1 rounded text-xs font-medium ${
                                  latestTransaction.status === 'Aborted' 
                                    ? 'bg-red-500/20 text-red-400 border border-red-500/30' 
                                    : latestTransaction.status === 'OutboundMined'
                                    ? 'bg-green-500/20 text-green-400 border border-green-500/30'
                                    : 'bg-yellow-500/20 text-yellow-400 border border-yellow-500/30'
                                }`}>
                                  {latestTransaction.status === 'Aborted' && '❌ '}
                                  {latestTransaction.status === 'OutboundMined' && '✅ '}
                                  {latestTransaction.status !== 'Aborted' && latestTransaction.status !== 'OutboundMined' && '⏳ '}
                                  {latestTransaction.status}
                                </div>
                              </div>
                              
                              {/* Transaction Hash */}
                              <div className="space-y-1 mb-2">
                                <div className="text-xs text-gray-400">Transaction Hash</div>
                                <div className="font-mono text-xs text-primary bg-bg-tertiary p-2 border border-border break-all overflow-hidden">
                                  {latestTransaction.transactionHash}
                                </div>
                              </div>
                              
                              {/* Status Message */}
                              {latestTransaction.status_message && (
                                <div className="space-y-1 mb-2">
                                  <div className="text-xs text-gray-400">Message</div>
                                  <div className="text-xs text-orange-400 bg-orange-500/10 p-2 rounded border border-orange-500/30">
                                    {latestTransaction.status_message}
                                  </div>
                                </div>
                              )}
                              
                              {/* Chain Information */}
                              <div className="grid grid-cols-2 gap-3 mb-2">
                                <div className="space-y-1">
                                  <div className="text-xs text-gray-400">From Chain</div>
                                  <div className="text-xs text-white font-medium">
                                    {getChainName(latestTransaction.sender_chain_id)}
                                  </div>
                                </div>
                                <div className="space-y-1">
                                  <div className="text-xs text-gray-400">To Chain</div>
                                  <div className="text-xs text-white font-medium">
                                    {getChainName(latestTransaction.receiver_chainId)}
                                  </div>
                                </div>
                              </div>
                              
                              {/* Confirmation Status */}
                              <div className="flex items-center justify-between mb-2">
                                <div className="text-xs text-gray-400">Confirmed on Destination</div>
                                <div className={`text-xs font-medium ${
                                  latestTransaction.confirmed_on_destination ? 'text-green-400' : 'text-red-400'
                                }`}>
                                  {latestTransaction.confirmed_on_destination ? '✅ Yes' : '❌ No'}
                                </div>
                              </div>
                              
                              {/* Timestamp */}
                              {latestTransaction.blockTimestamp && (
                                <div className="pt-2 border-t border-border">
                                  <div className="text-xs text-gray-500">
                                    Last updated: {(() => {
                                      const date = new Date(latestTransaction.blockTimestamp);
                                      return isNaN(date.getTime()) ? 'Unknown' : date.toLocaleTimeString();
                                    })()}
                                  </div>
                                </div>
                              )}
                            </div>
                          );
                        })()}
                      </div>
                    </div>
                  </div>
                )}
              </div>
            </div>
          </TabsContent>
          <TabsContent value="history" className="p-4 lg:p-6">
            <CrosschainHistory 
              baseContractAddress={sourceChain.id === 84532 ? (sourceChainContractAddress as string || '') : (targetChainContractAddress as string || '')}
              bscContractAddress={sourceChain.id === 97 ? (sourceChainContractAddress as string || '') : (targetChainContractAddress as string || '')}
            />
          </TabsContent>
          </div>
        </div>
      </Tabs>

      {/* Transaction Dialog */}
      <TransactionDialog
        isOpen={dialogState.isOpen}
        onOpenChange={onOpenChange}
        status={dialogState.status}
        hash={dialogState.hash}
        error={dialogState.error}
        title={
          pendingApprovalTokenId !== null 
            ? "Approve NFT Success" 
            : "Cross-Chain Transfer Success"
        }
        chainId={chainId}
      />
    </div>
  );
}
