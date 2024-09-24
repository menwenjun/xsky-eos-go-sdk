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

// checks if the DfsPathPerformanceStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsPathPerformanceStat{}

// DfsPathPerformanceStat DfsPathPerformanceStat records dfs path performance stat info
type DfsPathPerformanceStat struct {
	Create *time.Time `json:"create,omitempty"`
	DataCacheHitRate *float64 `json:"data_cache_hit_rate,omitempty"`
	DataDeleteBandwidthKbyte *int64 `json:"data_delete_bandwidth_kbyte,omitempty"`
	DataDeleteIops *int64 `json:"data_delete_iops,omitempty"`
	DataDeleteLatencyUs *int64 `json:"data_delete_latency_us,omitempty"`
	DataReadBandwidthKbyte *int64 `json:"data_read_bandwidth_kbyte,omitempty"`
	DataReadIops *int64 `json:"data_read_iops,omitempty"`
	DataReadLatencyUs *int64 `json:"data_read_latency_us,omitempty"`
	DataWriteBandwidthKbyte *int64 `json:"data_write_bandwidth_kbyte,omitempty"`
	DataWriteIops *int64 `json:"data_write_iops,omitempty"`
	DataWriteLatencyUs *int64 `json:"data_write_latency_us,omitempty"`
	MetaCacheHitRate *float64 `json:"meta_cache_hit_rate,omitempty"`
	MetaDeleteLatencyUs *int64 `json:"meta_delete_latency_us,omitempty"`
	MetaDeleteOps *int64 `json:"meta_delete_ops,omitempty"`
	MetaListLatencyUs *int64 `json:"meta_list_latency_us,omitempty"`
	MetaListOps *int64 `json:"meta_list_ops,omitempty"`
	MetaReadLatencyUs *int64 `json:"meta_read_latency_us,omitempty"`
	MetaReadOps *int64 `json:"meta_read_ops,omitempty"`
	MetaWriteLatencyUs *int64 `json:"meta_write_latency_us,omitempty"`
	MetaWriteOps *int64 `json:"meta_write_ops,omitempty"`
	ReadAheadBandwidthKbyte *int64 `json:"read_ahead_bandwidth_kbyte,omitempty"`
	ReadAheadIops *int64 `json:"read_ahead_iops,omitempty"`
}

// NewDfsPathPerformanceStat instantiates a new DfsPathPerformanceStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsPathPerformanceStat() *DfsPathPerformanceStat {
	this := DfsPathPerformanceStat{}
	return &this
}

// NewDfsPathPerformanceStatWithDefaults instantiates a new DfsPathPerformanceStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsPathPerformanceStatWithDefaults() *DfsPathPerformanceStat {
	this := DfsPathPerformanceStat{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsPathPerformanceStat) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDataCacheHitRate returns the DataCacheHitRate field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetDataCacheHitRate() float64 {
	if o == nil || IsNil(o.DataCacheHitRate) {
		var ret float64
		return ret
	}
	return *o.DataCacheHitRate
}

// GetDataCacheHitRateOk returns a tuple with the DataCacheHitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetDataCacheHitRateOk() (*float64, bool) {
	if o == nil || IsNil(o.DataCacheHitRate) {
		return nil, false
	}
	return o.DataCacheHitRate, true
}

// HasDataCacheHitRate returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasDataCacheHitRate() bool {
	if o != nil && !IsNil(o.DataCacheHitRate) {
		return true
	}

	return false
}

// SetDataCacheHitRate gets a reference to the given float64 and assigns it to the DataCacheHitRate field.
func (o *DfsPathPerformanceStat) SetDataCacheHitRate(v float64) {
	o.DataCacheHitRate = &v
}

// GetDataDeleteBandwidthKbyte returns the DataDeleteBandwidthKbyte field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetDataDeleteBandwidthKbyte() int64 {
	if o == nil || IsNil(o.DataDeleteBandwidthKbyte) {
		var ret int64
		return ret
	}
	return *o.DataDeleteBandwidthKbyte
}

