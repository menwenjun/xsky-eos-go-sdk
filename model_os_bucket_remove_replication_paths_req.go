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

// checks if the OSBucketRemoveReplicationPathsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSBucketRemoveReplicationPathsReq{}

// OSBucketRemoveReplicationPathsReq struct for OSBucketRemoveReplicationPathsReq
type OSBucketRemoveReplicationPathsReq struct {
	OsBucket OSBucketRemoveReplicationPathsReqBucket `json:"os_bucket"`
}

type _OSBucketRemoveReplicationPathsReq OSBucketRemoveReplicationPathsReq

// NewOSBucketRemoveReplicationPathsReq instantiates a new OSBucketRemoveReplicationPathsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSBucketRemoveReplicationPathsReq(osBucket OSBucketRemoveReplicationPathsReqBucket) *OSBucketRemoveReplicationPathsReq {
	this := OSBucketRemoveReplicationPathsReq{}
	this.OsBucket = osBucket
	return &this
}

// NewOSBucketRemoveReplicationPathsReqWithDefaults instantiates a new OSBucketRemoveReplicationPathsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSBucketRemoveReplicationPathsReqWithDefaults() *OSBucketRemoveReplicationPathsReq {
	this := OSBucketRemoveReplicationPathsReq{}
	return &this
}

// GetOsBucket returns the OsBucket field value
func (o *OSBucketRemoveReplicationPathsReq) GetOsBucket() OSBucketRemoveReplicationPathsReqBucket {
	if o == nil {
		var ret OSBucketRemoveReplicationPathsReqBucket
		return ret
	}

	return o.OsBucket
}

// GetOsBucketOk returns a tuple with the OsBucket field value
// and a boolean to check if the value has been set.
func (o *OSBucketRemoveReplicationPathsReq) GetOsBucketOk() (*OSBucketRemoveReplicationPathsReqBucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsBucket, true
}

// SetOsBucket sets field value
func (o *OSBucketRemoveReplicationPathsReq) SetOsBucket(v OSBucketRemoveReplicationPathsReqBucket) {
	o.OsBucket = v
}

func (o OSBucketRemoveReplicationPathsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSBucketRemoveReplicationPathsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_bucket"] = o.OsBucket
	return toSerialize, nil
}

func (o *OSBucketRemoveReplicationPathsReq) UnmarshalJSON(data []byte) (err error) {
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

	varOSBucketRemoveReplicationPathsReq := _OSBucketRemoveReplicationPathsReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSBucketRemoveReplicationPathsReq)

	if err != nil {
		return err
	}

	*o = OSBucketRemoveReplicationPathsReq(varOSBucketRemoveReplicationPathsReq)

	return err
}

type NullableOSBucketRemoveReplicationPathsReq struct {
	value *OSBucketRemoveReplicationPathsReq
	isSet bool
}

func (v NullableOSBucketRemoveReplicationPathsReq) Get() *OSBucketRemoveReplicationPathsReq {
	return v.value
}

func (v *NullableOSBucketRemoveReplicationPathsReq) Set(val *OSBucketRemoveReplicationPathsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableOSBucketRemoveReplicationPathsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableOSBucketRemoveReplicationPathsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSBucketRemoveReplicationPathsReq(val *OSBucketRemoveReplicationPathsReq) *NullableOSBucketRemoveReplicationPathsReq {
	return &NullableOSBucketRemoveReplicationPathsReq{value: val, isSet: true}
}

func (v NullableOSBucketRemoveReplicationPathsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSBucketRemoveReplicationPathsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


