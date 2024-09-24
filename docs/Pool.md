# Pool

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

## Methods

### NewPool

`func NewPool() *Pool`

NewPool instantiates a new Pool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolWithDefaults

`func NewPoolWithDefaults() *Pool`

NewPoolWithDefaults instantiates a new Pool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *Pool) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *Pool) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *Pool) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *Pool) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *Pool) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *Pool) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *Pool) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *Pool) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetBindOsdNum

`func (o *Pool) GetBindOsdNum() int64`

GetBindOsdNum returns the BindOsdNum field if non-nil, zero value otherwise.

### GetBindOsdNumOk

`func (o *Pool) GetBindOsdNumOk() (*int64, bool)`

GetBindOsdNumOk returns a tuple with the BindOsdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindOsdNum

`func (o *Pool) SetBindOsdNum(v int64)`

SetBindOsdNum sets BindOsdNum field to given value.

### HasBindOsdNum

`func (o *Pool) HasBindOsdNum() bool`

HasBindOsdNum returns a boolean if a field has been set.

### GetBlockVolumeNum

`func (o *Pool) GetBlockVolumeNum() int64`

GetBlockVolumeNum returns the BlockVolumeNum field if non-nil, zero value otherwise.

### GetBlockVolumeNumOk

`func (o *Pool) GetBlockVolumeNumOk() (*int64, bool)`

GetBlockVolumeNumOk returns a tuple with the BlockVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeNum

`func (o *Pool) SetBlockVolumeNum(v int64)`

SetBlockVolumeNum sets BlockVolumeNum field to given value.

### HasBlockVolumeNum

`func (o *Pool) HasBlockVolumeNum() bool`

HasBlockVolumeNum returns a boolean if a field has been set.

### GetCachePool

`func (o *Pool) GetCachePool() PoolNestview`

GetCachePool returns the CachePool field if non-nil, zero value otherwise.

### GetCachePoolOk

`func (o *Pool) GetCachePoolOk() (*PoolNestview, bool)`

GetCachePoolOk returns a tuple with the CachePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePool

`func (o *Pool) SetCachePool(v PoolNestview)`

SetCachePool sets CachePool field to given value.

### HasCachePool

`func (o *Pool) HasCachePool() bool`

HasCachePool returns a boolean if a field has been set.

### GetCluster

`func (o *Pool) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Pool) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Pool) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Pool) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCodingChunkNum

`func (o *Pool) GetCodingChunkNum() int64`

GetCodingChunkNum returns the CodingChunkNum field if non-nil, zero value otherwise.

### GetCodingChunkNumOk

`func (o *Pool) GetCodingChunkNumOk() (*int64, bool)`

GetCodingChunkNumOk returns a tuple with the CodingChunkNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodingChunkNum

`func (o *Pool) SetCodingChunkNum(v int64)`

SetCodingChunkNum sets CodingChunkNum field to given value.

### HasCodingChunkNum

`func (o *Pool) HasCodingChunkNum() bool`

HasCodingChunkNum returns a boolean if a field has been set.

### GetCompressAlgorithm

`func (o *Pool) GetCompressAlgorithm() string`

GetCompressAlgorithm returns the CompressAlgorithm field if non-nil, zero value otherwise.

### GetCompressAlgorithmOk

`func (o *Pool) GetCompressAlgorithmOk() (*string, bool)`

GetCompressAlgorithmOk returns a tuple with the CompressAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressAlgorithm

`func (o *Pool) SetCompressAlgorithm(v string)`

SetCompressAlgorithm sets CompressAlgorithm field to given value.

### HasCompressAlgorithm

`func (o *Pool) HasCompressAlgorithm() bool`

HasCompressAlgorithm returns a boolean if a field has been set.

### GetCompressed

`func (o *Pool) GetCompressed() bool`

GetCompressed returns the Compressed field if non-nil, zero value otherwise.

### GetCompressedOk

`func (o *Pool) GetCompressedOk() (*bool, bool)`

GetCompressedOk returns a tuple with the Compressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressed

`func (o *Pool) SetCompressed(v bool)`

SetCompressed sets Compressed field to given value.

### HasCompressed

`func (o *Pool) HasCompressed() bool`

HasCompressed returns a boolean if a field has been set.

### GetCreate

`func (o *Pool) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Pool) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Pool) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Pool) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataChunkNum

