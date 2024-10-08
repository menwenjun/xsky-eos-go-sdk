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

// checks if the MetadataServiceResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataServiceResp{}

// MetadataServiceResp struct for MetadataServiceResp
type MetadataServiceResp struct {
	MetadataService MetadataServiceRecord `json:"metadata_service"`
}

type _MetadataServiceResp MetadataServiceResp

// NewMetadataServiceResp instantiates a new MetadataServiceResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataServiceResp(metadataService MetadataServiceRecord) *MetadataServiceResp {
	this := MetadataServiceResp{}
	this.MetadataService = metadataService
	return &this
}

// NewMetadataServiceRespWithDefaults instantiates a new MetadataServiceResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataServiceRespWithDefaults() *MetadataServiceResp {
	this := MetadataServiceResp{}
	return &this
}

// GetMetadataService returns the MetadataService field value
func (o *MetadataServiceResp) GetMetadataService() MetadataServiceRecord {
	if o == nil {
		var ret MetadataServiceRecord
		return ret
	}

	return o.MetadataService
}

// GetMetadataServiceOk returns a tuple with the MetadataService field value
// and a boolean to check if the value has been set.
func (o *MetadataServiceResp) GetMetadataServiceOk() (*MetadataServiceRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataService, true
}

// SetMetadataService sets field value
func (o *MetadataServiceResp) SetMetadataService(v MetadataServiceRecord) {
	o.MetadataService = v
}

func (o MetadataServiceResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataServiceResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata_service"] = o.MetadataService
	return toSerialize, nil
}

func (o *MetadataServiceResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata_service",
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

	varMetadataServiceResp := _MetadataServiceResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetadataServiceResp)

	if err != nil {
		return err
	}

	*o = MetadataServiceResp(varMetadataServiceResp)

	return err
}

type NullableMetadataServiceResp struct {
	value *MetadataServiceResp
	isSet bool
}

func (v NullableMetadataServiceResp) Get() *MetadataServiceResp {
	return v.value
}

func (v *NullableMetadataServiceResp) Set(val *MetadataServiceResp) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataServiceResp) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataServiceResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataServiceResp(val *MetadataServiceResp) *NullableMetadataServiceResp {
	return &NullableMetadataServiceResp{value: val, isSet: true}
}

func (v NullableMetadataServiceResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataServiceResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


