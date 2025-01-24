import { Component } from '@angular/core';
import { Router, RouterModule } from '@angular/router';
import { ButtonModule } from '@syncfusion/ej2-angular-buttons';
import { AppBarModule } from '@syncfusion/ej2-angular-navigations';
import { ProtectedService } from './protected.service';
import { AppService } from '../app.service';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-protected',
  imports: [AppBarModule, ButtonModule, RouterModule],  
  templateUrl: './protected.component.html',
  styleUrl: './protected.component.css'
})
export class ProtectedComponent {

  constructor(private protectedService:ProtectedService,private cookieService:CookieService,private appService:AppService,private router:Router){}
  
  signOut(){
    this.appService.handleSignOut().then(() => {
      this.cookieService.deleteAll("/");
      this.appService.isLoggedIn = false;
      this.appService.isLoading = false;
      this.router.navigate(['login']);
    })
  }
}
