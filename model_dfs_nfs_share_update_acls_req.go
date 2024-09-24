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

// checks if the DfsNFSShareUpdateACLsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsNFSShareUpdateACLsReq{}

// DfsNFSShareUpdateACLsReq struct for DfsNFSShareUpdateACLsReq
type DfsNFSShareUpdateACLsReq struct {
	DfsNfsShare DfsNFSShareUpdateACLsReqShare `json:"dfs_nfs_share"`
}

type _DfsNFSShareUpdateACLsReq DfsNFSShareUpdateACLsReq

// NewDfsNFSShareUpdateACLsReq instantiates a new DfsNFSShareUpdateACLsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsNFSShareUpdateACLsReq(dfsNfsShare DfsNFSShareUpdateACLsReqShare) *DfsNFSShareUpdateACLsReq {
	this := DfsNFSShareUpdateACLsReq{}
	this.DfsNfsShare = dfsNfsShare
	return &this
}

// NewDfsNFSShareUpdateACLsReqWithDefaults instantiates a new DfsNFSShareUpdateACLsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsNFSShareUpdateACLsReqWithDefaults() *DfsNFSShareUpdateACLsReq {
	this := DfsNFSShareUpdateACLsReq{}
	return &this
}

// GetDfsNfsShare returns the DfsNfsShare field value
func (o *DfsNFSShareUpdateACLsReq) GetDfsNfsShare() DfsNFSShareUpdateACLsReqShare {
	if o == nil {
		var ret DfsNFSShareUpdateACLsReqShare
		return ret
	}

	return o.DfsNfsShare
}

// GetDfsNfsShareOk returns a tuple with the DfsNfsShare field value
// and a boolean to check if the value has been set.
func (o *DfsNFSShareUpdateACLsReq) GetDfsNfsShareOk() (*DfsNFSShareUpdateACLsReqShare, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsNfsShare, true
}

// SetDfsNfsShare sets field value
func (o *DfsNFSShareUpdateACLsReq) SetDfsNfsShare(v DfsNFSShareUpdateACLsReqShare) {
	o.DfsNfsShare = v
}

func (o DfsNFSShareUpdateACLsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsNFSShareUpdateACLsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_nfs_share"] = o.DfsNfsShare
	return toSerialize, nil
}

func (o *DfsNFSShareUpdateACLsReq) UnmarshalJSON(data []byte) (err error) {
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

	varDfsNFSShareUpdateACLsReq := _DfsNFSShareUpdateACLsReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsNFSShareUpdateACLsReq)

	if err != nil {
		return err
	}

	*o = DfsNFSShareUpdateACLsReq(varDfsNFSShareUpdateACLsReq)

	return err
}

type NullableDfsNFSShareUpdateACLsReq struct {
	value *DfsNFSShareUpdateACLsReq
	isSet bool
}

func (v NullableDfsNFSShareUpdateACLsReq) Get() *DfsNFSShareUpdateACLsReq {
	return v.value
}

func (v *NullableDfsNFSShareUpdateACLsReq) Set(val *DfsNFSShareUpdateACLsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsNFSShareUpdateACLsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsNFSShareUpdateACLsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsNFSShareUpdateACLsReq(val *DfsNFSShareUpdateACLsReq) *NullableDfsNFSShareUpdateACLsReq {
	return &NullableDfsNFSShareUpdateACLsReq{value: val, isSet: true}
}

func (v NullableDfsNFSShareUpdateACLsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsNFSShareUpdateACLsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


