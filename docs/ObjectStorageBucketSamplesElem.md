# ObjectStorageBucketSamplesElem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Samples** | Pointer to [**[]ObjectStorageBucketStat**](ObjectStorageBucketStat.md) |  | [optional] 

## Methods

### NewObjectStorageBucketSamplesElem

`func NewObjectStorageBucketSamplesElem() *ObjectStorageBucketSamplesElem`

NewObjectStorageBucketSamplesElem instantiates a new ObjectStorageBucketSamplesElem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketSamplesElemWithDefaults

`func NewObjectStorageBucketSamplesElemWithDefaults() *ObjectStorageBucketSamplesElem`

NewObjectStorageBucketSamplesElemWithDefaults instantiates a new ObjectStorageBucketSamplesElem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectStorageBucketSamplesElem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageBucketSamplesElem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageBucketSamplesElem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageBucketSamplesElem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaging

`func (o *ObjectStorageBucketSamplesElem) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ObjectStorageBucketSamplesElem) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ObjectStorageBucketSamplesElem) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *ObjectStorageBucketSamplesElem) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetSamples

`func (o *ObjectStorageBucketSamplesElem) GetSamples() []ObjectStorageBucketStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ObjectStorageBucketSamplesElem) GetSamplesOk() (*[]ObjectStorageBucketStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ObjectStorageBucketSamplesElem) SetSamples(v []ObjectStorageBucketStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ObjectStorageBucketSamplesElem) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