`func (o *Pool) GetDataChunkNum() int64`

GetDataChunkNum returns the DataChunkNum field if non-nil, zero value otherwise.

### GetDataChunkNumOk

`func (o *Pool) GetDataChunkNumOk() (*int64, bool)`

GetDataChunkNumOk returns a tuple with the DataChunkNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChunkNum

`func (o *Pool) SetDataChunkNum(v int64)`

SetDataChunkNum sets DataChunkNum field to given value.

### HasDataChunkNum

`func (o *Pool) HasDataChunkNum() bool`

HasDataChunkNum returns a boolean if a field has been set.

### GetDataPool

`func (o *Pool) GetDataPool() PoolNestview`

GetDataPool returns the DataPool field if non-nil, zero value otherwise.

### GetDataPoolOk

`func (o *Pool) GetDataPoolOk() (*PoolNestview, bool)`

GetDataPoolOk returns a tuple with the DataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPool

`func (o *Pool) SetDataPool(v PoolNestview)`

SetDataPool sets DataPool field to given value.

### HasDataPool

`func (o *Pool) HasDataPool() bool`

HasDataPool returns a boolean if a field has been set.

### GetDefaultManagedVolumeFormat

`func (o *Pool) GetDefaultManagedVolumeFormat() int64`

GetDefaultManagedVolumeFormat returns the DefaultManagedVolumeFormat field if non-nil, zero value otherwise.

### GetDefaultManagedVolumeFormatOk

`func (o *Pool) GetDefaultManagedVolumeFormatOk() (*int64, bool)`

GetDefaultManagedVolumeFormatOk returns a tuple with the DefaultManagedVolumeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultManagedVolumeFormat

`func (o *Pool) SetDefaultManagedVolumeFormat(v int64)`

SetDefaultManagedVolumeFormat sets DefaultManagedVolumeFormat field to given value.

### HasDefaultManagedVolumeFormat

`func (o *Pool) HasDefaultManagedVolumeFormat() bool`

HasDefaultManagedVolumeFormat returns a boolean if a field has been set.

### GetDeviceType

`func (o *Pool) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Pool) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Pool) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *Pool) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceTypeCheckDisabled

`func (o *Pool) GetDeviceTypeCheckDisabled() bool`

GetDeviceTypeCheckDisabled returns the DeviceTypeCheckDisabled field if non-nil, zero value otherwise.

### GetDeviceTypeCheckDisabledOk

`func (o *Pool) GetDeviceTypeCheckDisabledOk() (*bool, bool)`

GetDeviceTypeCheckDisabledOk returns a tuple with the DeviceTypeCheckDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCheckDisabled

`func (o *Pool) SetDeviceTypeCheckDisabled(v bool)`

SetDeviceTypeCheckDisabled sets DeviceTypeCheckDisabled field to given value.

### HasDeviceTypeCheckDisabled

`func (o *Pool) HasDeviceTypeCheckDisabled() bool`

HasDeviceTypeCheckDisabled returns a boolean if a field has been set.

### GetEncryptEnabled

`func (o *Pool) GetEncryptEnabled() bool`

GetEncryptEnabled returns the EncryptEnabled field if non-nil, zero value otherwise.

### GetEncryptEnabledOk

`func (o *Pool) GetEncryptEnabledOk() (*bool, bool)`

GetEncryptEnabledOk returns a tuple with the EncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEnabled

`func (o *Pool) SetEncryptEnabled(v bool)`

SetEncryptEnabled sets EncryptEnabled field to given value.

### HasEncryptEnabled

`func (o *Pool) HasEncryptEnabled() bool`

HasEncryptEnabled returns a boolean if a field has been set.

### GetFailureDomainType

`func (o *Pool) GetFailureDomainType() string`

GetFailureDomainType returns the FailureDomainType field if non-nil, zero value otherwise.

### GetFailureDomainTypeOk

`func (o *Pool) GetFailureDomainTypeOk() (*string, bool)`

GetFailureDomainTypeOk returns a tuple with the FailureDomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDomainType

`func (o *Pool) SetFailureDomainType(v string)`

SetFailureDomainType sets FailureDomainType field to given value.

### HasFailureDomainType

`func (o *Pool) HasFailureDomainType() bool`

HasFailureDomainType returns a boolean if a field has been set.

### GetGcQosPlan

`func (o *Pool) GetGcQosPlan() PoolGCPolicyPlan`

GetGcQosPlan returns the GcQosPlan field if non-nil, zero value otherwise.

### GetGcQosPlanOk

`func (o *Pool) GetGcQosPlanOk() (*PoolGCPolicyPlan, bool)`

GetGcQosPlanOk returns a tuple with the GcQosPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcQosPlan

`func (o *Pool) SetGcQosPlan(v PoolGCPolicyPlan)`

SetGcQosPlan sets GcQosPlan field to given value.

### HasGcQosPlan

`func (o *Pool) HasGcQosPlan() bool`

HasGcQosPlan returns a boolean if a field has been set.

### GetHidden

`func (o *Pool) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Pool) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Pool) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Pool) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *Pool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Pool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInnerPoolId

