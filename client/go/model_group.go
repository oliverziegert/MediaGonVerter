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

// Group information
type Group struct {
	// Unique identifier for the group
	Id int64 `json:"id"`
	// Group name
	Name string `json:"name"`
	// Creation date
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy *UserInfo `json:"createdBy"`
	// Amount of users
	CntUsers int32 `json:"cntUsers"`
	// Modification date
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UpdatedBy *UserInfo `json:"updatedBy,omitempty"`
	// Expiration date
	ExpireAt   time.Time `json:"expireAt,omitempty"`
	GroupRoles *RoleList `json:"groupRoles,omitempty"`
}
