import { Injectable, Logger } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { SimpleGraphQLAPI } from './api/simple-graphql';

interface ChainConfig {
  chainName: string;
  chainId: number;
  factoryAddress: string;
  registryAddress: string;
  tradeAddress: string;
  rpcUrl: string;
}

// Response interfaces
export interface PaginationInfo {
  total: number;
  page: number;
  pageSize: number;
  hasMore: boolean;
  totalPages: number;
}

export interface ServiceResponse<T> {
  data: T;
  pagination?: PaginationInfo;
}

@Injectable()
export class AppService {
  private readonly logger = new Logger(AppService.name);
  private baseGraphqlAPI: SimpleGraphQLAPI;
  private bscGraphqlAPI: SimpleGraphQLAPI;
  private chainConfigs: Map<number, ChainConfig> = new Map();

  constructor(private configService: ConfigService) {
    const chains = [
      // Mantle Sepolia
      {
        chainName: 'Mantle Sepolia',
        chainId: 5001,
        factoryAddress: this.configService.get<string>(
          'FACTORY_CONTRACT_ADDRESS',
        ),
        registryAddress: this.configService.get<string>(
          'REGISTRY_CONTRACT_ADDRESS',
        ),
        tradeAddress: this.configService.get<string>('TRADE_CONTRACT_ADDRESS'),
        rpcUrl: this.configService.get<string>('RPC_URL'),
      },
      // Pharos Testnet
      {
        chainName: 'Pharos Testnet',
        chainId: 688688,
        factoryAddress: '0xeF1858F3b4EA7412ab5B1F8B87aB1FFA7023a8C0',
        registryAddress: '0x625468C624C2878C7413349584B93d6696DeA5e3',
        tradeAddress: '0xafb267475caF612Bf275284C0cDB71A4562fEB1a',
        rpcUrl: 'https://testnet.dplabs-internal.com',
      },
      // Monad Testnet
      {
        chainName: 'Monad Testnet',
        chainId: 10143,
        factoryAddress: '0x582AEb4616F2D7AAdac262A14c88f2Bf1a6A3913',
        registryAddress: '0xFFa792f5f2BB3Cbe6B6e9AF676AE29efcEC1bb0d',
        tradeAddress: '0xCD0be2E5e9C53Aa0aF4B4e8aFb6a953C8F002Cea',
        rpcUrl:
          'https://powerful-proud-feather.monad-testnet.quiknode.pro/0b818d10efcd7181231937d99c0700b18d28b855/',
      },
      // MegaETH Testnet
      {
        chainName: 'MegaETH Testnet',
        chainId: 6342,
        factoryAddress: '0xeF1858F3b4EA7412ab5B1F8B87aB1FFA7023a8C0',
        registryAddress: '0x625468C624C2878C7413349584B93d6696DeA5e3',
        tradeAddress: '0xafb267475caF612Bf275284C0cDB71A4562fEB1a',
        rpcUrl: 'https://carrot.megaeth.com/rpc',
      },
    ];

    for (const chain of chains) {
      if (!chain.factoryAddress || !chain.registryAddress || !chain.rpcUrl) {
        throw new Error(
          `Missing required configuration for ${chain.chainName}`,
        );
      }
      this.chainConfigs.set(chain.chainId, chain);
    }

    // Initialize GraphQL client for subgraph
    const baseSubgraphUrl = this.configService.get<string>('BASE_SUBGRAPH_URL');
    const bscSubgraphUrl = this.configService.get<string>('BSC_SUBGRAPH_URL');
    if (!baseSubgraphUrl || !bscSubgraphUrl) {
      throw new Error(
        'Missing required configuration for base and bsc subgraphs',
      );
    }
    console.log('Base Subgraph URL:', baseSubgraphUrl);
    console.log('BSC Subgraph URL:', bscSubgraphUrl);
    this.baseGraphqlAPI = new SimpleGraphQLAPI(baseSubgraphUrl);
    this.bscGraphqlAPI = new SimpleGraphQLAPI(bscSubgraphUrl);
  }

