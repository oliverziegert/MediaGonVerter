# DeletedNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Node ID | [optional] [default to null]
**ParentId** | **int64** | Parent node ID (room or folder) | [default to null]
**ParentPath** | **string** | Parent node path  &#x60;/&#x60; if node is a root node (room) | [default to null]
**Type_** | **string** | Node type | [default to null]
**Name** | **string** | Node name | [default to null]
**ExpireAt** | [**time.Time**](time.Time.md) | Expiration date | [optional] [default to null]
**AccessedAt** | [**time.Time**](time.Time.md) | Last access date | [optional] [default to null]
**IsEncrypted** | **bool** | Encryption state | [optional] [default to null]
**Notes** | **string** | User notes | [optional] [default to null]
**Size** | **int64** | Node size in byte | [optional] [default to null]
**Classification** | **int32** | Classification ID:  * &#x60;1&#x60; - public  * &#x60;2&#x60; - internal  * &#x60;3&#x60; - confidential  * &#x60;4&#x60; - strictly confidential    (default: classification from parent room) | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date | [optional] [default to null]
**CreatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Modification date | [optional] [default to null]
**UpdatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**DeletedAt** | [**time.Time**](time.Time.md) | Deletion date | [optional] [default to null]
**DeletedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

