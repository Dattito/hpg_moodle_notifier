import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IndexSiteComponent } from './index-site.component';

describe('IndexSiteComponent', () => {
  let component: IndexSiteComponent;
  let fixture: ComponentFixture<IndexSiteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IndexSiteComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(IndexSiteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
