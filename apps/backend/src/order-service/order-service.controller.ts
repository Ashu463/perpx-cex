import { Body, Controller, Get, Patch, Post, Req } from '@nestjs/common';
import { OrderServiceService } from './order-service.service';
import { Decimal } from '@prisma/client/runtime/client';
import { PositionSide } from '@prisma/client';

@Controller('order-service')
export class OrderServiceController {
  constructor(
    private readonly orderService: OrderServiceService,
  ) {}

  @Post('/create')
  async createOrder(body: {
      marketId: string,
      userId: string,
      side: string, 
      positionSide: PositionSide,
      type: string, 
      price: Decimal,
      quantity: Decimal,
      leverage: number}){
    return this.orderService.createOrder(body)
  }

  @Patch('/update')
  async updateOrder(
    @Body()
    body: {
      orderId: string;
      price: number;
      quantity: number;
    },
  ) {
    return this.orderService.updateOrder({
      orderId: body.orderId,
      price: new Decimal(body.price),
      quantity: new Decimal(body.quantity),
    });
  }

  @Get('/all')
  async getAllOrders(
     userId: string
  ) {
    return this.orderService.getAllOrders(
        userId,
    );
  }

  @Get('/open')
  async getOpenOrders(
      userId: string
  ) {
    return this.orderService.getOpenOrders(
        userId
    );
  }

  @Post('/cancel')
  async cancelOrder(
    @Req() req,
    @Body()
    body: {
      orderId: string;
    },
  ) {
    return this.orderService.cancelOrder({
      orderId: body.orderId,
      userId: req.user.id,
    });
  }
}
