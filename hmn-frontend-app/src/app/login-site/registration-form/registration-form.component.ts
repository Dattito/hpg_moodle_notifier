import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-registration-form',
  templateUrl: './registration-form.component.html',
  styleUrls: ['./registration-form.component.css']
})
export class RegistrationFormComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }
  
  public verificationID = "";
  public phonenumber = "";
  public verificationCode = "";
  public moodleToken = "";

  @Input() logout: boolean = false; 

  public errors = "";
  public isSubmitting = false;

  public activeTab = 1;

  readonly navDisabled = true;

  onReceiveMoodleToken(event: any) {
    this.moodleToken = event;
    this.activeTab = 2;
  }

  onSignalVerificationSuccess() {
    this.activeTab = 3;
  }

}
