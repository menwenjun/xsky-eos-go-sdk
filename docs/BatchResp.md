# BatchResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiResponses** | Pointer to [**[]BatchRespAPIResponse**](BatchRespAPIResponse.md) | list of API responses in the batch. | [optional] 
**TotalCostSeconds** | Pointer to **float64** | the total time to execute all API requests. Unit is second. | [optional] 

## Methods

### NewBatchResp

`func NewBatchResp() *BatchResp`

NewBatchResp instantiates a new BatchResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRespWithDefaults

`func NewBatchRespWithDefaults() *BatchResp`

NewBatchRespWithDefaults instantiates a new BatchResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiResponses

`func (o *BatchResp) GetApiResponses() []BatchRespAPIResponse`

GetApiResponses returns the ApiResponses field if non-nil, zero value otherwise.

### GetApiResponsesOk

`func (o *BatchResp) GetApiResponsesOk() (*[]BatchRespAPIResponse, bool)`

GetApiResponsesOk returns a tuple with the ApiResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiResponses

`func (o *BatchResp) SetApiResponses(v []BatchRespAPIResponse)`

SetApiResponses sets ApiResponses field to given value.

### HasApiResponses

`func (o *BatchResp) HasApiResponses() bool`

HasApiResponses returns a boolean if a field has been set.

### GetTotalCostSeconds

`func (o *BatchResp) GetTotalCostSeconds() float64`

GetTotalCostSeconds returns the TotalCostSeconds field if non-nil, zero value otherwise.

### GetTotalCostSecondsOk

`func (o *BatchResp) GetTotalCostSecondsOk() (*float64, bool)`

GetTotalCostSecondsOk returns a tuple with the TotalCostSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostSeconds

`func (o *BatchResp) SetTotalCostSeconds(v float64)`

SetTotalCostSeconds sets TotalCostSeconds field to given value.

### HasTotalCostSeconds

`func (o *BatchResp) HasTotalCostSeconds() bool`

HasTotalCostSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


