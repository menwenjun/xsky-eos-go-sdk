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

// checks if the OSBucketMetadataSearchSetReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSBucketMetadataSearchSetReq{}

// OSBucketMetadataSearchSetReq struct for OSBucketMetadataSearchSetReq
type OSBucketMetadataSearchSetReq struct {
	OsBucket OSBucketMetadataSearchSetReqBucket `json:"os_bucket"`
}

type _OSBucketMetadataSearchSetReq OSBucketMetadataSearchSetReq

// NewOSBucketMetadataSearchSetReq instantiates a new OSBucketMetadataSearchSetReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSBucketMetadataSearchSetReq(osBucket OSBucketMetadataSearchSetReqBucket) *OSBucketMetadataSearchSetReq {
	this := OSBucketMetadataSearchSetReq{}
	this.OsBucket = osBucket
	return &this
}

// NewOSBucketMetadataSearchSetReqWithDefaults instantiates a new OSBucketMetadataSearchSetReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSBucketMetadataSearchSetReqWithDefaults() *OSBucketMetadataSearchSetReq {
	this := OSBucketMetadataSearchSetReq{}
	return &this
}

// GetOsBucket returns the OsBucket field value
func (o *OSBucketMetadataSearchSetReq) GetOsBucket() OSBucketMetadataSearchSetReqBucket {
	if o == nil {
		var ret OSBucketMetadataSearchSetReqBucket
		return ret
	}

	return o.OsBucket
}

// GetOsBucketOk returns a tuple with the OsBucket field value
// and a boolean to check if the value has been set.
func (o *OSBucketMetadataSearchSetReq) GetOsBucketOk() (*OSBucketMetadataSearchSetReqBucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsBucket, true
}

// SetOsBucket sets field value
func (o *OSBucketMetadataSearchSetReq) SetOsBucket(v OSBucketMetadataSearchSetReqBucket) {
	o.OsBucket = v
}

func (o OSBucketMetadataSearchSetReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSBucketMetadataSearchSetReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_bucket"] = o.OsBucket
	return toSerialize, nil
}

func (o *OSBucketMetadataSearchSetReq) UnmarshalJSON(data []byte) (err error) {
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

	varOSBucketMetadataSearchSetReq := _OSBucketMetadataSearchSetReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSBucketMetadataSearchSetReq)

	if err != nil {
		return err
	}

	*o = OSBucketMetadataSearchSetReq(varOSBucketMetadataSearchSetReq)

	return err
}

type NullableOSBucketMetadataSearchSetReq struct {
	value *OSBucketMetadataSearchSetReq
	isSet bool
}

func (v NullableOSBucketMetadataSearchSetReq) Get() *OSBucketMetadataSearchSetReq {
	return v.value
}

func (v *NullableOSBucketMetadataSearchSetReq) Set(val *OSBucketMetadataSearchSetReq) {
	v.value = val
	v.isSet = true
}

func (v NullableOSBucketMetadataSearchSetReq) IsSet() bool {
	return v.isSet
}

func (v *NullableOSBucketMetadataSearchSetReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSBucketMetadataSearchSetReq(val *OSBucketMetadataSearchSetReq) *NullableOSBucketMetadataSearchSetReq {
	return &NullableOSBucketMetadataSearchSetReq{value: val, isSet: true}
}

func (v NullableOSBucketMetadataSearchSetReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSBucketMetadataSearchSetReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


