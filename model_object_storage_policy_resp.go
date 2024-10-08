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

// checks if the ObjectStoragePolicyResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStoragePolicyResp{}

// ObjectStoragePolicyResp struct for ObjectStoragePolicyResp
type ObjectStoragePolicyResp struct {
	OsPolicy ObjectStoragePolicy `json:"os_policy"`
}

type _ObjectStoragePolicyResp ObjectStoragePolicyResp

// NewObjectStoragePolicyResp instantiates a new ObjectStoragePolicyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStoragePolicyResp(osPolicy ObjectStoragePolicy) *ObjectStoragePolicyResp {
	this := ObjectStoragePolicyResp{}
	this.OsPolicy = osPolicy
	return &this
}

// NewObjectStoragePolicyRespWithDefaults instantiates a new ObjectStoragePolicyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStoragePolicyRespWithDefaults() *ObjectStoragePolicyResp {
	this := ObjectStoragePolicyResp{}
	return &this
}

// GetOsPolicy returns the OsPolicy field value
func (o *ObjectStoragePolicyResp) GetOsPolicy() ObjectStoragePolicy {
	if o == nil {
		var ret ObjectStoragePolicy
		return ret
	}

	return o.OsPolicy
}

// GetOsPolicyOk returns a tuple with the OsPolicy field value
// and a boolean to check if the value has been set.
func (o *ObjectStoragePolicyResp) GetOsPolicyOk() (*ObjectStoragePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsPolicy, true
}

// SetOsPolicy sets field value
func (o *ObjectStoragePolicyResp) SetOsPolicy(v ObjectStoragePolicy) {
	o.OsPolicy = v
}

func (o ObjectStoragePolicyResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStoragePolicyResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_policy"] = o.OsPolicy
	return toSerialize, nil
}

func (o *ObjectStoragePolicyResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_policy",
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

	varObjectStoragePolicyResp := _ObjectStoragePolicyResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectStoragePolicyResp)

	if err != nil {
		return err
	}

	*o = ObjectStoragePolicyResp(varObjectStoragePolicyResp)

	return err
}

type NullableObjectStoragePolicyResp struct {
	value *ObjectStoragePolicyResp
	isSet bool
}

func (v NullableObjectStoragePolicyResp) Get() *ObjectStoragePolicyResp {
	return v.value
}

func (v *NullableObjectStoragePolicyResp) Set(val *ObjectStoragePolicyResp) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStoragePolicyResp) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStoragePolicyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStoragePolicyResp(val *ObjectStoragePolicyResp) *NullableObjectStoragePolicyResp {
	return &NullableObjectStoragePolicyResp{value: val, isSet: true}
}

func (v NullableObjectStoragePolicyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStoragePolicyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


