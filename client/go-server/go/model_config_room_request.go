/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Request model for configuring a room
type ConfigRoomRequest struct {
	// Retention period for deleted nodes in days
	RecycleBinRetentionPeriod int32 `json:"recycleBinRetentionPeriod,omitempty"`
	// Inherit permissions from parent room  (default: `false` if `parentId` is `0`; otherwise: `true`)
	InheritPermissions bool `json:"inheritPermissions,omitempty"`
	// Take over existing permissions
	TakeOverPermissions bool `json:"takeOverPermissions,omitempty"`
	// List of user ids  A room requires at least one admin (user or group)
	AdminIds []int64 `json:"adminIds,omitempty"`
	// List of group ids  A room requires at least one admin (user or group)
	AdminGroupIds []int64 `json:"adminGroupIds,omitempty"`
	// Behaviour when new users are added to the group:  * `autoallow`  * `pending`    Only relevant if `adminGroupIds` has items.
	NewGroupMemberAcceptance string `json:"newGroupMemberAcceptance,omitempty"`
	// Is activities log active (for rooms only)
	HasActivitiesLog bool `json:"hasActivitiesLog,omitempty"`
	// Classification ID:  * `1` - public  * `2` - internal  * `3` - confidential  * `4` - strictly confidential    Provided (or default) classification is taken from room  when file gets uploaded without any classification.
	Classification int32 `json:"classification,omitempty"`
	// &#128679; Deprecated since v4.10.0  Is recycle bin active (for rooms only)  Recycle bin is always on (disabling is not possible).
	HasRecycleBin bool `json:"hasRecycleBin,omitempty"`
}
