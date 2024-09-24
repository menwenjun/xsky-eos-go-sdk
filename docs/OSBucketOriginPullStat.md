# OSBucketOriginPullStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**ExpiredObjects** | Pointer to **int64** |  | [optional] 
**ExpiredSize** | Pointer to **int64** |  | [optional] 
**FailureOps** | Pointer to **int64** |  | [optional] 
**RxObjects** | Pointer to **int64** |  | [optional] 
**RxSize** | Pointer to **int64** |  | [optional] 
**StorageClasses** | Pointer to [**map[string]OSOriginPullStorageClassStat**](OSOriginPullStorageClassStat.md) |  | [optional] 
**SuccessOps** | Pointer to **int64** |  | [optional] 
**TotalOps** | Pointer to **int64** |  | [optional] 
**TotalRxObjects** | Pointer to **int64** |  | [optional] 
**TotalRxSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewOSBucketOriginPullStat

`func NewOSBucketOriginPullStat() *OSBucketOriginPullStat`

NewOSBucketOriginPullStat instantiates a new OSBucketOriginPullStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSBucketOriginPullStatWithDefaults

`func NewOSBucketOriginPullStatWithDefaults() *OSBucketOriginPullStat`

NewOSBucketOriginPullStatWithDefaults instantiates a new OSBucketOriginPullStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *OSBucketOriginPullStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSBucketOriginPullStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSBucketOriginPullStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSBucketOriginPullStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetExpiredObjects

`func (o *OSBucketOriginPullStat) GetExpiredObjects() int64`

GetExpiredObjects returns the ExpiredObjects field if non-nil, zero value otherwise.

### GetExpiredObjectsOk

`func (o *OSBucketOriginPullStat) GetExpiredObjectsOk() (*int64, bool)`

GetExpiredObjectsOk returns a tuple with the ExpiredObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredObjects

`func (o *OSBucketOriginPullStat) SetExpiredObjects(v int64)`

SetExpiredObjects sets ExpiredObjects field to given value.

### HasExpiredObjects

`func (o *OSBucketOriginPullStat) HasExpiredObjects() bool`

HasExpiredObjects returns a boolean if a field has been set.

### GetExpiredSize

`func (o *OSBucketOriginPullStat) GetExpiredSize() int64`

GetExpiredSize returns the ExpiredSize field if non-nil, zero value otherwise.

### GetExpiredSizeOk

`func (o *OSBucketOriginPullStat) GetExpiredSizeOk() (*int64, bool)`

GetExpiredSizeOk returns a tuple with the ExpiredSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredSize

`func (o *OSBucketOriginPullStat) SetExpiredSize(v int64)`

SetExpiredSize sets ExpiredSize field to given value.

### HasExpiredSize

`func (o *OSBucketOriginPullStat) HasExpiredSize() bool`

HasExpiredSize returns a boolean if a field has been set.

### GetFailureOps

`func (o *OSBucketOriginPullStat) GetFailureOps() int64`

GetFailureOps returns the FailureOps field if non-nil, zero value otherwise.

### GetFailureOpsOk

`func (o *OSBucketOriginPullStat) GetFailureOpsOk() (*int64, bool)`

GetFailureOpsOk returns a tuple with the FailureOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureOps

`func (o *OSBucketOriginPullStat) SetFailureOps(v int64)`

SetFailureOps sets FailureOps field to given value.

### HasFailureOps

`func (o *OSBucketOriginPullStat) HasFailureOps() bool`

HasFailureOps returns a boolean if a field has been set.

### GetRxObjects

`func (o *OSBucketOriginPullStat) GetRxObjects() int64`

GetRxObjects returns the RxObjects field if non-nil, zero value otherwise.

### GetRxObjectsOk

`func (o *OSBucketOriginPullStat) GetRxObjectsOk() (*int64, bool)`

GetRxObjectsOk returns a tuple with the RxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxObjects

`func (o *OSBucketOriginPullStat) SetRxObjects(v int64)`

SetRxObjects sets RxObjects field to given value.

### HasRxObjects

`func (o *OSBucketOriginPullStat) HasRxObjects() bool`

HasRxObjects returns a boolean if a field has been set.

### GetRxSize

