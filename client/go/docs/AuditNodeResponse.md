# AuditNodeResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **int64** | Node ID | [default to null]
**NodeName** | **string** | Node name | [default to null]
**NodeParentPath** | **string** | Parent node path  &#x60;/&#x60; if node is a root node (room) | [default to null]
**NodeCntChildren** | **int32** | Number of direct children  (no recursion; for rooms only) | [default to null]
**AuditUserPermissionList** | [**[]AuditUserPermission**](AuditUserPermission.md) | List of assigned users with permissions | [default to null]
**NodeParentId** | **int64** | Parent node ID (room or folder) | [optional] [default to null]
**NodeSize** | **int64** | Node size in byte | [optional] [default to null]
**NodeRecycleBinRetentionPeriod** | **int32** | Retention period for deleted nodes in days | [optional] [default to null]
**NodeQuota** | **int64** | Quota in byte | [optional] [default to null]
**NodeIsEncrypted** | **bool** | Encryption state | [optional] [default to null]
**NodeHasActivitiesLog** | **bool** | Is activities log active (for rooms only) | [optional] [default to true]
**NodeCreatedAt** | [**time.Time**](time.Time.md) | Creation date | [optional] [default to null]
**NodeCreatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**NodeUpdatedAt** | [**time.Time**](time.Time.md) | Modification date | [optional] [default to null]
**NodeUpdatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**NodeHasRecycleBin** | **bool** | &amp;#128679; Deprecated since v4.10.0  Is recycle bin active (for rooms only)  Recycle bin is always on (disabling is not possible). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

