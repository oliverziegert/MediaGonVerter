# OAuthAuthorization

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | ID of the OAuth client | [default to null]
**ClientName** | **string** | Name, which is shown at the client configuration and authorization. | [default to null]
**UserAgentCategory** | **string** | &amp;#128640; Since v4.12.0  User agent category. | [default to null]
**Id** | **int64** | &amp;#128640; Since v4.12.0  ID of the OAuth authorization | [optional] [default to null]
**IsStandard** | **bool** | &amp;#128640; Since v4.12.0  Determines whether client is a standard client. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | &amp;#128640; Since v4.13.0  Creation date of the authorization | [optional] [default to null]
**UsedAt** | [**time.Time**](time.Time.md) | &amp;#128640; Since v4.13.0  Usage date of the authorization  (Time of last usage.) | [optional] [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) | Expiration date of the authorization | [optional] [default to null]
**UserAgentType** | **string** | &amp;#128640; Since v4.12.0  User agent type. | [optional] [default to null]
**UserAgentOs** | **string** | &amp;#128640; Since v4.12.0  User agent OS. | [optional] [default to null]
**UserAgentInfo** | **string** | &amp;#128640; Since v4.12.0  User agent info. | [optional] [default to null]
**IsCurrentAuthorization** | **bool** | &amp;#128640; Since v4.25.0  Determines whether authorization matches the one from Authorization Header | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

