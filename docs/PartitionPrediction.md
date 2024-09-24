# PartitionPrediction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**UsedKbyte1day** | Pointer to **int64** |  | [optional] 
**UsedKbyte30days** | Pointer to **int64** |  | [optional] 
**UsedKbyte7days** | Pointer to **int64** |  | [optional] 
**UsedKbytePoints** | Pointer to [**[]PredictionPoint**](PredictionPoint.md) |  | [optional] 

## Methods

### NewPartitionPrediction

`func NewPartitionPrediction() *PartitionPrediction`

NewPartitionPrediction instantiates a new PartitionPrediction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionPredictionWithDefaults

`func NewPartitionPredictionWithDefaults() *PartitionPrediction`

NewPartitionPredictionWithDefaults instantiates a new PartitionPrediction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *PartitionPrediction) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PartitionPrediction) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PartitionPrediction) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PartitionPrediction) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *PartitionPrediction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartitionPrediction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartitionPrediction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PartitionPrediction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsedKbyte1day

`func (o *PartitionPrediction) GetUsedKbyte1day() int64`

GetUsedKbyte1day returns the UsedKbyte1day field if non-nil, zero value otherwise.

### GetUsedKbyte1dayOk

`func (o *PartitionPrediction) GetUsedKbyte1dayOk() (*int64, bool)`

GetUsedKbyte1dayOk returns a tuple with the UsedKbyte1day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte1day

`func (o *PartitionPrediction) SetUsedKbyte1day(v int64)`

SetUsedKbyte1day sets UsedKbyte1day field to given value.

### HasUsedKbyte1day

`func (o *PartitionPrediction) HasUsedKbyte1day() bool`

HasUsedKbyte1day returns a boolean if a field has been set.

### GetUsedKbyte30days

`func (o *PartitionPrediction) GetUsedKbyte30days() int64`

GetUsedKbyte30days returns the UsedKbyte30days field if non-nil, zero value otherwise.

### GetUsedKbyte30daysOk

`func (o *PartitionPrediction) GetUsedKbyte30daysOk() (*int64, bool)`

GetUsedKbyte30daysOk returns a tuple with the UsedKbyte30days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte30days

`func (o *PartitionPrediction) SetUsedKbyte30days(v int64)`

SetUsedKbyte30days sets UsedKbyte30days field to given value.

### HasUsedKbyte30days

`func (o *PartitionPrediction) HasUsedKbyte30days() bool`

HasUsedKbyte30days returns a boolean if a field has been set.

### GetUsedKbyte7days

`func (o *PartitionPrediction) GetUsedKbyte7days() int64`

GetUsedKbyte7days returns the UsedKbyte7days field if non-nil, zero value otherwise.

### GetUsedKbyte7daysOk

`func (o *PartitionPrediction) GetUsedKbyte7daysOk() (*int64, bool)`

GetUsedKbyte7daysOk returns a tuple with the UsedKbyte7days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte7days

`func (o *PartitionPrediction) SetUsedKbyte7days(v int64)`

SetUsedKbyte7days sets UsedKbyte7days field to given value.

### HasUsedKbyte7days

`func (o *PartitionPrediction) HasUsedKbyte7days() bool`

HasUsedKbyte7days returns a boolean if a field has been set.

### GetUsedKbytePoints

`func (o *PartitionPrediction) GetUsedKbytePoints() []PredictionPoint`

GetUsedKbytePoints returns the UsedKbytePoints field if non-nil, zero value otherwise.

### GetUsedKbytePointsOk

`func (o *PartitionPrediction) GetUsedKbytePointsOk() (*[]PredictionPoint, bool)`

GetUsedKbytePointsOk returns a tuple with the UsedKbytePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbytePoints

`func (o *PartitionPrediction) SetUsedKbytePoints(v []PredictionPoint)`

SetUsedKbytePoints sets UsedKbytePoints field to given value.

### HasUsedKbytePoints

`func (o *PartitionPrediction) HasUsedKbytePoints() bool`

HasUsedKbytePoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


