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

// checks if the DfsGatewayGroupNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsGatewayGroupNestview{}

// DfsGatewayGroupNestview struct for DfsGatewayGroupNestview
type DfsGatewayGroupNestview struct {
	// hdfs domain auth type
	HdfsSecurities []string `json:"hdfs_securities,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewDfsGatewayGroupNestview instantiates a new DfsGatewayGroupNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsGatewayGroupNestview() *DfsGatewayGroupNestview {
	this := DfsGatewayGroupNestview{}
	return &this
}

// NewDfsGatewayGroupNestviewWithDefaults instantiates a new DfsGatewayGroupNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsGatewayGroupNestviewWithDefaults() *DfsGatewayGroupNestview {
	this := DfsGatewayGroupNestview{}
	return &this
}

// GetHdfsSecurities returns the HdfsSecurities field value if set, zero value otherwise.
func (o *DfsGatewayGroupNestview) GetHdfsSecurities() []string {
	if o == nil || IsNil(o.HdfsSecurities) {
		var ret []string
		return ret
	}
	return o.HdfsSecurities
}

// GetHdfsSecuritiesOk returns a tuple with the HdfsSecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroupNestview) GetHdfsSecuritiesOk() ([]string, bool) {
	if o == nil || IsNil(o.HdfsSecurities) {
		return nil, false
	}
	return o.HdfsSecurities, true
}

// HasHdfsSecurities returns a boolean if a field has been set.
func (o *DfsGatewayGroupNestview) HasHdfsSecurities() bool {
	if o != nil && !IsNil(o.HdfsSecurities) {
		return true
	}

	return false
}

// SetHdfsSecurities gets a reference to the given []string and assigns it to the HdfsSecurities field.
func (o *DfsGatewayGroupNestview) SetHdfsSecurities(v []string) {
	o.HdfsSecurities = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsGatewayGroupNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroupNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsGatewayGroupNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsGatewayGroupNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DfsGatewayGroupNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroupNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DfsGatewayGroupNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DfsGatewayGroupNestview) SetName(v string) {
	o.Name = &v
}

func (o DfsGatewayGroupNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsGatewayGroupNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HdfsSecurities) {
		toSerialize["hdfs_securities"] = o.HdfsSecurities
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDfsGatewayGroupNestview struct {
	value *DfsGatewayGroupNestview
	isSet bool
}

func (v NullableDfsGatewayGroupNestview) Get() *DfsGatewayGroupNestview {
	return v.value
}

func (v *NullableDfsGatewayGroupNestview) Set(val *DfsGatewayGroupNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsGatewayGroupNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsGatewayGroupNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsGatewayGroupNestview(val *DfsGatewayGroupNestview) *NullableDfsGatewayGroupNestview {
	return &NullableDfsGatewayGroupNestview{value: val, isSet: true}
}

func (v NullableDfsGatewayGroupNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsGatewayGroupNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


