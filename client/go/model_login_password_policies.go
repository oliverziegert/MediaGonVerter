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

// Login password policies
type LoginPasswordPolicies struct {
	CharacterRules *CharacterRules `json:"characterRules"`
	// Minimum number of characters a password must contain
	MinLength int32 `json:"minLength"`
	// Determines whether a password must NOT contain word(s) from a dictionary
	RejectDictionaryWords bool `json:"rejectDictionaryWords"`
	// Determines whether a password must NOT contain user info (first name, last name, email, user name)
	RejectUserInfo bool `json:"rejectUserInfo"`
	// Determines whether a password must NOT contain keyboard patterns (e.g. `qwertz`, `asdf`)  (min. 4 character pattern)
	RejectKeyboardPatterns bool `json:"rejectKeyboardPatterns"`
	// Number of passwords to archive  (must be between `0` and `10`; `0` means that password history is disabled)
	NumberOfArchivedPasswords int32               `json:"numberOfArchivedPasswords"`
	PasswordExpiration        *PasswordExpiration `json:"passwordExpiration"`
	UserLockout               *UserLockout        `json:"userLockout"`
	// Modification date
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy *UserInfo `json:"updatedBy"`
}
