# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestSubscriptionScopes**](ResourcesApi.md#RequestSubscriptionScopes) | **Get** /v4/resources/user/notifications/scopes | Request list of subscription scopes
[**RequestUserAvatar**](ResourcesApi.md#RequestUserAvatar) | **Get** /v4/resources/users/{user_id}/avatar/{uuid} | Request user avatar

# **RequestSubscriptionScopes**
> NotificationScopeList RequestSubscriptionScopes(ctx, )
Request list of subscription scopes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description: Retrieve a list of subscription scopes.  ### Precondition: Authenticated user.  ### Postcondition: List of scopes is returned.  ### Further Information: None.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NotificationScopeList**](NotificationScopeList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUserAvatar**
> Avatar RequestUserAvatar(ctx, uuid, userId)
Request user avatar

### Description: Get user avatar.  ### Precondition: Valid user ID and avatar UUID  ### Postcondition: Avatar is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| UUID of the avatar | 
  **userId** | **int64**| User ID | 

### Return type

[**Avatar**](Avatar.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

