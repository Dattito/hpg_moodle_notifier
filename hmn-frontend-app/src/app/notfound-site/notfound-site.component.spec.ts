import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NotfoundSiteComponent } from './notfound-site.component';

describe('NotfoundSiteComponent', () => {
  let component: NotfoundSiteComponent;
  let fixture: ComponentFixture<NotfoundSiteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NotfoundSiteComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NotfoundSiteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
