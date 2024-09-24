# PoolPlacementNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]PoolPlacementNode**](PoolPlacementNode.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**PlacementNode** | Pointer to [**PlacementNode**](PlacementNode.md) |  | [optional] 
**ReplicateNum** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPoolPlacementNode

`func NewPoolPlacementNode() *PoolPlacementNode`

NewPoolPlacementNode instantiates a new PoolPlacementNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolPlacementNodeWithDefaults

`func NewPoolPlacementNodeWithDefaults() *PoolPlacementNode`

NewPoolPlacementNodeWithDefaults instantiates a new PoolPlacementNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *PoolPlacementNode) GetChildren() []PoolPlacementNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *PoolPlacementNode) GetChildrenOk() (*[]PoolPlacementNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *PoolPlacementNode) SetChildren(v []PoolPlacementNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *PoolPlacementNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCreate

`func (o *PoolPlacementNode) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PoolPlacementNode) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PoolPlacementNode) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PoolPlacementNode) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *PoolPlacementNode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolPlacementNode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolPlacementNode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PoolPlacementNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlacementNode

`func (o *PoolPlacementNode) GetPlacementNode() PlacementNode`

GetPlacementNode returns the PlacementNode field if non-nil, zero value otherwise.

### GetPlacementNodeOk

`func (o *PoolPlacementNode) GetPlacementNodeOk() (*PlacementNode, bool)`

GetPlacementNodeOk returns a tuple with the PlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementNode

`func (o *PoolPlacementNode) SetPlacementNode(v PlacementNode)`

SetPlacementNode sets PlacementNode field to given value.

### HasPlacementNode

`func (o *PoolPlacementNode) HasPlacementNode() bool`

HasPlacementNode returns a boolean if a field has been set.

### GetReplicateNum

`func (o *PoolPlacementNode) GetReplicateNum() int64`

GetReplicateNum returns the ReplicateNum field if non-nil, zero value otherwise.

### GetReplicateNumOk

`func (o *PoolPlacementNode) GetReplicateNumOk() (*int64, bool)`

GetReplicateNumOk returns a tuple with the ReplicateNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateNum

`func (o *PoolPlacementNode) SetReplicateNum(v int64)`

SetReplicateNum sets ReplicateNum field to given value.

### HasReplicateNum

`func (o *PoolPlacementNode) HasReplicateNum() bool`

HasReplicateNum returns a boolean if a field has been set.

### GetStatus

`func (o *PoolPlacementNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoolPlacementNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoolPlacementNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PoolPlacementNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *PoolPlacementNode) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PoolPlacementNode) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PoolPlacementNode) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PoolPlacementNode) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


