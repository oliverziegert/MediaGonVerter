# LastAdminUserRoom

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Room ID | [default to null]
**Name** | **string** | Room name | [default to null]
**ParentPath** | **string** | Parent node path  &#x60;/&#x60; if node is a root node (room) | [default to null]
**LastAdminInGroup** | **bool** | Determines whether user is last admin of a room due to being the last member of last admin group | [default to null]
**ParentId** | **int64** | Parent room ID | [optional] [default to null]
**LastAdminInGroupId** | **int64** | ID of the last admin group where the user is the only remaining member  (returned only if &#x60;lastAdminInGroup&#x60; is &#x60;true&#x60;) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

