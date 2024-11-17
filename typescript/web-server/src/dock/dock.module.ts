import { Module } from '@nestjs/common';
import { DockService } from './dock.service';
import { DockController } from './dock.controller';

@Module({
  providers: [DockService],
  controllers: [DockController]
})
export class DockModule {}
