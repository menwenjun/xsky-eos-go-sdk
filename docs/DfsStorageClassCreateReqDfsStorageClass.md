# DfsStorageClassCreateReqDfsStorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of custom class | [optional] 
**Name** | **string** | name of custom class | 
**PoolPolicies** | [**[]PoolPolicy**](PoolPolicy.md) | active pool policy array of custom class | 
**RootfsId** | **int64** | id of rootfs | 
**ScId** | **int64** | class unique id | 
**WritePolicy** | **string** | write policy of custom class | 

## Methods

### NewDfsStorageClassCreateReqDfsStorageClass

`func NewDfsStorageClassCreateReqDfsStorageClass(name string, poolPolicies []PoolPolicy, rootfsId int64, scId int64, writePolicy string, ) *DfsStorageClassCreateReqDfsStorageClass`

NewDfsStorageClassCreateReqDfsStorageClass instantiates a new DfsStorageClassCreateReqDfsStorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsStorageClassCreateReqDfsStorageClassWithDefaults

`func NewDfsStorageClassCreateReqDfsStorageClassWithDefaults() *DfsStorageClassCreateReqDfsStorageClass`

NewDfsStorageClassCreateReqDfsStorageClassWithDefaults instantiates a new DfsStorageClassCreateReqDfsStorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsStorageClassCreateReqDfsStorageClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsStorageClassCreateReqDfsStorageClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsStorageClassCreateReqDfsStorageClass) SetName(v string)`

SetName sets Name field to given value.


### GetPoolPolicies

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetPoolPolicies() []PoolPolicy`

GetPoolPolicies returns the PoolPolicies field if non-nil, zero value otherwise.

### GetPoolPoliciesOk

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetPoolPoliciesOk() (*[]PoolPolicy, bool)`

GetPoolPoliciesOk returns a tuple with the PoolPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPolicies

`func (o *DfsStorageClassCreateReqDfsStorageClass) SetPoolPolicies(v []PoolPolicy)`

SetPoolPolicies sets PoolPolicies field to given value.


### GetRootfsId

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetRootfsId() int64`

GetRootfsId returns the RootfsId field if non-nil, zero value otherwise.

### GetRootfsIdOk

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetRootfsIdOk() (*int64, bool)`

GetRootfsIdOk returns a tuple with the RootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsId

`func (o *DfsStorageClassCreateReqDfsStorageClass) SetRootfsId(v int64)`

SetRootfsId sets RootfsId field to given value.


### GetScId

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetScId() int64`

GetScId returns the ScId field if non-nil, zero value otherwise.

### GetScIdOk

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetScIdOk() (*int64, bool)`

GetScIdOk returns a tuple with the ScId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScId

`func (o *DfsStorageClassCreateReqDfsStorageClass) SetScId(v int64)`

SetScId sets ScId field to given value.


### GetWritePolicy

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetWritePolicy() string`

GetWritePolicy returns the WritePolicy field if non-nil, zero value otherwise.

### GetWritePolicyOk

`func (o *DfsStorageClassCreateReqDfsStorageClass) GetWritePolicyOk() (*string, bool)`

GetWritePolicyOk returns a tuple with the WritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicy

`func (o *DfsStorageClassCreateReqDfsStorageClass) SetWritePolicy(v string)`

SetWritePolicy sets WritePolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


