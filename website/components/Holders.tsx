'use client'

import React, { useState } from "react";
import { CopyIcon, CheckIcon } from '@radix-ui/react-icons';
import Pagination from '@/components/ui/pagination';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { useCollectionHolders, CollectionHolder } from '@/src/api/hooks';
import { sliceAddress } from "@/src/utils";
import { useChainId } from "wagmi";

const PAGE_SIZE = 10;

function formatTime(iso: string) {
  if (!iso) return '-';
  const date = new Date(iso);
  const now = new Date();
  const diff = (now.getTime() - date.getTime()) / 1000;
  if (diff < 60) return `${Math.floor(diff)}s ago`;
  if (diff < 3600) return `${Math.floor(diff / 60)}m ago`;
  if (diff < 86400) return `${Math.floor(diff / 3600)}h ago`;
  return date.toLocaleDateString();
}

function getHolderTier(count: number) {
  if (count >= 100) return { tier: "Whale", badge: "bg-purple-500 text-black" };
  if (count >= 10) return { tier: "Dolphin", badge: "bg-blue-500 text-black" };
  if (count >= 5) return { tier: "Fish", badge: "bg-[#16A34A] text-black" };
  return { tier: "Shrimp", badge: "bg-fluxus-primary/15 text-fluxus-primary" };
}

const getFirstAcquiredAt = (row: CollectionHolder) => {
  const anyRow = row as unknown as Record<string, unknown>;
  return (anyRow['first_buy_at'] as string | undefined) ?? row.firstAcquiredAt ?? '';
};

const getLastAcquiredAt = (row: CollectionHolder) => {
  const anyRow = row as unknown as Record<string, unknown>;
  return (anyRow['last_buy_at'] as string | undefined) ?? row.lastAcquiredAt ?? '';
};

export default function Holders({ collectionAddress }: { collectionAddress: string }) {
  const [page, setPage] = useState(1);
  const [copiedIdx, setCopiedIdx] = useState<number | null>(null);
  const chainId = useChainId();
  const { data: holders = [], pagination, isLoading } = useCollectionHolders(chainId, collectionAddress, PAGE_SIZE, page);
  const totalPages = pagination ? Math.ceil(Number(pagination.total) / PAGE_SIZE) : 1;
  const pagedData = holders;

  const handleCopy = async (address: string, idx: number) => {
    await navigator.clipboard.writeText(address);
    setCopiedIdx(idx);
    setTimeout(() => setCopiedIdx(null), 1200);
  };

  const renderRows = (rows: CollectionHolder[], isMobile = false) => {
    if (isLoading) {
      return (
        <div className="flex flex-col items-center justify-center py-12 text-secondary gap-3">
          <div className="w-8 h-8 border-2 border-fluxus-primary border-t-transparent rounded-full animate-spin" />
          <span>Loading holders...</span>
        </div>
      );
    }

    if (rows.length === 0) {
      return (
        <div className="flex flex-col items-center justify-center py-12 text-secondary gap-3">
          <div className="text-4xl">👥</div>
          <span>No holders found</span>
        </div>
      );
    }

    if (isMobile) {
      return rows.map((row: CollectionHolder, idx: number) => {
        const tier = getHolderTier(row.nft_count);
        const firstAcquiredAt = getFirstAcquiredAt(row);
        const lastAcquiredAt = getLastAcquiredAt(row);
        return (
          <div key={row.owner || idx} className="border border-border bg-bg-card p-4 space-y-3">
            <div className="flex items-start justify-between">
              <div>
                <p className="text-sm font-bold text-primary">{sliceAddress(row.owner)}</p>
                <p className="text-xs text-secondary">First: {formatTime(firstAcquiredAt)}</p>
              </div>
              <button
                className="p-1 border border-border text-secondary"
                onClick={() => handleCopy(row.owner, idx)}
                aria-label="Copy address"
              >
                {copiedIdx === idx ? <CheckIcon className="w-4 h-4 text-fluxus-primary" /> : <CopyIcon className="w-4 h-4" />}
              </button>
            </div>
            <div className="grid grid-cols-2 gap-3 text-sm">
              <div>
                <p className="text-xs text-secondary uppercase tracking-[0.2em] mb-1">Holdings</p>
                <p className="text-primary font-bold">{row.nft_count}</p>
              </div>
              <div>
                <p className="text-xs text-secondary uppercase tracking-[0.2em] mb-1">Tier</p>
                <span className={`inline-flex px-3 py-1 text-xs font-bold uppercase tracking-[0.2em] ${tier.badge}`}>
                  {tier.tier}
                </span>
              </div>
            </div>
            <div className="flex items-center justify-between text-xs text-secondary">
              <span>Last: {formatTime(lastAcquiredAt)}</span>
            </div>
          </div>
        );
      });
    }

    return (
      <Table>
        <TableHeader>
          <TableRow className="border-border">
            <TableHead className="text-primary">Holder</TableHead>
            <TableHead className="text-primary text-center">Holdings</TableHead>
            <TableHead className="text-primary text-center">Tier</TableHead>
            <TableHead className="text-primary text-right">First Acquired</TableHead>
            <TableHead className="text-primary text-right">Last Acquired</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {rows.map((row: CollectionHolder, idx: number) => {
            const tier = getHolderTier(row.nft_count);
            const firstAcquiredAt = getFirstAcquiredAt(row);
            const lastAcquiredAt = getLastAcquiredAt(row);
            return (
              <TableRow key={row.owner || idx} className="border-border">
                <TableCell>
                  <div className="flex items-center gap-3">
                    <span className="font-mono text-secondary text-sm">{sliceAddress(row.owner)}</span>
                    <button
                      className="p-1 border border-border text-secondary"
                      onClick={(e) => { e.stopPropagation(); handleCopy(row.owner, idx); }}
                      aria-label="Copy address"
                    >
                      {copiedIdx === idx ? <CheckIcon className="w-4 h-4 text-fluxus-primary" /> : <CopyIcon className="w-4 h-4" />}
                    </button>
                  </div>
                </TableCell>
                <TableCell className="text-center text-primary font-bold">
                  {row.nft_count}
                </TableCell>
                <TableCell className="text-center">
                  <span className={`inline-flex px-3 py-1 text-xs font-bold uppercase tracking-[0.2em] ${tier.badge}`}>
                    {tier.tier}
                  </span>
                </TableCell>
                <TableCell className="text-right text-secondary text-sm">
                  {formatTime(firstAcquiredAt)}
                </TableCell>
                <TableCell className="text-right text-secondary text-sm">
                  {formatTime(lastAcquiredAt)}
                </TableCell>
              </TableRow>
            );
          })}
        </TableBody>
      </Table>
    );
  };

  return (
    <div className="w-full px-2 sm:px-4 py-6 pb-24 space-y-6">
      <div className="flex flex-col gap-1">
        <h2 className="text-sm font-bold uppercase tracking-[0.2em] text-secondary">Top Holders</h2>
        <p className="text-xs text-secondary">Wallets holding this collection.</p>
      </div>

      <div className="hidden sm:block border border-border bg-bg-card">
        {renderRows(pagedData)}
      </div>

      <div className="sm:hidden space-y-3">
        {renderRows(pagedData, true)}
      </div>

      <div className="flex justify-center">
        <Pagination totalPages={totalPages} offset={page - 1} onPageChange={p => setPage(p)} />
      </div>
    </div>
  );
}
