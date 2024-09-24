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

// checks if the OSExternalStorageBucketInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSExternalStorageBucketInfo{}

// OSExternalStorageBucketInfo OSExternalStorageBucketInfo contains bucket info of external storage platform
type OSExternalStorageBucketInfo struct {
	Name string `json:"name"`
	Weight int64 `json:"weight"`
}

type _OSExternalStorageBucketInfo OSExternalStorageBucketInfo

// NewOSExternalStorageBucketInfo instantiates a new OSExternalStorageBucketInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSExternalStorageBucketInfo(name string, weight int64) *OSExternalStorageBucketInfo {
	this := OSExternalStorageBucketInfo{}
	this.Name = name
	this.Weight = weight
	return &this
}

// NewOSExternalStorageBucketInfoWithDefaults instantiates a new OSExternalStorageBucketInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSExternalStorageBucketInfoWithDefaults() *OSExternalStorageBucketInfo {
	this := OSExternalStorageBucketInfo{}
	return &this
}

// GetName returns the Name field value
func (o *OSExternalStorageBucketInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OSExternalStorageBucketInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OSExternalStorageBucketInfo) SetName(v string) {
	o.Name = v
}

// GetWeight returns the Weight field value
func (o *OSExternalStorageBucketInfo) GetWeight() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *OSExternalStorageBucketInfo) GetWeightOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *OSExternalStorageBucketInfo) SetWeight(v int64) {
	o.Weight = v
}

func (o OSExternalStorageBucketInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSExternalStorageBucketInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

func (o *OSExternalStorageBucketInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"weight",
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

	varOSExternalStorageBucketInfo := _OSExternalStorageBucketInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSExternalStorageBucketInfo)

	if err != nil {
		return err
	}

	*o = OSExternalStorageBucketInfo(varOSExternalStorageBucketInfo)

	return err
}

type NullableOSExternalStorageBucketInfo struct {
	value *OSExternalStorageBucketInfo
	isSet bool
}

func (v NullableOSExternalStorageBucketInfo) Get() *OSExternalStorageBucketInfo {
	return v.value
}

func (v *NullableOSExternalStorageBucketInfo) Set(val *OSExternalStorageBucketInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOSExternalStorageBucketInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOSExternalStorageBucketInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSExternalStorageBucketInfo(val *OSExternalStorageBucketInfo) *NullableOSExternalStorageBucketInfo {
	return &NullableOSExternalStorageBucketInfo{value: val, isSet: true}
}

func (v NullableOSExternalStorageBucketInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSExternalStorageBucketInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


