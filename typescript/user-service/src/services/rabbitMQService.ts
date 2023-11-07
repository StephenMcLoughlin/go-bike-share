import { Injectable } from "@nestjs/common"
import { ClientProxy, ClientProxyFactory, Transport } from "@nestjs/microservices"
import { CreateUserDto } from "src/user/create-user.dto"


@Injectable()
export class RabbitMQService {

    private client: ClientProxy

    constructor() {
        this.client = ClientProxyFactory.create({
            transport: Transport.RMQ,
            options:{
                urls: [process.env.RABBITMQ_URL],
                queue: 'tasks',
            }
        })
    }

    async sendUserMessage(user: CreateUserDto) {
        await this.client.emit('user_created', user)
    }
}