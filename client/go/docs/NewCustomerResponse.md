# NewCustomerResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the customer | [optional] [default to null]
**CompanyName** | **string** | Company name | [default to null]
**CustomerContractType** | **string** | Customer type | [default to null]
**QuotaMax** | **int64** | Maximal disc space which can be allocated by customer in bytes. -1 for unlimited | [default to null]
**UserMax** | **int32** | Maximal number of users | [default to null]
**IsLocked** | **bool** | Customer is locked:  * &#x60;false&#x60; - unlocked  * &#x60;true&#x60; - locked    All users of this customer will be blocked and can not login anymore. | [optional] [default to false]
**TrialDays** | **int32** | Number of days left for trial period (relevant only for type &#x60;demo&#x60;)  (not used) | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date | [optional] [default to null]
**FirstAdminUser** | [***FirstAdminUser**](FirstAdminUser.md) |  | [default to null]
**CustomerAttributes** | [***CustomerAttributes**](CustomerAttributes.md) |  | [optional] [default to null]
**ProviderCustomerId** | **string** | Provider customer ID | [optional] [default to null]
**WebhooksMax** | **int64** | &amp;#128640; Since v4.19.0  Maximal number of webhooks | [optional] [default to null]
**CustomerUuid** | **string** | &amp;#128640; Since v4.21.0  Customer UUID | [default to null]
**ActivationCode** | **string** | &amp;#128679; Deprecated since v4.8.0  Customer activation code string:  * valid only for types &#x60;free&#x60; and &#x60;demo&#x60;  * for &#x60;pay&#x60; customers it is empty | [optional] [default to null]
**LockStatus** | **bool** | &amp;#128679; Deprecated since v4.7.0  Customer lock status:  * &#x60;false&#x60; - unlocked  * &#x60;true&#x60; - locked    Please use &#x60;isLocked&#x60; instead.  All users of this customer will be blocked and can not login anymore. | [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

