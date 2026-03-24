"use client";

import { SearchResult } from "@/hooks/useSearch";
import { Search, ExternalLink, Users, TrendingUp } from "lucide-react";
import Link from "next/link";
import { formatEther } from "viem";

interface SearchResultsProps {
  results: SearchResult[];
  loading: boolean;
  error: string | null;
  isOpen: boolean;
  onClose: () => void;
}

export default function SearchResults({
  results,
  loading,
  error,
  isOpen,
  onClose,
}: SearchResultsProps) {
  if (!isOpen) return null;

  const formatPrice = (price: string) => {
    try {
      const value = parseFloat(formatEther(BigInt(price)));
      return value.toFixed(4);
    } catch {
      return "0.0000";
    }
  };

  const formatVolume = (volume: string) => {
    try {
      const value = parseFloat(formatEther(BigInt(volume)));
      if (value >= 1000000) {
        return `${(value / 1000000).toFixed(1)}M`;
      } else if (value >= 1000) {
        return `${(value / 1000).toFixed(1)}K`;
      }
      return value.toFixed(2);
    } catch {
      return "0";
    }
  };

  return (
    <div className="absolute left-0 right-0 top-full z-50 mt-2 max-h-96 overflow-y-auto border border-black/10 bg-[color:var(--bg-surface)] shadow-[0_18px_40px_rgba(17,24,39,0.08)]">
      {loading && (
        <div className="p-4 text-center font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">
          <div className="mx-auto mb-2 h-5 w-5 animate-spin border-2 border-[color:var(--color-primary)] border-t-transparent"></div>
          Searching...
        </div>
      )}

      {error && (
        <div className="p-4 text-center font-primary text-[10px] uppercase tracking-[0.18em] text-red-600">
          <Search className="w-5 h-5 mx-auto mb-2" />
          {error}
        </div>
      )}

      {!loading && !error && results.length === 0 && (
        <div className="p-4 text-center font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">
          <Search className="w-5 h-5 mx-auto mb-2" />
          No collections found
        </div>
      )}

      {!loading && !error && results.length > 0 && (
        <div className="py-2">
          {results.map((collection) => (
            <Link
              key={collection.id}
              href={`/collection/${collection.address}`}
              onClick={onClose}
              className="block border-b border-black/10 px-4 py-3 transition-colors last:border-b-0 hover:bg-[color:var(--bg-muted)]"
            >
              <div className="flex items-start gap-3">
                <div className="flex h-10 w-10 flex-shrink-0 items-center justify-center overflow-hidden border border-black/10 bg-[color:var(--bg-muted)]">
                  {collection.base_uri ? (
                    <img
                      src={collection.base_uri}
                      alt={collection.name}
                      className="h-full w-full object-cover"
                      onError={(e) => {
                        const target = e.target as HTMLImageElement;
                        target.style.display = 'none';
                        target.nextElementSibling?.classList.remove('hidden');
                      }}
                    />
                  ) : null}
                  <div className={`flex h-full w-full items-center justify-center font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--color-primary)] ${collection.base_uri ? 'hidden' : ''}`}>
                    {collection.symbol?.slice(0, 2).toUpperCase() || 'N/A'}
                  </div>
                </div>

                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 mb-1">
                    <h3 className="truncate font-heading text-[22px] leading-none text-[color:var(--text-primary)]">
                      {collection.name}
                    </h3>
                    <span className="border border-black/10 bg-[color:var(--bg-muted)] px-2 py-1 font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                      {collection.symbol}
                    </span>
                  </div>

                  <div className="flex items-center gap-4 font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                    <div className="flex items-center gap-1">
                      <TrendingUp className="w-3 h-3" />
                      <span>Floor: {formatPrice(collection.floor_price || collection.initial_price)} ETH</span>
                    </div>
                    <div className="flex items-center gap-1">
                      <Users className="w-3 h-3" />
                      <span>{collection.owners} owners</span>
                    </div>
                  </div>

                  <div className="flex items-center gap-2 mt-1">
                    <span className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                      {collection.address.slice(0, 6)}...{collection.address.slice(-4)}
                    </span>
                    <ExternalLink className="h-3 w-3 text-[color:var(--text-muted)]" />
                  </div>
                </div>

                <div className="text-right flex-shrink-0">
                  <div className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                    Vol: {formatVolume(collection.total_volume)} ETH
                  </div>
                  <div className="font-primary text-[10px] uppercase tracking-[0.14em] text-[color:var(--text-muted)]">
                    {collection.current_supply}/{collection.max_supply}
                  </div>
                </div>
              </div>
            </Link>
          ))}
        </div>
      )}
    </div>
  );
}
