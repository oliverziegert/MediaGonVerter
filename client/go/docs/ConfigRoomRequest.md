# ConfigRoomRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecycleBinRetentionPeriod** | **int32** | Retention period for deleted nodes in days | [optional] [default to null]
**InheritPermissions** | **bool** | Inherit permissions from parent room  (default: &#x60;false&#x60; if &#x60;parentId&#x60; is &#x60;0&#x60;; otherwise: &#x60;true&#x60;) | [optional] [default to null]
**TakeOverPermissions** | **bool** | Take over existing permissions | [optional] [default to null]
**AdminIds** | **[]int64** | List of user ids  A room requires at least one admin (user or group) | [optional] [default to null]
**AdminGroupIds** | **[]int64** | List of group ids  A room requires at least one admin (user or group) | [optional] [default to null]
**NewGroupMemberAcceptance** | **string** | Behaviour when new users are added to the group:  * &#x60;autoallow&#x60;  * &#x60;pending&#x60;    Only relevant if &#x60;adminGroupIds&#x60; has items. | [optional] [default to NEW_GROUP_MEMBER_ACCEPTANCE.AUTOALLOW]
**HasActivitiesLog** | **bool** | Is activities log active (for rooms only) | [optional] [default to true]
**Classification** | **int32** | Classification ID:  * &#x60;1&#x60; - public  * &#x60;2&#x60; - internal  * &#x60;3&#x60; - confidential  * &#x60;4&#x60; - strictly confidential    Provided (or default) classification is taken from room  when file gets uploaded without any classification. | [optional] [default to CLASSIFICATION.2_]
**HasRecycleBin** | **bool** | &amp;#128679; Deprecated since v4.10.0  Is recycle bin active (for rooms only)  Recycle bin is always on (disabling is not possible). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

