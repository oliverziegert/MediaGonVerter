# UserItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the user | [default to null]
**UserName** | **string** | &amp;#128640; Since v4.13.0  Username | [default to null]
**FirstName** | **string** | User first name | [default to null]
**LastName** | **string** | User last name | [default to null]
**IsLocked** | **bool** | User is locked:  * &#x60;false&#x60; - unlocked  * &#x60;true&#x60; - locked    User is locked and can not login anymore. | [default to false]
**HasManageableRooms** | **bool** | &amp;#128679; Deprecated since v4.27.0  User has manageable rooms | [optional] [default to null]
**AvatarUuid** | **string** | &amp;#128640; Since v4.11.0  Avatar UUID | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date | [optional] [default to null]
**LastLoginSuccessAt** | [**time.Time**](time.Time.md) | Last successful logon date | [optional] [default to null]
**ExpireAt** | [**time.Time**](time.Time.md) | Expiration date | [optional] [default to null]
**IsEncryptionEnabled** | **bool** | User has generated private key.  Possible if client-side encryption is active for this customer | [optional] [default to null]
**Email** | **string** | Email  | [optional] [default to null]
**Phone** | **string** | Phone number | [optional] [default to null]
**HomeRoomId** | **int64** | Homeroom ID | [optional] [default to null]
**UserRoles** | [***RoleList**](RoleList.md) |  | [optional] [default to null]
**UserAttributes** | [***UserAttributes**](UserAttributes.md) |  | [optional] [default to null]
**LockStatus** | **int32** | &amp;#128679; Deprecated since v4.7.0  User lock status:  * &#x60;0&#x60; - locked  * &#x60;1&#x60; - Web access allowed  * &#x60;2&#x60; - Web and mobile access allowed    Please use &#x60;isLocked&#x60; instead. | [default to null]
**Login** | **string** | &amp;#128679; Deprecated since v4.13.0  User login name | [default to null]
**Title** | **string** | &amp;#128679; Deprecated since v4.18.0  Job title | [optional] [default to null]
**Gender** | **string** | &amp;#128679; Deprecated since v4.12.0  Gender | [optional] [default to n]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

