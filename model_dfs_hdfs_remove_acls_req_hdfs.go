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

// checks if the DfsHdfsRemoveACLsReqHdfs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsHdfsRemoveACLsReqHdfs{}

// DfsHdfsRemoveACLsReqHdfs struct for DfsHdfsRemoveACLsReqHdfs
type DfsHdfsRemoveACLsReqHdfs struct {
	DfsHdfsAclIds []int64 `json:"dfs_hdfs_acl_ids"`
}

type _DfsHdfsRemoveACLsReqHdfs DfsHdfsRemoveACLsReqHdfs

// NewDfsHdfsRemoveACLsReqHdfs instantiates a new DfsHdfsRemoveACLsReqHdfs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsHdfsRemoveACLsReqHdfs(dfsHdfsAclIds []int64) *DfsHdfsRemoveACLsReqHdfs {
	this := DfsHdfsRemoveACLsReqHdfs{}
	this.DfsHdfsAclIds = dfsHdfsAclIds
	return &this
}

// NewDfsHdfsRemoveACLsReqHdfsWithDefaults instantiates a new DfsHdfsRemoveACLsReqHdfs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsHdfsRemoveACLsReqHdfsWithDefaults() *DfsHdfsRemoveACLsReqHdfs {
	this := DfsHdfsRemoveACLsReqHdfs{}
	return &this
}

// GetDfsHdfsAclIds returns the DfsHdfsAclIds field value
func (o *DfsHdfsRemoveACLsReqHdfs) GetDfsHdfsAclIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.DfsHdfsAclIds
}

// GetDfsHdfsAclIdsOk returns a tuple with the DfsHdfsAclIds field value
// and a boolean to check if the value has been set.
func (o *DfsHdfsRemoveACLsReqHdfs) GetDfsHdfsAclIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DfsHdfsAclIds, true
}

// SetDfsHdfsAclIds sets field value
func (o *DfsHdfsRemoveACLsReqHdfs) SetDfsHdfsAclIds(v []int64) {
	o.DfsHdfsAclIds = v
}

func (o DfsHdfsRemoveACLsReqHdfs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsHdfsRemoveACLsReqHdfs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_hdfs_acl_ids"] = o.DfsHdfsAclIds
	return toSerialize, nil
}

func (o *DfsHdfsRemoveACLsReqHdfs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_hdfs_acl_ids",
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

	varDfsHdfsRemoveACLsReqHdfs := _DfsHdfsRemoveACLsReqHdfs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsHdfsRemoveACLsReqHdfs)

	if err != nil {
		return err
	}

	*o = DfsHdfsRemoveACLsReqHdfs(varDfsHdfsRemoveACLsReqHdfs)

	return err
}

type NullableDfsHdfsRemoveACLsReqHdfs struct {
	value *DfsHdfsRemoveACLsReqHdfs
	isSet bool
}

func (v NullableDfsHdfsRemoveACLsReqHdfs) Get() *DfsHdfsRemoveACLsReqHdfs {
	return v.value
}

func (v *NullableDfsHdfsRemoveACLsReqHdfs) Set(val *DfsHdfsRemoveACLsReqHdfs) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsHdfsRemoveACLsReqHdfs) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsHdfsRemoveACLsReqHdfs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsHdfsRemoveACLsReqHdfs(val *DfsHdfsRemoveACLsReqHdfs) *NullableDfsHdfsRemoveACLsReqHdfs {
	return &NullableDfsHdfsRemoveACLsReqHdfs{value: val, isSet: true}
}

func (v NullableDfsHdfsRemoveACLsReqHdfs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsHdfsRemoveACLsReqHdfs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


