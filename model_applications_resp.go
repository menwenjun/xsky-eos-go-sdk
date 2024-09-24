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

// checks if the ApplicationsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationsResp{}

// ApplicationsResp struct for ApplicationsResp
type ApplicationsResp struct {
	Applications []Application `json:"applications,omitempty"`
}

// NewApplicationsResp instantiates a new ApplicationsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationsResp() *ApplicationsResp {
	this := ApplicationsResp{}
	return &this
}

// NewApplicationsRespWithDefaults instantiates a new ApplicationsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationsRespWithDefaults() *ApplicationsResp {
	this := ApplicationsResp{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *ApplicationsResp) GetApplications() []Application {
	if o == nil || IsNil(o.Applications) {
		var ret []Application
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResp) GetApplicationsOk() ([]Application, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *ApplicationsResp) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []Application and assigns it to the Applications field.
func (o *ApplicationsResp) SetApplications(v []Application) {
	o.Applications = v
}

func (o ApplicationsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	return toSerialize, nil
}

type NullableApplicationsResp struct {
	value *ApplicationsResp
	isSet bool
}

func (v NullableApplicationsResp) Get() *ApplicationsResp {
	return v.value
}

func (v *NullableApplicationsResp) Set(val *ApplicationsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationsResp(val *ApplicationsResp) *NullableApplicationsResp {
	return &NullableApplicationsResp{value: val, isSet: true}
}

func (v NullableApplicationsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


