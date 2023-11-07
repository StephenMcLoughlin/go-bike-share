import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { User } from './user.entity';
import { CreateUserDto } from './create-user.dto';

@Injectable()
export class UserService {
    constructor(@InjectRepository(User) private readonly userRepository: Repository<User>){}

    createUser(createUserDto: CreateUserDto): Promise<User> {
        const user: User = new User()
        user.firstName = createUserDto.firstName
        user.lastName = createUserDto.lastName
        user.email = createUserDto.email
        user.password = createUserDto.password

        return this.userRepository.save(user)
    }

    getUser(id: number): Promise<User> {
        return this.userRepository.findOneBy({id})
    }

    updateUser(id: number, createUserDto: CreateUserDto): Promise<User> {
        const user: User = new User()
        user.firstName = createUserDto.firstName
        user.lastName = createUserDto.lastName
        user.email = createUserDto.email
        user.password = createUserDto.password
        user.id = id

        return this.userRepository.save(user)
    }

    deleteUser(id: number): Promise<{ affected?: number }> {
        return this.userRepository.delete(id)
    }
}