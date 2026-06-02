import { Injectable } from "@nestjs/common";

@Injectable()
export class CacheService{
    constructor(){}
    private readonly prices = new Map<string, number>()

    setPrice(symbol: string, price: number){
        // console.log('Setting price for ', symbol, price)
        this.prices.set(symbol, price)
    }
    getPrice(symbol: string){
        return this.prices.get(symbol)
    }
    getAll(){
        return Object.fromEntries(this.prices)
    }
}