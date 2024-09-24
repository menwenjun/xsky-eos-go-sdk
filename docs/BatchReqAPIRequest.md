# BatchReqAPIRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **map[string]interface{}** | API request specification. | [optional] 
**Method** | **string** | API request method. | 
**PathAndParams** | **string** | API request path and query. | 

## Methods

### NewBatchReqAPIRequest

`func NewBatchReqAPIRequest(method string, pathAndParams string, ) *BatchReqAPIRequest`

NewBatchReqAPIRequest instantiates a new BatchReqAPIRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchReqAPIRequestWithDefaults

`func NewBatchReqAPIRequestWithDefaults() *BatchReqAPIRequest`

NewBatchReqAPIRequestWithDefaults instantiates a new BatchReqAPIRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *BatchReqAPIRequest) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BatchReqAPIRequest) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BatchReqAPIRequest) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *BatchReqAPIRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetMethod

`func (o *BatchReqAPIRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *BatchReqAPIRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *BatchReqAPIRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPathAndParams

`func (o *BatchReqAPIRequest) GetPathAndParams() string`

GetPathAndParams returns the PathAndParams field if non-nil, zero value otherwise.

### GetPathAndParamsOk

`func (o *BatchReqAPIRequest) GetPathAndParamsOk() (*string, bool)`

GetPathAndParamsOk returns a tuple with the PathAndParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathAndParams

`func (o *BatchReqAPIRequest) SetPathAndParams(v string)`

SetPathAndParams sets PathAndParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


