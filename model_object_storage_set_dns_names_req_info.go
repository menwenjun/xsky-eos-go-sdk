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

// checks if the ObjectStorageSetDNSNamesReqInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageSetDNSNamesReqInfo{}

// ObjectStorageSetDNSNamesReqInfo struct for ObjectStorageSetDNSNamesReqInfo
type ObjectStorageSetDNSNamesReqInfo struct {
	CnameEnabled *bool `json:"cname_enabled,omitempty"`
	DnsNames []string `json:"dns_names,omitempty"`
}

// NewObjectStorageSetDNSNamesReqInfo instantiates a new ObjectStorageSetDNSNamesReqInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageSetDNSNamesReqInfo() *ObjectStorageSetDNSNamesReqInfo {
	this := ObjectStorageSetDNSNamesReqInfo{}
	return &this
}

// NewObjectStorageSetDNSNamesReqInfoWithDefaults instantiates a new ObjectStorageSetDNSNamesReqInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageSetDNSNamesReqInfoWithDefaults() *ObjectStorageSetDNSNamesReqInfo {
	this := ObjectStorageSetDNSNamesReqInfo{}
	return &this
}

// GetCnameEnabled returns the CnameEnabled field value if set, zero value otherwise.
func (o *ObjectStorageSetDNSNamesReqInfo) GetCnameEnabled() bool {
	if o == nil || IsNil(o.CnameEnabled) {
		var ret bool
		return ret
	}
	return *o.CnameEnabled
}

// GetCnameEnabledOk returns a tuple with the CnameEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageSetDNSNamesReqInfo) GetCnameEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CnameEnabled) {
		return nil, false
	}
	return o.CnameEnabled, true
}

// HasCnameEnabled returns a boolean if a field has been set.
func (o *ObjectStorageSetDNSNamesReqInfo) HasCnameEnabled() bool {
	if o != nil && !IsNil(o.CnameEnabled) {
		return true
	}

	return false
}

// SetCnameEnabled gets a reference to the given bool and assigns it to the CnameEnabled field.
func (o *ObjectStorageSetDNSNamesReqInfo) SetCnameEnabled(v bool) {
	o.CnameEnabled = &v
}

// GetDnsNames returns the DnsNames field value if set, zero value otherwise.
func (o *ObjectStorageSetDNSNamesReqInfo) GetDnsNames() []string {
	if o == nil || IsNil(o.DnsNames) {
		var ret []string
		return ret
	}
	return o.DnsNames
}

// GetDnsNamesOk returns a tuple with the DnsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageSetDNSNamesReqInfo) GetDnsNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsNames) {
		return nil, false
	}
	return o.DnsNames, true
}

// HasDnsNames returns a boolean if a field has been set.
func (o *ObjectStorageSetDNSNamesReqInfo) HasDnsNames() bool {
	if o != nil && !IsNil(o.DnsNames) {
		return true
	}

	return false
}

// SetDnsNames gets a reference to the given []string and assigns it to the DnsNames field.
func (o *ObjectStorageSetDNSNamesReqInfo) SetDnsNames(v []string) {
	o.DnsNames = v
}

func (o ObjectStorageSetDNSNamesReqInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageSetDNSNamesReqInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CnameEnabled) {
		toSerialize["cname_enabled"] = o.CnameEnabled
	}
	if !IsNil(o.DnsNames) {
		toSerialize["dns_names"] = o.DnsNames
	}
	return toSerialize, nil
}

type NullableObjectStorageSetDNSNamesReqInfo struct {
	value *ObjectStorageSetDNSNamesReqInfo
	isSet bool
}

func (v NullableObjectStorageSetDNSNamesReqInfo) Get() *ObjectStorageSetDNSNamesReqInfo {
	return v.value
}

func (v *NullableObjectStorageSetDNSNamesReqInfo) Set(val *ObjectStorageSetDNSNamesReqInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageSetDNSNamesReqInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageSetDNSNamesReqInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageSetDNSNamesReqInfo(val *ObjectStorageSetDNSNamesReqInfo) *NullableObjectStorageSetDNSNamesReqInfo {
	return &NullableObjectStorageSetDNSNamesReqInfo{value: val, isSet: true}
}

func (v NullableObjectStorageSetDNSNamesReqInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageSetDNSNamesReqInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


