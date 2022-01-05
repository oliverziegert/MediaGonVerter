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

// Upload Share information
type PublicUploadShare struct {
	// Is share protected by password
	IsProtected bool `json:"isProtected"`
	// Creation date
	CreatedAt time.Time `json:"createdAt"`
	// Share display name (alias name)
	Name string `json:"name,omitempty"`
	// Encryption state
	IsEncrypted bool `json:"isEncrypted,omitempty"`
	// Expiration date
	ExpireAt time.Time `json:"expireAt,omitempty"`
	// User notes
	Notes string `json:"notes,omitempty"`
	// List of (public) uploaded files
	UploadedFiles         []PublicUploadedFileData `json:"uploadedFiles,omitempty"`
	UserUserPublicKeyList *UserUserPublicKeyList   `json:"userUserPublicKeyList,omitempty"`
	// Allow display of already uploaded files
	ShowUploadedFiles bool `json:"showUploadedFiles,omitempty"`
	// Remaining size
	RemainingSize int64 `json:"remainingSize,omitempty"`
	// Remaining slots
	RemainingSlots int32 `json:"remainingSlots,omitempty"`
	// &#128640; Since v4.11.0  Creator name
	CreatorName string `json:"creatorName"`
	// &#128640; Since v4.11.0  Creator username
	CreatorUsername string `json:"creatorUsername,omitempty"`
}
