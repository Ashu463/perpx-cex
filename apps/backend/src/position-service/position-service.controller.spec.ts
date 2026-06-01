import { Test, TestingModule } from '@nestjs/testing';
import { PositionServiceController } from './position-service.controller';
import { PositionServiceService } from './position-service.service';

describe('PositionServiceController', () => {
  let controller: PositionServiceController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [PositionServiceController],
      providers: [PositionServiceService],
    }).compile();

    controller = module.get<PositionServiceController>(PositionServiceController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
