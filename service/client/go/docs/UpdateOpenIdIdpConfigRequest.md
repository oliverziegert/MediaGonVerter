# UpdateOpenIdIdpConfigRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the IDP | [optional] [default to null]
**Issuer** | **string** | Issuer identifier of the IDP  The value is a case sensitive URL. | [optional] [default to null]
**AuthorizationEndPointUrl** | **string** | URL of the authorization endpoint | [optional] [default to null]
**TokenEndPointUrl** | **string** | URL of the token endpoint | [optional] [default to null]
**UserInfoEndPointUrl** | **string** | URL of the user info endpoint | [optional] [default to null]
**JwksEndPointUrl** | **string** | URL of the JWKS endpoint | [optional] [default to null]
**ClientId** | **string** | ID of the OpenID client | [optional] [default to null]
**ClientSecret** | **string** | Secret, which client uses at authentication. | [optional] [default to null]
**Flow** | **string** | &amp;#128640; Since v4.11.0  Flow, which is used at authentication | [optional] [default to null]
**Scopes** | **[]string** | List of requested scopes  Usually &#x60;openid&#x60; and the names of the requested claims. | [optional] [default to null]
**RedirectUris** | **[]string** | URIs, to which a user is redirected after authorization. | [optional] [default to null]
**PkceEnabled** | **bool** | Determines whether PKCE is enabled.  cf. [RFC 7636](https://tools.ietf.org/html/rfc7636) | [optional] [default to false]
**PkceChallengeMethod** | **string** | PKCE code challenge method.  cf. [RFC 7636](https://tools.ietf.org/html/rfc7636) | [optional] [default to null]
**MappingClaim** | **string** | Name of the claim which is used for the user mapping. | [optional] [default to null]
**FallbackMappingClaim** | **string** | Name of the claim which is used for the user mapping fallback. | [optional] [default to null]
**ResetFallbackMappingClaim** | **bool** | Set &#x60;true&#x60; to reset &#x60;fallbackMappingClaim&#x60;. | [optional] [default to null]
**UserInfoSource** | **string** | &amp;#128640; Since v4.23.0  Source, which is used to get user information at the import or update of a user. | [optional] [default to null]
**UserImportEnabled** | **bool** | Determines if a DRACOON account is automatically created for a new user  who successfully logs on with his / her AD / IDP account. | [optional] [default to false]
**UserImportGroup** | **int64** | User group that is assigned to users who are created by automatic import.  Reset with &#x60;0&#x60; | [optional] [default to null]
**UserUpdateEnabled** | **bool** | Determines if the DRACOON account is updated with data from AD / IDP.  For OpenID Connect, the scopes &#x60;email&#x60; and &#x60;profile&#x60; are needed. | [optional] [default to false]
**UserManagementUrl** | **string** | URL of the user management UI.  Use empty string to remove. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

