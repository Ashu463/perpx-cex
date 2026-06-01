import { Module } from '@nestjs/common';
import { AuthController } from './auth-service.controller';
import { AuthService } from './auth-service.service';
import { PrismaModule } from 'src/prisma/prisma.module';
import { JwtModule } from '@nestjs/jwt';

@Module({
  imports: [PrismaModule, JwtModule.register({
      secret: process.env.JWT_SECRET,
      signOptions: { expiresIn: '7d' },
    }),],
  controllers: [AuthController],
  providers: [AuthService],
})
export class AuthServiceModule {}
