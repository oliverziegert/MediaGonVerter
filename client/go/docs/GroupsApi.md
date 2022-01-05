# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupMembers**](GroupsApi.md#AddGroupMembers) | **Post** /v4/groups/{group_id}/users | Add group members
[**CreateGroup**](GroupsApi.md#CreateGroup) | **Post** /v4/groups | Create new user group
[**RemoveGroup**](GroupsApi.md#RemoveGroup) | **Delete** /v4/groups/{group_id} | Remove user group
[**RemoveGroupMembers**](GroupsApi.md#RemoveGroupMembers) | **Delete** /v4/groups/{group_id}/users | Remove group members
[**RequestGroup**](GroupsApi.md#RequestGroup) | **Get** /v4/groups/{group_id} | Request user group
[**RequestGroupMembers**](GroupsApi.md#RequestGroupMembers) | **Get** /v4/groups/{group_id}/users | Request group member users or / and users who can become a member
[**RequestGroupRoles**](GroupsApi.md#RequestGroupRoles) | **Get** /v4/groups/{group_id}/roles | Request list of roles assigned to the group
[**RequestGroupRooms**](GroupsApi.md#RequestGroupRooms) | **Get** /v4/groups/{group_id}/rooms | Request rooms granted to the group or / and rooms that can be granted
[**RequestGroups**](GroupsApi.md#RequestGroups) | **Get** /v4/groups | Request list of user groups
[**RequestLastAdminRoomsGroups**](GroupsApi.md#RequestLastAdminRoomsGroups) | **Get** /v4/groups/{group_id}/last_admin_rooms | Request rooms where the group is defined as last admin group
[**UpdateGroup**](GroupsApi.md#UpdateGroup) | **Put** /v4/groups/{group_id} | Update user group&#x27;s metadata

# **AddGroupMembers**
> Group AddGroupMembers(ctx, body, groupId, optional)
Add group members

### Description: Add members to a group.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change groups</span> required.  ### Postcondition:  New members are added to the group.  ### Further Information: Batch function.   The newly provided members will be added to the existing ones.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeGroupMembersRequest**](ChangeGroupMembersRequest.md)|  | 
  **groupId** | **int64**| Group ID | 
 **optional** | ***GroupsApiAddGroupMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiAddGroupMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> Group CreateGroup(ctx, body, optional)
Create new user group

### Description: Create a new user group.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change groups</span> required.  ### Postcondition:  A new user group is created.  ### Further Information: * If a group should **NOT** expire, leave `expireAt` empty. * Group names are limited to **150** characters * **All** characters are allowed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateGroupRequest**](CreateGroupRequest.md)|  | 
 **optional** | ***GroupsApiCreateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiCreateGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveGroup**
> RemoveGroup(ctx, groupId, optional)
Remove user group

### Description: Delete a user group.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; delete groups</span> required.  ### Postcondition:  User group is deleted.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 
 **optional** | ***GroupsApiRemoveGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiRemoveGroupOpts struct
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

# **RemoveGroupMembers**
> Group RemoveGroupMembers(ctx, body, groupId, optional)
Remove group members

### Description:   Remove group members.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change groups</span> required.  ### Postcondition:  Provided users are removed from the user group.  ### Further Information: Batch function.   The provided users are removed from the user group. Maximum number of users to remove in one request is 200.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeGroupMembersRequest**](ChangeGroupMembersRequest.md)|  | 
  **groupId** | **int64**| Group ID | 
 **optional** | ***GroupsApiRemoveGroupMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiRemoveGroupMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestGroup**
> Group RequestGroup(ctx, groupId, optional)
Request user group

### Description:   Retrieve detailed information about a user group.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read groups</span> required.  ### Postcondition:  User group is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 
 **optional** | ***GroupsApiRequestGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiRequestGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestGroupMembers**
> GroupUserList RequestGroupMembers(ctx, groupId, optional)
Request group member users or / and users who can become a member

### Description:   Retrieve a list of group member users or / and users who can become a member.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read groups</span> required.  ### Postcondition:  List of users is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `isMember:eq:false|user:cn:searchString`   Get all users that are **NOT** in this group **AND** whose (`firstName` **OR** `lastName` **OR** `email` **OR** `username`) is like `searchString`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `user` | User filter | `cn` | User contains value (`firstName` **OR** `lastName` **OR** `email` **OR** `username`). | `search String` | | `isMember` | Filter group members | `eq` |  | <ul><li>`true`</li><li>`false`</li><li>`any`</li></ul>default: `true` |  </details>  ### Deprecated filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | <del>`displayName`</del> | User display name filter (use `user` filter) | `cn` | User display name contains value (`firstName` **OR** `lastName` **OR** `email`). | `search String` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 
 **optional** | ***GroupsApiRequestGroupMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiRequestGroupMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**GroupUserList**](GroupUserList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestGroupRoles**
> RoleList RequestGroupRoles(ctx, groupId, optional)
Request list of roles assigned to the group

### Description:   Retrieve a list of all roles granted to a group.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read groups</span> required.  ### Postcondition:  List of granted roles is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 
 **optional** | ***GroupsApiRequestGroupRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiRequestGroupRolesOpts struct
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

# **RequestGroupRooms**
> RoomTreeDataList RequestGroupRooms(ctx, groupId, optional)
Request rooms granted to the group or / and rooms that can be granted

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.10.0</h3>  ### Description:   Retrieves a list of rooms granted to the group and / or that can be granted.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read groups</span> required.  ### Postcondition:  List of rooms is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `isGranted:eq:false|name:cn:searchString`   Get all rooms where the group is **NOT** granted **AND** whose name is like `searchString`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `name` | Room name filter | `cn` | Room name contains value. | `search String` | | `isGranted` | Filter rooms which the group is (not) granted | `eq` |  | <ul><li>`true`</li><li>`false`</li><li>`any`</li></ul>default: `true` | | `effectivePerm` | Filter rooms with DIRECT or DIRECT **AND** EFFECTIVE permissions<ul><li>`false`: DIRECT permissions</li><li>`true`:  DIRECT **AND** EFFECTIVE permissions</li></ul>DIRECT means: e.g. room administrator grants `read` permissions to group of users **directly** on desired room.<br>EFFECTIVE means: e.g. group of users gets `read` permissions on desired room through **inheritance**. | `eq` |  | `true or false`<br>default: `true` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 
 **optional** | ***GroupsApiRequestGroupRoomsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiRequestGroupRoomsOpts struct
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

# **RequestGroups**
> GroupList RequestGroups(ctx, optional)
Request list of user groups

### Description:   Returns a list of user groups.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read groups</span> required.  ### Postcondition:  List of user groups is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:cn:searchString`   Filter by group name containing `searchString`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `name` | Group name filter | `cn` | Group name contains value. | `search String` | | `hasRole` | (**`NEW`**) Group role filter<br>For more information about roles check **`GET /roles`** API | `eq` | Group role equals value. | <ul><li>`CONFIG_MANAGER` - Manages global configuration</li><li>`USER_MANAGER` - Manages users</li><li>`GROUP_MANAGER` - Manages user groups</li><li>`ROOM_MANAGER` - Manages top level rooms</li><li>`LOG_AUDITOR` - Reads audit logs</li><li>`NONMEMBER_VIEWER` - Views users and groups when having room _\"manage\"_ permission</li></ul> |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:asc|expireAt:desc`   Sort by `name` ascending **AND** by `expireAt` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `name` | Group name | | `createdAt` | Creation date | | `expireAt` | Expiration date | | `cntUsers` | Amount of users |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiRequestGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiRequestGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**GroupList**](GroupList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestLastAdminRoomsGroups**
> LastAdminGroupRoomList RequestLastAdminRoomsGroups(ctx, groupId, optional)
Request rooms where the group is defined as last admin group

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.10.0</h3>  ### Description:   Retrieve a list of all rooms where the group is defined as last admin group.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change groups</span> required.  ### Postcondition:  List of rooms is returned.   ### Further Information: An empty list is returned if no rooms were found where the group is defined as last admin group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 
 **optional** | ***GroupsApiRequestLastAdminRoomsGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiRequestLastAdminRoomsGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**LastAdminGroupRoomList**](LastAdminGroupRoomList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> Group UpdateGroup(ctx, body, groupId, optional)
Update user group's metadata

### Description:   Update user group's metadata .  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change groups</span> required.  ### Postcondition:  User group's metadata is changed.  ### Further Information: * If a group should **NOT** expire, leave `expireAt` empty. * Group names are limited to **150** characters * **All** characters are allowed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateGroupRequest**](UpdateGroupRequest.md)|  | 
  **groupId** | **int64**| Group ID | 
 **optional** | ***GroupsApiUpdateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiUpdateGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

