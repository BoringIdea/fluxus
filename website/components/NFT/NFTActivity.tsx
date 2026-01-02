'use client'
import React from 'react';
import { useChainId } from 'wagmi';
import { useTokenTxs, CollectionTx } from '@/src/api';
import { getChainSymbol, sliceAddress } from '@/src/utils';

function getActionDisplay(type: number) {
  switch (type) {
    case 1: 
      return { text: 'Mint', color: 'text-blue-400', bgColor: 'bg-blue-400/10', icon: '🎨' };
    case 2: 
      return { text: 'Buy', color: 'text-green-400', bgColor: 'bg-green-400/10', icon: '🛒' };
    case 3: 
      return { text: 'Sell', color: 'text-red-400', bgColor: 'bg-red-400/10', icon: '💰' };
    case 4: 
      return { text: 'Bulk Buy', color: 'text-green-400', bgColor: 'bg-green-400/10', icon: '📦' };
    case 5: 
      return { text: 'Bulk Sell', color: 'text-red-400', bgColor: 'bg-red-400/10', icon: '📦' };
    case 6: 
      return { text: 'Bulk Mint', color: 'text-blue-400', bgColor: 'bg-blue-400/10', icon: '🎨' };
    default: 
      return { text: 'Other', color: 'text-gray-400', bgColor: 'bg-gray-400/10', icon: '❓' };
  }
}

function formatPriceWeiToSymbol(p: string, chainId: number) {
  try {
    const v = BigInt(p || '0');
    const num = Number(v) / 1e18;
    return `${num.toFixed(4).replace(/\.0+$/, '')} ${getChainSymbol(chainId)}`;
  } catch {
    return `0 ${getChainSymbol(chainId)}`;
  }
}

function formatTime(iso: string) {
  if (!iso) return '-';
  // backend createdAt in Activity.tsx uses seconds → unify here
  const isNumeric = /^\d+$/.test(String(iso));
  const d = isNumeric ? new Date(Number(iso) * 1000) : new Date(iso);
  if (isNaN(d.getTime())) return '';
  const now = new Date();
  const diff = (now.getTime() - d.getTime()) / 1000;
  if (diff < 60) return `${Math.floor(diff)}s ago`;
  if (diff < 3600) return `${Math.floor(diff / 60)}m ago`;
  if (diff < 86400) return `${Math.floor(diff / 3600)}h ago`;
  return d.toLocaleDateString();
}

export default function NFTActivity({ collectionAddress, tokenId }: { collectionAddress: string; tokenId: number }) {
  const chainId = useChainId();
  const { data: txs, isLoading } = useTokenTxs(chainId, collectionAddress, tokenId, 50, 1);

  return (
    <div className="rounded-xl border border-neutral-800 bg-neutral-900 p-4 h-72 overflow-y-auto">
      {/* <div className="text-sm font-semibold text-gray-300 mb-3">ACTIVITY</div> */}
      {isLoading ? (
        <div className="text-gray-500 text-sm">Loading...</div>
      ) : txs && txs.length > 0 ? (
        <div className="flex flex-col gap-2">
          {txs.map((row: CollectionTx, idx: number) => {
            const actionDisplay = getActionDisplay(row.txType);
            return (
              <div key={row.txHash || idx} className="rounded-lg border border-neutral-800 bg-neutral-800/50 p-3 text-sm text-white">
                <div className="flex items-start justify-between mb-1">
                  <div className={`inline-flex items-center gap-2 px-2.5 py-1 rounded-md text-xs font-medium ${actionDisplay.bgColor} ${actionDisplay.color}`}>
                    <span className="text-xs">{actionDisplay.icon}</span>
                    {actionDisplay.text}
                  </div>
                  <div className="text-gray-500 text-xs">{formatTime(row.createdAt.toString())}</div>
                </div>
                <div className="flex items-center justify-between mt-1">
                  <div className="text-gray-300 font-mono">{sliceAddress(row.sender)}</div>
                  <div className="text-gray-200 font-semibold">{formatPriceWeiToSymbol(row.price, chainId)}</div>
                </div>
              </div>
            );
          })}
        </div>
      ) : (
        <div className="text-gray-500 text-sm">No activity</div>
      )}
    </div>
  );
}


