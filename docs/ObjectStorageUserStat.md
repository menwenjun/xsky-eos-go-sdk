# ObjectStorageUserStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedObjects** | Pointer to **int64** |  | [optional] 
**AllocatedSize** | Pointer to **int64** |  | [optional] 
**CacheAllocatedObjects** | Pointer to **int64** |  | [optional] 
**CacheAllocatedSize** | Pointer to **int64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DelOpsPm** | Pointer to **float64** |  | [optional] 
**ExternalAllocatedObjects** | Pointer to **int64** |  | [optional] 
**ExternalAllocatedSize** | Pointer to **int64** |  | [optional] 
**LocalAllocatedObjects** | Pointer to **int64** |  | [optional] 
**LocalAllocatedSize** | Pointer to **int64** |  | [optional] 
**Policies** | Pointer to [**map[string]ObjectStoragePolicyStat**](ObjectStoragePolicyStat.md) |  | [optional] 
**RxBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**RxOpsPm** | Pointer to **float64** |  | [optional] 
**TotalDelOps** | Pointer to **int64** |  | [optional] 
**TotalDelSuccessOps** | Pointer to **int64** |  | [optional] 
**TotalRxBytes** | Pointer to **int64** |  | [optional] 
**TotalRxOps** | Pointer to **int64** |  | [optional] 
**TotalRxSuccessOps** | Pointer to **int64** |  | [optional] 
**TotalTxBytes** | Pointer to **int64** |  | [optional] 
**TotalTxOps** | Pointer to **int64** |  | [optional] 
**TotalTxSuccessOps** | Pointer to **int64** |  | [optional] 
**TxBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**TxOpsPm** | Pointer to **float64** |  | [optional] 

## Methods

### NewObjectStorageUserStat

`func NewObjectStorageUserStat() *ObjectStorageUserStat`

NewObjectStorageUserStat instantiates a new ObjectStorageUserStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageUserStatWithDefaults

`func NewObjectStorageUserStatWithDefaults() *ObjectStorageUserStat`

NewObjectStorageUserStatWithDefaults instantiates a new ObjectStorageUserStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedObjects

`func (o *ObjectStorageUserStat) GetAllocatedObjects() int64`

GetAllocatedObjects returns the AllocatedObjects field if non-nil, zero value otherwise.

### GetAllocatedObjectsOk

`func (o *ObjectStorageUserStat) GetAllocatedObjectsOk() (*int64, bool)`

GetAllocatedObjectsOk returns a tuple with the AllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedObjects

`func (o *ObjectStorageUserStat) SetAllocatedObjects(v int64)`

SetAllocatedObjects sets AllocatedObjects field to given value.

### HasAllocatedObjects

`func (o *ObjectStorageUserStat) HasAllocatedObjects() bool`

HasAllocatedObjects returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *ObjectStorageUserStat) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *ObjectStorageUserStat) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *ObjectStorageUserStat) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *ObjectStorageUserStat) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetCacheAllocatedObjects

`func (o *ObjectStorageUserStat) GetCacheAllocatedObjects() int64`

GetCacheAllocatedObjects returns the CacheAllocatedObjects field if non-nil, zero value otherwise.

### GetCacheAllocatedObjectsOk

`func (o *ObjectStorageUserStat) GetCacheAllocatedObjectsOk() (*int64, bool)`

GetCacheAllocatedObjectsOk returns a tuple with the CacheAllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllocatedObjects

`func (o *ObjectStorageUserStat) SetCacheAllocatedObjects(v int64)`

SetCacheAllocatedObjects sets CacheAllocatedObjects field to given value.

### HasCacheAllocatedObjects

`func (o *ObjectStorageUserStat) HasCacheAllocatedObjects() bool`

HasCacheAllocatedObjects returns a boolean if a field has been set.

### GetCacheAllocatedSize

`func (o *ObjectStorageUserStat) GetCacheAllocatedSize() int64`

GetCacheAllocatedSize returns the CacheAllocatedSize field if non-nil, zero value otherwise.

### GetCacheAllocatedSizeOk

`func (o *ObjectStorageUserStat) GetCacheAllocatedSizeOk() (*int64, bool)`

GetCacheAllocatedSizeOk returns a tuple with the CacheAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllocatedSize

`func (o *ObjectStorageUserStat) SetCacheAllocatedSize(v int64)`

SetCacheAllocatedSize sets CacheAllocatedSize field to given value.

### HasCacheAllocatedSize

`func (o *ObjectStorageUserStat) HasCacheAllocatedSize() bool`

HasCacheAllocatedSize returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageUserStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageUserStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageUserStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageUserStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDelOpsPm

