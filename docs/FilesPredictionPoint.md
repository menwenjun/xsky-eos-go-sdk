# FilesPredictionPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**UsedFiles** | Pointer to **int64** |  | [optional] 

## Methods

### NewFilesPredictionPoint

`func NewFilesPredictionPoint() *FilesPredictionPoint`

NewFilesPredictionPoint instantiates a new FilesPredictionPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesPredictionPointWithDefaults

`func NewFilesPredictionPointWithDefaults() *FilesPredictionPoint`

NewFilesPredictionPointWithDefaults instantiates a new FilesPredictionPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *FilesPredictionPoint) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FilesPredictionPoint) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FilesPredictionPoint) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FilesPredictionPoint) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUsedFiles

`func (o *FilesPredictionPoint) GetUsedFiles() int64`

GetUsedFiles returns the UsedFiles field if non-nil, zero value otherwise.

### GetUsedFilesOk

`func (o *FilesPredictionPoint) GetUsedFilesOk() (*int64, bool)`

GetUsedFilesOk returns a tuple with the UsedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFiles

`func (o *FilesPredictionPoint) SetUsedFiles(v int64)`

SetUsedFiles sets UsedFiles field to given value.

### HasUsedFiles

`func (o *FilesPredictionPoint) HasUsedFiles() bool`

HasUsedFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


