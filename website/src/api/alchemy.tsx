import { getAlchemy, SupportedChainId } from '@/src/alchemy';
import { OwnedNft, NftOrdering, NftFilters } from 'alchemy-sdk';

// Alchemy NFT API Response Types based on documentation
export interface AlchemyNFT {
  contract: {
    address: string;
    name: string;
    symbol: string;
    totalSupply: string;
    tokenType: 'ERC721' | 'ERC1155';
    contractDeployer: string;
    deployedBlockNumber: number;
    spamClassifications: string[];
  };
  tokenId: string;
  tokenType: 'ERC721' | 'ERC1155';
  name?: string;
  description?: string;
  image?: {
    cachedUrl: string;
    thumbnailUrl: string;
    pngUrl: string;
    contentType: string;
    size: number;
    originalUrl: string;
  };
  raw: {
    tokenUri: string;
    metadata: Record<string, any>;
  };
  tokenUri?: string;
  timeLastUpdated: string;
}

export interface AlchemyNFTsResponse {
  ownedNfts: AlchemyNFT[];
  totalCount: number;
  pageKey?: string;
  validAt: {
    blockNumber: number;
    blockHash: string;
    blockTimestamp: string;
  };
}

// getNFTMetadata response (simplified for traits rendering)
export interface AlchemyNFTMetadataResponse {
  contract: {
    address: string;
    name?: string;
    symbol?: string;
  };
  tokenId: string;
  name?: string;
  description?: string;
  image?: {
    cachedUrl?: string;
    thumbnailUrl?: string;
    pngUrl?: string;
    contentType?: string;
    size?: number;
    originalUrl?: string;
  };
  raw?: {
    tokenUri?: string;
    metadata?: Record<string, any> & {
      attributes?: Array<{ trait_type?: string; value?: string | number }>;
    };
  };
  tokenUri?: string;
  timeLastUpdated?: string;
}

export interface GetNFTsByOwnerParams {
  chainId: SupportedChainId;
  owner: string;
  contractAddresses?: string[];
  withMetadata?: boolean;
  orderBy?: NftOrdering;
  excludeFilters?: NftFilters[];
  includeFilters?: NftFilters[];
  spamConfidenceLevel?: 'VERY_HIGH' | 'HIGH' | 'MEDIUM' | 'LOW';
  tokenUriTimeoutInMs?: number;
  pageKey?: string;
  pageSize?: number;
}

/**
 * Get NFTs owned by a specific address using Alchemy NFT API
 * @param params - Parameters for the API call
 * @returns Promise<AlchemyNFTsResponse>
 */
export async function getNFTsByOwner(params: GetNFTsByOwnerParams): Promise<AlchemyNFTsResponse> {
  const {
    chainId,
    owner,
    contractAddresses,
    withMetadata = true,
    orderBy,
    excludeFilters,
    includeFilters,
    spamConfidenceLevel,
    tokenUriTimeoutInMs,
    pageKey,
    pageSize = 100
  } = params;
  // Get Alchemy instance for the chain
  const alchemy = getAlchemy(chainId);
  if (!alchemy) {
    throw new Error(`Alchemy not supported for chain ID: ${chainId}`);
  }

  try {
    // Build query parameters
    const queryParams: Record<string, any> = {
      owner,
      withMetadata,
      pageSize: Math.min(pageSize, 100) // Max limit is 100
    };

    // Add optional parameters
    if (contractAddresses && contractAddresses.length > 0) {
      queryParams.contractAddresses = contractAddresses;
    }
    if (orderBy) {
      queryParams.orderBy = orderBy;
    }
    if (excludeFilters && excludeFilters.length > 0) {
      queryParams.excludeFilters = excludeFilters;
    }
    if (includeFilters && includeFilters.length > 0) {
      queryParams.includeFilters = includeFilters;
    }
    if (spamConfidenceLevel) {
      queryParams.spamConfidenceLevel = spamConfidenceLevel;
    }
    if (tokenUriTimeoutInMs !== undefined) {
      queryParams.tokenUriTimeoutInMs = tokenUriTimeoutInMs;
    }
    if (pageKey) {
      queryParams.pageKey = pageKey;
    }

    // Call Alchemy NFT API
    const response = await alchemy.nft.getNftsForOwner(owner, {
      contractAddresses,
      omitMetadata: !withMetadata,
      orderBy,
      excludeFilters,
      includeFilters,
      tokenUriTimeoutInMs,
      pageKey,
      pageSize: Math.min(pageSize, 100)
    });

    // Transform response to match our interface
    const result: AlchemyNFTsResponse = {
      ownedNfts: response.ownedNfts.map((nft: any) => {
        const contractAddress = nft?.contract?.address || nft?.contractAddress || '';
        const tokenType = (nft?.tokenType as 'ERC721' | 'ERC1155') || 'ERC721';
        const image = nft?.image
          ? {
              cachedUrl: nft.image.cachedUrl || '',
              thumbnailUrl: nft.image.thumbnailUrl || '',
              pngUrl: nft.image.pngUrl || '',
              contentType: nft.image.contentType || '',
              size: nft.image.size || 0,
              originalUrl: nft.image.originalUrl || ''
            }
          : undefined;
        const raw = nft?.raw
          ? {
              tokenUri: nft.raw?.tokenUri || '',
              metadata: nft.raw?.metadata || {}
            }
          : { tokenUri: '', metadata: {} };

        return {
          contract: {
            address: contractAddress,
            name: nft?.contract?.name || '',
            symbol: nft?.contract?.symbol || '',
            totalSupply: nft?.contract?.totalSupply || '0',
            tokenType,
            contractDeployer: nft?.contract?.contractDeployer || '',
            deployedBlockNumber: nft?.contract?.deployedBlockNumber || 0,
            spamClassifications: nft?.contract?.spamClassifications || []
          },
          tokenId: nft?.tokenId?.toString?.() || String(nft?.tokenId || ''),
          tokenType,
          name: nft?.name,
          description: nft?.description,
          image,
          raw,
          tokenUri: nft?.tokenUri,
          timeLastUpdated: nft?.timeLastUpdated || new Date().toISOString()
        } as AlchemyNFT;
      }),
      totalCount: response.totalCount,
      pageKey: response.pageKey,
      validAt: {
        blockNumber: response.validAt?.blockNumber || 0,
        blockHash: response.validAt?.blockHash || '',
        blockTimestamp: response.validAt?.blockTimestamp || new Date().toISOString()
      }
    };

    return result;
  } catch (error) {
    console.error('Error fetching NFTs by owner:', error);
    throw new Error(`Failed to fetch NFTs for owner ${owner} on chain ${chainId}: ${error instanceof Error ? error.message : 'Unknown error'}`);
  }
}

