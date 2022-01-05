# CreateFileUploadRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | **int64** | Parent node ID (room or folder) | [default to null]
**Name** | **string** | File name | [default to null]
**Classification** | **int32** | Classification ID:  * &#x60;1&#x60; - public  * &#x60;2&#x60; - internal  * &#x60;3&#x60; - confidential  * &#x60;4&#x60; - strictly confidential    (default: classification from parent room) | [optional] [default to null]
**Size** | **int64** | File size in byte | [optional] [default to null]
**Expiration** | [***ObjectExpiration**](ObjectExpiration.md) |  | [optional] [default to null]
**Notes** | **string** | User notes  Use empty string to remove. | [optional] [default to null]
**DirectS3Upload** | **bool** | &amp;#128640; Since v4.15.0  Upload direct to S3 | [optional] [default to false]
**TimestampCreation** | [**time.Time**](time.Time.md) | &amp;#128640; Since v4.22.0  Time the node was created on external file system  (default: current server datetime in UTC format) | [optional] [default to null]
**TimestampModification** | [**time.Time**](time.Time.md) | &amp;#128640; Since v4.22.0  Time the content of a node was last modified on external file system  (default: current server datetime in UTC format) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

