/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// First administrator user
type FirstAdminUser struct {
	// User first name
	FirstName string `json:"firstName"`
	// User last name
	LastName string `json:"lastName"`
	// &#128640; Since v4.13.0  Username
	UserName string        `json:"userName,omitempty"`
	AuthData *UserAuthData `json:"authData,omitempty"`
	// IETF language tag
	ReceiverLanguage string `json:"receiverLanguage,omitempty"`
	// Notify user about his new account  * default: `true` for `basic` auth type  * default: `false` for `active_directory`, `openid` and `radius` auth types
	NotifyUser bool `json:"notifyUser,omitempty"`
	// Email
	Email string `json:"email,omitempty"`
	// Phone number
	Phone string `json:"phone,omitempty"`
	// &#128679; Deprecated since v4.18.0  Job title
	Title string `json:"title,omitempty"`
	// &#128679; Deprecated since v4.7.0  Language ID or ISO 639-1 code
	Language string `json:"language,omitempty"`
	// &#128679; Deprecated since v4.13.0  Authentication methods:  * `sql`  * `active_directory`  * `radius`  * `openid`  use `authData` instead
	AuthMethods []UserAuthMethod `json:"authMethods,omitempty"`
	// &#128679; Deprecated since v4.13.0  If `true`, the user must change the `userName` at the first login.
	NeedsToChangeUserName bool `json:"needsToChangeUserName,omitempty"`
	// &#128679; Deprecated since v4.13.0  An initial password may be preset  use `authData` instead
	Password string `json:"password,omitempty"`
	// &#128679; Deprecated since v4.13.0  Determines whether user has to change his / her initial password.  use `authDate.mustChangePassword` instead
	NeedsToChangePassword bool `json:"needsToChangePassword,omitempty"`
	// &#128679; Deprecated since v4.13.0  User login name
	Login string `json:"login,omitempty"`
	// &#128679; Deprecated since v4.12.0  Gender
	Gender string `json:"gender,omitempty"`
}
