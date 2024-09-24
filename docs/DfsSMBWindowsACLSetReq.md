# DfsSMBWindowsACLSetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsRootfsId** | **int64** | id of rootfs | 
**DfsSmbWindowsAcls** | [**[]DfsSMBWindowsACLSet**](DfsSMBWindowsACLSet.md) | dfs smb windows acls | 
**InheritanceType** | Pointer to **string** | inheritance type | [optional] 
**LimitHandler** | Pointer to **bool** | limit handler for sub directory | [optional] 
**Path** | **string** | directory or subdirectory or sub file of share path | 
**ReplacePermission** | Pointer to **bool** | whether replace permission | [optional] 

## Methods

### NewDfsSMBWindowsACLSetReq

`func NewDfsSMBWindowsACLSetReq(dfsRootfsId int64, dfsSmbWindowsAcls []DfsSMBWindowsACLSet, path string, ) *DfsSMBWindowsACLSetReq`

NewDfsSMBWindowsACLSetReq instantiates a new DfsSMBWindowsACLSetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBWindowsACLSetReqWithDefaults

`func NewDfsSMBWindowsACLSetReqWithDefaults() *DfsSMBWindowsACLSetReq`

NewDfsSMBWindowsACLSetReqWithDefaults instantiates a new DfsSMBWindowsACLSetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsRootfsId

`func (o *DfsSMBWindowsACLSetReq) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsSMBWindowsACLSetReq) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsSMBWindowsACLSetReq) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetDfsSmbWindowsAcls

`func (o *DfsSMBWindowsACLSetReq) GetDfsSmbWindowsAcls() []DfsSMBWindowsACLSet`

GetDfsSmbWindowsAcls returns the DfsSmbWindowsAcls field if non-nil, zero value otherwise.

### GetDfsSmbWindowsAclsOk

`func (o *DfsSMBWindowsACLSetReq) GetDfsSmbWindowsAclsOk() (*[]DfsSMBWindowsACLSet, bool)`

GetDfsSmbWindowsAclsOk returns a tuple with the DfsSmbWindowsAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsSmbWindowsAcls

`func (o *DfsSMBWindowsACLSetReq) SetDfsSmbWindowsAcls(v []DfsSMBWindowsACLSet)`

SetDfsSmbWindowsAcls sets DfsSmbWindowsAcls field to given value.


### GetInheritanceType

`func (o *DfsSMBWindowsACLSetReq) GetInheritanceType() string`

GetInheritanceType returns the InheritanceType field if non-nil, zero value otherwise.

### GetInheritanceTypeOk

`func (o *DfsSMBWindowsACLSetReq) GetInheritanceTypeOk() (*string, bool)`

GetInheritanceTypeOk returns a tuple with the InheritanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceType

`func (o *DfsSMBWindowsACLSetReq) SetInheritanceType(v string)`

SetInheritanceType sets InheritanceType field to given value.

### HasInheritanceType

`func (o *DfsSMBWindowsACLSetReq) HasInheritanceType() bool`

HasInheritanceType returns a boolean if a field has been set.

### GetLimitHandler

`func (o *DfsSMBWindowsACLSetReq) GetLimitHandler() bool`

GetLimitHandler returns the LimitHandler field if non-nil, zero value otherwise.

### GetLimitHandlerOk

`func (o *DfsSMBWindowsACLSetReq) GetLimitHandlerOk() (*bool, bool)`

GetLimitHandlerOk returns a tuple with the LimitHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitHandler

`func (o *DfsSMBWindowsACLSetReq) SetLimitHandler(v bool)`

SetLimitHandler sets LimitHandler field to given value.

### HasLimitHandler

`func (o *DfsSMBWindowsACLSetReq) HasLimitHandler() bool`

HasLimitHandler returns a boolean if a field has been set.

### GetPath

`func (o *DfsSMBWindowsACLSetReq) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSMBWindowsACLSetReq) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSMBWindowsACLSetReq) SetPath(v string)`

SetPath sets Path field to given value.


### GetReplacePermission

`func (o *DfsSMBWindowsACLSetReq) GetReplacePermission() bool`

GetReplacePermission returns the ReplacePermission field if non-nil, zero value otherwise.

### GetReplacePermissionOk

`func (o *DfsSMBWindowsACLSetReq) GetReplacePermissionOk() (*bool, bool)`

GetReplacePermissionOk returns a tuple with the ReplacePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePermission

`func (o *DfsSMBWindowsACLSetReq) SetReplacePermission(v bool)`

SetReplacePermission sets ReplacePermission field to given value.

### HasReplacePermission

`func (o *DfsSMBWindowsACLSetReq) HasReplacePermission() bool`

HasReplacePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


