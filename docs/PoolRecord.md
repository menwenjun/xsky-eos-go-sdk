# PoolRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AllocatedSize** | Pointer to **int64** |  | [optional] 
**BindOsdNum** | Pointer to **int64** |  | [optional] 
**BlockVolumeNum** | Pointer to **int64** |  | [optional] 
**CachePool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**CodingChunkNum** | Pointer to **int64** |  | [optional] 
**CompressAlgorithm** | Pointer to **string** |  | [optional] 
**Compressed** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DataChunkNum** | Pointer to **int64** |  | [optional] 
**DataPool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**DefaultManagedVolumeFormat** | Pointer to **int64** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**DeviceTypeCheckDisabled** | Pointer to **bool** |  | [optional] 
**EncryptEnabled** | Pointer to **bool** |  | [optional] 
**FailureDomainType** | Pointer to **string** |  | [optional] 
**GcQosPlan** | Pointer to [**PoolGCPolicyPlan**](PoolGCPolicyPlan.md) |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InnerPoolId** | Pointer to **int32** |  | [optional] 
**IoBypassEnabled** | Pointer to **bool** |  | [optional] 
**IoBypassMode** | Pointer to **string** |  | [optional] 
**IoBypassThreshold** | Pointer to **int64** |  | [optional] 
**MinAllocSize** | Pointer to **int64** |  | [optional] 
**MinOsdNum** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumaApplyPolicy** | Pointer to **string** |  | [optional] 
**NumaBindBalanced** | Pointer to **bool** |  | [optional] 
**NumaBindPolicy** | Pointer to **string** |  | [optional] 
**NumaEnabled** | Pointer to **bool** |  | [optional] 
**OsdGroup** | Pointer to [**OsdGroupNestview**](OsdGroupNestview.md) |  | [optional] 
**OsdNum** | Pointer to **int64** |  | [optional] 
**OsdNumPerHost** | Pointer to **int64** |  | [optional] 
**OutFailureDomainNum** | Pointer to **int64** |  | [optional] 
**PoolId** | Pointer to **int32** |  | [optional] 
**PoolMode** | Pointer to **string** |  | [optional] 
**PoolName** | Pointer to **string** |  | [optional] 
**PoolRole** | Pointer to **string** |  | [optional] 
**PoolType** | Pointer to **string** |  | [optional] 
**PrimaryPlacementNode** | Pointer to [**PlacementNodeNestview**](PlacementNodeNestview.md) |  | [optional] 
**ProductType** | Pointer to **string** |  | [optional] 
**Property** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ProtectionDomain** | Pointer to [**ProtectionDomainNestview**](ProtectionDomainNestview.md) |  | [optional] 
**Qos** | Pointer to [**OsdQos**](OsdQos.md) |  | [optional] 
**ReplicateSize** | Pointer to **int64** |  | [optional] 
**ReservedPercent** | Pointer to **float64** |  | [optional] 
**SizePerOsd** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Stretched** | Pointer to **bool** |  | [optional] 
**StripeUnit** | Pointer to **int64** |  | [optional] 
**SubFailureDomainType** | Pointer to **string** |  | [optional] 
**SuggestedOmapSize** | Pointer to **int64** |  | [optional] 
**ThinProvisioned** | Pointer to **bool** |  | [optional] 
**TierStatus** | Pointer to **string** |  | [optional] 
**TransportMode** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]PoolStat**](PoolStat.md) |  | [optional] 

## Methods

### NewPoolRecord

`func NewPoolRecord() *PoolRecord`

NewPoolRecord instantiates a new PoolRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolRecordWithDefaults

`func NewPoolRecordWithDefaults() *PoolRecord`

NewPoolRecordWithDefaults instantiates a new PoolRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *PoolRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *PoolRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *PoolRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *PoolRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *PoolRecord) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *PoolRecord) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *PoolRecord) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *PoolRecord) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetBindOsdNum

`func (o *PoolRecord) GetBindOsdNum() int64`

GetBindOsdNum returns the BindOsdNum field if non-nil, zero value otherwise.

### GetBindOsdNumOk

`func (o *PoolRecord) GetBindOsdNumOk() (*int64, bool)`

GetBindOsdNumOk returns a tuple with the BindOsdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindOsdNum

