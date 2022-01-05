# LogEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Event ID | [default to null]
**Time** | [**time.Time**](time.Time.md) | Event timestamp | [default to null]
**UserId** | **int64** | Unique identifier for the user | [default to null]
**Message** | **string** | Event description | [default to null]
**OperationId** | **int32** | Operation type ID | [optional] [default to null]
**OperationName** | **string** | Operation name | [optional] [default to null]
**Status** | **int32** | Operation status:  * &#x60;0&#x60; - Success  * &#x60;2&#x60; - Error | [optional] [default to null]
**UserClient** | **string** | Client | [optional] [default to null]
**CustomerId** | **int64** | Unique identifier for the customer | [optional] [default to null]
**UserName** | **string** | Username | [optional] [default to null]
**UserIp** | **string** | User IP | [optional] [default to null]
**AuthParentSource** | **string** | Auth parent source ID | [optional] [default to null]
**AuthParentTarget** | **string** | Auth parent target ID | [optional] [default to null]
**ObjectId1** | **int64** | Object ID 1 | [optional] [default to null]
**ObjectType1** | **int32** | Object type 1 | [optional] [default to null]
**ObjectName1** | **string** | Object name 1 | [optional] [default to null]
**ObjectId2** | **int64** | Object ID 2 | [optional] [default to null]
**ObjectType2** | **int32** | Object type 2 | [optional] [default to null]
**ObjectName2** | **string** | Object name 2 | [optional] [default to null]
**Attribute1** | **string** | Attribute 1 | [optional] [default to null]
**Attribute2** | **string** | Attribute 2 | [optional] [default to null]
**Attribute3** | **string** | Attribute 3 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

