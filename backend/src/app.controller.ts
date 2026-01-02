import {
  Controller,
  Get,
  Param,
  Query,
  HttpStatus,
  HttpException,
  ValidationPipe,
  UsePipes,
} from '@nestjs/common';
import {
  ApiTags,
  ApiOperation,
  ApiResponse as SwaggerResponse,
  ApiParam,
} from '@nestjs/swagger';
import { AppService } from './app.service';

import { PaginationDto } from './dto/common.dto';
import { CrossChainStatusQueryDto } from './dto/crosschain.dto';
import {
  GetCollectionsDto,
  TransactionType,
  GetCollectionTransactionsDto,
  SearchCollectionsDto,
} from './dto/collection.dto';

// Response Interface
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

@Controller('api/v1')
@ApiTags('NFT Collections API')
@UsePipes(new ValidationPipe({ transform: true, whitelist: true }))
export class AppController {
  constructor(private readonly appService: AppService) {}

  // Transaction type mapping - moved from controller logic
  private readonly TRANSACTION_TYPE_MAPPING = {
    [TransactionType.MINT]: [1, 6], // mint and bulkMint
    [TransactionType.BUY]: [2, 4], // buy and bulkBuy
    [TransactionType.SELL]: [3, 5], // sell and bulkSell
  };

  private createSuccessResponse<T>(
    data: T,
    message?: string,
    pagination?: any,
  ): ApiResponse<T> {
    return {
      success: true,
      data,
      message,
      pagination,
    };
  }

