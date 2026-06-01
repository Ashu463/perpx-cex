import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { AuthServiceModule } from './auth-service/auth-service.module';
import { WalletServiceModule } from './wallet-service/wallet-service.module';
// import { OrderServiceModule } from './market-service/order-service/order-service.module';
import { OrderServiceModule } from './order-service/order-service.module';
import { PositionServiceModule } from './position-service/position-service.module';
import { TradeServiceModule } from './trade-service/trade-service.module';
import { PrismaModule } from './prisma/prisma.module';

import { MarketServiceModule } from './market-service/market-service.module';
import { CacheService } from './market-service/cache/price-cache.service';
import { RedisModule } from './common/redis/redis.module';

@Module({
  imports: [PrismaModule, AuthServiceModule, WalletServiceModule, OrderServiceModule, PositionServiceModule, TradeServiceModule, MarketServiceModule, RedisModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
