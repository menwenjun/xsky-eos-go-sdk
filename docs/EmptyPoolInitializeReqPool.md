# EmptyPoolInitializeReqPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**EmptyPoolInitializeReqPoolCache**](EmptyPoolInitializeReqPoolCache.md) |  | [optional] 
**OsdIds** | Pointer to **[]int64** |  | [optional] 
**PrimaryPlacementNodeId** | Pointer to **int64** |  | [optional] 
**Ruleset** | Pointer to [**[]PoolRuleReq**](PoolRuleReq.md) |  | [optional] 
**SubFailureDomainType** | Pointer to **string** |  | [optional] 

## Methods

### NewEmptyPoolInitializeReqPool

`func NewEmptyPoolInitializeReqPool() *EmptyPoolInitializeReqPool`

NewEmptyPoolInitializeReqPool instantiates a new EmptyPoolInitializeReqPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmptyPoolInitializeReqPoolWithDefaults

`func NewEmptyPoolInitializeReqPoolWithDefaults() *EmptyPoolInitializeReqPool`

NewEmptyPoolInitializeReqPoolWithDefaults instantiates a new EmptyPoolInitializeReqPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *EmptyPoolInitializeReqPool) GetCache() EmptyPoolInitializeReqPoolCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *EmptyPoolInitializeReqPool) GetCacheOk() (*EmptyPoolInitializeReqPoolCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *EmptyPoolInitializeReqPool) SetCache(v EmptyPoolInitializeReqPoolCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *EmptyPoolInitializeReqPool) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetOsdIds

`func (o *EmptyPoolInitializeReqPool) GetOsdIds() []int64`

GetOsdIds returns the OsdIds field if non-nil, zero value otherwise.

### GetOsdIdsOk

`func (o *EmptyPoolInitializeReqPool) GetOsdIdsOk() (*[]int64, bool)`

GetOsdIdsOk returns a tuple with the OsdIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdIds

`func (o *EmptyPoolInitializeReqPool) SetOsdIds(v []int64)`

SetOsdIds sets OsdIds field to given value.

### HasOsdIds

`func (o *EmptyPoolInitializeReqPool) HasOsdIds() bool`

HasOsdIds returns a boolean if a field has been set.

### GetPrimaryPlacementNodeId

`func (o *EmptyPoolInitializeReqPool) GetPrimaryPlacementNodeId() int64`

GetPrimaryPlacementNodeId returns the PrimaryPlacementNodeId field if non-nil, zero value otherwise.

### GetPrimaryPlacementNodeIdOk

`func (o *EmptyPoolInitializeReqPool) GetPrimaryPlacementNodeIdOk() (*int64, bool)`

GetPrimaryPlacementNodeIdOk returns a tuple with the PrimaryPlacementNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPlacementNodeId

`func (o *EmptyPoolInitializeReqPool) SetPrimaryPlacementNodeId(v int64)`

SetPrimaryPlacementNodeId sets PrimaryPlacementNodeId field to given value.

### HasPrimaryPlacementNodeId

`func (o *EmptyPoolInitializeReqPool) HasPrimaryPlacementNodeId() bool`

HasPrimaryPlacementNodeId returns a boolean if a field has been set.

### GetRuleset

`func (o *EmptyPoolInitializeReqPool) GetRuleset() []PoolRuleReq`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *EmptyPoolInitializeReqPool) GetRulesetOk() (*[]PoolRuleReq, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *EmptyPoolInitializeReqPool) SetRuleset(v []PoolRuleReq)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *EmptyPoolInitializeReqPool) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetSubFailureDomainType

`func (o *EmptyPoolInitializeReqPool) GetSubFailureDomainType() string`

GetSubFailureDomainType returns the SubFailureDomainType field if non-nil, zero value otherwise.

### GetSubFailureDomainTypeOk

`func (o *EmptyPoolInitializeReqPool) GetSubFailureDomainTypeOk() (*string, bool)`

GetSubFailureDomainTypeOk returns a tuple with the SubFailureDomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFailureDomainType

`func (o *EmptyPoolInitializeReqPool) SetSubFailureDomainType(v string)`

SetSubFailureDomainType sets SubFailureDomainType field to given value.

### HasSubFailureDomainType

`func (o *EmptyPoolInitializeReqPool) HasSubFailureDomainType() bool`

HasSubFailureDomainType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


