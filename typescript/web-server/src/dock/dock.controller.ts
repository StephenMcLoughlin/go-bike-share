import { Controller } from '@nestjs/common';
import { DockService } from './dock.service';

@Controller('api/dock')
export class DockController {
  constructor(private dockService: DockService){};

  
}
