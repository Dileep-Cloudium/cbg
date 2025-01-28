import { Component } from '@angular/core';
import { Router, RouterModule } from '@angular/router';
import { ButtonModule } from '@syncfusion/ej2-angular-buttons';
import { ProtectedService } from './protected.service';
import { AppService } from '../app.service';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-protected',
  imports: [ButtonModule, RouterModule],  
  templateUrl: './protected.component.html',
  styleUrl: './protected.component.css',
  standalone: true,
})
export class ProtectedComponent {

  constructor(private protectedService:ProtectedService,private cookieService:CookieService,private appService:AppService,private router:Router){
  }
  
  
  navigate(path: string) {
    this.router.navigate([path]);
  }

  signOut(){
    this.appService.isLoading = true;
    this.appService.handleSignOut().then(() => {
      this.cookieService.deleteAll("/");
      this.appService.isLoggedIn = false;
      this.appService.isLoading = false;
      this.router.navigate(['public']);
    })
  }
}
