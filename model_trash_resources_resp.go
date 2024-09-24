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

// checks if the TrashResourcesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrashResourcesResp{}

// TrashResourcesResp struct for TrashResourcesResp
type TrashResourcesResp struct {
	// trash resources
	TrashResources []TrashResource `json:"trash_resources"`
}

type _TrashResourcesResp TrashResourcesResp

// NewTrashResourcesResp instantiates a new TrashResourcesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrashResourcesResp(trashResources []TrashResource) *TrashResourcesResp {
	this := TrashResourcesResp{}
	this.TrashResources = trashResources
	return &this
}

// NewTrashResourcesRespWithDefaults instantiates a new TrashResourcesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrashResourcesRespWithDefaults() *TrashResourcesResp {
	this := TrashResourcesResp{}
	return &this
}

// GetTrashResources returns the TrashResources field value
func (o *TrashResourcesResp) GetTrashResources() []TrashResource {
	if o == nil {
		var ret []TrashResource
		return ret
	}

	return o.TrashResources
}

// GetTrashResourcesOk returns a tuple with the TrashResources field value
// and a boolean to check if the value has been set.
func (o *TrashResourcesResp) GetTrashResourcesOk() ([]TrashResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrashResources, true
}

// SetTrashResources sets field value
func (o *TrashResourcesResp) SetTrashResources(v []TrashResource) {
	o.TrashResources = v
}

func (o TrashResourcesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrashResourcesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trash_resources"] = o.TrashResources
	return toSerialize, nil
}

func (o *TrashResourcesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"trash_resources",
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

	varTrashResourcesResp := _TrashResourcesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTrashResourcesResp)

	if err != nil {
		return err
	}

	*o = TrashResourcesResp(varTrashResourcesResp)

	return err
}

type NullableTrashResourcesResp struct {
	value *TrashResourcesResp
	isSet bool
}

func (v NullableTrashResourcesResp) Get() *TrashResourcesResp {
	return v.value
}

func (v *NullableTrashResourcesResp) Set(val *TrashResourcesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableTrashResourcesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableTrashResourcesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrashResourcesResp(val *TrashResourcesResp) *NullableTrashResourcesResp {
	return &NullableTrashResourcesResp{value: val, isSet: true}
}

func (v NullableTrashResourcesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrashResourcesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


