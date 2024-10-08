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

// checks if the NetworkInterfaceStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkInterfaceStat{}

// NetworkInterfaceStat NetworkInterfaceStat define the statistics of network interface
type NetworkInterfaceStat struct {
	Create *time.Time `json:"create,omitempty"`
	DropPps *float64 `json:"drop_pps,omitempty"`
	ErrorPps *float64 `json:"error_pps,omitempty"`
	RxBandwidthKbyte *float64 `json:"rx_bandwidth_kbyte,omitempty"`
	RxDropPps *float64 `json:"rx_drop_pps,omitempty"`
	RxErrorPps *float64 `json:"rx_error_pps,omitempty"`
	RxPps *float64 `json:"rx_pps,omitempty"`
	TxBandwidthKbyte *float64 `json:"tx_bandwidth_kbyte,omitempty"`
	TxDropPps *float64 `json:"tx_drop_pps,omitempty"`
	TxErrorPps *float64 `json:"tx_error_pps,omitempty"`
	TxPps *float64 `json:"tx_pps,omitempty"`
}

// NewNetworkInterfaceStat instantiates a new NetworkInterfaceStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkInterfaceStat() *NetworkInterfaceStat {
	this := NetworkInterfaceStat{}
	return &this
}

// NewNetworkInterfaceStatWithDefaults instantiates a new NetworkInterfaceStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkInterfaceStatWithDefaults() *NetworkInterfaceStat {
	this := NetworkInterfaceStat{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *NetworkInterfaceStat) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDropPps returns the DropPps field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetDropPps() float64 {
	if o == nil || IsNil(o.DropPps) {
		var ret float64
		return ret
	}
	return *o.DropPps
}

// GetDropPpsOk returns a tuple with the DropPps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetDropPpsOk() (*float64, bool) {
	if o == nil || IsNil(o.DropPps) {
		return nil, false
	}
	return o.DropPps, true
}

// HasDropPps returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasDropPps() bool {
	if o != nil && !IsNil(o.DropPps) {
		return true
	}

	return false
}

// SetDropPps gets a reference to the given float64 and assigns it to the DropPps field.
func (o *NetworkInterfaceStat) SetDropPps(v float64) {
	o.DropPps = &v
}

// GetErrorPps returns the ErrorPps field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetErrorPps() float64 {
	if o == nil || IsNil(o.ErrorPps) {
		var ret float64
		return ret
	}
	return *o.ErrorPps
}

// GetErrorPpsOk returns a tuple with the ErrorPps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetErrorPpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ErrorPps) {
		return nil, false
	}
	return o.ErrorPps, true
}

// HasErrorPps returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasErrorPps() bool {
	if o != nil && !IsNil(o.ErrorPps) {
		return true
	}

	return false
}

// SetErrorPps gets a reference to the given float64 and assigns it to the ErrorPps field.
func (o *NetworkInterfaceStat) SetErrorPps(v float64) {
	o.ErrorPps = &v
}

// GetRxBandwidthKbyte returns the RxBandwidthKbyte field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetRxBandwidthKbyte() float64 {
	if o == nil || IsNil(o.RxBandwidthKbyte) {
		var ret float64
		return ret
	}
	return *o.RxBandwidthKbyte
}

// GetRxBandwidthKbyteOk returns a tuple with the RxBandwidthKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetRxBandwidthKbyteOk() (*float64, bool) {
	if o == nil || IsNil(o.RxBandwidthKbyte) {
		return nil, false
	}
	return o.RxBandwidthKbyte, true
}

// HasRxBandwidthKbyte returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasRxBandwidthKbyte() bool {
	if o != nil && !IsNil(o.RxBandwidthKbyte) {
		return true
	}

	return false
}

// SetRxBandwidthKbyte gets a reference to the given float64 and assigns it to the RxBandwidthKbyte field.
func (o *NetworkInterfaceStat) SetRxBandwidthKbyte(v float64) {
	o.RxBandwidthKbyte = &v
}

