# NodePermissions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manage** | **bool** | User / Group may grant all of the above permissions to other users and groups independently,  may update room metadata and create / update / delete subordinary rooms, has all permissions. | [default to null]
**Read** | **bool** | User / Group may see all rooms, files and folders in the room and download everything, copy files from this room. | [default to null]
**Create** | **bool** | User / Group may upload files, create folders and copy / move files to this room, overwriting is not possible. | [default to null]
**Change** | **bool** | User / Group may update metadata of nodes: rename files and folders, change classification, etc. | [default to null]
**Delete** | **bool** | User / Group may overwrite and remove files / folders, move files from this room. | [default to null]
**ManageDownloadShare** | **bool** | User / Group may create Download Shares for files and containers view all previously created Download Shares in this room. | [default to null]
**ManageUploadShare** | **bool** | User / Group may create Upload Shares for containers, view all previously created Upload Shares in this room. | [default to null]
**ReadRecycleBin** | **bool** | User / Group may look up files / folders in the recycle bin. | [default to null]
**RestoreRecycleBin** | **bool** | User / Group may restore files / folders from recycle bin - room permissions required. | [default to null]
**DeleteRecycleBin** | **bool** | User / Group may permanently remove files / folders from the recycle bin. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

