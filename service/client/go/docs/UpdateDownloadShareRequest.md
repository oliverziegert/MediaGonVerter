# UpdateDownloadShareRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Alias name | [optional] [default to null]
**Password** | **string** | Access password, not allowed for encrypted shares | [optional] [default to null]
**Expiration** | [***ObjectExpiration**](ObjectExpiration.md) |  | [optional] [default to null]
**Notes** | **string** | User notes | [optional] [default to null]
**InternalNotes** | **string** | &amp;#128640; Since v4.11.0  Internal notes | [optional] [default to null]
**ShowCreatorName** | **bool** | Show creator first and last name. | [optional] [default to null]
**ShowCreatorUsername** | **bool** | Show creator email address. | [optional] [default to null]
**MaxDownloads** | **int32** | Max allowed downloads | [optional] [default to null]
**TextMessageRecipients** | **[]string** | List of recipient FQTNs  E.123 / E.164 Format | [optional] [default to null]
**ReceiverLanguage** | **string** | Language tag for messages to receiver | [optional] [default to null]
**DefaultCountry** | **string** | Country shorthand symbol (cf. ISO 3166-2) | [optional] [default to null]
**ResetPassword** | **bool** | Set &#x27;true&#x27; to reset &#x27;password&#x27; for Download Share. | [optional] [default to null]
**ResetMaxDownloads** | **bool** | Set &#x27;true&#x27; to reset &#x27;maxDownloads&#x27; for Download Share. | [optional] [default to null]
**NotifyCreator** | **bool** | &amp;#128679; Deprecated since v4.20.0  Notify creator on every download. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

