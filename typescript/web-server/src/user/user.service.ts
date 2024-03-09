import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { User } from './user.entity';
import { CreateUserDto, LoginUserDto } from './user.dto';
import * as bcrypt from 'bcrypt';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(User) private readonly userRepository: Repository<User>,
  ) {}

  async createUser(createUserDto: CreateUserDto): Promise<User> {
    const user: User = new User();
    user.firstName = createUserDto.firstName;
    user.lastName = createUserDto.lastName;
    user.email = createUserDto.email;
    user.password = await bcrypt.hash(createUserDto.password, 12);

    return this.userRepository.save(user);
  }

  getUserByEmail(email: string): Promise<User> {
    return this.userRepository.findOneBy({ email });
  }

  updateUser(id: number, createUserDto: CreateUserDto): Promise<User> {
    const user: User = new User();
    user.firstName = createUserDto.firstName;
    user.lastName = createUserDto.lastName;
    user.email = createUserDto.email;
    user.password = createUserDto.password;
    user.id = id;

    return this.userRepository.save(user);
  }

  deleteUser(id: number): Promise<{ affected?: number }> {
    return this.userRepository.delete(id);
  }
}