/**
 * Get NFTs owned by a specific address for a specific collection
 * @param chainId - Chain ID
 * @param owner - Owner address
 * @param collectionAddress - Collection contract address
 * @param withMetadata - Whether to include metadata (default: true)
 * @param pageSize - Number of NFTs per page (default: 100, max: 100)
 * @returns Promise<AlchemyNFTsResponse>
 */
export async function getNFTsByOwnerForCollection(
  chainId: SupportedChainId,
  owner: string,
  collectionAddress: string,
  withMetadata: boolean = true,
  pageSize: number = 100
): Promise<AlchemyNFTsResponse> {
  return getNFTsByOwner({
    chainId,
    owner,
    contractAddresses: [collectionAddress],
    withMetadata,
    pageSize
  });
}

/**
 * Get all NFTs owned by a specific address (across all collections)
 * @param chainId - Chain ID
 * @param owner - Owner address
 * @param withMetadata - Whether to include metadata (default: true)
 * @param pageSize - Number of NFTs per page (default: 100, max: 100)
 * @returns Promise<AlchemyNFTsResponse>
 */
export async function getAllNFTsByOwner(
  chainId: SupportedChainId,
  owner: string,
  withMetadata: boolean = true,
  pageSize: number = 100
): Promise<AlchemyNFTsResponse> {
  return getNFTsByOwner({
    chainId,
    owner,
    withMetadata,
    pageSize
  });
}

/**
 * Get NFT metadata by contract and tokenId via Alchemy SDK
 * Minimal shape: includes image urls and raw.metadata.attributes for traits.
 * Docs: https://www.alchemy.com/docs/reference/nft-api-endpoints/nft-api-endpoints/nft-metadata-endpoints/get-nft-metadata-v-3
 */
export async function getNFTMetadataByTokenId(
  chainId: SupportedChainId,
  contractAddress: string,
  tokenId: string | number
): Promise<AlchemyNFTMetadataResponse> {
  const alchemy = getAlchemy(chainId);
  if (!alchemy) {
    throw new Error(`Alchemy not supported for chain ID: ${chainId}`);
  }

  try {
    const res: any = await alchemy.nft.getNftMetadata(contractAddress, tokenId);
    const image = res?.image
      ? {
          cachedUrl: res.image.cachedUrl,
          thumbnailUrl: res.image.thumbnailUrl,
          pngUrl: res.image.pngUrl,
          contentType: res.image.contentType,
          size: res.image.size,
          originalUrl: res.image.originalUrl,
        }
      : undefined;
    return {
      contract: {
        address: res?.contract?.address || contractAddress,
        name: res?.contract?.name,
        symbol: res?.contract?.symbol,
      },
      tokenId: res?.tokenId?.toString?.() || String(tokenId),
      name: res?.name,
      description: res?.description,
      image,
      raw: res?.raw,
      tokenUri: res?.tokenUri,
      timeLastUpdated: res?.timeLastUpdated,
    } as AlchemyNFTMetadataResponse;
  } catch (error) {
    console.error('Error fetching NFT metadata:', error);
    throw new Error(
      `Failed to fetch NFT metadata for ${contractAddress} #${tokenId} on chain ${chainId}: ${error instanceof Error ? error.message : 'Unknown error'}`
    );
  }
}
