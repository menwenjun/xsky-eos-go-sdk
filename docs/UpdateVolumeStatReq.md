# UpdateVolumeStatReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolName** | Pointer to **string** |  | [optional] 
**Stat** | Pointer to [**VolumeStat**](VolumeStat.md) |  | [optional] 
**VolumeName** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateVolumeStatReq

`func NewUpdateVolumeStatReq() *UpdateVolumeStatReq`

NewUpdateVolumeStatReq instantiates a new UpdateVolumeStatReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVolumeStatReqWithDefaults

`func NewUpdateVolumeStatReqWithDefaults() *UpdateVolumeStatReq`

NewUpdateVolumeStatReqWithDefaults instantiates a new UpdateVolumeStatReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolName

`func (o *UpdateVolumeStatReq) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *UpdateVolumeStatReq) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *UpdateVolumeStatReq) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *UpdateVolumeStatReq) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetStat

`func (o *UpdateVolumeStatReq) GetStat() VolumeStat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *UpdateVolumeStatReq) GetStatOk() (*VolumeStat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *UpdateVolumeStatReq) SetStat(v VolumeStat)`

SetStat sets Stat field to given value.

### HasStat

`func (o *UpdateVolumeStatReq) HasStat() bool`

HasStat returns a boolean if a field has been set.

### GetVolumeName

`func (o *UpdateVolumeStatReq) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *UpdateVolumeStatReq) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *UpdateVolumeStatReq) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *UpdateVolumeStatReq) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


