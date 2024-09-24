# BucketFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteTriggerArchiveTier** | Pointer to **bool** |  | [optional] 
**GetTriggerRestore** | Pointer to **bool** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**Recycle** | Pointer to **bool** |  | [optional] 
**TierCache** | Pointer to **bool** |  | [optional] 
**TierRetainExternalObject** | Pointer to **bool** | set bucket local delete mode | [optional] 
**TierWorm** | Pointer to **bool** |  | [optional] 
**Versioned** | Pointer to **bool** |  | [optional] 
**VersionsSuspended** | Pointer to **bool** |  | [optional] 
**Worm** | Pointer to **bool** |  | [optional] 

## Methods

### NewBucketFlag

`func NewBucketFlag() *BucketFlag`

NewBucketFlag instantiates a new BucketFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketFlagWithDefaults

`func NewBucketFlagWithDefaults() *BucketFlag`

NewBucketFlagWithDefaults instantiates a new BucketFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteTriggerArchiveTier

`func (o *BucketFlag) GetDeleteTriggerArchiveTier() bool`

GetDeleteTriggerArchiveTier returns the DeleteTriggerArchiveTier field if non-nil, zero value otherwise.

### GetDeleteTriggerArchiveTierOk

`func (o *BucketFlag) GetDeleteTriggerArchiveTierOk() (*bool, bool)`

GetDeleteTriggerArchiveTierOk returns a tuple with the DeleteTriggerArchiveTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteTriggerArchiveTier

`func (o *BucketFlag) SetDeleteTriggerArchiveTier(v bool)`

SetDeleteTriggerArchiveTier sets DeleteTriggerArchiveTier field to given value.

### HasDeleteTriggerArchiveTier

`func (o *BucketFlag) HasDeleteTriggerArchiveTier() bool`

HasDeleteTriggerArchiveTier returns a boolean if a field has been set.

### GetGetTriggerRestore

`func (o *BucketFlag) GetGetTriggerRestore() bool`

GetGetTriggerRestore returns the GetTriggerRestore field if non-nil, zero value otherwise.

### GetGetTriggerRestoreOk

`func (o *BucketFlag) GetGetTriggerRestoreOk() (*bool, bool)`

GetGetTriggerRestoreOk returns a tuple with the GetTriggerRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetTriggerRestore

`func (o *BucketFlag) SetGetTriggerRestore(v bool)`

SetGetTriggerRestore sets GetTriggerRestore field to given value.

### HasGetTriggerRestore

`func (o *BucketFlag) HasGetTriggerRestore() bool`

HasGetTriggerRestore returns a boolean if a field has been set.

### GetReadonly

`func (o *BucketFlag) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *BucketFlag) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *BucketFlag) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *BucketFlag) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetRecycle

`func (o *BucketFlag) GetRecycle() bool`

GetRecycle returns the Recycle field if non-nil, zero value otherwise.

### GetRecycleOk

`func (o *BucketFlag) GetRecycleOk() (*bool, bool)`

GetRecycleOk returns a tuple with the Recycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycle

`func (o *BucketFlag) SetRecycle(v bool)`

SetRecycle sets Recycle field to given value.

### HasRecycle

`func (o *BucketFlag) HasRecycle() bool`

HasRecycle returns a boolean if a field has been set.

### GetTierCache

`func (o *BucketFlag) GetTierCache() bool`

GetTierCache returns the TierCache field if non-nil, zero value otherwise.

### GetTierCacheOk

`func (o *BucketFlag) GetTierCacheOk() (*bool, bool)`

GetTierCacheOk returns a tuple with the TierCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierCache

`func (o *BucketFlag) SetTierCache(v bool)`

SetTierCache sets TierCache field to given value.

### HasTierCache

`func (o *BucketFlag) HasTierCache() bool`

HasTierCache returns a boolean if a field has been set.

### GetTierRetainExternalObject

`func (o *BucketFlag) GetTierRetainExternalObject() bool`

GetTierRetainExternalObject returns the TierRetainExternalObject field if non-nil, zero value otherwise.

### GetTierRetainExternalObjectOk

`func (o *BucketFlag) GetTierRetainExternalObjectOk() (*bool, bool)`

GetTierRetainExternalObjectOk returns a tuple with the TierRetainExternalObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierRetainExternalObject

`func (o *BucketFlag) SetTierRetainExternalObject(v bool)`

SetTierRetainExternalObject sets TierRetainExternalObject field to given value.

### HasTierRetainExternalObject

`func (o *BucketFlag) HasTierRetainExternalObject() bool`

HasTierRetainExternalObject returns a boolean if a field has been set.

### GetTierWorm

`func (o *BucketFlag) GetTierWorm() bool`

GetTierWorm returns the TierWorm field if non-nil, zero value otherwise.

### GetTierWormOk

`func (o *BucketFlag) GetTierWormOk() (*bool, bool)`

GetTierWormOk returns a tuple with the TierWorm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierWorm

`func (o *BucketFlag) SetTierWorm(v bool)`

SetTierWorm sets TierWorm field to given value.

### HasTierWorm

`func (o *BucketFlag) HasTierWorm() bool`

HasTierWorm returns a boolean if a field has been set.

### GetVersioned

`func (o *BucketFlag) GetVersioned() bool`

GetVersioned returns the Versioned field if non-nil, zero value otherwise.

### GetVersionedOk

`func (o *BucketFlag) GetVersionedOk() (*bool, bool)`

GetVersionedOk returns a tuple with the Versioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioned

`func (o *BucketFlag) SetVersioned(v bool)`

SetVersioned sets Versioned field to given value.

### HasVersioned

`func (o *BucketFlag) HasVersioned() bool`

HasVersioned returns a boolean if a field has been set.

### GetVersionsSuspended

`func (o *BucketFlag) GetVersionsSuspended() bool`

GetVersionsSuspended returns the VersionsSuspended field if non-nil, zero value otherwise.

### GetVersionsSuspendedOk

`func (o *BucketFlag) GetVersionsSuspendedOk() (*bool, bool)`

GetVersionsSuspendedOk returns a tuple with the VersionsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsSuspended

`func (o *BucketFlag) SetVersionsSuspended(v bool)`

SetVersionsSuspended sets VersionsSuspended field to given value.

### HasVersionsSuspended

`func (o *BucketFlag) HasVersionsSuspended() bool`

HasVersionsSuspended returns a boolean if a field has been set.

### GetWorm

`func (o *BucketFlag) GetWorm() bool`

GetWorm returns the Worm field if non-nil, zero value otherwise.

### GetWormOk

`func (o *BucketFlag) GetWormOk() (*bool, bool)`

GetWormOk returns a tuple with the Worm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorm

`func (o *BucketFlag) SetWorm(v bool)`

SetWorm sets Worm field to given value.

### HasWorm

`func (o *BucketFlag) HasWorm() bool`

HasWorm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


