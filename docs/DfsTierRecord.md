# DfsTierRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**ActivePoolIds** | Pointer to **[]int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**DfsStoragePolicies** | Pointer to [**[]DfsStoragePolicy**](DfsStoragePolicy.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PolicyNum** | Pointer to **int64** |  | [optional] 
**PoolIds** | Pointer to **[]int64** |  | [optional] 
**PoolNames** | Pointer to **[]string** |  | [optional] 
**PoolNum** | Pointer to **int64** |  | [optional] 
**PoolPolicies** | Pointer to [**[]DfsTierPoolPolicy**](DfsTierPoolPolicy.md) |  | [optional] 
**Pools** | Pointer to [**[]PoolNestview**](PoolNestview.md) |  | [optional] 
**ScId** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StoragePolicyIds** | Pointer to **[]int64** |  | [optional] 
**Stretched** | Pointer to **bool** | indicate that all pools under this class is stretched pool | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**WritePolicy** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to [**[]DfsTierStat**](DfsTierStat.md) |  | [optional] 

## Methods

### NewDfsTierRecord

`func NewDfsTierRecord() *DfsTierRecord`

NewDfsTierRecord instantiates a new DfsTierRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsTierRecordWithDefaults

`func NewDfsTierRecordWithDefaults() *DfsTierRecord`

NewDfsTierRecordWithDefaults instantiates a new DfsTierRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsTierRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsTierRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsTierRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsTierRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetActivePoolIds

`func (o *DfsTierRecord) GetActivePoolIds() []int64`

GetActivePoolIds returns the ActivePoolIds field if non-nil, zero value otherwise.

### GetActivePoolIdsOk

`func (o *DfsTierRecord) GetActivePoolIdsOk() (*[]int64, bool)`

GetActivePoolIdsOk returns a tuple with the ActivePoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePoolIds

`func (o *DfsTierRecord) SetActivePoolIds(v []int64)`

SetActivePoolIds sets ActivePoolIds field to given value.

### HasActivePoolIds

`func (o *DfsTierRecord) HasActivePoolIds() bool`

HasActivePoolIds returns a boolean if a field has been set.

### GetCluster

`func (o *DfsTierRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsTierRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsTierRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsTierRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsTierRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsTierRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsTierRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsTierRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsTierRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsTierRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsTierRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsTierRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsTierRecord) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsTierRecord) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsTierRecord) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsTierRecord) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetDfsStoragePolicies

`func (o *DfsTierRecord) GetDfsStoragePolicies() []DfsStoragePolicy`

GetDfsStoragePolicies returns the DfsStoragePolicies field if non-nil, zero value otherwise.

### GetDfsStoragePoliciesOk

`func (o *DfsTierRecord) GetDfsStoragePoliciesOk() (*[]DfsStoragePolicy, bool)`

GetDfsStoragePoliciesOk returns a tuple with the DfsStoragePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsStoragePolicies

`func (o *DfsTierRecord) SetDfsStoragePolicies(v []DfsStoragePolicy)`

SetDfsStoragePolicies sets DfsStoragePolicies field to given value.

### HasDfsStoragePolicies

`func (o *DfsTierRecord) HasDfsStoragePolicies() bool`

HasDfsStoragePolicies returns a boolean if a field has been set.

### GetId

`func (o *DfsTierRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsTierRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsTierRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsTierRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *DfsTierRecord) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *DfsTierRecord) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *DfsTierRecord) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *DfsTierRecord) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *DfsTierRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsTierRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsTierRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsTierRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyNum

`func (o *DfsTierRecord) GetPolicyNum() int64`

GetPolicyNum returns the PolicyNum field if non-nil, zero value otherwise.

### GetPolicyNumOk

`func (o *DfsTierRecord) GetPolicyNumOk() (*int64, bool)`

GetPolicyNumOk returns a tuple with the PolicyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNum

`func (o *DfsTierRecord) SetPolicyNum(v int64)`

SetPolicyNum sets PolicyNum field to given value.

### HasPolicyNum

`func (o *DfsTierRecord) HasPolicyNum() bool`

HasPolicyNum returns a boolean if a field has been set.

### GetPoolIds

`func (o *DfsTierRecord) GetPoolIds() []int64`

GetPoolIds returns the PoolIds field if non-nil, zero value otherwise.

### GetPoolIdsOk

`func (o *DfsTierRecord) GetPoolIdsOk() (*[]int64, bool)`

GetPoolIdsOk returns a tuple with the PoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIds

`func (o *DfsTierRecord) SetPoolIds(v []int64)`

SetPoolIds sets PoolIds field to given value.

### HasPoolIds

`func (o *DfsTierRecord) HasPoolIds() bool`

HasPoolIds returns a boolean if a field has been set.

### GetPoolNames

`func (o *DfsTierRecord) GetPoolNames() []string`

GetPoolNames returns the PoolNames field if non-nil, zero value otherwise.

