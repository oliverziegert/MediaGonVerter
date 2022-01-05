/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Request model for testing connection for Active Directory configuration
type TestActiveDirectoryConfigRequest struct {
	// IPv4 or IPv6 address or host name
	ServerIp string `json:"serverIp"`
	// Port
	ServerPort int32 `json:"serverPort"`
	// Distinguished Name (DN) of Active Directory administrative account
	ServerAdminName string `json:"serverAdminName"`
	// Password of Active Directory administrative account
	ServerAdminPassword string `json:"serverAdminPassword"`
	// Search scope of Active Directory; only users below this node can log on.
	LdapUsersDomain string `json:"ldapUsersDomain"`
	// Determines whether LDAPS should be used instead of plain LDAP.
	UseLdaps bool `json:"useLdaps,omitempty"`
	// SSL finger print of Active Directory server.  Mandatory for LDAPS connections.  Format: `Algorithm/Fingerprint`
	SslFingerPrint string `json:"sslFingerPrint,omitempty"`
}
