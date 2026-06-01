import { Module } from '@nestjs/common';
import { WalletService } from './wallet-service.service';
import { WalletController } from './wallet-service.controller';
import { PrismaModule } from '../prisma/prisma.module';

@Module({
  imports: [PrismaModule],
  controllers: [WalletController],
  providers: [WalletService],
})
export class WalletServiceModule {}