`func (o *Pool) GetInnerPoolId() int32`

GetInnerPoolId returns the InnerPoolId field if non-nil, zero value otherwise.

### GetInnerPoolIdOk

`func (o *Pool) GetInnerPoolIdOk() (*int32, bool)`

GetInnerPoolIdOk returns a tuple with the InnerPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerPoolId

`func (o *Pool) SetInnerPoolId(v int32)`

SetInnerPoolId sets InnerPoolId field to given value.

### HasInnerPoolId

`func (o *Pool) HasInnerPoolId() bool`

HasInnerPoolId returns a boolean if a field has been set.

### GetIoBypassEnabled

`func (o *Pool) GetIoBypassEnabled() bool`

GetIoBypassEnabled returns the IoBypassEnabled field if non-nil, zero value otherwise.

### GetIoBypassEnabledOk

`func (o *Pool) GetIoBypassEnabledOk() (*bool, bool)`

GetIoBypassEnabledOk returns a tuple with the IoBypassEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoBypassEnabled

`func (o *Pool) SetIoBypassEnabled(v bool)`

SetIoBypassEnabled sets IoBypassEnabled field to given value.

### HasIoBypassEnabled

`func (o *Pool) HasIoBypassEnabled() bool`

HasIoBypassEnabled returns a boolean if a field has been set.

### GetIoBypassMode

`func (o *Pool) GetIoBypassMode() string`

GetIoBypassMode returns the IoBypassMode field if non-nil, zero value otherwise.

### GetIoBypassModeOk

`func (o *Pool) GetIoBypassModeOk() (*string, bool)`

GetIoBypassModeOk returns a tuple with the IoBypassMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoBypassMode

`func (o *Pool) SetIoBypassMode(v string)`

SetIoBypassMode sets IoBypassMode field to given value.

### HasIoBypassMode

`func (o *Pool) HasIoBypassMode() bool`

HasIoBypassMode returns a boolean if a field has been set.

### GetIoBypassThreshold

`func (o *Pool) GetIoBypassThreshold() int64`

GetIoBypassThreshold returns the IoBypassThreshold field if non-nil, zero value otherwise.

### GetIoBypassThresholdOk

`func (o *Pool) GetIoBypassThresholdOk() (*int64, bool)`

GetIoBypassThresholdOk returns a tuple with the IoBypassThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoBypassThreshold

`func (o *Pool) SetIoBypassThreshold(v int64)`

SetIoBypassThreshold sets IoBypassThreshold field to given value.

### HasIoBypassThreshold

`func (o *Pool) HasIoBypassThreshold() bool`

HasIoBypassThreshold returns a boolean if a field has been set.

### GetMinAllocSize

`func (o *Pool) GetMinAllocSize() int64`

GetMinAllocSize returns the MinAllocSize field if non-nil, zero value otherwise.

