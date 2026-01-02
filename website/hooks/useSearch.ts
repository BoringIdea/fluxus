import { useState, useEffect, useCallback } from 'react';

export interface SearchResult {
  id: string;
  address: string;
  name: string;
  symbol: string;
  creator: string;
  base_uri: string;
  initial_price: string;
  max_supply: number;
  current_supply: number;
  total_volume: string;
  floor_price: string;
  owners: number;
  total_transactions: number;
  period_volume?: number;
  period_sales?: number;
}

export interface SearchResponse {
  success: boolean;
  data: SearchResult[];
  message?: string;
  pagination?: {
    page: number;
    pageSize: number;
    total?: number;
    totalPages?: number;
  };
}

export const useSearch = (chainId: number = 84532) => {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState<SearchResult[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isOpen, setIsOpen] = useState(false);

  const searchCollections = useCallback(
    async (searchQuery: string) => {
      if (!searchQuery.trim()) {
        setResults([]);
        setIsOpen(false);
        return;
      }

      setLoading(true);
      setError(null);

      try {
        const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;
        if (!backendUrl) {
          throw new Error('Backend URL not configured');
        }

        const response = await fetch(
          `${backendUrl}/api/v1/collections/search?chainId=${chainId}&query=${encodeURIComponent(
            searchQuery.trim()
          )}&page=1&pageSize=10`
        );

        if (!response.ok) {
          throw new Error(`Search failed: ${response.statusText}`);
        }

        const data: SearchResponse = await response.json();

        if (data.success) {
          setResults(data.data);
          setIsOpen(data.data.length > 0);
        } else {
          setError(data.message || 'Search failed');
          setResults([]);
          setIsOpen(false);
        }
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Search failed');
        setResults([]);
        setIsOpen(false);
      } finally {
        setLoading(false);
      }
    },
    [chainId]
  );

  // Debounced search
  useEffect(() => {
    const timeoutId = setTimeout(() => {
      if (query) {
        searchCollections(query);
      } else {
        setResults([]);
        setIsOpen(false);
      }
    }, 300);

    return () => clearTimeout(timeoutId);
  }, [query, searchCollections]);

  const clearSearch = useCallback(() => {
    setQuery('');
    setResults([]);
    setIsOpen(false);
    setError(null);
  }, []);

  const closeSearch = useCallback(() => {
    setIsOpen(false);
  }, []);

  return {
    query,
    setQuery,
    results,
    loading,
    error,
    isOpen,
    clearSearch,
    closeSearch,
    searchCollections,
  };
};
