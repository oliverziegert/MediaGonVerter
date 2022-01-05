# UpdateActiveDirectoryConfigRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Unique name for an Active Directory configuration | [optional] [default to null]
**ServerIp** | **string** | IPv4 or IPv6 address or host name | [optional] [default to null]
**ServerPort** | **int32** | Port | [optional] [default to null]
**ServerAdminName** | **string** | Distinguished Name (DN) of Active Directory administrative account | [optional] [default to null]
**ServerAdminPassword** | **string** | Password of Active Directory administrative account | [optional] [default to null]
**LdapUsersDomain** | **string** | Search scope of Active Directory; only users below this node can log on. | [optional] [default to null]
**UserFilter** | **string** | Name of Active Directory attribute that is used as login name. | [optional] [default to null]
**UserImport** | **bool** | Determines if a DRACOON account is automatically created for a new user  who successfully logs on with his / her AD / IDP account. | [optional] [default to null]
**AdExportGroup** | **string** | If &#x60;userImport&#x60; is set to &#x60;true&#x60;,  the user must be member of this Active Directory group to receive a newly created DRACOON account. | [optional] [default to null]
**SdsImportGroup** | **int64** | User group that is assigned to users who are created by automatic import.  Reset with &#x60;0&#x60; | [optional] [default to null]
**CreateHomeFolder** | **bool** | DEPRECATED, will be ignored  Determines whether a room is created for each user that is created by automatic import (like a home folder).  Room&#x27;s name will equal the user&#x27;s login name. | [optional] [default to false]
**HomeFolderParent** | **int64** | DEPRECATED, will be ignored  ID of the room in which the individual rooms for users will be created. | [optional] [default to null]
**UseLdaps** | **bool** | Determines whether LDAPS should be used instead of plain LDAP. | [optional] [default to null]
**SslFingerPrint** | **string** | SSL finger print of Active Directory server.  Mandatory for LDAPS connections.  Format: &#x60;Algorithm/Fingerprint&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

