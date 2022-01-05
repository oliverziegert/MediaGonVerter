# UpdateUploadSharesBulkRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | [***ObjectExpiration**](ObjectExpiration.md) |  | [optional] [default to null]
**ShowCreatorName** | **bool** | Show creator first and last name. | [optional] [default to null]
**ShowCreatorUsername** | **bool** | Show creator email address. | [optional] [default to null]
**ShowUploadedFiles** | **bool** | Allow display of already uploaded files | [optional] [default to null]
**MaxSlots** | **int32** | Maximal amount of files to upload | [optional] [default to null]
**ResetMaxSlots** | **bool** | Set &#x27;true&#x27; to reset &#x27;maxSlots&#x27; for Upload Share | [optional] [default to null]
**MaxSize** | **int64** | Maximal total size of uploaded files (in bytes) | [optional] [default to null]
**ResetMaxSize** | **bool** | Set &#x27;true&#x27; to reset &#x27;maxSize&#x27; for Upload Share | [optional] [default to null]
**FilesExpiryPeriod** | **int32** | Number of days after which uploaded files expire | [optional] [default to null]
**ResetFilesExpiryPeriod** | **bool** | Set &#x27;true&#x27; to reset &#x27;filesExpiryPeriod&#x27; for Upload Share | [optional] [default to null]
**ObjectIds** | **[]int64** | List of ids | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

