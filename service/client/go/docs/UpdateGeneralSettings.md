# UpdateGeneralSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharePasswordSmsEnabled** | **bool** | Allow sending of share passwords via SMS | [optional] [default to null]
**CryptoEnabled** | **bool** | Activation status of client-side encryption.  Can only be enabled once; disabling is not possible. | [optional] [default to null]
**EmailNotificationButtonEnabled** | **bool** | Enable email notification button | [optional] [default to null]
**EulaEnabled** | **bool** | Each user has to confirm the EULA at first login. | [optional] [default to null]
**S3TagsEnabled** | **bool** | &amp;#128640; Since v4.9.0  Defines if S3 tags are enabled | [optional] [default to null]
**AuthTokenRestrictions** | [***UpdateAuthTokenRestrictions**](UpdateAuthTokenRestrictions.md) |  | [optional] [default to null]
**HideLoginInputFields** | **bool** | &amp;#128679; Deprecated since v4.13.0  Defines if login fields should be hidden | [optional] [default to null]
**MediaServerEnabled** | **bool** | &amp;#128679; Deprecated since v4.12.0  Determines if the media server is enabled | [optional] [default to null]
**WeakPasswordEnabled** | **bool** | &amp;#128679; Deprecated since v4.14.0  Allow weak password  * A weak password has to fulfill the following criteria:     * is at least 8 characters long     * contains letters and numbers  * A strong password has to fulfill the following criteria in addition:     * contains at least one special character     * contains upper and lower case characters  Please use &#x60;PUT /system/config/policies/passwords&#x60; API to change configured password policies. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

