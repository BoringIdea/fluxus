import { GraphQLClient, gql } from 'graphql-request';

export interface GetCollectionsResponse {
  collectionStats_collection: CollectionStats[];
}

export interface CollectionInfo {
  id: string;
  address: string;
  name: string;
  symbol: string;
  creator: string;
  creator_fee: string;
  base_uri: string;
  initial_price: string;
  max_supply: number;
  max_price: string;
  support_mint: boolean;
  gas_limit: string;
  price_contract: string;
  is_registered: boolean;
  meta_data?: string;
  created_at: string;
  block_number: string;
}

export interface CollectionStats {
  id: string;
  current_supply: number;
  total_supply: number;
  owners: number;
  total_volume: string;
  floor_price?: string;
  last_updated: string;
  total_transactions: number;
  collection_info: CollectionInfo;
}

export interface Txs {
  id: string;
  price: string;
  sender: string;
  token_ids: string[];
  tx_type: string;
  blockTimestamp: string;
  blockNumber: string;
  collection_info: CollectionInfo;
}

export interface NftOwnership {
  id: string;
  token_id: number;
  owner: string;
  collection_info: CollectionInfo;
}

export interface OwnershipSummary {
  id: string;
  nft_count: number;
  owner: string;
  collection_info: CollectionInfo;
}

export interface GetUserOwnershipSummaryResponse {
  ownershipSummaries: OwnershipSummary[];
}

export interface GetUserOwnedTokensResponse {
  nftownerships: NftOwnership[];
}

export interface CollectionVolumeTransaction {
  price: string;
}

export interface GetCollectionVolumeResponse {
  txs_collection: CollectionVolumeTransaction[];
}

export interface GetUserTxsResponse {
  txs_collection: Txs[];
}

export interface GetCollectionHoldersResponse {
  ownershipSummaries: OwnershipSummary[];
}

export interface GraphQLResponse<T> {
  data: T;
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

export interface GetCrossChainTrasferStatusResponse {
  crossChainStatuses: CrossChainTransferStatus[];
}

export class SimpleGraphQLAPI {
  private client: GraphQLClient;

  constructor(url: string) {
    this.client = new GraphQLClient(url);
  }

  async getCollectionPeriodVolumeAndSales(
    address: string,
    fromTimestamp: number,
    toTimestamp: number,
  ): Promise<{ volume: number; sales: number }> {
    // For time range queries, we still need to fetch and sum individual transactions
    // since Graph Protocol doesn't support aggregation functions like SUM directly
    const query = gql`
      query GetCollectionVolume(
        $address: String!
        $fromTimestamp: BigInt!
        $toTimestamp: BigInt!
      ) {
        txs_collection(
          where: {
            blockTimestamp_gte: $fromTimestamp
            blockTimestamp_lte: $toTimestamp
            collection_info_: { address: $address }
          }
        ) {
          price
        }
      }
    `;

    const result = await this.client.request<GetCollectionVolumeResponse>(
      query,
      {
        address,
        fromTimestamp: Math.floor(fromTimestamp / 1000).toString(),
        toTimestamp: Math.floor(toTimestamp / 1000).toString(),
      },
    );

    // Calculate the sum of all transaction prices
    const totalVolume = result.txs_collection.reduce((sum, tx) => {
      return sum + parseFloat(tx.price);
    }, 0);

    return { volume: totalVolume, sales: result.txs_collection.length };
  }

  async getUserTxs(
    userAddress: string,
    first: number = 20,
    skip: number = 0,
  ): Promise<GetUserTxsResponse> {
    // Validate parameters to prevent GraphQL errors
    const safeFirst = Math.min(Math.max(1, first), 1000); // Between 1 and 1000
    const safeSkip = Math.min(Math.max(0, skip), 5000); // Between 0 and 5000
    const query = gql`
      query GetUserTxs($userAddress: String!, $first: Int!, $skip: Int!) {
        txs_collection(
          where: { sender: $userAddress }
          orderBy: blockTimestamp
          orderDirection: desc
          first: $first
          skip: $skip
        ) {
          id
          price
          sender
          token_ids
          tx_type
          blockTimestamp
          blockNumber
          collection_info {
            id
            address
            name
            symbol
            creator
            creator_fee
            base_uri
            initial_price
            max_supply
            price_contract
            is_registered
            meta_data
            created_at
            block_number
          }
        }
      }
    `;

    const result = await this.client.request<GetUserTxsResponse>(query, {
      userAddress,
      first: safeFirst,
      skip: safeSkip,
    });

    return result;
  }

