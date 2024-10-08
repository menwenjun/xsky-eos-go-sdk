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

// checks if the MetadataServiceStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataServiceStat{}

// MetadataServiceStat MetadataServiceStat defines the basic info of a metadata service
type MetadataServiceStat struct {
	CpuUtil *float64 `json:"cpu_util,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DeleteLatencyUs *float64 `json:"delete_latency_us,omitempty"`
	DeleteOps *float64 `json:"delete_ops,omitempty"`
	Disk *DiskStat `json:"disk,omitempty"`
	GetattrsLatencyUs *float64 `json:"getattrs_latency_us,omitempty"`
	ListLatencyUs *float64 `json:"list_latency_us,omitempty"`
	ListOps *float64 `json:"list_ops,omitempty"`
	MemUsagePercent *float64 `json:"mem_usage_percent,omitempty"`
	OpLatencyUs *float64 `json:"op_latency_us,omitempty"`
	OpenLatencyUs *float64 `json:"open_latency_us,omitempty"`
	ReadLatencyUs *float64 `json:"read_latency_us,omitempty"`
	ReadOps *float64 `json:"read_ops,omitempty"`
	RecoveryDone *int64 `json:"recovery_done,omitempty"`
	RecoveryLeftSecond *float64 `json:"recovery_left_second,omitempty"`
	RecoveryOps *float64 `json:"recovery_ops,omitempty"`
	// those fields are only for primary xmds
	RecoveryTotal *int64 `json:"recovery_total,omitempty"`
	WriteLatencyUs *float64 `json:"write_latency_us,omitempty"`
	WriteOps *float64 `json:"write_ops,omitempty"`
}

// NewMetadataServiceStat instantiates a new MetadataServiceStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataServiceStat() *MetadataServiceStat {
	this := MetadataServiceStat{}
	return &this
}

// NewMetadataServiceStatWithDefaults instantiates a new MetadataServiceStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataServiceStatWithDefaults() *MetadataServiceStat {
	this := MetadataServiceStat{}
	return &this
}

// GetCpuUtil returns the CpuUtil field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetCpuUtil() float64 {
	if o == nil || IsNil(o.CpuUtil) {
		var ret float64
		return ret
	}
	return *o.CpuUtil
}

// GetCpuUtilOk returns a tuple with the CpuUtil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetCpuUtilOk() (*float64, bool) {
	if o == nil || IsNil(o.CpuUtil) {
		return nil, false
	}
	return o.CpuUtil, true
}

// HasCpuUtil returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasCpuUtil() bool {
	if o != nil && !IsNil(o.CpuUtil) {
		return true
	}

	return false
}

// SetCpuUtil gets a reference to the given float64 and assigns it to the CpuUtil field.
func (o *MetadataServiceStat) SetCpuUtil(v float64) {
	o.CpuUtil = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *MetadataServiceStat) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDeleteLatencyUs returns the DeleteLatencyUs field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetDeleteLatencyUs() float64 {
	if o == nil || IsNil(o.DeleteLatencyUs) {
		var ret float64
		return ret
	}
	return *o.DeleteLatencyUs
}

// GetDeleteLatencyUsOk returns a tuple with the DeleteLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetDeleteLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.DeleteLatencyUs) {
		return nil, false
	}
	return o.DeleteLatencyUs, true
}

// HasDeleteLatencyUs returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasDeleteLatencyUs() bool {
	if o != nil && !IsNil(o.DeleteLatencyUs) {
		return true
	}

	return false
}

// SetDeleteLatencyUs gets a reference to the given float64 and assigns it to the DeleteLatencyUs field.
func (o *MetadataServiceStat) SetDeleteLatencyUs(v float64) {
	o.DeleteLatencyUs = &v
}

// GetDeleteOps returns the DeleteOps field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetDeleteOps() float64 {
	if o == nil || IsNil(o.DeleteOps) {
		var ret float64
		return ret
	}
	return *o.DeleteOps
}

// GetDeleteOpsOk returns a tuple with the DeleteOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetDeleteOpsOk() (*float64, bool) {
	if o == nil || IsNil(o.DeleteOps) {
		return nil, false
	}
	return o.DeleteOps, true
}

// HasDeleteOps returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasDeleteOps() bool {
	if o != nil && !IsNil(o.DeleteOps) {
		return true
	}

	return false
}

// SetDeleteOps gets a reference to the given float64 and assigns it to the DeleteOps field.
func (o *MetadataServiceStat) SetDeleteOps(v float64) {
	o.DeleteOps = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetDisk() DiskStat {
	if o == nil || IsNil(o.Disk) {
		var ret DiskStat
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetDiskOk() (*DiskStat, bool) {
	if o == nil || IsNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasDisk() bool {
	if o != nil && !IsNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given DiskStat and assigns it to the Disk field.
func (o *MetadataServiceStat) SetDisk(v DiskStat) {
	o.Disk = &v
}

// GetGetattrsLatencyUs returns the GetattrsLatencyUs field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetGetattrsLatencyUs() float64 {
	if o == nil || IsNil(o.GetattrsLatencyUs) {
		var ret float64
		return ret
	}
	return *o.GetattrsLatencyUs
}

// GetGetattrsLatencyUsOk returns a tuple with the GetattrsLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetGetattrsLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.GetattrsLatencyUs) {
		return nil, false
	}
	return o.GetattrsLatencyUs, true
}

// HasGetattrsLatencyUs returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasGetattrsLatencyUs() bool {
	if o != nil && !IsNil(o.GetattrsLatencyUs) {
		return true
	}

	return false
}

// SetGetattrsLatencyUs gets a reference to the given float64 and assigns it to the GetattrsLatencyUs field.
func (o *MetadataServiceStat) SetGetattrsLatencyUs(v float64) {
	o.GetattrsLatencyUs = &v
}

// GetListLatencyUs returns the ListLatencyUs field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetListLatencyUs() float64 {
	if o == nil || IsNil(o.ListLatencyUs) {
		var ret float64
		return ret
	}
	return *o.ListLatencyUs
}

// GetListLatencyUsOk returns a tuple with the ListLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetListLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.ListLatencyUs) {
		return nil, false
	}
	return o.ListLatencyUs, true
}

// HasListLatencyUs returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasListLatencyUs() bool {
	if o != nil && !IsNil(o.ListLatencyUs) {
		return true
	}

	return false
}

// SetListLatencyUs gets a reference to the given float64 and assigns it to the ListLatencyUs field.
func (o *MetadataServiceStat) SetListLatencyUs(v float64) {
	o.ListLatencyUs = &v
}

// GetListOps returns the ListOps field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetListOps() float64 {
	if o == nil || IsNil(o.ListOps) {
		var ret float64
		return ret
	}
	return *o.ListOps
}

// GetListOpsOk returns a tuple with the ListOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetListOpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ListOps) {
		return nil, false
	}
	return o.ListOps, true
}

// HasListOps returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasListOps() bool {
	if o != nil && !IsNil(o.ListOps) {
		return true
	}

	return false
}

// SetListOps gets a reference to the given float64 and assigns it to the ListOps field.
func (o *MetadataServiceStat) SetListOps(v float64) {
	o.ListOps = &v
}

// GetMemUsagePercent returns the MemUsagePercent field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetMemUsagePercent() float64 {
	if o == nil || IsNil(o.MemUsagePercent) {
		var ret float64
		return ret
	}
	return *o.MemUsagePercent
}

// GetMemUsagePercentOk returns a tuple with the MemUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetMemUsagePercentOk() (*float64, bool) {
	if o == nil || IsNil(o.MemUsagePercent) {
		return nil, false
	}
	return o.MemUsagePercent, true
}

// HasMemUsagePercent returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasMemUsagePercent() bool {
	if o != nil && !IsNil(o.MemUsagePercent) {
		return true
	}

	return false
}

// SetMemUsagePercent gets a reference to the given float64 and assigns it to the MemUsagePercent field.
func (o *MetadataServiceStat) SetMemUsagePercent(v float64) {
	o.MemUsagePercent = &v
}

// GetOpLatencyUs returns the OpLatencyUs field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetOpLatencyUs() float64 {
	if o == nil || IsNil(o.OpLatencyUs) {
		var ret float64
		return ret
	}
	return *o.OpLatencyUs
}

// GetOpLatencyUsOk returns a tuple with the OpLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetOpLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.OpLatencyUs) {
		return nil, false
	}
	return o.OpLatencyUs, true
}

// HasOpLatencyUs returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasOpLatencyUs() bool {
	if o != nil && !IsNil(o.OpLatencyUs) {
		return true
	}

	return false
}

// SetOpLatencyUs gets a reference to the given float64 and assigns it to the OpLatencyUs field.
func (o *MetadataServiceStat) SetOpLatencyUs(v float64) {
	o.OpLatencyUs = &v
}

// GetOpenLatencyUs returns the OpenLatencyUs field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetOpenLatencyUs() float64 {
	if o == nil || IsNil(o.OpenLatencyUs) {
		var ret float64
		return ret
	}
	return *o.OpenLatencyUs
}

// GetOpenLatencyUsOk returns a tuple with the OpenLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetOpenLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.OpenLatencyUs) {
		return nil, false
	}
	return o.OpenLatencyUs, true
}

// HasOpenLatencyUs returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasOpenLatencyUs() bool {
	if o != nil && !IsNil(o.OpenLatencyUs) {
		return true
	}

	return false
}

// SetOpenLatencyUs gets a reference to the given float64 and assigns it to the OpenLatencyUs field.
func (o *MetadataServiceStat) SetOpenLatencyUs(v float64) {
	o.OpenLatencyUs = &v
}

// GetReadLatencyUs returns the ReadLatencyUs field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetReadLatencyUs() float64 {
	if o == nil || IsNil(o.ReadLatencyUs) {
		var ret float64
		return ret
	}
	return *o.ReadLatencyUs
}

// GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetReadLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReadLatencyUs) {
		return nil, false
	}
	return o.ReadLatencyUs, true
}

// HasReadLatencyUs returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasReadLatencyUs() bool {
	if o != nil && !IsNil(o.ReadLatencyUs) {
		return true
	}

	return false
}

// SetReadLatencyUs gets a reference to the given float64 and assigns it to the ReadLatencyUs field.
func (o *MetadataServiceStat) SetReadLatencyUs(v float64) {
	o.ReadLatencyUs = &v
}

// GetReadOps returns the ReadOps field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetReadOps() float64 {
	if o == nil || IsNil(o.ReadOps) {
		var ret float64
		return ret
	}
	return *o.ReadOps
}

// GetReadOpsOk returns a tuple with the ReadOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetReadOpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReadOps) {
		return nil, false
	}
	return o.ReadOps, true
}

// HasReadOps returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasReadOps() bool {
	if o != nil && !IsNil(o.ReadOps) {
		return true
	}

	return false
}

// SetReadOps gets a reference to the given float64 and assigns it to the ReadOps field.
func (o *MetadataServiceStat) SetReadOps(v float64) {
	o.ReadOps = &v
}

// GetRecoveryDone returns the RecoveryDone field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetRecoveryDone() int64 {
	if o == nil || IsNil(o.RecoveryDone) {
		var ret int64
		return ret
	}
	return *o.RecoveryDone
}

// GetRecoveryDoneOk returns a tuple with the RecoveryDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetRecoveryDoneOk() (*int64, bool) {
	if o == nil || IsNil(o.RecoveryDone) {
		return nil, false
	}
	return o.RecoveryDone, true
}

// HasRecoveryDone returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasRecoveryDone() bool {
	if o != nil && !IsNil(o.RecoveryDone) {
		return true
	}

	return false
}

// SetRecoveryDone gets a reference to the given int64 and assigns it to the RecoveryDone field.
func (o *MetadataServiceStat) SetRecoveryDone(v int64) {
	o.RecoveryDone = &v
}

// GetRecoveryLeftSecond returns the RecoveryLeftSecond field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetRecoveryLeftSecond() float64 {
	if o == nil || IsNil(o.RecoveryLeftSecond) {
		var ret float64
		return ret
	}
	return *o.RecoveryLeftSecond
}

// GetRecoveryLeftSecondOk returns a tuple with the RecoveryLeftSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetRecoveryLeftSecondOk() (*float64, bool) {
	if o == nil || IsNil(o.RecoveryLeftSecond) {
		return nil, false
	}
	return o.RecoveryLeftSecond, true
}

// HasRecoveryLeftSecond returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasRecoveryLeftSecond() bool {
	if o != nil && !IsNil(o.RecoveryLeftSecond) {
		return true
	}

	return false
}

// SetRecoveryLeftSecond gets a reference to the given float64 and assigns it to the RecoveryLeftSecond field.
func (o *MetadataServiceStat) SetRecoveryLeftSecond(v float64) {
	o.RecoveryLeftSecond = &v
}

// GetRecoveryOps returns the RecoveryOps field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetRecoveryOps() float64 {
	if o == nil || IsNil(o.RecoveryOps) {
		var ret float64
		return ret
	}
	return *o.RecoveryOps
}

// GetRecoveryOpsOk returns a tuple with the RecoveryOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetRecoveryOpsOk() (*float64, bool) {
	if o == nil || IsNil(o.RecoveryOps) {
		return nil, false
	}
	return o.RecoveryOps, true
}

// HasRecoveryOps returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasRecoveryOps() bool {
	if o != nil && !IsNil(o.RecoveryOps) {
		return true
	}

	return false
}

// SetRecoveryOps gets a reference to the given float64 and assigns it to the RecoveryOps field.
func (o *MetadataServiceStat) SetRecoveryOps(v float64) {
	o.RecoveryOps = &v
}

// GetRecoveryTotal returns the RecoveryTotal field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetRecoveryTotal() int64 {
	if o == nil || IsNil(o.RecoveryTotal) {
		var ret int64
		return ret
	}
	return *o.RecoveryTotal
}

// GetRecoveryTotalOk returns a tuple with the RecoveryTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetRecoveryTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.RecoveryTotal) {
		return nil, false
	}
	return o.RecoveryTotal, true
}

// HasRecoveryTotal returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasRecoveryTotal() bool {
	if o != nil && !IsNil(o.RecoveryTotal) {
		return true
	}

	return false
}

// SetRecoveryTotal gets a reference to the given int64 and assigns it to the RecoveryTotal field.
func (o *MetadataServiceStat) SetRecoveryTotal(v int64) {
	o.RecoveryTotal = &v
}

// GetWriteLatencyUs returns the WriteLatencyUs field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetWriteLatencyUs() float64 {
	if o == nil || IsNil(o.WriteLatencyUs) {
		var ret float64
		return ret
	}
	return *o.WriteLatencyUs
}

// GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetWriteLatencyUsOk() (*float64, bool) {
	if o == nil || IsNil(o.WriteLatencyUs) {
		return nil, false
	}
	return o.WriteLatencyUs, true
}

// HasWriteLatencyUs returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasWriteLatencyUs() bool {
	if o != nil && !IsNil(o.WriteLatencyUs) {
		return true
	}

	return false
}

// SetWriteLatencyUs gets a reference to the given float64 and assigns it to the WriteLatencyUs field.
func (o *MetadataServiceStat) SetWriteLatencyUs(v float64) {
	o.WriteLatencyUs = &v
}

// GetWriteOps returns the WriteOps field value if set, zero value otherwise.
func (o *MetadataServiceStat) GetWriteOps() float64 {
	if o == nil || IsNil(o.WriteOps) {
		var ret float64
		return ret
	}
	return *o.WriteOps
}

// GetWriteOpsOk returns a tuple with the WriteOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataServiceStat) GetWriteOpsOk() (*float64, bool) {
	if o == nil || IsNil(o.WriteOps) {
		return nil, false
	}
	return o.WriteOps, true
}

// HasWriteOps returns a boolean if a field has been set.
func (o *MetadataServiceStat) HasWriteOps() bool {
	if o != nil && !IsNil(o.WriteOps) {
		return true
	}

	return false
}

// SetWriteOps gets a reference to the given float64 and assigns it to the WriteOps field.
func (o *MetadataServiceStat) SetWriteOps(v float64) {
	o.WriteOps = &v
}

func (o MetadataServiceStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataServiceStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpuUtil) {
		toSerialize["cpu_util"] = o.CpuUtil
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DeleteLatencyUs) {
		toSerialize["delete_latency_us"] = o.DeleteLatencyUs
	}
	if !IsNil(o.DeleteOps) {
		toSerialize["delete_ops"] = o.DeleteOps
	}
	if !IsNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !IsNil(o.GetattrsLatencyUs) {
		toSerialize["getattrs_latency_us"] = o.GetattrsLatencyUs
	}
	if !IsNil(o.ListLatencyUs) {
		toSerialize["list_latency_us"] = o.ListLatencyUs
	}
	if !IsNil(o.ListOps) {
		toSerialize["list_ops"] = o.ListOps
	}
	if !IsNil(o.MemUsagePercent) {
		toSerialize["mem_usage_percent"] = o.MemUsagePercent
	}
	if !IsNil(o.OpLatencyUs) {
		toSerialize["op_latency_us"] = o.OpLatencyUs
	}
	if !IsNil(o.OpenLatencyUs) {
		toSerialize["open_latency_us"] = o.OpenLatencyUs
	}
	if !IsNil(o.ReadLatencyUs) {
		toSerialize["read_latency_us"] = o.ReadLatencyUs
	}
	if !IsNil(o.ReadOps) {
		toSerialize["read_ops"] = o.ReadOps
	}
	if !IsNil(o.RecoveryDone) {
		toSerialize["recovery_done"] = o.RecoveryDone
	}
	if !IsNil(o.RecoveryLeftSecond) {
		toSerialize["recovery_left_second"] = o.RecoveryLeftSecond
	}
	if !IsNil(o.RecoveryOps) {
		toSerialize["recovery_ops"] = o.RecoveryOps
	}
	if !IsNil(o.RecoveryTotal) {
		toSerialize["recovery_total"] = o.RecoveryTotal
	}
	if !IsNil(o.WriteLatencyUs) {
		toSerialize["write_latency_us"] = o.WriteLatencyUs
	}
	if !IsNil(o.WriteOps) {
		toSerialize["write_ops"] = o.WriteOps
	}
	return toSerialize, nil
}

type NullableMetadataServiceStat struct {
	value *MetadataServiceStat
	isSet bool
}

func (v NullableMetadataServiceStat) Get() *MetadataServiceStat {
	return v.value
}

func (v *NullableMetadataServiceStat) Set(val *MetadataServiceStat) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataServiceStat) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataServiceStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataServiceStat(val *MetadataServiceStat) *NullableMetadataServiceStat {
	return &NullableMetadataServiceStat{value: val, isSet: true}
}

func (v NullableMetadataServiceStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataServiceStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


