# DfsDirectoryReqDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clean** | Pointer to **bool** | clean resources with directory, used when deleting | [optional] 
**DfsRootfsId** | **int64** | dfs rootfs id | 
**Path** | **string** | directory path | 
**Recursive** | Pointer to **bool** | recursive create directory, used when creating | [optional] 
**StoragePolicyIds** | Pointer to **[]int64** | storage policy id array | [optional] 

## Methods

### NewDfsDirectoryReqDirectory

`func NewDfsDirectoryReqDirectory(dfsRootfsId int64, path string, ) *DfsDirectoryReqDirectory`

NewDfsDirectoryReqDirectory instantiates a new DfsDirectoryReqDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsDirectoryReqDirectoryWithDefaults

`func NewDfsDirectoryReqDirectoryWithDefaults() *DfsDirectoryReqDirectory`

NewDfsDirectoryReqDirectoryWithDefaults instantiates a new DfsDirectoryReqDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClean

`func (o *DfsDirectoryReqDirectory) GetClean() bool`

GetClean returns the Clean field if non-nil, zero value otherwise.

### GetCleanOk

`func (o *DfsDirectoryReqDirectory) GetCleanOk() (*bool, bool)`

GetCleanOk returns a tuple with the Clean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClean

`func (o *DfsDirectoryReqDirectory) SetClean(v bool)`

SetClean sets Clean field to given value.

### HasClean

`func (o *DfsDirectoryReqDirectory) HasClean() bool`

HasClean returns a boolean if a field has been set.

### GetDfsRootfsId

`func (o *DfsDirectoryReqDirectory) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsDirectoryReqDirectory) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsDirectoryReqDirectory) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetPath

`func (o *DfsDirectoryReqDirectory) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsDirectoryReqDirectory) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsDirectoryReqDirectory) SetPath(v string)`

SetPath sets Path field to given value.


### GetRecursive

`func (o *DfsDirectoryReqDirectory) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *DfsDirectoryReqDirectory) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *DfsDirectoryReqDirectory) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *DfsDirectoryReqDirectory) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetStoragePolicyIds

`func (o *DfsDirectoryReqDirectory) GetStoragePolicyIds() []int64`

GetStoragePolicyIds returns the StoragePolicyIds field if non-nil, zero value otherwise.

### GetStoragePolicyIdsOk

`func (o *DfsDirectoryReqDirectory) GetStoragePolicyIdsOk() (*[]int64, bool)`

GetStoragePolicyIdsOk returns a tuple with the StoragePolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyIds

`func (o *DfsDirectoryReqDirectory) SetStoragePolicyIds(v []int64)`

SetStoragePolicyIds sets StoragePolicyIds field to given value.

### HasStoragePolicyIds

`func (o *DfsDirectoryReqDirectory) HasStoragePolicyIds() bool`

HasStoragePolicyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


