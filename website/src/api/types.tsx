export interface Collection {
  id: number;
  address: string;
  price_contract: string;
  creator: string;
  name: string;
  symbol: string;
  max_supply: number;
  max_price: string;
  creator_fee: string;
  initial_price: string;
  gas_limit: string;
  support_mint: boolean;
  base_uri: string;
  owners: number;
  current_supply: number;
  total_supply: number;
  floor_price: string;
  volume: string;
  period_volume?: string;
  total_volume: string;
  volume_1d?: string;
  sales_1d?: number;
  meta_data: CollectionMetadata;
}

export interface CollectionMetadata {
  name: string;
  symbol: string;
  description: string;
  image: string;
  banner_image: string;
  website?: string;
  twitter?: string;
  discord?: string;
}

export interface UserToken {
  token_id: number;
}

export interface Pagination {
  hasMore: boolean
  limit: number
  offset: number
  total: string | number
}

// API Response wrapper interface based on backend structure
export interface ApiResponse<T> {
  success: boolean;
  data: T;
  message?: string;
  pagination?: {
    page: number;
    pageSize: number;
    total?: number;
    totalPages?: number;
  };
}

// Transaction types based on backend enum
export enum TransactionType {
  MINT = 'mint',
  BUY = 'buy', 
  SELL = 'sell'
}

// Enhanced UserToken interface with more details
export interface UserTokenDetail extends UserToken {
  collectionAddress: string;
  collectionName?: string;
}

export interface CrossChainTransferStatus {
  id: string;
  destination: string;
  fluxusContract: string;
  isTransfered: boolean;
  receiver: string;
  sender: string;
  tokenId: number;
  blockTimestamp: number;
  transactionHash: string;
}