import { Component, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PublicService } from '../public.service';
import { ResponseModel } from '../../shared/response.model';
import { UserRegistrationModel, UserAWSRegistrationModel, UserAWSUpdateModel } from '../../shared/response.model';
import { SignUpOutput } from 'aws-amplify/auth';
import { TooltipComponent, TooltipModule } from '@syncfusion/ej2-angular-popups';
import { AppService } from '../../app.service';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { ButtonModule } from '@syncfusion/ej2-angular-buttons';
import { TextBoxModule } from '@syncfusion/ej2-angular-inputs';

/**
 * Component responsible for handling register related functionality
 */
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule,
    ButtonModule,
    TextBoxModule,
    TooltipModule
  ],
})

export class RegisterComponent {

  /**
   * Flag indicating whether the registration process is confirmed
   */
  isConfirm: boolean;

  /**
   * Formgroup for signUp form
   */
  signUpForm: FormGroup;

  /**
   * Flag to control visibility of button 
   */
  showButton: boolean;

  /**
   * Flag to control visibility of eye icon
   */
  showEye: boolean;

  /**
   * Reference to tooltip component
   */
  @ViewChild('tooltip')
  public tooltip!: TooltipComponent;

  /**
   * Injecting dependencies
   * @param router - router for navigation
   * @param formBuilder - Form Builder instance for signUp form
   * @param publicService - Service for handling public API
   * @param appService - Service for handling app related API
   */
  constructor(private router: Router, private formBuilder: FormBuilder, private publicService: PublicService, public appService: AppService) {
    this.isConfirm = false;
    this.showButton = false;
    this.showEye = false;
    this.signUpForm = this.formBuilder.group({
      firstName: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(100), Validators.pattern("[a-zA-Z ]*")]],
      lastName: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(100), Validators.pattern("[a-zA-Z ]*")]],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(14), Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*[^a-zA-Z0-9 ]).{14,}$')]],
      code: ['']
    })
  }

  /**
   * Lifecycle hook that executes when the component initializes, handling initializations, 
   * data retrieval
   */
  ngOnInit(): void {
    this.signUpForm.get('email')?.valueChanges.subscribe((value: string) => {
      if (value != "") {
        this.signUpForm.get('email')?.setValue(value.toLowerCase(), { emitEvent: false });
      }
    });
  }

  /**
   * Function to show/hide Password
   */
  showPassword() {
    this.showButton = !this.showButton;
    this.showEye = !this.showEye;
  }

  /**
   * Create a new record in application DB
   * @param obj - object with user details
   */
  createUser() {
    if (this.signUpForm.valid) {
      this.appService.isLoading = true;
      const obj: UserAWSRegistrationModel = this.signUpForm.value;
      // Currently iserting with default value
      obj.profileId = "b551df1a-6424-419e-8524-ae149858b8b0";
      this.publicService.signUp(obj)
        .then((resp: SignUpOutput) => {
          this.appService.isLoading = false;
          this.appService.openToaster("success", "Registration successful. You may login now.")
          this.navigate('login');
        }).catch((err) => {
          this.appService.isLoading = false;
          this.appService.openToaster("error", err.message)
        });
    }
  }

  /**
   * To confirm the user email, need to submit the code which was received
   */
  confirmSignUp() {
    // this.appService.isLoading = true;
    // this.publicService.confirmSignUp(this.signUpForm.value.email, this.signUpForm.value.code)
    //   .then(() => {
    //     this.appService.isLoading = false
    //     // this.appService.openToaster("success", "Registration successful. You may login now.")
    //     this.router.navigate(['login']);
    //   }).catch((e) => {
    //     this.appService.isLoading = false;
    //     // this.appService.openToaster("error", e.message)
    //   });
  }

  /**
   * Navigate to another page
   * @param val - path
   */
  navigate(val: string) {
    this.router.navigate([val]);
  }

  /**
   * Validate the fields of form
   * @param field - check for validation
   * @returns boolean
   */
  public isFieldValid(field: string) {
    return !(this.signUpForm.get(field)?.valid === true) && (this.signUpForm.get(field)?.dirty === true || this.signUpForm.get(field)?.touched === true);
  }
}
