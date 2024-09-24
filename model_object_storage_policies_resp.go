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

// checks if the ObjectStoragePoliciesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStoragePoliciesResp{}

// ObjectStoragePoliciesResp struct for ObjectStoragePoliciesResp
type ObjectStoragePoliciesResp struct {
	// object storage policies
	OsPolicies []ObjectStoragePolicy `json:"os_policies"`
}

type _ObjectStoragePoliciesResp ObjectStoragePoliciesResp

// NewObjectStoragePoliciesResp instantiates a new ObjectStoragePoliciesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStoragePoliciesResp(osPolicies []ObjectStoragePolicy) *ObjectStoragePoliciesResp {
	this := ObjectStoragePoliciesResp{}
	this.OsPolicies = osPolicies
	return &this
}

// NewObjectStoragePoliciesRespWithDefaults instantiates a new ObjectStoragePoliciesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStoragePoliciesRespWithDefaults() *ObjectStoragePoliciesResp {
	this := ObjectStoragePoliciesResp{}
	return &this
}

// GetOsPolicies returns the OsPolicies field value
func (o *ObjectStoragePoliciesResp) GetOsPolicies() []ObjectStoragePolicy {
	if o == nil {
		var ret []ObjectStoragePolicy
		return ret
	}

	return o.OsPolicies
}

// GetOsPoliciesOk returns a tuple with the OsPolicies field value
// and a boolean to check if the value has been set.
func (o *ObjectStoragePoliciesResp) GetOsPoliciesOk() ([]ObjectStoragePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsPolicies, true
}

// SetOsPolicies sets field value
func (o *ObjectStoragePoliciesResp) SetOsPolicies(v []ObjectStoragePolicy) {
	o.OsPolicies = v
}

func (o ObjectStoragePoliciesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStoragePoliciesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_policies"] = o.OsPolicies
	return toSerialize, nil
}

func (o *ObjectStoragePoliciesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_policies",
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

	varObjectStoragePoliciesResp := _ObjectStoragePoliciesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectStoragePoliciesResp)

	if err != nil {
		return err
	}

	*o = ObjectStoragePoliciesResp(varObjectStoragePoliciesResp)

	return err
}

type NullableObjectStoragePoliciesResp struct {
	value *ObjectStoragePoliciesResp
	isSet bool
}

func (v NullableObjectStoragePoliciesResp) Get() *ObjectStoragePoliciesResp {
	return v.value
}

func (v *NullableObjectStoragePoliciesResp) Set(val *ObjectStoragePoliciesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStoragePoliciesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStoragePoliciesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStoragePoliciesResp(val *ObjectStoragePoliciesResp) *NullableObjectStoragePoliciesResp {
	return &NullableObjectStoragePoliciesResp{value: val, isSet: true}
}

func (v NullableObjectStoragePoliciesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStoragePoliciesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