// GetDataDeleteBandwidthKbyteOk returns a tuple with the DataDeleteBandwidthKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetDataDeleteBandwidthKbyteOk() (*int64, bool) {
	if o == nil || IsNil(o.DataDeleteBandwidthKbyte) {
		return nil, false
	}
	return o.DataDeleteBandwidthKbyte, true
}

// HasDataDeleteBandwidthKbyte returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasDataDeleteBandwidthKbyte() bool {
	if o != nil && !IsNil(o.DataDeleteBandwidthKbyte) {
		return true
	}

	return false
}

// SetDataDeleteBandwidthKbyte gets a reference to the given int64 and assigns it to the DataDeleteBandwidthKbyte field.
func (o *DfsPathPerformanceStat) SetDataDeleteBandwidthKbyte(v int64) {
	o.DataDeleteBandwidthKbyte = &v
}

// GetDataDeleteIops returns the DataDeleteIops field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetDataDeleteIops() int64 {
	if o == nil || IsNil(o.DataDeleteIops) {
		var ret int64
		return ret
	}
	return *o.DataDeleteIops
}

// GetDataDeleteIopsOk returns a tuple with the DataDeleteIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetDataDeleteIopsOk() (*int64, bool) {
	if o == nil || IsNil(o.DataDeleteIops) {
		return nil, false
	}
	return o.DataDeleteIops, true
}

// HasDataDeleteIops returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasDataDeleteIops() bool {
	if o != nil && !IsNil(o.DataDeleteIops) {
		return true
	}

	return false
}

// SetDataDeleteIops gets a reference to the given int64 and assigns it to the DataDeleteIops field.
func (o *DfsPathPerformanceStat) SetDataDeleteIops(v int64) {
	o.DataDeleteIops = &v
}

// GetDataDeleteLatencyUs returns the DataDeleteLatencyUs field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetDataDeleteLatencyUs() int64 {
	if o == nil || IsNil(o.DataDeleteLatencyUs) {
		var ret int64
		return ret
	}
	return *o.DataDeleteLatencyUs
}

// GetDataDeleteLatencyUsOk returns a tuple with the DataDeleteLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetDataDeleteLatencyUsOk() (*int64, bool) {
	if o == nil || IsNil(o.DataDeleteLatencyUs) {
		return nil, false
	}
	return o.DataDeleteLatencyUs, true
}

// HasDataDeleteLatencyUs returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasDataDeleteLatencyUs() bool {
	if o != nil && !IsNil(o.DataDeleteLatencyUs) {
		return true
	}

	return false
}

// SetDataDeleteLatencyUs gets a reference to the given int64 and assigns it to the DataDeleteLatencyUs field.
func (o *DfsPathPerformanceStat) SetDataDeleteLatencyUs(v int64) {
	o.DataDeleteLatencyUs = &v
}

// GetDataReadBandwidthKbyte returns the DataReadBandwidthKbyte field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetDataReadBandwidthKbyte() int64 {
	if o == nil || IsNil(o.DataReadBandwidthKbyte) {
		var ret int64
		return ret
	}
	return *o.DataReadBandwidthKbyte
}

// GetDataReadBandwidthKbyteOk returns a tuple with the DataReadBandwidthKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetDataReadBandwidthKbyteOk() (*int64, bool) {
	if o == nil || IsNil(o.DataReadBandwidthKbyte) {
		return nil, false
	}
	return o.DataReadBandwidthKbyte, true
}

// HasDataReadBandwidthKbyte returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasDataReadBandwidthKbyte() bool {
	if o != nil && !IsNil(o.DataReadBandwidthKbyte) {
		return true
	}

	return false
}

// SetDataReadBandwidthKbyte gets a reference to the given int64 and assigns it to the DataReadBandwidthKbyte field.
func (o *DfsPathPerformanceStat) SetDataReadBandwidthKbyte(v int64) {
	o.DataReadBandwidthKbyte = &v
}

