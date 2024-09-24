# UpdateVolumeStatsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolName** | Pointer to **string** |  | [optional] 
**VolumeStats** | Pointer to [**map[string]VolumeStat**](VolumeStat.md) |  | [optional] 

## Methods

### NewUpdateVolumeStatsReq

`func NewUpdateVolumeStatsReq() *UpdateVolumeStatsReq`

NewUpdateVolumeStatsReq instantiates a new UpdateVolumeStatsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVolumeStatsReqWithDefaults

`func NewUpdateVolumeStatsReqWithDefaults() *UpdateVolumeStatsReq`

NewUpdateVolumeStatsReqWithDefaults instantiates a new UpdateVolumeStatsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolName

`func (o *UpdateVolumeStatsReq) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *UpdateVolumeStatsReq) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *UpdateVolumeStatsReq) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *UpdateVolumeStatsReq) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetVolumeStats

`func (o *UpdateVolumeStatsReq) GetVolumeStats() map[string]VolumeStat`

GetVolumeStats returns the VolumeStats field if non-nil, zero value otherwise.

### GetVolumeStatsOk

`func (o *UpdateVolumeStatsReq) GetVolumeStatsOk() (*map[string]VolumeStat, bool)`

GetVolumeStatsOk returns a tuple with the VolumeStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeStats

`func (o *UpdateVolumeStatsReq) SetVolumeStats(v map[string]VolumeStat)`

SetVolumeStats sets VolumeStats field to given value.

### HasVolumeStats

`func (o *UpdateVolumeStatsReq) HasVolumeStats() bool`

HasVolumeStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


