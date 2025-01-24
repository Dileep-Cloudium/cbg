import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BenifitsComponent } from './benifits.component';

describe('BenifitsComponent', () => {
  let component: BenifitsComponent;
  let fixture: ComponentFixture<BenifitsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BenifitsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BenifitsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