// GetDataReadIops returns the DataReadIops field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetDataReadIops() int64 {
	if o == nil || IsNil(o.DataReadIops) {
		var ret int64
		return ret
	}
	return *o.DataReadIops
}

// GetDataReadIopsOk returns a tuple with the DataReadIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetDataReadIopsOk() (*int64, bool) {
	if o == nil || IsNil(o.DataReadIops) {
		return nil, false
	}
	return o.DataReadIops, true
}

// HasDataReadIops returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasDataReadIops() bool {
	if o != nil && !IsNil(o.DataReadIops) {
		return true
	}

	return false
}

// SetDataReadIops gets a reference to the given int64 and assigns it to the DataReadIops field.
func (o *DfsPathPerformanceStat) SetDataReadIops(v int64) {
	o.DataReadIops = &v
}

// GetDataReadLatencyUs returns the DataReadLatencyUs field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetDataReadLatencyUs() int64 {
	if o == nil || IsNil(o.DataReadLatencyUs) {
		var ret int64
		return ret
	}
	return *o.DataReadLatencyUs
}

// GetDataReadLatencyUsOk returns a tuple with the DataReadLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetDataReadLatencyUsOk() (*int64, bool) {
	if o == nil || IsNil(o.DataReadLatencyUs) {
		return nil, false
	}
	return o.DataReadLatencyUs, true
}

// HasDataReadLatencyUs returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasDataReadLatencyUs() bool {
	if o != nil && !IsNil(o.DataReadLatencyUs) {
		return true
	}

	return false
}

// SetDataReadLatencyUs gets a reference to the given int64 and assigns it to the DataReadLatencyUs field.
func (o *DfsPathPerformanceStat) SetDataReadLatencyUs(v int64) {
	o.DataReadLatencyUs = &v
}

// GetDataWriteBandwidthKbyte returns the DataWriteBandwidthKbyte field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetDataWriteBandwidthKbyte() int64 {
	if o == nil || IsNil(o.DataWriteBandwidthKbyte) {
		var ret int64
		return ret
	}
	return *o.DataWriteBandwidthKbyte
}

// GetDataWriteBandwidthKbyteOk returns a tuple with the DataWriteBandwidthKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetDataWriteBandwidthKbyteOk() (*int64, bool) {
	if o == nil || IsNil(o.DataWriteBandwidthKbyte) {
		return nil, false
	}
	return o.DataWriteBandwidthKbyte, true
}

// HasDataWriteBandwidthKbyte returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasDataWriteBandwidthKbyte() bool {
	if o != nil && !IsNil(o.DataWriteBandwidthKbyte) {
		return true
	}

	return false
}

// SetDataWriteBandwidthKbyte gets a reference to the given int64 and assigns it to the DataWriteBandwidthKbyte field.
func (o *DfsPathPerformanceStat) SetDataWriteBandwidthKbyte(v int64) {
	o.DataWriteBandwidthKbyte = &v
}

// GetDataWriteIops returns the DataWriteIops field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetDataWriteIops() int64 {
	if o == nil || IsNil(o.DataWriteIops) {
		var ret int64
		return ret
	}
	return *o.DataWriteIops
}

// GetDataWriteIopsOk returns a tuple with the DataWriteIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetDataWriteIopsOk() (*int64, bool) {
	if o == nil || IsNil(o.DataWriteIops) {
		return nil, false
	}
	return o.DataWriteIops, true
}

// HasDataWriteIops returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasDataWriteIops() bool {
	if o != nil && !IsNil(o.DataWriteIops) {
		return true
	}

	return false
}

// SetDataWriteIops gets a reference to the given int64 and assigns it to the DataWriteIops field.
func (o *DfsPathPerformanceStat) SetDataWriteIops(v int64) {
	o.DataWriteIops = &v
}

