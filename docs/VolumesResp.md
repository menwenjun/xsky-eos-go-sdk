# VolumesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolumes** | [**[]VolumeRecord**](VolumeRecord.md) | volumes | 

## Methods

### NewVolumesResp

`func NewVolumesResp(blockVolumes []VolumeRecord, ) *VolumesResp`

NewVolumesResp instantiates a new VolumesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumesRespWithDefaults

`func NewVolumesRespWithDefaults() *VolumesResp`

NewVolumesRespWithDefaults instantiates a new VolumesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolumes

`func (o *VolumesResp) GetBlockVolumes() []VolumeRecord`

GetBlockVolumes returns the BlockVolumes field if non-nil, zero value otherwise.

### GetBlockVolumesOk

`func (o *VolumesResp) GetBlockVolumesOk() (*[]VolumeRecord, bool)`

GetBlockVolumesOk returns a tuple with the BlockVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumes

`func (o *VolumesResp) SetBlockVolumes(v []VolumeRecord)`

SetBlockVolumes sets BlockVolumes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


