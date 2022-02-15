# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestAuditNodeInfo**](EventlogApi.md#RequestAuditNodeInfo) | **Get** /v4/eventlog/audits/node_info | Request nodes
[**RequestAuditNodeUserData**](EventlogApi.md#RequestAuditNodeUserData) | **Get** /v4/eventlog/audits/nodes | Request node assigned users with permissions
[**RequestLogEventsAsJson**](EventlogApi.md#RequestLogEventsAsJson) | **Get** /v4/eventlog/events | Request system events
[**RequestLogOperations**](EventlogApi.md#RequestLogOperations) | **Get** /v4/eventlog/operations | Request allowed Log Operations

# **RequestAuditNodeInfo**
> AuditNodeInfoResponse RequestAuditNodeInfo(ctx, optional)
Request nodes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.31.0</h3>  ### Description:  Retrieve a list of all nodes of type room under a certain parent.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read audit log</span> required.  ### Postcondition: List of rooms.  ### Further Information: For rooms on root level, use parent_id = 0.  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `nodeName:cn:searchString_1|nodeIsEncrypted:eq:true`   Filter by node name containing `searchString_1` **AND** node is encrypted .  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `nodeId` | Node ID filter | `eq` | Node ID equals value. | `positive Integer` | | `nodeName` | Node name filter | `cn, eq, sw` | Node name contains / equals / starts with value. | `search String` | | `nodeIsEncrypted` | Encrypted node filter | `eq` |  | `true or false` |  </details>   ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `nodeName:asc`   Sort by `nodeName` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `nodeId` | Node ID | | `nodeName` | Node name |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventlogApiRequestAuditNodeInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventlogApiRequestAuditNodeInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentId** | **optional.Int64**| Parent node ID.  Only rooms can be parents.  Parent ID &#x60;0&#x60; or empty is the root node. | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**AuditNodeInfoResponse**](AuditNodeInfoResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestAuditNodeUserData**
> []AuditNodeResponse RequestAuditNodeUserData(ctx, optional)
Request node assigned users with permissions

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.32.0</h3>  ### Description:  Retrieve a list of all nodes of type room, and the room assignment users with permissions.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read audit log</span> required.  ### Postcondition: List of rooms and their assigned users is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Except for `userName`, `userFirstName` and  `userLastName` - these are connected via logical disjunction (**OR**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `userName:cn:searchString_1|userFirstName:cn:searchString_2|nodeId:eq:2`   Filter by user login containing `searchString_1` **OR** first name containing `searchString_2` **AND** node ID equals `2`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `nodeId` | Node ID filter | `eq` | Node ID equals value. | `positive Integer` | | `nodeName` | Node name filter | `cn, eq` | Node name contains / equals value. | `search String` | | `nodeParentId` | Node parent ID filter | `eq` | Parent ID equals value. | `positive Integer`<br>Parent ID `0` is the root node. | | `userId` | User ID filter | `eq` | User ID equals value. | `positive Integer` | | `userName` | Username (login) filter | `cn, eq` | Username contains / equals value. | `search String` | | `userFirstName` | User first name filter | `cn, eq` | User first name contains / equals value. | `search String` | | `userLastName` | User last name filter | `cn, eq` | User last name contains / equals value. | `search String` | | `permissionsManage` | Filter the users that do (not) have `manage` permissions in this room | `eq` |  | `true or false` | | `nodeIsEncrypted` | Encrypted node filter | `eq` |  | `true or false` | | `nodeHasActivitiesLog` | Activities log filter | `eq` |  | `true or false` |  </details>  ### Deprecated filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | <del>`nodeHasRecycleBin`</del> | Recycle bin filter<br>**Filter has no effect!** | `eq` |  | `true or false` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `nodeName:asc`   Sort by `nodeName` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `nodeId` | Node ID | | `nodeName` | Node name | | `nodeParentId` | Node parent ID | | `nodeSize` | Node size | | `nodeQuota` | Node quota |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventlogApiRequestAuditNodeUserDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventlogApiRequestAuditNodeUserDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**[]AuditNodeResponse**](AuditNodeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestLogEventsAsJson**
> LogEventList RequestLogEventsAsJson(ctx, optional)
Request system events

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.3.0</h3>  ### Description:  Retrieve eventlog (audit log) events.  ### Precondition: Role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Log Auditor</span> required.  ### Postcondition: List of audit log events is returned.  ### Further Information: Output is limited to **500** entries.   For more results please use filter criteria and paging (`offset` + `limit`).   Allowed `Accept-Header`: * `Accept: application/json` * `Accept: text/csv`    ---  Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `time:desc`   Sort by `time` descending (default sort option).  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `time` | Event timestamp |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventlogApiRequestLogEventsAsJsonOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventlogApiRequestLogEventsAsJsonOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **sort** | **optional.String**| Sort string | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **dateStart** | **optional.String**| Filter events from given date   e.g. &#x60;2015-12-31T23:59:00&#x60; | 
 **dateEnd** | **optional.String**| Filter events until given date   e.g. &#x60;2015-12-31T23:59:00&#x60; | 
 **type_** | **optional.Int32**| Operation ID   cf. &#x60;GET /eventlog/operations&#x60; | 
 **userId** | **optional.Int64**| User ID | 
 **status** | **optional.String**| Operation status:  * &#x60;0&#x60; - Success  * &#x60;2&#x60; - Error | 
 **userClient** | **optional.String**| User client | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**LogEventList**](LogEventList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestLogOperations**
> LogOperationList RequestLogOperations(ctx, optional)
Request allowed Log Operations

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.3.0</h3>  ### Description:  Retrieve eventlog (audit log) operation IDs and the associated log operation description.  ### Precondition: Role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Log Auditor</span> required.  ### Postcondition: List of available log operations is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventlogApiRequestLogOperationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventlogApiRequestLogOperationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDeprecated** | **optional.Bool**| Show only deprecated operations | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**LogOperationList**](LogOperationList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

