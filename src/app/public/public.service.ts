import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { signUp, signIn, confirmSignUp, confirmSignIn, SignInOutput, SignUpOutput, ConfirmSignUpOutput, resetPassword, confirmResetPassword, ResetPasswordOutput, ConfirmResetPasswordInput, resendSignUpCode, ResendSignUpCodeOutput } from 'aws-amplify/auth';
import { HttpClient } from '@angular/common/http';
import { ResponseModel } from '../shared/response.model';
import { environment } from '../../environments/environment.develop';
import { CognitoIdentityProviderClient, AssociateSoftwareTokenCommand, VerifySoftwareTokenCommand, SetUserMFAPreferenceCommand } from '@aws-sdk/client-cognito-identity-provider';
import { UserRegistrationModel, UserAWSUpdateModel, UserAWSRegistrationModel } from '../shared/response.model';

@Injectable({
  providedIn: 'root'
})
export class PublicService {

  // private cognito: AWS.CognitoIdentityServiceProvider;
  private cognito: CognitoIdentityProviderClient;

  constructor(private http: HttpClient) {
    console.log('service initialised');
    
    // Initialize CognitoIdentityProviderClient with region
    this.cognito = new CognitoIdentityProviderClient({
      region: "us-west-2"
    });
  }

  
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

  /**
   * 
   * @param accessToken - User Access token which generates after login
   * @param totpCode - 6 digit code
   * @returns 
   */
  validateTOTP(accessToken: string, totpCode: string) {
    const params = {
      AccessToken: accessToken,
      UserCode: totpCode
    };
    const command = new VerifySoftwareTokenCommand(params);
    return this.cognito.send(command);
  }


  setUserMFA(accessToken: string) {
    const params = {
      AccessToken: accessToken,
      SoftwareTokenMfaSettings: {
        Enabled: true,
        PreferredMfa: true
      }
    };
    const command = new SetUserMFAPreferenceCommand(params);
    return this.cognito.send(command);
  }

  tokenAssociation(accessToken: string) {
    const params = {
      AccessToken: accessToken
    };
    const command = new AssociateSoftwareTokenCommand(params);
    return this.cognito.send(command);
  }

  /**
   * Create the user record in application database
   * @param data - User details with name and email
   */
  register(data: UserRegistrationModel) {
    const url = environment.baseURL + 'user/register';
    return this.http.post<ResponseModel>(url, data)
      .pipe(map((response: ResponseModel) => {
        return response;
      }))
  }

  /**
   * Update the user record in application database with AWS user ID
   * @param data - User details with id and AWS User ID
   */
  updateAwsUserId(data: UserAWSUpdateModel) {
    const url = environment.baseURL + 'user/cognito_id';
    return this.http.put<ResponseModel>(url, data)
      .pipe(map((response: ResponseModel) => {
        return response;
      }))
  }


  /**
   * To confirm the user email
   * @param user
   */
  async confirmSignUp(userName: string, code: string): Promise<ConfirmSignUpOutput> {
    return confirmSignUp({
      username: userName,
      confirmationCode: code
    });
  }

  /**
   * To resend the verification code
   * @param userName - email
   * @returns
   */
  async resendSignUpCode(userName: string): Promise<ResendSignUpCodeOutput> {
    return resendSignUpCode({
      username: userName
    })
  }

  /**
   * Confirms the user sign-in with multi-factor authentication
   * @param code 
   * @returns 
   */
  async confirmSignin(code: string) {
    return confirmSignIn({ challengeResponse: code }).then((resp) => {
      return resp;
    }).catch((err) => {
      return err;
    });
  }

  /**
   * Authenticates the user by validating the credentials in AWS
   * @param user
   */
  async resetPassword(username: string): Promise<ResetPasswordOutput> {
    return resetPassword({ username });
  }

  /**
   * Confirms the user to reset password
   * @param inputData ConfirmResetPasswordInput
   * @returns response
   */
  async confirmResetPassword(inputData: ConfirmResetPasswordInput): Promise<void> {
    return confirmResetPassword(inputData).then((resp) => {
      return resp;
    }).catch((err) => {
      return err;
    })
  }

}
