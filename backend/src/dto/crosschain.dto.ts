import { IsOptional, IsInt, Min, Max, IsString } from 'class-validator';
import { ApiPropertyOptional } from '@nestjs/swagger';
import { Transform, Type } from 'class-transformer';

export class CrossChainStatusQueryDto {
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
    description: 'Items per page',
    default: 10,
    minimum: 1,
    maximum: 100,
  })
  @IsOptional()
  @Type(() => Number)
  @IsInt()
  @Min(1)
  @Max(100)
  pageSize: number = 10;

  @ApiPropertyOptional({ description: 'Filter by sender address' })
  @IsOptional()
  @IsString()
  @Transform(({ value }) => value?.toLowerCase())
  sender?: string;

  @ApiPropertyOptional({ description: 'Filter by receiver address' })
  @IsOptional()
  @IsString()
  @Transform(({ value }) => value?.toLowerCase())
  receiver?: string;
}
