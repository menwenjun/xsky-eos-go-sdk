# DfsDirectoryValidationRespDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpDfsSnapshotPolicyNum** | **int64** | count of data protection snapshot policy on the directory | 
**Existed** | **bool** | directory existed | 
**ParentChildBucketExisted** | **bool** | bucket existed on the parent or child directory | 
**ParentChildSnapshotExisted** | **bool** | snapshot existed on the parent or child directory | 
**QuotaExisted** | **bool** | quota existed on the directory | 
**RootDirSnapNum** | **int64** | count of snapshot on the root directory | 
**ShareExisted** | **bool** | share existed on the directory | 
**SnapshotExisted** | **bool** | snapshot existed on the directory | 

## Methods

### NewDfsDirectoryValidationRespDirectory

`func NewDfsDirectoryValidationRespDirectory(dpDfsSnapshotPolicyNum int64, existed bool, parentChildBucketExisted bool, parentChildSnapshotExisted bool, quotaExisted bool, rootDirSnapNum int64, shareExisted bool, snapshotExisted bool, ) *DfsDirectoryValidationRespDirectory`

NewDfsDirectoryValidationRespDirectory instantiates a new DfsDirectoryValidationRespDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsDirectoryValidationRespDirectoryWithDefaults

`func NewDfsDirectoryValidationRespDirectoryWithDefaults() *DfsDirectoryValidationRespDirectory`

NewDfsDirectoryValidationRespDirectoryWithDefaults instantiates a new DfsDirectoryValidationRespDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpDfsSnapshotPolicyNum

`func (o *DfsDirectoryValidationRespDirectory) GetDpDfsSnapshotPolicyNum() int64`

GetDpDfsSnapshotPolicyNum returns the DpDfsSnapshotPolicyNum field if non-nil, zero value otherwise.

### GetDpDfsSnapshotPolicyNumOk

`func (o *DfsDirectoryValidationRespDirectory) GetDpDfsSnapshotPolicyNumOk() (*int64, bool)`

GetDpDfsSnapshotPolicyNumOk returns a tuple with the DpDfsSnapshotPolicyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpDfsSnapshotPolicyNum

`func (o *DfsDirectoryValidationRespDirectory) SetDpDfsSnapshotPolicyNum(v int64)`

SetDpDfsSnapshotPolicyNum sets DpDfsSnapshotPolicyNum field to given value.


### GetExisted

`func (o *DfsDirectoryValidationRespDirectory) GetExisted() bool`

GetExisted returns the Existed field if non-nil, zero value otherwise.

### GetExistedOk

`func (o *DfsDirectoryValidationRespDirectory) GetExistedOk() (*bool, bool)`

GetExistedOk returns a tuple with the Existed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExisted

`func (o *DfsDirectoryValidationRespDirectory) SetExisted(v bool)`

SetExisted sets Existed field to given value.


### GetParentChildBucketExisted

`func (o *DfsDirectoryValidationRespDirectory) GetParentChildBucketExisted() bool`

GetParentChildBucketExisted returns the ParentChildBucketExisted field if non-nil, zero value otherwise.

### GetParentChildBucketExistedOk

`func (o *DfsDirectoryValidationRespDirectory) GetParentChildBucketExistedOk() (*bool, bool)`

GetParentChildBucketExistedOk returns a tuple with the ParentChildBucketExisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChildBucketExisted

`func (o *DfsDirectoryValidationRespDirectory) SetParentChildBucketExisted(v bool)`

SetParentChildBucketExisted sets ParentChildBucketExisted field to given value.


### GetParentChildSnapshotExisted

`func (o *DfsDirectoryValidationRespDirectory) GetParentChildSnapshotExisted() bool`

GetParentChildSnapshotExisted returns the ParentChildSnapshotExisted field if non-nil, zero value otherwise.

### GetParentChildSnapshotExistedOk

`func (o *DfsDirectoryValidationRespDirectory) GetParentChildSnapshotExistedOk() (*bool, bool)`

GetParentChildSnapshotExistedOk returns a tuple with the ParentChildSnapshotExisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChildSnapshotExisted

`func (o *DfsDirectoryValidationRespDirectory) SetParentChildSnapshotExisted(v bool)`

SetParentChildSnapshotExisted sets ParentChildSnapshotExisted field to given value.


### GetQuotaExisted

`func (o *DfsDirectoryValidationRespDirectory) GetQuotaExisted() bool`

GetQuotaExisted returns the QuotaExisted field if non-nil, zero value otherwise.

### GetQuotaExistedOk

`func (o *DfsDirectoryValidationRespDirectory) GetQuotaExistedOk() (*bool, bool)`

GetQuotaExistedOk returns a tuple with the QuotaExisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaExisted

`func (o *DfsDirectoryValidationRespDirectory) SetQuotaExisted(v bool)`

SetQuotaExisted sets QuotaExisted field to given value.


### GetRootDirSnapNum

`func (o *DfsDirectoryValidationRespDirectory) GetRootDirSnapNum() int64`

GetRootDirSnapNum returns the RootDirSnapNum field if non-nil, zero value otherwise.

### GetRootDirSnapNumOk

`func (o *DfsDirectoryValidationRespDirectory) GetRootDirSnapNumOk() (*int64, bool)`

GetRootDirSnapNumOk returns a tuple with the RootDirSnapNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDirSnapNum

`func (o *DfsDirectoryValidationRespDirectory) SetRootDirSnapNum(v int64)`

SetRootDirSnapNum sets RootDirSnapNum field to given value.


### GetShareExisted

`func (o *DfsDirectoryValidationRespDirectory) GetShareExisted() bool`

GetShareExisted returns the ShareExisted field if non-nil, zero value otherwise.

### GetShareExistedOk

`func (o *DfsDirectoryValidationRespDirectory) GetShareExistedOk() (*bool, bool)`

GetShareExistedOk returns a tuple with the ShareExisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareExisted

`func (o *DfsDirectoryValidationRespDirectory) SetShareExisted(v bool)`

SetShareExisted sets ShareExisted field to given value.


### GetSnapshotExisted

`func (o *DfsDirectoryValidationRespDirectory) GetSnapshotExisted() bool`

GetSnapshotExisted returns the SnapshotExisted field if non-nil, zero value otherwise.

### GetSnapshotExistedOk

`func (o *DfsDirectoryValidationRespDirectory) GetSnapshotExistedOk() (*bool, bool)`

GetSnapshotExistedOk returns a tuple with the SnapshotExisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExisted

`func (o *DfsDirectoryValidationRespDirectory) SetSnapshotExisted(v bool)`

SetSnapshotExisted sets SnapshotExisted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


