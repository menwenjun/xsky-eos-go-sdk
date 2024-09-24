# FSUserGroupRemoveUsersReqUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsUserIds** | Pointer to **[]int64** | ids of users, which are required when type is smb or ftp | [optional] 

## Methods

### NewFSUserGroupRemoveUsersReqUserGroup

`func NewFSUserGroupRemoveUsersReqUserGroup() *FSUserGroupRemoveUsersReqUserGroup`

NewFSUserGroupRemoveUsersReqUserGroup instantiates a new FSUserGroupRemoveUsersReqUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSUserGroupRemoveUsersReqUserGroupWithDefaults

`func NewFSUserGroupRemoveUsersReqUserGroupWithDefaults() *FSUserGroupRemoveUsersReqUserGroup`

NewFSUserGroupRemoveUsersReqUserGroupWithDefaults instantiates a new FSUserGroupRemoveUsersReqUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsUserIds

`func (o *FSUserGroupRemoveUsersReqUserGroup) GetFsUserIds() []int64`

GetFsUserIds returns the FsUserIds field if non-nil, zero value otherwise.

### GetFsUserIdsOk

`func (o *FSUserGroupRemoveUsersReqUserGroup) GetFsUserIdsOk() (*[]int64, bool)`

GetFsUserIdsOk returns a tuple with the FsUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserIds

`func (o *FSUserGroupRemoveUsersReqUserGroup) SetFsUserIds(v []int64)`

SetFsUserIds sets FsUserIds field to given value.

### HasFsUserIds

`func (o *FSUserGroupRemoveUsersReqUserGroup) HasFsUserIds() bool`

HasFsUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


