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
  if (count >= 100) return { tier: "Whale", badge: "border border-black/10 bg-[color:var(--fg-strong)] text-white" };
  if (count >= 10) return { tier: "Dolphin", badge: "border border-black/10 bg-[color:var(--bg-muted)] text-[color:var(--text-primary)]" };
  if (count >= 5) return { tier: "Fish", badge: "border border-transparent bg-[color:var(--color-primary)] text-white" };
  return { tier: "Shrimp", badge: "border border-black/10 bg-[color:var(--bg-muted)] text-[color:var(--color-primary)]" };
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
        <div className="flex flex-col items-center justify-center gap-3 py-12 text-[color:var(--text-muted)]">
          <div className="h-8 w-8 animate-spin border-2 border-[color:var(--color-primary)] border-t-transparent" />
          <span>Loading holders...</span>
        </div>
      );
    }

    if (rows.length === 0) {
      return (
        <div className="flex flex-col items-center justify-center gap-3 py-12 text-[color:var(--text-muted)]">
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
          <div key={row.owner || idx} className="space-y-3 border border-black/10 bg-[color:var(--bg-surface)] p-4">
            <div className="flex items-start justify-between">
              <div>
                <p className="font-heading text-[22px] leading-none text-[color:var(--text-primary)]">{sliceAddress(row.owner)}</p>
                <p className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">First: {formatTime(firstAcquiredAt)}</p>
              </div>
              <button
                className="border border-black/10 p-1 text-[color:var(--text-muted)]"
                onClick={() => handleCopy(row.owner, idx)}
                aria-label="Copy address"
              >
                {copiedIdx === idx ? <CheckIcon className="h-4 w-4 text-[color:var(--color-primary)]" /> : <CopyIcon className="h-4 w-4" />}
              </button>
            </div>
            <div className="grid grid-cols-2 gap-3 text-sm">
              <div>
                <p className="mb-1 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Holdings</p>
                <p className="font-heading text-[22px] leading-none text-[color:var(--text-primary)]">{row.nft_count}</p>
              </div>
              <div>
                <p className="mb-1 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Tier</p>
                <span className={`inline-flex px-3 py-1 font-primary text-[10px] uppercase tracking-[0.16em] ${tier.badge}`}>
                  {tier.tier}
                </span>
              </div>
            </div>
            <div className="flex items-center justify-between font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">
              <span>Last: {formatTime(lastAcquiredAt)}</span>
            </div>
          </div>
        );
      });
    }

    return (
      <Table>
        <TableHeader>
          <TableRow className="border-black/10">
            <TableHead>Holder</TableHead>
            <TableHead className="text-center">Holdings</TableHead>
            <TableHead className="text-center">Tier</TableHead>
            <TableHead className="text-right">First Acquired</TableHead>
            <TableHead className="text-right">Last Acquired</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {rows.map((row: CollectionHolder, idx: number) => {
            const tier = getHolderTier(row.nft_count);
            const firstAcquiredAt = getFirstAcquiredAt(row);
            const lastAcquiredAt = getLastAcquiredAt(row);
            return (
              <TableRow key={row.owner || idx} className="border-black/10">
                <TableCell>
                  <div className="flex items-center gap-3">
                    <span className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">{sliceAddress(row.owner)}</span>
                    <button
                      className="border border-black/10 p-1 text-[color:var(--text-muted)]"
                      onClick={(e) => { e.stopPropagation(); handleCopy(row.owner, idx); }}
                      aria-label="Copy address"
                    >
                      {copiedIdx === idx ? <CheckIcon className="h-4 w-4 text-[color:var(--color-primary)]" /> : <CopyIcon className="h-4 w-4" />}
                    </button>
                  </div>
                </TableCell>
                <TableCell className="text-center font-heading text-[20px] leading-none text-[color:var(--text-primary)]">
                  {row.nft_count}
                </TableCell>
                <TableCell className="text-center">
                  <span className={`inline-flex px-3 py-1 font-primary text-[10px] uppercase tracking-[0.16em] ${tier.badge}`}>
                    {tier.tier}
                  </span>
                </TableCell>
                <TableCell className="text-right font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                  {formatTime(firstAcquiredAt)}
                </TableCell>
                <TableCell className="text-right font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
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
    <div className="w-full space-y-6 px-2 py-6 pb-24 sm:px-4">
      <div className="flex flex-col gap-1">
        <h2 className="flux-kicker mb-2">Top Holders</h2>
        <p className="text-sm text-[color:var(--text-secondary)]">Wallets holding this collection.</p>
      </div>

      <div className="hidden border border-black/10 bg-[color:var(--bg-surface)] sm:block">
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
