# RestoreDeletedNodesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedNodeIds** | **[]int64** | List of deleted node IDs | [default to null]
**ResolutionStrategy** | **string** | Node conflict resolution strategy:  * &#x60;autorename&#x60;  * &#x60;overwrite&#x60;  * &#x60;fail&#x60; | [optional] [default to RESOLUTION_STRATEGY.AUTORENAME]
**KeepShareLinks** | **bool** | Preserve Download Share Links and point them to the new node. | [optional] [default to false]
**ParentId** | **int64** | Node parent ID  (default: previous parent ID) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

