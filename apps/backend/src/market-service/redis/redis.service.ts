import { Injectable } from '@nestjs/common';
import Redis from 'ioredis';

@Injectable()
export class RedisService {
    public publisher: Redis
    public subscriber: Redis

    constructor(){
        this.publisher = new Redis({
            port: 6379, 
            host: 'localhost'
        })
        this.subscriber = new Redis({
            port: 6379,
            host: 'localhost'
        })
    }

    async publish(channel: string, payload: string){
        return this.publisher.publish(channel, payload)
    }
}
