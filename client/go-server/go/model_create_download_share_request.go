/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Request model for creating a Download Share
type CreateDownloadShareRequest struct {
	// Source node ID
	NodeId int64 `json:"nodeId"`
	// Alias name  (default: name of the shared node)
	Name string `json:"name,omitempty"`
	// Access password, not allowed for encrypted shares
	Password string `json:"password,omitempty"`

	Expiration *ObjectExpiration `json:"expiration,omitempty"`
	// User notes
	Notes string `json:"notes,omitempty"`
	// &#128640; Since v4.11.0  Internal notes
	InternalNotes string `json:"internalNotes,omitempty"`
	// Show creator first and last name.
	ShowCreatorName bool `json:"showCreatorName,omitempty"`
	// Show creator email address.
	ShowCreatorUsername bool `json:"showCreatorUsername,omitempty"`
	// Max allowed downloads
	MaxDownloads int32 `json:"maxDownloads,omitempty"`

	KeyPair *UserKeyPairContainer `json:"keyPair,omitempty"`

	FileKey *FileKey `json:"fileKey,omitempty"`
	// Language tag for messages to receiver
	ReceiverLanguage string `json:"receiverLanguage,omitempty"`
	// &#128640; Since v4.11.0  List of recipient FQTNs  E.123 / E.164 Format
	TextMessageRecipients []string `json:"textMessageRecipients,omitempty"`
	// &#128679; Deprecated since v4.20.0  Notify creator on every download.
	NotifyCreator bool `json:"notifyCreator,omitempty"`
	// &#128679; Deprecated since v4.20.0  Language tag for messages to creator
	CreatorLanguage string `json:"creatorLanguage,omitempty"`
	// &#128679; Deprecated since v4.11.0  Notify recipients via email  Please use `POST /shares/downloads/{share_id}/email` API instead.
	SendMail bool `json:"sendMail,omitempty"`
	// &#128679; Deprecated since v4.11.0  CSV string of recipient email addresses
	MailRecipients string `json:"mailRecipients,omitempty"`
	// &#128679; Deprecated since v4.11.0  Notification email subject
	MailSubject string `json:"mailSubject,omitempty"`
	// &#128679; Deprecated since v4.11.0  Notification email content
	MailBody string `json:"mailBody,omitempty"`
	// &#128679; Deprecated since v4.11.0  Send share password via SMS  Please use `textMessageRecipients` attribute instead.
	SendSms bool `json:"sendSms,omitempty"`
	// &#128679; Deprecated since v4.11.0  CSV string of recipient MSISDNs
	SmsRecipients string `json:"smsRecipients,omitempty"`
}
