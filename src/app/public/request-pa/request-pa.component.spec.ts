import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RequestPaComponent } from './request-pa.component';

describe('RequestPaComponent', () => {
  let component: RequestPaComponent;
  let fixture: ComponentFixture<RequestPaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RequestPaComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RequestPaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
