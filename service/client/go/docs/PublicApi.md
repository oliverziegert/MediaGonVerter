# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFileUploadViaShare**](PublicApi.md#CancelFileUploadViaShare) | **Delete** /v4/public/shares/uploads/{access_key}/{upload_id} | Cancel file upload
[**CompleteFileUploadViaShare**](PublicApi.md#CompleteFileUploadViaShare) | **Put** /v4/public/shares/uploads/{access_key}/{upload_id} | Complete file upload
[**CompleteS3FileUploadViaShare**](PublicApi.md#CompleteS3FileUploadViaShare) | **Put** /v4/public/shares/uploads/{access_key}/{upload_id}/s3 | Complete S3 file upload
[**CreateShareUploadChannel**](PublicApi.md#CreateShareUploadChannel) | **Post** /v4/public/shares/uploads/{access_key} | Create new file upload channel
[**DownloadFileViaTokenPublic**](PublicApi.md#DownloadFileViaTokenPublic) | **Head** /v4/public/shares/downloads/{access_key}/{token} | Download file with token
[**DownloadFileViaTokenPublic1**](PublicApi.md#DownloadFileViaTokenPublic1) | **Get** /v4/public/shares/downloads/{access_key}/{token} | Download file with token
[**GenerateDownloadUrlPublic**](PublicApi.md#GenerateDownloadUrlPublic) | **Post** /v4/public/shares/downloads/{access_key} | Generate download URL
[**GeneratePresignedUrlsPublic**](PublicApi.md#GeneratePresignedUrlsPublic) | **Post** /v4/public/shares/uploads/{access_key}/{upload_id}/s3_urls | Generate presigned URLs for S3 file upload
[**RequestActiveDirectoryAuthInfo**](PublicApi.md#RequestActiveDirectoryAuthInfo) | **Get** /v4/public/system/info/auth/ad | Request Active Directory authentication information
[**RequestOpenIdAuthInfo**](PublicApi.md#RequestOpenIdAuthInfo) | **Get** /v4/public/system/info/auth/openid | Request OpenID Connect provider authentication information
[**RequestPublicDownloadShareInfo**](PublicApi.md#RequestPublicDownloadShareInfo) | **Get** /v4/public/shares/downloads/{access_key} | Request public Download Share information
[**RequestPublicUploadShareInfo**](PublicApi.md#RequestPublicUploadShareInfo) | **Get** /v4/public/shares/uploads/{access_key} | Request public Upload Share information
[**RequestSoftwareVersion**](PublicApi.md#RequestSoftwareVersion) | **Get** /v4/public/software/version | Request software version information
[**RequestSystemInfo**](PublicApi.md#RequestSystemInfo) | **Get** /v4/public/system/info | Request system information
[**RequestSystemTime**](PublicApi.md#RequestSystemTime) | **Get** /v4/public/time | Request system time
[**RequestThirdPartyDependencies**](PublicApi.md#RequestThirdPartyDependencies) | **Get** /v4/public/software/third_party_dependencies | Request third-party software dependencies
[**RequestUploadStatusPublic**](PublicApi.md#RequestUploadStatusPublic) | **Get** /v4/public/shares/uploads/{access_key}/{upload_id} | Request status of S3 file upload
[**UploadFileAsMultipartPublic1**](PublicApi.md#UploadFileAsMultipartPublic1) | **Post** /v4/public/shares/uploads/{access_key}/{upload_id} | Upload file

# **CancelFileUploadViaShare**
> CancelFileUploadViaShare(ctx, accessKey, uploadId)
Cancel file upload

### Description: Abort (chunked) upload via Upload Share.  ### Precondition: Valid Upload ID.  ### Postcondition: Aborts upload and invalidates upload ID / token.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessKey** | **string**| Access key | 
  **uploadId** | **string**| Upload channel ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompleteFileUploadViaShare**
> PublicUploadedFileData CompleteFileUploadViaShare(ctx, accessKey, uploadId, optional)
Complete file upload

### Description: Finalize (chunked) upload via Upload Share.  ### Precondition: Valid upload ID.   Only returns users that owns one of the following permissions: <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage upload share</span>  ### Postcondition: Finalizes upload.  ### Further Information: Chunked uploads (range requests) are supported.    Please ensure that all chunks have been transferred correctly before finishing the upload.   If file hash has been created in time a `201 Created` will be responded and hash will be part of response, otherwise it will be a `202 Accepted` without it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessKey** | **string**| Access key | 
  **uploadId** | **string**| Upload channel ID | 
 **optional** | ***PublicApiCompleteFileUploadViaShareOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiCompleteFileUploadViaShareOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UserFileKeyList**](UserFileKeyList.md)|  | 
 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 

### Return type

[**PublicUploadedFileData**](PublicUploadedFileData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompleteS3FileUploadViaShare**
> CompleteS3FileUploadViaShare(ctx, body, accessKey, uploadId)
Complete S3 file upload

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.15.0</h3>  ### Description: Finishes a S3 file upload and closes the corresponding upload channel.  ### Precondition: Valid upload ID.   Only returns users that owns one of the following permissions: <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage upload share</span>  ### Postcondition: Upload channel is closed. S3 multipart upload request is completed.  ### Further Information: None. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CompleteS3ShareUploadRequest**](CompleteS3ShareUploadRequest.md)|  | 
  **accessKey** | **string**| Access key | 
  **uploadId** | **string**| Upload channel ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateShareUploadChannel**
> CreateShareUploadChannelResponse CreateShareUploadChannel(ctx, body, accessKey)
Create new file upload channel

### Description:   Create a new upload channel.  ### Precondition: None.  ### Postcondition: Upload channel is created and corresponding upload URL, token & upload ID are returned.  ### Further Information: Use `uploadUrl` the upload `token` is deprecated.    Please provide the size of the intended upload so that the quota can be checked in advanced and no data is transferred unnecessarily.  ### Node naming convention: * Node (room, folder, file) names are limited to **150** characters. * Illegal names:   `'CON', 'PRN', 'AUX', 'NUL', 'COM1', 'COM2', 'COM3', 'COM4', 'COM5', 'COM6', 'COM7', 'COM8', 'COM9', 'LPT1', 'LPT2', 'LPT3', 'LPT4', 'LPT5', 'LPT6', 'LPT7', 'LPT8', 'LPT9', (and any of those with an extension)` * Illegal characters in names:   `'\\\\', '<','>', ':', '\\\"', '|', '?', '*', '/', leading '-', trailing '.' ` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateShareUploadChannelRequest**](CreateShareUploadChannelRequest.md)|  | 
  **accessKey** | **string**| Access key | 

### Return type

[**CreateShareUploadChannelResponse**](CreateShareUploadChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFileViaTokenPublic**
> DownloadFileViaTokenPublic(ctx, accessKey, token, optional)
Download file with token

### Description:   Download a file (or zip archive if target is a folder or room).  ### Precondition: Valid download token.  ### Postcondition: Stream is returned.  ### Further Information: Range requests are supported.   Range requests are illegal for zip archive download.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessKey** | **string**| Access key | 
  **token** | **string**| Download token | 
 **optional** | ***PublicApiDownloadFileViaTokenPublicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiDownloadFileViaTokenPublicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **range_** | **optional.String**| Range   e.g. &#x60;bytes&#x3D;0-999&#x60; | 
 **genericMimetype** | **optional.Bool**| Always return &#x60;application/octet-stream&#x60; instead of specific mimetype | 
 **inline** | **optional.Bool**| Use Content-Disposition: &#x60;inline&#x60; instead of &#x60;attachment&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFileViaTokenPublic1**
> DownloadFileViaTokenPublic1(ctx, accessKey, token, optional)
Download file with token

### Description:   Download a file (or zip archive if target is a folder or room).  ### Precondition: Valid download token.  ### Postcondition: Stream is returned.  ### Further Information: Range requests are supported.   Range requests are illegal for zip archive download.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessKey** | **string**| Access key | 
  **token** | **string**| Download token | 
 **optional** | ***PublicApiDownloadFileViaTokenPublic1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiDownloadFileViaTokenPublic1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **range_** | **optional.String**| Range   e.g. &#x60;bytes&#x3D;0-999&#x60; | 
 **genericMimetype** | **optional.Bool**| Always return &#x60;application/octet-stream&#x60; instead of specific mimetype | 
 **inline** | **optional.Bool**| Use Content-Disposition: &#x60;inline&#x60; instead of &#x60;attachment&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateDownloadUrlPublic**
> PublicDownloadTokenGenerateResponse GenerateDownloadUrlPublic(ctx, accessKey, optional)
Generate download URL

### Description: Generate a download URL to retrieve a shared file.  ### Precondition: None.  ### Postcondition: Download URL and token are generated and returned.  ### Further Information: Use `downloadUrl` the download `token` is deprecated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessKey** | **string**| Access key | 
 **optional** | ***PublicApiGenerateDownloadUrlPublicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiGenerateDownloadUrlPublicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PublicDownloadTokenGenerateRequest**](PublicDownloadTokenGenerateRequest.md)|  | 

### Return type

[**PublicDownloadTokenGenerateResponse**](PublicDownloadTokenGenerateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GeneratePresignedUrlsPublic**
> PresignedUrlList GeneratePresignedUrlsPublic(ctx, body, accessKey, uploadId, optional)
Generate presigned URLs for S3 file upload

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.15.0</h3>  ### Description: Generate presigned URLs for S3 file upload.  ### Precondition: Valid upload ID  ### Postcondition: List of presigned URLs is returned.  ### Further Information: The size for each part must be >= 5 MB, except for the last part.   The part number of the first part in S3 is 1 (not 0).   Use HTTP method `PUT` for uploading bytes via presigned URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GeneratePresignedUrlsRequest**](GeneratePresignedUrlsRequest.md)|  | 
  **accessKey** | **string**| Access key | 
  **uploadId** | **string**| Upload channel ID | 
 **optional** | ***PublicApiGeneratePresignedUrlsPublicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiGeneratePresignedUrlsPublicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 

### Return type

[**PresignedUrlList**](PresignedUrlList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestActiveDirectoryAuthInfo**
> ActiveDirectoryAuthInfo RequestActiveDirectoryAuthInfo(ctx, optional)
Request Active Directory authentication information

### Description:   Provides information about Active Directory authentication options.  ### Precondition: None.  ### Postcondition: Active Directory authentication options information is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicApiRequestActiveDirectoryAuthInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiRequestActiveDirectoryAuthInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isGlobalAvailable** | **optional.Bool**| Show only global available items | 

### Return type

[**ActiveDirectoryAuthInfo**](ActiveDirectoryAuthInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOpenIdAuthInfo**
> OpenIdAuthInfo RequestOpenIdAuthInfo(ctx, optional)
Request OpenID Connect provider authentication information

### Description:   Provides information about OpenID Connect authentication options.  ### Precondition: None.  ### Postcondition: OpenID Connect authentication options information is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicApiRequestOpenIdAuthInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiRequestOpenIdAuthInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isGlobalAvailable** | **optional.Bool**| Show only global available items | 

### Return type

[**OpenIdAuthInfo**](OpenIdAuthInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestPublicDownloadShareInfo**
> PublicDownloadShare RequestPublicDownloadShareInfo(ctx, accessKey, optional)
Request public Download Share information

### Description:   Retrieve the public information of a Download Share.  ### Precondition: None.  ### Postcondition: Download Share information is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessKey** | **string**| Access key | 
 **optional** | ***PublicApiRequestPublicDownloadShareInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiRequestPublicDownloadShareInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 

### Return type

[**PublicDownloadShare**](PublicDownloadShare.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestPublicUploadShareInfo**
> PublicUploadShare RequestPublicUploadShareInfo(ctx, accessKey, optional)
Request public Upload Share information

### Description:   Provides information about the desired Upload Share.  ### Precondition: Only `userUserPublicKeyList` is returned to the users who owns one of the following permissions: <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage upload share</span>  ### Postcondition: None.  ### Further Information: If no password is set, the returned information is reduced to the following attributes (if available):  * `name` * `createdAt` * `isProtected` * `isEncrypted` * `showUploadedFiles` * `userUserPublicKeyList` (if parent is end-to-end encrypted)  Only if the password is transmitted as `X-Sds-Share-Password` header, all values are returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessKey** | **string**| Access key | 
 **optional** | ***PublicApiRequestPublicUploadShareInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiRequestPublicUploadShareInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsSharePassword** | **optional.String**| Upload share password. Should be base64-encoded.  Plain X-Sds-Share-Passwords are *deprecated* and will be removed in the future | 
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 

### Return type

[**PublicUploadShare**](PublicUploadShare.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSoftwareVersion**
> SoftwareVersionData RequestSoftwareVersion(ctx, optional)
Request software version information

### Description:   Public software version information.  ### Precondition: None.  ### Postcondition: Sofware version information is returned.  ### Further Information: The version of DRACOON Server consists of two components: * **API** * **Core** (referred to as _\"Server\"_)  which are versioned individually.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicApiRequestSoftwareVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiRequestSoftwareVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 

### Return type

[**SoftwareVersionData**](SoftwareVersionData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSystemInfo**
> SystemInfo RequestSystemInfo(ctx, optional)
Request system information

### Description:   Provides information about system.  ### Precondition: None.  ### Postcondition: System information is returned.  ### Further Information: Authentication methods are sorted by **priority** attribute.   Smaller values have higher priority.   Authentication method with highest priority is considered as default.  ### System information: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | `languageDefault` | Defines which language should be default. | `ISO 639-1 code` | | `hideLoginInputFields` | Defines if login fields should be hidden. | `true or false` | | `s3Hosts` | List of available S3 hosts. | `String array` | | `s3EnforceDirectUpload` | Determines whether S3 direct upload is enforced or not. | `true or false` | | `useS3Storage` | Determines whether S3 Storage enabled and used. | `true or false` |  </details>  ### Authentication methods: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Authentication Method | Description | | :--- | :--- | | `basic` | **Basic** authentication globally allowed.<br>This option **MUST** be activated to allow users to log in with their credentials stored in the database.<br>Formerly known as `sql`. | | `active_directory` | **Active Directory** authentication globally allowed.<br>This option **MUST** be activated to allow users to log in with their Active Directory credentials. | | `radius` | **RADIUS** authentication globally allowed.<br>This option **MUST** be activated to allow users to log in with their RADIUS username, their PIN and a token password. | | `openid` | **OpenID Connect** authentication globally allowed.This option **MUST** be activated to allow users to log in with their OpenID Connect identity. | | `hideLoginInputFields` | Determines whether input fields for login should be enabled | `true or false` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicApiRequestSystemInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiRequestSystemInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isEnabled** | **optional.Bool**| Show only enabled authentication methods | 

### Return type

[**SystemInfo**](SystemInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSystemTime**
> SdsServerTime RequestSystemTime(ctx, optional)
Request system time

### Description:   Retrieve the actual server time.  ### Precondition: None.  ### Postcondition: Server time is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicApiRequestSystemTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiRequestSystemTimeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 

### Return type

[**SdsServerTime**](SdsServerTime.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestThirdPartyDependencies**
> []ThirdPartyDependenciesData RequestThirdPartyDependencies(ctx, )
Request third-party software dependencies

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.9.0</h3>  ### Description:   Provides information about used third-party software dependencies.  ### Precondition: None.  ### Postcondition: List of the third-party software dependencies used by **DRACOON Core** (referred to as _\"Server\"_) is returned.  ### Further Information: None.  

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ThirdPartyDependenciesData**](ThirdPartyDependenciesData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUploadStatusPublic**
> S3ShareUploadStatus RequestUploadStatusPublic(ctx, accessKey, uploadId)
Request status of S3 file upload

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.15.0</h3>  ### Description: Request status of a S3 file upload.  ### Precondition: An upload channel has been created and the user has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; create</span> permissions in the parent container (room or folder).  ### Postcondition: Status of S3 multipart upload request is returned.  ### Further Information: None.  ### Possible errors: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Http Status | Error Code | Description | | :--- | :--- | :--- | | `400 Bad Request` | `-80000` | Mandatory fields cannot be empty | | `400 Bad Request` | `-80001` | Invalid positive number | | `400 Bad Request` | `-80002` | Invalid number | | `400 Bad Request` | `-40001` | (Target) room is not encrypted | | `400 Bad Request` | `-40755` | Bad file name | | `400 Bad Request` | `-40763` | File key must be set for an upload into encrypted room | | `400 Bad Request` | `-50506` | Exceeds the number of files for this Upload Share | | `403 Forbidden` |  | Access denied | | `404 Not Found` | `-20501` | Upload not found | | `404 Not Found` | `-40000` | Container not found | | `404 Not Found` | `-41000` | Node not found | | `404 Not Found` | `-70501` | User not found | | `409 Conflict` | `-40010` | Container cannot be overwritten | | `409 Conflict` |  | File cannot be overwritten | | `500 Internal Server Error` |  | System Error | | `502 Bad Gateway` |  | S3 Error | | `502 Insufficient Storage` | `-50504` | Exceeds the quota for this Upload Share | | `502 Insufficient Storage` | `-40200` | Exceeds the free node quota in room | | `502 Insufficient Storage` | `-90200` | Exceeds the free customer quota | | `502 Insufficient Storage` | `-90201` | Exceeds the free customer physical disk space |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessKey** | **string**| Access key | 
  **uploadId** | **string**| Upload channel ID | 

### Return type

[**S3ShareUploadStatus**](S3ShareUploadStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFileAsMultipartPublic1**
> ChunkUploadResponse UploadFileAsMultipartPublic1(ctx, accessKey, uploadId, optional)
Upload file

### Description:   Chunked upload of files via Upload Share.  ### Precondition: Valid upload ID.  ### Postcondition: Chunk of file is uploaded.  ### Further Information: Chunked uploads (range requests) are supported.  Following `Content-Types` are supported by this API: * `multipart/form-data` * provided `Content-Type`    For both file upload types set the correct `Content-Type` header and body.    ### Examples:    * `multipart/form-data` ``` POST /api/v4/public/shares/uploads/{access_key}{upload_id} HTTP/1.1  Header: ... Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW ...  Body: ------WebKitFormBoundary7MA4YWxkTrZu0gW Content-Disposition: form-data; name=\"file\"; filename=\"file.txt\" Content-Type: text/plain  Content of file.txt ------WebKitFormBoundary7MA4YWxkTrZu0gW-- ```  * any other `Content-Type`   ``` POST /api/v4/public/shares/uploads/{access_key}{upload_id} HTTP/1.1  Header: ... Content-Type: { ... } ...  Body: raw content ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessKey** | **string**| Access key | 
  **uploadId** | **string**| Upload channel ID | 
 **optional** | ***PublicApiUploadFileAsMultipartPublic1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiUploadFileAsMultipartPublic1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.**|  | 
 **contentRange** | **optional.**| Content-Range   e.g. &#x60;bytes 0-999/3980&#x60; | 
 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 

### Return type

[**ChunkUploadResponse**](ChunkUploadResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

