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

// checks if the ObjectStorageLifecycleResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageLifecycleResp{}

// ObjectStorageLifecycleResp struct for ObjectStorageLifecycleResp
type ObjectStorageLifecycleResp struct {
	OsLifecycle ObjectStorageLifecycle `json:"os_lifecycle"`
}

type _ObjectStorageLifecycleResp ObjectStorageLifecycleResp

// NewObjectStorageLifecycleResp instantiates a new ObjectStorageLifecycleResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageLifecycleResp(osLifecycle ObjectStorageLifecycle) *ObjectStorageLifecycleResp {
	this := ObjectStorageLifecycleResp{}
	this.OsLifecycle = osLifecycle
	return &this
}

// NewObjectStorageLifecycleRespWithDefaults instantiates a new ObjectStorageLifecycleResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageLifecycleRespWithDefaults() *ObjectStorageLifecycleResp {
	this := ObjectStorageLifecycleResp{}
	return &this
}

// GetOsLifecycle returns the OsLifecycle field value
func (o *ObjectStorageLifecycleResp) GetOsLifecycle() ObjectStorageLifecycle {
	if o == nil {
		var ret ObjectStorageLifecycle
		return ret
	}

	return o.OsLifecycle
}

// GetOsLifecycleOk returns a tuple with the OsLifecycle field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageLifecycleResp) GetOsLifecycleOk() (*ObjectStorageLifecycle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsLifecycle, true
}

// SetOsLifecycle sets field value
func (o *ObjectStorageLifecycleResp) SetOsLifecycle(v ObjectStorageLifecycle) {
	o.OsLifecycle = v
}

func (o ObjectStorageLifecycleResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageLifecycleResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_lifecycle"] = o.OsLifecycle
	return toSerialize, nil
}

func (o *ObjectStorageLifecycleResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_lifecycle",
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

	varObjectStorageLifecycleResp := _ObjectStorageLifecycleResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectStorageLifecycleResp)

	if err != nil {
		return err
	}

	*o = ObjectStorageLifecycleResp(varObjectStorageLifecycleResp)

	return err
}

type NullableObjectStorageLifecycleResp struct {
	value *ObjectStorageLifecycleResp
	isSet bool
}

func (v NullableObjectStorageLifecycleResp) Get() *ObjectStorageLifecycleResp {
	return v.value
}

func (v *NullableObjectStorageLifecycleResp) Set(val *ObjectStorageLifecycleResp) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageLifecycleResp) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageLifecycleResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageLifecycleResp(val *ObjectStorageLifecycleResp) *NullableObjectStorageLifecycleResp {
	return &NullableObjectStorageLifecycleResp{value: val, isSet: true}
}

func (v NullableObjectStorageLifecycleResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageLifecycleResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


