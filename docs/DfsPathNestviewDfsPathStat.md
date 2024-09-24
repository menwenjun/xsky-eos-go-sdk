# DfsPathNestviewDfsPathStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **time.Time** |  | [optional] 
**Change** | Pointer to **time.Time** |  | [optional] 
**DpDfsSnapshotPolicies** | Pointer to **[]string** |  | [optional] 
**DpSnapshotNum** | Pointer to **int64** |  | [optional] 
**Files** | Pointer to **int64** |  | [optional] 
**LastSnapshotTime** | Pointer to **time.Time** |  | [optional] 
**Modify** | Pointer to **time.Time** |  | [optional] 
**QuotaNum** | Pointer to **int64** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**SnapshotNum** | Pointer to **int64** |  | [optional] 
**TotalSnapshotNum** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsPathNestviewDfsPathStat

`func NewDfsPathNestviewDfsPathStat() *DfsPathNestviewDfsPathStat`

NewDfsPathNestviewDfsPathStat instantiates a new DfsPathNestviewDfsPathStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsPathNestviewDfsPathStatWithDefaults

`func NewDfsPathNestviewDfsPathStatWithDefaults() *DfsPathNestviewDfsPathStat`

NewDfsPathNestviewDfsPathStatWithDefaults instantiates a new DfsPathNestviewDfsPathStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *DfsPathNestviewDfsPathStat) GetAccess() time.Time`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DfsPathNestviewDfsPathStat) GetAccessOk() (*time.Time, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DfsPathNestviewDfsPathStat) SetAccess(v time.Time)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *DfsPathNestviewDfsPathStat) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetChange

`func (o *DfsPathNestviewDfsPathStat) GetChange() time.Time`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *DfsPathNestviewDfsPathStat) GetChangeOk() (*time.Time, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *DfsPathNestviewDfsPathStat) SetChange(v time.Time)`

SetChange sets Change field to given value.

### HasChange

`func (o *DfsPathNestviewDfsPathStat) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetDpDfsSnapshotPolicies

`func (o *DfsPathNestviewDfsPathStat) GetDpDfsSnapshotPolicies() []string`

GetDpDfsSnapshotPolicies returns the DpDfsSnapshotPolicies field if non-nil, zero value otherwise.

### GetDpDfsSnapshotPoliciesOk

`func (o *DfsPathNestviewDfsPathStat) GetDpDfsSnapshotPoliciesOk() (*[]string, bool)`

GetDpDfsSnapshotPoliciesOk returns a tuple with the DpDfsSnapshotPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpDfsSnapshotPolicies

`func (o *DfsPathNestviewDfsPathStat) SetDpDfsSnapshotPolicies(v []string)`

SetDpDfsSnapshotPolicies sets DpDfsSnapshotPolicies field to given value.

### HasDpDfsSnapshotPolicies

`func (o *DfsPathNestviewDfsPathStat) HasDpDfsSnapshotPolicies() bool`

HasDpDfsSnapshotPolicies returns a boolean if a field has been set.

### GetDpSnapshotNum

`func (o *DfsPathNestviewDfsPathStat) GetDpSnapshotNum() int64`

GetDpSnapshotNum returns the DpSnapshotNum field if non-nil, zero value otherwise.

### GetDpSnapshotNumOk

`func (o *DfsPathNestviewDfsPathStat) GetDpSnapshotNumOk() (*int64, bool)`

GetDpSnapshotNumOk returns a tuple with the DpSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSnapshotNum

`func (o *DfsPathNestviewDfsPathStat) SetDpSnapshotNum(v int64)`

SetDpSnapshotNum sets DpSnapshotNum field to given value.

### HasDpSnapshotNum

`func (o *DfsPathNestviewDfsPathStat) HasDpSnapshotNum() bool`

HasDpSnapshotNum returns a boolean if a field has been set.

### GetFiles

`func (o *DfsPathNestviewDfsPathStat) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *DfsPathNestviewDfsPathStat) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *DfsPathNestviewDfsPathStat) SetFiles(v int64)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *DfsPathNestviewDfsPathStat) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLastSnapshotTime

