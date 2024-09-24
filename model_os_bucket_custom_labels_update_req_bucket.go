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

// checks if the OSBucketCustomLabelsUpdateReqBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSBucketCustomLabelsUpdateReqBucket{}

// OSBucketCustomLabelsUpdateReqBucket struct for OSBucketCustomLabelsUpdateReqBucket
type OSBucketCustomLabelsUpdateReqBucket struct {
	Labels []OSBucketCustomLabelsUpdateReqBucketLabelsElt `json:"labels"`
}

type _OSBucketCustomLabelsUpdateReqBucket OSBucketCustomLabelsUpdateReqBucket

// NewOSBucketCustomLabelsUpdateReqBucket instantiates a new OSBucketCustomLabelsUpdateReqBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSBucketCustomLabelsUpdateReqBucket(labels []OSBucketCustomLabelsUpdateReqBucketLabelsElt) *OSBucketCustomLabelsUpdateReqBucket {
	this := OSBucketCustomLabelsUpdateReqBucket{}
	this.Labels = labels
	return &this
}

// NewOSBucketCustomLabelsUpdateReqBucketWithDefaults instantiates a new OSBucketCustomLabelsUpdateReqBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSBucketCustomLabelsUpdateReqBucketWithDefaults() *OSBucketCustomLabelsUpdateReqBucket {
	this := OSBucketCustomLabelsUpdateReqBucket{}
	return &this
}

// GetLabels returns the Labels field value
func (o *OSBucketCustomLabelsUpdateReqBucket) GetLabels() []OSBucketCustomLabelsUpdateReqBucketLabelsElt {
	if o == nil {
		var ret []OSBucketCustomLabelsUpdateReqBucketLabelsElt
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *OSBucketCustomLabelsUpdateReqBucket) GetLabelsOk() ([]OSBucketCustomLabelsUpdateReqBucketLabelsElt, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *OSBucketCustomLabelsUpdateReqBucket) SetLabels(v []OSBucketCustomLabelsUpdateReqBucketLabelsElt) {
	o.Labels = v
}

func (o OSBucketCustomLabelsUpdateReqBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSBucketCustomLabelsUpdateReqBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
}

func (o *OSBucketCustomLabelsUpdateReqBucket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"labels",
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

	varOSBucketCustomLabelsUpdateReqBucket := _OSBucketCustomLabelsUpdateReqBucket{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSBucketCustomLabelsUpdateReqBucket)

	if err != nil {
		return err
	}

	*o = OSBucketCustomLabelsUpdateReqBucket(varOSBucketCustomLabelsUpdateReqBucket)

	return err
}

type NullableOSBucketCustomLabelsUpdateReqBucket struct {
	value *OSBucketCustomLabelsUpdateReqBucket
	isSet bool
}

func (v NullableOSBucketCustomLabelsUpdateReqBucket) Get() *OSBucketCustomLabelsUpdateReqBucket {
	return v.value
}

func (v *NullableOSBucketCustomLabelsUpdateReqBucket) Set(val *OSBucketCustomLabelsUpdateReqBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableOSBucketCustomLabelsUpdateReqBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableOSBucketCustomLabelsUpdateReqBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSBucketCustomLabelsUpdateReqBucket(val *OSBucketCustomLabelsUpdateReqBucket) *NullableOSBucketCustomLabelsUpdateReqBucket {
	return &NullableOSBucketCustomLabelsUpdateReqBucket{value: val, isSet: true}
}

func (v NullableOSBucketCustomLabelsUpdateReqBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSBucketCustomLabelsUpdateReqBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


