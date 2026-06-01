import { Injectable } from '@nestjs/common';
import Redis from 'ioredis';

@Injectable()
export class RedisService {
  public readonly publisher: Redis;
  public readonly subscriber: Redis;

  constructor() {
    this.publisher = new Redis({
      host: 'localhost',
      port: 6379,
    });

    this.subscriber = new Redis({
      host: 'localhost',
      port: 6379,
    });
  }

  async publish(
    channel: string,
    payload: unknown,
  ) {
    console.log('REDIS PUBLISH INPUT', payload);
    await this.publisher.publish(
      channel,
      JSON.stringify(payload),
    );
  }

  async subscribe(
    pattern: string,
  ) {
    return this.subscriber.psubscribe(pattern);
  }
}