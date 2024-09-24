# EventLogCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventLog** | Pointer to [**EventLogCreateReqEventLog**](EventLogCreateReqEventLog.md) |  | [optional] 

## Methods

### NewEventLogCreateReq

`func NewEventLogCreateReq() *EventLogCreateReq`

NewEventLogCreateReq instantiates a new EventLogCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLogCreateReqWithDefaults

`func NewEventLogCreateReqWithDefaults() *EventLogCreateReq`

NewEventLogCreateReqWithDefaults instantiates a new EventLogCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventLog

`func (o *EventLogCreateReq) GetEventLog() EventLogCreateReqEventLog`

GetEventLog returns the EventLog field if non-nil, zero value otherwise.

### GetEventLogOk

`func (o *EventLogCreateReq) GetEventLogOk() (*EventLogCreateReqEventLog, bool)`

GetEventLogOk returns a tuple with the EventLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLog

`func (o *EventLogCreateReq) SetEventLog(v EventLogCreateReqEventLog)`

SetEventLog sets EventLog field to given value.

### HasEventLog

`func (o *EventLogCreateReq) HasEventLog() bool`

HasEventLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


