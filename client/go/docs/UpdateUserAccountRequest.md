# UpdateUserAccountRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | &amp;#128640; Since v4.13.0  Username | [optional] [default to null]
**AcceptEULA** | **bool** | Accept EULA  Present, if EULA is system global active.  cf. &#x60;GET system/config/settings/general&#x60; - &#x60;eulaEnabled&#x60;  If accepted can not be undone. | [optional] [default to null]
**FirstName** | **string** | User first name | [optional] [default to null]
**LastName** | **string** | User last name | [optional] [default to null]
**Email** | **string** | Email  | [optional] [default to null]
**Phone** | **string** | Phone number | [optional] [default to null]
**Language** | **string** | &amp;#128640; Since v4.20.0  IETF language tag | [optional] [default to null]
**Title** | **string** | &amp;#128679; Deprecated since v4.18.0  Job title | [optional] [default to null]
**Login** | **string** | &amp;#128679; Deprecated since v4.13.0  User login name | [optional] [default to null]
**Gender** | **string** | &amp;#128679; Deprecated since v4.12.0  Gender  Do NOT use &#x60;gender&#x60;! It will be ignored. | [optional] [default to n]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

