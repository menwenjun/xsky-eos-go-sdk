# VolumeSamplesElem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Samples** | Pointer to [**[]VolumeStat**](VolumeStat.md) |  | [optional] 

## Methods

### NewVolumeSamplesElem

`func NewVolumeSamplesElem() *VolumeSamplesElem`

NewVolumeSamplesElem instantiates a new VolumeSamplesElem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSamplesElemWithDefaults

`func NewVolumeSamplesElemWithDefaults() *VolumeSamplesElem`

NewVolumeSamplesElemWithDefaults instantiates a new VolumeSamplesElem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VolumeSamplesElem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeSamplesElem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeSamplesElem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeSamplesElem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaging

`func (o *VolumeSamplesElem) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *VolumeSamplesElem) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *VolumeSamplesElem) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *VolumeSamplesElem) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetSamples

`func (o *VolumeSamplesElem) GetSamples() []VolumeStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *VolumeSamplesElem) GetSamplesOk() (*[]VolumeStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *VolumeSamplesElem) SetSamples(v []VolumeStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *VolumeSamplesElem) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


