import { Controller } from '@nestjs/common';
import { TradeServiceService } from './trade-service.service';

@Controller('trade-service')
export class TradeServiceController {
  constructor(private readonly tradeServiceService: TradeServiceService) {}
}
