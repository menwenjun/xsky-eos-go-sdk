# PoolCapacityReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodingChunkNum** | Pointer to **int64** |  | [optional] 
**DataChunkNum** | Pointer to **int64** |  | [optional] 
**OsdIds** | Pointer to **[]int64** |  | [optional] 
**PoolType** | Pointer to **string** |  | [optional] 
**ReplicateSize** | Pointer to **int64** |  | [optional] 
**ThinProvisioned** | Pointer to **bool** |  | [optional] 

## Methods

### NewPoolCapacityReq

`func NewPoolCapacityReq() *PoolCapacityReq`

NewPoolCapacityReq instantiates a new PoolCapacityReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolCapacityReqWithDefaults

`func NewPoolCapacityReqWithDefaults() *PoolCapacityReq`

NewPoolCapacityReqWithDefaults instantiates a new PoolCapacityReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodingChunkNum

`func (o *PoolCapacityReq) GetCodingChunkNum() int64`

GetCodingChunkNum returns the CodingChunkNum field if non-nil, zero value otherwise.

### GetCodingChunkNumOk

`func (o *PoolCapacityReq) GetCodingChunkNumOk() (*int64, bool)`

GetCodingChunkNumOk returns a tuple with the CodingChunkNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodingChunkNum

`func (o *PoolCapacityReq) SetCodingChunkNum(v int64)`

SetCodingChunkNum sets CodingChunkNum field to given value.

### HasCodingChunkNum

`func (o *PoolCapacityReq) HasCodingChunkNum() bool`

HasCodingChunkNum returns a boolean if a field has been set.

### GetDataChunkNum

`func (o *PoolCapacityReq) GetDataChunkNum() int64`

GetDataChunkNum returns the DataChunkNum field if non-nil, zero value otherwise.

### GetDataChunkNumOk

`func (o *PoolCapacityReq) GetDataChunkNumOk() (*int64, bool)`

GetDataChunkNumOk returns a tuple with the DataChunkNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChunkNum

`func (o *PoolCapacityReq) SetDataChunkNum(v int64)`

SetDataChunkNum sets DataChunkNum field to given value.

### HasDataChunkNum

`func (o *PoolCapacityReq) HasDataChunkNum() bool`

HasDataChunkNum returns a boolean if a field has been set.

### GetOsdIds

`func (o *PoolCapacityReq) GetOsdIds() []int64`

GetOsdIds returns the OsdIds field if non-nil, zero value otherwise.

### GetOsdIdsOk

`func (o *PoolCapacityReq) GetOsdIdsOk() (*[]int64, bool)`

GetOsdIdsOk returns a tuple with the OsdIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdIds

`func (o *PoolCapacityReq) SetOsdIds(v []int64)`

SetOsdIds sets OsdIds field to given value.

### HasOsdIds

`func (o *PoolCapacityReq) HasOsdIds() bool`

HasOsdIds returns a boolean if a field has been set.

### GetPoolType

`func (o *PoolCapacityReq) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *PoolCapacityReq) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *PoolCapacityReq) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *PoolCapacityReq) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetReplicateSize

`func (o *PoolCapacityReq) GetReplicateSize() int64`

GetReplicateSize returns the ReplicateSize field if non-nil, zero value otherwise.

### GetReplicateSizeOk

`func (o *PoolCapacityReq) GetReplicateSizeOk() (*int64, bool)`

GetReplicateSizeOk returns a tuple with the ReplicateSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateSize

`func (o *PoolCapacityReq) SetReplicateSize(v int64)`

SetReplicateSize sets ReplicateSize field to given value.

### HasReplicateSize

`func (o *PoolCapacityReq) HasReplicateSize() bool`

HasReplicateSize returns a boolean if a field has been set.

### GetThinProvisioned

`func (o *PoolCapacityReq) GetThinProvisioned() bool`

GetThinProvisioned returns the ThinProvisioned field if non-nil, zero value otherwise.

### GetThinProvisionedOk

`func (o *PoolCapacityReq) GetThinProvisionedOk() (*bool, bool)`

GetThinProvisionedOk returns a tuple with the ThinProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioned

`func (o *PoolCapacityReq) SetThinProvisioned(v bool)`

SetThinProvisioned sets ThinProvisioned field to given value.

### HasThinProvisioned

`func (o *PoolCapacityReq) HasThinProvisioned() bool`

HasThinProvisioned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


