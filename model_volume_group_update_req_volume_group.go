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

// checks if the VolumeGroupUpdateReqVolumeGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeGroupUpdateReqVolumeGroup{}

// VolumeGroupUpdateReqVolumeGroup struct for VolumeGroupUpdateReqVolumeGroup
type VolumeGroupUpdateReqVolumeGroup struct {
	Description *string `json:"description,omitempty"`
	Name string `json:"name"`
}

type _VolumeGroupUpdateReqVolumeGroup VolumeGroupUpdateReqVolumeGroup

// NewVolumeGroupUpdateReqVolumeGroup instantiates a new VolumeGroupUpdateReqVolumeGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeGroupUpdateReqVolumeGroup(name string) *VolumeGroupUpdateReqVolumeGroup {
	this := VolumeGroupUpdateReqVolumeGroup{}
	this.Name = name
	return &this
}

// NewVolumeGroupUpdateReqVolumeGroupWithDefaults instantiates a new VolumeGroupUpdateReqVolumeGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeGroupUpdateReqVolumeGroupWithDefaults() *VolumeGroupUpdateReqVolumeGroup {
	this := VolumeGroupUpdateReqVolumeGroup{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VolumeGroupUpdateReqVolumeGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeGroupUpdateReqVolumeGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VolumeGroupUpdateReqVolumeGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VolumeGroupUpdateReqVolumeGroup) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *VolumeGroupUpdateReqVolumeGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VolumeGroupUpdateReqVolumeGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VolumeGroupUpdateReqVolumeGroup) SetName(v string) {
	o.Name = v
}

func (o VolumeGroupUpdateReqVolumeGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeGroupUpdateReqVolumeGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *VolumeGroupUpdateReqVolumeGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varVolumeGroupUpdateReqVolumeGroup := _VolumeGroupUpdateReqVolumeGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVolumeGroupUpdateReqVolumeGroup)

	if err != nil {
		return err
	}

	*o = VolumeGroupUpdateReqVolumeGroup(varVolumeGroupUpdateReqVolumeGroup)

	return err
}

type NullableVolumeGroupUpdateReqVolumeGroup struct {
	value *VolumeGroupUpdateReqVolumeGroup
	isSet bool
}

func (v NullableVolumeGroupUpdateReqVolumeGroup) Get() *VolumeGroupUpdateReqVolumeGroup {
	return v.value
}

func (v *NullableVolumeGroupUpdateReqVolumeGroup) Set(val *VolumeGroupUpdateReqVolumeGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeGroupUpdateReqVolumeGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeGroupUpdateReqVolumeGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeGroupUpdateReqVolumeGroup(val *VolumeGroupUpdateReqVolumeGroup) *NullableVolumeGroupUpdateReqVolumeGroup {
	return &NullableVolumeGroupUpdateReqVolumeGroup{value: val, isSet: true}
}

func (v NullableVolumeGroupUpdateReqVolumeGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeGroupUpdateReqVolumeGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


