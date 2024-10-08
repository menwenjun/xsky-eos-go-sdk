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

// checks if the AdminVIPResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminVIPResult{}

// AdminVIPResult AdminVIPResult defines result for AdminVIPResp
type AdminVIPResult struct {
	AccessUrl *string `json:"access_url,omitempty"`
	Enable *bool `json:"enable,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Mask *int64 `json:"mask,omitempty"`
}

// NewAdminVIPResult instantiates a new AdminVIPResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminVIPResult() *AdminVIPResult {
	this := AdminVIPResult{}
	return &this
}

// NewAdminVIPResultWithDefaults instantiates a new AdminVIPResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminVIPResultWithDefaults() *AdminVIPResult {
	this := AdminVIPResult{}
	return &this
}

// GetAccessUrl returns the AccessUrl field value if set, zero value otherwise.
func (o *AdminVIPResult) GetAccessUrl() string {
	if o == nil || IsNil(o.AccessUrl) {
		var ret string
		return ret
	}
	return *o.AccessUrl
}

// GetAccessUrlOk returns a tuple with the AccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVIPResult) GetAccessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AccessUrl) {
		return nil, false
	}
	return o.AccessUrl, true
}

// HasAccessUrl returns a boolean if a field has been set.
func (o *AdminVIPResult) HasAccessUrl() bool {
	if o != nil && !IsNil(o.AccessUrl) {
		return true
	}

	return false
}

// SetAccessUrl gets a reference to the given string and assigns it to the AccessUrl field.
func (o *AdminVIPResult) SetAccessUrl(v string) {
	o.AccessUrl = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *AdminVIPResult) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVIPResult) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *AdminVIPResult) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *AdminVIPResult) SetEnable(v bool) {
	o.Enable = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *AdminVIPResult) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVIPResult) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *AdminVIPResult) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *AdminVIPResult) SetIp(v string) {
	o.Ip = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *AdminVIPResult) GetMask() int64 {
	if o == nil || IsNil(o.Mask) {
		var ret int64
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVIPResult) GetMaskOk() (*int64, bool) {
	if o == nil || IsNil(o.Mask) {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *AdminVIPResult) HasMask() bool {
	if o != nil && !IsNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int64 and assigns it to the Mask field.
func (o *AdminVIPResult) SetMask(v int64) {
	o.Mask = &v
}

func (o AdminVIPResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminVIPResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessUrl) {
		toSerialize["access_url"] = o.AccessUrl
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	return toSerialize, nil
}

type NullableAdminVIPResult struct {
	value *AdminVIPResult
	isSet bool
}

func (v NullableAdminVIPResult) Get() *AdminVIPResult {
	return v.value
}

func (v *NullableAdminVIPResult) Set(val *AdminVIPResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminVIPResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminVIPResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminVIPResult(val *AdminVIPResult) *NullableAdminVIPResult {
	return &NullableAdminVIPResult{value: val, isSet: true}
}

func (v NullableAdminVIPResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminVIPResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


