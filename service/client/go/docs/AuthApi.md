# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteOpenIdLogin**](AuthApi.md#CompleteOpenIdLogin) | **Post** /v4/auth/openid/login | Complete OpenID Connect authentication
[**InitiateOpenIdLogin**](AuthApi.md#InitiateOpenIdLogin) | **Get** /v4/auth/openid/login | Initiate OpenID Connect authentication
[**Login**](AuthApi.md#Login) | **Post** /v4/auth/login | Authenticate user (Login)
[**Ping**](AuthApi.md#Ping) | **Get** /v4/auth/ping | Ping
[**RecoverUserName**](AuthApi.md#RecoverUserName) | **Post** /v4/auth/recover_username | Recover username
[**RequestPasswordReset**](AuthApi.md#RequestPasswordReset) | **Post** /v4/auth/reset_password | Request password reset
[**ResetPassword**](AuthApi.md#ResetPassword) | **Put** /v4/auth/reset_password/{token} | Reset password
[**ValidateResetPasswordToken**](AuthApi.md#ValidateResetPasswordToken) | **Get** /v4/auth/reset_password/{token} | Validate information for password reset

# **CompleteOpenIdLogin**
> LoginResponse CompleteOpenIdLogin(ctx, code, state, optional)
Complete OpenID Connect authentication

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.14.0</h3>  ### Description:   This is the second step of the OpenID Connect authentication.   The user hands over the authorization code and is logged in.  ### Precondition: Existing user with activated OpenID Connect authentication that is **NOT** locked.  ### Postcondition: User is logged in.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| Authorization code | 
  **state** | **string**| Authentication state | 
 **optional** | ***AuthApiCompleteOpenIdLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthApiCompleteOpenIdLoginOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idToken** | **optional.String**| Identity token | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitiateOpenIdLogin**
> InitiateOpenIdLogin(ctx, issuer, redirectUri, language, test)
Initiate OpenID Connect authentication

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.14.0</h3>  ### Description: This is the first step of the OpenID Connect authentication.   The user is send to the OpenID Connect identity provider to authenticate himself and retrieve an authorization code.  ### Precondition: None.  ### Postcondition: User is redirected to OpenID Connect identity provider to authenticate himself.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issuer** | **string**| Issuer identifier of the OpenID Connect identity provider | 
  **redirectUri** | **string**| Redirect URI to complete the OpenID Connect authentication | 
  **language** | **string**| Language ID or ISO 639-1 code | 
  **test** | **bool**| Flag to test the authentication parameters.  If the request is valid, the API will respond with &#x60;204 No Content&#x60;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Login**
> LoginResponse Login(ctx, body)
Authenticate user (Login)

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.13.0</h3>  ### Description: Authenticates user and provides an authentication token (`X-Sds-Auth-Token`) that is required for the most operations.  ### Precondition: Existing user that is **NOT** locked.  ### Postcondition: User is logged in.  ### Further Information: The provided token is valid for **two hours**, every usage resets this period to two full hours again.   Logging off invalidates the token.    ### Available authentication methods: <details open style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Authentication Method (`authType`) | Description | | :--- | :--- | | `basic` | Log in with credentials stored in the database <br>Formerly known as `sql`.| | `active_directory` | Log in with Active Directory credentials | | `radius` | Log in with RADIUS username, PIN and token password.<br>Token (request parameter) may be set, otherwise this parameter is ignored. If token is set, password is optional. | | `openid` | Please use `POST /auth/openid/login` API to login with OpenID Connect identity |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoginRequest**](LoginRequest.md)|  | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ping**
> string Ping(ctx, )
Ping

### Description: Test connection to DRACOON Core Service.  ### Precondition: None.  ### Postcondition: `200 OK` with current date string is returned if successful.  ### Further Information: None.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecoverUserName**
> RecoverUserName(ctx, body)
Recover username

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.13.0</h3>  ### Description:   Request an email with the user names of all accounts connected to the email.  ### Precondition: Valid email address.  ### Postcondition: An email is sent to the provided address, with a list of account user names connected to it.  ### Further Information: None. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RecoverUserNameRequest**](RecoverUserNameRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestPasswordReset**
> RequestPasswordReset(ctx, body)
Request password reset

### Description:   Request an email with a password reset token for a certain user to reset password.  ### Precondition: Registered user account.  ### Postcondition: Provided user receives email with password reset token.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResetPasswordRequest**](ResetPasswordRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPassword**
> ResetPassword(ctx, body, token)
Reset password

### Description:   Resets user's password.  ### Precondition: User received a password reset token.  ### Postcondition: User's password is reset to the provided password.  ### Further Information: Forbidden characters in passwords: [`&`, `'`, `<`, `>`]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResetPasswordWithTokenRequest**](ResetPasswordWithTokenRequest.md)|  | 
  **token** | **string**| Password reset token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateResetPasswordToken**
> ResetPasswordTokenValidateResponse ValidateResetPasswordToken(ctx, token)
Validate information for password reset

### Description:   Request all information for a password change dialogue e.g. real name of user.  ### Precondition: User received a password reset token.  ### Postcondition: Context information is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Password reset token | 

### Return type

[**ResetPasswordTokenValidateResponse**](ResetPasswordTokenValidateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

