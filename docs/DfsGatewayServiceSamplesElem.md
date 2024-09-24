# DfsGatewayServiceSamplesElem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | id of dfs gateway service | [optional] 
**Name** | Pointer to **string** | name of dfs gateway service | [optional] 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Samples** | Pointer to [**[]DfsGatewayServiceStat**](DfsGatewayServiceStat.md) | sample list of dfs gateway service | [optional] 

## Methods

### NewDfsGatewayServiceSamplesElem

`func NewDfsGatewayServiceSamplesElem() *DfsGatewayServiceSamplesElem`

NewDfsGatewayServiceSamplesElem instantiates a new DfsGatewayServiceSamplesElem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayServiceSamplesElemWithDefaults

`func NewDfsGatewayServiceSamplesElemWithDefaults() *DfsGatewayServiceSamplesElem`

NewDfsGatewayServiceSamplesElemWithDefaults instantiates a new DfsGatewayServiceSamplesElem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DfsGatewayServiceSamplesElem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsGatewayServiceSamplesElem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsGatewayServiceSamplesElem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsGatewayServiceSamplesElem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsGatewayServiceSamplesElem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsGatewayServiceSamplesElem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsGatewayServiceSamplesElem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsGatewayServiceSamplesElem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaging

`func (o *DfsGatewayServiceSamplesElem) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *DfsGatewayServiceSamplesElem) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *DfsGatewayServiceSamplesElem) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *DfsGatewayServiceSamplesElem) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetSamples

`func (o *DfsGatewayServiceSamplesElem) GetSamples() []DfsGatewayServiceStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DfsGatewayServiceSamplesElem) GetSamplesOk() (*[]DfsGatewayServiceStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DfsGatewayServiceSamplesElem) SetSamples(v []DfsGatewayServiceStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DfsGatewayServiceSamplesElem) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


