# DeletedNodeSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | **int64** | Parent node ID (room or folder) | [default to null]
**ParentPath** | **string** | Parent node path  &#x60;/&#x60; if node is a root node (room) | [default to null]
**Name** | **string** | Node name | [default to null]
**Type_** | **string** | Node type | [default to null]
**CntVersions** | **int32** | Number of deleted versions of this file | [default to null]
**FirstDeletedAt** | [**time.Time**](time.Time.md) | First deleted version | [default to null]
**LastDeletedAt** | [**time.Time**](time.Time.md) | Last deleted version | [default to null]
**LastDeletedNodeId** | **int64** | Node ID of last deleted version | [default to null]
**TimestampCreation** | [**time.Time**](time.Time.md) | &amp;#128640; Since v4.22.0  Time the node was created on external file system | [optional] [default to null]
**TimestampModification** | [**time.Time**](time.Time.md) | &amp;#128640; Since v4.22.0  Time the content of a node was last modified on external file system | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

