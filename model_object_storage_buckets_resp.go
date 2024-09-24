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

// checks if the ObjectStorageBucketsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageBucketsResp{}

// ObjectStorageBucketsResp struct for ObjectStorageBucketsResp
type ObjectStorageBucketsResp struct {
	// records of object storage bucket
	OsBuckets []ObjectStorageBucketRecord `json:"os_buckets"`
}

type _ObjectStorageBucketsResp ObjectStorageBucketsResp

// NewObjectStorageBucketsResp instantiates a new ObjectStorageBucketsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageBucketsResp(osBuckets []ObjectStorageBucketRecord) *ObjectStorageBucketsResp {
	this := ObjectStorageBucketsResp{}
	this.OsBuckets = osBuckets
	return &this
}

// NewObjectStorageBucketsRespWithDefaults instantiates a new ObjectStorageBucketsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageBucketsRespWithDefaults() *ObjectStorageBucketsResp {
	this := ObjectStorageBucketsResp{}
	return &this
}

// GetOsBuckets returns the OsBuckets field value
func (o *ObjectStorageBucketsResp) GetOsBuckets() []ObjectStorageBucketRecord {
	if o == nil {
		var ret []ObjectStorageBucketRecord
		return ret
	}

	return o.OsBuckets
}

// GetOsBucketsOk returns a tuple with the OsBuckets field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketsResp) GetOsBucketsOk() ([]ObjectStorageBucketRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsBuckets, true
}

// SetOsBuckets sets field value
func (o *ObjectStorageBucketsResp) SetOsBuckets(v []ObjectStorageBucketRecord) {
	o.OsBuckets = v
}

func (o ObjectStorageBucketsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageBucketsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_buckets"] = o.OsBuckets
	return toSerialize, nil
}

func (o *ObjectStorageBucketsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_buckets",
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

	varObjectStorageBucketsResp := _ObjectStorageBucketsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectStorageBucketsResp)

	if err != nil {
		return err
	}

	*o = ObjectStorageBucketsResp(varObjectStorageBucketsResp)

	return err
}

type NullableObjectStorageBucketsResp struct {
	value *ObjectStorageBucketsResp
	isSet bool
}

func (v NullableObjectStorageBucketsResp) Get() *ObjectStorageBucketsResp {
	return v.value
}

func (v *NullableObjectStorageBucketsResp) Set(val *ObjectStorageBucketsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageBucketsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageBucketsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageBucketsResp(val *ObjectStorageBucketsResp) *NullableObjectStorageBucketsResp {
	return &NullableObjectStorageBucketsResp{value: val, isSet: true}
}

func (v NullableObjectStorageBucketsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageBucketsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


