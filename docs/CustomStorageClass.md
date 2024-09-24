# CustomStorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of custom class | [optional] 
**Name** | **string** | name of custom class | 
**PoolPolicies** | [**[]PoolPolicy**](PoolPolicy.md) | active pool policy array of custom class | 
**ScId** | **int64** | class unique id(1-7) | 
**WritePolicy** | **string** | write policy of custom class | 

## Methods

### NewCustomStorageClass

`func NewCustomStorageClass(name string, poolPolicies []PoolPolicy, scId int64, writePolicy string, ) *CustomStorageClass`

NewCustomStorageClass instantiates a new CustomStorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomStorageClassWithDefaults

`func NewCustomStorageClassWithDefaults() *CustomStorageClass`

NewCustomStorageClassWithDefaults instantiates a new CustomStorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CustomStorageClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomStorageClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomStorageClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomStorageClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CustomStorageClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomStorageClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomStorageClass) SetName(v string)`

SetName sets Name field to given value.


### GetPoolPolicies

`func (o *CustomStorageClass) GetPoolPolicies() []PoolPolicy`

GetPoolPolicies returns the PoolPolicies field if non-nil, zero value otherwise.

### GetPoolPoliciesOk

`func (o *CustomStorageClass) GetPoolPoliciesOk() (*[]PoolPolicy, bool)`

GetPoolPoliciesOk returns a tuple with the PoolPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPolicies

`func (o *CustomStorageClass) SetPoolPolicies(v []PoolPolicy)`

SetPoolPolicies sets PoolPolicies field to given value.


### GetScId

`func (o *CustomStorageClass) GetScId() int64`

GetScId returns the ScId field if non-nil, zero value otherwise.

### GetScIdOk

`func (o *CustomStorageClass) GetScIdOk() (*int64, bool)`

GetScIdOk returns a tuple with the ScId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScId

`func (o *CustomStorageClass) SetScId(v int64)`

SetScId sets ScId field to given value.


### GetWritePolicy

`func (o *CustomStorageClass) GetWritePolicy() string`

GetWritePolicy returns the WritePolicy field if non-nil, zero value otherwise.

### GetWritePolicyOk

`func (o *CustomStorageClass) GetWritePolicyOk() (*string, bool)`

GetWritePolicyOk returns a tuple with the WritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicy

`func (o *CustomStorageClass) SetWritePolicy(v string)`

SetWritePolicy sets WritePolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


