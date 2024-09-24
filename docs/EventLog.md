# EventLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** | create time | [optional] 
**Data** | Pointer to **map[string]interface{}** | dumped resource data | [optional] 
**Event** | Pointer to **string** | evnet type | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** | debug message | [optional] 
**OspCluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**ResourceId** | Pointer to **int64** | resource id | [optional] 
**ResourceType** | Pointer to **string** | resource type | [optional] 
**User** | Pointer to [**UserNestview**](UserNestview.md) |  | [optional] 

## Methods

### NewEventLog

`func NewEventLog() *EventLog`

NewEventLog instantiates a new EventLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLogWithDefaults

`func NewEventLogWithDefaults() *EventLog`

NewEventLogWithDefaults instantiates a new EventLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *EventLog) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *EventLog) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *EventLog) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *EventLog) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *EventLog) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *EventLog) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *EventLog) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *EventLog) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetData

`func (o *EventLog) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventLog) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventLog) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EventLog) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEvent

`func (o *EventLog) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventLog) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventLog) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *EventLog) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetHost

`func (o *EventLog) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EventLog) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EventLog) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *EventLog) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *EventLog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventLog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventLog) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EventLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *EventLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOspCluster

`func (o *EventLog) GetOspCluster() ClusterNestview`

GetOspCluster returns the OspCluster field if non-nil, zero value otherwise.

### GetOspClusterOk

`func (o *EventLog) GetOspClusterOk() (*ClusterNestview, bool)`

GetOspClusterOk returns a tuple with the OspCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspCluster

`func (o *EventLog) SetOspCluster(v ClusterNestview)`

SetOspCluster sets OspCluster field to given value.

### HasOspCluster

`func (o *EventLog) HasOspCluster() bool`

HasOspCluster returns a boolean if a field has been set.

### GetResourceId

`func (o *EventLog) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *EventLog) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *EventLog) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *EventLog) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *EventLog) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *EventLog) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *EventLog) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *EventLog) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetUser

`func (o *EventLog) GetUser() UserNestview`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventLog) GetUserOk() (*UserNestview, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventLog) SetUser(v UserNestview)`

SetUser sets User field to given value.

### HasUser

`func (o *EventLog) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


