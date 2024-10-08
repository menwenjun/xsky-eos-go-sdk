/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OsdCreateReqOsd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsdCreateReqOsd{}

// OsdCreateReqOsd struct for OsdCreateReqOsd
type OsdCreateReqOsd struct {
	// bind pool id
	BindPoolId *int64 `json:"bind_pool_id,omitempty"`
	// data disk id
	DiskId *int64 `json:"disk_id,omitempty"`
	// for xstore_min_alloc_size_hdd/xstore_min_alloc_size_ssd
	MinAllocSize *int64 `json:"min_alloc_size,omitempty"`
	// size of omap partition
	OmapByte *int64 `json:"omap_byte,omitempty"`
	// cache partition id
	PartitionId *int64 `json:"partition_id,omitempty"`
	// read cache size in bytes
	ReadCacheSize *int64 `json:"read_cache_size,omitempty"`
	// osd role: \"data\" or \"index\", default is \"data\"
	Role *string `json:"role,omitempty"`
}

// NewOsdCreateReqOsd instantiates a new OsdCreateReqOsd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsdCreateReqOsd() *OsdCreateReqOsd {
	this := OsdCreateReqOsd{}
	return &this
}

// NewOsdCreateReqOsdWithDefaults instantiates a new OsdCreateReqOsd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsdCreateReqOsdWithDefaults() *OsdCreateReqOsd {
	this := OsdCreateReqOsd{}
	return &this
}

// GetBindPoolId returns the BindPoolId field value if set, zero value otherwise.
func (o *OsdCreateReqOsd) GetBindPoolId() int64 {
	if o == nil || IsNil(o.BindPoolId) {
		var ret int64
		return ret
	}
	return *o.BindPoolId
}

// GetBindPoolIdOk returns a tuple with the BindPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdCreateReqOsd) GetBindPoolIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BindPoolId) {
		return nil, false
	}
	return o.BindPoolId, true
}

// HasBindPoolId returns a boolean if a field has been set.
func (o *OsdCreateReqOsd) HasBindPoolId() bool {
	if o != nil && !IsNil(o.BindPoolId) {
		return true
	}

	return false
}

// SetBindPoolId gets a reference to the given int64 and assigns it to the BindPoolId field.
func (o *OsdCreateReqOsd) SetBindPoolId(v int64) {
	o.BindPoolId = &v
}

// GetDiskId returns the DiskId field value if set, zero value otherwise.
func (o *OsdCreateReqOsd) GetDiskId() int64 {
	if o == nil || IsNil(o.DiskId) {
		var ret int64
		return ret
	}
	return *o.DiskId
}

// GetDiskIdOk returns a tuple with the DiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdCreateReqOsd) GetDiskIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DiskId) {
		return nil, false
	}
	return o.DiskId, true
}

// HasDiskId returns a boolean if a field has been set.
func (o *OsdCreateReqOsd) HasDiskId() bool {
	if o != nil && !IsNil(o.DiskId) {
		return true
	}

	return false
}

// SetDiskId gets a reference to the given int64 and assigns it to the DiskId field.
func (o *OsdCreateReqOsd) SetDiskId(v int64) {
	o.DiskId = &v
}

// GetMinAllocSize returns the MinAllocSize field value if set, zero value otherwise.
func (o *OsdCreateReqOsd) GetMinAllocSize() int64 {
	if o == nil || IsNil(o.MinAllocSize) {
		var ret int64
		return ret
	}
	return *o.MinAllocSize
}

// GetMinAllocSizeOk returns a tuple with the MinAllocSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdCreateReqOsd) GetMinAllocSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MinAllocSize) {
		return nil, false
	}
	return o.MinAllocSize, true
}

// HasMinAllocSize returns a boolean if a field has been set.
func (o *OsdCreateReqOsd) HasMinAllocSize() bool {
	if o != nil && !IsNil(o.MinAllocSize) {
		return true
	}

	return false
}

// SetMinAllocSize gets a reference to the given int64 and assigns it to the MinAllocSize field.
func (o *OsdCreateReqOsd) SetMinAllocSize(v int64) {
	o.MinAllocSize = &v
}

