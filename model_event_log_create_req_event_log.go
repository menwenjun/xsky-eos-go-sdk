/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EventLogCreateReqEventLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventLogCreateReqEventLog{}

// EventLogCreateReqEventLog struct for EventLogCreateReqEventLog
type EventLogCreateReqEventLog struct {
	Data *string `json:"data,omitempty"`
	Event string `json:"event"`
	HostId int64 `json:"host_id"`
	Message *string `json:"message,omitempty"`
	OspClusterId int64 `json:"osp_cluster_id"`
	ResourceId int64 `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	UserId int64 `json:"user_id"`
}

type _EventLogCreateReqEventLog EventLogCreateReqEventLog

// NewEventLogCreateReqEventLog instantiates a new EventLogCreateReqEventLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventLogCreateReqEventLog(event string, hostId int64, ospClusterId int64, resourceId int64, resourceType string, userId int64) *EventLogCreateReqEventLog {
	this := EventLogCreateReqEventLog{}
	this.Event = event
	this.HostId = hostId
	this.OspClusterId = ospClusterId
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.UserId = userId
	return &this
}

// NewEventLogCreateReqEventLogWithDefaults instantiates a new EventLogCreateReqEventLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventLogCreateReqEventLogWithDefaults() *EventLogCreateReqEventLog {
	this := EventLogCreateReqEventLog{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EventLogCreateReqEventLog) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLogCreateReqEventLog) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EventLogCreateReqEventLog) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *EventLogCreateReqEventLog) SetData(v string) {
	o.Data = &v
}

// GetEvent returns the Event field value
func (o *EventLogCreateReqEventLog) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EventLogCreateReqEventLog) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *EventLogCreateReqEventLog) SetEvent(v string) {
	o.Event = v
}

// GetHostId returns the HostId field value
func (o *EventLogCreateReqEventLog) GetHostId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value
// and a boolean to check if the value has been set.
func (o *EventLogCreateReqEventLog) GetHostIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostId, true
}

// SetHostId sets field value
func (o *EventLogCreateReqEventLog) SetHostId(v int64) {
	o.HostId = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EventLogCreateReqEventLog) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLogCreateReqEventLog) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EventLogCreateReqEventLog) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EventLogCreateReqEventLog) SetMessage(v string) {
	o.Message = &v
}

// GetOspClusterId returns the OspClusterId field value
func (o *EventLogCreateReqEventLog) GetOspClusterId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OspClusterId
}

// GetOspClusterIdOk returns a tuple with the OspClusterId field value
// and a boolean to check if the value has been set.
func (o *EventLogCreateReqEventLog) GetOspClusterIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OspClusterId, true
}

// SetOspClusterId sets field value
func (o *EventLogCreateReqEventLog) SetOspClusterId(v int64) {
	o.OspClusterId = v
}

// GetResourceId returns the ResourceId field value
func (o *EventLogCreateReqEventLog) GetResourceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *EventLogCreateReqEventLog) GetResourceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *EventLogCreateReqEventLog) SetResourceId(v int64) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value
func (o *EventLogCreateReqEventLog) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *EventLogCreateReqEventLog) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *EventLogCreateReqEventLog) SetResourceType(v string) {
	o.ResourceType = v
}

// GetUserId returns the UserId field value
func (o *EventLogCreateReqEventLog) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *EventLogCreateReqEventLog) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *EventLogCreateReqEventLog) SetUserId(v int64) {
	o.UserId = v
}

func (o EventLogCreateReqEventLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventLogCreateReqEventLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["event"] = o.Event
	toSerialize["host_id"] = o.HostId
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["osp_cluster_id"] = o.OspClusterId
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["user_id"] = o.UserId
	return toSerialize, nil
}

func (o *EventLogCreateReqEventLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
		"host_id",
		"osp_cluster_id",
		"resource_id",
		"resource_type",
		"user_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEventLogCreateReqEventLog := _EventLogCreateReqEventLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventLogCreateReqEventLog)

	if err != nil {
		return err
	}

	*o = EventLogCreateReqEventLog(varEventLogCreateReqEventLog)

	return err
}

type NullableEventLogCreateReqEventLog struct {
	value *EventLogCreateReqEventLog
	isSet bool
}

func (v NullableEventLogCreateReqEventLog) Get() *EventLogCreateReqEventLog {
	return v.value
}

func (v *NullableEventLogCreateReqEventLog) Set(val *EventLogCreateReqEventLog) {
	v.value = val
	v.isSet = true
}

func (v NullableEventLogCreateReqEventLog) IsSet() bool {
	return v.isSet
}

func (v *NullableEventLogCreateReqEventLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventLogCreateReqEventLog(val *EventLogCreateReqEventLog) *NullableEventLogCreateReqEventLog {
	return &NullableEventLogCreateReqEventLog{value: val, isSet: true}
}

func (v NullableEventLogCreateReqEventLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventLogCreateReqEventLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


