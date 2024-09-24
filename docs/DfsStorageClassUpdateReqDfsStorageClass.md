# DfsStorageClassUpdateReqDfsStorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of custom class | [optional] 
**Name** | Pointer to **string** | name of custom class | [optional] 
**PoolPolicies** | Pointer to [**[]PoolPolicy**](PoolPolicy.md) | active pool policy array of custom class | [optional] 
**WritePolicy** | Pointer to **string** | write policy of custom class | [optional] 

## Methods

### NewDfsStorageClassUpdateReqDfsStorageClass

`func NewDfsStorageClassUpdateReqDfsStorageClass() *DfsStorageClassUpdateReqDfsStorageClass`

NewDfsStorageClassUpdateReqDfsStorageClass instantiates a new DfsStorageClassUpdateReqDfsStorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsStorageClassUpdateReqDfsStorageClassWithDefaults

`func NewDfsStorageClassUpdateReqDfsStorageClassWithDefaults() *DfsStorageClassUpdateReqDfsStorageClass`

NewDfsStorageClassUpdateReqDfsStorageClassWithDefaults instantiates a new DfsStorageClassUpdateReqDfsStorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsStorageClassUpdateReqDfsStorageClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsStorageClassUpdateReqDfsStorageClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsStorageClassUpdateReqDfsStorageClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsStorageClassUpdateReqDfsStorageClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DfsStorageClassUpdateReqDfsStorageClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsStorageClassUpdateReqDfsStorageClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsStorageClassUpdateReqDfsStorageClass) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsStorageClassUpdateReqDfsStorageClass) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoolPolicies

`func (o *DfsStorageClassUpdateReqDfsStorageClass) GetPoolPolicies() []PoolPolicy`

GetPoolPolicies returns the PoolPolicies field if non-nil, zero value otherwise.

### GetPoolPoliciesOk

`func (o *DfsStorageClassUpdateReqDfsStorageClass) GetPoolPoliciesOk() (*[]PoolPolicy, bool)`

GetPoolPoliciesOk returns a tuple with the PoolPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPolicies

`func (o *DfsStorageClassUpdateReqDfsStorageClass) SetPoolPolicies(v []PoolPolicy)`

SetPoolPolicies sets PoolPolicies field to given value.

### HasPoolPolicies

`func (o *DfsStorageClassUpdateReqDfsStorageClass) HasPoolPolicies() bool`

HasPoolPolicies returns a boolean if a field has been set.

### GetWritePolicy

`func (o *DfsStorageClassUpdateReqDfsStorageClass) GetWritePolicy() string`

GetWritePolicy returns the WritePolicy field if non-nil, zero value otherwise.

### GetWritePolicyOk

`func (o *DfsStorageClassUpdateReqDfsStorageClass) GetWritePolicyOk() (*string, bool)`

GetWritePolicyOk returns a tuple with the WritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicy

`func (o *DfsStorageClassUpdateReqDfsStorageClass) SetWritePolicy(v string)`

SetWritePolicy sets WritePolicy field to given value.

### HasWritePolicy

`func (o *DfsStorageClassUpdateReqDfsStorageClass) HasWritePolicy() bool`

HasWritePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


