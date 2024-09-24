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

// checks if the PlacementNodeUpdateReqPlacementNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlacementNodeUpdateReqPlacementNode{}

// PlacementNodeUpdateReqPlacementNode struct for PlacementNodeUpdateReqPlacementNode
type PlacementNodeUpdateReqPlacementNode struct {
	Name string `json:"name"`
	ParentId int64 `json:"parent_id"`
}

type _PlacementNodeUpdateReqPlacementNode PlacementNodeUpdateReqPlacementNode

// NewPlacementNodeUpdateReqPlacementNode instantiates a new PlacementNodeUpdateReqPlacementNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlacementNodeUpdateReqPlacementNode(name string, parentId int64) *PlacementNodeUpdateReqPlacementNode {
	this := PlacementNodeUpdateReqPlacementNode{}
	this.Name = name
	this.ParentId = parentId
	return &this
}

// NewPlacementNodeUpdateReqPlacementNodeWithDefaults instantiates a new PlacementNodeUpdateReqPlacementNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlacementNodeUpdateReqPlacementNodeWithDefaults() *PlacementNodeUpdateReqPlacementNode {
	this := PlacementNodeUpdateReqPlacementNode{}
	return &this
}

// GetName returns the Name field value
func (o *PlacementNodeUpdateReqPlacementNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlacementNodeUpdateReqPlacementNode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlacementNodeUpdateReqPlacementNode) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value
func (o *PlacementNodeUpdateReqPlacementNode) GetParentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *PlacementNodeUpdateReqPlacementNode) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *PlacementNodeUpdateReqPlacementNode) SetParentId(v int64) {
	o.ParentId = v
}

func (o PlacementNodeUpdateReqPlacementNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlacementNodeUpdateReqPlacementNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["parent_id"] = o.ParentId
	return toSerialize, nil
}

func (o *PlacementNodeUpdateReqPlacementNode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"parent_id",
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

	varPlacementNodeUpdateReqPlacementNode := _PlacementNodeUpdateReqPlacementNode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlacementNodeUpdateReqPlacementNode)

	if err != nil {
		return err
	}

	*o = PlacementNodeUpdateReqPlacementNode(varPlacementNodeUpdateReqPlacementNode)

	return err
}

type NullablePlacementNodeUpdateReqPlacementNode struct {
	value *PlacementNodeUpdateReqPlacementNode
	isSet bool
}

func (v NullablePlacementNodeUpdateReqPlacementNode) Get() *PlacementNodeUpdateReqPlacementNode {
	return v.value
}

func (v *NullablePlacementNodeUpdateReqPlacementNode) Set(val *PlacementNodeUpdateReqPlacementNode) {
	v.value = val
	v.isSet = true
}

func (v NullablePlacementNodeUpdateReqPlacementNode) IsSet() bool {
	return v.isSet
}

func (v *NullablePlacementNodeUpdateReqPlacementNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlacementNodeUpdateReqPlacementNode(val *PlacementNodeUpdateReqPlacementNode) *NullablePlacementNodeUpdateReqPlacementNode {
	return &NullablePlacementNodeUpdateReqPlacementNode{value: val, isSet: true}
}

func (v NullablePlacementNodeUpdateReqPlacementNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlacementNodeUpdateReqPlacementNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


