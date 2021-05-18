import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MoodleFormComponent } from './moodle-form.component';

describe('MoodleFormComponent', () => {
  let component: MoodleFormComponent;
  let fixture: ComponentFixture<MoodleFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MoodleFormComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MoodleFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