`func (o *ObjectStorageUserStat) GetDelOpsPm() float64`

GetDelOpsPm returns the DelOpsPm field if non-nil, zero value otherwise.

### GetDelOpsPmOk

`func (o *ObjectStorageUserStat) GetDelOpsPmOk() (*float64, bool)`

GetDelOpsPmOk returns a tuple with the DelOpsPm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelOpsPm

`func (o *ObjectStorageUserStat) SetDelOpsPm(v float64)`

SetDelOpsPm sets DelOpsPm field to given value.

### HasDelOpsPm

`func (o *ObjectStorageUserStat) HasDelOpsPm() bool`

HasDelOpsPm returns a boolean if a field has been set.

### GetExternalAllocatedObjects

`func (o *ObjectStorageUserStat) GetExternalAllocatedObjects() int64`

GetExternalAllocatedObjects returns the ExternalAllocatedObjects field if non-nil, zero value otherwise.

### GetExternalAllocatedObjectsOk

`func (o *ObjectStorageUserStat) GetExternalAllocatedObjectsOk() (*int64, bool)`

GetExternalAllocatedObjectsOk returns a tuple with the ExternalAllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAllocatedObjects

`func (o *ObjectStorageUserStat) SetExternalAllocatedObjects(v int64)`

SetExternalAllocatedObjects sets ExternalAllocatedObjects field to given value.

### HasExternalAllocatedObjects

`func (o *ObjectStorageUserStat) HasExternalAllocatedObjects() bool`

HasExternalAllocatedObjects returns a boolean if a field has been set.

### GetExternalAllocatedSize

`func (o *ObjectStorageUserStat) GetExternalAllocatedSize() int64`

GetExternalAllocatedSize returns the ExternalAllocatedSize field if non-nil, zero value otherwise.

### GetExternalAllocatedSizeOk

`func (o *ObjectStorageUserStat) GetExternalAllocatedSizeOk() (*int64, bool)`

GetExternalAllocatedSizeOk returns a tuple with the ExternalAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAllocatedSize

`func (o *ObjectStorageUserStat) SetExternalAllocatedSize(v int64)`

SetExternalAllocatedSize sets ExternalAllocatedSize field to given value.

### HasExternalAllocatedSize

`func (o *ObjectStorageUserStat) HasExternalAllocatedSize() bool`

HasExternalAllocatedSize returns a boolean if a field has been set.

### GetLocalAllocatedObjects

`func (o *ObjectStorageUserStat) GetLocalAllocatedObjects() int64`

GetLocalAllocatedObjects returns the LocalAllocatedObjects field if non-nil, zero value otherwise.

### GetLocalAllocatedObjectsOk

`func (o *ObjectStorageUserStat) GetLocalAllocatedObjectsOk() (*int64, bool)`

GetLocalAllocatedObjectsOk returns a tuple with the LocalAllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAllocatedObjects

`func (o *ObjectStorageUserStat) SetLocalAllocatedObjects(v int64)`

SetLocalAllocatedObjects sets LocalAllocatedObjects field to given value.

### HasLocalAllocatedObjects

`func (o *ObjectStorageUserStat) HasLocalAllocatedObjects() bool`

HasLocalAllocatedObjects returns a boolean if a field has been set.

### GetLocalAllocatedSize

`func (o *ObjectStorageUserStat) GetLocalAllocatedSize() int64`

GetLocalAllocatedSize returns the LocalAllocatedSize field if non-nil, zero value otherwise.

### GetLocalAllocatedSizeOk

`func (o *ObjectStorageUserStat) GetLocalAllocatedSizeOk() (*int64, bool)`

GetLocalAllocatedSizeOk returns a tuple with the LocalAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAllocatedSize

`func (o *ObjectStorageUserStat) SetLocalAllocatedSize(v int64)`

SetLocalAllocatedSize sets LocalAllocatedSize field to given value.

### HasLocalAllocatedSize

`func (o *ObjectStorageUserStat) HasLocalAllocatedSize() bool`

HasLocalAllocatedSize returns a boolean if a field has been set.

### GetPolicies

`func (o *ObjectStorageUserStat) GetPolicies() map[string]ObjectStoragePolicyStat`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ObjectStorageUserStat) GetPoliciesOk() (*map[string]ObjectStoragePolicyStat, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ObjectStorageUserStat) SetPolicies(v map[string]ObjectStoragePolicyStat)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ObjectStorageUserStat) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetRxBandwidthKbyte

`func (o *ObjectStorageUserStat) GetRxBandwidthKbyte() float64`

