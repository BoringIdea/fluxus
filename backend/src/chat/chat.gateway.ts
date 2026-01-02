import {
  WebSocketGateway,
  WebSocketServer,
  SubscribeMessage,
  MessageBody,
  ConnectedSocket,
  OnGatewayConnection,
  OnGatewayDisconnect,
} from '@nestjs/websockets';
import { Server, Socket } from 'socket.io';
import { Injectable, Logger } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { ChatService } from './chat.service';

@Injectable()
@WebSocketGateway({
  cors: {
    origin: [
      'http://localhost:3000',
      'https://www.testnet.fluxusnft.xyz',
      'https://www.fluxusnft.xyz',
      process.env.FRONTEND_URL || 'http://localhost:3000',
    ].filter(Boolean),
    credentials: true,
    methods: ['GET', 'POST'],
  },
})
export class ChatGateway implements OnGatewayConnection, OnGatewayDisconnect {
  @WebSocketServer()
  server: Server;

  private readonly logger = new Logger(ChatGateway.name);
  private readonly connectedUsers = new Map<
    string,
    { userId: number; roomId: string }
  >();

  constructor(
    private readonly chatService: ChatService,
    private readonly jwtService: JwtService,
  ) {}

  /**
   * Handle client connection
   */
  async handleConnection(client: Socket) {
    this.logger.log(`Client connected: ${client.id}`);
  }

  /**
   * Handle client disconnection
   */
  async handleDisconnect(client: Socket) {
    this.logger.log(`Client disconnected: ${client.id}`);

    const userInfo = this.connectedUsers.get(client.id);
    if (userInfo) {
      try {
        // Remove from online users
        await this.chatService.removeOnlineUser(
          userInfo.userId,
          parseInt(userInfo.roomId),
          client.id,
        );

        // Notify room about user leaving
        const user = await this.chatService.getUserById(userInfo.userId);
        client.to(userInfo.roomId).emit('user_left', {
          userId: userInfo.userId,
          walletAddress: user?.wallet_address,
        });

        this.connectedUsers.delete(client.id);
      } catch (error) {
        this.logger.error(`Error during disconnect cleanup: ${error.message}`);
        // Clean up the connection even if there's an error
        this.connectedUsers.delete(client.id);
      }
    }
  }

  /**
   * Join chat room
   */
  @SubscribeMessage('join_room')
  async handleJoinRoom(
    @MessageBody() data: { roomId: string; token: string; chainId: number },
    @ConnectedSocket() client: Socket,
  ) {
    try {
      // Verify JWT token
      const payload = this.jwtService.verify(data.token);

      // Get chat room
      const room = await this.chatService.getOrCreateChatRoom(
        data.roomId,
        data.chainId,
      );

      // Add user to room
      await this.chatService.addUserToRoom(payload.userId, room.id);

      // Add to online users
      await this.chatService.addOnlineUser(payload.userId, room.id, client.id);

      // Join socket room
      await client.join(data.roomId);

      // Store user info
      this.connectedUsers.set(client.id, {
        userId: payload.userId,
        roomId: room.id, // Store the actual room ID from database, not the collection address
      });

      // Notify client
      client.emit('room_joined', {
        success: true,
        roomId: room.id,
        message: 'Successfully joined room',
      });

      // Notify other users
      const user = await this.chatService.getUserById(payload.userId);
      client.to(data.roomId).emit('user_joined', {
        userId: payload.userId,
        walletAddress: user?.wallet_address,
      });

      this.logger.log(`User ${payload.userId} joined room ${data.roomId}`);
    } catch (error) {
      this.logger.error(`Failed to join room: ${error.message}`, error.stack);
      client.emit('room_joined', {
        success: false,
        message: `Failed to join room: ${error.message}`,
      });
    }
  }

  /**
   * Leave chat room
   */
  @SubscribeMessage('leave_room')
  async handleLeaveRoom(
    @MessageBody() data: { roomId: string },
    @ConnectedSocket() client: Socket,
  ) {
    try {
      const userInfo = this.connectedUsers.get(client.id);
      if (userInfo) {
        // Remove from online users
        await this.chatService.removeOnlineUser(
          userInfo.userId,
          parseInt(data.roomId),
          client.id,
        );

        // Notify other users
        const user = await this.chatService.getUserById(userInfo.userId);
        client.to(data.roomId).emit('user_left', {
          userId: userInfo.userId,
          walletAddress: user?.wallet_address,
        });

        this.connectedUsers.delete(client.id);
      }

      // Leave socket room
      await client.leave(data.roomId);

      client.emit('room_left', {
        success: true,
        roomId: data.roomId,
        message: 'Successfully left room',
      });

      this.logger.log(`User left room ${data.roomId}`);
    } catch (error) {
      this.logger.error(`Failed to leave room: ${error.message}`);
      client.emit('room_left', {
        success: false,
        message: 'Failed to leave room',
      });
    }
  }

  /**
   * Send message
   */
  @SubscribeMessage('send_message')
  async handleSendMessage(
    @MessageBody()
    data: {
      roomId: string;
      messageType: string;
      content: string;
      replyToMessageId?: number;
      chainId: number;
    },
    @ConnectedSocket() client: Socket,
  ) {
    try {
      const userInfo = this.connectedUsers.get(client.id);
      if (!userInfo) {
        client.emit('message_sent', {
          success: false,
          message: 'User not authenticated',
        });
        return;
      }

      // Get chat room
      const room = await this.chatService.getOrCreateChatRoom(
        data.roomId,
        data.chainId,
      );

      // Save message to database
      const savedMessage = await this.chatService.saveMessage(
        room.id,
        userInfo.userId,
        data.messageType,
        data.content,
        data.replyToMessageId,
      );

      // Notify client
      client.emit('message_sent', {
        success: true,
        messageId: savedMessage.id,
        message: 'Message sent successfully',
      });

      // Broadcast to room
      this.server.to(data.roomId).emit('receive_message', savedMessage);

      this.logger.log(
        `Message sent in room ${data.roomId} by user ${userInfo.userId}`,
      );
    } catch (error) {
      this.logger.error(`Failed to send message: ${error.message}`);
      client.emit('message_sent', {
        success: false,
        message: 'Failed to send message',
      });
    }
  }

  /**
   * Handle typing start
   */
  @SubscribeMessage('typing_start')
  async handleTypingStart(
    @MessageBody() data: { roomId: string },
    @ConnectedSocket() client: Socket,
  ) {
    const userInfo = this.connectedUsers.get(client.id);
    if (userInfo) {
      const user = await this.chatService.getUserById(userInfo.userId);
      client.to(data.roomId).emit('user_typing', {
        userId: userInfo.userId,
        walletAddress: user?.wallet_address,
      });
    }
  }

  /**
   * Handle typing stop
   */
  @SubscribeMessage('typing_stop')
  async handleTypingStop(
    @MessageBody() data: { roomId: string },
    @ConnectedSocket() client: Socket,
  ) {
    const userInfo = this.connectedUsers.get(client.id);
    if (userInfo) {
      const user = await this.chatService.getUserById(userInfo.userId);
      client.to(data.roomId).emit('user_stop_typing', {
        userId: userInfo.userId,
        walletAddress: user?.wallet_address,
      });
    }
  }
}
