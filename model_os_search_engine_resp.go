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

// checks if the OSSearchEngineResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSSearchEngineResp{}

// OSSearchEngineResp struct for OSSearchEngineResp
type OSSearchEngineResp struct {
	OsSearchEngine OSSearchEngineRecord `json:"os_search_engine"`
}

type _OSSearchEngineResp OSSearchEngineResp

// NewOSSearchEngineResp instantiates a new OSSearchEngineResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSSearchEngineResp(osSearchEngine OSSearchEngineRecord) *OSSearchEngineResp {
	this := OSSearchEngineResp{}
	this.OsSearchEngine = osSearchEngine
	return &this
}

// NewOSSearchEngineRespWithDefaults instantiates a new OSSearchEngineResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSSearchEngineRespWithDefaults() *OSSearchEngineResp {
	this := OSSearchEngineResp{}
	return &this
}

// GetOsSearchEngine returns the OsSearchEngine field value
func (o *OSSearchEngineResp) GetOsSearchEngine() OSSearchEngineRecord {
	if o == nil {
		var ret OSSearchEngineRecord
		return ret
	}

	return o.OsSearchEngine
}

// GetOsSearchEngineOk returns a tuple with the OsSearchEngine field value
// and a boolean to check if the value has been set.
func (o *OSSearchEngineResp) GetOsSearchEngineOk() (*OSSearchEngineRecord, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsSearchEngine, true
}

// SetOsSearchEngine sets field value
func (o *OSSearchEngineResp) SetOsSearchEngine(v OSSearchEngineRecord) {
	o.OsSearchEngine = v
}

func (o OSSearchEngineResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSSearchEngineResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_search_engine"] = o.OsSearchEngine
	return toSerialize, nil
}

func (o *OSSearchEngineResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_search_engine",
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

	varOSSearchEngineResp := _OSSearchEngineResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSSearchEngineResp)

	if err != nil {
		return err
	}

	*o = OSSearchEngineResp(varOSSearchEngineResp)

	return err
}

type NullableOSSearchEngineResp struct {
	value *OSSearchEngineResp
	isSet bool
}

func (v NullableOSSearchEngineResp) Get() *OSSearchEngineResp {
	return v.value
}

func (v *NullableOSSearchEngineResp) Set(val *OSSearchEngineResp) {
	v.value = val
	v.isSet = true
}

func (v NullableOSSearchEngineResp) IsSet() bool {
	return v.isSet
}

func (v *NullableOSSearchEngineResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSSearchEngineResp(val *OSSearchEngineResp) *NullableOSSearchEngineResp {
	return &NullableOSSearchEngineResp{value: val, isSet: true}
}

func (v NullableOSSearchEngineResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSSearchEngineResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


