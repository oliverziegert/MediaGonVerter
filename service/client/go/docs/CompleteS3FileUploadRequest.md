# CompleteS3FileUploadRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parts** | [**[]S3FileUploadPart**](S3FileUploadPart.md) | List of S3 file upload parts | [default to null]
**ResolutionStrategy** | **string** | Node conflict resolution strategy:  * &#x60;autorename&#x60;  * &#x60;overwrite&#x60;  * &#x60;fail&#x60; | [optional] [default to RESOLUTION_STRATEGY.AUTORENAME]
**KeepShareLinks** | **bool** | Preserve Download Share Links and point them to the new node. | [optional] [default to false]
**FileName** | **string** | New file name to store with | [optional] [default to null]
**FileKey** | [***FileKey**](FileKey.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