### GetPoolNamesOk

`func (o *DfsTierRecord) GetPoolNamesOk() (*[]string, bool)`

GetPoolNamesOk returns a tuple with the PoolNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolNames

`func (o *DfsTierRecord) SetPoolNames(v []string)`

SetPoolNames sets PoolNames field to given value.

### HasPoolNames

`func (o *DfsTierRecord) HasPoolNames() bool`

HasPoolNames returns a boolean if a field has been set.

### GetPoolNum

`func (o *DfsTierRecord) GetPoolNum() int64`

GetPoolNum returns the PoolNum field if non-nil, zero value otherwise.

### GetPoolNumOk

`func (o *DfsTierRecord) GetPoolNumOk() (*int64, bool)`

GetPoolNumOk returns a tuple with the PoolNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolNum

`func (o *DfsTierRecord) SetPoolNum(v int64)`

SetPoolNum sets PoolNum field to given value.

### HasPoolNum

`func (o *DfsTierRecord) HasPoolNum() bool`

HasPoolNum returns a boolean if a field has been set.

### GetPoolPolicies

`func (o *DfsTierRecord) GetPoolPolicies() []DfsTierPoolPolicy`

GetPoolPolicies returns the PoolPolicies field if non-nil, zero value otherwise.

### GetPoolPoliciesOk

`func (o *DfsTierRecord) GetPoolPoliciesOk() (*[]DfsTierPoolPolicy, bool)`

GetPoolPoliciesOk returns a tuple with the PoolPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPolicies

`func (o *DfsTierRecord) SetPoolPolicies(v []DfsTierPoolPolicy)`

SetPoolPolicies sets PoolPolicies field to given value.

### HasPoolPolicies

`func (o *DfsTierRecord) HasPoolPolicies() bool`

HasPoolPolicies returns a boolean if a field has been set.

### GetPools

`func (o *DfsTierRecord) GetPools() []PoolNestview`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *DfsTierRecord) GetPoolsOk() (*[]PoolNestview, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *DfsTierRecord) SetPools(v []PoolNestview)`

SetPools sets Pools field to given value.

### HasPools

`func (o *DfsTierRecord) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetScId

`func (o *DfsTierRecord) GetScId() int64`

GetScId returns the ScId field if non-nil, zero value otherwise.

### GetScIdOk

`func (o *DfsTierRecord) GetScIdOk() (*int64, bool)`

GetScIdOk returns a tuple with the ScId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScId

`func (o *DfsTierRecord) SetScId(v int64)`

SetScId sets ScId field to given value.

### HasScId

`func (o *DfsTierRecord) HasScId() bool`

HasScId returns a boolean if a field has been set.

### GetStatus

`func (o *DfsTierRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsTierRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsTierRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsTierRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStoragePolicyIds

`func (o *DfsTierRecord) GetStoragePolicyIds() []int64`

GetStoragePolicyIds returns the StoragePolicyIds field if non-nil, zero value otherwise.

### GetStoragePolicyIdsOk

`func (o *DfsTierRecord) GetStoragePolicyIdsOk() (*[]int64, bool)`

GetStoragePolicyIdsOk returns a tuple with the StoragePolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyIds

`func (o *DfsTierRecord) SetStoragePolicyIds(v []int64)`

SetStoragePolicyIds sets StoragePolicyIds field to given value.

### HasStoragePolicyIds

`func (o *DfsTierRecord) HasStoragePolicyIds() bool`

HasStoragePolicyIds returns a boolean if a field has been set.

### GetStretched

`func (o *DfsTierRecord) GetStretched() bool`

GetStretched returns the Stretched field if non-nil, zero value otherwise.

### GetStretchedOk

`func (o *DfsTierRecord) GetStretchedOk() (*bool, bool)`

GetStretchedOk returns a tuple with the Stretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretched

`func (o *DfsTierRecord) SetStretched(v bool)`

SetStretched sets Stretched field to given value.

### HasStretched

`func (o *DfsTierRecord) HasStretched() bool`

HasStretched returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsTierRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsTierRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsTierRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsTierRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetWritePolicy

`func (o *DfsTierRecord) GetWritePolicy() string`

GetWritePolicy returns the WritePolicy field if non-nil, zero value otherwise.

### GetWritePolicyOk

`func (o *DfsTierRecord) GetWritePolicyOk() (*string, bool)`

GetWritePolicyOk returns a tuple with the WritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicy

`func (o *DfsTierRecord) SetWritePolicy(v string)`

SetWritePolicy sets WritePolicy field to given value.

### HasWritePolicy

`func (o *DfsTierRecord) HasWritePolicy() bool`

HasWritePolicy returns a boolean if a field has been set.

### GetSamples

`func (o *DfsTierRecord) GetSamples() []DfsTierStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DfsTierRecord) GetSamplesOk() (*[]DfsTierStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DfsTierRecord) SetSamples(v []DfsTierStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DfsTierRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


