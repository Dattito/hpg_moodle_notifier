import { TestBed } from '@angular/core/testing';

import { HmnBackendService } from './hmn-backend.service';

describe('HmnBackendService', () => {
  let service: HmnBackendService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HmnBackendService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
