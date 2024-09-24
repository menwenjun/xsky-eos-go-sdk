# PlacementNode

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

### NewPlacementNode

`func NewPlacementNode() *PlacementNode`

NewPlacementNode instantiates a new PlacementNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementNodeWithDefaults

`func NewPlacementNodeWithDefaults() *PlacementNode`

NewPlacementNodeWithDefaults instantiates a new PlacementNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *PlacementNode) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PlacementNode) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PlacementNode) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PlacementNode) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *PlacementNode) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PlacementNode) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PlacementNode) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PlacementNode) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *PlacementNode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlacementNode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlacementNode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PlacementNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PlacementNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlacementNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlacementNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlacementNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigId

`func (o *PlacementNode) GetOrigId() int64`

GetOrigId returns the OrigId field if non-nil, zero value otherwise.

### GetOrigIdOk

`func (o *PlacementNode) GetOrigIdOk() (*int64, bool)`

GetOrigIdOk returns a tuple with the OrigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigId

`func (o *PlacementNode) SetOrigId(v int64)`

SetOrigId sets OrigId field to given value.

### HasOrigId

`func (o *PlacementNode) HasOrigId() bool`

HasOrigId returns a boolean if a field has been set.

### GetParent

`func (o *PlacementNode) GetParent() PlacementNode`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PlacementNode) GetParentOk() (*PlacementNode, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PlacementNode) SetParent(v PlacementNode)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PlacementNode) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProperties

`func (o *PlacementNode) GetProperties() PlacementNodeProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PlacementNode) GetPropertiesOk() (*PlacementNodeProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PlacementNode) SetProperties(v PlacementNodeProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PlacementNode) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetType

`func (o *PlacementNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlacementNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlacementNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlacementNode) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *PlacementNode) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PlacementNode) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PlacementNode) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PlacementNode) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


