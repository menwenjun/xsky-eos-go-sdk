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

// checks if the DfsSMBWindowsACLInheritanceUpdateReqWindowsACL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsSMBWindowsACLInheritanceUpdateReqWindowsACL{}

// DfsSMBWindowsACLInheritanceUpdateReqWindowsACL struct for DfsSMBWindowsACLInheritanceUpdateReqWindowsACL
type DfsSMBWindowsACLInheritanceUpdateReqWindowsACL struct {
	// type of update action
	ActionType string `json:"action_type"`
	// id of rootfs
	DfsRootfsId int64 `json:"dfs_rootfs_id"`
	// directory path
	Path string `json:"path"`
}

type _DfsSMBWindowsACLInheritanceUpdateReqWindowsACL DfsSMBWindowsACLInheritanceUpdateReqWindowsACL

// NewDfsSMBWindowsACLInheritanceUpdateReqWindowsACL instantiates a new DfsSMBWindowsACLInheritanceUpdateReqWindowsACL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsSMBWindowsACLInheritanceUpdateReqWindowsACL(actionType string, dfsRootfsId int64, path string) *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL {
	this := DfsSMBWindowsACLInheritanceUpdateReqWindowsACL{}
	this.ActionType = actionType
	this.DfsRootfsId = dfsRootfsId
	this.Path = path
	return &this
}

// NewDfsSMBWindowsACLInheritanceUpdateReqWindowsACLWithDefaults instantiates a new DfsSMBWindowsACLInheritanceUpdateReqWindowsACL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsSMBWindowsACLInheritanceUpdateReqWindowsACLWithDefaults() *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL {
	this := DfsSMBWindowsACLInheritanceUpdateReqWindowsACL{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) GetActionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) SetActionType(v string) {
	o.ActionType = v
}

// GetDfsRootfsId returns the DfsRootfsId field value
func (o *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) GetDfsRootfsId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DfsRootfsId
}

// GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field value
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) GetDfsRootfsIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsRootfsId, true
}

// SetDfsRootfsId sets field value
func (o *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) SetDfsRootfsId(v int64) {
	o.DfsRootfsId = v
}

// GetPath returns the Path field value
func (o *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) SetPath(v string) {
	o.Path = v
}

func (o DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action_type"] = o.ActionType
	toSerialize["dfs_rootfs_id"] = o.DfsRootfsId
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

func (o *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action_type",
		"dfs_rootfs_id",
		"path",
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

	varDfsSMBWindowsACLInheritanceUpdateReqWindowsACL := _DfsSMBWindowsACLInheritanceUpdateReqWindowsACL{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsSMBWindowsACLInheritanceUpdateReqWindowsACL)

	if err != nil {
		return err
	}

	*o = DfsSMBWindowsACLInheritanceUpdateReqWindowsACL(varDfsSMBWindowsACLInheritanceUpdateReqWindowsACL)

	return err
}

type NullableDfsSMBWindowsACLInheritanceUpdateReqWindowsACL struct {
	value *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL
	isSet bool
}

func (v NullableDfsSMBWindowsACLInheritanceUpdateReqWindowsACL) Get() *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL {
	return v.value
}

func (v *NullableDfsSMBWindowsACLInheritanceUpdateReqWindowsACL) Set(val *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsSMBWindowsACLInheritanceUpdateReqWindowsACL) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsSMBWindowsACLInheritanceUpdateReqWindowsACL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsSMBWindowsACLInheritanceUpdateReqWindowsACL(val *DfsSMBWindowsACLInheritanceUpdateReqWindowsACL) *NullableDfsSMBWindowsACLInheritanceUpdateReqWindowsACL {
	return &NullableDfsSMBWindowsACLInheritanceUpdateReqWindowsACL{value: val, isSet: true}
}

func (v NullableDfsSMBWindowsACLInheritanceUpdateReqWindowsACL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsSMBWindowsACLInheritanceUpdateReqWindowsACL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


