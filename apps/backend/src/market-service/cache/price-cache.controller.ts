import { Controller, Get, Injectable } from "@nestjs/common";
import { CacheService } from "./price-cache.service";

@Controller('cache')
export class CacheController{
    constructor(private cacheService: CacheService){

    }
    
    @Get('/prices')
    getPrices() {
        return this.cacheService.getAll();
    }
}