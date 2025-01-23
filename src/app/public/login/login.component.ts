import { CommonModule } from '@angular/common';
import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { PublicService } from '../public.service';
import { HttpClientModule } from '@angular/common/http';
import { MaskedTextBoxModule, TextBoxModule } from '@syncfusion/ej2-angular-inputs';
import { ButtonModule } from '@syncfusion/ej2-angular-buttons';
import { AppService } from '../../app.service';
import { ButtonPropsModel, DialogComponent, DialogModule } from '@syncfusion/ej2-angular-popups';
import { AuthTokens, AuthUser, SignInOutput } from 'aws-amplify/auth';
// import { AssociateSoftwareTokenResponse } from '@aws-sdk/client-cognito-identity-provider';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule,
    ButtonModule,
    HttpClientModule,
    TextBoxModule,
    DialogModule,
    MaskedTextBoxModule
  ],
  
})

export class LoginComponent implements OnInit {

  /**
   * Reference to DialogComponent
   */
  @ViewChild('Dialog')
  public totpDialog: DialogComponent | undefined;


    /**
     * Form group for login form
     */
    loginForm: FormGroup;

  /**
   * Variable to store code entered for authentication
   */
  code: string;

  /**
   * Flag to control visibility of MFA setup popup
   */
  showMfaSetupPopup = false;

  isReset = false;

  /**
   * Flag to control visibility of TOTP popup
   */
  showTotpPopup = false;

  /**
   * Flag to indicate whether login was successful
   */
  showLoginSuccess = false;

  /**
   * Flag to control visibility of a button
   */
  showButton = false;

  /**
   * Flag to control visibility of eye icon
   */
  showEye = false;

  /**
   * Flag indicating whether the registration process is confirmed
   */
  isConfirm: boolean = false;

    /**
   * Adding button and its functionality to dialog
   */
    public dlgButtons: ButtonPropsModel[] = [{  buttonModel: { content: 'Ok', isPrimary: true, cssClass: "e-suucess" } }];

    /**
   * Injecting dependencies
   * @param router - Angular router for navigation
   * @param publicService - Service for handling public API
   * @param formBuilder - Form Builder instance for login form
   * @param cookieService - Handle cookie storage values
   * @param appService - Service for handling app related API
   */
    constructor(
      private router: Router, 
      private publicService: PublicService,
      private formBuilder: FormBuilder, 
      public appService: AppService) {
      this.code = "";
      this.loginForm = this.formBuilder.group({
        email: ['', [Validators.required, Validators.email]],
        password: ['', [Validators.required]],
        code: ['']
      })
    }
  

  ngOnInit(): void {
    console.log('LoginComponent initialized');
  }

  
  /**
   * Function to show/hide Password
   */
  showPassword() {
    this.showButton = !this.showButton;
    this.showEye = !this.showEye;
  }

  
  /**
   * Login function with AWS
   */
  public signIn(): void {
    this.appService.isLoading = true;
    this.publicService.signIn(this.loginForm.value.email, this.loginForm.value.password)
      .then(async (resp: SignInOutput) => {
        this.appService.isLoading = false;
        const { nextStep } = resp;
        console.log(nextStep);
        switch (nextStep.signInStep) {
          case 'CONTINUE_SIGN_IN_WITH_TOTP_SETUP':
            break;
          case 'CONFIRM_SIGN_IN_WITH_TOTP_CODE':
            // You need to get the code from the UI inputs
            // and then trigger the following function with a button click
            this.showTotpPopup = true;
            this.totpDialog?.show();
            break;
          case 'CONFIRM_SIGN_IN_WITH_SMS_CODE':
            break;
          case 'CONFIRM_SIGN_UP':
            break;
          case 'RESET_PASSWORD':
            break;
          case 'DONE':
            this.appService.openToaster("success", "Successfully logged in");
            this.router.navigate(['/policies']);
            break;
        }
      }).catch((e) => {
        this.appService.isLoading = false;
        console.log(e, "error");
        this.appService.openToaster("error", e.message ?? 'An error occurred');
      });
  }


  /**
   * After successful login, set the tokens into cookie storage and navigate to profile page
   * @param user - authenticated user details
   */
  async loginSuccess(user: AuthUser) {
  }


  getFieldCssClass(fieldName: string): string {
    const control = this.loginForm.get(fieldName);
    return control?.invalid && control?.touched ? 'error' : 'success';
  }

  onForgotPassword(): void {
    // Implement forgot password logic
  }


  /**
   * Validate the field to show / hide the error message
   * @param field - check for validation
   * @returns boolean
   */
  public isFieldValid(field: string) {
    return !(this.loginForm.get(field)?.valid === true) && (this.loginForm.get(field)?.dirty === true || this.loginForm.get(field)?.touched === true);
  }

  
  /**
   * Navigation functioanlity
   * @param val - path
   */
  navigate(val: string) {
    this.router.navigate([val]);
  }
}
