# GeneralSettingsInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharePasswordSmsEnabled** | **bool** | Allow sending of share passwords via SMS | [optional] [default to null]
**CryptoEnabled** | **bool** | Activation status of client-side encryption.  Can only be enabled once; disabling is not possible. | [optional] [default to null]
**EmailNotificationButtonEnabled** | **bool** | Enable email notification button | [optional] [default to null]
**EulaEnabled** | **bool** | Each user has to confirm the EULA at first login. | [optional] [default to null]
**WeakPasswordEnabled** | **bool** | Allow weak password  * A weak password has to fulfill the following criteria:     * is at least 8 characters long     * contains letters and numbers  * A strong password has to fulfill the following criteria in addition:     * contains at least one special character     * contains upper and lower case characters | [optional] [default to null]
**UseS3Storage** | **bool** | Defines if S3 is used as storage backend | [optional] [default to null]
**S3TagsEnabled** | **bool** | &amp;#128640; Since v4.9.0  Defines if S3 tags are enabled | [optional] [default to null]
**HomeRoomsActive** | **bool** | &amp;#128640; Since v4.10.0  Homerooms active | [optional] [default to null]
**HomeRoomParentId** | **int64** | &amp;#128640; Since v4.10.0  Homeroom Parent ID | [optional] [default to null]
**MediaServerEnabled** | **bool** | &amp;#128679; Deprecated since v4.12.0  Determines if the media server is enabled | [optional] [default to null]
**SubscriptionPlan** | **int32** | &amp;#128640; Since v4.30.0  Subscription Plan | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

