# ClientLunMapping

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

## Methods

### NewClientLunMapping

`func NewClientLunMapping() *ClientLunMapping`

NewClientLunMapping instantiates a new ClientLunMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientLunMappingWithDefaults

`func NewClientLunMappingWithDefaults() *ClientLunMapping`

NewClientLunMappingWithDefaults instantiates a new ClientLunMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolume

`func (o *ClientLunMapping) GetBlockVolume() Volume`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *ClientLunMapping) GetBlockVolumeOk() (*Volume, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *ClientLunMapping) SetBlockVolume(v Volume)`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *ClientLunMapping) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.

### GetClient

`func (o *ClientLunMapping) GetClient() Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ClientLunMapping) GetClientOk() (*Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ClientLunMapping) SetClient(v Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *ClientLunMapping) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCluster

`func (o *ClientLunMapping) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClientLunMapping) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClientLunMapping) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ClientLunMapping) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ClientLunMapping) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ClientLunMapping) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ClientLunMapping) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ClientLunMapping) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetHost

`func (o *ClientLunMapping) GetHost() Host`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ClientLunMapping) GetHostOk() (*Host, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ClientLunMapping) SetHost(v Host)`

SetHost sets Host field to given value.

### HasHost

`func (o *ClientLunMapping) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ClientLunMapping) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientLunMapping) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientLunMapping) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClientLunMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLun

`func (o *ClientLunMapping) GetLun() LunNestview`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *ClientLunMapping) GetLunOk() (*LunNestview, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *ClientLunMapping) SetLun(v LunNestview)`

SetLun sets Lun field to given value.

### HasLun

`func (o *ClientLunMapping) HasLun() bool`

HasLun returns a boolean if a field has been set.

### GetStatus

`func (o *ClientLunMapping) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientLunMapping) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientLunMapping) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClientLunMapping) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *ClientLunMapping) GetTarget() TargetNestview`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ClientLunMapping) GetTargetOk() (*TargetNestview, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ClientLunMapping) SetTarget(v TargetNestview)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ClientLunMapping) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUpdate

`func (o *ClientLunMapping) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ClientLunMapping) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ClientLunMapping) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ClientLunMapping) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


