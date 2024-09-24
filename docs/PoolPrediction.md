# PoolPrediction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**DataKbyte1day** | Pointer to **int64** |  | [optional] 
**DataKbyte30days** | Pointer to **int64** |  | [optional] 
**DataKbyte7days** | Pointer to **int64** |  | [optional] 
**DataKbytePoints** | Pointer to [**[]DataKbytePredictionPoint**](DataKbytePredictionPoint.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**UsedKbyte1day** | Pointer to **int64** |  | [optional] 
**UsedKbyte30days** | Pointer to **int64** |  | [optional] 
**UsedKbyte7days** | Pointer to **int64** |  | [optional] 
**UsedKbytePoints** | Pointer to [**[]PredictionPoint**](PredictionPoint.md) |  | [optional] 

## Methods

### NewPoolPrediction

`func NewPoolPrediction() *PoolPrediction`

NewPoolPrediction instantiates a new PoolPrediction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolPredictionWithDefaults

`func NewPoolPredictionWithDefaults() *PoolPrediction`

NewPoolPredictionWithDefaults instantiates a new PoolPrediction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *PoolPrediction) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *PoolPrediction) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *PoolPrediction) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *PoolPrediction) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataKbyte1day

`func (o *PoolPrediction) GetDataKbyte1day() int64`

GetDataKbyte1day returns the DataKbyte1day field if non-nil, zero value otherwise.

### GetDataKbyte1dayOk

`func (o *PoolPrediction) GetDataKbyte1dayOk() (*int64, bool)`

GetDataKbyte1dayOk returns a tuple with the DataKbyte1day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKbyte1day

`func (o *PoolPrediction) SetDataKbyte1day(v int64)`

SetDataKbyte1day sets DataKbyte1day field to given value.

### HasDataKbyte1day

`func (o *PoolPrediction) HasDataKbyte1day() bool`

HasDataKbyte1day returns a boolean if a field has been set.

### GetDataKbyte30days

`func (o *PoolPrediction) GetDataKbyte30days() int64`

GetDataKbyte30days returns the DataKbyte30days field if non-nil, zero value otherwise.

### GetDataKbyte30daysOk

`func (o *PoolPrediction) GetDataKbyte30daysOk() (*int64, bool)`

GetDataKbyte30daysOk returns a tuple with the DataKbyte30days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKbyte30days

`func (o *PoolPrediction) SetDataKbyte30days(v int64)`

SetDataKbyte30days sets DataKbyte30days field to given value.

### HasDataKbyte30days

`func (o *PoolPrediction) HasDataKbyte30days() bool`

HasDataKbyte30days returns a boolean if a field has been set.

### GetDataKbyte7days

`func (o *PoolPrediction) GetDataKbyte7days() int64`

GetDataKbyte7days returns the DataKbyte7days field if non-nil, zero value otherwise.

### GetDataKbyte7daysOk

`func (o *PoolPrediction) GetDataKbyte7daysOk() (*int64, bool)`

GetDataKbyte7daysOk returns a tuple with the DataKbyte7days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKbyte7days

`func (o *PoolPrediction) SetDataKbyte7days(v int64)`

SetDataKbyte7days sets DataKbyte7days field to given value.

### HasDataKbyte7days

`func (o *PoolPrediction) HasDataKbyte7days() bool`

HasDataKbyte7days returns a boolean if a field has been set.

### GetDataKbytePoints

`func (o *PoolPrediction) GetDataKbytePoints() []DataKbytePredictionPoint`

GetDataKbytePoints returns the DataKbytePoints field if non-nil, zero value otherwise.

### GetDataKbytePointsOk

`func (o *PoolPrediction) GetDataKbytePointsOk() (*[]DataKbytePredictionPoint, bool)`

GetDataKbytePointsOk returns a tuple with the DataKbytePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKbytePoints

`func (o *PoolPrediction) SetDataKbytePoints(v []DataKbytePredictionPoint)`

SetDataKbytePoints sets DataKbytePoints field to given value.

### HasDataKbytePoints

`func (o *PoolPrediction) HasDataKbytePoints() bool`

HasDataKbytePoints returns a boolean if a field has been set.

### GetId

`func (o *PoolPrediction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolPrediction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolPrediction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PoolPrediction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsedKbyte1day

`func (o *PoolPrediction) GetUsedKbyte1day() int64`

GetUsedKbyte1day returns the UsedKbyte1day field if non-nil, zero value otherwise.

### GetUsedKbyte1dayOk

`func (o *PoolPrediction) GetUsedKbyte1dayOk() (*int64, bool)`

GetUsedKbyte1dayOk returns a tuple with the UsedKbyte1day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte1day

`func (o *PoolPrediction) SetUsedKbyte1day(v int64)`

SetUsedKbyte1day sets UsedKbyte1day field to given value.

### HasUsedKbyte1day

`func (o *PoolPrediction) HasUsedKbyte1day() bool`

HasUsedKbyte1day returns a boolean if a field has been set.

### GetUsedKbyte30days

`func (o *PoolPrediction) GetUsedKbyte30days() int64`

GetUsedKbyte30days returns the UsedKbyte30days field if non-nil, zero value otherwise.

### GetUsedKbyte30daysOk

`func (o *PoolPrediction) GetUsedKbyte30daysOk() (*int64, bool)`

GetUsedKbyte30daysOk returns a tuple with the UsedKbyte30days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte30days

`func (o *PoolPrediction) SetUsedKbyte30days(v int64)`

SetUsedKbyte30days sets UsedKbyte30days field to given value.

### HasUsedKbyte30days

`func (o *PoolPrediction) HasUsedKbyte30days() bool`

HasUsedKbyte30days returns a boolean if a field has been set.

### GetUsedKbyte7days

`func (o *PoolPrediction) GetUsedKbyte7days() int64`

GetUsedKbyte7days returns the UsedKbyte7days field if non-nil, zero value otherwise.

### GetUsedKbyte7daysOk

`func (o *PoolPrediction) GetUsedKbyte7daysOk() (*int64, bool)`

GetUsedKbyte7daysOk returns a tuple with the UsedKbyte7days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte7days

`func (o *PoolPrediction) SetUsedKbyte7days(v int64)`

SetUsedKbyte7days sets UsedKbyte7days field to given value.

### HasUsedKbyte7days

`func (o *PoolPrediction) HasUsedKbyte7days() bool`

HasUsedKbyte7days returns a boolean if a field has been set.

### GetUsedKbytePoints

`func (o *PoolPrediction) GetUsedKbytePoints() []PredictionPoint`

GetUsedKbytePoints returns the UsedKbytePoints field if non-nil, zero value otherwise.

### GetUsedKbytePointsOk

`func (o *PoolPrediction) GetUsedKbytePointsOk() (*[]PredictionPoint, bool)`

GetUsedKbytePointsOk returns a tuple with the UsedKbytePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbytePoints

`func (o *PoolPrediction) SetUsedKbytePoints(v []PredictionPoint)`

SetUsedKbytePoints sets UsedKbytePoints field to given value.

### HasUsedKbytePoints

`func (o *PoolPrediction) HasUsedKbytePoints() bool`

HasUsedKbytePoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


