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

// checks if the DfsSMBWindowsACLSetReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsSMBWindowsACLSetReq{}

// DfsSMBWindowsACLSetReq struct for DfsSMBWindowsACLSetReq
type DfsSMBWindowsACLSetReq struct {
	// id of rootfs
	DfsRootfsId int64 `json:"dfs_rootfs_id"`
	// dfs smb windows acls
	DfsSmbWindowsAcls []DfsSMBWindowsACLSet `json:"dfs_smb_windows_acls"`
	// inheritance type
	InheritanceType *string `json:"inheritance_type,omitempty"`
	// limit handler for sub directory
	LimitHandler *bool `json:"limit_handler,omitempty"`
	// directory or subdirectory or sub file of share path
	Path string `json:"path"`
	// whether replace permission
	ReplacePermission *bool `json:"replace_permission,omitempty"`
}

type _DfsSMBWindowsACLSetReq DfsSMBWindowsACLSetReq

// NewDfsSMBWindowsACLSetReq instantiates a new DfsSMBWindowsACLSetReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsSMBWindowsACLSetReq(dfsRootfsId int64, dfsSmbWindowsAcls []DfsSMBWindowsACLSet, path string) *DfsSMBWindowsACLSetReq {
	this := DfsSMBWindowsACLSetReq{}
	this.DfsRootfsId = dfsRootfsId
	this.DfsSmbWindowsAcls = dfsSmbWindowsAcls
	this.Path = path
	return &this
}

// NewDfsSMBWindowsACLSetReqWithDefaults instantiates a new DfsSMBWindowsACLSetReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsSMBWindowsACLSetReqWithDefaults() *DfsSMBWindowsACLSetReq {
	this := DfsSMBWindowsACLSetReq{}
	return &this
}

// GetDfsRootfsId returns the DfsRootfsId field value
func (o *DfsSMBWindowsACLSetReq) GetDfsRootfsId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DfsRootfsId
}

// GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field value
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLSetReq) GetDfsRootfsIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsRootfsId, true
}

// SetDfsRootfsId sets field value
func (o *DfsSMBWindowsACLSetReq) SetDfsRootfsId(v int64) {
	o.DfsRootfsId = v
}

// GetDfsSmbWindowsAcls returns the DfsSmbWindowsAcls field value
func (o *DfsSMBWindowsACLSetReq) GetDfsSmbWindowsAcls() []DfsSMBWindowsACLSet {
	if o == nil {
		var ret []DfsSMBWindowsACLSet
		return ret
	}

	return o.DfsSmbWindowsAcls
}

// GetDfsSmbWindowsAclsOk returns a tuple with the DfsSmbWindowsAcls field value
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLSetReq) GetDfsSmbWindowsAclsOk() ([]DfsSMBWindowsACLSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.DfsSmbWindowsAcls, true
}

// SetDfsSmbWindowsAcls sets field value
func (o *DfsSMBWindowsACLSetReq) SetDfsSmbWindowsAcls(v []DfsSMBWindowsACLSet) {
	o.DfsSmbWindowsAcls = v
}

// GetInheritanceType returns the InheritanceType field value if set, zero value otherwise.
func (o *DfsSMBWindowsACLSetReq) GetInheritanceType() string {
	if o == nil || IsNil(o.InheritanceType) {
		var ret string
		return ret
	}
	return *o.InheritanceType
}

// GetInheritanceTypeOk returns a tuple with the InheritanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLSetReq) GetInheritanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InheritanceType) {
		return nil, false
	}
	return o.InheritanceType, true
}

// HasInheritanceType returns a boolean if a field has been set.
func (o *DfsSMBWindowsACLSetReq) HasInheritanceType() bool {
	if o != nil && !IsNil(o.InheritanceType) {
		return true
	}

	return false
}

// SetInheritanceType gets a reference to the given string and assigns it to the InheritanceType field.
func (o *DfsSMBWindowsACLSetReq) SetInheritanceType(v string) {
	o.InheritanceType = &v
}

