# MoveNodesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]MoveNode**](MoveNode.md) | List of nodes to be moved | [optional] [default to null]
**ResolutionStrategy** | **string** | Node conflict resolution strategy:  * &#x60;autorename&#x60;  * &#x60;overwrite&#x60;  * &#x60;fail&#x60; | [optional] [default to RESOLUTION_STRATEGY.AUTORENAME]
**KeepShareLinks** | **bool** | Preserve Download Share Links and point them to the new node. | [optional] [default to false]
**NodeIds** | **[]int64** | &amp;#128679; Deprecated since v4.5.0  Node IDs  Please use &#x60;items&#x60; instead. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