  async getUserTxsCount(userAddress: string): Promise<number> {
    const query = gql`
      query GetUserTxsCount($userAddress: String!) {
        txs_collection(where: { sender: $userAddress }) {
          id
        }
      }
    `;

    const result = await this.client.request<GetUserTxsResponse>(query, {
      userAddress,
    });

    return result.txs_collection.length;
  }

  async getTxs(
    userAddress?: string,
    collectionAddress?: string,
    txTypes?: number[],
    first: number = 20,
    skip: number = 0,
    tokenID?: string,
  ): Promise<GetUserTxsResponse> {
    // Validate parameters to prevent GraphQL errors
    const safeFirst = Math.min(Math.max(1, first), 1000); // Between 1 and 1000
    const safeSkip = Math.min(Math.max(0, skip), 5000); // Between 0 and 5000

    // Build dynamic where clause based on provided parameters
    const whereConditions: string[] = [];
    const variables: any = {
      first: safeFirst,
      skip: safeSkip,
    };

    if (userAddress) {
      whereConditions.push('sender: $userAddress');
      variables.userAddress = userAddress;
    }

    if (collectionAddress) {
      whereConditions.push('collection_info_: { address: $collectionAddress }');
      variables.collectionAddress = collectionAddress;
    }

    if (txTypes && txTypes.length > 0) {
      if (txTypes.length === 1) {
        whereConditions.push('tx_type: $txType');
        variables.txType = txTypes[0];
      } else {
        whereConditions.push('tx_type_in: $txTypes');
        variables.txTypes = txTypes;
      }
    }

    if (tokenID) {
      whereConditions.push('token_ids_contains: [$tokenID]');
      variables.tokenID = tokenID;
    }

    const whereClause =
      whereConditions.length > 0
        ? `where: { ${whereConditions.join(', ')} }`
        : '';

    // Build dynamic query parameters
    const queryParams: string[] = ['$first: Int!', '$skip: Int!'];
    if (userAddress) queryParams.push('$userAddress: String!');
    if (collectionAddress) queryParams.push('$collectionAddress: String!');
    if (txTypes && txTypes.length > 0) {
      if (txTypes.length === 1) {
        queryParams.push('$txType: Int!');
      } else {
        queryParams.push('$txTypes: [Int!]!');
      }
    }
    if (tokenID) queryParams.push('$tokenID: String!');

    const query = gql`
      query GetTxs(${queryParams.join(', ')}) {
        txs_collection(
          ${whereClause}
          orderBy: blockTimestamp
          orderDirection: desc
          first: $first
          skip: $skip
        ) {
          id
          price
          sender
          token_ids
          tx_type
          blockTimestamp
          blockNumber
          collection_info {
            id
            address
            name
            symbol
            creator
            creator_fee
            base_uri
            initial_price
            max_supply
            price_contract
            is_registered
            meta_data
            created_at
            block_number
          }
        }
      }
    `;

    const result = await this.client.request<GetUserTxsResponse>(
      query,
      variables,
    );
    return result;
  }

  async getTxsCount(
    userAddress?: string,
    collectionAddress?: string,
    txTypes?: number[],
    tokenID?: string,
  ): Promise<number> {
    // Build dynamic where clause based on provided parameters
    const whereConditions: string[] = [];
    const variables: any = {};

    if (userAddress) {
      whereConditions.push('sender: $userAddress');
      variables.userAddress = userAddress;
    }

    if (collectionAddress) {
      whereConditions.push('collection_info_: { address: $collectionAddress }');
      variables.collectionAddress = collectionAddress;
    }

    if (txTypes && txTypes.length > 0) {
      if (txTypes.length === 1) {
        whereConditions.push('tx_type: $txType');
        variables.txType = txTypes[0];
      } else {
        whereConditions.push('tx_type_in: $txTypes');
        variables.txTypes = txTypes;
      }
    }

    if (tokenID) {
      whereConditions.push('token_ids_contains: [$tokenID]');
      variables.tokenID = tokenID;
    }

    const whereClause =
      whereConditions.length > 0
        ? `where: { ${whereConditions.join(', ')} }`
        : '';

    // Build dynamic query parameters
    const queryParams: string[] = [];
    if (userAddress) queryParams.push('$userAddress: String!');
    if (collectionAddress) queryParams.push('$collectionAddress: String!');
    if (txTypes && txTypes.length > 0) {
      if (txTypes.length === 1) {
        queryParams.push('$txType: Int!');
      } else {
        queryParams.push('$txTypes: [Int!]!');
      }
    }
    if (tokenID) queryParams.push('$tokenID: String!');

    const paramString =
      queryParams.length > 0 ? `(${queryParams.join(', ')})` : '';

    const query = gql`
      query GetTxsCount${paramString} {
        txs_collection(${whereClause}) {
          id
        }
      }
    `;

    const result = await this.client.request<GetUserTxsResponse>(
      query,
      variables,
    );
    return result.txs_collection.length;
  }

