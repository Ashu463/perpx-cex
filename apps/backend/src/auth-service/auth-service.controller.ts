import {
  Body,
  Controller,
  Get,
  Post,
  Req,
} from '@nestjs/common';
import { AuthService } from './auth-service.service';

@Controller('auth')
export class AuthController {
  constructor(
    private readonly authService: AuthService,
  ) {}

  @Post('signup')
  signup(
    @Body()
    body: {
      username: string;
      password: string;
      email: string;
    },
  ) {
    return this.authService.signup(body);
  }

  @Post('login')
  login(
    @Body()
    body: {
      username: string;
      password: string;
    },
  ) {
    return this.authService.login(body);
  }

  @Post('refresh-token')
  refreshToken(
    @Body()
    body: {
      refreshToken: string;
    },
  ) {
    return this.authService.refreshToken(body);
  }

  @Post('logout')
  logout() {
    return this.authService.logout();
  }

  @Get('me')
  me(@Req() req: any) {
    return this.authService.me(
      req.user?.sub,
    );
  }
}