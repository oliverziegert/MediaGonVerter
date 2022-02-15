# UpdateLoginPasswordPolicies

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharacterRules** | [***CharacterRules**](CharacterRules.md) |  | [optional] [default to null]
**MinLength** | **int32** | Minimum number of characters a password must contain | [optional] [default to null]
**RejectDictionaryWords** | **bool** | Determines whether a password must NOT contain word(s) from a dictionary | [optional] [default to null]
**RejectUserInfo** | **bool** | Determines whether a password must NOT contain user info (first name, last name, email, user name) | [optional] [default to null]
**RejectKeyboardPatterns** | **bool** | Determines whether a password must NOT contain keyboard patterns (e.g. &#x60;qwertz&#x60;, &#x60;asdf&#x60;)  (min. 4 character pattern) | [optional] [default to null]
**NumberOfArchivedPasswords** | **int32** | Number of passwords to archive  (must be between &#x60;0&#x60; and &#x60;10&#x60;; &#x60;0&#x60; means that password history is disabled) | [optional] [default to null]
**PasswordExpiration** | [***PasswordExpiration**](PasswordExpiration.md) |  | [optional] [default to null]
**UserLockout** | [***UserLockout**](UserLockout.md) |  | [optional] [default to null]
**EnforceLoginPasswordChange** | **bool** | &amp;#128679; Deprecated since v4.24.0  Determines whether a login password change should be enforced for all users  Only takes effect, if login password policies get stricter | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

