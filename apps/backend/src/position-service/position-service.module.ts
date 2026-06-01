import { Module } from '@nestjs/common';
import { PositionServiceService } from './position-service.service';
import { PositionServiceController } from './position-service.controller';

@Module({
  controllers: [PositionServiceController],
  providers: [PositionServiceService],
})
export class PositionServiceModule {}
