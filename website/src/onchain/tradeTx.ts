import { TRADE_ABI, Fluxus_EVM_CROSS_CHAIN_ABI } from '@/src/contract';

export interface BuyTxParams {
  tradeContractAddress: `0x${string}` | string;
  collectionAddress: `0x${string}` | string;
  tokenId: number;
  buyPriceWei: bigint;
}

export function buildBuyTx({ tradeContractAddress, collectionAddress, tokenId, buyPriceWei }: BuyTxParams) {
  return {
    address: tradeContractAddress as any,
    abi: TRADE_ABI,
    functionName: 'buy' as const,
    args: [collectionAddress, tokenId],
    value: buyPriceWei,
  };
}

export interface QuickBuyTxParams {
  tradeContractAddress: `0x${string}` | string;
  collectionAddress: `0x${string}` | string;
  buyPriceWei: bigint;
}

export function buildQuickBuyTx({ tradeContractAddress, collectionAddress, buyPriceWei }: QuickBuyTxParams) {
  return {
    address: tradeContractAddress as any,
    abi: TRADE_ABI,
    functionName: 'quickBuy' as const,
    args: [collectionAddress],
    value: buyPriceWei,
  };
}

export function buildBulkQuickBuyTx({ tradeContractAddress, collectionAddress, amount, totalCostWei }: { tradeContractAddress: `0x${string}` | string; collectionAddress: `0x${string}` | string; amount: number; totalCostWei: bigint; }) {
  return {
    address: tradeContractAddress as any,
    abi: TRADE_ABI,
    functionName: 'bulkQuickBuy' as const,
    args: [collectionAddress, amount],
    value: totalCostWei,
  };
}

export function buildBulkBuyTx({ tradeContractAddress, collectionAddress, tokenIds, totalCostWei }: { tradeContractAddress: `0x${string}` | string; collectionAddress: `0x${string}` | string; tokenIds: number[]; totalCostWei: bigint; }) {
  return {
    address: tradeContractAddress as any,
    abi: TRADE_ABI,
    functionName: 'bulkBuy' as const,
    args: [collectionAddress, tokenIds],
    value: totalCostWei,
  };
}

export function buildSellTx({ tradeContractAddress, collectionAddress, tokenId }: { tradeContractAddress: `0x${string}` | string; collectionAddress: `0x${string}` | string; tokenId: number; }) {
  return {
    address: tradeContractAddress as any,
    abi: TRADE_ABI,
    functionName: 'sell' as const,
    args: [collectionAddress, tokenId],
  };
}

export function buildBulkSellTx({ tradeContractAddress, collectionAddress, tokenIds }: { tradeContractAddress: `0x${string}` | string; collectionAddress: `0x${string}` | string; tokenIds: number[]; }) {
  return {
    address: tradeContractAddress as any,
    abi: TRADE_ABI,
    functionName: 'bulkSell' as const,
    args: [collectionAddress, tokenIds],
  };
}

export function buildSetApprovalForAllTx({ nftContractAddress, operatorAddress, approved }: { nftContractAddress: `0x${string}` | string; operatorAddress: `0x${string}` | string; approved: boolean; }) {
  return {
    address: nftContractAddress as any,
    abi: Fluxus_EVM_CROSS_CHAIN_ABI,
    functionName: 'setApprovalForAll' as const,
    args: [operatorAddress, approved],
  };
}

export function buildTransferCrossChainTx({ 
  tradeContractAddress, 
  fluxusContractAddress, 
  tokenId, 
  receiver, 
  destination,
  gasFee,
}: { 
  tradeContractAddress: `0x${string}` | string; 
  fluxusContractAddress: `0x${string}` | string; 
  tokenId: number; 
  receiver: `0x${string}` | string; 
  destination: `0x${string}` | string; 
  gasFee: bigint;
}) {
  return {
    address: tradeContractAddress as any,
    abi: TRADE_ABI,
    functionName: 'transferCrossChain' as const,
    args: [fluxusContractAddress, tokenId, receiver, destination],
    value: gasFee, // May need to be adjusted if gas fees are required
  };
}


