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

// checks if the MetadataServicesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataServicesResp{}

// MetadataServicesResp struct for MetadataServicesResp
type MetadataServicesResp struct {
	MetadataServices []MetadataServiceRecord `json:"metadata_services"`
}

type _MetadataServicesResp MetadataServicesResp

// NewMetadataServicesResp instantiates a new MetadataServicesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataServicesResp(metadataServices []MetadataServiceRecord) *MetadataServicesResp {
	this := MetadataServicesResp{}
	this.MetadataServices = metadataServices
	return &this
}

// NewMetadataServicesRespWithDefaults instantiates a new MetadataServicesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataServicesRespWithDefaults() *MetadataServicesResp {
	this := MetadataServicesResp{}
	return &this
}

// GetMetadataServices returns the MetadataServices field value
func (o *MetadataServicesResp) GetMetadataServices() []MetadataServiceRecord {
	if o == nil {
		var ret []MetadataServiceRecord
		return ret
	}

	return o.MetadataServices
}

// GetMetadataServicesOk returns a tuple with the MetadataServices field value
// and a boolean to check if the value has been set.
func (o *MetadataServicesResp) GetMetadataServicesOk() ([]MetadataServiceRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataServices, true
}

// SetMetadataServices sets field value
func (o *MetadataServicesResp) SetMetadataServices(v []MetadataServiceRecord) {
	o.MetadataServices = v
}

func (o MetadataServicesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataServicesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata_services"] = o.MetadataServices
	return toSerialize, nil
}

func (o *MetadataServicesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata_services",
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

	varMetadataServicesResp := _MetadataServicesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetadataServicesResp)

	if err != nil {
		return err
	}

	*o = MetadataServicesResp(varMetadataServicesResp)

	return err
}

type NullableMetadataServicesResp struct {
	value *MetadataServicesResp
	isSet bool
}

func (v NullableMetadataServicesResp) Get() *MetadataServicesResp {
	return v.value
}

func (v *NullableMetadataServicesResp) Set(val *MetadataServicesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataServicesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataServicesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataServicesResp(val *MetadataServicesResp) *NullableMetadataServicesResp {
	return &NullableMetadataServicesResp{value: val, isSet: true}
}

func (v NullableMetadataServicesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataServicesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