  // Keep backward compatibility methods
  async getUserTxsByCollection(
    userAddress: string,
    collectionAddress: string,
    first: number = 20,
    skip: number = 0,
  ): Promise<GetUserTxsResponse> {
    return this.getTxs(userAddress, collectionAddress, undefined, first, skip);
  }

  async getUserTxsByCollectionCount(
    userAddress: string,
    collectionAddress: string,
  ): Promise<number> {
    return this.getTxsCount(userAddress, collectionAddress);
  }

  async getCollection(address: string): Promise<GetCollectionsResponse> {
    const query = gql`
      query GetCollection($address: String!) {
        collectionStats_collection(
          where: { collection_info_: { address: $address } }
        ) {
          collection_info {
            id
            address
            base_uri
            block_number
            created_at
            creator
            creator_fee
            initial_price
            is_registered
            max_supply
            max_price
            gas_limit
            meta_data
            name
            price_contract
            symbol
            support_mint
          }
          current_supply
          floor_price
          id
          last_updated
          owners
          total_supply
          total_transactions
          total_volume
        }
      }
    `;

    try {
      const result = await this.client.request<GetCollectionsResponse>(query, {
        address,
      });
      return result;
    } catch (error) {
      console.error('GraphQL Error:', error);
      throw error;
    }
  }

  async getCollections(
    first = 100,
    skip = 0,
    orderBy: string = 'collection_info__created_at',
    orderDirection: string = 'asc',
  ): Promise<GetCollectionsResponse> {
    try {
      const query = gql`
        query GetCollections(
          $first: Int!
          $skip: Int!
          $orderBy: String
          $orderDirection: String
        ) {
          collectionStats_collection(
            first: $first
            skip: $skip
            orderBy: $orderBy
            orderDirection: $orderDirection
          ) {
            collection_info {
              id
              address
              base_uri
              block_number
              created_at
              creator
              creator_fee
              initial_price
              is_registered
              max_supply
              meta_data
              name
              price_contract
              symbol
            }
            current_supply
            floor_price
            id
            last_updated
            owners
            total_supply
            total_transactions
            total_volume
          }
        }
      `;

      const result = await this.client.request<GetCollectionsResponse>(query, {
        first,
        skip,
        orderBy,
        orderDirection,
      });

      return result;
    } catch (error) {
      console.error('GraphQL Error:', error);
      throw error;
    }
  }

  async getCollectionsCount(): Promise<number> {
    const query = gql`
      query GetCollectionsCount {
        collectionStats_collection {
          id
        }
      }
    `;

    const result = await this.client.request<GetCollectionsResponse>(query);
    return result.collectionStats_collection.length;
  }

  async getUserOwnedTokens(
    userAddress: string,
    first: number = 20,
    skip: number = 0,
    collectionAddress?: string,
  ): Promise<GetUserOwnedTokensResponse> {
    let query;
    if (collectionAddress && collectionAddress !== '') {
      query = gql`
        query GetUserOwnedTokensByCollection(
          $userAddress: String!
          $collectionAddress: String!
          $first: Int!
          $skip: Int!
        ) {
          nftownerships(
            where: {
              owner: $userAddress
              collection_info_: { address: $collectionAddress }
            }
            orderBy: token_id
            orderDirection: desc
            first: $first
            skip: $skip
          ) {
            token_id
            collection_info {
              id
              address
              name
              symbol
              base_uri
            }
            owner
          }
        }
      `;
    } else {
      query = gql`
        query GetUserOwnedTokens(
          $userAddress: String!
          $first: Int!
          $skip: Int!
        ) {
          nftownerships(
            where: { owner: $userAddress }
            orderBy: token_id
            orderDirection: desc
            first: $first
            skip: $skip
          ) {
            token_id
            collection_info {
              id
              address
              name
              symbol
            }
            owner
          }
        }
      `;
    }

    const result = await this.client.request<GetUserOwnedTokensResponse>(
      query,
      {
        userAddress,
        collectionAddress,
        first,
        skip,
      },
    );

    return result;
  }

