# PdRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** | time of creating protection domain | [optional] 
**Description** | Pointer to **string** | description of protection domain | [optional] 
**Id** | Pointer to **int64** | id of protection domain | [optional] 
**Name** | Pointer to **string** | name of protection domain | [optional] 
**PlacementNode** | Pointer to [**PlacementNodeNestview**](PlacementNodeNestview.md) |  | [optional] 
**Status** | Pointer to **string** | status protection domain | [optional] 
**Update** | Pointer to **time.Time** | time of updating protection domain | [optional] 

## Methods

### NewPdRecord

`func NewPdRecord() *PdRecord`

NewPdRecord instantiates a new PdRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPdRecordWithDefaults

`func NewPdRecordWithDefaults() *PdRecord`

NewPdRecordWithDefaults instantiates a new PdRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *PdRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PdRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PdRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PdRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *PdRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PdRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PdRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PdRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PdRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PdRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PdRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PdRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PdRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PdRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PdRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PdRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacementNode

`func (o *PdRecord) GetPlacementNode() PlacementNodeNestview`

GetPlacementNode returns the PlacementNode field if non-nil, zero value otherwise.

### GetPlacementNodeOk

`func (o *PdRecord) GetPlacementNodeOk() (*PlacementNodeNestview, bool)`

GetPlacementNodeOk returns a tuple with the PlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementNode

`func (o *PdRecord) SetPlacementNode(v PlacementNodeNestview)`

SetPlacementNode sets PlacementNode field to given value.

### HasPlacementNode

`func (o *PdRecord) HasPlacementNode() bool`

HasPlacementNode returns a boolean if a field has been set.

### GetStatus

`func (o *PdRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PdRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PdRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PdRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *PdRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PdRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PdRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PdRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


