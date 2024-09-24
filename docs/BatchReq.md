# BatchReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionOnFailure** | Pointer to **string** | action of the remaining APIs in the batch when one API failed. Valid values are: \&quot;CONTINUE\&quot;, \&quot;ABORT\&quot;, default value is \&quot;CONTINUE\&quot;. If ExecutionOrder is \&quot;PARALLEL\&quot;, ActionOnFailure will be ignored. | [optional] 
**ApiRequests** | [**[]BatchReqAPIRequest**](BatchReqAPIRequest.md) | list of API requests in the batch. | 
**ExecutionOrder** | Pointer to **string** | ExecutionOrder is the order of execution of the APIs in the batch. Valid values are: \&quot;SEQUENTIAL\&quot;, \&quot;PARALLEL\&quot;, default value is \&quot;SEQUENTIAL\&quot;. | [optional] 

## Methods

### NewBatchReq

`func NewBatchReq(apiRequests []BatchReqAPIRequest, ) *BatchReq`

NewBatchReq instantiates a new BatchReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchReqWithDefaults

`func NewBatchReqWithDefaults() *BatchReq`

NewBatchReqWithDefaults instantiates a new BatchReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionOnFailure

`func (o *BatchReq) GetActionOnFailure() string`

GetActionOnFailure returns the ActionOnFailure field if non-nil, zero value otherwise.

### GetActionOnFailureOk

`func (o *BatchReq) GetActionOnFailureOk() (*string, bool)`

GetActionOnFailureOk returns a tuple with the ActionOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnFailure

`func (o *BatchReq) SetActionOnFailure(v string)`

SetActionOnFailure sets ActionOnFailure field to given value.

### HasActionOnFailure

`func (o *BatchReq) HasActionOnFailure() bool`

HasActionOnFailure returns a boolean if a field has been set.

### GetApiRequests

`func (o *BatchReq) GetApiRequests() []BatchReqAPIRequest`

GetApiRequests returns the ApiRequests field if non-nil, zero value otherwise.

### GetApiRequestsOk

`func (o *BatchReq) GetApiRequestsOk() (*[]BatchReqAPIRequest, bool)`

GetApiRequestsOk returns a tuple with the ApiRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRequests

`func (o *BatchReq) SetApiRequests(v []BatchReqAPIRequest)`

SetApiRequests sets ApiRequests field to given value.


### GetExecutionOrder

`func (o *BatchReq) GetExecutionOrder() string`

GetExecutionOrder returns the ExecutionOrder field if non-nil, zero value otherwise.

### GetExecutionOrderOk

`func (o *BatchReq) GetExecutionOrderOk() (*string, bool)`

GetExecutionOrderOk returns a tuple with the ExecutionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOrder

`func (o *BatchReq) SetExecutionOrder(v string)`

SetExecutionOrder sets ExecutionOrder field to given value.

### HasExecutionOrder

`func (o *BatchReq) HasExecutionOrder() bool`

HasExecutionOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


