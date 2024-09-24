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

// checks if the OSSearchEngineCreateReqSearchEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSSearchEngineCreateReqSearchEngine{}

// OSSearchEngineCreateReqSearchEngine struct for OSSearchEngineCreateReqSearchEngine
type OSSearchEngineCreateReqSearchEngine struct {
	GatewayDataSize int64 `json:"gateway_data_size"`
	GatewayFlavorId int64 `json:"gateway_flavor_id"`
	GatewayHttpPort *int64 `json:"gateway_http_port,omitempty"`
	GatewayTransportPort *int64 `json:"gateway_transport_port,omitempty"`
	OsSearchGateways []OSSearchGatewayReq `json:"os_search_gateways"`
}

type _OSSearchEngineCreateReqSearchEngine OSSearchEngineCreateReqSearchEngine

// NewOSSearchEngineCreateReqSearchEngine instantiates a new OSSearchEngineCreateReqSearchEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSSearchEngineCreateReqSearchEngine(gatewayDataSize int64, gatewayFlavorId int64, osSearchGateways []OSSearchGatewayReq) *OSSearchEngineCreateReqSearchEngine {
	this := OSSearchEngineCreateReqSearchEngine{}
	this.GatewayDataSize = gatewayDataSize
	this.GatewayFlavorId = gatewayFlavorId
	this.OsSearchGateways = osSearchGateways
	return &this
}

// NewOSSearchEngineCreateReqSearchEngineWithDefaults instantiates a new OSSearchEngineCreateReqSearchEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSSearchEngineCreateReqSearchEngineWithDefaults() *OSSearchEngineCreateReqSearchEngine {
	this := OSSearchEngineCreateReqSearchEngine{}
	return &this
}

// GetGatewayDataSize returns the GatewayDataSize field value
func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayDataSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayDataSize
}

// GetGatewayDataSizeOk returns a tuple with the GatewayDataSize field value
// and a boolean to check if the value has been set.
func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayDataSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayDataSize, true
}

// SetGatewayDataSize sets field value
func (o *OSSearchEngineCreateReqSearchEngine) SetGatewayDataSize(v int64) {
	o.GatewayDataSize = v
}

// GetGatewayFlavorId returns the GatewayFlavorId field value
func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayFlavorId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayFlavorId
}

// GetGatewayFlavorIdOk returns a tuple with the GatewayFlavorId field value
// and a boolean to check if the value has been set.
func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayFlavorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayFlavorId, true
}

// SetGatewayFlavorId sets field value
func (o *OSSearchEngineCreateReqSearchEngine) SetGatewayFlavorId(v int64) {
	o.GatewayFlavorId = v
}

// GetGatewayHttpPort returns the GatewayHttpPort field value if set, zero value otherwise.
func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayHttpPort() int64 {
	if o == nil || IsNil(o.GatewayHttpPort) {
		var ret int64
		return ret
	}
	return *o.GatewayHttpPort
}

// GetGatewayHttpPortOk returns a tuple with the GatewayHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayHttpPortOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayHttpPort) {
		return nil, false
	}
	return o.GatewayHttpPort, true
}

// HasGatewayHttpPort returns a boolean if a field has been set.
func (o *OSSearchEngineCreateReqSearchEngine) HasGatewayHttpPort() bool {
	if o != nil && !IsNil(o.GatewayHttpPort) {
		return true
	}

	return false
}

// SetGatewayHttpPort gets a reference to the given int64 and assigns it to the GatewayHttpPort field.
func (o *OSSearchEngineCreateReqSearchEngine) SetGatewayHttpPort(v int64) {
	o.GatewayHttpPort = &v
}

// GetGatewayTransportPort returns the GatewayTransportPort field value if set, zero value otherwise.
func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayTransportPort() int64 {
	if o == nil || IsNil(o.GatewayTransportPort) {
		var ret int64
		return ret
	}
	return *o.GatewayTransportPort
}

// GetGatewayTransportPortOk returns a tuple with the GatewayTransportPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchEngineCreateReqSearchEngine) GetGatewayTransportPortOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayTransportPort) {
		return nil, false
	}
	return o.GatewayTransportPort, true
}

// HasGatewayTransportPort returns a boolean if a field has been set.
func (o *OSSearchEngineCreateReqSearchEngine) HasGatewayTransportPort() bool {
	if o != nil && !IsNil(o.GatewayTransportPort) {
		return true
	}

	return false
}

// SetGatewayTransportPort gets a reference to the given int64 and assigns it to the GatewayTransportPort field.
func (o *OSSearchEngineCreateReqSearchEngine) SetGatewayTransportPort(v int64) {
	o.GatewayTransportPort = &v
}

// GetOsSearchGateways returns the OsSearchGateways field value
func (o *OSSearchEngineCreateReqSearchEngine) GetOsSearchGateways() []OSSearchGatewayReq {
	if o == nil {
		var ret []OSSearchGatewayReq
		return ret
	}

	return o.OsSearchGateways
}

// GetOsSearchGatewaysOk returns a tuple with the OsSearchGateways field value
// and a boolean to check if the value has been set.
func (o *OSSearchEngineCreateReqSearchEngine) GetOsSearchGatewaysOk() ([]OSSearchGatewayReq, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsSearchGateways, true
}

// SetOsSearchGateways sets field value
func (o *OSSearchEngineCreateReqSearchEngine) SetOsSearchGateways(v []OSSearchGatewayReq) {
	o.OsSearchGateways = v
}

func (o OSSearchEngineCreateReqSearchEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSSearchEngineCreateReqSearchEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gateway_data_size"] = o.GatewayDataSize
	toSerialize["gateway_flavor_id"] = o.GatewayFlavorId
	if !IsNil(o.GatewayHttpPort) {
		toSerialize["gateway_http_port"] = o.GatewayHttpPort
	}
	if !IsNil(o.GatewayTransportPort) {
		toSerialize["gateway_transport_port"] = o.GatewayTransportPort
	}
	toSerialize["os_search_gateways"] = o.OsSearchGateways
	return toSerialize, nil
}

func (o *OSSearchEngineCreateReqSearchEngine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gateway_data_size",
		"gateway_flavor_id",
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

	varOSSearchEngineCreateReqSearchEngine := _OSSearchEngineCreateReqSearchEngine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSSearchEngineCreateReqSearchEngine)

	if err != nil {
		return err
	}

	*o = OSSearchEngineCreateReqSearchEngine(varOSSearchEngineCreateReqSearchEngine)

	return err
}

type NullableOSSearchEngineCreateReqSearchEngine struct {
	value *OSSearchEngineCreateReqSearchEngine
	isSet bool
}

func (v NullableOSSearchEngineCreateReqSearchEngine) Get() *OSSearchEngineCreateReqSearchEngine {
	return v.value
}

func (v *NullableOSSearchEngineCreateReqSearchEngine) Set(val *OSSearchEngineCreateReqSearchEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableOSSearchEngineCreateReqSearchEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableOSSearchEngineCreateReqSearchEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSSearchEngineCreateReqSearchEngine(val *OSSearchEngineCreateReqSearchEngine) *NullableOSSearchEngineCreateReqSearchEngine {
	return &NullableOSSearchEngineCreateReqSearchEngine{value: val, isSet: true}
}

func (v NullableOSSearchEngineCreateReqSearchEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSSearchEngineCreateReqSearchEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


