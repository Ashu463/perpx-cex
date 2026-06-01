import { Module } from '@nestjs/common';
import { TradeServiceService } from './trade-service.service';
import { TradeServiceController } from './trade-service.controller';

@Module({
  controllers: [TradeServiceController],
  providers: [TradeServiceService],
})
export class TradeServiceModule {}
