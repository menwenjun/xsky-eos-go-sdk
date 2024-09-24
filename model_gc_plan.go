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

// checks if the GcPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcPlan{}

// GcPlan struct for GcPlan
type GcPlan struct {
	Days []int64 `json:"days"`
	Policies []GcPolicy `json:"policies"`
}

type _GcPlan GcPlan

// NewGcPlan instantiates a new GcPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcPlan(days []int64, policies []GcPolicy) *GcPlan {
	this := GcPlan{}
	this.Days = days
	this.Policies = policies
	return &this
}

// NewGcPlanWithDefaults instantiates a new GcPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcPlanWithDefaults() *GcPlan {
	this := GcPlan{}
	return &this
}

// GetDays returns the Days field value
func (o *GcPlan) GetDays() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *GcPlan) GetDaysOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Days, true
}

// SetDays sets field value
func (o *GcPlan) SetDays(v []int64) {
	o.Days = v
}

// GetPolicies returns the Policies field value
func (o *GcPlan) GetPolicies() []GcPolicy {
	if o == nil {
		var ret []GcPolicy
		return ret
	}

	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value
// and a boolean to check if the value has been set.
func (o *GcPlan) GetPoliciesOk() ([]GcPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policies, true
}

// SetPolicies sets field value
func (o *GcPlan) SetPolicies(v []GcPolicy) {
	o.Policies = v
}

func (o GcPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["days"] = o.Days
	toSerialize["policies"] = o.Policies
	return toSerialize, nil
}

func (o *GcPlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"days",
		"policies",
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

	varGcPlan := _GcPlan{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGcPlan)

	if err != nil {
		return err
	}

	*o = GcPlan(varGcPlan)

	return err
}

type NullableGcPlan struct {
	value *GcPlan
	isSet bool
}

func (v NullableGcPlan) Get() *GcPlan {
	return v.value
}

func (v *NullableGcPlan) Set(val *GcPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableGcPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableGcPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcPlan(val *GcPlan) *NullableGcPlan {
	return &NullableGcPlan{value: val, isSet: true}
}

func (v NullableGcPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