// GetOmapByte returns the OmapByte field value if set, zero value otherwise.
func (o *OsdCreateReqOsd) GetOmapByte() int64 {
	if o == nil || IsNil(o.OmapByte) {
		var ret int64
		return ret
	}
	return *o.OmapByte
}

// GetOmapByteOk returns a tuple with the OmapByte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdCreateReqOsd) GetOmapByteOk() (*int64, bool) {
	if o == nil || IsNil(o.OmapByte) {
		return nil, false
	}
	return o.OmapByte, true
}

// HasOmapByte returns a boolean if a field has been set.
func (o *OsdCreateReqOsd) HasOmapByte() bool {
	if o != nil && !IsNil(o.OmapByte) {
		return true
	}

	return false
}

// SetOmapByte gets a reference to the given int64 and assigns it to the OmapByte field.
func (o *OsdCreateReqOsd) SetOmapByte(v int64) {
	o.OmapByte = &v
}

// GetPartitionId returns the PartitionId field value if set, zero value otherwise.
func (o *OsdCreateReqOsd) GetPartitionId() int64 {
	if o == nil || IsNil(o.PartitionId) {
		var ret int64
		return ret
	}
	return *o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdCreateReqOsd) GetPartitionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PartitionId) {
		return nil, false
	}
	return o.PartitionId, true
}

// HasPartitionId returns a boolean if a field has been set.
func (o *OsdCreateReqOsd) HasPartitionId() bool {
	if o != nil && !IsNil(o.PartitionId) {
		return true
	}

	return false
}

// SetPartitionId gets a reference to the given int64 and assigns it to the PartitionId field.
func (o *OsdCreateReqOsd) SetPartitionId(v int64) {
	o.PartitionId = &v
}

// GetReadCacheSize returns the ReadCacheSize field value if set, zero value otherwise.
func (o *OsdCreateReqOsd) GetReadCacheSize() int64 {
	if o == nil || IsNil(o.ReadCacheSize) {
		var ret int64
		return ret
	}
	return *o.ReadCacheSize
}

// GetReadCacheSizeOk returns a tuple with the ReadCacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdCreateReqOsd) GetReadCacheSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadCacheSize) {
		return nil, false
	}
	return o.ReadCacheSize, true
}

// HasReadCacheSize returns a boolean if a field has been set.
func (o *OsdCreateReqOsd) HasReadCacheSize() bool {
	if o != nil && !IsNil(o.ReadCacheSize) {
		return true
	}

	return false
}

// SetReadCacheSize gets a reference to the given int64 and assigns it to the ReadCacheSize field.
func (o *OsdCreateReqOsd) SetReadCacheSize(v int64) {
	o.ReadCacheSize = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OsdCreateReqOsd) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsdCreateReqOsd) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OsdCreateReqOsd) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OsdCreateReqOsd) SetRole(v string) {
	o.Role = &v
}

func (o OsdCreateReqOsd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsdCreateReqOsd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BindPoolId) {
		toSerialize["bind_pool_id"] = o.BindPoolId
	}
	if !IsNil(o.DiskId) {
		toSerialize["disk_id"] = o.DiskId
	}
	if !IsNil(o.MinAllocSize) {
		toSerialize["min_alloc_size"] = o.MinAllocSize
	}
	if !IsNil(o.OmapByte) {
		toSerialize["omap_byte"] = o.OmapByte
	}
	if !IsNil(o.PartitionId) {
		toSerialize["partition_id"] = o.PartitionId
	}
	if !IsNil(o.ReadCacheSize) {
		toSerialize["read_cache_size"] = o.ReadCacheSize
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableOsdCreateReqOsd struct {
	value *OsdCreateReqOsd
	isSet bool
}

func (v NullableOsdCreateReqOsd) Get() *OsdCreateReqOsd {
	return v.value
}

func (v *NullableOsdCreateReqOsd) Set(val *OsdCreateReqOsd) {
	v.value = val
	v.isSet = true
}

func (v NullableOsdCreateReqOsd) IsSet() bool {
	return v.isSet
}

func (v *NullableOsdCreateReqOsd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsdCreateReqOsd(val *OsdCreateReqOsd) *NullableOsdCreateReqOsd {
	return &NullableOsdCreateReqOsd{value: val, isSet: true}
}

func (v NullableOsdCreateReqOsd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsdCreateReqOsd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


