/*
 * DRACOON API
 *
 * REST Web Services for DRACOON<br><br>This page provides an overview of all available and documented DRACOON APIs, which are grouped by tags.<br>Each tag provides a collection of APIs that are intended for a specific area of the DRACOON.<br><br><a title='Developer Information' href='https://developer.dracoon.com'>Developer Information</a>&emsp;&emsp;<a title='Get SDKs on GitHub' href='https://github.com/dracoon'>Get SDKs on GitHub</a><br><br><a title='Terms of service' href='https://www.dracoon.com/terms/general-terms-and-conditions/'>Terms of service</a>
 *
 * API version: 4.33.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Request model for updating a customer
type UpdateCustomerRequest struct {
	// Company name
	CompanyName string `json:"companyName,omitempty"`
	// Customer type
	CustomerContractType string `json:"customerContractType"`
	// Maximal disc space which can be allocated by customer in bytes. -1 for unlimited
	QuotaMax int64 `json:"quotaMax,omitempty"`
	// Maximal number of users
	UserMax int32 `json:"userMax,omitempty"`
	// Customer is locked:  * `false` - unlocked  * `true` - locked    All users of this customer will be blocked and can not login anymore.
	IsLocked bool `json:"isLocked,omitempty"`
	// Provider customer ID
	ProviderCustomerId string `json:"providerCustomerId,omitempty"`
	// &#128640; Since v4.19.0  Maximal number of webhooks
	WebhooksMax int64 `json:"webhooksMax,omitempty"`
	// &#128679; Deprecated since v4.7.0  Customer lock status:  * `false` - unlocked  * `true` - locked    Please use `isLocked` instead.  All users of this customer will be blocked and can not login anymore.
	LockStatus bool `json:"lockStatus,omitempty"`
}
