# VolumeDpBlockBackupPolicyMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpBlockBackupPolicy** | Pointer to [**DpBlockBackupPolicyNestview**](DpBlockBackupPolicyNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 

## Methods

### NewVolumeDpBlockBackupPolicyMapping

`func NewVolumeDpBlockBackupPolicyMapping() *VolumeDpBlockBackupPolicyMapping`

NewVolumeDpBlockBackupPolicyMapping instantiates a new VolumeDpBlockBackupPolicyMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeDpBlockBackupPolicyMappingWithDefaults

`func NewVolumeDpBlockBackupPolicyMappingWithDefaults() *VolumeDpBlockBackupPolicyMapping`

NewVolumeDpBlockBackupPolicyMappingWithDefaults instantiates a new VolumeDpBlockBackupPolicyMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpBlockBackupPolicy

`func (o *VolumeDpBlockBackupPolicyMapping) GetDpBlockBackupPolicy() DpBlockBackupPolicyNestview`

GetDpBlockBackupPolicy returns the DpBlockBackupPolicy field if non-nil, zero value otherwise.

### GetDpBlockBackupPolicyOk

`func (o *VolumeDpBlockBackupPolicyMapping) GetDpBlockBackupPolicyOk() (*DpBlockBackupPolicyNestview, bool)`

GetDpBlockBackupPolicyOk returns a tuple with the DpBlockBackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpBlockBackupPolicy

`func (o *VolumeDpBlockBackupPolicyMapping) SetDpBlockBackupPolicy(v DpBlockBackupPolicyNestview)`

SetDpBlockBackupPolicy sets DpBlockBackupPolicy field to given value.

### HasDpBlockBackupPolicy

`func (o *VolumeDpBlockBackupPolicyMapping) HasDpBlockBackupPolicy() bool`

HasDpBlockBackupPolicy returns a boolean if a field has been set.

### GetId

`func (o *VolumeDpBlockBackupPolicyMapping) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeDpBlockBackupPolicyMapping) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeDpBlockBackupPolicyMapping) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeDpBlockBackupPolicyMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *VolumeDpBlockBackupPolicyMapping) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeDpBlockBackupPolicyMapping) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeDpBlockBackupPolicyMapping) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeDpBlockBackupPolicyMapping) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolume

`func (o *VolumeDpBlockBackupPolicyMapping) GetVolume() VolumeNestview`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *VolumeDpBlockBackupPolicyMapping) GetVolumeOk() (*VolumeNestview, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *VolumeDpBlockBackupPolicyMapping) SetVolume(v VolumeNestview)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *VolumeDpBlockBackupPolicyMapping) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


