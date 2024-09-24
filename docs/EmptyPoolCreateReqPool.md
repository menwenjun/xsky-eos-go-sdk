# EmptyPoolCreateReqPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**EmptyPoolCreateReqPoolCache**](EmptyPoolCreateReqPoolCache.md) |  | [optional] 
**ClusterId** | Pointer to **int64** |  | [optional] 
**CodingChunkNum** | Pointer to **int64** |  | [optional] 
**DataChunkNum** | Pointer to **int64** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**EncryptEnabled** | Pointer to **bool** |  | [optional] 
**FailureDomainType** | Pointer to **string** |  | [optional] 
**MinAllocSize** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumaBindPolicy** | Pointer to **string** |  | [optional] 
**NumaEnabled** | Pointer to **bool** |  | [optional] 
**OsdNumPerHost** | Pointer to **int64** |  | [optional] 
**OutFailureDomainNum** | Pointer to **int64** |  | [optional] 
**PoolName** | Pointer to **string** |  | [optional] 
**PoolRole** | Pointer to **string** |  | [optional] 
**PoolType** | Pointer to **string** |  | [optional] 
**ProductType** | Pointer to **string** |  | [optional] 
**ProtectionDomainId** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**SizePerOsd** | Pointer to **int64** |  | [optional] 
**Stretched** | Pointer to **bool** |  | [optional] 
**ThinProvisioned** | Pointer to **bool** |  | [optional] 
**TransportMode** | Pointer to **string** |  | [optional] 

## Methods

### NewEmptyPoolCreateReqPool

`func NewEmptyPoolCreateReqPool() *EmptyPoolCreateReqPool`

NewEmptyPoolCreateReqPool instantiates a new EmptyPoolCreateReqPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmptyPoolCreateReqPoolWithDefaults

`func NewEmptyPoolCreateReqPoolWithDefaults() *EmptyPoolCreateReqPool`

NewEmptyPoolCreateReqPoolWithDefaults instantiates a new EmptyPoolCreateReqPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *EmptyPoolCreateReqPool) GetCache() EmptyPoolCreateReqPoolCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *EmptyPoolCreateReqPool) GetCacheOk() (*EmptyPoolCreateReqPoolCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *EmptyPoolCreateReqPool) SetCache(v EmptyPoolCreateReqPoolCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *EmptyPoolCreateReqPool) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetClusterId

`func (o *EmptyPoolCreateReqPool) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EmptyPoolCreateReqPool) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EmptyPoolCreateReqPool) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EmptyPoolCreateReqPool) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCodingChunkNum

`func (o *EmptyPoolCreateReqPool) GetCodingChunkNum() int64`

GetCodingChunkNum returns the CodingChunkNum field if non-nil, zero value otherwise.

### GetCodingChunkNumOk

`func (o *EmptyPoolCreateReqPool) GetCodingChunkNumOk() (*int64, bool)`

GetCodingChunkNumOk returns a tuple with the CodingChunkNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodingChunkNum

`func (o *EmptyPoolCreateReqPool) SetCodingChunkNum(v int64)`

SetCodingChunkNum sets CodingChunkNum field to given value.

### HasCodingChunkNum

`func (o *EmptyPoolCreateReqPool) HasCodingChunkNum() bool`

HasCodingChunkNum returns a boolean if a field has been set.

### GetDataChunkNum

`func (o *EmptyPoolCreateReqPool) GetDataChunkNum() int64`

GetDataChunkNum returns the DataChunkNum field if non-nil, zero value otherwise.

### GetDataChunkNumOk

`func (o *EmptyPoolCreateReqPool) GetDataChunkNumOk() (*int64, bool)`

GetDataChunkNumOk returns a tuple with the DataChunkNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChunkNum

`func (o *EmptyPoolCreateReqPool) SetDataChunkNum(v int64)`

SetDataChunkNum sets DataChunkNum field to given value.

### HasDataChunkNum

`func (o *EmptyPoolCreateReqPool) HasDataChunkNum() bool`

HasDataChunkNum returns a boolean if a field has been set.

### GetDeviceType

`func (o *EmptyPoolCreateReqPool) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *EmptyPoolCreateReqPool) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *EmptyPoolCreateReqPool) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *EmptyPoolCreateReqPool) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetEncryptEnabled

`func (o *EmptyPoolCreateReqPool) GetEncryptEnabled() bool`

GetEncryptEnabled returns the EncryptEnabled field if non-nil, zero value otherwise.

### GetEncryptEnabledOk

`func (o *EmptyPoolCreateReqPool) GetEncryptEnabledOk() (*bool, bool)`

GetEncryptEnabledOk returns a tuple with the EncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEnabled

`func (o *EmptyPoolCreateReqPool) SetEncryptEnabled(v bool)`

SetEncryptEnabled sets EncryptEnabled field to given value.

### HasEncryptEnabled

`func (o *EmptyPoolCreateReqPool) HasEncryptEnabled() bool`

