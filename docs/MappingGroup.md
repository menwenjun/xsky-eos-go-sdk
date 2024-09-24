# MappingGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPath** | Pointer to [**AccessPathNestview**](AccessPathNestview.md) |  | [optional] 
**ClientGroup** | Pointer to [**ClientGroupNestview**](ClientGroupNestview.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMappingGroup

`func NewMappingGroup() *MappingGroup`

NewMappingGroup instantiates a new MappingGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingGroupWithDefaults

`func NewMappingGroupWithDefaults() *MappingGroup`

NewMappingGroupWithDefaults instantiates a new MappingGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPath

`func (o *MappingGroup) GetAccessPath() AccessPathNestview`

GetAccessPath returns the AccessPath field if non-nil, zero value otherwise.

### GetAccessPathOk

`func (o *MappingGroup) GetAccessPathOk() (*AccessPathNestview, bool)`

GetAccessPathOk returns a tuple with the AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPath

`func (o *MappingGroup) SetAccessPath(v AccessPathNestview)`

SetAccessPath sets AccessPath field to given value.

### HasAccessPath

`func (o *MappingGroup) HasAccessPath() bool`

HasAccessPath returns a boolean if a field has been set.

### GetClientGroup

`func (o *MappingGroup) GetClientGroup() ClientGroupNestview`

GetClientGroup returns the ClientGroup field if non-nil, zero value otherwise.

### GetClientGroupOk

`func (o *MappingGroup) GetClientGroupOk() (*ClientGroupNestview, bool)`

GetClientGroupOk returns a tuple with the ClientGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGroup

`func (o *MappingGroup) SetClientGroup(v ClientGroupNestview)`

SetClientGroup sets ClientGroup field to given value.

### HasClientGroup

`func (o *MappingGroup) HasClientGroup() bool`

HasClientGroup returns a boolean if a field has been set.

### GetCluster

`func (o *MappingGroup) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *MappingGroup) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *MappingGroup) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *MappingGroup) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *MappingGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *MappingGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *MappingGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *MappingGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *MappingGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MappingGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MappingGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MappingGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *MappingGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MappingGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MappingGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MappingGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *MappingGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *MappingGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *MappingGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *MappingGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