### GetMinAllocSizeOk

`func (o *Pool) GetMinAllocSizeOk() (*int64, bool)`

GetMinAllocSizeOk returns a tuple with the MinAllocSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAllocSize

`func (o *Pool) SetMinAllocSize(v int64)`

SetMinAllocSize sets MinAllocSize field to given value.

### HasMinAllocSize

`func (o *Pool) HasMinAllocSize() bool`

HasMinAllocSize returns a boolean if a field has been set.

### GetMinOsdNum

`func (o *Pool) GetMinOsdNum() int64`

GetMinOsdNum returns the MinOsdNum field if non-nil, zero value otherwise.

### GetMinOsdNumOk

`func (o *Pool) GetMinOsdNumOk() (*int64, bool)`

GetMinOsdNumOk returns a tuple with the MinOsdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOsdNum

`func (o *Pool) SetMinOsdNum(v int64)`

SetMinOsdNum sets MinOsdNum field to given value.

### HasMinOsdNum

`func (o *Pool) HasMinOsdNum() bool`

HasMinOsdNum returns a boolean if a field has been set.

### GetName

`func (o *Pool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Pool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumaApplyPolicy

`func (o *Pool) GetNumaApplyPolicy() string`

GetNumaApplyPolicy returns the NumaApplyPolicy field if non-nil, zero value otherwise.

### GetNumaApplyPolicyOk

`func (o *Pool) GetNumaApplyPolicyOk() (*string, bool)`

GetNumaApplyPolicyOk returns a tuple with the NumaApplyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaApplyPolicy

`func (o *Pool) SetNumaApplyPolicy(v string)`

SetNumaApplyPolicy sets NumaApplyPolicy field to given value.

### HasNumaApplyPolicy

`func (o *Pool) HasNumaApplyPolicy() bool`

HasNumaApplyPolicy returns a boolean if a field has been set.

### GetNumaBindBalanced

`func (o *Pool) GetNumaBindBalanced() bool`

GetNumaBindBalanced returns the NumaBindBalanced field if non-nil, zero value otherwise.

### GetNumaBindBalancedOk

`func (o *Pool) GetNumaBindBalancedOk() (*bool, bool)`

GetNumaBindBalancedOk returns a tuple with the NumaBindBalanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaBindBalanced

`func (o *Pool) SetNumaBindBalanced(v bool)`

SetNumaBindBalanced sets NumaBindBalanced field to given value.

### HasNumaBindBalanced

`func (o *Pool) HasNumaBindBalanced() bool`

HasNumaBindBalanced returns a boolean if a field has been set.

### GetNumaBindPolicy

`func (o *Pool) GetNumaBindPolicy() string`

GetNumaBindPolicy returns the NumaBindPolicy field if non-nil, zero value otherwise.

### GetNumaBindPolicyOk

`func (o *Pool) GetNumaBindPolicyOk() (*string, bool)`

GetNumaBindPolicyOk returns a tuple with the NumaBindPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaBindPolicy

`func (o *Pool) SetNumaBindPolicy(v string)`

SetNumaBindPolicy sets NumaBindPolicy field to given value.

### HasNumaBindPolicy

`func (o *Pool) HasNumaBindPolicy() bool`

HasNumaBindPolicy returns a boolean if a field has been set.

### GetNumaEnabled

`func (o *Pool) GetNumaEnabled() bool`

GetNumaEnabled returns the NumaEnabled field if non-nil, zero value otherwise.

### GetNumaEnabledOk

`func (o *Pool) GetNumaEnabledOk() (*bool, bool)`

GetNumaEnabledOk returns a tuple with the NumaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaEnabled

`func (o *Pool) SetNumaEnabled(v bool)`

SetNumaEnabled sets NumaEnabled field to given value.

### HasNumaEnabled

`func (o *Pool) HasNumaEnabled() bool`

HasNumaEnabled returns a boolean if a field has been set.

### GetOsdGroup

`func (o *Pool) GetOsdGroup() OsdGroupNestview`

GetOsdGroup returns the OsdGroup field if non-nil, zero value otherwise.

### GetOsdGroupOk

`func (o *Pool) GetOsdGroupOk() (*OsdGroupNestview, bool)`

GetOsdGroupOk returns a tuple with the OsdGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdGroup

`func (o *Pool) SetOsdGroup(v OsdGroupNestview)`

SetOsdGroup sets OsdGroup field to given value.

### HasOsdGroup

`func (o *Pool) HasOsdGroup() bool`

HasOsdGroup returns a boolean if a field has been set.

### GetOsdNum

`func (o *Pool) GetOsdNum() int64`

GetOsdNum returns the OsdNum field if non-nil, zero value otherwise.

### GetOsdNumOk

`func (o *Pool) GetOsdNumOk() (*int64, bool)`

GetOsdNumOk returns a tuple with the OsdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdNum

`func (o *Pool) SetOsdNum(v int64)`

SetOsdNum sets OsdNum field to given value.

### HasOsdNum

`func (o *Pool) HasOsdNum() bool`

HasOsdNum returns a boolean if a field has been set.

### GetOsdNumPerHost

`func (o *Pool) GetOsdNumPerHost() int64`

GetOsdNumPerHost returns the OsdNumPerHost field if non-nil, zero value otherwise.

### GetOsdNumPerHostOk

`func (o *Pool) GetOsdNumPerHostOk() (*int64, bool)`

GetOsdNumPerHostOk returns a tuple with the OsdNumPerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdNumPerHost

`func (o *Pool) SetOsdNumPerHost(v int64)`

SetOsdNumPerHost sets OsdNumPerHost field to given value.

### HasOsdNumPerHost

`func (o *Pool) HasOsdNumPerHost() bool`

HasOsdNumPerHost returns a boolean if a field has been set.

### GetOutFailureDomainNum

`func (o *Pool) GetOutFailureDomainNum() int64`

GetOutFailureDomainNum returns the OutFailureDomainNum field if non-nil, zero value otherwise.

### GetOutFailureDomainNumOk

`func (o *Pool) GetOutFailureDomainNumOk() (*int64, bool)`

GetOutFailureDomainNumOk returns a tuple with the OutFailureDomainNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutFailureDomainNum

`func (o *Pool) SetOutFailureDomainNum(v int64)`

SetOutFailureDomainNum sets OutFailureDomainNum field to given value.

### HasOutFailureDomainNum

`func (o *Pool) HasOutFailureDomainNum() bool`

HasOutFailureDomainNum returns a boolean if a field has been set.

### GetPoolId

`func (o *Pool) GetPoolId() int32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *Pool) GetPoolIdOk() (*int32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *Pool) SetPoolId(v int32)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *Pool) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolMode

