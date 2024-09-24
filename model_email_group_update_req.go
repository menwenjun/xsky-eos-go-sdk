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

// checks if the EmailGroupUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailGroupUpdateReq{}

// EmailGroupUpdateReq struct for EmailGroupUpdateReq
type EmailGroupUpdateReq struct {
	EmailGroup EmailGroupUpdateReqEmailGroup `json:"email_group"`
}

type _EmailGroupUpdateReq EmailGroupUpdateReq

// NewEmailGroupUpdateReq instantiates a new EmailGroupUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailGroupUpdateReq(emailGroup EmailGroupUpdateReqEmailGroup) *EmailGroupUpdateReq {
	this := EmailGroupUpdateReq{}
	this.EmailGroup = emailGroup
	return &this
}

// NewEmailGroupUpdateReqWithDefaults instantiates a new EmailGroupUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailGroupUpdateReqWithDefaults() *EmailGroupUpdateReq {
	this := EmailGroupUpdateReq{}
	return &this
}

// GetEmailGroup returns the EmailGroup field value
func (o *EmailGroupUpdateReq) GetEmailGroup() EmailGroupUpdateReqEmailGroup {
	if o == nil {
		var ret EmailGroupUpdateReqEmailGroup
		return ret
	}

	return o.EmailGroup
}

// GetEmailGroupOk returns a tuple with the EmailGroup field value
// and a boolean to check if the value has been set.
func (o *EmailGroupUpdateReq) GetEmailGroupOk() (*EmailGroupUpdateReqEmailGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailGroup, true
}

// SetEmailGroup sets field value
func (o *EmailGroupUpdateReq) SetEmailGroup(v EmailGroupUpdateReqEmailGroup) {
	o.EmailGroup = v
}

func (o EmailGroupUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailGroupUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email_group"] = o.EmailGroup
	return toSerialize, nil
}

func (o *EmailGroupUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email_group",
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

	varEmailGroupUpdateReq := _EmailGroupUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailGroupUpdateReq)

	if err != nil {
		return err
	}

	*o = EmailGroupUpdateReq(varEmailGroupUpdateReq)

	return err
}

type NullableEmailGroupUpdateReq struct {
	value *EmailGroupUpdateReq
	isSet bool
}

func (v NullableEmailGroupUpdateReq) Get() *EmailGroupUpdateReq {
	return v.value
}

func (v *NullableEmailGroupUpdateReq) Set(val *EmailGroupUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailGroupUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailGroupUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailGroupUpdateReq(val *EmailGroupUpdateReq) *NullableEmailGroupUpdateReq {
	return &NullableEmailGroupUpdateReq{value: val, isSet: true}
}

func (v NullableEmailGroupUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailGroupUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


