/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Request model for creating an OAuth client
type CreateOAuthClientRequest struct {
	// Name, which is shown at the client configuration and authorization.
	ClientName string `json:"clientName"`
	// Authorized grant types  * `authorization_code`  * `implicit`  * `password`  * `client_credentials`  * `refresh_token`    cf. [RFC 6749](https://tools.ietf.org/html/rfc6749)
	GrantTypes []string `json:"grantTypes"`
	// ID of the OAuth client
	ClientId string `json:"clientId,omitempty"`
	// Secret, which client uses at authentication.
	ClientSecret string `json:"clientSecret,omitempty"`
	// Determines whether client is a confidential or public client.
	ClientType string `json:"clientType,omitempty"`
	// URIs, to which a user is redirected after authorization.
	RedirectUris []string `json:"redirectUris,omitempty"`
	// Validity of the access token in seconds.
	AccessTokenValidity int32 `json:"accessTokenValidity,omitempty"`
	// Validity of the refresh token in seconds.
	RefreshTokenValidity int32 `json:"refreshTokenValidity,omitempty"`
	// &#128640; Since v4.22.0  Validity of the approval interval in seconds.
	ApprovalValidity int32 `json:"approvalValidity,omitempty"`
}
