import {
  Body,
  Controller,
  Get,
  HttpCode,
  HttpStatus,
  Post,
  UseGuards,
  Request,
} from '@nestjs/common';
import { CreateUserDto, LoginUserDto } from 'src/user/user.dto';

import { RabbitMQService } from 'src/services/rabbitMqService';
import { JwtService } from '@nestjs/jwt';
import { AuthService } from './auth.service';
import { AuthGuard } from './auth.guard';
import { Public } from 'src/shared/public.decorator';

@Controller('api/auth')
export class AuthController {
  constructor(
    private authService: AuthService,
    private rabbitMQService: RabbitMQService,
    private jwtService: JwtService,
  ) {}

  @Public()
  @Post('register')
  async register(@Body() createUserDto: CreateUserDto) {
    return await this.authService.register(createUserDto);
    // this.userService.createUser(createUserDto);
    // this.rabbitMQService.sendUserMessage(createUserDto);
  }

  @Public()
  @HttpCode(HttpStatus.OK)
  @Post('login')
  async login(@Body() loginUserDto: LoginUserDto) {
    return await this.authService.login(
      loginUserDto.email,
      loginUserDto.password,
    );
  }
}
