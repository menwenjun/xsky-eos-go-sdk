# VolumeReplicationSetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolume** | Pointer to [**VolumeReplicationSetReqVolume**](VolumeReplicationSetReqVolume.md) |  | [optional] 

## Methods

### NewVolumeReplicationSetReq

`func NewVolumeReplicationSetReq() *VolumeReplicationSetReq`

NewVolumeReplicationSetReq instantiates a new VolumeReplicationSetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeReplicationSetReqWithDefaults

`func NewVolumeReplicationSetReqWithDefaults() *VolumeReplicationSetReq`

NewVolumeReplicationSetReqWithDefaults instantiates a new VolumeReplicationSetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolume

`func (o *VolumeReplicationSetReq) GetBlockVolume() VolumeReplicationSetReqVolume`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *VolumeReplicationSetReq) GetBlockVolumeOk() (*VolumeReplicationSetReqVolume, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *VolumeReplicationSetReq) SetBlockVolume(v VolumeReplicationSetReqVolume)`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *VolumeReplicationSetReq) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