`func (o *OSBucketOriginPullStat) GetRxSize() int64`

GetRxSize returns the RxSize field if non-nil, zero value otherwise.

### GetRxSizeOk

`func (o *OSBucketOriginPullStat) GetRxSizeOk() (*int64, bool)`

GetRxSizeOk returns a tuple with the RxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxSize

`func (o *OSBucketOriginPullStat) SetRxSize(v int64)`

SetRxSize sets RxSize field to given value.

### HasRxSize

`func (o *OSBucketOriginPullStat) HasRxSize() bool`

HasRxSize returns a boolean if a field has been set.

### GetStorageClasses

`func (o *OSBucketOriginPullStat) GetStorageClasses() map[string]OSOriginPullStorageClassStat`

GetStorageClasses returns the StorageClasses field if non-nil, zero value otherwise.

### GetStorageClassesOk

`func (o *OSBucketOriginPullStat) GetStorageClassesOk() (*map[string]OSOriginPullStorageClassStat, bool)`

GetStorageClassesOk returns a tuple with the StorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClasses

`func (o *OSBucketOriginPullStat) SetStorageClasses(v map[string]OSOriginPullStorageClassStat)`

SetStorageClasses sets StorageClasses field to given value.

### HasStorageClasses

`func (o *OSBucketOriginPullStat) HasStorageClasses() bool`

HasStorageClasses returns a boolean if a field has been set.

### GetSuccessOps

`func (o *OSBucketOriginPullStat) GetSuccessOps() int64`

GetSuccessOps returns the SuccessOps field if non-nil, zero value otherwise.

### GetSuccessOpsOk

`func (o *OSBucketOriginPullStat) GetSuccessOpsOk() (*int64, bool)`

GetSuccessOpsOk returns a tuple with the SuccessOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessOps

`func (o *OSBucketOriginPullStat) SetSuccessOps(v int64)`

SetSuccessOps sets SuccessOps field to given value.

### HasSuccessOps

`func (o *OSBucketOriginPullStat) HasSuccessOps() bool`

HasSuccessOps returns a boolean if a field has been set.

### GetTotalOps

`func (o *OSBucketOriginPullStat) GetTotalOps() int64`

GetTotalOps returns the TotalOps field if non-nil, zero value otherwise.

### GetTotalOpsOk

`func (o *OSBucketOriginPullStat) GetTotalOpsOk() (*int64, bool)`

GetTotalOpsOk returns a tuple with the TotalOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOps

`func (o *OSBucketOriginPullStat) SetTotalOps(v int64)`

SetTotalOps sets TotalOps field to given value.

### HasTotalOps

`func (o *OSBucketOriginPullStat) HasTotalOps() bool`

HasTotalOps returns a boolean if a field has been set.

### GetTotalRxObjects

`func (o *OSBucketOriginPullStat) GetTotalRxObjects() int64`

GetTotalRxObjects returns the TotalRxObjects field if non-nil, zero value otherwise.

### GetTotalRxObjectsOk

`func (o *OSBucketOriginPullStat) GetTotalRxObjectsOk() (*int64, bool)`

GetTotalRxObjectsOk returns a tuple with the TotalRxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRxObjects

`func (o *OSBucketOriginPullStat) SetTotalRxObjects(v int64)`

SetTotalRxObjects sets TotalRxObjects field to given value.

### HasTotalRxObjects

`func (o *OSBucketOriginPullStat) HasTotalRxObjects() bool`

HasTotalRxObjects returns a boolean if a field has been set.

### GetTotalRxSize

`func (o *OSBucketOriginPullStat) GetTotalRxSize() int64`

GetTotalRxSize returns the TotalRxSize field if non-nil, zero value otherwise.

### GetTotalRxSizeOk

`func (o *OSBucketOriginPullStat) GetTotalRxSizeOk() (*int64, bool)`

GetTotalRxSizeOk returns a tuple with the TotalRxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRxSize

`func (o *OSBucketOriginPullStat) SetTotalRxSize(v int64)`

SetTotalRxSize sets TotalRxSize field to given value.

### HasTotalRxSize

`func (o *OSBucketOriginPullStat) HasTotalRxSize() bool`

HasTotalRxSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


