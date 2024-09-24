# EventLogCreateReqEventLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** |  | [optional] 
**Event** | **string** |  | 
**HostId** | **int64** |  | 
**Message** | Pointer to **string** |  | [optional] 
**OspClusterId** | **int64** |  | 
**ResourceId** | **int64** |  | 
**ResourceType** | **string** |  | 
**UserId** | **int64** |  | 

## Methods

### NewEventLogCreateReqEventLog

`func NewEventLogCreateReqEventLog(event string, hostId int64, ospClusterId int64, resourceId int64, resourceType string, userId int64, ) *EventLogCreateReqEventLog`

NewEventLogCreateReqEventLog instantiates a new EventLogCreateReqEventLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLogCreateReqEventLogWithDefaults

`func NewEventLogCreateReqEventLogWithDefaults() *EventLogCreateReqEventLog`

NewEventLogCreateReqEventLogWithDefaults instantiates a new EventLogCreateReqEventLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventLogCreateReqEventLog) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventLogCreateReqEventLog) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventLogCreateReqEventLog) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *EventLogCreateReqEventLog) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEvent

`func (o *EventLogCreateReqEventLog) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventLogCreateReqEventLog) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventLogCreateReqEventLog) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetHostId

`func (o *EventLogCreateReqEventLog) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *EventLogCreateReqEventLog) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *EventLogCreateReqEventLog) SetHostId(v int64)`

SetHostId sets HostId field to given value.


### GetMessage

`func (o *EventLogCreateReqEventLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventLogCreateReqEventLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventLogCreateReqEventLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventLogCreateReqEventLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOspClusterId

`func (o *EventLogCreateReqEventLog) GetOspClusterId() int64`

GetOspClusterId returns the OspClusterId field if non-nil, zero value otherwise.

### GetOspClusterIdOk

`func (o *EventLogCreateReqEventLog) GetOspClusterIdOk() (*int64, bool)`

GetOspClusterIdOk returns a tuple with the OspClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterId

`func (o *EventLogCreateReqEventLog) SetOspClusterId(v int64)`

SetOspClusterId sets OspClusterId field to given value.


### GetResourceId

`func (o *EventLogCreateReqEventLog) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *EventLogCreateReqEventLog) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *EventLogCreateReqEventLog) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *EventLogCreateReqEventLog) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *EventLogCreateReqEventLog) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *EventLogCreateReqEventLog) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetUserId

`func (o *EventLogCreateReqEventLog) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventLogCreateReqEventLog) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventLogCreateReqEventLog) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


