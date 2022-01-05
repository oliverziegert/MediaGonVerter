# CreateUploadShareRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | **int64** | Target room or folder ID | [default to null]
**Name** | **string** | Alias name  (default: name of the shared node) | [optional] [default to null]
**Password** | **string** | Password | [optional] [default to null]
**Expiration** | [***ObjectExpiration**](ObjectExpiration.md) |  | [optional] [default to null]
**FilesExpiryPeriod** | **int32** | Number of days after which uploaded files expire | [optional] [default to null]
**Notes** | **string** | User notes | [optional] [default to null]
**InternalNotes** | **string** | &amp;#128640; Since v4.11.0  Internal notes | [optional] [default to null]
**ShowUploadedFiles** | **bool** | Allow display of already uploaded files | [optional] [default to false]
**MaxSlots** | **int32** | Maximal amount of files to upload | [optional] [default to null]
**MaxSize** | **int64** | Maximal total size of uploaded files (in bytes) | [optional] [default to null]
**ReceiverLanguage** | **string** | Language tag for messages to receiver | [optional] [default to null]
**TextMessageRecipients** | **[]string** | &amp;#128640; Since v4.11.0  List of recipient FQTNs  E.123 / E.164 Format | [optional] [default to null]
**ShowCreatorName** | **bool** | &amp;#128640; Since v4.11.0  Show creator first and last name. | [optional] [default to false]
**ShowCreatorUsername** | **bool** | &amp;#128640; Since v4.11.0  Show creator email address. | [optional] [default to false]
**NotifyCreator** | **bool** | &amp;#128679; Deprecated since v4.20.0  Notify creator on every upload. | [optional] [default to false]
**CreatorLanguage** | **string** | &amp;#128679; Deprecated since v4.20.0  Language tag for messages to creator | [optional] [default to null]
**SendMail** | **bool** | &amp;#128679; Deprecated since v4.11.0  Notify recipients via email  Please use &#x60;POST /shares/uploads/{share_id}/email&#x60; API instead. | [optional] [default to false]
**MailRecipients** | **string** | &amp;#128679; Deprecated since v4.11.0  CSV string of recipient email addresses | [optional] [default to null]
**MailSubject** | **string** | &amp;#128679; Deprecated since v4.11.0  Notification email subject | [optional] [default to null]
**MailBody** | **string** | &amp;#128679; Deprecated since v4.11.0  Notification email content | [optional] [default to null]
**SendSms** | **bool** | &amp;#128679; Deprecated since v4.11.0  Send share password via SMS  Please use &#x60;textMessageRecipients&#x60; attribute instead. | [optional] [default to false]
**SmsRecipients** | **string** | &amp;#128679; Deprecated since v4.11.0  CSV string of recipient MSISDNs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

