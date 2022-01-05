/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Request model for updating an Active Directory configuration
type UpdateActiveDirectoryConfigRequest struct {
	// Unique name for an Active Directory configuration
	Alias string `json:"alias,omitempty"`
	// IPv4 or IPv6 address or host name
	ServerIp string `json:"serverIp,omitempty"`
	// Port
	ServerPort int32 `json:"serverPort,omitempty"`
	// Distinguished Name (DN) of Active Directory administrative account
	ServerAdminName string `json:"serverAdminName,omitempty"`
	// Password of Active Directory administrative account
	ServerAdminPassword string `json:"serverAdminPassword,omitempty"`
	// Search scope of Active Directory; only users below this node can log on.
	LdapUsersDomain string `json:"ldapUsersDomain,omitempty"`
	// Name of Active Directory attribute that is used as login name.
	UserFilter string `json:"userFilter,omitempty"`
	// Determines if a DRACOON account is automatically created for a new user  who successfully logs on with his / her AD / IDP account.
	UserImport bool `json:"userImport,omitempty"`
	// If `userImport` is set to `true`,  the user must be member of this Active Directory group to receive a newly created DRACOON account.
	AdExportGroup string `json:"adExportGroup,omitempty"`
	// User group that is assigned to users who are created by automatic import.  Reset with `0`
	SdsImportGroup int64 `json:"sdsImportGroup,omitempty"`
	// DEPRECATED, will be ignored  Determines whether a room is created for each user that is created by automatic import (like a home folder).  Room's name will equal the user's login name.
	CreateHomeFolder bool `json:"createHomeFolder,omitempty"`
	// DEPRECATED, will be ignored  ID of the room in which the individual rooms for users will be created.
	HomeFolderParent int64 `json:"homeFolderParent,omitempty"`
	// Determines whether LDAPS should be used instead of plain LDAP.
	UseLdaps bool `json:"useLdaps,omitempty"`
	// SSL finger print of Active Directory server.  Mandatory for LDAPS connections.  Format: `Algorithm/Fingerprint`
	SslFingerPrint string `json:"sslFingerPrint,omitempty"`
}
