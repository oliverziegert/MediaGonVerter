# TestActiveDirectoryConfigResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerIp** | **string** | IPv4 or IPv6 address or host name | [default to null]
**ServerPort** | **int32** | Port | [default to null]
**ServerAdminName** | **string** | Distinguished Name (DN) of Active Directory administrative account | [default to null]
**ServerAdminPassword** | **string** | Password of Active Directory administrative account | [default to null]
**LdapUsersDomain** | **string** | Search scope of Active Directory; only users below this node can log on. | [default to null]
**UseLdaps** | **bool** | Determines whether LDAPS should be used instead of plain LDAP. | [default to null]
**SslFingerPrint** | **string** | SSL finger print of Active Directory server.  Mandatory for LDAPS connections.  Format: &#x60;Algorithm/Fingerprint&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

