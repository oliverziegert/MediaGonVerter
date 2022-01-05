# UploadShare

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Share ID | [default to null]
**Name** | **string** | Alias name | [default to null]
**TargetId** | **int64** | Target room or folder ID | [default to null]
**IsProtected** | **bool** | Is share protected by password | [default to null]
**AccessKey** | **string** | Share access key to generate secure link | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date | [default to null]
**CreatedBy** | [***UserInfo**](UserInfo.md) |  | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Modification date | [optional] [default to null]
**UpdatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**ExpireAt** | [**time.Time**](time.Time.md) | Expiration date | [optional] [default to null]
**TargetPath** | **string** | Path to shared upload node | [optional] [default to null]
**IsEncrypted** | **bool** | Encryption state | [optional] [default to null]
**Notes** | **string** | User notes | [optional] [default to null]
**InternalNotes** | **string** | &amp;#128640; Since v4.11.0  Internal notes | [optional] [default to null]
**FilesExpiryPeriod** | **int32** | Number of days after which uploaded files expire | [optional] [default to null]
**CntFiles** | **int32** | Total amount of existing files uploaded with this share. | [optional] [default to null]
**CntUploads** | **int32** | Total amount of uploads conducted with this share. | [optional] [default to null]
**ShowUploadedFiles** | **bool** | Allow display of already uploaded files | [optional] [default to null]
**DataUrl** | **string** | Upload Share URL | [optional] [default to null]
**MaxSlots** | **int32** | Maximal amount of files to upload | [optional] [default to null]
**MaxSize** | **int64** | Maximal total size of uploaded files (in bytes) | [optional] [default to null]
**TargetType** | **string** | Node type | [optional] [default to null]
**ShowCreatorName** | **bool** | &amp;#128640; Since v4.11.0  Show creator first and last name. | [optional] [default to null]
**ShowCreatorUsername** | **bool** | &amp;#128640; Since v4.11.0  Show creator email address. | [optional] [default to null]
**NotifyCreator** | **bool** | &amp;#128679; Deprecated since v4.20.0  Notify creator on every upload. | [default to null]
**Recipients** | **string** | &amp;#128679; Deprecated since v4.11.0  CSV string of recipient email addresses | [optional] [default to null]
**SmsRecipients** | **string** | &amp;#128679; Deprecated since v4.11.0  CSV string of recipient MSISDNs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

