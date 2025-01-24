import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RxClaimsComponent } from './rx-claims.component';

describe('RxClaimsComponent', () => {
  let component: RxClaimsComponent;
  let fixture: ComponentFixture<RxClaimsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RxClaimsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RxClaimsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
