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
    <div className="absolute top-full left-0 right-0 mt-1 bg-[#0f1115] border border-[#2a2f37] rounded-md shadow-lg z-50 max-h-96 overflow-y-auto">
      {loading && (
        <div className="p-4 text-center text-gray-400">
          <div className="animate-spin w-5 h-5 border-2 border-[#16A34A] border-t-transparent rounded-full mx-auto mb-2"></div>
          Searching...
        </div>
      )}

      {error && (
        <div className="p-4 text-center text-red-400">
          <Search className="w-5 h-5 mx-auto mb-2" />
          {error}
        </div>
      )}

      {!loading && !error && results.length === 0 && (
        <div className="p-4 text-center text-gray-400">
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
              className="block px-4 py-3 hover:bg-[#1a1d24] transition-colors border-b border-[#2a2f37] last:border-b-0"
            >
              <div className="flex items-start gap-3">
                <div className="w-10 h-10 bg-[#2a2f37] rounded-md flex items-center justify-center flex-shrink-0">
                  {collection.base_uri ? (
                    <img
                      src={collection.base_uri}
                      alt={collection.name}
                      className="w-full h-full object-cover rounded-md"
                      onError={(e) => {
                        const target = e.target as HTMLImageElement;
                        target.style.display = 'none';
                        target.nextElementSibling?.classList.remove('hidden');
                      }}
                    />
                  ) : null}
                  <div className={`w-full h-full flex items-center justify-center text-xs font-bold text-[#16A34A] ${collection.base_uri ? 'hidden' : ''}`}>
                    {collection.symbol?.slice(0, 2).toUpperCase() || 'N/A'}
                  </div>
                </div>

                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 mb-1">
                    <h3 className="font-medium text-white truncate">
                      {collection.name}
                    </h3>
                    <span className="text-xs text-gray-400 bg-[#2a2f37] px-2 py-0.5 rounded">
                      {collection.symbol}
                    </span>
                  </div>

                  <div className="flex items-center gap-4 text-xs text-gray-400">
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
                    <span className="text-xs text-gray-500 font-mono">
                      {collection.address.slice(0, 6)}...{collection.address.slice(-4)}
                    </span>
                    <ExternalLink className="w-3 h-3 text-gray-500" />
                  </div>
                </div>

                <div className="text-right flex-shrink-0">
                  <div className="text-xs text-gray-400">
                    Vol: {formatVolume(collection.total_volume)} ETH
                  </div>
                  <div className="text-xs text-gray-500">
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
