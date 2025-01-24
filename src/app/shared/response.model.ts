/**
 * Defines the structure of the response model for API interactions
 */
export interface ResponseModel {
    /**
     * Status of the API response
     */
    status: "success" | "fail" | "error",

    /**
     * Result data from API
     */
    data: any /* eslint-disable-line */
}

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

    /**
     * Custom attribute: To store the application generated user id
     */
    type: string
}

/**
 * Defines the structure to update AWS cognito user id in application databse
 */
export interface UserAWSUpdateModel {

    /**
     * user_id of user to be updated
     */
    profile_id: string,

    /**
     * AWS cognito user id generated in user pool
     */
    aws_cognito_user_id: string | undefined
}