/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Password reset information
type ResetPasswordTokenValidateResponse struct {
	// User first name
	FirstName string `json:"firstName"`
	// User last name
	LastName              string                 `json:"lastName"`
	LoginPasswordPolicies *LoginPasswordPolicies `json:"loginPasswordPolicies,omitempty"`
	// &#128679; Deprecated since v4.18.0  Job title
	Title string `json:"title,omitempty"`
	// &#128679; Deprecated since v4.12.0  Gender
	Gender string `json:"gender,omitempty"`
	// &#128679; Deprecated since v4.14.0  Allow weak password  Please use `loginPasswordPolicies` instead
	AllowSystemGlobalWeakPassword bool `json:"allowSystemGlobalWeakPassword,omitempty"`
}