GetRxBandwidthKbyte returns the RxBandwidthKbyte field if non-nil, zero value otherwise.

### GetRxBandwidthKbyteOk

`func (o *ObjectStorageUserStat) GetRxBandwidthKbyteOk() (*float64, bool)`

GetRxBandwidthKbyteOk returns a tuple with the RxBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBandwidthKbyte

`func (o *ObjectStorageUserStat) SetRxBandwidthKbyte(v float64)`

SetRxBandwidthKbyte sets RxBandwidthKbyte field to given value.

### HasRxBandwidthKbyte

`func (o *ObjectStorageUserStat) HasRxBandwidthKbyte() bool`

HasRxBandwidthKbyte returns a boolean if a field has been set.

### GetRxOpsPm

`func (o *ObjectStorageUserStat) GetRxOpsPm() float64`

GetRxOpsPm returns the RxOpsPm field if non-nil, zero value otherwise.

### GetRxOpsPmOk

`func (o *ObjectStorageUserStat) GetRxOpsPmOk() (*float64, bool)`

GetRxOpsPmOk returns a tuple with the RxOpsPm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxOpsPm

`func (o *ObjectStorageUserStat) SetRxOpsPm(v float64)`

SetRxOpsPm sets RxOpsPm field to given value.

### HasRxOpsPm

`func (o *ObjectStorageUserStat) HasRxOpsPm() bool`

HasRxOpsPm returns a boolean if a field has been set.

### GetTotalDelOps

`func (o *ObjectStorageUserStat) GetTotalDelOps() int64`

GetTotalDelOps returns the TotalDelOps field if non-nil, zero value otherwise.

### GetTotalDelOpsOk

`func (o *ObjectStorageUserStat) GetTotalDelOpsOk() (*int64, bool)`

GetTotalDelOpsOk returns a tuple with the TotalDelOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDelOps

`func (o *ObjectStorageUserStat) SetTotalDelOps(v int64)`

SetTotalDelOps sets TotalDelOps field to given value.

### HasTotalDelOps

`func (o *ObjectStorageUserStat) HasTotalDelOps() bool`

HasTotalDelOps returns a boolean if a field has been set.

### GetTotalDelSuccessOps

`func (o *ObjectStorageUserStat) GetTotalDelSuccessOps() int64`

GetTotalDelSuccessOps returns the TotalDelSuccessOps field if non-nil, zero value otherwise.

### GetTotalDelSuccessOpsOk

`func (o *ObjectStorageUserStat) GetTotalDelSuccessOpsOk() (*int64, bool)`

GetTotalDelSuccessOpsOk returns a tuple with the TotalDelSuccessOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDelSuccessOps

`func (o *ObjectStorageUserStat) SetTotalDelSuccessOps(v int64)`

SetTotalDelSuccessOps sets TotalDelSuccessOps field to given value.

### HasTotalDelSuccessOps

`func (o *ObjectStorageUserStat) HasTotalDelSuccessOps() bool`

HasTotalDelSuccessOps returns a boolean if a field has been set.

### GetTotalRxBytes

`func (o *ObjectStorageUserStat) GetTotalRxBytes() int64`

GetTotalRxBytes returns the TotalRxBytes field if non-nil, zero value otherwise.

### GetTotalRxBytesOk

`func (o *ObjectStorageUserStat) GetTotalRxBytesOk() (*int64, bool)`

GetTotalRxBytesOk returns a tuple with the TotalRxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRxBytes

`func (o *ObjectStorageUserStat) SetTotalRxBytes(v int64)`

SetTotalRxBytes sets TotalRxBytes field to given value.

### HasTotalRxBytes

`func (o *ObjectStorageUserStat) HasTotalRxBytes() bool`

HasTotalRxBytes returns a boolean if a field has been set.

### GetTotalRxOps

`func (o *ObjectStorageUserStat) GetTotalRxOps() int64`

GetTotalRxOps returns the TotalRxOps field if non-nil, zero value otherwise.

### GetTotalRxOpsOk

`func (o *ObjectStorageUserStat) GetTotalRxOpsOk() (*int64, bool)`

GetTotalRxOpsOk returns a tuple with the TotalRxOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRxOps

`func (o *ObjectStorageUserStat) SetTotalRxOps(v int64)`

SetTotalRxOps sets TotalRxOps field to given value.

### HasTotalRxOps

`func (o *ObjectStorageUserStat) HasTotalRxOps() bool`

HasTotalRxOps returns a boolean if a field has been set.

### GetTotalRxSuccessOps

`func (o *ObjectStorageUserStat) GetTotalRxSuccessOps() int64`

