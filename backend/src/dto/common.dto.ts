import { IsOptional, IsInt, Min, Max } from 'class-validator';
import { Transform, Type } from 'class-transformer';
import { ApiPropertyOptional } from '@nestjs/swagger';

export class PaginationDto {
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
    description: 'Page number (1-based)',
    default: 1,
    minimum: 1,
  })
  @IsOptional()
  @Type(() => Number)
  @IsInt()
  @Min(1)
  page: number = 1;

  @ApiPropertyOptional({
    description: 'Number of items per page',
    default: 20,
    minimum: 1,
    maximum: 100,
  })
  @IsOptional()
  @Type(() => Number)
  @IsInt()
  @Min(1)
  @Max(100)
  pageSize: number = 20;
}

export class AddressParamDto {
  @Transform(({ value }) => value?.toLowerCase())
  address: string;
}

export class UserAddressParamDto {
  @Transform(({ value }) => value?.toLowerCase())
  userAddress: string;
}

export class CollectionAddressParamDto {
  @Transform(({ value }) => value?.toLowerCase())
  collectionAddress: string;
}
