import { useState, useEffect, useMemo } from 'react';
import { calculateFluxusCrossChainAddress, DEFAULT_CONTRACT_PARAMS, ContractParams } from '@/lib/contractAddressCalculator';

export function useContractAddress(chainId: number, params?: Partial<ContractParams>) {
  const [contractAddress, setContractAddress] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  // Create a stable dependency key
  const paramsKey = useMemo(() => {
    if (!params) return 'default';
    return JSON.stringify(params, (key, value) => {
      // Convert BigInt to string for stable comparison
      return typeof value === 'bigint' ? value.toString() : value;
    });
  }, [params]);

  useEffect(() => {
    if (!chainId) return;

    const fetchContractAddress = async () => {
      setIsLoading(true);
      setError(null);
      
      try {
        // Merge default params with provided params
        const finalParams = { ...DEFAULT_CONTRACT_PARAMS, ...params };
        
        const address = await calculateFluxusCrossChainAddress(chainId, finalParams);
        setContractAddress(address);
      } catch (err) {
        const error = err instanceof Error ? err : new Error('Unknown error');
        setError(error);
        console.error(`Error fetching contract address for chain ${chainId}:`, error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchContractAddress();
  }, [chainId, paramsKey]);

  return {
    data: contractAddress,
    isLoading,
    error,
    refetch: () => {
      if (chainId) {
        const fetchContractAddress = async () => {
          setIsLoading(true);
          setError(null);
          
          try {
            const finalParams = { ...DEFAULT_CONTRACT_PARAMS, ...params };
            const address = await calculateFluxusCrossChainAddress(chainId, finalParams);
            setContractAddress(address);
          } catch (err) {
            const error = err instanceof Error ? err : new Error('Unknown error');
            setError(error);
          } finally {
            setIsLoading(false);
          }
        };
        fetchContractAddress();
      }
    }
  };
}
