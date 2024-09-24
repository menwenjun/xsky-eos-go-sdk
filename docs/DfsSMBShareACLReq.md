# DfsSMBShareACLReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of user or user group | [optional] 
**FsUserGroupId** | Pointer to **int64** | id of user group | [optional] 
**FsUserGroupName** | Pointer to **string** | id of user group | [optional] 
**FsUserId** | Pointer to **int64** | id of user | [optional] 
**FsUserName** | Pointer to **string** | id of user | [optional] 
**Id** | Pointer to **int64** | id of user group | [optional] 
**Permission** | Pointer to **string** | readonly or readwrite access | [optional] 
**Security** | Pointer to **string** | security of share acl | [optional] 
**Type** | Pointer to **string** | type of share acl | [optional] 

## Methods

### NewDfsSMBShareACLReq

`func NewDfsSMBShareACLReq() *DfsSMBShareACLReq`

NewDfsSMBShareACLReq instantiates a new DfsSMBShareACLReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBShareACLReqWithDefaults

`func NewDfsSMBShareACLReqWithDefaults() *DfsSMBShareACLReq`

NewDfsSMBShareACLReqWithDefaults instantiates a new DfsSMBShareACLReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsSMBShareACLReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsSMBShareACLReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsSMBShareACLReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsSMBShareACLReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFsUserGroupId

`func (o *DfsSMBShareACLReq) GetFsUserGroupId() int64`

GetFsUserGroupId returns the FsUserGroupId field if non-nil, zero value otherwise.

### GetFsUserGroupIdOk

`func (o *DfsSMBShareACLReq) GetFsUserGroupIdOk() (*int64, bool)`

GetFsUserGroupIdOk returns a tuple with the FsUserGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroupId

`func (o *DfsSMBShareACLReq) SetFsUserGroupId(v int64)`

SetFsUserGroupId sets FsUserGroupId field to given value.

### HasFsUserGroupId

`func (o *DfsSMBShareACLReq) HasFsUserGroupId() bool`

HasFsUserGroupId returns a boolean if a field has been set.

### GetFsUserGroupName

`func (o *DfsSMBShareACLReq) GetFsUserGroupName() string`

GetFsUserGroupName returns the FsUserGroupName field if non-nil, zero value otherwise.

### GetFsUserGroupNameOk

`func (o *DfsSMBShareACLReq) GetFsUserGroupNameOk() (*string, bool)`

GetFsUserGroupNameOk returns a tuple with the FsUserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroupName

`func (o *DfsSMBShareACLReq) SetFsUserGroupName(v string)`

SetFsUserGroupName sets FsUserGroupName field to given value.

### HasFsUserGroupName

`func (o *DfsSMBShareACLReq) HasFsUserGroupName() bool`

HasFsUserGroupName returns a boolean if a field has been set.

### GetFsUserId

`func (o *DfsSMBShareACLReq) GetFsUserId() int64`

GetFsUserId returns the FsUserId field if non-nil, zero value otherwise.

### GetFsUserIdOk

`func (o *DfsSMBShareACLReq) GetFsUserIdOk() (*int64, bool)`

GetFsUserIdOk returns a tuple with the FsUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserId

`func (o *DfsSMBShareACLReq) SetFsUserId(v int64)`

SetFsUserId sets FsUserId field to given value.

### HasFsUserId

`func (o *DfsSMBShareACLReq) HasFsUserId() bool`

HasFsUserId returns a boolean if a field has been set.

### GetFsUserName

`func (o *DfsSMBShareACLReq) GetFsUserName() string`

GetFsUserName returns the FsUserName field if non-nil, zero value otherwise.

### GetFsUserNameOk

`func (o *DfsSMBShareACLReq) GetFsUserNameOk() (*string, bool)`

GetFsUserNameOk returns a tuple with the FsUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserName

`func (o *DfsSMBShareACLReq) SetFsUserName(v string)`

SetFsUserName sets FsUserName field to given value.

### HasFsUserName

`func (o *DfsSMBShareACLReq) HasFsUserName() bool`

HasFsUserName returns a boolean if a field has been set.

### GetId

`func (o *DfsSMBShareACLReq) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsSMBShareACLReq) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsSMBShareACLReq) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsSMBShareACLReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermission

`func (o *DfsSMBShareACLReq) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *DfsSMBShareACLReq) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *DfsSMBShareACLReq) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *DfsSMBShareACLReq) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetSecurity

`func (o *DfsSMBShareACLReq) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DfsSMBShareACLReq) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DfsSMBShareACLReq) SetSecurity(v string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DfsSMBShareACLReq) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetType

`func (o *DfsSMBShareACLReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsSMBShareACLReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsSMBShareACLReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsSMBShareACLReq) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


