/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Auth token restrictions
type AuthTokenRestrictions struct {
	// &#128640; Since v4.13.0  Defines if OAuth token restrictions are enabled
	RestrictionEnabled bool `json:"restrictionEnabled,omitempty"`
	// &#128640; Since v4.13.0  Restricted OAuth access token validity (in seconds)
	AccessTokenValidity int32 `json:"accessTokenValidity,omitempty"`
	// &#128640; Since v4.13.0  Restricted OAuth refresh token validity (in seconds)
	RefreshTokenValidity int32 `json:"refreshTokenValidity,omitempty"`
}
