/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Request model for updating general settings
type UpdateGeneralSettings struct {
	// Allow sending of share passwords via SMS
	SharePasswordSmsEnabled bool `json:"sharePasswordSmsEnabled,omitempty"`
	// Activation status of client-side encryption.  Can only be enabled once; disabling is not possible.
	CryptoEnabled bool `json:"cryptoEnabled,omitempty"`
	// Enable email notification button
	EmailNotificationButtonEnabled bool `json:"emailNotificationButtonEnabled,omitempty"`
	// Each user has to confirm the EULA at first login.
	EulaEnabled bool `json:"eulaEnabled,omitempty"`
	// &#128640; Since v4.9.0  Defines if S3 tags are enabled
	S3TagsEnabled bool `json:"s3TagsEnabled,omitempty"`

	AuthTokenRestrictions *UpdateAuthTokenRestrictions `json:"authTokenRestrictions,omitempty"`
	// &#128679; Deprecated since v4.13.0  Defines if login fields should be hidden
	HideLoginInputFields bool `json:"hideLoginInputFields,omitempty"`
	// &#128679; Deprecated since v4.12.0  Determines if the media server is enabled
	MediaServerEnabled bool `json:"mediaServerEnabled,omitempty"`
	// &#128679; Deprecated since v4.14.0  Allow weak password  * A weak password has to fulfill the following criteria:     * is at least 8 characters long     * contains letters and numbers  * A strong password has to fulfill the following criteria in addition:     * contains at least one special character     * contains upper and lower case characters  Please use `PUT /system/config/policies/passwords` API to change configured password policies.
	WeakPasswordEnabled bool `json:"weakPasswordEnabled,omitempty"`
}
