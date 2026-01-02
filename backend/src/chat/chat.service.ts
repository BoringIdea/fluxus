import { Injectable, NotFoundException } from '@nestjs/common';
import { eq, and, desc, sql } from 'drizzle-orm';
import { db } from '../db/db';
import {
  users,
  chatRooms,
  chatRoomMembers,
  chatMessages,
  onlineUsers,
} from '../db/schema';
import { AppService } from '../app.service';

@Injectable()
export class ChatService {
  constructor(private readonly appService: AppService) {}

  /**
   * Get chat room by collection address only
   */
  async getChatRoomByAddress(collectionAddress: string): Promise<any> {
    const room = await db
      .select()
      .from(chatRooms)
      .where(eq(chatRooms.collection_address, collectionAddress.toLowerCase()))
      .limit(1);

    if (!room[0]) {
      throw new NotFoundException('Chat room not found');
    }

    return room[0];
  }

  /**
   * Get or create chat room for collection
   */
  async getOrCreateChatRoom(
    collectionAddress: string,
    chainId: number,
  ): Promise<any> {
    // Get collection info from GraphQL API
    const collection = await this.appService.getCollectionFromSubgraph(
      chainId,
      collectionAddress.toLowerCase(),
    );

    if (!collection) {
      throw new NotFoundException('Collection not found');
    }

    // Find existing chat room by collection address
    const existingRoom = await db
      .select()
      .from(chatRooms)
      .where(eq(chatRooms.collection_address, collectionAddress.toLowerCase()))
      .limit(1);

    if (existingRoom.length > 0) {
      return existingRoom[0];
    }

    // Create new chat room
    const roomData: any = {
      collection_address: collectionAddress.toLowerCase(),
      room_name: `${collection.name} Chat`,
      description: `Chat room for ${collection.name} collection`,
      is_active: true,
    };
    const newRoom = await db.insert(chatRooms).values(roomData).returning();

    return newRoom[0];
  }

  /**
   * Get chat room members
   */
  async getChatRoomMembers(roomId: number): Promise<any[]> {
    const members = await db
      .select({
        id: users.id,
        walletAddress: users.wallet_address,
      })
      .from(chatRoomMembers)
      .innerJoin(users, eq(chatRoomMembers.user_id, users.id))
      .where(
        and(
          eq(chatRoomMembers.room_id, roomId),
          eq(chatRoomMembers.is_active, true),
        ),
      );

    return members;
  }

  /**
   * Get chat room messages with pagination
   */
  async getChatMessages(
    roomId: number,
    page: number = 1,
    pageSize: number = 50,
    before?: string,
  ): Promise<{ messages: any[]; total: number }> {
    console.log(
      'Getting chat messages for room:',
      roomId,
      'page:',
      page,
      'pageSize:',
      pageSize,
    );
    const offset = (page - 1) * pageSize;

    let whereClause = eq(chatMessages.room_id, roomId);
    if (before) {
      whereClause = and(
        eq(chatMessages.room_id, roomId),
        sql`${chatMessages.created_at} < ${before}`,
      );
    }

    // Get messages
    const messages = await db
      .select({
        id: chatMessages.id,
        userId: chatMessages.user_id,
        walletAddress: users.wallet_address,
        messageType: chatMessages.message_type,
        content: chatMessages.content,
        replyToMessageId: chatMessages.reply_to_message_id,
        isEdited: chatMessages.is_edited,
        createdAt: chatMessages.created_at,
      })
      .from(chatMessages)
      .innerJoin(users, eq(chatMessages.user_id, users.id))
      .where(whereClause)
      .orderBy(desc(chatMessages.created_at))
      .limit(pageSize)
      .offset(offset);

    // Get total count
    const totalResult = await db
      .select({ count: sql<number>`count(*)` })
      .from(chatMessages)
      .where(eq(chatMessages.room_id, roomId));

    console.log(
      'Found messages:',
      messages.length,
      'Total count:',
      totalResult[0]?.count || 0,
    );

    return {
      messages: messages.reverse(), // Reverse to show oldest first
      total: totalResult[0]?.count || 0,
    };
  }