`func (o *PoolRecord) SetBindOsdNum(v int64)`

SetBindOsdNum sets BindOsdNum field to given value.

### HasBindOsdNum

`func (o *PoolRecord) HasBindOsdNum() bool`

HasBindOsdNum returns a boolean if a field has been set.

### GetBlockVolumeNum

`func (o *PoolRecord) GetBlockVolumeNum() int64`

GetBlockVolumeNum returns the BlockVolumeNum field if non-nil, zero value otherwise.

### GetBlockVolumeNumOk

`func (o *PoolRecord) GetBlockVolumeNumOk() (*int64, bool)`

GetBlockVolumeNumOk returns a tuple with the BlockVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeNum

`func (o *PoolRecord) SetBlockVolumeNum(v int64)`

SetBlockVolumeNum sets BlockVolumeNum field to given value.

### HasBlockVolumeNum

`func (o *PoolRecord) HasBlockVolumeNum() bool`

HasBlockVolumeNum returns a boolean if a field has been set.

### GetCachePool

`func (o *PoolRecord) GetCachePool() PoolNestview`

GetCachePool returns the CachePool field if non-nil, zero value otherwise.

### GetCachePoolOk

`func (o *PoolRecord) GetCachePoolOk() (*PoolNestview, bool)`

GetCachePoolOk returns a tuple with the CachePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePool

`func (o *PoolRecord) SetCachePool(v PoolNestview)`

SetCachePool sets CachePool field to given value.

### HasCachePool

`func (o *PoolRecord) HasCachePool() bool`

HasCachePool returns a boolean if a field has been set.

### GetCluster

