import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginSiteComponent } from './login-site/login-site.component';
import { RegistrationFormComponent } from './login-site/registration-form/registration-form.component';
import { NavigationBarComponent } from './navigation-bar/navigation-bar.component';
import { HowItWorksSiteComponent } from './how-it-works-site/how-it-works-site.component';
import { PrivacySiteComponent } from './privacy-site/privacy-site.component';
import { NotfoundSiteComponent } from './notfound-site/notfound-site.component';
import { IndexSiteComponent } from './index-site/index-site.component';
import { HmnBackendService } from './hmn-backend.service';
import { HttpErrorInterceptor } from './http-error.interceptor';
import { MoodleFormComponent } from './login-site/registration-form/moodle-form/moodle-form.component';
import { SignalFormComponent } from './login-site/registration-form/signal-form/signal-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { AlertService } from './alert.service';
import { AlertComponent } from './alert/alert.component';
import { VerificationFormComponent } from './login-site/registration-form/signal-form/verification-form/verification-form.component';
import { NgbCollapseModule, NgbDropdownModule, NgbNavbar, NgbNavModule } from '@ng-bootstrap/ng-bootstrap';
import { ContactSiteComponent } from './contact-site/contact-site.component';

@NgModule({
  declarations: [
    AppComponent,
    RegistrationFormComponent,
    NavigationBarComponent,
    LoginSiteComponent,
    HowItWorksSiteComponent,
    PrivacySiteComponent,
    NotfoundSiteComponent,
    IndexSiteComponent,
    MoodleFormComponent,
    SignalFormComponent,
    AlertComponent,
    VerificationFormComponent,
    ContactSiteComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    NgbNavModule,
    NgbCollapseModule,
    NgbDropdownModule
  ],
  providers: [
    HmnBackendService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: HttpErrorInterceptor,
      multi: true
    },
    AlertService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
