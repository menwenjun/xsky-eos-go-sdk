/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ObjectStorageBucketCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageBucketCreateReq{}

// ObjectStorageBucketCreateReq struct for ObjectStorageBucketCreateReq
type ObjectStorageBucketCreateReq struct {
	OsBucket *ObjectStorageBucketCreateReqBucket `json:"os_bucket,omitempty"`
}

// NewObjectStorageBucketCreateReq instantiates a new ObjectStorageBucketCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageBucketCreateReq() *ObjectStorageBucketCreateReq {
	this := ObjectStorageBucketCreateReq{}
	return &this
}

// NewObjectStorageBucketCreateReqWithDefaults instantiates a new ObjectStorageBucketCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageBucketCreateReqWithDefaults() *ObjectStorageBucketCreateReq {
	this := ObjectStorageBucketCreateReq{}
	return &this
}

// GetOsBucket returns the OsBucket field value if set, zero value otherwise.
func (o *ObjectStorageBucketCreateReq) GetOsBucket() ObjectStorageBucketCreateReqBucket {
	if o == nil || IsNil(o.OsBucket) {
		var ret ObjectStorageBucketCreateReqBucket
		return ret
	}
	return *o.OsBucket
}

// GetOsBucketOk returns a tuple with the OsBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageBucketCreateReq) GetOsBucketOk() (*ObjectStorageBucketCreateReqBucket, bool) {
	if o == nil || IsNil(o.OsBucket) {
		return nil, false
	}
	return o.OsBucket, true
}

// HasOsBucket returns a boolean if a field has been set.
func (o *ObjectStorageBucketCreateReq) HasOsBucket() bool {
	if o != nil && !IsNil(o.OsBucket) {
		return true
	}

	return false
}

// SetOsBucket gets a reference to the given ObjectStorageBucketCreateReqBucket and assigns it to the OsBucket field.
func (o *ObjectStorageBucketCreateReq) SetOsBucket(v ObjectStorageBucketCreateReqBucket) {
	o.OsBucket = &v
}

func (o ObjectStorageBucketCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageBucketCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OsBucket) {
		toSerialize["os_bucket"] = o.OsBucket
	}
	return toSerialize, nil
}

type NullableObjectStorageBucketCreateReq struct {
	value *ObjectStorageBucketCreateReq
	isSet bool
}

func (v NullableObjectStorageBucketCreateReq) Get() *ObjectStorageBucketCreateReq {
	return v.value
}

func (v *NullableObjectStorageBucketCreateReq) Set(val *ObjectStorageBucketCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageBucketCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageBucketCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageBucketCreateReq(val *ObjectStorageBucketCreateReq) *NullableObjectStorageBucketCreateReq {
	return &NullableObjectStorageBucketCreateReq{value: val, isSet: true}
}

func (v NullableObjectStorageBucketCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageBucketCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


