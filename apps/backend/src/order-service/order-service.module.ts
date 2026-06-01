import { Module } from '@nestjs/common';
import { OrderServiceService } from './order-service.service';
import { OrderServiceController } from './order-service.controller';
import { PrismaService } from 'src/prisma/prisma.service';

@Module({
  controllers: [OrderServiceController],
  providers: [OrderServiceService, PrismaService],
})
export class OrderServiceModule {}
