# DfsRootfsUpdateActivePoolReqRootfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePoolIds** | Pointer to **[]int64** | all active pool ids, include not changed ids | [optional] 
**PoolPolicies** | Pointer to [**[]PoolPolicy**](PoolPolicy.md) | changed active pool policy | [optional] 

## Methods

### NewDfsRootfsUpdateActivePoolReqRootfs

`func NewDfsRootfsUpdateActivePoolReqRootfs() *DfsRootfsUpdateActivePoolReqRootfs`

NewDfsRootfsUpdateActivePoolReqRootfs instantiates a new DfsRootfsUpdateActivePoolReqRootfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsRootfsUpdateActivePoolReqRootfsWithDefaults

`func NewDfsRootfsUpdateActivePoolReqRootfsWithDefaults() *DfsRootfsUpdateActivePoolReqRootfs`

NewDfsRootfsUpdateActivePoolReqRootfsWithDefaults instantiates a new DfsRootfsUpdateActivePoolReqRootfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePoolIds

`func (o *DfsRootfsUpdateActivePoolReqRootfs) GetActivePoolIds() []int64`

GetActivePoolIds returns the ActivePoolIds field if non-nil, zero value otherwise.

### GetActivePoolIdsOk

`func (o *DfsRootfsUpdateActivePoolReqRootfs) GetActivePoolIdsOk() (*[]int64, bool)`

GetActivePoolIdsOk returns a tuple with the ActivePoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePoolIds

`func (o *DfsRootfsUpdateActivePoolReqRootfs) SetActivePoolIds(v []int64)`

SetActivePoolIds sets ActivePoolIds field to given value.

### HasActivePoolIds

`func (o *DfsRootfsUpdateActivePoolReqRootfs) HasActivePoolIds() bool`

HasActivePoolIds returns a boolean if a field has been set.

### GetPoolPolicies

`func (o *DfsRootfsUpdateActivePoolReqRootfs) GetPoolPolicies() []PoolPolicy`

GetPoolPolicies returns the PoolPolicies field if non-nil, zero value otherwise.

### GetPoolPoliciesOk

`func (o *DfsRootfsUpdateActivePoolReqRootfs) GetPoolPoliciesOk() (*[]PoolPolicy, bool)`

GetPoolPoliciesOk returns a tuple with the PoolPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPolicies

`func (o *DfsRootfsUpdateActivePoolReqRootfs) SetPoolPolicies(v []PoolPolicy)`

SetPoolPolicies sets PoolPolicies field to given value.

### HasPoolPolicies

`func (o *DfsRootfsUpdateActivePoolReqRootfs) HasPoolPolicies() bool`

HasPoolPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


