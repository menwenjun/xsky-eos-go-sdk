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

// checks if the DfsXDCacheEnableReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsXDCacheEnableReq{}

// DfsXDCacheEnableReq struct for DfsXDCacheEnableReq
type DfsXDCacheEnableReq struct {
	DfsXdcache XDCache `json:"dfs_xdcache"`
}

type _DfsXDCacheEnableReq DfsXDCacheEnableReq

// NewDfsXDCacheEnableReq instantiates a new DfsXDCacheEnableReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsXDCacheEnableReq(dfsXdcache XDCache) *DfsXDCacheEnableReq {
	this := DfsXDCacheEnableReq{}
	this.DfsXdcache = dfsXdcache
	return &this
}

// NewDfsXDCacheEnableReqWithDefaults instantiates a new DfsXDCacheEnableReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsXDCacheEnableReqWithDefaults() *DfsXDCacheEnableReq {
	this := DfsXDCacheEnableReq{}
	return &this
}

// GetDfsXdcache returns the DfsXdcache field value
func (o *DfsXDCacheEnableReq) GetDfsXdcache() XDCache {
	if o == nil {
		var ret XDCache
		return ret
	}

	return o.DfsXdcache
}

// GetDfsXdcacheOk returns a tuple with the DfsXdcache field value
// and a boolean to check if the value has been set.
func (o *DfsXDCacheEnableReq) GetDfsXdcacheOk() (*XDCache, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsXdcache, true
}

// SetDfsXdcache sets field value
func (o *DfsXDCacheEnableReq) SetDfsXdcache(v XDCache) {
	o.DfsXdcache = v
}

func (o DfsXDCacheEnableReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsXDCacheEnableReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_xdcache"] = o.DfsXdcache
	return toSerialize, nil
}

func (o *DfsXDCacheEnableReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_xdcache",
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

	varDfsXDCacheEnableReq := _DfsXDCacheEnableReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsXDCacheEnableReq)

	if err != nil {
		return err
	}

	*o = DfsXDCacheEnableReq(varDfsXDCacheEnableReq)

	return err
}

type NullableDfsXDCacheEnableReq struct {
	value *DfsXDCacheEnableReq
	isSet bool
}

func (v NullableDfsXDCacheEnableReq) Get() *DfsXDCacheEnableReq {
	return v.value
}

func (v *NullableDfsXDCacheEnableReq) Set(val *DfsXDCacheEnableReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsXDCacheEnableReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsXDCacheEnableReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsXDCacheEnableReq(val *DfsXDCacheEnableReq) *NullableDfsXDCacheEnableReq {
	return &NullableDfsXDCacheEnableReq{value: val, isSet: true}
}

func (v NullableDfsXDCacheEnableReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsXDCacheEnableReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