  async getGraphqlAPI(chainId: number | string): Promise<SimpleGraphQLAPI> {
    const normalizedChainId =
      typeof chainId === 'string' ? Number(chainId.trim()) : chainId;
    if (!Number.isFinite(normalizedChainId)) {
      throw new Error(`Invalid chainId: ${chainId}`);
    }
    if (normalizedChainId === 84532) {
      return this.baseGraphqlAPI;
    } else if (normalizedChainId === 97) {
      return this.bscGraphqlAPI;
    }
    throw new Error(`Unsupported chainId: ${chainId}`);
  }

  // ===== Utility Methods =====

  /**
   * Create standardized pagination info
   */
  private createPaginationInfo(
    total: number,
    page: number,
    pageSize: number,
    currentDataLength: number,
    skip: number,
  ): PaginationInfo {
    return {
      total,
      page,
      pageSize,
      hasMore: skip + currentDataLength < total,
      totalPages: Math.ceil(total / pageSize),
    };
  }

  /**
   * Create standardized service response
   */
  private createServiceResponse<T>(
    data: T,
    pagination?: PaginationInfo,
  ): ServiceResponse<T> {
    return {
      data,
      ...(pagination && { pagination }),
    };
  }

  /**
   * Transform collection stats to unified format
   */
  private transformCollectionStats(stats: any, periodVolume?: any) {
    return {
      id: stats.id,
      address: stats.collection_info.address,
      price_contract: stats.collection_info.price_contract,
      creator: stats.collection_info.creator,
      name: stats.collection_info.name,
      symbol: stats.collection_info.symbol,
      max_supply: stats.collection_info.max_supply,
      initial_price: stats.collection_info.initial_price,
      creator_fee: stats.collection_info.creator_fee,
      base_uri: stats.collection_info.base_uri,
      is_registered: stats.collection_info.is_registered,
      meta_data: stats.collection_info.meta_data,
      created_at: stats.collection_info.created_at,
      block_number: stats.collection_info.block_number,
      current_supply: stats.current_supply,
      total_supply: stats.total_supply,
      owners: stats.owners,
      total_volume: stats.total_volume,
      floor_price: stats.floor_price,
      total_transactions: stats.total_transactions,
      last_updated: stats.last_updated,
      ...(periodVolume && {
        period_volume: periodVolume.volume,
        period_sales: periodVolume.sales,
      }),
    };
  }

  /**
   * Transform transaction data to unified format
   */
  private transformTransactionData(tx: any) {
    return {
      id: tx.id,
      txHash: tx.id,
      collectionId: tx.collection_info.id,
      sender: tx.sender,
      price: tx.price,
      tokenIds: tx.token_ids,
      txType: tx.tx_type,
      createdAt: tx.blockTimestamp,
    };
  }

  /**
   * Retry utility for resilient API calls
   */
  private async withRetry<T>(
    operation: () => Promise<T>,
    maxRetries = 3,
    delay = 1000,
    context = 'operation',
  ): Promise<T> {
    let lastError: Error;

    for (let attempt = 0; attempt <= maxRetries; attempt++) {
      try {
        return await operation();
      } catch (error) {
        lastError = error;

        if (attempt === maxRetries) {
          this.logger.error(
            `${context} failed after ${maxRetries + 1} attempts: ${error.message}`,
          );
          break;
        }

        const waitTime = delay * Math.pow(2, attempt);
        this.logger.warn(
          `${context} failed (attempt ${attempt + 1}/${maxRetries + 1}), retrying in ${waitTime}ms`,
        );
        await new Promise((resolve) => setTimeout(resolve, waitTime));
      }
    }

    throw lastError;
  }

  // ===== Subgraph Methods =====