HasEncryptEnabled returns a boolean if a field has been set.

### GetFailureDomainType

`func (o *EmptyPoolCreateReqPool) GetFailureDomainType() string`

GetFailureDomainType returns the FailureDomainType field if non-nil, zero value otherwise.

### GetFailureDomainTypeOk

`func (o *EmptyPoolCreateReqPool) GetFailureDomainTypeOk() (*string, bool)`

GetFailureDomainTypeOk returns a tuple with the FailureDomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDomainType

`func (o *EmptyPoolCreateReqPool) SetFailureDomainType(v string)`

SetFailureDomainType sets FailureDomainType field to given value.

### HasFailureDomainType

`func (o *EmptyPoolCreateReqPool) HasFailureDomainType() bool`

HasFailureDomainType returns a boolean if a field has been set.

### GetMinAllocSize

`func (o *EmptyPoolCreateReqPool) GetMinAllocSize() int64`

GetMinAllocSize returns the MinAllocSize field if non-nil, zero value otherwise.

### GetMinAllocSizeOk

`func (o *EmptyPoolCreateReqPool) GetMinAllocSizeOk() (*int64, bool)`

GetMinAllocSizeOk returns a tuple with the MinAllocSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAllocSize

`func (o *EmptyPoolCreateReqPool) SetMinAllocSize(v int64)`

SetMinAllocSize sets MinAllocSize field to given value.

### HasMinAllocSize

`func (o *EmptyPoolCreateReqPool) HasMinAllocSize() bool`

HasMinAllocSize returns a boolean if a field has been set.

### GetName

`func (o *EmptyPoolCreateReqPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmptyPoolCreateReqPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmptyPoolCreateReqPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmptyPoolCreateReqPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumaBindPolicy

`func (o *EmptyPoolCreateReqPool) GetNumaBindPolicy() string`

GetNumaBindPolicy returns the NumaBindPolicy field if non-nil, zero value otherwise.

### GetNumaBindPolicyOk

`func (o *EmptyPoolCreateReqPool) GetNumaBindPolicyOk() (*string, bool)`

GetNumaBindPolicyOk returns a tuple with the NumaBindPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaBindPolicy

`func (o *EmptyPoolCreateReqPool) SetNumaBindPolicy(v string)`

SetNumaBindPolicy sets NumaBindPolicy field to given value.

### HasNumaBindPolicy

`func (o *EmptyPoolCreateReqPool) HasNumaBindPolicy() bool`

HasNumaBindPolicy returns a boolean if a field has been set.

### GetNumaEnabled

`func (o *EmptyPoolCreateReqPool) GetNumaEnabled() bool`

GetNumaEnabled returns the NumaEnabled field if non-nil, zero value otherwise.

### GetNumaEnabledOk

`func (o *EmptyPoolCreateReqPool) GetNumaEnabledOk() (*bool, bool)`

GetNumaEnabledOk returns a tuple with the NumaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaEnabled

`func (o *EmptyPoolCreateReqPool) SetNumaEnabled(v bool)`

SetNumaEnabled sets NumaEnabled field to given value.

### HasNumaEnabled

`func (o *EmptyPoolCreateReqPool) HasNumaEnabled() bool`

HasNumaEnabled returns a boolean if a field has been set.

### GetOsdNumPerHost

`func (o *EmptyPoolCreateReqPool) GetOsdNumPerHost() int64`

GetOsdNumPerHost returns the OsdNumPerHost field if non-nil, zero value otherwise.

### GetOsdNumPerHostOk

`func (o *EmptyPoolCreateReqPool) GetOsdNumPerHostOk() (*int64, bool)`

GetOsdNumPerHostOk returns a tuple with the OsdNumPerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdNumPerHost

`func (o *EmptyPoolCreateReqPool) SetOsdNumPerHost(v int64)`

SetOsdNumPerHost sets OsdNumPerHost field to given value.

### HasOsdNumPerHost

`func (o *EmptyPoolCreateReqPool) HasOsdNumPerHost() bool`

HasOsdNumPerHost returns a boolean if a field has been set.

### GetOutFailureDomainNum

`func (o *EmptyPoolCreateReqPool) GetOutFailureDomainNum() int64`

GetOutFailureDomainNum returns the OutFailureDomainNum field if non-nil, zero value otherwise.

### GetOutFailureDomainNumOk

`func (o *EmptyPoolCreateReqPool) GetOutFailureDomainNumOk() (*int64, bool)`

GetOutFailureDomainNumOk returns a tuple with the OutFailureDomainNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutFailureDomainNum

`func (o *EmptyPoolCreateReqPool) SetOutFailureDomainNum(v int64)`

SetOutFailureDomainNum sets OutFailureDomainNum field to given value.

### HasOutFailureDomainNum

`func (o *EmptyPoolCreateReqPool) HasOutFailureDomainNum() bool`

HasOutFailureDomainNum returns a boolean if a field has been set.

### GetPoolName

`func (o *EmptyPoolCreateReqPool) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *EmptyPoolCreateReqPool) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *EmptyPoolCreateReqPool) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *EmptyPoolCreateReqPool) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolRole

