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

// Deleted node information (Deleted node can be a folder or file)
type DeletedNode struct {
	// Node ID
	Id int64 `json:"id,omitempty"`
	// Parent node ID (room or folder)
	ParentId int64 `json:"parentId"`
	// Parent node path  `/` if node is a root node (room)
	ParentPath string `json:"parentPath"`
	// Node type
	Type_ string `json:"type"`
	// Node name
	Name string `json:"name"`
	// Expiration date
	ExpireAt time.Time `json:"expireAt,omitempty"`
	// Last access date
	AccessedAt time.Time `json:"accessedAt,omitempty"`
	// Encryption state
	IsEncrypted bool `json:"isEncrypted,omitempty"`
	// User notes
	Notes string `json:"notes,omitempty"`
	// Node size in byte
	Size int64 `json:"size,omitempty"`
	// Classification ID:  * `1` - public  * `2` - internal  * `3` - confidential  * `4` - strictly confidential    (default: classification from parent room)
	Classification int32 `json:"classification,omitempty"`
	// Creation date
	CreatedAt time.Time `json:"createdAt,omitempty"`
	CreatedBy *UserInfo `json:"createdBy,omitempty"`
	// Modification date
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UpdatedBy *UserInfo `json:"updatedBy,omitempty"`
	// Deletion date
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	DeletedBy *UserInfo `json:"deletedBy,omitempty"`
}
