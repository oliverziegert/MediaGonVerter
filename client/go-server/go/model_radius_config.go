/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// RADIUS configuration
type RadiusConfig struct {
	// RADIUS Server IP Address
	IpAddress string `json:"ipAddress"`
	// RADIUS Server Port
	Port int32 `json:"port"`
	// Shared Secret to access the RADIUS server
	SharedSecret string `json:"sharedSecret"`
	// Sequence order of concatenated PIN and one-time token
	OtpPinFirst bool `json:"otpPinFirst"`

	FailoverServer *FailoverServer `json:"failoverServer,omitempty"`
}
