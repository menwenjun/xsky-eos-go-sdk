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

// checks if the OspMetadataClusterStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspMetadataClusterStat{}

// OspMetadataClusterStat OspMetadataClusterStat reprents osp metadata cluster stat info
type OspMetadataClusterStat struct {
	Create *time.Time `json:"create,omitempty"`
	MovingDataSizeByte *int64 `json:"moving_data_size_byte,omitempty"`
	ReadLatencyUs *float64 `json:"read_latency_us,omitempty"`
	ReadOps *float64 `json:"read_ops,omitempty"`
	TotalByte *int64 `json:"total_byte,omitempty"`
	TotalKvSizeByte *int64 `json:"total_kv_size_byte,omitempty"`
	UsedByte *int64 `json:"used_byte,omitempty"`
	WriteLatencyUs *float64 `json:"write_latency_us,omitempty"`
	WriteOps *float64 `json:"write_ops,omitempty"`
}

// NewOspMetadataClusterStat instantiates a new OspMetadataClusterStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspMetadataClusterStat() *OspMetadataClusterStat {
	this := OspMetadataClusterStat{}
	return &this
}

// NewOspMetadataClusterStatWithDefaults instantiates a new OspMetadataClusterStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspMetadataClusterStatWithDefaults() *OspMetadataClusterStat {
	this := OspMetadataClusterStat{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OspMetadataClusterStat) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterStat) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OspMetadataClusterStat) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OspMetadataClusterStat) SetCreate(v time.Time) {
	o.Create = &v
}

// GetMovingDataSizeByte returns the MovingDataSizeByte field value if set, zero value otherwise.
func (o *OspMetadataClusterStat) GetMovingDataSizeByte() int64 {
	if o == nil || IsNil(o.MovingDataSizeByte) {
		var ret int64
		return ret
	}
	return *o.MovingDataSizeByte
}

// GetMovingDataSizeByteOk returns a tuple with the MovingDataSizeByte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterStat) GetMovingDataSizeByteOk() (*int64, bool) {
	if o == nil || IsNil(o.MovingDataSizeByte) {
		return nil, false
	}
	return o.MovingDataSizeByte, true
}

// HasMovingDataSizeByte returns a boolean if a field has been set.
func (o *OspMetadataClusterStat) HasMovingDataSizeByte() bool {
	if o != nil && !IsNil(o.MovingDataSizeByte) {
		return true
	}

	return false
}

// SetMovingDataSizeByte gets a reference to the given int64 and assigns it to the MovingDataSizeByte field.
func (o *OspMetadataClusterStat) SetMovingDataSizeByte(v int64) {
	o.MovingDataSizeByte = &v
}

// GetReadLatencyUs returns the ReadLatencyUs field value if set, zero value otherwise.
func (o *OspMetadataClusterStat) GetReadLatencyUs() float64 {
	if o == nil || IsNil(o.ReadLatencyUs) {
		var ret float64
		return ret
	}
	return *o.ReadLatencyUs
}

// GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterStat) GetReadLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReadLatencyUs) {
		return nil, false
	}
	return o.ReadLatencyUs, true
}

// HasReadLatencyUs returns a boolean if a field has been set.
func (o *OspMetadataClusterStat) HasReadLatencyUs() bool {
	if o != nil && !IsNil(o.ReadLatencyUs) {
		return true
	}

	return false
}

// SetReadLatencyUs gets a reference to the given float64 and assigns it to the ReadLatencyUs field.
func (o *OspMetadataClusterStat) SetReadLatencyUs(v float64) {
	o.ReadLatencyUs = &v
}

// GetReadOps returns the ReadOps field value if set, zero value otherwise.
func (o *OspMetadataClusterStat) GetReadOps() float64 {
	if o == nil || IsNil(o.ReadOps) {
		var ret float64
		return ret
	}
	return *o.ReadOps
}

// GetReadOpsOk returns a tuple with the ReadOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterStat) GetReadOpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReadOps) {
		return nil, false
	}
	return o.ReadOps, true
}

// HasReadOps returns a boolean if a field has been set.
func (o *OspMetadataClusterStat) HasReadOps() bool {
	if o != nil && !IsNil(o.ReadOps) {
		return true
	}

	return false
}

// SetReadOps gets a reference to the given float64 and assigns it to the ReadOps field.
func (o *OspMetadataClusterStat) SetReadOps(v float64) {
	o.ReadOps = &v
}

// GetTotalByte returns the TotalByte field value if set, zero value otherwise.
func (o *OspMetadataClusterStat) GetTotalByte() int64 {
	if o == nil || IsNil(o.TotalByte) {
		var ret int64
		return ret
	}
	return *o.TotalByte
}

// GetTotalByteOk returns a tuple with the TotalByte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterStat) GetTotalByteOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalByte) {
		return nil, false
	}
	return o.TotalByte, true
}

// HasTotalByte returns a boolean if a field has been set.
func (o *OspMetadataClusterStat) HasTotalByte() bool {
	if o != nil && !IsNil(o.TotalByte) {
		return true
	}

	return false
}

// SetTotalByte gets a reference to the given int64 and assigns it to the TotalByte field.
func (o *OspMetadataClusterStat) SetTotalByte(v int64) {
	o.TotalByte = &v
}

// GetTotalKvSizeByte returns the TotalKvSizeByte field value if set, zero value otherwise.
func (o *OspMetadataClusterStat) GetTotalKvSizeByte() int64 {
	if o == nil || IsNil(o.TotalKvSizeByte) {
		var ret int64
		return ret
	}
	return *o.TotalKvSizeByte
}

