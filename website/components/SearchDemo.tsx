"use client";

import { useState } from "react";
import { useSearch } from "@/hooks/useSearch";
import SearchResults from "./SearchResults";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

/**
 * Demo component to showcase the search functionality
 * This can be used for testing or as a standalone search component
 */
export default function SearchDemo() {
  const [chainId, setChainId] = useState(84532);
  const {
    query,
    setQuery,
    results,
    loading,
    error,
    isOpen,
    clearSearch,
    closeSearch,
  } = useSearch(chainId);

  return (
    <div className="mx-auto max-w-2xl space-y-6 p-6">
      <div className="text-center">
        <h1 className="mb-2 font-heading text-[34px] leading-none text-[color:var(--text-primary)]">
          Fluxus Collection Search Demo
        </h1>
        <p className="text-[color:var(--text-secondary)]">
          Search for NFT collections by name or contract address
        </p>
      </div>

      <div className="space-y-4">
        <div className="flex gap-4 items-center">
          <label className="font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)]">Chain ID:</label>
          <select
            value={chainId}
            onChange={(e) => setChainId(Number(e.target.value))}
            className="border border-black/10 bg-[color:var(--bg-surface)] px-3 py-2 text-sm text-[color:var(--text-primary)]"
          >
            <option value={84532}>Base Sepolia (84532)</option>
            <option value={97}>BSC Testnet (97)</option>
          </select>
        </div>

        <div className="relative">
          <Input
            type="text"
            placeholder="Search collections by name or address..."
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            className="w-full border border-black/10 bg-[color:var(--bg-surface)] py-3 pl-10 pr-4 text-[color:var(--text-primary)] placeholder:text-[color:var(--text-muted)] focus-visible:ring-2 focus-visible:ring-[color:var(--color-primary)] focus-visible:border-transparent"
          />
          <div className="absolute inset-y-0 left-0 flex items-center pl-3">
            <svg
              className="h-5 w-5 text-[color:var(--text-muted)]"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              />
            </svg>
          </div>
        </div>

        <div className="flex gap-2">
          <Button
            onClick={clearSearch}
            variant="outline"
            className="text-sm"
          >
            Clear
          </Button>
          <Button
            onClick={() => setQuery("test")}
            variant="outline"
            className="text-sm"
          >
            Test Search
          </Button>
        </div>

        <div className="text-sm text-[color:var(--text-secondary)]">
          <p>Try searching for:</p>
          <ul className="list-disc list-inside mt-1 space-y-1">
            <li>Collection names (e.g., "Test Collection")</li>
            <li>Contract addresses (e.g., "0x123...")</li>
            <li>Partial matches work too!</li>
          </ul>
        </div>
      </div>

      <SearchResults
        results={results}
        loading={loading}
        error={error}
        isOpen={isOpen}
        onClose={closeSearch}
      />

      {error && (
        <div className="border border-red-200 bg-red-50 p-4">
          <h3 className="mb-2 text-red-600">Error</h3>
          <p className="text-sm text-red-500">{error}</p>
        </div>
      )}

      <div className="border border-black/10 bg-[color:var(--bg-surface)] p-4">
        <h3 className="mb-2 font-heading text-[24px] leading-none text-[color:var(--text-primary)]">Search Status</h3>
        <div className="space-y-1 text-sm text-[color:var(--text-secondary)]">
          <p>Query: "{query}"</p>
          <p>Loading: {loading ? "Yes" : "No"}</p>
          <p>Results: {results.length}</p>
          <p>Open: {isOpen ? "Yes" : "No"}</p>
          <p>Chain ID: {chainId}</p>
        </div>
      </div>
    </div>
  );
}
