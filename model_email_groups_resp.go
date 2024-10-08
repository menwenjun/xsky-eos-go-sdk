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

// checks if the EmailGroupsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailGroupsResp{}

// EmailGroupsResp struct for EmailGroupsResp
type EmailGroupsResp struct {
	// email groups
	EmailGroups []EmailGroupRecord `json:"email_groups"`
}

type _EmailGroupsResp EmailGroupsResp

// NewEmailGroupsResp instantiates a new EmailGroupsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailGroupsResp(emailGroups []EmailGroupRecord) *EmailGroupsResp {
	this := EmailGroupsResp{}
	this.EmailGroups = emailGroups
	return &this
}

// NewEmailGroupsRespWithDefaults instantiates a new EmailGroupsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailGroupsRespWithDefaults() *EmailGroupsResp {
	this := EmailGroupsResp{}
	return &this
}

// GetEmailGroups returns the EmailGroups field value
func (o *EmailGroupsResp) GetEmailGroups() []EmailGroupRecord {
	if o == nil {
		var ret []EmailGroupRecord
		return ret
	}

	return o.EmailGroups
}

// GetEmailGroupsOk returns a tuple with the EmailGroups field value
// and a boolean to check if the value has been set.
func (o *EmailGroupsResp) GetEmailGroupsOk() ([]EmailGroupRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailGroups, true
}

// SetEmailGroups sets field value
func (o *EmailGroupsResp) SetEmailGroups(v []EmailGroupRecord) {
	o.EmailGroups = v
}

func (o EmailGroupsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailGroupsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email_groups"] = o.EmailGroups
	return toSerialize, nil
}

func (o *EmailGroupsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email_groups",
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

	varEmailGroupsResp := _EmailGroupsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailGroupsResp)

	if err != nil {
		return err
	}

	*o = EmailGroupsResp(varEmailGroupsResp)

	return err
}

type NullableEmailGroupsResp struct {
	value *EmailGroupsResp
	isSet bool
}

func (v NullableEmailGroupsResp) Get() *EmailGroupsResp {
	return v.value
}

func (v *NullableEmailGroupsResp) Set(val *EmailGroupsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailGroupsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailGroupsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailGroupsResp(val *EmailGroupsResp) *NullableEmailGroupsResp {
	return &NullableEmailGroupsResp{value: val, isSet: true}
}

func (v NullableEmailGroupsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailGroupsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


