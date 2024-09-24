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

// checks if the DfsStorageClassPredictionsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsStorageClassPredictionsResp{}

// DfsStorageClassPredictionsResp struct for DfsStorageClassPredictionsResp
type DfsStorageClassPredictionsResp struct {
	// class predictions
	DfsStorageClassPredictions []DfsTierPrediction `json:"dfs_storage_class_predictions"`
}

type _DfsStorageClassPredictionsResp DfsStorageClassPredictionsResp

// NewDfsStorageClassPredictionsResp instantiates a new DfsStorageClassPredictionsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsStorageClassPredictionsResp(dfsStorageClassPredictions []DfsTierPrediction) *DfsStorageClassPredictionsResp {
	this := DfsStorageClassPredictionsResp{}
	this.DfsStorageClassPredictions = dfsStorageClassPredictions
	return &this
}

// NewDfsStorageClassPredictionsRespWithDefaults instantiates a new DfsStorageClassPredictionsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsStorageClassPredictionsRespWithDefaults() *DfsStorageClassPredictionsResp {
	this := DfsStorageClassPredictionsResp{}
	return &this
}

// GetDfsStorageClassPredictions returns the DfsStorageClassPredictions field value
func (o *DfsStorageClassPredictionsResp) GetDfsStorageClassPredictions() []DfsTierPrediction {
	if o == nil {
		var ret []DfsTierPrediction
		return ret
	}

	return o.DfsStorageClassPredictions
}

// GetDfsStorageClassPredictionsOk returns a tuple with the DfsStorageClassPredictions field value
// and a boolean to check if the value has been set.
func (o *DfsStorageClassPredictionsResp) GetDfsStorageClassPredictionsOk() ([]DfsTierPrediction, bool) {
	if o == nil {
		return nil, false
	}
	return o.DfsStorageClassPredictions, true
}

// SetDfsStorageClassPredictions sets field value
func (o *DfsStorageClassPredictionsResp) SetDfsStorageClassPredictions(v []DfsTierPrediction) {
	o.DfsStorageClassPredictions = v
}

func (o DfsStorageClassPredictionsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsStorageClassPredictionsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_storage_class_predictions"] = o.DfsStorageClassPredictions
	return toSerialize, nil
}

func (o *DfsStorageClassPredictionsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_storage_class_predictions",
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

	varDfsStorageClassPredictionsResp := _DfsStorageClassPredictionsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsStorageClassPredictionsResp)

	if err != nil {
		return err
	}

	*o = DfsStorageClassPredictionsResp(varDfsStorageClassPredictionsResp)

	return err
}

type NullableDfsStorageClassPredictionsResp struct {
	value *DfsStorageClassPredictionsResp
	isSet bool
}

func (v NullableDfsStorageClassPredictionsResp) Get() *DfsStorageClassPredictionsResp {
	return v.value
}

func (v *NullableDfsStorageClassPredictionsResp) Set(val *DfsStorageClassPredictionsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsStorageClassPredictionsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsStorageClassPredictionsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsStorageClassPredictionsResp(val *DfsStorageClassPredictionsResp) *NullableDfsStorageClassPredictionsResp {
	return &NullableDfsStorageClassPredictionsResp{value: val, isSet: true}
}

func (v NullableDfsStorageClassPredictionsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsStorageClassPredictionsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


