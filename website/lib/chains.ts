/**
 * Chain utilities for handling blockchain network information
 */

export interface ChainInfo {
  id: string;
  name: string;
  isTestnet: boolean;
  nativeCurrency: {
    name: string;
    symbol: string;
    decimals: number;
  };
  rpcUrls?: string[];
  blockExplorerUrls?: string[];
}

/**
 * Mapping of chain IDs to chain names
 */
const CHAIN_MAP: Record<string, string> = {
  // Testnets
  '84532': 'Base Sepolia',
  '97': 'BSC Testnet',
  '7001': 'ZetaChain Athens',
  '11155111': 'Sepolia',
  '80001': 'Polygon Mumbai',
  '43113': 'Avalanche Fuji',
  '420': 'Optimism Goerli',
  '421613': 'Arbitrum Goerli',
  
  // Mainnets
  '1': 'Ethereum Mainnet',
  '56': 'BSC Mainnet',
  '8453': 'Base Mainnet',
  '7000': 'ZetaChain Mainnet',
  '137': 'Polygon',
  '43114': 'Avalanche',
  '10': 'Optimism',
  '42161': 'Arbitrum One',
  '250': 'Fantom',
  '25': 'Cronos',
  '100': 'Gnosis Chain'
};

/**
 * Get chain name by chain ID
 * @param chainId - The chain ID as string
 * @returns The human-readable chain name
 */
export function getChainName(chainId: string): string {
  return CHAIN_MAP[chainId] || `Chain ${chainId}`;
}

/**
 * Check if a chain is a testnet
 * @param chainId - The chain ID as string
 * @returns true if the chain is a testnet
 */
export function isTestnet(chainId: string): boolean {
  const testnetChains = ['84532', '97', '7001', '11155111', '80001', '43113', '420', '421613'];
  return testnetChains.includes(chainId);
}

/**
 * Get all supported chain IDs
 * @returns Array of supported chain IDs
 */
export function getSupportedChainIds(): string[] {
  return Object.keys(CHAIN_MAP);
}

/**
 * Check if a chain ID is supported
 * @param chainId - The chain ID to check
 * @returns true if the chain is supported
 */
export function isChainSupported(chainId: string): boolean {
  return chainId in CHAIN_MAP;
}

/**
 * Map CCTX status to user-friendly transaction status
 * @param cctxStatus - The CCTX status string from ZetaChain
 * @returns Normalized transaction status
 */
export function mapCCTXStatus(cctxStatus: string): 'pending' | 'completed' | 'failed' {
  // Normalize status to handle case variations
  const normalizedStatus = cctxStatus?.toLowerCase() || '';
  
  // Check for completed states
  if (normalizedStatus.includes('outboundmined') || normalizedStatus === 'outboundmined') {
    return 'completed';
  }
  
  // Check for failed/reverted states
  if (normalizedStatus.includes('aborted') || 
      normalizedStatus.includes('reverted') || 
      normalizedStatus.includes('revert') ||
      normalizedStatus.includes('failed')) {
    return 'failed';
  }
  
  // Check for pending states (includes various pending statuses)
  if (normalizedStatus.includes('pending') ||
      normalizedStatus.includes('outboundmining') ||
      normalizedStatus.includes('inboundprocessed') ||
      normalizedStatus.includes('pending_inbound') ||
      normalizedStatus.includes('pending_outbound') ||
      normalizedStatus === '') {
    return 'pending';
  }
  
  // Default to pending for unknown statuses
  console.log(`⚠️ Unknown CCTX status: ${cctxStatus}, defaulting to pending`);
  return 'pending';
}

/**
 * Get chain info by chain ID (if we need more detailed info in the future)
 * @param chainId - The chain ID as string
 * @returns Basic chain information
 */
export function getChainInfo(chainId: string): ChainInfo | null {
  const name = getChainName(chainId);
  if (name.startsWith('Chain ')) {
    return null; // Unsupported chain
  }

  return {
    id: chainId,
    name,
    isTestnet: isTestnet(chainId),
    nativeCurrency: {
      name: 'ETH', // Default to ETH, can be expanded later
      symbol: 'ETH',
      decimals: 18
    }
  };
}
