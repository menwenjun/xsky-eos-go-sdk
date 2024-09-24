# PoolUpdateReqPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressAlgorithm** | Pointer to **string** |  | [optional] 
**Compressed** | Pointer to **bool** |  | [optional] 
**DefaultManagedVolumeFormat** | Pointer to **int64** | default managed volume format: 128 or 129 | [optional] 
**FailureDomainType** | Pointer to **string** |  | [optional] 
**IoBypassEnabled** | Pointer to **bool** |  | [optional] 
**IoBypassMode** | Pointer to **string** |  | [optional] 
**IoBypassThreshold** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OutFailureDomainNum** | Pointer to **int64** | ec k+m:h&lt;-&gt;k+m | [optional] 
**PrimaryPlacementNodeId** | Pointer to **int64** |  | [optional] 
**Property** | Pointer to **map[string][]string** |  | [optional] 
**Qos** | Pointer to [**OsdQos**](OsdQos.md) |  | [optional] 
**Ruleset** | Pointer to [**[]PoolRuleReq**](PoolRuleReq.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**ThinProvisioned** | Pointer to **bool** |  | [optional] 

## Methods

### NewPoolUpdateReqPool

`func NewPoolUpdateReqPool() *PoolUpdateReqPool`

NewPoolUpdateReqPool instantiates a new PoolUpdateReqPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolUpdateReqPoolWithDefaults

`func NewPoolUpdateReqPoolWithDefaults() *PoolUpdateReqPool`

NewPoolUpdateReqPoolWithDefaults instantiates a new PoolUpdateReqPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressAlgorithm

`func (o *PoolUpdateReqPool) GetCompressAlgorithm() string`

GetCompressAlgorithm returns the CompressAlgorithm field if non-nil, zero value otherwise.

### GetCompressAlgorithmOk

`func (o *PoolUpdateReqPool) GetCompressAlgorithmOk() (*string, bool)`

GetCompressAlgorithmOk returns a tuple with the CompressAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressAlgorithm

`func (o *PoolUpdateReqPool) SetCompressAlgorithm(v string)`

SetCompressAlgorithm sets CompressAlgorithm field to given value.

### HasCompressAlgorithm

`func (o *PoolUpdateReqPool) HasCompressAlgorithm() bool`

HasCompressAlgorithm returns a boolean if a field has been set.

### GetCompressed

`func (o *PoolUpdateReqPool) GetCompressed() bool`

GetCompressed returns the Compressed field if non-nil, zero value otherwise.

### GetCompressedOk

`func (o *PoolUpdateReqPool) GetCompressedOk() (*bool, bool)`

GetCompressedOk returns a tuple with the Compressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressed

`func (o *PoolUpdateReqPool) SetCompressed(v bool)`

SetCompressed sets Compressed field to given value.

### HasCompressed

`func (o *PoolUpdateReqPool) HasCompressed() bool`

HasCompressed returns a boolean if a field has been set.

### GetDefaultManagedVolumeFormat

`func (o *PoolUpdateReqPool) GetDefaultManagedVolumeFormat() int64`

GetDefaultManagedVolumeFormat returns the DefaultManagedVolumeFormat field if non-nil, zero value otherwise.

### GetDefaultManagedVolumeFormatOk

`func (o *PoolUpdateReqPool) GetDefaultManagedVolumeFormatOk() (*int64, bool)`

GetDefaultManagedVolumeFormatOk returns a tuple with the DefaultManagedVolumeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultManagedVolumeFormat

`func (o *PoolUpdateReqPool) SetDefaultManagedVolumeFormat(v int64)`

SetDefaultManagedVolumeFormat sets DefaultManagedVolumeFormat field to given value.

### HasDefaultManagedVolumeFormat

`func (o *PoolUpdateReqPool) HasDefaultManagedVolumeFormat() bool`

HasDefaultManagedVolumeFormat returns a boolean if a field has been set.

### GetFailureDomainType

`func (o *PoolUpdateReqPool) GetFailureDomainType() string`

GetFailureDomainType returns the FailureDomainType field if non-nil, zero value otherwise.

### GetFailureDomainTypeOk

`func (o *PoolUpdateReqPool) GetFailureDomainTypeOk() (*string, bool)`

GetFailureDomainTypeOk returns a tuple with the FailureDomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDomainType

`func (o *PoolUpdateReqPool) SetFailureDomainType(v string)`

SetFailureDomainType sets FailureDomainType field to given value.

### HasFailureDomainType

`func (o *PoolUpdateReqPool) HasFailureDomainType() bool`

HasFailureDomainType returns a boolean if a field has been set.

### GetIoBypassEnabled

`func (o *PoolUpdateReqPool) GetIoBypassEnabled() bool`

GetIoBypassEnabled returns the IoBypassEnabled field if non-nil, zero value otherwise.

### GetIoBypassEnabledOk

`func (o *PoolUpdateReqPool) GetIoBypassEnabledOk() (*bool, bool)`

GetIoBypassEnabledOk returns a tuple with the IoBypassEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoBypassEnabled

`func (o *PoolUpdateReqPool) SetIoBypassEnabled(v bool)`

SetIoBypassEnabled sets IoBypassEnabled field to given value.

### HasIoBypassEnabled

`func (o *PoolUpdateReqPool) HasIoBypassEnabled() bool`

HasIoBypassEnabled returns a boolean if a field has been set.

### GetIoBypassMode

`func (o *PoolUpdateReqPool) GetIoBypassMode() string`

GetIoBypassMode returns the IoBypassMode field if non-nil, zero value otherwise.

### GetIoBypassModeOk

`func (o *PoolUpdateReqPool) GetIoBypassModeOk() (*string, bool)`

GetIoBypassModeOk returns a tuple with the IoBypassMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoBypassMode

`func (o *PoolUpdateReqPool) SetIoBypassMode(v string)`

SetIoBypassMode sets IoBypassMode field to given value.

### HasIoBypassMode

`func (o *PoolUpdateReqPool) HasIoBypassMode() bool`

HasIoBypassMode returns a boolean if a field has been set.

### GetIoBypassThreshold

`func (o *PoolUpdateReqPool) GetIoBypassThreshold() int64`

GetIoBypassThreshold returns the IoBypassThreshold field if non-nil, zero value otherwise.

### GetIoBypassThresholdOk

`func (o *PoolUpdateReqPool) GetIoBypassThresholdOk() (*int64, bool)`

GetIoBypassThresholdOk returns a tuple with the IoBypassThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoBypassThreshold

`func (o *PoolUpdateReqPool) SetIoBypassThreshold(v int64)`

SetIoBypassThreshold sets IoBypassThreshold field to given value.

### HasIoBypassThreshold

`func (o *PoolUpdateReqPool) HasIoBypassThreshold() bool`

HasIoBypassThreshold returns a boolean if a field has been set.

### GetName

`func (o *PoolUpdateReqPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolUpdateReqPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolUpdateReqPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolUpdateReqPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutFailureDomainNum

`func (o *PoolUpdateReqPool) GetOutFailureDomainNum() int64`

GetOutFailureDomainNum returns the OutFailureDomainNum field if non-nil, zero value otherwise.

### GetOutFailureDomainNumOk

`func (o *PoolUpdateReqPool) GetOutFailureDomainNumOk() (*int64, bool)`

GetOutFailureDomainNumOk returns a tuple with the OutFailureDomainNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutFailureDomainNum

`func (o *PoolUpdateReqPool) SetOutFailureDomainNum(v int64)`

SetOutFailureDomainNum sets OutFailureDomainNum field to given value.

### HasOutFailureDomainNum

`func (o *PoolUpdateReqPool) HasOutFailureDomainNum() bool`

HasOutFailureDomainNum returns a boolean if a field has been set.

### GetPrimaryPlacementNodeId

`func (o *PoolUpdateReqPool) GetPrimaryPlacementNodeId() int64`

GetPrimaryPlacementNodeId returns the PrimaryPlacementNodeId field if non-nil, zero value otherwise.

### GetPrimaryPlacementNodeIdOk

`func (o *PoolUpdateReqPool) GetPrimaryPlacementNodeIdOk() (*int64, bool)`

GetPrimaryPlacementNodeIdOk returns a tuple with the PrimaryPlacementNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPlacementNodeId

`func (o *PoolUpdateReqPool) SetPrimaryPlacementNodeId(v int64)`

SetPrimaryPlacementNodeId sets PrimaryPlacementNodeId field to given value.

### HasPrimaryPlacementNodeId

`func (o *PoolUpdateReqPool) HasPrimaryPlacementNodeId() bool`

HasPrimaryPlacementNodeId returns a boolean if a field has been set.

### GetProperty

`func (o *PoolUpdateReqPool) GetProperty() map[string][]string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PoolUpdateReqPool) GetPropertyOk() (*map[string][]string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PoolUpdateReqPool) SetProperty(v map[string][]string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *PoolUpdateReqPool) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetQos

