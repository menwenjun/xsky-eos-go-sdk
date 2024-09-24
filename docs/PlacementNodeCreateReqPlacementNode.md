# PlacementNodeCreateReqPlacementNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **int64** |  | [optional] 
**IsWitness** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**ParentId** | **int64** |  | 
**Type** | **string** |  | 

## Methods

### NewPlacementNodeCreateReqPlacementNode

`func NewPlacementNodeCreateReqPlacementNode(name string, parentId int64, type_ string, ) *PlacementNodeCreateReqPlacementNode`

NewPlacementNodeCreateReqPlacementNode instantiates a new PlacementNodeCreateReqPlacementNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementNodeCreateReqPlacementNodeWithDefaults

`func NewPlacementNodeCreateReqPlacementNodeWithDefaults() *PlacementNodeCreateReqPlacementNode`

NewPlacementNodeCreateReqPlacementNodeWithDefaults instantiates a new PlacementNodeCreateReqPlacementNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *PlacementNodeCreateReqPlacementNode) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PlacementNodeCreateReqPlacementNode) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PlacementNodeCreateReqPlacementNode) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *PlacementNodeCreateReqPlacementNode) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetIsWitness

`func (o *PlacementNodeCreateReqPlacementNode) GetIsWitness() bool`

GetIsWitness returns the IsWitness field if non-nil, zero value otherwise.

### GetIsWitnessOk

`func (o *PlacementNodeCreateReqPlacementNode) GetIsWitnessOk() (*bool, bool)`

GetIsWitnessOk returns a tuple with the IsWitness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWitness

`func (o *PlacementNodeCreateReqPlacementNode) SetIsWitness(v bool)`

SetIsWitness sets IsWitness field to given value.

### HasIsWitness

`func (o *PlacementNodeCreateReqPlacementNode) HasIsWitness() bool`

HasIsWitness returns a boolean if a field has been set.

### GetName

`func (o *PlacementNodeCreateReqPlacementNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlacementNodeCreateReqPlacementNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlacementNodeCreateReqPlacementNode) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *PlacementNodeCreateReqPlacementNode) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PlacementNodeCreateReqPlacementNode) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PlacementNodeCreateReqPlacementNode) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### GetType

`func (o *PlacementNodeCreateReqPlacementNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlacementNodeCreateReqPlacementNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlacementNodeCreateReqPlacementNode) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


