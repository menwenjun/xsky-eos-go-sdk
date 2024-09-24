# PlacementNodeTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** | time of creating placement node | [optional] 
**Id** | Pointer to **int64** | id of placement node | [optional] 
**Name** | Pointer to **string** | name of placement node | [optional] 
**OrigId** | Pointer to **int64** | old id of placement node | [optional] 
**Parent** | Pointer to [**PlacementNode**](PlacementNode.md) |  | [optional] 
**Properties** | Pointer to [**PlacementNodeProperties**](PlacementNodeProperties.md) |  | [optional] 
**Type** | Pointer to **string** | type of placement node | [optional] 
**Update** | Pointer to **time.Time** | time of updating placement node | [optional] 
**AvailableOsdNum** | Pointer to **int64** |  | [optional] 
**Children** | Pointer to [**[]PlacementNodeTopology**](PlacementNodeTopology.md) |  | [optional] 
**HostNum** | Pointer to **int64** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPlacementNodeTopology

`func NewPlacementNodeTopology() *PlacementNodeTopology`

NewPlacementNodeTopology instantiates a new PlacementNodeTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementNodeTopologyWithDefaults

`func NewPlacementNodeTopologyWithDefaults() *PlacementNodeTopology`

NewPlacementNodeTopologyWithDefaults instantiates a new PlacementNodeTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *PlacementNodeTopology) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PlacementNodeTopology) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PlacementNodeTopology) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PlacementNodeTopology) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *PlacementNodeTopology) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PlacementNodeTopology) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PlacementNodeTopology) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PlacementNodeTopology) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *PlacementNodeTopology) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlacementNodeTopology) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlacementNodeTopology) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PlacementNodeTopology) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PlacementNodeTopology) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlacementNodeTopology) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlacementNodeTopology) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlacementNodeTopology) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigId

`func (o *PlacementNodeTopology) GetOrigId() int64`

GetOrigId returns the OrigId field if non-nil, zero value otherwise.

### GetOrigIdOk

`func (o *PlacementNodeTopology) GetOrigIdOk() (*int64, bool)`

GetOrigIdOk returns a tuple with the OrigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigId

`func (o *PlacementNodeTopology) SetOrigId(v int64)`

SetOrigId sets OrigId field to given value.

### HasOrigId

`func (o *PlacementNodeTopology) HasOrigId() bool`

HasOrigId returns a boolean if a field has been set.

### GetParent

`func (o *PlacementNodeTopology) GetParent() PlacementNode`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PlacementNodeTopology) GetParentOk() (*PlacementNode, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PlacementNodeTopology) SetParent(v PlacementNode)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PlacementNodeTopology) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProperties

`func (o *PlacementNodeTopology) GetProperties() PlacementNodeProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PlacementNodeTopology) GetPropertiesOk() (*PlacementNodeProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PlacementNodeTopology) SetProperties(v PlacementNodeProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PlacementNodeTopology) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetType

`func (o *PlacementNodeTopology) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlacementNodeTopology) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlacementNodeTopology) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlacementNodeTopology) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *PlacementNodeTopology) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PlacementNodeTopology) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PlacementNodeTopology) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PlacementNodeTopology) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetAvailableOsdNum

`func (o *PlacementNodeTopology) GetAvailableOsdNum() int64`

GetAvailableOsdNum returns the AvailableOsdNum field if non-nil, zero value otherwise.

### GetAvailableOsdNumOk

`func (o *PlacementNodeTopology) GetAvailableOsdNumOk() (*int64, bool)`

GetAvailableOsdNumOk returns a tuple with the AvailableOsdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOsdNum

`func (o *PlacementNodeTopology) SetAvailableOsdNum(v int64)`

SetAvailableOsdNum sets AvailableOsdNum field to given value.

### HasAvailableOsdNum

`func (o *PlacementNodeTopology) HasAvailableOsdNum() bool`

HasAvailableOsdNum returns a boolean if a field has been set.

### GetChildren

`func (o *PlacementNodeTopology) GetChildren() []PlacementNodeTopology`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *PlacementNodeTopology) GetChildrenOk() (*[]PlacementNodeTopology, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *PlacementNodeTopology) SetChildren(v []PlacementNodeTopology)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *PlacementNodeTopology) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetHostNum

`func (o *PlacementNodeTopology) GetHostNum() int64`

GetHostNum returns the HostNum field if non-nil, zero value otherwise.

### GetHostNumOk

`func (o *PlacementNodeTopology) GetHostNumOk() (*int64, bool)`

GetHostNumOk returns a tuple with the HostNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNum

`func (o *PlacementNodeTopology) SetHostNum(v int64)`

SetHostNum sets HostNum field to given value.

### HasHostNum

`func (o *PlacementNodeTopology) HasHostNum() bool`

HasHostNum returns a boolean if a field has been set.

### GetPayload

`func (o *PlacementNodeTopology) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PlacementNodeTopology) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PlacementNodeTopology) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *PlacementNodeTopology) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