`func (o *DfsPathNestviewDfsPathStat) GetLastSnapshotTime() time.Time`

GetLastSnapshotTime returns the LastSnapshotTime field if non-nil, zero value otherwise.

### GetLastSnapshotTimeOk

`func (o *DfsPathNestviewDfsPathStat) GetLastSnapshotTimeOk() (*time.Time, bool)`

GetLastSnapshotTimeOk returns a tuple with the LastSnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSnapshotTime

`func (o *DfsPathNestviewDfsPathStat) SetLastSnapshotTime(v time.Time)`

SetLastSnapshotTime sets LastSnapshotTime field to given value.

### HasLastSnapshotTime

`func (o *DfsPathNestviewDfsPathStat) HasLastSnapshotTime() bool`

HasLastSnapshotTime returns a boolean if a field has been set.

### GetModify

`func (o *DfsPathNestviewDfsPathStat) GetModify() time.Time`

GetModify returns the Modify field if non-nil, zero value otherwise.

### GetModifyOk

`func (o *DfsPathNestviewDfsPathStat) GetModifyOk() (*time.Time, bool)`

GetModifyOk returns a tuple with the Modify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModify

`func (o *DfsPathNestviewDfsPathStat) SetModify(v time.Time)`

SetModify sets Modify field to given value.

### HasModify

`func (o *DfsPathNestviewDfsPathStat) HasModify() bool`

HasModify returns a boolean if a field has been set.

### GetQuotaNum

`func (o *DfsPathNestviewDfsPathStat) GetQuotaNum() int64`

GetQuotaNum returns the QuotaNum field if non-nil, zero value otherwise.

### GetQuotaNumOk

`func (o *DfsPathNestviewDfsPathStat) GetQuotaNumOk() (*int64, bool)`

GetQuotaNumOk returns a tuple with the QuotaNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaNum

`func (o *DfsPathNestviewDfsPathStat) SetQuotaNum(v int64)`

SetQuotaNum sets QuotaNum field to given value.

### HasQuotaNum

`func (o *DfsPathNestviewDfsPathStat) HasQuotaNum() bool`

HasQuotaNum returns a boolean if a field has been set.

### GetShared

`func (o *DfsPathNestviewDfsPathStat) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *DfsPathNestviewDfsPathStat) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *DfsPathNestviewDfsPathStat) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *DfsPathNestviewDfsPathStat) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSize

`func (o *DfsPathNestviewDfsPathStat) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DfsPathNestviewDfsPathStat) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DfsPathNestviewDfsPathStat) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DfsPathNestviewDfsPathStat) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSnapshotNum

`func (o *DfsPathNestviewDfsPathStat) GetSnapshotNum() int64`

GetSnapshotNum returns the SnapshotNum field if non-nil, zero value otherwise.

### GetSnapshotNumOk

`func (o *DfsPathNestviewDfsPathStat) GetSnapshotNumOk() (*int64, bool)`

GetSnapshotNumOk returns a tuple with the SnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotNum

`func (o *DfsPathNestviewDfsPathStat) SetSnapshotNum(v int64)`

SetSnapshotNum sets SnapshotNum field to given value.

### HasSnapshotNum

`func (o *DfsPathNestviewDfsPathStat) HasSnapshotNum() bool`

HasSnapshotNum returns a boolean if a field has been set.

### GetTotalSnapshotNum

`func (o *DfsPathNestviewDfsPathStat) GetTotalSnapshotNum() int64`

GetTotalSnapshotNum returns the TotalSnapshotNum field if non-nil, zero value otherwise.

### GetTotalSnapshotNumOk

`func (o *DfsPathNestviewDfsPathStat) GetTotalSnapshotNumOk() (*int64, bool)`

GetTotalSnapshotNumOk returns a tuple with the TotalSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSnapshotNum

`func (o *DfsPathNestviewDfsPathStat) SetTotalSnapshotNum(v int64)`

SetTotalSnapshotNum sets TotalSnapshotNum field to given value.

### HasTotalSnapshotNum

`func (o *DfsPathNestviewDfsPathStat) HasTotalSnapshotNum() bool`

HasTotalSnapshotNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


