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

// checks if the MetadataClusterResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataClusterResp{}

// MetadataClusterResp struct for MetadataClusterResp
type MetadataClusterResp struct {
	MetadataCluster MetadataClusterRecord `json:"metadata_cluster"`
}

type _MetadataClusterResp MetadataClusterResp

// NewMetadataClusterResp instantiates a new MetadataClusterResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataClusterResp(metadataCluster MetadataClusterRecord) *MetadataClusterResp {
	this := MetadataClusterResp{}
	this.MetadataCluster = metadataCluster
	return &this
}

// NewMetadataClusterRespWithDefaults instantiates a new MetadataClusterResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataClusterRespWithDefaults() *MetadataClusterResp {
	this := MetadataClusterResp{}
	return &this
}

// GetMetadataCluster returns the MetadataCluster field value
func (o *MetadataClusterResp) GetMetadataCluster() MetadataClusterRecord {
	if o == nil {
		var ret MetadataClusterRecord
		return ret
	}

	return o.MetadataCluster
}

// GetMetadataClusterOk returns a tuple with the MetadataCluster field value
// and a boolean to check if the value has been set.
func (o *MetadataClusterResp) GetMetadataClusterOk() (*MetadataClusterRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataCluster, true
}

// SetMetadataCluster sets field value
func (o *MetadataClusterResp) SetMetadataCluster(v MetadataClusterRecord) {
	o.MetadataCluster = v
}

func (o MetadataClusterResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataClusterResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata_cluster"] = o.MetadataCluster
	return toSerialize, nil
}

func (o *MetadataClusterResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata_cluster",
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

	varMetadataClusterResp := _MetadataClusterResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetadataClusterResp)

	if err != nil {
		return err
	}

	*o = MetadataClusterResp(varMetadataClusterResp)

	return err
}

type NullableMetadataClusterResp struct {
	value *MetadataClusterResp
	isSet bool
}

func (v NullableMetadataClusterResp) Get() *MetadataClusterResp {
	return v.value
}

func (v *NullableMetadataClusterResp) Set(val *MetadataClusterResp) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataClusterResp) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataClusterResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataClusterResp(val *MetadataClusterResp) *NullableMetadataClusterResp {
	return &NullableMetadataClusterResp{value: val, isSet: true}
}

func (v NullableMetadataClusterResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataClusterResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


