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

// checks if the OSBucketSetAccessLoggingReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSBucketSetAccessLoggingReq{}

// OSBucketSetAccessLoggingReq struct for OSBucketSetAccessLoggingReq
type OSBucketSetAccessLoggingReq struct {
	OsBucket OSBucketSetAccessLoggingReqBucket `json:"os_bucket"`
}

type _OSBucketSetAccessLoggingReq OSBucketSetAccessLoggingReq

// NewOSBucketSetAccessLoggingReq instantiates a new OSBucketSetAccessLoggingReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSBucketSetAccessLoggingReq(osBucket OSBucketSetAccessLoggingReqBucket) *OSBucketSetAccessLoggingReq {
	this := OSBucketSetAccessLoggingReq{}
	this.OsBucket = osBucket
	return &this
}

// NewOSBucketSetAccessLoggingReqWithDefaults instantiates a new OSBucketSetAccessLoggingReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSBucketSetAccessLoggingReqWithDefaults() *OSBucketSetAccessLoggingReq {
	this := OSBucketSetAccessLoggingReq{}
	return &this
}

// GetOsBucket returns the OsBucket field value
func (o *OSBucketSetAccessLoggingReq) GetOsBucket() OSBucketSetAccessLoggingReqBucket {
	if o == nil {
		var ret OSBucketSetAccessLoggingReqBucket
		return ret
	}

	return o.OsBucket
}

// GetOsBucketOk returns a tuple with the OsBucket field value
// and a boolean to check if the value has been set.
func (o *OSBucketSetAccessLoggingReq) GetOsBucketOk() (*OSBucketSetAccessLoggingReqBucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsBucket, true
}

// SetOsBucket sets field value
func (o *OSBucketSetAccessLoggingReq) SetOsBucket(v OSBucketSetAccessLoggingReqBucket) {
	o.OsBucket = v
}

func (o OSBucketSetAccessLoggingReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSBucketSetAccessLoggingReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_bucket"] = o.OsBucket
	return toSerialize, nil
}

func (o *OSBucketSetAccessLoggingReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_bucket",
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

	varOSBucketSetAccessLoggingReq := _OSBucketSetAccessLoggingReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSBucketSetAccessLoggingReq)

	if err != nil {
		return err
	}

	*o = OSBucketSetAccessLoggingReq(varOSBucketSetAccessLoggingReq)

	return err
}

type NullableOSBucketSetAccessLoggingReq struct {
	value *OSBucketSetAccessLoggingReq
	isSet bool
}

func (v NullableOSBucketSetAccessLoggingReq) Get() *OSBucketSetAccessLoggingReq {
	return v.value
}

func (v *NullableOSBucketSetAccessLoggingReq) Set(val *OSBucketSetAccessLoggingReq) {
	v.value = val
	v.isSet = true
}

func (v NullableOSBucketSetAccessLoggingReq) IsSet() bool {
	return v.isSet
}

func (v *NullableOSBucketSetAccessLoggingReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSBucketSetAccessLoggingReq(val *OSBucketSetAccessLoggingReq) *NullableOSBucketSetAccessLoggingReq {
	return &NullableOSBucketSetAccessLoggingReq{value: val, isSet: true}
}

func (v NullableOSBucketSetAccessLoggingReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSBucketSetAccessLoggingReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


