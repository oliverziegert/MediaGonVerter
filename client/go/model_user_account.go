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

// User information
type UserAccount struct {
	// Unique identifier for the user
	Id int64 `json:"id"`
	// &#128640; Since v4.13.0  Username
	UserName string `json:"userName"`
	// User first name
	FirstName string `json:"firstName"`
	// User last name
	LastName string `json:"lastName"`
	// User is locked:  * `false` - unlocked  * `true` - locked    User is locked and can not login anymore.
	IsLocked bool `json:"isLocked"`
	// User has manageable rooms
	HasManageableRooms bool      `json:"hasManageableRooms"`
	UserRoles          *RoleList `json:"userRoles"`
	// &#128640; Since v4.20.0  IETF language tag
	Language string        `json:"language"`
	AuthData *UserAuthData `json:"authData"`
	// &#128640; Since v4.13.0  If `true`, the user must set the `email` at the first login.
	MustSetEmail bool `json:"mustSetEmail,omitempty"`
	// User has accepted EULA.  Present, if EULA is system global active.  cf. `GET system/config/settings/general` - `eulaEnabled`
	NeedsToAcceptEULA bool `json:"needsToAcceptEULA,omitempty"`
	// Expiration date
	ExpireAt time.Time `json:"expireAt,omitempty"`
	// User has generated private key.  Possible if client-side encryption is active for this customer
	IsEncryptionEnabled bool `json:"isEncryptionEnabled,omitempty"`
	// Last successful logon date
	LastLoginSuccessAt time.Time `json:"lastLoginSuccessAt,omitempty"`
	// Last failed logon date
	LastLoginFailAt time.Time `json:"lastLoginFailAt,omitempty"`
	// Email
	Email string `json:"email,omitempty"`
	// Phone number
	Phone string `json:"phone,omitempty"`
	// Homeroom ID
	HomeRoomId int64 `json:"homeRoomId,omitempty"`
	// All groups the user is member of
	UserGroups     []UserGroup     `json:"userGroups,omitempty"`
	UserAttributes *UserAttributes `json:"userAttributes,omitempty"`
	// &#128679; Deprecated since v4.18.0  Job title
	Title string `json:"title,omitempty"`
	// &#128679; Deprecated since v4.6.0  Last successful logon IP address
	LastLoginSuccessIp string `json:"lastLoginSuccessIp,omitempty"`
	// &#128679; Deprecated since v4.6.0  Last failed logon IP address
	LastLoginFailIp string `json:"lastLoginFailIp,omitempty"`
	// &#128679; Deprecated since v4.12.0  Gender
	Gender string `json:"gender,omitempty"`
	// &#128679; Deprecated since v4.13.0  If `true`, the user must change the `userName` at the first login.
	NeedsToChangeUserName bool `json:"needsToChangeUserName,omitempty"`
	// &#128679; Deprecated since v4.13.0  Authentication methods:  * `sql`  * `active_directory`  * `radius`  * `openid`  use `authData` instead
	AuthMethods []UserAuthMethod `json:"authMethods,omitempty"`
	// &#128679; Deprecated since v4.13.0  User login name
	Login string `json:"login,omitempty"`
	// &#128679; Deprecated since v4.7.0  User lock status:  * `0` - locked  * `1` - Web access allowed  * `2` - Web and mobile access allowed    Please use `isLocked` instead.
	LockStatus int32 `json:"lockStatus"`
	// &#128679; Deprecated since v4.13.0  Determines whether user has to change his / her password
	NeedsToChangePassword bool `json:"needsToChangePassword"`
}