// GetDataWriteLatencyUs returns the DataWriteLatencyUs field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetDataWriteLatencyUs() int64 {
	if o == nil || IsNil(o.DataWriteLatencyUs) {
		var ret int64
		return ret
	}
	return *o.DataWriteLatencyUs
}

// GetDataWriteLatencyUsOk returns a tuple with the DataWriteLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetDataWriteLatencyUsOk() (*int64, bool) {
	if o == nil || IsNil(o.DataWriteLatencyUs) {
		return nil, false
	}
	return o.DataWriteLatencyUs, true
}

// HasDataWriteLatencyUs returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasDataWriteLatencyUs() bool {
	if o != nil && !IsNil(o.DataWriteLatencyUs) {
		return true
	}

	return false
}

// SetDataWriteLatencyUs gets a reference to the given int64 and assigns it to the DataWriteLatencyUs field.
func (o *DfsPathPerformanceStat) SetDataWriteLatencyUs(v int64) {
	o.DataWriteLatencyUs = &v
}

// GetMetaCacheHitRate returns the MetaCacheHitRate field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetMetaCacheHitRate() float64 {
	if o == nil || IsNil(o.MetaCacheHitRate) {
		var ret float64
		return ret
	}
	return *o.MetaCacheHitRate
}

// GetMetaCacheHitRateOk returns a tuple with the MetaCacheHitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetMetaCacheHitRateOk() (*float64, bool) {
	if o == nil || IsNil(o.MetaCacheHitRate) {
		return nil, false
	}
	return o.MetaCacheHitRate, true
}

// HasMetaCacheHitRate returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasMetaCacheHitRate() bool {
	if o != nil && !IsNil(o.MetaCacheHitRate) {
		return true
	}

	return false
}

// SetMetaCacheHitRate gets a reference to the given float64 and assigns it to the MetaCacheHitRate field.
func (o *DfsPathPerformanceStat) SetMetaCacheHitRate(v float64) {
	o.MetaCacheHitRate = &v
}

// GetMetaDeleteLatencyUs returns the MetaDeleteLatencyUs field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetMetaDeleteLatencyUs() int64 {
	if o == nil || IsNil(o.MetaDeleteLatencyUs) {
		var ret int64
		return ret
	}
	return *o.MetaDeleteLatencyUs
}

// GetMetaDeleteLatencyUsOk returns a tuple with the MetaDeleteLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetMetaDeleteLatencyUsOk() (*int64, bool) {
	if o == nil || IsNil(o.MetaDeleteLatencyUs) {
		return nil, false
	}
	return o.MetaDeleteLatencyUs, true
}

// HasMetaDeleteLatencyUs returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasMetaDeleteLatencyUs() bool {
	if o != nil && !IsNil(o.MetaDeleteLatencyUs) {
		return true
	}

	return false
}

// SetMetaDeleteLatencyUs gets a reference to the given int64 and assigns it to the MetaDeleteLatencyUs field.
func (o *DfsPathPerformanceStat) SetMetaDeleteLatencyUs(v int64) {
	o.MetaDeleteLatencyUs = &v
}

// GetMetaDeleteOps returns the MetaDeleteOps field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetMetaDeleteOps() int64 {
	if o == nil || IsNil(o.MetaDeleteOps) {
		var ret int64
		return ret
	}
	return *o.MetaDeleteOps
}

// GetMetaDeleteOpsOk returns a tuple with the MetaDeleteOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetMetaDeleteOpsOk() (*int64, bool) {
	if o == nil || IsNil(o.MetaDeleteOps) {
		return nil, false
	}
	return o.MetaDeleteOps, true
}

// HasMetaDeleteOps returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasMetaDeleteOps() bool {
	if o != nil && !IsNil(o.MetaDeleteOps) {
		return true
	}

	return false
}

