# CreateFolderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | **int64** | Parent node ID (room or folder) | [default to null]
**Name** | **string** | Name | [default to null]
**Notes** | **string** | User notes  Use empty string to remove. | [optional] [default to null]
**Classification** | **int32** | &amp;#128640; Since v4.30.0  Classification ID:  * &#x60;1&#x60; - public  * &#x60;2&#x60; - internal  * &#x60;3&#x60; - confidential  * &#x60;4&#x60; - strictly confidential    Provided (or default) classification is taken from room  when file gets uploaded without any classification. | [optional] [default to null]
**TimestampCreation** | [**time.Time**](time.Time.md) | &amp;#128640; Since v4.22.0  Time the node was created on external file system  (default: current server datetime in UTC format) | [optional] [default to null]
**TimestampModification** | [**time.Time**](time.Time.md) | &amp;#128640; Since v4.22.0  Time the content of a node was last modified on external file system  (default: current server datetime in UTC format) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

