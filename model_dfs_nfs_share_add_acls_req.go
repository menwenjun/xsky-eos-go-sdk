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

// checks if the DfsNFSShareAddACLsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsNFSShareAddACLsReq{}

// DfsNFSShareAddACLsReq struct for DfsNFSShareAddACLsReq
type DfsNFSShareAddACLsReq struct {
	DfsNfsShare DfsNFSShareAddACLsReqShare `json:"dfs_nfs_share"`
}

type _DfsNFSShareAddACLsReq DfsNFSShareAddACLsReq

// NewDfsNFSShareAddACLsReq instantiates a new DfsNFSShareAddACLsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsNFSShareAddACLsReq(dfsNfsShare DfsNFSShareAddACLsReqShare) *DfsNFSShareAddACLsReq {
	this := DfsNFSShareAddACLsReq{}
	this.DfsNfsShare = dfsNfsShare
	return &this
}

// NewDfsNFSShareAddACLsReqWithDefaults instantiates a new DfsNFSShareAddACLsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsNFSShareAddACLsReqWithDefaults() *DfsNFSShareAddACLsReq {
	this := DfsNFSShareAddACLsReq{}
	return &this
}

// GetDfsNfsShare returns the DfsNfsShare field value
func (o *DfsNFSShareAddACLsReq) GetDfsNfsShare() DfsNFSShareAddACLsReqShare {
	if o == nil {
		var ret DfsNFSShareAddACLsReqShare
		return ret
	}

	return o.DfsNfsShare
}

// GetDfsNfsShareOk returns a tuple with the DfsNfsShare field value
// and a boolean to check if the value has been set.
func (o *DfsNFSShareAddACLsReq) GetDfsNfsShareOk() (*DfsNFSShareAddACLsReqShare, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsNfsShare, true
}

// SetDfsNfsShare sets field value
func (o *DfsNFSShareAddACLsReq) SetDfsNfsShare(v DfsNFSShareAddACLsReqShare) {
	o.DfsNfsShare = v
}

func (o DfsNFSShareAddACLsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsNFSShareAddACLsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_nfs_share"] = o.DfsNfsShare
	return toSerialize, nil
}

func (o *DfsNFSShareAddACLsReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_nfs_share",
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

	varDfsNFSShareAddACLsReq := _DfsNFSShareAddACLsReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsNFSShareAddACLsReq)

	if err != nil {
		return err
	}

	*o = DfsNFSShareAddACLsReq(varDfsNFSShareAddACLsReq)

	return err
}

type NullableDfsNFSShareAddACLsReq struct {
	value *DfsNFSShareAddACLsReq
	isSet bool
}

func (v NullableDfsNFSShareAddACLsReq) Get() *DfsNFSShareAddACLsReq {
	return v.value
}

func (v *NullableDfsNFSShareAddACLsReq) Set(val *DfsNFSShareAddACLsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsNFSShareAddACLsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsNFSShareAddACLsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsNFSShareAddACLsReq(val *DfsNFSShareAddACLsReq) *NullableDfsNFSShareAddACLsReq {
	return &NullableDfsNFSShareAddACLsReq{value: val, isSet: true}
}

func (v NullableDfsNFSShareAddACLsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsNFSShareAddACLsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