`func (o *PoolUpdateReqPool) GetQos() OsdQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *PoolUpdateReqPool) GetQosOk() (*OsdQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *PoolUpdateReqPool) SetQos(v OsdQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *PoolUpdateReqPool) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetRuleset

`func (o *PoolUpdateReqPool) GetRuleset() []PoolRuleReq`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *PoolUpdateReqPool) GetRulesetOk() (*[]PoolRuleReq, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *PoolUpdateReqPool) SetRuleset(v []PoolRuleReq)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *PoolUpdateReqPool) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetSize

`func (o *PoolUpdateReqPool) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PoolUpdateReqPool) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PoolUpdateReqPool) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PoolUpdateReqPool) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetThinProvisioned

`func (o *PoolUpdateReqPool) GetThinProvisioned() bool`

GetThinProvisioned returns the ThinProvisioned field if non-nil, zero value otherwise.

### GetThinProvisionedOk

`func (o *PoolUpdateReqPool) GetThinProvisionedOk() (*bool, bool)`

GetThinProvisionedOk returns a tuple with the ThinProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioned

`func (o *PoolUpdateReqPool) SetThinProvisioned(v bool)`

SetThinProvisioned sets ThinProvisioned field to given value.

### HasThinProvisioned

`func (o *PoolUpdateReqPool) HasThinProvisioned() bool`

HasThinProvisioned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


