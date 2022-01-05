# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestAlgorithms**](ConfigApi.md#RequestAlgorithms) | **Get** /v4/config/info/policies/algorithms | Request algorithms
[**RequestClassificationPoliciesConfigInfo**](ConfigApi.md#RequestClassificationPoliciesConfigInfo) | **Get** /v4/config/info/policies/classifications | Request classification policies
[**RequestGeneralSettingsInfo**](ConfigApi.md#RequestGeneralSettingsInfo) | **Get** /v4/config/info/general | Request general settings
[**RequestInfrastructurePropertiesInfo**](ConfigApi.md#RequestInfrastructurePropertiesInfo) | **Get** /v4/config/info/infrastructure | Request infrastructure properties
[**RequestNotificationChannelsInfo**](ConfigApi.md#RequestNotificationChannelsInfo) | **Get** /v4/config/info/notifications/channels | Request list of notification channels
[**RequestPasswordPoliciesConfigInfo**](ConfigApi.md#RequestPasswordPoliciesConfigInfo) | **Get** /v4/config/info/policies/passwords | Request password policies
[**RequestS3TagsInfo**](ConfigApi.md#RequestS3TagsInfo) | **Get** /v4/config/info/s3_tags | Request list of configured S3 tags
[**RequestSystemDefaultsInfo**](ConfigApi.md#RequestSystemDefaultsInfo) | **Get** /v4/config/info/defaults | Request default values
[**RequestSystemSettings**](ConfigApi.md#RequestSystemSettings) | **Get** /v4/config/settings | Request system settings
[**UpdateSystemSettings**](ConfigApi.md#UpdateSystemSettings) | **Put** /v4/config/settings | Update system settings

# **RequestAlgorithms**
> AlgorithmVersionInfoList RequestAlgorithms(ctx, optional)
Request algorithms

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description: Retrieve a list of available algorithms used for encryption.  ### Precondition: Authenticated user.  ### Postcondition: List of available algorithms is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiRequestAlgorithmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiRequestAlgorithmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**AlgorithmVersionInfoList**](AlgorithmVersionInfoList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestClassificationPoliciesConfigInfo**
> ClassificationPoliciesConfig RequestClassificationPoliciesConfigInfo(ctx, optional)
Request classification policies

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.30.0</h3>  ### Description: Retrieve a list of classification policies: * `shareClassificationPolicies`  ### Precondition: Authenticated user.  ### Postcondition: List of configured classification policies is returned.  ### Further Information: `classificationRequiresSharePassword`: When a node has this classification or higher, it can not be shared without a password. `0` means no password will be enforced. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiRequestClassificationPoliciesConfigInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiRequestClassificationPoliciesConfigInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**ClassificationPoliciesConfig**](ClassificationPoliciesConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestGeneralSettingsInfo**
> GeneralSettingsInfo RequestGeneralSettingsInfo(ctx, optional)
Request general settings

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.6.0</h3>  ### Description: Returns a list of configurable general settings.  ### Precondition: Authenticated user.  ### Postcondition: List of configurable general settings is returned.  ### Further Information: None.  ### Configurable general settings: <details open style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | `sharePasswordSmsEnabled` | Determines whether sending of share passwords via SMS is allowed. | `true or false` | | `cryptoEnabled` | Determines whether client-side encryption is enabled.<br>Can only be enabled once; disabling is **NOT** possible. | `true or false` | | `emailNotificationButtonEnabled` | Determines whether email notification button is enabled. | `true or false` | | `eulaEnabled` | Determines whether EULA is enabled.<br>Each user has to confirm the EULA at first login. | `true or false` | | `useS3Storage` | Defines if S3 is used as storage backend.<br>Can only be enabled once; disabling is **NOT** possible. | `true or false` | | `s3TagsEnabled` | Determines whether S3 tags are enabled | `true or false` | | `homeRoomsActive` | Determines whether each AD user has a personal home room | `true or false` | | `homeRoomParentId` | Defines a node under which all personal home rooms are located. **NULL** if `homeRoomsActive` is `false` | `Long` | | `subscriptionPlan` | Subscription Plan. <br> 0 = Standard, 1 = Premium, 2 = Free | `Integer` |  </details>  ### Deprecated general settings: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | <del>`mediaServerEnabled`</del> | Determines whether media server is enabled.<br>Returns boolean value dependent on conjunction of `mediaServerConfigEnabled` AND `mediaServerEnabled` | `true or false` | | <del>`weakPasswordEnabled`</del> | Determines whether weak password is allowed.<br>Use `GET /system/config/policies/passwords` API to get configured password policies. | `true or false` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiRequestGeneralSettingsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiRequestGeneralSettingsInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**GeneralSettingsInfo**](GeneralSettingsInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestInfrastructurePropertiesInfo**
> InfrastructureProperties RequestInfrastructurePropertiesInfo(ctx, optional)
Request infrastructure properties

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.6.0</h3>  ### Description:   Returns a list of read-only infrastructure properties.    ### Precondition: Authenticated user.  ### Postcondition: List of infrastructure properties is returned.  ### Further Information: Source: `core-service.properties`  ### Read-only infrastructure properties: <details open style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | `smsConfigEnabled` | Determines whether sending of share passwords via SMS is **system-wide** enabled. | `true or false` | | `mediaServerConfigEnabled` | Determines whether media server is **system-wide** enabled. | `true or false` | | `s3DefaultRegion` | Suggested S3 region | `Region name` | | `s3EnforceDirectUpload` | Enforce direct upload to S3 | `true or false` | | `isDracoonCloud` | Determines if the **DRACOON Core** is deployed in the cloud environment | `true or false` | | `tenantUuid` | Current tenant UUID | `UUID` |  </details> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiRequestInfrastructurePropertiesInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiRequestInfrastructurePropertiesInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**InfrastructureProperties**](InfrastructureProperties.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestNotificationChannelsInfo**
> NotificationChannelList RequestNotificationChannelsInfo(ctx, optional)
Request list of notification channels

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.20.0</h3>  ### Description: Retrieve a list of configured notification channels.  ### Precondition: Authenticated user.  ### Postcondition: List of notification channels is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiRequestNotificationChannelsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiRequestNotificationChannelsInfoOpts struct
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

# **RequestPasswordPoliciesConfigInfo**
> PasswordPoliciesConfig RequestPasswordPoliciesConfigInfo(ctx, optional)
Request password policies

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.14.0</h3>  ### Description:   Retrieve a list of configured password policies for all password types:   * `login` * `shares` * `encryption`  ### Precondition: Authenticated user.  ### Postcondition: List of configured password policies is returned.  ### Further Information: None.  ### Available password policies: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Name | Description | Value | Password Type | | :--- | :--- | :--- | :--- | | `mustContainCharacters` | Characters which a password must contain:<br><ul><li>`alpha` - at least one alphabetical character (`uppercase` **OR** `lowercase`)<pre>a b c d e f g h i j k l m n o p q r s t u v w x y z<br>A B C D E F G H I J K L M N O P Q R S T U V W X Y Z</pre></li><li>`uppercase` - at least one uppercase character<pre>A B C D E F G H I J K L M N O P Q R S T U V W X Y Z</pre></li><li>`lowercase` - at least one lowercase character<pre>a b c d e f g h i j k l m n o p q r s t u v w x y z</pre></li><li>`numeric` - at least one numeric character<pre>0 1 2 3 4 5 6 7 8 9</pre></li><li>`special` - at least one special character (letters and digits excluded)<pre>! \" # $ % ( ) * + , - . / : ; = ? @ [ \\ ] ^ _ { &#124; } ~</pre></li><li>`none` - none of the above</li></ul> | <ul><li>`alpha`</li><li>`uppercase`</li><li>`lowercase`</li><li>`numeric`</li><li>`special`</li><li>`none`</li></ul> | <ul><li>`login`</li><li>`shares`</li><li>`encryption`</li></ul> | | `numberOfCharacteristicsToEnforce` | Number of characteristics to enforce.<br>e.g. from `[\"uppercase\", \"lowercase\", \"numeric\", \"special\"]`<br>all 4 character sets can be enforced; but also only 2 of them | `Integer between 0 and 4` | <ul><li>`login`</li><li>`shares`</li><li>`encryption`</li></ul> | | `minLength` | Minimum number of characters a password must contain. | `Integer between 1 and 1024` | <ul><li>`login`</li><li>`shares`</li><li>`encryption`</li></ul> | | `rejectDictionaryWords` | Determines whether a password must **NOT** contain word(s) from a dictionary.<br>In `core-service.properties` a path to directory with dictionary files (`*.txt`) can be defined<br>cf. `policies.passwords.dictionary.directory`.<br><br>If this rule gets enabled `policies.passwords.dictionary.directory` must be defined and contain dictionary files.<br>Otherwise, the rule will not have any effect on password validation process. | `true or false` | <ul><li>`login`</li><li>`shares`</li></ul> | | `rejectUserInfo` | Determines whether a password must **NOT** contain user info.<br>Affects user's **first name**, **last name**, **email** and **user name**. | `true or false` | <ul><li>`login`</li><li>`shares`</li><li>`encryption`</li></ul> | | `rejectKeyboardPatterns` | Determines whether a password must **NOT** contain keyboard patterns.<br>e.g. `qwertz`, `asdf` (min. 4 character pattern) | `true or false` | <ul><li>`login`</li><li>`shares`</li><li>`encryption`</li></ul> | | `numberOfArchivedPasswords` | Number of passwords to archive.<br>Value `0` means that password history is disabled. | `Integer between 0 and 10` | <ul><li>`login`</li></ul> | | `passwordExpiration.enabled` | Determines whether password expiration is enabled. | `true or false` | <ul><li>`login`</li></ul> | | `maxPasswordAge` | Maximum allowed password age (in **days**) | `positive Integer` | <ul><li>`login`</li></ul> | | `userLockout.enabled` | Determines whether user lockout is enabled. | `true or false` | <ul><li>`login`</li></ul> | | `maxNumberOfLoginFailures` | Maximum allowed number of failed login attempts. | `positive Integer` | <ul><li>`login`</li></ul> | | `lockoutPeriod` | Amount of **minutes** a user has to wait to make another login attempt<br>after `maxNumberOfLoginFailures` has been exceeded. | `positive Integer` | <ul><li>`login`</li></ul> |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiRequestPasswordPoliciesConfigInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiRequestPasswordPoliciesConfigInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**PasswordPoliciesConfig**](PasswordPoliciesConfig.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestS3TagsInfo**
> S3TagList RequestS3TagsInfo(ctx, optional)
Request list of configured S3 tags

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.9.0</h3>  ### Description: Retrieve all configured S3 tags.  ### Precondition: Authenticated user.  ### Postcondition: List of configured S3 tags is returned.  ### Further Information: An empty list is returned if no S3 tags are found / configured.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiRequestS3TagsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiRequestS3TagsInfoOpts struct
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

# **RequestSystemDefaultsInfo**
> SystemDefaults RequestSystemDefaultsInfo(ctx, optional)
Request default values

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.6.0</h3>  ### Description:   Returns a list of configurable system default values.  ### Precondition: Authenticated user.  ### Postcondition: List of configurable default settings is returned.  ### Further Information: None.  ### Configurable default values: <details open style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | `languageDefault` | Defines which language should be default. | `ISO 639-1 code` | | `downloadShareDefaultExpirationPeriod` | Default expiration period for Download Shares in _days_. | `Integer between 0 and 9999` | | `uploadShareDefaultExpirationPeriod` | Default expiration period for Upload Shares in _days_. | `Integer between 0 and 9999` | | `fileDefaultExpirationPeriod` | Default expiration period for all uploaded files in _days_. | `Integer between 0 and 9999` | | `nonmemberViewerDefault` | Defines if new users get the role _Non Member Viewer_ by default | `true or false` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiRequestSystemDefaultsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiRequestSystemDefaultsInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**SystemDefaults**](SystemDefaults.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSystemSettings**
> ConfigOptionList RequestSystemSettings(ctx, optional)
Request system settings

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.6.0</h3>  ### Description:   Returns a list of configurable system settings.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read global config</span> required.  ### Postcondition: List of configurable settings is returned.  ### Further Information: Check for every settings key new corresponding API and key below.  If `eula_active` is true, but **NOT** accepted yet, or password **MUST** be changed, only the following two values are returned: * `allow_system_global_weak_password` * `eula_active`  ### Configurable settings <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | `branding_server_branding_id` | The branding UUID, which corresponds to _BRANDING-QUALIFIER_ in the new branding server.<br>cf. `GET /system/config/settings/branding` `BrandingConfig.brandingQualifier` | `String` | | `branding_portal_url` | Access URL to to the Branding Portal<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/branding` `BrandingConfig.brandingProviderUrl` | `String` | | `dblog` | Write logs to local database.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/eventlog` `EventlogConfig.enabled` | `true or false` | | `default_downloadshare_expiration_period` | Default expiration period for Download Shares in days<br>cf. `GET /system/config/settings/defaults` `SystemDefaults.downloadShareDefaultExpirationPeriod` | `Integer between 0 and 9999` | | `default_file_upload_expiration_date` | Default expiration period for all uploaded files in days<br>cf. `GET /system/config/settings/defaults` `SystemDefaults.fileDefaultExpirationPeriod` | `Integer between 0 and 9999` | | `default_language` | Define which language should be default.<br>cf. `GET /system/config/settings/defaults` `SystemDefaults.languageDefault` | cf. `GET /public/system/info` - `SystemInfo.languageDefault` | | `default_uploadshare_expiration_period` | Default expiration period for Upload Shares in days<br>cf. `GET /system/config/settings/defaults` `SystemDefaults.uploadShareDefaultExpirationPeriod` | `Integer between 0 and 9999` | | `enable_client_side_crypto` | Activation status of client-side encryption<br>Can only be enabled once; disabling is **NOT** possible.<br>cf. `GET /system/config/settings/general` `GeneralSettings.cryptoEnabled` | `true or false`<br>default: `false` | | `eula_active` | Each user has to confirm the EULA at first login.<br>cf. `GET /system/config/settings/general` `GeneralSettings.eulaEnabled` | `true or false` | | `eventlog_retention_period` | Retention period (in days) of event log entries<br>After that period, all entries are deleted.<br>cf. `GET /system/config/settings/eventlog` `EventlogConfig.retentionPeriod` | `Integer between 0 and 9999`<br>If set to `0`: no logs are deleted<br>Recommended value: `7` | | `ip_address_logging` | Determines whether a user's IP address is logged.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/eventlog` `EventlogConfig.logIpEnabled`<br>cf. `GET /system/config/settings/syslog` `SyslogConfig.logIpEnabled` | `true or false` | | `mailserver` | Email server to send emails.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/mail_server` `MailServerConfig.host` | `DNS name or IPv4 of an email server` | | `mailserver_authentication_necessary` | Set to `true` if the email server requires authentication.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/mail_server` `MailServerConfig.authenticationEnabled` | `true or false` | | `mailserver_password` | **Password is no longer returned.**<br>Check `mailserver_password_set` to determine whether password is set. |  | | `mailserver_password_set` | Indicates if a password is set for the mailserver (because `mailserver_password` is always returned empty).<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/mail_server` `MailServerConfig.passwordDefined` | `true or false` | | `mailserver_port` | Email server port<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/mail_server` `MailServerConfig.port` | `Valid port number` | | `mailserver_username` | User ame for email server<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/mail_server` `MailServerConfig.username` | `Username for authentication` | | `mailserver_use_ssl` | Email server requires SSL connection?<br>Only visible for _Config Manager_ of Provider Customer.<br>Requires `mailserver_use_starttls` to be `false`<br>cf. `GET /system/config/settings/mail_server` `MailServerConfig.username` | `true or false` | | `mailserver_use_starttls` | Email server requires StartTLS connection?<br>Only visible for _Config Manager_ of Provider Customer.<br>Requires `mailserver_use_ssl` to be `false`<br>cf. `GET /system/config/settings/mail_server` `MailServerConfig.starttlsEnabled` | `true or false` | | `syslog` | Write logs to a syslog interface.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/syslog` `SyslogConfig.enabled` | `true or false` | | `syslog_host` | Syslog server (IP or FQDN)<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/syslog` `SyslogConfig.host` | `DNS name or IPv4 of a syslog server` | | `syslog_port` | Syslog server port<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/syslog` `SyslogConfig.port` | `Valid port number` | | `syslog_protocol` | Protocol to connect to syslog server.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `GET /system/config/settings/syslog` `SyslogConfig.protocol` | `TCP or UDP` | | `enable_email_notification_button` | Enable mail notification button.<br>cf. `GET /system/config/settings/general` `GeneralSettings.emailNotificationButtonEnabled` | `true or false` | | `allow_share_password_sms` | Allow sending of share passwords via SMS.<br>cf. `GET /system/config/settings/general` `GeneralSettings.sharePasswordSmsEnabled` | `true or false` | | `globally_allow_share_password_sms` | Allow sending of share passwords via SMS **system-wide** (read-only).<br>cf. `GET /system/config/settings/infrastructure` `InfrastructureProperties.smsConfigEnabled` | `true or false` | | `use_s3_storage` | Defines if S3 is used as storage backend.<br>Can only be enabled once; disabling is **NOT** possible.<br>cf. `GET /system/config/settings/general` `GeneralSettings.useS3Storage` | `true or false` | | `s3_default_region` |Suggested S3 region (read-only)<br>cf. `GET /system/config/settings/infrastructure` `InfrastructureProperties.s3DefaultRegion` | `Region name` |  </details>  ### Deprecated settings <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | <del>`allow_system_global_weak_password`</del> | Determines whether weak password (cf. _Password Policy_ below) is allowed.<br>cf. `GET /system/config/settings/general` `GeneralSettings.weakPasswordEnabled`<br>Use `GET /system/config/policies/passwords` API to get configured password policies. | `true or false` | | <del>`branding_server_customer`</del> | The UUID of the branding server customer, which corresponds to customer key in the branding server. | `String` | | <del>`branding_server_url`</del> | Access URL to to the Branding Server.<br>Only visible for _Config Manager_ of Provider Customer. | `String` | | <del>`email_from`</del> | Sender of system-generated emails<br>Only visible for _Config Manager_ of Provider Customer.<br>**Moved to branding** | `Valid email address` | | <del>`email_to_sales`</del> | Contact email address for customers to request more user licenses or data volume.<br>**Moved to branding** | `Valid email address` | | <del>`email_to_support`</del> | Support email address for users<br>**Moved to branding** | `Valid email address` | | <del>`file_size_js`</del> | Maximum file size (in bytes) for downloads of encrypted files with JavaScript.<br>Bigger files will require a JavaApplet. | `Integer`<br>Recommended value: `10485760` (=`10MB`) | | <del>`system_name`</del> | System name<br>**Moved to branding** use `product.title` | `Display name of the DRACOON` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiRequestSystemSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiRequestSystemSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**ConfigOptionList**](ConfigOptionList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSystemSettings**
> UpdateSystemSettings(ctx, body, optional)
Update system settings

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.6.0</h3>  ### Description: Update configurable settings.  ### Precondition: Right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change global config</span> and role <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Config Manager</span> of the Provider Customer required.  ### Postcondition: One or more global settings gets changed.  ### Further Information: This API is deprecated and will be removed in the future.   Check for every settings key new corresponding API and key below.  ### Configurable settings: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | `branding_server_branding_id` | The branding UUID, which corresponds to _BRANDING-QUALIFIER_ in the new branding server.<br>cf. `PUT /system/config/settings/branding` `BrandingConfig.brandingQualifier` | `String` | | `branding_portal_url` | Access URL to to the Branding Portal<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/branding` `BrandingConfig.brandingProviderUrl` | `String` | | `dblog` | Write logs to local database.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/eventlog` `EventlogConfig.enabled` | `true or false` | | `default_downloadshare_expiration_period` | Default expiration period for Download Shares in days<br>cf. `PUT /system/config/settings/defaults` `SystemDefaults.downloadShareDefaultExpirationPeriod` | `Integer between 0 and 9999`<br>Set `0` to disable. | | `default_file_upload_expiration_date` | Default expiration period for all uploaded files in days<br>cf. `PUT /system/config/settings/defaults` `SystemDefaults.fileDefaultExpirationPeriod` | `Integer between 0 and 9999`<br>Set `0` to disable. | | `default_language` | Define which language should be default.<br>cf. `PUT /system/config/settings/defaults` `SystemDefaults.languageDefault` | cf. `GET /public/system/info` - `SystemInfo.languageDefault` | | `default_uploadshare_expiration_period` | Default expiration period for Upload Shares in days<br>cf. `PUT /system/config/settings/defaults` `SystemDefaults.uploadShareDefaultExpirationPeriod` | `Integer between 0 and 9999`<br>Set `0` to disable. | | `enable_client_side_crypto` | Activation status of client-side encryption<br>Can only be enabled once; disabling is **NOT** possible.<br>cf. `PUT /system/config/settings/general` `GeneralSettings.cryptoEnabled` | `true or false`<br>default: `false` | | `eula_active` | Each user has to confirm the EULA at first login.<br>cf. `PUT /system/config/settings/general` `GeneralSettings.eulaEnabled` | `true or false` | | `eventlog_retention_period` | Retention period (in days) of event log entries<br>After that period, all entries are deleted.<br>cf. `PUT /system/config/settings/eventlog` `EventlogConfig.retentionPeriod` | `Integer between 0 and 9999`<br>If set to `0`: no logs are deleted<br>Recommended value: `7` | | `ip_address_logging` | Determines whether a user's IP address is logged.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/eventlog` `EventlogConfig.logIpEnabled`<br>cf. `PUT /system/config/settings/syslog` `SyslogConfig.logIpEnabled` | `true or false` | | `mailserver` | Email server to send emails.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/mail_server` `MailServerConfig.host` | `DNS name or IPv4 of an email server` | | `mailserver_authentication_necessary` | Set to `true` if the email server requires authentication.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/mail_server` `MailServerConfig.authenticationEnabled` | `true or false` | | `mailserver_password` | Password for email server<br>cf. `PUT /system/config/settings/mail_server` `MailServerConfig.password` | `Password for authentication` | | `mailserver_port` | Email server port<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/mail_server` `MailServerConfig.port` | `Valid port number` | | `mailserver_username` | Username for email server<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/mail_server` `MailServerConfig.username` | `Username for authentication` | | `mailserver_use_ssl` | Email server requires SSL connection?<br>Only visible for _Config Manager_ of Provider Customer.<br>Requires `mailserver_use_starttls` to be `false`<br>cf. `PUT /system/config/settings/mail_server` `MailServerConfig.username` | `true or false` | | `mailserver_use_starttls` | Email server requires StartTLS connection?<br>Only visible for _Config Manager_ of Provider Customer.<br>Requires `mailserver_use_ssl` to be `false`<br>cf. `PUT /system/config/settings/mail_server` `MailServerConfig.starttlsEnabled` | `true or false` | | `syslog` | Write logs to a syslog interface.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/syslog` `SyslogConfig.enabled` | `true or false` | | `syslog_host` | Syslog server (IP or FQDN)<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/syslog` `SyslogConfig.host` | `DNS name or IPv4 of a syslog server` | | `syslog_port` | Syslog server port<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/syslog` `SyslogConfig.port` | `Valid port number` | | `syslog_protocol` | Protocol to connect to syslog server.<br>Only visible for _Config Manager_ of Provider Customer.<br>cf. `PUT /system/config/settings/syslog` `SyslogConfig.protocol` | `TCP or UDP` | | `enable_email_notification_button` | Enable mail notification button.<br>cf. `PUT /system/config/settings/general` `GeneralSettings.emailNotificationButtonEnabled` | `true or false` | | `allow_share_password_sms` | Allow sending of share passwords via SMS.<br>cf. `PUT /system/config/settings/general` `GeneralSettings.sharePasswordSmsEnabled` | `true or false` |  </details>  ### Deprecated settings: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Setting | Description | Value | | :--- | :--- | :--- | | <del>`allow_system_global_weak_password`</del> | Determines whether weak password (cf. _Password Policy_ below) is allowed.<br>cf. `PUT /system/config/settings/general` `GeneralSettings.weakPasswordEnabled`<br>Use `PUT /system/config/policies/passwords` API to change configured password policies. | `true or false` | | <del>`branding_server_customer`</del> | The UUID of the branding server customer, which corresponds to customer key in the branding server. | `String` | | <del>`branding_server_url`</del> | Access URL to to the Branding Server.<br>Only visible for _Config Manager_ of Provider Customer. | `String` | | <del>`email_from`</del> | Sender of system-generated emails<br>Only visible for _Config Manager_ of Provider Customer.<br>**Moved to branding** | `Valid email address` | | <del>`email_to_sales`</del> | Contact email address for customers to request more user licenses or data volume.<br>**Moved to branding** | `Valid email address` | | <del>`email_to_support`</del> | Support email address for users<br>**Moved to branding** | `Valid email address` | | <del>`file_size_js`</del> | Maximum file size (in bytes) for downloads of encrypted files with JavaScript.<br>Bigger files will require a JavaApplet. | `Integer`<br>Recommended value: `10485760` (=`10MB`) | | <del>`system_name`</del> | System name<br>**Moved to branding** use `product.title` | `Display name of the DRACOON` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigOptionList**](ConfigOptionList.md)|  | 
 **optional** | ***ConfigApiUpdateSystemSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiUpdateSystemSettingsOpts struct
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

