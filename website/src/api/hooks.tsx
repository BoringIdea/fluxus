import useSWR from 'swr';
import { Collection, CrossChainTransferStatus, Pagination, UserToken } from './types';

const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;

const fetcher = (url: string) => fetch(url).then(res => res.json());

export const useCollection = (chainId: number, contractAddress: string) => {
  return useSWR<{ data: Collection }>(
    contractAddress ? `${backendUrl}/api/v1/collections/${contractAddress}?chainId=${chainId}` : null,
    fetcher,
    { refreshInterval: 1000 }
  );
};

export const useUserTokensByCollection = (chainId: number, userAddress: string, collectionAddress: string, page = 1, pageSize = 100) => {
  return useSWR<{ data: UserToken[], pagination: Pagination }>(
    userAddress && collectionAddress ? `${backendUrl}/api/v1/users/${userAddress}/collections/${collectionAddress}/tokens?page=${page}&pageSize=${pageSize}&chainId=${chainId}` : null,
    fetcher,
    {
      refreshInterval: 1000,
      onError: (error) => {
        console.error('Error fetching user tokens:', error);
      },
      fallbackData: {
        data: [],
        pagination: {
          total: 0,
          limit: pageSize,
          offset: page * pageSize,
          hasMore: false
        }
      }
    }
  );
};

export interface CollectionTxsPagination {
  total: string;
  offset: number;
  limit: number;
  hasMore: boolean;
}

export interface CollectionTx {
  txHash: string;
  txType: number;
  sender: string;
  price: string;
  tokenIds: number[];
  status: number;
  createdAt: string;
}

export const useCollectionTxs = (chainId: number, collectionAddress: string, pageSize = 10, page = 0, type?: string) => {
  const swr = useSWR<{
    data: CollectionTx[];
    pagination: CollectionTxsPagination;
  }>(
    collectionAddress ? `${backendUrl}/api/v1/collections/${collectionAddress}/transactions?pageSize=${pageSize}&page=${page}&chainId=${chainId}${type ? `&type=${type}` : ''}` : null,
    fetcher,
    {
      refreshInterval: 10000
    }
  );

  return {
    data: swr.data?.data || [],
    pagination: swr.data?.pagination,
    isLoading: swr.isLoading,
    error: swr.error,
  };
};

// Token-level transactions by filtering collection transactions on client
export const useTokenTxs = (
  chainId: number,
  collectionAddress: string,
  tokenId: number,
  pageSize = 50,
  page = 0
) => {
  const swr = useSWR<{
    data: CollectionTx[];
    pagination: CollectionTxsPagination;
  }>(
    collectionAddress
      ? `${backendUrl}/api/v1/collections/${collectionAddress}/transactions?pageSize=${pageSize}&page=${page}&chainId=${chainId}&tokenID=${tokenId}`
      : null,
    fetcher,
    { refreshInterval: 10000 }
  );

  return {
    data: swr.data?.data || [],
    pagination: swr.data?.pagination,
    isLoading: swr.isLoading,
    error: swr.error,
  };
};

export interface CollectionHolder {
  owner: string;
  nft_count: number;
  firstAcquiredAt: string;
  lastAcquiredAt: string;
}

export const useCollectionHolders = (
  chainId: number,
  collectionAddress: string,
  pageSize = 10,
  page = 1
) => {
  const swr = useSWR<{
    data: CollectionHolder[];
    pagination: CollectionTxsPagination;
  }>(
    collectionAddress
      ? `${backendUrl}/api/v1/collections/${collectionAddress}/holders?pageSize=${pageSize}&page=${page}&chainId=${chainId}`
      : null,
    fetcher,
    { refreshInterval: 10000 }
  );

  return {
    data: swr.data?.data || [],
    pagination: swr.data?.pagination,
    isLoading: swr.isLoading,
    error: swr.error,
  };
};

export interface UserProfileCollection {
  name: string;
  symbol: string;
  address: string;
  initialPrice: string;
  maxSupply: number;
  nftCount: number;
  floorPrice: string;
}

export interface UserProfile {
  totalCollections: number;
  totalNFTs: number;
  collections: UserProfileCollection[];
}

export const useUserProfile = (chainId: number, userAddress: string) => {
  const { data, ...rest } = useSWR<{ data: UserProfile }>(
    userAddress ? `${backendUrl}/api/v1/users/${userAddress}/profile?chainId=${chainId}` : null,
    fetcher,
    {
      fallbackData: {
        data: {
        totalCollections: 0,
        totalNFTs: 0,
        collections: []
        }
      }
    }
  );

  return { ...data?.data, ...rest };
};

