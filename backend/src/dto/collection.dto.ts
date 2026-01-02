import { IsOptional, IsEnum, IsString, IsInt, Min, Max } from 'class-validator';
import { Transform, Type } from 'class-transformer';
import { ApiPropertyOptional } from '@nestjs/swagger';
import { PaginationDto } from './common.dto';

export enum TimePeriod {
  ONE_HOUR = '1h',
  TWENTY_FOUR_HOURS = '24h',
  SEVEN_DAYS = '7d',
}

export enum TransactionType {
  MINT = 'mint',
  BUY = 'buy',
  SELL = 'sell',
}

export class GetCollectionsDto extends PaginationDto {
  @ApiPropertyOptional({
    description: 'Chain ID',
    default: 84532,
    minimum: 1,
  })
  @IsOptional()
  @Type(() => Number)
  @IsInt()
  @Min(1)
  chainId: number = 84532;

  @ApiPropertyOptional({
    description: 'Number of collections to fetch',
    default: 100,
    minimum: 1,
    maximum: 1000,
  })
  @IsOptional()
  @Type(() => Number)
  @IsInt()
  @Min(1)
  @Max(1000)
  first: number = 100;

  @ApiPropertyOptional({
    description: 'Number of collections to skip',
    default: 0,
    minimum: 0,
  })
  @IsOptional()
  @Type(() => Number)
  @IsInt()
  @Min(0)
  skip: number = 0;

  @ApiPropertyOptional({
    description: 'Time period for volume calculation',
    enum: TimePeriod,
    default: TimePeriod.ONE_HOUR,
  })
  @IsOptional()
  @IsEnum(TimePeriod)
  period: TimePeriod = TimePeriod.ONE_HOUR;
}

export class GetCollectionTransactionsDto extends PaginationDto {
  @ApiPropertyOptional({
    description: 'Filter by transaction type',
    enum: TransactionType,
  })
  @IsOptional()
  @IsEnum(TransactionType)
  type?: TransactionType;

  @ApiPropertyOptional({
    description: 'Filter by specific token ID',
    type: 'string',
  })
  @IsOptional()
  @IsString()
  tokenID?: string;
}

export class GetUserOwnedTokensCountDto {
  @Transform(({ value }) => value?.toLowerCase())
  @IsString()
  userAddress: string;

  @Transform(({ value }) => value?.toLowerCase())
  @IsString()
  collectionAddress: string;
}

export class SearchCollectionsDto extends PaginationDto {
  @ApiPropertyOptional({
    description: 'Chain ID',
    default: 84532,
    minimum: 1,
  })
  @IsOptional()
  @Type(() => Number)
  @IsInt()
  @Min(1)
  chainId: number = 84532;

  @ApiPropertyOptional({
    description: 'Search query for collection name or address',
    type: 'string',
  })
  @IsOptional()
  @IsString()
  query?: string;
}