// SetMetaDeleteOps gets a reference to the given int64 and assigns it to the MetaDeleteOps field.
func (o *DfsPathPerformanceStat) SetMetaDeleteOps(v int64) {
	o.MetaDeleteOps = &v
}

// GetMetaListLatencyUs returns the MetaListLatencyUs field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetMetaListLatencyUs() int64 {
	if o == nil || IsNil(o.MetaListLatencyUs) {
		var ret int64
		return ret
	}
	return *o.MetaListLatencyUs
}

// GetMetaListLatencyUsOk returns a tuple with the MetaListLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetMetaListLatencyUsOk() (*int64, bool) {
	if o == nil || IsNil(o.MetaListLatencyUs) {
		return nil, false
	}
	return o.MetaListLatencyUs, true
}

// HasMetaListLatencyUs returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasMetaListLatencyUs() bool {
	if o != nil && !IsNil(o.MetaListLatencyUs) {
		return true
	}

	return false
}

// SetMetaListLatencyUs gets a reference to the given int64 and assigns it to the MetaListLatencyUs field.
func (o *DfsPathPerformanceStat) SetMetaListLatencyUs(v int64) {
	o.MetaListLatencyUs = &v
}

// GetMetaListOps returns the MetaListOps field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetMetaListOps() int64 {
	if o == nil || IsNil(o.MetaListOps) {
		var ret int64
		return ret
	}
	return *o.MetaListOps
}

// GetMetaListOpsOk returns a tuple with the MetaListOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetMetaListOpsOk() (*int64, bool) {
	if o == nil || IsNil(o.MetaListOps) {
		return nil, false
	}
	return o.MetaListOps, true
}

// HasMetaListOps returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasMetaListOps() bool {
	if o != nil && !IsNil(o.MetaListOps) {
		return true
	}

	return false
}

// SetMetaListOps gets a reference to the given int64 and assigns it to the MetaListOps field.
func (o *DfsPathPerformanceStat) SetMetaListOps(v int64) {
	o.MetaListOps = &v
}

// GetMetaReadLatencyUs returns the MetaReadLatencyUs field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetMetaReadLatencyUs() int64 {
	if o == nil || IsNil(o.MetaReadLatencyUs) {
		var ret int64
		return ret
	}
	return *o.MetaReadLatencyUs
}

// GetMetaReadLatencyUsOk returns a tuple with the MetaReadLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetMetaReadLatencyUsOk() (*int64, bool) {
	if o == nil || IsNil(o.MetaReadLatencyUs) {
		return nil, false
	}
	return o.MetaReadLatencyUs, true
}

// HasMetaReadLatencyUs returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasMetaReadLatencyUs() bool {
	if o != nil && !IsNil(o.MetaReadLatencyUs) {
		return true
	}

	return false
}

// SetMetaReadLatencyUs gets a reference to the given int64 and assigns it to the MetaReadLatencyUs field.
func (o *DfsPathPerformanceStat) SetMetaReadLatencyUs(v int64) {
	o.MetaReadLatencyUs = &v
}

// GetMetaReadOps returns the MetaReadOps field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetMetaReadOps() int64 {
	if o == nil || IsNil(o.MetaReadOps) {
		var ret int64
		return ret
	}
	return *o.MetaReadOps
}

// GetMetaReadOpsOk returns a tuple with the MetaReadOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetMetaReadOpsOk() (*int64, bool) {
	if o == nil || IsNil(o.MetaReadOps) {
		return nil, false
	}
	return o.MetaReadOps, true
}

// HasMetaReadOps returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasMetaReadOps() bool {
	if o != nil && !IsNil(o.MetaReadOps) {
		return true
	}

	return false
}

// SetMetaReadOps gets a reference to the given int64 and assigns it to the MetaReadOps field.
func (o *DfsPathPerformanceStat) SetMetaReadOps(v int64) {
	o.MetaReadOps = &v
}

