# PolicyLinkPathInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsPathId** | Pointer to **int64** | id of path | [optional] 
**HdfsNum** | Pointer to **int64** | hdfs num of path | [optional] 
**IsBucket** | Pointer to **bool** | if bucket of path | [optional] 
**Path** | Pointer to **string** | name of path | [optional] 
**Shares** | Pointer to **[]string** | share type of path | [optional] 
**StoragePolicyIds** | Pointer to **[]int64** | link storage policy ids | [optional] 

## Methods

### NewPolicyLinkPathInfo

`func NewPolicyLinkPathInfo() *PolicyLinkPathInfo`

NewPolicyLinkPathInfo instantiates a new PolicyLinkPathInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyLinkPathInfoWithDefaults

`func NewPolicyLinkPathInfoWithDefaults() *PolicyLinkPathInfo`

NewPolicyLinkPathInfoWithDefaults instantiates a new PolicyLinkPathInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsPathId

`func (o *PolicyLinkPathInfo) GetDfsPathId() int64`

GetDfsPathId returns the DfsPathId field if non-nil, zero value otherwise.

### GetDfsPathIdOk

`func (o *PolicyLinkPathInfo) GetDfsPathIdOk() (*int64, bool)`

GetDfsPathIdOk returns a tuple with the DfsPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPathId

`func (o *PolicyLinkPathInfo) SetDfsPathId(v int64)`

SetDfsPathId sets DfsPathId field to given value.

### HasDfsPathId

`func (o *PolicyLinkPathInfo) HasDfsPathId() bool`

HasDfsPathId returns a boolean if a field has been set.

### GetHdfsNum

`func (o *PolicyLinkPathInfo) GetHdfsNum() int64`

GetHdfsNum returns the HdfsNum field if non-nil, zero value otherwise.

### GetHdfsNumOk

`func (o *PolicyLinkPathInfo) GetHdfsNumOk() (*int64, bool)`

GetHdfsNumOk returns a tuple with the HdfsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsNum

`func (o *PolicyLinkPathInfo) SetHdfsNum(v int64)`

SetHdfsNum sets HdfsNum field to given value.

### HasHdfsNum

`func (o *PolicyLinkPathInfo) HasHdfsNum() bool`

HasHdfsNum returns a boolean if a field has been set.

### GetIsBucket

`func (o *PolicyLinkPathInfo) GetIsBucket() bool`

GetIsBucket returns the IsBucket field if non-nil, zero value otherwise.

### GetIsBucketOk

`func (o *PolicyLinkPathInfo) GetIsBucketOk() (*bool, bool)`

GetIsBucketOk returns a tuple with the IsBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBucket

`func (o *PolicyLinkPathInfo) SetIsBucket(v bool)`

SetIsBucket sets IsBucket field to given value.

### HasIsBucket

`func (o *PolicyLinkPathInfo) HasIsBucket() bool`

HasIsBucket returns a boolean if a field has been set.

### GetPath

`func (o *PolicyLinkPathInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PolicyLinkPathInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PolicyLinkPathInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PolicyLinkPathInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetShares

`func (o *PolicyLinkPathInfo) GetShares() []string`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *PolicyLinkPathInfo) GetSharesOk() (*[]string, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *PolicyLinkPathInfo) SetShares(v []string)`

SetShares sets Shares field to given value.

### HasShares

`func (o *PolicyLinkPathInfo) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetStoragePolicyIds

`func (o *PolicyLinkPathInfo) GetStoragePolicyIds() []int64`

GetStoragePolicyIds returns the StoragePolicyIds field if non-nil, zero value otherwise.

### GetStoragePolicyIdsOk

`func (o *PolicyLinkPathInfo) GetStoragePolicyIdsOk() (*[]int64, bool)`

GetStoragePolicyIdsOk returns a tuple with the StoragePolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyIds

`func (o *PolicyLinkPathInfo) SetStoragePolicyIds(v []int64)`

SetStoragePolicyIds sets StoragePolicyIds field to given value.

### HasStoragePolicyIds

`func (o *PolicyLinkPathInfo) HasStoragePolicyIds() bool`

HasStoragePolicyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


