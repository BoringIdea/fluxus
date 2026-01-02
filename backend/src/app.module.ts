import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { ScheduleModule } from '@nestjs/schedule';
import { JwtModule } from '@nestjs/jwt';
import { AppService } from './app.service';
import { AppController } from './app.controller';
import { AuthController } from './auth/auth.controller';
import { ChatController } from './chat/chat.controller';
import { ChatGateway } from './chat/chat.gateway';
import { AuthService } from './auth/auth.service';
import { ChatService } from './chat/chat.service';

@Module({
  imports: [
    ScheduleModule.forRoot(),
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    JwtModule.register({
      secret: process.env.JWT_SECRET || 'your-secret-key',
      signOptions: { expiresIn: '24h' },
    }),
  ],
  controllers: [AppController, AuthController, ChatController],
  providers: [AppService, AuthService, ChatService, ChatGateway],
})
export class AppModule {}