// GetRxDropPps returns the RxDropPps field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetRxDropPps() float64 {
	if o == nil || IsNil(o.RxDropPps) {
		var ret float64
		return ret
	}
	return *o.RxDropPps
}

// GetRxDropPpsOk returns a tuple with the RxDropPps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetRxDropPpsOk() (*float64, bool) {
	if o == nil || IsNil(o.RxDropPps) {
		return nil, false
	}
	return o.RxDropPps, true
}

// HasRxDropPps returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasRxDropPps() bool {
	if o != nil && !IsNil(o.RxDropPps) {
		return true
	}

	return false
}

// SetRxDropPps gets a reference to the given float64 and assigns it to the RxDropPps field.
func (o *NetworkInterfaceStat) SetRxDropPps(v float64) {
	o.RxDropPps = &v
}

// GetRxErrorPps returns the RxErrorPps field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetRxErrorPps() float64 {
	if o == nil || IsNil(o.RxErrorPps) {
		var ret float64
		return ret
	}
	return *o.RxErrorPps
}

// GetRxErrorPpsOk returns a tuple with the RxErrorPps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetRxErrorPpsOk() (*float64, bool) {
	if o == nil || IsNil(o.RxErrorPps) {
		return nil, false
	}
	return o.RxErrorPps, true
}

// HasRxErrorPps returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasRxErrorPps() bool {
	if o != nil && !IsNil(o.RxErrorPps) {
		return true
	}

	return false
}

// SetRxErrorPps gets a reference to the given float64 and assigns it to the RxErrorPps field.
func (o *NetworkInterfaceStat) SetRxErrorPps(v float64) {
	o.RxErrorPps = &v
}

// GetRxPps returns the RxPps field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetRxPps() float64 {
	if o == nil || IsNil(o.RxPps) {
		var ret float64
		return ret
	}
	return *o.RxPps
}

// GetRxPpsOk returns a tuple with the RxPps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetRxPpsOk() (*float64, bool) {
	if o == nil || IsNil(o.RxPps) {
		return nil, false
	}
	return o.RxPps, true
}

// HasRxPps returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasRxPps() bool {
	if o != nil && !IsNil(o.RxPps) {
		return true
	}

	return false
}

// SetRxPps gets a reference to the given float64 and assigns it to the RxPps field.
func (o *NetworkInterfaceStat) SetRxPps(v float64) {
	o.RxPps = &v
}

// GetTxBandwidthKbyte returns the TxBandwidthKbyte field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetTxBandwidthKbyte() float64 {
	if o == nil || IsNil(o.TxBandwidthKbyte) {
		var ret float64
		return ret
	}
	return *o.TxBandwidthKbyte
}

// GetTxBandwidthKbyteOk returns a tuple with the TxBandwidthKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetTxBandwidthKbyteOk() (*float64, bool) {
	if o == nil || IsNil(o.TxBandwidthKbyte) {
		return nil, false
	}
	return o.TxBandwidthKbyte, true
}

// HasTxBandwidthKbyte returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasTxBandwidthKbyte() bool {
	if o != nil && !IsNil(o.TxBandwidthKbyte) {
		return true
	}

	return false
}

// SetTxBandwidthKbyte gets a reference to the given float64 and assigns it to the TxBandwidthKbyte field.
func (o *NetworkInterfaceStat) SetTxBandwidthKbyte(v float64) {
	o.TxBandwidthKbyte = &v
}

// GetTxDropPps returns the TxDropPps field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetTxDropPps() float64 {
	if o == nil || IsNil(o.TxDropPps) {
		var ret float64
		return ret
	}
	return *o.TxDropPps
}

// GetTxDropPpsOk returns a tuple with the TxDropPps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetTxDropPpsOk() (*float64, bool) {
	if o == nil || IsNil(o.TxDropPps) {
		return nil, false
	}
	return o.TxDropPps, true
}

