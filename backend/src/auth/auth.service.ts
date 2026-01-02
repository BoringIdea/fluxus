import {
  Injectable,
  UnauthorizedException,
  BadRequestException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { ethers } from 'ethers';
import { eq, and } from 'drizzle-orm';
import { db } from '../db/db';
import { users, userTokens, collections } from '../db/schema';
import { AppService } from '../app.service';

@Injectable()
export class AuthService {
  constructor(
    private jwtService: JwtService,
    private appService: AppService,
  ) {}

  /**
   * Generate signature message for wallet authentication
   */
  async generateSignatureMessage(
    walletAddress: string,
  ): Promise<{ message: string; nonce: string }> {
    const nonce = Math.random().toString(36).substring(2, 15);
    const message = `Sign this message to authenticate with Fluxus Chat.\nNonce: ${nonce}\nTimestamp: ${Date.now()}`;

    // Store nonce for verification
    await this.upsertUser(walletAddress, nonce);

    return { message, nonce };
  }

  /**
   * Verify wallet signature and NFT ownership
   */
  async verifySignature(
    walletAddress: string,
    signature: string,
    message: string,
    collectionAddress: string,
    chainId: number,
  ): Promise<{ token: string; user: any }> {
    try {
      if (!signature || typeof signature !== 'string') {
        throw new UnauthorizedException(
          'Signature should not be empty, signature must be a string',
        );
      }

      // Verify signature
      const recoveredAddress = ethers.verifyMessage(message, signature);
      if (recoveredAddress.toLowerCase() !== walletAddress.toLowerCase()) {
        throw new UnauthorizedException('Invalid signature');
      }

      // Verify NFT ownership
      const hasNFT = await this.verifyNFTOwnership(
        walletAddress,
        collectionAddress,
        chainId,
      );
      if (!hasNFT) {
        throw new UnauthorizedException(
          'User does not own NFT from this collection',
        );
      }

      // Get or create user
      const user = await this.upsertUser(walletAddress);

      // Generate JWT token
      const token = this.jwtService.sign(
        {
          userId: user.id,
          walletAddress: user.wallet_address,
          collectionAddress,
        },
        { expiresIn: '24h' },
      );

      // Store token in database
      await this.storeToken(user.id, token);

      return { token, user };
    } catch (error) {
      if (error instanceof UnauthorizedException) {
        throw error;
      }
      throw new BadRequestException('Signature verification failed');
    }
  }

  /**
   * Verify JWT token
   */
  async verifyToken(token: string): Promise<any> {
    try {
      const payload = this.jwtService.verify(token);

      // Check if token exists in database and is active
      const tokenRecord = await db
        .select()
        .from(userTokens)
        .where(
          and(
            eq(userTokens.user_id, payload.userId),
            eq(userTokens.is_active, true),
          ),
        )
        .limit(1);

      if (tokenRecord.length === 0) {
        throw new UnauthorizedException('Token not found or inactive');
      }

      // Check if token is expired
      if (new Date() > new Date(tokenRecord[0].expires_at)) {
        throw new UnauthorizedException('Token expired');
      }

      return payload;
    } catch {
      throw new UnauthorizedException('Invalid token');
    }
  }

  /**
   * Verify NFT ownership using existing subgraph service
   */
  private async verifyNFTOwnership(
    walletAddress: string,
    collectionAddress: string,
    chainId: number,
  ): Promise<boolean> {
    try {
      // Get collection info from GraphQL API
      const collection = await this.appService.getCollectionFromSubgraph(
        chainId,
        collectionAddress.toLowerCase(),
      );

      if (!collection) {
        return false;
      }

      // Use existing service to check NFT ownership
      const result =
        await this.appService.getUserOwnedTokensByCollectionFromSubgraph(
          chainId,
          walletAddress.toLowerCase(),
          collectionAddress.toLowerCase(),
          1,
          1,
        );

      return result.tokens.length > 0;
    } catch (error) {
      console.error('Error verifying NFT ownership:', error);
      return false;
    }
  }

  /**
   * Get or create user
   */
  private async upsertUser(
    walletAddress: string,
    nonce?: string,
  ): Promise<any> {
    const existingUser = await db
      .select()
      .from(users)
      .where(eq(users.wallet_address, walletAddress.toLowerCase()))
      .limit(1);

    if (existingUser.length > 0) {
      // Update last active time
      const updateData: any = {
        last_active_at: new Date(),
      };
      if (nonce) {
        updateData.nonce = nonce;
      }
      await db
        .update(users)
        .set(updateData)
        .where(eq(users.id, existingUser[0].id));

      return existingUser[0];
    }

    // Create new user
    const insertData: any = {
      wallet_address: walletAddress.toLowerCase(),
      last_active_at: new Date(),
    };
    if (nonce) {
      insertData.nonce = nonce;
    }
    const newUser = await db.insert(users).values(insertData).returning();

    return newUser[0];
  }

  /**
   * Store JWT token in database
   */
  private async storeToken(userId: number, token: string): Promise<void> {
    const expiresAt = new Date();
    expiresAt.setHours(expiresAt.getHours() + 24); // 24 hours

    const tokenData: any = {
      user_id: userId,
      token_hash: token, // In production, hash the token
      expires_at: expiresAt,
      is_active: true,
    };
    await db.insert(userTokens).values(tokenData);
  }
}
