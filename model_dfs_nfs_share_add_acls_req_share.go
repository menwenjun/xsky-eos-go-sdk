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

// checks if the DfsNFSShareAddACLsReqShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsNFSShareAddACLsReqShare{}

// DfsNFSShareAddACLsReqShare struct for DfsNFSShareAddACLsReqShare
type DfsNFSShareAddACLsReqShare struct {
	// access control array
	DfsNfsShareAcls []DfsNFSShareACLReq `json:"dfs_nfs_share_acls"`
}

type _DfsNFSShareAddACLsReqShare DfsNFSShareAddACLsReqShare

// NewDfsNFSShareAddACLsReqShare instantiates a new DfsNFSShareAddACLsReqShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsNFSShareAddACLsReqShare(dfsNfsShareAcls []DfsNFSShareACLReq) *DfsNFSShareAddACLsReqShare {
	this := DfsNFSShareAddACLsReqShare{}
	this.DfsNfsShareAcls = dfsNfsShareAcls
	return &this
}

// NewDfsNFSShareAddACLsReqShareWithDefaults instantiates a new DfsNFSShareAddACLsReqShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsNFSShareAddACLsReqShareWithDefaults() *DfsNFSShareAddACLsReqShare {
	this := DfsNFSShareAddACLsReqShare{}
	return &this
}

// GetDfsNfsShareAcls returns the DfsNfsShareAcls field value
func (o *DfsNFSShareAddACLsReqShare) GetDfsNfsShareAcls() []DfsNFSShareACLReq {
	if o == nil {
		var ret []DfsNFSShareACLReq
		return ret
	}

	return o.DfsNfsShareAcls
}

// GetDfsNfsShareAclsOk returns a tuple with the DfsNfsShareAcls field value
// and a boolean to check if the value has been set.
func (o *DfsNFSShareAddACLsReqShare) GetDfsNfsShareAclsOk() ([]DfsNFSShareACLReq, bool) {
	if o == nil {
		return nil, false
	}
	return o.DfsNfsShareAcls, true
}

// SetDfsNfsShareAcls sets field value
func (o *DfsNFSShareAddACLsReqShare) SetDfsNfsShareAcls(v []DfsNFSShareACLReq) {
	o.DfsNfsShareAcls = v
}

func (o DfsNFSShareAddACLsReqShare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsNFSShareAddACLsReqShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_nfs_share_acls"] = o.DfsNfsShareAcls
	return toSerialize, nil
}

func (o *DfsNFSShareAddACLsReqShare) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_nfs_share_acls",
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

	varDfsNFSShareAddACLsReqShare := _DfsNFSShareAddACLsReqShare{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsNFSShareAddACLsReqShare)

	if err != nil {
		return err
	}

	*o = DfsNFSShareAddACLsReqShare(varDfsNFSShareAddACLsReqShare)

	return err
}

type NullableDfsNFSShareAddACLsReqShare struct {
	value *DfsNFSShareAddACLsReqShare
	isSet bool
}

func (v NullableDfsNFSShareAddACLsReqShare) Get() *DfsNFSShareAddACLsReqShare {
	return v.value
}

func (v *NullableDfsNFSShareAddACLsReqShare) Set(val *DfsNFSShareAddACLsReqShare) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsNFSShareAddACLsReqShare) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsNFSShareAddACLsReqShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsNFSShareAddACLsReqShare(val *DfsNFSShareAddACLsReqShare) *NullableDfsNFSShareAddACLsReqShare {
	return &NullableDfsNFSShareAddACLsReqShare{value: val, isSet: true}
}

func (v NullableDfsNFSShareAddACLsReqShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsNFSShareAddACLsReqShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


