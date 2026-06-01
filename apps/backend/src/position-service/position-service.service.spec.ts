import { Test, TestingModule } from '@nestjs/testing';
import { PositionServiceService } from './position-service.service';

describe('PositionServiceService', () => {
  let service: PositionServiceService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [PositionServiceService],
    }).compile();

    service = module.get<PositionServiceService>(PositionServiceService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
