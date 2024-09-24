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

// checks if the GcPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcPolicy{}

// GcPolicy struct for GcPolicy
type GcPolicy struct {
	Bandwidth *int64 `json:"bandwidth,omitempty"`
	Enable bool `json:"enable"`
	EndTime *string `json:"end_time,omitempty"`
	Level *string `json:"level,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
}

type _GcPolicy GcPolicy

// NewGcPolicy instantiates a new GcPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcPolicy(enable bool) *GcPolicy {
	this := GcPolicy{}
	this.Enable = enable
	return &this
}

// NewGcPolicyWithDefaults instantiates a new GcPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcPolicyWithDefaults() *GcPolicy {
	this := GcPolicy{}
	return &this
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *GcPolicy) GetBandwidth() int64 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret int64
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcPolicy) GetBandwidthOk() (*int64, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *GcPolicy) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int64 and assigns it to the Bandwidth field.
func (o *GcPolicy) SetBandwidth(v int64) {
	o.Bandwidth = &v
}

// GetEnable returns the Enable field value
func (o *GcPolicy) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *GcPolicy) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value
func (o *GcPolicy) SetEnable(v bool) {
	o.Enable = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GcPolicy) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcPolicy) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GcPolicy) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *GcPolicy) SetEndTime(v string) {
	o.EndTime = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GcPolicy) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcPolicy) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GcPolicy) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *GcPolicy) SetLevel(v string) {
	o.Level = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GcPolicy) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcPolicy) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GcPolicy) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *GcPolicy) SetStartTime(v string) {
	o.StartTime = &v
}

func (o GcPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	toSerialize["enable"] = o.Enable
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	return toSerialize, nil
}

func (o *GcPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enable",
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

	varGcPolicy := _GcPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGcPolicy)

	if err != nil {
		return err
	}

	*o = GcPolicy(varGcPolicy)

	return err
}

type NullableGcPolicy struct {
	value *GcPolicy
	isSet bool
}

func (v NullableGcPolicy) Get() *GcPolicy {
	return v.value
}

func (v *NullableGcPolicy) Set(val *GcPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableGcPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableGcPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcPolicy(val *GcPolicy) *NullableGcPolicy {
	return &NullableGcPolicy{value: val, isSet: true}
}

func (v NullableGcPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


