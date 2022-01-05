# CreateShareUploadChannelRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | File name | [default to null]
**Size** | **int64** | File size in byte | [optional] [default to null]
**Password** | **string** | Password | [optional] [default to null]
**DirectS3Upload** | **bool** | &amp;#128640; Since v4.15.0  Upload direct to S3 | [optional] [default to false]
**TimestampCreation** | [**time.Time**](time.Time.md) | &amp;#128640; Since v4.22.0  Time the node was created on external file system  (default: current server datetime in UTC format) | [optional] [default to null]
**TimestampModification** | [**time.Time**](time.Time.md) | &amp;#128640; Since v4.22.0  Time the content of a node was last modified on external file system  (default: current server datetime in UTC format) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

