import { BadRequestException, ConflictException, ForbiddenException, Injectable, NotFoundException } from '@nestjs/common';
import { Decimal } from '@prisma/client/runtime/client';
import { PrismaService } from 'src/prisma/prisma.service';
import { OrderType, OrderSide, PositionSide } from '@prisma/client';


enum OrderIntent{
    OPEN_LONG,
    CLOSE_LONG,
    OPEN_SHORT,
    CLOSE_SHORT
}
@Injectable()
export class OrderServiceService {
    constructor(
        private readonly prismaService: PrismaService
    ){}
    
    async validateClosePosition(
        userId: string,
        marketId: string,
        quantity: Decimal,
        positionSide: PositionSide,
    ) {
        const position = await this.prismaService.position.findFirst({
            where: {
                userId,
                marketId,
                side: positionSide,
            },
        });

        if (!position) {
            throw new ForbiddenException(
                'Position does not exist',
            );
        }

        if (position.quantity.lessThan(quantity)) {
            throw new ForbiddenException(
                'Insufficient quantity',
            );
        }
    }
    async validateMargin(
        userId: string,
        margin: Decimal,
    ) {
        const wallet = await this.prismaService.wallet.findUnique({
            where: {
                userId,
            },
        });

        if (!wallet) {
            throw new ForbiddenException(
                'Wallet not found',
            );
        }

        if (wallet.availableBalance.lessThan(margin)) {
            throw new ForbiddenException(
                'Insufficient balance',
            );
        }
    }
    async createRespectiveOrders(){

    }
    async createOrder(body: {
    marketId: string,
    userId: string,
    side: string, 
    positionSide: PositionSide,
    type: string, 
    price: Decimal,
    quantity: Decimal,
    leverage: number
  }){
    const {userId, marketId, side, positionSide, type, price, quantity, leverage} = body

    const market = await this.prismaService.market.findUnique({where: {id: marketId}})

    if(!market){
        throw new NotFoundException('Market not found with id', marketId)
    }

    // margin calculation 
    const orderValue = price.mul(quantity)
    const margin: Decimal = orderValue.dividedBy(leverage)

    let intent = "";
    if(side === 'BUY'){
        if(positionSide === 'LONG') intent = "OPEN_LONG";
        else intent = "OPEN_SHORT"
    }
    else{
        if(positionSide === 'LONG') intent = "CLOSE_LONG"
        else intent = "CLOSE_SHORT"
    }
    let data: any = {
        userId: userId,
        marketId: marketId,
        type: type as OrderType, 
        orderPrice: price,
        quantity: quantity,
        leverage: leverage,
        margin: margin,
        status: 'PENDING',
        remainingQuantity: quantity,
        orderCost: Decimal(margin),
        orderValue: orderValue
    }
    switch(intent){
        case "OPEN_LONG":
            this.validateMargin(userId, margin);
            data = {
                side: OrderSide.BUY,
                positionSide: PositionSide.LONG,
                ...data
            }

        case "OPEN_SHORT":
            this.validateMargin(userId, margin);
            data = {
                side: OrderSide.BUY,
                positionSide: PositionSide.SHORT,
                ...data
            }

        case "CLOSE_LONG":
            this.validateClosePosition(userId, marketId, quantity, positionSide);
            data = {
                side: OrderSide.SELL,
                positionSide: PositionSide.LONG,
                ...data
            }

        case "CLOSE_SHORT":
            this.validateClosePosition(userId, marketId, quantity, positionSide);
            data = {
                side: OrderSide.SELL,
                positionSide: PositionSide.SHORT,
                ...data
            }

    }
    const newMarketOrder = await this.prismaService.order.create({
        data: data
    })


    if(!newMarketOrder){
        throw  new ConflictException('Error in creating market order to DB')
    }
    
    
    // publishToRedis(); #TODO


  }

  async getAll(body: {
    userId: number,
    orderId: number,
    market: string,
    side: string,
    type: string, 
    price: Decimal, 
    filledQuantity: Decimal,
    status: string
  }){
    const {userId, orderId, market, side, type, price, filledQuantity, status} = body

    // validate whether the user exist or not ya fer push it to the sessions and maintain it. 
    // if yes then return all orders associated with this user
    
  }
  
  async getOpenOrders(body:{
    userId: number,
    orderId: number,
    market: string,
    side: string,
    type: string, 
    price: Decimal, 
    filledQuantity: Decimal
  }){
    // validate whether the user exist or not
    // if so return all orders which are pending
  }

  async updateOrder(body:{
    userId: string, 
    orderId: string, 
    price: Decimal,
    quantity: Decimal,
    leverage: number
  }){
    
    // cancel existing 
    // create new
    // publish event

  }

  async cancelOrder(body: {
    userId: string, 
    orderId: string,
  }){
    // validate ownership 
    // validate it's open or not
    // publish ORDER_CANCELLED
    // release margin 
    // delete it from orders table and orderbook
  }
}
