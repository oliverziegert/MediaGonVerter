# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdConfig**](SystemAuthConfigApi.md#CreateAdConfig) | **Post** /v4/system/config/auth/ads | Create Active Directory configuration
[**CreateOAuthClient**](SystemAuthConfigApi.md#CreateOAuthClient) | **Post** /v4/system/config/oauth/clients | Create OAuth client
[**CreateOpenIdIdpConfig**](SystemAuthConfigApi.md#CreateOpenIdIdpConfig) | **Post** /v4/system/config/auth/openid/idps | Create OpenID Connect IDP configuration
[**CreateRadiusConfig**](SystemAuthConfigApi.md#CreateRadiusConfig) | **Post** /v4/system/config/auth/radius | Create RADIUS configuration
[**RemoveAdConfig**](SystemAuthConfigApi.md#RemoveAdConfig) | **Delete** /v4/system/config/auth/ads/{ad_id} | Remove Active Directory configuration
[**RemoveOAuthClient**](SystemAuthConfigApi.md#RemoveOAuthClient) | **Delete** /v4/system/config/oauth/clients/{client_id} | Remove OAuth client
[**RemoveOpenIdIdpConfig**](SystemAuthConfigApi.md#RemoveOpenIdIdpConfig) | **Delete** /v4/system/config/auth/openid/idps/{idp_id} | Remove OpenID Connect IDP configuration
[**RemoveRadiusConfig**](SystemAuthConfigApi.md#RemoveRadiusConfig) | **Delete** /v4/system/config/auth/radius | Remove RADIUS configuration
[**RequestAdConfig**](SystemAuthConfigApi.md#RequestAdConfig) | **Get** /v4/system/config/auth/ads/{ad_id} | Request Active Directory configuration
[**RequestAdConfigs**](SystemAuthConfigApi.md#RequestAdConfigs) | **Get** /v4/system/config/auth/ads | Request list of Active Directory configurations
[**RequestOAuthClient**](SystemAuthConfigApi.md#RequestOAuthClient) | **Get** /v4/system/config/oauth/clients/{client_id} | Request OAuth client
[**RequestOAuthClients**](SystemAuthConfigApi.md#RequestOAuthClients) | **Get** /v4/system/config/oauth/clients | Request list of OAuth clients
[**RequestOpenIdIdpConfig**](SystemAuthConfigApi.md#RequestOpenIdIdpConfig) | **Get** /v4/system/config/auth/openid/idps/{idp_id} | Request OpenID Connect IDP configuration
[**RequestOpenIdIdpConfigs**](SystemAuthConfigApi.md#RequestOpenIdIdpConfigs) | **Get** /v4/system/config/auth/openid/idps | Request list of OpenID Connect IDP configurations
[**RequestRadiusConfig**](SystemAuthConfigApi.md#RequestRadiusConfig) | **Get** /v4/system/config/auth/radius | Request RADIUS configuration
[**TestAdConfig**](SystemAuthConfigApi.md#TestAdConfig) | **Post** /v4/system/config/actions/test/ad | Test Active Directory configuration
[**TestRadiusConfig**](SystemAuthConfigApi.md#TestRadiusConfig) | **Post** /v4/system/config/actions/test/radius | Test RADIUS server availability
[**UpdateAdConfig**](SystemAuthConfigApi.md#UpdateAdConfig) | **Put** /v4/system/config/auth/ads/{ad_id} | Update Active Directory configuration
[**UpdateOAuthClient**](SystemAuthConfigApi.md#UpdateOAuthClient) | **Put** /v4/system/config/oauth/clients/{client_id} | Update OAuth client
[**UpdateOpenIdIdpConfig**](SystemAuthConfigApi.md#UpdateOpenIdIdpConfig) | **Put** /v4/system/config/auth/openid/idps/{idp_id} | Update OpenID Connect IDP configuration
[**UpdateRadiusConfig**](SystemAuthConfigApi.md#UpdateRadiusConfig) | **Put** /v4/system/config/auth/radius | Update RADIUS configuration

# **CreateAdConfig**
> ActiveDirectoryConfig CreateAdConfig(ctx, body, optional)
Create Active Directory configuration

### Description: Create a new Active Directory configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: New Active Directory configuration created.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateActiveDirectoryConfigRequest**](CreateActiveDirectoryConfigRequest.md)|  | 
 **optional** | ***SystemAuthConfigApiCreateAdConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiCreateAdConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**ActiveDirectoryConfig**](ActiveDirectoryConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOAuthClient**
> OAuthClient CreateOAuthClient(ctx, body, optional)
Create OAuth client

### Description: Create a new OAuth client.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: New OAuth client created.  ### Further Information:   Client secret **MUST** have:   * at least 12 characters, at most 32 characters   * only lower case characters, upper case characters and digits   * at least 1 lower case character, 1 upper case character and 1 digit    The client secret is optional and will be generated if it is left empty.    Valid grant types are:   * `authorization_code`   * `implicit`   * `password`   * `client_credentials`   * `refresh_token`    Grant type `client_credentials` is currently **NOT** permitted!  Allowed characters for client ID are: `[a-zA-Z0-9_-]`  If grant types `authorization_code` or `implicit` are used, a redirect URI **MUST** be provided!  Default access token validity: **8 hours**   Default refresh token validity: **30 days** Default approval validity: **½ year**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOAuthClientRequest**](CreateOAuthClientRequest.md)|  | 
 **optional** | ***SystemAuthConfigApiCreateOAuthClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiCreateOAuthClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOpenIdIdpConfig**
> OpenIdIdpConfig CreateOpenIdIdpConfig(ctx, body, optional)
Create OpenID Connect IDP configuration

### Description: Create new OpenID Connect IDP configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: New OpenID Connect IDP configuration is created.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOpenIdIdpConfigRequest**](CreateOpenIdIdpConfigRequest.md)|  | 
 **optional** | ***SystemAuthConfigApiCreateOpenIdIdpConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiCreateOpenIdIdpConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**OpenIdIdpConfig**](OpenIdIdpConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRadiusConfig**
> RadiusConfig CreateRadiusConfig(ctx, body, optional)
Create RADIUS configuration

### Description:   Create new RADIUS configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: New RADIUS configuration is created.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RadiusConfigCreateRequest**](RadiusConfigCreateRequest.md)|  | 
 **optional** | ***SystemAuthConfigApiCreateRadiusConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiCreateRadiusConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**RadiusConfig**](RadiusConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAdConfig**
> RemoveAdConfig(ctx, adId, optional)
Remove Active Directory configuration

### Description: Delete an existing Active Directory configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: Active Directory configuration is removed.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **adId** | **int32**| Active Directory ID | 
 **optional** | ***SystemAuthConfigApiRemoveAdConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRemoveAdConfigOpts struct
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

# **RemoveOAuthClient**
> RemoveOAuthClient(ctx, clientId, optional)
Remove OAuth client

### Description: Delete an existing OAuth client.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: OAuth client is removed.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| OAuth client ID | 
 **optional** | ***SystemAuthConfigApiRemoveOAuthClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRemoveOAuthClientOpts struct
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

# **RemoveOpenIdIdpConfig**
> RemoveOpenIdIdpConfig(ctx, idpId, optional)
Remove OpenID Connect IDP configuration

### Description: Delete an existing OpenID Connect IDP configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: OpenID Connect IDP configuration is removed.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **int32**| OpenID Connect IDP configuration ID | 
 **optional** | ***SystemAuthConfigApiRemoveOpenIdIdpConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRemoveOpenIdIdpConfigOpts struct
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

# **RemoveRadiusConfig**
> RemoveRadiusConfig(ctx, optional)
Remove RADIUS configuration

### Description:   Delete existing RADIUS configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: RADIUS configuration is deleted.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAuthConfigApiRemoveRadiusConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRemoveRadiusConfigOpts struct
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

# **RequestAdConfig**
> ActiveDirectoryConfig RequestAdConfig(ctx, adId, optional)
Request Active Directory configuration

### Description:   Retrieve the configuration of an Active Directory.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: Active Directory configuration is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **adId** | **int32**| Active Directory ID | 
 **optional** | ***SystemAuthConfigApiRequestAdConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRequestAdConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**ActiveDirectoryConfig**](ActiveDirectoryConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestAdConfigs**
> ActiveDirectoryConfigList RequestAdConfigs(ctx, optional)
Request list of Active Directory configurations

### Description:   Retrieve a list of configured Active Directories.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: List of Active Directory configurations is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAuthConfigApiRequestAdConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRequestAdConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**ActiveDirectoryConfigList**](ActiveDirectoryConfigList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOAuthClient**
> OAuthClient RequestOAuthClient(ctx, clientId, optional)
Request OAuth client

### Description:   Retrieve the configuration of an OAuth client.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: OAuth client is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| OAuth client ID | 
 **optional** | ***SystemAuthConfigApiRequestOAuthClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRequestOAuthClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOAuthClients**
> []OAuthClient RequestOAuthClients(ctx, optional)
Request list of OAuth clients

### Description:   Retrieve a list of configured OAuth clients.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: List of OAuth clients is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `isStandard:eq:true`   Get standard OAuth clients.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `isStandard` | Standard client filter | `eq` |  | `true or false` | | `isExternal` | External client filter | `eq` |  | `true or false` | | `isEnabled` | Enabled/disabled clients filter | `eq` |  | `true or false` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `clientName:desc|isStandard:asc`   Sort by `clientName` descending **AND** `isStandard` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `clientName` | Client name | | `isStandard` | Is a standard client | | `isExternal` | Is a external client | | `isEnabled` | Is a enabled client |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAuthConfigApiRequestOAuthClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRequestOAuthClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**[]OAuthClient**](OAuthClient.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOpenIdIdpConfig**
> OpenIdIdpConfig RequestOpenIdIdpConfig(ctx, idpId, optional)
Request OpenID Connect IDP configuration

### Description:   Retrieve an OpenID Connect IDP configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: OpenID Connect IDP configuration is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idpId** | **int32**| OpenID Connect IDP configuration ID | 
 **optional** | ***SystemAuthConfigApiRequestOpenIdIdpConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRequestOpenIdIdpConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**OpenIdIdpConfig**](OpenIdIdpConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestOpenIdIdpConfigs**
> []OpenIdIdpConfig RequestOpenIdIdpConfigs(ctx, optional)
Request list of OpenID Connect IDP configurations

### Description:   Retrieve a list of configured OpenID Connect IDPs.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: List of OpenID Connect IDP configurations is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAuthConfigApiRequestOpenIdIdpConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRequestOpenIdIdpConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**[]OpenIdIdpConfig**](OpenIdIdpConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRadiusConfig**
> RadiusConfig RequestRadiusConfig(ctx, optional)
Request RADIUS configuration

### Description:   Retrieve a RADIUS configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: RADIUS configuration is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAuthConfigApiRequestRadiusConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiRequestRadiusConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**RadiusConfig**](RadiusConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestAdConfig**
> TestActiveDirectoryConfigResponse TestAdConfig(ctx, body, optional)
Test Active Directory configuration

### Description:   Test Active Directory configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: Active Directory configuration is returned if successful.  ### Further Information: DRACOON tries to establish a connection with the provided information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestActiveDirectoryConfigRequest**](TestActiveDirectoryConfigRequest.md)|  | 
 **optional** | ***SystemAuthConfigApiTestAdConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiTestAdConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**TestActiveDirectoryConfigResponse**](TestActiveDirectoryConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestRadiusConfig**
> TestRadiusConfig(ctx, optional)
Test RADIUS server availability

### Description:   Test RADIUS configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: RADIUS configuration is returned if successful.  ### Further Information: DRACOON tries to establish a connection with the provided information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAuthConfigApiTestRadiusConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiTestRadiusConfigOpts struct
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

# **UpdateAdConfig**
> ActiveDirectoryConfig UpdateAdConfig(ctx, body, adId, optional)
Update Active Directory configuration

### Description:   Update an existing Active Directory configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: Active Directory configuration updated.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateActiveDirectoryConfigRequest**](UpdateActiveDirectoryConfigRequest.md)|  | 
  **adId** | **int32**| Active Directory ID | 
 **optional** | ***SystemAuthConfigApiUpdateAdConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiUpdateAdConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**ActiveDirectoryConfig**](ActiveDirectoryConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOAuthClient**
> OAuthClient UpdateOAuthClient(ctx, body, clientId, optional)
Update OAuth client

### Description:   Update an existing OAuth client.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: OAuth client updated.  ### Further Information:   Client secret **MUST** have:   * at least 12 characters, at most 32 characters   * only lower case characters, upper case characters and digits   * at least 1 lower case character, 1 upper case character and 1 digit    The client secret is optional and will be generated if it is left empty.    Valid grant types are:   * `authorization_code`   * `implicit`   * `password`   * `client_credentials`   * `refresh_token`    Grant type `client_credentials` is currently **NOT** permitted!  If grant types `authorization_code` or `implicit` are used, a redirect URI **MUST** be provided! 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOAuthClientRequest**](UpdateOAuthClientRequest.md)|  | 
  **clientId** | **string**| OAuth client ID | 
 **optional** | ***SystemAuthConfigApiUpdateOAuthClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiUpdateOAuthClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOpenIdIdpConfig**
> OpenIdIdpConfig UpdateOpenIdIdpConfig(ctx, body, idpId, optional)
Update OpenID Connect IDP configuration

### Description:   Update an existing OpenID Connect IDP configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: OpenID Connect IDP configuration is updated.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOpenIdIdpConfigRequest**](UpdateOpenIdIdpConfigRequest.md)|  | 
  **idpId** | **int32**| OpenID Connect IDP configuration ID | 
 **optional** | ***SystemAuthConfigApiUpdateOpenIdIdpConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiUpdateOpenIdIdpConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**OpenIdIdpConfig**](OpenIdIdpConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRadiusConfig**
> RadiusConfig UpdateRadiusConfig(ctx, body, optional)
Update RADIUS configuration

### Description:   Update existing RADIUS configuration.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: RADIUS configuration is updated.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RadiusConfigUpdateRequest**](RadiusConfigUpdateRequest.md)|  | 
 **optional** | ***SystemAuthConfigApiUpdateRadiusConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAuthConfigApiUpdateRadiusConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**RadiusConfig**](RadiusConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

