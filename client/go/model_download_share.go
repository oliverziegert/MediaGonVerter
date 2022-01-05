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

// Download Share information
type DownloadShare struct {
	// Share ID
	Id int64 `json:"id"`
	// Alias name
	Name string `json:"name"`
	// Source node ID
	NodeId int64 `json:"nodeId"`
	// Share access key to generate secure link
	AccessKey string `json:"accessKey"`
	// Downloads counter (incremented on each download)
	CntDownloads int32 `json:"cntDownloads"`
	// Creation date
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy *UserInfo `json:"createdBy"`
	// Modification date
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UpdatedBy *UserInfo `json:"updatedBy,omitempty"`
	// User notes
	Notes string `json:"notes,omitempty"`
	// &#128640; Since v4.11.0  Internal notes
	InternalNotes string `json:"internalNotes,omitempty"`
	// Show creator first and last name.
	ShowCreatorName bool `json:"showCreatorName,omitempty"`
	// Show creator email address.
	ShowCreatorUsername bool `json:"showCreatorUsername,omitempty"`
	// Is share protected by password
	IsProtected bool `json:"isProtected,omitempty"`
	// Expiration date
	ExpireAt time.Time `json:"expireAt,omitempty"`
	// Max allowed downloads
	MaxDownloads int32 `json:"maxDownloads,omitempty"`
	// Path to shared download node
	NodePath string `json:"nodePath,omitempty"`
	// Path to shared download node
	DataUrl string `json:"dataUrl,omitempty"`
	// Encrypted share  (this only applies to shared files, not folders)
	IsEncrypted bool `json:"isEncrypted,omitempty"`
	// Node type
	NodeType string `json:"nodeType,omitempty"`
	// &#128679; Deprecated since v4.11.0  Classification ID:  * `1` - public  * `2` - internal  * `3` - confidential  * `4` - strictly confidential    (default: classification from parent room)
	Classification int32 `json:"classification,omitempty"`
	// &#128679; Deprecated since v4.20.0  Notify creator on every download.
	NotifyCreator bool `json:"notifyCreator"`
	// &#128679; Deprecated since v4.11.0  CSV string of recipient email addresses
	Recipients string `json:"recipients,omitempty"`
	// &#128679; Deprecated since v4.11.0  CSV string of recipient MSISDNs
	SmsRecipients string `json:"smsRecipients,omitempty"`
}
