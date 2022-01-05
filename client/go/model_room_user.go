/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// User information
type RoomUser struct {
	UserInfo *UserInfo `json:"userInfo"`
	// Is user granted room permissions
	IsGranted          bool                `json:"isGranted"`
	Permissions        *NodePermissions    `json:"permissions,omitempty"`
	PublicKeyContainer *PublicKeyContainer `json:"publicKeyContainer,omitempty"`
	// &#128679; Deprecated since v4.11.0  Unique identifier for the user  use `id` from `UserInfo` instead
	Id int64 `json:"id"`
	// &#128679; Deprecated since v4.11.0  User login name
	Login string `json:"login"`
	// &#128679; Deprecated since v4.11.0  Display name  use information from `UserInfo` instead to combine a display name
	DisplayName string `json:"displayName"`
	// &#128679; Deprecated since v4.11.0  Email   use `email` from `UserInfo` instead
	Email string `json:"email"`
}
