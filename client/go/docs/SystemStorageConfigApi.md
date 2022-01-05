# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateS3Config**](SystemStorageConfigApi.md#CreateS3Config) | **Post** /v4/system/config/storage/s3 | Create S3 storage configuration
[**CreateS3Tag**](SystemStorageConfigApi.md#CreateS3Tag) | **Post** /v4/system/config/storage/s3/tags | Create S3 tag
[**RemoveS3Tag**](SystemStorageConfigApi.md#RemoveS3Tag) | **Delete** /v4/system/config/storage/s3/tags/{id} | Remove S3 tag
[**Request3Config**](SystemStorageConfigApi.md#Request3Config) | **Get** /v4/system/config/storage/s3 | Request S3 storage configuration
[**RequestS3Tag**](SystemStorageConfigApi.md#RequestS3Tag) | **Get** /v4/system/config/storage/s3/tags/{id} | Request S3 tag
[**RequestS3TagList**](SystemStorageConfigApi.md#RequestS3TagList) | **Get** /v4/system/config/storage/s3/tags | Request list of configured S3 tags
[**UpdateS3Config**](SystemStorageConfigApi.md#UpdateS3Config) | **Put** /v4/system/config/storage/s3 | Update S3 storage configuration

# **CreateS3Config**
> S3Config CreateS3Config(ctx, body, optional)
Create S3 storage configuration

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.3.0</h3>  ### Description:   Create new S3 configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: New S3 configuration is created.  ### Further Information: Forbidden characters in bucket names: [`.`]   `bucketName` and `endpointUrl` are deprecated, use `bucketUrl` instead.  ### Virtual hosted style access  Example: https://<span style=\"color:red;\">bucket-name</span>.s3.<span style=\"color:red;\">region</span>.amazonaws.com/<span style=\"color:red;\">key-name</span> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**S3ConfigCreateRequest**](S3ConfigCreateRequest.md)|  | 
 **optional** | ***SystemStorageConfigApiCreateS3ConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemStorageConfigApiCreateS3ConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**S3Config**](S3Config.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateS3Tag**
> S3Tag CreateS3Tag(ctx, body, optional)
Create S3 tag

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.9.0</h3>  ### Description:   Create new S3 tag.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: New S3 tag is created.  ### Further Information: * Maximum key length: **128** characters.   * Maximum value length: **256** characters.   * Both S3 tag key and value are **case-sensitive** strings.   * Maximum of **20 mandatory S3 tags** is allowed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**S3TagCreateRequest**](S3TagCreateRequest.md)|  | 
 **optional** | ***SystemStorageConfigApiCreateS3TagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemStorageConfigApiCreateS3TagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**S3Tag**](S3Tag.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveS3Tag**
> RemoveS3Tag(ctx, id, optional)
Remove S3 tag

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.9.0</h3>  ### Description:   Delete S3 tag.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: S3 tag gets deleted.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| S3 tag ID | 
 **optional** | ***SystemStorageConfigApiRemoveS3TagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemStorageConfigApiRemoveS3TagOpts struct
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

# **Request3Config**
> S3Config Request3Config(ctx, optional)
Request S3 storage configuration

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.3.0</h3>  ### Description:   Retrieve S3 configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: S3 configuration is returned.  ### Further Information: None.  ### Virtual hosted style access  Example: https://<span style=\"color:red;\">bucket-name</span>.s3.<span style=\"color:red;\">region</span>.amazonaws.com/<span style=\"color:red;\">key-name</span> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemStorageConfigApiRequest3ConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemStorageConfigApiRequest3ConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**S3Config**](S3Config.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestS3Tag**
> S3Tag RequestS3Tag(ctx, id, optional)
Request S3 tag

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.9.0</h3>  ### Description:   Retrieve single S3 tag.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: S3 tag is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| S3 tag ID | 
 **optional** | ***SystemStorageConfigApiRequestS3TagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemStorageConfigApiRequestS3TagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**S3Tag**](S3Tag.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestS3TagList**
> S3TagList RequestS3TagList(ctx, optional)
Request list of configured S3 tags

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.9.0</h3>  ### Description:   Retrieve all configured S3 tags.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: S3 tags are returned.  ### Further Information: An empty list is returned if no S3 tags are found / configured.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemStorageConfigApiRequestS3TagListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemStorageConfigApiRequestS3TagListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**S3TagList**](S3TagList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateS3Config**
> S3Config UpdateS3Config(ctx, body, optional)
Update S3 storage configuration

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.3.0</h3>  ### Description:   Update existing S3 configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: S3 configuration is updated.  ### Further Information: Forbidden characters in bucket names: [`.`]   `bucketName` and `endpointUrl` are deprecated, use `bucketUrl` instead.  ### Virtual hosted style access  Example: https://<span style=\"color:red;\">bucket-name</span>.s3.<span style=\"color:red;\">region</span>.amazonaws.com/<span style=\"color:red;\">key-name</span> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**S3ConfigUpdateRequest**](S3ConfigUpdateRequest.md)|  | 
 **optional** | ***SystemStorageConfigApiUpdateS3ConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemStorageConfigApiUpdateS3ConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**S3Config**](S3Config.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

