# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFileUploadByToken**](UploadsApi.md#CancelFileUploadByToken) | **Delete** /v4/uploads/{token} | Cancel file upload
[**CompleteFileUploadByToken**](UploadsApi.md#CompleteFileUploadByToken) | **Put** /v4/uploads/{token} | Complete file upload
[**UploadFileByTokenAsMultipart1**](UploadsApi.md#UploadFileByTokenAsMultipart1) | **Post** /v4/uploads/{token} | Upload file

# **CancelFileUploadByToken**
> CancelFileUploadByToken(ctx, token)
Cancel file upload

### Description: Cancel file upload.  ### Precondition: Valid upload token.  ### Postcondition: Upload canceled, token invalidated and all already transfered chunks removed.  ### Further Information: It is recommended to notify the API about cancelled uploads if possible.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Upload token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompleteFileUploadByToken**
> Node CompleteFileUploadByToken(ctx, token, optional)
Complete file upload

### Description: Finish uploading a file.  ### Precondition: Valid upload token.  ### Postcondition: File created.  ### Further Information: The provided file name might be changed in accordance with the resolution strategy:  * **autorename**: changes the file name and adds a number to avoid conflicts. * **overwrite**: deletes any old file with the same file name. * **fail**: returns an error; in this case, another `PUT` request with a different file name may be sent.  Please ensure that all chunks have been transferred correctly before finishing the upload.  Download share id (if exists) gets changed if: - node with the same name exists in the target container - `resolutionStrategy` is `overwrite` - `keepShareLinks` is `true`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Upload token | 
 **optional** | ***UploadsApiCompleteFileUploadByTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadsApiCompleteFileUploadByTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CompleteUploadRequest**](CompleteUploadRequest.md)|  | 
 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 

### Return type

[**Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFileByTokenAsMultipart1**
> ChunkUploadResponse UploadFileByTokenAsMultipart1(ctx, token, optional)
Upload file

### Description:   Upload a (chunk of a) file.  ### Precondition: Valid upload token.  ### Postcondition: Chunk uploaded.  ### Further Information: Range requests are supported.    Following `Content-Types` are supported by this API: * `multipart/form-data` * provided `Content-Type`  For both file upload types set the correct `Content-Type` header and body.    ### Examples:    * `multipart/form-data` ``` POST /api/v4/uploads/{token} HTTP/1.1  Header: ... Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW ...  Body: ------WebKitFormBoundary7MA4YWxkTrZu0gW Content-Disposition: form-data; name=\"file\"; filename=\"file.txt\" Content-Type: text/plain  Content of file.txt ------WebKitFormBoundary7MA4YWxkTrZu0gW-- ```  * any other `Content-Type`  ``` POST /api/v4/uploads/{token} HTTP/1.1  Header: ... Content-Type: { ... } ...  Body: raw content ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Upload token | 
 **optional** | ***UploadsApiUploadFileByTokenAsMultipart1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadsApiUploadFileByTokenAsMultipart1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 
 **contentRange** | **optional.**| Content-Range   e.g. &#x60;bytes 0-999/3980&#x60; | 

### Return type

[**ChunkUploadResponse**](ChunkUploadResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

