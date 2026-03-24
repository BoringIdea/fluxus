'use client'
import React from 'react';
import { useChainId } from 'wagmi';
import { useTokenTxs, CollectionTx } from '@/src/api';
import { getChainSymbol, sliceAddress } from '@/src/utils';

function getActionDisplay(type: number) {
  switch (type) {
    case 1: 
      return { text: 'Mint', color: 'text-[color:var(--text-secondary)]', bgColor: 'bg-[color:var(--bg-muted)]', icon: 'MN' };
    case 2: 
      return { text: 'Buy', color: 'text-[color:var(--color-primary-dark)]', bgColor: 'bg-[color:var(--color-primary)]/8', icon: 'BY' };
    case 3: 
      return { text: 'Sell', color: 'text-[color:var(--text-secondary)]', bgColor: 'bg-[color:var(--bg-muted)]', icon: 'SL' };
    case 4: 
      return { text: 'Bulk Buy', color: 'text-[color:var(--color-primary-dark)]', bgColor: 'bg-[color:var(--color-primary)]/8', icon: 'BB' };
    case 5: 
      return { text: 'Bulk Sell', color: 'text-[color:var(--text-secondary)]', bgColor: 'bg-[color:var(--bg-muted)]', icon: 'BS' };
    case 6: 
      return { text: 'Bulk Mint', color: 'text-[color:var(--text-secondary)]', bgColor: 'bg-[color:var(--bg-muted)]', icon: 'BM' };
    default: 
      return { text: 'Other', color: 'text-[color:var(--text-muted)]', bgColor: 'bg-[color:var(--bg-muted)]', icon: 'OT' };
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
    <div className="h-72 overflow-y-auto border border-black/10 bg-[color:var(--bg-surface)] p-4">
      {isLoading ? (
        <div className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Loading...</div>
      ) : txs && txs.length > 0 ? (
        <div className="flex flex-col gap-2">
          {txs.map((row: CollectionTx, idx: number) => {
            const actionDisplay = getActionDisplay(row.txType);
            return (
              <div key={row.txHash || idx} className="border border-black/10 bg-[color:var(--bg-muted)] p-3">
                <div className="flex items-start justify-between mb-1">
                  <div className={`inline-flex items-center gap-2 px-2.5 py-1 font-primary text-[10px] uppercase tracking-[0.16em] ${actionDisplay.bgColor} ${actionDisplay.color}`}>
                    <span className="text-[10px]">{actionDisplay.icon}</span>
                    {actionDisplay.text}
                  </div>
                  <div className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">{formatTime(row.createdAt.toString())}</div>
                </div>
                <div className="flex items-center justify-between mt-1">
                  <div className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">{sliceAddress(row.sender)}</div>
                  <div className="font-serif text-[16px] font-semibold leading-none tracking-[-0.01em] text-[color:var(--text-secondary)]">{formatPriceWeiToSymbol(row.price, chainId)}</div>
                </div>
              </div>
            );
          })}
        </div>
      ) : (
        <div className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">No activity</div>
      )}
    </div>
  );
}
