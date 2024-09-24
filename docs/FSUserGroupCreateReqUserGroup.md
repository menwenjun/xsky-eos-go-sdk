# FSUserGroupCreateReqUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of user group | [optional] 
**FsUserIds** | Pointer to **[]int64** | ids of users, which are required when type is smb or ftp | [optional] 
**GroupId** | Pointer to **int64** | gid of user group | [optional] 
**Name** | **string** | name of user group | 

## Methods

### NewFSUserGroupCreateReqUserGroup

`func NewFSUserGroupCreateReqUserGroup(name string, ) *FSUserGroupCreateReqUserGroup`

NewFSUserGroupCreateReqUserGroup instantiates a new FSUserGroupCreateReqUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSUserGroupCreateReqUserGroupWithDefaults

`func NewFSUserGroupCreateReqUserGroupWithDefaults() *FSUserGroupCreateReqUserGroup`

NewFSUserGroupCreateReqUserGroupWithDefaults instantiates a new FSUserGroupCreateReqUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FSUserGroupCreateReqUserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FSUserGroupCreateReqUserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FSUserGroupCreateReqUserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FSUserGroupCreateReqUserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFsUserIds

`func (o *FSUserGroupCreateReqUserGroup) GetFsUserIds() []int64`

GetFsUserIds returns the FsUserIds field if non-nil, zero value otherwise.

### GetFsUserIdsOk

`func (o *FSUserGroupCreateReqUserGroup) GetFsUserIdsOk() (*[]int64, bool)`

GetFsUserIdsOk returns a tuple with the FsUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserIds

`func (o *FSUserGroupCreateReqUserGroup) SetFsUserIds(v []int64)`

SetFsUserIds sets FsUserIds field to given value.

### HasFsUserIds

`func (o *FSUserGroupCreateReqUserGroup) HasFsUserIds() bool`

HasFsUserIds returns a boolean if a field has been set.

### GetGroupId

`func (o *FSUserGroupCreateReqUserGroup) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FSUserGroupCreateReqUserGroup) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FSUserGroupCreateReqUserGroup) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *FSUserGroupCreateReqUserGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *FSUserGroupCreateReqUserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSUserGroupCreateReqUserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSUserGroupCreateReqUserGroup) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


