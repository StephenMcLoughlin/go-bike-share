import { Body, Controller, Delete, Get, Param, Post, Put } from '@nestjs/common';
import { CreateUserDto } from './create-user.dto';
import { UserService } from './user.service';
import { RabbitMQService } from 'src/services/rabbitMqService';

@Controller('user')
export class UserController {

    constructor(private userService: UserService, private rabbitMQService: RabbitMQService){}

    @Get(':id')
    getUser(@Param('id') id: number) {
        this.userService.getUser(id)
    }

    @Post()
    createUser(@Body() createUserDto: CreateUserDto) {
        this.userService.createUser(createUserDto)
        this.rabbitMQService.sendUserMessage(createUserDto)
    }

    @Put(':id')
    updateUser(@Param('id') id: number, @Body()createUserDto: CreateUserDto) {
        return this.userService.updateUser(id, createUserDto)
    }

    @Delete(':id')
    deleteUser(@Param('id') id: number) {
        return this.userService.deleteUser(id)
    }
}
