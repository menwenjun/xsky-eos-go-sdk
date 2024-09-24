# BatchRespAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **map[string]interface{}** | http response body. | [optional] 
**CostSeconds** | Pointer to **float64** | the time to execute the API request. Unit is second. | [optional] 
**StatusCode** | Pointer to **int64** | http response status code. | [optional] 

## Methods

### NewBatchRespAPIResponse

`func NewBatchRespAPIResponse() *BatchRespAPIResponse`

NewBatchRespAPIResponse instantiates a new BatchRespAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRespAPIResponseWithDefaults

`func NewBatchRespAPIResponseWithDefaults() *BatchRespAPIResponse`

NewBatchRespAPIResponseWithDefaults instantiates a new BatchRespAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *BatchRespAPIResponse) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BatchRespAPIResponse) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BatchRespAPIResponse) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *BatchRespAPIResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCostSeconds

`func (o *BatchRespAPIResponse) GetCostSeconds() float64`

GetCostSeconds returns the CostSeconds field if non-nil, zero value otherwise.

### GetCostSecondsOk

`func (o *BatchRespAPIResponse) GetCostSecondsOk() (*float64, bool)`

GetCostSecondsOk returns a tuple with the CostSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostSeconds

`func (o *BatchRespAPIResponse) SetCostSeconds(v float64)`

SetCostSeconds sets CostSeconds field to given value.

### HasCostSeconds

`func (o *BatchRespAPIResponse) HasCostSeconds() bool`

HasCostSeconds returns a boolean if a field has been set.

### GetStatusCode

`func (o *BatchRespAPIResponse) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BatchRespAPIResponse) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BatchRespAPIResponse) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BatchRespAPIResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


