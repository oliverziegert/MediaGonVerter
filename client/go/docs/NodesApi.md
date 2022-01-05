# {{classname}}

All URIs are relative to *https://cloud.pc-ziegert.de/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFavorite**](NodesApi.md#AddFavorite) | **Post** /v4/nodes/{node_id}/favorite | Mark a node (room, folder or file) as favorite
[**CancelFileUpload**](NodesApi.md#CancelFileUpload) | **Delete** /v4/nodes/files/uploads/{upload_id} | Cancel file upload
[**ChangePendingAssignments**](NodesApi.md#ChangePendingAssignments) | **Put** /v4/nodes/rooms/pending | Handle user-room assignments per group
[**CompleteFileUpload**](NodesApi.md#CompleteFileUpload) | **Put** /v4/nodes/files/uploads/{upload_id} | Complete file upload
[**CompleteS3FileUpload**](NodesApi.md#CompleteS3FileUpload) | **Put** /v4/nodes/files/uploads/{upload_id}/s3 | Complete S3 file upload
[**ConfigureRoom**](NodesApi.md#ConfigureRoom) | **Put** /v4/nodes/rooms/{room_id}/config | Configure room
[**CopyNodes**](NodesApi.md#CopyNodes) | **Post** /v4/nodes/{node_id}/copy_to | Copy node(s)
[**CreateAndPreserveRoomRescueKeyPair**](NodesApi.md#CreateAndPreserveRoomRescueKeyPair) | **Post** /v4/nodes/rooms/{room_id}/keypairs | Create key pair and preserve copy of old private key
[**CreateFileUploadChannel**](NodesApi.md#CreateFileUploadChannel) | **Post** /v4/nodes/files/uploads | Create new file upload channel
[**CreateFolder**](NodesApi.md#CreateFolder) | **Post** /v4/nodes/folders | Create new folder
[**CreateNodeComment**](NodesApi.md#CreateNodeComment) | **Post** /v4/nodes/{node_id}/comments | Create node comment
[**CreateRoom**](NodesApi.md#CreateRoom) | **Post** /v4/nodes/rooms | Create new room
[**DownloadZipArchive**](NodesApi.md#DownloadZipArchive) | **Post** /v4/nodes/zip/download | Download files / folders as ZIP archive
[**EmptyDeletedNodes**](NodesApi.md#EmptyDeletedNodes) | **Delete** /v4/nodes/{node_id}/deleted_nodes | Empty recycle bin
[**EncryptRoom**](NodesApi.md#EncryptRoom) | **Put** /v4/nodes/rooms/{room_id}/encrypt | Encrypt room
[**GenerateDownloadUrl**](NodesApi.md#GenerateDownloadUrl) | **Post** /v4/nodes/files/{file_id}/downloads | Generate download URL
[**GenerateDownloadUrlForZipArchive**](NodesApi.md#GenerateDownloadUrlForZipArchive) | **Post** /v4/nodes/zip | Generate download URL for ZIP download
[**GeneratePresignedUrlsFiles**](NodesApi.md#GeneratePresignedUrlsFiles) | **Post** /v4/nodes/files/uploads/{upload_id}/s3_urls | Generate presigned URLs for S3 file upload
[**HandleRoomWebhookAssignments**](NodesApi.md#HandleRoomWebhookAssignments) | **Put** /v4/nodes/rooms/{room_id}/webhooks | Assign or unassign webhooks to room
[**MoveNodes**](NodesApi.md#MoveNodes) | **Post** /v4/nodes/{node_id}/move_to | Move node(s)
[**RemoveDeletedNodes**](NodesApi.md#RemoveDeletedNodes) | **Delete** /v4/nodes/deleted_nodes | Remove nodes from recycle bin
[**RemoveFavorite**](NodesApi.md#RemoveFavorite) | **Delete** /v4/nodes/{node_id}/favorite | Unmark a node (room, folder or file) as favorite
[**RemoveNode**](NodesApi.md#RemoveNode) | **Delete** /v4/nodes/{node_id} | Remove node
[**RemoveNodeComment**](NodesApi.md#RemoveNodeComment) | **Delete** /v4/nodes/comments/{comment_id} | Remove node comment
[**RemoveNodes**](NodesApi.md#RemoveNodes) | **Delete** /v4/nodes | Remove nodes
[**RemoveRoomRescueKeyPair**](NodesApi.md#RemoveRoomRescueKeyPair) | **Delete** /v4/nodes/rooms/{room_id}/keypair | Remove rooms&#x27;s rescue key pair
[**RequestDeletedNode**](NodesApi.md#RequestDeletedNode) | **Get** /v4/nodes/deleted_nodes/{deleted_node_id} | Request deleted node
[**RequestDeletedNodeVersions**](NodesApi.md#RequestDeletedNodeVersions) | **Get** /v4/nodes/{node_id}/deleted_nodes/versions | Request deleted versions of nodes
[**RequestDeletedNodesSummary**](NodesApi.md#RequestDeletedNodesSummary) | **Get** /v4/nodes/{node_id}/deleted_nodes | Request list of deleted nodes
[**RequestListOfWebhooksForRoom**](NodesApi.md#RequestListOfWebhooksForRoom) | **Get** /v4/nodes/rooms/{room_id}/webhooks | Request list of webhooks that are assigned or can be assigned to this room
[**RequestMissingFileKeys**](NodesApi.md#RequestMissingFileKeys) | **Get** /v4/nodes/missingFileKeys | Request files without user&#x27;s file key
[**RequestNode**](NodesApi.md#RequestNode) | **Get** /v4/nodes/{node_id} | Request node
[**RequestNodeComments**](NodesApi.md#RequestNodeComments) | **Get** /v4/nodes/{node_id}/comments | Request list of node comments
[**RequestNodeParents**](NodesApi.md#RequestNodeParents) | **Get** /v4/nodes/{node_id}/parents | Request list of parent nodes
[**RequestNodes**](NodesApi.md#RequestNodes) | **Get** /v4/nodes | Request list of nodes
[**RequestPendingAssignments**](NodesApi.md#RequestPendingAssignments) | **Get** /v4/nodes/rooms/pending | Request user-room assignments per group
[**RequestRoomActivitiesLogAsJson**](NodesApi.md#RequestRoomActivitiesLogAsJson) | **Get** /v4/nodes/rooms/{room_id}/events | Request events of a room
[**RequestRoomGroups**](NodesApi.md#RequestRoomGroups) | **Get** /v4/nodes/rooms/{room_id}/groups | Request room granted group(s) or / and group(s) that can be granted
[**RequestRoomPolicies**](NodesApi.md#RequestRoomPolicies) | **Get** /v4/nodes/rooms/{room_id}/policies | Request Room Policies
[**RequestRoomRescueKey**](NodesApi.md#RequestRoomRescueKey) | **Get** /v4/nodes/files/{file_id}/data_room_file_key | Request room rescue key
[**RequestRoomRescueKeyPair**](NodesApi.md#RequestRoomRescueKeyPair) | **Get** /v4/nodes/rooms/{room_id}/keypair | Request room rescue key
[**RequestRoomRescueKeyPairs**](NodesApi.md#RequestRoomRescueKeyPairs) | **Get** /v4/nodes/rooms/{room_id}/keypairs | Request all room rescue key pairs
[**RequestRoomS3Tags**](NodesApi.md#RequestRoomS3Tags) | **Get** /v4/nodes/rooms/{room_id}/s3_tags | Request list of all assigned S3 tags to the room
[**RequestRoomUsers**](NodesApi.md#RequestRoomUsers) | **Get** /v4/nodes/rooms/{room_id}/users | Request room granted user(s) or / and user(s) that can be granted
[**RequestSystemRescueKey**](NodesApi.md#RequestSystemRescueKey) | **Get** /v4/nodes/files/{file_id}/data_space_file_key | Request system rescue key
[**RequestUploadStatusFiles**](NodesApi.md#RequestUploadStatusFiles) | **Get** /v4/nodes/files/uploads/{upload_id} | Request status of S3 file upload
[**RequestUserFileKey**](NodesApi.md#RequestUserFileKey) | **Get** /v4/nodes/files/{file_id}/user_file_key | Request user&#x27;s file key
[**RestoreNodes**](NodesApi.md#RestoreNodes) | **Post** /v4/nodes/deleted_nodes/actions/restore | Restore deleted nodes
[**RevokeRoomGroups**](NodesApi.md#RevokeRoomGroups) | **Delete** /v4/nodes/rooms/{room_id}/groups | Revoke granted group(s) from room
[**RevokeRoomUsers**](NodesApi.md#RevokeRoomUsers) | **Delete** /v4/nodes/rooms/{room_id}/users | Revoke granted user(s) from room
[**SearchNodes**](NodesApi.md#SearchNodes) | **Get** /v4/nodes/search | Search nodes
[**SetRoomPolicies**](NodesApi.md#SetRoomPolicies) | **Put** /v4/nodes/rooms/{room_id}/policies | Set room policies
[**SetRoomRescueKeyPair**](NodesApi.md#SetRoomRescueKeyPair) | **Post** /v4/nodes/rooms/{room_id}/keypair | Set room&#x27;s rescue key pair
[**SetRoomS3Tags**](NodesApi.md#SetRoomS3Tags) | **Post** /v4/nodes/rooms/{room_id}/s3_tags | Set S3 tags for a room
[**SetUserFileKeys**](NodesApi.md#SetUserFileKeys) | **Post** /v4/nodes/files/keys | Set file keys for a list of users and files
[**UpdateFavorites**](NodesApi.md#UpdateFavorites) | **Put** /v4/nodes/favorites | Mark or unmark a list of nodes (room, folder or file) as favorite
[**UpdateFile**](NodesApi.md#UpdateFile) | **Put** /v4/nodes/files/{file_id} | Updates a file’s metadata
[**UpdateFiles**](NodesApi.md#UpdateFiles) | **Put** /v4/nodes/files | Updates a list of  file’s metadata
[**UpdateFolder**](NodesApi.md#UpdateFolder) | **Put** /v4/nodes/folders/{folder_id} | Updates folder’s metadata
[**UpdateNodeComment**](NodesApi.md#UpdateNodeComment) | **Put** /v4/nodes/comments/{comment_id} | Edit node comment
[**UpdateRoom**](NodesApi.md#UpdateRoom) | **Put** /v4/nodes/rooms/{room_id} | Updates room’s metadata
[**UpdateRoomGroups**](NodesApi.md#UpdateRoomGroups) | **Put** /v4/nodes/rooms/{room_id}/groups | Add or change room granted group(s)
[**UpdateRoomUsers**](NodesApi.md#UpdateRoomUsers) | **Put** /v4/nodes/rooms/{room_id}/users | Add or change room granted user(s)
[**UploadFileAsMultipart**](NodesApi.md#UploadFileAsMultipart) | **Post** /v4/nodes/files/uploads/{upload_id} | Upload file

# **AddFavorite**
> Node AddFavorite(ctx, nodeId, optional)
Mark a node (room, folder or file) as favorite

### Description:   Marks a node (room, folder or file) as favorite.  ### Precondition: Authenticated user is allowed to <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128065; see</span> the node (i.e. `isBrowsable = true`).  ### Postcondition: A node gets marked as favorite.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Node ID | 
 **optional** | ***NodesApiAddFavoriteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiAddFavoriteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelFileUpload**
> CancelFileUpload(ctx, uploadId, optional)
Cancel file upload

### Description: Cancel a (S3) file upload and destroy the upload channel.  ### Precondition: An upload channel has been created and user has to be the creator of the upload channel.  ### Postcondition: The upload channel is removed and all temporary uploaded data is purged.  ### Further Information: It is recommended to notify the API about cancelled uploads if possible.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uploadId** | **string**| Upload channel ID | 
 **optional** | ***NodesApiCancelFileUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCancelFileUploadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangePendingAssignments**
> ChangePendingAssignments(ctx, body, optional)
Handle user-room assignments per group

### Description:   Handles a list of user-room assignments by groups that have **NOT** been approved yet   **WAITING** or **DENIED** assignments can be **ACCEPTED**.  ### Precondition: None.  ### Postcondition: User-room assignment is approved and the user gets access to the group.  ### Further Information: Room administrators should **SHOULD** handle pending assignments to provide access to rooms for other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PendingAssignmentsRequest**](PendingAssignmentsRequest.md)|  | 
 **optional** | ***NodesApiChangePendingAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiChangePendingAssignmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompleteFileUpload**
> Node CompleteFileUpload(ctx, uploadId, optional)
Complete file upload

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.9.0</h3>  ### Use `uploads` API  ### Description: Finishes an upload and closes the corresponding upload channel.  ### Precondition: An upload channel has been created and data has been transmitted.  ### Postcondition: The upload is finished and the temporary file is moved to the productive environment.  ### Further Information: The provided file name might be changed in accordance with the resolution strategy:   * **autorename**: changes the file name and adds a number to avoid conflicts. * **overwrite**: deletes any old file with the same file name. * **fail**: returns an error; in this case, another `PUT` request with a different file name may be sent.  Please ensure that all chunks have been transferred correctly before finishing the upload.   Download share id (if exists) gets changed if: - node with the same name exists in the target container - `resolutionStrategy` is `overwrite` - `keepShareLinks` is `true`  ### Node naming convention: * Node (room, folder, file) names are limited to **150** characters. * Illegal names:   `'CON', 'PRN', 'AUX', 'NUL', 'COM1', 'COM2', 'COM3', 'COM4', 'COM5', 'COM6', 'COM7', 'COM8', 'COM9', 'LPT1', 'LPT2', 'LPT3', 'LPT4', 'LPT5', 'LPT6', 'LPT7', 'LPT8', 'LPT9', (and any of those with an extension)` * Illegal characters in names:   `'\\\\', '<','>', ':', '\\\"', '|', '?', '*', '/', leading '-', trailing '.' `

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uploadId** | **string**| Upload channel ID | 
 **optional** | ***NodesApiCompleteFileUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCompleteFileUploadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CompleteUploadRequest**](CompleteUploadRequest.md)|  | 
 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompleteS3FileUpload**
> CompleteS3FileUpload(ctx, body, uploadId, optional)
Complete S3 file upload

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.15.0</h3>  ### Description: Finishes a S3 file upload and closes the corresponding upload channel.  ### Precondition: An upload channel has been created, data has been transmitted and user has to be the creator of the upload channel  ### Postcondition: Upload channel is closed. S3 multipart upload request is completed.  ### Further Information: Download share id (if exists) gets changed if: - node with the same name exists in the target container - `resolutionStrategy` is `overwrite` - `keepShareLinks` is `true`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CompleteS3FileUploadRequest**](CompleteS3FileUploadRequest.md)|  | 
  **uploadId** | **string**| Upload channel ID | 
 **optional** | ***NodesApiCompleteS3FileUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCompleteS3FileUploadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigureRoom**
> Node ConfigureRoom(ctx, body, roomId, optional)
Configure room

### Description: Configure a room.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: Room's configuration is changed.  ### Further Information: Provided (or default) classification is taken from room when file gets uploaded without any classification.    To set `adminIds` or `adminGroupIds` the `inheritPermissions` value has to be `false`. Otherwise use: * `PUT /nodes/rooms/{room_id}/groups` * `PUT /nodes/rooms/{room_id}/users `    APIs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigRoomRequest**](ConfigRoomRequest.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiConfigureRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiConfigureRoomOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyNodes**
> Node CopyNodes(ctx, body, nodeId, optional)
Copy node(s)

### Description: Copies nodes (folder, file) to another parent.  ### Precondition: Authenticated user with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions in the source parent and <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; create</span> permissions in the target parent node.  ### Postcondition: Nodes are copied to target parent.  ### Further Information: Nodes **MUST** be in same source parent.   **Rooms **CANNOT** be copied.**  Download share id (if exists) gets changed if: - node with the same name exists in the target container - `resolutionStrategy` is `overwrite` - `keepShareLinks` is `true`  ### Node naming convention: * Node (room, folder, file) names are limited to **150** characters. * Illegal names:   `'CON', 'PRN', 'AUX', 'NUL', 'COM1', 'COM2', 'COM3', 'COM4', 'COM5', 'COM6', 'COM7', 'COM8', 'COM9', 'LPT1', 'LPT2', 'LPT3', 'LPT4', 'LPT5', 'LPT6', 'LPT7', 'LPT8', 'LPT9', (and any of those with an extension)` * Illegal characters in names:   `'\\\\', '<','>', ':', '\\\"', '|', '?', '*', '/', leading '-', trailing '.' ` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CopyNodesRequest**](CopyNodesRequest.md)|  | 
  **nodeId** | **int64**| Target parent node ID | 
 **optional** | ***NodesApiCopyNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCopyNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAndPreserveRoomRescueKeyPair**
> CreateAndPreserveRoomRescueKeyPair(ctx, body, roomId, optional)
Create key pair and preserve copy of old private key

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Create room rescue key pair and preserve copy of old private key.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: Room rescue key pair is created.   Copy of old private key is preserved.  ### Further Information: You can submit your old private key, encrypted with your current password.   This allows migrating file keys encrypted with your old key pair to the new one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateKeyPairRequest**](CreateKeyPairRequest.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiCreateAndPreserveRoomRescueKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCreateAndPreserveRoomRescueKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFileUploadChannel**
> CreateFileUploadResponse CreateFileUploadChannel(ctx, body, optional)
Create new file upload channel

### Description: This endpoint creates a new upload channel which is the first step in any file upload workflow.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; create</span> permissions in the parent container (room or folder).  ### Postcondition: A new upload channel for a file is created.   Its ID and an upload token are returned.  ### Further Information: The upload ID is used for uploads with `X-Sds-Auth-Token` header, the upload token can be used for uploads without authentication header.  Please provide the size of the intended upload so that the quota can be checked in advanced and no data is transferred unnecessarily.  Notes are limited to **255** characters.  ### Node naming convention: * Node (room, folder, file) names are limited to **150** characters. * Illegal names:   `'CON', 'PRN', 'AUX', 'NUL', 'COM1', 'COM2', 'COM3', 'COM4', 'COM5', 'COM6', 'COM7', 'COM8', 'COM9', 'LPT1', 'LPT2', 'LPT3', 'LPT4', 'LPT5', 'LPT6', 'LPT7', 'LPT8', 'LPT9', (and any of those with an extension)` * Illegal characters in names:   `'\\\\', '<','>', ':', '\\\"', '|', '?', '*', '/', leading '-', trailing '.' ` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFileUploadRequest**](CreateFileUploadRequest.md)|  | 
 **optional** | ***NodesApiCreateFileUploadChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCreateFileUploadChannelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**CreateFileUploadResponse**](CreateFileUploadResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFolder**
> Node CreateFolder(ctx, body, optional)
Create new folder

### Description: Create a new folder.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; create</span> permissions in current room.  ### Postcondition: New folder is created.  ### Further Information: Folders **CANNOT** be created on top level (without parent element).   Notes are limited to **255** characters.  ### Node naming convention: * Node (room, folder, file) names are limited to **150** characters. * Illegal names:   `'CON', 'PRN', 'AUX', 'NUL', 'COM1', 'COM2', 'COM3', 'COM4', 'COM5', 'COM6', 'COM7', 'COM8', 'COM9', 'LPT1', 'LPT2', 'LPT3', 'LPT4', 'LPT5', 'LPT6', 'LPT7', 'LPT8', 'LPT9', (and any of those with an extension)` * Illegal characters in names:   `'\\\\', '<','>', ':', '\\\"', '|', '?', '*', '/', leading '-', trailing '.' ` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFolderRequest**](CreateFolderRequest.md)|  | 
 **optional** | ***NodesApiCreateFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCreateFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNodeComment**
> Comment CreateNodeComment(ctx, body, nodeId, optional)
Create node comment

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.10.0</h3>  ### Description: Create a comment for a specific node.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions on the node.  ### Postcondition: Comment is created.  ### Further Information: Maximum allowed text length: **65535** characters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateNodeCommentRequest**](CreateNodeCommentRequest.md)|  | 
  **nodeId** | **int64**| Node ID | 
 **optional** | ***NodesApiCreateNodeCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCreateNodeCommentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Comment**](Comment.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRoom**
> Node CreateRoom(ctx, body, optional)
Create new room

### Description: Creates a new room at the provided parent node.   Creation of top level rooms provided.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage</span> permissions in the parent room.  ### Postcondition: A new room is created.  ### Further Information:   Rooms may only have other rooms as parent.   Rooms on top level do **NOT** have any parent.   Rooms may have rooms as children on n hierarchy levels.   If permission inheritance is disabled, there **MUST** be at least one admin user / group (with neither the group nor the user having an expiration date).  Notes are limited to **255** characters.  Provided (or default) classification is taken from room when file gets uploaded without any classification.  ### Node naming convention: * Node (room, folder, file) names are limited to **150** characters. * Illegal names:   `'CON', 'PRN', 'AUX', 'NUL', 'COM1', 'COM2', 'COM3', 'COM4', 'COM5', 'COM6', 'COM7', 'COM8', 'COM9', 'LPT1', 'LPT2', 'LPT3', 'LPT4', 'LPT5', 'LPT6', 'LPT7', 'LPT8', 'LPT9', (and any of those with an extension)` * Illegal characters in names:   `'\\\\', '<','>', ':', '\\\"', '|', '?', '*', '/', leading '-', trailing '.' `

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoomRequest**](CreateRoomRequest.md)|  | 
 **optional** | ***NodesApiCreateRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiCreateRoomOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadZipArchive**
> DownloadZipArchive(ctx, body, optional)
Download files / folders as ZIP archive

### Description:   Download multiple files in a ZIP archive.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions in auth parent room.  ### Postcondition: Stream is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ZipDownloadRequest**](ZipDownloadRequest.md)|  | 
 **optional** | ***NodesApiDownloadZipArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiDownloadZipArchiveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmptyDeletedNodes**
> EmptyDeletedNodes(ctx, nodeId, optional)
Empty recycle bin

### Description:   Empty a recycle bin.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; delete recycle bin</span> permissions in parent room.  ### Postcondition: All files in the recycle bin are permanently removed.  ### Further Information: Actually removes the previously deleted files from the system.   **This action is irreversible.**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Room ID | 
 **optional** | ***NodesApiEmptyDeletedNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiEmptyDeletedNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EncryptRoom**
> Node EncryptRoom(ctx, body, roomId, optional)
Encrypt room

### Description:   Activates the client-side encryption for a room.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: Encryption of room is activated.  ### Further Information: Only empty rooms at the top level may be encrypted.   This endpoint may also be used to disable encryption of an empty room.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EncryptRoomRequest**](EncryptRoomRequest.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiEncryptRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiEncryptRoomOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateDownloadUrl**
> DownloadTokenGenerateResponse GenerateDownloadUrl(ctx, fileId, optional)
Generate download URL

### Description: Create a download URL to retrieve a file without `X-Sds-Auth-Token` Header.  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions in parent room.  ### Postcondition: Download token is generated and returned.  ### Further Information: The token is necessary to access `downloads` ressources.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **int64**| File ID | 
 **optional** | ***NodesApiGenerateDownloadUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiGenerateDownloadUrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**DownloadTokenGenerateResponse**](DownloadTokenGenerateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateDownloadUrlForZipArchive**
> DownloadTokenGenerateResponse GenerateDownloadUrlForZipArchive(ctx, body, optional)
Generate download URL for ZIP download

### Description:   Create a download URL to retrieve several files in one ZIP archive.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions in parent room.  ### Postcondition: Download URL is generated and returned.  ### Further Information: The token is necessary to access `downloads` resources.   ZIP download is only available for files and folders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ZipDownloadRequest**](ZipDownloadRequest.md)|  | 
 **optional** | ***NodesApiGenerateDownloadUrlForZipArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiGenerateDownloadUrlForZipArchiveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**DownloadTokenGenerateResponse**](DownloadTokenGenerateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GeneratePresignedUrlsFiles**
> PresignedUrlList GeneratePresignedUrlsFiles(ctx, body, uploadId, optional)
Generate presigned URLs for S3 file upload

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.15.0</h3>  ### Description: Generate presigned URLs for S3 file upload.  ### Precondition: An upload channel has been created and user has to be the creator of the upload channel.  ### Postcondition: List of presigned URLs is returned.  ### Further Information: The size for each part must be >= 5 MB, except for the last part.   The part number of the first part in S3 is 1 (not 0).   Use HTTP method `PUT` for uploading bytes via presigned URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GeneratePresignedUrlsRequest**](GeneratePresignedUrlsRequest.md)|  | 
  **uploadId** | **string**| Upload channel ID | 
 **optional** | ***NodesApiGeneratePresignedUrlsFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiGeneratePresignedUrlsFilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**PresignedUrlList**](PresignedUrlList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HandleRoomWebhookAssignments**
> RoomWebhookList HandleRoomWebhookAssignments(ctx, body, roomId, optional)
Assign or unassign webhooks to room

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Handle room webhook assignments.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: List of webhooks is returned.  ### Further Information: None.  ### Available event types:  <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Name | Description | Scope | | :--- | :--- | :--- | | **`downloadshare.created`** | Triggered when a new download share is created in affected room | Node Webhook | | **`downloadshare.deleted`** | Triggered when a download share is deleted in affected room | Node Webhook | | **`downloadshare.used`** | Triggered when a download share is utilized in affected room | Node Webhook | | **`uploadshare.created`** | Triggered when a new upload share is created in affected room | Node Webhook | | **`uploadshare.deleted`** | Triggered when a upload share is deleted in affected room | Node Webhook | | **`uploadshare.used`** | Triggered when a new file is uploaded via the upload share in affected room | Node Webhook | | **`file.created`** | Triggered when a new file is uploaded in affected room | Node Webhook | | **`folder.created`** | Triggered when a new folder is created in affected room | Node Webhook | | **`room.created`** | Triggered when a new room is created (in affected room) | Node Webhook | | **`file.deleted`** | Triggered when a file is deleted in affected room | Node Webhook | | **`folder.deleted`** | Triggered when a folder is deleted in affected room | Node Webhook | | **`room.deleted`** | Triggered when a room is deleted in affected room | Node Webhook |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateRoomWebhookRequest**](UpdateRoomWebhookRequest.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiHandleRoomWebhookAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiHandleRoomWebhookAssignmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**RoomWebhookList**](RoomWebhookList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveNodes**
> Node MoveNodes(ctx, body, nodeId, optional)
Move node(s)

### Description:   Moves nodes (folder, file) to another parent.  ### Precondition: Authenticated user with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> and <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; delete</span> permissions in the source parent and <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; create</span> permissions in the target parent node.  ### Postcondition: Nodes are moved to target parent.  ### Further Information: Nodes **MUST** be in same source parent.   **Rooms **CANNOT** be moved.**  Download share id (if exists) gets changed if: - node with the same name exists in the target container - `resolutionStrategy` is `overwrite` - `keepShareLinks` is `true`  ### Node naming convention: * Node (room, folder, file) names are limited to **150** characters. * Illegal names:   `'CON', 'PRN', 'AUX', 'NUL', 'COM1', 'COM2', 'COM3', 'COM4', 'COM5', 'COM6', 'COM7', 'COM8', 'COM9', 'LPT1', 'LPT2', 'LPT3', 'LPT4', 'LPT5', 'LPT6', 'LPT7', 'LPT8', 'LPT9', (and any of those with an extension)` * Illegal characters in names:   `'\\\\', '<','>', ':', '\\\"', '|', '?', '*', '/', leading '-', trailing '.' ` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoveNodesRequest**](MoveNodesRequest.md)|  | 
  **nodeId** | **int64**| Target parent node ID | 
 **optional** | ***NodesApiMoveNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiMoveNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveDeletedNodes**
> RemoveDeletedNodes(ctx, body, optional)
Remove nodes from recycle bin

### Description: Permanently remove a list of nodes from the recycle bin.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; delete recycle bin</span> permissions in parent room.  ### Postcondition: All provided nodes are removed.  ### Further Information: The removal of deleted nodes from the recycle bin is irreversible.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteDeletedNodesRequest**](DeleteDeletedNodesRequest.md)|  | 
 **optional** | ***NodesApiRemoveDeletedNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRemoveDeletedNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFavorite**
> RemoveFavorite(ctx, nodeId, optional)
Unmark a node (room, folder or file) as favorite

### Description: Unmarks a node (room, folder or file) as favorite.  ### Precondition: Authenticated user is allowed to <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128065; see</span> the node (i.e. `isBrowsable = true`).  ### Postcondition: A node gets unmarked as favorite.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Node ID | 
 **optional** | ***NodesApiRemoveFavoriteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRemoveFavoriteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveNode**
> RemoveNode(ctx, nodeId, optional)
Remove node

### Description: Delete node (room, folder or file).  ### Precondition: Authenticated user with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; delete</span> permissions on supplied nodes (for folders or files) or on superordinated node (for rooms).  ### Postcondition: Node gets deleted.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Node ID | 
 **optional** | ***NodesApiRemoveNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRemoveNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveNodeComment**
> RemoveNodeComment(ctx, commentId, optional)
Remove node comment

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.10.0</h3>  ### Description: Delete an existing comment for a specific node.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions on the node and is the creator of the comment **OR** <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span> in auth parent room.  ### Postcondition: Comment is deleted.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commentId** | **int64**| Comment ID | 
 **optional** | ***NodesApiRemoveNodeCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRemoveNodeCommentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveNodes**
> RemoveNodes(ctx, body, optional)
Remove nodes

### Description: Delete nodes (room, folder or file).  ### Precondition: Authenticated user with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; delete</span> permissions on supplied nodes (for folders or files) or on superordinated node (for rooms).  ### Postcondition: Nodes are deleted.  ### Further Information: Nodes **MUST** be in same parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteNodesRequest**](DeleteNodesRequest.md)|  | 
 **optional** | ***NodesApiRemoveNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRemoveNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoomRescueKeyPair**
> RemoveRoomRescueKeyPair(ctx, roomId, optional)
Remove rooms's rescue key pair

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Delete room rescue key pair.  ### Precondition: Authenticated user.  ### Postcondition: Key pair is removed (cf. further information below).  ### Further Information: Please set a new room rescue key pair first and re-encrypt file keys with it.   If no version is set, deleted key pair with lowest preference value.   Although, `version` **SHOULD** be set. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRemoveRoomRescueKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRemoveRoomRescueKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| Version (NEW) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestDeletedNode**
> DeletedNode RequestDeletedNode(ctx, deletedNodeId, optional)
Request deleted node

### Description:   Get metadata of a deleted node.  ### Precondition: User can access parent room and has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read recycle bin</span> permissions.  ### Postcondition: Requested deleted node is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deletedNodeId** | **int64**| Deleted node ID | 
 **optional** | ***NodesApiRequestDeletedNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestDeletedNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**DeletedNode**](DeletedNode.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestDeletedNodeVersions**
> DeletedNodeVersionsList RequestDeletedNodeVersions(ctx, nodeId, type_, name, optional)
Request deleted versions of nodes

### Description:   Retrieve all deleted versions of a node.  ### Precondition: User can access parent room and has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read recycle bin</span> permissions.  ### Postcondition: List of deleted versions of a node is returned.  ### Further Information: The node is identified by three parameters: * parent ID * name * type (file, folder).  Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `expireAt:desc|size:asc`   Sort by `expireAt` descending **AND** `size` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `expireAt` | Expiration date | | `accessedAt` | Last access date | | `size` | Node size | | `classification` | Classification ID:<ul><li>1 - public</li><li>2 - internal</li><li>3 - confidential</li><li>4 - strictly confidential</li></ul> | | `createdAt` | Creation date | | `createdBy` | Creator first name, last name | | `updatedAt` | Last modification date | | `updatedBy` | Last modifier first name, last name | | `deletedAt` | Deleted date | | `deletedBy` | Deleter first name, last name |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Parent ID (room or folder ID) | 
  **type_** | **string**| Node type | 
  **name** | **string**| Node name | 
 **optional** | ***NodesApiRequestDeletedNodeVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestDeletedNodeVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **sort** | **optional.String**| Sort string | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**DeletedNodeVersionsList**](DeletedNodeVersionsList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestDeletedNodesSummary**
> DeletedNodeSummaryList RequestDeletedNodesSummary(ctx, nodeId, optional)
Request list of deleted nodes

### Description:   Retrieve a list of deleted nodes in a recycle bin.  ### Precondition: User can access parent room and has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read recycle bin</span> permissions.  ### Postcondition: List of deleted nodes is returned.  ### Further Information: Only room IDs are accepted as parent ID since only rooms may have a recycle bin.  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `type:eq:file:folder|name:cn:searchString_1|parentPath:cn:searchString_2`   Get deleted nodes where type equals (`file` **OR** `folder`) **AND** deleted node name containing `searchString_1` **AND** deleted node parent path containing `searchString 2`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `type` | Node type filter | `eq` | Node type equals value(s).<br>Multiple values are allowed and will be connected via logical disjunction (**OR**).<br>e.g. `type:eq:folder:file` | <ul><li>`folder`</li><li>`file`</li></ul> | | `name` | Node name filter | `cn` | Node name contains value. | `search String` | | `parentPath` | Parent path filter | `cn` | Parent path contains value. | `search String` | | `timestampCreation` | Creation timestamp filter | `ge, le` | Creation timestamp is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `timestampCreation:ge:2016-12-31`&#124;<br>`timestampCreation:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `timestampModification` | Modification timestamp filter | `ge, le` | Modification timestamp is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `timestampModification:ge:2016-12-31T23:00:00.123`&#124;<br>`timestampModification:le:2018-01-01T11:00:00.540` | `Date (yyyy-MM-dd)` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.   Nodes are sorted by type first, then by sent sort string.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:desc|timestampCreation:asc`   Sort by `name` descending **AND** `timestampCreation` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `name` | Node name | | `cntVersions` | Number of deleted versions of this file | | `firstDeletedAt` | First deleted version | | `lastDeletedAt` | Last deleted version | | `parentPath` | Parent path of deleted node | | `timestampCreation` | Creation timestamp | | `timestampModification` | Modification timestamp |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Parent ID (can only be a room ID) | 
 **optional** | ***NodesApiRequestDeletedNodesSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestDeletedNodesSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**DeletedNodeSummaryList**](DeletedNodeSummaryList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestListOfWebhooksForRoom**
> RoomWebhookList RequestListOfWebhooksForRoom(ctx, roomId, optional)
Request list of webhooks that are assigned or can be assigned to this room

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.19.0</h3>  ### Description:   Get a list of webhooks for the room scope with their assignment status.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: List of webhooks is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `isAssigned:eq:true`   Get a list of assigned webhooks to the room.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | **`isAssigned`** | Assigned/unassigned webhooks filter | `eq` |  | `true or false` |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRequestListOfWebhooksForRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestListOfWebhooksForRoomOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**RoomWebhookList**](RoomWebhookList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestMissingFileKeys**
> MissingKeysResponse RequestMissingFileKeys(ctx, optional)
Request files without user's file key

### Description:   Requests a list of missing file keys that may be generated by the current user.    ### Precondition: User has a key pair.   Only returns users that owns one of the following permissions: <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span>  ### Postcondition: None.  ### Further Information: Clients **SHOULD** regularly request missing file keys to provide access to files for other users.   The returned list is ordered by priority (emergency passwords / rescue keys are returned first). There is an enforced limit of **100** items per request. A total value greater than limit signals that there are more entries but does not necessarily reflect the precise number of total items. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NodesApiRequestMissingFileKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestMissingFileKeysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **roomId** | **optional.Int64**| Room ID | 
 **fileId** | **optional.Int64**| File ID | 
 **userId** | **optional.Int64**| User ID | 
 **useKey** | **optional.String**| Determines which key should be used (NEW) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**MissingKeysResponse**](MissingKeysResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestNode**
> Node RequestNode(ctx, nodeId, optional)
Request node

### Description:   Get node (room, folder or file).  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions in auth parent room.  ### Postcondition: Requested node is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Node ID | 
 **optional** | ***NodesApiRequestNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestNodeComments**
> CommentList RequestNodeComments(ctx, nodeId, optional)
Request list of node comments

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.10.0</h3>  ### Description: Get comments for a specific node.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions on the node.  ### Postcondition: List with comments (sorted by `createdAt` timestamp) is returned.  ### Further Information: An empty list is returned if no comments were found.   Output is limited to **500** entries.   For more results please use filter criteria and paging (`offset` + `limit`).  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Node ID | 
 **optional** | ***NodesApiRequestNodeCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestNodeCommentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **hideDeleted** | **optional.Bool**| Hide deleted comments (default: false) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**CommentList**](CommentList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestNodeParents**
> NodeParentList RequestNodeParents(ctx, nodeId, optional)
Request list of parent nodes

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.10.0</h3>  ### Description:   Requests a list of node ancestors, sorted from root node to the node's direct parent node.  ### Precondition: User is allowed to browse through the node tree until the requested node.  ### Postcondition: List of parent nodes is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **int64**| Node ID | 
 **optional** | ***NodesApiRequestNodeParentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestNodeParentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**NodeParentList**](NodeParentList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestNodes**
> NodeList RequestNodes(ctx, optional)
Request list of nodes

### Description:   Provides a hierarchical list of file system nodes (rooms, folders or files) of a given parent that are accessible by the current user.  ### Precondition: Authenticated user.  ### Postcondition: List of nodes is returned.  ### Further Information: `EncryptionInfo` is **NOT** provided.  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `type:eq:room:folder|perm:eq:read`   Get nodes where type equals (`room` **OR** `folder`) **AND** user has `read` permissions.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `type` | Node type filter | `eq` | Node type equals value.<br>Multiple values are allowed and will be connected via logical disjunction (**OR**).<br>e.g. `type:eq:room:folder` | <ul><li>`room`</li><li>`folder`</li><li>`file`</li></ul> | | `perm` | Permission filter | `eq` | Permission equals value.<br>Multiple values are allowed and will be connected via logical disjunction (**OR**).<br>e.g. `perm:eq:read:create:delete` | <ul><li>`manage`</li><li>`read`</li><li>`change`</li><li>`create`</li><li>`delete`</li><li>`manageDownloadShare`</li><li>`manageUploadShare`</li><li>`canReadRecycleBin`</li><li>`canRestoreRecycleBin`</li><li>`canDeleteRecycleBin`</li></ul> | | `childPerm` | Same as `perm`, but less restrictive (applies to child nodes only).<br>Child nodes of the parent node which do not meet the filter condition<br>are **NOT** returned. | `eq` | cf. `perm` | cf. `perm` | | `name` | Node name filter | `cn, eq` | Node name contains / equals value. | `search String` | | `encrypted` | Node encryption status filter | `eq` |  | `true or false` | | `branchVersion` | Node branch version filter | `ge, le` | Branch version is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `branchVersion:ge:1423280937404`&#124;`branchVersion:le:1523280937404` | `version number` | | `timestampCreation` | Creation timestamp filter | `ge, le` | Creation timestamp is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `timestampCreation:ge:2016-12-31T23:00:00.123`&#124;<br>`timestampCreation:le:2018-01-01T11:00:00.540` | `Date (yyyy-MM-dd)` | | `timestampModification` | Modification timestamp filter | `ge, le` | Modification timestamp is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `timestampModification:ge:2016-12-31T23:00:00.123`&#124;<br>`timestampModification:le:2018-01-01T11:00:00.540` | `Date (yyyy-MM-dd)` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.   Nodes are sorted by type first, then by sent sort string.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:desc|fileType:asc`   Sort by `name` descending **AND** `fileType` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `name` | Node name | | `createdAt` | Creation date | | `createdBy` | Creator first name, last name | | `updatedAt` | Last modification date | | `updatedBy` | Last modifier first name, last name | | `fileType` | File type (extension) | | `classification` | Classification ID:<ul><li>1 - public</li><li>2 - internal</li><li>3 - confidential</li><li>4 - strictly confidential</li></ul> | | `size` | Node size | | `cntDeletedVersions` | Number of deleted versions of this file / folder (**NOT** recursive; for files and folders only) | | `timestampCreation` | Creation timestamp | | `timestampModification` | Modification timestamp |  </details>  ### Deprecated sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | <del>`cntChildren`</del> | Number of direct children (**NOT** recursive; for rooms and folders only) |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NodesApiRequestNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **depthLevel** | **optional.Int32**| * &#x60;0&#x60; - top level nodes only  * &#x60;n&#x60; (any positive number) - include &#x60;n&#x60; levels starting from the current node | 
 **parentId** | **optional.Int64**| Parent node ID.  Only rooms and folders can be parents.  Parent ID &#x60;0&#x60; or empty is the root node. | 
 **roomManager** | **optional.Bool**| Show all rooms for management perspective.  Only possible for _Rooms Managers_.  For all other users, it will be ignored. | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**NodeList**](NodeList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestPendingAssignments**
> PendingAssignmentList RequestPendingAssignments(ctx, optional)
Request user-room assignments per group

### Description:   Requests a list of user-room assignments by groups that have **NOT** been approved yet   These can have the state: * **WAITING**   * **DENIED**   * **ACCEPTED**    **ACCEPTED** assignments are already removed from the list.  ### Precondition: None.  ### Postcondition: List of user-room assignments is returned.  ### Further Information: Room administrators **SHOULD** regularly request pending assingments to provide access to rooms for other users.  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `state:eq:WAITING`   Filter assignments by state `WAITING`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `userId` | User ID filter | `eq` | User ID equals value. | `positive Integer` | | `groupId` | Group ID filter | `eq` | Group ID equals value. | `positive Integer` | | `roomId` | Room ID filter | `eq` | Room ID equals value. | `positive Integer` | | `state` | Assignment state | `eq` | Assignment state equals value. | `WAITING or DENIED` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `userId:desc|state:asc`   Sort by `userId` descending **AND** `state` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `userId` | User ID | | `groupId` | Group ID | | `roomId` | Room ID | | `state` | State |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NodesApiRequestPendingAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestPendingAssignmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**PendingAssignmentList**](PendingAssignmentList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoomActivitiesLogAsJson**
> LogEventList RequestRoomActivitiesLogAsJson(ctx, roomId, optional)
Request events of a room

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.3.0</h3>  ### Description: Retrieve syslog (audit log) events related to a room.  ### Precondition: Requires <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions on that room.  ### Postcondition: List of events is returned.  ### Further Information: Output may be limited to a certain number of entries.   Please use filter criteria and paging.  Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort fields are supported.    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `time:desc`   Sort by `time` descending (default sort option).  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `time` | Event timestamp |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRequestRoomActivitiesLogAsJsonOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestRoomActivitiesLogAsJsonOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **sort** | **optional.String**| Sort string | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **dateStart** | **optional.String**| Filter events from given date   e.g. &#x60;2015-12-31T23:59:00&#x60; | 
 **dateEnd** | **optional.String**| Filter events until given date   e.g. &#x60;2015-12-31T23:59:00&#x60; | 
 **type_** | **optional.Int32**| Operation ID   cf. &#x60;GET /eventlog/operations&#x60; | 
 **userId** | **optional.Int64**| User ID | 
 **status** | **optional.Int32**| Operation status:  * &#x60;0&#x60; - Success  * &#x60;2&#x60; - Error | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**LogEventList**](LogEventList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoomGroups**
> RoomGroupList RequestRoomGroups(ctx, roomId, optional)
Request room granted group(s) or / and group(s) that can be granted

### Description:   Retrieve a list of groups that are and / or can be granted to the room.  ### Precondition: Any permissions on target room.  ### Postcondition: List of groups is returned.  ### Further Information:  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `isGranted:eq:false|name:cn:searchString`   Get all groups that are **NOT** granted to this room **AND** whose name is like `searchString`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `name` | Group name filter | `cn` | Group name contains value. | `search String` | | `groupId` | Group ID filter | `eq` | Group ID equals value. | `positive Integer` | | `isGranted` | Filter the groups that have (no) access to this room.<br>**This filter is only available for room administrators.**<br>**Other users can only look for groups in their rooms, so this filter is `true` and **CANNOT** be overridden.** | `eq` |  | <ul><li>`true`</li><li>`false`</li><li>`any`</li></ul>default: `true` | | `permissionsManage` | Filter the groups that do (not) have `manage` permissions in this room. | `eq` |  | `true or false` | | `effectivePerm` | Filter groups with DIRECT or DIRECT **AND** EFFECTIVE permissions<ul><li>`false`: DIRECT permissions</li><li>`true`: DIRECT **AND** EFFECTIVE permissions</li></ul>DIRECT means: e.g. room administrator grants `read` permissions to group of users **directly** on desired room.<br>EFFECTIVE means: e.g. group of users gets `read` permissions on desired room through **inheritance**. | `eq` |  | `true or false`<br>default: `false` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:desc`   Sort by `name` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `name` | Group name |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRequestRoomGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestRoomGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**RoomGroupList**](RoomGroupList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoomPolicies**
> RoomPolicies RequestRoomPolicies(ctx, roomId, optional)
Request Room Policies

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.32.0</h3>  ### Description:   Retrieve the room policies: * `defaultExpirationPeriod`  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions in that room.  ### Postcondition: Room Policies returned.  ### Further Information: `defaultExpirationPeriod`: Default policy room expiration period in seconds. All existing and future files in a room will have their expiration date set to this period after their respective upload. Existing files can be set to expire earlier afterwards. `0` means no default expiration policy will be enforced.    

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRequestRoomPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestRoomPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**RoomPolicies**](RoomPolicies.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoomRescueKey**
> FileKey RequestRoomRescueKey(ctx, fileId, optional)
Request room rescue key

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.24.0</h3>  ### Description:   Returns the file key for the room emergency password / rescue key of a certain file (if available).  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions in parent room.  ### Postcondition: File key is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **int64**| File ID | 
 **optional** | ***NodesApiRequestRoomRescueKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestRoomRescueKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| Version (NEW) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**FileKey**](FileKey.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoomRescueKeyPair**
> UserKeyPairContainer RequestRoomRescueKeyPair(ctx, roomId, optional)
Request room rescue key

### Description:   Retrieve the room rescue key pair.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions in that room.  ### Postcondition: Key pair is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRequestRoomRescueKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestRoomRescueKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **version** | **optional.String**| Version (NEW) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**UserKeyPairContainer**](UserKeyPairContainer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoomRescueKeyPairs**
> []UserKeyPairContainer RequestRoomRescueKeyPairs(ctx, roomId, optional)
Request all room rescue key pairs

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Retrieve all room rescue key pairs to allow migrating room-rescue-key-encrypted file keys.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions in that room.  ### Postcondition: List of key pairs is returned.  ### Further Information: In the case of an algorithm migration to a room rescue key pair, one should create the new key pair before deleting the old one. This allows re-encrypting file keys with the new key pair, using the old one.  This API allows to retrieve both key pairs, in contrast to `GET /nodes/rooms/{room_id}/keypair`, which only delivers the preferred one. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRequestRoomRescueKeyPairsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestRoomRescueKeyPairsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**[]UserKeyPairContainer**](UserKeyPairContainer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoomS3Tags**
> S3TagList RequestRoomS3Tags(ctx, roomId, optional)
Request list of all assigned S3 tags to the room

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.9.0</h3>  ### Description:   Retrieve a list of S3 tags assigned to a room.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: List of assigned S3 tags is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRequestRoomS3TagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestRoomS3TagsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**S3TagList**](S3TagList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRoomUsers**
> RoomUserList RequestRoomUsers(ctx, roomId, optional)
Request room granted user(s) or / and user(s) that can be granted

### Description:   Retrieve a list of users that are and / or can be granted to the room.  ### Precondition: Any permissions on target room.  ### Postcondition: None.  ### Further Information: List of users is returned.  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  > `permissionsManage:eq:true|user:cn:searchString`   Get all users that have `manage` permissions to this room **AND** whose (`firstName` **OR** `lastName` **OR** `email` **OR** `username`) is like `searchString`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `user` | User filter | `cn` | User contains value (`firstName` **OR** `lastName` **OR** `email` **OR** `username`). | `search String` | | `userId` | User ID filter | `eq` | User ID equals value. | `positive Integer` | | `isGranted` | Filter the users that have (no) access to this room.<br>**This filter is only available for room administrators.**<br>**Other users can only look for users in their rooms, so this filter is `true` and **CANNOT** be overridden.** | `eq` |  | <ul><li>`true`</li><li>`false`</li><li>`any`</li></ul>default: `true` | | `permissionsManage` | Filter the users that do (not) have `manage` permissions in this room. | `eq` |  | `true or false` | | `effectivePerm` | Filter users with DIRECT or DIRECT **AND** EFFECTIVE permissions<ul><li>`false`: DIRECT permissions</li><li>`true`: DIRECT **AND** EFFECTIVE permissions</li><li>`any`: DIRECT **AND** EFFECTIVE **AND** OVER GROUP permissions</li></ul>DIRECT means: e.g. room administrator grants `read` permissions to group of users **directly** on desired room.<br>EFFECTIVE means: e.g. group of users gets `read` permissions on desired room through **inheritance**.<br>OVER GROUP means: e.g. user gets `read` permissions on desired room through **group membership**. | `eq` |  | <ul><li>`true`</li><li>`false`</li><li>`any`</li></ul>default: `false` |  </details>  ### Deprecated filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | <del>`displayName`</del> | User display name filter (use `user` filter) | `cn` | User display name contains value (`firstName` **OR** `lastName` **OR** `email`). | `search String` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `user:desc`   Sort by `user` descending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | **`user`** | User - sort by `firstName`, `lastName`, `username`, `email` (in this order) |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRequestRoomUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestRoomUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**RoomUserList**](RoomUserList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSystemRescueKey**
> FileKey RequestSystemRescueKey(ctx, fileId, optional)
Request system rescue key

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.24.0</h3>  ### Description:   Returns the file key for the system emergency password / rescue key of a certain file (if available).  ### Precondition: User with <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions in parent room.  ### Postcondition: File key is returned.  ### Further Information: None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **int64**| File ID | 
 **optional** | ***NodesApiRequestSystemRescueKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestSystemRescueKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| Version (NEW) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**FileKey**](FileKey.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUploadStatusFiles**
> S3FileUploadStatus RequestUploadStatusFiles(ctx, uploadId, optional)
Request status of S3 file upload

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.15.0</h3>  ### Description: Request status of a S3 file upload.  ### Precondition: An upload channel has been created and user has to be the creator of the upload channel.  ### Postcondition: Status of S3 multipart upload request is returned.  ### Further Information: None.  ### Possible errors: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | Http Status | Error Code | Description | | :--- | :--- | :--- | | `400 Bad Request` | `-80000` | Mandatory fields cannot be empty | | `400 Bad Request` | `-80001` | Invalid positive number | | `400 Bad Request` | `-80002` | Invalid number | | `400 Bad Request` | `-40001` | (Target) room is not encrypted | | `400 Bad Request` | `-40755` | Bad file name | | `400 Bad Request` | `-40763` | File key must be set for an upload into encrypted room | | `400 Bad Request` | `-50506` | Exceeds the number of files for this Upload Share | | `403 Forbidden` |  | Access denied | | `404 Not Found` | `-20501` | Upload not found | | `404 Not Found` | `-40000` | Container not found | | `404 Not Found` | `-41000` | Node not found | | `404 Not Found` | `-70501` | User not found | | `409 Conflict` | `-40010` | Container cannot be overwritten | | `409 Conflict` |  | File cannot be overwritten | | `500 Internal Server Error` |  | System Error | | `502 Bad Gateway` |  | S3 Error | | `502 Insufficient Storage` | `-50504` | Exceeds the quota for this Upload Share | | `502 Insufficient Storage` | `-40200` | Exceeds the free node quota in room | | `502 Insufficient Storage` | `-90200` | Exceeds the free customer quota | | `502 Insufficient Storage` | `-90201` | Exceeds the free customer physical disk space |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uploadId** | **string**| Upload channel ID | 
 **optional** | ***NodesApiRequestUploadStatusFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestUploadStatusFilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**S3FileUploadStatus**](S3FileUploadStatus.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestUserFileKey**
> FileKey RequestUserFileKey(ctx, fileId, optional)
Request user's file key

### Description:   Returns the file key for the current user (if available).  ### Precondition: User with one of the following permissions in parent room: <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span>  ### Postcondition: File key is returned.  ### Further Information: The symmetric file key is encrypted with the user's public key.   File keys are generated with the workflow _\"Generate file keys\"_ that starts at `GET /nodes/missingFileKeys`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **int64**| File ID | 
 **optional** | ***NodesApiRequestUserFileKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRequestUserFileKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| Version (NEW) | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**FileKey**](FileKey.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreNodes**
> RestoreNodes(ctx, body, optional)
Restore deleted nodes

### Description:   Restore a list of deleted nodes.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; create</span> permissions in parent room and <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; restore recycle bin</span> permissions.  ### Postcondition: The selected files are moved from the recycle bin to the chosen productive container.  ### Further Information: If no parent ID is provided, the node is restored to its previous location.   The default resolution strategy is `autorename` that adds numbers to the file name until the conflict is solved.   If an existing file is overwritten, it is moved to the recycle bin instead of the restored one.  Download share id (if exists) gets changed if: - node with the same name exists in the target container - `resolutionStrategy` is `overwrite` - `keepShareLinks` is `true`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RestoreDeletedNodesRequest**](RestoreDeletedNodesRequest.md)|  | 
 **optional** | ***NodesApiRestoreNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRestoreNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeRoomGroups**
> RevokeRoomGroups(ctx, body, roomId, optional)
Revoke granted group(s) from room

### Description:   Revoke granted groups from room.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: Group's permissions are revoked.  ### Further Information: Batch function.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoomGroupsDeleteBatchRequest**](RoomGroupsDeleteBatchRequest.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRevokeRoomGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRevokeRoomGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeRoomUsers**
> RevokeRoomUsers(ctx, body, roomId, optional)
Revoke granted user(s) from room

### Description:   Revoke granted users from room.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: User's permissions are revoked.  ### Further Information: Batch function.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoomUsersDeleteBatchRequest**](RoomUsersDeleteBatchRequest.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiRevokeRoomUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiRevokeRoomUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchNodes**
> NodeList SearchNodes(ctx, searchString, optional)
Search nodes

### Description:   Provides a flat list of file system nodes (rooms, folders or files) of a given parent that are accessible by the current user.  ### Precondition: Authenticated user is allowed to <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128065; see</span> nodes (i.e. `isBrowsable = true`).  ### Postcondition: List of nodes is returned.  ### Further Information:   Output is limited to **500** entries.   For more results please use filter criteria and paging (`offset` + `limit`).   `EncryptionInfo` is **NOT** provided.   Wildcard character is the asterisk character: `*`  ### Filtering: All filter fields are connected via logical conjunction (**AND**)   Filter string syntax: `FIELD_NAME:OPERATOR:VALUE[:VALUE...]`    <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `type:eq:file|createdAt:ge:2015-01-01`   Get nodes where type equals `file` **AND** file creation date is **>=** `2015-01-01`.  </details>  ### Filtering options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Filter Description | `OPERATOR` | Operator Description | `VALUE` | | :--- | :--- | :--- | :--- | :--- | | `type` | Node type filter | `eq` | Node type equals value.<br>Multiple values are allowed and will be connected via logical disjunction (**OR**).<br>e.g. `type:eq:room:folder` | <ul><li>`room`</li><li>`folder`</li><li>`file`</li></ul> | | `fileType` | File type filter (file extension) | `cn, eq` | File type contains / equals value. | `search String` | | `classification` | Classification filter | `eq` | Classification equals value. | <ul><li>`1` - public</li><li>`2` - internal</li><li>`3` - confidential</li><li>`4` - strictly confidential</li></ul> | | `createdBy` | Creator login filter | `cn, eq` | Creator login contains / equals value (`firstName` **OR** `lastName` **OR** `email` **OR** `username`). | `search String` | | `createdById` | Creator ID filter | `eq` | Creator ID equals value. | `positive Integer  or -1 for external user` | | `createdAt` | Creation date filter | `ge, le` | Creation date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `createdAt:ge:2016-12-31`&#124;`createdAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `updatedBy` | Last modifier login filter | `cn, eq` | Last modifier login contains / equals value (`firstName` **OR** `lastName` **OR** `email` **OR** `username`). | `search String` | | `updatedById` | Last modifier ID filter | `eq` | Modifier ID equals value. | `positive Integer or -1 for external user` | | `updatedAt` | Last modification date filter | `ge, le` | Last modification date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `updatedAt:ge:2016-12-31`&#124;`updatedAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `expireAt` | Expiration date filter | `ge, le` | Expiration date is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `expireAt:ge:2016-12-31`&#124;`expireAt:le:2018-01-01` | `Date (yyyy-MM-dd)` | | `size` | Node size filter | `ge, le` | Node size is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `size:ge:5`&#124;`size:le:10` | `size in bytes` | | `isFavorite` | Favorite filter | `eq` |  | `true or false` | | `branchVersion` | Node branch version filter | `ge, le` | Branch version is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `branchVersion:ge:1423280937404`&#124;`branchVersion:le:1523280937404` | `version number` | | `parentPath` | Parent path | `cn, eq` | Parent path contains / equals  value. | `search String` | | `timestampCreation` | Creation timestamp filter | `ge, le` | Creation timestamp is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `timestampCreation:ge:2016-12-31T23:00:00.123`&#124;<br>`timestampCreation:le:2018-01-01T11:00:00.540` | `Date (yyyy-MM-dd)` | | `timestampModification` | Modification timestamp filter | `ge, le` | Modification timestamp is greater / less equals than value.<br>Multiple operator values are allowed and will be connected via logical conjunction (**AND**).<br>e.g. `timestampModification:ge:2016-12-31T23:00:00.123`&#124;<br>`timestampModification:le:2018-01-01T11:00:00.540` | `Date (yyyy-MM-dd)` |  </details>  ---  ### Sorting: Sort string syntax: `FIELD_NAME:ORDER`   `ORDER` can be `asc` or `desc`.   Multiple sort criteria are possible.   Fields are connected via logical conjunction **AND**.  <details style=\"padding-left: 10px\"> <summary style=\"cursor: pointer; outline: none\"><strong>Example</strong></summary>  `name:desc|size:asc`   Sort by `name` descending **AND** `size` ascending.  </details>  ### Sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | `name` | Node name | | `createdAt` | Creation date | | `createdBy` | Creator first name, last name | | `updatedAt` | Last modification date | | `updatedBy` | Last modifier first name, last name | | `fileType` | File type (extension) | | `classification` | Classification ID:<ul><li>1 - public</li><li>2 - internal</li><li>3 - confidential</li><li>4 - strictly confidential</li></ul> | | `size` | Node size | | `cntDeletedVersions` | Number of deleted versions of this file / folder (**NOT** recursive; for files and folders only) | | `type` | Node type (room, folder, file) | | `parentPath` | Parent path | | `timestampCreation` | Creation timestamp | | `timestampModification` | Modification timestamp |  </details>  ### Deprecated sorting options: <details style=\"padding: 10px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px;\"> <summary style=\"cursor: pointer; outline: none\"><strong>Expand</strong></summary>  | `FIELD_NAME` | Description | | :--- | :--- | | <del>`cntChildren`</del> | Number of direct children (**NOT** recursive; for rooms and folders only) |  </details>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchString** | **string**| Search string | 
 **optional** | ***NodesApiSearchNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiSearchNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsDateFormat** | **optional.String**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **depthLevel** | **optional.Int32**| * &#x60;0&#x60; - top level nodes only (default)  * &#x60;-1&#x60; - full tree  * &#x60;n&#x60; (any positive number) - include &#x60;n&#x60; levels starting from the current node | 
 **parentId** | **optional.Int64**| Parent node ID.  Only rooms and folders can be parents.  Parent ID &#x60;0&#x60; or empty is the root node. | 
 **filter** | **optional.String**| Filter string | 
 **sort** | **optional.String**| Sort string | 
 **offset** | **optional.Int32**| Range offset | 
 **limit** | **optional.Int32**| Range limit.  Maximum 500.   For more results please use paging (&#x60;offset&#x60; + &#x60;limit&#x60;). | 
 **xSdsAuthToken** | **optional.String**| Authentication token | 

### Return type

[**NodeList**](NodeList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRoomPolicies**
> SetRoomPolicies(ctx, body, roomId, optional)
Set room policies

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.32.0</h3>  ### Description:   Retrieve the room policies: * `defaultExpirationPeriod`  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: Room policy is set.  ### Further Information: `defaultExpirationPeriod`: Default policy room expiration period in seconds. All existing and future files in a room will have their expiration date set to this period after their respective upload. Existing files can be set to expire earlier afterwards. `0` means no default expiration policy will be enforced. This removes all expiration dates from existing files.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoomPoliciesRequest**](RoomPoliciesRequest.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiSetRoomPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiSetRoomPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRoomRescueKeyPair**
> SetRoomRescueKeyPair(ctx, body, roomId, optional)
Set room's rescue key pair

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.24.0</h3>  ### Description:   Set room rescue key pair.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: Key pair is set.  ### Further Information: Room rescue key pair can be used to upgrade algorithm.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserKeyPairContainer**](UserKeyPairContainer.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiSetRoomRescueKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiSetRoomRescueKeyPairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRoomS3Tags**
> S3TagList SetRoomS3Tags(ctx, body, roomId, optional)
Set S3 tags for a room

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.9.0</h3>  ### Description:   Set S3 tags to a room.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>.  ### Postcondition: Provided S3 tags are assigned to a room.  ### Further Information: Every request overrides current S3 tags.   Mandatory S3 tag IDs **MUST** be sent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**S3TagIds**](S3TagIds.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiSetRoomS3TagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiSetRoomS3TagsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**S3TagList**](S3TagList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetUserFileKeys**
> SetUserFileKeys(ctx, body, optional)
Set file keys for a list of users and files

### Description:   Sets symmetric file keys for several users and files.  ### Precondition: User has file keys for the files.   Only settable by users that own one of the following permissions: <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; manage download share</span>, <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change config</span>  ### Postcondition: Stores new file keys for other users.  ### Further Information: Only users with copies of the file key (encrypted with their public keys) can access a certain file.   This endpoint is used for the distribution of file keys amongst an authorized user base.   User can set file key for himself.   The users who already have a file key are ignored and keep the distributed file key 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserFileKeySetBatchRequest**](UserFileKeySetBatchRequest.md)|  | 
 **optional** | ***NodesApiSetUserFileKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiSetUserFileKeysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFavorites**
> UpdateFavorites(ctx, body, optional)
Mark or unmark a list of nodes (room, folder or file) as favorite

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.25.0</h3>  ### Description:   Marks or unmarks a list of nodes (room, folder or file) as favorite.  ### Precondition: Authenticated user is allowed to <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128065; see</span> the node (i.e. `isBrowsable = true`).  ### Postcondition: Nodes gets marked as favorite.  ### Further Information: Maximum number of nodes is 200.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFavoritesBulkRequest**](UpdateFavoritesBulkRequest.md)|  | 
 **optional** | ***NodesApiUpdateFavoritesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiUpdateFavoritesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFile**
> Node UpdateFile(ctx, body, fileId, optional)
Updates a file’s metadata

### Description: Updates a list of file’s metadata.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change</span> permissions in parent room.  ### Postcondition: File's metadata is changed.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFileRequest**](UpdateFileRequest.md)|  | 
  **fileId** | **int64**| File ID | 
 **optional** | ***NodesApiUpdateFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiUpdateFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFiles**
> UpdateFiles(ctx, body, optional)
Updates a list of  file’s metadata

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.25.0</h3>  ### Description:   Updates a list of file’s metadata.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change</span> permissions in parent room.  ### Postcondition: File's metadata is changed.  ### Further Information: Maximum number of shares is 200 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFilesBulkRequest**](UpdateFilesBulkRequest.md)|  | 
 **optional** | ***NodesApiUpdateFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiUpdateFilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFolder**
> Node UpdateFolder(ctx, body, folderId, optional)
Updates folder’s metadata

### Description:   Updates folder’s metadata.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; change</span> permissions in parent room.  ### Postcondition: Folder's metadata is changed.  ### Further Information: Notes are limited to **255** characters.  ### Node naming convention: * Node (room, folder, file) names are limited to **150** characters. * Illegal names:   `'CON', 'PRN', 'AUX', 'NUL', 'COM1', 'COM2', 'COM3', 'COM4', 'COM5', 'COM6', 'COM7', 'COM8', 'COM9', 'LPT1', 'LPT2', 'LPT3', 'LPT4', 'LPT5', 'LPT6', 'LPT7', 'LPT8', 'LPT9', (and any of those with an extension)` * Illegal characters in names:   `'\\\\', '<','>', ':', '\\\"', '|', '?', '*', '/', leading '-', trailing '.' ` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFolderRequest**](UpdateFolderRequest.md)|  | 
  **folderId** | **int64**| Folder ID | 
 **optional** | ***NodesApiUpdateFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiUpdateFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeComment**
> Comment UpdateNodeComment(ctx, body, commentId, optional)
Edit node comment

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128640; Since v4.10.0</h3>  ### Description: Edit the text of an existing comment for a specific node.  ### Precondition: User has <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; read</span> permissions on the node and is the creator of the comment.  ### Postcondition: Comments text gets changed.  ### Further Information: Maximum allowed text length: **65535** characters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeNodeCommentRequest**](ChangeNodeCommentRequest.md)|  | 
  **commentId** | **int64**| Comment ID | 
 **optional** | ***NodesApiUpdateNodeCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiUpdateNodeCommentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Comment**](Comment.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoom**
> Node UpdateRoom(ctx, body, roomId, optional)
Updates room’s metadata

### Description:   Updates room’s metadata.  ### Precondition: User is a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span> at superordinated level.  ### Postcondition: Room's metadata is changed.  ### Further Information: Notes are limited to **255** characters.  ### Node naming convention: * Node (room, folder, file) names are limited to **150** characters. * Illegal names:   `'CON', 'PRN', 'AUX', 'NUL', 'COM1', 'COM2', 'COM3', 'COM4', 'COM5', 'COM6', 'COM7', 'COM8', 'COM9', 'LPT1', 'LPT2', 'LPT3', 'LPT4', 'LPT5', 'LPT6', 'LPT7', 'LPT8', 'LPT9', (and any of those with an extension)` * Illegal characters in names:   `'\\\\', '<','>', ':', '\\\"', '|', '?', '*', '/', leading '-', trailing '.' `

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateRoomRequest**](UpdateRoomRequest.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiUpdateRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiUpdateRoomOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsDateFormat** | **optional.**| Date time format (cf. [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) &amp; [leettime.de](http://leettime.de/)) | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**Node**](Node.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoomGroups**
> UpdateRoomGroups(ctx, body, roomId, optional)
Add or change room granted group(s)

### Description: All existing group permissions will be overwritten.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>. To add new members, the user needs the right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; non-members add</span>, which is included in any role.  ### Postcondition: Group's permissions are changed.  ### Further Information: Batch function.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoomGroupsAddBatchRequest**](RoomGroupsAddBatchRequest.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiUpdateRoomGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiUpdateRoomGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoomUsers**
> UpdateRoomUsers(ctx, body, roomId, optional)
Add or change room granted user(s)

### Description: All existing user permissions will be overwritten.  ### Precondition: User needs to be a <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128100; Room Administrator</span>. To add new members, the user needs the right <span style='padding: 3px; background-color: #F6F7F8; border: 1px solid #000; border-radius: 5px; display: inline;'>&#128275; non-members add</span>, which is included in any role.  ### Postcondition: User's permissions are changed.  ### Further Information: Batch function.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoomUsersAddBatchRequest**](RoomUsersAddBatchRequest.md)|  | 
  **roomId** | **int64**| Room ID | 
 **optional** | ***NodesApiUpdateRoomUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiUpdateRoomUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFileAsMultipart**
> ChunkUploadResponse UploadFileAsMultipart(ctx, uploadId, optional)
Upload file

<h3 style='padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;'>&#128679; Deprecated since v4.9.0</h3>  ### Use `uploads` API  ### Description:   Uploads a file or parts of it in an active upload channel.  ### Precondition: An upload channel has been created.  ### Postcondition: A file or parts of it are uploaded to a temporary location.  ### Further Information: This endpoints supports chunked upload.    Following `Content-Types` are supported by this API: * `multipart/form-data` * provided `Content-Type`     For both file upload types set the correct `Content-Type` header and body.    ### Examples:    * `multipart/form-data` ``` POST /api/v4/nodes/files/uploads/{upload_id} HTTP/1.1  Header: ... Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW ...  Body: ------WebKitFormBoundary7MA4YWxkTrZu0gW Content-Disposition: form-data; name=\"file\"; filename=\"file.txt\" Content-Type: text/plain  Content of file.txt ------WebKitFormBoundary7MA4YWxkTrZu0gW-- ```  * any other `Content-Type`   ``` POST /api/v4/nodes/files/uploads/{upload_id}  HTTP/1.1  Header: ... Content-Type: { ... } ...  Body: raw content ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uploadId** | **string**| Upload channel ID | 
 **optional** | ***NodesApiUploadFileAsMultipartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodesApiUploadFileAsMultipartOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 
 **contentRange** | **optional.**| Content-Range   e.g. &#x60;bytes 0-999/3980&#x60; | 
 **xSdsAuthToken** | **optional.**| Authentication token | 

### Return type

[**ChunkUploadResponse**](ChunkUploadResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

