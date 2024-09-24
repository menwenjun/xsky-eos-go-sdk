# Lun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPath** | Pointer to [**AccessPathNestview**](AccessPathNestview.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LunId** | Pointer to **int64** |  | [optional] 
**Volume** | Pointer to [**VolumeNestview**](VolumeNestview.md) |  | [optional] 

## Methods

### NewLun

`func NewLun() *Lun`

NewLun instantiates a new Lun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLunWithDefaults

`func NewLunWithDefaults() *Lun`

NewLunWithDefaults instantiates a new Lun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPath

`func (o *Lun) GetAccessPath() AccessPathNestview`

GetAccessPath returns the AccessPath field if non-nil, zero value otherwise.

### GetAccessPathOk

`func (o *Lun) GetAccessPathOk() (*AccessPathNestview, bool)`

GetAccessPathOk returns a tuple with the AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPath

`func (o *Lun) SetAccessPath(v AccessPathNestview)`

SetAccessPath sets AccessPath field to given value.

### HasAccessPath

`func (o *Lun) HasAccessPath() bool`

HasAccessPath returns a boolean if a field has been set.

### GetCluster

`func (o *Lun) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Lun) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Lun) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Lun) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *Lun) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Lun) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Lun) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Lun) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *Lun) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Lun) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Lun) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Lun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLunId

`func (o *Lun) GetLunId() int64`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *Lun) GetLunIdOk() (*int64, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *Lun) SetLunId(v int64)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *Lun) HasLunId() bool`

HasLunId returns a boolean if a field has been set.

### GetVolume

`func (o *Lun) GetVolume() VolumeNestview`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Lun) GetVolumeOk() (*VolumeNestview, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Lun) SetVolume(v VolumeNestview)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Lun) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


