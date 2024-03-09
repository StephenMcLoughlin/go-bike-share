import {
  BadRequestException,
  Body,
  Controller,
  Post,
  UnauthorizedException,
} from '@nestjs/common';
import { CreateUserDto, LoginUserDto } from 'src/user/user.dto';
import * as bcrypt from 'bcrypt';

import { RabbitMQService } from 'src/services/rabbitMqService';
import { User } from 'src/user/user.entity';
import { JwtService } from '@nestjs/jwt';
import { AuthService } from './auth.service';

type LoggedInUser = Omit<User, 'password'>;

@Controller('api/auth')
export class AuthController {
  constructor(
    private authService: AuthService,
    private rabbitMQService: RabbitMQService,
    private jwtService: JwtService,
  ) {}

  @Post('register')
  async register(@Body() createUserDto: CreateUserDto) {
    return await this.authService.register(createUserDto);
    // this.userService.createUser(createUserDto);
    // this.rabbitMQService.sendUserMessage(createUserDto);
  }

  @Post('login')
  async login(@Body() loginUserDto: LoginUserDto) {
    return await this.authService.login(
      loginUserDto.email,
      loginUserDto.password,
    );
  }
}
