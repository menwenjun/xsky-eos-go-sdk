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

// checks if the DfsNFSShareSetACLsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsNFSShareSetACLsReq{}

// DfsNFSShareSetACLsReq struct for DfsNFSShareSetACLsReq
type DfsNFSShareSetACLsReq struct {
	DfsNfsShare DfsNFSShareSetACLsReqShare `json:"dfs_nfs_share"`
}

type _DfsNFSShareSetACLsReq DfsNFSShareSetACLsReq

// NewDfsNFSShareSetACLsReq instantiates a new DfsNFSShareSetACLsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsNFSShareSetACLsReq(dfsNfsShare DfsNFSShareSetACLsReqShare) *DfsNFSShareSetACLsReq {
	this := DfsNFSShareSetACLsReq{}
	this.DfsNfsShare = dfsNfsShare
	return &this
}

// NewDfsNFSShareSetACLsReqWithDefaults instantiates a new DfsNFSShareSetACLsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsNFSShareSetACLsReqWithDefaults() *DfsNFSShareSetACLsReq {
	this := DfsNFSShareSetACLsReq{}
	return &this
}

// GetDfsNfsShare returns the DfsNfsShare field value
func (o *DfsNFSShareSetACLsReq) GetDfsNfsShare() DfsNFSShareSetACLsReqShare {
	if o == nil {
		var ret DfsNFSShareSetACLsReqShare
		return ret
	}

	return o.DfsNfsShare
}

// GetDfsNfsShareOk returns a tuple with the DfsNfsShare field value
// and a boolean to check if the value has been set.
func (o *DfsNFSShareSetACLsReq) GetDfsNfsShareOk() (*DfsNFSShareSetACLsReqShare, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsNfsShare, true
}

// SetDfsNfsShare sets field value
func (o *DfsNFSShareSetACLsReq) SetDfsNfsShare(v DfsNFSShareSetACLsReqShare) {
	o.DfsNfsShare = v
}

func (o DfsNFSShareSetACLsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsNFSShareSetACLsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_nfs_share"] = o.DfsNfsShare
	return toSerialize, nil
}

func (o *DfsNFSShareSetACLsReq) UnmarshalJSON(data []byte) (err error) {
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

	varDfsNFSShareSetACLsReq := _DfsNFSShareSetACLsReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsNFSShareSetACLsReq)

	if err != nil {
		return err
	}

	*o = DfsNFSShareSetACLsReq(varDfsNFSShareSetACLsReq)

	return err
}

type NullableDfsNFSShareSetACLsReq struct {
	value *DfsNFSShareSetACLsReq
	isSet bool
}

func (v NullableDfsNFSShareSetACLsReq) Get() *DfsNFSShareSetACLsReq {
	return v.value
}

func (v *NullableDfsNFSShareSetACLsReq) Set(val *DfsNFSShareSetACLsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsNFSShareSetACLsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsNFSShareSetACLsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsNFSShareSetACLsReq(val *DfsNFSShareSetACLsReq) *NullableDfsNFSShareSetACLsReq {
	return &NullableDfsNFSShareSetACLsReq{value: val, isSet: true}
}

func (v NullableDfsNFSShareSetACLsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsNFSShareSetACLsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


