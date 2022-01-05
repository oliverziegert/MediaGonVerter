# Webhook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID | [default to null]
**Name** | **string** | Name | [default to null]
**Url** | **string** | URL | [default to null]
**Secret** | **string** | Secret; used for event message signatures | [optional] [default to null]
**IsEnabled** | **bool** | Is enabled | [default to null]
**ExpireAt** | [**time.Time**](time.Time.md) | Expiration date / time | [default to null]
**EventTypeNames** | **[]string** | List of names of event types | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date | [default to null]
**CreatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Modification date | [default to null]
**UpdatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**FailStatus** | **int32** | Last HTTP status code when a webhook is disabled due to delivery failures | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

