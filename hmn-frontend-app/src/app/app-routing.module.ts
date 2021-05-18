import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HowItWorksSiteComponent } from './how-it-works-site/how-it-works-site.component';
import { IndexSiteComponent } from './index-site/index-site.component';
import { LoginSiteComponent } from './login-site/login-site.component';
import { NotfoundSiteComponent } from './notfound-site/notfound-site.component';
import { PrivacySiteComponent } from './privacy-site/privacy-site.component';

const routes: Routes = [
  {path: "how-it-works", component: HowItWorksSiteComponent},
  {path: "privacy", component: PrivacySiteComponent},
  {path: "login", component: LoginSiteComponent},
  {path: "", component: IndexSiteComponent},
  {path: "**", component: NotfoundSiteComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
