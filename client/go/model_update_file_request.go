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

// Request model for updating file's metadata
type UpdateFileRequest struct {
	// File name
	Name       string            `json:"name,omitempty"`
	Expiration *ObjectExpiration `json:"expiration,omitempty"`
	// Classification ID:  * `1` - public  * `2` - internal  * `3` - confidential  * `4` - strictly confidential
	Classification int32 `json:"classification,omitempty"`
	// User notes  Use empty string to remove.
	Notes string `json:"notes,omitempty"`
	// &#128640; Since v4.22.0  Time the node was created on external file system  (default: current server datetime in UTC format)
	TimestampCreation time.Time `json:"timestampCreation,omitempty"`
	// &#128640; Since v4.22.0  Time the content of a node was last modified on external file system  (default: current server datetime in UTC format)
	TimestampModification time.Time `json:"timestampModification,omitempty"`
}
