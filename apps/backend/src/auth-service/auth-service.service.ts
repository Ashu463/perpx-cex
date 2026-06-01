import {
  Injectable,
  UnauthorizedException,
  ConflictException,
  NotFoundException,
} from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcrypt';
import { randomUUID } from 'crypto';

@Injectable()
export class AuthService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly jwtService: JwtService,
  ) {}

  async signup(body: {
    username: string;
    password: string;
    email: string
  }) {
    const { username, password, email } = body;
    
    const existingUser = await this.prisma.user.findUnique({
      where: {
        username,
      },
    });
    
    if (existingUser) {
      throw new ConflictException('Username already exists');
    }
    // captcha implementation to be expected on front end
    // email verification

    const hashedPassword = await bcrypt.hash(password, 10);

    const user = await this.prisma.user.create({
      data: {
        id: randomUUID(),
        username,
        password: hashedPassword,
        email: email,
        createdAt: new Date(),
        updatedAt:  new Date(),
        isActive: true

      },
    });

    return {
      id: user.id,
      username: user.username,
      createdAt: user.createdAt,
    };
  }

  async login(body: {
    username: string;
    password: string;
  }) {
    // KYC to be done
    const { username, password } = body;

    const user = await this.prisma.user.findUnique({
      where: {
        username,
      },
    });

    if (!user) {
      throw new UnauthorizedException('Invalid credentials');
    }

    const isPasswordValid = await bcrypt.compare(
      password,
      user.password,
    );

    if (!isPasswordValid) {
      throw new UnauthorizedException('Invalid credentials');
    }

    const accessToken = await this.jwtService.signAsync({
      sub: user.id,
      username: user.username,
    });

    const refreshToken = await this.jwtService.signAsync(
      {
        sub: user.id,
      },
      {
        expiresIn: '7d',
      },
    );

    return {
      accessToken,
      refreshToken,
    };
  }

  async refreshToken(body: {
    refreshToken: string;
  }) {
    const payload = await this.jwtService.verifyAsync(
      body.refreshToken,
    );

    const accessToken = await this.jwtService.signAsync({
      sub: payload.sub,
    });

    return {
      accessToken,
    };
  }

  async logout() {
    return {
      message: 'Logged out successfully',
    };
  }

  async me(userId: string) {
    const user = await this.prisma.user.findUnique({
      where: {
        id: userId,
      },
    });

    if (!user) {
      throw new NotFoundException();
    }

    return {
      id: user.id,
      username: user.username,
    };
  }
}