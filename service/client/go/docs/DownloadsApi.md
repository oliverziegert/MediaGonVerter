# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadAvatar**](DownloadsApi.md#DownloadAvatar) | **Get** /v4/downloads/avatar/{user_id}/{uuid} | Download avatar
[**DownloadFileViaToken**](DownloadsApi.md#DownloadFileViaToken) | **Head** /v4/downloads/{token} | Download file
[**DownloadFileViaToken1**](DownloadsApi.md#DownloadFileViaToken1) | **Get** /v4/downloads/{token} | Download file
[**DownloadZipArchiveViaToken**](DownloadsApi.md#DownloadZipArchiveViaToken) | **Get** /v4/downloads/zip/{token} | Download ZIP archive

# **DownloadAvatar**
> string DownloadAvatar(ctx, userId, uuid)
Download avatar

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.11.0</h3>  ### Description: Download avatar for given user ID and UUID.  ### Precondition: Valid UUID.  ### Postcondition: Stream is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| User ID | 
  **uuid** | **string**| UUID of the avatar | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFileViaToken**
> DownloadFileViaToken(ctx, token, optional)
Download file

### Description: Download a file.  ### Precondition: Valid download token.  ### Postcondition: Stream is returned.  ### Further Information: Range requests are supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Download token | 
 **optional** | ***DownloadsApiDownloadFileViaTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DownloadsApiDownloadFileViaTokenOpts struct
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

# **DownloadFileViaToken1**
> DownloadFileViaToken1(ctx, token, optional)
Download file

### Description: Download a file.  ### Precondition: Valid download token.  ### Postcondition: Stream is returned.  ### Further Information: Range requests are supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Download token | 
 **optional** | ***DownloadsApiDownloadFileViaToken1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DownloadsApiDownloadFileViaToken1Opts struct
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

# **DownloadZipArchiveViaToken**
> DownloadZipArchiveViaToken(ctx, token)
Download ZIP archive

### Description: Download multiple files in a ZIP archive.  ### Precondition: Valid download token.  ### Postcondition: Stream is returned.  ### Further Information: Create a download token with `POST /nodes/zip` API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Download token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

