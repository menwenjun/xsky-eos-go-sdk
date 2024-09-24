# DfsStoragePolicyUpdateReqDfsStoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | descirption of storage policy | [optional] 
**Filter** | Pointer to [**DfsPolicyFilter**](DfsPolicyFilter.md) |  | [optional] 
**Name** | Pointer to **string** | name of storage policy | [optional] 
**RootfsId** | **int64** | id of rootfs | 
**ScId** | **int64** | scid of storage class | 

## Methods

### NewDfsStoragePolicyUpdateReqDfsStoragePolicy

`func NewDfsStoragePolicyUpdateReqDfsStoragePolicy(rootfsId int64, scId int64, ) *DfsStoragePolicyUpdateReqDfsStoragePolicy`

NewDfsStoragePolicyUpdateReqDfsStoragePolicy instantiates a new DfsStoragePolicyUpdateReqDfsStoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsStoragePolicyUpdateReqDfsStoragePolicyWithDefaults

`func NewDfsStoragePolicyUpdateReqDfsStoragePolicyWithDefaults() *DfsStoragePolicyUpdateReqDfsStoragePolicy`

NewDfsStoragePolicyUpdateReqDfsStoragePolicyWithDefaults instantiates a new DfsStoragePolicyUpdateReqDfsStoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilter

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) GetFilter() DfsPolicyFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) GetFilterOk() (*DfsPolicyFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) SetFilter(v DfsPolicyFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetName

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootfsId

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) GetRootfsId() int64`

GetRootfsId returns the RootfsId field if non-nil, zero value otherwise.

### GetRootfsIdOk

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) GetRootfsIdOk() (*int64, bool)`

GetRootfsIdOk returns a tuple with the RootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsId

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) SetRootfsId(v int64)`

SetRootfsId sets RootfsId field to given value.


### GetScId

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) GetScId() int64`

GetScId returns the ScId field if non-nil, zero value otherwise.

### GetScIdOk

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) GetScIdOk() (*int64, bool)`

GetScIdOk returns a tuple with the ScId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScId

`func (o *DfsStoragePolicyUpdateReqDfsStoragePolicy) SetScId(v int64)`

SetScId sets ScId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


