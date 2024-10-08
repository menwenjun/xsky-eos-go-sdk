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

// checks if the DfsCSIDirectoryResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsCSIDirectoryResp{}

// DfsCSIDirectoryResp struct for DfsCSIDirectoryResp
type DfsCSIDirectoryResp struct {
	DfsDirectory DfsCSIDirectoryRespDirectory `json:"dfs_directory"`
	DfsQuota *DfsQuota `json:"dfs_quota,omitempty"`
}

type _DfsCSIDirectoryResp DfsCSIDirectoryResp

// NewDfsCSIDirectoryResp instantiates a new DfsCSIDirectoryResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsCSIDirectoryResp(dfsDirectory DfsCSIDirectoryRespDirectory) *DfsCSIDirectoryResp {
	this := DfsCSIDirectoryResp{}
	this.DfsDirectory = dfsDirectory
	return &this
}

// NewDfsCSIDirectoryRespWithDefaults instantiates a new DfsCSIDirectoryResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsCSIDirectoryRespWithDefaults() *DfsCSIDirectoryResp {
	this := DfsCSIDirectoryResp{}
	return &this
}

// GetDfsDirectory returns the DfsDirectory field value
func (o *DfsCSIDirectoryResp) GetDfsDirectory() DfsCSIDirectoryRespDirectory {
	if o == nil {
		var ret DfsCSIDirectoryRespDirectory
		return ret
	}

	return o.DfsDirectory
}

// GetDfsDirectoryOk returns a tuple with the DfsDirectory field value
// and a boolean to check if the value has been set.
func (o *DfsCSIDirectoryResp) GetDfsDirectoryOk() (*DfsCSIDirectoryRespDirectory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsDirectory, true
}

// SetDfsDirectory sets field value
func (o *DfsCSIDirectoryResp) SetDfsDirectory(v DfsCSIDirectoryRespDirectory) {
	o.DfsDirectory = v
}

// GetDfsQuota returns the DfsQuota field value if set, zero value otherwise.
func (o *DfsCSIDirectoryResp) GetDfsQuota() DfsQuota {
	if o == nil || IsNil(o.DfsQuota) {
		var ret DfsQuota
		return ret
	}
	return *o.DfsQuota
}

// GetDfsQuotaOk returns a tuple with the DfsQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsCSIDirectoryResp) GetDfsQuotaOk() (*DfsQuota, bool) {
	if o == nil || IsNil(o.DfsQuota) {
		return nil, false
	}
	return o.DfsQuota, true
}

// HasDfsQuota returns a boolean if a field has been set.
func (o *DfsCSIDirectoryResp) HasDfsQuota() bool {
	if o != nil && !IsNil(o.DfsQuota) {
		return true
	}

	return false
}

// SetDfsQuota gets a reference to the given DfsQuota and assigns it to the DfsQuota field.
func (o *DfsCSIDirectoryResp) SetDfsQuota(v DfsQuota) {
	o.DfsQuota = &v
}

func (o DfsCSIDirectoryResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsCSIDirectoryResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_directory"] = o.DfsDirectory
	if !IsNil(o.DfsQuota) {
		toSerialize["dfs_quota"] = o.DfsQuota
	}
	return toSerialize, nil
}

func (o *DfsCSIDirectoryResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_directory",
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

	varDfsCSIDirectoryResp := _DfsCSIDirectoryResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsCSIDirectoryResp)

	if err != nil {
		return err
	}

	*o = DfsCSIDirectoryResp(varDfsCSIDirectoryResp)

	return err
}

type NullableDfsCSIDirectoryResp struct {
	value *DfsCSIDirectoryResp
	isSet bool
}

func (v NullableDfsCSIDirectoryResp) Get() *DfsCSIDirectoryResp {
	return v.value
}

func (v *NullableDfsCSIDirectoryResp) Set(val *DfsCSIDirectoryResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsCSIDirectoryResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsCSIDirectoryResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsCSIDirectoryResp(val *DfsCSIDirectoryResp) *NullableDfsCSIDirectoryResp {
	return &NullableDfsCSIDirectoryResp{value: val, isSet: true}
}

func (v NullableDfsCSIDirectoryResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsCSIDirectoryResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