`func (o *PoolRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PoolRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PoolRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PoolRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCodingChunkNum

`func (o *PoolRecord) GetCodingChunkNum() int64`

GetCodingChunkNum returns the CodingChunkNum field if non-nil, zero value otherwise.

### GetCodingChunkNumOk

`func (o *PoolRecord) GetCodingChunkNumOk() (*int64, bool)`

GetCodingChunkNumOk returns a tuple with the CodingChunkNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodingChunkNum

`func (o *PoolRecord) SetCodingChunkNum(v int64)`

SetCodingChunkNum sets CodingChunkNum field to given value.

### HasCodingChunkNum

`func (o *PoolRecord) HasCodingChunkNum() bool`

HasCodingChunkNum returns a boolean if a field has been set.

### GetCompressAlgorithm

`func (o *PoolRecord) GetCompressAlgorithm() string`

GetCompressAlgorithm returns the CompressAlgorithm field if non-nil, zero value otherwise.

### GetCompressAlgorithmOk

`func (o *PoolRecord) GetCompressAlgorithmOk() (*string, bool)`

GetCompressAlgorithmOk returns a tuple with the CompressAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressAlgorithm

`func (o *PoolRecord) SetCompressAlgorithm(v string)`

SetCompressAlgorithm sets CompressAlgorithm field to given value.

### HasCompressAlgorithm

`func (o *PoolRecord) HasCompressAlgorithm() bool`

HasCompressAlgorithm returns a boolean if a field has been set.

### GetCompressed

`func (o *PoolRecord) GetCompressed() bool`

GetCompressed returns the Compressed field if non-nil, zero value otherwise.

### GetCompressedOk

`func (o *PoolRecord) GetCompressedOk() (*bool, bool)`

GetCompressedOk returns a tuple with the Compressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressed

`func (o *PoolRecord) SetCompressed(v bool)`

SetCompressed sets Compressed field to given value.

### HasCompressed

`func (o *PoolRecord) HasCompressed() bool`

HasCompressed returns a boolean if a field has been set.

### GetCreate

`func (o *PoolRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PoolRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PoolRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PoolRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataChunkNum

`func (o *PoolRecord) GetDataChunkNum() int64`

GetDataChunkNum returns the DataChunkNum field if non-nil, zero value otherwise.

### GetDataChunkNumOk

`func (o *PoolRecord) GetDataChunkNumOk() (*int64, bool)`

GetDataChunkNumOk returns a tuple with the DataChunkNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChunkNum

`func (o *PoolRecord) SetDataChunkNum(v int64)`

SetDataChunkNum sets DataChunkNum field to given value.

### HasDataChunkNum

`func (o *PoolRecord) HasDataChunkNum() bool`

HasDataChunkNum returns a boolean if a field has been set.

### GetDataPool

`func (o *PoolRecord) GetDataPool() PoolNestview`

GetDataPool returns the DataPool field if non-nil, zero value otherwise.

### GetDataPoolOk

`func (o *PoolRecord) GetDataPoolOk() (*PoolNestview, bool)`

GetDataPoolOk returns a tuple with the DataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPool

`func (o *PoolRecord) SetDataPool(v PoolNestview)`

SetDataPool sets DataPool field to given value.

### HasDataPool

`func (o *PoolRecord) HasDataPool() bool`

HasDataPool returns a boolean if a field has been set.

### GetDefaultManagedVolumeFormat

`func (o *PoolRecord) GetDefaultManagedVolumeFormat() int64`

GetDefaultManagedVolumeFormat returns the DefaultManagedVolumeFormat field if non-nil, zero value otherwise.

### GetDefaultManagedVolumeFormatOk

`func (o *PoolRecord) GetDefaultManagedVolumeFormatOk() (*int64, bool)`

GetDefaultManagedVolumeFormatOk returns a tuple with the DefaultManagedVolumeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultManagedVolumeFormat

`func (o *PoolRecord) SetDefaultManagedVolumeFormat(v int64)`

SetDefaultManagedVolumeFormat sets DefaultManagedVolumeFormat field to given value.

### HasDefaultManagedVolumeFormat

`func (o *PoolRecord) HasDefaultManagedVolumeFormat() bool`

HasDefaultManagedVolumeFormat returns a boolean if a field has been set.

### GetDeviceType

`func (o *PoolRecord) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PoolRecord) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PoolRecord) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PoolRecord) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceTypeCheckDisabled

`func (o *PoolRecord) GetDeviceTypeCheckDisabled() bool`

GetDeviceTypeCheckDisabled returns the DeviceTypeCheckDisabled field if non-nil, zero value otherwise.

### GetDeviceTypeCheckDisabledOk

`func (o *PoolRecord) GetDeviceTypeCheckDisabledOk() (*bool, bool)`

GetDeviceTypeCheckDisabledOk returns a tuple with the DeviceTypeCheckDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCheckDisabled

`func (o *PoolRecord) SetDeviceTypeCheckDisabled(v bool)`

SetDeviceTypeCheckDisabled sets DeviceTypeCheckDisabled field to given value.

### HasDeviceTypeCheckDisabled

`func (o *PoolRecord) HasDeviceTypeCheckDisabled() bool`

HasDeviceTypeCheckDisabled returns a boolean if a field has been set.

### GetEncryptEnabled

`func (o *PoolRecord) GetEncryptEnabled() bool`

GetEncryptEnabled returns the EncryptEnabled field if non-nil, zero value otherwise.

### GetEncryptEnabledOk

`func (o *PoolRecord) GetEncryptEnabledOk() (*bool, bool)`

GetEncryptEnabledOk returns a tuple with the EncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEnabled

`func (o *PoolRecord) SetEncryptEnabled(v bool)`

SetEncryptEnabled sets EncryptEnabled field to given value.

### HasEncryptEnabled

`func (o *PoolRecord) HasEncryptEnabled() bool`

HasEncryptEnabled returns a boolean if a field has been set.

### GetFailureDomainType

`func (o *PoolRecord) GetFailureDomainType() string`

GetFailureDomainType returns the FailureDomainType field if non-nil, zero value otherwise.

### GetFailureDomainTypeOk

`func (o *PoolRecord) GetFailureDomainTypeOk() (*string, bool)`

GetFailureDomainTypeOk returns a tuple with the FailureDomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDomainType

`func (o *PoolRecord) SetFailureDomainType(v string)`

SetFailureDomainType sets FailureDomainType field to given value.

### HasFailureDomainType

`func (o *PoolRecord) HasFailureDomainType() bool`

HasFailureDomainType returns a boolean if a field has been set.

### GetGcQosPlan

`func (o *PoolRecord) GetGcQosPlan() PoolGCPolicyPlan`

GetGcQosPlan returns the GcQosPlan field if non-nil, zero value otherwise.

### GetGcQosPlanOk

`func (o *PoolRecord) GetGcQosPlanOk() (*PoolGCPolicyPlan, bool)`

GetGcQosPlanOk returns a tuple with the GcQosPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcQosPlan

`func (o *PoolRecord) SetGcQosPlan(v PoolGCPolicyPlan)`

SetGcQosPlan sets GcQosPlan field to given value.

### HasGcQosPlan

`func (o *PoolRecord) HasGcQosPlan() bool`

HasGcQosPlan returns a boolean if a field has been set.

### GetHidden

`func (o *PoolRecord) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PoolRecord) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PoolRecord) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PoolRecord) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *PoolRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PoolRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInnerPoolId

