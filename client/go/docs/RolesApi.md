# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoleGroups**](RolesApi.md#AddRoleGroups) | **Post** /v4/roles/{role_id}/groups | Assign group(s) to the role
[**AddRoleUsers**](RolesApi.md#AddRoleUsers) | **Post** /v4/roles/{role_id}/users | Assign user(s) to the role
[**RequestRoleGroups**](RolesApi.md#RequestRoleGroups) | **Get** /v4/roles/{role_id}/groups | Request groups with specific role
[**RequestRoleUsers**](RolesApi.md#RequestRoleUsers) | **Get** /v4/roles/{role_id}/users | Request users with specific role
[**RequestRoles**](RolesApi.md#RequestRoles) | **Get** /v4/roles | Request all roles with assigned rights
[**RevokeRoleGroups**](RolesApi.md#RevokeRoleGroups) | **Delete** /v4/roles/{role_id}/groups | Revoke granted role from group(s)
[**RevokeRoleUsers**](RolesApi.md#RevokeRoleUsers) | **Delete** /v4/roles/{role_id}/users | Revoke granted role from user(s)

# **AddRoleGroups**
> RoleGroupList AddRoleGroups(ctx, body, roleId, optional)
Assign group(s) to the role

### Description: Assign group(s) to a role.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; grant permission on desired role</span> required.  ### Postcondition: One or more groups will be added to a role.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupIds**](GroupIds.md)|  | 
  **roleId** | **int32**| Role ID | 
 **optional** | ***RolesApiAddRoleGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiAddRoleGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**RoleGroupList**](RoleGroupList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRoleUsers**
> RoleUserList AddRoleUsers(ctx, body, roleId, optional)
Assign user(s) to the role

### Description: Assign user(s) to a role.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; grant permission on desired role</span> required.  ### Postcondition: One or more users will be added to a role.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserIds**](UserIds.md)|  | 
  **roleId** | **int32**| Role ID | 
 **optional** | ***RolesApiAddRoleUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiAddRoleUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**RoleUserList**](RoleUserList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoleGroups**
> RoleGroupList RequestRoleGroups(ctx, roleId, optional)
Request groups with specific role

### Description:   Get all groups with a specific role.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read groups</span> required.  ### Postcondition: List of to the role assigned groups is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `isMember:eq:false|name:cn:searchString`   Get all groups that are **NOT** a member of that role **AND** whose name contains `searchString`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `isMember` | Filter the groups which are (not) member of that role | `eq` |  | <ul><li>`true`</li><li>`false`</li><li>`any`</li></ul>default: `true` | | `name` | Group name filter | `cn` | Group name contains value. | `search String` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **int32**| Role ID | 
 **optional** | ***RolesApiRequestRoleGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiRequestRoleGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**RoleGroupList**](RoleGroupList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoleUsers**
> RoleUserList RequestRoleUsers(ctx, roleId, optional)
Request users with specific role

### Description:   Get all users with a specific role.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read users</span> required.  ### Postcondition: List of users is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `isMember:eq:false|user:cn:searchString`   Get all users that are **NOT** member of that role **AND** whose (`firstName` **OR** `lastName` **OR** `email` **OR** `username`) is like `searchString`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `user` | User filter | `cn` | User contains value (`firstName` **OR** `lastName` **OR** `email` **OR** `username`). | `search String` | | `isMember` | Filter the users which are (not) member of that role | `eq` |  | <ul><li>`true`</li><li>`false`</li><li>`any`</li></ul>default: `true` |  </details>  ### Deprecated filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | <del>`displayName`</del> | User display name filter (use `user` filter) | `cn` | User display name contains value (`firstName` **OR** `lastName` **OR** `email`). | `search String` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **int32**| Role ID | 
 **optional** | ***RolesApiRequestRoleUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiRequestRoleUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**RoleUserList**](RoleUserList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoles**
> RoleList RequestRoles(ctx, optional)
Request all roles with assigned rights

### Description:   Retrieve a list of all roles with assigned rights.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read users</span> required.  ### Postcondition: List of roles with assigned rights is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RolesApiRequestRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiRequestRolesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**RoleList**](RoleList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeRoleGroups**
> RoleGroupList RevokeRoleGroups(ctx, body, roleId, optional)
Revoke granted role from group(s)

### Description:   Revoke granted group(s) from a role.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; grant permission on desired role</span> required.   For each role, at least one non-expiring user **MUST** remain who may grant the role.  ### Postcondition: One or more groups will be removed from a role.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupIds**](GroupIds.md)|  | 
  **roleId** | **int32**| Role ID | 
 **optional** | ***RolesApiRevokeRoleGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiRevokeRoleGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**RoleGroupList**](RoleGroupList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeRoleUsers**
> RoleUserList RevokeRoleUsers(ctx, body, roleId, optional)
Revoke granted role from user(s)

### Description:   Revoke granted user(s) from a role.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; grant permission on desired role</span> required.   For each role, at least one non-expiring user **MUST** remain who may grant the role.  ### Postcondition: One or more users will be removed from a role.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserIds**](UserIds.md)|  | 
  **roleId** | **int32**| Role ID | 
 **optional** | ***RolesApiRevokeRoleUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RolesApiRevokeRoleUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**RoleUserList**](RoleUserList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

