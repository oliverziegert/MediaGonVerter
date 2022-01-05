# SharesPasswordPolicies

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharacterRules** | [***CharacterRules**](CharacterRules.md) |  | [optional] [default to null]
**MinLength** | **int32** | Minimum number of characters a password must contain | [optional] [default to null]
**RejectDictionaryWords** | **bool** | Determines whether a password must NOT contain word(s) from a dictionary | [optional] [default to null]
**RejectUserInfo** | **bool** | Determines whether a password must NOT contain user info (first name, last name, email, user name) | [optional] [default to null]
**RejectKeyboardPatterns** | **bool** | Determines whether a password must NOT contain keyboard patterns (e.g. &#x60;qwertz&#x60;, &#x60;asdf&#x60;)  (min. 4 character pattern) | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Modification date | [optional] [default to null]
**UpdatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