`func (o *Pool) GetPoolMode() string`

GetPoolMode returns the PoolMode field if non-nil, zero value otherwise.

### GetPoolModeOk

`func (o *Pool) GetPoolModeOk() (*string, bool)`

GetPoolModeOk returns a tuple with the PoolMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMode

`func (o *Pool) SetPoolMode(v string)`

SetPoolMode sets PoolMode field to given value.

### HasPoolMode

`func (o *Pool) HasPoolMode() bool`

HasPoolMode returns a boolean if a field has been set.

### GetPoolName

`func (o *Pool) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *Pool) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *Pool) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *Pool) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolRole

`func (o *Pool) GetPoolRole() string`

GetPoolRole returns the PoolRole field if non-nil, zero value otherwise.

### GetPoolRoleOk

`func (o *Pool) GetPoolRoleOk() (*string, bool)`

GetPoolRoleOk returns a tuple with the PoolRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolRole

`func (o *Pool) SetPoolRole(v string)`

SetPoolRole sets PoolRole field to given value.

### HasPoolRole

`func (o *Pool) HasPoolRole() bool`

HasPoolRole returns a boolean if a field has been set.

### GetPoolType

`func (o *Pool) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *Pool) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *Pool) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *Pool) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetPrimaryPlacementNode

`func (o *Pool) GetPrimaryPlacementNode() PlacementNodeNestview`

