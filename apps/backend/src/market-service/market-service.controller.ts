import { Controller, Get } from '@nestjs/common';
import { MarketServiceService } from './services/market-service.service';
import { CacheService } from './cache/price-cache.service';

@Controller('market-service')
export class MarketServiceController {
  constructor(private readonly marketServiceService: MarketServiceService, private cacheService: CacheService) {}

}