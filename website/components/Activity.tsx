'use client'

import React, { useEffect, useMemo, useRef, useState } from "react";
import { MixerHorizontalIcon, CopyIcon, CheckIcon } from '@radix-ui/react-icons';
import Pagination from '@/components/ui/pagination';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { useCollectionTxs, CollectionTx } from '@/src/api/hooks';
import { sliceAddress, formatNumberWithMaxDecimals, getChainSymbol } from "@/src/utils";
import { useChainId } from "wagmi";

const actionOptions = ["All", "Sell", "Buy", "Mint"];
const PAGE_SIZE = 10;

function getActionDisplay(type: number) {
  switch (type) {
    case 1:
      return { text: "Mint", badge: "border border-black/10 bg-[color:var(--bg-muted)] text-[color:var(--text-secondary)]" };
    case 2:
      return { text: "Buy", badge: "border border-[color:var(--color-primary)]/20 bg-[color:var(--color-primary)]/8 text-[color:var(--color-primary-dark)]" };
    case 3:
      return { text: "Sell", badge: "border border-black/10 bg-[color:var(--bg-muted)] text-[color:var(--text-secondary)]" };
    case 4:
      return { text: "Bulk Buy", badge: "border border-[color:var(--color-primary)]/20 bg-[color:var(--color-primary)]/8 text-[color:var(--color-primary-dark)]" };
    case 5:
      return { text: "Bulk Sell", badge: "border border-black/10 bg-[color:var(--bg-muted)] text-[color:var(--text-secondary)]" };
    case 6:
      return { text: "Bulk Mint", badge: "border border-black/10 bg-[color:var(--bg-muted)] text-[color:var(--text-secondary)]" };
    default:
      return { text: "Other", badge: "border border-black/10 bg-[color:var(--bg-muted)] text-[color:var(--text-muted)]" };
  }
}

function formatPrice(price: string, chainId: number) {
  if (!price) return '-';
  const ethValue = (Number(price) / 1e18).toString();
  const formatted = formatNumberWithMaxDecimals(ethValue);
  const symbol = getChainSymbol(chainId);
  return `${formatted} ${symbol}`;
}

function formatTime(iso: string) {
  if (!iso) return '-';
  const date = new Date(Number(iso) * 1000);
  const now = new Date();
  const diff = (now.getTime() - date.getTime()) / 1000;
  if (diff < 60) return `${Math.floor(diff)}s ago`;
  if (diff < 3600) return `${Math.floor(diff / 60)}m ago`;
  if (diff < 86400) return `${Math.floor(diff / 3600)}h ago`;
  return date.toLocaleDateString();
}

