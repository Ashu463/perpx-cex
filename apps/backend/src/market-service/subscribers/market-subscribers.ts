import { Injectable, OnModuleInit } from "@nestjs/common";
import Redis from "ioredis";
import { CacheService } from "../cache/price-cache.service";
import { RedisService } from "../../common/redis/redis.service";

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
      console.log('MESSAGE RECEIVED', channel, message);

      const payload = JSON.parse(message);
      const actualPayload = JSON.parse(payload)
      // double parsing is working IDK why #TODO - enhancement

      this.priceCache.setPrice(
        actualPayload.symbol,
        actualPayload.price,
      );

      console.log(
        'CACHE UPDATED',
        actualPayload.symbol,
        actualPayload.price,
      );
    },
  );
}
}