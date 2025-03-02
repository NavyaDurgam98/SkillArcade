import { TestBed } from '@angular/core/testing';

import { TakequizService } from './takequiz.service';

describe('TakequizService', () => {
  let service: TakequizService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TakequizService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
