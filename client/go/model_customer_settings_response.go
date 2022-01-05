/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Customer settings
type CustomerSettingsResponse struct {
	// Homerooms active
	HomeRoomsActive bool `json:"homeRoomsActive"`
	// Homeroom Parent ID
	HomeRoomParentId int64 `json:"homeRoomParentId,omitempty"`
	// Homeroom Parent Name
	HomeRoomParentName string `json:"homeRoomParentName,omitempty"`
	// Homeroom Quota in bytes
	HomeRoomQuota int64 `json:"homeRoomQuota,omitempty"`
}
