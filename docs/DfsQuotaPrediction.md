# DfsQuotaPrediction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**UsedFiles1day** | Pointer to **int64** |  | [optional] 
**UsedFiles30days** | Pointer to **int64** |  | [optional] 
**UsedFiles7days** | Pointer to **int64** |  | [optional] 
**UsedFilesPoints** | Pointer to [**[]FilesPredictionPoint**](FilesPredictionPoint.md) |  | [optional] 
**UsedKbyte1day** | Pointer to **int64** |  | [optional] 
**UsedKbyte30days** | Pointer to **int64** |  | [optional] 
**UsedKbyte7days** | Pointer to **int64** |  | [optional] 
**UsedKbytePoints** | Pointer to [**[]PredictionPoint**](PredictionPoint.md) |  | [optional] 

## Methods

### NewDfsQuotaPrediction

`func NewDfsQuotaPrediction() *DfsQuotaPrediction`

NewDfsQuotaPrediction instantiates a new DfsQuotaPrediction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQuotaPredictionWithDefaults

`func NewDfsQuotaPredictionWithDefaults() *DfsQuotaPrediction`

NewDfsQuotaPredictionWithDefaults instantiates a new DfsQuotaPrediction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DfsQuotaPrediction) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsQuotaPrediction) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsQuotaPrediction) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsQuotaPrediction) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *DfsQuotaPrediction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsQuotaPrediction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsQuotaPrediction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsQuotaPrediction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsedFiles1day

`func (o *DfsQuotaPrediction) GetUsedFiles1day() int64`

GetUsedFiles1day returns the UsedFiles1day field if non-nil, zero value otherwise.

### GetUsedFiles1dayOk

`func (o *DfsQuotaPrediction) GetUsedFiles1dayOk() (*int64, bool)`

GetUsedFiles1dayOk returns a tuple with the UsedFiles1day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFiles1day

`func (o *DfsQuotaPrediction) SetUsedFiles1day(v int64)`

SetUsedFiles1day sets UsedFiles1day field to given value.

### HasUsedFiles1day

`func (o *DfsQuotaPrediction) HasUsedFiles1day() bool`

HasUsedFiles1day returns a boolean if a field has been set.

### GetUsedFiles30days

`func (o *DfsQuotaPrediction) GetUsedFiles30days() int64`

GetUsedFiles30days returns the UsedFiles30days field if non-nil, zero value otherwise.

### GetUsedFiles30daysOk

`func (o *DfsQuotaPrediction) GetUsedFiles30daysOk() (*int64, bool)`

GetUsedFiles30daysOk returns a tuple with the UsedFiles30days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFiles30days

`func (o *DfsQuotaPrediction) SetUsedFiles30days(v int64)`

SetUsedFiles30days sets UsedFiles30days field to given value.

### HasUsedFiles30days

`func (o *DfsQuotaPrediction) HasUsedFiles30days() bool`

HasUsedFiles30days returns a boolean if a field has been set.

### GetUsedFiles7days

`func (o *DfsQuotaPrediction) GetUsedFiles7days() int64`

GetUsedFiles7days returns the UsedFiles7days field if non-nil, zero value otherwise.

### GetUsedFiles7daysOk

`func (o *DfsQuotaPrediction) GetUsedFiles7daysOk() (*int64, bool)`

GetUsedFiles7daysOk returns a tuple with the UsedFiles7days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFiles7days

`func (o *DfsQuotaPrediction) SetUsedFiles7days(v int64)`

SetUsedFiles7days sets UsedFiles7days field to given value.

### HasUsedFiles7days

`func (o *DfsQuotaPrediction) HasUsedFiles7days() bool`

HasUsedFiles7days returns a boolean if a field has been set.

### GetUsedFilesPoints

`func (o *DfsQuotaPrediction) GetUsedFilesPoints() []FilesPredictionPoint`

GetUsedFilesPoints returns the UsedFilesPoints field if non-nil, zero value otherwise.

### GetUsedFilesPointsOk

`func (o *DfsQuotaPrediction) GetUsedFilesPointsOk() (*[]FilesPredictionPoint, bool)`

GetUsedFilesPointsOk returns a tuple with the UsedFilesPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFilesPoints

`func (o *DfsQuotaPrediction) SetUsedFilesPoints(v []FilesPredictionPoint)`

SetUsedFilesPoints sets UsedFilesPoints field to given value.

### HasUsedFilesPoints

`func (o *DfsQuotaPrediction) HasUsedFilesPoints() bool`

HasUsedFilesPoints returns a boolean if a field has been set.

### GetUsedKbyte1day

`func (o *DfsQuotaPrediction) GetUsedKbyte1day() int64`

GetUsedKbyte1day returns the UsedKbyte1day field if non-nil, zero value otherwise.

### GetUsedKbyte1dayOk

`func (o *DfsQuotaPrediction) GetUsedKbyte1dayOk() (*int64, bool)`

GetUsedKbyte1dayOk returns a tuple with the UsedKbyte1day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte1day

`func (o *DfsQuotaPrediction) SetUsedKbyte1day(v int64)`

SetUsedKbyte1day sets UsedKbyte1day field to given value.

### HasUsedKbyte1day

`func (o *DfsQuotaPrediction) HasUsedKbyte1day() bool`

HasUsedKbyte1day returns a boolean if a field has been set.

### GetUsedKbyte30days

`func (o *DfsQuotaPrediction) GetUsedKbyte30days() int64`

GetUsedKbyte30days returns the UsedKbyte30days field if non-nil, zero value otherwise.

### GetUsedKbyte30daysOk

`func (o *DfsQuotaPrediction) GetUsedKbyte30daysOk() (*int64, bool)`

GetUsedKbyte30daysOk returns a tuple with the UsedKbyte30days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte30days

`func (o *DfsQuotaPrediction) SetUsedKbyte30days(v int64)`

SetUsedKbyte30days sets UsedKbyte30days field to given value.

### HasUsedKbyte30days

`func (o *DfsQuotaPrediction) HasUsedKbyte30days() bool`

HasUsedKbyte30days returns a boolean if a field has been set.

### GetUsedKbyte7days

`func (o *DfsQuotaPrediction) GetUsedKbyte7days() int64`

GetUsedKbyte7days returns the UsedKbyte7days field if non-nil, zero value otherwise.

### GetUsedKbyte7daysOk

`func (o *DfsQuotaPrediction) GetUsedKbyte7daysOk() (*int64, bool)`

GetUsedKbyte7daysOk returns a tuple with the UsedKbyte7days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte7days

`func (o *DfsQuotaPrediction) SetUsedKbyte7days(v int64)`

SetUsedKbyte7days sets UsedKbyte7days field to given value.

### HasUsedKbyte7days

`func (o *DfsQuotaPrediction) HasUsedKbyte7days() bool`

HasUsedKbyte7days returns a boolean if a field has been set.

### GetUsedKbytePoints

`func (o *DfsQuotaPrediction) GetUsedKbytePoints() []PredictionPoint`

GetUsedKbytePoints returns the UsedKbytePoints field if non-nil, zero value otherwise.

### GetUsedKbytePointsOk

`func (o *DfsQuotaPrediction) GetUsedKbytePointsOk() (*[]PredictionPoint, bool)`

GetUsedKbytePointsOk returns a tuple with the UsedKbytePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbytePoints

`func (o *DfsQuotaPrediction) SetUsedKbytePoints(v []PredictionPoint)`

SetUsedKbytePoints sets UsedKbytePoints field to given value.

### HasUsedKbytePoints

`func (o *DfsQuotaPrediction) HasUsedKbytePoints() bool`

HasUsedKbytePoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


