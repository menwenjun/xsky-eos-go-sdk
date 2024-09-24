# VolumeBackupProtectionReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolume** | Pointer to [**VolumeBackupProtectionReqVolume**](VolumeBackupProtectionReqVolume.md) |  | [optional] 

## Methods

### NewVolumeBackupProtectionReq

`func NewVolumeBackupProtectionReq() *VolumeBackupProtectionReq`

NewVolumeBackupProtectionReq instantiates a new VolumeBackupProtectionReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeBackupProtectionReqWithDefaults

`func NewVolumeBackupProtectionReqWithDefaults() *VolumeBackupProtectionReq`

NewVolumeBackupProtectionReqWithDefaults instantiates a new VolumeBackupProtectionReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolume

`func (o *VolumeBackupProtectionReq) GetBlockVolume() VolumeBackupProtectionReqVolume`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *VolumeBackupProtectionReq) GetBlockVolumeOk() (*VolumeBackupProtectionReqVolume, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *VolumeBackupProtectionReq) SetBlockVolume(v VolumeBackupProtectionReqVolume)`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *VolumeBackupProtectionReq) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


