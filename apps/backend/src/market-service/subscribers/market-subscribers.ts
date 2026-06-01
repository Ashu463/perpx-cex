import { Injectable, OnModuleInit } from "@nestjs/common";
import Redis from "ioredis";
import { CacheService } from "../cache/price-cache.service";
import { RedisService } from "../redis/redis.service";

@Injectable()
export class MarketSubscriber implements OnModuleInit {
    constructor (
        private readonly redisService: RedisService,
        private readonly priceCache: CacheService
    ){}

    async onModuleInit() {
  console.log('MARKET SUBSCRIBER STARTED');

  const subscriber = this.redisService.subscriber;

  await subscriber.psubscribe('ticker:*');

  console.log('Subscribed to ticker:*');

  subscriber.on(
    'pmessage',
    (_pattern, channel, message) => {
      console.log('MESSAGE RECEIVED', channel);

      const payload = JSON.parse(message);

      this.priceCache.setPrice(
        payload.symbol,
        payload.price,
      );

      console.log(
        'CACHE UPDATED',
        payload.symbol,
        payload.price,
      );
    },
  );
}
}