import {
  BadRequestException,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { CreateUserDto } from 'src/user/user.dto';
import { UserService } from 'src/user/user.service';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AuthService {
  constructor(
    private userService: UserService,
    private jwtService: JwtService,
  ) {}

  async register(createUserDto: CreateUserDto) {
    try {
      const user = await this.userService.createUser(createUserDto);

      const { password, ...result } = user;
      return result;
    } catch (error) {
      throw new BadRequestException();
    }
  }

  async login(
    email: string,
    password: string,
  ): Promise<{ access_token: string }> {
    const user = await this.userService.getUserByEmail(email);

    if (!user) {
      throw new UnauthorizedException();
    }

    const isMatch = await bcrypt.compare(password, user.password);

    if (!isMatch) {
      throw new UnauthorizedException();
    }

    const payload = {
      sub: user.id,
      firstName: user.firstName,
      lastName: user.lastName,
    };

    return {
      ...payload,
      access_token: await this.jwtService.signAsync(payload),
    };
  }
}
