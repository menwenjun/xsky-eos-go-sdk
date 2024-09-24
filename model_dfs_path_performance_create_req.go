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

// checks if the DfsPathPerformanceCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsPathPerformanceCreateReq{}

// DfsPathPerformanceCreateReq struct for DfsPathPerformanceCreateReq
type DfsPathPerformanceCreateReq struct {
	DfsPathPerformance DfsPathPerformanceCreateReqPathPerformance `json:"dfs_path_performance"`
}

type _DfsPathPerformanceCreateReq DfsPathPerformanceCreateReq

// NewDfsPathPerformanceCreateReq instantiates a new DfsPathPerformanceCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsPathPerformanceCreateReq(dfsPathPerformance DfsPathPerformanceCreateReqPathPerformance) *DfsPathPerformanceCreateReq {
	this := DfsPathPerformanceCreateReq{}
	this.DfsPathPerformance = dfsPathPerformance
	return &this
}

// NewDfsPathPerformanceCreateReqWithDefaults instantiates a new DfsPathPerformanceCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsPathPerformanceCreateReqWithDefaults() *DfsPathPerformanceCreateReq {
	this := DfsPathPerformanceCreateReq{}
	return &this
}

// GetDfsPathPerformance returns the DfsPathPerformance field value
func (o *DfsPathPerformanceCreateReq) GetDfsPathPerformance() DfsPathPerformanceCreateReqPathPerformance {
	if o == nil {
		var ret DfsPathPerformanceCreateReqPathPerformance
		return ret
	}

	return o.DfsPathPerformance
}

// GetDfsPathPerformanceOk returns a tuple with the DfsPathPerformance field value
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceCreateReq) GetDfsPathPerformanceOk() (*DfsPathPerformanceCreateReqPathPerformance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsPathPerformance, true
}

// SetDfsPathPerformance sets field value
func (o *DfsPathPerformanceCreateReq) SetDfsPathPerformance(v DfsPathPerformanceCreateReqPathPerformance) {
	o.DfsPathPerformance = v
}

func (o DfsPathPerformanceCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsPathPerformanceCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_path_performance"] = o.DfsPathPerformance
	return toSerialize, nil
}

func (o *DfsPathPerformanceCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_path_performance",
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

	varDfsPathPerformanceCreateReq := _DfsPathPerformanceCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsPathPerformanceCreateReq)

	if err != nil {
		return err
	}

	*o = DfsPathPerformanceCreateReq(varDfsPathPerformanceCreateReq)

	return err
}

type NullableDfsPathPerformanceCreateReq struct {
	value *DfsPathPerformanceCreateReq
	isSet bool
}

func (v NullableDfsPathPerformanceCreateReq) Get() *DfsPathPerformanceCreateReq {
	return v.value
}

func (v *NullableDfsPathPerformanceCreateReq) Set(val *DfsPathPerformanceCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsPathPerformanceCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsPathPerformanceCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsPathPerformanceCreateReq(val *DfsPathPerformanceCreateReq) *NullableDfsPathPerformanceCreateReq {
	return &NullableDfsPathPerformanceCreateReq{value: val, isSet: true}
}

func (v NullableDfsPathPerformanceCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsPathPerformanceCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


