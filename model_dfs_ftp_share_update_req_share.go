/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DfsFTPShareUpdateReqShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsFTPShareUpdateReqShare{}

// DfsFTPShareUpdateReqShare struct for DfsFTPShareUpdateReqShare
type DfsFTPShareUpdateReqShare struct {
	// description of share
	Description *string `json:"description,omitempty"`
	// new dfs gateway group id
	DfsGatewayGroupId *int64 `json:"dfs_gateway_group_id,omitempty"`
}

// NewDfsFTPShareUpdateReqShare instantiates a new DfsFTPShareUpdateReqShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsFTPShareUpdateReqShare() *DfsFTPShareUpdateReqShare {
	this := DfsFTPShareUpdateReqShare{}
	return &this
}

// NewDfsFTPShareUpdateReqShareWithDefaults instantiates a new DfsFTPShareUpdateReqShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsFTPShareUpdateReqShareWithDefaults() *DfsFTPShareUpdateReqShare {
	this := DfsFTPShareUpdateReqShare{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DfsFTPShareUpdateReqShare) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareUpdateReqShare) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DfsFTPShareUpdateReqShare) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DfsFTPShareUpdateReqShare) SetDescription(v string) {
	o.Description = &v
}

// GetDfsGatewayGroupId returns the DfsGatewayGroupId field value if set, zero value otherwise.
func (o *DfsFTPShareUpdateReqShare) GetDfsGatewayGroupId() int64 {
	if o == nil || IsNil(o.DfsGatewayGroupId) {
		var ret int64
		return ret
	}
	return *o.DfsGatewayGroupId
}

// GetDfsGatewayGroupIdOk returns a tuple with the DfsGatewayGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFTPShareUpdateReqShare) GetDfsGatewayGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DfsGatewayGroupId) {
		return nil, false
	}
	return o.DfsGatewayGroupId, true
}

// HasDfsGatewayGroupId returns a boolean if a field has been set.
func (o *DfsFTPShareUpdateReqShare) HasDfsGatewayGroupId() bool {
	if o != nil && !IsNil(o.DfsGatewayGroupId) {
		return true
	}

	return false
}

// SetDfsGatewayGroupId gets a reference to the given int64 and assigns it to the DfsGatewayGroupId field.
func (o *DfsFTPShareUpdateReqShare) SetDfsGatewayGroupId(v int64) {
	o.DfsGatewayGroupId = &v
}

func (o DfsFTPShareUpdateReqShare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsFTPShareUpdateReqShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DfsGatewayGroupId) {
		toSerialize["dfs_gateway_group_id"] = o.DfsGatewayGroupId
	}
	return toSerialize, nil
}

type NullableDfsFTPShareUpdateReqShare struct {
	value *DfsFTPShareUpdateReqShare
	isSet bool
}

func (v NullableDfsFTPShareUpdateReqShare) Get() *DfsFTPShareUpdateReqShare {
	return v.value
}

func (v *NullableDfsFTPShareUpdateReqShare) Set(val *DfsFTPShareUpdateReqShare) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsFTPShareUpdateReqShare) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsFTPShareUpdateReqShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsFTPShareUpdateReqShare(val *DfsFTPShareUpdateReqShare) *NullableDfsFTPShareUpdateReqShare {
	return &NullableDfsFTPShareUpdateReqShare{value: val, isSet: true}
}

func (v NullableDfsFTPShareUpdateReqShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsFTPShareUpdateReqShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


