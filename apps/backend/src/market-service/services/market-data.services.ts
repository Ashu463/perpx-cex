import { Injectable, OnModuleInit } from "@nestjs/common";
import Redis from "ioredis";
import { WebSocket } from 'ws'
import { RedisService } from "../redis/redis.service";

@Injectable()
export class MarketDataService implements OnModuleInit {
    constructor(
  private readonly redisService: RedisService,
) {
  console.log('redisService', redisService);
}
  private ws: WebSocket
  async onModuleInit() {
    this.connectToBinance();
  }

  private connectToBinance() {
    // Binance WS
    const url = 'wss://stream.binance.com:9443/stream?streams=btcusdt@miniTicker/ethusdt@miniTicker/solusdt@miniTicker';
    this.ws = new WebSocket(url)

    this.ws.on('open', () => {
        console.log('Connected to Binance')
    })
    this.ws.on('message', (data: any) => {
        this.handleMessage(data.toString())
        // console.log('message recieved ', data)
    })
    this.ws.on('error', (err: any) => {
        console.log('error occurred ', err)
    })
    this.ws.on('close', () =>{
        console.log('connection closed')
        setTimeout(() => {
            this.connectToBinance();
        }, 5000);
    })
  }
  redis = new Redis();

  private async handlePriceUpdate(data: any) {
    await this.redis.publish(
      'market:prices',
      JSON.stringify(data),
    );
  }
  private async handleMessage(message: string) {
    const parsed = JSON.parse(message);

    const ticker = parsed.data;

    const payload = {
        symbol: ticker.s,
        price: Number(ticker.c),
        timestamp: Date.now(),
    };

    await this.redisService.publish(
        `ticker:${ticker.s}`,
        JSON.stringify(payload),
    );

    console.log('published', payload);
    }
}