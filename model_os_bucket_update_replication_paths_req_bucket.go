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

// checks if the OSBucketUpdateReplicationPathsReqBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSBucketUpdateReplicationPathsReqBucket{}

// OSBucketUpdateReplicationPathsReqBucket struct for OSBucketUpdateReplicationPathsReqBucket
type OSBucketUpdateReplicationPathsReqBucket struct {
	OsReplicationPathUuids []string `json:"os_replication_path_uuids"`
}

type _OSBucketUpdateReplicationPathsReqBucket OSBucketUpdateReplicationPathsReqBucket

// NewOSBucketUpdateReplicationPathsReqBucket instantiates a new OSBucketUpdateReplicationPathsReqBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSBucketUpdateReplicationPathsReqBucket(osReplicationPathUuids []string) *OSBucketUpdateReplicationPathsReqBucket {
	this := OSBucketUpdateReplicationPathsReqBucket{}
	this.OsReplicationPathUuids = osReplicationPathUuids
	return &this
}

// NewOSBucketUpdateReplicationPathsReqBucketWithDefaults instantiates a new OSBucketUpdateReplicationPathsReqBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSBucketUpdateReplicationPathsReqBucketWithDefaults() *OSBucketUpdateReplicationPathsReqBucket {
	this := OSBucketUpdateReplicationPathsReqBucket{}
	return &this
}

// GetOsReplicationPathUuids returns the OsReplicationPathUuids field value
func (o *OSBucketUpdateReplicationPathsReqBucket) GetOsReplicationPathUuids() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OsReplicationPathUuids
}

// GetOsReplicationPathUuidsOk returns a tuple with the OsReplicationPathUuids field value
// and a boolean to check if the value has been set.
func (o *OSBucketUpdateReplicationPathsReqBucket) GetOsReplicationPathUuidsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsReplicationPathUuids, true
}

// SetOsReplicationPathUuids sets field value
func (o *OSBucketUpdateReplicationPathsReqBucket) SetOsReplicationPathUuids(v []string) {
	o.OsReplicationPathUuids = v
}

func (o OSBucketUpdateReplicationPathsReqBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSBucketUpdateReplicationPathsReqBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_replication_path_uuids"] = o.OsReplicationPathUuids
	return toSerialize, nil
}

func (o *OSBucketUpdateReplicationPathsReqBucket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_replication_path_uuids",
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

	varOSBucketUpdateReplicationPathsReqBucket := _OSBucketUpdateReplicationPathsReqBucket{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSBucketUpdateReplicationPathsReqBucket)

	if err != nil {
		return err
	}

	*o = OSBucketUpdateReplicationPathsReqBucket(varOSBucketUpdateReplicationPathsReqBucket)

	return err
}

type NullableOSBucketUpdateReplicationPathsReqBucket struct {
	value *OSBucketUpdateReplicationPathsReqBucket
	isSet bool
}

func (v NullableOSBucketUpdateReplicationPathsReqBucket) Get() *OSBucketUpdateReplicationPathsReqBucket {
	return v.value
}

func (v *NullableOSBucketUpdateReplicationPathsReqBucket) Set(val *OSBucketUpdateReplicationPathsReqBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableOSBucketUpdateReplicationPathsReqBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableOSBucketUpdateReplicationPathsReqBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSBucketUpdateReplicationPathsReqBucket(val *OSBucketUpdateReplicationPathsReqBucket) *NullableOSBucketUpdateReplicationPathsReqBucket {
	return &NullableOSBucketUpdateReplicationPathsReqBucket{value: val, isSet: true}
}

func (v NullableOSBucketUpdateReplicationPathsReqBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSBucketUpdateReplicationPathsReqBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


