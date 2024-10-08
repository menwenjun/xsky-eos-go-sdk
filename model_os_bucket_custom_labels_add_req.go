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

// checks if the OSBucketCustomLabelsAddReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSBucketCustomLabelsAddReq{}

// OSBucketCustomLabelsAddReq struct for OSBucketCustomLabelsAddReq
type OSBucketCustomLabelsAddReq struct {
	OsBucket OSBucketCustomLabelsAddReqBucket `json:"os_bucket"`
}

type _OSBucketCustomLabelsAddReq OSBucketCustomLabelsAddReq

// NewOSBucketCustomLabelsAddReq instantiates a new OSBucketCustomLabelsAddReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSBucketCustomLabelsAddReq(osBucket OSBucketCustomLabelsAddReqBucket) *OSBucketCustomLabelsAddReq {
	this := OSBucketCustomLabelsAddReq{}
	this.OsBucket = osBucket
	return &this
}

// NewOSBucketCustomLabelsAddReqWithDefaults instantiates a new OSBucketCustomLabelsAddReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSBucketCustomLabelsAddReqWithDefaults() *OSBucketCustomLabelsAddReq {
	this := OSBucketCustomLabelsAddReq{}
	return &this
}

// GetOsBucket returns the OsBucket field value
func (o *OSBucketCustomLabelsAddReq) GetOsBucket() OSBucketCustomLabelsAddReqBucket {
	if o == nil {
		var ret OSBucketCustomLabelsAddReqBucket
		return ret
	}

	return o.OsBucket
}

// GetOsBucketOk returns a tuple with the OsBucket field value
// and a boolean to check if the value has been set.
func (o *OSBucketCustomLabelsAddReq) GetOsBucketOk() (*OSBucketCustomLabelsAddReqBucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsBucket, true
}

// SetOsBucket sets field value
func (o *OSBucketCustomLabelsAddReq) SetOsBucket(v OSBucketCustomLabelsAddReqBucket) {
	o.OsBucket = v
}

func (o OSBucketCustomLabelsAddReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSBucketCustomLabelsAddReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_bucket"] = o.OsBucket
	return toSerialize, nil
}

func (o *OSBucketCustomLabelsAddReq) UnmarshalJSON(data []byte) (err error) {
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

	varOSBucketCustomLabelsAddReq := _OSBucketCustomLabelsAddReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSBucketCustomLabelsAddReq)

	if err != nil {
		return err
	}

	*o = OSBucketCustomLabelsAddReq(varOSBucketCustomLabelsAddReq)

	return err
}

type NullableOSBucketCustomLabelsAddReq struct {
	value *OSBucketCustomLabelsAddReq
	isSet bool
}

func (v NullableOSBucketCustomLabelsAddReq) Get() *OSBucketCustomLabelsAddReq {
	return v.value
}

func (v *NullableOSBucketCustomLabelsAddReq) Set(val *OSBucketCustomLabelsAddReq) {
	v.value = val
	v.isSet = true
}

func (v NullableOSBucketCustomLabelsAddReq) IsSet() bool {
	return v.isSet
}

func (v *NullableOSBucketCustomLabelsAddReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSBucketCustomLabelsAddReq(val *OSBucketCustomLabelsAddReq) *NullableOSBucketCustomLabelsAddReq {
	return &NullableOSBucketCustomLabelsAddReq{value: val, isSet: true}
}

func (v NullableOSBucketCustomLabelsAddReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSBucketCustomLabelsAddReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