GetTotalRxSuccessOps returns the TotalRxSuccessOps field if non-nil, zero value otherwise.

### GetTotalRxSuccessOpsOk

`func (o *ObjectStorageUserStat) GetTotalRxSuccessOpsOk() (*int64, bool)`

GetTotalRxSuccessOpsOk returns a tuple with the TotalRxSuccessOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRxSuccessOps

`func (o *ObjectStorageUserStat) SetTotalRxSuccessOps(v int64)`

SetTotalRxSuccessOps sets TotalRxSuccessOps field to given value.

### HasTotalRxSuccessOps

`func (o *ObjectStorageUserStat) HasTotalRxSuccessOps() bool`

HasTotalRxSuccessOps returns a boolean if a field has been set.

### GetTotalTxBytes

`func (o *ObjectStorageUserStat) GetTotalTxBytes() int64`

GetTotalTxBytes returns the TotalTxBytes field if non-nil, zero value otherwise.

### GetTotalTxBytesOk

`func (o *ObjectStorageUserStat) GetTotalTxBytesOk() (*int64, bool)`

GetTotalTxBytesOk returns a tuple with the TotalTxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTxBytes

`func (o *ObjectStorageUserStat) SetTotalTxBytes(v int64)`

SetTotalTxBytes sets TotalTxBytes field to given value.

### HasTotalTxBytes

`func (o *ObjectStorageUserStat) HasTotalTxBytes() bool`

HasTotalTxBytes returns a boolean if a field has been set.

### GetTotalTxOps

`func (o *ObjectStorageUserStat) GetTotalTxOps() int64`

GetTotalTxOps returns the TotalTxOps field if non-nil, zero value otherwise.

### GetTotalTxOpsOk

`func (o *ObjectStorageUserStat) GetTotalTxOpsOk() (*int64, bool)`

GetTotalTxOpsOk returns a tuple with the TotalTxOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTxOps

`func (o *ObjectStorageUserStat) SetTotalTxOps(v int64)`

SetTotalTxOps sets TotalTxOps field to given value.

### HasTotalTxOps

`func (o *ObjectStorageUserStat) HasTotalTxOps() bool`

HasTotalTxOps returns a boolean if a field has been set.

### GetTotalTxSuccessOps

`func (o *ObjectStorageUserStat) GetTotalTxSuccessOps() int64`

GetTotalTxSuccessOps returns the TotalTxSuccessOps field if non-nil, zero value otherwise.

### GetTotalTxSuccessOpsOk

`func (o *ObjectStorageUserStat) GetTotalTxSuccessOpsOk() (*int64, bool)`

GetTotalTxSuccessOpsOk returns a tuple with the TotalTxSuccessOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTxSuccessOps

`func (o *ObjectStorageUserStat) SetTotalTxSuccessOps(v int64)`

SetTotalTxSuccessOps sets TotalTxSuccessOps field to given value.

### HasTotalTxSuccessOps

`func (o *ObjectStorageUserStat) HasTotalTxSuccessOps() bool`

HasTotalTxSuccessOps returns a boolean if a field has been set.

### GetTxBandwidthKbyte

`func (o *ObjectStorageUserStat) GetTxBandwidthKbyte() float64`

GetTxBandwidthKbyte returns the TxBandwidthKbyte field if non-nil, zero value otherwise.

### GetTxBandwidthKbyteOk

`func (o *ObjectStorageUserStat) GetTxBandwidthKbyteOk() (*float64, bool)`

GetTxBandwidthKbyteOk returns a tuple with the TxBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBandwidthKbyte

`func (o *ObjectStorageUserStat) SetTxBandwidthKbyte(v float64)`

SetTxBandwidthKbyte sets TxBandwidthKbyte field to given value.

### HasTxBandwidthKbyte

`func (o *ObjectStorageUserStat) HasTxBandwidthKbyte() bool`

HasTxBandwidthKbyte returns a boolean if a field has been set.

### GetTxOpsPm

`func (o *ObjectStorageUserStat) GetTxOpsPm() float64`

GetTxOpsPm returns the TxOpsPm field if non-nil, zero value otherwise.

### GetTxOpsPmOk

`func (o *ObjectStorageUserStat) GetTxOpsPmOk() (*float64, bool)`

GetTxOpsPmOk returns a tuple with the TxOpsPm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxOpsPm

`func (o *ObjectStorageUserStat) SetTxOpsPm(v float64)`

SetTxOpsPm sets TxOpsPm field to given value.

### HasTxOpsPm

`func (o *ObjectStorageUserStat) HasTxOpsPm() bool`

HasTxOpsPm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


