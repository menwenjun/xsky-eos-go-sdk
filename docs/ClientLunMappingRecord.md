# ClientLunMappingRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolume** | Pointer to [**Volume**](Volume.md) |  | [optional] 
**Client** | Pointer to [**Client**](Client.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to [**Host**](Host.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Lun** | Pointer to [**LunNestview**](LunNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**TargetNestview**](TargetNestview.md) |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]VolumeStat**](VolumeStat.md) |  | [optional] 
**SessionStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewClientLunMappingRecord

`func NewClientLunMappingRecord() *ClientLunMappingRecord`

NewClientLunMappingRecord instantiates a new ClientLunMappingRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientLunMappingRecordWithDefaults

`func NewClientLunMappingRecordWithDefaults() *ClientLunMappingRecord`

NewClientLunMappingRecordWithDefaults instantiates a new ClientLunMappingRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolume

`func (o *ClientLunMappingRecord) GetBlockVolume() Volume`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *ClientLunMappingRecord) GetBlockVolumeOk() (*Volume, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *ClientLunMappingRecord) SetBlockVolume(v Volume)`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *ClientLunMappingRecord) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.

### GetClient

`func (o *ClientLunMappingRecord) GetClient() Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ClientLunMappingRecord) GetClientOk() (*Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ClientLunMappingRecord) SetClient(v Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *ClientLunMappingRecord) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCluster

`func (o *ClientLunMappingRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClientLunMappingRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClientLunMappingRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ClientLunMappingRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ClientLunMappingRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ClientLunMappingRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ClientLunMappingRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ClientLunMappingRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetHost

`func (o *ClientLunMappingRecord) GetHost() Host`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ClientLunMappingRecord) GetHostOk() (*Host, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ClientLunMappingRecord) SetHost(v Host)`

SetHost sets Host field to given value.

### HasHost

`func (o *ClientLunMappingRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ClientLunMappingRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientLunMappingRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientLunMappingRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClientLunMappingRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLun

`func (o *ClientLunMappingRecord) GetLun() LunNestview`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *ClientLunMappingRecord) GetLunOk() (*LunNestview, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *ClientLunMappingRecord) SetLun(v LunNestview)`

SetLun sets Lun field to given value.

### HasLun

`func (o *ClientLunMappingRecord) HasLun() bool`

HasLun returns a boolean if a field has been set.

### GetStatus

`func (o *ClientLunMappingRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientLunMappingRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientLunMappingRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClientLunMappingRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *ClientLunMappingRecord) GetTarget() TargetNestview`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ClientLunMappingRecord) GetTargetOk() (*TargetNestview, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ClientLunMappingRecord) SetTarget(v TargetNestview)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ClientLunMappingRecord) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUpdate

`func (o *ClientLunMappingRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ClientLunMappingRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ClientLunMappingRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ClientLunMappingRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *ClientLunMappingRecord) GetSamples() []VolumeStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ClientLunMappingRecord) GetSamplesOk() (*[]VolumeStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ClientLunMappingRecord) SetSamples(v []VolumeStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ClientLunMappingRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSessionStatus

`func (o *ClientLunMappingRecord) GetSessionStatus() string`

GetSessionStatus returns the SessionStatus field if non-nil, zero value otherwise.

### GetSessionStatusOk

`func (o *ClientLunMappingRecord) GetSessionStatusOk() (*string, bool)`

GetSessionStatusOk returns a tuple with the SessionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStatus

`func (o *ClientLunMappingRecord) SetSessionStatus(v string)`

SetSessionStatus sets SessionStatus field to given value.

### HasSessionStatus

`func (o *ClientLunMappingRecord) HasSessionStatus() bool`

HasSessionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


