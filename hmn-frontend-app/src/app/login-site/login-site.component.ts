import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HmnBackendService } from '../hmn-backend.service';

@Component({
  selector: 'app-login-site',
  templateUrl: './login-site.component.html',
  styleUrls: ['./login-site.component.css']
})
export class LoginSiteComponent implements OnInit {

  constructor(
    private hbs: HmnBackendService,
    private route: ActivatedRoute  
  ) { }

  logout: boolean = false;

  ngOnInit(): void {
    this.route.queryParams.subscribe(
      params => {
        const value = params.logout
        this.logout = value ? value.toLocaleLowerCase() === 'true': false;
      }
    )
        
  }
}
