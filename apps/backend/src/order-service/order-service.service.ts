import { BadRequestException, ConflictException, ForbiddenException, Injectable, NotFoundException, NotImplementedException } from '@nestjs/common';
import { Decimal } from '@prisma/client/runtime/client';
import { PrismaService } from 'src/prisma/prisma.service';
import { OrderType, OrderSide, PositionSide, OrderStatus, Prisma } from '@prisma/client';
import { RedisService } from 'src/common/redis/redis.service';
import { STREAMS } from 'src/common/redis/redis.contants';
import { OrderEvents } from './order-service.dto';


enum OrderIntent{
    OPEN_LONG,
    CLOSE_LONG,
    OPEN_SHORT,
    CLOSE_SHORT
}
@Injectable()
export class OrderServiceService {
    constructor(
        private readonly prismaService: PrismaService,
        private readonly redisService: RedisService
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
            break;

        case "OPEN_SHORT":
            this.validateMargin(userId, margin);
            data = {
                side: OrderSide.BUY,
                positionSide: PositionSide.SHORT,
                ...data
            }
            break;

        case "CLOSE_LONG":
            this.validateClosePosition(userId, marketId, quantity, positionSide);
            data = {
                side: OrderSide.SELL,
                positionSide: PositionSide.LONG,
                ...data
            }
            break;

        case "CLOSE_SHORT":
            this.validateClosePosition(userId, marketId, quantity, positionSide);
            data = {
                side: OrderSide.SELL,
                positionSide: PositionSide.SHORT,
                ...data
            }
            break;

    }
    const newMarketOrder = await this.prismaService.order.create({
        data: data
    })

    if(!newMarketOrder){
        throw  new ConflictException('Error in creating market order to DB')
    }
    
    
    // publishToRedis(); #TODO as streams
    const publishToRedis = await this.redisService.addToStream(
        STREAMS.ORDER_SUBMISSIONS,
        {
            event: 'ORDER_CREATED',
            orderId: newMarketOrder.id,
            userId: newMarketOrder.userId,
            marketId: newMarketOrder.marketId,
            side: newMarketOrder.side,
            positionSide: newMarketOrder.positionSide,
            type: newMarketOrder.type,
            price: newMarketOrder.orderPrice,
            quantity: newMarketOrder.quantity,
            leverage: newMarketOrder.leverage,
            margin: newMarketOrder.margin
        }
    )
    if(!publishToRedis){
        throw new NotImplementedException('Redis stream failed to publish')
    }

