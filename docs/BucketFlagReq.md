# BucketFlagReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteTriggerArchiveTier** | Pointer to **bool** |  | [optional] 
**GetTriggerRestore** | Pointer to **bool** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**TierCache** | Pointer to **bool** |  | [optional] 
**TierRetainExternalObject** | Pointer to **bool** | set bucket local delete mode | [optional] 
**TierWorm** | Pointer to **bool** |  | [optional] 
**Versioned** | Pointer to **bool** |  | [optional] 
**VersionsSuspended** | Pointer to **bool** |  | [optional] 
**Worm** | Pointer to **bool** |  | [optional] 

## Methods

### NewBucketFlagReq

`func NewBucketFlagReq() *BucketFlagReq`

NewBucketFlagReq instantiates a new BucketFlagReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketFlagReqWithDefaults

`func NewBucketFlagReqWithDefaults() *BucketFlagReq`

NewBucketFlagReqWithDefaults instantiates a new BucketFlagReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteTriggerArchiveTier

`func (o *BucketFlagReq) GetDeleteTriggerArchiveTier() bool`

GetDeleteTriggerArchiveTier returns the DeleteTriggerArchiveTier field if non-nil, zero value otherwise.

### GetDeleteTriggerArchiveTierOk

`func (o *BucketFlagReq) GetDeleteTriggerArchiveTierOk() (*bool, bool)`

GetDeleteTriggerArchiveTierOk returns a tuple with the DeleteTriggerArchiveTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteTriggerArchiveTier

`func (o *BucketFlagReq) SetDeleteTriggerArchiveTier(v bool)`

SetDeleteTriggerArchiveTier sets DeleteTriggerArchiveTier field to given value.

### HasDeleteTriggerArchiveTier

`func (o *BucketFlagReq) HasDeleteTriggerArchiveTier() bool`

HasDeleteTriggerArchiveTier returns a boolean if a field has been set.

### GetGetTriggerRestore

`func (o *BucketFlagReq) GetGetTriggerRestore() bool`

GetGetTriggerRestore returns the GetTriggerRestore field if non-nil, zero value otherwise.

### GetGetTriggerRestoreOk

`func (o *BucketFlagReq) GetGetTriggerRestoreOk() (*bool, bool)`

GetGetTriggerRestoreOk returns a tuple with the GetTriggerRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetTriggerRestore

`func (o *BucketFlagReq) SetGetTriggerRestore(v bool)`

SetGetTriggerRestore sets GetTriggerRestore field to given value.

### HasGetTriggerRestore

`func (o *BucketFlagReq) HasGetTriggerRestore() bool`

HasGetTriggerRestore returns a boolean if a field has been set.

### GetReadonly

`func (o *BucketFlagReq) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *BucketFlagReq) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *BucketFlagReq) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *BucketFlagReq) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetTierCache

`func (o *BucketFlagReq) GetTierCache() bool`

GetTierCache returns the TierCache field if non-nil, zero value otherwise.

### GetTierCacheOk

`func (o *BucketFlagReq) GetTierCacheOk() (*bool, bool)`

GetTierCacheOk returns a tuple with the TierCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierCache

`func (o *BucketFlagReq) SetTierCache(v bool)`

SetTierCache sets TierCache field to given value.

### HasTierCache

`func (o *BucketFlagReq) HasTierCache() bool`

HasTierCache returns a boolean if a field has been set.

### GetTierRetainExternalObject

`func (o *BucketFlagReq) GetTierRetainExternalObject() bool`

GetTierRetainExternalObject returns the TierRetainExternalObject field if non-nil, zero value otherwise.

### GetTierRetainExternalObjectOk

`func (o *BucketFlagReq) GetTierRetainExternalObjectOk() (*bool, bool)`

GetTierRetainExternalObjectOk returns a tuple with the TierRetainExternalObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierRetainExternalObject

`func (o *BucketFlagReq) SetTierRetainExternalObject(v bool)`

SetTierRetainExternalObject sets TierRetainExternalObject field to given value.

### HasTierRetainExternalObject

`func (o *BucketFlagReq) HasTierRetainExternalObject() bool`

HasTierRetainExternalObject returns a boolean if a field has been set.

### GetTierWorm

`func (o *BucketFlagReq) GetTierWorm() bool`

GetTierWorm returns the TierWorm field if non-nil, zero value otherwise.

### GetTierWormOk

`func (o *BucketFlagReq) GetTierWormOk() (*bool, bool)`

GetTierWormOk returns a tuple with the TierWorm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierWorm

`func (o *BucketFlagReq) SetTierWorm(v bool)`

SetTierWorm sets TierWorm field to given value.

### HasTierWorm

`func (o *BucketFlagReq) HasTierWorm() bool`

HasTierWorm returns a boolean if a field has been set.

### GetVersioned

`func (o *BucketFlagReq) GetVersioned() bool`

GetVersioned returns the Versioned field if non-nil, zero value otherwise.

### GetVersionedOk

`func (o *BucketFlagReq) GetVersionedOk() (*bool, bool)`

GetVersionedOk returns a tuple with the Versioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioned

`func (o *BucketFlagReq) SetVersioned(v bool)`

SetVersioned sets Versioned field to given value.

### HasVersioned

`func (o *BucketFlagReq) HasVersioned() bool`

HasVersioned returns a boolean if a field has been set.

### GetVersionsSuspended

`func (o *BucketFlagReq) GetVersionsSuspended() bool`

GetVersionsSuspended returns the VersionsSuspended field if non-nil, zero value otherwise.

### GetVersionsSuspendedOk

`func (o *BucketFlagReq) GetVersionsSuspendedOk() (*bool, bool)`

GetVersionsSuspendedOk returns a tuple with the VersionsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsSuspended

`func (o *BucketFlagReq) SetVersionsSuspended(v bool)`

SetVersionsSuspended sets VersionsSuspended field to given value.

### HasVersionsSuspended

`func (o *BucketFlagReq) HasVersionsSuspended() bool`

HasVersionsSuspended returns a boolean if a field has been set.

### GetWorm

`func (o *BucketFlagReq) GetWorm() bool`

GetWorm returns the Worm field if non-nil, zero value otherwise.

### GetWormOk

`func (o *BucketFlagReq) GetWormOk() (*bool, bool)`

GetWormOk returns a tuple with the Worm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorm

`func (o *BucketFlagReq) SetWorm(v bool)`

SetWorm sets Worm field to given value.

### HasWorm

`func (o *BucketFlagReq) HasWorm() bool`

HasWorm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


