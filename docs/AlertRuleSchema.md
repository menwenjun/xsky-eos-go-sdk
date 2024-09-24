# AlertRuleSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to **[]string** |  | [optional] 
**Types** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAlertRuleSchema

`func NewAlertRuleSchema() *AlertRuleSchema`

NewAlertRuleSchema instantiates a new AlertRuleSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleSchemaWithDefaults

`func NewAlertRuleSchemaWithDefaults() *AlertRuleSchema`

NewAlertRuleSchemaWithDefaults instantiates a new AlertRuleSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *AlertRuleSchema) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AlertRuleSchema) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AlertRuleSchema) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AlertRuleSchema) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTypes

`func (o *AlertRuleSchema) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *AlertRuleSchema) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *AlertRuleSchema) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *AlertRuleSchema) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


