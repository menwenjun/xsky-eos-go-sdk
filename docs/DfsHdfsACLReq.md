# DfsHdfsACLReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of user or user group | [optional] 
**FsUserGroupId** | Pointer to **int64** | fs user group id | [optional] 
**FsUserGroupName** | Pointer to **string** | name of user group | [optional] 
**FsUserId** | Pointer to **int64** | fs user id | [optional] 
**FsUserName** | Pointer to **string** | name of user | [optional] 
**Permission** | Pointer to **string** | permission of user or group | [optional] 
**Security** | Pointer to **string** | acl security type | [optional] 
**Type** | Pointer to **string** | acl type | [optional] 

## Methods

### NewDfsHdfsACLReq

`func NewDfsHdfsACLReq() *DfsHdfsACLReq`

NewDfsHdfsACLReq instantiates a new DfsHdfsACLReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsACLReqWithDefaults

`func NewDfsHdfsACLReqWithDefaults() *DfsHdfsACLReq`

NewDfsHdfsACLReqWithDefaults instantiates a new DfsHdfsACLReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsHdfsACLReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsHdfsACLReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsHdfsACLReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsHdfsACLReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFsUserGroupId

`func (o *DfsHdfsACLReq) GetFsUserGroupId() int64`

GetFsUserGroupId returns the FsUserGroupId field if non-nil, zero value otherwise.

### GetFsUserGroupIdOk

`func (o *DfsHdfsACLReq) GetFsUserGroupIdOk() (*int64, bool)`

GetFsUserGroupIdOk returns a tuple with the FsUserGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroupId

`func (o *DfsHdfsACLReq) SetFsUserGroupId(v int64)`

SetFsUserGroupId sets FsUserGroupId field to given value.

### HasFsUserGroupId

`func (o *DfsHdfsACLReq) HasFsUserGroupId() bool`

HasFsUserGroupId returns a boolean if a field has been set.

### GetFsUserGroupName

`func (o *DfsHdfsACLReq) GetFsUserGroupName() string`

GetFsUserGroupName returns the FsUserGroupName field if non-nil, zero value otherwise.

### GetFsUserGroupNameOk

`func (o *DfsHdfsACLReq) GetFsUserGroupNameOk() (*string, bool)`

GetFsUserGroupNameOk returns a tuple with the FsUserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroupName

`func (o *DfsHdfsACLReq) SetFsUserGroupName(v string)`

SetFsUserGroupName sets FsUserGroupName field to given value.

### HasFsUserGroupName

`func (o *DfsHdfsACLReq) HasFsUserGroupName() bool`

HasFsUserGroupName returns a boolean if a field has been set.

### GetFsUserId

`func (o *DfsHdfsACLReq) GetFsUserId() int64`

GetFsUserId returns the FsUserId field if non-nil, zero value otherwise.

### GetFsUserIdOk

`func (o *DfsHdfsACLReq) GetFsUserIdOk() (*int64, bool)`

GetFsUserIdOk returns a tuple with the FsUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserId

`func (o *DfsHdfsACLReq) SetFsUserId(v int64)`

SetFsUserId sets FsUserId field to given value.

### HasFsUserId

`func (o *DfsHdfsACLReq) HasFsUserId() bool`

HasFsUserId returns a boolean if a field has been set.

### GetFsUserName

`func (o *DfsHdfsACLReq) GetFsUserName() string`

GetFsUserName returns the FsUserName field if non-nil, zero value otherwise.

### GetFsUserNameOk

`func (o *DfsHdfsACLReq) GetFsUserNameOk() (*string, bool)`

GetFsUserNameOk returns a tuple with the FsUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserName

`func (o *DfsHdfsACLReq) SetFsUserName(v string)`

SetFsUserName sets FsUserName field to given value.

### HasFsUserName

`func (o *DfsHdfsACLReq) HasFsUserName() bool`

HasFsUserName returns a boolean if a field has been set.

### GetPermission

`func (o *DfsHdfsACLReq) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *DfsHdfsACLReq) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *DfsHdfsACLReq) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *DfsHdfsACLReq) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetSecurity

`func (o *DfsHdfsACLReq) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DfsHdfsACLReq) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DfsHdfsACLReq) SetSecurity(v string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DfsHdfsACLReq) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetType

`func (o *DfsHdfsACLReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsHdfsACLReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsHdfsACLReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsHdfsACLReq) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