GetPrimaryPlacementNode returns the PrimaryPlacementNode field if non-nil, zero value otherwise.

### GetPrimaryPlacementNodeOk

`func (o *Pool) GetPrimaryPlacementNodeOk() (*PlacementNodeNestview, bool)`

GetPrimaryPlacementNodeOk returns a tuple with the PrimaryPlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPlacementNode

`func (o *Pool) SetPrimaryPlacementNode(v PlacementNodeNestview)`

SetPrimaryPlacementNode sets PrimaryPlacementNode field to given value.

### HasPrimaryPlacementNode

`func (o *Pool) HasPrimaryPlacementNode() bool`

HasPrimaryPlacementNode returns a boolean if a field has been set.

### GetProductType

`func (o *Pool) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Pool) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Pool) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *Pool) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProperty

`func (o *Pool) GetProperty() map[string]map[string]interface{}`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *Pool) GetPropertyOk() (*map[string]map[string]interface{}, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *Pool) SetProperty(v map[string]map[string]interface{})`

SetProperty sets Property field to given value.

### HasProperty

`func (o *Pool) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetProtectionDomain

`func (o *Pool) GetProtectionDomain() ProtectionDomainNestview`

GetProtectionDomain returns the ProtectionDomain field if non-nil, zero value otherwise.

### GetProtectionDomainOk

`func (o *Pool) GetProtectionDomainOk() (*ProtectionDomainNestview, bool)`

GetProtectionDomainOk returns a tuple with the ProtectionDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDomain

`func (o *Pool) SetProtectionDomain(v ProtectionDomainNestview)`

SetProtectionDomain sets ProtectionDomain field to given value.

### HasProtectionDomain

`func (o *Pool) HasProtectionDomain() bool`

HasProtectionDomain returns a boolean if a field has been set.

### GetQos

`func (o *Pool) GetQos() OsdQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Pool) GetQosOk() (*OsdQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Pool) SetQos(v OsdQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Pool) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReplicateSize

`func (o *Pool) GetReplicateSize() int64`

GetReplicateSize returns the ReplicateSize field if non-nil, zero value otherwise.

### GetReplicateSizeOk

`func (o *Pool) GetReplicateSizeOk() (*int64, bool)`

GetReplicateSizeOk returns a tuple with the ReplicateSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateSize

`func (o *Pool) SetReplicateSize(v int64)`

SetReplicateSize sets ReplicateSize field to given value.

### HasReplicateSize

`func (o *Pool) HasReplicateSize() bool`

HasReplicateSize returns a boolean if a field has been set.

### GetReservedPercent

`func (o *Pool) GetReservedPercent() float64`

GetReservedPercent returns the ReservedPercent field if non-nil, zero value otherwise.

### GetReservedPercentOk

`func (o *Pool) GetReservedPercentOk() (*float64, bool)`

GetReservedPercentOk returns a tuple with the ReservedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedPercent

`func (o *Pool) SetReservedPercent(v float64)`

SetReservedPercent sets ReservedPercent field to given value.

### HasReservedPercent

`func (o *Pool) HasReservedPercent() bool`

HasReservedPercent returns a boolean if a field has been set.

### GetSizePerOsd

`func (o *Pool) GetSizePerOsd() int64`

GetSizePerOsd returns the SizePerOsd field if non-nil, zero value otherwise.

### GetSizePerOsdOk

`func (o *Pool) GetSizePerOsdOk() (*int64, bool)`

GetSizePerOsdOk returns a tuple with the SizePerOsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizePerOsd

`func (o *Pool) SetSizePerOsd(v int64)`

SetSizePerOsd sets SizePerOsd field to given value.

### HasSizePerOsd

`func (o *Pool) HasSizePerOsd() bool`

HasSizePerOsd returns a boolean if a field has been set.

### GetStatus

