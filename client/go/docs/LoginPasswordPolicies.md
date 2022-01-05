# LoginPasswordPolicies

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharacterRules** | [***CharacterRules**](CharacterRules.md) |  | [default to null]
**MinLength** | **int32** | Minimum number of characters a password must contain | [default to null]
**RejectDictionaryWords** | **bool** | Determines whether a password must NOT contain word(s) from a dictionary | [default to null]
**RejectUserInfo** | **bool** | Determines whether a password must NOT contain user info (first name, last name, email, user name) | [default to null]
**RejectKeyboardPatterns** | **bool** | Determines whether a password must NOT contain keyboard patterns (e.g. &#x60;qwertz&#x60;, &#x60;asdf&#x60;)  (min. 4 character pattern) | [default to null]
**NumberOfArchivedPasswords** | **int32** | Number of passwords to archive  (must be between &#x60;0&#x60; and &#x60;10&#x60;; &#x60;0&#x60; means that password history is disabled) | [default to null]
**PasswordExpiration** | [***PasswordExpiration**](PasswordExpiration.md) |  | [default to null]
**UserLockout** | [***UserLockout**](UserLockout.md) |  | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Modification date | [default to null]
**UpdatedBy** | [***UserInfo**](UserInfo.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

