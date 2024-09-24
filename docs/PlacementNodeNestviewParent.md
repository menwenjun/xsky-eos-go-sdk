# PlacementNodeNestviewParent

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

## Methods

### NewPlacementNodeNestviewParent

`func NewPlacementNodeNestviewParent() *PlacementNodeNestviewParent`

NewPlacementNodeNestviewParent instantiates a new PlacementNodeNestviewParent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementNodeNestviewParentWithDefaults

`func NewPlacementNodeNestviewParentWithDefaults() *PlacementNodeNestviewParent`

NewPlacementNodeNestviewParentWithDefaults instantiates a new PlacementNodeNestviewParent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *PlacementNodeNestviewParent) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PlacementNodeNestviewParent) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PlacementNodeNestviewParent) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PlacementNodeNestviewParent) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *PlacementNodeNestviewParent) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PlacementNodeNestviewParent) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PlacementNodeNestviewParent) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PlacementNodeNestviewParent) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *PlacementNodeNestviewParent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlacementNodeNestviewParent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlacementNodeNestviewParent) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PlacementNodeNestviewParent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PlacementNodeNestviewParent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlacementNodeNestviewParent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlacementNodeNestviewParent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlacementNodeNestviewParent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigId

`func (o *PlacementNodeNestviewParent) GetOrigId() int64`

GetOrigId returns the OrigId field if non-nil, zero value otherwise.

### GetOrigIdOk

`func (o *PlacementNodeNestviewParent) GetOrigIdOk() (*int64, bool)`

GetOrigIdOk returns a tuple with the OrigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigId

`func (o *PlacementNodeNestviewParent) SetOrigId(v int64)`

SetOrigId sets OrigId field to given value.

### HasOrigId

`func (o *PlacementNodeNestviewParent) HasOrigId() bool`

HasOrigId returns a boolean if a field has been set.

### GetParent

`func (o *PlacementNodeNestviewParent) GetParent() PlacementNode`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PlacementNodeNestviewParent) GetParentOk() (*PlacementNode, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PlacementNodeNestviewParent) SetParent(v PlacementNode)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PlacementNodeNestviewParent) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProperties

`func (o *PlacementNodeNestviewParent) GetProperties() PlacementNodeProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PlacementNodeNestviewParent) GetPropertiesOk() (*PlacementNodeProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PlacementNodeNestviewParent) SetProperties(v PlacementNodeProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PlacementNodeNestviewParent) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetType

`func (o *PlacementNodeNestviewParent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlacementNodeNestviewParent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlacementNodeNestviewParent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlacementNodeNestviewParent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *PlacementNodeNestviewParent) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PlacementNodeNestviewParent) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PlacementNodeNestviewParent) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PlacementNodeNestviewParent) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


