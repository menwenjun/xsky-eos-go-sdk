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

// checks if the ErrorRecordResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorRecordResp{}

// ErrorRecordResp struct for ErrorRecordResp
type ErrorRecordResp struct {
	ErrorRecord ErrorRecord `json:"error_record"`
}

type _ErrorRecordResp ErrorRecordResp

// NewErrorRecordResp instantiates a new ErrorRecordResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorRecordResp(errorRecord ErrorRecord) *ErrorRecordResp {
	this := ErrorRecordResp{}
	this.ErrorRecord = errorRecord
	return &this
}

// NewErrorRecordRespWithDefaults instantiates a new ErrorRecordResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorRecordRespWithDefaults() *ErrorRecordResp {
	this := ErrorRecordResp{}
	return &this
}

// GetErrorRecord returns the ErrorRecord field value
func (o *ErrorRecordResp) GetErrorRecord() ErrorRecord {
	if o == nil {
		var ret ErrorRecord
		return ret
	}

	return o.ErrorRecord
}

// GetErrorRecordOk returns a tuple with the ErrorRecord field value
// and a boolean to check if the value has been set.
func (o *ErrorRecordResp) GetErrorRecordOk() (*ErrorRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorRecord, true
}

// SetErrorRecord sets field value
func (o *ErrorRecordResp) SetErrorRecord(v ErrorRecord) {
	o.ErrorRecord = v
}

func (o ErrorRecordResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorRecordResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_record"] = o.ErrorRecord
	return toSerialize, nil
}

func (o *ErrorRecordResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error_record",
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

	varErrorRecordResp := _ErrorRecordResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorRecordResp)

	if err != nil {
		return err
	}

	*o = ErrorRecordResp(varErrorRecordResp)

	return err
}

type NullableErrorRecordResp struct {
	value *ErrorRecordResp
	isSet bool
}

func (v NullableErrorRecordResp) Get() *ErrorRecordResp {
	return v.value
}

func (v *NullableErrorRecordResp) Set(val *ErrorRecordResp) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorRecordResp) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorRecordResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorRecordResp(val *ErrorRecordResp) *NullableErrorRecordResp {
	return &NullableErrorRecordResp{value: val, isSet: true}
}

func (v NullableErrorRecordResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorRecordResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


