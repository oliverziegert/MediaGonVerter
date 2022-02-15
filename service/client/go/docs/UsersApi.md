# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /v4/users | Create new user
[**RemoveUser**](UsersApi.md#RemoveUser) | **Delete** /v4/users/{user_id} | Remove user
[**RemoveUserAttribute**](UsersApi.md#RemoveUserAttribute) | **Delete** /v4/users/{user_id}/userAttributes/{key} | Remove custom user attribute
[**RequestLastAdminRoomsUsers**](UsersApi.md#RequestLastAdminRoomsUsers) | **Get** /v4/users/{user_id}/last_admin_rooms | Request rooms where the user is last admin
[**RequestUser**](UsersApi.md#RequestUser) | **Get** /v4/users/{user_id} | Request user
[**RequestUserAttributes**](UsersApi.md#RequestUserAttributes) | **Get** /v4/users/{user_id}/userAttributes | Request custom user attributes
[**RequestUserGroups**](UsersApi.md#RequestUserGroups) | **Get** /v4/users/{user_id}/groups | Request groups that user is a member of or / and can become a member
[**RequestUserRoles**](UsersApi.md#RequestUserRoles) | **Get** /v4/users/{user_id}/roles | Request user&#x27;s granted roles
[**RequestUsers**](UsersApi.md#RequestUsers) | **Get** /v4/users | Request users
[**RequestUsersRooms**](UsersApi.md#RequestUsersRooms) | **Get** /v4/users/{user_id}/rooms | Request rooms granted to the user or / and rooms that can be granted
[**SetUserAttributes**](UsersApi.md#SetUserAttributes) | **Post** /v4/users/{user_id}/userAttributes | Set custom user attributes
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /v4/users/{user_id} | Update user&#x27;s metadata
[**UpdateUserAttributes**](UsersApi.md#UpdateUserAttributes) | **Put** /v4/users/{user_id}/userAttributes | Add or edit custom user attributes

# **CreateUser**
> UserData CreateUser(ctx, body, optional)
Create new user

### Description: Create a new user.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change users</span> required.  ### Postcondition: New user is created.  ### Further Information: * If a user should **NOT** expire, leave `expireAt` empty. * All input fields are limited to **150** characters * Forbidden characters in passwords: [`&`, `'`, `<`, `>`]  ### Authentication Method Options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Authentication Method | Option Key | Option Value | | :--- | :--- | :--- | | `basic` / `sql` | `username` | Unique user identifier | | `active_directory` | `ad_config_id` (optional) | Active Directory configuration ID | |  | `username` | Active Directory username according to authentication setting `userFilter` | | `radius` | `username` | RADIUS username | | `openid` | `openid_config_id` (optional) | OpenID Connect configuration ID | |  | `username` | OpenID Connect username according to authentication setting `mappingClaim` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserRequest**](CreateUserRequest.md)|  | 
 **optional** | ***UsersApiCreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiCreateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**UserData**](UserData.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUser**
> RemoveUser(ctx, userId, optional)
Remove user

### Description: Delete a user.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; delete users</span> required.  ### Postcondition: User is deleted.  ### Further Information: User **CANNOT** be deleted if he is a last room administrator of any room.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| User ID | 
 **optional** | ***UsersApiRemoveUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiRemoveUserOpts struct
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

# **RemoveUserAttribute**
> RemoveUserAttribute(ctx, userId, key, optional)
Remove custom user attribute

### Description: Delete custom user attribute.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change users</span> required.  ### Postcondition: Custom user attribute is deleted.  ### Further Information: * Allowed characters for keys are: `[a-zA-Z0-9_-]`   * Characters are **case-insensitive**.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| User ID | 
  **key** | **string**| Key | 
 **optional** | ***UsersApiRemoveUserAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiRemoveUserAttributeOpts struct
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

# **RequestLastAdminRoomsUsers**
> LastAdminUserRoomList RequestLastAdminRoomsUsers(ctx, userId, optional)
Request rooms where the user is last admin

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.10.0</h3>  ### Description:   Retrieve a list of all rooms where the user is last admin (except homeroom and its subordinary rooms).  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change users</span> required.  ### Postcondition: List of rooms is returned.   ### Further Information: An empty list is returned if no rooms were found where the user is last admin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| User ID | 
 **optional** | ***UsersApiRequestLastAdminRoomsUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiRequestLastAdminRoomsUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**LastAdminUserRoomList**](LastAdminUserRoomList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUser**
> UserData RequestUser(ctx, userId, optional)
Request user

### Description:   Retrieve detailed information about a single user.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read users</span> required.  ### Postcondition: User information is returned.  ### Further Information: None.  ### Authentication Method Options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Authentication Method | Option Key | Option Value | | :--- | :--- | :--- | | `basic` / `sql` | `username` | Unique user identifier | | `active_directory` | `ad_config_id` (optional) | Active Directory configuration ID | |  | `username` | Active Directory username according to authentication setting `userFilter` | | `radius` | `username` | RADIUS username | | `openid` | `openid_config_id` (optional) | OpenID Connect configuration ID | |  | `username` | OpenID Connect username according to authentication setting `mappingClaim` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| User ID | 
 **optional** | ***UsersApiRequestUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiRequestUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **effectiveRoles** | **optional.Bool**| Filter users with DIRECT or DIRECT **AND** EFFECTIVE roles.  * &#x60;false&#x60;: DIRECT roles  * &#x60;true&#x60;: DIRECT **AND** EFFECTIVE roles  DIRECT means: e.g. user gets role **directly** granted from someone with _grant permission_ right.  EFFECTIVE means: e.g. user gets role through **group membership**. | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UserData**](UserData.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUserAttributes**
> AttributesResponse RequestUserAttributes(ctx, userId, optional)
Request custom user attributes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.12.0</h3>  ### Description:   Retrieve a list of user attributes.  ### Precondition: None.  ### Postcondition: List of attributes is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `key:cn:searchString_1|value:cn:searchString_2`   Filter by attribute key contains `searchString_1` **AND** attribute value contains `searchString_2`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `key` | User attribute key filter | `cn, eq, sw` | Attribute key contains / equals / starts with value. | `search String` | | `value` | User attribute value filter | `cn, eq, sw` | Attribute value contains / equals / starts with value. | `search String` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `key:asc|value:desc`   Sort by `key` ascending **AND** by `value` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `key` | User attribute key | | `value` | User attribute value |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| User ID | 
 **optional** | ***UsersApiRequestUserAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiRequestUserAttributesOpts struct
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

# **RequestUserGroups**
> UserGroupList RequestUserGroups(ctx, userId, optional)
Request groups that user is a member of or / and can become a member

### Description:   Retrieves a list of groups a user is member of and / or can become a member.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read users</span> required.  ### Postcondition: List of groups is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `isMember:eq:false|name:cn:searchString`   Get all groups that the user is **NOT** member of **AND** whose name is like `searchString`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `name` | Group name filter | `cn` | Group name contains value. | `search String` | | `isMember` | Filter the groups which the user is (not) member of | `eq` |  | <ul><li>`true`</li><li>`false`</li><li>`any`</li></ul>default: `true` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| User ID | 
 **optional** | ***UsersApiRequestUserGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiRequestUserGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UserGroupList**](UserGroupList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUserRoles**
> RoleList RequestUserRoles(ctx, userId, optional)
Request user's granted roles

### Description:   Retrieve a list of all roles granted to a user.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read users</span> required.  ### Postcondition: List of granted roles is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| User ID | 
 **optional** | ***UsersApiRequestUserRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiRequestUserRolesOpts struct
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

# **RequestUsers**
> UserList RequestUsers(ctx, optional)
Request users

### Description:   Returns a list of DRACOON users.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read users</span> required.  ### Postcondition: List of users is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Except for `login`, `firstName` and  `lastName` - these are connected via logical disjunction (**OR**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `login:cn:searchString_1|firstName:cn:searchString_2|lockStatus:eq:2`   Filter users by login contains `searchString_1` **OR** firstName contains `searchString_2` **AND** those who are **NOT** locked.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `email` | Email filter | `eq`, `cn` | Email contains value. | `search String` | | `userName` | User name filter | `eq`, `cn` | UserName contains value. | `search String` | | `firstName` | User first name filter | `cn` | User first name contains value. | `search String` | | `lastName` | User last name filter | `cn` | User last name contains value. | `search String` | | `isLocked` | User lock status filter | `eq` |  | `true or false` | | `effectiveRoles` | Filter users with DIRECT or DIRECT **AND** EFFECTIVE roles<ul><li>`false`: DIRECT roles</li><li>`true`: DIRECT **AND** EFFECTIVE roles</li></ul>DIRECT means: e.g. user gets role **directly** granted from someone with _grant permission_ right.<br>EFFECTIVE means: e.g. user gets role through **group membership**. | `eq` |  | `true or false`<br>default: `false` | | `createdAt` | Creation date filter | `ge, le` | Creation date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `createdAt:ge:2016-12-31`&#124;`createdAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `phone` | Phone filter | `eq` | Phone equals value. | `search String` | | `isEncryptionEnabled` | Encryption status filter<ul><li>client-side encryption</li><li>private key possession</li></ul> | `eq` |  | `true or false` | | `hasRole` | User role filter<br>Depends on **effectiveRoles**.<br>For more Roles information please call `GET /roles API` | `eq` | User role  equals value. | <ul><li>`CONFIG_MANAGER` - Manage global configs</li><li>`USER_MANAGER` - Manage Users</li><li>`GROUP_MANAGER` - Manage User-Groups</li><li>`ROOM_MANAGER` - Manage top level Data Rooms</li><li>`LOG_AUDITOR` - Read logs</li><li>`NONMEMBER_VIEWER` - View users and groups when having room manage permission</li></ul> |  </details>  ### Deprecated filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | <del>`lockStatus`</del> | User lock status filter | `eq` | User lock status equals value. | <ul><li>`0` - Locked</li><li>`1` - Web access allowed</li><li>`2` - Web and mobile access allowed</li></ul> | | <del>`login`</del> | User login filter | `cn` | User login contains value. | `search String` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `firstName:asc|lastLoginSuccessAt:desc`   Sort by `firstName` ascending **AND** by `lastLoginSuccessAt` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `userName` | User name | | `email` | User email | | `firstName` | User first name | | `lastName` | User last name | | `isLocked` | User lock status | | `lastLoginSuccessAt` | Last successful login date | | `expireAt` | Expiration date | | `createdAt` | Creation date |  </details>  ### Deprecated sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | <del>`gender`</del> | Gender | | <del>`lockStatus`</del> | User lock status | | <del>`login`</del> | User login |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiRequestUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiRequestUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **includeAttributes** | **optional.Bool**| Include custom user attributes. | 
 **includeRoles** | **optional.Bool**| Include roles | 
 **includeManageableRooms** | **optional.Bool**| Include hasManageableRooms (deprecated) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UserList**](UserList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUsersRooms**
> RoomTreeDataList RequestUsersRooms(ctx, userId, optional)
Request rooms granted to the user or / and rooms that can be granted

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.10.0</h3>  ### Description:   Retrieves a list of rooms granted to the user and / or that can be granted.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read users</span> required.  ### Postcondition: List of rooms is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `isGranted:eq:true|isLastAdmin:eq:true|name:cn:searchString`   Get all rooms that the user is granted **AND** is last admin **AND** whose name is like `searchString`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `name` | Room name filter | `cn` | Room name contains value. | `search String` | | `isGranted` | Filter the rooms which the user is (not) granted. | `eq` |  | <ul><li>`true`</li><li>`false`</li><li>`any`</li></ul>default: `true` | | `isLastAdmin` | Filter the rooms which the user is last room administrator.<br>Only in connection with `isGranted:eq:true` filter possible. | `eq` |  | `true` | | `effectivePerm` | Filter rooms with DIRECT or DIRECT **AND** EFFECTIVE permissions<ul><li>`false`: DIRECT permissions</li><li>`true`: DIRECT **AND** EFFECTIVE permissions</li><li>`any`: DIRECT **AND** EFFECTIVE **AND** OVER GROUP permissions</li></ul>DIRECT means: e.g. room administrator grants `read` permissions to group of users **directly** on desired room.<br>EFFECTIVE means: e.g. group of users gets `read` permissions on desired room through **inheritance**.<br>OVER GROUP means: e.g. user gets `read` permissions on desired room through **group membership**. | `eq` |  | <ul><li>`true`</li><li>`false`</li><li>`any`</li></ul>default: `false` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| User ID | 
 **optional** | ***UsersApiRequestUsersRoomsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiRequestUsersRoomsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**RoomTreeDataList**](RoomTreeDataList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetUserAttributes**
> UserData SetUserAttributes(ctx, body, userId, optional)
Set custom user attributes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.28.0</h3>  ### Description:   Set custom user attributes.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change users</span> required.  ### Postcondition: Custom user attributes are set.  ### Further Information: Batch function.   All existing user attributes will be deleted.    * Allowed characters for keys are: `[a-zA-Z0-9_-]`   * Characters are **case-insensitive**.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserAttributes**](UserAttributes.md)|  | 
  **userId** | **int64**| User ID | 
 **optional** | ***UsersApiSetUserAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiSetUserAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**UserData**](UserData.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UserData UpdateUser(ctx, body, userId, optional)
Update user's metadata

### Description:   Update user's metadata.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change users</span> required.  ### Postcondition: User's metadata is updated.  ### Further Information: * If a user should **NOT** expire, leave `expireAt` empty. * All input fields are limited to **150** characters * **All** characters are allowed.  ### Authentication Method Options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Authentication Method | Option Key | Option Value | | :--- | :--- | :--- | | `basic` / `sql` | `username` | Unique user identifier | | `active_directory` | `ad_config_id` (optional) | Active Directory configuration ID | |  | `username` | Active Directory username according to authentication setting `userFilter` | | `radius` | `username` | RADIUS username | | `openid` | `openid_config_id` (optional) | OpenID Connect configuration ID | |  | `username` | OpenID Connect username according to authentication setting `mappingClaim` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserRequest**](UpdateUserRequest.md)|  | 
  **userId** | **int64**| User ID | 
 **optional** | ***UsersApiUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUpdateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**UserData**](UserData.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserAttributes**
> UserData UpdateUserAttributes(ctx, body, userId, optional)
Add or edit custom user attributes

### Description:   Add or edit custom user attributes. <br/><br/><span style=\"font-weight: bold; color: red;\"> &#128679; **Warning: Please note that the response with HTTP status code 200 (OK) is deprecated and will be replaced with HTTP status code 204 (No content)!**</span><br/>  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change users</span> required.  ### Postcondition: Custom user attributes gets added or edited.  ### Further Information: Batch function.   If an entry exists before, it will be overwritten.    * Allowed characters for keys are: `[a-zA-Z0-9_-]`   * Characters are **case-insensitive**.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserAttributes**](UserAttributes.md)|  | 
  **userId** | **int64**| User ID | 
 **optional** | ***UsersApiUpdateUserAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUpdateUserAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**UserData**](UserData.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

