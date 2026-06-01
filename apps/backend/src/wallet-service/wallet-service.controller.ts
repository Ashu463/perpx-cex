import {
  Body,
  Controller,
  Get,
  Param,
  Post,
} from '@nestjs/common';
import { WalletService } from './wallet-service.service';

@Controller('wallet')
export class WalletController {
  constructor(
    private readonly walletService: WalletService,
  ) {}

  @Post('create')
  create(
    @Body()
    body: {
      userId: string;
    },
  ) {
    return this.walletService.create(body);
  }
  @Get('balance/:userId')
  getBalance(
    @Param('userId') userId: string,
  ) {
    return this.walletService.getBalance(
      userId,
    );
  }
  @Post('deposit')
  deposit(
    @Body()
    body: {
      userId: string;
      amount: number;
    },
  ) {
    return this.walletService.deposit(
      body.userId,
      body.amount,
    );
  }
  @Post('withdraw')
  withdraw(
    @Body()
    body: {
      userId: string;
      amount: number;
    },
  ) {
    return this.walletService.withdraw(
      body.userId,
      body.amount,
    );
  }
  @Get('transactions/:userId')
  transactions(
    @Param('userId') userId: string,
  ) {
    return this.walletService.transactions(
      userId,
    );
  }
}