import { Body, Controller } from '@nestjs/common';
import { OrderServiceService } from './order-service.service';
import { Decimal } from '@prisma/client/runtime/client';

@Controller('order-service')
export class OrderServiceController {
  constructor(
    private readonly orderService: OrderServiceService,
  ) {}

  


}