    return {
        message: 'Order created successfully',
        orderId: newMarketOrder.id,
    }

  }

  async getAllOrders(userId: string) {

        const orders = await this.prismaService.order.findMany({
            where: {
                userId,
            },
            orderBy: {
                createdAt: 'desc',
            },
        });

        return orders;
    }
  
  async getOpenOrders(userId: string) {
    return this.prismaService.order.findMany({
            where: {
                userId,
                OR: [
                    {
                        status: OrderStatus.PENDING,
                    },
                    {
                        status:
                            OrderStatus.PARTIALLY_FIILED,
                    },
                ],
            },
            orderBy: {
                createdAt: 'desc',
            },
        });
    }

  async updateOrder(body:{
    orderId: string, 
    price: Decimal,
    quantity: Decimal,
  }){

    const {orderId, price, quantity} = body;
    if(!orderId){
        throw new BadRequestException('orderId missing')
    }

    const order = await this.prismaService.order.findUnique({where: {id: orderId}})

    if(!order){
        throw new BadRequestException('Order with this id does not exist')
    }
    if(order.status !== OrderStatus.PENDING && order.status !== OrderStatus.PARTIALLY_FIILED){
        throw new BadRequestException('Order is no longer pending, might be allocated')
    }

    if(order.remainingQuantity.greaterThan(quantity)){
        throw new BadRequestException('Placing a new order would be better rather than updating old one, just cancel to delete the rest of units')
    }
    
    const oldMargin = order.margin;
    const newOrderValue: Decimal = price.mul(order.remainingQuantity)
    const newMargin: Decimal = newOrderValue.dividedBy(order.leverage)

    const diff = newMargin.sub(oldMargin)
    const wallet = await this.prismaService.wallet.findUnique({where: {id: order.userId}})
    const availableBalance = wallet!.availableBalance

    if(diff.lessThan(0) || availableBalance.lessThan(diff)){
    }
    
    if(diff.greaterThan(0)){
        if(availableBalance.lessThan(diff))
        throw new NotFoundException('Unavailable resource, either available balance is low or diff between margins is negative')
        
        
    }

    await this.prismaService.wallet.update({
        where: {id: order.userId},
        data:{
            availableBalance: availableBalance.sub(diff),
            lockedBalance: wallet!.lockedBalance.add(diff)
        }
    })
    // mark the sold trade as filled and then open new order with remaining trade at asked price. 
    await this.prismaService.order.update({
        where: {id: order.id},
        data:{
            status: OrderStatus.CANCELLED
        }
    })
    const newMarketOrder = await this.prismaService.order.create({
        data:{
            userId: order.userId,
            marketId: order.marketId,
            type: order.type,
            side: order.side,
            positionSide: order.positionSide,
            orderPrice: price,
            quantity: order.remainingQuantity,
            leverage: order.leverage,
            margin: newMargin,
            status: OrderStatus.PENDING,
            remainingQuantity: order.remainingQuantity,
            orderCost: new Prisma.Decimal(newMargin),
            orderValue: newOrderValue
        }
    })
    if(!newMarketOrder){
        throw new NotImplementedException('failed to create new order')
    }
    const publishToRedis = await this.redisService.addToStream(
        STREAMS.ORDER_SUBMISSIONS,
        {
            event: 'ORDER_CREATED',
            orderId: newMarketOrder.id,
            userId: newMarketOrder.userId,
            marketId: newMarketOrder.marketId,
            side: newMarketOrder.side,
            positionSide: newMarketOrder.positionSide,
            type: newMarketOrder.type,
            price: newMarketOrder.orderPrice,
            quantity: newMarketOrder.quantity,
            leverage: newMarketOrder.leverage,
            margin: newMarketOrder.margin
        }
    )
    if(!publishToRedis){
        throw new NotImplementedException('Failed to publish order')
    }
    return {
        orderId: newMarketOrder.id
    }
    // cancel existing 
    // create new
    // publish event

  }

  async cancelOrder(body: {
        orderId: string;
        userId: string;
    }) {
        const { orderId, userId } = body;

        const order = await this.prismaService.order.findUnique({
            where: {
                id: orderId,
            },
        });

        if (!order) {
            throw new NotFoundException(
                'Order not found',
            );
        }

        if (order.userId !== userId) {
            throw new ForbiddenException(
                'Order does not belong to user',
            );
        }

        if (
            order.status !== OrderStatus.PENDING &&
            order.status !== OrderStatus.PARTIALLY_FIILED
        ) {
            throw new BadRequestException(
                'Order cannot be cancelled',
            );
        }

        const wallet =
            await this.prismaService.wallet.findUnique({
                where: {
                    id: userId,
                },
            });

        if (!wallet) {
            throw new NotFoundException(
                'Wallet not found',
            );
        }

        const refund = order.margin.mul(
            order.remainingQuantity.div(
                order.quantity,
            ),
        );

        await this.prismaService.wallet.update({
            where: {
                id: userId,
            },
            data: {
                availableBalance:
                    wallet.availableBalance.add(
                        refund,
                    ),

                lockedBalance:
                    wallet.lockedBalance.sub(
                        refund,
                    ),
            },
        });

        await this.prismaService.order.update({
            where: {
                id: orderId,
            },
            data: {
                status:
                    OrderStatus.CANCELLED,

                remainingQuantity:
                    new Prisma.Decimal(0),
            },
        });

        await this.redisService.addToStream(
            STREAMS.ORDER_SUBMISSIONS,
            {
                event:
                    OrderEvents.ORDER_CANCEL_REQUESTED,

                orderId: order.id,

                marketId:
                    order.marketId,

                side: order.side,

                price:
                    order.orderPrice.toString(),

                remainingQuantity:
                    order.remainingQuantity.toString(),
            },
        );

        /**
         * TODO:
         * Ideal architecture:
         *
         * 1. Publish ORDER_CANCEL_REQUESTED
         * 2. Matching Engine removes order
         * 3. Matching Engine emits ORDER_CANCELLED
         * 4. Backend consumes ORDER_CANCELLED
         * 5. Unlock margin
         * 6. Update DB status
         */

        return {
            message:
                'Order cancelled successfully',
            refund,
        };
    }

}