  async getUserAllOwneredTokens(
    userAddress: string,
  ): Promise<GetUserOwnedTokensResponse> {
    const query = gql`
      query GetUserAllOwneredTokens($userAddress: String!) {
        nftownerships(where: { owner: $userAddress }) {
          token_id
          collection_info {
            id
            address
            name
            symbol
            initial_price
            max_supply
          }
        }
      }
    `;

    const result = await this.client.request<GetUserOwnedTokensResponse>(
      query,
      { userAddress },
    );
    return result;
  }

  async getUserOwnedTokensCount(
    userAddress: string,
    collectionAddress?: string,
  ): Promise<number> {
    const query = gql`
      query GetUserOwnedTokensCount(
        $userAddress: String!
        $collectionAddress: String!
      ) {
        nftownerships(
          where: {
            owner: $userAddress
            collection_info_: { address: $collectionAddress }
          }
        ) {
          id
        }
      }
    `;

    const result = await this.client.request<GetUserOwnedTokensResponse>(
      query,
      {
        userAddress,
        collectionAddress,
      },
    );

    return result.nftownerships.length;
  }

  async getUserOwnershiSummary(
    userAddress: string,
  ): Promise<GetUserOwnershipSummaryResponse> {
    const query = gql`
      query GetUserOwnershiSummary($userAddress: String!) {
        ownershipSummaries(where: { owner: $userAddress }) {
          id
          nft_count
          owner
          collection_info {
            id
            address
            name
            symbol
            initial_price
            max_supply
          }
        }
      }
    `;

    const result = await this.client.request<GetUserOwnershipSummaryResponse>(
      query,
      { userAddress },
    );
    return result;
  }

  async getCollectionHolders(
    collectionAddress: string,
    first: number,
    skip: number,
  ): Promise<GetCollectionHoldersResponse> {
    const query = gql`
      query getTopHolders(
        $collectionAddress: Bytes!
        $first: Int!
        $skip: Int!
      ) {
        ownershipSummaries(
          where: {
            collection_info_: { address: $collectionAddress }
            nft_count_gt: 0
          }
          orderBy: nft_count
          orderDirection: desc
          first: $first
          skip: $skip
        ) {
          owner
          nft_count
        }
      }
    `;

    const result = await this.client.request<GetCollectionHoldersResponse>(
      query,
      { collectionAddress, first, skip },
    );
    return result;
  }

  async getCollectionHoldersCount(collectionAddress: string): Promise<number> {
    const query = gql`
      query GetCollectionHoldersCount($collectionAddress: Bytes!) {
        ownershipSummaries(
          where: { collection_info_: { address: $collectionAddress } }
        ) {
          id
        }
      }
    `;

    const result = await this.client.request<GetCollectionHoldersResponse>(
      query,
      { collectionAddress },
    );
    return result.ownershipSummaries.length;
  }

  async getCrossChainTransferStatus(
    fluxusContract: string,
    first: number,
    skip: number,
    sender?: string,
    receiver?: string,
  ): Promise<GetCrossChainTrasferStatusResponse> {
    const variables: any = { fluxusContract, first, skip };
    const queryParams: string[] = [
      '$fluxusContract: String!',
      '$first: Int!',
      '$skip: Int!',
    ];

    const filters: string[] = ['fluxusContract: $fluxusContract'];
    if (sender && !receiver) {
      variables.sender = sender;
      queryParams.push('$sender: String!');
      filters.push('sender: $sender');
    } else if (receiver && !sender) {
      variables.receiver = receiver;
      queryParams.push('$receiver: String!');
      filters.push('receiver: $receiver');
    }

    const whereClause = `where: { ${filters.join(', ')} }`;

    const query = gql`
      query GetCrossChainTransferStatus(${queryParams.join(', ')}) {
        crossChainStatuses(
          ${whereClause}
          orderBy: blockTimestamp
          orderDirection: desc
          first: $first
          skip: $skip
        ) {
          id
          destination
          fluxusContract
          isTransfered
          receiver
          sender
          tokenId
          blockTimestamp
          transactionHash
        }
      }
    `;

    const result =
      await this.client.request<GetCrossChainTrasferStatusResponse>(
        query,
        variables,
      );
    return result;
  }

