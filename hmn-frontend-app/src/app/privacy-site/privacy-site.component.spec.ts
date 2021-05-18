import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PrivacySiteComponent } from './privacy-site.component';

describe('PrivacySiteComponent', () => {
  let component: PrivacySiteComponent;
  let fixture: ComponentFixture<PrivacySiteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PrivacySiteComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PrivacySiteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
