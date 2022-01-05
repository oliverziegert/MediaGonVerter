# SystemInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LanguageDefault** | **string** | System default language  cf. [RFC 5646](https://tools.ietf.org/html/rfc5646) | [default to null]
**HideLoginInputFields** | **bool** | &amp;#128640; Since v4.13.0  Defines if login fields should be hidden | [default to null]
**S3Hosts** | **[]string** | &amp;#128640; Since v4.14.0  List of S3 Hosts for CSP header | [default to null]
**S3EnforceDirectUpload** | **bool** | &amp;#128640; Since v4.15.0  Determines whether S3 direct upload is enforced or not | [default to null]
**UseS3Storage** | **bool** | &amp;#128640; Since v4.21.0  Defines if S3 is used as storage backend | [default to null]
**AuthMethods** | [**[]AuthMethod**](AuthMethod.md) | &amp;#128679; Deprecated since v4.13.0  Authentication methods:  * &#x60;sql&#x60;  * &#x60;active_directory&#x60;  * &#x60;radius&#x60;  * &#x60;openid&#x60;  use &#x60;authData&#x60; instead | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