`func (o *PoolRecord) GetInnerPoolId() int32`

GetInnerPoolId returns the InnerPoolId field if non-nil, zero value otherwise.

### GetInnerPoolIdOk

`func (o *PoolRecord) GetInnerPoolIdOk() (*int32, bool)`

GetInnerPoolIdOk returns a tuple with the InnerPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerPoolId

`func (o *PoolRecord) SetInnerPoolId(v int32)`

SetInnerPoolId sets InnerPoolId field to given value.

### HasInnerPoolId

`func (o *PoolRecord) HasInnerPoolId() bool`

HasInnerPoolId returns a boolean if a field has been set.

### GetIoBypassEnabled

`func (o *PoolRecord) GetIoBypassEnabled() bool`

GetIoBypassEnabled returns the IoBypassEnabled field if non-nil, zero value otherwise.

### GetIoBypassEnabledOk

`func (o *PoolRecord) GetIoBypassEnabledOk() (*bool, bool)`

GetIoBypassEnabledOk returns a tuple with the IoBypassEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoBypassEnabled

`func (o *PoolRecord) SetIoBypassEnabled(v bool)`

SetIoBypassEnabled sets IoBypassEnabled field to given value.

### HasIoBypassEnabled

`func (o *PoolRecord) HasIoBypassEnabled() bool`

HasIoBypassEnabled returns a boolean if a field has been set.

### GetIoBypassMode

`func (o *PoolRecord) GetIoBypassMode() string`

GetIoBypassMode returns the IoBypassMode field if non-nil, zero value otherwise.

### GetIoBypassModeOk

`func (o *PoolRecord) GetIoBypassModeOk() (*string, bool)`

GetIoBypassModeOk returns a tuple with the IoBypassMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoBypassMode

`func (o *PoolRecord) SetIoBypassMode(v string)`

SetIoBypassMode sets IoBypassMode field to given value.

### HasIoBypassMode

`func (o *PoolRecord) HasIoBypassMode() bool`

HasIoBypassMode returns a boolean if a field has been set.

### GetIoBypassThreshold

`func (o *PoolRecord) GetIoBypassThreshold() int64`

GetIoBypassThreshold returns the IoBypassThreshold field if non-nil, zero value otherwise.

### GetIoBypassThresholdOk

`func (o *PoolRecord) GetIoBypassThresholdOk() (*int64, bool)`

GetIoBypassThresholdOk returns a tuple with the IoBypassThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoBypassThreshold

`func (o *PoolRecord) SetIoBypassThreshold(v int64)`

SetIoBypassThreshold sets IoBypassThreshold field to given value.

### HasIoBypassThreshold

`func (o *PoolRecord) HasIoBypassThreshold() bool`

HasIoBypassThreshold returns a boolean if a field has been set.

### GetMinAllocSize

`func (o *PoolRecord) GetMinAllocSize() int64`

GetMinAllocSize returns the MinAllocSize field if non-nil, zero value otherwise.

### GetMinAllocSizeOk

`func (o *PoolRecord) GetMinAllocSizeOk() (*int64, bool)`

GetMinAllocSizeOk returns a tuple with the MinAllocSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAllocSize

`func (o *PoolRecord) SetMinAllocSize(v int64)`

SetMinAllocSize sets MinAllocSize field to given value.

### HasMinAllocSize

`func (o *PoolRecord) HasMinAllocSize() bool`

HasMinAllocSize returns a boolean if a field has been set.

### GetMinOsdNum