`func (o *EmptyPoolCreateReqPool) GetPoolRole() string`

GetPoolRole returns the PoolRole field if non-nil, zero value otherwise.

### GetPoolRoleOk

`func (o *EmptyPoolCreateReqPool) GetPoolRoleOk() (*string, bool)`

GetPoolRoleOk returns a tuple with the PoolRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolRole

`func (o *EmptyPoolCreateReqPool) SetPoolRole(v string)`

SetPoolRole sets PoolRole field to given value.

### HasPoolRole

`func (o *EmptyPoolCreateReqPool) HasPoolRole() bool`

HasPoolRole returns a boolean if a field has been set.

### GetPoolType

`func (o *EmptyPoolCreateReqPool) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *EmptyPoolCreateReqPool) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *EmptyPoolCreateReqPool) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *EmptyPoolCreateReqPool) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetProductType

`func (o *EmptyPoolCreateReqPool) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *EmptyPoolCreateReqPool) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *EmptyPoolCreateReqPool) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *EmptyPoolCreateReqPool) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProtectionDomainId

`func (o *EmptyPoolCreateReqPool) GetProtectionDomainId() int64`

GetProtectionDomainId returns the ProtectionDomainId field if non-nil, zero value otherwise.

### GetProtectionDomainIdOk

`func (o *EmptyPoolCreateReqPool) GetProtectionDomainIdOk() (*int64, bool)`

GetProtectionDomainIdOk returns a tuple with the ProtectionDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDomainId

`func (o *EmptyPoolCreateReqPool) SetProtectionDomainId(v int64)`

SetProtectionDomainId sets ProtectionDomainId field to given value.

### HasProtectionDomainId

`func (o *EmptyPoolCreateReqPool) HasProtectionDomainId() bool`

HasProtectionDomainId returns a boolean if a field has been set.

### GetSize

`func (o *EmptyPoolCreateReqPool) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EmptyPoolCreateReqPool) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EmptyPoolCreateReqPool) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *EmptyPoolCreateReqPool) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizePerOsd

`func (o *EmptyPoolCreateReqPool) GetSizePerOsd() int64`

GetSizePerOsd returns the SizePerOsd field if non-nil, zero value otherwise.

### GetSizePerOsdOk

`func (o *EmptyPoolCreateReqPool) GetSizePerOsdOk() (*int64, bool)`

GetSizePerOsdOk returns a tuple with the SizePerOsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizePerOsd

`func (o *EmptyPoolCreateReqPool) SetSizePerOsd(v int64)`

SetSizePerOsd sets SizePerOsd field to given value.

### HasSizePerOsd

`func (o *EmptyPoolCreateReqPool) HasSizePerOsd() bool`

HasSizePerOsd returns a boolean if a field has been set.

### GetStretched

`func (o *EmptyPoolCreateReqPool) GetStretched() bool`

GetStretched returns the Stretched field if non-nil, zero value otherwise.

### GetStretchedOk

`func (o *EmptyPoolCreateReqPool) GetStretchedOk() (*bool, bool)`

GetStretchedOk returns a tuple with the Stretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretched

`func (o *EmptyPoolCreateReqPool) SetStretched(v bool)`

SetStretched sets Stretched field to given value.

### HasStretched

`func (o *EmptyPoolCreateReqPool) HasStretched() bool`

HasStretched returns a boolean if a field has been set.

### GetThinProvisioned

`func (o *EmptyPoolCreateReqPool) GetThinProvisioned() bool`

GetThinProvisioned returns the ThinProvisioned field if non-nil, zero value otherwise.

### GetThinProvisionedOk

`func (o *EmptyPoolCreateReqPool) GetThinProvisionedOk() (*bool, bool)`

GetThinProvisionedOk returns a tuple with the ThinProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioned

`func (o *EmptyPoolCreateReqPool) SetThinProvisioned(v bool)`

SetThinProvisioned sets ThinProvisioned field to given value.

### HasThinProvisioned

`func (o *EmptyPoolCreateReqPool) HasThinProvisioned() bool`

HasThinProvisioned returns a boolean if a field has been set.

### GetTransportMode

`func (o *EmptyPoolCreateReqPool) GetTransportMode() string`

GetTransportMode returns the TransportMode field if non-nil, zero value otherwise.

### GetTransportModeOk

`func (o *EmptyPoolCreateReqPool) GetTransportModeOk() (*string, bool)`

GetTransportModeOk returns a tuple with the TransportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMode

`func (o *EmptyPoolCreateReqPool) SetTransportMode(v string)`

SetTransportMode sets TransportMode field to given value.

### HasTransportMode

`func (o *EmptyPoolCreateReqPool) HasTransportMode() bool`

HasTransportMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


