import { Module } from '@nestjs/common';
import { MarketServiceService } from './services/market-service.service';
import { MarketServiceController } from './market-service.controller';
import { MarketDataService } from './services/market-data.services';
import { RedisService } from '../common/redis/redis.service';
import { CacheService } from './cache/price-cache.service';
import { CacheController } from './cache/price-cache.controller';
import { MarketSubscriber } from './subscribers/market-subscribers';

@Module({
  imports: [],
  controllers: [MarketServiceController, CacheController],
  providers: [MarketServiceService, MarketDataService, RedisService, CacheService, MarketSubscriber],
  exports: [CacheService]
})
export class MarketServiceModule {}
