# GeneratePresignedUrlsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int64** | &#x60;Content-Length&#x60; header size for each presigned URL (in bytes)  *MUST* be &gt;&#x3D; 5 MB except the last part. | [default to null]
**FirstPartNumber** | **int32** | First part number of a range of requested presigned URLs (for S3 it is: &#x60;1&#x60;) | [default to null]
**LastPartNumber** | **int32** | Last part number of a range of requested presigned URLs | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

