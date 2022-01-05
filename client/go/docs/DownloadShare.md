# DownloadShare

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Share ID | [default to null]
**Name** | **string** | Alias name | [default to null]
**NodeId** | **int64** | Source node ID | [default to null]
**AccessKey** | **string** | Share access key to generate secure link | [default to null]
**CntDownloads** | **int32** | Downloads counter (incremented on each download) | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date | [default to null]
**CreatedBy** | [***UserInfo**](UserInfo.md) |  | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Modification date | [optional] [default to null]
**UpdatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**Notes** | **string** | User notes | [optional] [default to null]
**InternalNotes** | **string** | &amp;#128640; Since v4.11.0  Internal notes | [optional] [default to null]
**ShowCreatorName** | **bool** | Show creator first and last name. | [optional] [default to null]
**ShowCreatorUsername** | **bool** | Show creator email address. | [optional] [default to null]
**IsProtected** | **bool** | Is share protected by password | [optional] [default to null]
**ExpireAt** | [**time.Time**](time.Time.md) | Expiration date | [optional] [default to null]
**MaxDownloads** | **int32** | Max allowed downloads | [optional] [default to null]
**NodePath** | **string** | Path to shared download node | [optional] [default to null]
**DataUrl** | **string** | Path to shared download node | [optional] [default to null]
**IsEncrypted** | **bool** | Encrypted share  (this only applies to shared files, not folders) | [optional] [default to null]
**NodeType** | **string** | Node type | [optional] [default to null]
**Classification** | **int32** | &amp;#128679; Deprecated since v4.11.0  Classification ID:  * &#x60;1&#x60; - public  * &#x60;2&#x60; - internal  * &#x60;3&#x60; - confidential  * &#x60;4&#x60; - strictly confidential    (default: classification from parent room) | [optional] [default to null]
**NotifyCreator** | **bool** | &amp;#128679; Deprecated since v4.20.0  Notify creator on every download. | [default to null]
**Recipients** | **string** | &amp;#128679; Deprecated since v4.11.0  CSV string of recipient email addresses | [optional] [default to null]
**SmsRecipients** | **string** | &amp;#128679; Deprecated since v4.11.0  CSV string of recipient MSISDNs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

