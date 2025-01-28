import { EventEmitter, Injectable } from '@angular/core';
import { AuthTokens, AuthUser, fetchAuthSession } from 'aws-amplify/auth';
import { ToasterModel, ToasterType, ToasterPositionX, ToasterPositionY } from './shared/toaster/toaster.model';
import { environment } from '../environments/environment.develop';
import { getCurrentUser } from 'aws-amplify/auth';
import { Amplify } from 'aws-amplify';
import { CognitoIdentityProviderClient, GetUserCommand } from '@aws-sdk/client-cognito-identity-provider';
import { signOut } from 'aws-amplify/auth';

@Injectable({
  providedIn: 'root'
})
export class AppService {
  
  /**
   * AWS CognitoIdentityServiceProvider instance for managing user-related actions
   */
  private cognito!: CognitoIdentityProviderClient;

  /**
   * To show / hide loader icon
   */
  isLoading = false;

  
  /**
   * Handles toaster object
   */
  toasterObject!: ToasterModel

  
  /**
   * Output event emitter for toaster
   */
  openToasterEmit = new EventEmitter();
  
  /**
   * User authentication status
   */
  isLoggedIn = false;

  constructor() {
    Amplify.configure({
      Auth: {
        Cognito: environment.cognito
      }
    });
    this.cognito = new CognitoIdentityProviderClient({
      region: "us-west-2"
    });
   }
   
  /**
   * Opens the toaster and display the message
   * @param type - Type of toaster. Ex: "success" | "warning" | "info" | "error"
   * @param message - Message to display in the toaster
   * @param title - Title to display in the toaster
   * @param positionX - Horizontal position of toaster
   * @param positionY - Vertical position of toaster
   */
  openToaster(type: ToasterType, message: string | null = null, title?: string, positionX: ToasterPositionX | null = null, positionY: ToasterPositionY | null = null) {
    this.toasterObject = {
      type: type,
      message: message != null ? message : "No message",
      title: title,
      position: {
        X: positionX != null ? positionX : "Right", // Defulat, RIGHT
        Y: positionY != null ? positionY : "Top" // Default, TOP
      }
    }
    this.openToasterEmit.emit();
  }


   
  /**
   * Returns access token and id token if refresh token exist and valid
   * @returns AuthTokens
   */
  async getAuthTokens(): Promise<AuthTokens> {
    return (await fetchAuthSession()).tokens ?? {} as AuthTokens;
  }

  
  /**
   * Get the logged in user detail if token is exist and valid
   * @returns AuthUser
   */
  async getCurrentUser(): Promise<AuthUser> {
    return await getCurrentUser();
  }

  
  /**
   * Method to get user details from user pool
   * @param token - access_token
   * @returns user details
   */
  async getUser(token: string) {
    const params = {
      "AccessToken": token
    };
    return await this.cognito.send(new GetUserCommand(params));
  }
  
    /**
   * Handles the sign-out process, securely terminating the user's session and
   * ensuring a clean logout from the application
   * @returns 
   */
      async handleSignOut() {
        return signOut().then((resp) => {
          return resp;
        }).catch((err) => {
          return err;
        });
      }
}
