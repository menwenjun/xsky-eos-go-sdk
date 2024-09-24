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

// checks if the MetadataClustersResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataClustersResp{}

// MetadataClustersResp struct for MetadataClustersResp
type MetadataClustersResp struct {
	MetadataClusters []MetadataClusterRecord `json:"metadata_clusters"`
}

type _MetadataClustersResp MetadataClustersResp

// NewMetadataClustersResp instantiates a new MetadataClustersResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataClustersResp(metadataClusters []MetadataClusterRecord) *MetadataClustersResp {
	this := MetadataClustersResp{}
	this.MetadataClusters = metadataClusters
	return &this
}

// NewMetadataClustersRespWithDefaults instantiates a new MetadataClustersResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataClustersRespWithDefaults() *MetadataClustersResp {
	this := MetadataClustersResp{}
	return &this
}

// GetMetadataClusters returns the MetadataClusters field value
func (o *MetadataClustersResp) GetMetadataClusters() []MetadataClusterRecord {
	if o == nil {
		var ret []MetadataClusterRecord
		return ret
	}

	return o.MetadataClusters
}

// GetMetadataClustersOk returns a tuple with the MetadataClusters field value
// and a boolean to check if the value has been set.
func (o *MetadataClustersResp) GetMetadataClustersOk() ([]MetadataClusterRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataClusters, true
}

// SetMetadataClusters sets field value
func (o *MetadataClustersResp) SetMetadataClusters(v []MetadataClusterRecord) {
	o.MetadataClusters = v
}

func (o MetadataClustersResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataClustersResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata_clusters"] = o.MetadataClusters
	return toSerialize, nil
}

func (o *MetadataClustersResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata_clusters",
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

	varMetadataClustersResp := _MetadataClustersResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetadataClustersResp)

	if err != nil {
		return err
	}

	*o = MetadataClustersResp(varMetadataClustersResp)

	return err
}

type NullableMetadataClustersResp struct {
	value *MetadataClustersResp
	isSet bool
}

func (v NullableMetadataClustersResp) Get() *MetadataClustersResp {
	return v.value
}

func (v *NullableMetadataClustersResp) Set(val *MetadataClustersResp) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataClustersResp) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataClustersResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataClustersResp(val *MetadataClustersResp) *NullableMetadataClustersResp {
	return &NullableMetadataClustersResp{value: val, isSet: true}
}

func (v NullableMetadataClustersResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataClustersResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


