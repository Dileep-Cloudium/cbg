import { CommonModule } from '@angular/common';
import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { MaskedTextBoxModule, TextBoxModule } from '@syncfusion/ej2-angular-inputs';
import { PublicService } from '../public.service';
import { ButtonModule } from '@syncfusion/ej2-angular-buttons';
import { AppService } from '../../app.service';
import { ButtonPropsModel, DialogComponent, DialogModule } from '@syncfusion/ej2-angular-popups';
import { AuthTokens, AuthUser, SignInOutput } from 'aws-amplify/auth';
import { QRCodeGeneratorAllModule } from '@syncfusion/ej2-angular-barcode-generator';
import { environment } from '../../../environments/environment.develop';

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
    TextBoxModule,
    DialogModule,
    MaskedTextBoxModule,
    QRCodeGeneratorAllModule
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
  public dlgButtons: ButtonPropsModel[] = [{ click: this.confirmTotp.bind(this), buttonModel: { content: 'Ok', isPrimary: true, cssClass: "e-suucess" } }];

  /**
   * Injecting dependencies
   * @param router - Angular router for navigation
   * @param publicService - Service for handling public API
   * @param formBuilder - Form Builder instance for login form
   * @param appService - Service for handling app related API
   */
  constructor(private router: Router, private publicService: PublicService,
    private formBuilder: FormBuilder, public appService: AppService) {
    this.code = "";
    this.loginForm = this.formBuilder.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required]],
      code: ['']
    })
  }

  /**
   * Lifecycle hook that executes when the component initializes, handling initializations, 
   * data retrieval, or subscriptions for optimal component setup
   */
  ngOnInit(): void {
    this.loginForm.get('email')?.valueChanges.subscribe((value: string) => {
      if (value != "") {
        this.loginForm.get('email')?.setValue(value.toLowerCase(), { emitEvent: false });
      }
    });

    // Check whether the user logged in or not
    this.appService.getCurrentUser().then((resp: AuthUser) => {
      if ("userId" in resp) {
        this.loginSuccess(resp)
      }
    })
  }

  /**
   * Function to show/hide Password
   */
  showPassword() {
    this.showButton = !this.showButton;
    this.showEye = !this.showEye;
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
   * Login function with AWS
   */
  public signIn(): void {
    this.appService.isLoading = true;
    this.publicService.signIn(this.loginForm.value.email, this.loginForm.value.password)
      .then(async (resp: SignInOutput) => {
        this.appService.isLoading = false;
        const { nextStep } = resp;
        console.log('response',nextStep)
        switch (nextStep.signInStep) {
          case 'CONTINUE_SIGN_IN_WITH_TOTP_SETUP':
            // This happens when the MFA method is TOTP
            // The user needs to setup the TOTP before using it
            this.code = nextStep.totpSetupDetails.getSetupUri("SagePA" + ((environment.production === true) ? "" : `-${environment.name}`)).href;
            this.code = this.code.replace(this.code.split(":")[2].slice(0, 36), this.loginForm.value.email);
            this.showMfaSetupPopup = true;
            break;
          case 'CONFIRM_SIGN_IN_WITH_TOTP_CODE':
            // You need to get the code from the UI inputs
            // and then trigger the following function with a button click
            this.showTotpPopup = true;
            this.totpDialog?.show();
            break;
          case 'CONFIRM_SIGN_IN_WITH_SMS_CODE':
            // If MFA is enabled, sign-in should be confirmed with the confirmation code
            this.showTotpPopup = true;
            this.totpDialog?.show();
            break;
          case 'CONFIRM_SIGN_UP':
            this.resendSignUpCode()
            break;
          case 'RESET_PASSWORD':
            this.router.navigate(['forgot-password']);
          break;
          case 'DONE': {
            const session: AuthTokens = await this.appService.getAuthTokens();
            this.publicService.tokenAssociation(session.accessToken.toString()).then(async (resp: any) => {
              this.code = `otpauth://totp/${this.loginForm.value.email}?secret=${resp.SecretCode}&issuer=SagePA-${environment.name}`;
              this.showMfaSetupPopup = true;
              this.isReset = true;
            }).catch((e: { message: string | null | undefined; }) => {
              this.appService.isLoading = false;
              this.appService.openToaster("error", e.message);
            });
          }
        }
      }).catch((e) => {
        this.appService.isLoading = false;
        this.appService.openToaster("error", e.message);
      });
  }

  /**
   * Method to resend sign-up verification code to the user's email
   */
  resendSignUpCode(): void {
    this.appService.isLoading = true;
    this.publicService.resendSignUpCode(this.loginForm.value.email)
      .then(() => {
        this.appService.isLoading = false;
        this.appService.openToaster("success", "Verification code sent to email");
        this.isConfirm = true;
      }).catch((e: { message: string | null | undefined; }) => {
        this.appService.isLoading = false;
        this.appService.openToaster("error", e.message);
      });
  }

  /**
   * To confirm the user email, need to submit the code which was received
   */
  confirmSignUp() {
    this.appService.isLoading = true;
    this.publicService.confirmSignUp(this.loginForm.value.email, this.loginForm.value.code)
      .then(() => {
        this.appService.isLoading = false
        this.appService.openToaster("success", "Email Verified")
        this.isConfirm = false;
        this.signIn()
      }).catch((e) => {
        this.appService.isLoading = false;
        this.appService.openToaster("error", e.message)
      });
  }

  /**
   * Show the login form
   */
  backToLogin() {
    this.showMfaSetupPopup = false;
  }

  /**
   * On submit of TOTP code
   */
  async confirmTotp() {
    const value = (document.getElementById("totp") as HTMLInputElement).value.replace(/ |_/g, '');
    if (value == "" || value.length !== 6) {
      this.appService.openToaster("warning", "Please enter six digit code from your authenticator app");
    }
    else {
      this.appService.isLoading = true;
      if (this.isReset === true) {
        const session: AuthTokens = await this.appService.getAuthTokens();
        this.publicService.validateTOTP(session.accessToken.toString(), value).then(async (resp:any) => {
          console.log(resp)
          this.appService.isLoading = false;
          this.publicService.setUserMFA(session.accessToken.toString()).then(async (resp:any) => {
            console.log(resp)
            const user: AuthUser = await this.appService.getCurrentUser();
            this.showLoginSuccess = true;
            this.loginSuccess(user);
          }).catch((e:any) => {
            this.appService.isLoading = false;
            this.appService.openToaster("error", e.message)
          });
        }).catch((e:any) => {
          this.appService.isLoading = false;
          this.appService.openToaster("error", e.message)
        });
      } else {
        this.publicService.confirmSignin(value).then(async (resp) => {
          this.appService.isLoading = false;
          if ("isSignedIn" in resp) {
            const user: AuthUser = await this.appService.getCurrentUser();
            this.showLoginSuccess = true;
            this.loginSuccess(user);
          } else {
            this.appService.openToaster("error", ("message" in resp) ? resp?.message : "Login failed");
          }
        })
      }
      this.totpDialog?.hide();
      this.showTotpPopup = false;
    }
  }

  /**
   * After successful login, set the tokens into cookie storage and navigate to profile page
   * @param user - authenticated user details
   */
  async loginSuccess(user: AuthUser) {
    const session: AuthTokens = await this.appService.getAuthTokens();
    this.appService.getUser(session.accessToken.toString()).then(async (userResp:any) => {
      if ("UserMFASettingList" in userResp) {
        // this.cookieService.set('access_token', session.accessToken.toString(), 1, "/");
        // this.cookieService.set('id_token', session?.idToken?.toString() ?? "", 1, "/");
        // this.cookieService.set('email', user?.signInDetails?.loginId ?? "", 1, "/");
        // localStorage.setItem('email', user?.signInDetails?.loginId ?? "");
        // this.cookieService.set('id', user.userId, 1, "/");
        this.appService.isLoggedIn = true;
        this.appService.isLoading = false;

        this.router.navigate(['/dashboard']);
      }
      else {
        this.publicService.tokenAssociation(session.accessToken.toString()).then((resp: any) => {
          this.code = `otpauth://totp/${userResp.UserAttributes[0].Value}?secret=${resp.SecretCode}&issuer=SagePA-${environment.name}`;
          this.showMfaSetupPopup = true;
          this.isReset = true;
        }).catch((e: { message: string | null | undefined; }) => {
          this.appService.isLoading = false;
          this.appService.openToaster("error", e.message);
        });
      }
    }).catch((error) => {
      this.appService.openToaster("error", error.message);
    });
  }

  /**
   * Navigation functioanlity
   * @param val - path
   */
  navigate(val: string) {
    this.router.navigate([val]);
  }

}
