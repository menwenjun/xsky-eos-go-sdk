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

// checks if the MetadataClusterPredictionsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataClusterPredictionsResp{}

// MetadataClusterPredictionsResp struct for MetadataClusterPredictionsResp
type MetadataClusterPredictionsResp struct {
	// metadata cluster predictions
	MetadataClusterPredictions []MetadataClusterPrediction `json:"metadata_cluster_predictions"`
}

type _MetadataClusterPredictionsResp MetadataClusterPredictionsResp

// NewMetadataClusterPredictionsResp instantiates a new MetadataClusterPredictionsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataClusterPredictionsResp(metadataClusterPredictions []MetadataClusterPrediction) *MetadataClusterPredictionsResp {
	this := MetadataClusterPredictionsResp{}
	this.MetadataClusterPredictions = metadataClusterPredictions
	return &this
}

// NewMetadataClusterPredictionsRespWithDefaults instantiates a new MetadataClusterPredictionsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataClusterPredictionsRespWithDefaults() *MetadataClusterPredictionsResp {
	this := MetadataClusterPredictionsResp{}
	return &this
}

// GetMetadataClusterPredictions returns the MetadataClusterPredictions field value
func (o *MetadataClusterPredictionsResp) GetMetadataClusterPredictions() []MetadataClusterPrediction {
	if o == nil {
		var ret []MetadataClusterPrediction
		return ret
	}

	return o.MetadataClusterPredictions
}

// GetMetadataClusterPredictionsOk returns a tuple with the MetadataClusterPredictions field value
// and a boolean to check if the value has been set.
func (o *MetadataClusterPredictionsResp) GetMetadataClusterPredictionsOk() ([]MetadataClusterPrediction, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataClusterPredictions, true
}

// SetMetadataClusterPredictions sets field value
func (o *MetadataClusterPredictionsResp) SetMetadataClusterPredictions(v []MetadataClusterPrediction) {
	o.MetadataClusterPredictions = v
}

func (o MetadataClusterPredictionsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataClusterPredictionsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata_cluster_predictions"] = o.MetadataClusterPredictions
	return toSerialize, nil
}

func (o *MetadataClusterPredictionsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata_cluster_predictions",
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

	varMetadataClusterPredictionsResp := _MetadataClusterPredictionsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetadataClusterPredictionsResp)

	if err != nil {
		return err
	}

	*o = MetadataClusterPredictionsResp(varMetadataClusterPredictionsResp)

	return err
}

type NullableMetadataClusterPredictionsResp struct {
	value *MetadataClusterPredictionsResp
	isSet bool
}

func (v NullableMetadataClusterPredictionsResp) Get() *MetadataClusterPredictionsResp {
	return v.value
}

func (v *NullableMetadataClusterPredictionsResp) Set(val *MetadataClusterPredictionsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataClusterPredictionsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataClusterPredictionsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataClusterPredictionsResp(val *MetadataClusterPredictionsResp) *NullableMetadataClusterPredictionsResp {
	return &NullableMetadataClusterPredictionsResp{value: val, isSet: true}
}

func (v NullableMetadataClusterPredictionsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataClusterPredictionsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


