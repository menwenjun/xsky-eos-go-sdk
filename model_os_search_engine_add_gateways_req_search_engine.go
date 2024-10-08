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

// checks if the OSSearchEngineAddGatewaysReqSearchEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSSearchEngineAddGatewaysReqSearchEngine{}

// OSSearchEngineAddGatewaysReqSearchEngine struct for OSSearchEngineAddGatewaysReqSearchEngine
type OSSearchEngineAddGatewaysReqSearchEngine struct {
	OsSearchGateways []OSSearchGatewayReq `json:"os_search_gateways"`
}

type _OSSearchEngineAddGatewaysReqSearchEngine OSSearchEngineAddGatewaysReqSearchEngine

// NewOSSearchEngineAddGatewaysReqSearchEngine instantiates a new OSSearchEngineAddGatewaysReqSearchEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSSearchEngineAddGatewaysReqSearchEngine(osSearchGateways []OSSearchGatewayReq) *OSSearchEngineAddGatewaysReqSearchEngine {
	this := OSSearchEngineAddGatewaysReqSearchEngine{}
	this.OsSearchGateways = osSearchGateways
	return &this
}

// NewOSSearchEngineAddGatewaysReqSearchEngineWithDefaults instantiates a new OSSearchEngineAddGatewaysReqSearchEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSSearchEngineAddGatewaysReqSearchEngineWithDefaults() *OSSearchEngineAddGatewaysReqSearchEngine {
	this := OSSearchEngineAddGatewaysReqSearchEngine{}
	return &this
}

// GetOsSearchGateways returns the OsSearchGateways field value
func (o *OSSearchEngineAddGatewaysReqSearchEngine) GetOsSearchGateways() []OSSearchGatewayReq {
	if o == nil {
		var ret []OSSearchGatewayReq
		return ret
	}

	return o.OsSearchGateways
}

// GetOsSearchGatewaysOk returns a tuple with the OsSearchGateways field value
// and a boolean to check if the value has been set.
func (o *OSSearchEngineAddGatewaysReqSearchEngine) GetOsSearchGatewaysOk() ([]OSSearchGatewayReq, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsSearchGateways, true
}

// SetOsSearchGateways sets field value
func (o *OSSearchEngineAddGatewaysReqSearchEngine) SetOsSearchGateways(v []OSSearchGatewayReq) {
	o.OsSearchGateways = v
}

func (o OSSearchEngineAddGatewaysReqSearchEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSSearchEngineAddGatewaysReqSearchEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["os_search_gateways"] = o.OsSearchGateways
	return toSerialize, nil
}

func (o *OSSearchEngineAddGatewaysReqSearchEngine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"os_search_gateways",
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

	varOSSearchEngineAddGatewaysReqSearchEngine := _OSSearchEngineAddGatewaysReqSearchEngine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSSearchEngineAddGatewaysReqSearchEngine)

	if err != nil {
		return err
	}

	*o = OSSearchEngineAddGatewaysReqSearchEngine(varOSSearchEngineAddGatewaysReqSearchEngine)

	return err
}

type NullableOSSearchEngineAddGatewaysReqSearchEngine struct {
	value *OSSearchEngineAddGatewaysReqSearchEngine
	isSet bool
}

func (v NullableOSSearchEngineAddGatewaysReqSearchEngine) Get() *OSSearchEngineAddGatewaysReqSearchEngine {
	return v.value
}

func (v *NullableOSSearchEngineAddGatewaysReqSearchEngine) Set(val *OSSearchEngineAddGatewaysReqSearchEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableOSSearchEngineAddGatewaysReqSearchEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableOSSearchEngineAddGatewaysReqSearchEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSSearchEngineAddGatewaysReqSearchEngine(val *OSSearchEngineAddGatewaysReqSearchEngine) *NullableOSSearchEngineAddGatewaysReqSearchEngine {
	return &NullableOSSearchEngineAddGatewaysReqSearchEngine{value: val, isSet: true}
}

func (v NullableOSSearchEngineAddGatewaysReqSearchEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSSearchEngineAddGatewaysReqSearchEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


