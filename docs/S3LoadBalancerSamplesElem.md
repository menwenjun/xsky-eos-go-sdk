# S3LoadBalancerSamplesElem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Samples** | Pointer to [**[]S3LoadBalancerStat**](S3LoadBalancerStat.md) |  | [optional] 

## Methods

### NewS3LoadBalancerSamplesElem

`func NewS3LoadBalancerSamplesElem() *S3LoadBalancerSamplesElem`

NewS3LoadBalancerSamplesElem instantiates a new S3LoadBalancerSamplesElem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerSamplesElemWithDefaults

`func NewS3LoadBalancerSamplesElemWithDefaults() *S3LoadBalancerSamplesElem`

NewS3LoadBalancerSamplesElemWithDefaults instantiates a new S3LoadBalancerSamplesElem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3LoadBalancerSamplesElem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3LoadBalancerSamplesElem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3LoadBalancerSamplesElem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *S3LoadBalancerSamplesElem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaging

`func (o *S3LoadBalancerSamplesElem) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *S3LoadBalancerSamplesElem) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *S3LoadBalancerSamplesElem) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *S3LoadBalancerSamplesElem) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetSamples

`func (o *S3LoadBalancerSamplesElem) GetSamples() []S3LoadBalancerStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *S3LoadBalancerSamplesElem) GetSamplesOk() (*[]S3LoadBalancerStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *S3LoadBalancerSamplesElem) SetSamples(v []S3LoadBalancerStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *S3LoadBalancerSamplesElem) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


