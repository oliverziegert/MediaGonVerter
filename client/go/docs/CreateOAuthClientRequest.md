# CreateOAuthClientRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | **string** | Name, which is shown at the client configuration and authorization. | [default to null]
**GrantTypes** | **[]string** | Authorized grant types  * &#x60;authorization_code&#x60;  * &#x60;implicit&#x60;  * &#x60;password&#x60;  * &#x60;client_credentials&#x60;  * &#x60;refresh_token&#x60;    cf. [RFC 6749](https://tools.ietf.org/html/rfc6749) | [default to null]
**ClientId** | **string** | ID of the OAuth client | [optional] [default to null]
**ClientSecret** | **string** | Secret, which client uses at authentication. | [optional] [default to null]
**ClientType** | **string** | Determines whether client is a confidential or public client. | [optional] [default to null]
**RedirectUris** | **[]string** | URIs, to which a user is redirected after authorization. | [optional] [default to null]
**AccessTokenValidity** | **int32** | Validity of the access token in seconds. | [optional] [default to null]
**RefreshTokenValidity** | **int32** | Validity of the refresh token in seconds. | [optional] [default to null]
**ApprovalValidity** | **int32** | &amp;#128640; Since v4.22.0  Validity of the approval interval in seconds. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

