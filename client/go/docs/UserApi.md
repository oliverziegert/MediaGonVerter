# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeUserPassword**](UserApi.md#ChangeUserPassword) | **Put** /v4/user/account/password | Change user&#x27;s password
[**CreateAndPreserveUserKeyPair**](UserApi.md#CreateAndPreserveUserKeyPair) | **Post** /v4/user/account/keypairs | Create key pair and preserve copy of old private key
[**EnableCustomerEncryption**](UserApi.md#EnableCustomerEncryption) | **Put** /v4/user/account/customer | Activate client-side encryption for customer
[**ListDownloadShareSubscriptions**](UserApi.md#ListDownloadShareSubscriptions) | **Get** /v4/user/subscriptions/download_shares | List Download Share subscriptions
[**ListNodeSubscriptions**](UserApi.md#ListNodeSubscriptions) | **Get** /v4/user/subscriptions/nodes | List node subscriptions
[**ListUploadShareSubscriptions**](UserApi.md#ListUploadShareSubscriptions) | **Get** /v4/user/subscriptions/upload_shares | List Upload Share subscriptions
[**Logout**](UserApi.md#Logout) | **Post** /v4/user/logout | Invalidate authentication token
[**PingUser**](UserApi.md#PingUser) | **Get** /v4/user/ping | (authenticated) Ping
[**RemoveOAuthApproval**](UserApi.md#RemoveOAuthApproval) | **Delete** /v4/user/oauth/approvals/{client_id} | Remove OAuth client approval
[**RemoveOAuthAuthorization**](UserApi.md#RemoveOAuthAuthorization) | **Delete** /v4/user/oauth/authorizations/{client_id}/{authorization_id} | Remove a OAuth authorization
[**RemoveOAuthAuthorizations**](UserApi.md#RemoveOAuthAuthorizations) | **Delete** /v4/user/oauth/authorizations/{client_id} | Remove all OAuth authorizations of a client
[**RemoveProfileAttribute**](UserApi.md#RemoveProfileAttribute) | **Delete** /v4/user/profileAttributes/{key} | Remove user profile attribute
[**RemoveUserKeyPair**](UserApi.md#RemoveUserKeyPair) | **Delete** /v4/user/account/keypair | Remove user&#x27;s key pair
[**RequestAvatar**](UserApi.md#RequestAvatar) | **Get** /v4/user/account/avatar | Request avatar
[**RequestCustomerInfo**](UserApi.md#RequestCustomerInfo) | **Get** /v4/user/account/customer | Request customer information for user
[**RequestCustomerKeyPair**](UserApi.md#RequestCustomerKeyPair) | **Get** /v4/user/account/customer/keypair | Request customer&#x27;s key pair
[**RequestListOfNotificationConfigs**](UserApi.md#RequestListOfNotificationConfigs) | **Get** /v4/user/notifications/config | Request list of notification configurations
[**RequestOAuthApprovals**](UserApi.md#RequestOAuthApprovals) | **Get** /v4/user/oauth/approvals | Request list of OAuth client approvals
[**RequestOAuthAuthorizations**](UserApi.md#RequestOAuthAuthorizations) | **Get** /v4/user/oauth/authorizations | Request list of OAuth client authorizations
[**RequestProfileAttributes**](UserApi.md#RequestProfileAttributes) | **Get** /v4/user/profileAttributes | Request user profile attributes
[**RequestUserInfo**](UserApi.md#RequestUserInfo) | **Get** /v4/user/account | Request user account information
[**RequestUserKeyPair**](UserApi.md#RequestUserKeyPair) | **Get** /v4/user/account/keypair | Request user&#x27;s key pair
[**RequestUserKeyPairs**](UserApi.md#RequestUserKeyPairs) | **Get** /v4/user/account/keypairs | Request all user key pairs
[**ResetAvatar**](UserApi.md#ResetAvatar) | **Delete** /v4/user/account/avatar | Reset avatar
[**SetProfileAttributes**](UserApi.md#SetProfileAttributes) | **Post** /v4/user/profileAttributes | Set user profile attributes
[**SetUserKeyPair**](UserApi.md#SetUserKeyPair) | **Post** /v4/user/account/keypair | Set user&#x27;s key pair
[**SubscribeDownloadShare**](UserApi.md#SubscribeDownloadShare) | **Post** /v4/user/subscriptions/download_shares/{share_id} | Subscribe Download Share for notifications
[**SubscribeDownloadShares**](UserApi.md#SubscribeDownloadShares) | **Put** /v4/user/subscriptions/download_shares | Subscribe or Unsubscribe a List of Download Shares for notifications
[**SubscribeNode**](UserApi.md#SubscribeNode) | **Post** /v4/user/subscriptions/nodes/{node_id} | Subscribe node for notifications
[**SubscribeUploadShare**](UserApi.md#SubscribeUploadShare) | **Post** /v4/user/subscriptions/upload_shares/{share_id} | Subscribe Upload Share for notifications
[**SubscribeUploadShares**](UserApi.md#SubscribeUploadShares) | **Put** /v4/user/subscriptions/upload_shares | Subscribe or Unsubscribe a List of Upload Shares for notifications
[**UnsubscribeDownloadShare**](UserApi.md#UnsubscribeDownloadShare) | **Delete** /v4/user/subscriptions/download_shares/{share_id} | Unsubscribe Download Share from notifications
[**UnsubscribeNode**](UserApi.md#UnsubscribeNode) | **Delete** /v4/user/subscriptions/nodes/{node_id} | Unsubscribe node from notifications
[**UnsubscribeUploadShare**](UserApi.md#UnsubscribeUploadShare) | **Delete** /v4/user/subscriptions/upload_shares/{share_id} | Unsubscribe Upload Share from notifications
[**UpdateNodeSubscriptions**](UserApi.md#UpdateNodeSubscriptions) | **Put** /v4/user/subscriptions/nodes | Subscribe or Unsubscribe a List of nodes for notifications
[**UpdateNotificationConfig**](UserApi.md#UpdateNotificationConfig) | **Put** /v4/user/notifications/config/{id} | Update notification configuration
[**UpdateProfileAttributes**](UserApi.md#UpdateProfileAttributes) | **Put** /v4/user/profileAttributes | Add or edit user profile attributes
[**UpdateUserAccount**](UserApi.md#UpdateUserAccount) | **Put** /v4/user/account | Update user account
[**UploadAvatarAsMultipart**](UserApi.md#UploadAvatarAsMultipart) | **Post** /v4/user/account/avatar | Change avatar

# **ChangeUserPassword**
> ChangeUserPassword(ctx, body, optional)
Change user's password

### Description: Change the user's password.  ### Precondition: Authenticated user.  ### Postcondition: User's password is changed.  ### Further Information: The password **MUST** comply to configured password policies.    Forbidden characters in passwords: [`&`, `'`, `<`, `>`]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeUserPasswordRequest**](ChangeUserPasswordRequest.md)|  | 
 **optional** | ***UserApiChangeUserPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiChangeUserPasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAndPreserveUserKeyPair**
> CreateAndPreserveUserKeyPair(ctx, body, optional)
Create key pair and preserve copy of old private key

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Create user key pair and preserve copy of old private key.  ### Precondition: Authenticated user.  ### Postcondition: Key pair is created.   Copy of old private key is preserved.  ### Further Information: You can submit your old private key, encrypted with your current password.   This allows migrating file keys encrypted with your old key pair to the new one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateKeyPairRequest**](CreateKeyPairRequest.md)|  | 
 **optional** | ***UserApiCreateAndPreserveUserKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateAndPreserveUserKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableCustomerEncryption**
> CustomerData EnableCustomerEncryption(ctx, body, optional)
Activate client-side encryption for customer

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.24.0</h3>  ### Use `POST /settings/keypair` API  ### Description:   Activate client-side encryption for according customer.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change config</span> required.  ### Postcondition: Client-side encryption is enabled.  ### Further Information: Sets the ability for this customer to encrypt rooms.   Once enabled on customer level, it **CANNOT** be unset.   On activation, a customer rescue key pair **MUST** be set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnableCustomerEncryptionRequest**](EnableCustomerEncryptionRequest.md)|  | 
 **optional** | ***UserApiEnableCustomerEncryptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiEnableCustomerEncryptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**CustomerData**](CustomerData.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDownloadShareSubscriptions**
> SubscribedDownloadShareList ListDownloadShareSubscriptions(ctx, optional)
List Download Share subscriptions

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description:   Retrieve a list of subscribed Download Shares for current user.   ### Precondition: Authenticated user.  ### Postcondition: List of subscribed Download Shares is returned.  ### Further Information: None.  ### Filtering All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `authParentId:eq:#`   Get download shares where `authParentId` equals `#`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | **`downloadShareId`** | Download Share ID filter | `eq` | Download Share ID equals value. | `long value` | | **`authParentId`** | Auth parent ID filter | `eq` | Auth parent ID equals value. | `long value` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `downloadShareId:desc|authParentId:asc`   Sort by `downloadShareId` descending **AND** `authParentId` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | **`downloadShareId`** | Download Share ID | | **`authParentId`** | Auth parent ID |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiListDownloadShareSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListDownloadShareSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter string | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **offset** | **optional.Int32**| Range offset | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**SubscribedDownloadShareList**](SubscribedDownloadShareList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeSubscriptions**
> SubscribedNodeList ListNodeSubscriptions(ctx, optional)
List node subscriptions

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description:   Retrieve a list of subscribed nodes for current user.   ### Precondition: Authenticated user.  ### Postcondition: List of subscribed nodes is returned.  ### Further Information: None.  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `authParentId:eq:#`   Get nodes where `authParentId` equals `#`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | **`nodeId`** | Node ID filter | `eq` | Node ID equals value. | `long value` | | **`authParentId`** | Auth parent ID filter | `eq` | Auth parent ID equals value. | `long value` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `nodeId:desc|authParentId:asc`   Sort by `nodeId` descending **AND** `authParentId` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | **`nodeId`** | Node ID | | **`authParentId`** | Auth parent ID |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiListNodeSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListNodeSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter string | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **offset** | **optional.Int32**| Range offset | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**SubscribedNodeList**](SubscribedNodeList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUploadShareSubscriptions**
> SubscribedUploadShareList ListUploadShareSubscriptions(ctx, optional)
List Upload Share subscriptions

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Retrieve a list of subscribed Upload Shares for current user.   ### Precondition: Authenticated user.  ### Postcondition: List of subscribed Upload Shares is returned.  ### Further Information: None.  ### Filtering All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `targetNodeId:eq:#`   Get upload shares where `targetNodeId` equals `#`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | **`uploadShareId`** | Upload Share ID filter | `eq` | Upload Share ID equals value. | `long value` | | **`targetNodeId`** | Target node ID filter | `eq` | Target node ID equals value. | `long value` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `uploadShareId:desc|targetNodeId:asc`   Sort by `uploadShareId` descending **AND** `targetNodeId` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | **`uploadShareId`** | Upload Share ID | | **`targetNodeId`** | Target node ID |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiListUploadShareSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListUploadShareSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter string | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **offset** | **optional.Int32**| Range offset | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**SubscribedUploadShareList**](SubscribedUploadShareList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Logout**
> Logout(ctx, optional)
Invalidate authentication token

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.12.0</h3>  ### Description:   Log out a user.  ### Precondition: Authenticated user.  ### Postcondition: * User is logged out   * Authentication token gets invalidated.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiLogoutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiLogoutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **everywhere** | **optional.Bool**| Invalidate all tokens | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingUser**
> string PingUser(ctx, optional)
(authenticated) Ping

### Description: Test connection to DRACOON Server (while authenticated).  ### Precondition: Authenticated user.  ### Postcondition: `200 OK` with principal information is returned if successful.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiPingUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiPingUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOAuthApproval**
> RemoveOAuthApproval(ctx, clientId, optional)
Remove OAuth client approval

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.22.0</h3>  ### Functional Description: Delete an OAuth client approval.  ### Precondition: Authenticated user and valid client ID  ### Postcondition: OAuth Client approval is revoked.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| OAuth client ID | 
 **optional** | ***UserApiRemoveOAuthApprovalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRemoveOAuthApprovalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOAuthAuthorization**
> RemoveOAuthAuthorization(ctx, clientId, authorizationId, optional)
Remove a OAuth authorization

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.12.0</h3>  ### Description: Delete an authorization.  ### Precondition: Authenticated user and valid client ID, authorization ID  ### Postcondition: Authorization is revoked.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| OAuth client ID | 
  **authorizationId** | **int64**| OAuth authorization ID | 
 **optional** | ***UserApiRemoveOAuthAuthorizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRemoveOAuthAuthorizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOAuthAuthorizations**
> RemoveOAuthAuthorizations(ctx, clientId, optional)
Remove all OAuth authorizations of a client

### Description: Delete all authorizations of a client.  ### Precondition: Authenticated user and valid client ID  ### Postcondition: All authorizations for the client are revoked.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| OAuth client ID | 
 **optional** | ***UserApiRemoveOAuthAuthorizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRemoveOAuthAuthorizationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveProfileAttribute**
> RemoveProfileAttribute(ctx, key, optional)
Remove user profile attribute

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.7.0</h3>  ### Description:   Delete custom user profile attribute.  ### Precondition: None.  ### Postcondition: Custom user profile attribute is deleted.  ### Further Information: Allowed characters for keys are: `[a-zA-Z0-9_-]`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Key | 
 **optional** | ***UserApiRemoveProfileAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRemoveProfileAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserKeyPair**
> RemoveUserKeyPair(ctx, optional)
Remove user's key pair

### Description:   Delete user key pair.  ### Precondition: Authenticated user.  ### Postcondition: Key pair is deleted.  ### Further Information: If parameter `version` is not set and two key versions exist, this API deletes version A.       If two keys with the same version are set, this API deletes the older one.  This will also remove all file keys that were encrypted with the user public key. If the user had exclusive access to some files, those are removed as well since decrypting them became impossible.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRemoveUserKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRemoveUserKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **optional.String**| Version (NEW) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestAvatar**
> Avatar RequestAvatar(ctx, optional)
Request avatar

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.11.0</h3>  ### Description: Get the avatar.  ### Precondition: Authenticated user.  ### Postcondition: Avatar is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRequestAvatarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRequestAvatarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**Avatar**](Avatar.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestCustomerInfo**
> CustomerData RequestCustomerInfo(ctx, optional)
Request customer information for user

### Description:   Use this API to get:  * customer name * used / free space * used / available * user account info  of the according customer.  ### Precondition: Authenticated user.  ### Postcondition: Customer information is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRequestCustomerInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRequestCustomerInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**CustomerData**](CustomerData.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestCustomerKeyPair**
> UserKeyPairContainer RequestCustomerKeyPair(ctx, optional)
Request customer's key pair

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.24.0</h3>  ### Use `GET /settings/keypair` API  ### Description:   Retrieve the customer rescue key pair.  ### Precondition: Authenticated user.  ### Postcondition: Key pair is returned.  ### Further Information: The private key is password-based encrypted with `AES256` / `PBKDF2`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRequestCustomerKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRequestCustomerKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UserKeyPairContainer**](UserKeyPairContainer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestListOfNotificationConfigs**
> NotificationConfigList RequestListOfNotificationConfigs(ctx, optional)
Request list of notification configurations

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description:   Retrieve a list of notification configurations for current user.   ### Precondition: Authenticated user.  ### Postcondition: List of available notification configurations is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRequestListOfNotificationConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRequestListOfNotificationConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**NotificationConfigList**](NotificationConfigList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOAuthApprovals**
> []OAuthApproval RequestOAuthApprovals(ctx, optional)
Request list of OAuth client approvals

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.22.0</h3>  ### Functional Description:   Retrieve information about all OAuth client approvals.  ### Precondition: Authenticated user.  ### Postcondition: None.  ### Further Information: None.  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `clientName:desc`   Sort by `clientName` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `clientName` | Client name |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRequestOAuthApprovalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRequestOAuthApprovalsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**[]OAuthApproval**](OAuthApproval.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOAuthAuthorizations**
> []OAuthAuthorization RequestOAuthAuthorizations(ctx, optional)
Request list of OAuth client authorizations

### Description:   Retrieve information about all OAuth client authorizations.  ### Precondition: Authenticated user.  ### Postcondition: List of OAuth client authorizations is returned.  ### Further Information:  ### Filtering: Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `isStandard:eq:true`   Get standard OAuth clients.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `isStandard` | Standard client filter | `eq` |  | `true or false` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `clientName:desc`   Sort by `clientName` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `clientName` | Client name |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRequestOAuthAuthorizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRequestOAuthAuthorizationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**[]OAuthAuthorization**](OAuthAuthorization.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestProfileAttributes**
> AttributesResponse RequestProfileAttributes(ctx, optional)
Request user profile attributes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.7.0</h3>  ### Description:   Retrieve a list of user profile attributes.  ### Precondition: None.  ### Postcondition: List of attributes is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `key:cn:searchString_1|value:cn:searchString_2`   Filter by attribute key contains `searchString_1` **AND** attribute value contains `searchString_2`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `key` | User profile attribute key filter | `cn, eq, sw` | Attribute key contains / equals / starts with value. | `search String` | | `value` | User profile attribute value filter | `cn, eq, sw` | Attribute value contains / equals / starts with value. | `search String` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `key:asc|value:desc`   Sort by `key` ascending **AND** by `value` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `key` | User profile attribute key | | `value` | User profile attribute value |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRequestProfileAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRequestProfileAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**AttributesResponse**](AttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUserInfo**
> UserAccount RequestUserInfo(ctx, optional)
Request user account information

### Description:   Retrieves all information regarding the current user's account.  ### Precondition: Authenticated user.  ### Postcondition: User information is returned.  ### Further Information: Setting the query parameter `more_info` to `true`, causes the API to return more details e.g. the user's groups.    `customer` (`CustomerData`) attribute in `UserAccount` response model is deprecated. Please use response from `GET /user/account/customer` instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRequestUserInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRequestUserInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **moreInfo** | **optional.Bool**| Get more info for this user  e.g. list of user groups | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UserAccount**](UserAccount.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUserKeyPair**
> UserKeyPairContainer RequestUserKeyPair(ctx, optional)
Request user's key pair

### Description:   Retrieve the user key pair.  ### Precondition: Authenticated user.  ### Postcondition: Key pair is returned.   ### Further Information: The private key is password-based encrypted with `AES256` / `PBKDF2`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRequestUserKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRequestUserKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **version** | **optional.String**| Version (NEW) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UserKeyPairContainer**](UserKeyPairContainer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUserKeyPairs**
> []UserKeyPairContainer RequestUserKeyPairs(ctx, optional)
Request all user key pairs

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Retrieve all user key pairs to allow re-encrypting file keys without need for a second distributor.  ### Precondition: Authenticated user.  ### Postcondition: List of key pairs is returned.   ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiRequestUserKeyPairsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRequestUserKeyPairsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**[]UserKeyPairContainer**](UserKeyPairContainer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetAvatar**
> Avatar ResetAvatar(ctx, optional)
Reset avatar

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.11.0</h3>  ### Description:   Reset (custom) avatar to default avatar.  ### Precondition: Authenticated user.  ### Postcondition: * User's avatar gets deleted.   * Default avatar is set.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiResetAvatarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiResetAvatarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**Avatar**](Avatar.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetProfileAttributes**
> ProfileAttributes SetProfileAttributes(ctx, body, optional)
Set user profile attributes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.12.0</h3>  ### Description:   Set custom user profile attributes.  ### Precondition: None.  ### Postcondition: Custom user profile attributes are set.  ### Further Information: Batch function.   All existing user profile attributes will be deleted.    * Allowed characters for keys are: `[a-zA-Z0-9_-]`   * Characters are **case-insensitive**   * Maximum key length is **255**   * Maximum value length is **4096**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProfileAttributesRequest**](ProfileAttributesRequest.md)|  | 
 **optional** | ***UserApiSetProfileAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiSetProfileAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**ProfileAttributes**](ProfileAttributes.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetUserKeyPair**
> SetUserKeyPair(ctx, body, optional)
Set user's key pair

### Description:   Set the user key pair.  ### Precondition: Authenticated user.  ### Postcondition: Key pair is set.  ### Further Information: Overwriting an existing key pair is **NOT** possible.   Please delete the existing key pair first.   The private key is password-based encrypted with `AES256` / `PBKDF2`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserKeyPairContainer**](UserKeyPairContainer.md)|  | 
 **optional** | ***UserApiSetUserKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiSetUserKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeDownloadShare**
> SubscribedDownloadShare SubscribeDownloadShare(ctx, shareId, optional)
Subscribe Download Share for notifications

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description:   Subscribe Download Share for notifications.  ### Precondition: User with _\"manage download share\"_ permissions on target node.  ### Postcondition: Download Share is subscribed.   Notifications for this Download Share will be triggered in the future.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shareId** | **int64**| Share ID | 
 **optional** | ***UserApiSubscribeDownloadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiSubscribeDownloadShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**SubscribedDownloadShare**](SubscribedDownloadShare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeDownloadShares**
> SubscribeDownloadShares(ctx, body, optional)
Subscribe or Unsubscribe a List of Download Shares for notifications

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.25.0</h3>  ### Description:   Subscribe/Unsubscribe download shares for notifications.  ### Precondition: User with _\"manage download share\"_ permissions on target node.    ### Postcondition: Download shares are subscribed or unsubscribed. Notifications for these download shares will be triggered in the future.  ### Further Information: Maximum number of subscriptions is 200.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSubscriptionsBulkRequest**](UpdateSubscriptionsBulkRequest.md)|  | 
 **optional** | ***UserApiSubscribeDownloadSharesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiSubscribeDownloadSharesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeNode**
> SubscribedNode SubscribeNode(ctx, nodeId, optional)
Subscribe node for notifications

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description: Subscribe node for notifications.  ### Precondition: User has _\"read\"_ permissions in auth parent room.  ### Postcondition: Node is subscribed. Notifications for this node will be triggered in the future.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Node ID | 
 **optional** | ***UserApiSubscribeNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiSubscribeNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**SubscribedNode**](SubscribedNode.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeUploadShare**
> SubscribedUploadShare SubscribeUploadShare(ctx, shareId, optional)
Subscribe Upload Share for notifications

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Subscribe Upload Share for notifications.  ### Precondition: User with _\"manage upload share\"_ permissions on target node.  ### Postcondition: Upload Share is subscribed.   Notifications for this Upload Share will be triggered in the future.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shareId** | **int64**| Share ID | 
 **optional** | ***UserApiSubscribeUploadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiSubscribeUploadShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**SubscribedUploadShare**](SubscribedUploadShare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeUploadShares**
> SubscribeUploadShares(ctx, body, optional)
Subscribe or Unsubscribe a List of Upload Shares for notifications

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.25.0</h3>  ### Description:   Subscribe/Unsubscribe upload shares for notifications.  ### Precondition: User with _\"manage upload share\"_ permissions on target node.    ### Postcondition: Upload shares are subscribed or unsubscribed. Notifications for these upload shares will be triggered in the future.  ### Further Information: Maximum number of subscriptions is 200.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSubscriptionsBulkRequest**](UpdateSubscriptionsBulkRequest.md)|  | 
 **optional** | ***UserApiSubscribeUploadSharesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiSubscribeUploadSharesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribeDownloadShare**
> UnsubscribeDownloadShare(ctx, shareId, optional)
Unsubscribe Download Share from notifications

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description:   Unsubscribe Download Share from notifications.  ### Precondition: User with _\"manage download share\"_ permissions on target node.  ### Postcondition: Download Share is unsubscribed.   Notifications for this Download Share are disabled.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shareId** | **int64**| Share ID | 
 **optional** | ***UserApiUnsubscribeDownloadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUnsubscribeDownloadShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribeNode**
> UnsubscribeNode(ctx, nodeId, optional)
Unsubscribe node from notifications

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description:   Unsubscribe node from notifications.  ### Precondition: User has _\"read\"_ permissions in auth parent room.  ### Postcondition: Node is unsubscribed.   Notifications for this node are disabled.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Node ID | 
 **optional** | ***UserApiUnsubscribeNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUnsubscribeNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribeUploadShare**
> UnsubscribeUploadShare(ctx, shareId, optional)
Unsubscribe Upload Share from notifications

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Unsubscribe Upload Share from notifications.  ### Precondition: User with _\"manage upload share\"_ permissions on target node.  ### Postcondition: Upload Share is unsubscribed.   Notifications for this Upload Share are disabled.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shareId** | **int64**| Share ID | 
 **optional** | ***UserApiUnsubscribeUploadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUnsubscribeUploadShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeSubscriptions**
> UpdateNodeSubscriptions(ctx, body, optional)
Subscribe or Unsubscribe a List of nodes for notifications

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.25.0</h3>  ### Description:   Subscribe/Unsubscribe nodes for notifications.  ### Precondition: User has _\"read\"_ permissions in auth parent room.  ### Postcondition: Nodes are subscribed or unsubscribed. Notifications for these nodes will be triggered in the future.  ### Further Information: Maximum number of subscriptions is 200.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSubscriptionsBulkRequest**](UpdateSubscriptionsBulkRequest.md)|  | 
 **optional** | ***UserApiUpdateNodeSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateNodeSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNotificationConfig**
> NotificationConfig UpdateNotificationConfig(ctx, body, id, optional)
Update notification configuration

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description:   Update notification configuration for current user.   ### Precondition: Authenticated user.  ### Postcondition: Notification configuration is updated.  ### Further Information: Leave `channelIds` empty to disable notifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationConfigChangeRequest**](NotificationConfigChangeRequest.md)|  | 
  **id** | **int64**| Unique identifier for a notification configuration | 
 **optional** | ***UserApiUpdateNotificationConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateNotificationConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**NotificationConfig**](NotificationConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProfileAttributes**
> ProfileAttributes UpdateProfileAttributes(ctx, body, optional)
Add or edit user profile attributes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.7.0</h3>  ### Description:   Add or edit custom user profile attributes. <br/><br/><span style=\"font-weight: bold; color: red;\"> &#128679; **Warning: Please note that the response with HTTP status code 200 (OK) is deprecated and will be replaced with HTTP status code 204 (No content)!**</span><br/>  ### Precondition: None.  ### Postcondition: Custom user profile attributes are added or edited.  ### Further Information: Batch function.   If an entry existed before, it will be overwritten.   Range submodel is never returned.  * Allowed characters for keys are: `[a-zA-Z0-9_-]`   * Characters are **case-insensitive**   * Maximum key length is **255**   * Maximum value length is **4096**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProfileAttributesRequest**](ProfileAttributesRequest.md)|  | 
 **optional** | ***UserApiUpdateProfileAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateProfileAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**ProfileAttributes**](ProfileAttributes.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserAccount**
> UserAccount UpdateUserAccount(ctx, body, optional)
Update user account

### Description:   Update current user's account.  ### Precondition: Authenticated user.  ### Postcondition: User's account is updated.  ### Further Information: * All input fields are limited to **150** characters.   * **All** characters are allowed.    `customer` (`CustomerData`) attribute in `UserAccount` response model is deprecated. Please use response from `GET /user/account/customer` instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserAccountRequest**](UpdateUserAccountRequest.md)|  | 
 **optional** | ***UserApiUpdateUserAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateUserAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**UserAccount**](UserAccount.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadAvatarAsMultipart**
> Avatar UploadAvatarAsMultipart(ctx, optional)
Change avatar

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.11.0</h3>  ### Description: Change the avatar.  ### Precondition: Authenticated user.  ### Postcondition: Avatar is changed.  ### Further Information: * Media type **MUST** be `jpeg` or `png` * File size **MUST** bei less than `5 MB` * Dimensions **MUST** be `256x256 px`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiUploadAvatarAsMultipartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUploadAvatarAsMultipartOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Avatar**](Avatar.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

