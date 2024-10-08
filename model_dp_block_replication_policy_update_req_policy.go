/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DpBlockReplicationPolicyUpdateReqPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpBlockReplicationPolicyUpdateReqPolicy{}

// DpBlockReplicationPolicyUpdateReqPolicy struct for DpBlockReplicationPolicyUpdateReqPolicy
type DpBlockReplicationPolicyUpdateReqPolicy struct {
	// name
	Name *string `json:"name,omitempty"`
	// replication timeout seconds
	TimeoutSeconds *int64 `json:"timeout_seconds,omitempty"`
}

// NewDpBlockReplicationPolicyUpdateReqPolicy instantiates a new DpBlockReplicationPolicyUpdateReqPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpBlockReplicationPolicyUpdateReqPolicy() *DpBlockReplicationPolicyUpdateReqPolicy {
	this := DpBlockReplicationPolicyUpdateReqPolicy{}
	return &this
}

// NewDpBlockReplicationPolicyUpdateReqPolicyWithDefaults instantiates a new DpBlockReplicationPolicyUpdateReqPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpBlockReplicationPolicyUpdateReqPolicyWithDefaults() *DpBlockReplicationPolicyUpdateReqPolicy {
	this := DpBlockReplicationPolicyUpdateReqPolicy{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DpBlockReplicationPolicyUpdateReqPolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockReplicationPolicyUpdateReqPolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DpBlockReplicationPolicyUpdateReqPolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DpBlockReplicationPolicyUpdateReqPolicy) SetName(v string) {
	o.Name = &v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *DpBlockReplicationPolicyUpdateReqPolicy) GetTimeoutSeconds() int64 {
	if o == nil || IsNil(o.TimeoutSeconds) {
		var ret int64
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockReplicationPolicyUpdateReqPolicy) GetTimeoutSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeoutSeconds) {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *DpBlockReplicationPolicyUpdateReqPolicy) HasTimeoutSeconds() bool {
	if o != nil && !IsNil(o.TimeoutSeconds) {
		return true
	}

	return false
}

// SetTimeoutSeconds gets a reference to the given int64 and assigns it to the TimeoutSeconds field.
func (o *DpBlockReplicationPolicyUpdateReqPolicy) SetTimeoutSeconds(v int64) {
	o.TimeoutSeconds = &v
}

func (o DpBlockReplicationPolicyUpdateReqPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpBlockReplicationPolicyUpdateReqPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TimeoutSeconds) {
		toSerialize["timeout_seconds"] = o.TimeoutSeconds
	}
	return toSerialize, nil
}

type NullableDpBlockReplicationPolicyUpdateReqPolicy struct {
	value *DpBlockReplicationPolicyUpdateReqPolicy
	isSet bool
}

func (v NullableDpBlockReplicationPolicyUpdateReqPolicy) Get() *DpBlockReplicationPolicyUpdateReqPolicy {
	return v.value
}

func (v *NullableDpBlockReplicationPolicyUpdateReqPolicy) Set(val *DpBlockReplicationPolicyUpdateReqPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDpBlockReplicationPolicyUpdateReqPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDpBlockReplicationPolicyUpdateReqPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpBlockReplicationPolicyUpdateReqPolicy(val *DpBlockReplicationPolicyUpdateReqPolicy) *NullableDpBlockReplicationPolicyUpdateReqPolicy {
	return &NullableDpBlockReplicationPolicyUpdateReqPolicy{value: val, isSet: true}
}

func (v NullableDpBlockReplicationPolicyUpdateReqPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpBlockReplicationPolicyUpdateReqPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


