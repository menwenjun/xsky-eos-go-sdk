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

// checks if the LicenseResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseResp{}

// LicenseResp struct for LicenseResp
type LicenseResp struct {
	License License `json:"license"`
}

type _LicenseResp LicenseResp

// NewLicenseResp instantiates a new LicenseResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseResp(license License) *LicenseResp {
	this := LicenseResp{}
	this.License = license
	return &this
}

// NewLicenseRespWithDefaults instantiates a new LicenseResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseRespWithDefaults() *LicenseResp {
	this := LicenseResp{}
	return &this
}

// GetLicense returns the License field value
func (o *LicenseResp) GetLicense() License {
	if o == nil {
		var ret License
		return ret
	}

	return o.License
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
func (o *LicenseResp) GetLicenseOk() (*License, bool) {
	if o == nil {
		return nil, false
	}
	return &o.License, true
}

// SetLicense sets field value
func (o *LicenseResp) SetLicense(v License) {
	o.License = v
}

func (o LicenseResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["license"] = o.License
	return toSerialize, nil
}

func (o *LicenseResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"license",
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

	varLicenseResp := _LicenseResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLicenseResp)

	if err != nil {
		return err
	}

	*o = LicenseResp(varLicenseResp)

	return err
}

type NullableLicenseResp struct {
	value *LicenseResp
	isSet bool
}

func (v NullableLicenseResp) Get() *LicenseResp {
	return v.value
}

func (v *NullableLicenseResp) Set(val *LicenseResp) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseResp) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseResp(val *LicenseResp) *NullableLicenseResp {
	return &NullableLicenseResp{value: val, isSet: true}
}

func (v NullableLicenseResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