`func (o *PoolRecord) GetMinOsdNum() int64`

GetMinOsdNum returns the MinOsdNum field if non-nil, zero value otherwise.

### GetMinOsdNumOk

`func (o *PoolRecord) GetMinOsdNumOk() (*int64, bool)`

GetMinOsdNumOk returns a tuple with the MinOsdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOsdNum

`func (o *PoolRecord) SetMinOsdNum(v int64)`

SetMinOsdNum sets MinOsdNum field to given value.

### HasMinOsdNum

`func (o *PoolRecord) HasMinOsdNum() bool`

HasMinOsdNum returns a boolean if a field has been set.

### GetName

`func (o *PoolRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumaApplyPolicy

`func (o *PoolRecord) GetNumaApplyPolicy() string`

GetNumaApplyPolicy returns the NumaApplyPolicy field if non-nil, zero value otherwise.

### GetNumaApplyPolicyOk

`func (o *PoolRecord) GetNumaApplyPolicyOk() (*string, bool)`

GetNumaApplyPolicyOk returns a tuple with the NumaApplyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaApplyPolicy

`func (o *PoolRecord) SetNumaApplyPolicy(v string)`

SetNumaApplyPolicy sets NumaApplyPolicy field to given value.

### HasNumaApplyPolicy

`func (o *PoolRecord) HasNumaApplyPolicy() bool`

HasNumaApplyPolicy returns a boolean if a field has been set.

### GetNumaBindBalanced

`func (o *PoolRecord) GetNumaBindBalanced() bool`

GetNumaBindBalanced returns the NumaBindBalanced field if non-nil, zero value otherwise.

### GetNumaBindBalancedOk

`func (o *PoolRecord) GetNumaBindBalancedOk() (*bool, bool)`

GetNumaBindBalancedOk returns a tuple with the NumaBindBalanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaBindBalanced

`func (o *PoolRecord) SetNumaBindBalanced(v bool)`

SetNumaBindBalanced sets NumaBindBalanced field to given value.

### HasNumaBindBalanced

`func (o *PoolRecord) HasNumaBindBalanced() bool`

HasNumaBindBalanced returns a boolean if a field has been set.

### GetNumaBindPolicy

`func (o *PoolRecord) GetNumaBindPolicy() string`

GetNumaBindPolicy returns the NumaBindPolicy field if non-nil, zero value otherwise.

### GetNumaBindPolicyOk

`func (o *PoolRecord) GetNumaBindPolicyOk() (*string, bool)`

GetNumaBindPolicyOk returns a tuple with the NumaBindPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaBindPolicy

`func (o *PoolRecord) SetNumaBindPolicy(v string)`

SetNumaBindPolicy sets NumaBindPolicy field to given value.

### HasNumaBindPolicy

`func (o *PoolRecord) HasNumaBindPolicy() bool`

HasNumaBindPolicy returns a boolean if a field has been set.

### GetNumaEnabled

`func (o *PoolRecord) GetNumaEnabled() bool`

GetNumaEnabled returns the NumaEnabled field if non-nil, zero value otherwise.

### GetNumaEnabledOk

`func (o *PoolRecord) GetNumaEnabledOk() (*bool, bool)`

GetNumaEnabledOk returns a tuple with the NumaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaEnabled

`func (o *PoolRecord) SetNumaEnabled(v bool)`

SetNumaEnabled sets NumaEnabled field to given value.

### HasNumaEnabled

`func (o *PoolRecord) HasNumaEnabled() bool`

HasNumaEnabled returns a boolean if a field has been set.

### GetOsdGroup

`func (o *PoolRecord) GetOsdGroup() OsdGroupNestview`

GetOsdGroup returns the OsdGroup field if non-nil, zero value otherwise.

### GetOsdGroupOk

`func (o *PoolRecord) GetOsdGroupOk() (*OsdGroupNestview, bool)`

GetOsdGroupOk returns a tuple with the OsdGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdGroup

`func (o *PoolRecord) SetOsdGroup(v OsdGroupNestview)`

SetOsdGroup sets OsdGroup field to given value.

### HasOsdGroup

`func (o *PoolRecord) HasOsdGroup() bool`

HasOsdGroup returns a boolean if a field has been set.

### GetOsdNum

`func (o *PoolRecord) GetOsdNum() int64`

GetOsdNum returns the OsdNum field if non-nil, zero value otherwise.

### GetOsdNumOk

`func (o *PoolRecord) GetOsdNumOk() (*int64, bool)`

GetOsdNumOk returns a tuple with the OsdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdNum

`func (o *PoolRecord) SetOsdNum(v int64)`

SetOsdNum sets OsdNum field to given value.

### HasOsdNum

`func (o *PoolRecord) HasOsdNum() bool`

HasOsdNum returns a boolean if a field has been set.

### GetOsdNumPerHost

`func (o *PoolRecord) GetOsdNumPerHost() int64`

GetOsdNumPerHost returns the OsdNumPerHost field if non-nil, zero value otherwise.

### GetOsdNumPerHostOk

`func (o *PoolRecord) GetOsdNumPerHostOk() (*int64, bool)`

GetOsdNumPerHostOk returns a tuple with the OsdNumPerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdNumPerHost

`func (o *PoolRecord) SetOsdNumPerHost(v int64)`

SetOsdNumPerHost sets OsdNumPerHost field to given value.

### HasOsdNumPerHost

`func (o *PoolRecord) HasOsdNumPerHost() bool`

HasOsdNumPerHost returns a boolean if a field has been set.

### GetOutFailureDomainNum

`func (o *PoolRecord) GetOutFailureDomainNum() int64`

GetOutFailureDomainNum returns the OutFailureDomainNum field if non-nil, zero value otherwise.

### GetOutFailureDomainNumOk

`func (o *PoolRecord) GetOutFailureDomainNumOk() (*int64, bool)`

GetOutFailureDomainNumOk returns a tuple with the OutFailureDomainNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutFailureDomainNum

`func (o *PoolRecord) SetOutFailureDomainNum(v int64)`

SetOutFailureDomainNum sets OutFailureDomainNum field to given value.

### HasOutFailureDomainNum

`func (o *PoolRecord) HasOutFailureDomainNum() bool`

HasOutFailureDomainNum returns a boolean if a field has been set.

### GetPoolId

`func (o *PoolRecord) GetPoolId() int32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PoolRecord) GetPoolIdOk() (*int32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PoolRecord) SetPoolId(v int32)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *PoolRecord) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolMode

