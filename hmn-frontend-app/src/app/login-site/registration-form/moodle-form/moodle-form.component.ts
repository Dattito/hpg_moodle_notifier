import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AlertService } from 'src/app/alert.service';
import { HmnBackendService } from 'src/app/hmn-backend.service';

@Component({
  selector: 'app-moodle-form',
  templateUrl: './moodle-form.component.html',
  styleUrls: ['./moodle-form.component.css']
})
export class MoodleFormComponent implements OnInit {
  

  constructor(
    private formBuilder: FormBuilder,
    private alertService: AlertService,
    private hmn: HmnBackendService
  ) { }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    })
  }

  get f() {return this.form.controls}

  submitted = false;
  loading = false;

  form!: FormGroup;

  @Output() receiveMoodleToken = new EventEmitter();

  onFormSubmit() {
    this.submitted = true;

    this.alertService.clear();

    if (this.form.invalid) return;

    this.loading = true;

    this.hmn.getMoodleToken(this.f.username.value, this.f.password.value)
    .subscribe(resp => {
      this.loading = false;
      this.receiveMoodleToken.emit(resp.token);
    }, err => {
      this.loading = false;
    });
  }
}
