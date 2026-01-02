import { IsString, IsNotEmpty, IsOptional, IsNumber } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class SignatureMessageDto {
  @ApiProperty({ description: 'Wallet address to get signature message for' })
  @IsString()
  @IsNotEmpty()
  walletAddress: string;
}

export class VerifySignatureDto {
  @ApiProperty({ description: 'Wallet address' })
  @IsString()
  @IsNotEmpty()
  walletAddress: string;

  @ApiProperty({ description: 'Signature from wallet' })
  @IsString()
  @IsNotEmpty()
  signature: string;

  @ApiProperty({ description: 'Message that was signed' })
  @IsString()
  @IsNotEmpty()
  message: string;

  @ApiProperty({
    description: 'Collection address for NFT ownership verification',
  })
  @IsString()
  @IsNotEmpty()
  collectionAddress: string;

  @ApiProperty({
    description: 'Chain ID for the collection',
  })
  @IsNumber()
  @IsNotEmpty()
  chainId: number;
}

export class VerifyTokenDto {
  @ApiProperty({ description: 'JWT token to verify' })
  @IsString()
  @IsNotEmpty()
  token: string;
}

export class AuthResponseDto {
  @ApiProperty({ description: 'Success status' })
  success: boolean;

  @ApiProperty({
    description: 'JWT token if authentication successful',
    required: false,
  })
  token?: string;

  @ApiProperty({ description: 'User information', required: false })
  user?: {
    id: number;
    walletAddress: string;
  };

  @ApiProperty({ description: 'Response message' })
  message: string;
}