// HasTxDropPps returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasTxDropPps() bool {
	if o != nil && !IsNil(o.TxDropPps) {
		return true
	}

	return false
}

// SetTxDropPps gets a reference to the given float64 and assigns it to the TxDropPps field.
func (o *NetworkInterfaceStat) SetTxDropPps(v float64) {
	o.TxDropPps = &v
}

// GetTxErrorPps returns the TxErrorPps field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetTxErrorPps() float64 {
	if o == nil || IsNil(o.TxErrorPps) {
		var ret float64
		return ret
	}
	return *o.TxErrorPps
}

// GetTxErrorPpsOk returns a tuple with the TxErrorPps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetTxErrorPpsOk() (*float64, bool) {
	if o == nil || IsNil(o.TxErrorPps) {
		return nil, false
	}
	return o.TxErrorPps, true
}

// HasTxErrorPps returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasTxErrorPps() bool {
	if o != nil && !IsNil(o.TxErrorPps) {
		return true
	}

	return false
}

// SetTxErrorPps gets a reference to the given float64 and assigns it to the TxErrorPps field.
func (o *NetworkInterfaceStat) SetTxErrorPps(v float64) {
	o.TxErrorPps = &v
}

// GetTxPps returns the TxPps field value if set, zero value otherwise.
func (o *NetworkInterfaceStat) GetTxPps() float64 {
	if o == nil || IsNil(o.TxPps) {
		var ret float64
		return ret
	}
	return *o.TxPps
}

// GetTxPpsOk returns a tuple with the TxPps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceStat) GetTxPpsOk() (*float64, bool) {
	if o == nil || IsNil(o.TxPps) {
		return nil, false
	}
	return o.TxPps, true
}

// HasTxPps returns a boolean if a field has been set.
func (o *NetworkInterfaceStat) HasTxPps() bool {
	if o != nil && !IsNil(o.TxPps) {
		return true
	}

	return false
}

// SetTxPps gets a reference to the given float64 and assigns it to the TxPps field.
func (o *NetworkInterfaceStat) SetTxPps(v float64) {
	o.TxPps = &v
}

func (o NetworkInterfaceStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkInterfaceStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DropPps) {
		toSerialize["drop_pps"] = o.DropPps
	}
	if !IsNil(o.ErrorPps) {
		toSerialize["error_pps"] = o.ErrorPps
	}
	if !IsNil(o.RxBandwidthKbyte) {
		toSerialize["rx_bandwidth_kbyte"] = o.RxBandwidthKbyte
	}
	if !IsNil(o.RxDropPps) {
		toSerialize["rx_drop_pps"] = o.RxDropPps
	}
	if !IsNil(o.RxErrorPps) {
		toSerialize["rx_error_pps"] = o.RxErrorPps
	}
	if !IsNil(o.RxPps) {
		toSerialize["rx_pps"] = o.RxPps
	}
	if !IsNil(o.TxBandwidthKbyte) {
		toSerialize["tx_bandwidth_kbyte"] = o.TxBandwidthKbyte
	}
	if !IsNil(o.TxDropPps) {
		toSerialize["tx_drop_pps"] = o.TxDropPps
	}
	if !IsNil(o.TxErrorPps) {
		toSerialize["tx_error_pps"] = o.TxErrorPps
	}
	if !IsNil(o.TxPps) {
		toSerialize["tx_pps"] = o.TxPps
	}
	return toSerialize, nil
}

type NullableNetworkInterfaceStat struct {
	value *NetworkInterfaceStat
	isSet bool
}

func (v NullableNetworkInterfaceStat) Get() *NetworkInterfaceStat {
	return v.value
}

func (v *NullableNetworkInterfaceStat) Set(val *NetworkInterfaceStat) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkInterfaceStat) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkInterfaceStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkInterfaceStat(val *NetworkInterfaceStat) *NullableNetworkInterfaceStat {
	return &NullableNetworkInterfaceStat{value: val, isSet: true}
}

func (v NullableNetworkInterfaceStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkInterfaceStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


