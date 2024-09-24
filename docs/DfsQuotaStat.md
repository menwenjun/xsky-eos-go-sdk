# DfsQuotaStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**UsedFiles** | Pointer to **int64** |  | [optional] 
**UsedKbyte** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsQuotaStat

`func NewDfsQuotaStat() *DfsQuotaStat`

NewDfsQuotaStat instantiates a new DfsQuotaStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQuotaStatWithDefaults

`func NewDfsQuotaStatWithDefaults() *DfsQuotaStat`

NewDfsQuotaStatWithDefaults instantiates a new DfsQuotaStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DfsQuotaStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsQuotaStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsQuotaStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsQuotaStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUsedFiles

`func (o *DfsQuotaStat) GetUsedFiles() int64`

GetUsedFiles returns the UsedFiles field if non-nil, zero value otherwise.

### GetUsedFilesOk

`func (o *DfsQuotaStat) GetUsedFilesOk() (*int64, bool)`

GetUsedFilesOk returns a tuple with the UsedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFiles

`func (o *DfsQuotaStat) SetUsedFiles(v int64)`

SetUsedFiles sets UsedFiles field to given value.

### HasUsedFiles

`func (o *DfsQuotaStat) HasUsedFiles() bool`

HasUsedFiles returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *DfsQuotaStat) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *DfsQuotaStat) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *DfsQuotaStat) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *DfsQuotaStat) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


