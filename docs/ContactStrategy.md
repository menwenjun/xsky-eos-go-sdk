# ContactStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelStrategies** | Pointer to [**[]AlertStrategy**](AlertStrategy.md) |  | [optional] 
**DeleteStrategies** | Pointer to [**[]AlertStrategy**](AlertStrategy.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewContactStrategy

`func NewContactStrategy() *ContactStrategy`

NewContactStrategy instantiates a new ContactStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactStrategyWithDefaults

`func NewContactStrategyWithDefaults() *ContactStrategy`

NewContactStrategyWithDefaults instantiates a new ContactStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelStrategies

`func (o *ContactStrategy) GetCancelStrategies() []AlertStrategy`

GetCancelStrategies returns the CancelStrategies field if non-nil, zero value otherwise.

### GetCancelStrategiesOk

`func (o *ContactStrategy) GetCancelStrategiesOk() (*[]AlertStrategy, bool)`

GetCancelStrategiesOk returns a tuple with the CancelStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelStrategies

`func (o *ContactStrategy) SetCancelStrategies(v []AlertStrategy)`

SetCancelStrategies sets CancelStrategies field to given value.

### HasCancelStrategies

`func (o *ContactStrategy) HasCancelStrategies() bool`

HasCancelStrategies returns a boolean if a field has been set.

### GetDeleteStrategies

`func (o *ContactStrategy) GetDeleteStrategies() []AlertStrategy`

GetDeleteStrategies returns the DeleteStrategies field if non-nil, zero value otherwise.

### GetDeleteStrategiesOk

`func (o *ContactStrategy) GetDeleteStrategiesOk() (*[]AlertStrategy, bool)`

GetDeleteStrategiesOk returns a tuple with the DeleteStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteStrategies

`func (o *ContactStrategy) SetDeleteStrategies(v []AlertStrategy)`

SetDeleteStrategies sets DeleteStrategies field to given value.

### HasDeleteStrategies

`func (o *ContactStrategy) HasDeleteStrategies() bool`

HasDeleteStrategies returns a boolean if a field has been set.

### GetType

`func (o *ContactStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactStrategy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContactStrategy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


