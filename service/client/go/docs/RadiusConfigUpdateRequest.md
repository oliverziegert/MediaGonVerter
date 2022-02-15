# RadiusConfigUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | RADIUS Server IP Address | [optional] [default to null]
**Port** | **int32** | RADIUS Server Port | [optional] [default to null]
**SharedSecret** | **string** | Shared Secret to access the RADIUS server | [optional] [default to null]
**OtpPinFirst** | **bool** | Sequence order of concatenated PIN and one-time token | [optional] [default to true]
**FailoverServer** | [***FailoverServer**](FailoverServer.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

