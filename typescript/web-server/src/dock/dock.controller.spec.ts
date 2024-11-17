import { Test, TestingModule } from '@nestjs/testing';
import { DockController } from './dock.controller';

describe('DockController', () => {
  let controller: DockController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [DockController],
    }).compile();

    controller = module.get<DockController>(DockController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
