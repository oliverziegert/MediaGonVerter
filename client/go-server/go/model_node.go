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

// Node information (Node can be a room, folder or file)
type Node struct {
	// Node ID
	Id int64 `json:"id"`
	// Node type
	Type_ string `json:"type"`
	// Name
	Name string `json:"name"`
	// &#128640; Since v4.22.0  Time the node was created on external file system
	TimestampCreation time.Time `json:"timestampCreation,omitempty"`
	// &#128640; Since v4.22.0  Time the content of a node was last modified on external file system
	TimestampModification time.Time `json:"timestampModification,omitempty"`
	// Parent node ID (room or folder)
	ParentId int64 `json:"parentId,omitempty"`
	// Parent node path  `/` if node is a root node (room)
	ParentPath string `json:"parentPath,omitempty"`
	// Creation date
	CreatedAt time.Time `json:"createdAt,omitempty"`

	CreatedBy *UserInfo `json:"createdBy,omitempty"`
	// Modification date
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	UpdatedBy *UserInfo `json:"updatedBy,omitempty"`
	// Expiration date
	ExpireAt time.Time `json:"expireAt,omitempty"`
	// MD5 hash of file
	Hash string `json:"hash,omitempty"`
	// File type / extension (for files only)
	FileType string `json:"fileType,omitempty"`
	// File media type (for files only)
	MediaType string `json:"mediaType,omitempty"`
	// Node size in byte
	Size int64 `json:"size,omitempty"`
	// Classification ID:  * `1` - public  * `2` - internal  * `3` - confidential  * `4` - strictly confidential
	Classification int32 `json:"classification,omitempty"`
	// User notes
	Notes string `json:"notes,omitempty"`

	Permissions *NodePermissions `json:"permissions,omitempty"`
	// Inherit permissions from parent room  (default: `false` if `parentId` is `0`; otherwise: `true`)
	InheritPermissions bool `json:"inheritPermissions,omitempty"`
	// Encryption state
	IsEncrypted bool `json:"isEncrypted,omitempty"`

	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`
	// Number of deleted versions of this file / folder  (for rooms / folders only)
	CntDeletedVersions int32 `json:"cntDeletedVersions,omitempty"`
	// Returns the number of comments of this node.
	CntComments int32 `json:"cntComments,omitempty"`
	// Returns the number of Download Shares of this node.
	CntDownloadShares int32 `json:"cntDownloadShares,omitempty"`
	// Returns the number of Upload Shares of this node.
	CntUploadShares int32 `json:"cntUploadShares,omitempty"`
	// Retention period for deleted nodes in days
	RecycleBinRetentionPeriod int32 `json:"recycleBinRetentionPeriod,omitempty"`
	// Is activities log active (for rooms only)
	HasActivitiesLog bool `json:"hasActivitiesLog,omitempty"`
	// Quota in byte
	Quota int64 `json:"quota,omitempty"`
	// Node is marked as favorite (for rooms / folders only)
	IsFavorite bool `json:"isFavorite,omitempty"`
	// Version of last change in this node or a node further down the tree.
	BranchVersion int64 `json:"branchVersion,omitempty"`
	// Media server media token
	MediaToken string `json:"mediaToken,omitempty"`
	// &#128640; Since v4.11.0  Determines whether node is browsable by client (for rooms only)
	IsBrowsable bool `json:"isBrowsable,omitempty"`
	// &#128640; Since v4.11.0  Amount of direct child rooms where this node is the parent node  (no recursion; for rooms only)
	CntRooms int32 `json:"cntRooms,omitempty"`
	// &#128640; Since v4.11.0  Amount of direct child folders where this node is the parent node  (no recursion; for rooms / folders only)
	CntFolders int32 `json:"cntFolders,omitempty"`
	// &#128640; Since v4.11.0  Amount of direct child files where this node is the parent node  (no recursion; for rooms / folders only)
	CntFiles int32 `json:"cntFiles,omitempty"`
	// &#128640; Since v4.15.0  Auth parent room ID
	AuthParentId int64 `json:"authParentId,omitempty"`
	// &#128679; Deprecated since v4.11.0  Number of direct children  (no recursion; for rooms / folders only)
	CntChildren int32 `json:"cntChildren,omitempty"`
	// &#128679; Deprecated since v4.10.0  Is recycle bin active (for rooms only)  Recycle bin is always on (disabling is not possible).
	HasRecycleBin bool `json:"hasRecycleBin,omitempty"`
	// &#128679; Deprecated since v4.10.0  Child nodes list (if requested)  (for rooms / folders only)
	Children []Node `json:"children,omitempty"`
}
