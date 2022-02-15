# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAndPreserveKeyPair**](SettingsApi.md#CreateAndPreserveKeyPair) | **Post** /v4/settings/keypairs | Create system rescue key pair and preserve copy of old private key
[**CreateWebhook**](SettingsApi.md#CreateWebhook) | **Post** /v4/settings/webhooks | Create webhook
[**RemoveSystemRescueKeyPair**](SettingsApi.md#RemoveSystemRescueKeyPair) | **Delete** /v4/settings/keypair | Remove system rescue key pair
[**RemoveWebhook**](SettingsApi.md#RemoveWebhook) | **Delete** /v4/settings/webhooks/{webhook_id} | Remove webhook
[**RequestAllSystemRescueKeyPairs**](SettingsApi.md#RequestAllSystemRescueKeyPairs) | **Get** /v4/settings/keypairs | Request all system rescue key pairs
[**RequestListOfEventTypesForConfigManager**](SettingsApi.md#RequestListOfEventTypesForConfigManager) | **Get** /v4/settings/webhooks/event_types | Request list of event types
[**RequestListOfWebhooks**](SettingsApi.md#RequestListOfWebhooks) | **Get** /v4/settings/webhooks | Request list of webhooks
[**RequestNotificationChannels**](SettingsApi.md#RequestNotificationChannels) | **Get** /v4/settings/notifications/channels | Request list of notification channels
[**RequestSettings**](SettingsApi.md#RequestSettings) | **Get** /v4/settings | Request customer settings
[**RequestSystemRescueKeyPair**](SettingsApi.md#RequestSystemRescueKeyPair) | **Get** /v4/settings/keypair | Request system rescue key pair
[**RequestWebhook**](SettingsApi.md#RequestWebhook) | **Get** /v4/settings/webhooks/{webhook_id} | Request webhook
[**ResetWebhookLifetime**](SettingsApi.md#ResetWebhookLifetime) | **Post** /v4/settings/webhooks/{webhook_id}/reset_lifetime | Reset webhook lifetime
[**SetSettings**](SettingsApi.md#SetSettings) | **Put** /v4/settings | Set customer settings
[**SetSystemRescueKeyPair**](SettingsApi.md#SetSystemRescueKeyPair) | **Post** /v4/settings/keypair | Activate client-side encryption for customer
[**ToggleNotificationChannels**](SettingsApi.md#ToggleNotificationChannels) | **Put** /v4/settings/notifications/channels | Toggle notification channels
[**UpdateWebhook**](SettingsApi.md#UpdateWebhook) | **Put** /v4/settings/webhooks/{webhook_id} | Update webhook

# **CreateAndPreserveKeyPair**
> CreateAndPreserveKeyPair(ctx, body, optional)
Create system rescue key pair and preserve copy of old private key

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Create system rescue key pair and preserve copy of old private key.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> required.  ### Postcondition: System rescue key pair is created.   Copy of old private key is preserved.  ### Further Information: You can submit your old private key, encrypted with your current password.   This allows migrating file keys encrypted with your old key pair to the new one. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateKeyPairRequest**](CreateKeyPairRequest.md)|  | 
 **optional** | ***SettingsApiCreateAndPreserveKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiCreateAndPreserveKeyPairOpts struct
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

# **CreateWebhook**
> Webhook CreateWebhook(ctx, body, optional)
Create webhook

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Create a new webhook for the customer scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change config</span> required.  ### Postcondition: Webhook is created for given event types.  ### Further Information: URL must begin with the `HTTPS` scheme.   Webhook names are limited to 150 characters.  ### Available event types: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Name | Description | Scope | | :--- | :--- | :--- | | **`user.created`** | Triggered when a new user is created | Customer Admin Webhook | | **`user.deleted`** | Triggered when a user is deleted | Customer Admin Webhook | | **`user.locked`** | Triggered when a user gets locked | Customer Admin Webhook | |  |  |  | | **`webhook.expiring`** | Triggered 30/20/10/1 days before a webhook expires |  Customer Admin Webhook | |  |  |  | | **`downloadshare.created`** | Triggered when a new download share is created in affected room | Node Webhook | | **`downloadshare.deleted`** | Triggered when a download share is deleted in affected room | Node Webhook | | **`downloadshare.used`** | Triggered when a download share is utilized in affected room | Node Webhook | | **`uploadshare.created`** | Triggered when a new upload share is created in affected room | Node Webhook | | **`uploadshare.deleted`** | Triggered when a upload share is deleted in affected room | Node Webhook | | **`uploadshare.used`** | Triggered when a new file is uploaded via the upload share in affected room | Node Webhook | | **`file.created`** | Triggered when a new file is uploaded in affected room | Node Webhook | | **`folder.created`** | Triggered when a new folder is created in affected room | Node Webhook | | **`room.created`** | Triggered when a new room is created (in affected room) | Node Webhook | | **`file.deleted`** | Triggered when a file is deleted in affected room | Node Webhook | | **`folder.deleted`** | Triggered when a folder is deleted in affected room | Node Webhook | | **`room.deleted`** | Triggered when a room is deleted in affected room | Node Webhook |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateWebhookRequest**](CreateWebhookRequest.md)|  | 
 **optional** | ***SettingsApiCreateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiCreateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveSystemRescueKeyPair**
> RemoveSystemRescueKeyPair(ctx, optional)
Remove system rescue key pair

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Remove the system rescue key pair.  ### Precondition: * Authenticated user * Existence of own key pair  ### Postcondition: Key pair is removed (cf. further information below).  ### Further Information: Please set a new system rescue key pair first and re-encrypt file keys with it.   If no version is set, deleted key pair with lowest preference value.   Although, `version` **SHOULD** be set. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiRemoveSystemRescueKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiRemoveSystemRescueKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **optional.String**| Version (NEW) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveWebhook**
> RemoveWebhook(ctx, webhookId, optional)
Remove webhook

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Delete a webhook for the customer scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change config</span> required.  ### Postcondition: Webhook is deleted.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***SettingsApiRemoveWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiRemoveWebhookOpts struct
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

# **RequestAllSystemRescueKeyPairs**
> []UserKeyPairContainer RequestAllSystemRescueKeyPairs(ctx, optional)
Request all system rescue key pairs

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Retrieve all system rescue key pairs to allow migrating system-rescue-key-encrypted file keys.  ### Precondition: * Authenticated user * Existence of own key pair  ### Postcondition: List of key pairs is returned.  ### Further Information: In the case of an algorithm migration of a system rescue key, one should create the new key pair before deleting the old one.   This allows re-encrypting file keys with the new key pair, using the old one.    This API allows to retrieve both key pairs, in contrast to `GET /settings/keypair`, which only delivers the preferred one. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiRequestAllSystemRescueKeyPairsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiRequestAllSystemRescueKeyPairsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**[]UserKeyPairContainer**](UserKeyPairContainer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestListOfEventTypesForConfigManager**
> EventTypeList RequestListOfEventTypesForConfigManager(ctx, optional)
Request list of event types

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Get a list of available (for <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span>) event types.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change config</span> required.  ### Postcondition: List of available event types is returned.  ### Further Information: None. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiRequestListOfEventTypesForConfigManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiRequestListOfEventTypesForConfigManagerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**EventTypeList**](EventTypeList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestListOfWebhooks**
> WebhookList RequestListOfWebhooks(ctx, optional)
Request list of webhooks

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Get a list of webhooks for the customer scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change config</span> required.  ### Postcondition: List of webhooks is returned.  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:cn:goo|createdAt:ge:2015-01-01`   Get webhooks where name contains `goo` **AND** webhook creation date is **>=** `2015-01-01`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | **`id`** | Webhook id filter | `eq` | Webhook id equals value.<br>Multiple values are allowed and will be connected via logical disjunction (**OR**). |`positive number`| | **`name`** | Webhook type name| `cn, eq` | Webhook name contains / equals value. | `search String` | | **`isEnabled`** | Webhook isEnabled filter | `eq` |  | `true or false` | | **`createdAt`** | Creation date filter | `ge, le` | Creation date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `createdAt:ge:2016-12-31`&#124;`createdAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | **`updatedAt`** | Last modification date filter | `ge, le` | Last modification date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `updatedAt:ge:2016-12-31`&#124;`updatedAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | **`expiration`** | Expiration date filter | `ge, le, eq` | Expiration date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `expiration:ge:2016-12-31`&#124;`expiration:le:2018-01-01` | `Date (yyyy-MM-dd)` | | **`lastFailStatus`** | Failure status filter | `eq` | Last HTTP status code. Set when a webhook is auto-disabled due to repeated delivery failures |`positive number`|  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:desc|isEnabled:asc`   Sort by `name` descending and `isEnabled` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | **`id`** | Webhook id | | **`name`** | Webhook name | | **`isEnabled`** | Webhook isEnabled | | **`createdAt`** | Creation date | | **`updatedAt`** | Last modification date | | **`expiration`** | Expiration date |  </details> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiRequestListOfWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiRequestListOfWebhooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**WebhookList**](WebhookList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestNotificationChannels**
> NotificationChannelList RequestNotificationChannels(ctx, optional)
Request list of notification channels

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description:   Retrieve a list of configured notification channels.  ### Precondition: Right _\"change config\"_ required.  ### Postcondition: List of notification channels is returned.  ### Further Information: None. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiRequestNotificationChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiRequestNotificationChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**NotificationChannelList**](NotificationChannelList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSettings**
> CustomerSettingsResponse RequestSettings(ctx, optional)
Request customer settings

### Description:   Retrieve customer related settings.   ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read config</span> required.  ### Postcondition: List of available settings is returned.  ### Further Information: None.  ### Configurable customer settings: <details open style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | `homeRoomParentName` | Name of the container in which all user's home rooms are located.<br>`null` if `homeRoomsActive` is `false`. | `String` | | `homeRoomQuota` | Refers to the quota of each single user's home room.<br>`0` represents no quota.<br>`null` if `homeRoomsActive` is `false`. | `positive Long` | | `homeRoomsActive` | If set to `true`, every user with an Active Directory account gets a personal homeroom.<br>Once activated, this **CANNOT** be deactivated. | `true or false` |   </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiRequestSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiRequestSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**CustomerSettingsResponse**](CustomerSettingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSystemRescueKeyPair**
> UserKeyPairContainer RequestSystemRescueKeyPair(ctx, optional)
Request system rescue key pair

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Retrieve the system rescue key pair.  ### Precondition: * Authenticated user * Existence of own key pair  ### Postcondition: Key pair is returned.  ### Further Information: If more than one key pair exists the one with highest preference value is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiRequestSystemRescueKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiRequestSystemRescueKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **version** | **optional.String**| Version (NEW) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UserKeyPairContainer**](UserKeyPairContainer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestWebhook**
> Webhook RequestWebhook(ctx, webhookId, optional)
Request webhook

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Get a specific webhook for the customer scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change config</span> required.  ### Postcondition: Webhook is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***SettingsApiRequestWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiRequestWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetWebhookLifetime**
> Webhook ResetWebhookLifetime(ctx, webhookId, optional)
Reset webhook lifetime

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Reset the lifetime of a webhook for the customer scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change config</span> required.  ### Postcondition: Lifetime of the webhook is reset.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***SettingsApiResetWebhookLifetimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiResetWebhookLifetimeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSettings**
> CustomerSettingsResponse SetSettings(ctx, body, optional)
Set customer settings

### Description:   Set customer related settings.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> required.  ### Postcondition: Provided settings are updated.  ### Further Information: None.  ### Configurable customer settings <details open style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | `homeRoomParentName` | Name of the container in which all user's home rooms are located.<br>`null` if `homeRoomsActive` is `false`. | `String` | | `homeRoomQuota` | Refers to the quota of each single user's home room.<br>`0` represents no quota.<br>`null` if `homeRoomsActive` is `false`. | `positive Long` | | `homeRoomsActive` | If set to `true`, every user with an Active Directory account gets a personal homeroom.<br>Once activated, this **CANNOT** be deactivated. | `true or false` |  </details>  ### Node naming convention: * Node (room, folder, file) names are limited to **150** characters. * Illegal names:   `'CON', 'PRN', 'AUX', 'NUL', 'COM1', 'COM2', 'COM3', 'COM4', 'COM5', 'COM6', 'COM7', 'COM8', 'COM9', 'LPT1', 'LPT2', 'LPT3', 'LPT4', 'LPT5', 'LPT6', 'LPT7', 'LPT8', 'LPT9', (and any of those with an extension)` * Illegal characters in names:   `'\\\\', '<','>', ':', '\\\"', '|', '?', '*', '/', leading '-', trailing '.' ` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomerSettingsRequest**](CustomerSettingsRequest.md)|  | 
 **optional** | ***SettingsApiSetSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiSetSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**CustomerSettingsResponse**](CustomerSettingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSystemRescueKeyPair**
> SetSystemRescueKeyPair(ctx, body, optional)
Activate client-side encryption for customer

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Set the system rescue key pair and activate client-side encryption for according customer.  ### Precondition: Role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> required.  ### Postcondition: System rescue key pair is set and client-side encryption is enabled.  ### Further Information: Sets the ability for this customer to encrypt rooms.   Once enabled on customer level, it **CANNOT** be unset.   On activation, a customer rescue key pair **MUST** be set. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserKeyPairContainer**](UserKeyPairContainer.md)|  | 
 **optional** | ***SettingsApiSetSystemRescueKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiSetSystemRescueKeyPairOpts struct
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

# **ToggleNotificationChannels**
> NotificationChannelList ToggleNotificationChannels(ctx, body, optional)
Toggle notification channels

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description:   Toggle configured notification channels.  ### Precondition: Right _\"change config\"_ required.  ### Postcondition: Channel status is switched.  ### Further Information: None. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationChannelActivationRequest**](NotificationChannelActivationRequest.md)|  | 
 **optional** | ***SettingsApiToggleNotificationChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiToggleNotificationChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**NotificationChannelList**](NotificationChannelList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> Webhook UpdateWebhook(ctx, body, webhookId, optional)
Update webhook

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Update an existing webhook for the customer scope.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change config</span> required.  ### Postcondition: Webhook is updated.  ### Further Information: URL must begin with the `HTTPS` scheme. Webhook names are limited to 150 characters. Webhook event types can not be changed from Customer Admin Webhook types to Node Webhook types and vice versa    ### Available event types: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Name | Description | Scope | | :--- | :--- | :--- | | **`user.created`** | Triggered when a new user is created | Customer Admin Webhook | | **`user.deleted`** | Triggered when a user is deleted | Customer Admin Webhook | | **`user.locked`** | Triggered when a user gets locked | Customer Admin Webhook | |  |  |  | | **`webhook.expiring`** | Triggered 30/20/10/1 days before a webhook expires |  Customer Admin Webhook | |  |  |  | | **`downloadshare.created`** | Triggered when a new download share is created in affected room | Node Webhook | | **`downloadshare.deleted`** | Triggered when a download share is deleted in affected room | Node Webhook | | **`downloadshare.used`** | Triggered when a download share is utilized in affected room | Node Webhook | | **`uploadshare.created`** | Triggered when a new upload share is created in affected room | Node Webhook | | **`uploadshare.deleted`** | Triggered when a upload share is deleted in affected room | Node Webhook | | **`uploadshare.used`** | Triggered when a new file is uploaded via the upload share in affected room | Node Webhook | | **`file.created`** | Triggered when a new file is uploaded in affected room | Node Webhook | | **`folder.created`** | Triggered when a new folder is created in affected room | Node Webhook | | **`room.created`** | Triggered when a new room is created (in affected room) | Node Webhook | | **`file.deleted`** | Triggered when a file is deleted in affected room | Node Webhook | | **`folder.deleted`** | Triggered when a folder is deleted in affected room | Node Webhook | | **`room.deleted`** | Triggered when a room is deleted in affected room | Node Webhook |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateWebhookRequest**](UpdateWebhookRequest.md)|  | 
  **webhookId** | **int64**| Webhook ID | 
 **optional** | ***SettingsApiUpdateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiUpdateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

