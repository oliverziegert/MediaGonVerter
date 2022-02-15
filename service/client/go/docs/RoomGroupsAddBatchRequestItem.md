# RoomGroupsAddBatchRequestItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the group | [default to null]
**Permissions** | [***NodePermissions**](NodePermissions.md) |  | [default to null]
**NewGroupMemberAcceptance** | **string** | Behaviour when new users are added to the group:  * &#x60;autoallow&#x60;  * &#x60;pending&#x60;    Only relevant if &#x60;adminGroupIds&#x60; has items. | [optional] [default to NEW_GROUP_MEMBER_ACCEPTANCE.AUTOALLOW]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