// GetMetaWriteLatencyUs returns the MetaWriteLatencyUs field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetMetaWriteLatencyUs() int64 {
	if o == nil || IsNil(o.MetaWriteLatencyUs) {
		var ret int64
		return ret
	}
	return *o.MetaWriteLatencyUs
}

// GetMetaWriteLatencyUsOk returns a tuple with the MetaWriteLatencyUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetMetaWriteLatencyUsOk() (*int64, bool) {
	if o == nil || IsNil(o.MetaWriteLatencyUs) {
		return nil, false
	}
	return o.MetaWriteLatencyUs, true
}

// HasMetaWriteLatencyUs returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasMetaWriteLatencyUs() bool {
	if o != nil && !IsNil(o.MetaWriteLatencyUs) {
		return true
	}

	return false
}

// SetMetaWriteLatencyUs gets a reference to the given int64 and assigns it to the MetaWriteLatencyUs field.
func (o *DfsPathPerformanceStat) SetMetaWriteLatencyUs(v int64) {
	o.MetaWriteLatencyUs = &v
}

// GetMetaWriteOps returns the MetaWriteOps field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetMetaWriteOps() int64 {
	if o == nil || IsNil(o.MetaWriteOps) {
		var ret int64
		return ret
	}
	return *o.MetaWriteOps
}

// GetMetaWriteOpsOk returns a tuple with the MetaWriteOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetMetaWriteOpsOk() (*int64, bool) {
	if o == nil || IsNil(o.MetaWriteOps) {
		return nil, false
	}
	return o.MetaWriteOps, true
}

// HasMetaWriteOps returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasMetaWriteOps() bool {
	if o != nil && !IsNil(o.MetaWriteOps) {
		return true
	}

	return false
}

// SetMetaWriteOps gets a reference to the given int64 and assigns it to the MetaWriteOps field.
func (o *DfsPathPerformanceStat) SetMetaWriteOps(v int64) {
	o.MetaWriteOps = &v
}

// GetReadAheadBandwidthKbyte returns the ReadAheadBandwidthKbyte field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetReadAheadBandwidthKbyte() int64 {
	if o == nil || IsNil(o.ReadAheadBandwidthKbyte) {
		var ret int64
		return ret
	}
	return *o.ReadAheadBandwidthKbyte
}

// GetReadAheadBandwidthKbyteOk returns a tuple with the ReadAheadBandwidthKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetReadAheadBandwidthKbyteOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadAheadBandwidthKbyte) {
		return nil, false
	}
	return o.ReadAheadBandwidthKbyte, true
}

// HasReadAheadBandwidthKbyte returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasReadAheadBandwidthKbyte() bool {
	if o != nil && !IsNil(o.ReadAheadBandwidthKbyte) {
		return true
	}

	return false
}

// SetReadAheadBandwidthKbyte gets a reference to the given int64 and assigns it to the ReadAheadBandwidthKbyte field.
func (o *DfsPathPerformanceStat) SetReadAheadBandwidthKbyte(v int64) {
	o.ReadAheadBandwidthKbyte = &v
}

// GetReadAheadIops returns the ReadAheadIops field value if set, zero value otherwise.
func (o *DfsPathPerformanceStat) GetReadAheadIops() int64 {
	if o == nil || IsNil(o.ReadAheadIops) {
		var ret int64
		return ret
	}
	return *o.ReadAheadIops
}

// GetReadAheadIopsOk returns a tuple with the ReadAheadIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathPerformanceStat) GetReadAheadIopsOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadAheadIops) {
		return nil, false
	}
	return o.ReadAheadIops, true
}

// HasReadAheadIops returns a boolean if a field has been set.
func (o *DfsPathPerformanceStat) HasReadAheadIops() bool {
	if o != nil && !IsNil(o.ReadAheadIops) {
		return true
	}

	return false
}

// SetReadAheadIops gets a reference to the given int64 and assigns it to the ReadAheadIops field.
func (o *DfsPathPerformanceStat) SetReadAheadIops(v int64) {
	o.ReadAheadIops = &v
}

