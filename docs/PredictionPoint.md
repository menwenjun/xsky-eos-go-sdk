# PredictionPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**UsedKbyte** | Pointer to **int64** |  | [optional] 

## Methods

### NewPredictionPoint

`func NewPredictionPoint() *PredictionPoint`

NewPredictionPoint instantiates a new PredictionPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionPointWithDefaults

`func NewPredictionPointWithDefaults() *PredictionPoint`

NewPredictionPointWithDefaults instantiates a new PredictionPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *PredictionPoint) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PredictionPoint) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PredictionPoint) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PredictionPoint) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *PredictionPoint) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *PredictionPoint) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *PredictionPoint) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *PredictionPoint) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


