# UserAuthData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | Authentication method    Authentication methods:  * &#x60;basic&#x60;  * &#x60;active_directory&#x60;  * &#x60;radius&#x60;  * &#x60;openid&#x60; | [default to null]
**Login** | **string** | User login name | [optional] [default to null]
**Password** | **string** | Password (only relevant for &#x60;basic&#x60; authentication type)  *NOT* your Active Directory, OpenID or RADIUS password! | [optional] [default to null]
**MustChangePassword** | **bool** | Determines whether user has to change his / her password  * default: &#x60;true&#x60; for &#x60;basic&#x60; auth type  * default: &#x60;false&#x60; for &#x60;active_directory&#x60;, &#x60;openid&#x60; and &#x60;radius&#x60; auth types | [optional] [default to null]
**AdConfigId** | **int32** | ID of the user&#x27;s Active Directory. | [optional] [default to null]
**OidConfigId** | **int32** | ID of the user&#x27;s OIDC provider. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

