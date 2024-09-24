# VolumeUpdateReqVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**BlockSnapshotId** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Flattened** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PerformancePriority** | Pointer to **int64** |  | [optional] 
**Qos** | Pointer to [**VolumeQosSpec**](VolumeQosSpec.md) |  | [optional] 
**QosEnabled** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 

## Methods

### NewVolumeUpdateReqVolume

`func NewVolumeUpdateReqVolume() *VolumeUpdateReqVolume`

NewVolumeUpdateReqVolume instantiates a new VolumeUpdateReqVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeUpdateReqVolumeWithDefaults

`func NewVolumeUpdateReqVolumeWithDefaults() *VolumeUpdateReqVolume`

NewVolumeUpdateReqVolumeWithDefaults instantiates a new VolumeUpdateReqVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VolumeUpdateReqVolume) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VolumeUpdateReqVolume) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VolumeUpdateReqVolume) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *VolumeUpdateReqVolume) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBlockSnapshotId

`func (o *VolumeUpdateReqVolume) GetBlockSnapshotId() int64`

GetBlockSnapshotId returns the BlockSnapshotId field if non-nil, zero value otherwise.

### GetBlockSnapshotIdOk

`func (o *VolumeUpdateReqVolume) GetBlockSnapshotIdOk() (*int64, bool)`

GetBlockSnapshotIdOk returns a tuple with the BlockSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSnapshotId

`func (o *VolumeUpdateReqVolume) SetBlockSnapshotId(v int64)`

SetBlockSnapshotId sets BlockSnapshotId field to given value.

### HasBlockSnapshotId

`func (o *VolumeUpdateReqVolume) HasBlockSnapshotId() bool`

HasBlockSnapshotId returns a boolean if a field has been set.

### GetDescription

`func (o *VolumeUpdateReqVolume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeUpdateReqVolume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeUpdateReqVolume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeUpdateReqVolume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFlattened

`func (o *VolumeUpdateReqVolume) GetFlattened() bool`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *VolumeUpdateReqVolume) GetFlattenedOk() (*bool, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *VolumeUpdateReqVolume) SetFlattened(v bool)`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *VolumeUpdateReqVolume) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### GetName

`func (o *VolumeUpdateReqVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeUpdateReqVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeUpdateReqVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeUpdateReqVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPerformancePriority

`func (o *VolumeUpdateReqVolume) GetPerformancePriority() int64`

GetPerformancePriority returns the PerformancePriority field if non-nil, zero value otherwise.

### GetPerformancePriorityOk

`func (o *VolumeUpdateReqVolume) GetPerformancePriorityOk() (*int64, bool)`

GetPerformancePriorityOk returns a tuple with the PerformancePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformancePriority

`func (o *VolumeUpdateReqVolume) SetPerformancePriority(v int64)`

SetPerformancePriority sets PerformancePriority field to given value.

### HasPerformancePriority

`func (o *VolumeUpdateReqVolume) HasPerformancePriority() bool`

HasPerformancePriority returns a boolean if a field has been set.

### GetQos

`func (o *VolumeUpdateReqVolume) GetQos() VolumeQosSpec`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *VolumeUpdateReqVolume) GetQosOk() (*VolumeQosSpec, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *VolumeUpdateReqVolume) SetQos(v VolumeQosSpec)`

SetQos sets Qos field to given value.

### HasQos

`func (o *VolumeUpdateReqVolume) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetQosEnabled

`func (o *VolumeUpdateReqVolume) GetQosEnabled() bool`

GetQosEnabled returns the QosEnabled field if non-nil, zero value otherwise.

### GetQosEnabledOk

`func (o *VolumeUpdateReqVolume) GetQosEnabledOk() (*bool, bool)`

GetQosEnabledOk returns a tuple with the QosEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosEnabled

`func (o *VolumeUpdateReqVolume) SetQosEnabled(v bool)`

SetQosEnabled sets QosEnabled field to given value.

### HasQosEnabled

`func (o *VolumeUpdateReqVolume) HasQosEnabled() bool`

HasQosEnabled returns a boolean if a field has been set.

### GetSize

`func (o *VolumeUpdateReqVolume) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeUpdateReqVolume) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeUpdateReqVolume) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumeUpdateReqVolume) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


