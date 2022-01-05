# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomer**](ProvisioningApi.md#CreateCustomer) | **Post** /v4/provisioning/customers | Create customer
[**CreateTenantWebhook**](ProvisioningApi.md#CreateTenantWebhook) | **Post** /v4/provisioning/webhooks | Create tenant webhook
[**RemoveCustomer**](ProvisioningApi.md#RemoveCustomer) | **Delete** /v4/provisioning/customers/{customer_id} | Remove customer
[**RemoveCustomerAttribute**](ProvisioningApi.md#RemoveCustomerAttribute) | **Delete** /v4/provisioning/customers/{customer_id}/customerAttributes/{key} | Remove customer attribute
[**RemoveTenantWebhook**](ProvisioningApi.md#RemoveTenantWebhook) | **Delete** /v4/provisioning/webhooks/{webhook_id} | Remove tenant webhook
[**RequestCustomer**](ProvisioningApi.md#RequestCustomer) | **Get** /v4/provisioning/customers/{customer_id} | Get customer
[**RequestCustomerAttributes**](ProvisioningApi.md#RequestCustomerAttributes) | **Get** /v4/provisioning/customers/{customer_id}/customerAttributes | Request customer attributes
[**RequestCustomerUsers**](ProvisioningApi.md#RequestCustomerUsers) | **Get** /v4/provisioning/customers/{customer_id}/users | Request list of customer users
[**RequestCustomers**](ProvisioningApi.md#RequestCustomers) | **Get** /v4/provisioning/customers | Request list of customers
[**RequestListOfEventTypesForTenant**](ProvisioningApi.md#RequestListOfEventTypesForTenant) | **Get** /v4/provisioning/webhooks/event_types | Request list of event types
[**RequestListOfTenantWebhooks**](ProvisioningApi.md#RequestListOfTenantWebhooks) | **Get** /v4/provisioning/webhooks | Request list of tenant webhooks
[**RequestTenantWebhook**](ProvisioningApi.md#RequestTenantWebhook) | **Get** /v4/provisioning/webhooks/{webhook_id} | Request tenant webhook
[**ResetTenantWebhookLifetime**](ProvisioningApi.md#ResetTenantWebhookLifetime) | **Post** /v4/provisioning/webhooks/{webhook_id}/reset_lifetime | Reset tenant webhook lifetime
[**SetCustomerAttributes**](ProvisioningApi.md#SetCustomerAttributes) | **Post** /v4/provisioning/customers/{customer_id}/customerAttributes | Set customer attributes
[**UpdateCustomer**](ProvisioningApi.md#UpdateCustomer) | **Put** /v4/provisioning/customers/{customer_id} | Update customer
[**UpdateCustomerAttributes**](ProvisioningApi.md#UpdateCustomerAttributes) | **Put** /v4/provisioning/customers/{customer_id}/customerAttributes | Add or edit customer attributes
[**UpdateTenantWebhook**](ProvisioningApi.md#UpdateTenantWebhook) | **Put** /v4/provisioning/webhooks/{webhook_id} | Update tenant webhook

# **CreateCustomer**
> NewCustomerResponse CreateCustomer(ctx, body, optional)
Create customer

### Description: Create a new customer.  ### Precondition: Authentication with `X-Sds-Service-Token` required.    ### Postcondition: A new customer is created.  ### Further Information: If no company name is set, first letter of the first name separated by dot following by last name of the first administrator is used (e.g. `J.Doe`).   Max quota has to be at least `1 MB` (= `1.048.576 B`).  If `basic` authentication is enabled, the first administrator will get `basic` authentication by default.   To create a first administrator without `basic` authentication it **MUST** be disabled explicitly.    Forbidden characters in passwords: [`&`, `'`, `<`, `>`]  ### Authentication Method Options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Authentication Method | Option Key | Option Value | | :--- | :--- | :--- | | `basic` / `sql` | `username` | Unique user identifier | | `active_directory` | `ad_config_id` (optional) | Active Directory configuration ID | |  | `username` | Active Directory username according to authentication setting `userFilter` | | `radius` | `username` | RADIUS username | | `openid` | `openid_config_id` (optional) | OpenID Connect configuration ID | |  | `username` | OpenID Connect username according to authentication setting `mappingClaim` |  </details> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NewCustomerRequest**](NewCustomerRequest.md)|  | 
 **optional** | ***ProvisioningApiCreateCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiCreateCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsServiceToken** | **optional.**| Service Authentication token | 

### Return type

[**NewCustomerResponse**](NewCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTenantWebhook**
> Webhook CreateTenantWebhook(ctx, body, optional)
Create tenant webhook

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Create a new webhook for the tenant scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage webhook</span> required.  ### Postcondition: Webhook is created for given event types.  ### Further Information: URL must begin with the `HTTPS` scheme. Webhook names are limited to 150 characters.  ### Available event types: <details open style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Name | Description | Scope | | :--- | :--- | :--- | | **`customer.created`** | Triggered when a new customer is created | Tenant Webhook | | **`customer.deleted`** | Triggered when a user is deleted | Tenant Webhook | | **`webhook.expiring`** | Triggered 30/20/10/1 days before a webhook expires |  Tenant Webhook |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateWebhookRequest**](CreateWebhookRequest.md)|  | 
 **optional** | ***ProvisioningApiCreateTenantWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiCreateTenantWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsServiceToken** | **optional.**| Service Authentication token | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveCustomer**
> RemoveCustomer(ctx, customerId, optional)
Remove customer

### Description: Delete a customer.  ### Precondition: Authentication with `X-Sds-Service-Token` required.  ### Postcondition: Customer is deleted.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **int64**| Customer ID | 
 **optional** | ***ProvisioningApiRemoveCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRemoveCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveCustomerAttribute**
> RemoveCustomerAttribute(ctx, customerId, key, optional)
Remove customer attribute

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.4.0</h3>  ### Description: Delete a custom customer attribute.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> required.  ### Postcondition: Custom customer attribute gets deleted.  ### Further Information: * Allowed characters for keys are: `[a-zA-Z0-9_-]`   * Characters are **case-insensitive**.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **int64**| Customer ID | 
  **key** | **string**| Key | 
 **optional** | ***ProvisioningApiRemoveCustomerAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRemoveCustomerAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTenantWebhook**
> RemoveTenantWebhook(ctx, webhookId, optional)
Remove tenant webhook

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Delete a webhook for the tenant scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage webhook</span> required.  ### Postcondition: Webhook is deleted.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***ProvisioningApiRemoveTenantWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRemoveTenantWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestCustomer**
> Customer RequestCustomer(ctx, customerId, optional)
Get customer

### Description:   Receive details of a selected customer.  ### Precondition: Authentication with `X-Sds-Service-Token` required.  ### Postcondition: Customer details are returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **int64**| Customer ID | 
 **optional** | ***ProvisioningApiRequestCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRequestCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **includeAttributes** | **optional.Bool**| Include custom customer attributes. | 
 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestCustomerAttributes**
> AttributesResponse RequestCustomerAttributes(ctx, customerId, optional)
Request customer attributes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.4.0</h3>  ### Description:   Retrieve a list of customer attributes.  ### Precondition: Authentication with `X-Sds-Service-Token` required.   Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read all customers</span> required.  ### Postcondition: List of attributes is returned.  ### Further Information:  ### Filtering: Filters are case insensitive.   All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `key:cn:searchString_1|value:cn:searchString_2`   Filter by attribute key contains `searchString_1` **AND** attribute value contains `searchString_2`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `key` | Customer attribute key filter | `cn, eq, sw` | Attribute key contains / equals / starts with value. | `search String` | | `value` | Customer attribute value filter | `cn, eq, sw` | Attribute value contains / equals / starts with value. | `search String` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `key:asc|value:desc`   Sort by `key` ascending **AND** by `value` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `key` | Customer attribute key | | `value` | Customer attribute value |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **int64**| Customer ID | 
 **optional** | ***ProvisioningApiRequestCustomerAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRequestCustomerAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

[**AttributesResponse**](AttributesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestCustomerUsers**
> UserList RequestCustomerUsers(ctx, customerId, optional)
Request list of customer users

### Description:   Receive a list of users associated with a certain customer.  ### Precondition: Authentication with `X-Sds-Service-Token` required.  ### Postcondition: List of customer users is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Except for `login`, `firstName` and  `lastName` - these are connected via logical disjunction (**OR**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `login:cn:searchString_1|firstName:cn:searchString_2|lockStatus:eq:2`   Filter users by login contains `searchString_1` **OR** firstName contains `searchString_2` **AND** those who are **NOT** locked.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `email` | Email filter | `eq`, `cn` | Email contains value. | `search String` | | `userName` | User name filter | `eq`, `cn` | UserName contains value. | `search String` | | `firstName` | User first name filter | `cn` | User first name contains value. | `search String` | | `lastName` | User last name filter | `cn` | User last name contains value. | `search String` | | `isLocked` | User lock status filter | `eq` |  | `true or false` | | `effectiveRoles` | Filter users with DIRECT or DIRECT **AND** EFFECTIVE roles<ul><li>`false`: DIRECT roles</li><li>`true`: DIRECT **AND** EFFECTIVE roles</li></ul>DIRECT means: e.g. user gets role **directly** granted from someone with _grant permission_ right.<br>EFFECTIVE means: e.g. user gets role through **group membership**. | `eq` |  | `true or false`<br>default: `false` | | `createdAt` | Creation date filter | `ge, le` | Creation date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `createdAt:ge:2016-12-31`&#124;`createdAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `phone` | Phone filter | `eq` | Phone equals value. | `search String` | | `isEncryptionEnabled` | Encryption status filter<ul><li>client-side encryption</li><li>private key possession</li></ul> | `eq` |  | `true or false` | | `hasRole` | (**`NEW`**) User role filter<br>Depends on **effectiveRoles**.<br>For more information about roles check **`GET /roles`** API | `eq` | User role equals value. | <ul><li>`CONFIG_MANAGER` - Manages global configuration</li><li>`USER_MANAGER` - Manages users</li><li>`GROUP_MANAGER` - Manages user groups</li><li>`ROOM_MANAGER` - Manages top level rooms</li><li>`LOG_AUDITOR` - Reads audit logs</li><li>`NONMEMBER_VIEWER` - Views users and groups when having room _\"manage\"_ permission</li></ul> |  </details>  ### Deprecated filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | <del>`lockStatus`</del> | User lock status filter | `eq` | User lock status equals value. | <ul><li>`0` - Locked</li><li>`1` - Web access allowed</li><li>`2` - Web and mobile access allowed</li></ul> | | <del>`login`</del> |  User login filter | `cn` | User login contains value. | `search String` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `firstName:asc|lastLoginSuccessAt:desc`   Sort by `firstName` ascending **AND** by `lastLoginSuccessAt` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `userName` | User name | | `email` | User email | | `firstName` | User first name | | `lastName` | User last name | | `isLocked` | User lock status | | `lastLoginSuccessAt` | Last successful login date | | `expireAt` | Expiration date | | `createdAt` | Creation date |  </details>  ### Deprecated sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | <del>`gender`</del> | Gender | | <del>`lockStatus`</del> | User lock status | | <del>`login`</del> | User login |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **int64**| Customer ID | 
 **optional** | ***ProvisioningApiRequestCustomerUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRequestCustomerUsersOpts struct
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
 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

[**UserList**](UserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestCustomers**
> CustomerList RequestCustomers(ctx, optional)
Request list of customers

### Description:   Receive a list of customers.  ### Precondition: Authentication with `X-Sds-Service-Token` required.  ### Postcondition: List of customers is returned.  ### Further Information: This list returns a maximum of **1000** entries.    ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `trialDaysLeft:le:10|userMax:le:100`   Get all customers with `10` trial days left **AND** user maximum **<=** `100`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `id` | Customer ID filter | `eq` | Customer ID equals value. | `positive Integer` | | `companyName` | Company name filter | `cn` | Company name contains value. | `search String` | | `customerContractType` | Customer contract type filter | `eq` | Customer contract type equals value. | <ul><li>`demo`</li><li>`free`</li><li>`pay`</li></ul> | | `trialDaysLeft` | Left trial days filter | `ge, le` | Left trial days are greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `trialDaysLeft:ge:5`&#124;`trialDaysLeft:le:10` | | `providerCustomerId` | Provider Customer ID filter | `cn, eq` | Provider Customer ID contains / equals value. | `search String` | | `quotaMax` | Maximum quota filter | `ge, le` | Maximum quota is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `quotaMax:ge:1024`&#124;`quotaMax:le:1073741824` | `positive Integer` | | `quotaUsed` | Used quota filter | `ge, le` | Used quota is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `quotaUsed:ge:1024`&#124;`quotaUsed:le:1073741824` | `positive Integer` | | `userMax` | User maximum filter | `ge, le` | User maxiumum is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `userMax:ge:10`&#124;`userMax:le:100` | `positive Integer` | | `userUsed` | Number of registered users filter | `ge, le` | Number of registered users is is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `userUsed:ge:10`&#124;`userUsed:le:100` | `positive Integer` | | `isLocked` | Lock status filter | `eq` |  | `true or false` | | `createdAt` | Creation date filter | `ge, le` | Creation date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `createdAt:ge:2016-12-31`&#124;`createdAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `updatedAt` | Last modification date filter | `ge, le` | Last modification date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `updatedAt:ge:2016-12-31`&#124;`updatedAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `lastLoginAt` | Last login date filter | `ge, le` | Last login date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `lastLoginAt:ge:2016-12-31`&#124;`lastLoginAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `userLogin` | User login filter | `eq` | User login name equals value.<br>Search user all logins e.g. `basic`, `active_directory`, `radius`. | `search String` | | `attributeKey` | Customer attribute key filter | `eq`, `nex` | Customer attribute key equals value / Customer attribute does **NOT** exist at customer | `search String` | | `attributeValue` | Customer attribute value filter | `eq` | Customer attribute value equals value. | `search String` |  </details>  ### Deprecated filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | <del>`activationCode`</del> | Activation code filter | `cn, eq` | Activation code contains / equals value. | `search String` | | <del>`lockStatus`</del> | Lock status filter | `eq` |  | <ul><li>`0` - unlocked</li><li>`1` - locked</li></ul> |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `companyName:desc|userUsed:asc`   Sort by `companyName` descending **AND** `userUsed` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `companyName` | Company name | | `customerContractType` | Customer contract type | | `trialDaysLeft` | Number of remaining trial days (demo customers) | | `providerCustomerId` | Provider Customer ID | | `quotaMax` | Maximum quota | | `quotaUsed` | Currently used quota | | `userMax` | Maximum user number | | `userUsed` | Number of registered users | | `isLocked` | Lock status of customer | | `createdAt` | Creation date | | `updatedAt` | Last modification date | | `lastLoginAt` | Last login date of any user of this customer |  </details>  ### Deprecated sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | <del>`lockStatus`</del> | Lock status of customer |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisioningApiRequestCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRequestCustomersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **includeAttributes** | **optional.Bool**| Include custom customer attributes. | 
 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

[**CustomerList**](CustomerList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestListOfEventTypesForTenant**
> EventTypeList RequestListOfEventTypesForTenant(ctx, optional)
Request list of event types

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Get a list of available event types.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage webhook</span> required.  ### Postcondition: List of available event types is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisioningApiRequestListOfEventTypesForTenantOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRequestListOfEventTypesForTenantOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

[**EventTypeList**](EventTypeList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestListOfTenantWebhooks**
> WebhookList RequestListOfTenantWebhooks(ctx, optional)
Request list of tenant webhooks

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Get a list of webhooks for the tenant scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage webhook</span> required.  ### Postcondition: List of webhooks is returned.  ### Further Information:   Output is limited to **500** entries.   For more results please use filter criteria and paging (`offset` + `limit`).   `EncryptionInfo` is **NOT** provided.   Wildcard character is the asterisk character: **`*`**  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:cn:goo|createdAt:ge:2015-01-01`   Get webhooks where name contains `goo` **AND** webhook creation date is **>=** `2015-01-01`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | **`id`** | Webhook id filter | `eq` | Webhook id equals value.<br>Multiple values are allowed and will be connected via logical disjunction (**OR**). |`positive number`| | **`name`** | Webhook type name| `cn, eq` | Webhook name contains / equals value. | `search String` | | **`isEnabled`** | Webhook isEnabled filter | `eq` |  | `true or false` | | **`createdAt`** | Creation date filter | `ge, le` | Creation date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `createdAt:ge:2016-12-31`&#124;`createdAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | **`updatedAt`** | Last modification date filter | `ge, le` | Last modification date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `updatedAt:ge:2016-12-31`&#124;`updatedAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | **`expiration`** | Expiration date filter | `ge, le, eq` | Expiration date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `expiration:ge:2016-12-31`&#124;`expiration:le:2018-01-01` | `Date (yyyy-MM-dd)` | | **`lastFailStatus`** | Failure status filter | `eq` | Last HTTP status code. Set when a webhook is auto-disabled due to repeated delivery failures |`positive number`|  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:desc|isEnabled:asc`   Sort by `name` descending and `isEnabled` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | **`id`** | Webhook id | | **`name`** | Webhook name | | **`isEnabled`** | Webhook isEnabled | | **`createdAt`** | Creation date | | **`updatedAt`** | Last modification date | | **`expiration`** | Expiration date |  </details> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisioningApiRequestListOfTenantWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRequestListOfTenantWebhooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

[**WebhookList**](WebhookList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestTenantWebhook**
> Webhook RequestTenantWebhook(ctx, webhookId, optional)
Request tenant webhook

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Get a specific webhook for the tenant scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage webhook</span> required.  ### Postcondition: Webhook is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***ProvisioningApiRequestTenantWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRequestTenantWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetTenantWebhookLifetime**
> Webhook ResetTenantWebhookLifetime(ctx, webhookId, optional)
Reset tenant webhook lifetime

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Reset the lifetime of a webhook for the tenant scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage webhook</span> required.  ### Postcondition: Lifetime of the webhook is reset.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***ProvisioningApiResetTenantWebhookLifetimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiResetTenantWebhookLifetimeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsServiceToken** | **optional.String**| Service Authentication token | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetCustomerAttributes**
> Customer SetCustomerAttributes(ctx, body, customerId, optional)
Set customer attributes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.28.0</h3>  ### Description:   Set custom customer attributes.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> required.  ### Postcondition: Custom customer attributes gets set.  ### Further Information: Batch function.   All existing customer attributes will be deleted.    * Allowed characters for keys are: `[a-zA-Z0-9_-]`   * Characters are **case-insensitive**. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomerAttributes**](CustomerAttributes.md)|  | 
  **customerId** | **int64**| Customer ID | 
 **optional** | ***ProvisioningApiSetCustomerAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiSetCustomerAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsServiceToken** | **optional.**| Service Authentication token | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomer**
> UpdateCustomerResponse UpdateCustomer(ctx, body, customerId, optional)
Update customer

### Description:   Change selected attributes of a customer.  ### Precondition: Authentication with `X-Sds-Service-Token` required.  ### Postcondition: Selected attributes of customer are updated.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCustomerRequest**](UpdateCustomerRequest.md)|  | 
  **customerId** | **int64**| Customer ID | 
 **optional** | ***ProvisioningApiUpdateCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiUpdateCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsServiceToken** | **optional.**| Service Authentication token | 

### Return type

[**UpdateCustomerResponse**](UpdateCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerAttributes**
> Customer UpdateCustomerAttributes(ctx, body, customerId, optional)
Add or edit customer attributes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.4.0</h3>  ### Description:   Add or edit custom customer attributes. <br/><br/><span style=\"font-weight: bold; color: red;\"> &#128679; **Warning: Please note that the response with HTTP status code 200 (OK) is deprecated and will be replaced with HTTP status code 204 (No content)!**</span><br/>  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> required.  ### Postcondition: Custom customer attributes get added or edited.  ### Further Information: Batch function.   If an entry exists before, it will be overwritten.    * Allowed characters for keys are: `[a-zA-Z0-9_-]`   * Characters are **case-insensitive**.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomerAttributes**](CustomerAttributes.md)|  | 
  **customerId** | **int64**| Customer ID | 
 **optional** | ***ProvisioningApiUpdateCustomerAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiUpdateCustomerAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsServiceToken** | **optional.**| Service Authentication token | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTenantWebhook**
> Webhook UpdateTenantWebhook(ctx, body, webhookId, optional)
Update tenant webhook

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Update an existing webhook for the tenant scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage webhook</span> required.  ### Postcondition: Webhook is updated.  ### Further Information: URL must begin with the `HTTPS` scheme. Webhook names are limited to 150 characters.  ### Available event types:  <details open style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Name | Description | Scope | | :--- | :--- | :--- | | **`customer.created`** | Triggered when a new customer is created | Tenant Webhook | | **`customer.deleted`** | Triggered when a user is deleted | Tenant Webhook | | **`webhook.expiring`** | Triggered 30/20/10/1 days before a webhook expires |  Tenant Webhook |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateWebhookRequest**](UpdateWebhookRequest.md)|  | 
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***ProvisioningApiUpdateTenantWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiUpdateTenantWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsServiceToken** | **optional.**| Service Authentication token | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

