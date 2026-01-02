import {
  Controller,
  Get,
  Post,
  Body,
  Query,
  HttpStatus,
  HttpException,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiQuery } from '@nestjs/swagger';
import { AuthService } from './auth.service';
import { VerifySignatureDto, AuthResponseDto } from '../dto/auth.dto';

@Controller('api/v1/auth')
@ApiTags('Authentication API')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  /**
   * Get signature message for wallet authentication
   */
  @Get('signature-message')
  @ApiOperation({ summary: 'Get signature message for wallet authentication' })
  @ApiQuery({
    name: 'walletAddress',
    description: 'Wallet address to get signature message for',
  })
  @ApiResponse({
    status: 200,
    description: 'Signature message retrieved successfully',
    schema: {
      type: 'object',
      properties: {
        success: { type: 'boolean' },
        message: { type: 'string' },
        nonce: { type: 'string' },
      },
    },
  })
  async getSignatureMessage(@Query('walletAddress') walletAddress: string) {
    try {
      if (!walletAddress) {
        throw new HttpException(
          'Wallet address is required',
          HttpStatus.BAD_REQUEST,
        );
      }

      const result =
        await this.authService.generateSignatureMessage(walletAddress);

      return {
        success: true,
        message: result.message,
        nonce: result.nonce,
      };
    } catch (error) {
      throw new HttpException(
        `Failed to generate signature message: ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  /**
   * Verify wallet signature and authenticate user
   */
  @Post('verify-signature')
  @ApiOperation({ summary: 'Verify wallet signature and authenticate user' })
  @ApiResponse({
    status: 200,
    description: 'Signature verified successfully',
    type: AuthResponseDto,
  })
  @ApiResponse({
    status: 401,
    description: 'Invalid signature or user does not own NFT from collection',
  })
  async verifySignature(@Body() verifySignatureDto: VerifySignatureDto) {
    try {
      const result = await this.authService.verifySignature(
        verifySignatureDto.walletAddress,
        verifySignatureDto.signature,
        verifySignatureDto.message,
        verifySignatureDto.collectionAddress,
        verifySignatureDto.chainId,
      );

      return {
        success: true,
        token: result.token,
        user: {
          id: result.user.id,
          walletAddress: result.user.wallet_address,
        },
        message: 'Authentication successful',
      };
    } catch (error) {
      if (error instanceof HttpException) {
        throw error;
      }
      throw new HttpException(
        `Authentication failed: ${error.message}`,
        HttpStatus.UNAUTHORIZED,
      );
    }
  }

  /**
   * Verify JWT token
   */
  @Get('verify-token')
  @ApiOperation({ summary: 'Verify JWT token' })
  @ApiQuery({ name: 'token', description: 'JWT token to verify' })
  @ApiResponse({
    status: 200,
    description: 'Token verified successfully',
    schema: {
      type: 'object',
      properties: {
        success: { type: 'boolean' },
        user: {
          type: 'object',
          properties: {
            id: { type: 'number' },
            walletAddress: { type: 'string' },
          },
        },
        message: { type: 'string' },
      },
    },
  })
  async verifyToken(@Query('token') token: string) {
    try {
      if (!token) {
        throw new HttpException('Token is required', HttpStatus.BAD_REQUEST);
      }

      const payload = await this.authService.verifyToken(token);

      return {
        success: true,
        user: {
          id: payload.userId,
          walletAddress: payload.walletAddress,
        },
        message: 'Token verified successfully',
      };
    } catch (error) {
      throw new HttpException(
        `Token verification failed: ${error.message}`,
        HttpStatus.UNAUTHORIZED,
      );
    }
  }
}
