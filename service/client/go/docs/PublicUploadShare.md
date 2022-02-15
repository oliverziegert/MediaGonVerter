# PublicUploadShare

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsProtected** | **bool** | Is share protected by password | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date | [default to null]
**Name** | **string** | Share display name (alias name) | [optional] [default to null]
**IsEncrypted** | **bool** | Encryption state | [optional] [default to null]
**ExpireAt** | [**time.Time**](time.Time.md) | Expiration date | [optional] [default to null]
**Notes** | **string** | User notes | [optional] [default to null]
**UploadedFiles** | [**[]PublicUploadedFileData**](PublicUploadedFileData.md) | List of (public) uploaded files | [optional] [default to null]
**UserUserPublicKeyList** | [***UserUserPublicKeyList**](UserUserPublicKeyList.md) |  | [optional] [default to null]
**ShowUploadedFiles** | **bool** | Allow display of already uploaded files | [optional] [default to null]
**RemainingSize** | **int64** | Remaining size | [optional] [default to null]
**RemainingSlots** | **int32** | Remaining slots | [optional] [default to null]
**CreatorName** | **string** | &amp;#128640; Since v4.11.0  Creator name | [default to null]
**CreatorUsername** | **string** | &amp;#128640; Since v4.11.0  Creator username | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

