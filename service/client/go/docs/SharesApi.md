# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDownloadShare**](SharesApi.md#CreateDownloadShare) | **Post** /v4/shares/downloads | Create new Download Share
[**CreateUploadShare**](SharesApi.md#CreateUploadShare) | **Post** /v4/shares/uploads | Create new Upload Share
[**DeleteDownloadShares**](SharesApi.md#DeleteDownloadShares) | **Delete** /v4/shares/downloads | Remove Download Shares
[**DeleteUploadShares**](SharesApi.md#DeleteUploadShares) | **Delete** /v4/shares/uploads | Remove Upload Shares
[**RemoveDownloadShare**](SharesApi.md#RemoveDownloadShare) | **Delete** /v4/shares/downloads/{share_id} | Remove Download Share
[**RemoveUploadShare**](SharesApi.md#RemoveUploadShare) | **Delete** /v4/shares/uploads/{share_id} | Remove Upload Share
[**RequestDownloadShare**](SharesApi.md#RequestDownloadShare) | **Get** /v4/shares/downloads/{share_id} | Request Download Share
[**RequestDownloadShareQr**](SharesApi.md#RequestDownloadShareQr) | **Get** /v4/shares/downloads/{share_id}/qr | Request Download Share via QR Code
[**RequestDownloadShares**](SharesApi.md#RequestDownloadShares) | **Get** /v4/shares/downloads | Request list of Download Shares
[**RequestUploadShare**](SharesApi.md#RequestUploadShare) | **Get** /v4/shares/uploads/{share_id} | Request Upload Share
[**RequestUploadShareQr**](SharesApi.md#RequestUploadShareQr) | **Get** /v4/shares/uploads/{share_id}/qr | Request Upload Share via QR Code
[**RequestUploadShares**](SharesApi.md#RequestUploadShares) | **Get** /v4/shares/uploads | Request list of Upload Shares
[**SendDownloadShareLinkViaEmail**](SharesApi.md#SendDownloadShareLinkViaEmail) | **Post** /v4/shares/downloads/{share_id}/email | Send an existing Download Share link via email
[**SendUploadShareLinkViaEmail**](SharesApi.md#SendUploadShareLinkViaEmail) | **Post** /v4/shares/uploads/{share_id}/email | Send an existing Upload Share link via email
[**UpdateDownloadShare**](SharesApi.md#UpdateDownloadShare) | **Put** /v4/shares/downloads/{share_id} | Update Download Share
[**UpdateDownloadShares**](SharesApi.md#UpdateDownloadShares) | **Put** /v4/shares/downloads | Update a list of Download Shares
[**UpdateUploadShare**](SharesApi.md#UpdateUploadShare) | **Put** /v4/shares/uploads/{share_id} | Update Upload Share
[**UpdateUploadShares**](SharesApi.md#UpdateUploadShares) | **Put** /v4/shares/uploads | Update List of Upload Shares

# **CreateDownloadShare**
> DownloadShare CreateDownloadShare(ctx, body, optional)
Create new Download Share

### Description: Create a new Download Share.  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span> permissions on target node.  ### Postcondition: Download Share is created.  ### Further Information:  If the target node is a room: subordinary rooms are excluded from a Download Share.  * `name` is limited to **150** characters. * `notes` are limited to **255** characters. * `password` is limited to **1024** characters.  Use `POST /shares/downloads/{share_id}/email` API for sending emails.    Forbidden characters in passwords: [`&`, `'`, `<`, `>`]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDownloadShareRequest**](CreateDownloadShareRequest.md)|  | 
 **optional** | ***SharesApiCreateDownloadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiCreateDownloadShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**DownloadShare**](DownloadShare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUploadShare**
> UploadShare CreateUploadShare(ctx, body, optional)
Create new Upload Share

### Description: Create a new Upload Share (aka File Request).  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage upload share</span> permissions on target container.  ### Postcondition: Upload Share is created.  ### Further Information:  * `name` is limited to **150** characters. * `notes` are limited to **255** characters. * `password` is limited to **1024** characters.  Forbidden characters in passwords: [`&`, `'`, `<`, `>`]    Use `POST /shares/uploads/{share_id}/email` API for sending emails. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUploadShareRequest**](CreateUploadShareRequest.md)|  | 
 **optional** | ***SharesApiCreateUploadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiCreateUploadShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**UploadShare**](UploadShare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDownloadShares**
> DeleteDownloadShares(ctx, body, optional)
Remove Download Shares

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.21.0</h3>  ### Functional Description: Delete multiple Download Shares.  ### Precondition: User with _\"manage download share\"_ permissions on target nodes.  ### Postcondition: Download Shares are deleted.  ### Further Information: Only the Download Shares are removed; the referenced files or containers persists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteDownloadSharesRequest**](DeleteDownloadSharesRequest.md)|  | 
 **optional** | ***SharesApiDeleteDownloadSharesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiDeleteDownloadSharesOpts struct
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

# **DeleteUploadShares**
> DeleteUploadShares(ctx, body, optional)
Remove Upload Shares

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.21.0</h3>  ### Functional Description: Delete multiple Upload Shares (aka Upload Accounts).  ### Precondition: User has _\"manage upload share\"_ permissions on target containers.  ### Postcondition: Upload Shares are deleted.  ### Further Information: Only the Upload Shares are removed; already uploaded files and the target container persist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteUploadSharesRequest**](DeleteUploadSharesRequest.md)|  | 
 **optional** | ***SharesApiDeleteUploadSharesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiDeleteUploadSharesOpts struct
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

# **RemoveDownloadShare**
> RemoveDownloadShare(ctx, shareId, optional)
Remove Download Share

### Description: Delete a Download Share.  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span> permissions on target node.  ### Postcondition: Download Share is deleted.  ### Further Information: Only the Download Share is removed; the referenced file or container persists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shareId** | **int64**| Share ID | 
 **optional** | ***SharesApiRemoveDownloadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiRemoveDownloadShareOpts struct
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

# **RemoveUploadShare**
> RemoveUploadShare(ctx, shareId, optional)
Remove Upload Share

### Description: Delete an Upload Share (aka File Request).  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage upload share</span> permissions on target container.  ### Postcondition: Upload Share is deleted.  ### Further Information: Only the Upload Share is removed; already uploaded files and the target container persist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shareId** | **int64**| Share ID | 
 **optional** | ***SharesApiRemoveUploadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiRemoveUploadShareOpts struct
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

# **RequestDownloadShare**
> DownloadShare RequestDownloadShare(ctx, shareId, optional)
Request Download Share

### Description:   Retrieve detailed information about one Download Share.  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span> permissions on target node.  ### Postcondition: Download Share is returned  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shareId** | **int64**| Share ID | 
 **optional** | ***SharesApiRequestDownloadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiRequestDownloadShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**DownloadShare**](DownloadShare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestDownloadShareQr**
> DownloadShare RequestDownloadShareQr(ctx, shareId, optional)
Request Download Share via QR Code

### Description:   Retrieve detailed information about one Download Share.  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span> permissions on target node.  ### Postcondition: Download Share is returned  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shareId** | **int64**| Share ID | 
 **optional** | ***SharesApiRequestDownloadShareQrOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiRequestDownloadShareQrOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**DownloadShare**](DownloadShare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestDownloadShares**
> DownloadShareList RequestDownloadShares(ctx, optional)
Request list of Download Shares

### Description:   Retrieve a list of Download Shares.  ### Precondition: Authenticated user.  ### Postcondition: List of available Download Shares is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical (**AND**). createdBy and updatedBy searches several user-related attributes.  Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:cn:searchString_1|createdBy:cn:searchString_2` Filter by file name contains `searchString_1` **AND** creator info (`firstName` **OR** `lastName` **OR** `email` **OR** `username`) contains `searchString_2`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `name` | Alias or node name filter | `cn` | Alias or node name contains value. | `search String` | | `createdAt` | Creation date filter | `ge, le` | Creation date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `createdAt:ge:2016-12-31`&#124;`createdAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `createdBy` | Creator info filter | `cn, eq` | Creator info (`firstName` **OR** `lastName` **OR** `email` **OR** `username`) contains value. | `search String` | | `createdById` | Creator ID filter | `eq` | Creator ID equals value. | `positive Integer` | | `accessKey` | Share access key filter | `cn` | Share access key contains values. | `search String` | | `nodeId` | Source node ID | `eq` | Source node (room, folder, file) ID equals value. | `positive Integer` | | `updatedBy` | Modifier info filter | `cn, eq` | Modifier info (`firstName` **OR** `lastName` **OR** `email` **OR** `username`) contains value. | `search String` | | `updatedById` | Modifier ID filter | `eq` | Modifier ID equals value. | `positive Integer` |  </details>  ### Deprecated filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | <del>`userId`</del>  | Creator user ID | `eq` | Creator user ID equals value. Use `createdById` instead | `positive Integer` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:asc|expireAt:desc`   Sort by `name` ascending **AND** by `expireAt` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `name` | Alias or node name | | `notifyCreator` | Notify creator on every download | | `expireAt` | Expiration date | | `createdAt` | Creation date | | `createdBy` | Creator first name, last name | | `classification` | Classification ID:<ul><li>1 - public</li><li>2 - internal</li><li>3 - confidential</li><li>4 - strictly confidential</li></ul> |  </details> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SharesApiRequestDownloadSharesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiRequestDownloadSharesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**DownloadShareList**](DownloadShareList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUploadShare**
> UploadShare RequestUploadShare(ctx, shareId, optional)
Request Upload Share

### Description:   Retrieve detailed information about one Upload Share (aka File Request).  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage upload share</span> permissions on target container.  ### Postcondition: Upload Share is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shareId** | **int64**| Share ID | 
 **optional** | ***SharesApiRequestUploadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiRequestUploadShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UploadShare**](UploadShare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUploadShareQr**
> UploadShare RequestUploadShareQr(ctx, shareId, optional)
Request Upload Share via QR Code

### Description:   Retrieve detailed information about one Upload Share (aka File Request).  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage upload share</span> permissions on target container.  ### Postcondition: Upload Share is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shareId** | **int64**| Share ID | 
 **optional** | ***SharesApiRequestUploadShareQrOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiRequestUploadShareQrOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UploadShare**](UploadShare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUploadShares**
> UploadShareList RequestUploadShares(ctx, optional)
Request list of Upload Shares

### Description:   Retrieve a list of Upload Shares (aka File Requests).  ### Precondition: Authenticated user.  ### Postcondition: List of available Upload Shares is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical (**AND**). createdBy and updatedBy searches several user-related attributes. Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:cn:searchString_1|createdBy:cn:searchString_2`   Filter by alias name contains `searchString_1` **AND** creator info (`firstName` **OR** `lastName` **OR** `email` **OR** `username`) contains `searchString_2`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `name` | Alias name filter | `cn` | Alias name contains value. | `search String` | | `createdAt` | Creation date filter | `ge, le` | Creation date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `createdAt:ge:2016-12-31`&#124;`createdAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `createdBy` | Creator info filter | `cn, eq` | Creator info (`firstName` **OR** `lastName` **OR** `email` **OR** `username`) contains value. | `search String` | | `createdById` | Creator ID filter | `eq` | Creator ID equals value. | `positive Integer` | | `accessKey` | Share access key filter | `cn` | Share access key contains values. | `search String` | | `userId` | Creator user ID | `eq` | Creator user ID equals value. | `positive Integer` | | `targetId` | Target node ID | `eq` | Target node (room, folder) ID equals value. | `positive Integer` | | `updatedBy` | Modifier info filter | `cn, eq` | Modifier info (`firstName` **OR** `lastName` **OR** `email` **OR** `username`) contains value. | `search String` | | `updatedById` | Modifier ID filter | `eq` | Modifier ID equals value. | `positive Integer` |  </details>  ### Deprecated filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | <del>`targetId`</del> | Target node ID | `cn` | Target node (room, folder) ID equals value. | `positive Integer` | | <del>`userId` </del>| Creator user ID | `eq` | Creator user ID equals value. Use `createdById` instead. | `positive Integer` |  </details>  ---  Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:asc|expireAt:desc`   Sort by `name` ascending **AND** by `expireAt` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `name` | Alias name | | `notifyCreator` | Notify creator on every upload | | `expireAt` | Expiration date | | `createdAt` | Creation date | | `createdBy` | Creator first name, last name |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SharesApiRequestUploadSharesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiRequestUploadSharesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UploadShareList**](UploadShareList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendDownloadShareLinkViaEmail**
> SendDownloadShareLinkViaEmail(ctx, body, shareId, optional)
Send an existing Download Share link via email

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.11.0</h3>  ### Description: Send an email to specific recipients for existing Download Share.  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span> permissions on target node.  ### Postcondition: Download Share link successfully sent.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DownloadShareLinkEmail**](DownloadShareLinkEmail.md)|  | 
  **shareId** | **int64**| Share ID | 
 **optional** | ***SharesApiSendDownloadShareLinkViaEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiSendDownloadShareLinkViaEmailOpts struct
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

# **SendUploadShareLinkViaEmail**
> SendUploadShareLinkViaEmail(ctx, body, shareId, optional)
Send an existing Upload Share link via email

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.11.0</h3>  ### Description: Send an email to specific recipients for existing Upload Share.  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage upload share</span> permissions on target container.  ### Postcondition: Upload Share link successfully sent.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UploadShareLinkEmail**](UploadShareLinkEmail.md)|  | 
  **shareId** | **int64**| Share ID | 
 **optional** | ***SharesApiSendUploadShareLinkViaEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiSendUploadShareLinkViaEmailOpts struct
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

# **UpdateDownloadShare**
> DownloadShare UpdateDownloadShare(ctx, body, shareId, optional)
Update Download Share

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.11.0</h3>  ### Description: Update an existing Download Share.  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span> permissions on target node.  ### Postcondition: Download Share is successfully updated.  ### Further Information: * `name` is limited to **150** characters. * `notes` are limited to **255** characters. * `password` is limited to **1024** characters.  Forbidden characters in passwords: [`&`, `'`, `<`, `>`]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDownloadShareRequest**](UpdateDownloadShareRequest.md)|  | 
  **shareId** | **int64**| Share ID | 
 **optional** | ***SharesApiUpdateDownloadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiUpdateDownloadShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**DownloadShare**](DownloadShare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDownloadShares**
> UpdateDownloadShares(ctx, body, optional)
Update a list of Download Shares

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.25.0</h3>  ### Description: Update a list of existing Download Shares.  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span> permissions on target node.  ### Postcondition: Download Shares are successfully updated.  ### Further Information: Maximum number of shares is 200

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDownloadSharesBulkRequest**](UpdateDownloadSharesBulkRequest.md)|  | 
 **optional** | ***SharesApiUpdateDownloadSharesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiUpdateDownloadSharesOpts struct
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

# **UpdateUploadShare**
> UploadShare UpdateUploadShare(ctx, body, shareId, optional)
Update Upload Share

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.11.0</h3>  ### Description: Update existing Upload Share (aka File Request).  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage upload share</span> permissions on target container.  ### Postcondition: Upload Share successfully updated.  ### Further Information:  * `name` is limited to **150** characters. * `notes` are limited to **255** characters. * `password` is limited to **1024** characters.  Forbidden characters in passwords: [`&`, `'`, `<`, `>`]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUploadShareRequest**](UpdateUploadShareRequest.md)|  | 
  **shareId** | **int64**| Share ID | 
 **optional** | ***SharesApiUpdateUploadShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiUpdateUploadShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**UploadShare**](UploadShare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUploadShares**
> UpdateUploadShares(ctx, body, optional)
Update List of Upload Shares

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.25.0</h3>  ### Description: Update a list of existing Upload Shares (aka File Request).  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage upload share</span> permissions on target container.  ### Postcondition: Upload Shares successfully updated.  ### Further Information: Maximum number of shares is 200

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUploadSharesBulkRequest**](UpdateUploadSharesBulkRequest.md)|  | 
 **optional** | ***SharesApiUpdateUploadSharesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SharesApiUpdateUploadSharesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

