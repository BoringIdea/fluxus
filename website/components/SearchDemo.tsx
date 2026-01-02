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
    <div className="max-w-2xl mx-auto p-6 space-y-6">
      <div className="text-center">
        <h1 className="text-2xl font-bold text-white mb-2">
          Fluxus Collection Search Demo
        </h1>
        <p className="text-gray-400">
          Search for NFT collections by name or contract address
        </p>
      </div>

      <div className="space-y-4">
        <div className="flex gap-4 items-center">
          <label className="text-white text-sm">Chain ID:</label>
          <select
            value={chainId}
            onChange={(e) => setChainId(Number(e.target.value))}
            className="bg-[#0f1115] border border-[#2a2f37] rounded-md px-3 py-2 text-white text-sm"
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
            className="w-full bg-[#0f1115]/90 border border-[#2a2f37] rounded-md pl-10 pr-4 py-3 text-white placeholder:text-gray-500 focus-visible:ring-2 focus-visible:ring-[#16A34A] focus-visible:border-transparent"
          />
          <div className="absolute inset-y-0 left-0 flex items-center pl-3">
            <svg
              className="w-5 h-5 text-gray-500"
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

        <div className="text-sm text-gray-400">
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
        <div className="bg-red-900/20 border border-red-500/50 rounded-md p-4">
          <h3 className="text-red-400 font-medium mb-2">Error</h3>
          <p className="text-red-300 text-sm">{error}</p>
        </div>
      )}

      <div className="bg-[#0f1115] border border-[#2a2f37] rounded-md p-4">
        <h3 className="text-white font-medium mb-2">Search Status</h3>
        <div className="text-sm text-gray-400 space-y-1">
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
