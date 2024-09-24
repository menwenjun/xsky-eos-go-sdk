# PoolsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pools** | [**[]PoolRecord**](PoolRecord.md) | pools | 

## Methods

### NewPoolsResp

`func NewPoolsResp(pools []PoolRecord, ) *PoolsResp`

NewPoolsResp instantiates a new PoolsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolsRespWithDefaults

`func NewPoolsRespWithDefaults() *PoolsResp`

NewPoolsRespWithDefaults instantiates a new PoolsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPools

`func (o *PoolsResp) GetPools() []PoolRecord`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *PoolsResp) GetPoolsOk() (*[]PoolRecord, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *PoolsResp) SetPools(v []PoolRecord)`

SetPools sets Pools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


