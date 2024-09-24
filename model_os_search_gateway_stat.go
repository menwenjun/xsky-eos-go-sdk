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

// checks if the OSSearchGatewayStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSSearchGatewayStat{}

// OSSearchGatewayStat OSSearchGatewayStat defines model os OS search gateway stat
type OSSearchGatewayStat struct {
	CpuUtil *float64 `json:"cpu_util,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	MemUsagePercent *float64 `json:"mem_usage_percent,omitempty"`
	ReadBandwidthKbyte *float64 `json:"read_bandwidth_kbyte,omitempty"`
	ReadIops *float64 `json:"read_iops,omitempty"`
	ReadLatencyUs *float64 `json:"read_latency_us,omitempty"`
	WriteBandwidthKbyte *float64 `json:"write_bandwidth_kbyte,omitempty"`
	WriteIops *float64 `json:"write_iops,omitempty"`
	WriteLatencyUs *float64 `json:"write_latency_us,omitempty"`
}

// NewOSSearchGatewayStat instantiates a new OSSearchGatewayStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSSearchGatewayStat() *OSSearchGatewayStat {
	this := OSSearchGatewayStat{}
	return &this
}

// NewOSSearchGatewayStatWithDefaults instantiates a new OSSearchGatewayStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSSearchGatewayStatWithDefaults() *OSSearchGatewayStat {
	this := OSSearchGatewayStat{}
	return &this
}

// GetCpuUtil returns the CpuUtil field value if set, zero value otherwise.
func (o *OSSearchGatewayStat) GetCpuUtil() float64 {
	if o == nil || IsNil(o.CpuUtil) {
		var ret float64
		return ret
	}
	return *o.CpuUtil
}

// GetCpuUtilOk returns a tuple with the CpuUtil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayStat) GetCpuUtilOk() (*float64, bool) {
	if o == nil || IsNil(o.CpuUtil) {
		return nil, false
	}
	return o.CpuUtil, true
}

// HasCpuUtil returns a boolean if a field has been set.
func (o *OSSearchGatewayStat) HasCpuUtil() bool {
	if o != nil && !IsNil(o.CpuUtil) {
		return true
	}

	return false
}

// SetCpuUtil gets a reference to the given float64 and assigns it to the CpuUtil field.
func (o *OSSearchGatewayStat) SetCpuUtil(v float64) {
	o.CpuUtil = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OSSearchGatewayStat) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayStat) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OSSearchGatewayStat) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OSSearchGatewayStat) SetCreate(v time.Time) {
	o.Create = &v
}

// GetMemUsagePercent returns the MemUsagePercent field value if set, zero value otherwise.
func (o *OSSearchGatewayStat) GetMemUsagePercent() float64 {
	if o == nil || IsNil(o.MemUsagePercent) {
		var ret float64
		return ret
	}
	return *o.MemUsagePercent
}

// GetMemUsagePercentOk returns a tuple with the MemUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayStat) GetMemUsagePercentOk() (*float64, bool) {
	if o == nil || IsNil(o.MemUsagePercent) {
		return nil, false
	}
	return o.MemUsagePercent, true
}

// HasMemUsagePercent returns a boolean if a field has been set.
func (o *OSSearchGatewayStat) HasMemUsagePercent() bool {
	if o != nil && !IsNil(o.MemUsagePercent) {
		return true
	}

	return false
}

// SetMemUsagePercent gets a reference to the given float64 and assigns it to the MemUsagePercent field.
func (o *OSSearchGatewayStat) SetMemUsagePercent(v float64) {
	o.MemUsagePercent = &v
}

// GetReadBandwidthKbyte returns the ReadBandwidthKbyte field value if set, zero value otherwise.
func (o *OSSearchGatewayStat) GetReadBandwidthKbyte() float64 {
	if o == nil || IsNil(o.ReadBandwidthKbyte) {
		var ret float64
		return ret
	}
	return *o.ReadBandwidthKbyte
}

// GetReadBandwidthKbyteOk returns a tuple with the ReadBandwidthKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayStat) GetReadBandwidthKbyteOk() (*float64, bool) {
	if o == nil || IsNil(o.ReadBandwidthKbyte) {
		return nil, false
	}
	return o.ReadBandwidthKbyte, true
}

// HasReadBandwidthKbyte returns a boolean if a field has been set.
func (o *OSSearchGatewayStat) HasReadBandwidthKbyte() bool {
	if o != nil && !IsNil(o.ReadBandwidthKbyte) {
		return true
	}

	return false
}

// SetReadBandwidthKbyte gets a reference to the given float64 and assigns it to the ReadBandwidthKbyte field.
func (o *OSSearchGatewayStat) SetReadBandwidthKbyte(v float64) {
	o.ReadBandwidthKbyte = &v
}

// GetReadIops returns the ReadIops field value if set, zero value otherwise.
func (o *OSSearchGatewayStat) GetReadIops() float64 {
	if o == nil || IsNil(o.ReadIops) {
		var ret float64
		return ret
	}
	return *o.ReadIops
}

// GetReadIopsOk returns a tuple with the ReadIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayStat) GetReadIopsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReadIops) {
		return nil, false
	}
	return o.ReadIops, true
}

// HasReadIops returns a boolean if a field has been set.
func (o *OSSearchGatewayStat) HasReadIops() bool {
	if o != nil && !IsNil(o.ReadIops) {
		return true
	}

	return false
}

// SetReadIops gets a reference to the given float64 and assigns it to the ReadIops field.
func (o *OSSearchGatewayStat) SetReadIops(v float64) {
	o.ReadIops = &v
}

// GetReadLatencyUs returns the ReadLatencyUs field value if set, zero value otherwise.
func (o *OSSearchGatewayStat) GetReadLatencyUs() float64 {
	if o == nil || IsNil(o.ReadLatencyUs) {
		var ret float64
		return ret
	}
	return *o.ReadLatencyUs
}

// GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayStat) GetReadLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReadLatencyUs) {
		return nil, false
	}
	return o.ReadLatencyUs, true
}

// HasReadLatencyUs returns a boolean if a field has been set.
func (o *OSSearchGatewayStat) HasReadLatencyUs() bool {
	if o != nil && !IsNil(o.ReadLatencyUs) {
		return true
	}

	return false
}

// SetReadLatencyUs gets a reference to the given float64 and assigns it to the ReadLatencyUs field.
func (o *OSSearchGatewayStat) SetReadLatencyUs(v float64) {
	o.ReadLatencyUs = &v
}

// GetWriteBandwidthKbyte returns the WriteBandwidthKbyte field value if set, zero value otherwise.
func (o *OSSearchGatewayStat) GetWriteBandwidthKbyte() float64 {
	if o == nil || IsNil(o.WriteBandwidthKbyte) {
		var ret float64
		return ret
	}
	return *o.WriteBandwidthKbyte
}

// GetWriteBandwidthKbyteOk returns a tuple with the WriteBandwidthKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayStat) GetWriteBandwidthKbyteOk() (*float64, bool) {
	if o == nil || IsNil(o.WriteBandwidthKbyte) {
		return nil, false
	}
	return o.WriteBandwidthKbyte, true
}

// HasWriteBandwidthKbyte returns a boolean if a field has been set.
func (o *OSSearchGatewayStat) HasWriteBandwidthKbyte() bool {
	if o != nil && !IsNil(o.WriteBandwidthKbyte) {
		return true
	}

	return false
}

// SetWriteBandwidthKbyte gets a reference to the given float64 and assigns it to the WriteBandwidthKbyte field.
func (o *OSSearchGatewayStat) SetWriteBandwidthKbyte(v float64) {
	o.WriteBandwidthKbyte = &v
}

// GetWriteIops returns the WriteIops field value if set, zero value otherwise.
func (o *OSSearchGatewayStat) GetWriteIops() float64 {
	if o == nil || IsNil(o.WriteIops) {
		var ret float64
		return ret
	}
	return *o.WriteIops
}

// GetWriteIopsOk returns a tuple with the WriteIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayStat) GetWriteIopsOk() (*float64, bool) {
	if o == nil || IsNil(o.WriteIops) {
		return nil, false
	}
	return o.WriteIops, true
}

// HasWriteIops returns a boolean if a field has been set.
func (o *OSSearchGatewayStat) HasWriteIops() bool {
	if o != nil && !IsNil(o.WriteIops) {
		return true
	}

	return false
}

// SetWriteIops gets a reference to the given float64 and assigns it to the WriteIops field.
func (o *OSSearchGatewayStat) SetWriteIops(v float64) {
	o.WriteIops = &v
}

// GetWriteLatencyUs returns the WriteLatencyUs field value if set, zero value otherwise.
func (o *OSSearchGatewayStat) GetWriteLatencyUs() float64 {
	if o == nil || IsNil(o.WriteLatencyUs) {
		var ret float64
		return ret
	}
	return *o.WriteLatencyUs
}

// GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSSearchGatewayStat) GetWriteLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.WriteLatencyUs) {
		return nil, false
	}
	return o.WriteLatencyUs, true
}

// HasWriteLatencyUs returns a boolean if a field has been set.
func (o *OSSearchGatewayStat) HasWriteLatencyUs() bool {
	if o != nil && !IsNil(o.WriteLatencyUs) {
		return true
	}

	return false
}

// SetWriteLatencyUs gets a reference to the given float64 and assigns it to the WriteLatencyUs field.
func (o *OSSearchGatewayStat) SetWriteLatencyUs(v float64) {
	o.WriteLatencyUs = &v
}

func (o OSSearchGatewayStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSSearchGatewayStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpuUtil) {
		toSerialize["cpu_util"] = o.CpuUtil
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.MemUsagePercent) {
		toSerialize["mem_usage_percent"] = o.MemUsagePercent
	}
	if !IsNil(o.ReadBandwidthKbyte) {
		toSerialize["read_bandwidth_kbyte"] = o.ReadBandwidthKbyte
	}
	if !IsNil(o.ReadIops) {
		toSerialize["read_iops"] = o.ReadIops
	}
	if !IsNil(o.ReadLatencyUs) {
		toSerialize["read_latency_us"] = o.ReadLatencyUs
	}
	if !IsNil(o.WriteBandwidthKbyte) {
		toSerialize["write_bandwidth_kbyte"] = o.WriteBandwidthKbyte
	}
	if !IsNil(o.WriteIops) {
		toSerialize["write_iops"] = o.WriteIops
	}
	if !IsNil(o.WriteLatencyUs) {
		toSerialize["write_latency_us"] = o.WriteLatencyUs
	}
	return toSerialize, nil
}

type NullableOSSearchGatewayStat struct {
	value *OSSearchGatewayStat
	isSet bool
}

func (v NullableOSSearchGatewayStat) Get() *OSSearchGatewayStat {
	return v.value
}

func (v *NullableOSSearchGatewayStat) Set(val *OSSearchGatewayStat) {
	v.value = val
	v.isSet = true
}

func (v NullableOSSearchGatewayStat) IsSet() bool {
	return v.isSet
}

func (v *NullableOSSearchGatewayStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSSearchGatewayStat(val *OSSearchGatewayStat) *NullableOSSearchGatewayStat {
	return &NullableOSSearchGatewayStat{value: val, isSet: true}
}

func (v NullableOSSearchGatewayStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSSearchGatewayStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


