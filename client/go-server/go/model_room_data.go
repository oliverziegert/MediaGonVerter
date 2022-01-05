/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"time"
)

// Room information
type RoomData struct {
	// Node type
	Type_ string `json:"type,omitempty"`
	// Room ID
	Id int64 `json:"id"`
	// Is user granted room permissions
	IsGranted bool `json:"isGranted"`
	// Name
	Name string `json:"name"`
	// Encryption state
	IsEncrypted bool `json:"isEncrypted"`
	// Retention period for deleted nodes in days
	RecycleBinRetentionPeriod int32 `json:"recycleBinRetentionPeriod"`
	// Parent node ID (room or folder)
	ParentId int64 `json:"parentId,omitempty"`
	// Room size
	Size int64 `json:"size,omitempty"`

	Permissions *NodePermissions `json:"permissions,omitempty"`
	// Expiration date
	CreatedAt time.Time `json:"createdAt,omitempty"`

	CreatedBy *UserInfo `json:"createdBy,omitempty"`
	// Modification date
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	UpdatedBy *UserInfo `json:"updatedBy,omitempty"`
	// Quota in byte
	Quota int64 `json:"quota,omitempty"`
	// Returns the number of Download Shares of this node.
	CntDownloadShares int32 `json:"cntDownloadShares,omitempty"`
	// Returns the number of Upload Shares of this node.
	CntUploadShares int32 `json:"cntUploadShares,omitempty"`
	// Node is marked as favorite (for rooms / folders only)
	IsFavorite bool `json:"isFavorite,omitempty"`
	// &#128679; Deprecated since v4.10.0  Is recycle bin active (for rooms only)  Recycle bin is always on (disabling is not possible).
	HasRecycleBin bool `json:"hasRecycleBin"`
	// &#128679; Deprecated since v4.10.0  List of rooms, where this room is a parent (if exist)
	Children []RoomData `json:"children,omitempty"`
}
