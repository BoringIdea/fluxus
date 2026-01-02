import {
  IsString,
  IsNotEmpty,
  IsOptional,
  IsNumber,
  Min,
  Max,
} from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';
import { Type } from 'class-transformer';

export class GetChatRoomDto {
  @ApiProperty({ description: 'Collection address' })
  @IsString()
  @IsNotEmpty()
  collectionAddress: string;
}

export class GetChatMessagesDto {
  @ApiProperty({ description: 'Page number', required: false, default: 1 })
  @IsOptional()
  @Type(() => Number)
  @IsNumber()
  @Min(1)
  page?: number = 1;

  @ApiProperty({ description: 'Page size', required: false, default: 50 })
  @IsOptional()
  @Type(() => Number)
  @IsNumber()
  @Min(1)
  @Max(100)
  pageSize?: number = 50;

  @ApiProperty({
    description: 'Get messages before this timestamp',
    required: false,
  })
  @IsOptional()
  @IsString()
  before?: string;
}

export class SendMessageDto {
  @ApiProperty({ description: 'Room ID' })
  @IsNumber()
  roomId: number;

  @ApiProperty({ description: 'Message type', default: 'text' })
  @IsString()
  @IsNotEmpty()
  messageType: string;

  @ApiProperty({ description: 'Message content' })
  @IsString()
  @IsNotEmpty()
  content: string;

  @ApiProperty({ description: 'Reply to message ID', required: false })
  @IsOptional()
  @IsNumber()
  replyToMessageId?: number;
}

export class ChatRoomDto {
  @ApiProperty({ description: 'Room ID' })
  id: number;

  @ApiProperty({ description: 'Collection ID' })
  collectionId: number;

  @ApiProperty({ description: 'Room name' })
  roomName: string;

  @ApiProperty({ description: 'Room description', required: false })
  description?: string;

  @ApiProperty({ description: 'Member count' })
  memberCount: number;

  @ApiProperty({ description: 'Is room active' })
  isActive: boolean;
}

export class ChatMessageDto {
  @ApiProperty({ description: 'Message ID' })
  id: number;

  @ApiProperty({ description: 'User ID' })
  userId: number;

  @ApiProperty({ description: 'Wallet address' })
  walletAddress: string;

  @ApiProperty({ description: 'Message type' })
  messageType: string;

  @ApiProperty({ description: 'Message content' })
  content: string;

  @ApiProperty({ description: 'Reply to message ID', required: false })
  replyToMessageId?: number;

  @ApiProperty({ description: 'Is message edited' })
  isEdited: boolean;

  @ApiProperty({ description: 'Created at timestamp' })
  createdAt: string;
}

export class OnlineUserDto {
  @ApiProperty({ description: 'User ID' })
  id: number;

  @ApiProperty({ description: 'Wallet address' })
  walletAddress: string;

  @ApiProperty({ description: 'Connected at timestamp' })
  connectedAt: string;
}