  // Collections Endpoints
  @Get('collections')
  @ApiOperation({ summary: 'Get collections with statistics' })
  @SwaggerResponse({
    status: 200,
    description: 'Collections retrieved successfully',
  })
  async getCollections(@Query() query: GetCollectionsDto) {
    try {
      const result = await this.appService.getCollectionsFromSubgraph(
        query.chainId,
        query.first,
        query.skip,
        query.period,
      );

      return this.createSuccessResponse(
        result.collections,
        'Collections retrieved successfully',
        result.pagination,
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve collections, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('collections/recommended')
  @ApiOperation({ summary: 'Get recommended collections' })
  @SwaggerResponse({
    status: 200,
    description: 'Recommended collections retrieved successfully',
  })
  async getRecommendedCollections(@Query() query: { chainId: number }) {
    try {
      const result = await this.appService.getRecommendedCollectionFromSubgraph(
        query.chainId,
      );
      return this.createSuccessResponse(
        result,
        'Recommended collections retrieved successfully',
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve recommended collections, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('collections/search')
  @ApiOperation({ summary: 'Search collections by name or address' })
  @SwaggerResponse({
    status: 200,
    description: 'Collections search results retrieved successfully',
  })
  async searchCollections(@Query() query: SearchCollectionsDto) {
    try {
      if (!query.query || query.query.trim().length === 0) {
        return this.createSuccessResponse([], 'No search query provided');
      }

      const result = await this.appService.searchCollectionsFromSubgraph(
        query.chainId,
        query.query.trim(),
        query.page,
        query.pageSize,
      );

      return this.createSuccessResponse(
        result.collections,
        'Collections search results retrieved successfully',
        result.pagination,
      );
    } catch (error) {
      throw new HttpException(
        `Failed to search collections, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('collections/:address')
  @ApiOperation({ summary: 'Get collection details by address' })
  @ApiParam({ name: 'address', description: 'Collection contract address' })
  @SwaggerResponse({
    status: 200,
    description: 'Collection details retrieved successfully',
  })
  async getCollection(
    @Param('address') address: string,
    @Query() query: { chainId: number },
  ) {
    try {
      const result = await this.appService.getCollectionFromSubgraph(
        query.chainId,
        address.toLowerCase(),
      );
      return this.createSuccessResponse(
        result,
        'Collection details retrieved successfully',
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve collection details, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('collections/:collectionAddress/transactions')
  @ApiOperation({ summary: 'Get collection transactions' })
  @ApiParam({
    name: 'collectionAddress',
    description: 'Collection contract address',
  })
  @SwaggerResponse({
    status: 200,
    description: 'Collection transactions retrieved successfully',
  })
  async getCollectionTransactions(
    @Param('collectionAddress') collectionAddress: string,
    @Query() query: GetCollectionTransactionsDto,
  ) {
    try {
      const txTypes = query.type
        ? this.TRANSACTION_TYPE_MAPPING[query.type]
        : undefined;

      const result = await this.appService.getTxsFromSubgraph(
        query.chainId,
        undefined,
        collectionAddress.toLowerCase(),
        txTypes,
        query.page,
        query.pageSize,
        query.tokenID,
      );

      return this.createSuccessResponse(
        result.transactions,
        'Collection transactions retrieved successfully',
        result.pagination,
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve collection transactions, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('collections/:collectionAddress/holders')
  @ApiOperation({ summary: 'Get collection holders' })
  @ApiParam({
    name: 'collectionAddress',
    description: 'Collection contract address',
  })
  @SwaggerResponse({
    status: 200,
    description: 'Collection holders retrieved successfully',
  })
  async getCollectionHolders(
    @Param('collectionAddress') collectionAddress: string,
    @Query() query: PaginationDto,
  ) {
    try {
      const result = await this.appService.getCollectionHoldersFromSubgraph(
        query.chainId,
        collectionAddress.toLowerCase(),
        query.page,
        query.pageSize,
      );

      return this.createSuccessResponse(
        result.holders,
        'Collection holders retrieved successfully',
        result.pagination,
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve collection holders, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  // Users Endpoints
  @Get('users/:userAddress/profile')
  @ApiOperation({ summary: 'Get user profile' })
  @ApiParam({ name: 'userAddress', description: 'User wallet address' })
  @SwaggerResponse({
    status: 200,
    description: 'User profile retrieved successfully',
  })
  async getUserProfile(
    @Param('userAddress') userAddress: string,
    @Query() query: { chainId: number },
  ) {
    try {
      const result = await this.appService.getUserProfileFromSubgraph(
        query.chainId,
        userAddress.toLowerCase(),
      );
      return this.createSuccessResponse(
        result,
        'User profile retrieved successfully',
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve user profile, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('users/:userAddress/transactions')
  @ApiOperation({ summary: 'Get user transactions' })
  @ApiParam({ name: 'userAddress', description: 'User wallet address' })
  @SwaggerResponse({
    status: 200,
    description: 'User transactions retrieved successfully',
  })
  async getUserTransactions(
    @Param('userAddress') userAddress: string,
    @Query() query: PaginationDto,
  ) {
    try {
      const result = await this.appService.getTxsFromSubgraph(
        query.chainId,
        userAddress.toLowerCase(),
        undefined,
        undefined,
        query.page,
        query.pageSize,
      );

      return this.createSuccessResponse(
        result.transactions,
        'User transactions retrieved successfully',
        result.pagination,
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve user transactions, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('users/:userAddress/tokens')
  @ApiOperation({ summary: 'Get all tokens owned by user' })
  @ApiParam({ name: 'userAddress', description: 'User wallet address' })
  @SwaggerResponse({
    status: 200,
    description: 'User tokens retrieved successfully',
  })
  async getUserTokens(
    @Param('userAddress') userAddress: string,
    @Query() query: PaginationDto,
  ) {
    try {
      const result = await this.appService.getUserOwnedTokensFromSubgraph(
        query.chainId,
        userAddress.toLowerCase(),
        query.page,
        query.pageSize,
      );

      return this.createSuccessResponse(
        result.tokens,
        'User tokens retrieved successfully',
        result.pagination,
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve user tokens, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('users/:userAddress/collections/:collectionAddress/tokens')
  @ApiOperation({ summary: 'Get user tokens in specific collection' })
  @ApiParam({ name: 'userAddress', description: 'User wallet address' })
  @ApiParam({
    name: 'collectionAddress',
    description: 'Collection contract address',
  })
  @SwaggerResponse({
    status: 200,
    description: 'User collection tokens retrieved successfully',
  })
  async getUserCollectionTokens(
    @Param('userAddress') userAddress: string,
    @Param('collectionAddress') collectionAddress: string,
    @Query() query: PaginationDto,
  ) {
    try {
      const result =
        await this.appService.getUserOwnedTokensByCollectionFromSubgraph(
          query.chainId,
          userAddress.toLowerCase(),
          collectionAddress.toLowerCase(),
          query.page,
          query.pageSize,
        );

      return this.createSuccessResponse(
        result.tokens,
        'User collection tokens retrieved successfully',
        result.pagination,
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve user collection tokens, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('users/:userAddress/collections/:collectionAddress/transactions')
  @ApiOperation({ summary: 'Get user transactions in specific collection' })
  @ApiParam({ name: 'userAddress', description: 'User wallet address' })
  @ApiParam({
    name: 'collectionAddress',
    description: 'Collection contract address',
  })
  @SwaggerResponse({
    status: 200,
    description: 'User collection transactions retrieved successfully',
  })
  async getUserCollectionTransactions(
    @Param('userAddress') userAddress: string,
    @Param('collectionAddress') collectionAddress: string,
    @Query() query: PaginationDto,
  ) {
    try {
      const result = await this.appService.getTxsFromSubgraph(
        query.chainId,
        userAddress.toLowerCase(),
        collectionAddress.toLowerCase(),
        undefined,
        query.page,
        query.pageSize,
      );

      return this.createSuccessResponse(
        result.transactions,
        'User collection transactions retrieved successfully',
        result.pagination,
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve user collection transactions, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('users/:userAddress/collections/:collectionAddress/tokens/count')
  @ApiOperation({ summary: 'Get count of tokens owned by user in collection' })
  @ApiParam({ name: 'userAddress', description: 'User wallet address' })
  @ApiParam({
    name: 'collectionAddress',
    description: 'Collection contract address',
  })
  @SwaggerResponse({
    status: 200,
    description: 'Token count retrieved successfully',
  })
  async getUserOwnedTokensCount(
    @Param('userAddress') userAddress: string,
    @Param('collectionAddress') collectionAddress: string,
    @Query() query: { chainId: number },
  ) {
    try {
      const result = await this.appService.getUserOwnedTokensCountFromSubgraph(
        query.chainId,
        userAddress.toLowerCase(),
        collectionAddress.toLowerCase(),
      );

      return this.createSuccessResponse(
        result,
        'Token count retrieved successfully',
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve token count, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  @Get('crosschain/status/:fluxusContract')
  @ApiOperation({ summary: 'Get cross chain transfer status' })
  @ApiParam({ name: 'fluxusContract', description: 'Fluxus contract address' })
  @SwaggerResponse({
    status: 200,
    description: 'Cross chain transfer status retrieved successfully',
  })
  async getCrossChainTransferStatus(
    @Param('fluxusContract') fluxusContract: string,
    @Query() query: CrossChainStatusQueryDto,
  ) {
    try {
      const result =
        await this.appService.getCrossChainTransferStatusFromSubgraph(
          query.chainId,
          fluxusContract,
          query.page ?? 1,
          query.pageSize ?? 10,
          query.sender,
          query.receiver,
        );
      return this.createSuccessResponse(
        result.statuses,
        'Cross chain transfer status retrieved successfully',
        result.pagination,
      );
    } catch (error) {
      throw new HttpException(
        `Failed to retrieve cross chain transfer status, ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }
}
