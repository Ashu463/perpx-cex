import { Test, TestingModule } from '@nestjs/testing';
import { TradeServiceService } from './trade-service.service';

describe('TradeServiceService', () => {
  let service: TradeServiceService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [TradeServiceService],
    }).compile();

    service = module.get<TradeServiceService>(TradeServiceService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
