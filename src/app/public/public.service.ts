import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { signUp, signIn, confirmSignUp, confirmSignIn, SignInOutput, SignUpOutput, ConfirmSignUpOutput, resetPassword, confirmResetPassword, ResetPasswordOutput, ConfirmResetPasswordInput, resendSignUpCode, ResendSignUpCodeOutput } from 'aws-amplify/auth';
import { HttpClient } from '@angular/common/http';
import { ResponseModel } from '../shared/response.model';
import { environment } from '../../environments/environment.develop';

/**
 * Defines the structure of user registration for application databse
 */
export interface UserRegistrationModel {

  /**
   * First name of user
   */
  first_name: string,

  /**
   * Last name of user
   */
  last_name: string,

  /**
   * email of user
   */
  email: string
}

/**
* Defines the structure of user registration for AWS
*/
export interface UserAWSRegistrationModel {

  /**
   * Standard attribute: given_name
   */
  firstName: string,

  /**
   * Standard attribute: family_name
   */
  lastName: string,

  /**
   * Standard attribute: email
   */
  email: string,

  /**
   * Standard & required attribute: username
   */
  userName: string,

  /**
   * Standard & required attribute: password
   */
  password: string,

  /**
   * Custom attribute: To store the application generated user id
   */
  profileId: string
}

/**
* Defines the structure to update AWS cognito user id in application databse
*/
export interface UserAWSUpdateModel {

  /**
   * user_id of user to be updated
   */
  id: number,

  /**
   * AWS cognito user id generated in user pool
   */
  aws_cognito_user_id: string | undefined
}

@Injectable({
  providedIn: 'root'
})
export class PublicService {

  
  /**
   * Authenticates the user by validating the credentials in AWS
   * @param user
   */
  async signIn(username: string, password: string): Promise<SignInOutput> {
    return signIn({ username, password });
  }

  /**
   * Registers a new user based on the provided information in AWS
   * @param user
   */
  async signUp(user: UserAWSRegistrationModel): Promise<SignUpOutput> {
    return signUp({
      username: user.email,
      password: user.password,
      options: {
        userAttributes: {
          "email": user.email,
          "given_name": user.firstName,
          "family_name": user.lastName,
          "custom:profile_id": user.profileId
        }
      }
    });
  }

}
