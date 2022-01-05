# PublicDownloadShare

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsProtected** | **bool** | Is share protected by password | [default to null]
**FileName** | **string** | File name | [default to null]
**Size** | **int64** | File size or container size not compressed (in bytes) | [default to null]
**LimitReached** | **bool** | Downloads limit reached | [default to null]
**CreatorName** | **string** | Creator name | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date | [default to null]
**HasDownloadLimit** | **bool** | &amp;#128640; Since v4.11.0  Determines whether Download Share has a limit for amount of downloads | [default to null]
**MediaType** | **string** | &amp;#128640; Since v4.11.0  * &#x60;application/zip&#x60; (for folders and rooms)  * actual file media type (for files only) | [default to null]
**Name** | **string** | Share display name (alias name) | [optional] [default to null]
**CreatorUsername** | **string** | Creator username | [optional] [default to null]
**ExpireAt** | [**time.Time**](time.Time.md) | Expiration date | [optional] [default to null]
**Notes** | **string** | User notes | [optional] [default to null]
**IsEncrypted** | **bool** | Encryption state | [optional] [default to null]
**FileKey** | [***FileKey**](FileKey.md) |  | [optional] [default to null]
**PrivateKeyContainer** | [***PrivateKeyContainer**](PrivateKeyContainer.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

