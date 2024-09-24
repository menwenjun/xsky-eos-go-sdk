# DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootfsId** | **int64** |  | 
**UnlinkPaths** | **[]string** | storage policy to cancel link with paths | 

## Methods

### NewDfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy

`func NewDfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy(rootfsId int64, unlinkPaths []string, ) *DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy`

NewDfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy instantiates a new DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicyWithDefaults

`func NewDfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicyWithDefaults() *DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy`

NewDfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicyWithDefaults instantiates a new DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootfsId

`func (o *DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy) GetRootfsId() int64`

GetRootfsId returns the RootfsId field if non-nil, zero value otherwise.

### GetRootfsIdOk

`func (o *DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy) GetRootfsIdOk() (*int64, bool)`

GetRootfsIdOk returns a tuple with the RootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsId

`func (o *DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy) SetRootfsId(v int64)`

SetRootfsId sets RootfsId field to given value.


### GetUnlinkPaths

`func (o *DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy) GetUnlinkPaths() []string`

GetUnlinkPaths returns the UnlinkPaths field if non-nil, zero value otherwise.

### GetUnlinkPathsOk

`func (o *DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy) GetUnlinkPathsOk() (*[]string, bool)`

GetUnlinkPathsOk returns a tuple with the UnlinkPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlinkPaths

`func (o *DfsStoragePolicyUnlinkDfsPathReqDfsStoragePolicy) SetUnlinkPaths(v []string)`

SetUnlinkPaths sets UnlinkPaths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


