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
      return { text: "Mint", badge: "bg-blue-500 text-black" };
    case 2:
      return { text: "Buy", badge: "bg-[#16A34A] text-black" };
    case 3:
      return { text: "Sell", badge: "bg-red-500 text-black" };
    case 4:
      return { text: "Bulk Buy", badge: "bg-[#16A34A] text-black" };
    case 5:
      return { text: "Bulk Sell", badge: "bg-red-500 text-black" };
    case 6:
      return { text: "Bulk Mint", badge: "bg-blue-500 text-black" };
    default:
      return { text: "Other", badge: "bg-secondary text-black" };
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
        className={`flex items-center justify-center gap-2 border border-border px-3 py-2 text-xs font-bold uppercase tracking-[0.2em] ${filter !== "All" ? 'text-fluxus-primary bg-bg-tertiary' : 'text-secondary'}`}
        onClick={() => setFilterOpen(prev => !prev)}
        aria-expanded={filterOpen}
      >
        <MixerHorizontalIcon className="w-4 h-4" />
        Filter
      </button>
      {filterOpen && (
        <div className="absolute right-0 mt-2 w-64 border border-border bg-bg-card p-4 z-50 space-y-3">
          <p className="text-sm font-bold text-primary">Filter Actions</p>
          <div className="space-y-2">
            {actionOptions.map(option => (
              <label key={option} className="flex items-center gap-3 text-sm text-secondary cursor-pointer select-none">
                <input
                  type="checkbox"
                  checked={checkedActions.includes(option)}
                  onChange={() => handleCheck(option)}
                  className="accent-fluxus-primary h-4 w-4 border border-border bg-bg-tertiary"
                />
                <span>{option}</span>
              </label>
            ))}
          </div>
          <div className="flex justify-end gap-2 text-xs">
            <button className="px-3 py-2 text-secondary border border-border" onClick={resetFilter}>
              Reset
            </button>
            <button className="px-3 py-2 bg-fluxus-primary text-black" onClick={applyFilter}>
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
        <div className="flex flex-col items-center justify-center py-12 text-secondary gap-3">
          <div className="w-8 h-8 border-2 border-fluxus-primary border-t-transparent rounded-full animate-spin" />
          <span>Loading transactions...</span>
        </div>
      );
    }
    if (rows.length === 0) {
      return (
        <div className="flex flex-col items-center justify-center py-12 text-secondary gap-3">
          <div className="text-4xl">📊</div>
          <span>No transactions found</span>
        </div>
      );
    }

    if (isMobile) {
      return rows.map((row: CollectionTx, idx: number) => {
        const actionDisplay = getActionDisplay(row.txType);
        return (
          <div key={row.txHash || idx} className="border border-border bg-bg-card p-4 space-y-3">
            <div className="flex items-start justify-between">
              <span className={`px-3 py-1 text-xs font-bold uppercase tracking-[0.2em] ${actionDisplay.badge}`}>
                {actionDisplay.text}
              </span>
              <span className="text-xs text-secondary">{formatTime(row.createdAt.toString())}</span>
            </div>
            <div className="grid grid-cols-2 gap-3 text-sm">
              <div>
                <p className="text-xs text-secondary uppercase tracking-[0.2em] mb-1">Items</p>
                <p className="font-mono text-primary">
                  {Array.isArray(row.tokenIds) ? row.tokenIds.join(', ') : row.tokenIds}
                </p>
              </div>
              <div>
                <p className="text-xs text-secondary uppercase tracking-[0.2em] mb-1">Price</p>
                <p className="font-bold text-primary">{formatPrice(row.price, chainId)}</p>
              </div>
            </div>
            <div className="flex items-center justify-between text-sm">
              <span className="font-mono text-secondary">{sliceAddress(row.sender)}</span>
              <button
                className="p-1 border border-border text-secondary"
                onClick={() => handleCopy(row.sender, idx)}
                aria-label="Copy address"
              >
                {copiedIdx === idx ? <CheckIcon className="w-4 h-4 text-fluxus-primary" /> : <CopyIcon className="w-4 h-4" />}
              </button>
            </div>
          </div>
        );
      });
    }

    return (
      <Table>
        <TableHeader>
          <TableRow className="border-border">
            <TableHead className="text-primary">Action</TableHead>
            <TableHead className="text-primary text-center">Items</TableHead>
            <TableHead className="text-primary text-right">Price</TableHead>
            <TableHead className="text-primary text-center">Wallet</TableHead>
            <TableHead className="text-primary text-right">Time</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {rows.map((row: CollectionTx, idx: number) => {
            const actionDisplay = getActionDisplay(row.txType);
            return (
              <TableRow key={row.txHash || idx} className="border-border">
                <TableCell>
                  <span className={`inline-flex px-3 py-1 text-xs font-bold uppercase tracking-[0.2em] ${actionDisplay.badge}`}>
                    {actionDisplay.text}
                  </span>
                </TableCell>
                <TableCell className="text-center font-mono text-secondary">
                  {Array.isArray(row.tokenIds) ? row.tokenIds.join(', ') : row.tokenIds}
                </TableCell>
                <TableCell className="text-right font-bold text-primary">
                  {formatPrice(row.price, chainId)}
                </TableCell>
                <TableCell className="text-center">
                  <div className="flex items-center justify-center gap-2">
                    <span className="font-mono text-secondary text-sm">{sliceAddress(row.sender)}</span>
                    <button
                      className="p-1 border border-border text-secondary"
                      onClick={(e) => { e.stopPropagation(); handleCopy(row.sender, idx); }}
                      aria-label="Copy address"
                    >
                      {copiedIdx === idx ? <CheckIcon className="w-4 h-4 text-fluxus-primary" /> : <CopyIcon className="w-4 h-4" />}
                    </button>
                  </div>
                </TableCell>
                <TableCell className="text-right text-secondary text-sm">
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
    <div className="w-full px-2 sm:px-4 py-6 pb-24 space-y-6">
      <div className="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-3">
        <div>
          <h2 className="text-sm font-bold uppercase tracking-[0.2em] text-secondary">Activity</h2>
          <p className="text-xs text-secondary">Latest on-chain interactions for this collection.</p>
        </div>
        <div className="w-full sm:w-auto flex justify-end">
          {filterButton}
        </div>
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
