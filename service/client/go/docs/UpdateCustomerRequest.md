# UpdateCustomerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | **string** | Company name | [optional] [default to null]
**CustomerContractType** | **string** | Customer type | [default to null]
**QuotaMax** | **int64** | Maximal disc space which can be allocated by customer in bytes. -1 for unlimited | [optional] [default to null]
**UserMax** | **int32** | Maximal number of users | [optional] [default to null]
**IsLocked** | **bool** | Customer is locked:  * &#x60;false&#x60; - unlocked  * &#x60;true&#x60; - locked    All users of this customer will be blocked and can not login anymore. | [optional] [default to false]
**ProviderCustomerId** | **string** | Provider customer ID | [optional] [default to null]
**WebhooksMax** | **int64** | &amp;#128640; Since v4.19.0  Maximal number of webhooks | [optional] [default to null]
**LockStatus** | **bool** | &amp;#128679; Deprecated since v4.7.0  Customer lock status:  * &#x60;false&#x60; - unlocked  * &#x60;true&#x60; - locked    Please use &#x60;isLocked&#x60; instead.  All users of this customer will be blocked and can not login anymore. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