export default function Activity({ collectionAddress }: { collectionAddress: string }) {
  const chainId = useChainId();
  const [filterOpen, setFilterOpen] = useState(false);
  const [checkedActions, setCheckedActions] = useState<string[]>(["All"]);
  const [filter, setFilter] = useState("All");
  const [page, setPage] = useState(1);
  const [copiedIdx, setCopiedIdx] = useState<number | null>(null);
  const filterRef = useRef<HTMLDivElement>(null);

  const { data: txs = [], pagination, isLoading } = useCollectionTxs(chainId, collectionAddress, PAGE_SIZE, page);

  const filteredData = useMemo(() => {
    if (filter === "multi" && checkedActions.length > 0) {
      return txs.filter((d: CollectionTx) => checkedActions.includes(getActionDisplay(d.txType).text));
    }
    if (filter !== "All" && filter !== "multi") {
      return txs.filter((d: CollectionTx) => getActionDisplay(d.txType).text === filter);
    }
    return txs;
  }, [txs, filter, checkedActions]);

  const totalPages = pagination ? Math.ceil(Number(pagination.total) / PAGE_SIZE) : 1;
  const pagedData = filteredData;

  useEffect(() => {
    if (!filterOpen) return;
    const handleClick = (event: MouseEvent) => {
      if (filterRef.current && !filterRef.current.contains(event.target as Node)) {
        setFilterOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClick);
    return () => document.removeEventListener("mousedown", handleClick);
  }, [filterOpen]);

  const handleCheck = (action: string) => {
    if (action === "All") {
      setCheckedActions(["All"]);
    } else {
      setCheckedActions(prev => {
        const newChecked = prev.includes(action)
          ? prev.filter(a => a !== action)
          : [...prev.filter(a => a !== "All"), action];
        return newChecked.length === 0 ? ["All"] : newChecked;
      });
    }
  };

  const applyFilter = () => {
    if (checkedActions.includes("All") || checkedActions.length === 0) {
      setFilter("All");
    } else if (checkedActions.length === 1) {
      setFilter(checkedActions[0]);
    } else {
      setFilter("multi");
    }
    setFilterOpen(false);
    setPage(1);
  };

  const resetFilter = () => {
    setCheckedActions(["All"]);
    setFilter("All");
    setFilterOpen(false);
    setPage(1);
  };

  const handleCopy = async (address: string, idx: number) => {
    await navigator.clipboard.writeText(address);
    setCopiedIdx(idx);
    setTimeout(() => setCopiedIdx(null), 1200);
  };

  const filterButton = (
    <div className="relative w-full sm:w-auto" ref={filterRef}>
      <button
        className={`flex items-center justify-center gap-2 border border-black/10 px-3 py-2 font-primary text-[10px] uppercase tracking-[0.18em] ${filter !== "All" ? 'bg-[color:var(--bg-muted)] text-[color:var(--color-primary)]' : 'text-[color:var(--text-muted)]'}`}
        onClick={() => setFilterOpen(prev => !prev)}
        aria-expanded={filterOpen}
      >
        <MixerHorizontalIcon className="w-4 h-4" />
        Filter
      </button>
      {filterOpen && (
        <div className="absolute right-0 z-50 mt-2 w-64 space-y-3 border border-black/10 bg-[color:var(--bg-surface)] p-4">
          <p className="font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-primary)]">Filter Actions</p>
          <div className="space-y-2">
            {actionOptions.map(option => (
              <label key={option} className="flex cursor-pointer select-none items-center gap-3 text-sm text-[color:var(--text-secondary)]">
                <input
                  type="checkbox"
                  checked={checkedActions.includes(option)}
                  onChange={() => handleCheck(option)}
                  className="h-4 w-4 accent-[color:var(--color-primary)]"
                />
                <span>{option}</span>
              </label>
            ))}
          </div>
          <div className="flex justify-end gap-2 font-primary text-[10px] uppercase tracking-[0.16em]">
            <button className="border border-black/10 px-3 py-2 text-[color:var(--text-muted)]" onClick={resetFilter}>
              Reset
            </button>
            <button className="border border-[color:var(--color-primary)]/20 bg-[color:var(--color-primary)]/8 px-3 py-2 text-[color:var(--color-primary-dark)]" onClick={applyFilter}>
              Apply
            </button>
          </div>
        </div>
      )}
    </div>
  );

  const renderRows = (rows: CollectionTx[], isMobile = false) => {
    if (isLoading) {
      return (
        <div className="flex flex-col items-center justify-center gap-3 py-12 text-[color:var(--text-muted)]">
          <div className="h-8 w-8 animate-spin border-2 border-[color:var(--color-primary)] border-t-transparent" />
          <span>Loading transactions...</span>
        </div>
      );
    }
    if (rows.length === 0) {
      return (
        <div className="flex flex-col items-center justify-center gap-3 py-12 text-[color:var(--text-muted)]">
          <div className="text-4xl">📊</div>
          <span>No transactions found</span>
        </div>
      );
    }

    if (isMobile) {
      return rows.map((row: CollectionTx, idx: number) => {
        const actionDisplay = getActionDisplay(row.txType);
        return (
          <div key={row.txHash || idx} className="space-y-3 border border-black/10 bg-[color:var(--bg-surface)] p-4">
            <div className="flex items-start justify-between">
              <span className={`px-3 py-1 text-xs font-bold uppercase tracking-[0.2em] ${actionDisplay.badge}`}>
                {actionDisplay.text}
              </span>
              <span className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">{formatTime(row.createdAt.toString())}</span>
            </div>
            <div className="grid grid-cols-2 gap-3 text-sm">
              <div>
                <p className="mb-1 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Items</p>
                <p className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-primary)]">
                  {Array.isArray(row.tokenIds) ? row.tokenIds.join(', ') : row.tokenIds}
                </p>
              </div>
              <div>
                <p className="mb-1 font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Price</p>
                <p className="font-serif text-[16px] font-semibold leading-none tracking-[-0.01em] text-[color:var(--text-secondary)]">{formatPrice(row.price, chainId)}</p>
              </div>
            </div>
            <div className="flex items-center justify-between text-sm">
              <span className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">{sliceAddress(row.sender)}</span>
              <button
                className="border border-black/10 p-1 text-[color:var(--text-muted)]"
                onClick={() => handleCopy(row.sender, idx)}
                aria-label="Copy address"
              >
                {copiedIdx === idx ? <CheckIcon className="h-4 w-4 text-[color:var(--color-primary)]" /> : <CopyIcon className="h-4 w-4" />}
              </button>
            </div>
          </div>
        );
      });
    }

    return (
      <Table>
        <TableHeader>
          <TableRow className="border-black/10">
            <TableHead>Action</TableHead>
            <TableHead className="text-center">Items</TableHead>
            <TableHead className="text-right">Price</TableHead>
            <TableHead className="text-center">Wallet</TableHead>
            <TableHead className="text-right">Time</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {rows.map((row: CollectionTx, idx: number) => {
            const actionDisplay = getActionDisplay(row.txType);
            return (
              <TableRow key={row.txHash || idx} className="border-black/10">
                <TableCell>
                  <span className={`inline-flex px-3 py-1 font-primary text-[10px] uppercase tracking-[0.16em] ${actionDisplay.badge}`}>
                    {actionDisplay.text}
                  </span>
                </TableCell>
                <TableCell className="text-center font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                  {Array.isArray(row.tokenIds) ? row.tokenIds.join(', ') : row.tokenIds}
                </TableCell>
                <TableCell className="text-right font-serif text-[16px] font-semibold leading-none tracking-[-0.01em] text-[color:var(--text-secondary)]">
                  {formatPrice(row.price, chainId)}
                </TableCell>
                <TableCell className="text-center">
                  <div className="flex items-center justify-center gap-2">
                    <span className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">{sliceAddress(row.sender)}</span>
                    <button
                      className="border border-black/10 p-1 text-[color:var(--text-muted)]"
                      onClick={(e) => { e.stopPropagation(); handleCopy(row.sender, idx); }}
                      aria-label="Copy address"
                    >
                      {copiedIdx === idx ? <CheckIcon className="h-4 w-4 text-[color:var(--color-primary)]" /> : <CopyIcon className="h-4 w-4" />}
                    </button>
                  </div>
                </TableCell>
                <TableCell className="text-right font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                  {formatTime(row.createdAt.toString())}
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
      <div className="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-3">
        <div>
          <h2 className="flux-kicker mb-2">Activity</h2>
          <p className="text-sm text-[color:var(--text-secondary)]">Latest on-chain interactions for this collection.</p>
        </div>
        <div className="w-full sm:w-auto flex justify-end">
          {filterButton}
        </div>
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