`func (o *PoolRecord) GetPoolMode() string`

GetPoolMode returns the PoolMode field if non-nil, zero value otherwise.

### GetPoolModeOk

`func (o *PoolRecord) GetPoolModeOk() (*string, bool)`

GetPoolModeOk returns a tuple with the PoolMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMode

`func (o *PoolRecord) SetPoolMode(v string)`

SetPoolMode sets PoolMode field to given value.

### HasPoolMode

`func (o *PoolRecord) HasPoolMode() bool`

HasPoolMode returns a boolean if a field has been set.

### GetPoolName

`func (o *PoolRecord) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *PoolRecord) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *PoolRecord) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *PoolRecord) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolRole

`func (o *PoolRecord) GetPoolRole() string`

GetPoolRole returns the PoolRole field if non-nil, zero value otherwise.

### GetPoolRoleOk

`func (o *PoolRecord) GetPoolRoleOk() (*string, bool)`

GetPoolRoleOk returns a tuple with the PoolRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolRole

`func (o *PoolRecord) SetPoolRole(v string)`

SetPoolRole sets PoolRole field to given value.

### HasPoolRole

`func (o *PoolRecord) HasPoolRole() bool`

HasPoolRole returns a boolean if a field has been set.

### GetPoolType

`func (o *PoolRecord) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *PoolRecord) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *PoolRecord) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *PoolRecord) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetPrimaryPlacementNode

`func (o *PoolRecord) GetPrimaryPlacementNode() PlacementNodeNestview`

GetPrimaryPlacementNode returns the PrimaryPlacementNode field if non-nil, zero value otherwise.

### GetPrimaryPlacementNodeOk

`func (o *PoolRecord) GetPrimaryPlacementNodeOk() (*PlacementNodeNestview, bool)`

GetPrimaryPlacementNodeOk returns a tuple with the PrimaryPlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPlacementNode

`func (o *PoolRecord) SetPrimaryPlacementNode(v PlacementNodeNestview)`

