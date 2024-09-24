# DfsDirectoryRespDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsRootfs** | [**DfsRootfs**](DfsRootfs.md) |  | 
**DirectoryResult** | [**DfsDirectoryResultResp**](DfsDirectoryResultResp.md) |  | 
**StoragePolicyIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewDfsDirectoryRespDirectory

`func NewDfsDirectoryRespDirectory(dfsRootfs DfsRootfs, directoryResult DfsDirectoryResultResp, ) *DfsDirectoryRespDirectory`

NewDfsDirectoryRespDirectory instantiates a new DfsDirectoryRespDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsDirectoryRespDirectoryWithDefaults

`func NewDfsDirectoryRespDirectoryWithDefaults() *DfsDirectoryRespDirectory`

NewDfsDirectoryRespDirectoryWithDefaults instantiates a new DfsDirectoryRespDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsRootfs

`func (o *DfsDirectoryRespDirectory) GetDfsRootfs() DfsRootfs`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsDirectoryRespDirectory) GetDfsRootfsOk() (*DfsRootfs, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsDirectoryRespDirectory) SetDfsRootfs(v DfsRootfs)`

SetDfsRootfs sets DfsRootfs field to given value.


### GetDirectoryResult

`func (o *DfsDirectoryRespDirectory) GetDirectoryResult() DfsDirectoryResultResp`

GetDirectoryResult returns the DirectoryResult field if non-nil, zero value otherwise.

### GetDirectoryResultOk

`func (o *DfsDirectoryRespDirectory) GetDirectoryResultOk() (*DfsDirectoryResultResp, bool)`

GetDirectoryResultOk returns a tuple with the DirectoryResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResult

`func (o *DfsDirectoryRespDirectory) SetDirectoryResult(v DfsDirectoryResultResp)`

SetDirectoryResult sets DirectoryResult field to given value.


### GetStoragePolicyIds

`func (o *DfsDirectoryRespDirectory) GetStoragePolicyIds() []int64`

GetStoragePolicyIds returns the StoragePolicyIds field if non-nil, zero value otherwise.

### GetStoragePolicyIdsOk

`func (o *DfsDirectoryRespDirectory) GetStoragePolicyIdsOk() (*[]int64, bool)`

GetStoragePolicyIdsOk returns a tuple with the StoragePolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyIds

`func (o *DfsDirectoryRespDirectory) SetStoragePolicyIds(v []int64)`

SetStoragePolicyIds sets StoragePolicyIds field to given value.

### HasStoragePolicyIds

`func (o *DfsDirectoryRespDirectory) HasStoragePolicyIds() bool`

HasStoragePolicyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


