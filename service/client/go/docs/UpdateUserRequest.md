# UpdateUserRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | User first name | [optional] [default to null]
**LastName** | **string** | User last name | [optional] [default to null]
**UserName** | **string** | &amp;#128640; Since v4.13.0  Username | [optional] [default to null]
**Email** | **string** | Email  | [optional] [default to null]
**IsLocked** | **bool** | User is locked:  * &#x60;false&#x60; - unlocked  * &#x60;true&#x60; - locked    User is locked and can not login anymore. | [optional] [default to false]
**Phone** | **string** | Phone number | [optional] [default to null]
**ReceiverLanguage** | **string** | IETF language tag | [optional] [default to null]
**Expiration** | [***ObjectExpiration**](ObjectExpiration.md) |  | [optional] [default to null]
**AuthData** | [***UserAuthDataUpdateRequest**](UserAuthDataUpdateRequest.md) |  | [optional] [default to null]
**Title** | **string** | &amp;#128679; Deprecated since v4.18.0  Job title | [optional] [default to null]
**LockStatus** | **int32** | &amp;#128679; Deprecated since v4.7.0  User lock status:  * &#x60;0&#x60; - locked  * &#x60;1&#x60; - Web access allowed  * &#x60;2&#x60; - Web and mobile access allowed    Please use &#x60;isLocked&#x60; instead. | [optional] [default to null]
**Gender** | **string** | &amp;#128679; Deprecated since v4.12.0  Gender  Do NOT use &#x60;gender&#x60;! It will be ignored. | [optional] [default to n]
**AuthMethods** | [**[]UserAuthMethod**](UserAuthMethod.md) | &amp;#128679; Deprecated since v4.13.0  Authentication methods:  * &#x60;sql&#x60;  * &#x60;active_directory&#x60;  * &#x60;radius&#x60;  * &#x60;openid&#x60;  use &#x60;authData&#x60; instead | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