export const useCollections = (chainId: number, skip: number, first: number, period: string) => {
  return useSWR<{ data: Collection[], pagination: Pagination }>(
    `${backendUrl}/api/v1/collections?skip=${skip}&first=${first}&period=${period}&chainId=${chainId}`,
    fetcher,
    { refreshInterval: 1000, fallbackData: { data: [], pagination: { hasMore: false, limit: 0, offset: 0, total: 0 } } }
  );
};

export const useRecommendedCollection = (chainId: number) => {
  return useSWR<{ data: Collection[] }>(
    `${backendUrl}/api/v1/collections/recommended?chainId=${chainId}`,
    fetcher,
    { refreshInterval: 1000, fallbackData: { data: [] } }
  );
};

// New hooks for additional user endpoints
export const useUserTransactions = (chainId: number, userAddress: string, page = 1, pageSize = 10) => {
  const swr = useSWR<{
    data: CollectionTx[];
    pagination: CollectionTxsPagination;
  }>(
    userAddress ? `${backendUrl}/api/v1/users/${userAddress}/transactions?page=${page}&pageSize=${pageSize}&chainId=${chainId}` : null,
    fetcher,
    {
      refreshInterval: 10000
    }
  );

  return {
    data: swr.data?.data || [],
    pagination: swr.data?.pagination,
    isLoading: swr.isLoading,
    error: swr.error,
  };
};

export const useUserTokens = (chainId: number, userAddress: string, page = 1, pageSize = 10) => {
  const swr = useSWR<{
    data: UserToken[];
    pagination: CollectionTxsPagination;
  }>(
    userAddress ? `${backendUrl}/api/v1/users/${userAddress}/tokens?page=${page}&pageSize=${pageSize}&chainId=${chainId}` : null,
    fetcher,
    {
      refreshInterval: 10000
    }
  );

  return {
    data: swr.data?.data || [],
    pagination: swr.data?.pagination,
    isLoading: swr.isLoading,
    error: swr.error,
  };
};

export const useUserCollectionTransactions = (
  chainId: number,
  userAddress: string,
  collectionAddress: string,
  page = 1,
  pageSize = 10
) => {
  const swr = useSWR<{
    data: CollectionTx[];
    pagination: CollectionTxsPagination;
  }>(
    userAddress && collectionAddress 
      ? `${backendUrl}/api/v1/users/${userAddress}/collections/${collectionAddress}/transactions?page=${page}&pageSize=${pageSize}&chainId=${chainId}` 
      : null,
    fetcher,
    {
      refreshInterval: 10000
    }
  );

  return {
    data: swr.data?.data || [],
    pagination: swr.data?.pagination,
    isLoading: swr.isLoading,
    error: swr.error,
  };
};

export const useUserOwnedTokensCount = (chainId: number, userAddress: string, collectionAddress: string) => {
  return useSWR<{ data: number }>(
    userAddress && collectionAddress 
      ? `${backendUrl}/api/v1/users/${userAddress}/collections/${collectionAddress}/tokens/count?chainId=${chainId}` 
      : null,
    fetcher,
    { refreshInterval: 10000 }
  );
};

export const useCrossChainStatus = (chainId: number, contractAddress: string, userAddress?: string) => {
  // Reduce refresh rate and avoid refetch on window focus to prevent heavy reloads when navigating
  const swrConfig = {
    refreshInterval: 15000,
    revalidateOnFocus: false,
    revalidateIfStale: false,
    revalidateOnReconnect: true,
    fallbackData: { data: [], pagination: { hasMore: false, limit: 0, offset: 0, total: 0 } }
  } as const;

  return userAddress ? useSWR<{ data: CrossChainTransferStatus[], pagination: Pagination }>(
    userAddress && contractAddress 
      ? `${backendUrl}/api/v1/crosschain/status/${contractAddress}?sender=${userAddress}&chainId=${chainId}` 
      : null,
    fetcher,
    swrConfig as any
  ) : useSWR<{ data: CrossChainTransferStatus[], pagination: Pagination }>(
    contractAddress 
      ? `${backendUrl}/api/v1/crosschain/status/${contractAddress}?chainId=${chainId}` 
      : null,
    fetcher,
    swrConfig as any
  );
};
