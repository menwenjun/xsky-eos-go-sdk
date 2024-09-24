# DfsStoragePolicyCreateReqDfsStoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | descirption of storage policy | [optional] 
**Filter** | Pointer to [**DfsPolicyFilter**](DfsPolicyFilter.md) |  | [optional] 
**Name** | **string** | name of storage policy | 
**RootfsId** | **int64** | id of rootfs | 
**ScId** | **int64** | scid of storage class | 

## Methods

### NewDfsStoragePolicyCreateReqDfsStoragePolicy

`func NewDfsStoragePolicyCreateReqDfsStoragePolicy(name string, rootfsId int64, scId int64, ) *DfsStoragePolicyCreateReqDfsStoragePolicy`

NewDfsStoragePolicyCreateReqDfsStoragePolicy instantiates a new DfsStoragePolicyCreateReqDfsStoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsStoragePolicyCreateReqDfsStoragePolicyWithDefaults

`func NewDfsStoragePolicyCreateReqDfsStoragePolicyWithDefaults() *DfsStoragePolicyCreateReqDfsStoragePolicy`

NewDfsStoragePolicyCreateReqDfsStoragePolicyWithDefaults instantiates a new DfsStoragePolicyCreateReqDfsStoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilter

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) GetFilter() DfsPolicyFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) GetFilterOk() (*DfsPolicyFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) SetFilter(v DfsPolicyFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetName

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) SetName(v string)`

SetName sets Name field to given value.


### GetRootfsId

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) GetRootfsId() int64`

GetRootfsId returns the RootfsId field if non-nil, zero value otherwise.

### GetRootfsIdOk

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) GetRootfsIdOk() (*int64, bool)`

GetRootfsIdOk returns a tuple with the RootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsId

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) SetRootfsId(v int64)`

SetRootfsId sets RootfsId field to given value.


### GetScId

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) GetScId() int64`

GetScId returns the ScId field if non-nil, zero value otherwise.

### GetScIdOk

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) GetScIdOk() (*int64, bool)`

GetScIdOk returns a tuple with the ScId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScId

`func (o *DfsStoragePolicyCreateReqDfsStoragePolicy) SetScId(v int64)`

SetScId sets ScId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