  async getCrossChainTransferStatusCount(
    fluxusContract: string,
    sender?: string,
    receiver?: string,
  ): Promise<number> {
    const variables: any = { fluxusContract };
    const queryParams: string[] = ['$fluxusContract: String!'];

    const filters: string[] = ['fluxusContract: $fluxusContract'];
    if (sender && !receiver) {
      variables.sender = sender;
      queryParams.push('$sender: String!');
      filters.push('sender: $sender');
    } else if (receiver && !sender) {
      variables.receiver = receiver;
      queryParams.push('$receiver: String!');
      filters.push('receiver: $receiver');
    }

    const whereClause = `where: { ${filters.join(', ')} }`;

    const paramString = `(${queryParams.join(', ')})`;
    const query = gql`
      query GetCrossChainTransferStatusCount${paramString} {
        crossChainStatuses(${whereClause}) {
          id
        }
      }
    `;

    const result =
      await this.client.request<GetCrossChainTrasferStatusResponse>(
        query,
        variables,
      );

    return result.crossChainStatuses?.length ?? 0;
  }

  async getCrossChainTransferStatusBothCount(
    fluxusContract: string,
    sender: string,
    receiver: string,
  ): Promise<number> {
    const query = gql`
      query GetCrossChainTransferStatusBothCount(
        $fluxusContract: String!
        $sender: String!
        $receiver: String!
      ) {
        crossChainStatuses(
          where: {
            fluxusContract: $fluxusContract
            sender: $sender
            receiver: $receiver
          }
        ) {
          id
        }
      }
    `;

    const result =
      await this.client.request<GetCrossChainTrasferStatusResponse>(query, {
        fluxusContract,
        sender,
        receiver,
      });

    return result.crossChainStatuses?.length ?? 0;
  }

  async searchCollectionsByName(
    query: string,
    first: number = 20,
    skip: number = 0,
  ): Promise<GetCollectionsResponse> {
    const searchQuery = gql`
      query SearchCollectionsByName(
        $query: String!
        $first: Int!
        $skip: Int!
      ) {
        collectionStats_collection(
          where: { collection_info_: { name_contains_nocase: $query } }
          first: $first
          skip: $skip
          orderBy: collection_info__created_at
          orderDirection: desc
        ) {
          id
          current_supply
          total_supply
          owners
          total_volume
          floor_price
          last_updated
          total_transactions
          collection_info {
            id
            address
            name
            symbol
            creator
            creator_fee
            base_uri
            initial_price
            max_supply
            max_price
            support_mint
            gas_limit
            price_contract
            is_registered
            meta_data
            created_at
            block_number
          }
        }
      }
    `;

    return await this.client.request<GetCollectionsResponse>(searchQuery, {
      query,
      first,
      skip,
    });
  }

  async searchCollectionsByAddress(
    query: string,
    first: number = 20,
    skip: number = 0,
  ): Promise<GetCollectionsResponse> {
    const searchQuery = gql`
      query SearchCollectionsByAddress(
        $query: String!
        $first: Int!
        $skip: Int!
      ) {
        collectionStats_collection(
          where: { collection_info_: { address_contains: $query } }
          first: $first
          skip: $skip
          orderBy: collection_info__created_at
          orderDirection: desc
        ) {
          id
          current_supply
          total_supply
          owners
          total_volume
          floor_price
          last_updated
          total_transactions
          collection_info {
            id
            address
            name
            symbol
            creator
            creator_fee
            base_uri
            initial_price
            max_supply
            max_price
            support_mint
            gas_limit
            price_contract
            is_registered
            meta_data
            created_at
            block_number
          }
        }
      }
    `;

    return await this.client.request<GetCollectionsResponse>(searchQuery, {
      query,
      first,
      skip,
    });
  }
}
