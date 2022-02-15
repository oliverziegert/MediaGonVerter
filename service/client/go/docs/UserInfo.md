# UserInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the user | [default to null]
**UserType** | **string** | &amp;#128640; Since v4.11.0  User type:  * &#x60;internal&#x60; - ordinary DRACOON user  * &#x60;external&#x60; - external user without DRACOON account  * &#x60;system&#x60; - system user (non human &amp;#128125;)  * &#x60;deleted&#x60; - deleted DRACOON user | [default to null]
**AvatarUuid** | **string** | &amp;#128640; Since v4.11.0  Avatar UUID | [default to null]
**UserName** | **string** | &amp;#128640; Since v4.13.0  Username (only returned for &#x60;internal&#x60; users) | [default to null]
**FirstName** | **string** | &amp;#128640; Since v4.11.0  User first name (mandatory if &#x60;userType&#x60; is &#x60;internal&#x60;) | [default to null]
**LastName** | **string** | &amp;#128640; Since v4.11.0  User last name (mandatory if &#x60;userType&#x60; is &#x60;internal&#x60;) | [default to null]
**Email** | **string** | &amp;#128640; Since v4.11.0  Email  | [optional] [default to null]
**Title** | **string** | &amp;#128679; Deprecated since v4.18.0  Job title | [optional] [default to null]
**DisplayName** | **string** | &amp;#128679; Deprecated since v4.11.0  Display name  use other fields from &#x60;UserInfo&#x60; instead to combine a display name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

