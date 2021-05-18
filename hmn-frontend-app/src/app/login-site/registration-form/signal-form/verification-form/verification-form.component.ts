import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AlertService } from 'src/app/alert.service';
import { HmnBackendService } from 'src/app/hmn-backend.service';

@Component({
  selector: 'app-verification-form',
  templateUrl: './verification-form.component.html',
  styleUrls: ['./verification-form.component.css']
})
export class VerificationFormComponent implements OnInit {

  constructor(
    private hmn: HmnBackendService,
    private alertService: AlertService,
    private formBuilder: FormBuilder
  ) { }

  submitted = false;
  loading = false;
  @Input() verificationID = "";
  @Input() logout: boolean = false;
  @Output() onSuccess = new EventEmitter();

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      verificationCode: ['',[ Validators.required, Validators.min(111111), Validators.max(999999)]]
    })
  }

  form!: FormGroup;
  get f() {return this.form.controls}

  onFormSubmit() {
    this.submitted = true;

    this.alertService.clear();

    if (this.form.invalid) return;

    this.loading = true;

    if (this.logout) {
      this.hmn.deleteAssignments(this.verificationID, this.f.verificationCode.value)
      .subscribe(resp => {
        this.onSuccess.emit();
        this.loading = false;
      }, err => {
        this.loading = false;
      })
    } else {
      this.hmn.registerAssignments(this.verificationID, this.f.verificationCode.value)
      .subscribe(resp => {
        this.onSuccess.emit();
        this.loading = false;
      }, err => {
        this.loading = false;
      })
    }
    
  }
}