SetPrimaryPlacementNode sets PrimaryPlacementNode field to given value.

### HasPrimaryPlacementNode

`func (o *PoolRecord) HasPrimaryPlacementNode() bool`

HasPrimaryPlacementNode returns a boolean if a field has been set.

### GetProductType

`func (o *PoolRecord) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PoolRecord) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PoolRecord) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *PoolRecord) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProperty

`func (o *PoolRecord) GetProperty() map[string]map[string]interface{}`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PoolRecord) GetPropertyOk() (*map[string]map[string]interface{}, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PoolRecord) SetProperty(v map[string]map[string]interface{})`

SetProperty sets Property field to given value.

### HasProperty

`func (o *PoolRecord) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetProtectionDomain

`func (o *PoolRecord) GetProtectionDomain() ProtectionDomainNestview`

GetProtectionDomain returns the ProtectionDomain field if non-nil, zero value otherwise.

### GetProtectionDomainOk

`func (o *PoolRecord) GetProtectionDomainOk() (*ProtectionDomainNestview, bool)`

GetProtectionDomainOk returns a tuple with the ProtectionDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDomain

`func (o *PoolRecord) SetProtectionDomain(v ProtectionDomainNestview)`

SetProtectionDomain sets ProtectionDomain field to given value.

### HasProtectionDomain

`func (o *PoolRecord) HasProtectionDomain() bool`

HasProtectionDomain returns a boolean if a field has been set.

### GetQos

`func (o *PoolRecord) GetQos() OsdQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *PoolRecord) GetQosOk() (*OsdQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *PoolRecord) SetQos(v OsdQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *PoolRecord) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReplicateSize

`func (o *PoolRecord) GetReplicateSize() int64`

GetReplicateSize returns the ReplicateSize field if non-nil, zero value otherwise.

### GetReplicateSizeOk

`func (o *PoolRecord) GetReplicateSizeOk() (*int64, bool)`

GetReplicateSizeOk returns a tuple with the ReplicateSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateSize

`func (o *PoolRecord) SetReplicateSize(v int64)`

SetReplicateSize sets ReplicateSize field to given value.

### HasReplicateSize

`func (o *PoolRecord) HasReplicateSize() bool`

HasReplicateSize returns a boolean if a field has been set.

### GetReservedPercent

`func (o *PoolRecord) GetReservedPercent() float64`

GetReservedPercent returns the ReservedPercent field if non-nil, zero value otherwise.

### GetReservedPercentOk

`func (o *PoolRecord) GetReservedPercentOk() (*float64, bool)`

GetReservedPercentOk returns a tuple with the ReservedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedPercent

`func (o *PoolRecord) SetReservedPercent(v float64)`

SetReservedPercent sets ReservedPercent field to given value.

### HasReservedPercent

`func (o *PoolRecord) HasReservedPercent() bool`

HasReservedPercent returns a boolean if a field has been set.

### GetSizePerOsd

`func (o *PoolRecord) GetSizePerOsd() int64`

GetSizePerOsd returns the SizePerOsd field if non-nil, zero value otherwise.

### GetSizePerOsdOk

`func (o *PoolRecord) GetSizePerOsdOk() (*int64, bool)`

GetSizePerOsdOk returns a tuple with the SizePerOsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizePerOsd

`func (o *PoolRecord) SetSizePerOsd(v int64)`

SetSizePerOsd sets SizePerOsd field to given value.

### HasSizePerOsd

`func (o *PoolRecord) HasSizePerOsd() bool`

HasSizePerOsd returns a boolean if a field has been set.

### GetStatus

