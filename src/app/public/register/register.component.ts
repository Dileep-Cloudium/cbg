import { Component, EventEmitter, Output, ViewChild } from '@angular/core';
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
import { MaskedTextBoxModule, TextBoxModule } from '@syncfusion/ej2-angular-inputs';
import { DatePickerModule } from '@syncfusion/ej2-angular-calendars';
import { DatePipe } from '@angular/common';

/**
 * Component responsible for handling register related functionality
 */
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css','../public.component.css'],
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule,
    ButtonModule,
    TextBoxModule,
    TooltipModule,
    MaskedTextBoxModule,
    DatePickerModule
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
   * To keep Today's date as maximum for date of birth field
   */
  maxDate: Date = new Date();

  /**
   * Reference to tooltip component
   */
  @ViewChild('tooltip')
  public tooltip!: TooltipComponent;

  
  /**
   * Output event for action
   */
  @Output() action = new EventEmitter();
  
  /**
   * Injecting dependencies
   * @param router - router for navigation
   * @param formBuilder - Form Builder instance for signUp form
   * @param publicService - Service for handling public API
   * @param appService - Service for handling app related API
   */
  constructor(private router: Router, private formBuilder: FormBuilder, private publicService: PublicService, public appService: AppService, private datePipe: DatePipe) {
    this.isConfirm = false;
    this.showButton = false;
    this.showEye = false;
    this.signUpForm = this.formBuilder.group({
      firstName: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(100), Validators.pattern("[a-zA-Z ]*")]],
      lastName: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(100), Validators.pattern("[a-zA-Z ]*")]],
      memberId: ['', [Validators.required]],
      memberDOB: ['', [Validators.required]],
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
      const obj: UserRegistrationModel = {
        first_name: this.signUpForm.value.firstName,
        last_name: this.signUpForm.value.lastName,
        member_id: this.signUpForm.value.memberId,
        member_dob: this.datePipe.transform(this.signUpForm.value.memberDOB, 'MMddyyyy') ?? '',
        email: this.signUpForm.value.email
      }
      console.log('signup form',this.signUpForm.value);
      return
      this.publicService.register(obj).subscribe((data: ResponseModel) => {
        this.appService.isLoading = false;
        this.signUp(data)
      }, error => {
        this.appService.openToaster("error", "data" in error.error ? error.error.data[0] : "Failed to load response data");
        this.appService.isLoading = false;
      });
    }
  }

  /**
   * After successfull insertion in DB, register the user in AWS user pool with 
   * profile_id genereated from application database
   * @param data - ResponseModel
   */
  signUp(data: ResponseModel) {
    this.appService.isLoading = true;
    const obj: UserAWSRegistrationModel = this.signUpForm.value;
    // Currently iserting with default value
    obj.profileId = data.data?.id?.replace("user#", "");
    obj.type = "M";
    this.publicService.signUp(obj)
      .then((resp: SignUpOutput) => {
        this.appService.isLoading = false;
        this.updateUser(data, resp);
      }).catch((err) => {
        this.appService.isLoading = false;
        this.appService.openToaster("error", err.message)
      });
  }

  /**
   * After user creation in AWS user pool, get the AWS user id
   * and update in application DB
   * @param userData - Application DB record
   * @param awsUserData - AWS user record
   */
  updateUser(userData: ResponseModel, awsUserData: SignUpOutput) {
    this.appService.isLoading = true;
    const obj: UserAWSUpdateModel = {
      aws_cognito_user_id: awsUserData.userId,
      profile_id: userData.data?.id?.replace("user#", "")
    }
    this.publicService.updateAwsUserId(obj).subscribe((data: ResponseModel) => {
      this.appService.isLoading = false;
      this.isConfirm = true;
    }, error => {
      this.appService.openToaster("error", "data" in error.error ? error.error.detail[0] : "Failed to load response data");
    });
  }

  /**
   * To confirm the user email, need to submit the code which was received
   */
  confirmSignUp() {
    this.appService.isLoading = true;
    this.publicService.confirmSignUp(this.signUpForm.value.email, this.signUpForm.value.code)
      .then(() => {
        this.appService.isLoading = false
        this.appService.openToaster("success", "Registration successful. You may login now.")
        this.action.emit('login');
      }).catch((e) => {
        this.appService.isLoading = false;
        this.appService.openToaster("error", e.message)
      });
  }

  /**
   * Navigate to another page
   * @param val - path
   */
  navigate(val: string) {
    // this.tooltip.opensOn = "Auto"
    this.router.navigate([val]);
  }


  outputAction(val: any) {
    this.action.emit(val);
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
