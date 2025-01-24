import { TestBed } from '@angular/core/testing';

import { Interceptor } from './interceptor.interceptor';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('InterceptorInterceptor', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    providers: [
      Interceptor
    ]
  }));

  it('should be created', () => {
    const interceptor: Interceptor = TestBed.inject(Interceptor);
    expect(interceptor).toBeTruthy();
  });
});
