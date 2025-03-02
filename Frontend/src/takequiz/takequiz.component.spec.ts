import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TakequizComponent } from './takequiz.component';

describe('TakequizComponent', () => {
  let component: TakequizComponent;
  let fixture: ComponentFixture<TakequizComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TakequizComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TakequizComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