`func (o *PoolRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoolRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoolRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PoolRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStretched

`func (o *PoolRecord) GetStretched() bool`

GetStretched returns the Stretched field if non-nil, zero value otherwise.

### GetStretchedOk

`func (o *PoolRecord) GetStretchedOk() (*bool, bool)`

GetStretchedOk returns a tuple with the Stretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretched

`func (o *PoolRecord) SetStretched(v bool)`

SetStretched sets Stretched field to given value.

### HasStretched

`func (o *PoolRecord) HasStretched() bool`

HasStretched returns a boolean if a field has been set.

### GetStripeUnit

`func (o *PoolRecord) GetStripeUnit() int64`

GetStripeUnit returns the StripeUnit field if non-nil, zero value otherwise.

### GetStripeUnitOk

`func (o *PoolRecord) GetStripeUnitOk() (*int64, bool)`

GetStripeUnitOk returns a tuple with the StripeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeUnit

`func (o *PoolRecord) SetStripeUnit(v int64)`

SetStripeUnit sets StripeUnit field to given value.

### HasStripeUnit

`func (o *PoolRecord) HasStripeUnit() bool`

HasStripeUnit returns a boolean if a field has been set.

### GetSubFailureDomainType

`func (o *PoolRecord) GetSubFailureDomainType() string`

GetSubFailureDomainType returns the SubFailureDomainType field if non-nil, zero value otherwise.

### GetSubFailureDomainTypeOk

`func (o *PoolRecord) GetSubFailureDomainTypeOk() (*string, bool)`

GetSubFailureDomainTypeOk returns a tuple with the SubFailureDomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFailureDomainType

`func (o *PoolRecord) SetSubFailureDomainType(v string)`

SetSubFailureDomainType sets SubFailureDomainType field to given value.

### HasSubFailureDomainType

`func (o *PoolRecord) HasSubFailureDomainType() bool`

HasSubFailureDomainType returns a boolean if a field has been set.

### GetSuggestedOmapSize

`func (o *PoolRecord) GetSuggestedOmapSize() int64`

GetSuggestedOmapSize returns the SuggestedOmapSize field if non-nil, zero value otherwise.

### GetSuggestedOmapSizeOk

`func (o *PoolRecord) GetSuggestedOmapSizeOk() (*int64, bool)`

GetSuggestedOmapSizeOk returns a tuple with the SuggestedOmapSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedOmapSize

`func (o *PoolRecord) SetSuggestedOmapSize(v int64)`

SetSuggestedOmapSize sets SuggestedOmapSize field to given value.

### HasSuggestedOmapSize

`func (o *PoolRecord) HasSuggestedOmapSize() bool`

HasSuggestedOmapSize returns a boolean if a field has been set.

### GetThinProvisioned

`func (o *PoolRecord) GetThinProvisioned() bool`

GetThinProvisioned returns the ThinProvisioned field if non-nil, zero value otherwise.

### GetThinProvisionedOk

`func (o *PoolRecord) GetThinProvisionedOk() (*bool, bool)`

GetThinProvisionedOk returns a tuple with the ThinProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioned

`func (o *PoolRecord) SetThinProvisioned(v bool)`

SetThinProvisioned sets ThinProvisioned field to given value.

### HasThinProvisioned

`func (o *PoolRecord) HasThinProvisioned() bool`

HasThinProvisioned returns a boolean if a field has been set.

### GetTierStatus

`func (o *PoolRecord) GetTierStatus() string`

GetTierStatus returns the TierStatus field if non-nil, zero value otherwise.

### GetTierStatusOk

`func (o *PoolRecord) GetTierStatusOk() (*string, bool)`

GetTierStatusOk returns a tuple with the TierStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierStatus

`func (o *PoolRecord) SetTierStatus(v string)`

SetTierStatus sets TierStatus field to given value.

### HasTierStatus

`func (o *PoolRecord) HasTierStatus() bool`

HasTierStatus returns a boolean if a field has been set.

### GetTransportMode

`func (o *PoolRecord) GetTransportMode() string`

GetTransportMode returns the TransportMode field if non-nil, zero value otherwise.

### GetTransportModeOk

`func (o *PoolRecord) GetTransportModeOk() (*string, bool)`

GetTransportModeOk returns a tuple with the TransportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMode

`func (o *PoolRecord) SetTransportMode(v string)`

SetTransportMode sets TransportMode field to given value.

### HasTransportMode

`func (o *PoolRecord) HasTransportMode() bool`

HasTransportMode returns a boolean if a field has been set.

### GetUpdate

`func (o *PoolRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PoolRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PoolRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PoolRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *PoolRecord) GetSamples() []PoolStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *PoolRecord) GetSamplesOk() (*[]PoolStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *PoolRecord) SetSamples(v []PoolStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *PoolRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


