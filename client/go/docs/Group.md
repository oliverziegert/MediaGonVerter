# Group

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the group | [default to null]
**Name** | **string** | Group name | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date | [default to null]
**CreatedBy** | [***UserInfo**](UserInfo.md) |  | [default to null]
**CntUsers** | **int32** | Amount of users | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Modification date | [optional] [default to null]
**UpdatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**ExpireAt** | [**time.Time**](time.Time.md) | Expiration date | [optional] [default to null]
**GroupRoles** | [***RoleList**](RoleList.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

