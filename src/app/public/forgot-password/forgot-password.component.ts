import { Component } from '@angular/core';
import { AbstractControl, FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { ConfirmResetPasswordInput, ResetPasswordOutput } from 'aws-amplify/auth';
import { AppService } from '../../app.service';
import { PublicService } from '../public.service';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-forgot-password',
  standalone: true,
  imports: [CommonModule,FormsModule,ReactiveFormsModule],
  templateUrl: './forgot-password.component.html',
  styleUrl: './forgot-password.component.css'
})
export class ForgotPasswordComponent {
/**
   * Password Email form
   */
passwordEmailForm: FormGroup

/**
 * Password Code form
 */
passwordCodeForm: FormGroup

/**
 * To show passwordCodeForm
 */
isBoolean: boolean = false

/**
 * Variable indicating text to show or hide
 */
showButton !: boolean;

/**
 * Variable indicating text to show or hide
 */
showEye !: boolean;

/**
 * Variable indicating text to show or hide
 */
showCButton !: boolean;

/**
 * Variable indicating text to show or hide
 */
showCEye !: boolean

/**
 * Injecting dependencies
 * @param formbuilder - formbuilder for creating form
 * @param publicService - service for public module related functionality
 * @param appService - service for app level functionality
 * @param router - router service for navigation
 */
constructor(private formbuilder: FormBuilder, private publicService: PublicService, private appService: AppService, private router: Router) {

  /**
   * Password Email form
   */
  this.passwordEmailForm = this.formbuilder.group({
    email: ['', [Validators.required, Validators.email]]
  })

  /**
   * Password Code form
   */
  this.passwordCodeForm = this.formbuilder.group({
    confirmationCode: ['', [Validators.required, Validators.pattern('[0-9]{6}')]],
    newPassword: ['', [Validators.required, Validators.minLength(14), Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*[^a-zA-Z0-9 ]).{14,}$')]],
    confirmPassword: ['', [Validators.required]],
  }, { validators: this.confirmPasswordValidator })
}

/**
 * Method to handle resetPassword
 */
async handleResetPassword() {
  if (this.passwordEmailForm.value.email != null) {
    this.appService.isLoading = true;
    console.log(this.passwordEmailForm.value.email)
    this.publicService.resetPassword(this.passwordEmailForm.value.email)
      .then((resp: ResetPasswordOutput) => {
        console.log(resp);
        const { nextStep } = resp;
        console.log(nextStep.resetPasswordStep)
        switch (nextStep.resetPasswordStep) {
          case 'CONFIRM_RESET_PASSWORD_WITH_CODE':
            this.appService.isLoading = false;
            this.appService.openToaster("success", "Confirmation code sent to email")
            this.isBoolean = true;
            break;
        }
      }).catch((e) => {
        this.appService.isLoading = false;
        console.log(e)
        this.appService.openToaster("error", e.message)
      });
  }

}

/**
 * Method to confirm resetPassword using confirmation code
 */
async handleConfirmResetPassword() {
  this.appService.isLoading = true;
  const inputObj: ConfirmResetPasswordInput = {
    username: this.passwordEmailForm.value.email,
    confirmationCode: this.passwordCodeForm.value.confirmationCode,
    newPassword: this.passwordCodeForm.value.newPassword
  }
  this.publicService.confirmResetPassword(inputObj)
    .then(resp => {
      this.appService.isLoading = false;
      console.log(resp);
      if (resp == undefined) {
        this.appService.isLoading = false;
        this.appService.openToaster("success", "Password reset successfully");
        this.router.navigate(['login']);
      }
      else {
        this.appService.isLoading = false;
        this.appService.openToaster("error", resp['message'])
      }
      this.passwordCodeForm.reset();
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
 * Function to show/hide Password
 */
showConfirmPassword() {
  this.showCButton = !this.showCButton;
  this.showCEye = !this.showCEye
}

/**
 * Validator function to check if the passwords match
 * @param control - form control containing 'newPassword' and 'confirmPassword' fields
 * @returns validation error if passwords do not match, otherwise null
 */
confirmPasswordValidator(control: AbstractControl) {
  const newPassword = control.get('newPassword')?.value;
  const confirmPassword = control.get('confirmPassword')?.value;

  return newPassword === confirmPassword ? null : { PasswordNoMatch: true };
}

/**
 * Validate the field to show / hide the error message
 * @param field - check for validation
 * @returns boolean
 */
public isFieldValid(field: string) {
  return !(this.passwordCodeForm.get(field)?.valid === true) && (this.passwordCodeForm.get(field)?.dirty === true || this.passwordCodeForm.get(field)?.touched === true);
}
}
