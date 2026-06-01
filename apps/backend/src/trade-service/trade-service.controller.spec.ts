import { Test, TestingModule } from '@nestjs/testing';
import { TradeServiceController } from './trade-service.controller';
import { TradeServiceService } from './trade-service.service';

describe('TradeServiceController', () => {
  let controller: TradeServiceController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [TradeServiceController],
      providers: [TradeServiceService],
    }).compile();

    controller = module.get<TradeServiceController>(TradeServiceController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
