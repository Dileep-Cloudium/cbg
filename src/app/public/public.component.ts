import { Component, ViewChild } from '@angular/core';
import { ButtonModule } from '@syncfusion/ej2-angular-buttons';
import { Router, RouterModule } from '@angular/router';
import { DialogComponent, DialogModule } from '@syncfusion/ej2-angular-popups';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { MaskedTextBoxModule } from '@syncfusion/ej2-angular-inputs';
import { ForgotPasswordComponent } from './forgot-password/forgot-password.component';

@Component({
  selector: 'app-public',
  standalone: true,
  imports: [ 
    ButtonModule, 
    RouterModule,
    DialogModule,
    CommonModule,
    MaskedTextBoxModule,
    LoginComponent,
    RegisterComponent,
    ForgotPasswordComponent
  ],
  templateUrl: './public.component.html',
  styleUrl: './public.component.css'
})
export class PublicComponent {
  
  // Reference the Dialog element
  @ViewChild('ejDialog') ejDialog: DialogComponent | any;

  displayComponent: string = '';

  constructor(private router: Router) {
    console.log('public component constructor');
    
  }

  ngOnInit(): void {}

  navigate(path: string) {
    console.log('navigate', path);
    
    this.router.navigate([path]);
  }


  showDialog(path: any) {
    console.log('path',path);
    console.log('ejDialog',this.ejDialog);
    
    this.ejDialog.show();
    if(path == 'login'){
      this.displayComponent = 'login';
    }else if(path == 'register'){
      this.displayComponent = 'register';
    }else if(path == 'forgot-password'){
      this.displayComponent = 'forgot-password';
    }
    console.log('displayComponent',this.displayComponent);
  }
  
  closeDialog(path: any){
    this.ejDialog.hide();
    this.router.navigate([path]);
  }
}
