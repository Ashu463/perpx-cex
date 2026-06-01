import {
    BadRequestException,
  ConflictException,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';

@Injectable()
export class WalletService {
  constructor(
    private readonly prisma: PrismaService,
  ) {}

  async create(body: { userId: string }) {
    const { userId } = body;

    const user = await this.prisma.user.findUnique({
      where: {
        id: userId,
      },
    });

    if (!user) {
      throw new NotFoundException('User not found');
    }

    const existingWallet =
      await this.prisma.wallet.findUnique({
        where: {
          userId,
        },
      });

    if (existingWallet) {
      throw new ConflictException(
        'Wallet already exists',
      );
    }

    const wallet = await this.prisma.wallet.create({
      data: {
        userId,
      },
    });

    return {
      walletId: wallet.id,
      availableBalance: wallet.availableBalance,
      lockedBalance: wallet.lockedBalance,
    };
  }

  async getBalance(userId: string) {
    const wallet = await this.prisma.wallet.findUnique({
        where: {
        userId,
        },
    });

    if (!wallet) {
        throw new NotFoundException('Wallet not found');
    }

    const unrealizedPnl = 0;

    return {
        availableBalance: wallet.availableBalance,
        lockedBalance: wallet.lockedBalance,
        equity:
        Number(wallet.availableBalance) +
        unrealizedPnl,
    };
    }
    
    async deposit(
    userId: string,
    amount: number,
    ) {
    const wallet = await this.prisma.wallet.findUnique({
        where: {
        userId,
        },
    });

    if (!wallet) {
        throw new NotFoundException();
    }

    const updatedWallet =
        await this.prisma.wallet.update({
        where: {
            userId,
        },
        data: {
            availableBalance: {
            increment: amount,
            },
        },
        });

    await this.prisma.transaction.create({
        data: {
        walletId: wallet.id,
        amount,
        type: 'DEPOSIT',
        },
    });

    return {
        message: 'Deposit successful',
        availableBalance:
        updatedWallet.availableBalance,
    };
    }

    async withdraw(
    userId: string,
    amount: number,
    ) {
    const wallet = await this.prisma.wallet.findUnique({
        where: {
        userId,
        },
    });

    if (!wallet) {
        throw new NotFoundException();
    }

    if (
        Number(wallet.availableBalance) <
        amount
    ) {
        throw new BadRequestException(
        'Insufficient balance',
        );
    }

    const updatedWallet =
        await this.prisma.wallet.update({
        where: {
            userId,
        },
        data: {
            availableBalance: {
            decrement: amount,
            },
        },
        });

    await this.prisma.transaction.create({
        data: {
        walletId: wallet.id,
        amount,
        type: 'WITHDRAWAL',
        },
    });

    return {
        message: 'Withdrawal successful',
        availableBalance:
        updatedWallet.availableBalance,
    };
    }

    async transactions(userId: string) {
    const wallet = await this.prisma.wallet.findUnique({
        where: {
        userId,
        },
    });

    if (!wallet) {
        throw new NotFoundException();
    }

    return this.prisma.transaction.findMany({
        where: {
        walletId: wallet.id,
        },
        select: {
        id: true,
        type: true,
        amount: true,
        },
        orderBy: {
        createdAt: 'desc',
        },
    });
    }
}