`func (o *Pool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Pool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Pool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Pool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStretched

`func (o *Pool) GetStretched() bool`

GetStretched returns the Stretched field if non-nil, zero value otherwise.

### GetStretchedOk

`func (o *Pool) GetStretchedOk() (*bool, bool)`

GetStretchedOk returns a tuple with the Stretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretched

`func (o *Pool) SetStretched(v bool)`

SetStretched sets Stretched field to given value.

### HasStretched

`func (o *Pool) HasStretched() bool`

HasStretched returns a boolean if a field has been set.

### GetStripeUnit

`func (o *Pool) GetStripeUnit() int64`

GetStripeUnit returns the StripeUnit field if non-nil, zero value otherwise.

### GetStripeUnitOk

`func (o *Pool) GetStripeUnitOk() (*int64, bool)`

GetStripeUnitOk returns a tuple with the StripeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeUnit

`func (o *Pool) SetStripeUnit(v int64)`

SetStripeUnit sets StripeUnit field to given value.

### HasStripeUnit

`func (o *Pool) HasStripeUnit() bool`

HasStripeUnit returns a boolean if a field has been set.

### GetSubFailureDomainType

`func (o *Pool) GetSubFailureDomainType() string`

GetSubFailureDomainType returns the SubFailureDomainType field if non-nil, zero value otherwise.

### GetSubFailureDomainTypeOk

`func (o *Pool) GetSubFailureDomainTypeOk() (*string, bool)`

GetSubFailureDomainTypeOk returns a tuple with the SubFailureDomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFailureDomainType

`func (o *Pool) SetSubFailureDomainType(v string)`

SetSubFailureDomainType sets SubFailureDomainType field to given value.

### HasSubFailureDomainType

`func (o *Pool) HasSubFailureDomainType() bool`

HasSubFailureDomainType returns a boolean if a field has been set.

### GetSuggestedOmapSize

`func (o *Pool) GetSuggestedOmapSize() int64`

GetSuggestedOmapSize returns the SuggestedOmapSize field if non-nil, zero value otherwise.

### GetSuggestedOmapSizeOk

`func (o *Pool) GetSuggestedOmapSizeOk() (*int64, bool)`

GetSuggestedOmapSizeOk returns a tuple with the SuggestedOmapSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedOmapSize

`func (o *Pool) SetSuggestedOmapSize(v int64)`

SetSuggestedOmapSize sets SuggestedOmapSize field to given value.

### HasSuggestedOmapSize

`func (o *Pool) HasSuggestedOmapSize() bool`

HasSuggestedOmapSize returns a boolean if a field has been set.

### GetThinProvisioned

`func (o *Pool) GetThinProvisioned() bool`

GetThinProvisioned returns the ThinProvisioned field if non-nil, zero value otherwise.

### GetThinProvisionedOk

`func (o *Pool) GetThinProvisionedOk() (*bool, bool)`

GetThinProvisionedOk returns a tuple with the ThinProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioned

`func (o *Pool) SetThinProvisioned(v bool)`

SetThinProvisioned sets ThinProvisioned field to given value.

### HasThinProvisioned

`func (o *Pool) HasThinProvisioned() bool`

HasThinProvisioned returns a boolean if a field has been set.

### GetTierStatus

`func (o *Pool) GetTierStatus() string`

GetTierStatus returns the TierStatus field if non-nil, zero value otherwise.

### GetTierStatusOk

`func (o *Pool) GetTierStatusOk() (*string, bool)`

GetTierStatusOk returns a tuple with the TierStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierStatus

`func (o *Pool) SetTierStatus(v string)`

SetTierStatus sets TierStatus field to given value.

### HasTierStatus

`func (o *Pool) HasTierStatus() bool`

HasTierStatus returns a boolean if a field has been set.

### GetTransportMode

`func (o *Pool) GetTransportMode() string`

GetTransportMode returns the TransportMode field if non-nil, zero value otherwise.

### GetTransportModeOk

`func (o *Pool) GetTransportModeOk() (*string, bool)`

GetTransportModeOk returns a tuple with the TransportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMode

`func (o *Pool) SetTransportMode(v string)`

SetTransportMode sets TransportMode field to given value.

### HasTransportMode

`func (o *Pool) HasTransportMode() bool`

HasTransportMode returns a boolean if a field has been set.

### GetUpdate

`func (o *Pool) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Pool) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Pool) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Pool) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


