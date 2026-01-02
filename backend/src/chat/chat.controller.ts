import {
  Controller,
  Get,
  Param,
  Query,
  HttpStatus,
  HttpException,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiParam } from '@nestjs/swagger';
import { ChatService } from './chat.service';
import { AuthService } from '../auth/auth.service';
import { GetChatMessagesDto } from '../dto/chat.dto';

// Simple JWT guard for chat endpoints
@Controller('api/v1/chat')
@ApiTags('Chat API')
export class ChatController {
  constructor(
    private readonly chatService: ChatService,
    private readonly authService: AuthService,
  ) {}

  /**
   * Get chat room information
   */
  @Get('rooms/:collectionAddress')
  @ApiOperation({ summary: 'Get chat room information for collection' })
  @ApiParam({
    name: 'collectionAddress',
    description: 'Collection contract address',
  })
  @ApiResponse({
    status: 200,
    description: 'Chat room information retrieved successfully',
    schema: {
      type: 'object',
      properties: {
        success: { type: 'boolean' },
        room: { $ref: '#/components/schemas/ChatRoomDto' },
        members: {
          type: 'array',
          items: {
            type: 'object',
            properties: {
              id: { type: 'number' },
              walletAddress: { type: 'string' },
              isOnline: { type: 'boolean' },
            },
          },
        },
        message: { type: 'string' },
      },
    },
  })
  async getChatRoom(
    @Param('collectionAddress') collectionAddress: string,
    @Query('chainId') chainId: number,
  ) {
    try {
      if (!chainId) {
        throw new HttpException('Chain ID is required', HttpStatus.BAD_REQUEST);
      }

      // Get or create chat room
      const room = await this.chatService.getOrCreateChatRoom(
        collectionAddress,
        chainId,
      );

      // Get room members
      const members = await this.chatService.getChatRoomMembers(room.id);

      // Get online users
      const onlineUsers = await this.chatService.getOnlineUsers(room.id);
      const onlineUserIds = new Set(onlineUsers.map((u) => u.id));

      // Add online status to members
      const membersWithStatus = members.map((member) => ({
        ...member,
        isOnline: onlineUserIds.has(member.id),
      }));

      return {
        success: true,
        room: {
          id: room.id,
          collectionId: room.collection_id,
          roomName: room.room_name,
          description: room.description,
          memberCount: members.length,
          isActive: room.is_active,
        },
        members: membersWithStatus,
        message: 'Chat room information retrieved successfully',
      };
    } catch (error) {
      throw new HttpException(
        `Failed to get chat room: ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  /**
   * Get chat room messages
   */
  @Get('rooms/:collectionAddress/messages')
  @ApiOperation({ summary: 'Get chat room messages' })
  @ApiParam({
    name: 'collectionAddress',
    description: 'Collection contract address',
  })
  @ApiResponse({
    status: 200,
    description: 'Chat messages retrieved successfully',
    schema: {
      type: 'object',
      properties: {
        success: { type: 'boolean' },
        messages: {
          type: 'array',
          items: { $ref: '#/components/schemas/ChatMessageDto' },
        },
        pagination: {
          type: 'object',
          properties: {
            page: { type: 'number' },
            pageSize: { type: 'number' },
            total: { type: 'number' },
            totalPages: { type: 'number' },
          },
        },
        message: { type: 'string' },
      },
    },
  })
  async getChatMessages(
    @Param('collectionAddress') collectionAddress: string,
    @Query() query: GetChatMessagesDto,
  ) {
    try {
      if (!collectionAddress) {
        throw new HttpException(
          'Collection address is required',
          HttpStatus.BAD_REQUEST,
        );
      }

      // Get chat room by collection address only
      const room =
        await this.chatService.getChatRoomByAddress(collectionAddress);

      // Get messages
      const result = await this.chatService.getChatMessages(
        room.id,
        query.page || 1,
        query.pageSize || 50,
        query.before,
      );

      const totalPages = Math.ceil(result.total / (query.pageSize || 50));

      return {
        success: true,
        messages: result.messages,
        pagination: {
          page: query.page || 1,
          pageSize: query.pageSize || 50,
          total: result.total,
          totalPages,
        },
        message: 'Chat messages retrieved successfully',
      };
    } catch (error) {
      throw new HttpException(
        `Failed to get chat messages: ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }

  /**
   * Get online users for chat room
   */
  @Get('rooms/:collectionAddress/online-users')
  @ApiOperation({ summary: 'Get online users for chat room' })
  @ApiParam({
    name: 'collectionAddress',
    description: 'Collection contract address',
  })
  @ApiResponse({
    status: 200,
    description: 'Online users retrieved successfully',
    schema: {
      type: 'object',
      properties: {
        success: { type: 'boolean' },
        onlineUsers: {
          type: 'array',
          items: { $ref: '#/components/schemas/OnlineUserDto' },
        },
        message: { type: 'string' },
      },
    },
  })
  async getOnlineUsers(
    @Param('collectionAddress') collectionAddress: string,
    @Query('chainId') chainId: number,
  ) {
    try {
      if (!chainId) {
        throw new HttpException('Chain ID is required', HttpStatus.BAD_REQUEST);
      }

      // Get chat room
      const room = await this.chatService.getOrCreateChatRoom(
        collectionAddress,
        chainId,
      );

      // Get online users
      const onlineUsers = await this.chatService.getOnlineUsers(room.id);

      return {
        success: true,
        onlineUsers,
        message: 'Online users retrieved successfully',
      };
    } catch (error) {
      throw new HttpException(
        `Failed to get online users: ${error.message}`,
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    }
  }
}
