import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AlertService } from 'src/app/alert.service';
import { HmnBackendService } from 'src/app/hmn-backend.service';

@Component({
  selector: 'app-signal-form',
  templateUrl: './signal-form.component.html',
  styleUrls: ['./signal-form.component.css']
})
export class SignalFormComponent implements OnInit {

  constructor(
    private formBuilder: FormBuilder,
    private alertService: AlertService,
    private hmn: HmnBackendService
  ) { }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      phoneNumber: ['', Validators.required]
    })
  }

  get f() {return this.form.controls}

  submitted = false;
  loading = false;
  enterVerificationCode = false;
  verificationID = "";
  @Input() moodleToken: string = "";
  @Input() logout: boolean = false;
  @Output() onSuccess = new EventEmitter();

  form!: FormGroup;

  onVerificationCodeSuccess() {
    this.onSuccess.emit();
  }

  onFormSubmit() {
    this.enterVerificationCode = false;
    this.submitted = true;

    this.alertService.clear();

    if (this.form.invalid) return;

    this.loading = true;

    this.hmn.getSignalVerificationCode(this.moodleToken, "+49" + this.f.phoneNumber.value)
    .subscribe(resp => {
      this.loading = false;
      this.verificationID = resp.id;
      this.enterVerificationCode = true;
    }, err => {
      this.loading = false;
    });
  }
}