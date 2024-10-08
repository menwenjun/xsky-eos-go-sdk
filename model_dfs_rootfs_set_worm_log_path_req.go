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

// checks if the DfsRootfsSetWormLogPathReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsRootfsSetWormLogPathReq{}

// DfsRootfsSetWormLogPathReq struct for DfsRootfsSetWormLogPathReq
type DfsRootfsSetWormLogPathReq struct {
	DfsRootfs DfsRootfsSetWormLogPathReqRootfs `json:"dfs_rootfs"`
}

type _DfsRootfsSetWormLogPathReq DfsRootfsSetWormLogPathReq

// NewDfsRootfsSetWormLogPathReq instantiates a new DfsRootfsSetWormLogPathReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsRootfsSetWormLogPathReq(dfsRootfs DfsRootfsSetWormLogPathReqRootfs) *DfsRootfsSetWormLogPathReq {
	this := DfsRootfsSetWormLogPathReq{}
	this.DfsRootfs = dfsRootfs
	return &this
}

// NewDfsRootfsSetWormLogPathReqWithDefaults instantiates a new DfsRootfsSetWormLogPathReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsRootfsSetWormLogPathReqWithDefaults() *DfsRootfsSetWormLogPathReq {
	this := DfsRootfsSetWormLogPathReq{}
	return &this
}

// GetDfsRootfs returns the DfsRootfs field value
func (o *DfsRootfsSetWormLogPathReq) GetDfsRootfs() DfsRootfsSetWormLogPathReqRootfs {
	if o == nil {
		var ret DfsRootfsSetWormLogPathReqRootfs
		return ret
	}

	return o.DfsRootfs
}

// GetDfsRootfsOk returns a tuple with the DfsRootfs field value
// and a boolean to check if the value has been set.
func (o *DfsRootfsSetWormLogPathReq) GetDfsRootfsOk() (*DfsRootfsSetWormLogPathReqRootfs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsRootfs, true
}

// SetDfsRootfs sets field value
func (o *DfsRootfsSetWormLogPathReq) SetDfsRootfs(v DfsRootfsSetWormLogPathReqRootfs) {
	o.DfsRootfs = v
}

func (o DfsRootfsSetWormLogPathReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsRootfsSetWormLogPathReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_rootfs"] = o.DfsRootfs
	return toSerialize, nil
}

func (o *DfsRootfsSetWormLogPathReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_rootfs",
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

	varDfsRootfsSetWormLogPathReq := _DfsRootfsSetWormLogPathReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsRootfsSetWormLogPathReq)

	if err != nil {
		return err
	}

	*o = DfsRootfsSetWormLogPathReq(varDfsRootfsSetWormLogPathReq)

	return err
}

type NullableDfsRootfsSetWormLogPathReq struct {
	value *DfsRootfsSetWormLogPathReq
	isSet bool
}

func (v NullableDfsRootfsSetWormLogPathReq) Get() *DfsRootfsSetWormLogPathReq {
	return v.value
}

func (v *NullableDfsRootfsSetWormLogPathReq) Set(val *DfsRootfsSetWormLogPathReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsRootfsSetWormLogPathReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsRootfsSetWormLogPathReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsRootfsSetWormLogPathReq(val *DfsRootfsSetWormLogPathReq) *NullableDfsRootfsSetWormLogPathReq {
	return &NullableDfsRootfsSetWormLogPathReq{value: val, isSet: true}
}

func (v NullableDfsRootfsSetWormLogPathReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsRootfsSetWormLogPathReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


