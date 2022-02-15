# RoomData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Node type | [optional] [default to null]
**Id** | **int64** | Room ID | [default to null]
**IsGranted** | **bool** | Is user granted room permissions | [default to null]
**Name** | **string** | Name | [default to null]
**IsEncrypted** | **bool** | Encryption state | [default to null]
**RecycleBinRetentionPeriod** | **int32** | Retention period for deleted nodes in days | [default to null]
**ParentId** | **int64** | Parent node ID (room or folder) | [optional] [default to null]
**Size** | **int64** | Room size | [optional] [default to null]
**Permissions** | [***NodePermissions**](NodePermissions.md) |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Expiration date | [optional] [default to null]
**CreatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Modification date | [optional] [default to null]
**UpdatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**Quota** | **int64** | Quota in byte | [optional] [default to null]
**CntDownloadShares** | **int32** | Returns the number of Download Shares of this node. | [optional] [default to null]
**CntUploadShares** | **int32** | Returns the number of Upload Shares of this node. | [optional] [default to null]
**IsFavorite** | **bool** | Node is marked as favorite (for rooms / folders only) | [optional] [default to null]
**HasRecycleBin** | **bool** | &amp;#128679; Deprecated since v4.10.0  Is recycle bin active (for rooms only)  Recycle bin is always on (disabling is not possible). | [default to null]
**Children** | [**[]RoomData**](RoomData.md) | &amp;#128679; Deprecated since v4.10.0  List of rooms, where this room is a parent (if exist) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