  /**
   * Add user to chat room
   */
  async addUserToRoom(userId: number, roomId: number): Promise<void> {
    // Check if user is already a member
    const existingMember = await db
      .select()
      .from(chatRoomMembers)
      .where(
        and(
          eq(chatRoomMembers.user_id, userId),
          eq(chatRoomMembers.room_id, roomId),
        ),
      )
      .limit(1);

    if (existingMember.length > 0) {
      // Update existing membership
      await db
        .update(chatRoomMembers)
        .set({
          is_active: true,
          joined_at: new Date(),
        })
        .where(eq(chatRoomMembers.id, existingMember[0].id));
    } else {
      // Add new member
      await db.insert(chatRoomMembers).values({
        room_id: roomId,
        user_id: userId,
        is_active: true,
      });
    }
  }

  /**
   * Remove user from chat room
   */
  async removeUserFromRoom(userId: number, roomId: number): Promise<void> {
    await db
      .update(chatRoomMembers)
      .set({ is_active: false })
      .where(
        and(
          eq(chatRoomMembers.user_id, userId),
          eq(chatRoomMembers.room_id, roomId),
        ),
      );
  }

  /**
   * Save chat message
   */
  async saveMessage(
    roomId: number,
    userId: number,
    messageType: string,
    content: string,
    replyToMessageId?: number,
  ): Promise<any> {
    const newMessage = await db
      .insert(chatMessages)
      .values({
        room_id: roomId,
        user_id: userId,
        message_type: messageType,
        content,
        reply_to_message_id: replyToMessageId,
      } as any)
      .returning();

    // Get message with user info
    const messageWithUser = await db
      .select({
        id: chatMessages.id,
        userId: chatMessages.user_id,
        walletAddress: users.wallet_address,
        messageType: chatMessages.message_type,
        content: chatMessages.content,
        replyToMessageId: chatMessages.reply_to_message_id,
        isEdited: chatMessages.is_edited,
        createdAt: chatMessages.created_at,
      })
      .from(chatMessages)
      .innerJoin(users, eq(chatMessages.user_id, users.id))
      .where(eq(chatMessages.id, newMessage[0].id))
      .limit(1);

    return messageWithUser[0];
  }

  /**
   * Get online users for room
   */
  async getOnlineUsers(roomId: number): Promise<any[]> {
    const onlineUsersList = await db
      .select({
        id: users.id,
        walletAddress: users.wallet_address,
        connectedAt: onlineUsers.connected_at,
      })
      .from(onlineUsers)
      .innerJoin(users, eq(onlineUsers.user_id, users.id))
      .where(eq(onlineUsers.room_id, roomId));

    return onlineUsersList;
  }

  /**
   * Add user to online users
   */
  async addOnlineUser(
    userId: number,
    roomId: number,
    socketId: string,
  ): Promise<void> {
    // Remove existing connections for this user in this room
    await db
      .delete(onlineUsers)
      .where(
        and(eq(onlineUsers.user_id, userId), eq(onlineUsers.room_id, roomId)),
      );

    // Add new connection
    const onlineUserData: any = {
      user_id: userId,
      room_id: roomId,
      socket_id: socketId,
    };
    await db.insert(onlineUsers).values(onlineUserData);
  }

  /**
   * Remove user from online users
   */
  async removeOnlineUser(
    userId: number,
    roomId: number,
    socketId: string,
  ): Promise<void> {
    try {
      // Use a simpler approach - just try to delete without checking first
      // This avoids potential issues with connection state
      console.log('Removing online user:', userId, roomId, socketId);
      // await db
      //   .delete(onlineUsers)
      //   .where(
      //     and(
      //       eq(onlineUsers.user_id, userId),
      //       eq(onlineUsers.room_id, roomId),
      //       eq(onlineUsers.socket_id, socketId),
      //     ),
      //   );
    } catch (error) {
      console.error('Error removing online user:', error);
      // Don't throw the error, just log it
    }
  }

  /**
   * Get user by ID
   */
  async getUserById(userId: number): Promise<any> {
    try {
      const user = await db
        .select()
        .from(users)
        .where(eq(users.id, userId))
        .limit(1);

      return user[0];
    } catch (error) {
      console.error('Error getting user by ID:', error);
      return null;
    }
  }
}
