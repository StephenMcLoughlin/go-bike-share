import { Test, TestingModule } from '@nestjs/testing';
import { DockService } from './dock.service';

describe('DockService', () => {
  let service: DockService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [DockService],
    }).compile();

    service = module.get<DockService>(DockService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
