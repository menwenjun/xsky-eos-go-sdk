# DfsTierAddPoolReqTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePoolIds** | Pointer to **[]int64** | active pool ids when add | [optional] 
**PoolIds** | Pointer to **[]int64** | pool ids to add/remove | [optional] 
**PoolPolicies** | Pointer to [**[]PoolPolicy**](PoolPolicy.md) | active pool policy to add | [optional] 

## Methods

### NewDfsTierAddPoolReqTier

`func NewDfsTierAddPoolReqTier() *DfsTierAddPoolReqTier`

NewDfsTierAddPoolReqTier instantiates a new DfsTierAddPoolReqTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsTierAddPoolReqTierWithDefaults

`func NewDfsTierAddPoolReqTierWithDefaults() *DfsTierAddPoolReqTier`

NewDfsTierAddPoolReqTierWithDefaults instantiates a new DfsTierAddPoolReqTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePoolIds

`func (o *DfsTierAddPoolReqTier) GetActivePoolIds() []int64`

GetActivePoolIds returns the ActivePoolIds field if non-nil, zero value otherwise.

### GetActivePoolIdsOk

`func (o *DfsTierAddPoolReqTier) GetActivePoolIdsOk() (*[]int64, bool)`

GetActivePoolIdsOk returns a tuple with the ActivePoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePoolIds

`func (o *DfsTierAddPoolReqTier) SetActivePoolIds(v []int64)`

SetActivePoolIds sets ActivePoolIds field to given value.

### HasActivePoolIds

`func (o *DfsTierAddPoolReqTier) HasActivePoolIds() bool`

HasActivePoolIds returns a boolean if a field has been set.

### GetPoolIds

`func (o *DfsTierAddPoolReqTier) GetPoolIds() []int64`

GetPoolIds returns the PoolIds field if non-nil, zero value otherwise.

### GetPoolIdsOk

`func (o *DfsTierAddPoolReqTier) GetPoolIdsOk() (*[]int64, bool)`

GetPoolIdsOk returns a tuple with the PoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIds

`func (o *DfsTierAddPoolReqTier) SetPoolIds(v []int64)`

SetPoolIds sets PoolIds field to given value.

### HasPoolIds

`func (o *DfsTierAddPoolReqTier) HasPoolIds() bool`

HasPoolIds returns a boolean if a field has been set.

### GetPoolPolicies

`func (o *DfsTierAddPoolReqTier) GetPoolPolicies() []PoolPolicy`

GetPoolPolicies returns the PoolPolicies field if non-nil, zero value otherwise.

### GetPoolPoliciesOk

`func (o *DfsTierAddPoolReqTier) GetPoolPoliciesOk() (*[]PoolPolicy, bool)`

GetPoolPoliciesOk returns a tuple with the PoolPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPolicies

`func (o *DfsTierAddPoolReqTier) SetPoolPolicies(v []PoolPolicy)`

SetPoolPolicies sets PoolPolicies field to given value.

### HasPoolPolicies

`func (o *DfsTierAddPoolReqTier) HasPoolPolicies() bool`

HasPoolPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


