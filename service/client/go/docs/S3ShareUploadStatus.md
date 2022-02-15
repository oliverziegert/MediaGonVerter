# S3ShareUploadStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | S3 file upload status:  * &#x60;transfer&#x60; - upload in progress  * &#x60;finishing&#x60; - completing file upload  * &#x60;done&#x60; - file upload successully done  * &#x60;error&#x60; - an error occurred while file upload | [default to null]
**FileName** | **string** | File name | [default to null]
**Size** | **int64** | File size in byte | [optional] [default to null]
**ErrorDetails** | [***ErrorResponse**](ErrorResponse.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