func (o DfsPathPerformanceStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsPathPerformanceStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DataCacheHitRate) {
		toSerialize["data_cache_hit_rate"] = o.DataCacheHitRate
	}
	if !IsNil(o.DataDeleteBandwidthKbyte) {
		toSerialize["data_delete_bandwidth_kbyte"] = o.DataDeleteBandwidthKbyte
	}
	if !IsNil(o.DataDeleteIops) {
		toSerialize["data_delete_iops"] = o.DataDeleteIops
	}
	if !IsNil(o.DataDeleteLatencyUs) {
		toSerialize["data_delete_latency_us"] = o.DataDeleteLatencyUs
	}
	if !IsNil(o.DataReadBandwidthKbyte) {
		toSerialize["data_read_bandwidth_kbyte"] = o.DataReadBandwidthKbyte
	}
	if !IsNil(o.DataReadIops) {
		toSerialize["data_read_iops"] = o.DataReadIops
	}
	if !IsNil(o.DataReadLatencyUs) {
		toSerialize["data_read_latency_us"] = o.DataReadLatencyUs
	}
	if !IsNil(o.DataWriteBandwidthKbyte) {
		toSerialize["data_write_bandwidth_kbyte"] = o.DataWriteBandwidthKbyte
	}
	if !IsNil(o.DataWriteIops) {
		toSerialize["data_write_iops"] = o.DataWriteIops
	}
	if !IsNil(o.DataWriteLatencyUs) {
		toSerialize["data_write_latency_us"] = o.DataWriteLatencyUs
	}
	if !IsNil(o.MetaCacheHitRate) {
		toSerialize["meta_cache_hit_rate"] = o.MetaCacheHitRate
	}
	if !IsNil(o.MetaDeleteLatencyUs) {
		toSerialize["meta_delete_latency_us"] = o.MetaDeleteLatencyUs
	}
	if !IsNil(o.MetaDeleteOps) {
		toSerialize["meta_delete_ops"] = o.MetaDeleteOps
	}
	if !IsNil(o.MetaListLatencyUs) {
		toSerialize["meta_list_latency_us"] = o.MetaListLatencyUs
	}
	if !IsNil(o.MetaListOps) {
		toSerialize["meta_list_ops"] = o.MetaListOps
	}
	if !IsNil(o.MetaReadLatencyUs) {
		toSerialize["meta_read_latency_us"] = o.MetaReadLatencyUs
	}
	if !IsNil(o.MetaReadOps) {
		toSerialize["meta_read_ops"] = o.MetaReadOps
	}
	if !IsNil(o.MetaWriteLatencyUs) {
		toSerialize["meta_write_latency_us"] = o.MetaWriteLatencyUs
	}
	if !IsNil(o.MetaWriteOps) {
		toSerialize["meta_write_ops"] = o.MetaWriteOps
	}
	if !IsNil(o.ReadAheadBandwidthKbyte) {
		toSerialize["read_ahead_bandwidth_kbyte"] = o.ReadAheadBandwidthKbyte
	}
	if !IsNil(o.ReadAheadIops) {
		toSerialize["read_ahead_iops"] = o.ReadAheadIops
	}
	return toSerialize, nil
}

type NullableDfsPathPerformanceStat struct {
	value *DfsPathPerformanceStat
	isSet bool
}

func (v NullableDfsPathPerformanceStat) Get() *DfsPathPerformanceStat {
	return v.value
}

func (v *NullableDfsPathPerformanceStat) Set(val *DfsPathPerformanceStat) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsPathPerformanceStat) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsPathPerformanceStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsPathPerformanceStat(val *DfsPathPerformanceStat) *NullableDfsPathPerformanceStat {
	return &NullableDfsPathPerformanceStat{value: val, isSet: true}
}

func (v NullableDfsPathPerformanceStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsPathPerformanceStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


