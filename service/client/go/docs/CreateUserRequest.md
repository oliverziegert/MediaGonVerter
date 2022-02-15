# CreateUserRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | User first name | [default to null]
**LastName** | **string** | User last name | [default to null]
**UserName** | **string** | &amp;#128640; Since v4.13.0  Username | [optional] [default to null]
**Phone** | **string** | Phone number | [optional] [default to null]
**Expiration** | [***ObjectExpiration**](ObjectExpiration.md) |  | [optional] [default to null]
**ReceiverLanguage** | **string** | IETF language tag | [optional] [default to null]
**Email** | **string** | Email  | [optional] [default to null]
**NotifyUser** | **bool** | &amp;#128640; Since v4.9.0  Notify user about his new account  * default: &#x60;true&#x60; for &#x60;basic&#x60; auth type  * default: &#x60;false&#x60; for &#x60;active_directory&#x60;, &#x60;openid&#x60; and &#x60;radius&#x60; auth types | [optional] [default to null]
**AuthData** | [***UserAuthData**](UserAuthData.md) |  | [optional] [default to null]
**IsNonmemberViewer** | **bool** | &amp;#128640; Since v4.12.0  Determines whether user has the role NONMEMBER_VIEWER | [optional] [default to null]
**Title** | **string** | &amp;#128679; Deprecated since v4.18.0  Job title | [optional] [default to null]
**Password** | **string** | &amp;#128679; Deprecated since v4.13.0  An initial password may be preset  use &#x60;authData&#x60; instead | [optional] [default to null]
**Login** | **string** | &amp;#128679; Deprecated since v4.13.0  User login name | [optional] [default to null]
**AuthMethods** | [**[]UserAuthMethod**](UserAuthMethod.md) | &amp;#128679; Deprecated since v4.13.0  Authentication methods:  * &#x60;sql&#x60;  * &#x60;active_directory&#x60;  * &#x60;radius&#x60;  * &#x60;openid&#x60;  use &#x60;authData&#x60; instead | [optional] [default to null]
**NeedsToChangePassword** | **bool** | &amp;#128679; Deprecated since v4.13.0  Determines whether user has to change his / her initial password.  use &#x60;authDate.mustChangePassword&#x60; instead | [optional] [default to null]
**Gender** | **string** | &amp;#128679; Deprecated since v4.12.0  Gender  Do NOT use &#x60;gender&#x60;! It will be ignored. | [optional] [default to n]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

