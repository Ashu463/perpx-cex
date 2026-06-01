import { Controller } from '@nestjs/common';
import { PositionServiceService } from './position-service.service';

@Controller('position-service')
export class PositionServiceController {
  constructor(private readonly positionServiceService: PositionServiceService) {}
}
