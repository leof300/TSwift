import { TestBed } from '@angular/core/testing';

import { GoApiService } from './go-api.service';

describe('GoApiService', () => {
  let service: GoApiService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GoApiService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