// GetTotalKvSizeByteOk returns a tuple with the TotalKvSizeByte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterStat) GetTotalKvSizeByteOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalKvSizeByte) {
		return nil, false
	}
	return o.TotalKvSizeByte, true
}

// HasTotalKvSizeByte returns a boolean if a field has been set.
func (o *OspMetadataClusterStat) HasTotalKvSizeByte() bool {
	if o != nil && !IsNil(o.TotalKvSizeByte) {
		return true
	}

	return false
}

// SetTotalKvSizeByte gets a reference to the given int64 and assigns it to the TotalKvSizeByte field.
func (o *OspMetadataClusterStat) SetTotalKvSizeByte(v int64) {
	o.TotalKvSizeByte = &v
}

// GetUsedByte returns the UsedByte field value if set, zero value otherwise.
func (o *OspMetadataClusterStat) GetUsedByte() int64 {
	if o == nil || IsNil(o.UsedByte) {
		var ret int64
		return ret
	}
	return *o.UsedByte
}

// GetUsedByteOk returns a tuple with the UsedByte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterStat) GetUsedByteOk() (*int64, bool) {
	if o == nil || IsNil(o.UsedByte) {
		return nil, false
	}
	return o.UsedByte, true
}

// HasUsedByte returns a boolean if a field has been set.
func (o *OspMetadataClusterStat) HasUsedByte() bool {
	if o != nil && !IsNil(o.UsedByte) {
		return true
	}

	return false
}

// SetUsedByte gets a reference to the given int64 and assigns it to the UsedByte field.
func (o *OspMetadataClusterStat) SetUsedByte(v int64) {
	o.UsedByte = &v
}

// GetWriteLatencyUs returns the WriteLatencyUs field value if set, zero value otherwise.
func (o *OspMetadataClusterStat) GetWriteLatencyUs() float64 {
	if o == nil || IsNil(o.WriteLatencyUs) {
		var ret float64
		return ret
	}
	return *o.WriteLatencyUs
}

// GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterStat) GetWriteLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.WriteLatencyUs) {
		return nil, false
	}
	return o.WriteLatencyUs, true
}

// HasWriteLatencyUs returns a boolean if a field has been set.
func (o *OspMetadataClusterStat) HasWriteLatencyUs() bool {
	if o != nil && !IsNil(o.WriteLatencyUs) {
		return true
	}

	return false
}

// SetWriteLatencyUs gets a reference to the given float64 and assigns it to the WriteLatencyUs field.
func (o *OspMetadataClusterStat) SetWriteLatencyUs(v float64) {
	o.WriteLatencyUs = &v
}

// GetWriteOps returns the WriteOps field value if set, zero value otherwise.
func (o *OspMetadataClusterStat) GetWriteOps() float64 {
	if o == nil || IsNil(o.WriteOps) {
		var ret float64
		return ret
	}
	return *o.WriteOps
}

// GetWriteOpsOk returns a tuple with the WriteOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataClusterStat) GetWriteOpsOk() (*float64, bool) {
	if o == nil || IsNil(o.WriteOps) {
		return nil, false
	}
	return o.WriteOps, true
}

// HasWriteOps returns a boolean if a field has been set.
func (o *OspMetadataClusterStat) HasWriteOps() bool {
	if o != nil && !IsNil(o.WriteOps) {
		return true
	}

	return false
}

// SetWriteOps gets a reference to the given float64 and assigns it to the WriteOps field.
func (o *OspMetadataClusterStat) SetWriteOps(v float64) {
	o.WriteOps = &v
}

func (o OspMetadataClusterStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspMetadataClusterStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.MovingDataSizeByte) {
		toSerialize["moving_data_size_byte"] = o.MovingDataSizeByte
	}
	if !IsNil(o.ReadLatencyUs) {
		toSerialize["read_latency_us"] = o.ReadLatencyUs
	}
	if !IsNil(o.ReadOps) {
		toSerialize["read_ops"] = o.ReadOps
	}
	if !IsNil(o.TotalByte) {
		toSerialize["total_byte"] = o.TotalByte
	}
	if !IsNil(o.TotalKvSizeByte) {
		toSerialize["total_kv_size_byte"] = o.TotalKvSizeByte
	}
	if !IsNil(o.UsedByte) {
		toSerialize["used_byte"] = o.UsedByte
	}
	if !IsNil(o.WriteLatencyUs) {
		toSerialize["write_latency_us"] = o.WriteLatencyUs
	}
	if !IsNil(o.WriteOps) {
		toSerialize["write_ops"] = o.WriteOps
	}
	return toSerialize, nil
}

type NullableOspMetadataClusterStat struct {
	value *OspMetadataClusterStat
	isSet bool
}

func (v NullableOspMetadataClusterStat) Get() *OspMetadataClusterStat {
	return v.value
}

func (v *NullableOspMetadataClusterStat) Set(val *OspMetadataClusterStat) {
	v.value = val
	v.isSet = true
}

func (v NullableOspMetadataClusterStat) IsSet() bool {
	return v.isSet
}

func (v *NullableOspMetadataClusterStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspMetadataClusterStat(val *OspMetadataClusterStat) *NullableOspMetadataClusterStat {
	return &NullableOspMetadataClusterStat{value: val, isSet: true}
}

func (v NullableOspMetadataClusterStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspMetadataClusterStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