// GetLimitHandler returns the LimitHandler field value if set, zero value otherwise.
func (o *DfsSMBWindowsACLSetReq) GetLimitHandler() bool {
	if o == nil || IsNil(o.LimitHandler) {
		var ret bool
		return ret
	}
	return *o.LimitHandler
}

// GetLimitHandlerOk returns a tuple with the LimitHandler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLSetReq) GetLimitHandlerOk() (*bool, bool) {
	if o == nil || IsNil(o.LimitHandler) {
		return nil, false
	}
	return o.LimitHandler, true
}

// HasLimitHandler returns a boolean if a field has been set.
func (o *DfsSMBWindowsACLSetReq) HasLimitHandler() bool {
	if o != nil && !IsNil(o.LimitHandler) {
		return true
	}

	return false
}

// SetLimitHandler gets a reference to the given bool and assigns it to the LimitHandler field.
func (o *DfsSMBWindowsACLSetReq) SetLimitHandler(v bool) {
	o.LimitHandler = &v
}

// GetPath returns the Path field value
func (o *DfsSMBWindowsACLSetReq) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLSetReq) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DfsSMBWindowsACLSetReq) SetPath(v string) {
	o.Path = v
}

// GetReplacePermission returns the ReplacePermission field value if set, zero value otherwise.
func (o *DfsSMBWindowsACLSetReq) GetReplacePermission() bool {
	if o == nil || IsNil(o.ReplacePermission) {
		var ret bool
		return ret
	}
	return *o.ReplacePermission
}

// GetReplacePermissionOk returns a tuple with the ReplacePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBWindowsACLSetReq) GetReplacePermissionOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplacePermission) {
		return nil, false
	}
	return o.ReplacePermission, true
}

// HasReplacePermission returns a boolean if a field has been set.
func (o *DfsSMBWindowsACLSetReq) HasReplacePermission() bool {
	if o != nil && !IsNil(o.ReplacePermission) {
		return true
	}

	return false
}

// SetReplacePermission gets a reference to the given bool and assigns it to the ReplacePermission field.
func (o *DfsSMBWindowsACLSetReq) SetReplacePermission(v bool) {
	o.ReplacePermission = &v
}

func (o DfsSMBWindowsACLSetReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsSMBWindowsACLSetReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_rootfs_id"] = o.DfsRootfsId
	toSerialize["dfs_smb_windows_acls"] = o.DfsSmbWindowsAcls
	if !IsNil(o.InheritanceType) {
		toSerialize["inheritance_type"] = o.InheritanceType
	}
	if !IsNil(o.LimitHandler) {
		toSerialize["limit_handler"] = o.LimitHandler
	}
	toSerialize["path"] = o.Path
	if !IsNil(o.ReplacePermission) {
		toSerialize["replace_permission"] = o.ReplacePermission
	}
	return toSerialize, nil
}

func (o *DfsSMBWindowsACLSetReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_rootfs_id",
		"dfs_smb_windows_acls",
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

	varDfsSMBWindowsACLSetReq := _DfsSMBWindowsACLSetReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsSMBWindowsACLSetReq)

	if err != nil {
		return err
	}

	*o = DfsSMBWindowsACLSetReq(varDfsSMBWindowsACLSetReq)

	return err
}

type NullableDfsSMBWindowsACLSetReq struct {
	value *DfsSMBWindowsACLSetReq
	isSet bool
}

func (v NullableDfsSMBWindowsACLSetReq) Get() *DfsSMBWindowsACLSetReq {
	return v.value
}

func (v *NullableDfsSMBWindowsACLSetReq) Set(val *DfsSMBWindowsACLSetReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsSMBWindowsACLSetReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsSMBWindowsACLSetReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsSMBWindowsACLSetReq(val *DfsSMBWindowsACLSetReq) *NullableDfsSMBWindowsACLSetReq {
	return &NullableDfsSMBWindowsACLSetReq{value: val, isSet: true}
}

func (v NullableDfsSMBWindowsACLSetReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsSMBWindowsACLSetReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