  /**
   * Get user owned tokens from subgraph - FIXED DOUBLE WRAPPING
   */
  async getUserOwnedTokensFromSubgraph(
    chainId: number,
    userAddress: string,
    page: number,
    pageSize: number,
  ): Promise<{ tokens: any[]; pagination: PaginationInfo }> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    try {
      const { first, skip } = this.calculatePagination(page, pageSize);

      // Get both tokens and summary for total count
      const [result, summary] = await Promise.all([
        this.withRetry(
          () => graphqlAPI.getUserOwnedTokens(userAddress, first, skip),
          2,
          1000,
          'getUserOwnedTokens',
        ),
        this.withRetry(
          () => graphqlAPI.getUserOwnershiSummary(userAddress),
          2,
          1000,
          'getUserOwnershiSummary',
        ),
      ]);

      // Transform the result to match expected format
      const transformedData = result.nftownerships.map((item) => ({
        token_id: item.token_id,
        owner: item.owner,
        collection_id: item.collection_info.id,
        name: item.collection_info.name,
        address: item.collection_info.address,
        price_contract: item.collection_info.price_contract,
        base_uri: item.collection_info.base_uri,
        symbol: item.collection_info.symbol,
        creator: item.collection_info.creator,
        max_supply: item.collection_info.max_supply,
        initial_price: item.collection_info.initial_price,
      }));

      // Calculate total from ownership summary
      const totalTokens = summary.ownershipSummaries.reduce(
        (total, ownership) => total + ownership.nft_count,
        0,
      );

      const pagination = this.createPaginationInfo(
        totalTokens,
        page,
        pageSize,
        transformedData.length,
        skip,
      );

      return { tokens: transformedData, pagination };
    } catch (error) {
      this.logger.error(`Error getting user owned tokens: ${error.message}`);
      throw new Error(`Failed to fetch user tokens: ${error.message}`);
    }
  }

  /**
   * Get user owned tokens by collection from subgraph - FIXED DOUBLE WRAPPING
   */
  async getUserOwnedTokensByCollectionFromSubgraph(
    chainId: number,
    userAddress: string,
    collectionAddress: string,
    page: number,
    pageSize: number,
  ): Promise<{ tokens: any[]; pagination: PaginationInfo }> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    try {
      const { first, skip } = this.calculatePagination(page, pageSize);

      const [result, count] = await Promise.all([
        this.withRetry(
          () =>
            graphqlAPI.getUserOwnedTokens(
              userAddress,
              first,
              skip,
              collectionAddress,
            ),
          2,
          1000,
          'getUserOwnedTokensByCollection',
        ),
        this.withRetry(
          () =>
            graphqlAPI.getUserOwnedTokensCount(userAddress, collectionAddress),
          2,
          1000,
          'getUserOwnedTokensCount',
        ),
      ]);

      const transformedData =
        result.nftownerships.length > 0
          ? result.nftownerships.map((item) => ({
              token_id: item.token_id,
              collection_id: item.collection_info.id,
              name: item.collection_info.name,
              address: item.collection_info.address,
              price_contract: item.collection_info.price_contract,
              base_uri: item.collection_info.base_uri,
            }))
          : [];

      const pagination = this.createPaginationInfo(
        count,
        page,
        pageSize,
        transformedData.length,
        skip,
      );

      return { tokens: transformedData, pagination };
    } catch (error) {
      this.logger.error(
        `Error getting user owned tokens by collection: ${error.message}`,
      );
      throw new Error(
        `Failed to fetch user collection tokens: ${error.message}`,
      );
    }
  }

  /**
   * Get user owned tokens count from subgraph
   */
  async getUserOwnedTokensCountFromSubgraph(
    chainId: number,
    userAddress: string,
    collectionAddress: string,
  ): Promise<number> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    try {
      return await this.withRetry(
        () =>
          graphqlAPI.getUserOwnedTokensCount(userAddress, collectionAddress),
        2,
        1000,
        'getUserOwnedTokensCount',
      );
    } catch (error) {
      this.logger.error(
        `Error getting user owned tokens count: ${error.message}`,
      );
      throw new Error(`Failed to fetch token count: ${error.message}`);
    }
  }

  /**
   * Get user profile from subgraph - OPTIMIZED
   */
  async getUserProfileFromSubgraph(
    chainId: number,
    userAddress: string,
  ): Promise<any> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    try {
      const userOwnershipSummary = await this.withRetry(
        () => graphqlAPI.getUserOwnershiSummary(userAddress),
        2,
        1000,
        'getUserProfile',
      );

      // query collection floor price
      const collectonsWithFloorPrice = await Promise.all(
        userOwnershipSummary.ownershipSummaries.map(async (summary) => {
          const collection = await graphqlAPI.getCollection(
            summary.collection_info.address,
          );
          return {
            ...summary,
            floor_price: collection.collectionStats_collection[0].floor_price,
          };
        }),
      );

      let totalCollections = 0;
      let totalNFTs = 0;
      const collections = [];

      for (const summary of collectonsWithFloorPrice) {
        totalCollections += 1;
        totalNFTs += summary.nft_count;
        collections.push({
          id: summary.collection_info.id,
          address: summary.collection_info.address,
          name: summary.collection_info.name,
          symbol: summary.collection_info.symbol,
          initialPrice: summary.collection_info.initial_price,
          maxSupply: summary.collection_info.max_supply,
          nftCount: summary.nft_count,
          floorPrice: summary.floor_price,
        });
      }

      return {
        totalCollections,
        totalNFTs,
        collections,
      };
    } catch (error) {
      this.logger.error(`Error getting user profile: ${error.message}`);
      throw new Error(`Failed to fetch user profile: ${error.message}`);
    }
  }

  /**
   * Get collection holders from subgraph - FIXED DOUBLE WRAPPING
   */
  async getCollectionHoldersFromSubgraph(
    chainId: number,
    collectionAddress: string,
    page: number,
    pageSize: number,
  ): Promise<{ holders: any[]; pagination: PaginationInfo }> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    try {
      const { first, skip } = this.calculatePagination(page, pageSize);

      const [result, total] = await Promise.all([
        this.withRetry(
          () => graphqlAPI.getCollectionHolders(collectionAddress, first, skip),
          2,
          1000,
          'getCollectionHolders',
        ),
        this.withRetry(
          () => graphqlAPI.getCollectionHoldersCount(collectionAddress),
          2,
          1000,
          'getCollectionHoldersCount',
        ),
      ]);

      const pagination = this.createPaginationInfo(
        total,
        page,
        pageSize,
        result.ownershipSummaries.length,
        skip,
      );

      return { holders: result.ownershipSummaries, pagination };
    } catch (error) {
      this.logger.error(`Error getting collection holders: ${error.message}`);
      throw new Error(`Failed to fetch collection holders: ${error.message}`);
    }
  }

  /**
   * Get collections from subgraph with stats - FIXED DOUBLE WRAPPING
   */
  async getCollectionsFromSubgraph(
    chainId: number,
    first = 100,
    skip = 0,
    period: string = '1h',
  ): Promise<{ collections: any[]; pagination: PaginationInfo }> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    const { fromTimestamp, toTimestamp } = this.getFromAndToTimestamp(period);

    try {
      // Fetch basic data concurrently
      const [result, total] = await Promise.all([
        this.withRetry(
          () => graphqlAPI.getCollections(first, skip),
          2,
          1000,
          'getCollections',
        ),
        this.withRetry(
          () => graphqlAPI.getCollectionsCount(),
          2,
          1000,
          'getCollectionsCount',
        ),
      ]);

      // Early return if no collections
      if (!result.collectionStats_collection.length) {
        const pagination = this.createPaginationInfo(
          total,
          Math.floor(skip / first) + 1,
          first,
          0,
          skip,
        );
        return { collections: [], pagination };
      }

      // Fetch period volumes with individual error handling
      const volumePromises = result.collectionStats_collection.map(
        async (stats) => {
          try {
            return await graphqlAPI.getCollectionPeriodVolumeAndSales(
              stats.collection_info.address,
              fromTimestamp,
              toTimestamp,
            );
          } catch (error) {
            this.logger.warn(
              `Failed to get period volume for ${stats.collection_info.address}: ${error.message}`,
            );
            return { volume: 0, sales: 0 }; // fallback
          }
        },
      );

      const periodVolumes = await Promise.all(volumePromises);

      // Transform data using the helper method
      const collections = result.collectionStats_collection.map(
        (stats, index) =>
          this.transformCollectionStats(stats, periodVolumes[index]),
      );

      const pagination = this.createPaginationInfo(
        total,
        Math.floor(skip / first) + 1,
        first,
        collections.length,
        skip,
      );

      return { collections, pagination };
    } catch (error) {
      this.logger.error(
        `Error getting collections from subgraph: ${error.message}`,
      );
      throw new Error(`Failed to fetch collections: ${error.message}`);
    }
  }

  /**
   * Get single collection from subgraph - OPTIMIZED
   */
  async getCollectionFromSubgraph(
    chainId: number,
    address: string,
  ): Promise<any> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    try {
      const { fromTimestamp, toTimestamp } = this.getFromAndToTimestamp('24h');

      const [collectionResult, periodVolumeAndSales] = await Promise.all([
        this.withRetry(
          () => graphqlAPI.getCollection(address),
          2,
          1000,
          'getCollection',
        ),
        this.withRetry(
          () =>
            graphqlAPI.getCollectionPeriodVolumeAndSales(
              address,
              fromTimestamp,
              toTimestamp,
            ),
          2,
          1000,
          'getCollectionPeriodVolume',
        ),
      ]);

      if (collectionResult.collectionStats_collection.length === 0) {
        return null;
      }

      const collectionStats = collectionResult.collectionStats_collection[0];
      const collectionInfo = collectionStats.collection_info;

      return {
        id: collectionInfo.id,
        address: collectionInfo.address,
        price_contract: collectionInfo.price_contract,
        creator: collectionInfo.creator,
        name: collectionInfo.name,
        symbol: collectionInfo.symbol,
        max_supply: collectionInfo.max_supply,
        max_price: collectionInfo.max_price,
        support_mint: collectionInfo.support_mint,
        gas_limit: collectionInfo.gas_limit,
        initial_price: collectionInfo.initial_price,
        creator_fee: collectionInfo.creator_fee,
        base_uri: collectionInfo.base_uri,
        is_registered: collectionInfo.is_registered,
        meta_data: collectionInfo.meta_data,
        created_at: collectionInfo.created_at,
        current_supply: collectionStats?.current_supply || 0,
        total_supply: collectionStats?.total_supply || 0,
        owners: collectionStats?.owners || 0,
        total_volume: collectionStats?.total_volume || '0',
        volume_1d: periodVolumeAndSales.volume,
        sales_1d: periodVolumeAndSales.sales,
        floor_price:
          collectionStats?.floor_price || collectionInfo.initial_price,
        total_transactions: collectionStats?.total_transactions || 0,
      };
    } catch (error) {
      this.logger.error(
        `Error getting collection from subgraph: ${error.message}`,
      );
      throw new Error(`Failed to fetch collection: ${error.message}`);
    }
  }

  /**
   * Search collections from subgraph by name or address
   */
  async searchCollectionsFromSubgraph(
    chainId: number,
    query: string,
    page: number = 1,
    pageSize: number = 20,
  ): Promise<{ collections: any[]; pagination: PaginationInfo }> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    try {
      const { first, skip } = this.calculatePagination(page, pageSize);

      // Search by both name and address
      const [nameResults, addressResults] = await Promise.all([
        this.withRetry(
          () => graphqlAPI.searchCollectionsByName(query, first, skip),
          2,
          1000,
          'searchCollectionsByName',
        ),
        this.withRetry(
          () => graphqlAPI.searchCollectionsByAddress(query, first, skip),
          2,
          1000,
          'searchCollectionsByAddress',
        ),
      ]);

      // Merge and deduplicate results
      const collectionsMap = new Map();

      // Add name search results
      if (nameResults.collectionStats_collection) {
        nameResults.collectionStats_collection.forEach((collection) => {
          collectionsMap.set(collection.collection_info.address, collection);
        });
      }

      // Add address search results (will overwrite duplicates)
      if (addressResults.collectionStats_collection) {
        addressResults.collectionStats_collection.forEach((collection) => {
          collectionsMap.set(collection.collection_info.address, collection);
        });
      }

      const collections = Array.from(collectionsMap.values()).slice(0, first);

      // Get period volume for each collection
      const { fromTimestamp, toTimestamp } = this.getFromAndToTimestamp('24h');
      const volumePromises = collections.map(async (stats) => {
        try {
          return await graphqlAPI.getCollectionPeriodVolumeAndSales(
            stats.collection_info.address,
            fromTimestamp,
            toTimestamp,
          );
        } catch (error) {
          this.logger.warn(
            `Failed to get period volume for ${stats.collection_info.address}: ${error.message}`,
          );
          return { volume: 0, sales: 0 };
        }
      });

      const periodVolumes = await Promise.all(volumePromises);

      // Transform data
      const transformedCollections = collections.map((stats, index) =>
        this.transformCollectionStats(stats, periodVolumes[index]),
      );

      const pagination = this.createPaginationInfo(
        collections.length,
        page,
        pageSize,
        transformedCollections.length,
        skip,
      );

      return { collections: transformedCollections, pagination };
    } catch (error) {
      this.logger.error(`Error searching collections: ${error.message}`);
      throw new Error(`Failed to search collections: ${error.message}`);
    }
  }

  /**
   * Get recommended collection from subgraph - OPTIMIZED
   */
  async getRecommendedCollectionFromSubgraph(chainId: number): Promise<any[]> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    try {
      const result = await this.withRetry(
        () =>
          graphqlAPI.getCollections(1, 0, 'collection_info__created_at', 'asc'),
        2,
        1000,
        'getRecommendedCollection',
      );

      const collectionStats = result.collectionStats_collection[0];
      const collectionInfo = collectionStats.collection_info;
      const { fromTimestamp, toTimestamp } = this.getFromAndToTimestamp('24h');

      const periodVolumeAndSales = await this.withRetry(
        () =>
          graphqlAPI.getCollectionPeriodVolumeAndSales(
            collectionInfo.address,
            fromTimestamp,
            toTimestamp,
          ),
        2,
        1000,
        'getRecommendedCollectionVolume',
      );

      return [
        {
          id: collectionInfo.id,
          address: collectionInfo.address,
          name: collectionInfo.name,
          symbol: collectionInfo.symbol,
          image: collectionInfo.base_uri,
          period_volume: periodVolumeAndSales.volume,
          period_sales: periodVolumeAndSales.sales,
          total_volume: collectionStats.total_volume,
          floor_price: collectionStats.floor_price,
          total_transactions: collectionStats.total_transactions,
          last_updated: collectionStats.last_updated,
          created_at: collectionInfo.created_at,
          block_number: collectionInfo.block_number,
          current_supply: collectionStats.current_supply,
          total_supply: collectionStats.total_supply,
          owners: collectionStats.owners,
          is_registered: collectionInfo.is_registered,
          meta_data: collectionInfo.meta_data,
          creator: collectionInfo.creator,
          creator_fee: collectionInfo.creator_fee,
          initial_price: collectionInfo.initial_price,
          max_supply: collectionInfo.max_supply,
          price_contract: collectionInfo.price_contract,
        },
      ];
    } catch (error) {
      this.logger.error(
        `Error getting recommended collection: ${error.message}`,
      );
      throw new Error(
        `Failed to fetch recommended collection: ${error.message}`,
      );
    }
  }

  /**
   * Get transactions from subgraph - FIXED DOUBLE WRAPPING
   */
  async getTxsFromSubgraph(
    chainId: number,
    userAddress?: string,
    collectionAddress?: string,
    txTypes?: number[],
    page: number = 1,
    pageSize: number = 20,
    tokenID?: string,
  ): Promise<{ transactions: any[]; pagination: PaginationInfo }> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    try {
      const { safePage, safePageSize, skip } = this.calculatePagination(
        page,
        pageSize,
        1000,
      );

      // Fetch data concurrently with individual error handling
      const [txsResult, totalCount] = await Promise.allSettled([
        this.withRetry(
          () =>
            graphqlAPI.getTxs(
              userAddress,
              collectionAddress,
              txTypes,
              safePageSize,
              skip,
              tokenID,
            ),
          2,
          1000,
          'getTxs',
        ),
        this.withRetry(
          () =>
            graphqlAPI.getTxsCount(
              userAddress,
              collectionAddress,
              txTypes,
              tokenID,
            ),
          2,
          1000,
          'getTxsCount',
        ),
      ]);

      // Handle potential failures
      if (txsResult.status === 'rejected') {
        throw new Error(`Failed to fetch transactions: ${txsResult.reason}`);
      }

      const transactions = txsResult.value.txs_collection.map(
        this.transformTransactionData,
      );
      const total =
        totalCount.status === 'fulfilled'
          ? totalCount.value
          : transactions.length;

      if (totalCount.status === 'rejected') {
        this.logger.warn(
          'Failed to get total count, using current data length as estimate',
        );
      }

      const pagination = this.createPaginationInfo(
        total,
        safePage,
        safePageSize,
        transactions.length,
        skip,
      );

      return { transactions, pagination };
    } catch (error) {
      this.logger.error(`Error getting txs from subgraph: ${error.message}`);
      throw new Error(`Failed to fetch transactions: ${error.message}`);
    }
  }

  async getCrossChainTransferStatusFromSubgraph(
    chainId: number,
    fluxusContract: string,
    page: number = 1,
    pageSize: number = 10,
    sender?: string,
    receiver?: string,
  ): Promise<{ statuses: any[]; pagination: PaginationInfo }> {
    const graphqlAPI = await this.getGraphqlAPI(chainId);
    try {
      const { safePage, safePageSize, skip, first } = this.calculatePagination(
        page,
        pageSize,
        100,
      );

      // If both sender and receiver provided, emulate OR by merging two queries
      if (sender && receiver) {
        const [
          senderResult,
          receiverResult,
          bothCount,
          senderCount,
          receiverCount,
        ] = await Promise.all([
          this.withRetry(
            () =>
              graphqlAPI.getCrossChainTransferStatus(
                fluxusContract,
                Math.ceil(first / 2),
                skip,
                sender,
                undefined,
              ),
            2,
            1000,
            'getCrossChainTransferStatus_sender',
          ),
          this.withRetry(
            () =>
              graphqlAPI.getCrossChainTransferStatus(
                fluxusContract,
                Math.floor(first / 2),
                skip,
                undefined,
                receiver,
              ),
            2,
            1000,
            'getCrossChainTransferStatus_receiver',
          ),
          this.withRetry(
            () =>
              graphqlAPI.getCrossChainTransferStatusBothCount(
                fluxusContract,
                sender,
                receiver,
              ),
            2,
            1000,
            'getCrossChainTransferStatusBothCount',
          ),
          this.withRetry(
            () =>
              graphqlAPI.getCrossChainTransferStatusCount(
                fluxusContract,
                sender,
                undefined,
              ),
            2,
            1000,
            'getCrossChainTransferStatusCount_sender',
          ),
          this.withRetry(
            () =>
              graphqlAPI.getCrossChainTransferStatusCount(
                fluxusContract,
                undefined,
                receiver,
              ),
            2,
            1000,
            'getCrossChainTransferStatusCount_receiver',
          ),
        ]);

        // Merge and de-duplicate by id
        const map = new Map<string, any>();
        for (const item of senderResult?.crossChainStatuses ?? []) {
          map.set(item.id, item);
        }
        for (const item of receiverResult?.crossChainStatuses ?? []) {
          if (!map.has(item.id)) map.set(item.id, item);
        }
        const merged = Array.from(map.values());

        const total = senderCount + receiverCount - bothCount;
        const pagination = this.createPaginationInfo(
          total,
          safePage,
          safePageSize,
          merged.length,
          skip,
        );
        return { statuses: merged, pagination };
      }

      const [result, total] = await Promise.all([
        this.withRetry(
          () =>
            graphqlAPI.getCrossChainTransferStatus(
              fluxusContract,
              first,
              skip,
              sender,
              receiver,
            ),
          2,
          1000,
          'getCrossChainTransferStatus',
        ),
        this.withRetry(
          () =>
            graphqlAPI.getCrossChainTransferStatusCount(
              fluxusContract,
              sender,
              receiver,
            ),
          2,
          1000,
          'getCrossChainTransferStatusCount',
        ),
      ]);

      const statuses = result?.crossChainStatuses ?? [];

      const pagination = this.createPaginationInfo(
        total,
        safePage,
        safePageSize,
        statuses.length,
        skip,
      );

      return { statuses, pagination };
    } catch (error) {
      this.logger.error(
        `Error getting cross chain transfer status: ${error.message}`,
      );
      throw new Error(
        `Failed to fetch cross chain transfer status: ${error.message}`,
      );
    }
  }

  // ===== Helper Methods =====

  private calculatePagination(
    page: number,
    pageSize: number,
    maxPageSize: number = 100,
  ) {
    const safePage = Math.max(1, page);
    const safePageSize = Math.min(maxPageSize, Math.max(1, pageSize));
    const skip = (safePage - 1) * safePageSize;
    return {
      safePage,
      safePageSize,
      skip,
      first: safePageSize,
    };
  }

  private getFromAndToTimestamp(period: string): {
    fromTimestamp: number;
    toTimestamp: number;
  } {
    const now = new Date();
    const hours: { [key: string]: number } = {
      '1h': 1,
      '24h': 24,
      '7d': 7 * 24,
    };

    const hourCount = hours[period] || 1;
    const fromTimestamp = now.getTime() - hourCount * 60 * 60 * 1000;
    const toTimestamp = now.getTime();

    return { fromTimestamp, toTimestamp };
  }
}
