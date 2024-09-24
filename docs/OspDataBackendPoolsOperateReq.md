# OspDataBackendPoolsOperateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **int64** |  | [optional] 
**DataBackendId** | Pointer to **int64** |  | [optional] 
**PoolIds** | Pointer to **[]int64** |  | [optional] 
**Pools** | Pointer to **map[string]bool** |  | [optional] 

## Methods

### NewOspDataBackendPoolsOperateReq

`func NewOspDataBackendPoolsOperateReq() *OspDataBackendPoolsOperateReq`

NewOspDataBackendPoolsOperateReq instantiates a new OspDataBackendPoolsOperateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspDataBackendPoolsOperateReqWithDefaults

`func NewOspDataBackendPoolsOperateReqWithDefaults() *OspDataBackendPoolsOperateReq`

NewOspDataBackendPoolsOperateReqWithDefaults instantiates a new OspDataBackendPoolsOperateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *OspDataBackendPoolsOperateReq) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *OspDataBackendPoolsOperateReq) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *OspDataBackendPoolsOperateReq) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *OspDataBackendPoolsOperateReq) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetDataBackendId

`func (o *OspDataBackendPoolsOperateReq) GetDataBackendId() int64`

GetDataBackendId returns the DataBackendId field if non-nil, zero value otherwise.

### GetDataBackendIdOk

`func (o *OspDataBackendPoolsOperateReq) GetDataBackendIdOk() (*int64, bool)`

GetDataBackendIdOk returns a tuple with the DataBackendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBackendId

`func (o *OspDataBackendPoolsOperateReq) SetDataBackendId(v int64)`

SetDataBackendId sets DataBackendId field to given value.

### HasDataBackendId

`func (o *OspDataBackendPoolsOperateReq) HasDataBackendId() bool`

HasDataBackendId returns a boolean if a field has been set.

### GetPoolIds

`func (o *OspDataBackendPoolsOperateReq) GetPoolIds() []int64`

GetPoolIds returns the PoolIds field if non-nil, zero value otherwise.

### GetPoolIdsOk

`func (o *OspDataBackendPoolsOperateReq) GetPoolIdsOk() (*[]int64, bool)`

GetPoolIdsOk returns a tuple with the PoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIds

`func (o *OspDataBackendPoolsOperateReq) SetPoolIds(v []int64)`

SetPoolIds sets PoolIds field to given value.

### HasPoolIds

`func (o *OspDataBackendPoolsOperateReq) HasPoolIds() bool`

HasPoolIds returns a boolean if a field has been set.

### GetPools

`func (o *OspDataBackendPoolsOperateReq) GetPools() map[string]bool`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *OspDataBackendPoolsOperateReq) GetPoolsOk() (*map[string]bool, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *OspDataBackendPoolsOperateReq) SetPools(v map[string]bool)`

SetPools sets Pools field to given value.

### HasPools

`func (o *OspDataBackendPoolsOperateReq) HasPools() bool`

HasPools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


