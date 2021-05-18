import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HowItWorksSiteComponent } from './how-it-works-site.component';

describe('HowItWorksSiteComponent', () => {
  let component: HowItWorksSiteComponent;
  let fixture: ComponentFixture<HowItWorksSiteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HowItWorksSiteComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HowItWorksSiteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
