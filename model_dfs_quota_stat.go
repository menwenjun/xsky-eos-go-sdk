/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the DfsQuotaStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsQuotaStat{}

// DfsQuotaStat DfsQuotaStat records dfs gateway stat info
type DfsQuotaStat struct {
	Create *time.Time `json:"create,omitempty"`
	UsedFiles *int64 `json:"used_files,omitempty"`
	UsedKbyte *int64 `json:"used_kbyte,omitempty"`
}

// NewDfsQuotaStat instantiates a new DfsQuotaStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsQuotaStat() *DfsQuotaStat {
	this := DfsQuotaStat{}
	return &this
}

// NewDfsQuotaStatWithDefaults instantiates a new DfsQuotaStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsQuotaStatWithDefaults() *DfsQuotaStat {
	this := DfsQuotaStat{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsQuotaStat) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuotaStat) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsQuotaStat) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsQuotaStat) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUsedFiles returns the UsedFiles field value if set, zero value otherwise.
func (o *DfsQuotaStat) GetUsedFiles() int64 {
	if o == nil || IsNil(o.UsedFiles) {
		var ret int64
		return ret
	}
	return *o.UsedFiles
}

// GetUsedFilesOk returns a tuple with the UsedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuotaStat) GetUsedFilesOk() (*int64, bool) {
	if o == nil || IsNil(o.UsedFiles) {
		return nil, false
	}
	return o.UsedFiles, true
}

// HasUsedFiles returns a boolean if a field has been set.
func (o *DfsQuotaStat) HasUsedFiles() bool {
	if o != nil && !IsNil(o.UsedFiles) {
		return true
	}

	return false
}

// SetUsedFiles gets a reference to the given int64 and assigns it to the UsedFiles field.
func (o *DfsQuotaStat) SetUsedFiles(v int64) {
	o.UsedFiles = &v
}

// GetUsedKbyte returns the UsedKbyte field value if set, zero value otherwise.
func (o *DfsQuotaStat) GetUsedKbyte() int64 {
	if o == nil || IsNil(o.UsedKbyte) {
		var ret int64
		return ret
	}
	return *o.UsedKbyte
}

// GetUsedKbyteOk returns a tuple with the UsedKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQuotaStat) GetUsedKbyteOk() (*int64, bool) {
	if o == nil || IsNil(o.UsedKbyte) {
		return nil, false
	}
	return o.UsedKbyte, true
}

// HasUsedKbyte returns a boolean if a field has been set.
func (o *DfsQuotaStat) HasUsedKbyte() bool {
	if o != nil && !IsNil(o.UsedKbyte) {
		return true
	}

	return false
}

// SetUsedKbyte gets a reference to the given int64 and assigns it to the UsedKbyte field.
func (o *DfsQuotaStat) SetUsedKbyte(v int64) {
	o.UsedKbyte = &v
}

func (o DfsQuotaStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsQuotaStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.UsedFiles) {
		toSerialize["used_files"] = o.UsedFiles
	}
	if !IsNil(o.UsedKbyte) {
		toSerialize["used_kbyte"] = o.UsedKbyte
	}
	return toSerialize, nil
}

type NullableDfsQuotaStat struct {
	value *DfsQuotaStat
	isSet bool
}

func (v NullableDfsQuotaStat) Get() *DfsQuotaStat {
	return v.value
}

func (v *NullableDfsQuotaStat) Set(val *DfsQuotaStat) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsQuotaStat) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsQuotaStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsQuotaStat(val *DfsQuotaStat) *NullableDfsQuotaStat {
	return &NullableDfsQuotaStat{value: val, isSet: true}
}

func (v NullableDfsQuotaStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsQuotaStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


