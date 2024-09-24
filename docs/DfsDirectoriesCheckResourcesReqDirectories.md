# DfsDirectoriesCheckResourcesReqDirectories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckCur** | Pointer to **bool** | check cur dir | [optional] 
**CheckParent** | Pointer to **bool** | check parent dir | [optional] 
**CheckSub** | Pointer to **bool** | check sub dir | [optional] 
**DfsRootfsId** | **int64** | dfs rootfs id | 
**Directories** | **[]string** |  | 
**Resources** | Pointer to **[]string** | check resources | [optional] 

## Methods

### NewDfsDirectoriesCheckResourcesReqDirectories

`func NewDfsDirectoriesCheckResourcesReqDirectories(dfsRootfsId int64, directories []string, ) *DfsDirectoriesCheckResourcesReqDirectories`

NewDfsDirectoriesCheckResourcesReqDirectories instantiates a new DfsDirectoriesCheckResourcesReqDirectories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsDirectoriesCheckResourcesReqDirectoriesWithDefaults

`func NewDfsDirectoriesCheckResourcesReqDirectoriesWithDefaults() *DfsDirectoriesCheckResourcesReqDirectories`

NewDfsDirectoriesCheckResourcesReqDirectoriesWithDefaults instantiates a new DfsDirectoriesCheckResourcesReqDirectories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckCur

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetCheckCur() bool`

GetCheckCur returns the CheckCur field if non-nil, zero value otherwise.

### GetCheckCurOk

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetCheckCurOk() (*bool, bool)`

GetCheckCurOk returns a tuple with the CheckCur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCur

`func (o *DfsDirectoriesCheckResourcesReqDirectories) SetCheckCur(v bool)`

SetCheckCur sets CheckCur field to given value.

### HasCheckCur

`func (o *DfsDirectoriesCheckResourcesReqDirectories) HasCheckCur() bool`

HasCheckCur returns a boolean if a field has been set.

### GetCheckParent

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetCheckParent() bool`

GetCheckParent returns the CheckParent field if non-nil, zero value otherwise.

### GetCheckParentOk

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetCheckParentOk() (*bool, bool)`

GetCheckParentOk returns a tuple with the CheckParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckParent

`func (o *DfsDirectoriesCheckResourcesReqDirectories) SetCheckParent(v bool)`

SetCheckParent sets CheckParent field to given value.

### HasCheckParent

`func (o *DfsDirectoriesCheckResourcesReqDirectories) HasCheckParent() bool`

HasCheckParent returns a boolean if a field has been set.

### GetCheckSub

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetCheckSub() bool`

GetCheckSub returns the CheckSub field if non-nil, zero value otherwise.

### GetCheckSubOk

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetCheckSubOk() (*bool, bool)`

GetCheckSubOk returns a tuple with the CheckSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSub

`func (o *DfsDirectoriesCheckResourcesReqDirectories) SetCheckSub(v bool)`

SetCheckSub sets CheckSub field to given value.

### HasCheckSub

`func (o *DfsDirectoriesCheckResourcesReqDirectories) HasCheckSub() bool`

HasCheckSub returns a boolean if a field has been set.

### GetDfsRootfsId

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsDirectoriesCheckResourcesReqDirectories) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetDirectories

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetDirectories() []string`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetDirectoriesOk() (*[]string, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *DfsDirectoriesCheckResourcesReqDirectories) SetDirectories(v []string)`

SetDirectories sets Directories field to given value.


### GetResources

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DfsDirectoriesCheckResourcesReqDirectories) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DfsDirectoriesCheckResourcesReqDirectories) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DfsDirectoriesCheckResourcesReqDirectories) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


